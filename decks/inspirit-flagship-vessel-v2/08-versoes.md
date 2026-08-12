# Composição por versão — Torneio (≤ R$200) × Mesão (≤ R$350)

> Correção de 2026-08-11. O `report.md` da v2 apresentava **uma única lista de exportação híbrida** — ela tinha 6 slots na versão mesão (Padeem, Reckoner Bankbuster, Thousand Moons Smithy, Cloud Key, Restoration Magic, Cyberdrive Awakener) e 1 na versão torneio (Stoic Rebuttal), e não listava Reconnaissance Mission nem Bident of Thassa em lugar nenhum. Todas as estatísticas (tipos, curva, draw, proteção) foram calculadas sobre esse híbrido, que **não corresponde a nenhum dos dois decks físicos**. Este arquivo reconstrói as duas composições do zero, com tipos e custos verificados na Scryfall.

## Os 7 slots que mudam de versão

93 cartas são idênticas nas duas versões (incluindo o comandante e os 38 terrenos). Só estes 7 slots trocam:

| # | Torneio | US$ | Mesão | US$ | Origem da divergência |
|---|---|---|---|---|---|
| 1 | Reconnaissance Mission — Encantamento, {2}{U}{U} | 0,28 | Padeem, Consul of Innovation — Criatura Lendária, {3}{U} | 3,24 | v1, fase draw (slot da Frogmyr Enforcer) |
| 2 | Bident of Thassa — Encantamento **Artefato** Lendário, {2}{U}{U} | 0,39 | Reckoner Bankbuster — Artefato — Veículo, {2} | 0,61 | v1, fase draw (slot da Voyage Home) |
| 3 | Stoic Rebuttal — Instantâneo, {1}{U}{U} | 0,32 | Access Denied — Instantâneo, {3}{U}{U} | 2,62 | v2, fase interação (slot da Chainsaw) |
| 4 | Master Trinketeer — Criatura, {2}{W} | 0,29 | Thousand Moons Smithy — Artefato Lendário (transforma em terreno), 4 | ~2,81 | v2, corte de custo |
| 5 | Prophetic Prism — Artefato, {2} | 0,21 | Cloud Key — Artefato, {3} | ~1,00 | v2, corte de custo |
| 6 | One with the Machine — **Feitiço**, {3}{U} | 0,25 | Restoration Magic — Instantâneo, {W} | ~0,30 | v2, corte de custo |
| 7 | Chief of the Foundry — Criatura Artefato, {3} | 0,21 | Cyberdrive Awakener — Criatura Artefato, {5}{U} | ~1,50 | v2, corte de custo |

## Composição por tipo

| Categoria | Torneio | Mesão | Δ |
|---|---|---|---|
| Comandante (Spacecraft — criatura só a 8+) | 1 | 1 | — |
| **Criaturas de verdade** (entram como criatura) | **22** | **22** | — |
| **Veículos e Spacecraft** (só viram criatura com crew/station) | **3** | **4** | −1 |
| Artefatos não-criatura (fora veículos/spacecraft) | 15 | 15 | — |
| Encantamentos | 3 | 2 | +1 |
| Planeswalkers | 1 | 1 | — |
| Instantâneos | 10 | 11 | −1 |
| Feitiços | 7 | 6 | +1 |
| Terrenos | 38 | 38 | — |
| **Total** | **100** | **100** | — |

> Bident of Thassa é Encantamento **Artefato** — está contada em "Artefatos não-criatura" (não nos encantamentos), por ser esse o tipo que importa aqui (affinity, metalcraft, Master of Etherium e a estática do Inspirit).

### Corpos de verdade × permanentes que *viram* criatura

Os arquivos anteriores contavam **25 criaturas** porque agrupavam veículos e spacecraft junto das criaturas. Eles são **artefatos não-criatura em repouso** — precisam de outras criaturas para crew/station:

