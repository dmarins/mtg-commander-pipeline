# Relatório de Otimização — Inspirit, Flagship Vessel (v2)

> Segunda auditoria, feita sobre o resultado final da v1 (`decks/inspirit-flagship-vessel/report.md`), antes de qualquer goldfishing. Histórico completo da v1 preservado naquela pasta — este relatório documenta apenas o estado resultante da v2 e as trocas adicionais desta rodada.

> 🚨 **Orçamento invalidado (2026-08-12)**: todos os valores em US$ e as conversões `× 5,5` deste relatório são **estimativa por proxy da Scryfall e estão errados**. A régua do torneio é o **menor valor da LigaMagic**. Precificação real em **`09-precos-ligamagic.md`**: o deck atual custa **≈ R$ 220**, contra teto de R$ 200. Ignore as seções de custo abaixo até que sejam refeitas.

> ⚠️ **Correção de 2026-08-11**: este relatório originalmente descrevia o deck como **uma lista só**, misturando slots das versões torneio e mesão. Na prática existem **dois decks de 100 cartas** que compartilham 93 cartas e divergem em 7 slots — com composições, curvas e wincons diferentes. As seções abaixo passam a ser explícitas por versão; a composição completa, as duas listas de exportação e a comparação detalhada estão em **`08-versoes.md`**.

## Informações Gerais do Deck

- **Comandante**: Inspirit, Flagship Vessel
- **Cores do Deck**: Branco/Azul/Vermelho (Jeskai, identidade WUR)
- **Número de Cartas**: 100 (99 + Comandante) — **em cada versão** (torneio e mesão)
- **Foco do Deck**: Artefatos em massa (go-wide de thopters/golems/constructs) + marcadores de carga/+1/+1 + proteção estática do comandante ("other artifacts you control have hexproof and indestructible"). Subtemas: proliferate, token-makers, redução de custo, affinity, e um cluster identificado nesta v2 em torno de "cast a noncreature spell" (Third Path Iconoclast, Chrome Host Seedshark, Saheeli, Whirlwind of Thought).
- Quantidade de cartas por tipo:

| Tipo | Torneio | Mesão |
|---|---|---|
| Comandante (Spacecraft — **vira criatura só a 8+**) | 1 | 1 |
| Subcomandantes | 0 | 0 |
| Planeswalkers | 1 | 1 |
| **Criaturas de verdade** (entram como criatura, incl. artefatos-criatura) | **22** | **22** |
| **Veículos e Spacecraft** (artefatos; **só viram criatura com crew/station**) | **3** | **4** |
| Artefatos (não-criatura, fora veículos/spacecraft) | 15 | 15 |
| Instantâneos | 10 | 11 |
| Feitiços | 7 | 6 |
| Encantamentos | 3 | 2 |
| Terrenos Básicos por Cor | 6 Plains · 9 Island · 5 Mountain (20) | idem |
| Terrenos Não-Básicos | 18 | 18 |

> **Correção de 2026-08-12**: versões anteriores deste relatório contavam **25 criaturas** porque agrupavam Brotherhood Vertibird (Veículo), Uthros Research Craft e Warmaker Gunship (Spacecraft) junto das criaturas. Eles são **artefatos não-criatura em repouso** e dependem de outras criaturas para virar criatura. O deck tem **22 corpos de verdade**, não 25 — e o próprio comandante é um Spacecraft.

**Veículos e Spacecraft (o que precisam para virar criatura):**

| Carta | Requisito | Vale a pena atingir? |
|---|---|---|
| Inspirit, Flagship Vessel (comandante) | Station 8+ | Opcional — a estática (hexproof/indestructible) e o gatilho 1+ funcionam **desde 0 contadores**; 8+ só acrescenta corpo 5/5 com voar |
| Warmaker Gunship | Station 6+ | Sim — 4/3 voador; o ETB de dano já acontece sem station |
| Uthros Research Craft | Station 12+ | Na prática **não** — o valor está no limiar **3+** (compra por artefato conjurado); 12+ é inalcançável na maioria das partidas |
| Brotherhood Vertibird | Crew 2 | Sim — barato, e o poder é igual ao nº de artefatos |
| Reckoner Bankbuster *(só mesão)* | Crew 3 | Opcional — o draw por contador de carga não exige crew |

  - **Tokens**: gerados dinamicamente em jogo (thopters via Sai/Pinnacle Emissary/Thopter Spy Network/Access Denied na versão mesão, Golems 3/3, Servos 1/1 via Saheeli e Master Trinketeer na torneio, incubadores via Chrome Host Seedshark) — nenhum token pré-listado no 99
  - Bident of Thassa (torneio) é Encantamento **Artefato** — contada nos artefatos, não nos encantamentos.
