package store

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dmarins/mtg-commander-pipeline/internal/scryfall"
)

type bulkFace struct {
	Name       string `json:"name"`
	ManaCost   string `json:"mana_cost"`
	TypeLine   string `json:"type_line"`
	OracleText string `json:"oracle_text"`
	Power      string `json:"power"`
	Toughness  string `json:"toughness"`
}

type bulkCard struct {
	OracleID      string            `json:"oracle_id"`
	Name          string            `json:"name"`
	ManaCost      string            `json:"mana_cost"`
	CMC           float64           `json:"cmc"`
	TypeLine      string            `json:"type_line"`
	OracleText    string            `json:"oracle_text"`
	Power         string            `json:"power"`
	Toughness     string            `json:"toughness"`
	Loyalty       string            `json:"loyalty"`
	Colors        []string          `json:"colors"`
	ColorIdentity []string          `json:"color_identity"`
	Keywords      []string          `json:"keywords"`
	Layout        string            `json:"layout"`
	Legalities    map[string]string `json:"legalities"`
	Set           string            `json:"set"`
	ReleasedAt    string            `json:"released_at"`
	EDHRecRank    int               `json:"edhrec_rank"`
	ScryfallURI   string            `json:"scryfall_uri"`
	Faces         []bulkFace        `json:"card_faces"`
}

// Progress recebe mensagens de andamento do build.
type Progress func(format string, args ...any)

