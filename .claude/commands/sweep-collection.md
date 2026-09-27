---
description: Varredura da coleção — cruza cada sobressalente com os decks, ranqueia onde cada uma rende mais e propõe as trocas dos finalistas
argument-hint: [@ <deck alvo>]
---

Você é o **orquestrador de uma varredura da coleção**. Argumento recebido: `$ARGUMENTS`.

É o `/swap-card` em lote. Lá a entrada é uma carta que o usuário trouxe; aqui as entradas são
**todas as sobressalentes** de `data/collection.tsv`. O trabalho é descobrir quais delas fazem
algum deck ganhar e qual carta sai para abrir cada slot. Nenhuma busca de carta nova no
Scryfall acontece aqui: a caixa é o universo de entradas.

A varredura é um **funil**, e o custo de cada etapa sobe à medida que o número de pares cai:

```
coleção × decks  →  corte por cor (mecânico, sem agente)
                 →  triagem pelo orquestrador (sinergia, lacuna, histórico)
                 →  finalistas: troca completa com especialista + gate de swap
                 →  um report.md por deck afetado  →  tabela única de trocas com ID  →  veredito
```

Siga as regras do CLAUDE.md. Leia `references/card-evaluation-checklist.md` antes da Fase 4:
cada saída é um corte como qualquer outro e passa pela ficha F1–F7 completa (regra 4).

## Modo autônomo — como este comando conversa com o usuário

O usuário **dispara a varredura e sai**. Ele volta no fim para ler a tabela de trocas e dar o
veredito. Isso define o comando:

- **Todas as perguntas acontecem na Fase 0**, num único AskUserQuestion. Da Fase 1 até a
  Fase 5, **nenhuma pergunta, nenhuma pausa, nenhum "posso seguir?"**. Mensagens no chat
  entre as fases são de progresso, curtas, e não esperam resposta.
- **Dúvida no meio do caminho não vira pergunta: vira regra padrão + pendência.** Cada caso
  ambíguo previsto tem o seu padrão escrito na fase correspondente. Caso não previsto: escolha
  a opção **mais conservadora** (a que não troca), registre como pendência no `report.md` do
  deck e siga.
- **Nada que dependa da aprovação é escrito antes dela** (regra 9). A varredura produz
  proposta (pasta da rodada, arquivos de fase, `report.md`). `deck.md`, `decisions.md`, o
  índice de rodadas e a coleção só mudam na Fase 6.
- **Sem cotação no navegador.** Entradas são sobressalentes e custam R$ 0,00. O total do deck
  usa `mtgdb prices`; o que não tiver cotação sai como `a cotar` (regra 2), sem abrir o
  browser. Isso fica para depois da aprovação.

A sessão é descartável, porque o usuário dá `/clear` ao fim. **Tudo que precisa sobreviver vai
para arquivo**; nada fica só no chat.

## Fase 0 — Preparação e perguntas (única interação antes do fim)

Levante tudo primeiro e pergunte uma vez só.

1. **Estado da coleção.** A varredura inteira responde a uma pergunta sobre a caixa física. Com
   `data/collection.tsv` defasado, o resultado propõe carta que já foi vendida e ignora carta
   que o usuário tem.
   - total: `bin/mtgdb collection -list | tail -1`;
   - última sincronização: `git log -1 --format='%ad %s' --date=short -- data/collection.tsv`
     e a maior data da coluna `added_at`;
   - alterações não commitadas: `git status --short data/collection.tsv`;
   - **o que pode ter virado sobressalente depois disso**: rodadas fechadas após a última
     sincronização (índice de rodadas de cada `00-briefing.md`) e cortes do `decisions.md`
     com data posterior. Em deck **montado**, o corte aprovado sai da caixa do deck, mas só
     entra na coleção pela `/update-collection` (regra 7: posse nunca é inferida). Liste
     esses cortes por deck. São as cartas que a varredura não vai enxergar.
2. **Decks e escopo.** Argumento `[@ <deck>]`: com deck alvo (slug ou nome do comandante), a
   varredura olha só ele; sem argumento, todos em `ls decks/`. Para cada deck, leia no
   `00-briefing.md`: **Comandante** atual, **Identidade de cor**, **Modo**, **Status físico**
   (e qual lista ele declara como real), régua de orçamento/mesa-alvo e **alvos protegidos**.
   Deck **sem o campo Status físico** entra nas perguntas do item 3. Posse é declarada, não
   deduzida.
