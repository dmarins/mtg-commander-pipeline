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
