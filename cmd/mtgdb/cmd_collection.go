package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dmarins/mtg-commander-pipeline/internal/deck"
	"github.com/dmarins/mtg-commander-pipeline/internal/store"
)

func cmdCollection(args []string) error {
	fs := flag.NewFlagSet("collection", flag.ExitOnError)
	list := fs.Bool("list", false, "lista a coleção inteira")
	add := fs.Bool("add", false, "acrescenta as cartas informadas à coleção")
	sync := fs.Bool("sync", false, "substitui a coleção pela lista informada (o que não estiver nela sai)")
	file := fs.String("file", "", "arquivo com a lista para -sync (\"-\" lê da entrada padrão)")
	dry := fs.Bool("dry-run", false, "com -sync, mostra o que mudaria sem escrever")
	force := fs.Bool("force", false, "com -sync, aceita esvaziar a coleção")
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

	if *sync {
		return syncCollection(db, *file, fs.Args(), *note, *dry, *force)
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

// syncCollection substitui data/collection.tsv pela lista recebida.
//
// A lista é lida com o mesmo parser das decklists: o usuário exporta a caixa no
// formato MTGO ("1 Bombard (EOE) 129") e não deve precisar aprender um segundo
// formato só para atualizar a coleção.
func syncCollection(db *store.DB, file string, args []string, note string, dry, force bool) error {
	want, src, err := collectionInput(file, args)
	if err != nil {
		return err
	}
	if len(want) == 0 && !force {
		return fmt.Errorf("a lista de %s não tem nenhuma carta — nada foi escrito.\n"+
			"Uma lista vazia apagaria a coleção inteira; se for isso mesmo, repita com -force", src)
	}

	// Regra 8: o TSV guarda o nome oficial em inglês. Resolver aqui evita que a
	// mesma carta entre duas vezes por grafias diferentes entre exportações.
	entries := make([]store.CollectionEntry, 0, len(want))
	var unresolved []string
	for _, e := range want {
		name := e.Name
		if db.Built() {
			if m, err := db.Resolve(name); err == nil && m.Found {
				name = m.Card.Name
			} else {
				unresolved = append(unresolved, e.Name)
			}
		}
		entries = append(entries, store.CollectionEntry{Name: name, Qty: e.Qty})
	}

	diff, err := store.SyncCollection(dataDir(), entries, note, time.Now().Format("2006-01-02"), dry)
	if err != nil {
		return err
	}

	label := "aplicado"
	if dry {
		label = "SIMULAÇÃO — nada foi escrito"
	}
	fmt.Printf("sincronização da coleção · fonte: %s · %s\n\n", src, label)
	fmt.Printf("  entram     %3d\n", len(diff.Added))
	fmt.Printf("  permanecem %3d\n", diff.Kept)
	fmt.Printf("  saem       %3d\n", len(diff.Removed))
	if len(diff.Changed) > 0 {
		fmt.Printf("  requantif. %3d\n", len(diff.Changed))
	}

	printEntries("\nENTRAM", diff.Added)
	printEntries("\nSAEM", diff.Removed)
	if len(diff.Changed) > 0 {
		fmt.Println("\nQUANTIDADE ALTERADA")
		for _, c := range diff.Changed {
			fmt.Printf("  %-40s %d → %d\n", truncate(c.Name, 40), c.From, c.To)
		}
	}

	if len(unresolved) > 0 {
		fmt.Printf("\nAVISO — %d nome(s) não resolveram no banco e foram gravados como vieram:\n", len(unresolved))
		for _, n := range unresolved {
			fmt.Printf("  %s\n", n)
		}
		fmt.Println("Confira a grafia (nomes em inglês) ou rode 'make refresh' se a carta for de um set recente.")
	}

	if dry {
		fmt.Println("\nRepita sem -dry-run para aplicar.")
		return nil
	}
	_, err = db.LoadCollection(dataDir())
	return err
}

func printEntries(title string, es []store.CollectionEntry) {
	if len(es) == 0 {
		return
	}
	fmt.Println(title)
	for _, e := range es {
		meta := e.Note
		if e.AddedAt != "" {
			meta = strings.TrimSpace(meta + " (desde " + e.AddedAt + ")")
		}
		fmt.Printf("  %2d× %-40s %s\n", e.Qty, truncate(e.Name, 40), meta)
	}
}

// collectionInput resolve de onde vem a lista: um arquivo, a entrada padrão ou
// os próprios argumentos. Devolve também um rótulo da origem, que entra no
// relatório — numa operação que apaga linhas, saber o que foi lido importa.
func collectionInput(file string, args []string) ([]deck.Entry, string, error) {
	switch {
	case file == "-":
		tmp, err := os.CreateTemp("", "mtgdb-collection-*.txt")
		if err != nil {
			return nil, "", err
		}
		defer os.Remove(tmp.Name())
		if _, err := io.Copy(tmp, bufio.NewReader(os.Stdin)); err != nil {
			tmp.Close()
			return nil, "", err
		}
		tmp.Close()
		es, err := deck.ParseList(tmp.Name())
		return es, "entrada padrão", err

	case file != "":
		es, err := deck.ParseList(file)
		if err != nil {
			return nil, "", err
		}
		return es, filepath.Base(file), nil

	case len(args) > 0:
		es := make([]deck.Entry, 0, len(args))
		for _, a := range args {
			es = append(es, deck.Entry{Qty: 1, Name: a})
		}
		return es, "argumentos da linha de comando", nil
	}
	return nil, "", fmt.Errorf("uso: mtgdb collection -sync -file <lista.txt> [-dry-run]")
}