3. **Pergunte uma vez**, com AskUserQuestion. Junte as perguntas numa chamada só (até 4):
   - **Coleção** (sempre), com os números do item 1 no texto da pergunta:
     - **"Está atualizada, pode seguir"**;
     - **"Vou atualizar agora (colo a lista)"**: execute o fluxo de
       `.claude/commands/update-collection.md` **inteiro**, com `-dry-run`, diff e aprovação,
       sem pular a conferência por estar dentro de outro comando. Aplicada a sincronização,
       refaça o item 1 e siga;
     - **"Atualizo depois, encerra aqui"**: instrua a rodar `/update-collection <lista>` e
       depois `/sweep-collection` de novo. Encerre sem abrir rodada nem criar arquivo.

     Não sugira nem monte a lista a partir dos cortes do item 1. Eles servem para o usuário
     lembrar o que conferir na caixa, não para você registrar posse.
   - **Status físico** (só se algum deck não tiver o campo): montado ou só no papel, e qual
     lista é a real.
4. Respondida a Fase 0, avise em uma linha que a varredura segue sozinha até a tabela final,
   e **não pergunte mais nada até a Fase 5**.

## Fase 1 — Regras de escopo por deck (sem perguntas)

- **Lista de referência:** a troca é feita contra a lista que o campo Status físico declara
  como real. Se `deck.md` divergir dela (deck em playtest de proxy, versão nova só no papel),
  **use a declarada no Status físico** e registre a divergência como pendência no `report.md`
  daquele deck. Não pergunte.
- **Deck em `build` com rodada aberta** não tem troca. As sobressalentes que passarem na
  triagem viram **candidatas ao pool daquela rodada**: recomendação no relatório dele, sem ID
  na tabela de trocas e sem rodada nova.

## Fase 2 — Corte por cor (mecânico)

A identidade de cor elimina pares **sem análise**. Faça em lote, não carta a carta:

```bash
tail -n +2 data/collection.tsv | cut -f1 | tr '\n' '\0' \
  | xargs -0 bin/mtgdb oracle -json \
  | jq -r '.[] | [.query, .found, (.card.ColorIdentity // ""), (.card.LegalCommander // ""), (.card.TypeLine // "")] | @tsv'
```

- `found=false`: siga a escada da regra 6 (`make refresh` → MCP). Nome que nem assim resolve é
  carta fantasma na coleção: sai da varredura e vai para as **pendências gerais** da tabela
  final, com a sugestão de corrigir pela `/update-collection`. Não pare por isso.
- `LegalCommander` diferente de `legal`: fora de todos os decks.
- **Elegível** num deck = cada letra da identidade da carta está na identidade do deck.
  Identidade vazia (incolor) é elegível em todos.
- Descarte o par se o deck **já tem a carta** (`grep -iF` na lista de referência).
- Respeite `qty`: a coleção diz quantas cópias existem. Uma cópia vai para **um** deck só
  (regra 7).

Registre a matriz (deck · identidade · nº de sobressalentes elegíveis, mais as cartas sem deck
elegível). Ela vai para os relatórios. No chat, só uma linha de progresso.

## Fase 3 — Triagem (orquestrador)

Com a lista de pares elegíveis, **você** faz a triagem. Invocar especialista por par é caro
e desnecessário para decidir *se* a carta tem lugar. Para ter os dados em mãos:

- `bin/mtgdb deck <slug>` de cada deck: agregados, curva e contagem por categoria. Anote as
  **lacunas** contra as metas (draw 12–13, ramp 10–11 + 2–3, interação ~10, wipes 2–4,
  terrenos pela fórmula, wincons 3+).
- `bin/mtgdb oracle <nomes...>` das sobressalentes elegíveis. **Leia o texto oracle**; não
  julgue pelo nome nem pelas tags (regra 6). Uma leitura por carta serve para todos os decks
  em que ela é elegível.

Para cada par carta × deck:

1. **Histórico primeiro (regra 5).** Muitas sobressalentes **saíram destes mesmos decks**.
   Procure a carta no `decisions.md` do deck. Se ela foi cortada ali e o motivo continua
   válido, o par cai, com uma linha dizendo quem cortou, quando e por quê. Também caem as
   **reprovações do usuário** registradas ali: veredito dele não se repropõe sem fato novo.
   Só volta se algo mudou (comandante novo, lacuna nova, peça que ela alimentava entrou), e a
   mudança é dita.