// BuildCards recarrega as tabelas de cartas a partir do bulk oracle_cards.
func (d *DB) BuildCards(info scryfall.BulkInfo, log Progress) (int, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	for _, stmt := range []string{
		`DELETE FROM cards`, `DELETE FROM card_faces`,
		`DELETE FROM card_names`, `DELETE FROM cards_fts`,
	} {
		if _, err := tx.Exec(stmt); err != nil {
			return 0, err
		}
	}

	insCard, err := tx.Prepare(`INSERT OR REPLACE INTO cards
		(oracle_id,name,mana_cost,cmc,type_line,oracle_text,power,toughness,loyalty,
		 colors,color_identity,keywords,layout,legal_commander,edhrec_rank,set_code,
		 released_at,scryfall_uri)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer insCard.Close()

	insFace, err := tx.Prepare(`INSERT OR REPLACE INTO card_faces
		(oracle_id,idx,name,mana_cost,type_line,oracle_text,power,toughness)
		VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer insFace.Close()

	insName, err := tx.Prepare(`INSERT OR IGNORE INTO card_names (norm,oracle_id,is_face) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer insName.Close()

	insFTS, err := tx.Prepare(`INSERT INTO cards_fts (name,type_line,oracle_text,oracle_id) VALUES (?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer insFTS.Close()

	n := 0
	err = scryfall.Stream(info, func(line []byte) error {
		var c bulkCard
		if err := json.Unmarshal(line, &c); err != nil {
			return nil // registro isolado ilegível não deve abortar o build
		}
		if c.OracleID == "" || c.Name == "" {
			return nil
		}

		// MDFCs trazem o texto nas faces; concatena para busca e heurísticas.
		text, types := c.OracleText, c.TypeLine
		if len(c.Faces) > 0 {
			var tb, yb []string
			for _, f := range c.Faces {
				if f.OracleText != "" {
					tb = append(tb, f.OracleText)
				}
				if f.TypeLine != "" {
					yb = append(yb, f.TypeLine)
				}
			}
			if text == "" {
				text = strings.Join(tb, "\n--\n")
			}
			if types == "" {
				types = strings.Join(yb, " // ")
			}
		}

		if _, err := insCard.Exec(c.OracleID, c.Name, c.ManaCost, c.CMC, types, text,
			c.Power, c.Toughness, c.Loyalty,
			strings.Join(c.Colors, ""), strings.Join(c.ColorIdentity, ""),
			strings.Join(c.Keywords, ", "), c.Layout, c.Legalities["commander"],
			c.EDHRecRank, c.Set, c.ReleasedAt, c.ScryfallURI); err != nil {
			return err
		}

		if _, err := insName.Exec(Normalize(c.Name), c.OracleID, 0); err != nil {
			return err
		}
		for i, f := range c.Faces {
			if _, err := insFace.Exec(c.OracleID, i, f.Name, f.ManaCost, f.TypeLine,
				f.OracleText, f.Power, f.Toughness); err != nil {
				return err
			}
			if f.Name != "" && !strings.EqualFold(f.Name, c.Name) {
				if _, err := insName.Exec(Normalize(f.Name), c.OracleID, 1); err != nil {
					return err
				}
			}
		}

		if _, err := insFTS.Exec(c.Name, types, text, c.OracleID); err != nil {
			return err
		}

		n++
		if log != nil && n%5000 == 0 {
			log("  %d cartas...", n)
		}
		return nil
	})
	if err != nil {
		return n, err
	}

	if err := tx.Commit(); err != nil {
		return n, err
	}
	_ = d.SetMeta("cards_updated_at", info.UpdatedAt)
	_ = d.SetMeta("cards_built_at", time.Now().Format(time.RFC3339))
	return n, nil
}

type bulkTag struct {
	Label       string `json:"label"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Taggings    []struct {
		OracleID string `json:"oracle_id"`
		Weight   string `json:"weight"`
	} `json:"taggings"`
}

// BuildTags recarrega as tags curadas do Scryfall Tagger (as de otag:).
func (d *DB) BuildTags(info scryfall.BulkInfo, log Progress) (int, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	for _, stmt := range []string{`DELETE FROM card_tags`, `DELETE FROM tags`} {
		if _, err := tx.Exec(stmt); err != nil {
			return 0, err
		}
	}

	insTag, err := tx.Prepare(`INSERT OR REPLACE INTO tags (slug,description) VALUES (?,?)`)
	if err != nil {
		return 0, err
	}
	defer insTag.Close()

	insLink, err := tx.Prepare(`INSERT OR IGNORE INTO card_tags (oracle_id,tag,weight) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer insLink.Close()

	tags, links := 0, 0
	err = scryfall.Stream(info, func(line []byte) error {
		var t bulkTag
		if err := json.Unmarshal(line, &t); err != nil {
			return nil
		}
		slug := t.Slug
		if slug == "" {
			slug = t.Label
		}
		if slug == "" {
			return nil
		}
		if _, err := insTag.Exec(slug, t.Description); err != nil {
			return err
		}
		tags++
		for _, tg := range t.Taggings {
			if tg.OracleID == "" {
				continue
			}
			if _, err := insLink.Exec(tg.OracleID, slug, tg.Weight); err != nil {
				return err
			}
			links++
		}
		if log != nil && tags%2000 == 0 {
			log("  %d tags, %d vínculos...", tags, links)
		}
		return nil
	})
	if err != nil {
		return tags, err
	}
	if err := tx.Commit(); err != nil {
		return tags, err
	}
	_ = d.SetMeta("tags_updated_at", info.UpdatedAt)
	return tags, nil
}

type bulkRuling struct {
	OracleID    string `json:"oracle_id"`
	Source      string `json:"source"`
	PublishedAt string `json:"published_at"`
	Comment     string `json:"comment"`
}

// BuildRulings recarrega as rulings oficiais.
func (d *DB) BuildRulings(info scryfall.BulkInfo, log Progress) (int, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM rulings`); err != nil {
		return 0, err
	}
	ins, err := tx.Prepare(`INSERT INTO rulings (oracle_id,source,published_at,comment) VALUES (?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer ins.Close()

	n := 0
	err = scryfall.Stream(info, func(line []byte) error {
		var r bulkRuling
		if err := json.Unmarshal(line, &r); err != nil {
			return nil
		}
		if r.OracleID == "" {
			return nil
		}
		if _, err := ins.Exec(r.OracleID, r.Source, r.PublishedAt, r.Comment); err != nil {
			return err
		}
		n++
		if log != nil && n%20000 == 0 {
			log("  %d rulings...", n)
		}
		return nil
	})
	if err != nil {
		return n, err
	}
	if err := tx.Commit(); err != nil {
		return n, err
	}
	_ = d.SetMeta("rulings_updated_at", info.UpdatedAt)
	return n, nil
}

// BuildAll reconstrói cartas, tags e rulings.
func (d *DB) BuildAll(log Progress) error {
	idx, err := scryfall.Index()
	if err != nil {
		return fmt.Errorf("consultando bulk-data: %w", err)
	}

	steps := []struct {
		kind string
		run  func(scryfall.BulkInfo, Progress) (int, error)
	}{
		{"oracle_cards", d.BuildCards},
		{"oracle_tags", d.BuildTags},
		{"rulings", d.BuildRulings},
	}

	for _, s := range steps {
		info, ok := idx[s.kind]
		if !ok {
			return fmt.Errorf("bulk %q não encontrado no índice do Scryfall", s.kind)
		}
		if log != nil {
			log("%s (%.1f MB compactado, atualizado em %s)",
				s.kind, float64(info.CompressedSize)/(1<<20), info.UpdatedAt[:10])
		}
		n, err := s.run(info, log)
		if err != nil {
			return fmt.Errorf("importando %s: %w", s.kind, err)
		}
		if log != nil {
			log("  %d registros ✓", n)
		}
	}
	return d.SetMeta("schema_version", SchemaVersion)
}
