# MTG Commander Pipeline

Pipeline de subagentes do Claude Code para **construir e otimizar decks de Commander (EDH)**, seguindo o processo de 7 passos: comandante → cartas temáticas → draw → ramp → interação → terrenos e cortes → condições de vitória e goldfishing. As buscas de cartas usam o [MCP do Scryfall](https://mcpmarket.com/server/scryfall) rodando localmente via Docker.

## Pré-requisitos

- Docker com a imagem do MCP: `docker pull mcp/scryfall`
- Claude Code

## Uso

```bash
cd ~/mtg-commander-pipeline
claude
```

Na primeira execução, aprove o servidor MCP `scryfall` do projeto. Depois:

- `/build-deck Krenko, Mob Boss` — constrói um deck novo (aceita comandante ou só um tema, ex.: `/build-deck goblins agressivo`)
- `/improve-deck meu-deck.txt` — audita e otimiza uma decklist existente (formato `1 Nome da Carta` por linha)

O orquestrador conduz você por checkpoints em cada fase — nenhuma carta entra no deck sem sua aprovação. Todo o estado fica em `decks/<nome-do-comandante>/`, incluindo o relatório final (`report.md`) com a lista de exportação padrão MTG Online. Se a sessão cair, rode o comando de novo: ele retoma da última fase concluída.

## Arquitetura

```
/build-deck | /improve-deck  (orquestrador = sessão principal)
        │  delega via Agent tool, 1 fase por vez, checkpoint com o usuário entre fases
        ▼
┌─────────────────────────────────────────────────────────────┐
│ commander-scout → theme-analyst → draw-specialist →          │
│ ramp-specialist → interaction-specialist →                   │
│ manabase-engineer → wincon-tester                            │
└─────────────────────────────────────────────────────────────┘
        │  cada especialista consulta o Scryfall diretamente,
        ▼  seguindo o guia central de buscas
bin/mtgdb                               ←  dados de carta locais (SQLite + bulk Scryfall)
references/scryfall-search-guide.md     ←  lógica de consulta centralizada
references/card-evaluation-checklist.md ←  como julgar uma carta (ficha F1–F7, corte, histórico)
decks/<slug>/*.md                       ←  estado compartilhado entre agentes
```

## Começando

```bash
make db     # compila bin/mtgdb e constrói o banco (~15 s, ~103 MB)
make help   # lista os alvos
```

Depois, no Claude Code: `/build-deck <tema ou comandante>` ou `/improve-deck <caminho da decklist>`.

Duas adaptações em relação ao desenho original de "agente de busca centralizado":

1. **Subagentes não podem invocar outros subagentes** no Claude Code — a centralização da busca virou o guia `references/scryfall-search-guide.md`, que todos os especialistas leem antes de consultar o MCP (mesmo benefício: lógica de busca em um único lugar).
2. **O orquestrador não pode ser um subagente** (subagentes não interagem com o usuário) — ele é executado pela sessão principal via comandos `/build-deck` e `/improve-deck`.

### O custo de ter especialistas por categoria

Dividir o pipeline por função (draw, ramp, interação…) dá profundidade, mas cria um viés: **cada agente julga a carta pela sua própria lente**, e num deck de Commander a mesma carta costuma exercer 3–5 funções. Uma peça que parece fraca em "draw" pode ser o corpo que paga um custo de tap, a contagem que liga um affinity e o redutor que segura a curva. Cortada por uma lente, ela volta na rodada seguinte por outra — vai-e-vem que parece mudança de estratégia mas é análise parcial.

`references/card-evaluation-checklist.md` existe para fechar essa brecha: ficha de funções obrigatória antes de qualquer veredito, protocolo de corte com cobertura declarada, simetria de critério entre os dois lados da troca, e `decks/<slug>/decisions.md` como registro contra reposições sem justificativa. Os especialistas ficam proibidos de cortar carta cujo valor está fora da sua especialidade — devolvem a decisão ao orquestrador, que é quem vê o deck inteiro.

### mtgdb — dados de carta locais

Uma checklist mais rigorosa só se cumpre se levantar os fatos for barato. `bin/mtgdb` (Go + SQLite, ver `references/mtgdb.md`) mantém o **bulk data do Scryfall** local — 38 mil cartas, 4,5 mil tags do Tagger e 78 mil rulings — e responde sem rede:

```bash
mtgdb oracle "Solar Array"        # ficha F2–F7 + texto oracle
mtgdb deck <slug>                 # o deck inteiro, com os agregados da checklist
mtgdb search "charge counter" -id WUR -type Artifact
mtgdb tag -list proliferate       # tags curadas, as mesmas por trás de otag:
mtgdb rulings "Deepglow Skate"
```

`mtgdb deck` separa **criaturas de verdade** de **veículos e spacecraft** — a distinção que já custou uma rodada inteira de análise errada, porque veículo não é corpo, é consumidor de corpo.

Sobre preço: cotação **não** tem validade previsível, então `data/prices.tsv` é uma série **append-only** de observações datadas. `mtgdb prices -stale 30` mostra a idade de cada uma, e `-volatile` mostra quais cartas de fato oscilam — as únicas que precisam ser reconferidas antes de um torneio.

O banco (`data/scryfall.db`) é derivado e **não versionado**: `make db` reconstrói em ~15 s. Já `data/collection.tsv` e `data/prices.tsv` são versionados, porque são dados do usuário e não saem de fonte pública — cada linha de preço custou uma navegação manual na LigaMagic.

## Permissões (opcional)

Para não receber prompts de permissão do MCP, crie `.claude/settings.local.json`:

```json
{
  "enableAllProjectMcpServers": true,
  "permissions": {
    "allow": ["mcp__scryfall"]
  }
}
```

## Estrutura

```
.claude/agents/      # 7 subagentes especialistas
.claude/commands/    # /build-deck e /improve-deck (orquestradores)
references/          # guia de buscas Scryfall + checklist de avaliação + template de relatório
decks/               # um diretório por deck (estado + decisions.md + relatório final)
CLAUDE.md            # regras do pipeline (metas, formato de estado, convenções)
```
