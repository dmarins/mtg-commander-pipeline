package deck

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

// Fiche é a ficha de funções de uma carta (F1–F7 do checklist).
//
// ATENÇÃO: os campos abaixo são SINAIS extraídos por casamento de padrão sobre o
// texto oracle, não vereditos. Servem para dirigir a leitura — nunca para
// substituí-la. Quem decide se a carta entra ou sai é quem leu o texto inteiro.
type Fiche struct {
	Card store.Card
	Tags []string

	// F2 — corpo
	IsCreature   bool
	IsVehicle    bool
	IsSpacecraft bool
	Power        string
	PowerNum     int
	PowerDynamic bool // poder definido por characteristic (*)
	CrewN        int  // 0 = não tem crew
	HasStation   bool

	// F3 — tipo alimenta contagens
	Types []string

	// F4 — recebe
	HasCounters   bool // menciona contadores no próprio texto
	CounterKinds  []string
	CanBeAnthemed bool // criatura, logo alvo de lordes

	// F5 — facilita
	ReducesCost bool
	MakesMana   bool
	GrantsToOthers bool

	// F6 — curva
	CMC float64

	// F7 — atrito
	UsesTapSymbol bool // tem {T} no custo de ativação: disputa o tap com crew/station
	Sacrifices    bool
}

var (
	reCrew     = regexp.MustCompile(`(?i)\bcrew\s+(\d+)`)
	reCostLess = regexp.MustCompile(`(?i)cost(?:s)?\s+\{[^}]+\}\s+less`)

	// Tipos de contador: ou o padrão +N/+N e -N/-N, ou um dos nomes conhecidos.
	// Casar "a palavra antes de counter" pega ruído ("a counter", "many counters"),
	// então a lista fechada é mais confiável — e o que faltar aparece como
	// "contadores", que já basta como sinal.
	rePTCounter   = regexp.MustCompile(`([+-]\d+/[+-]\d+) counters?\b`)
	reNamedCntr   = regexp.MustCompile(`(?i)\b(charge|loyalty|hour|rev|verse|lore|oil|energy|poison|stun|shield|incubate|time|fade|quest|experience|ki|spore|storage|blood|page|bounty|defense|level|coin|net|omen|pressure|shred|ticket)\s+counters?\b`)

	// \bother\b evita casar com "another creature you control" (Station, crew).
	reOther     = regexp.MustCompile(`(?i)\bother\b`)
	reGrantVerb = regexp.MustCompile(`(?i)\b(get|gets|have|has|gain|gains)\b`)
)

// Build monta a ficha de uma carta.
func Build(c store.Card, tags []string) Fiche {
	text := c.AllText()
	types := c.AllTypes()
	lower := strings.ToLower(text)
	ltypes := strings.ToLower(types)

	f := Fiche{
		Card:         c,
		Tags:         tags,
		CMC:          c.CMC,
		Power:        c.Power,
		IsCreature:   strings.Contains(ltypes, "creature"),
		IsVehicle:    strings.Contains(ltypes, "vehicle"),
		IsSpacecraft: strings.Contains(ltypes, "spacecraft"),
		HasStation:   strings.Contains(lower, "station"),
	}

	// Poder: distingue número real de poder dinâmico (Master of Etherium, Brotherhood
	// Vertibird). Um "*" não é poder 0 — é poder que depende do board.
	if c.Power != "" {
		if n, err := strconv.Atoi(c.Power); err == nil {
			f.PowerNum = n
		} else {
			f.PowerDynamic = true
		}
	}

	if m := reCrew.FindStringSubmatch(text); m != nil {
		f.CrewN, _ = strconv.Atoi(m[1])
	}

	// F3 — tipos que alimentam contagens (affinity, metalcraft, improvise, lordes).
	for _, t := range []string{"Artifact", "Creature", "Enchantment", "Land",
		"Instant", "Sorcery", "Planeswalker", "Vehicle", "Spacecraft", "Battle"} {
		if strings.Contains(types, t) {
			f.Types = append(f.Types, t)
		}
	}

	// F4 — contadores citados no próprio texto.
	// "counter target spell" não é contador: exige a forma "<tipo> counter".
	if strings.Contains(lower, "counter") {
		seen := map[string]bool{}
		addKind := func(k string) {
			k = strings.ToLower(strings.TrimSpace(k))
			if k != "" && !seen[k] {
				seen[k] = true
				f.CounterKinds = append(f.CounterKinds, k)
			}
		}
		for _, m := range rePTCounter.FindAllStringSubmatch(text, -1) {
			addKind(m[1])
		}
		for _, m := range reNamedCntr.FindAllStringSubmatch(text, -1) {
			addKind(m[1])
		}
		// Menciona contador de forma genérica (proliferate, "counters on it").
		f.HasCounters = len(f.CounterKinds) > 0 ||
			strings.Contains(lower, "proliferate") ||
			strings.Contains(lower, "counters on") ||
			strings.Contains(lower, "counter on")
	}
	f.CanBeAnthemed = f.IsCreature

	// F5 — o que a carta dá a outras.
	f.ReducesCost = reCostLess.MatchString(text)
	f.MakesMana = strings.Contains(lower, "add {") || strings.Contains(lower, "add one mana") ||
		strings.Contains(lower, "add x mana") || strings.Contains(lower, "mana of any")
	f.GrantsToOthers = reOther.MatchString(text) && reGrantVerb.MatchString(text)

	// F7 — atrito. {T} num custo de ativação compete com crew e station pelo mesmo tap.
	f.UsesTapSymbol = strings.Contains(text, "{T}")
	f.Sacrifices = strings.Contains(lower, "sacrifice")

	return f
}

// TapConflict informa se a carta é um corpo que também quer usar o próprio tap —
// o atrito clássico em decks com crew/station/convoke/improvise.
func (f Fiche) TapConflict() bool { return f.IsCreature && f.UsesTapSymbol }

// BodyNote resume o eixo F2 em uma expressão curta.
func (f Fiche) BodyNote() string {
	var parts []string
	switch {
	case f.IsCreature && f.PowerDynamic:
		parts = append(parts, "corpo "+f.Power+"(dinâmico)")
	case f.IsCreature:
		parts = append(parts, "corpo "+f.Power)
	case f.IsVehicle:
		parts = append(parts, "veículo")
	case f.IsSpacecraft:
		parts = append(parts, "spacecraft")
	default:
		parts = append(parts, "—")
	}
	if f.CrewN > 0 {
		parts = append(parts, "crew "+strconv.Itoa(f.CrewN))
	}
	if f.HasStation {
		parts = append(parts, "station")
	}
	if f.TapConflict() {
		parts = append(parts, "usa {T}")
	}
	return strings.Join(parts, ", ")
}

// FacilitaNote resume o eixo F5.
func (f Fiche) FacilitaNote() string {
	var parts []string
	if f.ReducesCost {
		parts = append(parts, "reduz custo")
	}
	if f.MakesMana {
		parts = append(parts, "mana")
	}
	if f.GrantsToOthers {
		parts = append(parts, "buffa outros")
	}
	if len(parts) == 0 {
		return "—"
	}
	return strings.Join(parts, ", ")
}

// RecebeNote resume o eixo F4.
func (f Fiche) RecebeNote() string {
	var parts []string
	if len(f.CounterKinds) > 0 {
		parts = append(parts, strings.Join(f.CounterKinds, "/"))
	} else if f.HasCounters {
		parts = append(parts, "contadores")
	}
	if f.CanBeAnthemed {
		parts = append(parts, "anthems")
	}
	if len(parts) == 0 {
		return "—"
	}
	return strings.Join(parts, ", ")
}
