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

// AddToCollection acrescenta cartas a data/collection.tsv e ao banco, mantendo
// o arquivo ordenado por nome para o diff do git ficar estável.
func (d *DB) AddToCollection(dataDir string, names []string, note, when string) (int, error) {
	if _, err := d.LoadCollection(dataDir); err != nil {
		return 0, err
	}

	type entry struct {
		qty        int
		note, when string
	}
	current := map[string]entry{}
	rows, err := d.sql.Query(`SELECT name, qty, COALESCE(note,''), COALESCE(added_at,'') FROM collection`)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		var n, no, ad string
		var q int
		if err := rows.Scan(&n, &q, &no, &ad); err != nil {
			rows.Close()
			return 0, err
		}
		current[n] = entry{q, no, ad}
	}
	rows.Close()

	added := 0
	for _, n := range names {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if e, ok := current[n]; ok {
			e.qty++
			current[n] = e
		} else {
			current[n] = entry{1, note, when}
		}
		added++
	}

	keys := make([]string, 0, len(current))
	for k := range current {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	var b strings.Builder
	b.WriteString("name\tqty\tnote\tadded_at\n")
	for _, k := range keys {
		e := current[k]
		b.WriteString(fmt.Sprintf("%s\t%d\t%s\t%s\n", k, e.qty, e.note, e.when))
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return 0, err
	}
	if err := os.WriteFile(filepath.Join(dataDir, collectionFile), []byte(b.String()), 0o644); err != nil {
		return 0, err
	}

	_, err = d.LoadCollection(dataDir)
	return added, err
}
