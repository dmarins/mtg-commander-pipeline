# Relatório de Otimização — Krenko, Mob Boss

**Rodada v3 · 2026-08-22 · modo `improve`**
Estado: **26 trocas aplicadas.** Uma decisão pendente (ver "Pendências").

---

## Mudanças aplicadas

| # | Sai | Entra | Fase | Motivo em uma linha |
|---|---|---|---|---|
| 1 | Crash Through | **Impact Tremors** | lordes | trample em 1/1 = 0 de excesso → 10 dano a cada oponente por ativação |
| 2 | Warlord's Fury | **Goblin Bombardment** | lordes | 1º sac outlet grátis; acende Boggart Shenanigans, que era caminho morto |
| 3 | Goblin Rabblemaster | **Goblin Chieftain** | lordes | sai o ataque compulsório; entra lorde +1/+1 e **haste em Krenko** |
| 4 | Ambitious Assault | **Shared Animosity** | lordes | 100% das criaturas são Goblin: +9/+0 em cada, **todo turno** |
| 5 | Bravado | **Purphoros, God of the Forge** | lordes | 20 dano a cada oponente por ativação; indestructible |
| 6 | Tidings of War | **Rundvelt Hordemaster** | lordes | amass empilhava num só Army → lorde + motor de carta |
| 7 | Akki Ember-Keeper | **Skirk Prospector** | lordes | gatilho morto → 2º sac outlet, Goblin de 1 mana |
| 8 | Mob Mentality | **Hobgoblin Bandit Lord** | lordes | exigia ataque total → lorde + conversor repetível |
| 9 | Shiny Impetus | **Urza's Incubator** | ramp | dava +2/+2 a criatura adversária → `{2}` off em 33 cartas · **coleção** |
| 10 | Mycosynth Wellspring | **Umbral Mantle** | ramp | falso ramp que agravava flood → mana sink sem fundo |
| 11 | Goblin Glasswright | **Hammer of Purphoros** | ramp | Treasure adiado → haste a todos + terreno vira Golem 3/3 |
| 12 | Rummaging Goblin | **Skullclamp** | draw | loot net 0 → token 1/1 vira 2 cartas por 1 mana |
| 13 | Grotag Night-Runner | **Idol of Oblivion** | draw | exigia conectar → 1 carta/turno sem gastar mana |
| 14 | Goblin Sky Raider | **Reckless Impulse** | draw | 3 manas por 1 de dano evasivo → 2 cartas por 2 |
| 15 | Gundabad Opportunist | **Light Up the Stage** | draw | impulso one-shot por 4 → 2 cartas por `{R}` efetivo |
| 16 | Fissure Wizard | **Goblin Engineer** | draw | loot net 0 → Goblin que tutora e recompra Skullclamp |
| 17 | Innocent Bystander | **Goblin Matron** | draw | gatilho impossível de acionar → tutor de qualquer Goblin |
| 18 | Dragon Mantle | **Outpost Siege** | draw | cantrip net 0 → draw repetível que sobrevive a wipe |
| 19 | Misty Mountains Raider | ~~Deflecting Swat~~ → **Kuldotha Rebirth** | proteção | amass empilhava → 3 Goblins por `{R}` · **coleção** |
| 20 | Punishing Fire | **Goblin Chirurgeon** | proteção | Shock sem lifegain → regeneração repetível e grátis |
| 21 | Pinecone Strike | **Swiftfoot Boots** | proteção | Smite cobre melhor → hexproof + haste |
| 22 | Cosmotronic Wave | **Lightning Greaves** | proteção | 1 de dano não responde nada pós-T4 → equip {0} |
| 23 | Redcap Thief | **Chandra, Torch of Defiance** | proteção | taxa errada no CMC 3 → PW que atravessa wipe |
| 24 | Memorial to War | **Den of the Bugbear** | manabase | MLD sai por **regra 11** → terreno que vira Goblin e gera tokens |
| 25 | Looming Spires | **Castle Embereth** | manabase | 0 sinergia, entrava virado → desvirado com Mountain + anthem |
| 26 | Mountain (1 de 32) | **Volley Veteran** | manabase | 36→35 terrenos → remoção que escala com o enxame |

**Reversão registrada:** `Mudbutton Torchrunner` estava marcado para corte e **foi mantido** — o deck
ganhou três sac outlets nesta rodada e ele passou de "1/1 por 3 manas" a 3 de dano dirigível à vontade.

---

