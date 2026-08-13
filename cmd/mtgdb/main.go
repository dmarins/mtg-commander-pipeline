// Command mtgdb é o acesso local do pipeline aos dados de carta.
//
// Ele mantém um SQLite construído a partir do bulk data do Scryfall (cartas,
// tags do Tagger e rulings oficiais) e responde as consultas que as fases de
// análise fazem — sem rede e sem gastar uma requisição por carta.
//
// O banco é derivado e descartável (`mtgdb build` reconstrói). O que é do
// usuário — coleção e observações de preço — vive em data/*.tsv, versionado.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const usage = `mtgdb — dados locais de carta para o pipeline de Commander

USO
  mtgdb <comando> [argumentos]

COMANDOS
  build                    baixa o bulk data do Scryfall e (re)constrói o banco
  oracle <nomes...>        ficha F1–F7 das cartas (substitui o curl por carta)
  deck <slug>              ficha do deck inteiro + agregados da checklist
  search <termo>           busca full-text no texto oracle (FTS5)
  tag <slug-da-tag>        cartas com uma tag curada do Scryfall Tagger
  rulings <nome>           rulings oficiais da carta
  prices <nomes...>        cotações conhecidas, com a idade de cada uma
  collection <nomes...>    verifica o que já está na coleção
  status                   estado do banco e datas dos dumps

Rode 'mtgdb <comando> -h' para as opções de cada um.

Os nomes de carta são sempre em inglês (regra 8 do CLAUDE.md). Nomes que não
resolverem são reportados — nunca substituídos por um palpite silencioso.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(2)
	}

	cmd := os.Args[1]
	args := reorderArgs(os.Args[2:])

	var err error
	switch cmd {
	case "build":
		err = cmdBuild(args)
	case "oracle":
		err = cmdOracle(args)
	case "deck":
		err = cmdDeck(args)
	case "search":
		err = cmdSearch(args)
	case "tag":
		err = cmdTag(args)
	case "rulings":
		err = cmdRulings(args)
	case "prices":
		err = cmdPrices(args)
	case "collection":
		err = cmdCollection(args)
	case "status":
		err = cmdStatus(args)
	case "-h", "--help", "help":
		fmt.Println(usage)
		return
	default:
		fmt.Fprintf(os.Stderr, "comando desconhecido: %s\n\n%s\n", cmd, usage)
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
}

// reorderArgs move as flags para a frente dos argumentos posicionais.
//
// O pacote flag para de interpretar assim que encontra o primeiro não-flag, o
// que quebraria usos naturais como `mtgdb deck <slug> -missing` ou
// `mtgdb prices "Sol Ring" -history`. Reordenar é mais amigável que exigir que
// o usuário lembre da ordem.
//
// Flags que consomem valor (`-limit 20`, `-kind menor`) precisam levar o valor
// junto; a lista abaixo é a das flags com argumento separado que os comandos
// declaram. As booleanas nunca consomem o próximo token.
func reorderArgs(args []string) []string {
	takesValue := map[string]bool{
		"-limit": true, "-id": true, "-type": true, "-cmc-max": true, "-cmc-min": true,
		"-stale": true, "-source": true, "-kind": true, "-date": true, "-note": true,
	}
	var flags, positional []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			positional = append(positional, args[i+1:]...)
			break
		}
		if strings.HasPrefix(a, "-") && len(a) > 1 {
			flags = append(flags, a)
			// -limit=20 já carrega o valor; -limit 20 precisa do próximo token.
			base := a
			if j := strings.Index(a, "="); j >= 0 {
				base = a[:j]
			} else if takesValue[a] && i+1 < len(args) {
				i++
				flags = append(flags, args[i])
			}
			_ = base
			continue
		}
		positional = append(positional, a)
	}
	return append(flags, positional...)
}

// root localiza a raiz do repositório subindo até achar go.mod.
// MTGDB_ROOT tem precedência, para uso fora da árvore.
func root() string {
	if v := os.Getenv("MTGDB_ROOT"); v != "" {
		return v
	}
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "."
		}
		dir = parent
	}
}

func dbPath() string   { return filepath.Join(root(), "data", "scryfall.db") }
func dataDir() string  { return filepath.Join(root(), "data") }
func decksDir() string { return filepath.Join(root(), "decks") }

// truncate encurta um texto para caber numa coluna, sem cortar no meio de um rune.
func truncate(s string, n int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	if n <= 1 {
		return string(r[:n])
	}
	return string(r[:n-1]) + "…"
}
