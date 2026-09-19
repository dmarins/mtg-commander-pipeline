package store

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Os TSV de data/ são a FONTE de collection e prices — o banco é só um espelho
// consultável. São dados do usuário, não reconstruíveis do Scryfall: a coleção é
// o que ele possui, e cada linha de preço custou uma navegação manual na
// LigaMagic (a página monta o valor por JS). Ficam versionados em texto para ter
// diff legível e merge por linha; o .db fica fora do git porque é derivado.

const (
	collectionFile = "collection.tsv"
	pricesFile     = "prices.tsv"
)

func readTSV(path string, wantCols int, fn func(rec []string, line int) error) error {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil // arquivo opcional
	}
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64<<10), 1<<20)
	line := 0
	for sc.Scan() {
		line++
		raw := strings.TrimRight(sc.Text(), "\r")
		if raw == "" || strings.HasPrefix(raw, "#") {
			continue
		}
		rec := strings.Split(raw, "\t")
		if line == 1 && strings.EqualFold(strings.TrimSpace(rec[0]), "name") {
			continue // cabeçalho
		}
		if len(rec) < wantCols {
			return fmt.Errorf("%s:%d: esperava %d colunas separadas por TAB, achei %d",
				filepath.Base(path), line, wantCols, len(rec))
		}
		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}
		if err := fn(rec, line); err != nil {
			return fmt.Errorf("%s:%d: %w", filepath.Base(path), line, err)
		}
	}
	return sc.Err()
}

// LoadCollection substitui a tabela collection pelo conteúdo de data/collection.tsv.
// Colunas: name, qty, note
func (d *DB) LoadCollection(dataDir string) (int, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM collection`); err != nil {
		return 0, err
	}
	ins, err := tx.Prepare(`INSERT OR REPLACE INTO collection (name,qty,note,added_at) VALUES (?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer ins.Close()

	n := 0
	err = readTSV(filepath.Join(dataDir, collectionFile), 1, func(rec []string, _ int) error {
		name := rec[0]
		if name == "" {
			return nil
		}
		qty := 1
		if len(rec) > 1 && rec[1] != "" {
			if v, err := strconv.Atoi(rec[1]); err == nil {
				qty = v
			}
		}
		note := ""
		if len(rec) > 2 {
			note = rec[2]
		}
		added := ""
		if len(rec) > 3 {
			added = rec[3]
		}
		if _, err := ins.Exec(name, qty, note, added); err != nil {
			return err
		}
		n++
		return nil
	})
	if err != nil {
		return n, err
	}
	return n, tx.Commit()
}

// LoadPrices substitui a tabela prices pelo conteúdo de data/prices.tsv.
// Colunas: name, source, kind, value, currency, captured_at, note
func (d *DB) LoadPrices(dataDir string) (int, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM prices`); err != nil {
		return 0, err
	}
	ins, err := tx.Prepare(`INSERT OR REPLACE INTO prices
		(name,source,kind,value,currency,captured_at,note) VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer ins.Close()

	n := 0
	err = readTSV(filepath.Join(dataDir, pricesFile), 6, func(rec []string, _ int) error {
		// Aceita vírgula decimal: os preços vêm da LigaMagic em formato BR.
		val, err := strconv.ParseFloat(strings.ReplaceAll(rec[3], ",", "."), 64)
		if err != nil {
			return fmt.Errorf("valor %q não é numérico", rec[3])
		}
		note := ""
		if len(rec) > 6 {
			note = rec[6]
		}
		if _, err := ins.Exec(rec[0], rec[1], rec[2], val, rec[4], rec[5], note); err != nil {
			return err
		}
		n++
		return nil
	})
	if err != nil {
		return n, err
	}
	return n, tx.Commit()
}

// AppendPrice acrescenta uma observação ao data/prices.tsv e ao banco.
// Nunca sobrescreve: preço é observação datada, não um campo que se atualiza.
func (d *DB) AppendPrice(dataDir string, p PriceObs) error {
	path := filepath.Join(dataDir, pricesFile)
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return err
	}

	// Cria com cabeçalho se ainda não existir.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		header := "name\tsource\tkind\tvalue\tcurrency\tcaptured_at\tnote\n"
		if err := os.WriteFile(path, []byte(header), 0o644); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("%s\t%s\t%s\t%.2f\t%s\t%s\t%s\n",
		p.Name, p.Source, p.Kind, p.Value, p.Currency, p.CapturedAt, p.Note)
	if _, err := f.WriteString(line); err != nil {
		return err
	}

	_, err = d.sql.Exec(`INSERT OR REPLACE INTO prices
		(name,source,kind,value,currency,captured_at,note) VALUES (?,?,?,?,?,?,?)`,
		p.Name, p.Source, p.Kind, p.Value, p.Currency, p.CapturedAt, p.Note)
	return err
}

// CollectionEntry é uma linha de data/collection.tsv.
type CollectionEntry struct {
	Name    string
	Qty     int
	Note    string
	AddedAt string
}

// CollectionDiff descreve o que um Sync fez (ou faria, em simulação).
type CollectionDiff struct {
	Added    []CollectionEntry // não estavam no TSV
	Removed  []CollectionEntry // estavam e saíram — com note e added_at antigos
	Changed  []CollectionChange
	Kept     int
	Unparsed int // linhas da entrada sem nome aproveitável
}

