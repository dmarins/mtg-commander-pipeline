package deck

import "testing"

// A exportação MTGO carimba set e número do colecionador no nome. Antes disso
// ser tratado, uma lista.txt inteira caía como "não resolvido" — o nome ia para
// o banco como "Bombard (EOE) 129".
func TestCleanNameStripsMTGOTrailer(t *testing.T) {
	cases := map[string]string{
		"Bombard (EOE) 129":                        "Bombard",
		"Dragon Throne of Tarkir (PKTK) 219 *F*":   "Dragon Throne of Tarkir",
		"Tori D'Avenant, Fury Rider (DMU) 223":     "Tori D'Avenant, Fury Rider",
		"Dire Flail // Dire Blunderbuss (LCI) 145": "Dire Flail",
		"Sol Ring":                                 "Sol Ring",
		// Nome real cujo parêntese é texto e encerra a linha: o sufixo de set não
		// pode engoli-lo. (reTrailer ainda o remove — comportamento pré-existente.)
		"Mercadia's Downfall (MMQ) 205": "Mercadia's Downfall",
	}
	for in, want := range cases {
		if got := cleanName(in); got != want {
			t.Errorf("cleanName(%q) = %q, quero %q", in, got, want)
		}
	}
}