2. **Sinergia (regra 3):** 2+ pontos com o comandante e/ou com peças do deck. Menos que isso,
   o par cai.
3. **Lacuna:** carta que cobre categoria abaixo da meta vale mais que carta que só substitui
   uma equivalente.
4. **Ganho plausível:** existe no deck uma carta visivelmente pior na mesma função? Aqui é só
   a hipótese de saída para ranquear. O veredito do corte é da Fase 4, com ficha completa.
   Nunca aponte alvo protegido como hipótese.

Custo não entra na triagem: sobressalente já está na caixa e não custa nada. É a razão de a
varredura existir.

**Ranqueie e resolva conflitos:**

- Carta boa em dois decks, com uma cópia só: fica no deck onde o ganho é maior. O segundo vira
  **alternativa** registrada. Se o usuário reprovar a troca no primeiro deck, é ela que se
  oferece (Fase 6).
- **Finalistas:** no máximo **3 trocas por deck** e **8 no total** por varredura. O que passou
  na triagem e ficou fora do corte vai para o relatório como "próximas da fila", com uma linha
  de motivo cada.
- Nenhum par sobreviveu: encerre sem abrir rodada, com um resumo curto da triagem no chat.
  "Nada da caixa melhora os decks hoje" é uma resposta válida.

Uma linha de progresso no chat (quantos finalistas, em quais decks) e siga direto.

## Fase 4 — As trocas nos decks (especialistas)

Para cada deck com finalistas:

1. **Abra a rodada:** maior `v<N>` do índice de rodadas do briefing → crie
   `decks/<slug>/rounds/v<N+1>-<hoje>/`. É uma rodada por deck, com todas as trocas da
   varredura dentro dela. O índice do briefing só é tocado na Fase 6.
2. Invoque o especialista da função principal de cada finalista, em modo `improve`, com a mesma
   tabela de roteamento da Fase 4 do `.claude/commands/swap-card.md`. Finalistas do mesmo deck
   que caem no mesmo especialista vão **numa invocação só**. Passe: o diretório do deck, a
   pasta da rodada, a lista de referência, a ficha de cada entrada e a instrução de **troca
   pontual em lote, sem interação com o usuário**: *entradas fixas = `<cartas>`; para cada
   uma, proponha **uma** saída (no máximo duas candidatas ranqueadas), com ficha F1–F7
   completa, cobertura de cada função da carta que sai e simetria de critério; duas entradas
   não podem disputar a mesma saída; não proponha outras entradas; não faça perguntas: dúvida
   vira pendência escrita no arquivo da fase.*
   Especialistas de decks diferentes são independentes. Invoque-os em paralelo.
3. **Gate de swap:** rode os 5 itens do gate de `.claude/commands/improve-deck.md` em cada
   troca. O que falhar volta ao especialista (SendMessage) **uma vez**, com o item que falhou.
   Falhou de novo: a troca **cai** e vai para o relatório como "reprovada no gate", com o
   motivo. Não há terceira tentativa, e uma troca travada não segura as outras. Troca que o
   especialista conclui não valer a pena (a carta não ganha de nenhuma do deck) cai com o
   motivo dito. Isso é resultado, não falha.
4. **Recontagem depois de todas as trocas do deck juntas:** categorias, curva, tipos que
   alimentam contagens do deck. O deck continua com exatamente 100 cartas. Trocas que se
   somam podem abrir buraco que nenhuma abre sozinha, por exemplo duas saídas da mesma
   categoria. Se isso acontecer, derrube a troca de menor ganho e registre por quê.

## Fase 5 — Relatórios e tabela de veredito

### IDs das trocas

Cada troca que sobreviveu ao gate recebe um **ID estável**: prefixo curto do deck + número,
na ordem do ranking do deck. O prefixo é a primeira palavra do comandante, minúscula
(`krenko-1`, `thorin-2`, `inspirit-1`). O mesmo ID aparece no chat, no `report.md` e depois no
`decisions.md`. É por ele que o usuário responde.

### `report.md` por deck

Um por deck afetado, em `decks/<slug>/rounds/v<N+1>-<hoje>/report.md`, seguindo
`references/deck-report-template.md`, com no topo uma seção **"Varredura da coleção"**:

- data da varredura e estado da coleção usado (total, data da última sincronização);
- tabela **Trocas propostas**: ID · sai → entra · função · motivo em uma linha;
- para cada troca: a ficha da carta que sai e quem cobre cada função dela;
- o **corte por cor** e a **triagem** deste deck: pares que caíram por histórico (com o
  registro de `decisions.md`), trocas reprovadas no gate, "próximas da fila" e cartas que eram
  boas aqui mas foram para outro deck por disputa de cópia;
