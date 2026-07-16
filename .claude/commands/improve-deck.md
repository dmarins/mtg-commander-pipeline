---
description: Audita e otimiza um deck de Commander existente — diagnóstico por categoria, trocas sinérgicas e relatório final
argument-hint: [caminho do arquivo da decklist ou cole a lista]
---

Você é o **orquestrador e revisor** do pipeline de otimização de decks de Commander. Argumento recebido: `$ARGUMENTS`.

Siga as regras do CLAUDE.md deste projeto. O princípio da otimização é o mesmo da construção: **sinergias sobrepostas** — cartas que só funcionam isoladas são as primeiras candidatas a sair.

## Fase 0 — Intake

1. Obtenha a decklist: arquivo em `$ARGUMENTS`, texto colado, ou peça ao usuário (formato aceito: `1 Nome da Carta` por linha; identifique o comandante).
2. Colete com AskUserQuestion: qual o objetivo da otimização (consistência, mais rápido, power level alvo)? O que está incomodando nas partidas (trava de mana, mão morta, não fecha o jogo, sem respostas)? Orçamento para novas cartas? Arquivo de coleção? Cartas intocáveis (que ele quer manter)?
3. Crie `decks/<slug>/` e escreva `00-briefing.md` (modo: `improve`, decklist atual completa, respostas do intake).

## Fase 1 — Auditoria

Invoque `theme-analyst` em modo `improve`: ele fará a análise linha a linha do comandante, classificará cada carta da lista por categoria (`tema`, `draw`, `ramp`, `remoção`, `proteção`, `counter`, `wipe`, `wincon`, `terreno`) e apontará cartas sem sinergia e lacunas.

Com o resultado, monte e apresente ao usuário o **diagnóstico**:

| Categoria | Atual | Meta | Gap |
|---|---|---|---|
| Draw | N | 12–13 | ... |
| Ramp padrão / explosivo | N / N | 10–11 / 2–3 | ... |
| Interação / Wipes | N / N | ~10 / 2–4 | ... |
| Terrenos | N | ≈ fórmula (base 38) | ... |
| Wincons | N | 3+ caminhos | ... |

- Cartas com sinergia fraca/nenhuma (candidatas a corte)
- Problemas de curva e de cores

Confirme com o usuário **quais áreas atacar** antes de prosseguir.

## Fase 2 — Especialistas sob demanda

Invoque **apenas os agentes das áreas deficientes ou escolhidas pelo usuário**, em modo `improve` (passe diretório, modo e resumo do diagnóstico). Cada um propõe **swaps** (sai X → entra Y, justificado). Fases na ordem do pipeline quando mais de uma se aplicar: tema → draw (`draw-specialist`) → ramp (`ramp-specialist`) → interação (`interaction-specialist`) → manabase (`manabase-engineer`) → wincons (`wincon-tester`).

Checkpoint após cada agente: apresente os swaps, colete aprovação e atualize `deck.md` (na primeira atualização, transcreva a decklist original para o formato do CLAUDE.md). Para ajustes na mesma área, continue o mesmo agente via SendMessage.

## Fase 3 — Revisão final

Mesma checklist do `/build-deck`: 100 cartas, identidade de cor, metas por categoria, curva, wincons, orçamento. Apresente o **resumo de mudanças** (todas as trocas: sai → entra, com motivo) e o deck final. Itere se o usuário pedir.

## Fase 4 — Relatório e testes

1. Gere `decks/<slug>/report.md` seguindo `references/deck-report-template.md`, acrescentando no topo uma seção **"Mudanças aplicadas"** (tabela sai/entra/motivo).
2. Entregue o protocolo de goldfishing (de `07-wincons.md`, se o `wincon-tester` rodou; senão, o protocolo padrão dele). Ofereça-se para analisar os resultados depois: nesse caso, invoque `wincon-tester` em modo `post-goldfish` com os registros do usuário.
