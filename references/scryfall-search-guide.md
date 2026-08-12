# Guia Central de Buscas no Scryfall

Todos os subagentes usam este guia como fonte única de verdade para consultar cartas. Ele substitui o "agente de busca centralizado": a lógica de consulta mora aqui.

## Ferramentas MCP disponíveis

| Ferramenta | Uso |
|---|---|
| `mcp__scryfall__search_cards` | Busca full-text com sintaxe Scryfall (`query`) |
| `mcp__scryfall__get_card_by_name` | Dados completos de uma carta pelo nome exato em inglês |
| `mcp__scryfall__get_card_by_id` | Dados por Scryfall ID |
| `mcp__scryfall__get_rulings` | Rulings oficiais (por Scryfall/Oracle ID) |
| `mcp__scryfall__get_prices_by_name` / `get_prices_by_id` | Preços (usd, eur, tix) |
| `mcp__scryfall__random_card` | Carta aleatória (inspiração) |

## ⚠️ Controle de volume (crítico)

`search_cards` retorna **a primeira página inteira do Scryfall (até 175 cartas em JSON completo, ~1 MB)**. Resultados grandes estouram o contexto e são salvos em arquivo pelo harness.

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
| `usd<` | Preço em dólar — **filtro de busca apenas, não é régua de orçamento** (ver abaixo) | `usd<5` |
| `produces:` | Produz mana de | `produces:rg t:land` |
| `is:dual`, `is:fetchland`, `is:bounceland` | Classes de terrenos | |
| `order:edhrec` | Ordena por popularidade EDHREC | |
| `-o:"..."` | Excluir texto | |

Combine com espaço = AND; `(a OR b)` = OU.

## Tags funcionais confirmadas (testadas na API)

Funcionam: `otag:ramp`, `otag:mana-rock`, `otag:card-advantage`, `otag:draw`, `otag:removal`, `otag:boardwipe` (= `otag:sweeper`), `otag:counterspell`, `otag:protection`, `otag:tutor`, `otag:burn`, `otag:extra-turn`.

**Não existem**: `otag:wincon`, `otag:finisher`, `otag:token-generator` — para esses, use receitas de `o:` abaixo. Se uma tag retornar erro, caia para a busca `o:` equivalente.

## Receitas por categoria

Em todas, prefixe com `id<=<identidade> legal:commander order:edhrec` (+ `usd<X` se houver orçamento).

## ⚠️ Preço: `usd<` peneira, LigaMagic decide

O `usd<X` da Scryfall serve **só para reduzir o volume de candidatas na busca**. Ele **não** determina se uma carta ou um deck cabem no orçamento — a régua é o **menor valor da LigaMagic** (regra 2 do `CLAUDE.md`).

- Consulta: `https://www.ligamagic.com.br/?view=cards/card&card=<Nome+Em+Ingles>`
- Use o **primeiro** dos três números de "Preço Médio de Venda no Marketplace" (menor / médio / maior). Confira também a linha Foil — às vezes o foil é mais barato que o normal.
- A página monta o preço por JS: **WebFetch não pega** (retorna só o gif de loading). Use as ferramentas de browser (`claude-in-chrome`): navegar, esperar ~2,5s, ler o texto da página.
- Se o usuário mantiver o deck cadastrado na LigaMagic (`?view=dks/deck&id=<id>`), a página do deck já traz o preço carta a carta e o total — muito mais rápido que consultar uma a uma.

**Por que isso importa** (medições reais de 2026-08-12, conversão de proxy US$ × 5,5):

| Carta | Proxy Scryfall | LigaMagic (menor) | Erro |
|---|---|---|---|
| Restoration Magic | ~R$ 1,65 | R$ 10,75 | 6,5× para mais |
| Loran's Escape | ~R$ 2,75 | R$ 15,45 | 5,6× para mais |
| Invisible Force Field | ~R$ 2,42 | R$ 10,94 | 4,5× para mais |
| Chief of the Foundry | ~R$ 1,16 | R$ 0,08 | 14× para menos |
| Reckoner Bankbuster | ~R$ 3,63 | R$ 1,90 | 1,9× para menos |

Os erros vão para os dois lados — **não existe fator de correção**. Ao apresentar qualquer total, rotule a origem: `estimativa (Scryfall)` ou `LigaMagic (menor)`.

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

- Confirme detalhes de carta individual com `get_card_by_name` (barato) em vez de nova busca.
- Anote sempre: nome exato, custo de mana, CMC, tipo e por que sinergiza (2+ pontos).
- Rate limit: o Scryfall pede ~100 ms entre chamadas; não dispare buscas em rajada.