## Diagnóstico → resultado

| Categoria | Antes | Depois | Meta |
|---|---|---|---|
| Draw real | 5 | **12** | 12–13 ✅ |
| Economia de mana | 2 | **6** | 6–7 ✅ |
| Proteção | **0** | **3** | — ✅ |
| Permanentes que atravessam wipe | 2,5 | **12** | — ✅ |
| Lordes de Goblin | **0** | **3** | 4–5 ⚠️ |
| Haste permanente p/ Krenko | 2 (≈15%) | **5** | — ✅ |
| Wincons que fecham **sem combate** | **0** | **3** | 3+ ✅ |
| Resposta a encantamento | 0 | **0** | ❌ aberta |
| Terrenos | 36 (4/4 virados) | **35** (2/4) | ✅ |

---

## Informações Gerais do Deck

- **Comandante**: Krenko, Mob Boss
- **Cores**: Mono-vermelho (identidade R)
- **Número de Cartas**: 100 (99 + Comandante)
- **Foco**: Goblin tribal — enxame de tokens convertido em dano por **gatilhos de ETB**, não por combate
- Por tipo: Comandante 1 · Planeswalkers 1 · Criaturas 30 · Artefatos 9 · Encantamentos 8 ·
  Instantâneos 8 · Feitiços 8 · Terrenos 35 (31 Mountain + 4 não-básicos)
- Draw **12** · Economia de mana **6** · Interação **10** · Board wipes **1** (assimétrico) · Proteção **3**

## Curva de Mana

- Básicos **31** · Não-básicos **4** · **CMC médio 2,42**
- Distribuição: 0–1: **13** · 2: **21** · 3: **22** · 4: **9** · 5+: **0**
- 100% das cartas são mono-R ou incolores.

> **Ressalva honesta:** CMC 3 continua alto (22). A melhora é **efetiva**, não estrutural: com
> `Urza's Incubator` em jogo, 13 dessas cartas são feitiços de criatura Goblin e passam a custar 1.
> Sem o Incubator, o engarrafamento do turno 3 permanece.

## Sinergias e Estratégias

- **Estratégia**: montar largura de Goblins e converter **quantidade de gatilhos** em dano simultâneo
  a todos os oponentes. O deck deixou de precisar que o dano *atravesse* bloqueadores.
- **Sinergias principais**:
  - Krenko cria N tokens → `Impact Tremors` (1 cada) + `Purphoros` (2 cada) = **3N a cada oponente**
  - `Skullclamp` + token 1/1 → 2 cartas, e o token que morre dispara `Boggart Shenanigans` e `Rundvelt Hordemaster`
  - `Goblin Bombardment` transforma o enxame — e o board wipe adversário — em dano dirigível
  - `Shared Animosity`: todas as 30 criaturas são Goblin, então cada atacante recebe +1/+0 por atacante
- **Cartas-chave**: Krenko · Goblin Chieftain (haste nele) · Impact Tremors · Purphoros · Skullclamp
- **Pontos fortes**: três caminhos que fecham sem combate; draw que escala com o enxame; 12 permanentes
  que sobrevivem a wipe de criaturas
- **Pontos fracos**: **0 resposta a encantamento** · CMC 3 alto · 8 dos 10 encantamentos são payoff,
  então um `Farewell` apaga o plano · combo dependente de Krenko com haste

## Combos e Condições de Vitória

1. **Krenko + Impact Tremors + Purphoros** — T6–7, mata a **mesa inteira**, sem combate
2. **Goblin Bombardment + Boggart Shenanigans** — T7–8, 1 oponente, sem combate
3. **Umbral Mantle + Skirk Prospector + Krenko** — T5–7, mesa inteira, **combo aprovado pelo usuário**
4. **Shared Animosity + Ferocity of the Wilds** — T6–7, mesa inteira, via combate
5. **Assault on Osgiliath** — T6+, 1–2 oponentes, via combate

---

## Custo da rodada

**R$ 488,95** — 25 compras pelo **menor valor da LigaMagic**, cotações de **2026-08-22**,
registradas em `data/prices.tsv`. Duas entradas vieram da coleção a custo zero
(`Urza's Incubator`, `Kuldotha Rebirth`).

Mais caras: Goblin Chieftain R$ 44,90 · Goblin Engineer R$ 39,60 · Umbral Mantle R$ 36,90 ·
Hobgoblin Bandit Lord R$ 29,97 · Lightning Greaves R$ 28,35 · Skullclamp R$ 26,89 ·
Shared Animosity R$ 25,99 · Purphoros R$ 129,99.

