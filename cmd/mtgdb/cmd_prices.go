package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

func cmdPrices(args []string) error {
	fs := flag.NewFlagSet("prices", flag.ExitOnError)
	stale := fs.Int("stale", -1, "lista apenas cotações com mais de N dias")
	history := fs.Bool("history", false, "mostra todas as observações, não só a mais recente")
	volatile := fs.Bool("volatile", false, "ranqueia as cartas por oscilação observada")
	add := fs.Bool("add", false, "acrescenta uma observação: -add <carta> <valor> [-kind menor] [-date AAAA-MM-DD]")
	source := fs.String("source", "ligamagic", "fonte da cotação")
	kind := fs.String("kind", "menor", "tipo: menor|medio|maior|foil_menor…")
	date := fs.String("date", time.Now().Format("2006-01-02"), "data da captura (AAAA-MM-DD)")
	note := fs.String("note", "", "observação livre")
	if err := fs.Parse(args); err != nil {
		return err
	}

	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// data/prices.tsv é a fonte; o banco é espelho. Recarrega para refletir
	// edições manuais feitas no arquivo desde o último build.
	if _, err := db.LoadPrices(dataDir()); err != nil {
		return fmt.Errorf("lendo data/prices.tsv: %w", err)
	}

	if *add {
		if fs.NArg() < 2 {
			return fmt.Errorf("uso: mtgdb prices -add \"Nome da Carta\" <valor>")
		}
		name := strings.Join(fs.Args()[:fs.NArg()-1], " ")
		raw := fs.Arg(fs.NArg() - 1)
		val, err := strconv.ParseFloat(strings.ReplaceAll(raw, ",", "."), 64)
		if err != nil {
			return fmt.Errorf("valor %q não é numérico", raw)
		}
		// Confere o nome oficial quando o banco está construído, mas não bloqueia:
		// registrar a cotação é mais importante que a grafia.
		if db.Built() {
			if m, err := db.Resolve(name); err == nil && m.Found {
				name = m.Card.Name
			} else {
				fmt.Printf("aviso: %q não resolveu para uma carta conhecida — gravando assim mesmo\n", name)
			}
		}
		obs := store.PriceObs{
			Name: name, Source: *source, Kind: *kind, Value: val,
			Currency: "BRL", CapturedAt: *date, Note: *note,
		}
		if err := db.AppendPrice(dataDir(), obs); err != nil {
			return err
		}
		fmt.Printf("registrado: %s · %s %s · R$ %.2f · %s\n", name, *source, *kind, val, *date)
		fmt.Println("(acrescentado a data/prices.tsv — observações nunca são sobrescritas)")
		return nil
	}

	if *volatile {
		return printVolatility(db)
	}

	names := fs.Args()
	if len(names) == 0 && *stale < 0 {
		return fmt.Errorf("informe nomes de carta, ou use -stale N, -volatile ou -add")
	}

	if *stale >= 0 && len(names) == 0 {
		return printStale(db, *stale)
	}

	for _, n := range names {
		obs, err := db.PriceHistory(n)
		if err != nil {
			return err
		}
		if len(obs) == 0 {
			fmt.Printf("%-34s sem cotação registrada\n", truncate(n, 34))
			continue
		}
		if *history {
			fmt.Printf("%s\n", obs[0].Name)
			for _, o := range obs {
				fmt.Printf("  %s  %-10s %-12s R$ %8.2f  (%d dias)\n",
					o.CapturedAt, o.Source, o.Kind, o.Value, o.AgeDays())
			}
			fmt.Println()
			continue
		}
		o := obs[0]
		fmt.Printf("%-34s R$ %8.2f  %-10s %-8s %s (%d dias)\n",
			truncate(o.Name, 34), o.Value, o.Source, o.Kind, o.CapturedAt, o.AgeDays())
	}
	return nil
}

