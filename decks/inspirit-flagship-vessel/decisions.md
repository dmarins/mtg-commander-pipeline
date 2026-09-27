# Registro de decisões — Inspirit, Flagship Vessel

> Arquivo aberto em 2026-09-18, na reconstrução do deck a partir do histórico do git.
> As rodadas v1 e v2 foram conduzidas **antes** de a regra 5 existir e não produziram registro
> cronológico de cortes e entradas. **Nada foi inventado aqui**: o histórico delas é o que está
> escrito nos relatórios de cada rodada.

- **v1 (2026-07-25)** — trocas e justificativas em [`rounds/v1-2026-07-25/report.md`](rounds/v1-2026-07-25/report.md).
- **v2 (2026-07-25 → 2026-08-12)** — trocas e justificativas em [`rounds/v2-2026-08-12/report.md`](rounds/v2-2026-08-12/report.md);
  divergência torneio × mesão em [`rounds/v2-2026-08-12/08-versoes.md`](rounds/v2-2026-08-12/08-versoes.md).
- **v3 (2026-08-13)** — 6 trocas aplicadas. O relatório original **se perdeu** (a pasta existiu em
  disco e nunca foi commitada); resumo reconstruído em
  [`rounds/v3-2026-08-13/NOTAS-RECONSTRUIDAS.md`](rounds/v3-2026-08-13/NOTAS-RECONSTRUIDAS.md).
  Ficou **aberta** a troca `Rip Apart` → `Sunder the Gateway` (custo declarado: Rip Apart é o único
  removedor dedicado de planeswalker do deck).
- **pós-v3** — o usuário evoluiu o deck por conta própria. O que ele montou **não está documentado**
  e será registrado quando ele passar a lista atual.

⚠️ **Consequência para a regra 5:** ao propor uma carta que já esteve neste deck, o "quem cortou e
por quê" precisa ser procurado nos dois relatórios acima — e, para o período pós-v2, **perguntado ao
usuário**, porque não há registro. A partir da v3 este arquivo passa a ser append-only por rodada.

---

## Registro pós-v3 — lista montada informada em 2026-09-18

O usuário passou a lista do deck **físico**. Não houve rodada de pipeline: é **registro**, não
otimização. As mudanças abaixo foram feitas por ele, fora do processo, entre 2026-08-13 e
2026-09-18. **O motivo de cada uma não foi declarado** — o que está aqui é o *fato* medido, não a
justificativa. Para a regra 5, tratar como "cortada pelo usuário, motivo desconhecido": ao
repropor qualquer uma delas, **pergunte a ele** antes.

> ⚠️ O diff é **v2 → hoje**, não v3 → hoje: o relatório da v3 se perdeu e não existe lista dela em
> disco. As 6 trocas da v3 estão dentro deste intervalo e aparecem misturadas às do usuário.

### Saíram (15 cartas + 1 Mountain)

| Carta | Função que exercia na v2 | Observação |
|---|---|---|
| Blast Zone | wipe escalável em terreno | só proliferate a alimentava (o gatilho 1+ mira artefato) |
| Boros Signet | fixação R/W | entrou na v2 justamente para socorrer o R; hoje o R é a cor mais magra (4 Mountain) |
| Disruption Protocol | counter | **o deck ficou com 0 counterspells** |
| Stoic Rebuttal | counter | idem |
| Fumigate | wipe assimétrico | wipes caíram de 3 (com Blast Zone) para 2 |
| Glass Casket | remoção por exílio | — |
| Leonin Abunas | 2ª camada de hexproof no comandante | sobrou só Padeem cobrindo o Inspirit |
| Lux Artillery | wincon B (queima aos 30 counters) | o caminho B de vitória deixou de existir |
| Lux Cannon | remoção recorrente por charge counters | — |
| Reckoner Bankbuster | draw recarregável pelo gatilho 1+ | — |
| Rip Apart | única remoção dedicada de planeswalker | a troca aberta da v3 **foi resolvida**: `Sunder the Gateway` está no deck e `Rip Apart` não |
| Sphere of the Suns | ramp com charges | corte da v3 (→ Solar Array) |
| Thirst for Knowledge | draw 3 descartando artefato | — |
| Unwanted Remake | remoção instantânea por {W} | corte da v3 (→ Paladin Danse, que **não** está na lista final) |
| Temple of Epiphany | terreno U/R | — |
| Mountain (5 → 4) | fonte de R | — |

### Entraram (16 cartas)

