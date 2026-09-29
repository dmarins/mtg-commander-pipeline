package store

import (
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Sol Ring", "sol ring"},
		{"Loran's Escape", "lorans escape"},
		{"Alibou, Ancient Witness", "alibou ancient witness"},
		{"Kilo, Apogee Mind", "kilo apogee mind"},
		// Acentos aparecem em nomes oficiais (Márton, Jötun) e em texto digitado.
		{"Márton Stromgald", "marton stromgald"},
		{"Jötun Grunt", "jotun grunt"},
		{"  Extra   Espaços  ", "extra espacos"},
		{"Thousand Moons Smithy", "thousand moons smithy"},
		{"Ratchet Bomb", "ratchet bomb"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Normalize(c.in); got != c.want {
			t.Errorf("Normalize(%q) = %q, quer %q", c.in, got, c.want)
		}
	}
}

// Grafias diferentes da mesma carta têm de colapsar na mesma chave — é isso que
// permite casar o que o usuário digita com o nome oficial do Scryfall.
func TestNormalizeColapsaVariacoes(t *testing.T) {
	groups := [][]string{
		{"Loran's Escape", "Lorans Escape", "loran's escape", "LORAN'S ESCAPE"},
		{"Alibou, Ancient Witness", "Alibou Ancient Witness", "alibou,ancient witness"},
		{"Sai, Master Thopterist", "Sai Master Thopterist"},
	}
	for _, g := range groups {
		first := Normalize(g[0])
		for _, v := range g[1:] {
			if got := Normalize(v); got != first {
				t.Errorf("Normalize(%q) = %q, quer %q (igual a %q)", v, got, first, g[0])
			}
		}
	}
}

func TestFTSQuote(t *testing.T) {
	if got := ftsQuote(`sol ring`); got != `"sol ring"` {
		t.Errorf("ftsQuote = %q", got)
	}
	// Aspas internas precisam ser duplicadas, senão a query FTS5 quebra.
	if got := ftsQuote(`a"b`); got != `"a""b"` {
		t.Errorf("ftsQuote com aspas = %q", got)
	}
}

func TestPriceObsAgeDays(t *testing.T) {
	invalid := PriceObs{CapturedAt: "não é data"}
	if got := invalid.AgeDays(); got != -1 {
		t.Errorf("data inválida deveria dar -1, deu %d", got)
	}
	old := PriceObs{CapturedAt: "2020-01-01"}
	if got := old.AgeDays(); got < 365 {
		t.Errorf("cotação de 2020 deveria ter centenas de dias, deu %d", got)
	}
}

func TestCardAllTextIncluiFaces(t *testing.T) {
	// MDFCs trazem o texto nas faces; perder isso faz uma carta parecer vazia.
	c := Card{
		Name:       "Thousand Moons Smithy",
		OracleText: "",
		Faces: []Face{
			{Name: "Thousand Moons Smithy", OracleText: "create a Gnome Soldier"},
			{Name: "Barracks of the Thousand", OracleText: "tap five untapped artifacts"},
		},
	}
	all := c.AllText()
	for _, want := range []string{"Gnome Soldier", "five untapped artifacts"} {
		if !strings.Contains(all, want) {
			t.Errorf("AllText não trouxe %q: %q", want, all)
		}
	}
}

// seedCard grava uma carta mínima com os aliases que o build geraria: o nome
// completo e cada face com nome diferente dele.
func seedCard(t *testing.T, d *DB, id, name, layout string, faces ...string) {
	t.Helper()
	if _, err := d.sql.Exec(`INSERT INTO cards (oracle_id,name,layout) VALUES (?,?,?)`,
		id, name, layout); err != nil {
		t.Fatal(err)
	}
	if _, err := d.sql.Exec(`INSERT OR IGNORE INTO card_names (norm,oracle_id,is_face) VALUES (?,?,0)`,
		Normalize(name), id); err != nil {
		t.Fatal(err)
	}
	for i, f := range faces {
		if _, err := d.sql.Exec(`INSERT INTO card_faces (oracle_id,idx,name) VALUES (?,?,?)`,
			id, i, f); err != nil {
			t.Fatal(err)
		}
		if !strings.EqualFold(f, name) {
			if _, err := d.sql.Exec(`INSERT OR IGNORE INTO card_names (norm,oracle_id,is_face) VALUES (?,?,1)`,
				Normalize(f), id); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func openTestDB(t *testing.T) *DB {
	t.Helper()
	d, err := Open(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { d.Close() })
	return d
}

// O bulk oracle_cards traz entradas que não são cartas de jogo — cartas de arte
// ("Goblin Glasswright // Goblin Glasswright", layout art_series), tokens,
// emblemas. Elas compartilham o nome da face com a carta real, e a lista do
// usuário chega só com a primeira face ("Goblin Glasswright // Craft with Pride"
// vira "Goblin Glasswright"). Sem desempate, a resolução ficava ambígua e a
// coleção gravava um nome fantasma.
func TestResolvePrefereCartaDeJogo(t *testing.T) {
	d := openTestDB(t)
	seedCard(t, d, "real", "Goblin Glasswright // Craft with Pride", "prepare",
		"Goblin Glasswright", "Craft with Pride")
	seedCard(t, d, "arte", "Goblin Glasswright // Goblin Glasswright", "art_series",
		"Goblin Glasswright", "Goblin Glasswright")

	m, err := d.Resolve("Goblin Glasswright")
	if err != nil {
		t.Fatal(err)
	}
	if !m.Found || m.Card.OracleID != "real" {
		t.Fatalf("quer a carta de jogo, veio found=%v card=%q ambígua=%v", m.Found, m.Card.Name, m.Ambigs)
	}
	if m.How != "face" {
		t.Errorf("How = %q, quer face", m.How)
	}
}

// Quando duas cartas de jogo disputam o nome, o desempate não pode escolher
// nenhuma: a ambiguidade é real e o usuário precisa ver os candidatos.
func TestResolveAmbiguidadeEntreCartasDeJogo(t *testing.T) {
	d := openTestDB(t)
	seedCard(t, d, "a", "Fire // Ice", "split", "Fire", "Ice")
	seedCard(t, d, "b", "Fire // Brimstone", "split", "Fire", "Brimstone")
	seedCard(t, d, "arte", "Fire // Fire", "art_series", "Fire", "Fire")

	m, err := d.Resolve("Fire")
	if err != nil {
		t.Fatal(err)
	}
	if m.Found {
		t.Fatalf("não devia resolver, resolveu %q", m.Card.Name)
	}
	if len(m.Ambigs) != 2 {
		t.Errorf("candidatos = %v, quer só as 2 cartas de jogo", m.Ambigs)
	}
}

// Uma entrada fora de jogo que é a única dona do nome continua resolvendo:
// o desempate só age quando há disputa.
func TestResolveEntradaUnicaForaDeJogo(t *testing.T) {
	d := openTestDB(t)
	seedCard(t, d, "tok", "Marit Lage", "token")

	m, err := d.Resolve("marit lage")
	if err != nil {
		t.Fatal(err)
	}
	if !m.Found || m.Card.OracleID != "tok" {
		t.Fatalf("quer o token, veio found=%v card=%q", m.Found, m.Card.Name)
	}
}
