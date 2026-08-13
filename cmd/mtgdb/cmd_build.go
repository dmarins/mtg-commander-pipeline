package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

func openDB() (*store.DB, error) {
	db, err := store.Open(dbPath())
	if err != nil {
		return nil, err
	}
	return db, nil
}

// openBuilt abre o banco exigindo que já tenha passado por um build.
func openBuilt() (*store.DB, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	if !db.Built() {
		db.Close()
		return nil, fmt.Errorf("banco vazio — rode 'make db' (ou 'mtgdb build') primeiro")
	}
	// Um banco construído por uma versão anterior pode ter índices de nome
	// inconsistentes com a normalização atual — o sintoma seria uma carta
	// "não encontrada" sem motivo aparente. Melhor falhar explicitamente.
	if v := db.Meta("schema_version"); v != store.SchemaVersion {
		db.Close()
		return nil, fmt.Errorf(
			"banco na versão de schema %q, o binário espera %q — rode 'make refresh'",
			orNone(v), store.SchemaVersion)
	}
	return db, nil
}

func orNone(v string) string {
	if v == "" {
		return "desconhecida"
	}
	return v
}

func cmdBuild(args []string) error {
	fs := flag.NewFlagSet("build", flag.ExitOnError)
	quiet := fs.Bool("quiet", false, "não imprime o andamento")
	if err := fs.Parse(args); err != nil {
		return err
	}

	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	log := store.Progress(func(format string, a ...any) {
		if !*quiet {
			fmt.Printf(format+"\n", a...)
		}
	})

	start := time.Now()
	if !*quiet {
		fmt.Printf("banco: %s\n\n", dbPath())
	}

	if err := db.BuildAll(log); err != nil {
		return err
	}

	// collection e prices vêm dos TSV versionados — o banco é só espelho deles.
	nc, err := db.LoadCollection(dataDir())
	if err != nil {
		return fmt.Errorf("importando collection.tsv: %w", err)
	}
	np, err := db.LoadPrices(dataDir())
	if err != nil {
		return fmt.Errorf("importando prices.tsv: %w", err)
	}
	log("data/collection.tsv: %d cartas ✓", nc)
	log("data/prices.tsv: %d observações ✓", np)

	if fi, err := os.Stat(dbPath()); err == nil {
		log("\nconcluído em %s · banco com %.1f MB",
			time.Since(start).Round(time.Second), float64(fi.Size())/(1<<20))
	}
	return nil
}

func cmdStatus(args []string) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Printf("banco:      %s\n", dbPath())
	if fi, err := os.Stat(dbPath()); err == nil {
		fmt.Printf("tamanho:    %.1f MB\n", float64(fi.Size())/(1<<20))
	}
	fmt.Printf("cartas:     %d\n", db.CardCount())

	count := func(table string) int {
		var n int
		_ = db.SQL().QueryRow(`SELECT count(*) FROM ` + table).Scan(&n)
		return n
	}
	fmt.Printf("tags:       %d (%d vínculos)\n", count("tags"), count("card_tags"))
	fmt.Printf("rulings:    %d\n", count("rulings"))
	fmt.Printf("coleção:    %d cartas\n", count("collection"))
	fmt.Printf("preços:     %d observações\n", count("prices"))

	fmt.Println()
	for _, k := range []string{"cards_updated_at", "tags_updated_at", "rulings_updated_at", "cards_built_at"} {
		if v := db.Meta(k); v != "" {
			fmt.Printf("%-20s %s\n", k+":", v)
		}
	}
	if !db.Built() {
		fmt.Println("\nbanco vazio — rode 'make db'")
	}
	return nil
}
