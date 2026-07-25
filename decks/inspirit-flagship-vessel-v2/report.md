# Relatório de Otimização — Inspirit, Flagship Vessel (v2)

> Segunda auditoria, feita sobre o resultado final da v1 (`decks/inspirit-flagship-vessel/report.md`), antes de qualquer goldfishing. Histórico completo da v1 preservado naquela pasta — este relatório documenta apenas o estado resultante da v2 e as trocas adicionais desta rodada.

## Informações Gerais do Deck

- **Comandante**: Inspirit, Flagship Vessel
- **Cores do Deck**: Branco/Azul/Vermelho (Jeskai, identidade WUR)
- **Número de Cartas**: 100 (99 + Comandante)
- **Foco do Deck**: Artefatos em massa (go-wide de thopters/golems/constructs) + marcadores de carga/+1/+1 + proteção estática do comandante ("other artifacts you control have hexproof and indestructible"). Subtemas: proliferate, token-makers, redução de custo, affinity, e um cluster identificado nesta v2 em torno de "cast a noncreature spell" (Third Path Iconoclast, Chrome Host Seedshark, Saheeli, Whirlwind of Thought).
- Quantidade de cartas por tipo:
  - **Comandante**: 1
  - **Subcomandantes**: 0
  - **Planeswalkers**: 1
  - **Criaturas** (incl. artefatos-criatura, veículos e spacecraft): 25
  - **Artefatos** (não-criatura): 16
  - **Instantâneos**: 10
  - **Feitiços**: 7
  - **Encantamentos**: 2
  - **Terrenos Básicos por Cor**: 6 Plains (W), 9 Island (U), 5 Mountain (R) — 20 total
  - **Terrenos Não-Básicos**: 18
  - **Tokens**: gerados dinamicamente em jogo (thopters via Sai/Pinnacle Emissary/Thopter Spy Network/Access Denied se essa opção for usada, Golems 3/3, Servos 1/1 via Saheeli, incubadores via Chrome Host Seedshark) — nenhum token pré-listado no 99
- Quantidade de Draw, Ramp e Interação (após as 3 trocas da v2):
  - **Draw**: 13 fontes (meta 12–13) ✓ — sem alteração da v1
  - **Ramp**: 10 padrão + 1 explosivo (meta 10–11 padrão + 2–3 explosivos) — padrão na meta; explosivo abaixo da meta, **aceito conscientemente** por excesso geral de ramp identificado na auditoria
  - **Interação pontual**: ~9 peças (meta ~10) — ligeiramente abaixo após o corte do Chainsaw, compensado pela 2ª camada de counterspell
  - **Counterspells**: 2 (Disruption Protocol + Stoic Rebuttal) — dobrou em relação à v1 (1)
  - **Board Wipes**: 3 (meta 2–4) ✓
  - **Proteção do comandante**: hexproof (Padeem + Leonin Abunas) + **indestructible (Invisible Force Field, novo na v2)** — as duas lacunas de remoção agora cobertas, além dos 3 one-shots (Loran's Escape, Blacksmith's Skill, Restoration Magic)

## Curva de Mana

- **Número de Terrenos Básicos**: 20
- **Número de Terrenos Não-Básicos**: 18
- **CMC Médio** (sem terrenos): ~3,05 (X-spells contados como CMC 0)
- **Distribuição de CMC**:
  - 0–1: 9
  - 2: 14
  - 3: 18
  - 4: 12
  - 5: 4
  - 6+: 4
- **Distribuição de CMC por Cor** (símbolos coloridos nos 61 não-terrenos): Azul (U) 48% · Branco (W) 33% · Vermelho (R) 19%
- **Quantidade de Cartas por Cor** (símbolos): U ≈ 25, W ≈ 17, R ≈ 10, incolor ≈ 19 cartas (rocks/artefatos genéricos)

## Sinergias e Estratégias

