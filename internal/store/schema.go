package store

// schemaSQL define o banco. Tudo aqui, exceto collection e prices, é derivado do
// bulk data do Scryfall e pode ser reconstruído a qualquer momento com `mtgdb build`.
//
// collection e prices vêm dos TSV versionados em data/ — são dados do usuário,
// não reconstruíveis, e por isso o banco é sempre um espelho deles, nunca a fonte.
const schemaSQL = `
PRAGMA journal_mode = WAL;
PRAGMA synchronous = NORMAL;

CREATE TABLE IF NOT EXISTS cards (
  oracle_id       TEXT PRIMARY KEY,
  name            TEXT NOT NULL,
  mana_cost       TEXT,
  cmc             REAL,
  type_line       TEXT,
  oracle_text     TEXT,
  power           TEXT,
  toughness       TEXT,
  loyalty         TEXT,
  colors          TEXT,
  color_identity  TEXT,
  keywords        TEXT,
  layout          TEXT,
  legal_commander TEXT,
  edhrec_rank     INTEGER,
  set_code        TEXT,
  released_at     TEXT,
  scryfall_uri    TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_cards_name  ON cards(name COLLATE NOCASE);
CREATE INDEX        IF NOT EXISTS idx_cards_cmc   ON cards(cmc);
CREATE INDEX        IF NOT EXISTS idx_cards_ci    ON cards(color_identity);
CREATE INDEX        IF NOT EXISTS idx_cards_type  ON cards(type_line);

-- Faces de cartas modais/transformáveis (MDFC, transform, split).
-- Thousand Moons Smithy // Barracks of the Thousand tem o texto útil na face 0;
-- sem esta tabela, MDFCs aparecem com oracle_text vazio.
CREATE TABLE IF NOT EXISTS card_faces (
  oracle_id   TEXT NOT NULL,
  idx         INTEGER NOT NULL,
  name        TEXT,
  mana_cost   TEXT,
  type_line   TEXT,
  oracle_text TEXT,
  power       TEXT,
  toughness   TEXT,
  PRIMARY KEY (oracle_id, idx)
);

-- Nomes normalizados (minúsculas, sem acentos e sem pontuação) para resolver
-- entradas digitadas à mão. Cobre também cada face isolada de um MDFC.
CREATE TABLE IF NOT EXISTS card_names (
  norm      TEXT NOT NULL,
  oracle_id TEXT NOT NULL,
  is_face   INTEGER NOT NULL DEFAULT 0,
  PRIMARY KEY (norm, oracle_id)
);
CREATE INDEX IF NOT EXISTS idx_card_names_norm ON card_names(norm);

CREATE VIRTUAL TABLE IF NOT EXISTS cards_fts USING fts5(
  name, type_line, oracle_text,
  oracle_id UNINDEXED,
  tokenize = 'unicode61 remove_diacritics 2'
);

-- Tags curadas do Scryfall Tagger (as mesmas por trás de otag: nas buscas).
CREATE TABLE IF NOT EXISTS card_tags (
  oracle_id TEXT NOT NULL,
  tag       TEXT NOT NULL,
  weight    TEXT,
  PRIMARY KEY (oracle_id, tag)
);
CREATE INDEX IF NOT EXISTS idx_card_tags_tag ON card_tags(tag);

CREATE TABLE IF NOT EXISTS tags (
  slug        TEXT PRIMARY KEY,
  description TEXT
);

CREATE TABLE IF NOT EXISTS rulings (
  oracle_id    TEXT NOT NULL,
  source       TEXT,
  published_at TEXT,
  comment      TEXT
);
CREATE INDEX IF NOT EXISTS idx_rulings_oracle ON rulings(oracle_id);

CREATE TABLE IF NOT EXISTS collection (
  name     TEXT PRIMARY KEY,
  qty      INTEGER NOT NULL DEFAULT 1,
  note     TEXT,
  added_at TEXT
);

-- Preços são APPEND-ONLY: uma observação datada, nunca um valor que se atualiza.
-- "Em 12/08 o menor da LigaMagic para Loran's Escape era R$ 15,45" continua
-- verdadeiro para sempre, mesmo que hoje seja outro número. É essa série que
-- permite saber a IDADE de uma cotação e quais cartas de fato oscilam.
CREATE TABLE IF NOT EXISTS prices (
  name        TEXT NOT NULL,
  source      TEXT NOT NULL,
  kind        TEXT NOT NULL,
  value       REAL NOT NULL,
  currency    TEXT NOT NULL DEFAULT 'BRL',
  captured_at TEXT NOT NULL,
  note        TEXT,
  PRIMARY KEY (name, source, kind, captured_at)
);
CREATE INDEX IF NOT EXISTS idx_prices_name ON prices(name COLLATE NOCASE);

CREATE TABLE IF NOT EXISTS meta (
  key   TEXT PRIMARY KEY,
  value TEXT
);
`