| Carta | Tipo real | Requisito | Presente em |
|---|---|---|---|
| Inspirit, Flagship Vessel | Artefato Lendário — Spacecraft, 5/5 | Station **8+** | ambas (comandante) |
| Warmaker Gunship | Artefato — Spacecraft, 4/3 | Station **6+** | ambas |
| Uthros Research Craft | Artefato — Spacecraft, 0/8 | Station **12+** | ambas |
| Brotherhood Vertibird | Artefato — Veículo, */4 | **Crew 2** | ambas |
| Reckoner Bankbuster | Artefato — Veículo, 4/4 | **Crew 3** | só mesão |

**Nenhum deles precisa virar criatura para dar valor**, o que atenua a contagem: a estática do Inspirit (hexproof + indestructible nos outros artefatos) e o gatilho 1+ funcionam **desde 0 contadores**; o 3+ do Uthros (compra por artefato conjurado) é o limiar que importa — o 12+ é praticamente inalcançável; e o ETB de dano do Gunship e o draw do Bankbuster independem de station/crew. Os limiares que valem a pena perseguir são o **6+ do Gunship** e o **8+ do Inspirit**.

**Economia de Station**: Station tapa **outra criatura sua** (velocidade de feitiço) e rende contadores iguais ao **poder** dela — tokens 1/1 rendem 1 por tap, o que torna os 8 do Inspirit lentos por essa via. Atalhos reais no deck:

