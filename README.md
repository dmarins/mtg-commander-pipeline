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
references/scryfall-search-guide.md     ←  lógica de consulta centralizada
references/card-evaluation-checklist.md ←  como julgar uma carta (ficha F1–F7, corte, histórico)
decks/<slug>/*.md                       ←  estado compartilhado entre agentes
```

Duas adaptações em relação ao desenho original de "agente de busca centralizado":

1. **Subagentes não podem invocar outros subagentes** no Claude Code — a centralização da busca virou o guia `references/scryfall-search-guide.md`, que todos os especialistas leem antes de consultar o MCP (mesmo benefício: lógica de busca em um único lugar).
2. **O orquestrador não pode ser um subagente** (subagentes não interagem com o usuário) — ele é executado pela sessão principal via comandos `/build-deck` e `/improve-deck`.

### O custo de ter especialistas por categoria

Dividir o pipeline por função (draw, ramp, interação…) dá profundidade, mas cria um viés: **cada agente julga a carta pela sua própria lente**, e num deck de Commander a mesma carta costuma exercer 3–5 funções. Uma peça que parece fraca em "draw" pode ser o corpo que paga um custo de tap, a contagem que liga um affinity e o redutor que segura a curva. Cortada por uma lente, ela volta na rodada seguinte por outra — vai-e-vem que parece mudança de estratégia mas é análise parcial.

`references/card-evaluation-checklist.md` existe para fechar essa brecha: ficha de funções obrigatória antes de qualquer veredito, protocolo de corte com cobertura declarada, simetria de critério entre os dois lados da troca, e `decks/<slug>/decisions.md` como registro contra reposições sem justificativa. Os especialistas ficam proibidos de cortar carta cujo valor está fora da sua especialidade — devolvem a decisão ao orquestrador, que é quem vê o deck inteiro.

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
