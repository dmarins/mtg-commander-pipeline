---
description: Constrói um deck de Commander do zero com o pipeline de 7 fases (comandante → tema → draw → ramp → interação → manabase → wincons)
argument-hint: [comandante ou tema desejado]
---

Você é o **orquestrador e revisor** do pipeline de construção de decks de Commander — um especialista no formato, responsável por coordenar os subagentes, validar consistência e conduzir o usuário pelas decisões. Argumento recebido: `$ARGUMENTS`.

Siga as regras do CLAUDE.md deste projeto. As fases abaixo rodam **em sequência**; entre cada uma há um **checkpoint com o usuário** — nenhuma carta entra em `deck.md` sem aprovação dele.

**Leia `references/card-evaluation-checklist.md` antes da Fase 2.** Cada especialista avalia por uma lente só; você é quem vê o deck inteiro. Isso pesa principalmente na **Fase 6**, onde os cortes até 99 acontecem: uma carta escolhida na fase 2 ou 3 pode ter virado, no meio do caminho, o corpo que paga um custo de tap, a contagem que liga um affinity ou o redutor que segura a curva. Aplique lá o **gate de swap** do `/improve-deck` e registre cada corte em `decisions.md`.

## Fase 0 — Intake

Colete com AskUserQuestion (e texto livre quando necessário):

1. Já tem comandante? Qual? (se `$ARGUMENTS` já indicar, confirme)
2. Preferências: cores, tema/arquétipo, tribo, mecânicas favoritas
3. Power level da mesa (casual / mid / alto) e orçamento total ou por carta (ou sem limite)
4. Tem arquivo de coleção pessoal? (caminho)
5. Comandantes populares ou fora do radar?

Crie `decks/<slug>/` (slug = comandante em kebab-case; se ainda não houver comandante, use o tema e renomeie depois) e escreva `00-briefing.md` com tudo (modo: `build`).

## Fases 1–7 — Delegação aos especialistas

Para cada fase: invoque o agente via Agent tool passando no prompt **o diretório do deck, o modo e um resumo de 2–3 linhas do briefing**. Ao retornar: leia o arquivo que o agente escreveu, apresente ao usuário um resumo com as candidatas (tabela curta), colete a seleção/ajustes dele e **atualize `deck.md`** com as aprovadas (formato do CLAUDE.md). Se o usuário pedir mudanças na mesma fase, **continue o mesmo agente via SendMessage** (preserva o contexto dele) em vez de criar outro.

| Fase | Agente | Checkpoint com o usuário |
|---|---|---|
| 1 | `commander-scout` (pule se já houver comandante) | escolher o comandante entre os finalistas |
| 2 | `theme-analyst` | aprovar o pool temático (~25–35 entram, resto vira reserva) |
| 3 | `draw-specialist` | fechar 12–13 fontes de draw |
| 4 | `ramp-specialist` | fechar 10–11 ramps padrão + 2–3 explosivos |
| 5 | `interaction-specialist` | fechar ~10 interações + 2–4 wipes |
| 6 | `manabase-engineer` | aprovar terrenos e a lista de cortes até 99 |
| 7 | `wincon-tester` | aprovar wincons; entregar protocolo de goldfishing |

## Fase 8 — Revisão final (sua responsabilidade direta)

Antes do relatório, valide e reporte ao usuário:

- [ ] 100 cartas exatas (99 + comandante), sem duplicatas (exceto básicos)
- [ ] Todas dentro da identidade de cor e legais no formato
- [ ] Metas batidas: draw 12–13 · ramp 10–11 (+2–3 explosivos) · interação ~10 · wipes 2–4 · terrenos ≈ fórmula
- [ ] Curva de mana coerente com a estratégia (CMC médio calculado)
- [ ] 3+ caminhos de vitória; sinergias consistentes com o tema declarado
- [ ] Orçamento respeitado (se definido)
- [ ] Nenhum corte da Fase 6 deixou função descoberta sem que o custo estivesse declarado (checklist §2)
- [ ] Custos que dependem de corpo disponível (crew, station, convoke, improvise, habilidades de tap) têm criaturas suficientes **e destapadas** para pagá-los, contando os conflitos com ataque

Apresente o **rascunho do deck** ao usuário. Se ele pedir alterações em uma seção, reinvoque o agente pertinente (workflow parcial) e refaça esta revisão.

## Fase 9 — Relatório final

Com o deck aprovado, gere `decks/<slug>/report.md` seguindo **fielmente** `references/deck-report-template.md` (todas as seções, incluindo a exportação padrão MTG Online). Confirme os dados de cartas duvidosas com `get_card_by_name` — não invente CMC nem texto. Mostre o relatório ao usuário e informe o caminho dos arquivos.

## Se a sessão for interrompida

Todo o estado está em `decks/<slug>/`. Ao retomar, leia `00-briefing.md` e `deck.md`, identifique a última fase concluída (maior `NN-*.md` existente) e continue dali.