- **Master of Etherium** — poder = nº de artefatos; sozinho costuma pagar os 8 do Inspirit.
- **Brotherhood Vertibird** — crewar por 2 de poder e então tapá-lo para Station: o poder dele também é o nº de artefatos. Crew e station são ambos ações de fase principal, então cabem no mesmo turno.
- **Deepglow Skate** — dobra os contadores já acumulados (4 → 8 fecha o Inspirit).
- **Coretapper + proliferate** (Kilo, Surge Conductor, Tezzeret's Gambit) — sobem contadores **sem tapar criatura nenhuma**.
- Criaturas com poder ≥ 3 disponíveis para station: **9** nas duas versões.

Station funciona com criaturas recém-jogadas — summoning sickness não impede, exatamente como em crew, porque é a habilidade do Spacecraft que tapa a criatura, não a criatura tapando a si mesma ([Draftsim](https://draftsim.com/mtg-station/), [TheGamer](https://www.thegamer.com/magic-the-gathering-mtg-spacecraft-stationing-guide/)).

**Consequência para o caminho A (enxame de combate)**: **Alibou** só dispara quando **artefatos-criatura atacam**. São **14 artefatos-criatura de verdade** nas duas versões; os Spacecraft entram nessa conta só depois de estacionados, e o Inspirit só ataca a partir de 8+.

**Correções de classificação em relação ao `deck.md`/`report.md` anteriores** (verificadas na Scryfall):

- **Stern Lesson** é **Instantâneo** ({2}{U}), não Feitiço — estava na seção errada. Por isso as contagens acima dão 10/7 e 11/6, e não 10/7 como no report antigo.
- **Cyberdrive Awakener** é **Artifact Creature — Construct**, não criatura simples: **é protegida** pela estática do Inspirit (o report tratava como exposta).
- **Kilo, Apogee Mind** é **Legendary Artifact Creature — Robot Artificer**: também protegida.
- **Thousand Moons Smithy** é MDFC (`Thousand Moons Smithy // Barracks of the Thousand`) — a face de trás é Terreno Artefato Lendário; segue contando como não-terreno no 99.

## Permanentes-artefato (o que a estática do Inspirit protege)

| | Torneio | Mesão |
|---|---|---|
| Artefatos no 99 (criaturas + não-criaturas) | 32 | 33 |
| Terrenos-artefato (Bridges) | 3 | 3 |
| Comandante (artefato) | 1 | 1 |
| **Total de permanentes-artefato** | **36** | **37** |

A diferença de 1 vem do slot 4: Thousand Moons Smithy (artefato) → Master Trinketeer (criatura não-artefato). Metalcraft, affinity e improvise seguem triviais nas duas versões.

**Permanentes não-artefato expostos** (não recebem hexproof/indestructible do comandante):

- **Comuns às duas**: Third Path Iconoclast, Malcator, Sai, Chrome Host Seedshark, Jhoira, Leonin Abunas, Deepglow Skate, Saheeli, Thopter Spy Network, Whirlwind of Thought.
- **Só torneio**: Master Trinketeer, **Reconnaissance Mission** (encantamento puro).
- **Só mesão**: Padeem.

Detalhe relevante: no torneio, o motor de draw por dano de combate é **Bident of Thassa**, que é artefato → **fica protegido**. No mesão o slot equivalente (Reckoner Bankbuster) também é artefato. Mas o slot 1 inverte: Reconnaissance Mission (torneio) é encantamento exposto, e Padeem (mesão) é criatura exposta — empate em fragilidade.

## Curva de mana

| Faixa de CMC | Torneio | Mesão |
|---|---|---|
| 0–1 | 8 | 9 |
| 2 | 14 | 14 |
| 3 | **19** | 17 |
| 4 | **13** | 12 |
| 5 | 4 | 5 |
| 6+ | 3 | 4 |
| **CMC médio** (sem terrenos, X = 0) | **3,05** | **3,08** |

Médias praticamente iguais, mas os perfis diferem: a torneio é mais **plana e concentrada em 3–4** (32 cartas), a mesão tem mais **top-end** (9 cartas em 5+, contra 7). O topo da mesão são justamente as peças que a torneio abriu mão: Cyberdrive Awakener (6) e Access Denied (5).

## Símbolos coloridos

| | Torneio | Mesão |
|---|---|---|
| Azul (U) | 32 pips (50%) | 29 pips (46%) |
| Branco (W) | 19 pips (30%) | 21 pips (33%) |
| Vermelho (R) | 11 pips (17%) | 11 pips (17%) |
| Híbrido U/R (Saheeli) | 2 | 2 |
| **Cartas com custo duplo-azul** | **6** | **4** |

Duplo-azul comum às duas: Disruption Protocol {U}{U}, Thopter Spy Network {2}{U}{U}, Reverse Engineer {3}{U}{U}. A torneio soma **Reconnaissance Mission {2}{U}{U}, Bident of Thassa {2}{U}{U} e Stoic Rebuttal {1}{U}{U}**; a mesão soma só Access Denied {3}{U}{U}.

⚠️ **Ponto de atenção da versão torneio**: a manabase (9 Island + ~11 não-básicos que produzem azul ≈ 20 fontes confiáveis) foi dimensionada para o perfil da lista híbrida. Com 6 cartas duplo-azul, a torneio fica **no limite inferior** da conta de Karsten para {U}{U} no turno 3 (~20–21 fontes). Mitigadores já presentes: Prophetic Prism (exclusiva da torneio, fixa qualquer cor), Talisman of Progress, Arcane Signet, Sphere of the Suns, Pentad Prism, Crystalline Crawler, Spire of Industry. Não é um erro, mas se o goldfishing travar em azul, o ajuste natural é **Mountain → Island** (o vermelho tem 11 pips e Boros Signet como muleta).

## Draw, ramp, interação e proteção

| Categoria | Torneio | Mesão | Meta do pipeline |
|---|---|---|---|
| Fontes de draw | **14** | 13 | 12–13 |
| Ramp padrão | 10 | 10 | 10–11 ✓ |
| Ramp explosivo | 1 | 1 | 2–3 (abaixo, aceito na v2) |
| Redutores de custo | **3** | 4 | — |
| Counterspells | 2 | 2 | — |
| Board wipes | 3 | 3 | 2–4 ✓ |
| Interação pontual | ~9 | ~9 | ~10 |
| Hexproof estático no comandante | **1** (Leonin Abunas) | **2** (Padeem + Leonin Abunas) | — |
| Indestructible sob demanda | 1 (Invisible Force Field) | 1 (Invisible Force Field) | — |
| One-shots de proteção | **2** (Loran's Escape, Blacksmith's Skill) | **3** (+ Restoration Magic) | — |

**Fontes de draw comuns (11)**: Sai, Uthros Research Craft, Jhoira, Thought Monitor, Midnight Clock, Thopter Spy Network, Whirlwind of Thought, Thirst for Knowledge, Stern Lesson, Tezzeret's Gambit, Reverse Engineer.
**Torneio (+3)**: Reconnaissance Mission, Bident of Thassa, One with the Machine → **14**. (Prophetic Prism repõe a si mesma, não conta como card advantage.)
**Mesão (+2)**: Padeem, Reckoner Bankbuster → **13**.

## Condições de vitória

| Caminho | Torneio | Mesão |
|---|---|---|
| **A. Enxame de combate** | ✅ e **reforçado** — dois anthems (Chief of the Foundry: +1/+1 em outras criaturas-artefato; Master Trinketeer: +1/+1 em Servos e Thopters) somados a Master of Etherium | ✅ — só Master of Etherium como lord |
| **B. Queima via marcadores** (Lux Artillery aos 30+) | ✅ idêntico | ✅ idêntico |
| **C. Conversão explosiva de rocks** (Cyberdrive Awakener) | ❌ **cortado** | ✅ |
| Token */* do Thousand Moons Smithy | ❌ | ✅ |
| Evasão em massa (voar concedido pelo Cyberdrive) | ❌ | ✅ |

Leitura prática: a **torneio é mais consistente e mais barata de operar** (curva menor, dois anthems, +1 fonte de draw), com **teto mais baixo** — perde o turno explosivo do Cyberdrive Awakener e a evasão coletiva. A **mesão tem mais teto e mais proteção** (hexproof duplo, 3 one-shots, wincon C), ao custo de curva mais pesada e 1 fonte de draw a menos.

## Diferença de custo (já apurada em `report.md`)

| Versão | Total | Teto | Situação |
|---|---|---|---|
| Torneio | US$36,85 (~R$202,68) | R$200 | no limite — falta ~1 corte de US$1 para folga real |
| Mesão | US$49,89 (~R$274,40) | R$350 | folgado |

## Listas de exportação

### Torneio (≤ R$200)

```
1 Hangarback Walker
1 Coretapper
1 Etherium Sculptor
1 Third Path Iconoclast
1 Enthusiastic Mechanaut
1 Brotherhood Vertibird
1 Foundry Inspector
1 Kilo, Apogee Mind
1 Malcator, Purity Overseer
1 Master of Etherium
1 Pinnacle Emissary
1 Sai, Master Thopterist
1 Surge Conductor
1 Chrome Host Seedshark
1 Uthros Research Craft
1 Warmaker Gunship
1 Jhoira, Weatherlight Captain
1 Leonin Abunas
1 Crystalline Crawler
1 Alibou, Ancient Witness
1 Deepglow Skate
1 Kappa Cannoneer
1 Thought Monitor
1 Master Trinketeer
1 Chief of the Foundry
1 Everflowing Chalice
1 Sol Ring
1 Arcane Signet
1 Sphere of the Suns
1 Talisman of Progress
1 Empowered Autogenerator
1 Glass Casket
1 Pentad Prism
1 Boros Signet
1 Perilous Snare
1 Midnight Clock
1 Lux Artillery
1 Lux Cannon
1 Bident of Thassa
1 Prophetic Prism
1 Thopter Spy Network
1 Whirlwind of Thought
1 Reconnaissance Mission
1 Saheeli, Sublime Artificer
1 Dispatch
1 Swords to Plowshares
1 Loran's Escape
1 Blacksmith's Skill
1 Unwanted Remake
1 Disruption Protocol
1 Invisible Force Field
1 Thirst for Knowledge
1 Stern Lesson
1 Stoic Rebuttal
1 Rip Apart
1 Tezzeret's Gambit
1 Chain Reaction
1 Fumigate
1 Reverse Engineer
1 Organic Extinction
1 One with the Machine
1 Command Tower
1 Spire of Industry
1 Exotic Orchard
1 Battlefield Forge
1 Rugged Prairie
1 Skycloud Expanse
1 Port Town
1 Glacial Fortress
1 Clifftop Retreat
1 Irrigated Farmland
1 Mystic Monastery
1 Temple of Epiphany
1 Temple of Enlightenment
1 Blast Zone
1 Rustvale Bridge
1 Razortide Bridge
1 Silverbluff Bridge
1 Perilous Landscape
6 Plains
9 Island
5 Mountain

1 Inspirit, Flagship Vessel
```

### Mesão (≤ R$350)

```
1 Hangarback Walker
1 Coretapper
1 Etherium Sculptor
1 Third Path Iconoclast
1 Enthusiastic Mechanaut
1 Brotherhood Vertibird
1 Foundry Inspector
1 Kilo, Apogee Mind
1 Malcator, Purity Overseer
1 Master of Etherium
1 Pinnacle Emissary
1 Sai, Master Thopterist
1 Surge Conductor
1 Chrome Host Seedshark
1 Uthros Research Craft
1 Warmaker Gunship
1 Jhoira, Weatherlight Captain
1 Leonin Abunas
1 Crystalline Crawler
1 Alibou, Ancient Witness
1 Deepglow Skate
1 Kappa Cannoneer
1 Thought Monitor
1 Padeem, Consul of Innovation
1 Cyberdrive Awakener
1 Everflowing Chalice
1 Sol Ring
1 Arcane Signet
1 Sphere of the Suns
1 Talisman of Progress
1 Empowered Autogenerator
1 Glass Casket
1 Pentad Prism
1 Boros Signet
1 Perilous Snare
1 Midnight Clock
1 Lux Artillery
1 Lux Cannon
1 Reckoner Bankbuster
1 Cloud Key
1 Thousand Moons Smithy
1 Thopter Spy Network
1 Whirlwind of Thought
1 Saheeli, Sublime Artificer
1 Dispatch
1 Swords to Plowshares
1 Loran's Escape
1 Blacksmith's Skill
1 Unwanted Remake
1 Disruption Protocol
1 Invisible Force Field
1 Thirst for Knowledge
1 Stern Lesson
1 Restoration Magic
1 Access Denied
1 Rip Apart
1 Tezzeret's Gambit
1 Chain Reaction
1 Fumigate
1 Reverse Engineer
1 Organic Extinction
1 Command Tower
1 Spire of Industry
1 Exotic Orchard
1 Battlefield Forge
1 Rugged Prairie
1 Skycloud Expanse
1 Port Town
1 Glacial Fortress
1 Clifftop Retreat
1 Irrigated Farmland
1 Mystic Monastery
1 Temple of Epiphany
1 Temple of Enlightenment
1 Blast Zone
1 Rustvale Bridge
1 Razortide Bridge
1 Silverbluff Bridge
1 Perilous Landscape
6 Plains
9 Island
5 Mountain

1 Inspirit, Flagship Vessel
```

## Impacto no goldfishing

O protocolo de `decks/inspirit-flagship-vessel/07-wincons.md` continua valendo, mas **as duas versões precisam ser testadas separadamente** — elas não têm a mesma curva, o mesmo número de fontes de draw, nem os mesmos caminhos de vitória. Pontos a observar em cada uma:

- **Torneio**: travas em azul (6 cartas duplo-azul, ~20 fontes); o comandante morrer a remoção pontual quando Leonin Abunas não está em jogo (só 1 hexproof estático + 2 one-shots); vitórias saindo mais tarde por falta do caminho C.
- **Mesão**: mãos travadas com peças de 5+ (9 cartas); se Access Denied for lenta demais na prática ({3}{U}{U} contra {U}{U} da Stoic Rebuttal), vale considerar trazer Stoic Rebuttal também para o mesão.
