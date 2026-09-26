---
description: Troca pontual — recebe uma carta, descobre em qual deck ela encaixa melhor e propõe a saída correspondente
argument-hint: <nome da carta em inglês> [@ <deck alvo>]
---

Você é o **orquestrador de uma otimização cirúrgica**. Argumento recebido: `$ARGUMENTS`.

A entrada é **fixa**: uma carta que o usuário trouxe. O trabalho não é buscar cartas novas —
é descobrir **qual deck ganha mais com ela** e **qual carta sai** para abrir o slot. Nenhuma
busca de carta nova no Scryfall acontece aqui.

Siga as regras do CLAUDE.md. Leia `references/card-evaluation-checklist.md` antes da Fase 3:
a saída é um corte como qualquer outro e passa pela ficha F1–F7 completa (regra 4).

A sessão é descartável — o usuário dá `/clear` ao fim. **Tudo que precisa sobreviver vai para
arquivo** (`report.md` da rodada, `decisions.md`, `deck.md`); nada fica só no chat.

## Fase 0 — Leia o argumento

- Formato: `<carta> [@ <deck>]`. Tudo antes do `@` é o nome da carta; o que vem depois é o
  deck alvo — slug (`krenko-mob-boss`) ou nome do comandante (`Krenko`). Se o usuário
  escrever em linguagem natural ("Swords to Plowshares no Thorin"), interprete.
- Sem argumento: peça o nome da carta e pare.
- Nome em português: a regra 8 exige o nome oficial em inglês. Não traduza de memória —
  peça o nome em inglês ou confirme o candidato que o banco/MCP devolver.

## Fase 1 — Ficha da carta (banco local → refresh → MCP)

O `mtgdb` tem a base inteira do Scryfall (uma entrada por carta, de Alpha até as cartas já
reveladas), atualizada pelo `make refresh`. Carta ausente quase sempre significa **banco
desatualizado**, não carta desconhecida — por isso o refresh vem antes do MCP.

1. **Banco local:** `bin/mtgdb oracle -rulings "<carta>"`. Encontrou → siga para o item 4.
   Se o banco não existir, `make db`. Resolvido com `how` diferente de `exact` (prefixo, fts,
   face): confirme que é a carta certa antes de seguir.
2. **Não encontrou → `make refresh`** (~15 s: rebaixa o bulk do Scryfall e reconstrói o
   banco) e rode `bin/mtgdb oracle -rulings "<carta>"` de novo. Encontrou → a carta agora está
   no banco, completa (tags e rulings); siga para o item 4.
3. **Nem o refresh trouxe** (a carta entrou no Scryfall há menos de um dia e ainda não está
   no dump diário) → use o MCP `mcp__mtg__get_card_details` e `mcp__mtg__get_card_rulings`
   para a análise. A análise segue com o que veio do MCP, e o relatório marca a
   **pendência**: "ficha via MCP — sem tags do Tagger; carta fora do banco local".
   **Avise o usuário** de que será preciso um novo `make refresh` no futuro (a partir de
   amanhã) para a carta entrar no `mtgdb`.
   - Se o MCP também não achar, pare e reporte o nome — **nunca avalie de memória** (regra 6).
   - O MCP **não traz o texto das faces** de carta de duas faces (só nome, tipo e link).
     Nesse caso diga isso ao usuário e leia o texto na ficha do Scryfall pelo link — não siga
     com o texto incompleto.
4. **Legalidade:** `legal_commander` diferente de `legal` → pare aqui. Carta banida não entra
   em deck nenhum.
5. Anote da ficha: **identidade de cor**, CMC, tipo, função(ões) — e rode
   `bin/mtgdb collection "<carta>"` (está nas sobressalentes?) e `bin/mtgdb prices "<carta>"`.

**Demais ferramentas do MCP `mtg`** (EDHREC, Archidekt, busca, regras, legalidade de deck):
o orquestrador **pode** usá-las, mas por padrão não precisa — a consulta de carta é feita na
base local. Use-as quando o usuário pedir explicitamente. Nos especialistas, elas seguem
fazendo parte do workflow de cada fase, como definido na própria definição do agente.

## Fase 2 — Corte por cor (antes de qualquer análise)

A identidade de cor elimina decks **mecanicamente** — não se analisa sinergia de carta branca
num deck mono-vermelho.

1. Liste os decks: `ls decks/`. Para cada um, leia no `00-briefing.md` as linhas
   **Comandante** (a atual — alguns briefings registram também o anterior),
   **Identidade de cor**, **Modo**, **Status físico** e a régua de orçamento/mesa-alvo.
2. **Elegível** = cada cor da identidade da carta está na identidade do deck. Carta incolor
   é elegível em todos.
3. Descarte também os decks que **já têm a carta** (`grep -i` no `deck.md`/`lista.txt`, ou
   `bin/mtgdb deck <slug>`).
4. **Deck alvo informado:** avalie só ele. Se ele falhar no corte por cor, diga por quê numa
   linha e pergunte se o usuário quer que os outros decks sejam avaliados — não redirecione
   por conta própria.
5. **Nenhum elegível:** diga isso com a identidade de cada deck e **encerre** — sem abrir
   rodada, sem arquivo.

Mostre o corte em uma tabela curta (deck · identidade · elegível? · motivo) e siga.

## Fase 3 — Triagem entre os elegíveis (orquestrador)