| Carta | Função | Observação |
|---|---|---|
| Astral Cornucopia | ramp escalável em charge counters | — |
| Chief of the Foundry | anthem de criaturas-artefato | era a peça da versão **torneio** da v2; virou fixa |
| Dawnsire, Sunstar Dreadnought | remoção/wincon (100 de dano no 10+) | 4º spacecraft — aumenta a disputa por corpos para Station |
| Emissary Escort | corpo grande e barato para Station | corte→entrada da v3 (saiu Cloud Key) — mas **Cloud Key voltou** |
| Ethersworn Sphinx | draw via cascade + 4/4 voador por affinity | — |
| Golem Foundry | fábrica de Golems 3/3 por charge counters | entrada da v3 (saiu Organic Extinction) — mas **Organic Extinction voltou** |
| Marketback Walker | corpo X + draw ao morrer | — |
| Metastatic Evangel | proliferate por criatura nontoken | entrada da v3 (saiu Crystalline Crawler) — mas **Crystalline Crawler voltou** |
| Recon Craft Theta | proliferate ao atacar + traz o próprio tripulante | 3ª fonte repetível de contadores |
| Solar Array | mana + sunburst (arranque frio de contadores) | entrada da v3 (saiu Sphere of the Suns) |
| Spring-Loaded Sawblades | remoção de criatura virada em flash; craft → veículo | — |
| Stone by Sunlight | remoção modal / indestrutível + artificializa | entrada da v3 (saiu Swords to Plowshares) — mas **Swords to Plowshares voltou** |
| Sunder the Gateway | remoção de artefato/encantamento + Incubate 2 | resolve a questão aberta de 2026-08-13 |
| Vivid Crag | terreno de qualquer cor com charge counters | — |
| Voyage Home | draw 3 + 3 de vida por affinity | — |
| Voyager Quickwelder | redutor de custo nº 5 | — |

### Cartas que voltaram depois de cortadas na v3

`Cloud Key`, `Crystalline Crawler`, `Organic Extinction` e `Swords to Plowshares` foram cortadas na
v3 e **estão de volta** na lista montada. Nenhuma das entradas da v3 que as substituiu foi desfeita
— elas convivem. `Paladin Danse, Steel Maverick` é a **única** entrada da v3 que não chegou à lista
final. Motivos não declarados; **perguntar antes de tratar esses cortes da v3 como válidos** (regra 5).

---

## Troca avulsa — 2026-09-26

O usuário trouxe uma carta que **já possui** e perguntou onde ela rende mais. Pela identidade
(`{U}{R}{W}`), só o Inspirit a comporta. Ficha completa feita, troca aprovada por ele. Não é rodada
de pipeline.

| Data | Carta | Ação | Fase/agente | Motivo | Funções descobertas pelo corte |
|---|---|---|---|---|---|
| 2026-09-26 | Vraska, Soul of Stone | entrada | avulsa · orquestrador | Sculpture Treasure 1/1 (criatura-artefato) por mágica não-criatura (~36 no deck): corpo para Station/crew **ou** mana de qualquer cor; conta como artefato (Master of Etherium, affinity, improvise, Malcator). **Motivo que o usuário destacou:** vigilância às criaturas-artefato — atacam e ainda estacionam na 2ª fase principal (Station é velocidade de feitiço). Sculpture recém-criada pode estacionar no mesmo turno (o `{T}` é custo da nave, não habilidade da criatura). | — |
| 2026-09-26 | Third Path Iconoclast | corte | avulsa · orquestrador | F1 contido na Vraska: mesmo token 1/1 criatura-artefato por mágica não-criatura, sem o Treasure. F2 corpo 2/1 → Vraska 3/3. F3/F4 tokens-artefato que recebem os anthems → Sculptures, mesma taxa e mesmos anthems. F5 nenhum. | **F6: o motor de tokens deixa de começar no turno 2** (2-drops 13→12, 3-drops 21→22). Custo aceito. |

**Atritos da entrada, aceitos pelo usuário:**

- **Alibou, Ancient Witness perde dano.** O X dela é o nº de artefatos **virados** quando o gatilho
  resolve, e o ruling diz explicitamente que atacantes com vigilância não contam. Com a Vraska em
  campo, X vem só do que virou antes do combate (rocks, crew, Station na 1ª fase principal). A cada
  turno o usuário escolhe: *modo Station* (atacar e estacionar depois) ou *modo Alibou* (virar tudo
  antes do combate). Compensação: a haste da Alibou deixa a Sculpture recém-criada atacar e ainda
  estacionar depois.