- custo: entradas **R$ 0,00 (sobressalente)**; total do deck pela regra 2 (LigaMagic menor, com
  data; `a cotar` onde faltar);
- **pendências**: lista de referência divergente, fichas via MCP, cotações, casos resolvidos
  por regra padrão.

A lista final completa e o bloco de importação em texto puro assumem **todas as trocas do deck
aprovadas**. Na Fase 6 o relatório é reconsolidado com o que de fato passou. Rode
`bin/linkify <report.md>` em cada um.

Deck em `build`: o relatório é o da rodada aberta, com as sobressalentes aprovadas na triagem
registradas como candidatas ao pool. Não mexa na lista.

### A mensagem final no chat

É o que o usuário lê quando volta. Autossuficiente, sem precisar rolar a sessão:

1. uma linha de contexto: coleção usada (total/data), quantos pares avaliados, quantas trocas
   propostas em quantos decks;
2. **uma tabela por deck**, com o caminho do `report.md` no título:

   | ID | Sai | Entra | Por quê |
   |---|---|---|---|

3. candidatas ao pool de decks em `build`, se houver (sem ID, são só recomendação);
4. pendências gerais (cartas fantasma na coleção, fichas via MCP, `make refresh` pendente);
5. **como responder**, literalmente:

   > Responda do jeito que preferir, por exemplo:
   > - `tudo`: aprova todas as trocas de todos os decks
   > - `tudo no krenko`: aprova todas as trocas daquele deck (os outros ficam pendentes)
   > - `krenko-2 não`: reprova só essa troca (dá para juntar: `tudo, menos krenko-2 e thorin-1`)
   > - `thorin nenhuma`: reprova todas as trocas daquele deck
   > - motivo opcional depois de `—`: `krenko-2 não — quero manter o Goblin Chieftain`

   Interprete linguagem natural equivalente. Troca não mencionada fica **pendente**. Não é
   aprovada por omissão.

## Fase 6 — Aplicar o veredito

Aplique exatamente o que foi dito, deck a deck. Para cada troca:

- **Aprovada:**
  1. `deck.md`: sai/entra, com categorias e sinergias.
  2. `decisions.md`: **as duas pontas** (a entrada, com origem "sobressalente, varredura de
     `<data>`, `<ID>`", e o corte), com o motivo e a data.
- **Reprovada:** registre em `decisions.md` como "proposta reprovada pelo usuário" com o ID, a
  data e o motivo, se ele deu (ou "sem motivo informado"). É esse registro que impede a mesma
  troca de voltar na próxima varredura (Fase 3, item 1). **Não reinvoque especialista.** Se
  a carta tinha **alternativa** em outro deck, mencione no fechamento como opção para uma
  próxima rodada (`/swap-card <carta> @ <deck>`), sem executar.
- **Pendente:** nada é escrito sobre ela. Continua no `report.md` e é listada no fechamento.

Depois, por deck:

1. Deck com alguma troca aprovada: **reconsolide o `report.md`** (lista final, bloco de
   importação e recontagem só com as aprovadas; reprovadas e pendentes marcadas na tabela),
   rode `bin/linkify` de novo e acrescente a rodada ao índice do `00-briefing.md` (rodada,
   data, estado, link).
2. Deck com tudo reprovado: apague a pasta da rodada. Rodada sem entrega não entra no índice;
   as reprovações ficam no `decisions.md`.
3. Deck com tudo pendente: mantenha a pasta, marque o estado `pendente` no índice do briefing
   e liste no fechamento.

**Coleção:** a sobressalente que entrou num deck montado **sai** da coleção, e a carta cortada
**pode** virar sobressalente. As duas mudanças são físicas, então nenhuma é feita aqui. No
fechamento, liste por deck o que tirar da caixa de sobressalentes e o que devolver a ela, e
diga que, depois de montar as trocas, ele roda `/update-collection` com a lista nova. Posse
nunca é inferida (regra 7).

Feche com um resumo curto (aprovadas, reprovadas e pendentes por deck, e as movimentações
físicas) e ofereça o commit (decks afetados + `data/prices.tsv` se mudou), sem fazê-lo por
conta própria. Se alguma ficha veio do MCP, repita o aviso do `make refresh` pendente.