- Quantidade de Draw, Ramp e Interação (após as 3 trocas da v2 + o corte de custo da torneio):

| Categoria | Torneio | Mesão | Meta |
|---|---|---|---|
| Fontes de draw | **14** | 13 | 12–13 ✓ |
| Ramp padrão | 10 | 10 | 10–11 ✓ |
| Ramp explosivo | 1 | 1 | 2–3 (abaixo, **aceito conscientemente** por excesso geral de ramp) |
| Redutores de custo | 3 | 4 | — |
| Interação pontual | ~9 | ~9 | ~10 (ligeiramente abaixo após o corte do Chainsaw) |
| Counterspells | 2 (Disruption Protocol + Stoic Rebuttal) | 2 (Disruption Protocol + Access Denied) | dobrou em relação à v1 (1) |
| Board wipes | 3 | 3 | 2–4 ✓ |
| Hexproof estático no comandante | **1** (Leonin Abunas) | **2** (Padeem + Leonin Abunas) | — |
| Indestructible sob demanda | 1 (Invisible Force Field) | 1 (Invisible Force Field) | — |
| One-shots de proteção | **2** (Loran's Escape, Blacksmith's Skill) | **3** (+ Restoration Magic) | — |

> A versão torneio troca **proteção por consistência**: perde uma camada de hexproof e um one-shot (Padeem e Restoration Magic saíram no corte de custo), mas ganha uma fonte de draw e uma curva mais baixa.

## Curva de Mana

- **Número de Terrenos Básicos**: 20 (nas duas versões)
- **Número de Terrenos Não-Básicos**: 18 (nas duas versões)

| Faixa de CMC | Torneio | Mesão |
|---|---|---|
| 0–1 | 8 | 9 |
| 2 | 14 | 14 |
| 3 | 19 | 17 |
| 4 | 13 | 12 |
| 5 | 4 | 5 |
| 6+ | 3 | 4 |
| **CMC médio** (sem terrenos, X = 0) | **3,05** | **3,08** |

- **Símbolos coloridos nos 61 não-terrenos**: torneio — U 32 (50%) · W 19 (30%) · R 11 (17%) · U/R híbrido 2 · mesão — U 29 (46%) · W 21 (33%) · R 11 (17%) · U/R híbrido 2
- **Cartas com custo duplo-azul**: torneio **6** · mesão **4** — ver alerta de manabase em `08-versoes.md` (a torneio fica no limite inferior da conta de Karsten para {U}{U} no turno 3)

## Sinergias e Estratégias

