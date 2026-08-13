package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dmarins/mtg-commander-pipeline/internal/deck"
	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

func cmdOracle(args []string) error {
	fs := flag.NewFlagSet("oracle", flag.ExitOnError)
	asJSON := fs.Bool("json", false, "saída em JSON")
	short := fs.Bool("short", false, "uma linha por carta, sem o texto oracle")
	withRulings := fs.Bool("rulings", false, "inclui as rulings oficiais")
	if err := fs.Parse(args); err != nil {
		return err
	}
	names := fs.Args()
	if len(names) == 0 {
		return fmt.Errorf("informe ao menos um nome de carta")
	}

	db, err := openBuilt()
	if err != nil {
		return err
	}
	defer db.Close()

	var found []store.Match
	var missing []store.Match
	for _, n := range names {
		m, err := db.Resolve(n)
		if err != nil {
			return err
		}
		if m.Found {
			found = append(found, m)
		} else {
			missing = append(missing, m)
		}
	}

	if *asJSON {
		type out struct {
			Query   string      `json:"query"`
			How     string      `json:"how,omitempty"`
			Found   bool        `json:"found"`
			Card    *store.Card `json:"card,omitempty"`
			Tags    []string    `json:"tags,omitempty"`
			Ambigs  []string    `json:"ambiguous,omitempty"`
		}
		var res []out
		for _, m := range append(found, missing...) {
			o := out{Query: m.Query, How: m.How, Found: m.Found, Ambigs: m.Ambigs}
			if m.Found {
				c := m.Card
				o.Card = &c
				o.Tags, _ = db.Tags(c.OracleID)
			}
			res = append(res, o)
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(res)
	}

	for i, m := range found {
		if *short {
			printShort(db, m)
			continue
		}
		if i > 0 {
			fmt.Println()
		}
		printFiche(db, m, *withRulings)
	}
	reportMissing(missing)
	return nil
}

func printShort(db *store.DB, m store.Match) {
	c := m.Card
	pt := ""
	if c.Power != "" {
		pt = " " + c.Power + "/" + c.Toughness
	}
	fmt.Printf("%-34s %-12s %s%s\n", truncate(c.Name, 34), c.ManaCost, truncate(c.TypeLine, 40), pt)
}

func printFiche(db *store.DB, m store.Match, withRulings bool) {
	c := m.Card
	tags, _ := db.Tags(c.OracleID)
	f := deck.Build(c, tags)

	head := c.Name
	if c.ManaCost != "" {
		head += "  " + c.ManaCost
	}
	if c.Power != "" {
		head += "  " + c.Power + "/" + c.Toughness
	}
	fmt.Println(head)
	fmt.Println(strings.Repeat("─", len([]rune(head))))
	fmt.Printf("%s · CMC %.0f · identidade %s\n", c.TypeLine, c.CMC, orDash(c.ColorIdentity))
	if m.How != "exact" {
		fmt.Printf("  (resolvido por %s a partir de %q — confira se é a carta certa)\n", m.How, m.Query)
	}

	fmt.Println()
	for _, line := range strings.Split(c.AllText(), "\n") {
		if strings.TrimSpace(line) != "" {
			fmt.Println("  " + line)
		}
	}

	fmt.Println()
	fmt.Printf("  F2 corpo      %s\n", f.BodyNote())
	fmt.Printf("  F3 tipos      %s\n", orDash(strings.Join(f.Types, ", ")))
	fmt.Printf("  F4 recebe     %s\n", f.RecebeNote())
	fmt.Printf("  F5 facilita   %s\n", f.FacilitaNote())
	fmt.Printf("  F6 curva      CMC %.0f\n", f.CMC)
	fmt.Printf("  F7 atrito     %s\n", atritoNote(f))

	if len(tags) > 0 {
		fmt.Printf("  tags          %s\n", truncate(strings.Join(tags, ", "), 100))
	}
	if qty, ok, _ := db.InCollection(c.Name); ok {
		fmt.Printf("  coleção       tem %d\n", qty)
	}
	if p, ok, _ := db.LatestPrice(c.Name, "ligamagic", "menor"); ok {
		fmt.Printf("  preço         R$ %.2f (LigaMagic menor, %s — %d dias)\n", p.Value, p.CapturedAt, p.AgeDays())
	}

	if withRulings {
		rs, _ := db.Rulings(c.OracleID)
		if len(rs) > 0 {
			fmt.Printf("\n  rulings (%d):\n", len(rs))
			for _, r := range rs {
				fmt.Printf("   · [%s] %s\n", r.PublishedAt, r.Comment)
			}
		}
	}

	fmt.Println("\n  ⚠ F2–F7 acima são sinais extraídos por padrão de texto, não vereditos.")
	fmt.Println("    O texto oracle está impresso acima justamente para ser lido.")
}

func atritoNote(f deck.Fiche) string {
	var parts []string
	if f.TapConflict() {
		parts = append(parts, "corpo que também quer o próprio {T}")
	} else if f.UsesTapSymbol {
		parts = append(parts, "usa {T}")
	}
	if f.Sacrifices {
		parts = append(parts, "sacrifica")
	}
	if len(parts) == 0 {
		return "—"
	}
	return strings.Join(parts, "; ")
}

func orDash(s string) string {
	if strings.TrimSpace(s) == "" {
		return "—"
	}
	return s
}

func reportMissing(missing []store.Match) {
	if len(missing) == 0 {
		return
	}
	fmt.Fprintf(os.Stderr, "\n%d não resolvido(s) — nomes de carta são em inglês:\n", len(missing))
	for _, m := range missing {
		if len(m.Ambigs) > 0 {
			fmt.Fprintf(os.Stderr, "  %-30s ambíguo: %s\n", m.Query, truncate(strings.Join(m.Ambigs, " | "), 90))
		} else {
			fmt.Fprintf(os.Stderr, "  %s\n", m.Query)
		}
	}
}
