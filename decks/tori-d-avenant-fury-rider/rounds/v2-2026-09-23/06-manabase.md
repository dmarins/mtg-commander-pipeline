# Manabase e Cortes — Phlage, Titan of Fire's Fury (v2 · controle de atrito RW)

> **Natureza desta fase.** Não existe `deck.md`: a v1 (Otharri) foi reprovada e nunca comprada. Isto é
> um `build` na prática — **eu consolido os 99** a partir dos titulares e reservas das Fases 2, 3, 4, 5
> e 7, dimensiono a base de mana e aplico as decisões de preço que o orquestrador me delegou.
>
> **Regra 5 — conferido.** `decisions.md` registra só movimentos de zona de comando (Otharri entra e
> sai, Tori sai, Phlage entra). **Nenhuma carta destes 99 tem histórico de corte** e nenhuma entrada é
> reposição de carta cortada.
>
> **Regra 6 — oracle.** Todo texto citado foi puxado com `bin/mtgdb oracle` nesta sessão — das 77
> cartas não-básicas dos 99, de todas as cartas que saem e dos ~30 terrenos candidatos. Os 99 conferidos por
> `mtgdb oracle -json`: **todas resolvem por nome exato, todas `legal` em Commander, todas com
> identidade ⊆ RW.**
>
> **Regra 7 — caixa e `lista.txt` antes de compra.** `bin/mtgdb collection -list` (116) e `lista.txt`
> (69 não-básicas + 30 básicos) varridas primeiro. **50 das 99 cartas já existem fisicamente** (28
> não-básicas + 22 básicos). As 49 compras foram conferidas uma a uma com `mtgdb collection`: **nenhuma
> está na caixa nem na `lista.txt`.** Dispensas de terreno da caixa/lista estão justificadas na §2.4.
>
> **Regra 2 — preço.** Só `mtgdb prices` (LigaMagic, menor). Nenhum número do Scryfall em lugar
> nenhum. O que não tem cotação está como **`a cotar`**, sem estimativa.
>
> **Travas do escape e regra 11.** Nenhuma reanimação, nenhum blink, nenhum MLD/stax/lock nos 99.

---

## 1. Cálculo

**CMC médio: 2,52** (soma 156 em 62 não-terrenos; X conta 0) · **draws+ramps mv ≤ 2: 16** ·
Fórmula: 31,42 + 3,13 × 2,52 − 0,28 × 16 = **34,8 → 35 terrenos**

- Os 16: 8 ramps (`Sol Ring`, `Arcane Signet`, `Boros Signet`, `Talisman of Conviction`, `Myr Convert`,
  `Zookeeper Mechan`, `Ornithopter of Paradise`, `Millikin`) + 8 fontes de carta mv ≤ 2 (`Inti`,
  `Tome of Legends`, `Reckoner Bankbuster`, `Thrill of Possibility`, `Cathartic Reunion`,
  `Faithless Looting`, `Cathartic Pyre`, `Demand Answers`). Se contar só vantagem de carta real (sem os
  5 loots), N = 11 e a fórmula dá **36,2**.