// CollectionChange é uma carta que ficou, mas mudou de quantidade.
type CollectionChange struct {
	Name     string
	From, To int
}

// ReadCollection lê data/collection.tsv. Arquivo ausente devolve lista vazia —
// a primeira sincronização de uma máquina nova é um caso normal, não um erro.
func ReadCollection(dataDir string) ([]CollectionEntry, error) {
	var out []CollectionEntry
	err := readTSV(filepath.Join(dataDir, collectionFile), 1, func(rec []string, _ int) error {
		if rec[0] == "" {
			return nil
		}
		e := CollectionEntry{Name: rec[0], Qty: 1}
		if len(rec) > 1 && rec[1] != "" {
			if v, err := strconv.Atoi(rec[1]); err == nil {
				e.Qty = v
			}
		}
		if len(rec) > 2 {
			e.Note = rec[2]
		}
		if len(rec) > 3 {
			e.AddedAt = rec[3]
		}
		out = append(out, e)
		return nil
	})
	return out, err
}

// WriteCollection reescreve data/collection.tsv ordenado por nome, para o diff
// do git ficar estável entre execuções.
func WriteCollection(dataDir string, entries []CollectionEntry) error {
	sorted := append([]CollectionEntry(nil), entries...)
	sort.Slice(sorted, func(i, j int) bool {
		return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
	})

	var b strings.Builder
	b.WriteString("name\tqty\tnote\tadded_at\n")
	for _, e := range sorted {
		b.WriteString(fmt.Sprintf("%s\t%d\t%s\t%s\n", e.Name, e.Qty, e.Note, e.AddedAt))
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dataDir, collectionFile), []byte(b.String()), 0o644)
}

// SyncCollection substitui a coleção pela lista informada: o que não está em
// want sai do arquivo.
//
// É uma substituição, não um merge — a lista recebida é o retrato atual da
// caixa. Mas note e added_at de quem permanece são preservados: essas colunas
// são anotação manual do usuário, e uma lista de nomes não as carrega. Quem
// entra ganha note e a data de hoje.
//
// dryRun calcula o diff sem tocar no disco, porque uma substituição derivada de
// uma lista colada só é segura depois que o usuário viu o que sairia.
func SyncCollection(dataDir string, want []CollectionEntry, note, when string, dryRun bool) (CollectionDiff, error) {
	var diff CollectionDiff

	current, err := ReadCollection(dataDir)
	if err != nil {
		return diff, err
	}
	byKey := make(map[string]CollectionEntry, len(current))
	order := make([]string, 0, len(current))
	for _, e := range current {
		k := strings.ToLower(e.Name)
		if _, dup := byKey[k]; !dup {
			order = append(order, k)
		}
		byKey[k] = e
	}

	final := make([]CollectionEntry, 0, len(want))
	seen := make(map[string]bool, len(want))
	for _, w := range want {
		if strings.TrimSpace(w.Name) == "" {
			diff.Unparsed++
			continue
		}
		k := strings.ToLower(w.Name)
		if w.Qty < 1 {
			w.Qty = 1
		}
		// Nome repetido na entrada: soma, como duas linhas "1 Forest" somariam.
		if seen[k] {
			for i := range final {
				if strings.EqualFold(final[i].Name, w.Name) {
					final[i].Qty += w.Qty
					break
				}
			}
			continue
		}
		seen[k] = true

		if old, ok := byKey[k]; ok {
			kept := CollectionEntry{Name: old.Name, Qty: w.Qty, Note: old.Note, AddedAt: old.AddedAt}
			if old.Qty != w.Qty {
				diff.Changed = append(diff.Changed, CollectionChange{Name: old.Name, From: old.Qty, To: w.Qty})
			}
			diff.Kept++
			final = append(final, kept)
			continue
		}
		e := CollectionEntry{Name: w.Name, Qty: w.Qty, Note: note, AddedAt: when}
		if w.Note != "" {
			e.Note = w.Note
		}
		diff.Added = append(diff.Added, e)
		final = append(final, e)
	}

	for _, k := range order {
		if !seen[k] {
			diff.Removed = append(diff.Removed, byKey[k])
		}
	}

	if dryRun {
		return diff, nil
	}
	if err := WriteCollection(dataDir, final); err != nil {
		return diff, err
	}
	return diff, nil
}

// AddToCollection acrescenta cartas a data/collection.tsv e ao banco, mantendo
// o arquivo ordenado por nome para o diff do git ficar estável.
func (d *DB) AddToCollection(dataDir string, names []string, note, when string) (int, error) {
	current, err := ReadCollection(dataDir)
	if err != nil {
		return 0, err
	}
	index := make(map[string]int, len(current))
	for i, e := range current {
		index[strings.ToLower(e.Name)] = i
	}

	added := 0
	for _, n := range names {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if i, ok := index[strings.ToLower(n)]; ok {
			current[i].Qty++
		} else {
			index[strings.ToLower(n)] = len(current)
			current = append(current, CollectionEntry{Name: n, Qty: 1, Note: note, AddedAt: when})
		}
		added++
	}

	if err := WriteCollection(dataDir, current); err != nil {
		return 0, err
	}
	_, err = d.LoadCollection(dataDir)
	return added, err
}
