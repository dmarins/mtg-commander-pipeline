package store

import (
	"database/sql"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Normalize reduz um nome de carta a uma chave comparável: minúsculas, sem
// acentos e sem pontuação. Serve para casar entradas digitadas à mão
// ("Loran's Escape", "loran escape", "LORANS ESCAPE") com o nome oficial.
func Normalize(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	out, _, err := transform.String(t, s)
	if err != nil {
		out = s
	}
	out = strings.ToLower(out)

	var b strings.Builder
	prevSpace := false
	for _, r := range out {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			prevSpace = false
		case r == '\'' || r == '’' || r == '´' || r == '`':
			// Apóstrofo é ELIDIDO, não vira espaço: quem digita sem ele
			// ("Lorans Escape") precisa casar com o nome oficial
			// ("Loran's Escape"). Trocar por espaço quebraria justamente isso.
		default:
			// Os demais separadores (vírgula, hífen, barra) viram um espaço.
			if !prevSpace {
				b.WriteRune(' ')
				prevSpace = true
			}
		}
	}
	return strings.TrimSpace(b.String())
}

// Match é o resultado de resolver um nome digitado para uma carta do banco.
type Match struct {
	Query  string
	Card   Card
	Found  bool
	How    string   // exact | normalized | face | prefix | fts
	Ambigs []string // candidatos, quando a resolução foi ambígua
}

// Resolve traduz um nome digitado no registro oracle correspondente.
//
// A escada de tentativas vai da mais segura para a mais frouxa; a coluna How
// registra por onde passou, para que o chamador possa desconfiar de um acerto
// obtido por busca textual. Nomes de carta são sempre em inglês (regra 8 do
// CLAUDE.md), então entradas em português não resolvem aqui de propósito — elas
// aparecem como não encontradas em vez de casar com a carta errada.
func (d *DB) Resolve(query string) (Match, error) {
	m := Match{Query: query}
	q := strings.TrimSpace(query)
	if q == "" {
		return m, nil
	}

	// 1. Nome oficial exato (case-insensitive).
	row := d.sql.QueryRow(`SELECT `+cardCols+` FROM cards WHERE name = ? COLLATE NOCASE`, q)
	if c, err := scanCard(row); err == nil {
		return d.finish(&m, c, "exact")
	} else if err != sql.ErrNoRows {
		return m, err
	}

	// 2. Nome normalizado — cobre apóstrofos, vírgulas e acentos, e também
	//    cada face isolada de um MDFC ("Barracks of the Thousand").
	norm := Normalize(q)
	ids, err := d.idsByNorm(norm)
	if err != nil {
		return m, err
	}
	if len(ids) == 1 {
		c, ok, err := d.ByOracleID(ids[0])
		if err != nil {
			return m, err
		}
		if ok {
			how := "normalized"
			if !strings.EqualFold(Normalize(c.Name), norm) {
				how = "face"
			}
			m.Card, m.Found, m.How = c, true, how
			return m, nil
		}
	} else if len(ids) > 1 {
		return d.ambiguous(&m, ids)
	}

	// 3. Prefixo — resolve abreviações ("Alibou", "Kilo").
	ids, err = d.idsByPrefix(norm)
	if err != nil {
		return m, err
	}
	if len(ids) == 1 {
		c, ok, err := d.ByOracleID(ids[0])
		if err != nil {
			return m, err
		}
		if ok {
			m.Card, m.Found, m.How = c, true, "prefix"
			return m, nil
		}
	} else if len(ids) > 1 {
		return d.ambiguous(&m, ids)
	}

	// 4. Última tentativa: busca textual no nome. Menos confiável — o chamador
	//    deve conferir o resultado antes de usar.
	rows, err := d.sql.Query(
		`SELECT oracle_id FROM cards_fts WHERE cards_fts MATCH ? LIMIT 5`,
		`name : `+ftsQuote(norm))
	if err != nil {
		return m, nil // FTS falhou (sintaxe): trata como não encontrado
	}
	defer rows.Close()
	var found []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return m, err
		}
		found = append(found, id)
	}
	if len(found) == 1 {
		c, ok, err := d.ByOracleID(found[0])
		if err != nil {
			return m, err
		}
		if ok {
			m.Card, m.Found, m.How = c, true, "fts"
			return m, nil
		}
	} else if len(found) > 1 {
		return d.ambiguous(&m, found)
	}
	return m, nil
}

func (d *DB) finish(m *Match, c Card, how string) (Match, error) {
	if err := d.loadFaces(&c); err != nil {
		return *m, err
	}
	m.Card, m.Found, m.How = c, true, how
	return *m, nil
}

func (d *DB) idsByNorm(norm string) ([]string, error) {
	rows, err := d.sql.Query(`SELECT DISTINCT oracle_id FROM card_names WHERE norm = ?`, norm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanIDs(rows)
}

func (d *DB) idsByPrefix(norm string) ([]string, error) {
	rows, err := d.sql.Query(
		`SELECT DISTINCT oracle_id FROM card_names WHERE norm LIKE ? LIMIT 6`, norm+" %")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanIDs(rows)
}

func scanIDs(rows *sql.Rows) ([]string, error) {
	var out []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

func (d *DB) ambiguous(m *Match, ids []string) (Match, error) {
	for _, id := range ids {
		if c, ok, err := d.ByOracleID(id); err == nil && ok {
			m.Ambigs = append(m.Ambigs, c.Name)
		}
	}
	return *m, nil
}

// ftsQuote escapa um termo para uso literal numa query FTS5.
func ftsQuote(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