- **Kilo, Apogee Mind não perde:** deixa de proliferar ao atacar, mas prolifera ao ser virada para
  Station depois do combate — ataque + proliferate + Station, onde antes era ataque + proliferate.
- **Vraska não é artefato:** sem a proteção do comandante; morre para Organic Extinction e Chain
  Reaction do próprio deck (as Sculptures sobrevivem). Pede `{U}{R}{W}` no turno 3, com R sendo a
  cor mais magra da base.

**Alternativas avaliadas e mantidas:** Chrome Host Seedshark (voar e tokens de tamanho X — Incubate 5
do Dawnsire é Station de 5 — funções exclusivas) e Saheeli, Sublime Artificer (é mágica não-criatura,
dispara Vraska/Seedshark/Whirlwind; `−2` de cópia sem par no deck). Malcator ganha com a entrada.

**Custo:** Third Path Iconoclast R$ 3,58 (LigaMagic menor, cotação de 2026-09-18) sai; Vraska
**a cotar**. Sem compra — a carta já é do usuário.

---

## v4 — troca pontual (`/swap-card Requisition Raid`) · 2026-09-26

O usuário trouxe `Requisition Raid` (sobressalente, custo zero) e pediu o deck onde ela rende mais.
Corte por cor: Inspirit, Thorin e Phlage elegíveis; Krenko e Satoru fora (sem W). Inspirit venceu
a triagem (interação 9/~10, só 2 respostas a artefato/encantamento, pacote de contadores e de
mágica não-criatura). Ficha completa dos dois lados em
[`rounds/v4-2026-09-26/05-interaction.md`](rounds/v4-2026-09-26/05-interaction.md); relatório em
[`rounds/v4-2026-09-26/report.md`](rounds/v4-2026-09-26/report.md). **Aprovada pelo usuário.**

| Data | Carta | Ação | Fase/agente | Motivo | Funções descobertas pelo corte |
|---|---|---|---|---|---|
| 2026-09-26 | Requisition Raid | entrada | v4 · swap-card · interaction-specialist | R$ 0,44 (LigaMagic menor, 2026-09-23), **da coleção** — compra zero. 3ª resposta a artefato/encantamento (pega token e artefato de qualquer jogador; os dois numa mágica por `{3}{W}`); interação 9 → 10. Modo 3: +1/+1 permanente no time × 5 proliferates + Deepglow Skate; Hangarback/Marketback/Crystalline Crawler convertem o contador. Mágica não-criatura de MV 1: dispara Vraska, Saheeli e Whirlwind, e os tokens desses gatilhos recebem o contador | — |
| 2026-09-26 | Reverse Engineer | corte | v4 · swap-card · interaction-specialist | draw acima da meta (14 → 13, dentro de 12–13); feitiço não-artefato por feitiço não-artefato — artefatos (42), criaturas (27) e gatilhos de mágica não-criatura inalterados; alivia `{U}{U}`. Entrou na v1 pelo pipeline no slot da Ethersworn Sphinx, que o usuário trouxe de volta — o motivo da entrada caducou | compra 3 → Voyage Home, Tezzeret's Gambit, Stern Lesson, Thought Monitor, Ethersworn Sphinx + motores; gatilho de mágica não-criatura → Requisition Raid; **Incubate 5 via Chrome Host Seedshark → descoberto em parte** (Raid dá Incubate 1; custo aceito); **improvise → artefatos virados antes do combate para o X da Alibou → descoberto em parte** (custo aceito; Organic Extinction e Kappa Cannoneer seguem com improvise) |

**Reserva não usada:** Voyage Home — cortada pelo pipeline na v1 e trazida de volta pelo usuário sem
motivo declarado; não foi tocada. **Posse:** Reverse Engineer só vira sobressalente se o usuário o
listar na próxima `/update-collection` (regra 7). **Custo da lista:** R$ 252,71 → R$ 252,95 +
Vraska `a cotar` (LigaMagic menor, cotações de 2026-09-18 a 2026-09-23; a diferença para os R$ 252,22
anteriores é a recotação da Chain Reaction em 2026-09-23).

---

## v5 — troca pontual (`/swap-card Traxos, Academy Guardian`) · 2026-09-27