`Deflecting Swat` foi cortada ao custar **R$ 414,00** — 46% do total da rodada.

---

## Pendências

1. **`Lightning Greaves` concede shroud** — impede o Equip de `Umbral Mantle` em Krenko e impede
   `Goblin Chirurgeon` de regenerá-lo. Contornável pela ordem de jogo (Mantle antes, Greaves depois).
   **Decisão do usuário pendente:** manter, cortar por `General Kreat, the Boltbringer`, ou cortar e
   devolver o slot.
2. **Resposta a encantamento continua em 0.** `Chaos Warp` e `Vandalblast` registradas para a próxima rodada.
3. **Lordes em 3, meta era 4–5.**
4. **Goldfishing não executado.** Protocolo pronto em `07-wincons.md`.

## Lista de Cartas por Tipo

### Comandante (1)

- **Krenko, Mob Boss** — CMC 4 3/3 · Legendary Creature — Goblin Warrior · {T}: Create X 1/1 red Goblin creature tokens, where X is the number of Goblins you control.

### Planeswalkers (1)

- **Chandra, Torch of Defiance** — CMC 4 · Legendary Planeswalker — Chandra · +1: Exile the top card of your library. You may cast that card. If you don't, Chandra deals 2 damage to each o

### Criaturas (30)

- **Fanatical Firebrand** — CMC 1 1/1 · Creature — Goblin Pirate · Haste (This creature can attack and {T} as soon as it comes under your control.)
- **Goblin Chirurgeon** — CMC 1 0/2 · Creature — Goblin Shaman · Sacrifice a Goblin: Regenerate target creature.
- **Mogg Sentry** — CMC 1 1/1 · Creature — Goblin Warrior · Whenever an opponent casts a spell, this creature gets +2/+2 until end of turn.
- **Mudbutton Clanger** — CMC 1 1/1 · Creature — Goblin Warrior · Kinship — At the beginning of your upkeep, you may look at the top card of your library. If it shares a creatu
- **Raging Goblin** — CMC 1 1/1 · Creature — Goblin Berserker · Haste (This creature can attack and {T} as soon as it comes under your control.)
- **Reckless Lackey** — CMC 1 1/2 · Creature — Goblin Pirate · First strike, haste
- **Skirk Prospector** — CMC 1 1/1 · Creature — Goblin · Sacrifice a Goblin: Add {R}.
- **Tin Street Dodger** — CMC 1 1/1 · Creature — Goblin Rogue · Haste
- **Conspicuous Snoop** — CMC 2 2/2 · Creature — Goblin Rogue · Play with the top card of your library revealed.
- **Ember Hauler** — CMC 2 2/2 · Creature — Goblin · {1}, Sacrifice this creature: It deals 2 damage to any target.
- **Goblin Cratermaker** — CMC 2 2/2 · Creature — Goblin Warrior · {1}, Sacrifice this creature: Choose one —
- **Goblin Engineer** — CMC 2 1/2 · Creature — Goblin Artificer · When this creature enters, you may search your library for an artifact card, put it into your graveyard, then 
- **Goblin Wardriver** — CMC 2 2/2 · Creature — Goblin Warrior · Battle cry (Whenever this creature attacks, each other attacking creature gets +1/+0 until end of turn.)
- **Goblin-town Flunkies** — CMC 2 1/1 · Creature — Goblin Soldier · Haste
- **Goro-Goro, Disciple of Ryusei** — CMC 2 2/2 · Legendary Creature — Goblin Samurai · {R}: Creatures you control gain haste until end of turn.
- **Rundvelt Hordemaster** — CMC 2 1/1 · Creature — Goblin Warrior · Other Goblins you control get +1/+1.
- **Clamor Shaman** — CMC 3 1/1 · Creature — Goblin Shaman · Riot (This creature enters with your choice of a +1/+1 counter or haste.)
- **Goblin Chainwhirler** — CMC 3 3/3 · Creature — Goblin Warrior · First strike
- **Goblin Chieftain** — CMC 3 2/2 · Creature — Goblin · Haste (This creature can attack and {T} as soon as it comes under your control.)
- **Goblin Matron** — CMC 3 1/1 · Creature — Goblin · When this creature enters, you may search your library for a Goblin card, reveal that card, put it into your h
- **Goblin Warchief** — CMC 3 2/2 · Creature — Goblin Warrior · Goblin spells you cast cost {1} less to cast.
- **Grishnákh, Brash Instigator** — CMC 3 1/1 · Legendary Creature — Goblin Soldier · When Grishnákh enters, amass Orcs 2. When you do, until end of turn, gain control of target nonlegendary creat
- **Guttersnipe** — CMC 3 2/2 · Creature — Goblin Shaman · Whenever you cast an instant or sorcery spell, this creature deals 2 damage to each opponent.
- **Hobgoblin Bandit Lord** — CMC 3 2/3 · Creature — Goblin Rogue · Other Goblins you control get +1/+1.
- **Krenko, Baron of Tin Street** — CMC 3 3/3 · Legendary Creature — Goblin · Haste
- **Mudbutton Torchrunner** — CMC 3 1/1 · Creature — Goblin Warrior · When this creature dies, it deals 3 damage to any target.
- **Squee, Goblin Nabob** — CMC 3 1/1 · Legendary Creature — Goblin · At the beginning of your upkeep, you may return this card from your graveyard to your hand.
- **Purphoros, God of the Forge** — CMC 4 6/5 · Legendary Enchantment Creature — God · Indestructible
- **Volley Veteran** — CMC 4 4/2 · Creature — Goblin Warrior · When this creature enters, it deals damage to target creature an opponent controls equal to the number of Gobl
- **Zada, Hedron Grinder** — CMC 4 3/3 · Legendary Creature — Goblin Ally · Whenever you cast an instant or sorcery spell that targets only Zada, copy that spell for each other creature 