Com mais de um deck elegível, **você** faz a triagem — invocar um especialista por deck é
caro e desnecessário para decidir *onde* a carta rende mais. Para cada deck elegível:

- `bin/mtgdb deck <slug>` — agregados, curva e contagem por categoria.
- **Sinergia (regra 3):** quantos pontos a carta soma com o comandante e com as peças do
  deck? Menos de 2 → o deck sai da disputa, com o motivo dito.
- **Lacuna:** o deck está abaixo da meta na categoria da carta (draw 12–13, ramp 10–11 + 2–3,
  interação ~10, wipes 2–4, terrenos pela fórmula, wincons 3+)? Carta que cobre lacuna vale
  mais que carta que só substitui uma equivalente.
- **Histórico (regra 5):** `decisions.md` do deck — a carta já foi cortada ali? Se o motivo
  original continua válido, ela não volta; diga o que mudou, se algo mudou.
- **Orçamento:** a régua do briefing (torneio/mesão) comporta a carta? Preço é **só** o menor
  da LigaMagic (regra 2): sem cotação em `mtgdb prices`, cote no navegador e registre com
  `mtgdb prices -add`. Sem cotação possível → `a cotar`, nunca estimativa.
- **Modo do deck:** deck em `build` com rodada aberta não tem troca — a carta vira candidata
  ao pool daquela rodada. Registre isso no relatório como recomendação e não abra rodada nova.

Ranqueie. **Uma cópia vai para um deck só** (regra 7 — deck montado é coleção fechada): o
vencedor recebe a proposta completa; um segundo deck muito próximo aparece como alternativa
(que exigiria outra cópia). Se nenhum deck ganha de verdade com a carta, **diga isso** e
encerre — "não entra em lugar nenhum" é uma resposta válida.

## Fase 4 — A troca no deck escolhido (especialista)

1. **Abra a rodada** do deck: maior `v<N>` do índice de rodadas do briefing → crie
   `decks/<slug>/rounds/v<N+1>-<hoje>/`.
2. Confira qual lista é a real (campo **Status físico**): a troca é feita contra ela. Se
   `deck.md` e a lista física divergirem (deck em playtest de proxy, v nova só no papel),
   pergunte contra qual lista trocar antes de invocar o especialista.
3. Invoque **o especialista da função principal da carta**, em modo `improve`:

   | Função da carta | Especialista |
   |---|---|
   | tema / sinergia do comandante | `theme-analyst` |
   | draw | `draw-specialist` |
   | ramp | `ramp-specialist` |
   | remoção, counter, proteção, wipe | `interaction-specialist` |
   | terreno | `manabase-engineer` |
   | finisher / wincon | `wincon-tester` |

   Passe na invocação: o diretório do deck, a pasta da rodada, o modo `improve`, a ficha da
   carta (Fase 1) e a instrução de **troca pontual** — *entrada fixa = `<carta>`; proponha
   **uma** saída (no máximo duas candidatas ranqueadas), com ficha F1–F7 completa, cobertura
   de cada função da carta que sai e simetria de critério; não proponha outras entradas.*
   O especialista escreve o arquivo da sua fase na pasta da rodada, como no `/improve-deck`.
4. **Gate de swap:** rode os 5 itens do gate de `.claude/commands/improve-deck.md` (ficha
   completa, cobertura, simetria, histórico, alvo protegido). O que falhar volta ao
   especialista — não chega ao usuário.
5. Recontagem: categoria, curva, tipos que alimentam contagens do deck, e o deck continua com
   exatamente 100 cartas.

## Fase 5 — Relatório e decisão

Gere `decks/<slug>/rounds/v<N+1>-<hoje>/report.md` seguindo `references/deck-report-template.md`,
com no topo uma seção **"Troca pontual"**:

- a carta que entra (ficha resumida, preço, se está nas sobressalentes);
- o **corte por cor** e o **ranking dos decks elegíveis**, com o motivo de cada um;
- a troca: **sai → entra**, com a ficha da carta que sai e quem cobre cada função dela;
- impacto no custo (LigaMagic menor, com data) e pendências (cotação, ficha via MCP, posse).

**Posse:** carta fora de `mtgdb collection` é tratada como **compra**. Se o usuário já tem a
carta em mãos, ele declara — posse nunca é inferida (regra 7).

O relatório leva a lista final completa e o bloco de importação em texto puro, como todo
`report.md`. Rode `bin/linkify <report.md>` no fim. Apresente a proposta no chat de forma
curta, apontando o `report.md`, e peça a aprovação.

## Fase 6 — Aplicar e fechar

Só com a aprovação do usuário (regra 9):

1. Atualize `deck.md` (sai/entra, com categorias e sinergias) e registre em `decisions.md`
   **as duas pontas**: a entrada e o corte, com o motivo e a data.
2. Acrescente a rodada no índice do `00-briefing.md` (rodada, data, estado, link).
3. Deck **montado**: lembre que a carta cortada só vira sobressalente pelo `/update-collection`
   — corte não entra na coleção automaticamente.
4. Reprovação: o motivo vai para `decisions.md` e a proposta volta ao mesmo especialista
   (SendMessage) — ou, se a reprovação for do deck escolhido, ao segundo do ranking.
5. Ofereça o commit (deck + `data/prices.tsv` se houve cotação nova), sem fazê-lo por conta
   própria. Se a ficha veio do MCP (Fase 1, item 3), repita o aviso do `make refresh`
   pendente. Depois disso a sessão pode ser limpa com `/clear`.