O usuário trouxe `Traxos, Academy Guardian` (fora das sobressalentes → compra, R$ 1,89 LigaMagic menor,
2026-09-27) e pediu o deck onde ela rende mais. Corte por cor: Inspirit e Satoru elegíveis (U); Krenko,
Thorin e Phlage fora. Satoru descartado (em `build`; ativador de ninjutsu de 4 manas contra os de 0–2 do
pool). Ficha completa dos dois lados em
[`rounds/v5-2026-09-27/02-theme.md`](rounds/v5-2026-09-27/02-theme.md); relatório em
[`rounds/v5-2026-09-27/report.md`](rounds/v5-2026-09-27/report.md). **Aprovada pelo usuário em 2026-09-27.**

| Data | Carta | Ação | Fase/agente | Motivo | Funções descobertas pelo corte |
|---|---|---|---|---|---|
| 2026-09-27 | Traxos, Academy Guardian | entrada | v5 · swap-card · theme-analyst | Corpo permanente e protegido (artefato → hexproof + indestrutível do comandante) por `{1}{U}`/`{U}` depois de uma mágica não-criatura; 1/5 voadora; anthems de Chief/Master → 3/7; prowess × 35 mágicas não-criatura = poder extra para Station no turno; artifact spell lendária → Pinnacle, Sai, Golem Foundry, Uthros, Jhoira, redutores; nontoken que entra → Surge Conductor, Metastatic Evangel. **Pontos fracos aceitos:** poder base 1; vigilância redundante com a Vraska; atacar com vigilância não conta no X da Alibou | — |
| 2026-09-27 | Stern Lesson | corte | v5 · swap-card · theme-analyst | Fonte de draw de saldo líquido 0 (loot 2/1) e a única do deck que não põe corpo nem motor em campo. A v1 a manteve com draw em 8 e ~7 ramps; hoje draw 13 → **12** e ramp 11 → **10** + 1, as duas metas cumpridas. Estava na lista original do usuário; nunca foi cortada nem reposta | loot → 12 fontes de draw restantes; Powerstone (ramp) → 10 rocks/fontes + 5 redutores; +1 artefato → a própria Traxos; **mágica não-criatura (Vraska/Saheeli/Whirlwind/Seedshark/prowess) 36 → 35 e a única compra em instant speed que gera token → descobertas (custo aceito)** |

**Alternativa não usada:** Talisman of Progress (tiraria aceleração de T2 e dispara mais gatilhos do deck
que a Stern Lesson). **Posse:** Stern Lesson só vira sobressalente se o usuário a listar na próxima
`/update-collection` (regra 7); Traxos tratada como compra — posse não declarada. **Custo da lista:**
R$ 252,95 → R$ 254,74 + Vraska `a cotar` (LigaMagic menor, cotações de 2026-09-18 a 2026-09-27).

## 2026-09-27 — v6 (`/swap-card Room of Refuge`): corte da Temple of Enlightenment reprovado

Proposta do `manabase-engineer`: Temple of Enlightenment → Room of Refuge (tapland por tapland), com o
scry 1 declarado como função descoberta ("custo aceito"). **Reprovada pelo usuário em 2026-09-27:**
*"Temple of Enlightenment faz scry 1, isso ajuda o deck"*. A Temple fica, e o scry dela conta como
função a proteger. Com a v5 aplicada, a Stern Lesson (uma das coberturas de seleção citadas) também
saiu. A proposta voltou ao `manabase-engineer`, contra o `deck.md` pós-v5.

| Data | Carta | Ação | Fase/agente | Motivo | Funções descobertas pelo corte |
|---|---|---|---|---|---|
| 2026-09-27 | Temple of Enlightenment | corte **reprovado** (fica) | v6 · swap-card · manabase-engineer | usuário: o scry 1 ajuda o deck; a seleção de início de jogo não é custo aceitável | — |
| 2026-09-27 | Room of Refuge | entrada **recusada** (v6 fechada sem troca) | v6 · swap-card · manabase-engineer | refeita contra o `deck.md` pós-v5 e com a Temple protegida, nenhuma saída passa. Cycling e filtro têm a mesma régua do scry; básicos, destapados e condicionais levariam os virados de 8 para 9; as Bridges são artefato; a Mystic Monastery fixa melhor; a Vivid Crag é troca lateral. Como 37º terreno, desfaria os 36 e todas as categorias estão no piso. Para voltar a ser proposta, precisa de fato novo: mudança na contagem de virados, falta de R medida em jogo, ou um eixo de sacrificar terreno | — |
