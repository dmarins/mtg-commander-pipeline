# Guia Central de Buscas no Scryfall

Todos os subagentes usam este guia como fonte única de verdade para consultar cartas. Ele substitui o "agente de busca centralizado": a lógica de consulta mora aqui.

⚠️ **Scryfall é repositório de cartas, não de preço.** Daqui saem oracle, tipo, CMC, cores, tags e rulings — nunca valor. Preço vem do **menor valor da LigaMagic** (regra 2 do `CLAUDE.md`), e nem o filtro `usd<` entra nas queries. Ver "Preço não se busca aqui" abaixo.

## Antes de tudo: o banco local

**Consulte `bin/mtgdb` primeiro** — ver `references/mtgdb.md`. Ele tem o bulk data do Scryfall em SQLite (cartas, tags do Tagger, rulings), responde sem rede e não gasta uma requisição por carta:

| Em vez de | Use |
|---|---|
| `get_card_details` por carta | `mtgdb oracle "<nome>" "<nome>" ...` |
| ler a decklist carta a carta | `mtgdb deck <slug>` |
| `search_cards` com `o:<termo>` | `mtgdb search "<termo>" -id <cores>` |
| `search_cards` com `otag:<tag>` | `mtgdb tag <slug> -id <cores>` |
| `get_card_rulings` | `mtgdb rulings "<nome>"` |

Recorra ao MCP quando o banco não bastar: **carta mais nova que o último dump**, ou uma consulta que precise da sintaxe completa do Scryfall. Se o banco não existir, `make db` (~15 s).

## Ferramentas MCP disponíveis