- **Listas de referência (Archidekt, bracket 3, comandante Phlage confirmado):**
  #14466984 *Forth Phlage!* **36** · #8068510 *Usain Bolt* **35** · #8145344 *Phlage – Well Done* **31**.
  (#25832205 *Budget Phlag* soma 28 — descartada da calibragem: parece lista incompleta.)
  O bracket não está no briefing; usei **3** por ser o que corresponde a "competitivo sob teto de preço".
  A #8068510 é um deck de blink/reanimação (Sun Titan, Cloudshift) — eixo que este deck **proíbe** — e
  nenhuma das três joga as 9 rochas deste; a contagem baixa delas não é molde.

### Recomendação: **37 terrenos** (+2 sobre a fórmula, −1 sobre a base de 38)

A leitura que fecha o número: **37 = 35 terrenos que produzem cor + 2 utilitários incolores que
funcionam como mágica** (`Blast Zone` = wipe, `Throne of the High City` = 3ª entrada da coroa). A fórmula
dá 35 fontes de mana; os dois utilitários não são fonte de `{R}`/`{W}` e por isso vêm **por cima**, como a
Fase 3 pediu (03 §9: "entram como utilidade sobre os 37, não no lugar de um terreno que produz `R`/`W`" —
na prática, sobre os 35).

Por que não ir a 38, como o método sugere de partida:
1. **A fórmula não enxerga o requisito de cor**, mas ele não se resolve com terreno a mais — resolve-se
   com fonte dupla. A Fase 4 mediu (04 §8.1): 36→38 terrenos vale +2,2 pp no T5; 6→14 fontes duplas
   vale **+6,1 pp** a custo zero de slot. Gastei o esforço onde ele rende: **12 fontes duplas** (§2).
2. **O 38º terreno sairia de um slot que já está no piso** — criaturas (15, piso 14), remoção (20 no
   ponto), draw (12, piso da meta) ou combustível (14, config D). Não há slot "sobrando" para ele.
3. **Inundar aqui é mais barato que o normal:** terreno excedente é descartado pelos 8 loots
   (combustível), é munição do `Flame Jab` (retrace), e o `Glittering Massif` faz cycling.

Por que não descer a 35–36, como a fórmula e as listas sugerem:
1. **Quatro pips coloridos sem genérico no escape** (`{R}{R}{W}{W}`), mais `{1}{R}{R}{R}` (`Torbran`),
   `{X}{X}{X}{R}{R}` (`Crackle`, X=3 = 11 manas), `{6}{W}` duas vezes (`Approach`) e `{4}{R}{W}{W}`
   (`Gisela`). Três dos quatro fechos pedem 7+ manas.
2. A Fase 4 mediu o Phlage no T3 em 85,0% com o pacote de 9 rochas (36 terrenos / 6 duplas), 87,5% com
   10 duplas e 89,8% com 38 / 12 duplas — a base mais dupla é o que compra os últimos pontos da única
   jogada garantida do deck. Com 37 / 12 o número fica entre os dois últimos; **não re-medi**.

**Correção a carregar para o `report.md` (Fase 4 §9.5):** o Phlage da zona de comando no T3 é
**~85%**, não 100% como dizem 02 §4.2 e 07 §7. O gargalo é land drop + os pips `{R}` e `{W}` juntos.

---

## 2. Terrenos recomendados

### 2.1 Não-básicos (15) — 12 fontes duplas, 4 viradas

| Terreno | Produz | Entra virado? | Sinergia/Utilidade | Na coleção? | Preço |
|---|---|---|---|---|---|
| Command Tower | `{R}`/`{W}` | não | intocável | **lista** | R$ 1,50 |
| Battlefield Forge | `{C}`; `{R}`/`{W}` com 1 de dano | não | dual desvirada; o Helix repõe a vida | compra | R$ 3,79 |
| Clifftop Retreat | `{R}`/`{W}` | só sem Mountain/Plains | 22 básicos + 2 Mountain Plains → quase sempre desvirada | compra | R$ 2,90 |
| Rugged Prairie | `{C}`; `{R/W}`→`{R}{R}`/`{R}{W}`/`{W}{W}` | não | **filtro**: Mountain + Prairie = `{W}{W}`. É a peça que converte uma mão vermelha no escape | compra | R$ 2,73 |
| Furycalm Snarl | `{R}`/`{W}` | só sem básico na mão | com 22 básicos + Massif/Summit (tipos Mountain Plains), falha só em mão que já é mulligan | compra | R$ 1,10 |
| Sunbillow Verge | `{W}`; `{R}` com Mountain/Plains | não | desvirada sempre | compra | `a cotar` |
| Radiant Summit | `{R}`/`{W}` | só com < 2 básicos | T1–T2 às vezes virada, depois nunca | compra | `a cotar` |
| Spectator Seating | `{R}`/`{W}` | não (mesa de 4) | desvirada sempre em multiplayer | compra | `a cotar` |
| Glittering Massif | `{R}`/`{W}` | **sim** | cycling `{2}` = combustível + tipo Land no cemitério (delirium da `Unholy Heat`) + seguro de inundação | compra | `a cotar` |
| Fields of Strife | `{R}`/`{W}` | **sim** | surveil 1 por `{2}{R}{W}` = mana sink que abastece | **lista** | `a cotar` |
| Stone Quarry | `{R}`/`{W}` | **sim** | — | **lista** | `a cotar` |
| Wind-Scarred Crag | `{R}`/`{W}` | **sim** | 1 de vida | **lista** | `a cotar` |
| Arena of Glory | `{R}`; exert: `{R}{R}` + ímpeto | só sem Mountain | **o Phlage escapado com `{R}{R}` da Arena ganha ímpeto e ataca no mesmo turno = 2º Helix (3 de dano + 3 de vida) + 6 de combate.** Também dá ímpeto a `Torbran`/`Gisela`. Cobre a função de ímpeto do `Bitter Reunion`, que sai (§4) | compra | `a cotar` |
| Blast Zone | `{C}` | não | wipe por MV no slot de terreno (Fase 5 §3 recomendou "forte") | **caixa** | R$ 0,89 |
| Throne of the High City | `{C}` | não | 3ª entrada da coroa (`Palace Jailer`, `Court of Ire`); conta nas 12 de draw da Fase 3 | compra | R$ 0,65 |

**Contagem contra os pedidos das fases (04 §8):**

| Pedido | Alvo | Entregue |
|---|---|---|
| Fontes duplas | mín. 10, alvo 12–14 | **12** (11 no critério estrito — `Rugged Prairie` é filtro e precisa de outra fonte colorida para entregar duas cores) |
| Duplas entrando viradas | ≤ 4–5 | **4 sempre** (Massif, Fields, Quarry, Crag) + 3 condicionais que quase nunca falham (Snarl, Clifftop, Summit) |
| Terrenos só `{C}` | — | **2** (Blast Zone, Throne). Recusei o 3º (`War Room`, §4) |
| Terreno com cycling | considerar (orquestrador, Fase 5) | **1** (`Glittering Massif`, que é dual) |

### 2.2 Básicos: **11 Mountain + 11 Plains** (22, todos da `lista.txt`, que tem 15 de cada)

Pips das 62 não-terreno: **`{R}` = 41 · `{W}` = 22** (66% / 34%). A proporção pura daria **14/8** —
e 14/8 está errado para este deck, pela mesma razão que a v1 já registrou em sentido inverso: o vermelho
tem muitos pips, **quase todos simples e baratos** (`Flame Slash`, `Magma Spray`, `Unholy Heat`,
`Faithless Looting`, os loots de `{1}{R}`), e o branco tem poucos pips **concentrados em duplos**
(escape `{W}{W}`, `Fumigate` `{3}{W}{W}`, `Slaughter the Strong` `{1}{W}{W}`, `Palace Jailer` `{2}{W}{W}`,
`Gisela` `{W}{W}`).

Medi as divisões (Monte Carlo próprio, 40–60 mil mãos, `on the play`, 37 terrenos com os 15 não-básicos
acima, rochas coloridas e incolores separadas). **O modelo ignora terreno virado e mulligan — serve para
comparar divisões entre si, não para substituir os números absolutos da Fase 4.**

| Básicos | `{R}{R}{W}{W}` T5 | `{R}{R}{W}{W}` T6 | `{R}{R}{R}` T4 | `{W}{W}` T3 | `{W}{W}` T4 | `{R}` até T2 | `{W}` até T2 |
|---|---|---|---|---|---|---|---|
| 14 M / 8 P (pips puros) | 68,4% | 76,7% | 69,2% | 62,4% | 74,4% | 92,7% | 78,0% |
| 12 M / 10 P | 70,2% | 78,2% | 64,0% | 67,3% | 78,8% | 90,7% | 81,9% |
| **11 M / 11 P** | **70,9%** | **78,6%** | 61,9% | **69,9%** | **80,8%** | **90,2%** | 83,6% |
| 10 M / 12 P | 70,3% | 78,4% | 58,6% | 71,2% | 81,9% | 88,6% | 84,7% |
| 8 M / 14 P | 70,0% | 77,7% | 52,8% | 76,5% | 85,8% | 85,8% | 88,3% |

Leitura: **o escape é praticamente insensível à divisão** entre 12/10 e 8/14 (a fixação vem das 12 duplas
e das 6 rochas de cor à escolha). O que a divisão troca é `{R}{R}{R}` do `Torbran` contra `{W}{W}` cedo e
`{R}` no T1–T2. **11/11** fica no topo do escape, mantém `{R}` cedo em ~90% (o vermelho joga a maioria
das mágicas de 1–2 manas) e dá `{W}{W}` no T4 em ~81%. `Torbran` é 1 carta em 99 e é jogada de T5+ na
prática; os `{W}{W}` são 5 cartas.

**Fontes por cor:** `{W}` = 11 Plains + 12 duplas = **23 em terreno** · `{R}` = 11 Mountain + 12 duplas +
Arena = **24 em terreno**. + 6 rochas de cor à escolha (`Arcane Signet`, `Boros Signet`, `Talisman`,
`Myr Convert`, `Commander's Sphere`, `Ornithopter`) + `Zookeeper Mechan` (`{R}`) + Treasures
(`Big Score`, `Seize the Spoils`, `Reckoner Bankbuster`).

### 2.3 Radar do EDHREC — seções `Lands` e `Utility Lands`

A seção `Radar do EDHREC` **não existe** no `02-theme.md` da v2; chamei `get_edhrec_recommendations`
uma vez (inclusão veio quebrada, como o guia avisa — usei só a nota de sinergia). Passaram pelo filtro de
2+ sinergias e entraram: `Arena of Glory` (0,28 — maior nota de Utility Lands), `Rugged Prairie`,
`Clifftop Retreat`, `Battlefield Forge`, `Spectator Seating`, `Radiant Summit`, `Sunbillow Verge`. Ficaram
fora, com motivo na §4.3: `Eiganjo`, `Sokenzan`, `War Room`, `Forgotten Cave`/`Secluded Steppe`/Deserts,
`Elegant Parlor`, `Sacred Foundry`, `Sunbaked Canyon`, fetches (`Arid Mesa`, `Fabled Passage`),
`Emeria`, `Rogue's Passage`, `Reliquary Tower`.

### 2.4 Caixa e `lista.txt` — terrenos dispensados (regra 7)

| Terreno | Origem | Por que não entra |
|---|---|---|
| `Boros Guildgate` | lista | Dual virada idêntica a `Stone Quarry`. Perde a vaga para `Glittering Massif` (também virada, **mais** cycling = combustível + delirium). **É o fallback do Massif** se a cotação dele falhar (§6) |
| `Sandstone Bridge` | lista | Mono-`{W}` **e** virada; o ETB (+1/+1 e vigilância até o fim do turno) não tem alvo útil num deck que não ataca em massa. Pior que um Plains em todos os eixos |
| 4 Mountain + 4 Plains excedentes | lista | 22 básicos bastam; ficam fisicamente sobrando (decisão de coleção é do orquestrador) |
| `Invigorating Hot Spring` (caixa) | caixa | Não é terreno — é encantamento `{1}{R}{G}`, fora da identidade |

---

## 3. Os 99 — lista final consolidada

**Exatamente 99 + comandante = 100.** 62 não-terrenos + 37 terrenos. Tudo RW, tudo `legal`.
Origem: **caixa** = sobressalente (`data/collection.tsv`) · **lista** = Tori físico (`lista.txt`) ·
**compra**. Preço = LigaMagic (menor), `mtgdb prices`.

**Comandante (fora da régua):** `Phlage, Titan of Fire's Fury` · `{1}{R}{W}` · 6/6 · escape `{R}{R}{W}{W}`.

### Criaturas (15)

| Carta | CMC | Tipo | Cores | Categorias | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| Erebor Flamesmith | 2 | Creature — Dwarf Artificer 2/1 | R | wincon | 1 a cada oponente por instantâneo/feitiço; corpo 2/1 | compra | R$ 0,33 |
| Firebrand Archer | 2 | Creature — Human Archer 2/1 | R | wincon | 1 a cada oponente por não-criatura | compra | R$ 0,99 |
| Inti, Seneschal of the Sun | 2 | Legendary Creature — Human Knight 2/2 | R | draw, tema | todo descarte vira impulse (janela longa) — multiplica os loots; corpo 2/2 | compra | R$ 4,90 |
| Lesser Masticore | 2 | Artifact Creature — Masticore 2/2 | incolor | tema, remoção | descarte no custo = combustível; ping {4} repetível; persist; corpo 2/2 | caixa | R$ 0,09 |
| Millikin | 2 | Artifact Creature — Construct 0/1 | incolor | ramp, tema | motor: moe 1/turno em velocidade de instantâneo + mana + corpo | compra | R$ 0,18 |
| Myr Convert | 2 | Artifact Creature — Phyrexian Myr 2/1 | incolor | ramp, tema | rocha de qualquer cor + corpo 2/1 (piso de criaturas) | caixa | R$ 0,15 |
| Ornithopter of Paradise | 2 | Artifact Creature — Thopter 0/2 | incolor | ramp, tema | mv2 qualquer cor + 0/2 voador (céu) + corpo; entra no lugar do Boros Locket (Fase 4 R1) | compra | `a cotar` |
| Thermo-Alchemist | 2 | Creature — Human Shaman 0/3 | R | wincon | 0/3 defender; ping que destapa a cada mágica | compra | R$ 0,55 |
| Zookeeper Mechan | 2 | Artifact Creature — Robot 1/3 | R | ramp, tema | rocha {R} + corpo 1/3 que segura T2–T4; {6}{R}: +4/+0 no Phlage (Fase 4 R2) | caixa | R$ 0,09 |
| Guttersnipe | 3 | Creature — Goblin Shaman 2/2 | R | wincon | 2 a cada oponente por instantâneo/feitiço | compra | R$ 5,40 |
| Sin Prodder | 3 | Creature — Devil 3/2 | R | draw, tema | carta ou moinho+dano todo upkeep; 3/2 menace; fonte vermelha | compra | R$ 1,00 |
| Palace Jailer | 4 | Creature — Human Soldier 2/2 | W | draw, remoção | monarca + exila criatura (0,5 remoção); corpo 2/2 | compra | R$ 0,30 |
| Torbran, Thane of Red Fell | 4 | Legendary Creature — Dwarf Noble 2/4 | R | wincon | +2 por fonte vermelha: Helix 5, pingadores ×3, Court 9 | compra | R$ 6,98 |
| Syr Carah, the Bold | 5 | Legendary Creature — Human Knight 3/3 | R | draw, wincon | mágica que fere jogador → impulse; {T}: 1 de dano; corpo 3/3 | lista | R$ 0,09 |
| Gisela, Blade of Goldnight | 7 | Legendary Creature — Angel 5/5 | RW | wincon | dobra dano a oponentes (inclui combate) e corta pela metade o recebido; 5/5 voadora | compra | R$ 15,00 |

### Artefatos (10)

| Carta | CMC | Tipo | Cores | Categorias | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| Sol Ring | 1 | Artifact | incolor | ramp | intocável; paga a 2ª ação do turno do escape (+4,8 pp) | lista | R$ 7,95 |
| Arcane Signet | 2 | Artifact | incolor | ramp | intocável; {R}/{W} no T2 → Phlage T3 | lista | R$ 4,00 |
| Boros Signet | 2 | Artifact | incolor | ramp | {1}→{R}{W}: metade do escape numa ativação; converte o {C} do Sol Ring | compra | R$ 1,50 |
| Glass Casket | 2 | Artifact | W | remoção | exílio permanente MV≤3 | caixa | R$ 0,09 |
| Perpetual Timepiece | 2 | Artifact | incolor | tema | motor: moe 2/turno só tapando; 2ª habilidade responde a ódio de cemitério | compra | R$ 0,65 |
| Ratchet Bomb | 2 | Artifact | incolor | wipe | seletivo por MV; mata fichas com 0 contadores | caixa | R$ 1,00 |
| Reckoner Bankbuster | 2 | Artifact — Vehicle 4/4 | incolor | draw | 3 cartas por {2} cada; Treasure no fim | compra | R$ 1,90 |
| Talisman of Conviction | 2 | Artifact | incolor | ramp | mv2 de cor à escolha; 1 de vida pago pelo Helix | compra | R$ 3,80 |
| Tome of Legends | 2 | Artifact — Book | incolor | draw | página a cada entrada/ataque do comandante — o escape recarrega | compra | R$ 1,20 |
| Commander's Sphere | 3 | Artifact | incolor | ramp | {R}/{W}; sacrifica p/ comprar → vai ao cemitério (combustível) | caixa | R$ 0,50 |

### Encantamentos (6)

| Carta | CMC | Tipo | Cores | Categorias | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| Lorehold Excavation | 2 | Enchantment | RW | tema, wincon | motor: moe 1/turno automático + 1 a cada oponente/turno (fonte vermelha → Torbran) | compra | R$ 0,08 |
| Seal of Cleansing | 2 | Enchantment | W | remoção | pré-pago; vai ao cemitério ao usar | lista | R$ 0,30 |
| Case of the Crimson Pulse | 3 | Enchantment — Case | R | draw, tema | ETB descarta 1/compra 2; resolvido: todo upkeep descarta a mão e compra 2 (motor) | compra | R$ 0,90 |
| Celebrate the Mountain-king | 4 | Enchantment | W | remoção, tema | 1 não-terreno por oponente + recruit (loot) | lista | R$ 0,23 |
| Outpost Siege | 4 | Enchantment | R | draw | impulse de upkeep todo turno (Khans) | compra | R$ 0,90 |
| Court of Ire | 5 | Enchantment | R | draw, remoção, wincon | monarca + 7 de dano/upkeep em qualquer alvo (9 c/ Torbran) | compra | R$ 4,95 |

### Instantâneos (16)

| Carta | CMC | Tipo | Cores | Categorias | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| Gods Willing | 1 | Instant | W | proteção | protege Torbran/Gisela por {W}; scry 1 | lista | `a cotar` |
| Magma Spray | 1 | Instant | R | remoção | 2 + exila (mata recursão alheia) | caixa | R$ 0,05 |
| Unholy Heat | 1 | Instant | R | remoção | delirium → 6; cresce com o cemitério | compra | R$ 0,50 |
| Abrade | 2 | Instant | R | remoção | 3 em criatura OU destrói artefato | compra | R$ 0,30 |
| Cathartic Pyre | 2 | Instant | R | tema, remoção | modal: 3 de dano em criatura/PW OU loot 2 — instantâneo | compra | R$ 0,84 |
| Demand Answers | 2 | Instant | R | tema | loot instantâneo; pode sacrificar Treasure/rocha no lugar do descarte | compra | R$ 2,80 |
| Disenchant | 2 | Instant | W | remoção | artefato/encantamento em instantâneo (ódio de cemitério) | lista | R$ 0,25 |
| Lightning Helix | 2 | Instant | RW | remoção, wincon | 3 dano + 3 vida; fonte vermelha | compra | R$ 0,45 |
| Lightning Strike | 2 | Instant | R | remoção, wincon | 3 em qualquer alvo; fonte vermelha | caixa | R$ 0,10 |
| Smite the Deathless | 2 | Instant | R | remoção | 3 + remove indestrutível + exila | caixa | R$ 0,12 |
| Thrill of Possibility | 2 | Instant | R | tema | loot instantâneo: 2 ao cemitério | caixa | R$ 0,50 |
| Chaos Warp | 3 | Instant | R | remoção | irrestrita, instantânea | compra | R$ 4,12 |
| Generous Gift | 3 | Instant | W | remoção | irrestrita, instantânea | compra | R$ 5,00 |
| Seize Opportunity | 3 | Instant | R | draw | impulse 2 instantâneo, janela até o fim do próximo turno | caixa | R$ 0,05 |
| Wear // Tear | 3 | Instant // Instant | RW | remoção | fuse: artefato E encantamento em instantâneo | compra | R$ 2,91 |
| Big Score | 4 | Instant | R | tema, ramp | loot + 2 Treasures (paga {W}{W} do escape), instantâneo | compra | R$ 7,65 |

### Feitiços (15)

| Carta | CMC | Tipo | Cores | Categorias | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| Conflagrate | 1 | Sorcery | R | tema, wipe, wincon | X dividido (wipe cirúrgico/fecho); flashback descarta X = combustível; fonte vermelha | compra | R$ 0,07 |
| Faithless Looting | 1 | Sorcery | R | tema | 3 ao cemitério por {R}; flashback | compra | R$ 5,40 |
| Flame Jab | 1 | Sorcery | R | tema, remoção | retrace: nunca sai do cemitério, converte terreno excedente em 1 de dano (3 c/ Torbran) | compra | R$ 0,05 |
| Flame Slash | 1 | Sorcery | R | remoção | 4 dano por {R}; fonte vermelha | caixa | R$ 2,48 |
| Requisition Raid | 1 | Sorcery | W | remoção | artefato E encantamento no mesmo card (ódio de cemitério) | caixa | R$ 0,44 |
| Cathartic Reunion | 2 | Sorcery | R | tema | 3 ao cemitério, saca 3 | compra | R$ 0,14 |
| Crackle with Power | 2 | Sorcery | R | wincon, remoção | 5X a cada um de até X alvos: X=3 → 15 em cada oponente | compra | R$ 26,40 |
| Mizzium Mortars | 2 | Sorcery | R | remoção, wipe | 4 numa criatura alheia; overload = wipe unilateral | compra | R$ 1,10 |
| Swift Reckoning | 2 | Sorcery | W | remoção | destrói criatura virada; spell mastery → instantâneo | lista | R$ 0,10 |
| Light Up the Stage | 3 | Sorcery | R | draw | impulse 2 por {R}: o Helix liga o spectacle; repõe Monument (fila 03 §8) | compra | R$ 0,74 |
| Reduce to Memory | 3 | Sorcery — Lesson | W | remoção | exila qualquer não-terreno | lista | R$ 0,13 |
| Seize the Spoils | 3 | Sorcery | R | tema, ramp | loot + 1 Treasure; mv3 | compra | R$ 0,06 |
| Slaughter the Strong | 3 | Sorcery | W | wipe | sacrifício por poder: poupa os pingadores 0–2 | compra | R$ 0,20 |
| Fumigate | 5 | Sorcery | W | wipe | vida por criatura; enche o cemitério | caixa | R$ 1,37 |
| Approach of the Second Sun | 7 | Sorcery | W | wincon | alt-win: conjurada 2× = vitória; 7 de vida por conjura | compra | R$ 12,02 |

### Terrenos (37)

| Carta | CMC | Tipo | Cores | Categorias | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| Arena of Glory | 0 | Land | R | terreno, tema | exert: {R}{R} + ímpeto na criatura → Phlage escapado ataca no mesmo turno (2º Helix); cobre o ímpeto do Bitter Reunion · origem: EDHREC Utility Lands | compra | `a cotar` |
| Battlefield Forge | 0 | Land | RW | terreno | dual desvirada; {C} grátis; 1 de dano pago pelo Helix · origem: EDHREC/Archidekt | compra | R$ 3,79 |
| Blast Zone | 0 | Land | — | terreno, wipe | wipe por MV no slot de terreno | caixa | R$ 0,89 |
| Clifftop Retreat | 0 | Land | RW | terreno | dual desvirada com Mountain/Plains (22 básicos + 2 Mountain Plains) · origem: EDHREC/Archidekt | compra | R$ 2,90 |
| Command Tower | 0 | Land | RW (via comandante) | terreno | intocável; {R}/{W} desvirado | lista | R$ 1,50 |
| Fields of Strife | 0 | Land | RW | terreno | dual virada + surveil 1 por {2}{R}{W} | lista | `a cotar` |
| Furycalm Snarl | 0 | Land | RW | terreno | dual; desvirada revelando básico | compra | R$ 1,10 |
| Glittering Massif | 0 | Land — Mountain Plains | RW | terreno, tema | dual virada + cycling {2}: combustível, tipo Land no cemitério p/ delirium da Unholy Heat, seguro contra inundação | compra | `a cotar` |
| 11× Mountain | 0 | Basic Land — Mountain | R | terreno | básico | lista | `a cotar` |
| 11× Plains | 0 | Basic Land — Plains | W | terreno | básico | lista | `a cotar` |
| Radiant Summit | 0 | Land — Mountain Plains | RW | terreno | desvirada com 2+ básicos (22 no deck) · origem: EDHREC/Archidekt | compra | `a cotar` |
| Rugged Prairie | 0 | Land | RW | terreno | filtro {R/W}→{R}{R}/{R}{W}/{W}{W}: Mountain + Prairie = {W}{W} do escape · origem: EDHREC/Archidekt | compra | R$ 2,73 |
| Spectator Seating | 0 | Land | RW | terreno | desvirada com 2+ oponentes (sempre, em mesa de 4) · origem: EDHREC/Archidekt | compra | `a cotar` |
| Stone Quarry | 0 | Land | RW | terreno | dual virada | lista | `a cotar` |
| Sunbillow Verge | 0 | Land | RW | terreno | sempre desvirada; {R} com Mountain/Plains · origem: EDHREC | compra | `a cotar` |
| Throne of the High City | 0 | Land | — | terreno, draw | 3ª entrada para a coroa (Jailer, Court); conta nas 12 de draw da Fase 3 | compra | R$ 0,65 |
| Wind-Scarred Crag | 0 | Land | RW | terreno | dual virada + 1 de vida | lista | `a cotar` |

---

## 4. Contagens finais contra as metas

| Meta | Alvo | Final | Cartas / observação |
|---|---|---|---|
| **Criaturas** | piso 14, meta 16–17 | **15** ✅ piso · ❌ meta | `Zookeeper Mechan`, `Myr Convert`, `Ornithopter of Paradise`, `Millikin`, `Lesser Masticore`, `Palace Jailer`, `Inti`, `Sin Prodder`, `Syr Carah`, `Torbran`, `Gisela`, `Guttersnipe`, `Thermo-Alchemist`, `Erebor Flamesmith`, `Firebrand Archer`. **A meta de 16–17 não coube** — §7, pendência 2 |
| **Remoção** | 20 no ponto (Fase 5) | **20** ✅ | as 18 da Fase 5 + `Court of Ire` + `Palace Jailer` (0,5). **Parciais não contadas:** `Cathartic Pyre` (modal), `Flame Jab`, `Lesser Masticore` (ping), `Crackle` (X=1), `Conflagrate` |
| **Wipes** | 3–4 | **3 + 1 em terreno** ✅ | `Slaughter the Strong`, `Ratchet Bomb`, `Fumigate` + `Blast Zone`. Efeitos de massa extras: `Mizzium Mortars` (overload, unilateral), `Conflagrate` |
| **Combustível** | 14 (config D) | **14** ✅ | Motores (6): `Millikin`, `Lorehold Excavation`, `Perpetual Timepiece`, `Flame Jab`, `Lesser Masticore`, `Case of the Crimson Pulse` · Loots (8): `Thrill`, `Cathartic Reunion`, `Faithless Looting`, `Cathartic Pyre`, `Demand Answers`, `Big Score`, `Seize the Spoils`, `Conflagrate`. **Custo declarado abaixo** |
| **Ramp** | 9, nunca < 8 | **9** ✅ | `Sol Ring`, `Arcane Signet`, `Boros Signet`, `Talisman of Conviction`, `Myr Convert`, `Zookeeper Mechan`, `Commander's Sphere`, `Ornithopter of Paradise`, `Millikin`. **8 de 9 em mv ≤ 2** (mín. 6) · **6 de cor à escolha** (mín. 6, no limite) · explosivo: `Big Score` + `Seize the Spoils` (2 de 3) |
| **Draw** (vantagem real, critério da Fase 3) | 12–13 nominais | **12** ✅ (piso) | `Palace Jailer`, `Case of the Crimson Pulse`, `Inti`, `Sin Prodder`, `Tome of Legends`, `Court of Ire`, `Outpost Siege`, `Seize Opportunity`, `Reckoner Bankbuster`, `Syr Carah`, `Light Up the Stage`, `Throne of the High City` |
| **Proteção** | 2 (Fase 5) | **1** ⚠ | `Gods Willing`. `Boros Charm` saiu por preço (§5) — função de indestrutível em massa **descoberta** |
| **Counters** | — | 0 | buraco de cor declarado pela Fase 5 §5.1 |
| **Caminhos de vitória** (≠ atacar com o comandante) | 3+ | **4** ✅ | §4.1 |
| **Terrenos** | 37–38 | **37** ✅ | 12 duplas, 4 viradas, 2 incolores |

**Custo declarado do combustível.** O 14 bate a config D **no número**, mas não na composição: a Fase 2
dimensionou 10 loots + 4 motores e eu entrego **8 loots + 6 motores** (`Case of the Crimson Pulse` e
`Lesser Masticore` ocupam o lugar de `Tormenting Voice`/`Wild Guess`/`Electric Revelation`/`Bitter
Reunion`). A tabela da §3.1 da Fase 2 foi medida com loots despejando 2–3 cartas cada — um motor despeja
menos no T5 e mais do T6 em diante. **Não re-simulei**; espero P(1º escape no T5) ligeiramente abaixo dos
35,5% da config D e o 2º escape mais cedo (02 §3.2: motor em campo = 2º escape em ~2,5 turnos contra 4–5).
O goldfishing (07 §6.2, métricas 3 e 4) mede os dois.

### 4.1 Caminhos de vitória restantes — **4**

| # | Caminho | Peças | Status |
|---|---|---|---|
| R1 | Relógio focado | `Court of Ire` (7/upkeep, 9 c/ Torbran) + Helix de cada entrada/ataque + `Syr Carah` | intacto |
| R2 | Dreno agregado multiplicado | `Guttersnipe`, `Thermo-Alchemist`, `Erebor Flamesmith`, `Firebrand Archer`, `Lorehold Excavation` × `Torbran`/`Gisela` | intacto |
| R3′ | Mágica X na mesa inteira | **`Crackle with Power`** (X=3 → 15 em cada oponente; 30 com `Gisela`) + **`Conflagrate`** (1º da fila 07 §8, substituindo `Repercussion`) | `Repercussion` saiu por preço; a fila da Fase 7 foi seguida — `Conflagrate` já estava nos 99 pela Fase 2, então a substituição não custou slot |
| R4 | Alt-win | `Approach of the Second Sun` | intacto |

**Não parei o fluxo:** restam 4 caminhos (≥ 3). `Chain Reaction`, 2º da fila, **não serve como fecho**
sem `Repercussion` — o oracle dá dano **só a criaturas** — e saiu por slot (§5).

### 4.2 Curva final (62 não-terrenos; X conta 0)

**0–1: 9 · 2: 31 · 3: 12 · 4: 5 · 5: 3 · 6+: 2** — CMC médio 2,52.

- 0–1: `Sol Ring`, `Flame Jab`, `Faithless Looting`, `Conflagrate`, `Flame Slash`, `Magma Spray`, `Unholy Heat`, `Requisition Raid`, `Gods Willing`
- 4: `Big Score`, `Palace Jailer`, `Outpost Siege`, `Celebrate the Mountain-king`, `Torbran`
- 5: `Court of Ire`, `Syr Carah`, `Fumigate` · 6+: `Gisela` (7), `Approach of the Second Sun` (7)
- Sumidouros tardios de mana (fora da curva nominal): `Crackle` X=3 (11), `Conflagrate` X, `Mizzium Mortars` overload (6), `Fields of Strife` (4), `Zookeeper Mechan` `{6}{R}`, `Lesser Masticore` `{4}`, `Blast Zone`, `Throne` (5)

### 4.3 Criaturas na mão — a composição final

| Criaturas | P(mão de 7 **sem** criatura) | P(0 vistas até o T4, na jogada) | P(≥ 2 vistas até o T4) |
|---|---|---|---|
| 14 (piso) | 33,2% | 20,1% | 42,9% |
| **15 (final)** | **30,4%** | **17,7%** | **46,8%** |
| 16 (meta) | 27,9% | 15,6% | 50,6% |
| 17 | 25,5% | 13,7% | 54,3% |

Com **31 permanentes não-terreno** (15 criaturas + 10 artefatos + 6 encantamentos), **P(mão de 7 sem
nenhum permanente não-terreno) = 6,5%** — a métrica que a Fase 2 (02 §4.2) defendeu como a certa. Mas o
usuário reclamou de **criatura**, não de permanente, e 30,4% é **2,5×** os 12,1% da v1 corrigida.

---

## 5. Decisões de preço e cortes — com ficha (regra 4)

### 5.1 A tabela de custo × função que o orquestrador pediu

| Carta | Preço | % do teto | Veredito | Por quê, em uma linha |
|---|---|---|---|---|
| `Repercussion` | R$ 47,99 | 24,0% | **SAI** | fecho de 1 turno que exige parceiro (`Blasphemous Act`, +R$ 12,10 = 30% do teto); simétrico, e com 15 criaturas próprias o risco de auto-dano cresce |
| `Boros Charm` | R$ 33,17 | 16,6% | **SAI** | única proteção em massa, mas não cabe junto do `Crackle` e o deck escolhe o turno dos próprios wipes |
| `Crackle with Power` | R$ 26,40 | 13,2% | **ENTRA** | único fecho simultâneo que escala com o mana de T9+ (7–9 sobrando, 07 §3.4); a dor nº 1 do intake é "não fecha" |
| `War Room` | R$ 19,99 | 10,0% | **SAI** | seria o 3º terreno só `{C}`; a 12ª fonte de draw sai por R$ 0,74 (`Light Up the Stage`) |
| `Unexpected Windfall` | R$ 19,90 | 10,0% | **SAI** | texto de `Big Score` com `{R}{R}`; `Big Score` + `Seize the Spoils` cobrem |
| `Gisela, Blade of Goldnight` | R$ 15,00 | 7,5% | **ENTRA** | um dos **dois** multiplicadores viáveis de RW no teto (07 §4.3); dobra R2 e R3′, cobre o céu, corta pela metade o dano recebido |
| `Tablet of Discovery` | R$ 12,49 | 6,2% | **SAI** | motor de 1 carta (só ETB); `Perpetual Timepiece` (R$ 0,65) mói 2/turno |
| `Blasphemous Act` | R$ 12,10 | 6,1% | **SAI** | sem `Repercussion` é só um wipe, e o deck tem 3 + `Blast Zone` + `Mizzium` |
| `Approach of the Second Sun` | R$ 12,02 | 6,0% | **ENTRA** | o único caminho que ignora a aritmética de 3×40 (R4) |
| `Bitter Reunion` | R$ 9,90 | 5,0% | **SAI** | loot coberto por 8; ímpeto coberto por `Arena of Glory` |
| `Electrostatic Field` | R$ 9,90 | 5,0% | **SAI** | gatilho idêntico ao `Erebor Flamesmith` (R$ 0,33); o corpo que faltava veio de `Zookeeper Mechan` (caixa, R$ 0,09) |
| `Big Score` | R$ 7,65 | 3,8% | **ENTRA** | loot + 2 Treasures instantâneo com 1 pip; é o explosivo nº 1 da Fase 4 |

**Entram: R$ 61,07 · Saem: R$ 165,44.**

### 5.2 Fichas dos cortes — funções e quem cobre cada uma

Condições assumidas **iguais** para quem sai e quem entra (checklist §3): Phlage escapado em campo,
`Torbran` e/ou `Gisela` em campo quando a peça for de fecho, **15 criaturas próprias** no deck (não as 4
que a Fase 7 assumiu ao dimensionar `Repercussion`).

**C1 · `Repercussion`** `{1}{R}{R}` Enchantment — **SAI (preço)**
F1 *sempre que uma criatura sofre dano, causa esse dano ao controlador dela* — simétrico. F2 sem corpo.
F3 encantamento (nada no deck conta encantamentos). F4 nada. F5 converte cada queima em criatura em dano
na cara; é fonte vermelha → `Torbran` soma +2 por instância. F6 T3, fecha a partir do T7–T9 com um wipe.
F7 **simétrico**: com as 15 criaturas deste deck (Fase 7 dimensionou com 4), um `Blasphemous Act` com 3
criaturas minhas em campo me dá 39 (18 com `Gisela`); exige um 2º card caro para fechar.
→ *matar a mesa num turno* → **parcialmente** coberto por `Crackle` (11 manas, não 1–2) e `Conflagrate`;
*converter remoção em dano na cara* → `Torbran` (+2 por evento, só fonte vermelha) — parcial.
**Descoberto (custo aceito):** o fecho de mesa **barato** com a mesa cheia. Caminhos restantes: 4 (§4.1).

**C2 · `Boros Charm`** `{R}{W}` Instant — **SAI (preço)**
F1 modal: 4 a jogador/PW · **seus permanentes ganham indestrutível** · double strike. F2 sem corpo.
F3 instantâneo (spell mastery, pingadores, combustível). F4 nada. F5 protege motores, pingadores e
multiplicadores de wipe de destruição. F6 T2, instantâneo. F7 não protege de dano/`-X/-X`/exílio/
sacrifício nem de `Repercussion`.
→ *4 na cara* → `Lightning Strike`, `Lightning Helix`, `Court of Ire`; *gatilho de mágica/combustível* →
31 instantâneos/feitiços; *double strike no Phlage* → nada (marginal); *indestrutível em massa* → **nada**.
**Descoberto (custo aceito):** wipe **alheio** de destruição leva motores e pingadores. Mitigações reais:
o Phlage volta por escape e o wipe enche o cemitério; `Lesser Masticore` tem persist; os 3 wipes do deck
são meus, no meu turno. **Proteção dedicada fica em 1** (`Gods Willing`) contra 2 da Fase 5.

**C3 · `War Room`** Land — **SAI (preço + cor)**
F1 `{C}`; `{3}`, `{T}`, 2 de vida: compre 1. F2–F4 nada. F5 draw repetível em slot de terreno. F6 T1.
F7 seria o **3º terreno só `{C}`** numa base de `{R}{R}{W}{W}` + `{R}{R}{R}` + `{W}{W}`×5; o `{3}` compete
com o escape todo turno.
→ *fonte de draw* → `Light Up the Stage` (R$ 0,74, 1º da fila 03 §8); *slot de terreno* → 1 básico;
*mana sink tardio* → `Crackle`, `Fields of Strife`, `Zookeeper Mechan`, `Lesser Masticore`, `Blast Zone`.
**Descoberto:** draw fica em 12 nominais, não 13.

**C4 · `Unexpected Windfall`** `{2}{R}{R}` Instant — **SAI (preço)**
F1 descarte + compre 2 + 2 Treasures. F2–F5 = `Big Score`. F6 T4. F7 `{R}{R}` na base que briga por `{W}{W}`.
→ tudo coberto por `Big Score` (texto idêntico, 1 pip) e `Seize the Spoils` (1 Treasure).
**Descoberto:** o 3º loot-com-tesouro da Fase 4 (04 §6: 2 → 3 compra +1,9 pp de "escapar com 2 de sobra"
no T6 e **custa** 2,4 pp de escape no T5). Custo aceito.

**C5 · `Tablet of Discovery`** `{2}{R}` Artifact — **SAI (preço)**
F1 ETB moe 1 e deixa jogar; `{T}`: `{R}`; `{T}`: `{R}{R}` só para instantâneo/feitiço. F3 artefato
(combustível do `Demand Answers`). F5 rampa `{R}`. F6 T3. F7 **só vermelho** (o gargalo é branco) e o
combustível é **uma carta, uma vez** — não é motor.
→ *motor* → `Perpetual Timepiece` (2/turno), `Millikin`, `Lorehold`, `Flame Jab`, `Masticore`, `Case`;
*rampa* → não estava nas 9 da Fase 4; *`{R}{R}` para mágica* → duplas + `Boros Signet`. **Nada descoberto.**

**C6 · `Blasphemous Act`** `{8}{R}` Sorcery — **SAI (preço)**
F1 −1 por criatura em campo; 13 a cada criatura. F3 feitiço. F5 fonte vermelha. F6 1–3 manas com mesa
cheia. F7 leva **todas** as minhas 15 criaturas, `Gisela` e `Torbran` inclusos.
→ *wipe total* → `Fumigate`, `Slaughter the Strong`, `Ratchet Bomb`, `Blast Zone`; *wipe de fonte
vermelha* → `Mizzium Mortars` (overload, e unilateral); *parceira da `Repercussion`* → sem objeto.
**Descoberto:** wipe barato que ignora tamanho em mesa cheia. Custo aceito.

**C7 · `Bitter Reunion`** `{1}{R}` Enchantment — **SAI (preço)**
F1 ETB descarte → compre 2; `{1}`, sacrifique: suas criaturas ganham ímpeto. F3 encantamento que vai ao
cemitério. F5 **ímpeto ao Phlage escapado = 2º Helix no turno**. F6 T2.
→ *loot/combustível* → 8 loots + 6 motores; *ímpeto* → `Arena of Glory` (`a cotar`).
**Condição:** se `Arena of Glory` cair para o fallback (§6), o ímpeto fica **descoberto**.

**C8 · `Electrostatic Field`** `{1}{R}` 0/4 defender — **SAI (preço)** · corte condicionado da Fase 7 §10
F1 1 a cada oponente por instantâneo/feitiço. F2 **0/4 defender** — melhor bloqueador barato do pool.
F3 **criatura** (conta no piso). F4 nada útil. F5 alcance. F6 T2. F7 não ataca.
→ *gatilho* → `Erebor Flamesmith` (idêntico, R$ 0,33), `Guttersnipe`, `Thermo-Alchemist`, `Firebrand
Archer`; *corpo que segura T2–T5* → `Zookeeper Mechan` 1/3 (caixa), `Thermo-Alchemist` 0/3, `Torbran`
2/4, `Ornithopter` 0/2 voador. **Simetria:** dei ao `Mechan` crédito de corpo e o nego ao `Field` — a
diferença é R$ 9,81 **e** o slot: o `Mechan` entra no slot de uma rocha (§5.3), o `Field` precisaria tirar
uma mágica. **Descoberto:** a 16ª criatura (27,9% contra 30,4%) — pendência 2.

**C9 · `Boros Locket`** `{3}` Artifact — **SAI (slot → `Ornithopter of Paradise`)** · 1ª cedência da Fase 4 §7
F1 `{R}`/`{W}`; `{R/W}`×4, `{T}`, sacrifique: compre 2. F3 artefato. F5 fixação. F6 **T3** (−3,9 pp no T5
contra uma mv 2, 04 §5). F7 saque custa 4 pips.
→ *fixação* → `Ornithopter` (qualquer cor, **mv 2**, +2,7 pp no T4 e +2,0 pp no T5 pela Fase 4);
*saque tardio + ir ao cemitério* → `Commander's Sphere` (1 carta). **Descoberto:** 1 carta tardia.
**Ganho:** +1 criatura (0/2 voadora, cobre o céu).

**C10 · `Sphere of the Suns`** `{2}` Artifact — **SAI (slot → `Zookeeper Mechan`)**
F1 entra virada com 3 contadores; `{T}`, tire um: qualquer cor. F3 artefato (fodder do `Demand Answers`
quando esgota). F5 fixação por 3 turnos. F6 T2 (usável no T3). F7 **expira** e entra virada.
→ *rampa* → `Zookeeper Mechan` (`{R}`, permanente, desvirado); *fixação de cor à escolha* → 6 rochas de
cor à escolha + 12 duplas; *artefato para o `Demand Answers`* → `Mechan` também é artefato, + Treasures.
**Custo medido (§2.2):** −0,5 pp no escape T5/T6 e −1,1 pp em `{W}{W}` no T4 — o modelo trata a Sphere
como permanente e desvirada, então o custo real é **menor**. **Ganho:** +1 criatura (1/3 que bloqueia
T2–T4) a R$ 0 e sem slot. A Fase 4 deixou o `Mechan` como R2 "se a Fase 2 pedir corpo antes de cor" — a
Fase 2 pediu (meta 16–17), e a cor é da minha fase.

**C11 · `Electric Revelation`** `{2}{R}` Instant — **SAI (slot → `Lesser Masticore`)**
F1 descarte → compre 2; flashback `{3}{R}`. F3 instantâneo. F5 dois loots. F6 **T3**. F7 o flashback
exila (02 §3.4 limita o deck a 3 cartas de flashback — ficam 2: `Faithless Looting`, `Conflagrate`).
→ *loot instantâneo* → `Thrill`, `Demand Answers`, `Cathartic Pyre`, `Big Score`; *combustível* →
`Lesser Masticore` (descarte no custo). **Descoberto:** o 2º loot do flashback (seleção). **Ganho:** +1
criatura 2/2 com persist e ping repetível, mv 2 no lugar de mv 3.

**C12 · `Chain Reaction`** `{2}{R}{R}` Sorcery — **SAI (slot)** · era o 4º wipe da Fase 5
F1 X = nº de criaturas em campo, em cada criatura (**não** em jogadores). F3 feitiço. F5 fonte vermelha
(`Torbran` +2). F6 T4. F7 mata meus x/3 e x/4 (`Mechan`, `Thermo`, `Torbran`) e custa `{R}{R}`.
→ *wipe* → `Fumigate`, `Slaughter the Strong`, `Ratchet Bomb` + `Blast Zone` (4 com o terreno);
*wipe vermelho* → `Mizzium Mortars` overload; *fecho com `Repercussion`* → sem objeto.
**Nada descoberto.** O slot era o 63º: cortar aqui mantém remoção em 20, draw em 12 e combustível em 14.

**C13 · `Tormenting Voice`** `{1}{R}` e **C14 · `Wild Guess`** `{R}{R}` — **SAEM** · cortes condicionados da Fase 3 §7
Ficha completa da Fase 3 (03 §7), oracle reconferido: texto idêntico (descarte → compre 2), feitiço, T2,
`Wild Guess` com `{R}{R}`. → *loot* → `Thrill`, `Faithless Looting`, `Cathartic Reunion`; *combustível* →
`Case of the Crimson Pulse` (descarta a mão todo upkeep); *slot* → `Seize the Spoils` (loot + Treasure,
Fase 4 §6). **Nada descoberto.**

### 5.3 Cortes condicionados devolvidos à Fase 6 que **não entram**

**`Hedron Crawler`** `{2}` 0/1 (caixa, R$ 0,10) — Fase 4 §9.3
F1 `{T}`: `{C}`. F2 0/1 — não bloqueia nada relevante. F3 artefato-criatura. F5 mana incolor. F6 T2.
F7 **zero pip** (04 §2.3). Entraria como corpo (a Fase 4 não o conta como ramp), e para isso precisaria
tirar uma mágica. → *corpo* → `Zookeeper Mechan` 1/3 e `Ornithopter` 0/2 fazem o mesmo **e** produzem cor;
*`{C}`* → `Sol Ring`. Fica como opção da pendência 2.

**`Solemn Simulacrum`** `{4}` 2/2 (`a cotar`) — Fase 4 §9.3
F1 ETB busca básico **virado**; morre → compre 1. F2 2/2. F3 artefato-criatura. F5 fixa a cor que falta.
F6 **CMC 4 = o custo do escape** (joga um, não joga o outro). → *corpo* → 15 criaturas; *fixação* → 12
duplas + `Rugged Prairie`; *carta ao morrer* → 12 fontes de draw; *preço* desconhecido. Não entra.

### 5.4 Promovidos das reservas (entradas — ficha nas fases de origem)

| Entra | De onde | No lugar de | Sinergias (2+) |
|---|---|---|---|
| `Ornithopter of Paradise` | Fase 4 R1 | `Boros Locket` | rocha mv 2 de qualquer cor · corpo 0/2 voador (céu, piso de criaturas) · artefato |
| `Zookeeper Mechan` | Fase 4 R2 (caixa) | `Sphere of the Suns` | rocha `{R}` · corpo 1/3 · `{6}{R}`: +4/+0 no Phlage (10/6) · artefato |
| `Light Up the Stage` | Fase 3 §5, 1º da fila | `Monument to Endurance` (fora por preço) | impulse 2 · o Helix do Phlage liga o spectacle sozinho (custa `{R}`) · feitiço = combustível |
| `Lesser Masticore` | Fase 2 §9.1 (caixa) | `Electric Revelation` | combustível no custo · ping repetível · persist · corpo |

`Palantír of Orthanc` (fora por preço) **não foi reposto**: a queda de 14 → 12 fontes nominais da Fase 3
está declarada (os dois outros saíram por preço — `Monument` — e por cor/preço — `War Room`).

---

## 6. Custo — LigaMagic (menor), cotações de 2026-08-12 a 2026-09-23

| Bloco | Cartas | Cotado | `a cotar` |
|---|---|---|---|
| **Compras** | 49 | **R$ 133,43** (43 cartas) | 6: `Ornithopter of Paradise`, `Sunbillow Verge`, `Radiant Summit`, `Spectator Seating`, `Glittering Massif`, `Arena of Glory` |
| **Caixa + `lista.txt`** (já existem, contam na régua) | 50 | **R$ 22,47** (24 cartas) | 26: `Gods Willing`, `Fields of Strife`, `Stone Quarry`, `Wind-Scarred Crag`, 11 Mountain, 11 Plains |
| **Total das 99** | **99** | **R$ 155,90** (67 cartas) | **32 cartas** |
| **Folga até R$ 200,00** | | **R$ 44,10** — reservada inteira para os 32 itens `a cotar` | |

Comandante fora da régua. As cotações mais velhas (43 dias, 12/08) são de peças baratas da caixa e de
`Boros Signet`/`Myr Convert`/`Flame Slash`/`Fumigate`/`Reckoner Bankbuster`; `mtgdb prices -volatile` diz
quais precisam de reconferência antes do torneio.

**Regra de fechamento (decisão por gatilho, não estimativa):** cada uma das 6 compras `a cotar` tem teto
de **R$ 5,00**; se passar, cai para o fallback nomeado:

| Compra `a cotar` | Fallback | Preço do fallback | Efeito na base |
|---|---|---|---|
| `Ornithopter of Paradise` | `Boros Locket` (caixa) | R$ 0,18 | criaturas 15 → **14** (piso) |
| `Sunbillow Verge` | `Exotic Orchard` | R$ 1,00 | dupla condicional à mesa |
| `Radiant Summit` | `Rustvale Bridge` | R$ 1,09 | 5ª dupla virada (limite da Fase 4) |
| `Spectator Seating` | Mountain (lista) | a cotar | duplas 12 → 11 |
| `Glittering Massif` | `Boros Guildgate` (lista) | a cotar | perde o cycling |
| `Arena of Glory` | Mountain (lista) | a cotar | ímpeto do Phlage fica descoberto (C7) |

No pior caso (as 6 no teto), sobram R$ 14,10 para os 26 itens de caixa/lista `a cotar` (22 básicos e 4
comuns). **Se não couber**, a ordem de alavancas é: (1) os fallbacks acima; (2) `Big Score` → `Tormenting
Voice` (R$ 0,20; −R$ 7,45; perde 2 Treasures); (3) `Crackle with Power` → fila da Fase 7 (`Comet Storm`/
`Jaya's Immolating Inferno`, `a cotar` → `Darksteel Reactor`, caixa) — −até R$ 26,40, e **restam 3
caminhos** (R1, R2, R4 + `Conflagrate`).

**Lista a cotar para o orquestrador (12 nomes):** `Ornithopter of Paradise` · `Sunbillow Verge` ·
`Radiant Summit` · `Spectator Seating` · `Glittering Massif` · `Arena of Glory` · `Gods Willing` ·
`Fields of Strife` · `Stone Quarry` · `Wind-Scarred Crag` · `Mountain` · `Plains`.
(`Twin Bolt` e `Ornithopter` estavam na sua lista; `Twin Bolt` **não entrou**, não precisa de cotação.)

---

## 7. O que ficou fora

### 7.1 Por preço

| Carta | Preço | Quem decidiu | Função e quem cobre |
|---|---|---|---|
| `Smothering Tithe` | R$ 233,91 | orquestrador | Treasures → `Big Score`, `Seize the Spoils`, `Bankbuster` |
| `Solphim, Mayhem Dominus` | R$ 118,90 | orquestrador / Fase 7 | multiplicador → `Torbran`, `Gisela` |
| `Palantír of Orthanc` | R$ 117,98 | orquestrador | draw + moinho → **não reposto** (draw 14 → 12) |
| `Monument to Endurance` | R$ 110,00 | orquestrador | draw por descarte → `Light Up the Stage` (1º da fila 03 §8) |
| `Fiery Emancipation` | R$ 54,54 | orquestrador / Fase 7 | multiplicador → `Torbran`, `Gisela` |
| `Repercussion` | R$ 47,99 | **eu** (C1) | fecho → `Crackle` + `Conflagrate` |
| `Chandra's Ignition` | R$ 37,90 | orquestrador | wipe/fecho → `Fumigate`, `Crackle` |
| `Boros Charm` | R$ 33,17 | **eu** (C2) | proteção em massa → **descoberta** |
| `Kessig Flamebreather` | R$ 24,79 | orquestrador | pingador → `Firebrand Archer` |
| `War Room` | R$ 19,99 | **eu** (C3) | draw → `Light Up the Stage` |
| `Unexpected Windfall` | R$ 19,90 | **eu** (C4) | → `Big Score`, `Seize the Spoils` |
| `Winds of Abandon` | R$ 19,41 | Fase 5 (reserva de upgrade) | wipe unilateral → `Mizzium Mortars` overload |
| `Felidar Sovereign` | R$ 15,99 | Fase 7 (reserva R5) | alt-win → `Approach` |
| `Swords to Plowshares` | R$ 13,99 | Fase 5 | exílio de criatura → `Magma Spray`, `Smite`, `Glass Casket`, `Reduce to Memory` |
| `Tablet of Discovery` | R$ 12,49 | **eu** (C5) | → `Perpetual Timepiece` |
| `Blasphemous Act` | R$ 12,10 | **eu** (C6) | → 3 wipes + `Blast Zone` + `Mizzium` |
| `Bitter Reunion` | R$ 9,90 | **eu** (C7) | → loots + `Arena of Glory` |
| `Electrostatic Field` | R$ 9,90 | **eu** (C8) | → `Erebor Flamesmith` + `Zookeeper Mechan` |

### 7.2 Por slot ou por ficha

| Carta | Motivo |
|---|---|
| `Boros Locket`, `Sphere of the Suns`, `Electric Revelation`, `Chain Reaction`, `Tormenting Voice`, `Wild Guess` | fichas C9–C14 (§5.2) |
| `Hedron Crawler`, `Solemn Simulacrum` | cortes condicionados que não entram (§5.3) |
| `Faithless Salvaging`, `Firebolt` (Fase 2) · `Oracle's Vault`, `Bargaining Table`, `Endless Atlas`, `Dawn of Hope`, `Valakut Awakening`, `Magus of the Wheel`, `Court of Embereth` (Fase 3) · `Mind Stone`, `Gold Myr`/`Iron Myr`, `Strike It Rich` (Fase 4) · `Twin Bolt`, `Invoke the Divine`, `Fateful End`, `Seal of Fire`, `Fateful Absence`, `Deflecting Palm`, `Lux Cannon` (Fase 5) · `Comet Storm`, `Jaya's Immolating Inferno`, `Darksteel Reactor` (fila da Fase 7) | reservas das fases, sem alteração — continuam na fila na ordem em que cada fase as pôs |
| `Sunforger` (R$ 6,90) | 02 §9.7 — ficou sem slot em todas as fases; o equip `{3}` e o corpo magro seguem valendo como F7 |

### 7.3 Terrenos considerados e não usados

| Terreno | Motivo |
|---|---|
| `Eiganjo, Seat of the Empire`, `Sunbaked Canyon` | **upgrades a cotar**: Eiganjo = Plains desvirado com channel (remoção + descarte = combustível); Canyon = dual desvirada que se sacrifica por carta. Entram no lugar de 1 Plains / `Stone Quarry` se cotarem baratos |
| `Elegant Parlor` | alternativa equivalente ao `Glittering Massif` (virada + surveil 1) |
| `Sacred Foundry`, `Inspiring Vantage`, `Sundown Pass`, `Sunscorched Divide` | duplas a cotar; com 12 duplas o ganho marginal não justifica mais itens sem preço. `Sundown Pass` é a primeira se uma das 3 duplas `a cotar` falhar e o usuário quiser manter 12 |
| `Forgotten Cave`, `Secluded Steppe`, `Desert of the Fervent`/`of the True`, `Smoldering Crater`, `Drifting Meadow` | cycling mono e **virado**: seriam a 5ª e 6ª virada (acima do limite de 4–5) sem ser fonte dupla. O cycling já está no Massif |
| `Karn's Bastion` (R$ 9,85), `Sequestered Stash`, `Rogue's Passage`, `Reliquary Tower`, `Command Beacon` | 3º terreno só `{C}`; nenhum tem 2 sinergias com o eixo (o Beacon não serve: o Phlage volta do cemitério, não da zona de comando) |
| `Sokenzan` (R$ 13,16), `Emeria`, fetches | Sokenzan faz fichas (go-wide); Emeria reanima (trava do escape); fetch não adiciona fonte num deck de 2 cores com 22 básicos |
| `Rustvale Bridge`, `Exotic Orchard`, `Boros Guildgate` | fallbacks (§6) |

---

## 8. Pendências que exigem decisão do usuário (não decidi por ele)

1. **A pergunta da paciência (07 §7).** O deck ameaça no **T6–T7** e mata a mesa na mediana do **T13**.
   Sem `Repercussion` não existe mais fecho de mesa barato: o `Crackle` precisa de **11 manas** (T9–T11) e
   o `Approach` de 7 duas vezes. Corrigir no `report.md`: o T3 é **~85%**, não 100%.
2. **Mão sem criatura (02 §4.2).** Final: **15 criaturas, 30,4%** de mãos de 7 sem nenhuma — 2,5× a v1
   corrigida. A meta de 16–17 não coube sem rebaixar outra meta. Opções, todas com custo declarado:
   (a) `Electrostatic Field` (R$ 9,90) no lugar de `Light Up the Stage` → 16 criaturas, draw 12 → 11;
   (b) `Hedron Crawler` (caixa, R$ 0) no lugar de `Seize Opportunity` → 16 criaturas, draw 12 → 11;
   (c) as duas → 17 criaturas (25,5%), draw 10; (d) manter 15.
3. **Fecho de 2 cartas (`Repercussion` + wipe) — power level.** A pergunta deixou de existir nesta lista
   porque a peça saiu por preço. Registro para ele decidir se quer revertê-la: custa **R$ 47,99 + R$ 12,10**
   (`Blasphemous Act`) = R$ 60,09 **e 2 slots**. Tirando `Crackle` (R$ 26,40, 1 slot) ainda faltam
   **R$ 33,69 e 1 slot** — que sairiam da folga de R$ 44,10 reservada aos 32 itens `a cotar` e de uma meta
   já no piso. E o auto-dano agora é medido com 15 criaturas próprias, não 4. Nota de power level que continua valendo:
   **`Gisela` + `Crackle` X=3 = 30 em cada oponente** num feitiço (07 §5).
4. **Proteção em 1 peça.** `Boros Charm` (R$ 33,17) só volta no lugar do `Crackle` (+R$ 6,77 de custo,
   mesmo slot) — troca fecho por proteção. Os caminhos de vitória cairiam para 3 (R1, R2, R4, com
   `Conflagrate` em R3′). Decisão de gosto: dor 1 (não fecha) × dor 4 (morre para wipe).
5. **Básicos contam na régua** — assumido que sim (briefing); os 22 estão `a cotar`.

---

## 9. Validação

- **99 + comandante = 100.** 62 não-terrenos + 37 terrenos (15 não-básicos + 22 básicos). Sem duplicatas
  além de básicos.
- **Identidade e legalidade:** `bin/mtgdb oracle -json` sobre os 77 nomes não-básicos + os 2 básicos —
  todos `how = exact`, `LegalCommander = legal`, `ColorIdentity ⊆ RW`. Não usei `validate_deck` (não
  confere identidade nem banimento, e o `mtgdb` já cobriu as duas coisas).
- **Travas:** nenhuma reanimação/blink; `Perpetual Timepiece` embaralha o cemitério de volta só se você
  quiser (02 §3.4). Regra 11: nenhum MLD/stax/lock.
- **Intocáveis:** `Sol Ring`, `Arcane Signet`, `Command Tower` — presentes.
- **Orçamento:** R$ 155,90 cotados + 32 `a cotar` ≤ R$ 200,00 sob a regra de gatilho da §6.