func printStale(db *store.DB, days int) error {
	rows, err := db.SQL().Query(`
		SELECT name, source, kind, value, captured_at
		  FROM prices p
		 WHERE captured_at = (SELECT max(captured_at) FROM prices q
		                       WHERE q.name = p.name AND q.source = p.source AND q.kind = p.kind)
		 ORDER BY captured_at, value DESC`)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("Cotações com mais de %d dias:\n\n", days)
	n := 0
	for rows.Next() {
		var name, src, kind, at string
		var val float64
		if err := rows.Scan(&name, &src, &kind, &val, &at); err != nil {
			return err
		}
		p := store.PriceObs{CapturedAt: at}
		age := p.AgeDays()
		if age < days {
			continue
		}
		fmt.Printf("%-34s R$ %8.2f  %-10s %-8s %s (%d dias)\n",
			truncate(name, 34), val, src, kind, at, age)
		n++
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if n == 0 {
		fmt.Println("(nenhuma)")
		return nil
	}
	fmt.Printf("\n%d cotação(ões) a reconferir.\n", n)
	fmt.Println("A captura é manual: a página da LigaMagic monta o preço por JS, então é")
	fmt.Println("navegador carta a carta. Registre o resultado com 'mtgdb prices -add'.")
	return nil
}

// printVolatility mostra quais cartas de fato oscilam, para que a reconferência
// antes de um torneio mire nas poucas que podem mover o total do deck.
func printVolatility(db *store.DB) error {
	rows, err := db.SQL().Query(`
		SELECT name, source, kind, count(*) n, min(value) lo, max(value) hi,
		       min(captured_at) first_at, max(captured_at) last_at
		  FROM prices
		 GROUP BY name, source, kind
		HAVING n > 1
		 ORDER BY (hi - lo) DESC`)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Print("Oscilação observada (só cartas com 2+ capturas):\n\n")
	fmt.Printf("%-34s %6s %10s %10s %8s  %s\n", "carta", "obs", "mín", "máx", "var", "período")
	fmt.Println(strings.Repeat("─", 92))
	n := 0
	for rows.Next() {
		var name, src, kind, firstAt, lastAt string
		var cnt int
		var lo, hi float64
		if err := rows.Scan(&name, &src, &kind, &cnt, &lo, &hi, &firstAt, &lastAt); err != nil {
			return err
		}
		var pct float64
		if lo > 0 {
			pct = (hi - lo) / lo * 100
		}
		fmt.Printf("%-34s %6d %10.2f %10.2f %7.0f%%  %s → %s\n",
			truncate(name, 34), cnt, lo, hi, pct, firstAt, lastAt)
		n++
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if n == 0 {
		fmt.Println("(ainda não há carta com duas capturas — a série começa a valer na 2ª rodada de precificação)")
	}
	return nil
}

func cmdCollection(args []string) error {
	fs := flag.NewFlagSet("collection", flag.ExitOnError)
	list := fs.Bool("list", false, "lista a coleção inteira")
	add := fs.Bool("add", false, "acrescenta as cartas informadas à coleção")
	note := fs.String("note", "", "observação para as cartas adicionadas")
	if err := fs.Parse(args); err != nil {
		return err
	}

	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.LoadCollection(dataDir()); err != nil {
		return err
	}

	if *add {
		if fs.NArg() == 0 {
			return fmt.Errorf("uso: mtgdb collection -add \"Nome da Carta\" ...")
		}
		names := fs.Args()
		if db.Built() {
			for i, n := range names {
				if m, err := db.Resolve(n); err == nil && m.Found {
					names[i] = m.Card.Name
				} else {
					fmt.Printf("aviso: %q não resolveu — gravando assim mesmo\n", n)
				}
			}
		}
		n, err := db.AddToCollection(dataDir(), names, *note, time.Now().Format("2006-01-02"))
		if err != nil {
			return err
		}
		fmt.Printf("%d carta(s) registradas em data/collection.tsv\n", n)
		return nil
	}

	if *list {
		rows, err := db.SQL().Query(`SELECT name, qty, COALESCE(note,'') FROM collection ORDER BY name COLLATE NOCASE`)
		if err != nil {
			return err
		}
		defer rows.Close()
		n := 0
		for rows.Next() {
			var name, note string
			var qty int
			if err := rows.Scan(&name, &qty, &note); err != nil {
				return err
			}
			fmt.Printf("%3d× %-40s %s\n", qty, truncate(name, 40), note)
			n++
		}
		fmt.Printf("\n%d cartas na coleção\n", n)
		return rows.Err()
	}

	if fs.NArg() == 0 {
		return fmt.Errorf("informe nomes de carta, ou use -list / -add")
	}
	for _, n := range fs.Args() {
		name := n
		if db.Built() {
			if m, err := db.Resolve(n); err == nil && m.Found {
				name = m.Card.Name
			}
		}
		qty, ok, err := db.InCollection(name)
		if err != nil {
			return err
		}
		if ok {
			fmt.Printf("%-40s ✓ tem %d\n", truncate(name, 40), qty)
		} else {
			fmt.Printf("%-40s — não está na coleção\n", truncate(name, 40))
		}
	}
	return nil
}
