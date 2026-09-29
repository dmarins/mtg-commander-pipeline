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

	// 1. Nome oficial exato (case-insensitive). Se um token ou carta de arte
	//    tiver o mesmo nome de uma carta de jogo, a carta de jogo vem primeiro.
	row := d.sql.QueryRow(`SELECT `+cardCols+` FROM cards WHERE name = ? COLLATE NOCASE
		ORDER BY layout IN (`+nonGameLayoutsSQL+`) LIMIT 1`, q)
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
	if done, err := d.settle(&m, ids, func(c Card) string {
		if !strings.EqualFold(Normalize(c.Name), norm) {
			return "face"
		}
		return "normalized"
	}); done {
		return m, err
	}

	// 3. Prefixo — resolve abreviações ("Alibou", "Kilo").
	ids, err = d.idsByPrefix(norm)
	if err != nil {
		return m, err
	}
	if done, err := d.settle(&m, ids, fixedHow("prefix")); done {
		return m, err
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
	found, err := scanIDs(rows)
	if err != nil {
		return m, err
	}
	_, err = d.settle(&m, found, fixedHow("fts"))
	return m, err
}

// nonGameLayouts são os layouts do bulk oracle_cards que não são cartas de jogo:
// cartas de arte, tokens, emblemas e as peças de variantes casuais. Elas repetem
// o nome de cartas reais — a carta de arte "Goblin Glasswright // Goblin
// Glasswright" tem a mesma face que "Goblin Glasswright // Craft with Pride" —
// e, sem desempate, tornavam ambígua a resolução pela face.
var nonGameLayouts = map[string]bool{
	"art_series": true, "token": true, "double_faced_token": true, "emblem": true,
	"front_card": true, "vanguard": true, "planar": true, "scheme": true,
}

const nonGameLayoutsSQL = `'art_series','token','double_faced_token','emblem',
	'front_card','vanguard','planar','scheme'`

func fixedHow(how string) func(Card) string {
	return func(Card) string { return how }
}

// settle decide entre os candidatos de um degrau da escada e preenche m.
// done indica que a escada para aqui — por acerto ou por ambiguidade; sem
// candidatos, a resolução segue para o próximo degrau.
func (d *DB) settle(m *Match, ids []string, how func(Card) string) (done bool, err error) {
	if len(ids) == 0 {
		return false, nil
	}
	c, ok, err := d.pick(m, ids)
	if err != nil {
		return true, err
	}
	if ok {
		m.Card, m.Found, m.How = c, true, how(c)
		return true, nil
	}
	return len(m.Ambigs) > 0, nil
}

// pick escolhe entre cartas que casaram o mesmo nome. Uma entrada fora de jogo
// só vence quando não há carta de jogo na disputa; entre cartas de jogo não há
// desempate — os candidatos vão para m.Ambigs e o chamador decide.
func (d *DB) pick(m *Match, ids []string) (Card, bool, error) {
	var game, other []Card
	for _, id := range ids {
		c, ok, err := d.ByOracleID(id)
		if err != nil {
			return Card{}, false, err
		}
		if !ok {
			continue
		}
		if nonGameLayouts[c.Layout] {
			other = append(other, c)
		} else {
			game = append(game, c)
		}
	}
	cands := game
	if len(cands) == 0 {
		cands = other
	}
	if len(cands) == 1 {
		return cands[0], true, nil
	}
	for _, c := range cands {
		m.Ambigs = append(m.Ambigs, c.Name)
	}
	return Card{}, false, nil
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

// ftsQuote escapa um termo para uso literal numa query FTS5.
func ftsQuote(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
