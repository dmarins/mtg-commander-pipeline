// Package store persiste o bulk data do Scryfall em SQLite e responde as
// consultas que o pipeline faz sobre cartas, tags, rulings, coleção e preços.
package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// SchemaVersion identifica o formato dos dados derivados.
//
// INCREMENTE sempre que mudar o schema OU a função Normalize: os nomes
// normalizados ficam materializados em card_names durante o build, então uma
// mudança na normalização deixa o índice inconsistente com o código — e a falha
// aparece como "carta não encontrada", que é fácil confundir com erro de grafia.
const SchemaVersion = "2"

// DB embrulha a conexão SQLite.
type DB struct{ sql *sql.DB }

// Face é uma face de carta modal/transformável.
type Face struct {
	Name       string
	ManaCost   string
	TypeLine   string
	OracleText string
	Power      string
	Toughness  string
}

// Card é o registro oracle de uma carta.
type Card struct {
	OracleID       string
	Name           string
	ManaCost       string
	CMC            float64
	TypeLine       string
	OracleText     string
	Power          string
	Toughness      string
	Loyalty        string
	Colors         string
	ColorIdentity  string
	Keywords       string
	Layout         string
	LegalCommander string
	EDHRecRank     int
	ScryfallURI    string
	Faces          []Face
}

// AllText devolve o texto da carta somado ao de todas as faces, para as
// heurísticas não perderem o conteúdo de MDFCs (cujo oracle_text raiz é vazio).
func (c Card) AllText() string {
	var b strings.Builder
	b.WriteString(c.OracleText)
	for _, f := range c.Faces {
		b.WriteString("\n")
		b.WriteString(f.OracleText)
	}
	return b.String()
}

// AllTypes devolve a type line somada à das faces.
func (c Card) AllTypes() string {
	var b strings.Builder
	b.WriteString(c.TypeLine)
	for _, f := range c.Faces {
		b.WriteString(" // ")
		b.WriteString(f.TypeLine)
	}
	return b.String()
}

// Open abre (criando se preciso) o banco no caminho informado.
func Open(path string) (*DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	h, err := sql.Open("sqlite", path+"?_pragma=busy_timeout(10000)")
	if err != nil {
		return nil, err
	}
	if _, err := h.Exec(schemaSQL); err != nil {
		h.Close()
		return nil, fmt.Errorf("aplicando schema: %w", err)
	}
	return &DB{sql: h}, nil
}

func (d *DB) Close() error { return d.sql.Close() }

// SQL expõe a conexão para os comandos que montam consultas próprias.
func (d *DB) SQL() *sql.DB { return d.sql }

// Meta lê um valor da tabela meta.
func (d *DB) Meta(key string) string {
	var v string
	_ = d.sql.QueryRow(`SELECT value FROM meta WHERE key = ?`, key).Scan(&v)
	return v
}

// SetMeta grava um valor na tabela meta.
func (d *DB) SetMeta(key, value string) error {
	_, err := d.sql.Exec(
		`INSERT INTO meta(key,value) VALUES(?,?) ON CONFLICT(key) DO UPDATE SET value=excluded.value`,
		key, value)
	return err
}

// Built informa se o banco já recebeu um build de cartas.
func (d *DB) Built() bool {
	var n int
	_ = d.sql.QueryRow(`SELECT count(*) FROM cards`).Scan(&n)
	return n > 0
}

// CardCount devolve quantas cartas oracle estão carregadas.
func (d *DB) CardCount() int {
	var n int
	_ = d.sql.QueryRow(`SELECT count(*) FROM cards`).Scan(&n)
	return n
}

const cardCols = `oracle_id, name, mana_cost, cmc, type_line, oracle_text, power,
	toughness, loyalty, colors, color_identity, keywords, layout,
	legal_commander, edhrec_rank, scryfall_uri`

func scanCard(rows interface{ Scan(...any) error }) (Card, error) {
	var c Card
	var mc, tl, ot, pw, tg, ly, co, ci, kw, la, lc, uri sql.NullString
	var cmc sql.NullFloat64
	var rank sql.NullInt64
	err := rows.Scan(&c.OracleID, &c.Name, &mc, &cmc, &tl, &ot, &pw, &tg, &ly,
		&co, &ci, &kw, &la, &lc, &rank, &uri)
	if err != nil {
		return c, err
	}
	c.ManaCost, c.CMC, c.TypeLine, c.OracleText = mc.String, cmc.Float64, tl.String, ot.String
	c.Power, c.Toughness, c.Loyalty = pw.String, tg.String, ly.String
	c.Colors, c.ColorIdentity, c.Keywords, c.Layout = co.String, ci.String, kw.String, la.String
	c.LegalCommander, c.ScryfallURI, c.EDHRecRank = lc.String, uri.String, int(rank.Int64)
	return c, nil
}

