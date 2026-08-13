// Package deck lê a decklist de um diretório decks/<slug>/ e monta a ficha de
// funções exigida por references/card-evaluation-checklist.md.
package deck

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Entry é uma linha da decklist.
type Entry struct {
	Qty     int
	Name    string
	Section string // seção do deck.md de onde veio, quando houver
}

var (
	// "1 Sol Ring", "1x Sol Ring", "Sol Ring"
	reQty = regexp.MustCompile(`^\s*(\d+)\s*[xX]?\s+(.+?)\s*$`)
	// "Plains ×6" / "Island x9" — como os básicos aparecem agrupados no deck.md
	reSuffixQty = regexp.MustCompile(`^(.+?)\s*[×xX]\s*(\d+)$`)
	// "## Criaturas (22)" ou "### Artefatos"
	reHeading = regexp.MustCompile(`^#{2,}\s*(.+?)\s*$`)
)

// candidateFiles devolve, em ordem de preferência, os arquivos que podem conter
// a lista: a fonte de verdade (deck.md) primeiro, depois exportações e briefing.
func candidateFiles(dir string) []string {
	return []string{
		filepath.Join(dir, "deck.md"),
		filepath.Join(dir, "lista.txt"),
		filepath.Join(dir, "00-briefing.md"),
	}
}

// Load lê a decklist do diretório do deck, tentando os arquivos conhecidos.
// Devolve também qual arquivo foi usado.
func Load(dir string) ([]Entry, string, error) {
	var lastErr error
	for _, path := range candidateFiles(dir) {
		entries, err := parseFile(path)
		if err != nil {
			if !os.IsNotExist(err) {
				lastErr = err
			}
			continue
		}
		if len(entries) > 0 {
			return entries, path, nil
		}
	}
	if lastErr != nil {
		return nil, "", lastErr
	}
	return nil, "", fmt.Errorf("nenhuma decklist encontrada em %s (procurei deck.md, lista.txt, 00-briefing.md)", dir)
}

func parseFile(path string) ([]Entry, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []Entry
	seen := map[string]bool{}
	section := ""
	inFence := false

	add := func(name, sec string, qty int) {
		name = cleanName(name)
		// Básicos agrupados: "Plains ×6" traz a quantidade no fim do nome.
		if m := reSuffixQty.FindStringSubmatch(name); m != nil {
			if n, err := strconv.Atoi(m[2]); err == nil {
				name, qty = strings.TrimSpace(m[1]), n
			}
		}
		if name == "" || isNoise(name) {
			return
		}
		key := strings.ToLower(name)
		if seen[key] {
			return
		}
		seen[key] = true
		out = append(out, Entry{Qty: qty, Name: name, Section: sec})
	}

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64<<10), 4<<20)
	for sc.Scan() {
		line := strings.TrimRight(sc.Text(), "\r")
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "```") {
			inFence = !inFence
			continue
		}
		if trimmed == "" {
			continue
		}
		if m := reHeading.FindStringSubmatch(trimmed); m != nil {
			section = stripCount(m[1])
			continue
		}

		// Tabela markdown: a primeira coluna é o nome da carta.
		if strings.HasPrefix(trimmed, "|") {
			cells := splitRow(trimmed)
			if len(cells) == 0 {
				continue
			}
			first := cells[0]
			if isSeparator(first) || strings.EqualFold(first, "carta") || strings.EqualFold(first, "sai") {
				continue
			}
			add(first, section, 1)
			continue
		}

		// Linha "N Nome" (exportação MTGO) ou lista com marcador.
		body := strings.TrimLeft(trimmed, "-*• \t")
		if m := reQty.FindStringSubmatch(body); m != nil {
			qty := 1
			fmt.Sscanf(m[1], "%d", &qty)
			add(m[2], section, qty)
			continue
		}
		// Dentro de bloco de código sem quantidade, aceita o nome puro.
		if inFence {
			add(body, section, 1)
		}
	}
	return out, sc.Err()
}

func splitRow(line string) []string {
	line = strings.Trim(line, "|")
	parts := strings.Split(line, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	// Descarta linhas totalmente vazias.
	for _, p := range parts {
		if p != "" {
			return parts
		}
	}
	return nil
}

func isSeparator(s string) bool {
	if s == "" {
		return true
	}
	for _, r := range s {
		if r != '-' && r != ':' && r != ' ' {
			return false
		}
	}
	return true
}

var (
	reMarkup  = regexp.MustCompile(`[*_` + "`" + `]+`)
	reLink    = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
	reTrailer = regexp.MustCompile(`\s*\((?:[^()]*)\)\s*$`)
	reCount   = regexp.MustCompile(`\s*\(\d+\)\s*$`)
)

// cleanName remove marcação markdown, emojis de anotação e sufixos como "(22)".
func cleanName(s string) string {
	s = reLink.ReplaceAllString(s, "$1")
	s = reMarkup.ReplaceAllString(s, "")
	s = strings.Map(func(r rune) rune {
		if r == '🔀' || r == '✅' || r == '❌' || r == '⚠' || r == '️' {
			return -1
		}
		return r
	}, s)
	s = strings.TrimSpace(s)
	// "Thousand Moons Smithy // Barracks of the Thousand" → primeira face
	if i := strings.Index(s, " // "); i > 0 {
		s = s[:i]
	}
	s = strings.TrimSpace(reTrailer.ReplaceAllString(s, ""))
	return s
}

func stripCount(s string) string {
	s = reMarkup.ReplaceAllString(s, "")
	return strings.TrimSpace(reCount.ReplaceAllString(s, ""))
}

// isNoise descarta cabeçalhos e rótulos que aparecem em tabelas de análise mas
// não são cartas.
func isNoise(s string) bool {
	l := strings.ToLower(s)
	switch l {
	case "carta", "cartas", "nome", "total", "sai", "entra", "categoria", "métrica",
		"metrica", "antes", "depois", "meta", "fase", "data", "ação", "acao", "—", "-":
		return true
	}
	if strings.HasPrefix(l, "total ") || strings.HasPrefix(l, "soma ") {
		return true
	}
	// Números puros (contagens de tabela).
	if strings.IndexFunc(l, func(r rune) bool { return r < '0' || r > '9' }) == -1 {
		return true
	}
	return false
}
