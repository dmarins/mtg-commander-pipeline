# mtgdb — dados locais de carta

Binário Go que mantém um SQLite construído a partir do **bulk data do Scryfall** e responde as consultas do pipeline sem rede.

**Use `mtgdb` antes de qualquer consulta ao MCP do Scryfall.** O MCP continua útil para o que o bulk não cobre (cartas lançadas depois do último dump, preços em USD), mas ficha de carta, busca por texto, tags e rulings saem daqui — instantâneos e sem gastar uma requisição por carta.

## Primeiro uso

```bash
make db        # compila o binário e constrói o banco (~15 s, ~103 MB)
make refresh   # rebaixa o bulk e reconstrói do zero
make status    # estado do banco e data dos dumps
```

O binário fica em `bin/mtgdb`. O banco fica em `data/scryfall.db` e **não é versionado** — é derivado e reconstruível. Se `mtgdb` reclamar de versão de schema, rode `make refresh`.

## Comandos

### `oracle` — ficha de uma ou mais cartas

Substitui o `curl` carta a carta. Imprime o texto oracle completo e os eixos F2–F7 do checklist.

```bash
mtgdb oracle "Solar Array" "Moxite Refinery"
mtgdb oracle -short "Sol Ring" "Arcane Signet"    # uma linha por carta
mtgdb oracle -rulings "Deepglow Skate"            # inclui rulings oficiais
mtgdb oracle -json "Emissary Escort"              # saída estruturada
```

### `deck` — ficha do deck inteiro

Lê `decks/<slug>/deck.md` (ou `lista.txt`, ou `00-briefing.md`), cruza com o banco e devolve os agregados que o checklist pede + a ficha por carta.

```bash
mtgdb deck inspirit-flagship-vessel-v2
mtgdb deck inspirit-flagship-vessel-v2 -full      # + texto oracle de cada carta
mtgdb deck inspirit-flagship-vessel-v2 -missing   # só o que não resolveu
```

Os agregados incluem a distinção que já custou uma rodada de análise errada: **criaturas de verdade** separadas de **veículos e spacecraft**, que não são corpos — são consumidores de corpo.

### `search` — busca full-text (FTS5) no texto oracle

```bash
mtgdb search "charge counter" -id WUR -type Artifact -cmc-max 3
mtgdb search "proliferate" -id WUR
mtgdb search '"enters tapped"' -limit 20          # frase exata
```

Sintaxe FTS5: `AND`, `OR`, `NOT`, `"frase exata"`, `prefixo*`.
Filtros: `-id` (identidade de cor, como `id<=` do Scryfall), `-type`, `-cmc-min`, `-cmc-max`, `-limit`, `-commander=false`.

### `tag` — tags curadas do Scryfall Tagger

As mesmas por trás de `otag:` na busca do site. São curadas por humanos que entendem o jogo — melhores que qualquer busca por palavra-chave.

```bash
mtgdb tag -list proliferate      # descobre o slug certo
mtgdb tag ramp -id WUR -limit 30
```

Nem todo termo é um slug: `proliferate` não existe, mas `pseudo-proliferate` e `synergy-proliferate` sim. Sempre confirme com `-list` antes de concluir que "não há cartas".

### `rulings` — decisões oficiais

```bash
mtgdb rulings "Deepglow Skate"
```

Nem toda carta tem ruling. Quando não tiver, o comando diz isso explicitamente — não invente a regra nem conclua que a interpretação está livre. Vá às Comprehensive Rules ou a uma fonte de juiz.

### `prices` — cotações datadas

Preço **não** tem validade previsível, então o banco guarda uma **série append-only**: cada linha é uma observação com data, nunca um valor que se sobrescreve.

```bash
mtgdb prices "Loran's Escape" "Cyberdrive Awakener"
mtgdb prices "Loran's Escape" -history        # todas as observações
mtgdb prices -stale 30                        # cotações com mais de 30 dias
mtgdb prices -volatile                        # o que de fato oscila
mtgdb prices -add "Loran's Escape" 12,90      # registra nova observação
```

`-volatile` é o comando que muda a rotina pré-torneio: em vez de reconferir 100 cartas, ele mostra as poucas que historicamente se mexem — que são justamente as que podem empurrar o total para fora do teto.

**A captura continua manual.** A página da LigaMagic monta o preço por JS, então é navegador carta a carta. O `mtgdb` evita *perder* e evita *recapturar* — não substitui a primeira captura.

### `collection` — o que o usuário já possui

```bash
mtgdb collection "Emissary Escort" "Sol Ring"
mtgdb collection -list
mtgdb collection -add "Titan Forge" -note "maybeboard"
```

Torna a regra 7 do `CLAUDE.md` ("coleção pessoal primeiro") uma consulta em vez de leitura de arquivo.

## Resolução de nomes

Nomes de carta são sempre em inglês (regra 8). A resolução tenta, em ordem: nome exato → normalizado (sem acento, sem apóstrofo, sem pontuação) → face de carta dupla → prefixo → busca textual. A coluna `how` diz por onde passou; qualquer coisa diferente de `exact` merece uma conferida.

**Nome que não resolve é reportado, nunca substituído por palpite.** Ambiguidade também: `mtgdb` lista os candidatos e não escolhe por você.

## O que é derivado e o que é seu

| Caminho | Versionado? | Por quê |
|---|---|---|
| `data/scryfall.db` | ❌ | derivado do bulk público; cada rebuild gera um blob binário novo e merge de binário é irresolúvel |
| `data/collection.tsv` | ✅ | o que você possui — não reconstruível |
| `data/prices.tsv` | ✅ | série de cotações; cada linha custou uma navegação manual |

Perder o banco não custa nada (`make db` reconstrói). Perder os TSV custa trabalho real — por isso eles são texto, com diff legível e merge por linha.

## Manutenção

O Scryfall atualiza os dumps diariamente. Cartas de sets muito recentes podem faltar até o próximo `make refresh` — nesse caso, e só nesse, caia para o MCP do Scryfall.

`SchemaVersion` em `internal/store/store.go` deve ser incrementada sempre que o schema **ou a função `Normalize`** mudar: os nomes normalizados ficam materializados no build, e uma mudança silenciosa faria cartas sumirem das buscas sem erro aparente.
