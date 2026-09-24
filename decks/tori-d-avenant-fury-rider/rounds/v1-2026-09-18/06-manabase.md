# Manabase e Cortes — Otharri, Suns' Glory (v1 · 2026-09-19 · modo `improve`)

> Fonte de todo texto oracle: `bin/mtgdb` (dump Scryfall), puxado **nesta sessão** — inclusive o
> dos 6 terrenos atuais e o das cartas que proponho cortar (regra 6). Preços em US$ são
> **estimativa (Scryfall)** e **não** decidem orçamento; a régua é o **menor valor da LigaMagic**
> (regra 2), e 32 das entradas **não têm cotação** (`mtgdb prices` conferido nesta sessão).
>
> **Consulta ao `decisions.md` (regra 5):** o registro tem duas linhas, ambas de 2026-09-19 e ambas
> sobre o comandante (entra Otharri · Tori sai da zona de comando). **Nenhuma** carta deste
> documento — cortada ou proposta — tem histórico de corte e retorno. Nenhuma entrada minha é
> reposição de carta previamente cortada. `Bond of Discipline` é o único caso de vai-e-vem em
> potencial e está tratado explicitamente em §5.

---

## Cálculo

Composição final projetada: **63 não-terrenos** (soma dos CMCs = 168) · **36 terrenos**.

CMC médio: **2,67** · draws+ramps mv≤2: **15–17** · Fórmula:

```
N = 17 (todos):  31,42 + 3,13 × 2,667 − 0,28 × 17 = 35,01
N = 15 (sem os 2 rituais condicionados ao combate): 31,42 + 3,13 × 2,667 − 0,28 × 15 = 35,57
```

**→ 36 terrenos.**

**Como cheguei ao N.** Contei como "ramp/draw mv ≤ 2" as 17 peças que de fato substituem um
terreno na mão: `Sol Ring`, `Arcane Signet`, `Boros Signet`, `Talisman of Conviction`,
`Mind Stone`, `Sphere of the Suns`, `Fellwar Stone`, `Everflowing Chalice`, `Sword of the Animist`,
`Knight of the White Orchid`, `Battle Hymn`, `Great Train Heist`, `Belladonna Took`,
`Recruitment Officer`, `Tome of Legends`, `Idol of Oblivion`, `Staff of the Storyteller`.
`Battle Hymn` e `Great Train Heist` são **rituais condicionados ao combate** — não ajudam a fazer
o land drop nem a jogar na curva — então rodei a conta também sem eles. A diferença é de meio
terreno; a decisão não muda.

**Arredondo 35,0–35,6 para 36, não para 35.** Três motivos, nesta ordem:
1. o comandante custa **5** e o plano inteiro é atacar com ele a partir do T4–T5 — travar em 4
   terrenos custa um ciclo inteiro de contadores de experiência;
2. `Sword of the Animist` busca **terreno básico do grimório** — ela **paga** por uma base
   básica e densa, e recompensa acertar o land drop em vez de puni-lo;
3. a regra do método: na dúvida, para cima. Inundar é menos punitivo que travar.

> ### ⚠ Correção do briefing: o deck **não** está −2 na base
> O briefing registra "**36 terrenos**, −2 sobre 38". **O −2 não existe.** 38 é o ponto de partida
> do método, não o alvo; a fórmula, rodada sobre a composição que as cinco fases propõem, dá
> **35,0–35,6**. O deck já tem **exatamente o número certo de terrenos**.
> **O problema da manabase é 100% de qualidade, 0% de quantidade** — e é grave: **5 dos 6
> não-básicos entram virados** e três deles não fazem absolutamente nada além de produzir mana
> um turno tarde demais. Minha proposta é **6 não-básicos por 6 não-básicos, 30 básicos por 30
> básicos, 36 por 36** — zero slot de feitiço consumido pela manabase, o que devolve os 2 slots
> que o briefing imaginava perdidos.

---

## Terrenos recomendados

### Os 6 não-básicos (1 sai virado, contra 5 hoje)

| Terreno | Produz | Entra virado? | Sinergia/Utilidade | Na coleção? |
|---|---|---|---|---|
| **Command Tower** | `{R}` `{W}` | **não** | intocável (briefing). Melhor terreno do deck | **já no deck** |
| **Battlefield Forge** | `{C}` grátis · `{R}`/`{W}` por 1 de vida | **não** | substitui `Boros Guildgate` 1:1 sem o turno perdido. O modo `{C}` grátis paga o `{3}` genérico do Otharri sem custo nenhum de vida | compra · **R$ 3,79** (LigaMagic menor, 2026-09-18 — 1 dia) |
| **Furycalm Snarl** | `{R}` `{W}` | **quase nunca** — desvirada se você revelar um Mountain/Plains da mão | com **30 básicos** na lista, a condição falha só em mãos sem terreno — que já são mulligan. Substitui `Stone Quarry` | compra · s/ cotação · US$ 0,30 *(est. Scryfall)* |
| **Rustvale Bridge** | `{R}` `{W}` | **sim** | **é `Artifact Land`** → conta para o **metalcraft do `Jor Kadeen`** e para o contador do `Luxknight Breacher`, e é **indestrutível**. É o único terreno virado que compro: troco `Wind-Scarred Crag` (vira + 1 de vida) por "vira + permanece um artefato imune a `Vandalblast`/`Austere Command`" | compra · **R$ 1,09** (LigaMagic menor, 2026-09-18 — 1 dia) |
| **Blast Zone** | `{C}` | **não** | **wipe modular em slot de terreno.** Oracle conferido: entra **com 1 contador de carga** e só sobe com `{X}{X}` — é **fisicamente incapaz de chegar a 0** e portanto **não pode matar as fichas do Otharri (MV 0)**. Responde a prisões, a enxames de 1–2 drops e a `Rhystic Study` | **CAIXA** · R$ 0,89 (já paga) |
| **Karn's Bastion** | `{C}` | **não** | `{4},{T}`: proliferate. **+1 contador de experiência = +1 ficha em todo ataque futuro, para sempre.** E prolifera ainda `Tome of Legends` (páginas), `Wedding Announcement` (convites), `Staff of the Storyteller` (histórias), `Everflowing Chalice` (cargas), `Blast Zone` (cargas) e a lealdade do `Basri Ket` — **6 alvos**, não um. É o *mana sink* de jogo longo que o deck não tem | compra · s/ cotação · US$ 2,39 *(est. Scryfall)* |

**Saldo:** hoje **5 de 6** não-básicos entram virados; na proposta, **1 de 6**. Dois deles
(`Blast Zone`, `Karn's Bastion`) produzem só `{C}` — é o teto que aceito num deck que precisa de
`{R}{W}` e tem `{W}{W}` em cinco cartas. **Por isso recusei um terceiro incolor** (ver §2.2).

### Básicos: **17 Plains + 13 Mountain** (30 — os mesmos 30 que ele já tem)