Servidor `mtg` ([nathanmartins/mtg-mcp](https://github.com/nathanmartins/mtg-mcp), binário `mtg-mcp`).

Estado conferido em 2026-09-23 contra o `mtg-mcp` v2.2.1. Ferramenta marcada como quebrada **não se chama** —
é requisição perdida e resposta que parece dado. Ao atualizar o binário, reteste e atualize esta tabela.

| Ferramenta | Uso | Estado |
|---|---|---|
| `mcp__mtg__search_cards` | Busca com sintaxe Scryfall (`query`); até 50 resultados em texto compacto | ok |
| `mcp__mtg__get_card_details` | Dados completos de uma carta pelo nome exato em inglês | ok |
| `mcp__mtg__get_card_rulings` | Rulings oficiais | ok |
| `mcp__mtg__check_commander_legality` | Legalidade de **uma** carta em Commander | ok |
| `mcp__mtg__get_banned_list` | Lista de banidas do Commander, ao vivo | ok |
| `mcp__mtg__validate_deck` | Contagem de 100 cartas e singleton | **parcial** — apesar da descrição, **não** confere identidade de cor nem banimento (testado: `Counterspell` e `Mana Crypt` num deck mono-R passaram) |
| `mcp__mtg__get_edhrec_recommendations` | Cartas do EDHREC para um comandante, por seção | ok, **inclusão quebrada** (ver abaixo) |
| `mcp__mtg__search_archidekt_decks` | Listas públicas por comandante, ordenadas por visualizações | ok, com ressalvas (ver abaixo) |
| `mcp__mtg__get_archidekt_deck` | Lista completa de um deck do Archidekt (URL ou ID); `lands_only: true` traz só os terrenos | ok |
| `mcp__mtg__get_archidekt_user_decks` | Decks públicos de um usuário do Archidekt | não testada |
| `mcp__mtg__get_rule` / `search_rules` | Comprehensive Rules, por número ou palavra-chave | ok |
| `mcp__mtg__get_glossary_term` | Glossário das regras | não testada |
| ~~`mcp__mtg__get_edhrec_combos`~~ | Combos por identidade | **quebrada** — HTTP 403 |
| ~~`mcp__mtg__search_moxfield_decks`~~ / ~~`get_moxfield_deck`~~ / ~~`get_moxfield_user_decks`~~ | Moxfield | **quebradas** — HTTP 404 em todas |
| ~~`mcp__mtg__get_card_price`~~ | USD do Scryfall convertido por câmbio | **fora de uso e bloqueada** — viola a regra 2 |
| `mcp__mtg__get_card_image` | Imagem da carta | sem uso no pipeline |

## Fontes de meta: EDHREC e Archidekt

EDHREC e listas públicas mostram o que **outros jogadores** põem no deck — popularidade e correlação
estatística, não análise do plano deste deck. Servem para **descobrir** candidatas que a busca por termo não
achou e para **calibrar** números (terrenos, contagens) contra decks reais do mesmo bracket. Nunca para decidir.

**Regras comuns a toda fonte de meta:**

- **Mesmo filtro de qualquer candidata.** Carta vinda daqui entra só com oracle puxado na hora
  (`bin/mtgdb oracle`), ficha F1–F7 e **2+ pontos de sinergia** justificados por você (regra 3). "Sinergia alta
  no EDHREC", "X decks usam" ou "está na lista mais vista do Archidekt" **não é** ponto de sinergia — é o
  motivo de você ter olhado a carta.
- **O plano do briefing manda.** A meta reflete o arquétipo médio do comandante; se o briefing declara outro
  eixo, a carta popular que serve ao eixo errado fica de fora, com o motivo por escrito.
- **Preço não vem daqui.** Campo de preço em retorno de EDHREC ou Archidekt (inclusive "$100 budget" no nome
  do deck) é ignorado (regra 2). Cotação só por `mtgdb prices`; o que faltar vai como `a cotar`.
- **Regra 11 vale.** Combo ou peça de lock/stax/MLD que apareça na meta só entra se o briefing pedir.
- **Rastreie a origem.** Na tabela de candidatas, cite na coluna de sinergias quando a carta veio de fonte de
  meta (`origem: EDHREC alta sinergia`, `origem: Archidekt #3047743`). O orquestrador precisa saber o que
  veio de estatística e o que veio de busca por termo.

### EDHREC — `get_edhrec_recommendations`

- **Uma chamada por comandante, por rodada, e quem faz é o `theme-analyst`.** Ele grava o retorno na seção
  `Radar do EDHREC` do `02-theme.md`; os especialistas seguintes **leem de lá**. Só chame de novo se a seção
  não existir (rodada antiga) ou se faltar uma seção do retorno que a sua fase precisa. O EDHREC pede
  ~1 requisição/segundo.
- Use `limit` 15–20: com o padrão (10) cada seção corta em 5 e esconde o resto (`...and 45 more cards`).
- **Seções do retorno e quem as lê:**

  | Seção | Fase |
  |---|---|
  | `High Synergy Cards`, `New Cards`, `Top Cards`, `Creatures`, `Enchantments` | tema (02) |
  | `Instants`, `Sorceries`, `Utility Artifacts` | draw (03) e interação (05) — as seções são por **tipo**, não por função: classifique pela oracle |
  | `Mana Artifacts` | ramp (04) |
  | `Lands`, `Utility Lands` | manabase (06) |
  | `Game Changers` | orquestrador e wincons (07) — cota do bracket |

- **Leia a nota de sinergia, não a inclusão.** Na v2.2.1 a inclusão volta quebrada (`Total Decks: 0`,
  `0 decks (NaN%)`); a nota `Synergy` vem correta. Não cite nem conclua nada a partir de inclusão — trate
  como ausente. Sinergia negativa (`Sol Ring`, `Arcane Signet`) quer dizer "genérica, não específica deste
  comandante", não "ruim".
- **Game Changers contam para o bracket.** Marque-as (`game changer`) para o orquestrador controlar a cota
  declarada no briefing.

### Archidekt — `search_archidekt_decks` + `get_archidekt_deck`

- **Filtre por `bracket`** igual ao power level do briefing. Comparar com lista de bracket 4 empurra o deck
  para cima sem ninguém ter pedido.
- **`limit` é ignorado:** a busca sempre devolve ~60 decks (lista curta, mas ocupa contexto). Faça uma busca
  por comandante e escolha **2–3 listas** para abrir.
- **A busca é aproximada:** para comandantes com nome parecido ou partes com vários comandantes, aparecem decks
  de outro comandante (buscando Otharri veio `Esika`, `Niv`, `Kellan`). Confirme o comandante no cabeçalho do
  `get_archidekt_deck` antes de usar a lista.
- **`lands_only: true`** traz só os terrenos — é a chamada da manabase. Atenção à contagem: o cabeçalho
  (`Lands (7)`) conta **linhas**, não cartas; `29x Mountain` é uma linha. Some as quantidades.
- Lista pública é **referência de calibragem**, não molde: a meta de 38 terrenos, a fórmula e as metas de
  categoria do pipeline continuam valendo; a comparação só entra no relatório como contexto.

### Moxfield

Quebrado na v2.2.1 (404 em busca, deck e usuário). Se o usuário passar um link do Moxfield, peça a exportação
em texto (`Export → Copy for MTGO`) — não tente o MCP nem WebFetch.

### Validação do deck

`validate_deck` cobre **só** contagem e singleton. A checagem completa da revisão final é:

1. `validate_deck` — 100 cartas, sem duplicatas além de básicos;
2. `get_banned_list` — cruze com a lista do deck;
3. identidade de cor — `bin/mtgdb oracle <nomes...>` (linha `identidade`) contra a do comandante; o
   `mtgdb deck` não mostra identidade.

## ⚠️ Controle de volume (crítico)

`search_cards` retorna **até 50 cartas** em texto (nome, custo, tipo, oracle, set, legalidade). Mais leve que o JSON completo, mas 50 oracles ainda pesam no contexto — mantenha as queries estreitas.

Regras:

1. **Mire em consultas que retornem ≤ 40 cartas.** Restrinja sempre por: identidade (`id<=`), custo (`mv<=N` ou faixas `mv=2`, `mv=3`...), tipo (`t:` / `-t:`), e tag ou texto de oracle.
2. **Sempre acrescente `order:edhrec`** — ordena por popularidade no EDHREC; mesmo se o resultado for truncado, o topo é o que interessa.
3. **Divida buscas largas em fatias** (por CMC, por tipo) em vez de uma busca gigante.
4. Se o resultado for salvo em arquivo (mensagem "exceeds maximum allowed tokens"), **não leia o arquivo inteiro**. Extraia só o essencial com python3 (não há `jq` nesta máquina):

```bash
python3 -c "
import json
d = json.load(open('CAMINHO_DO_ARQUIVO'))
print('TOTAL:', d['total_cards'])
for c in d['data'][:40]:
    print(f\"{c['name']} | {c.get('mana_cost','')} | mv{c['cmc']:.0f} | {c['type_line']} | {(c.get('oracle_text') or '').replace(chr(10),' // ')[:200]}\")
"
```

5. Alternativa quando precisar de mais controle (página 2, campos específicos): API direta com `curl` + python3:

```bash
curl -s --get "https://api.scryfall.com/cards/search" \
  --data-urlencode "q=QUERY AQUI" --data-urlencode "order=edhrec" | python3 -c "..."
```

## Sintaxe essencial

| Operador | Significado | Exemplo |
|---|---|---|
| `id<=wub` | Identidade de cor dentro de Esper (**use sempre**) | `id<=bg` |
| `legal:commander` | Legal no formato (exclui banidas) (**use sempre**) | |
| `c:g` / `c>=rg` | Cor da carta | |
| `t:` / `-t:` | Tipo / excluir tipo | `t:creature`, `-t:land` |
| `o:"texto"` | Texto de oracle (aceita `~` como nome da carta) | `o:"whenever ~ attacks"` |
| `kw:` | Palavra-chave de habilidade | `kw:flying`, `kw:landfall` |
| `mv=` `mv<=` `mv>=` | Valor de mana | `mv<=2` |
| `pow` / `tou` | Poder / resistência | `pow>=4` |
| `is:commander` | Pode ser comandante | `is:commander id<=ur` |
| `otag:` | Tag funcional (Scryfall Tagger) | `otag:ramp` |
| `produces:` | Produz mana de | `produces:rg t:land` |
| `is:dual`, `is:fetchland`, `is:bounceland` | Classes de terrenos | |
| `order:edhrec` | Ordena por popularidade EDHREC | |
| `-o:"..."` | Excluir texto | |

Combine com espaço = AND; `(a OR b)` = OU.

## Tags funcionais confirmadas (testadas na API)

Funcionam: `otag:ramp`, `otag:mana-rock`, `otag:card-advantage`, `otag:draw`, `otag:removal`, `otag:boardwipe` (= `otag:sweeper`), `otag:counterspell`, `otag:protection`, `otag:tutor`, `otag:burn`, `otag:extra-turn`.

**Não existem**: `otag:wincon`, `otag:finisher`, `otag:token-generator` — para esses, use receitas de `o:` abaixo. Se uma tag retornar erro, caia para a busca `o:` equivalente.

## Receitas por categoria

Em todas, prefixe com `id<=<identidade> legal:commander order:edhrec`. **Nunca acrescente `usd<X`** — preço não entra em query (ver abaixo).

## ⚠️ Preço não se busca aqui

**O Scryfall não é fonte de preço neste projeto — em nenhum papel.** Não use `get_card_price` (bloqueada),
não use `usd<` / `eur<` / `tix<` em query, e ignore campos de preço que venham junto de outro retorno. Nem como
estimativa, nem como "ordem de grandeza", nem para peneirar candidatas antes de cotar.

A única cotação válida é o **menor valor da LigaMagic** (regra 2 do `CLAUDE.md`):

- Consulte primeiro o que já foi capturado: `mtgdb prices <nomes...>`.
- O que faltar: `https://www.ligamagic.com.br/?view=cards/card&card=<Nome+Em+Ingles>`
- Use o **primeiro** dos três números de "Preço Médio de Venda no Marketplace" (menor / médio / maior). Confira também a linha Foil — às vezes o foil é mais barato que o normal.
- A página monta o preço por JS: **WebFetch não pega** (retorna só o gif de loading). Use as ferramentas de browser (`claude-in-chrome`): navegar, esperar ~2,5s, ler o texto da página.
- Registre o resultado com `mtgdb prices -add "<carta>" <valor>`.
- Se o usuário mantiver o deck cadastrado na LigaMagic (`?view=dks/deck&id=<id>`), a página do deck já traz o preço carta a carta e o total — muito mais rápido que consultar uma a uma.

**Por que o proxy saiu de vez** (medições reais de 2026-08-12, conversão de proxy US$ × 5,5):

| Carta | Proxy Scryfall | LigaMagic (menor) | Erro |
|---|---|---|---|
| Restoration Magic | ~R$ 1,65 | R$ 10,75 | 6,5× para mais |
| Loran's Escape | ~R$ 2,75 | R$ 15,45 | 5,6× para mais |
| Invisible Force Field | ~R$ 2,42 | R$ 10,94 | 4,5× para mais |
| Chief of the Foundry | ~R$ 1,16 | R$ 0,08 | 14× para menos |
| Reckoner Bankbuster | ~R$ 3,63 | R$ 1,90 | 1,9× para menos |

Os erros vão para os dois lados e **não existe fator de correção**. As duas últimas linhas são o motivo de o filtro
sair também da *busca*: `usd<X` descarta silenciosamente carta que na LigaMagic custa centavos. Sem o filtro,
a triagem por custo acontece **depois** da busca — sobre cotação real, e não sobre proxy.

Ao apresentar qualquer total, rotule origem e idade: `LigaMagic (menor), cotações de <data>`. Carta sem cotação
aparece como `a cotar` — nunca com número estimado.

**Comandantes** (commander-scout):
- Por tema: `is:commander id<=br o:sacrifice`
- Menos conhecidos: adicione `-is:reprint` e revise além do topo do `order:edhrec`

**Temáticas** (theme-analyst) — derive dos termos do comandante:
- `o:"whenever a creature you control dies"`, `kw:landfall`, `o:"enters"` etc.
- Sinergia dupla: combine dois termos — `o:sacrifice o:"draw"`

**Draw** (draw-specialist):
- Geral: `otag:card-advantage -t:land mv<=4`
- Sinérgico: `otag:draw o:<palavra-chave do tema>`
- Evite loot puro (troca sem ganho) a menos que o tema use cemitério.

**Ramp** (ramp-specialist):
- Padrão: `otag:ramp mv<=2` e `mv=3` (fatias separadas)
- Rochas: `otag:mana-rock mv<=3`
- Explosivo: `otag:ramp mv>=4 o:/add .*(three|four|five|six|X)/` ou `o:"double" o:mana`

**Interação** (interaction-specialist):
- Remoção: `otag:removal mv<=3 (t:instant OR t:sorcery)`
- Counters: `otag:counterspell mv<=3`
- Proteção: `otag:protection mv<=2`
- Wipes: `otag:boardwipe mv<=6`

**Terrenos** (manabase-engineer):
- Duais: `t:land produces:bg is:dual` / `is:fetchland id<=bg`
- Utilidade: `t:land otag:card-advantage` ou `t:land o:<tema>`

**Wincons** (wincon-tester):
- Finishers de massa: `o:"creatures you control get" o:"+" mv>=4`
- Dano direto: `otag:burn o:"each opponent"`
- Alt-win: `o:"you win the game"`

## Boas práticas

- Confirme detalhes de carta individual com `get_card_details` (barato) em vez de nova busca.
- Anote sempre: nome exato, custo de mana, CMC, tipo e por que sinergiza (2+ pontos).
- Rate limit: o Scryfall pede ~100 ms entre chamadas e o EDHREC ~1 s; não dispare buscas em rajada.
