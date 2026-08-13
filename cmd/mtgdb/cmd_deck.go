package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dmarins/mtg-commander-pipeline/internal/deck"
	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

func cmdDeck(args []string) error {
	fs := flag.NewFlagSet("deck", flag.ExitOnError)
	full := fs.Bool("full", false, "imprime o texto oracle de cada carta")
	onlyMissing := fs.Bool("missing", false, "lista apenas as cartas não resolvidas")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return fmt.Errorf("uso: mtgdb deck <slug>   (diretórios em %s)", decksDir())
	}
	slug := fs.Arg(0)
	dir := filepath.Join(decksDir(), slug)
	if _, err := os.Stat(dir); err != nil {
		return fmt.Errorf("deck %q não encontrado em %s", slug, decksDir())
	}

	entries, srcFile, err := deck.Load(dir)
	if err != nil {
		return err
	}

	db, err := openBuilt()
	if err != nil {
		return err
	}
	defer db.Close()

	var fiches []deck.Fiche
	var sections []string
	var missing []store.Match
	for _, e := range entries {
		m, err := db.Resolve(e.Name)
		if err != nil {
			return err
		}
		if !m.Found {
			missing = append(missing, m)
			continue
		}
		tags, _ := db.Tags(m.Card.OracleID)
		fiches = append(fiches, deck.Build(m.Card, tags))
		sections = append(sections, e.Section)
	}

	if *onlyMissing {
		reportMissing(missing)
		return nil
	}

	rel, _ := filepath.Rel(root(), srcFile)
	fmt.Printf("deck %s — %d entradas lidas de %s, %d resolvidas\n\n", slug, len(entries), rel, len(fiches))

	printAggregates(fiches)

	fmt.Println("\nFICHA DE FUNÇÕES (F2–F7)")
	fmt.Printf("%-32s %-5s %-26s %-22s %-18s %s\n", "carta", "cmc", "F2 corpo", "F4 recebe", "F5 facilita", "F7 atrito")
	fmt.Println(strings.Repeat("─", 130))
	for _, f := range fiches {
		fmt.Printf("%-32s %-5.0f %-26s %-22s %-18s %s\n",
			truncate(f.Card.Name, 32), f.CMC,
			truncate(f.BodyNote(), 26), truncate(f.RecebeNote(), 22),
			truncate(f.FacilitaNote(), 18), truncate(atritoNote(f), 30))
	}

	if *full {
		fmt.Println("\nTEXTO ORACLE")
		for _, f := range fiches {
			fmt.Printf("\n%s  %s\n", f.Card.Name, f.Card.ManaCost)
			for _, line := range strings.Split(f.Card.AllText(), "\n") {
				if strings.TrimSpace(line) != "" {
					fmt.Println("  " + line)
				}
			}
		}
	}

	fmt.Println("\n⚠ Os eixos acima são sinais extraídos por padrão de texto, não vereditos.")
	fmt.Println("  Antes de cortar qualquer carta, leia o oracle dela (mtgdb oracle <nome>)")
	fmt.Println("  e siga o protocolo de references/card-evaluation-checklist.md §2.")

	reportMissing(missing)
	return nil
}

func printAggregates(fiches []deck.Fiche) {
	var creatures, vehicles, spacecraft, artifacts int
	var tapConflict, stationable int
	var powerGE3, dynamicPower int
	curve := map[int]int{}
	typeCount := map[string]int{}

	for _, f := range fiches {
		for _, t := range f.Types {
			typeCount[t]++
		}
		if strings.Contains(f.Card.AllTypes(), "Artifact") {
			artifacts++
		}
		switch {
		case f.IsCreature:
			creatures++
			// Corpo capaz de pagar crew/station: criatura de verdade.
			stationable++
			if f.PowerDynamic {
				dynamicPower++
			} else if f.PowerNum >= 3 {
				powerGE3++
			}
			if f.TapConflict() {
				tapConflict++
			}
		case f.IsVehicle:
			vehicles++
		case f.IsSpacecraft:
			spacecraft++
		}
		if !strings.Contains(f.Card.TypeLine, "Land") {
			b := int(f.CMC)
			if b > 6 {
				b = 7
			}
			curve[b]++
		}
	}

	fmt.Println("AGREGADOS")
	fmt.Printf("  criaturas de verdade      %d\n", creatures)
	fmt.Printf("  veículos                  %d   (precisam de crew — consomem corpos)\n", vehicles)
	fmt.Printf("  spacecraft                %d   (precisam de station — consomem corpos)\n", spacecraft)
	fmt.Printf("  artefatos (qualquer tipo) %d\n", artifacts)
	fmt.Printf("  corpos com poder ≥3       %d   (+%d com poder dinâmico)\n", powerGE3, dynamicPower)
	fmt.Printf("  corpos que usam o próprio {T}  %d   (atrito com crew/station)\n", tapConflict)

	var types []string
	for t := range typeCount {
		types = append(types, t)
	}
	sort.Strings(types)
	var tparts []string
	for _, t := range types {
		tparts = append(tparts, fmt.Sprintf("%s %d", t, typeCount[t]))
	}
	fmt.Printf("  tipos                     %s\n", strings.Join(tparts, " · "))

	fmt.Printf("  curva (não-terrenos)      ")
	for i := 0; i <= 7; i++ {
		label := fmt.Sprintf("%d", i)
		if i == 7 {
			label = "7+"
		}
		fmt.Printf("%s:%d ", label, curve[i])
	}
	fmt.Println()
}
