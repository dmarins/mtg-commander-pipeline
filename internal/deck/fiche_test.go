package deck

import (
	"strings"
	"testing"

	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

// Cartas reais, com o oracle text exato — as heurísticas existem para dirigir a
// leitura de casos assim, então é contra eles que precisam ser verificadas.
var (
	etheriumSculptor = store.Card{
		Name: "Etherium Sculptor", ManaCost: "{1}{U}", CMC: 2,
		TypeLine: "Artifact Creature — Vedalken Artificer", Power: "1", Toughness: "2",
		OracleText: "Artifact spells you cast cost {1} less to cast.",
	}
	masterOfEtherium = store.Card{
		Name: "Master of Etherium", ManaCost: "{2}{U}", CMC: 3,
		TypeLine: "Artifact Creature — Vedalken Wizard", Power: "*", Toughness: "*",
		OracleText: "Master of Etherium's power and toughness are each equal to the number of artifacts you control.\nOther artifact creatures you control get +1/+1.",
	}
	uthros = store.Card{
		Name: "Uthros Research Craft", ManaCost: "{2}{U}", CMC: 3,
		TypeLine: "Artifact — Spacecraft", Power: "0", Toughness: "8",
		OracleText: "Station (Tap another creature you control: Put charge counters equal to its power on this Spacecraft. Station only as a sorcery. It's an artifact creature at 12+.)\n3+ | Whenever you cast an artifact spell, draw a card. Put a charge counter on this Spacecraft.\n12+ | Flying\nThis Spacecraft gets +1/+0 for each artifact you control.",
	}
	vertibird = store.Card{
		Name: "Brotherhood Vertibird", ManaCost: "{3}", CMC: 3,
		TypeLine: "Artifact — Vehicle", Power: "*", Toughness: "4",
		OracleText: "Flying\nBrotherhood Vertibird's power is equal to the number of artifacts you control.\nCrew 2 (Tap any number of creatures you control with total power 2 or greater: This Vehicle becomes an artifact creature until end of turn.)",
	}
	coretapper = store.Card{
		Name: "Coretapper", ManaCost: "{2}", CMC: 2,
		TypeLine: "Artifact Creature — Myr", Power: "1", Toughness: "1",
		OracleText: "{T}: Put a charge counter on target artifact.\nSacrifice this creature: Put two charge counters on target artifact.",
	}
	accessDenied = store.Card{
		Name: "Access Denied", ManaCost: "{3}{U}{U}", CMC: 5,
		TypeLine:   "Instant",
		OracleText: "Counter target spell. Create X 1/1 colorless Thopter artifact creature tokens with flying, where X is that spell's mana value.",
	}
)

// O caso que motivou o checklist: a carta foi julgada "criatura 1/2 fraca",
// ignorando que ela reduz custo, é um corpo tapável e conta como artefato.
func TestFicheEtheriumSculptorMostraTodasAsFuncoes(t *testing.T) {
	f := Build(etheriumSculptor, nil)

	if !f.IsCreature || f.PowerNum != 1 {
		t.Errorf("F2: esperava criatura de poder 1, veio %+v", f.BodyNote())
	}
	if !f.ReducesCost {
		t.Error("F5: redução de custo não foi detectada — é a função que a análise parcial perdeu")
	}
	if !contains(f.Types, "Artifact") || !contains(f.Types, "Creature") {
		t.Errorf("F3: esperava Artifact e Creature, veio %v", f.Types)
	}
	if !f.CanBeAnthemed {
		t.Error("F4: criatura-artefato recebe anthems (Chief of the Foundry, Master of Etherium)")
	}
	if f.TapConflict() {
		t.Error("F7: ela não tem habilidade de {T}, então não disputa o tap com station/crew")
	}
}

// Poder "*" não é poder 0. Confundir os dois subestima o maior estacionador do deck.
func TestFichePoderDinamico(t *testing.T) {
	f := Build(masterOfEtherium, nil)
	if !f.PowerDynamic {
		t.Error("Master of Etherium tem poder dinâmico (*)")
	}
	if f.PowerNum != 0 {
		t.Errorf("poder dinâmico não deve virar número, veio %d", f.PowerNum)
	}
	if !strings.Contains(f.BodyNote(), "dinâmico") {
		t.Errorf("BodyNote deveria sinalizar poder dinâmico, veio %q", f.BodyNote())
	}
	if !f.GrantsToOthers {
		t.Error("F5: 'Other artifact creatures you control get +1/+1' é buff a outros")
	}
}

// "Tap another creature you control" contém "another" — não pode ser lido como
// "other ... get", senão toda carta com Station vira "buffa outros".
func TestFicheAnotherNaoContaComoOutros(t *testing.T) {
	f := Build(uthros, nil)
	if f.GrantsToOthers {
		t.Error("F5: 'another creature' não é concessão de bônus a outras cartas")
	}
	if !f.HasStation {
		t.Error("F2: Station não foi detectado")
	}
	if !f.IsSpacecraft {
		t.Error("F2: type line diz Spacecraft")
	}
	if f.IsCreature {
		t.Error("Spacecraft em repouso NÃO é criatura — contá-lo como corpo infla a conta de quem pode pagar station/crew")
	}
	if !contains(f.CounterKinds, "charge") {
		t.Errorf("F4: esperava charge counters, veio %v", f.CounterKinds)
	}
}

func TestFicheVeiculoNaoEhCriatura(t *testing.T) {
	f := Build(vertibird, nil)
	if f.IsCreature {
		t.Error("Veículo sem crew não é criatura")
	}
	if !f.IsVehicle || f.CrewN != 2 {
		t.Errorf("esperava veículo com crew 2, veio %q", f.BodyNote())
	}
}

// Corpo que tem a própria habilidade de {T} disputa o tap com station e crew.
func TestFicheAtritoDeTap(t *testing.T) {
	f := Build(coretapper, nil)
	if !f.TapConflict() {
		t.Error("Coretapper é criatura com {T}: deve acusar atrito de tap")
	}
	if !f.Sacrifices {
		t.Error("F7: a carta se sacrifica")
	}

	// Uma carta sem corpo que usa {T} não gera esse atrito específico.
	if Build(uthros, nil).TapConflict() {
		t.Error("Spacecraft não é corpo, logo não disputa tap com crew/station")
	}
}

// "Counter target spell" não é um contador — é anular mágica.
func TestFicheCounterspellNaoEhContador(t *testing.T) {
	f := Build(accessDenied, nil)
	if len(f.CounterKinds) != 0 {
		t.Errorf("counterspell não tem tipo de contador, veio %v", f.CounterKinds)
	}
}

func contains(hay []string, needle string) bool {
	for _, h := range hay {
		if h == needle {
			return true
		}
	}
	return false
}