Contagem de pips do deck final (63 não-terrenos **+ o comandante**): **`{W}` = 41 · `{R}` = 25**
→ 62% branco / 38% vermelho. A proporção pura daria 19/11, e **19/11 é errado**: o vermelho tem
poucos pips **mas todos precisam sair cedo** (`Crash Through` `{R}`, `Flame Slash` `{R}`,
`Vandalblast` `{R}`, `Great Train Heist` `{R}`, `Hexgold Halberd` `{1}{R}`, `Battle Hymn` `{1}{R}`,
`Smite the Deathless` `{1}{R}`, e o bloco inteiro de `{2}{R}`). O branco concentra os **pips
duplos** (`Knight of the White Orchid` `{W}{W}`, `Basri Ket` `{1}{W}{W}`, `Palace Jailer` `{2}{W}{W}`,
`The Circle of Loyalty` `{4}{W}{W}`, `Hour of Reckoning` `{4}{W}{W}{W}`). **17/13** atende os dois.

| | Fontes em terreno | + rochas de cor (Arcane Signet, Boros Signet, Talisman, Commander's Sphere, Sphere of the Suns, Fellwar Stone) | Total efetivo |
|---|---|---|---|
| **Branco** | 17 Plains + 4 duais = **21** | +6 | **~27** |
| **Vermelho** | 13 Mountain + 4 duais = **17** | +6 | **~23** |
| Só incolor | Blast Zone, Karn's Bastion = **2** | — | — |

**Custo:** ele tem 15/15. A mudança pede **comprar 2 Plains** (~R$ 0,40) e deixar 2 Mountain de
sobressalente. Se preferir não mexer, **15/15 é aceitável** — perde ~1 ponto percentual de
consistência em `{W}{W}`, não inverte nenhuma decisão.

### 2.1 Por que **não** compro dual caro

`Sacred Foundry`, `Inspiring Vantage` (US$ 2,54), `Sundown Pass` (US$ 2,67) e `Sunbaked Canyon`
(US$ 4,57) são melhores que `Battlefield Forge`/`Furycalm Snarl` — em um deck de **duas cores com
30 básicos e 6 rochas de fixação**, a melhora é de poucos pontos percentuais, e cada uma custa o
equivalente a 2–4 cartas do pacote de draw ou de interação, que resolvem dores declaradas pelo
usuário. **Dispensa por régua de custo, não por função.** Se sobrar orçamento no fechamento,
`Inspiring Vantage` é a primeira a entrar, no lugar de `Rustvale Bridge`.

### 2.2 Os três terrenos devolvidos a mim — dois entram, dois não

| Devolvido por | Terreno | Veredito |
|---|---|---|
| Fase 2 / Fase 5 | **Blast Zone** | **ENTRA.** Caixa, custo zero, e o risco que se imagina nele não existe (§tabela acima) |
| Fase 2 | **Karn's Bastion** | **ENTRA.** Único terreno do pool que alimenta o motor de contadores de experiência |
| Fase 3 (draw) | **Bonders' Enclave** | **NÃO ENTRA.** `{3},{T}`: compre 1, **só se você controlar criatura de poder 4+**. Duas objeções: (i) seria o **terceiro** terreno que só produz `{C}` numa base que precisa de `{R}{W}` no T4–T5 e de `{W}{W}` em cinco cartas — é exatamente aí que o color screw começa; (ii) `{3}` por carta compete, todo turno, com o mana do plano, e o pacote de draw já entrega **11 fontes dedicadas** com custo marginal ≤ `{1}`. Fica como **primeira reserva de terreno** se o usuário cortar duas fontes de draw por preço |
| Fase 3 (draw) | **Throne of the High City** | **NÃO ENTRA.** `{4},{T}, sacrifique este terreno: você se torna o monarca`. Quarto incolor, **e sacrifica um terreno** num deck que quer chegar a 5–7 terrenos e mantê-los; a coroa já vem do `Palace Jailer` por 4 manas **com corpo e com exílio de uma criatura junto**. Redundância cara em slot de terreno |
| Fase 2 | **Legion's Landing // Adanto, the First Fort** | **NÃO É TERRENO.** É `Legendary Enchantment` `{W}` — só vira terreno depois de transformar. **Não conta nos 36** e não cobre lacuna de manabase nenhuma. Continua ótima como carta de feitiço (ficha no T1 → vira gerador de fichas imune a wipe); devolvo ao orquestrador na **fila de permutáveis** (§4), não na manabase |

### 2.3 Conciliação com `Sword of the Animist` (pedido no prompt)

`Sword of the Animist` busca **terreno básico**, não qualquer terreno. Isso empurra a base para
**mais básicos, não menos** — e é exatamente o que proponho: **30 básicos em 36 terrenos (83%)**,
contra os 30 em 36 de hoje. **Não há conflito a resolver:** a proposta troca **não-básico por
não-básico** e deixa a contagem de básicos intacta. O único ajuste é a proporção interna
(15/15 → 17/13), que não muda o número de alvos da Sword.

Corolário: quanto mais eu enchesse a base de não-básicos (duais caros, utilitários), **pior** a
Sword ficaria. Com 30 básicos ela acha alvo em 100% das partidas até o 4º ativação.

### 2.4 Varredura da caixa (regra 7)

`bin/mtgdb collection -list` rodado nesta sessão: das **113 sobressalentes, exatamente uma é
terreno — `Blast Zone`** — e ela entra. Não há nenhuma dual, nenhum fetch, nenhum utilitário de
RW na caixa. As 4 compras de terreno são compras **porque a função não existe na caixa**, não por
preferência.

---

## 3. O problema aritmético, resolvido

As cinco fases propuseram **~45 entradas**. O deck tem **99 cartas** e **36 delas são terreno**,
o que deixa **63 slots de feitiço**. Como 23 cartas de `lista.txt` merecem ficar, os slots reais
para entradas são **40**, não 45 — e 5 candidatas precisam ficar de fora.

| Bloco | Cartas |
|---|---|
| **Terrenos** | 36 (30 básicos + 6 não-básicos; 1 da caixa, 4 compras) |
| **Ficam de `lista.txt`** | **23** |
| **Entram da caixa** (custo zero) | **6** |
| **Entram por compra** | **32** |
| **Slots permutáveis** (§4) | **2** |
| **Total** | **99** ✅ |

### 3.1 Contagem de metas, com a dupla e tripla função marcada

| Meta | Alvo | Deck final | Cartas |
|---|---|---|---|
| **Draw** | 12 | **11 dedicadas + 2 rochas que viram carta = 13** | Belladonna Took · Recruitment Officer · Tome of Legends · Idol of Oblivion · Staff of the Storyteller *(tb. artefato + ficha)* · Wedding Announcement *(tb. anthem)* · Tocasia's Welcome · Outpost Siege · Palace Jailer *(tb. remoção)* · **Neyali** *(tb. tema + wincon)* · Light Up the Stage ‖ + Mind Stone e Commander's Sphere (sac → compra) |
| **Ramp padrão** | 10–11 | **11** | Sol Ring · Arcane Signet · Boros Signet · Talisman of Conviction · Mind Stone · Commander's Sphere · Sphere of the Suns · Fellwar Stone · Everflowing Chalice · **Sword of the Animist** *(tb. artefato + equipamento)* · Knight of the White Orchid |
| **Ramp explosivo** | 2–3 | **2** | Battle Hymn · **Great Train Heist** *(tb. combate extra = Rota 4)* |
| **Interação** | ~10–14 | **14** | Glass Casket *(tb. artefato)* · Smite the Deathless · Flame Slash · Molten Blast · Generous Gift · Chaos Warp · Celebrate the Mountain-king · Disenchant · Seal of Cleansing · **Palace Jailer** (ETB exila) · Unbreakable Formation · Duty Beyond Death · Gods Willing · Feat of Resistance |
| **Board wipes** | 2–4 | **3 + 1 em terreno** | Hour of Reckoning (poupa as fichas) · Winds of Abandon (unilateral) · Vandalblast (unilateral) ‖ **Blast Zone** (terreno) |
| **Artefatos (metalcraft `Jor Kadeen`)** | 13 | **18 + 1 terreno-artefato = 19** | as 10 rochas/equipamentos de ramp + The Circle of Loyalty · Ancestral Blade · Glass Casket · Tome of Legends · Idol of Oblivion · Staff of the Storyteller · **Molten Gatekeeper** *(tb. wincon + criatura + unearth)* · **Hexgold Halberd** *(tb. ficha Rebel desvirada + trample)* ‖ **Rustvale Bridge** |
| **Rotas de vitória** | 3 | **4** | **R1** alfa com estáticas (Neyali · Jor Kadeen · Circle of Loyalty · Warleader's Call · Radiant Destiny · Valor in Akros · Wedding Festivity) · **R2** dano de *enters* (Warleader's Call · Molten Gatekeeper · Witty Roastmaster) · **R3** furo de bloqueio (Goblin War Drums · Barrage of Boulders · Winds of Abandon · Hour of Reckoning · Crash Through) · **R4** combate extra (Éomer Marshal · Combat Celebrant · Great Train Heist) |

**Metalcraft ligado por padrão a partir do T3** — e `Jor Kadeen` deixa de depender de equipamentos
que estão sendo cortados.

### 3.2 Uma correção de manabase dentro de uma entrada de outra fase

**`Knight of the White Orchid` custa `{W}{W}` no mv 2.** Com **21 fontes brancas em terreno**, a
chance de ter `{W}{W}` no turno 2 é baixa — na prática ela é uma carta de T3–T4, e aí o gatilho
("se um oponente controla mais terrenos") já falha com frequência. **Recomendo trocá-la por
`Loyal Warhound`** `{1}{W}` (3/1 vigilância; ETB busca um Plains **para o campo**, virado, sob a
mesma condição): **um pip só**, mesma função de land-ramp que sobrevive a wipe, corpo com
vigilância que **defende a coroa do `Palace Jailer`**, e US$ 0,66 contra US$ 4,21 *(est.)*.
O próprio `ramp-specialist` já nomeou `Loyal Warhound` como a substituta (§4 do `04-ramp.md`) —
eu só antecipo a troca, pelo motivo que é da minha lente: **exigência de cor**.

---

## 4. Plano de cortes (deck em 99 → 99, com 44 cartas trocadas)

### 4.1 Cortes já protocolados pelas Fases 2–5 — **endosso, não repito ficha** (22)

| Fase | Cortes |
|---|---|
| **Fase 2 · limpos (12)** | Vigilante Justice · Sanctuary Lockdown · Parhelion Patrol · Inspiring Roar · Shoulder to Shoulder · Zealous Display · Knight of Sorrows · Rohirrim Lancer · Cavalry Drillmaster · Lake-town Lookout · **Bond of Discipline** (§5) · Gideon's Triumph |
| **Fase 2 · condicionados, condição **cumprida** (3)** | **True-Faith Censer** (K3) e **Squire's Lightblade** (K15) — a condição era "repor 2 artefatos"; reponho **18**, `Hexgold Halberd` inclusa · **Kwende** (K4) — condição era "Neyali entra"; **entra** |
| **Fase 3 · draw (2)** | Warlord's Fury · Djeru's Renunciation *(I5 da Fase 5 confirmou: com 11+ fontes de draw reais, o cycling é redundante e o corte fica limpo)* |
| **Fase 4 · ramp (1)** | Trailblazer's Torch |
| **Fase 5 · interação (4)** | **Paladin Danse** (condição "Unbreakable Formation + Glass Casket entram" — as duas **entram**) · Adamant Will · Invoke the Divine · Expose to Daylight |

### 4.2 Cortes que **eu** proponho — ficha F1–F7 completa (regra 4)

Condições assumidas **idênticas** para quem sai e para quem entra (§3 do checklist): Otharri em
campo com 3–4 contadores, 6–10 fichas Rebel, `The Circle of Loyalty` e/ou `Jor Kadeen` em campo,
metalcraft ligado.

#### M1 · `Boros Guildgate` · `Stone Quarry` · `Wind-Scarred Crag` — Land — **corte limpo (3)**
- **F1** `This land enters tapped. {T}: Add {R} or {W}.` (`Wind-Scarred Crag` acrescenta `você ganha 1 vida`).
- **F2** Sem corpo. **F3** Terreno; `Boros Guildgate` é `Land — Gate`, e o deck não tem nenhuma carta que leia Gate. **F4** Não recebe nada. **F5** Fixação `{R}`/`{W}`.
- **F6** **Custa um turno inteiro.** Num deck cujo plano é `{3}{R}{W}` no T4–T5 e um ataque por turno a partir dali, entrar virado no T4 empurra o comandante para o T5 e o primeiro contador de experiência para o T6.
- **F7** Três cópias funcionais do mesmo terreno virado, e 1 de vida não é função.
- **Protocolo:** funções = [fonte `{R}`, fonte `{W}`, 1 de vida (Crag), tipo Gate (Guildgate)].
  Fontes → **cobertas e melhoradas** por `Battlefield Forge` e `Furycalm Snarl` (mesmas duas cores,
  **desviradas**) e por `Rustvale Bridge`. 1 de vida → **descoberta**, custo aceito e declarado
  (o comandante tem vínculo com a vida). Tipo Gate → nunca foi lida por nada. **Corte limpo.**

#### M2 · `Fields of Strife` — Land — **corte limpo**
- **F1** `This land enters tapped. {T}: Add {R} or {W}. {2}{R}{W}, {T}: Surveil 1.`
- **F2–F4** Sem corpo, terreno, não recebe nada. **F5** Fixação `{R}`/`{W}` + surveil.
- **F6** Entra virada. **F7** **Atrito fatal:** a ativação custa `{2}{R}{W}` **e** vira o próprio terreno — são **4 manas, incluindo os dois coloridos do comandante**, por um surveil 1. É literalmente o custo do Otharri. Não existe turno neste deck em que esse botão seja a melhor jogada.
- **Protocolo:** funções = [fonte RW, surveil repetível]. Fonte → coberta por `Furycalm Snarl`.
  Surveil → **descoberta**, custo aceito: o pacote de draw entrega 11 fontes de carta de verdade;
  trocar 4 manas por um surveil é o pior preço de filtragem do deck. **Corte limpo.**

#### M3 · `Sandstone Bridge` — Land — **corte limpo** (condicionado K18 da Fase 2, agora resolvido)
- **F1** `This land enters tapped. When this land enters, target creature gets +1/+1 and gains vigilance until end of turn. {T}: Add {W}.`
- **F2–F4** Sem corpo, terreno, não recebe nada. **F5** +1/+1 e **vigilância** até o fim do turno a **uma** criatura.
- **F6** Entra virada. **F7** Produz **só `{W}`** — é a pior das seis para pagar `{3}{R}{W}`. E o ETB foi desenhado para o eixo de untap da Tori, que saiu da zona de comando.
- **Protocolo:** funções = [fonte `{W}`, buff de alvo único, vigilância de alvo único por um turno].
  Fonte → coberta por Plains e pelas três duais desviradas. Buff → coberto por `Basri's Solidarity`,
  `Duty Beyond Death`, `Valor in Akros`. **Vigilância → coberta e ampliada por `Radiant Destiny`
  nomeando Rebel**, que dá vigilância ao **enxame inteiro, permanentemente** — e é justamente o que
  destrava a recursão `{2}{R}{W}` do Otharri (sempre há Rebel desvirado). **Nenhuma função
  descoberta. K18 resolvido: corte limpo.**

#### M4 · `Inspiring Veteran` `{R}{W}` 2/2 — **corte limpo** (K5 da Fase 2, condição resolvida pela contagem)
- **F1** `Other Knights you control get +1/+1.` **F2** Corpo 2/2 por 2, RW. **F3** Human Knight → alimenta o affinity do `The Circle of Loyalty`. **F4** Recebe anthems e contadores. **F5** Anthem typal de Knight. **F6** Turno 2. **F7** As fichas são **Rebel**.
- **A contagem que faltava.** A Fase 2 escreveu a regra: *"≤10 Knights, sai; ≥14, fica"*. O deck
  final tem **3 Knights** (`Dawnstrike Vanguard`, `Éomer, Marshal of Rohan`, `Knight of the White
  Orchid`) — 4 com a Tori. O anthem dele buffa **duas** criaturas. **Condição resolvida por
  margem enorme: sai.**
- **Protocolo:** funções = [anthem typal, corpo RW mv 2, contagem Knight]. Anthem → coberto e
  ampliado por `The Circle of Loyalty`, `Radiant Destiny`, `Warleader's Call` (todos leem "creatures
  you control"). Corpo mv 2 RW → coberto por `Hexgold Halberd` (2 manas, **traz um corpo 2/2
  vermelho desvirado** e é artefato). Contagem Knight → **encolhe e está declarada**: `The Circle
  of Loyalty` passa a custar ~3 manas em vez de ~2. Custo aceito.

#### M5 · `Adriana, Captain of the Guard` `{3}{R}{W}` 4/4 — **corte limpo**
- **F1** `Melee.` / `Other creatures you control have melee.` **F2** Corpo 4/4; com Circle + Jor Kadeen é 8/5. **F3** Human Knight legendária → ficha do Circle. **F4** Recebe tudo. **F5** Concede melee. **F6** Turno 5 — **o turno do comandante**.
- **F7** **Atrito fatal, já diagnosticado pela Fase 2 e confirmado no ruling (2016-08-23):** as
  fichas do Otharri **entram atacando mas nunca são declaradas atacantes**, então **melee não
  dispara para elas**. A cláusula que justifica o slot não enxerga 80% do board futuro. E compete
  de frente com o Otharri no mv 5.
- **Protocolo:** funções = [corpo 4/4, melee própria, **concessão de melee**, contagem Knight/legendária].
  Corpo e contagens → cobertos por `Neyali` (3/3 legendária, Human **Rebel**). Concessão de melee →
  **coberta e superada por `Neyali`** (`attacking tokens have double strike` — **estática**, alcança
  as fichas que entram atacando, que é exatamente o que melee não faz). Melee própria → descoberta,
  custo declarado: um bônus de +3/+3 num corpo só, por 5 manas.

#### M6 · `Syr Alin, the Lion's Claw` `{3}{W}{W}` 4/4 — **corte limpo com custo declarado**
- **F1** `First strike.` / `Whenever Syr Alin attacks, other creatures you control get +1/+1 until end of turn.` **F2** Corpo 4/4 com first strike. **F3** Human Knight legendária. **F4** Recebe tudo. **F5** Pump em massa no ataque — **alcança as fichas** por ordenação de gatilhos (§1.3b da Fase 2). **F6** Turno 5.
- **F7** `{W}{W}` no mv 5 na faixa mais congestionada; o efeito é **temporário** e exige **ela** atacando (expõe um 4/4 a bloqueio e remoção).
- **Protocolo:** funções = [corpo, first strike, **pump em massa no ataque**, contagem].
  Pump → **coberto e superado, três vezes**: `Valor in Akros` (dispara **N vezes**, uma por ficha),
  `Warleader's Call` e `Radiant Destiny` (+1/+1 **estático e permanente**, sem precisar atacar).
  First strike → descoberto e declarado dispensável (o pagamento, `Kwende`, já saiu).
  Corpo/contagem → cobertos por `Neyali` e `Hero of Bladehold` (fila §4).

#### M7 · `Éomer of the Riddermark` `{4}{R}` 5/4 — **corte limpo**
- **F1** `Haste.` / `Whenever Éomer attacks, **if you control a creature with the greatest power among creatures on the battlefield**, create a 1/1 white Human Soldier token.` **F2** Corpo 5/4 com ímpeto. **F3** Human Knight legendária. **F4** Recebe tudo. **F5** 1 ficha por ataque, condicionada. **F6** Turno 5.
- **F7** A cláusula-*if* é **verificada duas vezes** (ao disparar e ao resolver) e depende do board **da mesa inteira** — num pod com um fatty de 6+ poder, o gatilho simplesmente não acontece. Produz **1** ficha por ataque contra as **N** do comandante no mesmo passo.
- **Protocolo:** funções = [corpo 5/4 com ímpeto, ficha condicionada por ataque, contagem legendária].
  Ficha por ataque → **coberta e multiplicada** pelo próprio Otharri e por `Hero of Bladehold`
  (2 fichas, **sem condição**). Corpo + ímpeto → coberto por `Combat Celebrant` (4/1 por 3, que
  ainda dá **combate extra**). Contagem legendária → coberta por Neyali/Jor Kadeen/Belladonna.

#### M8 · `Fireborn Knight` `{R/W}{R/W}{R/W}{R/W}` 2/3 — **corte limpo**
- **F1** `Double strike.` / `{R/W}×4: +1/+1 até o fim do turno.` **F2** Corpo 2/3 com double strike; com Circle + Jor Kadeen é 6/4 double strike = 12 de dano. **F3** Human Knight. **F4** Recebe tudo. **F5** Nada a outros. **F6** Turno 4, custo híbrido flexível (bom). **F7** O *mana sink* custa **4 manas por +1/+1** — o pior preço de mana do deck; e double strike **num corpo só** contra `Neyali`, que dá double strike a **todo o enxame**.
- **Protocolo:** funções = [corpo, **double strike**, flexibilidade de cor, mana sink, contagem Knight].
  Double strike → **coberto e multiplicado por `Neyali`** (todas as fichas atacantes). Mana sink →
  coberto por `Karn's Bastion` (proliferate) e `Idol of Oblivion` (`{8}`: 10/10). Flexibilidade de
  cor → **descoberta e declarada**: é a única carta do deck castável só com híbrido. Custo aceito —
  a base 17/13 + 6 rochas de cor resolve o problema que ela contornava.

#### M9 · `Relentless Rohirrim` `{3}{R}` 4/3 — **corte limpo**
- **F1** `When this creature enters, the Ring tempts you.` **F2** Corpo 4/3. **F3** Human Knight. **F4** Recebe tudo. **F5** Nada. **F6** Turno 4. **F7** O deck **não tem nenhuma outra carta de the Ring** — o Portador do Anel dá "não pode ser bloqueada por criatura de poder maior" a **uma** criatura, e nada mais. É um 4/3 sem texto relevante no turno mais disputado da curva.
- **Protocolo:** funções = [corpo 4/3, contagem Knight/Human, evasão parcial de um alvo].
  Corpo → coberto pelo motor de fichas e por `Combat Celebrant`/`Hero of Bladehold`. Evasão →
  **coberta e ampliada** por `Goblin War Drums` (menace ao time) e `Crash Through` (trample ao time).
  Contagem → coberta. **Nenhuma função real perdida.**

#### M10 · `Fervent Cathar` `{2}{R}` 2/1 — **corte limpo com custo declarado**
- **F1** `Haste.` / `When this creature enters, target creature can't block this turn.` **F2** 2/1 com ímpeto. **F3** Human Knight. **F4** Recebe tudo. **F5** Remove **um** bloqueador, **uma** vez. **F6** Turno 3. **F7** Remover 1 bloqueador contra um board de 8–10 atacantes resolve ~10% do problema.
- **Protocolo:** funções = [corpo com ímpeto, contagem, **furo de bloqueio pontual**].
  Furo → **coberto e ampliado por `Barrage of Boulders`** (*ferocious*: **nenhuma** criatura pode
  bloquear) e por `Goblin War Drums` (menace permanente). Corpo com ímpeto → coberto por
  `Combat Celebrant` (mesmo mv, corpo maior, combate extra).

#### M11 · `Inspiring Captain` `{3}{W}` 3/3 — **corte limpo**
- **F1** `When this creature enters, creatures you control get +1/+1 until end of turn.` **F2** 3/3. **F3** Human Knight. **F4** Recebe tudo. **F5** Anthem **one-shot**. **F6** Turno 4. **F7** O anthem dura **um turno** e dispara **uma vez na vida**; `Valor in Akros`, no mesmo custo, dispara **uma vez por criatura que entra** — com 4 contadores, **4 vezes num único combate**, para sempre.
- **Protocolo:** funções = [corpo 3/3, contagem, anthem temporário one-shot]. Todas cobertas —
  a terceira, **estritamente superada por `Valor in Akros`, que já está no deck**.

#### M12 · `Knight Luminary` `{3}{W}` 3/2 — **corte limpo com custo declarado**
- **F1** `When this creature enters, create a 1/1 white Human Soldier token.` / `Warp {1}{W}`. **F2** 3/2 + ficha. **F3** Human Knight. **F4** Recebe tudo. **F5** 1 ficha (dispara `Belladonna Took`/`Valor in Akros`/`Warleader's Call`). **F6** Turno 4, ou **turno 2 via warp** (e recompra depois) — o melhor eixo dela. **F7** Dois corpos por 4 manas num deck que fabrica 4 corpos por ataque de graça.
- **Protocolo:** funções = [corpo, **ficha que liga os gatilhos de *enters***, contagem, warp].
  Gatilho de *enters* → coberto com folga por `Ancestral Blade` (2 manas, ficha **e** artefato),
  `Staff of the Storyteller` (2 manas, ficha **e** saque), `Raise the Alarm` (2 fichas em instant)
  e `Hexgold Halberd`. Warp → **descoberto**, custo declarado: é a única carta com warp do deck.

#### M13 · `Sheriff of Safe Passage` `{2}{W}` 0/0 — **corte limpo com custo declarado**
- **F1** Entra com um contador +1/+1 **mais um por outra criatura que você controla**. / `Plot {1}{W}`. **F2** Corpo que escala: com 6 fichas é um 7/7 por 3 manas. **F3** Human Knight. **F4** Recebe anthems **e contadores**. **F5** Nada. **F6** Turno 3, ou **plot** no T3 para sair de graça depois. **F7** **Atrito estrutural:** ela só é grande **se o board já está grande** — ou seja, ela é ótima quando você já está ganhando e é um 1/1 quando você acabou de tomar um wipe. É a definição de win-more.
- **Protocolo:** funções = [corpo escalonado, contagem, **plot como guarda contra wipe**, receptor de contadores].
  Corpo → coberto. Guarda contra wipe → **coberta e superada** pelos **contadores de experiência**
  (ficam no jogador) + `Molten Gatekeeper` (**unearth `{R}`**) + `Assemble the Legion`/
  `Wedding Announcement` (não-criaturas que reconstroem). Receptor de contadores → coberto
  (`Karn's Bastion` prolifera 6 alvos melhores).

#### M14 · `Luxknight Breacher` `{3}{W}` 2/2 — **corte limpo**
- **F1** Entra com um contador +1/+1 **por outra criatura e/ou artefato** que você controla. **F2** Corpo que escala; com 6 fichas + 10 artefatos é um 18/18. **F3** Human Knight. **F4** Recebe anthems e contadores. **F5** Nada. **F6** Turno 4. **F7** Mesmo atrito do M13, **agravado**: um 18/18 sozinho morre para `Swords to Plowshares` e não tem evasão. O deck ganha por **largura**, e concentrar 18 pontos de poder num corpo removível é o oposto da tese — e da dor 4.
- **Protocolo:** funções = [corpo escalonado, contagem artefato+criatura, receptor de contadores].
  Todas cobertas: o mesmo poder distribuído em 6 fichas 5/2 (com `Jor Kadeen`) é **imune a remoção
  pontual**. Nenhuma função descoberta.

#### M15 · `Elite Interceptor // Rejoinder` `{W}` 1/2 — **corte limpo com custo declarado**
- **F1** Entra *prepared*; enquanto preparada, você pode conjurar uma cópia de `Rejoinder` (`{1}{W}`: vire **ou desvire** uma criatura alvo; compre 1). **F2** Corpo 1/2 por 1. **F3** Human Wizard — **não é Knight**, não alimenta nenhuma contagem do deck. **F4** Recebe anthems. **F5** **Desvirar** — pode destravar a recursão Rebel do Otharri. **F6** Turno 1. **F7** `Rejoinder` é **troca de mana por carta**, não vantagem; e o modo de desvirar é **uma vez**.
- **Protocolo:** funções = [corpo mv 1, **desvirar um Rebel**, cantrip].
  Desvirar → **coberto e tornado permanente por `Radiant Destiny`** (vigilância às fichas: **sempre**
  há Rebel desvirado) e por `Hexgold Halberd` (ficha Rebel **desvirada** ao entrar). Cantrip →
  **descoberto**, custo aceito e declarado (o pacote de draw troca 1-por-1 por saque recorrente).
  Corpo mv 1 → coberto pelo motor.

#### M16 · `Miraculous Recovery` `{4}{W}` — **corte limpo**
- **F1** `Return target creature card from your graveyard to the battlefield. Put a +1/+1 counter on it.` **F2–F4** Sem corpo, instantâneo sem contagem, não recebe nada. **F5** Reanima **uma** criatura. **F6** Turno 5 — o turno do comandante. **F7** **Redundância tripla:** o Otharri se devolve sozinho por `{2}{R}{W}` + virar um Rebel (**sem gastar carta e quantas vezes quiser**), `Late to Dinner` o devolve **desvirado** por 4, e `Molten Gatekeeper` tem **unearth `{R}`**.
- **Protocolo:** função única = [reanimação de 1 criatura] → **coberta três vezes**, duas delas mais
  baratas. O único diferencial (ser instantâneo) não paga 5 manas. Nenhuma função descoberta.

#### M17 · `Remember the Fallen` `{2}{W}` — **corte limpo**
- **F1** Escolha um ou ambos: devolva um card de criatura **e/ou** de artefato do cemitério **para a mão**. **F2–F4** Sem corpo, feitiço, não recebe nada. **F5** Recursão **para a mão** (você ainda paga o custo depois). **F6** Turno 3. **F7** Recursão para a mão é 3 manas por **zero efeito no board** — e o deck tem 19 artefatos justamente porque eles são baratos e substituíveis.
- **Protocolo:** funções = [recursão de criatura, recursão de artefato].
  Criatura → coberta por `Late to Dinner`, pela recursão do próprio Otharri e pelo unearth do
  `Molten Gatekeeper`. Artefato → **descoberta**, custo declarado e aceito: com **19 artefatos**,
  perder um não desliga o metalcraft (3 bastam), e `Vandalblast`/`Austere Command` são as ameaças
  que importam — e contra elas a resposta é `Rustvale Bridge` (terreno-artefato **indestrutível**,
  que sozinho mantém a contagem viva depois de um `Vandalblast` adversário), não recomprar uma
  peça do cemitério por 3 manas.

#### M18 · `Swift Reckoning` `{1}{W}` — **corte limpo com custo declarado**
- **F1** `Spell mastery — se houver 2+ instantâneos/feitiços no seu cemitério, pode conjurar como se tivesse flash.` / `Destroy target **tapped** creature.` **F2–F4** Sem corpo, feitiço, não recebe nada. **F5** Destrói criatura **virada**. **F6** Turno 2. **F7** **Só pega criatura virada** — ou seja, só funciona **depois** que o oponente ataca, ou contra criaturas com habilidade de tap. Contra o alvo que mais importa (o comandante do oponente parado bloqueando, um `Rhystic Study`, um fatty defensivo), **não faz nada**.
- **Protocolo:** função única = [remoção condicionada de criatura]. **Coberta e superada** por
  `Glass Casket` (exila MV≤3, sem condição, **e é artefato**), `Smite the Deathless` (instant,
  fura indestrutível, exila), `Flame Slash` (4 de dano por `{R}`), `Generous Gift` e `Chaos Warp`.
  Custo declarado: o deck perde uma remoção de 2 manas e ganha seis melhores. A Fase 5 já havia
  sinalizado esta carta como um dos dois slots mais fracos da categoria, **deixando o corte comigo**.

#### M19 · `Reduce to Memory` `{1}{W}{W}` — **corte limpo com custo declarado**
- **F1** `Exile target nonland permanent. Its controller creates a 3/2 red and white Spirit creature token.` **F2–F4** Sem corpo, feitiço-Lesson, não recebe nada. **F5** Exílio de **qualquer** não-terreno. **F6** Turno 3. **F7** Duas objeções: (i) é **feitiço** — não responde a nada; (ii) `{W}{W}` no mv 3, e **devolve um 3/2 ao oponente** — num deck que quer atacar, dar um bloqueador 3/2 de brinde é caro.
- **Protocolo:** função = [remoção flexível de não-terreno]. **Coberta e superada por `Generous Gift`**
  (instant, destrói **qualquer permanente incluindo terreno**, devolve 3/3 — pior corpo, mas em
  velocidade de instantâneo) e por `Chaos Warp` (instant, qualquer permanente, **sem compensação**).
  A Fase 5 sinalizou e deixou o corte comigo. Custo declarado: sai a 3ª remoção flexível, ficam 4.

### 4.3 Resumo do plano de cortes

| # | Corte proposto | CMC | Motivo |
|---|---|---|---|
| 1–12 | Vigilante Justice · Sanctuary Lockdown · Parhelion Patrol · Inspiring Roar · Shoulder to Shoulder · Zealous Display · Knight of Sorrows · Rohirrim Lancer · Cavalry Drillmaster · Lake-town Lookout · Bond of Discipline · Gideon's Triumph | 4·3·4·4·3·3·5·1·2·1·5·2 | Fase 2, cortes limpos com ficha — endosso |
| 13–15 | True-Faith Censer · Squire's Lightblade · Kwende, Pride of Femeref | 2·1·4 | Fase 2, condicionados — **condições cumpridas** (19 artefatos · Neyali entra) |
| 16–17 | Warlord's Fury · Djeru's Renunciation | 1·2 | Fase 3 (e I5 da Fase 5) — cantrip/cycling trocados por saque real |
| 18 | Trailblazer's Torch | 4 | Fase 4 — mv 4 no slot do comandante; contagem de artefato reposta 18× |
| 19–22 | Paladin Danse · Adamant Will · Invoke the Divine · Expose to Daylight | 3·2·3·3 | Fase 5 — condição do Paladin cumprida (Unbreakable Formation + Glass Casket entram) |
| **23–25** | **Boros Guildgate · Stone Quarry · Wind-Scarred Crag** | 0 | **M1** — entram viradas sem upside; trocadas por Battlefield Forge e Furycalm Snarl (desviradas) e Rustvale Bridge (virada, mas **artefato indestrutível**) |
| **26** | **Fields of Strife** | 0 | **M2** — entra virada e o botão custa `{2}{R}{W}` (o custo do próprio comandante) por um surveil |
| **27** | **Sandstone Bridge** | 0 | **M3** (K18 resolvido) — entra virada, produz **só `{W}`**, e a vigilância de alvo único é coberta por Radiant Destiny no enxame inteiro |
| **28** | **Inspiring Veteran** | 2 | **M4** (K5 resolvido) — o deck final tem **3 Knights**; o anthem buffa 2 criaturas |
| **29** | **Adriana, Captain of the Guard** | 5 | **M5** — melee **não dispara** para fichas que entram atacando (ruling 2016-08-23) |
| **30** | **Syr Alin, the Lion's Claw** | 5 | **M6** — pump temporário superado por Valor in Akros (N disparos) e pelos anthems estáticos |
| **31** | **Éomer of the Riddermark** | 5 | **M7** — 1 ficha por ataque, **condicionada ao maior poder da mesa** |
| **32** | **Fireborn Knight** | 4 | **M8** — double strike num corpo só; Neyali dá ao enxame inteiro |
| **33** | **Relentless Rohirrim** | 4 | **M9** — 4/3 sem texto relevante (o deck não tem outro card de the Ring) |
| **34** | **Fervent Cathar** | 3 | **M10** — tira **1** bloqueador; Barrage of Boulders tira **todos** |
| **35** | **Inspiring Captain** | 4 | **M11** — anthem one-shot estritamente superado por `Valor in Akros`, que já está no deck |
| **36** | **Knight Luminary** | 4 | **M12** — 2 corpos por 4 manas num deck que fabrica 4 por ataque |
| **37** | **Sheriff of Safe Passage** | 3 | **M13** — win-more: grande só quando o board já está grande |
| **38** | **Luxknight Breacher** | 4 | **M14** — concentra 18 de poder num corpo removível; anti-tese de um deck de largura |
| **39** | **Elite Interceptor // Rejoinder** | 1 | **M15** — Rejoinder é mana por carta; o desvirar vira permanente com Radiant Destiny |
| **40** | **Miraculous Recovery** | 5 | **M16** — reanimação redundante (Otharri se devolve sozinho por 4, sem gastar carta) |
| **41** | **Remember the Fallen** | 3 | **M17** — recursão **para a mão** por 3 manas, zero efeito no board |
| **42** | **Swift Reckoning** | 2 | **M18** — só mata criatura **virada** |
| **43** | **Reduce to Memory** | 3 | **M19** — feitiço `{W}{W}` que **dá um 3/2 ao oponente**; Generous Gift e Chaos Warp fazem melhor em instant |

**Total: 43 cortes** (22 endossados + 21 meus, sendo 5 de terreno) · **+ 1 carta de decisão do
usuário** (`Esgaroth Garrison`, §4.4) = até 44 slots liberados.

### 4.4 As duas decisões do usuário — **os 2 slots permutáveis**

O deck fecha em **97 cartas de núcleo + 2 slots permutáveis + o comandante = 100**. Os dois slots
absorvem as duas respostas sem exigir nenhum recálculo: **o CMC médio (2,67) e a conta de
terrenos (36) são idênticos com ou sem elas** — foi assim que montei a lista.

**Fila de preenchimento dos 2 slots, por prioridade:**

| # | Carta | Custo | Por que nesta posição |
|---|---|---|---|
| **P1** | **Tori D'Avenant, Fury Rider** `{1}{R}{R}{W}` 3/3 | **R$ 0** (é dele) | **Recomendo que fique.** É a **2ª fonte de trample em massa** do deck (a 1ª é `Crash Through`, que é *one-shot*), e trample é a resposta ao chump-block que trava fichas 2/2 terrestres. Ordenando o gatilho do Otharri por último na pilha, o efeito dela **alcança as fichas recém-criadas** — e elas são **vermelhas**, que é exatamente o público da cláusula. 3 pontos de sinergia, custo zero, e corpo 3/3 com **vigilância** que bloqueia o crack-back. Atrito declarado: `{R}{R}` no mv 4, a exigência de cor mais dura do deck depois do `Hour of Reckoning` |
| **P2** | **Esgaroth Garrison** `{4}{W}` */5 | **R$ 0** (é dele) | **Fica se o usuário quiser o bloqueador.** Corpo = nº de criaturas → 8/5–11/5 com o enxame, e **resistência 5** é a maior do deck: é o único corpo que segura um contra-ataque sem morrer. O *recruit* é loot (não conta como draw — a Fase 3 já declarou). Atrito: mv **5**, competindo com o turno do comandante, e ela **não faz nada** no turno em que o board está vazio |
| **P3** | **Syr Carah, the Bold** `{3}{R}{R}` 3/3 | **R$ 0** (já no deck) | Entra se o usuário dispensar P1/P2. É a **12ª fonte dedicada de draw** (impulse repetível pelo `{T}`, mesmo sem atacar). Atrito: mv 5 e o `{T}` compete com atacar |
| **P4** | **Bond of Discipline** `{4}{W}` | **R$ 0** (já no deck) | **Volta apenas se `Goblin War Drums` E `Barrage of Boulders` caírem por preço** (§5) |
| **P5** | **Hero of Bladehold** `{2}{W}{W}` 3/4 | compra · US$ 0,61 *(est.)* | 2 fichas viradas e atacando **por ataque**, sem condição + battle cry que alcança as fichas por ordenação |
| **P6** | **Intangible Virtue** `{1}{W}` | compra · US$ 0,32 *(est.)* | +1/+1 **e vigilância a fichas** — redundância barata do `Radiant Destiny` |
| **P7** | **Legion's Landing // Adanto** `{W}` | compra · US$ 3,46 *(est.)* | Ficha no T1 → vira **terreno gerador de fichas** imune a wipe. Excelente, mas é o 3º efeito de "reconstrói sozinho" |
| **P8** | **Fumigate** `{4}{W}` | **R$ 0** (caixa) | 4º wipe, se o usuário quiser mais resposta de massa |

---

## 5. O desacordo `Bond of Discipline` — decidido

O `wincon-tester` pediu **manter** `Bond of Discipline` **se** `Goblin War Drums` e
`Barrage of Boulders` não entrarem. **As duas entram no núcleo.** Portanto:

**Veredito: `Bond of Discipline` sai (K16 confirmado), e a condição está cumprida por escrito.**

Aplicando as **mesmas** condições aos dois lados (6–10 fichas, um anthem em campo):

| | **Sai:** Bond of Discipline `{4}{W}` | **Entram:** Goblin War Drums `{2}{R}` + Barrage of Boulders `{2}{R}` |
|---|---|---|
| Custo | 5 manas, **uma vez** | 3 manas cada; o War Drums é **permanente** |
| Efeito | vira todas as criaturas dos oponentes; suas criaturas ganham vínculo | menace **estático e permanente** ao time inteiro + um turno em que **ninguém pode bloquear** e os x/1 do oponente morrem |
| Turno | **5** — colide de frente com o turno do Otharri | **3** — e sobram 2 manas para `Gods Willing`/`Feat of Resistance` |
| Resiliência | feitiço: usa-se e acaba | `Goblin War Drums` é **encantamento** — sobrevive ao wipe e continua valendo quando o Otharri reconstrói |

- **Falter em massa** → coberto por `Barrage of Boulders` (mais barato, mesmo turno) e
  **substituído por algo melhor**: menace permanente, que não gasta carta nenhuma.
- **Vínculo em massa** → **descoberto**. Custo declarado e aceito: o comandante tem vínculo, e o
  deck corre para vencer no T6–T7.

> **Gatilho de reversão, escrito:** se o orquestrador cortar `Goblin War Drums` **e**
> `Barrage of Boulders` por preço na LigaMagic, `Bond of Discipline` **volta** pelo slot P4 — é
> custo zero e o deck **não pode ficar com zero furo de bloqueio** (é a dor nº 1 do usuário). Se
> cair **só uma** das duas, a outra basta e o Bond não volta.

---

## 6. Prioridade de compra — núcleo inegociável × cauda sacrificável

> **Nenhuma afirmação de orçamento.** Das 36 compras, apenas **9 têm cotação LigaMagic**
> (`mtgdb prices`, nesta sessão): Battlefield Forge R$ 3,79 · Rustvale Bridge R$ 1,09 · Boros
> Signet R$ 1,50 · Mind Stone R$ 2,25 · Everflowing Chalice R$ 0,99 · Battle Hymn R$ 6,89 ·
> Chaos Warp R$ 3,81 · Outpost Siege R$ 0,90 · Idol of Oblivion R$ 8,94 · Light Up the Stage
> R$ 0,74. **As outras 27 precisam de captura manual antes de qualquer total** (regra 2). Os
> básicos e as 7 cartas da caixa (`Blast Zone`, `Commander's Sphere`, `Sphere of the Suns`,
> `Glass Casket`, `Smite the Deathless`, `Flame Slash`, `Molten Blast`) já estão pagos.

### 6.1 Núcleo inegociável (24 compras) — se uma destas cair, uma **meta** quebra

| Função que quebra | Cartas |
|---|---|
| **Rota 1 (multiplicador estático)** | Neyali, Suns' Vanguard · Warleader's Call · Radiant Destiny |
| **Rota 2 (dano de *enters*)** | Molten Gatekeeper *(+ artefato + unearth)* |
| **Rota 3 (furo de bloqueio)** | Goblin War Drums · Barrage of Boulders |
| **Rota 4 (combate extra)** | Combat Celebrant |
| **Draw imune a wipe (dor 2 + dor 4)** | Tome of Legends · Staff of the Storyteller · Wedding Announcement · Tocasia's Welcome · Outpost Siege · Palace Jailer |
| **Ramp mv 2 (Otharri no T4)** | Boros Signet · Talisman of Conviction · Mind Stone · Fellwar Stone · Everflowing Chalice · Sword of the Animist |
| **Proteção / wipe (dor 3 + dor 4)** | Unbreakable Formation · Winds of Abandon · Hour of Reckoning · Vandalblast |
| **Manabase** | Battlefield Forge · Furycalm Snarl |

### 6.2 Cauda sacrificável — **ordem exata de corte** se o total estourar R$ 200

| Ordem | Sai | Substituto (custo R$ 0, salvo indicação) | O que se perde |
|---|---|---|---|
| **1º** | **Idol of Oblivion** (R$ 8,94 — a mais cara cotada) | **Oracle's Vault** (caixa) | saque de custo marginal zero vira impulse de `{2}`; draw continua em 11 |
| **2º** | **Battle Hymn** (R$ 6,89) | **Manalith** ou **Hedron Crawler** (caixa) | ramp explosivo cai para 1; ramp padrão vira 12 |
| **3º** | **Witty Roastmaster** | — (nada) | 2º pingador da Rota 2; o 1º (Molten Gatekeeper) é inegociável e a rota sobrevive |
| **4º** | **Chaos Warp** (R$ 3,81) | **Pinecone Strike** ou **Molten Blast** (caixa) | única resposta **vermelha** a encantamento |
| **5º** | **Knight of the White Orchid** | **Loyal Warhound** (US$ 0,66) **ou Myr Convert** (caixa) | já recomendado em §3.2 por exigência de cor — esta troca **deveria ser feita de qualquer forma** |
| **6º** | **Generous Gift** | **Joust** ou **Smite the Deathless** (caixa) | "destrói qualquer permanente" em instant |
| **7º** | **Hexgold Halberd** | **Ancestral Blade** (já no deck) | Rebel desvirado avulso; `Radiant Destiny` já cobre por vigilância |
| **8º** | **Battlefield Forge** (R$ 3,79) | **Boros Guildgate** (volta do corte, R$ 0) | 1 terreno volta a entrar virado |
| **9º** | **Karn's Bastion** | 1 Mountain básico | proliferate do contador de experiência |
| **10º** | **Great Train Heist** | — | ramp explosivo vai a 0; `Combat Celebrant` sustenta a Rota 4 sozinho |
| **11º** | **Barrage of Boulders** | **Bond of Discipline** (já no deck, §5) | — (a função volta com custo zero) |

**Leitura:** cortando a cauda inteira (11 itens), o deck perde **~R$ 30 em cotações conhecidas e
~US$ 20 em estimativas**, **não quebra nenhuma meta** e continua com as 4 rotas. É essa a margem
de manobra do orquestrador depois da captura de preços.

---

## Curva final projetada

**0–1: 8 · 2: 26 · 3: 16 · 4: 8 · 5: 2 · 6+: 3** — 63 não-terrenos · CMC médio **2,67**

| mv | Nº | Cartas |
|---|---|---|
| **0–1** | 8 | Everflowing Chalice (0) · Sol Ring · Crash Through · Gods Willing · Recruitment Officer · Flame Slash · Vandalblast · Great Train Heist |
| **2** | 26 | Arcane Signet · Boros Signet · Talisman of Conviction · Mind Stone · Fellwar Stone · Sphere of the Suns · Sword of the Animist · Knight of the White Orchid *(→ Loyal Warhound)* · Battle Hymn · Tome of Legends · Idol of Oblivion · Staff of the Storyteller · Ancestral Blade · Hexgold Halberd · Glass Casket · Smite the Deathless · Belladonna Took · Duty Beyond Death · Feat of Resistance · Disenchant · Seal of Cleansing · Basri's Solidarity · Pride of Conquerors · Raise the Alarm · Honored Crop-Captain · Winds of Abandon |
| **3** | 16 | Commander's Sphere · Basri Ket · Wedding Announcement · Tocasia's Welcome · Light Up the Stage · Molten Blast · Generous Gift · Chaos Warp · Unbreakable Formation · Warleader's Call · Molten Gatekeeper · Witty Roastmaster · Goblin War Drums · Barrage of Boulders · Radiant Destiny · Combat Celebrant |
| **4** | 8 | Valor in Akros · Late to Dinner · Éomer, Marshal of Rohan · Celebrate the Mountain-king · Outpost Siege · Palace Jailer · Neyali, Suns' Vanguard · **[slot permutável 1 — Tori D'Avenant]** |
| **5** | 2 | Jor Kadeen, the Prevailer · **[slot permutável 2 — Esgaroth Garrison]** |
| **6+** | 3 | The Circle of Loyalty (6) · Dawnstrike Vanguard (6) · Hour of Reckoning (7, convoke) |

**Comparação com hoje:** a curva atual tem **15 cartas em mv 4 e 9 em mv 5**. A projetada tem
**8 e 2**. **34 das 63 não-terrenos custam ≤ 2** — o que é o que torna `Otharri` no T4 uma jogada
rotineira em vez de excepcional, e o que faz a fórmula de terrenos fechar em 36 em vez de 38.