func (d *DB) loadFaces(c *Card) error {
	rows, err := d.sql.Query(
		`SELECT name, mana_cost, type_line, oracle_text, power, toughness
		   FROM card_faces WHERE oracle_id = ? ORDER BY idx`, c.OracleID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var f Face
		var n, mc, tl, ot, pw, tg sql.NullString
		if err := rows.Scan(&n, &mc, &tl, &ot, &pw, &tg); err != nil {
			return err
		}
		f.Name, f.ManaCost, f.TypeLine = n.String, mc.String, tl.String
		f.OracleText, f.Power, f.Toughness = ot.String, pw.String, tg.String
		c.Faces = append(c.Faces, f)
	}
	return rows.Err()
}

// ByOracleID busca uma carta pelo oracle_id.
func (d *DB) ByOracleID(id string) (Card, bool, error) {
	row := d.sql.QueryRow(`SELECT `+cardCols+` FROM cards WHERE oracle_id = ?`, id)
	c, err := scanCard(row)
	if err == sql.ErrNoRows {
		return Card{}, false, nil
	}
	if err != nil {
		return Card{}, false, err
	}
	if err := d.loadFaces(&c); err != nil {
		return c, true, err
	}
	return c, true, nil
}

// Tags devolve as tags curadas (Scryfall Tagger) de uma carta.
func (d *DB) Tags(oracleID string) ([]string, error) {
	rows, err := d.sql.Query(`SELECT tag FROM card_tags WHERE oracle_id = ? ORDER BY tag`, oracleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

// Ruling é uma decisão oficial publicada sobre a carta.
type Ruling struct {
	Source      string
	PublishedAt string
	Comment     string
}

// Rulings devolve as rulings oficiais de uma carta.
func (d *DB) Rulings(oracleID string) ([]Ruling, error) {
	rows, err := d.sql.Query(
		`SELECT source, published_at, comment FROM rulings
		  WHERE oracle_id = ? ORDER BY published_at, rowid`, oracleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Ruling
	for rows.Next() {
		var r Ruling
		if err := rows.Scan(&r.Source, &r.PublishedAt, &r.Comment); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// PriceObs é uma observação de preço — sempre datada, nunca sobrescrita.
type PriceObs struct {
	Name       string
	Source     string
	Kind       string
	Value      float64
	Currency   string
	CapturedAt string
	Note       string
}

// AgeDays devolve a idade da observação em dias.
func (p PriceObs) AgeDays() int {
	t, err := time.Parse("2006-01-02", p.CapturedAt)
	if err != nil {
		return -1
	}
	return int(time.Since(t).Hours() / 24)
}

// LatestPrice devolve a observação mais recente de (carta, fonte, tipo).
func (d *DB) LatestPrice(name, source, kind string) (PriceObs, bool, error) {
	row := d.sql.QueryRow(
		`SELECT name, source, kind, value, currency, captured_at, COALESCE(note,'')
		   FROM prices
		  WHERE name = ? COLLATE NOCASE AND source = ? AND kind = ?
		  ORDER BY captured_at DESC LIMIT 1`, name, source, kind)
	var p PriceObs
	err := row.Scan(&p.Name, &p.Source, &p.Kind, &p.Value, &p.Currency, &p.CapturedAt, &p.Note)
	if err == sql.ErrNoRows {
		return p, false, nil
	}
	if err != nil {
		return p, false, err
	}
	return p, true, nil
}

// PriceHistory devolve todas as observações de uma carta, da mais recente para a mais antiga.
func (d *DB) PriceHistory(name string) ([]PriceObs, error) {
	rows, err := d.sql.Query(
		`SELECT name, source, kind, value, currency, captured_at, COALESCE(note,'')
		   FROM prices WHERE name = ? COLLATE NOCASE
		  ORDER BY captured_at DESC, source, kind`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PriceObs
	for rows.Next() {
		var p PriceObs
		if err := rows.Scan(&p.Name, &p.Source, &p.Kind, &p.Value, &p.Currency, &p.CapturedAt, &p.Note); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// InCollection informa a quantidade que o usuário possui da carta.
func (d *DB) InCollection(name string) (int, bool, error) {
	var qty int
	err := d.sql.QueryRow(`SELECT qty FROM collection WHERE name = ? COLLATE NOCASE`, name).Scan(&qty)
	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	return qty, true, nil
}