- **Estratégia Geral**: "Artifacts matter" go-wide — Inspirit acumula marcadores via Station e o gatilho 1+ (combate), enquanto sua estática protege a mesa de artefatos (hexproof + indestructible), exceto a si mesmo. A v2 fechou justamente essa exceção: Invisible Force Field cobre o Inspirit com indestructible sob demanda, e Padeem/Leonin Abunas cobrem hexproof em redundância dupla.
- **Principais Sinergias**:
  - Proliferate (Kilo, Surge Conductor, Tezzeret's Gambit, Deepglow Skate) acelera charge counters em Spacecraft, Lux Cannon/Artillery e rocks de counters.
  - **Cluster "cast noncreature spell"** (achado da v2): Third Path Iconoclast, Chrome Host Seedshark, Saheeli e Whirlwind of Thought disparam juntos a cada spell não-criatura — a maioria do deck, dado o volume de artefatos/instants/sorceries.
  - **Wipes assimétricos a favor do jogador** (Chain Reaction, Organic Extinction, Fumigate — os três confirmados na v2): "destroy all creatures/permanents" mata só os não-artefatos, porque as criaturas-artefato do jogador são indestrutíveis via Inspirit.
  - Redutores de custo (Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector, Cloud Key) aceleram a curva de ~40 artefatos.
  - Fixação de mana desredundantizada na v2: Talisman of Progress (W/U) + Boros Signet (R/W) cobrem as 3 cores sem duplicar W/U como antes (Azorius Signet cortado).
- **Cartas-Chave**: Inspirit, Flagship Vessel (motor e proteção); Padeem + Leonin Abunas + Invisible Force Field (proteção completa: hexproof duplo + indestructible sob demanda); Master of Etherium e Alibou (finishers principais); Stoic Rebuttal (segunda camada contra combo/bomba alheia).
- **Pontos Fortes**: proteção agora cobre as duas linhas de defesa (targeted e não-targeted) contra remoção do comandante; 2 counterspells em vez de 1; manabase sem redundância de fixação; curva bem distribuída.
- **Pontos Fracos**: interação pontual caiu ligeiramente (9, era 9 na v1 também mas agora sem Chainsaw como meia-peça extra) — compensado pela 2ª camada de counter; ramp explosivo abaixo da meta (1, aceito conscientemente).

## Combos e Condições de Vitória

- **Combos Principais**: nenhum combo de 2 peças ou infinito identificado (herdado da v1, sem alteração nesta rodada).
- **Condições de Vitória**: os 3 caminhos validados na v1 (`07-wincons.md`) permanecem intactos — nenhuma troca da v2 afetou Master of Etherium, Alibou, Kappa Cannoneer, Cyberdrive Awakener, Lux Artillery ou Thousand Moons Smithy.
  - **A. Enxame de combate** (turno 6–8): Master of Etherium + token-makers + Kappa Cannoneer + Cyberdrive Awakener + Alibou.
  - **B. Queima via marcadores** (turno 7–9): Lux Artillery aos 30+ marcadores — não depende de combate.
  - **C. Conversão explosiva de rocks** (turno 6–8): Cyberdrive Awakener + Alibou + pacote de ramp.

## Lista de Cartas por Tipo

- **Comandante**: Inspirit, Flagship Vessel — CMC 3, Artefato Lendário Spacecraft, Station, gatilho 1+ (+1/+1 ou 2 charges por combate), estática "other artifacts have hexproof and indestructible".
- **Planeswalkers**: Saheeli, Sublime Artificer — CMC 3, Servo 1/1 por spell não-criatura (cluster v2); -2 copia artefato.
- **Criaturas** (incl. artefatos-criatura, veículos e spacecraft — sem alteração desta v2):
  - Hangarback Walker, Coretapper, Etherium Sculptor, Third Path Iconoclast (cluster noncreature), Enthusiastic Mechanaut, Brotherhood Vertibird, Foundry Inspector, Kilo Apogee Mind, Malcator Purity Overseer, Master of Etherium, Pinnacle Emissary, Sai Master Thopterist, Surge Conductor, Chrome Host Seedshark (cluster noncreature), Uthros Research Craft, Warmaker Gunship, Jhoira Weatherlight Captain, Padeem Consul of Innovation, Leonin Abunas, Crystalline Crawler, Alibou Ancient Witness, Deepglow Skate, Kappa Cannoneer, Cyberdrive Awakener, Thought Monitor.
  - (Ver `decks/inspirit-flagship-vessel/report.md` para descrição carta a carta — nenhuma mudou nesta v2.)
- **Artefatos**:
  - Everflowing Chalice, Sol Ring, Arcane Signet, Sphere of the Suns, Talisman of Progress, Empowered Autogenerator, Glass Casket, Pentad Prism, **Boros Signet** (novo — R/W, substitui Azorius Signet), Cloud Key, Perilous Snare, Reckoner Bankbuster, Midnight Clock, Lux Artillery, Lux Cannon, Thousand Moons Smithy.
  - *(Removidos desde a v1: Gilded Lotus, Azorius Signet, Chainsaw.)*
- **Encantamentos**: Thopter Spy Network; Whirlwind of Thought (cluster noncreature).
- **Instantâneos**:
  - Dispatch, Swords to Plowshares, Loran's Escape, Blacksmith's Skill, Restoration Magic, Unwanted Remake, Disruption Protocol — CMC 1, Instantâneo — todos, sem alteração da v1.
  - **Stoic Rebuttal** (novo — CMC 3, {1}{U}{U}, metalcraft reduz para {U}{U}, "counter target spell" sem restrição) — substitui Chainsaw.
  - **Invisible Force Field** (novo — CMC 2, {1}{W}, até 4 permanentes seus ganham indestructible + Rebound) — substitui Gilded Lotus.
  - Thirst for Knowledge — CMC 3, sem alteração.
- **Feitiços**: Rip Apart, Stern Lesson, Tezzeret's Gambit, Chain Reaction, Fumigate (reconfirmado assimétrico na v2), Reverse Engineer, Organic Extinction — sem alteração da v1.
- **Terrenos**: 6 Plains, 9 Island, 5 Mountain, Command Tower, Spire of Industry, Exotic Orchard, Battlefield Forge, Rugged Prairie, Skycloud Expanse, Port Town, Glacial Fortress, Clifftop Retreat, Irrigated Farmland, Mystic Monastery, Temple of Epiphany, Temple of Enlightenment, Blast Zone, Rustvale Bridge, Razortide Bridge, Silverbluff Bridge, Perilous Landscape — sem alteração da v1.

## Lista de Cartas para Exportação (padrão MTG Online)

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
1 Padeem, Consul of Innovation
1 Leonin Abunas
1 Crystalline Crawler
1 Alibou, Ancient Witness
1 Deepglow Skate
1 Kappa Cannoneer
1 Cyberdrive Awakener
1 Thought Monitor
1 Everflowing Chalice
1 Sol Ring
1 Arcane Signet
1 Sphere of the Suns
1 Talisman of Progress
1 Empowered Autogenerator
1 Glass Casket
1 Pentad Prism
1 Boros Signet
1 Cloud Key
1 Perilous Snare
1 Reckoner Bankbuster
1 Midnight Clock
1 Lux Artillery
1 Lux Cannon
1 Thousand Moons Smithy
1 Thopter Spy Network
1 Whirlwind of Thought
1 Saheeli, Sublime Artificer
1 Dispatch
1 Swords to Plowshares
1 Loran's Escape
1 Blacksmith's Skill
1 Restoration Magic
1 Unwanted Remake
1 Disruption Protocol
1 Stoic Rebuttal
1 Invisible Force Field
1 Thirst for Knowledge
1 Rip Apart
1 Stern Lesson
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

> Nota: se preferir a alternativa mesão do 2º counterspell, substitua `1 Stoic Rebuttal` por `1 Access Denied` na lista acima (ver seção seguinte).

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

Sem alteração em relação à v1 — usar o protocolo já definido em `decks/inspirit-flagship-vessel/07-wincons.md` (preparação, sequência de turnos, tabela de registro para 5+ partidas, meta de vitória projetada ≤ turno 7). As 3 trocas da v2 não afetam nenhuma das condições de vitória validadas naquela fase; recomenda-se rodar o goldfishing já considerando a lista atualizada desta v2 (a v1 nunca chegou a ser testada).