### Artefatos (9)

- **Skullclamp** — CMC 1 · Artifact — Equipment · Equipped creature gets +1/-1.
- **Sol Ring** — CMC 1 · Artifact · {T}: Add {C}{C}.
- **Arcane Signet** — CMC 2 · Artifact · {T}: Add one mana of any color in your commander's color identity.
- **Idol of Oblivion** — CMC 2 · Artifact · {T}: Draw a card. Activate only if you created a token this turn.
- **Lightning Greaves** — CMC 2 · Artifact — Equipment · Equipped creature has haste and shroud. (It can't be the target of spells or abilities.)
- **Swiftfoot Boots** — CMC 2 · Artifact — Equipment · Equipped creature has hexproof and haste. (It can't be the target of spells or abilities your opponents contro
- **Hammer of Purphoros** — CMC 3 · Legendary Enchantment Artifact · Creatures you control have haste.
- **Umbral Mantle** — CMC 3 · Artifact — Equipment · Equipped creature has "{3}, {Q}: This creature gets +2/+2 until end of turn." ({Q} is the untap symbol.)
- **Urza's Incubator** — CMC 3 · Artifact · As this artifact enters, choose a creature type.

### Encantamentos (8)

- **Quest for the Goblin Lord** — CMC 1 · Enchantment · Whenever a Goblin you control enters, you may put a quest counter on this enchantment.
- **Goblin Bombardment** — CMC 2 · Enchantment · Sacrifice a creature: This enchantment deals 1 damage to any target.
- **Impact Tremors** — CMC 2 · Enchantment · Whenever a creature you control enters, this enchantment deals 1 damage to each opponent.
- **Boggart Shenanigans** — CMC 3 · Kindred Enchantment — Goblin · Whenever another Goblin you control is put into a graveyard from the battlefield, you may have this enchantmen
- **Ferocity of the Wilds** — CMC 3 · Enchantment · Attacking non-Human creatures you control get +1/+0 and have trample.
- **Shared Animosity** — CMC 3 · Enchantment · Whenever a creature you control attacks, it gets +1/+0 until end of turn for each other attacking creature tha
- **Burn, Burn, Tree and Fern** — CMC 4 · Enchantment — Saga · (As this Saga enters and after your draw step, add a lore counter. Sacrifice after IV.)
- **Outpost Siege** — CMC 4 · Enchantment · As this enchantment enters, choose Khans or Dragons.

### Instantâneos (7)

- **Blitz of the Thunder-Raptor** — CMC 2 · Instant · Blitz of the Thunder-Raptor deals damage to target creature or planeswalker equal to the number of instant and
- **Fists of Flame** — CMC 2 · Instant · Draw a card. Until end of turn, target creature gains trample and gets +1/+0 for each card you've drawn this t
- **Skullcrack** — CMC 2 · Instant · Players can't gain life this turn. Damage can't be prevented this turn. Skullcrack deals 3 damage to target pl
- **Smite the Deathless** — CMC 2 · Instant · Smite the Deathless deals 3 damage to target creature. That creature loses indestructible until end of turn. I
- **Goblin Surprise** — CMC 3 · Instant · Choose one —
- **Massive Raid** — CMC 3 · Instant · Massive Raid deals damage to any target equal to the number of creatures you control.
- **Lightning Volley** — CMC 4 · Instant · Until end of turn, creatures you control gain "{T}: This creature deals 1 damage to any target."

### Feitiços (9)

- **Kick in the Door** — CMC 1 · Sorcery · Put a +1/+1 counter on target creature. That creature gains haste until end of turn and can't be blocked by Wa
- **Kuldotha Rebirth** — CMC 1 · Sorcery · As an additional cost to cast this spell, sacrifice an artifact.
- **Krenko's Command** — CMC 2 · Sorcery · Create two 1/1 red Goblin creature tokens.
- **Reckless Impulse** — CMC 2 · Sorcery · Exile the top two cards of your library. Until the end of your next turn, you may play those cards.
- **Tormenting Voice** — CMC 2 · Sorcery · As an additional cost to cast this spell, discard a card.
- **Assault on Osgiliath** — CMC 3 · Sorcery · Amass Orcs X, then Goblins and Orcs you control gain double strike and haste until end of turn. (To amass Orcs
- **Hordeling Outburst** — CMC 3 · Sorcery · Create three 1/1 red Goblin creature tokens.
- **Light Up the Stage** — CMC 3 · Sorcery · Spectacle {R} (You may cast this spell for its spectacle cost rather than its mana cost if an opponent lost li
- **Goblin War Party** — CMC 4 · Sorcery · Choose one —

### Terrenos (35)

- **Castle Embereth** — — · Land · This land enters tapped unless you control a Mountain.
- **Den of the Bugbear** — — · Land · If you control two or more other lands, this land enters tapped.
- **Forgotten Cave** — — · Land · This land enters tapped.
- **31× Mountain** — — · Basic Land — Mountain · ({T}: Add {R}.)
- **The Autonomous Furnace** — — · Land — Sphere · This land enters tapped.

## Lista de Cartas para Exportação (padrão MTG Online)

```
1 Skirk Prospector
1 Shared Animosity
1 Arcane Signet
1 Assault on Osgiliath
1 Blitz of the Thunder-Raptor
1 Boggart Shenanigans
1 Purphoros, God of the Forge
1 Burn, Burn, Tree and Fern
1 Clamor Shaman
1 Conspicuous Snoop
1 Lightning Greaves
1 Impact Tremors
1 Outpost Siege
1 Ember Hauler
1 Fanatical Firebrand
1 Ferocity of the Wilds
1 Goblin Engineer
1 Fists of Flame
1 Forgotten Cave
1 Goblin Chainwhirler
1 Goblin Cratermaker
1 Hammer of Purphoros
1 Goblin Chieftain
1 Reckless Impulse
1 Goblin Surprise
1 Goblin War Party
1 Goblin Warchief
1 Goblin Wardriver
1 Goblin-town Flunkies
1 Goro-Goro, Disciple of Ryusei
1 Grishnákh, Brash Instigator
1 Idol of Oblivion
1 Light Up the Stage
1 Guttersnipe
1 Hordeling Outburst
1 Goblin Matron
1 Kick in the Door
1 Krenko's Command
1 Krenko, Baron of Tin Street
1 Lightning Volley
1 Castle Embereth
1 Massive Raid
1 Den of the Bugbear
1 Kuldotha Rebirth
1 Hobgoblin Bandit Lord
1 Mogg Sentry
1 Mudbutton Clanger
1 Mudbutton Torchrunner
1 Umbral Mantle
1 Swiftfoot Boots
1 Goblin Chirurgeon
1 Quest for the Goblin Lord
1 Raging Goblin
1 Reckless Lackey
1 Chandra, Torch of Defiance
1 Skullclamp
1 Urza's Incubator
1 Skullcrack
1 Smite the Deathless
1 Sol Ring
1 Squee, Goblin Nabob
1 The Autonomous Furnace
1 Rundvelt Hordemaster
1 Tin Street Dodger
1 Tormenting Voice
1 Goblin Bombardment
1 Zada, Hedron Grinder
1 Volley Veteran
31 Mountain

1 Krenko, Mob Boss
```
