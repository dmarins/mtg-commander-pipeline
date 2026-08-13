package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

func cmdSearch(args []string) error {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	ci := fs.String("id", "", "identidade de cor permitida, ex.: WUR (subconjunto, como id<= no Scryfall)")
	typ := fs.String("type", "", "filtra por trecho da type line, ex.: Artifact")
	maxCMC := fs.Float64("cmc-max", -1, "CMC máximo")
	minCMC := fs.Float64("cmc-min", -1, "CMC mínimo")
	commander := fs.Bool("commander", true, "apenas cartas legais em Commander")
	limit := fs.Int("limit", 40, "máximo de resultados")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() == 0 {
		return fmt.Errorf("uso: mtgdb search <termos>   (sintaxe FTS5: AND, OR, NOT, \"frase exata\", prefixo*)")
	}

	db, err := openBuilt()
	if err != nil {
		return err
	}
	defer db.Close()

	query := strings.Join(fs.Args(), " ")

	sb := strings.Builder{}
	sb.WriteString(`SELECT c.oracle_id, c.name, c.mana_cost, c.cmc, c.type_line, c.color_identity
	                  FROM cards_fts f JOIN cards c ON c.oracle_id = f.oracle_id
	                 WHERE cards_fts MATCH ?`)
	params := []any{"oracle_text : " + query}

	if *commander {
		sb.WriteString(` AND c.legal_commander = 'legal'`)
	}
	if *typ != "" {
		sb.WriteString(` AND c.type_line LIKE ?`)
		params = append(params, "%"+*typ+"%")
	}
	if *maxCMC >= 0 {
		sb.WriteString(` AND c.cmc <= ?`)
		params = append(params, *maxCMC)
	}
	if *minCMC >= 0 {
		sb.WriteString(` AND c.cmc >= ?`)
		params = append(params, *minCMC)
	}
	if *ci != "" {
		// Identidade da carta contida na informada: nenhuma cor fora do conjunto.
		for _, color := range "WUBRG" {
			if !strings.ContainsRune(strings.ToUpper(*ci), color) {
				sb.WriteString(fmt.Sprintf(` AND c.color_identity NOT LIKE '%%%c%%'`, color))
			}
		}
	}
	sb.WriteString(` ORDER BY c.cmc, c.name LIMIT ?`)
	params = append(params, *limit)

	rows, err := db.SQL().Query(sb.String(), params...)
	if err != nil {
		return fmt.Errorf("busca inválida (sintaxe FTS5): %w", err)
	}
	defer rows.Close()

	n := 0
	for rows.Next() {
		var id, name, cost, tl, cid string
		var cmc float64
		if err := rows.Scan(&id, &name, &cost, &cmc, &tl, &cid); err != nil {
			return err
		}
		fmt.Printf("%-34s %-12s %-4.0f %-38s %s\n",
			truncate(name, 34), cost, cmc, truncate(tl, 38), orDash(cid))
		n++
	}
	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Printf("\n%d resultado(s)\n", n)
	if n == *limit {
		fmt.Println("(limite atingido — refine os filtros ou use -limit)")
	}
	return nil
}

func cmdTag(args []string) error {
	fs := flag.NewFlagSet("tag", flag.ExitOnError)
	ci := fs.String("id", "", "identidade de cor permitida, ex.: WUR")
	limit := fs.Int("limit", 60, "máximo de resultados")
	list := fs.Bool("list", false, "lista as tags que casam com o termo, em vez das cartas")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() == 0 {
		return fmt.Errorf("uso: mtgdb tag <slug>   (ex.: 'proliferate', 'ramp'; use -list para descobrir slugs)")
	}
	slug := fs.Arg(0)

	db, err := openBuilt()
	if err != nil {
		return err
	}
	defer db.Close()

	if *list {
		rows, err := db.SQL().Query(
			`SELECT t.slug, COALESCE(t.description,''), count(ct.oracle_id)
			   FROM tags t LEFT JOIN card_tags ct ON ct.tag = t.slug
			  WHERE t.slug LIKE ? GROUP BY t.slug ORDER BY count(ct.oracle_id) DESC LIMIT ?`,
			"%"+slug+"%", *limit)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var s, d string
			var n int
			if err := rows.Scan(&s, &d, &n); err != nil {
				return err
			}
			fmt.Printf("%-42s %5d  %s\n", s, n, truncate(d, 60))
		}
		return rows.Err()
	}

	q := strings.Builder{}
	q.WriteString(`SELECT c.name, c.mana_cost, c.cmc, c.type_line, c.color_identity
	                 FROM card_tags ct JOIN cards c ON c.oracle_id = ct.oracle_id
	                WHERE ct.tag = ? AND c.legal_commander = 'legal'`)
	params := []any{slug}
	if *ci != "" {
		for _, color := range "WUBRG" {
			if !strings.ContainsRune(strings.ToUpper(*ci), color) {
				q.WriteString(fmt.Sprintf(` AND c.color_identity NOT LIKE '%%%c%%'`, color))
			}
		}
	}
	q.WriteString(` ORDER BY c.cmc, c.name LIMIT ?`)
	params = append(params, *limit)

	rows, err := db.SQL().Query(q.String(), params...)
	if err != nil {
		return err
	}
	defer rows.Close()

	n := 0
	for rows.Next() {
		var name, cost, tl, cid string
		var cmc float64
		if err := rows.Scan(&name, &cost, &cmc, &tl, &cid); err != nil {
			return err
		}
		fmt.Printf("%-34s %-12s %-4.0f %-38s %s\n",
			truncate(name, 34), cost, cmc, truncate(tl, 38), orDash(cid))
		n++
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if n == 0 {
		fmt.Printf("nenhuma carta com a tag %q — use 'mtgdb tag -list %s' para descobrir o slug certo\n", slug, slug)
	} else {
		fmt.Printf("\n%d carta(s) com a tag %q\n", n, slug)
	}
	return nil
}

func cmdRulings(args []string) error {
	fs := flag.NewFlagSet("rulings", flag.ExitOnError)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() == 0 {
		return fmt.Errorf("uso: mtgdb rulings <nome da carta>")
	}

	db, err := openBuilt()
	if err != nil {
		return err
	}
	defer db.Close()

	m, err := db.Resolve(strings.Join(fs.Args(), " "))
	if err != nil {
		return err
	}
	if !m.Found {
		reportMissing([]store.Match{m})
		return nil
	}

	rs, err := db.Rulings(m.Card.OracleID)
	if err != nil {
		return err
	}
	fmt.Printf("%s — %d ruling(s)\n\n", m.Card.Name, len(rs))
	for _, r := range rs {
		fmt.Printf("[%s · %s]\n  %s\n\n", r.PublishedAt, r.Source, r.Comment)
	}
	if len(rs) == 0 {
		fmt.Println("Sem rulings oficiais. Para dúvida de regra, consulte as Comprehensive Rules")
		fmt.Println("ou uma fonte de juiz — não decida de memória.")
	}
	return nil
}