- **Estratégia Geral**: "Artifacts matter" go-wide — Inspirit acumula marcadores via Station e o gatilho 1+ (combate), enquanto sua estática protege a mesa de artefatos (hexproof + indestructible), exceto a si mesmo. A v2 fechou justamente essa exceção: Invisible Force Field cobre o Inspirit com indestructible sob demanda (nas duas versões), e o hexproof vem de Leonin Abunas — **em redundância dupla com Padeem só na versão mesão**, já que Padeem saiu no corte de custo da torneio.
- **Principais Sinergias**:
  - Proliferate (Kilo, Surge Conductor, Tezzeret's Gambit, Deepglow Skate) acelera charge counters em Spacecraft, Lux Cannon/Artillery e rocks de counters.
  - **Economia de Station** (o deck tem 22 corpos, não 25): Station custa **tapar outra criatura sua**, em velocidade de feitiço, e rende contadores iguais ao **poder** dela. Como a maioria dos corpos é pequena (1/1 de thopter, servo, drone), pagar os 8 contadores do Inspirit tapando tokens é lento. Os atalhos reais são: **Master of Etherium** (poder = nº de artefatos, costuma valer 8–12 sozinho); **Brotherhood Vertibird** crewado por 2 de poder e então tapado para Station (poder = nº de artefatos — crew primeiro, station depois, tudo na fase principal); **Deepglow Skate** dobrando os contadores já acumulados; e o pacote de Coretapper/proliferate, que sobe contadores **sem tapar criatura nenhuma**. Criaturas com poder ≥ 3 disponíveis para station: 9 nas duas versões. Station funciona com criaturas recém-jogadas — a summoning sickness não impede, exatamente como em crew ([Draftsim](https://draftsim.com/mtg-station/), [TheGamer](https://www.thegamer.com/magic-the-gathering-mtg-spacecraft-stationing-guide/)).
  - **Alibou** exige que **artefatos-criatura ataquem** para disparar (X = artefatos virados). São 14 artefatos-criatura de verdade nas duas versões — os Spacecraft só entram nessa conta depois de estacionados, e o Inspirit só ataca a partir de 8+.
  - **Cluster "cast noncreature spell"** (achado da v2): Third Path Iconoclast, Chrome Host Seedshark, Saheeli e Whirlwind of Thought disparam juntos a cada spell não-criatura — a maioria do deck, dado o volume de artefatos/instants/sorceries.
  - **Wipes assimétricos a favor do jogador** (Chain Reaction, Organic Extinction, Fumigate — os três confirmados na v2): "destroy all creatures/permanents" mata só os não-artefatos, porque as criaturas-artefato do jogador são indestrutíveis via Inspirit.
  - Redutores de custo aceleram a curva de ~40 artefatos: Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector nas duas versões, **+ Cloud Key só no mesão** (na torneio o slot virou Prophetic Prism, que fixa cor e repõe a si mesma em vez de reduzir custo).
  - Fixação de mana desredundantizada na v2: Talisman of Progress (W/U) + Boros Signet (R/W) cobrem as 3 cores sem duplicar W/U como antes (Azorius Signet cortado).
- **Cartas-Chave**: Inspirit, Flagship Vessel (motor e proteção); Leonin Abunas + Invisible Force Field nas duas versões (**+ Padeem só no mesão**, que é o único com hexproof duplo); Master of Etherium e Alibou (finishers principais); o 2º counterspell (Stoic Rebuttal na torneio, Access Denied no mesão) como camada contra combo/bomba alheia.
- **Pontos Fortes**: proteção cobre as duas linhas de defesa (targeted e não-targeted) contra remoção do comandante; 2 counterspells em vez de 1; manabase sem redundância de fixação; curva bem distribuída. Na torneio somam-se dois anthems (Chief of the Foundry, Master Trinketeer) e 1 fonte de draw a mais.
- **Pontos Fracos**: interação pontual caiu ligeiramente (9, era 9 na v1 também mas agora sem Chainsaw como meia-peça extra) — compensado pela 2ª camada de counter; ramp explosivo abaixo da meta (1, aceito conscientemente). **Só na torneio**: uma camada de hexproof a menos, um one-shot de proteção a menos, perda do wincon C e 6 cartas duplo-azul contra ~20 fontes de azul.

## Combos e Condições de Vitória

- **Combos Principais**: nenhum combo de 2 peças ou infinito identificado (herdado da v1, sem alteração nesta rodada).
- **Condições de Vitória**: as 3 trocas da v2 não afetaram nenhum caminho, mas **o corte de custo da versão torneio afeta** — Cyberdrive Awakener e Thousand Moons Smithy só existem no mesão.

| Caminho | Torneio | Mesão |
|---|---|---|
| **A. Enxame de combate** (turno 6–8): Master of Etherium + token-makers + Kappa Cannoneer + Alibou | ✅ e **reforçado** — dois anthems (Chief of the Foundry, Master Trinketeer) | ✅ — com Cyberdrive Awakener somando corpos e evasão |
| **B. Queima via marcadores** (turno 7–9): Lux Artillery aos 30+ marcadores, sem depender de combate | ✅ idêntico | ✅ idêntico |
| **C. Conversão explosiva de rocks** (turno 6–8): Cyberdrive Awakener + Alibou + pacote de ramp | ❌ **cortado** no corte de custo | ✅ |

> Torneio = mais consistência, teto mais baixo (2 caminhos, curva menor, +1 draw). Mesão = mais teto e mais proteção (3 caminhos, hexproof duplo), com curva mais pesada.

## Lista de Cartas por Tipo

**Base comum às duas versões (93 cartas, incluindo comandante e terrenos):**

- **Comandante**: Inspirit, Flagship Vessel — CMC 3, Artefato Lendário Spacecraft, Station, gatilho 1+ (+1/+1 ou 2 charges por combate), estática "other artifacts have hexproof and indestructible".
- **Planeswalkers**: Saheeli, Sublime Artificer — CMC 3, Servo 1/1 por spell não-criatura (cluster v2); -2 copia artefato.
- **Criaturas de verdade** (20 comuns — entram como criatura):
  - Hangarback Walker, Coretapper, Etherium Sculptor, Third Path Iconoclast (cluster noncreature), Enthusiastic Mechanaut, Foundry Inspector, Kilo Apogee Mind, Malcator Purity Overseer, Master of Etherium, Pinnacle Emissary, Sai Master Thopterist, Surge Conductor, Chrome Host Seedshark (cluster noncreature), Jhoira Weatherlight Captain, Leonin Abunas, Crystalline Crawler, Alibou Ancient Witness, Deepglow Skate, Kappa Cannoneer, Thought Monitor.
  - (Ver `decks/inspirit-flagship-vessel/report.md` para descrição carta a carta — nenhuma mudou nesta v2.)
- **Veículos e Spacecraft** (3 comuns — artefatos, **não** criaturas em repouso): Brotherhood Vertibird (Veículo, crew 2, poder = nº de artefatos), Uthros Research Craft (Spacecraft, 3+ compra por artefato conjurado, criatura só a 12+), Warmaker Gunship (Spacecraft, ETB de dano, criatura a 6+).
- **Artefatos** (13 comuns): Everflowing Chalice, Sol Ring, Arcane Signet, Sphere of the Suns, Talisman of Progress, Empowered Autogenerator, Glass Casket, Pentad Prism, **Boros Signet** (novo — R/W, substitui Azorius Signet), Perilous Snare, Midnight Clock, Lux Artillery, Lux Cannon.
  - *(Removidos desde a v1: Gilded Lotus, Azorius Signet, Chainsaw.)*
- **Encantamentos** (2 comuns): Thopter Spy Network; Whirlwind of Thought (cluster noncreature).
- **Instantâneos** (9 comuns):
  - Dispatch, Swords to Plowshares, Loran's Escape, Blacksmith's Skill, Unwanted Remake — CMC 1; Disruption Protocol — CMC 2.
  - **Invisible Force Field** (novo — CMC 2, {1}{W}, até 4 permanentes seus ganham indestructible + Rebound) — substitui Gilded Lotus.
  - Thirst for Knowledge — CMC 3; **Stern Lesson** — CMC 3 (é **Instantâneo**, não Feitiço: estava classificada errado nas versões anteriores deste relatório).
- **Feitiços** (6 comuns): Rip Apart, Tezzeret's Gambit, Chain Reaction, Fumigate (reconfirmado assimétrico na v2), Reverse Engineer, Organic Extinction.
- **Terrenos** (38, idênticos nas duas versões): 6 Plains, 9 Island, 5 Mountain, Command Tower, Spire of Industry, Exotic Orchard, Battlefield Forge, Rugged Prairie, Skycloud Expanse, Port Town, Glacial Fortress, Clifftop Retreat, Irrigated Farmland, Mystic Monastery, Temple of Epiphany, Temple of Enlightenment, Blast Zone, Rustvale Bridge, Razortide Bridge, Silverbluff Bridge, Perilous Landscape — sem alteração da v1.

**Os 7 slots que divergem:**

| # | Torneio | Tipo | Mesão | Tipo |
|---|---|---|---|---|
| 1 | Reconnaissance Mission ({2}{U}{U}) | Encantamento — draw | Padeem, Consul of Innovation ({3}{U}) | Criatura Lendária — draw + hexproof |
| 2 | Bident of Thassa ({2}{U}{U}) | Encantamento **Artefato** — draw (protegido pelo Inspirit) | Reckoner Bankbuster ({2}) | Artefato — Veículo — draw |
| 3 | Stoic Rebuttal ({1}{U}{U}) | Instantâneo — counter (metalcraft → {U}{U}) | Access Denied ({3}{U}{U}) | Instantâneo — counter + X Thopters |
| 4 | Master Trinketeer ({2}{W}) | Criatura — anthem de Servos/Thopters | Thousand Moons Smithy (4) | Artefato Lendário — token */* |
| 5 | Prophetic Prism ({2}) | Artefato — fixação + cantrip | Cloud Key ({3}) | Artefato — redutor de custo |
| 6 | One with the Machine ({3}{U}) | **Feitiço** — draw | Restoration Magic ({W}) | Instantâneo — proteção |
| 7 | Chief of the Foundry ({3}) | Criatura Artefato — anthem | Cyberdrive Awakener ({5}{U}) | Criatura Artefato — wincon C |

Outras classificações corrigidas nesta revisão: **Cyberdrive Awakener** e **Kilo, Apogee Mind** são artefatos-criatura (portanto **protegidos** pela estática do Inspirit); **Thousand Moons Smithy** é MDFC (`// Barracks of the Thousand`, Terreno Artefato Lendário na face de trás).

## Listas de Cartas para Exportação (padrão MTG Online)

São **duas listas de 100 cartas**, não uma. Ambas estão em `08-versoes.md`, que é a fonte de verdade das versões:

- **Torneio (≤ R$200)** — `08-versoes.md`, seção "Listas de exportação → Torneio"
- **Mesão (≤ R$350)** — `08-versoes.md`, seção "Listas de exportação → Mesão"

Para montar uma a partir da outra, troque os 7 slots da tabela em "Lista de Cartas por Tipo" acima. As outras 93 cartas (comandante + 54 não-terrenos + 38 terrenos) são idênticas.


## Mudanças Aplicadas (auditoria v2 — segunda rodada, sobre o resultado já otimizado da v1)

Esta v2 partiu do deck **já otimizado e ainda não testado** da v1 (`decks/inspirit-flagship-vessel/report.md`) e fez uma auditoria cética do zero, sem aceitar automaticamente nenhuma escolha anterior.

| # | Sai | Entra | CMC | US$ | Faixa | Motivo |
|---|---|---|---|---|---|---|
| 1 | Gilded Lotus | Invisible Force Field | 2 | 0.44 | torneio | Gilded Lotus era o rock mais fraco (sem contadores/corpo, redundante com as outras 10 fontes de ramp). Invisible Force Field fecha a lacuna real: nenhuma carta dava indestructible ao próprio Inspirit (só hexproof via Padeem/Leonin Abunas) — protege até 4 permanentes + Rebound, cobrindo também o cluster de criaturas não-artefato expostas. |
| 2 | Azorius Signet | Boros Signet | 2 | 0.42 | torneio | Cortava a redundância W/U com Talisman of Progress (mantido, mais eficiente). Boros Signet fixa R/W — as duas cores mais escassas do deck (5 Mountain, 6 Plains) — sem reforçar ainda mais o U, já a cor mais bem servida. |
| 3 | Chainsaw | Stoic Rebuttal (alt. mesão: Access Denied) | 3 | 0.32 (Access Denied: 2.62) | torneio (mesão) | Chainsaw era a peça mais fraca/isolada da remoção — confirmado via regra oficial que o rev counter rende só 1 por evento de morte simultânea, não 1 por criatura (motor mais lento do que um relato de partida sugeria). Fecha a lacuna real de counterspell (1→2). Access Denied é a alternativa mesão: mais cara, mas gera X Thopters (X = mana value do spell anulado) que alimentam o próprio tema. |

**Achados que corrigiram entendimentos da v1** (sem trocar cartas): Fumigate é wipe assimétrico a favor do jogador — a v1 não tinha deixado isso explícito; a razão original para cogitar cortar Chainsaw na v1 (rev counter não seria alimentado por mortes alheias) estava errada — o rev counter conta qualquer morte, mas mesmo corrigido o entendimento, a carta acabou saindo por outro motivo (peça mais fraca da remoção, para abrir espaço ao 2º counter).

**Observação registrada sem ação**: possível ganho adicional ao revisar o pacote de wincons/finishers com o cluster "cast noncreature spell" recém-identificado — não avaliado nesta rodada, fica como sugestão para uma eventual v3.

**Custo total dos 3 swaps da v2**: ≈ US$1,18 (versão 100% torneio) — Invisible Force Field (0.44) + Boros Signet (0.42) + Stoic Rebuttal (0.32).

## Montagem do Deck — Sai/Entra por Versão (pós-compra, 2026-07-25)

Compra já realizada pelo usuário. Esta seção registra, de forma consolidada, quais cartas saem da decklist **original** (antes de qualquer otimização) e quais entram no lugar — cobrindo todo o histórico (v1 + v2) — separadas em trocas que valem para as duas versões e trocas exclusivas da versão torneio (corte de custo pra caber no teto de R$200).

### 1. Trocas que valem para as duas versões (mesão e torneio)

| # | Sai (decklist original) | Entra (torneio) | Entra (mesão) |
|---|---|---|---|
| 1 | Astral Cornucopia | Arcane Signet | Arcane Signet |
| 2 | Azorius Signet | Boros Signet | Boros Signet |
| 3 | Boros Garrison | Blast Zone | Blast Zone |
| 4 | Captain Storm, Cosmium Raider | Empowered Autogenerator | Empowered Autogenerator |
| 5 | Cargo Ship | Talisman of Progress | Talisman of Progress |
| 6 | Chainsaw | **Stoic Rebuttal** | **Access Denied** |
| 7 | Diversion Unit | Disruption Protocol | Disruption Protocol |
| 8 | Emissary Escort | *(corte simples, financia terreno extra)* | *(idem)* |
| 9 | Ethersworn Sphinx | Reverse Engineer | Reverse Engineer |
| 10 | Frogmyr Enforcer | **Reconnaissance Mission** | **Padeem, Consul of Innovation** |
| 11 | Golem Foundry | Invisible Force Field *(via Gilded Lotus na v1, cortada de novo na v2 antes de comprar)* | Invisible Force Field |
| 12 | Marketback Walker | Thopter Spy Network | Thopter Spy Network |
| 13 | Memory Guardian | Whirlwind of Thought | Whirlwind of Thought |
| 14 | Phyrexian Revoker | Leonin Abunas | Leonin Abunas |
| 15 | Solar Array | Midnight Clock | Midnight Clock |
| 16 | Temple of Triumph | +1 Island | +1 Island |
| 17 | Thopter Fabricator | Fumigate | Fumigate |
| 18 | Voyage Home | **Bident of Thassa** | **Reckoner Bankbuster** |
| 19 | Voyager Quickwelder | *(corte simples, financia terreno extra)* | *(idem)* |
| 20 | Zookeeper Mechan | Sphere of the Suns | Sphere of the Suns |
| — | *(ajuste de básicos)* | +2 Island | +2 Island |

As 3 linhas em **negrito** (6, 10, 18) são as únicas que mudam de carta dependendo da versão — todo o resto é idêntico entre torneio e mesão. Island sobe de 6 para 9 no total (+1 da troca da Temple of Triumph, +2 do ajuste de básicos).

### 2. Trocas exclusivas da versão torneio (corte de custo, para caber em ≤ R$200)

A versão mesão **mantém as 4 cartas originais abaixo** — só a torneio troca, para tirar o deck de R$202,68 (acima do teto) e trazê-lo para R$193,93.

| Sai (só torneio — mesão mantém) | Entra (só torneio) | US$ | Motivo |
|---|---|---|---|
| Thousand Moons Smithy | Master Trinketeer | 0,29 | Única impressão existente (US$2,56+); substituto barateia e ainda buffa Servos/Thopters |
| Cloud Key | Prophetic Prism | 0,21 | Redutor de custo redundante (já há 3 outros: Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector); substituto dá draw + fixação de mana |
| Restoration Magic | One with the Machine | 0,25 | Proteção redundante (Loran's Escape + Blacksmith's Skill já cobrem); substituto já era reserva aprovada em `03-draw.md` |
| Cyberdrive Awakener | Chief of the Foundry | 0,20 | Corta o wincon C só na torneio (A e B seguem intactos); substituto já era pendência aprovada da v1 |

### Custo final do deck completo (100 cartas), por versão

| Versão | Total | Teto |
|---|---|---|
| Torneio | US$36,85 (~R$202,68) | R$200 — praticamente no limite (diferença dentro do ruído da conversão ×5,5) |
| Mesão | US$49,89 (~R$274,40) | R$350 — folgado |

Nota: não se aplica troca de impressão/edição em cartas já possuídas antes da v1 (Kilo Apogee Mind, Sol Ring, Etherium Sculptor, Pinnacle Emissary) — só faz sentido buscar a impressão mais barata em compras novas de fato (ex.: Leonin Abunas, comprada na edição Mirrodin 2003 por US$1,33).

## Protocolo de Goldfishing

O protocolo em si não muda — usar o já definido em `decks/inspirit-flagship-vessel/07-wincons.md` (preparação, sequência de turnos, tabela de registro para 5+ partidas, meta de vitória projetada ≤ turno 7). As 3 trocas da v2 não afetam nenhuma das condições de vitória validadas naquela fase; a v1 nunca chegou a ser testada, então o goldfishing deve usar as listas desta v2.

**Rode as duas versões separadamente** — elas não têm a mesma curva, o mesmo número de fontes de draw nem os mesmos caminhos de vitória. O que observar em cada uma está em `08-versoes.md`, seção "Impacto no goldfishing".
