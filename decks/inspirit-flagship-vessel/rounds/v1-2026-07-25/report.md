# Relatório de Otimização — Inspirit, Flagship Vessel

## Informações Gerais do Deck

- **Comandante**: Inspirit, Flagship Vessel
- **Cores do Deck**: Branco/Azul/Vermelho (Jeskai, identidade WUR)
- **Número de Cartas**: 100 (99 + Comandante)
- **Foco do Deck**: Artefatos em massa (go-wide de thopters/golems/constructs) + marcadores de carga/+1/+1 + proteção estática do comandante ("other artifacts you control have hexproof and indestructible"). Subtemas: proliferate, token-makers, redução de custo, affinity.
- Quantidade de cartas por tipo:
  - **Comandante**: 1
  - **Subcomandantes**: 0
  - **Planeswalkers**: 1
  - **Criaturas** (incl. artefatos-criatura, veículos e spacecraft): 25
  - **Artefatos** (não-criatura): 18
  - **Instantâneos**: 8
  - **Feitiços**: 7
  - **Encantamentos**: 2
  - **Terrenos Básicos por Cor**: 6 Plains (W), 9 Island (U), 5 Mountain (R) — 20 total
  - **Terrenos Não-Básicos**: 18
  - **Tokens**: gerados dinamicamente em jogo (thopters via Sai/Pinnacle Emissary/Thopter Spy Network, Golems 3/3, Servos 1/1 via Saheeli, incubadores via Chrome Host Seedshark) — nenhum token pré-listado no 99
- Quantidade de Draw, Ramp e Interação (metas do pipeline, batidas nas fases de otimização):
  - **Draw**: 13 fontes (meta 12–13) ✓
  - **Ramp**: 10 padrão + 2 explosivos (meta 10–11 padrão + 2–3 explosivos) ✓
  - **Interação pontual**: ~10 peças (meta ~10) ✓
  - **Board Wipes**: 3 (meta 2–4) ✓

## Curva de Mana

- **Número de Terrenos Básicos**: 20
- **Número de Terrenos Não-Básicos**: 18
- **CMC Médio** (sem terrenos): ~3,03 (X-spells contados como CMC 0)
- **Distribuição de CMC**:
  - 0–1: 9
  - 2: 14
  - 3: 17
  - 4: 12
  - 5: 5
  - 6+: 4
- **Distribuição de CMC por Cor** (símbolos coloridos nos 61 não-terrenos):
  - Azul (U): 46% dos símbolos — cor mais exigida, refletida nos 9 Island
  - Branco (W): 31% dos símbolos
  - Vermelho (R): 23% dos símbolos — cor de suporte, menor exigência
- **Quantidade de Cartas por Cor** (aprox., contagem de identidade nos não-terrenos): U ≈ 24 símbolos, W ≈ 16 símbolos, R ≈ 11 símbolos, incolor ≈ 20 cartas (rocks/artefatos genéricos)

## Sinergias e Estratégias

- **Estratégia Geral**: "Artifacts matter" go-wide — Inspirit acumula marcadores de carga via Station e o gatilho 1+ (combate), enquanto sua estática protege toda a mesa de artefatos (hexproof + indestructible) menos ele mesmo. O deck constrói um exército de artefatos-criatura protegidos e fecha o jogo por combate, queima via marcadores ou conversão explosiva de mana rocks em ameaças.
- **Principais Sinergias**:
  - Proliferate (Kilo, Surge Conductor, Tezzeret's Gambit, Deepglow Skate) acelera charge counters em Spacecraft, Lux Cannon/Artillery e mana rocks de counters (Everflowing Chalice, Sphere of the Suns, Pentad Prism).
  - Token-makers (Sai, Saheeli, Pinnacle Emissary, Thopter Spy Network) alimentam Master of Etherium (lord) e Kappa Cannoneer (cresce por artefato).
  - Redutores de custo (Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector, Cloud Key) aceleram a curva de ~40 artefatos do deck.
  - A estática do Inspirit protege quase todos os finishers, que são eles mesmos artefatos (Master of Etherium, Alibou, Kappa Cannoneer, Cyberdrive Awakener, Thousand Moons Smithy).
- **Cartas-Chave**: Inspirit, Flagship Vessel (motor e proteção); Padeem + Leonin Abunas (redundância de proteção); Master of Etherium e Alibou (finishers principais); Lux Artillery (via alternativa de vitória).
- **Pontos Fortes**: proteção estática ampla, 3 caminhos de vitória distintos, curva bem distribuída após otimização, manabase corrigida para a demanda real de pips.
- **Pontos Fracos**: dependência de ter o Inspirit em jogo para a proteção mais ampla (mitigado por 5 camadas de proteção); Caminho B (queima via marcadores) depende de acumular volume ao longo de vários turnos.

## Combos e Condições de Vitória

- **Combos Principais**: nenhum combo de 2 peças ou infinito identificado — condizente com o power level casual do briefing e com a regra da diversão do pipeline.
- **Condições de Vitória** (3 caminhos, detalhados em `07-wincons.md`):
  - **A. Enxame de combate** (turno 6–8, consistência alta): Master of Etherium como lord + token-makers + Kappa Cannoneer (unblockable) + Cyberdrive Awakener (voo permanente) + Alibou (haste ao time).
  - **B. Queima via marcadores** (turno 7–9, consistência média): Lux Artillery dispara 10 dano/oponente ao acumular 30+ marcadores entre artefatos/criaturas — não depende de combate.
  - **C. Conversão explosiva de rocks** (turno 6–8, consistência média-alta): Cyberdrive Awakener transforma o pacote de 10+ mana rocks em exército de 4/4 protegidos, com Alibou dando haste e dano extra.

## Lista de Cartas por Tipo

- **Comandante**: Inspirit, Flagship Vessel — CMC 3, Artefato Lendário Spacecraft, Station (marcadores = poder da criatura virada), gatilho 1+ (combate: +1/+1 ou 2 charges), estática "other artifacts have hexproof and indestructible".
- **Subcomandantes**: nenhum.
- **Planeswalkers**: Saheeli, Sublime Artificer — CMC 3, Servo 1/1 por spell não-criatura; -2 copia artefato.
- **Criaturas** (incl. artefatos-criatura, veículos e spacecraft):
  - Hangarback Walker — CMC X, +1/+1 counters escalam com trigger/proliferate.
  - Coretapper — CMC 2, charge counters em Spacecraft/Lux Cannon; sacrifício = 2 charges.
  - Etherium Sculptor — CMC 2, reduz custo de artefatos em {1}.
  - Third Path Iconoclast — CMC 2, token artefato por spell não-criatura.
  - Enthusiastic Mechanaut — CMC 2, redutor de custo, flying.
  - Brotherhood Vertibird (Veículo) — CMC 3, poder = nº de artefatos, combustível de Station.
  - Foundry Inspector — CMC 3, redutor de custo.
  - Kilo, Apogee Mind — CMC 3, proliferate ao ficar virado; haste.
  - Malcator, Purity Overseer — CMC 3, cria Golems 3/3; recompensa 3 artefatos/turno.
  - Master of Etherium — CMC 3, lord (poder = nº de artefatos).
  - Pinnacle Emissary — CMC 3, drone por artifact spell.
  - Sai, Master Thopterist — CMC 3, thopter por artifact spell; sacrifica 2 → compra carta.
  - Surge Conductor — CMC 3, proliferate por artefato nontoken entrando.
  - Chrome Host Seedshark — CMC 3, incubate = tokens artefato com counters.
  - Uthros Research Craft (Spacecraft) — CMC 3, compra por artifact spell + charges do gatilho 1+.
  - Warmaker Gunship (Spacecraft) — CMC 3, ETB: dano = nº de artefatos; recebe charges.
  - Jhoira, Weatherlight Captain — CMC 4, compra por spell histórica.
  - Padeem, Consul of Innovation — CMC 4, compra no upkeep pelo maior MV de artefato + hexproof ao Inspirit.
  - Leonin Abunas — CMC 4, 2/5, "artifacts you control have hexproof" (cobre o próprio Inspirit).
  - Crystalline Crawler — CMC 4, +1/+1 counters viram mana colorida.
  - Alibou, Ancient Witness — CMC 5, haste ao time; dano + scry por artefatos virados.
  - Deepglow Skate — CMC 5, dobra marcadores.
  - Kappa Cannoneer — CMC 6, cresce por artefato; unblockable; ward 4; improvise.
  - Cyberdrive Awakener — CMC 6, anima mana rocks/Bridges em 4/4 hexproof indestrutíveis.
  - Thought Monitor — CMC 7, affinity; compra 2 + voador.
- **Artefatos**:
  - Everflowing Chalice — CMC X, charge counters crescem com trigger/proliferate.
  - Sol Ring — CMC 1, melhor rock de mana.
  - Arcane Signet — CMC 2, fixação WUR, entra destravado.
  - Sphere of the Suns — CMC 2, gatilho 1+ repõe 2 charges/turno.
  - Talisman of Progress — CMC 2, mv2 destravado (U/W).
  - Chainsaw (Equipamento) — CMC 2, 3 dano no ETB; rev counters equipam atacantes.
  - Empowered Autogenerator — CMC 4, explosivo de charge counters.
  - Glass Casket — CMC 2, exílio de criatura, protegido pelo Inspirit.
  - Pentad Prism — CMC 2, charge counters recarregáveis = mana colorida.
  - Azorius Signet — CMC 2, fixação W/U.
  - Cloud Key — CMC 3, redutor de custo.
  - Perilous Snare — CMC 3, exílio de não-terreno.
  - Reckoner Bankbuster (Veículo) — CMC 2, gatilho 1+ recarrega charges → compra perpétua protegida.
  - Midnight Clock — CMC 3, rock + hour counters (proliferate acelera o "compre 7").
  - Gilded Lotus — CMC 5, {T}: 3 maná de uma cor — 2º ramp explosivo.
  - Lux Artillery — CMC 4, sunburst + 10 dano/oponente aos 30 marcadores.
  - Lux Cannon — CMC 4, "destroy target permanent" recorrente via 2 charges/turno.
  - Thousand Moons Smithy — CMC 4, token */* proporcional; vira terreno-fábrica.
- **Instantâneos**:
  - Dispatch — CMC 1, metalcraft trivial → remoção quase gratuita.
  - Swords to Plowshares — CMC 1, melhor remoção branca.
  - Loran's Escape — CMC 1, indestrutível + scry, funciona no Inspirit.
  - Blacksmith's Skill — CMC 1, indestrutível, funciona no Inspirit.
  - Restoration Magic — CMC 1, indestrutível, funciona no Inspirit; salva de wipes.
  - Unwanted Remake — CMC 1, destrói criatura por {W}.
  - Disruption Protocol — CMC 2, counter duro (sem escape); custo extra = virar 1 artefato.
  - Thirst for Knowledge — CMC 3, compra 3 descartando artefato.
- **Feitiços**:
  - Rip Apart — CMC 2, remoção flexível (criatura/PW ou artefato/encantamento).
  - Stern Lesson — CMC 3, compra 2 + Powerstone.
  - Tezzeret's Gambit — CMC 4, compra 2 + proliferate.
  - Chain Reaction — CMC 4, wipe assimétrico (artefatos-criatura sobrevivem).
  - Fumigate — CMC 5, "destroy all creatures" + ganha vida por morte.
  - Reverse Engineer — CMC 5, improvise → compra 3.
  - Organic Extinction — CMC 10 (improvise), wipe muito assimétrico (só não-artefatos morrem).
- **Terrenos**:
  - 6 Plains, 9 Island, 5 Mountain (básicos).
  - Command Tower, Spire of Industry, Exotic Orchard — fixação tricolor.
  - Battlefield Forge, Rugged Prairie, Skycloud Expanse, Port Town, Glacial Fortress, Clifftop Retreat — fixação de par de cor, majoritariamente destravados.
  - Irrigated Farmland, Mystic Monastery — taplands de fixação/cycling.
  - Temple of Epiphany, Temple of Enlightenment — taplands + scry.
  - Blast Zone — wipe-em-terreno, escala com o pacote de proliferate.
  - Rustvale Bridge, Razortide Bridge, Silverbluff Bridge — terrenos-artefato (protegidos pelo Inspirit, alimentam Spire of Industry/Cyberdrive Awakener).
  - Perilous Landscape — fetch lenta + cycling.

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
1 Chainsaw
1 Empowered Autogenerator
1 Glass Casket
1 Pentad Prism
1 Azorius Signet
1 Cloud Key
1 Perilous Snare
1 Reckoner Bankbuster
1 Midnight Clock
1 Gilded Lotus
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

## Mudanças Aplicadas (auditoria `/improve-deck`)

Dores originais relatadas pelo usuário e como foram endereçadas:

| # | Dor relatada | Resolução |
|---|---|---|
| 1 | Sem respostas / refém das ameaças | Fase interação: ~10 peças pontuais + 3 board wipes (era 2). |
| 2 | Mão morta / cartas que não engrenam sozinhas | Cortes ao longo das 4 fases: Emissary Escort, Voyager Quickwelder, Golem Foundry, Diversion Unit, Phyrexian Revoker, Thopter Fabricator. |
| 3 | Comandante vira alvo imediato | 5 camadas de proteção: Padeem + Leonin Abunas (estáticos, hexproof a todos os artefatos incl. o próprio Inspirit) + Loran's Escape/Blacksmith's Skill/Restoration Magic (one-shots de indestrutível). |
| 4 | Deck lento / sensação de pouco terreno e draw/ramp | Draw 8→13 fontes; ramp 7→10 padrão +0→2 explosivos; terrenos 36→38 (fórmula de Karsten confirmou déficit real de 2 fontes). |
| 5 | Cartas sem sentido / candidatas a corte | Identificadas e substituídas fase a fase (ver tabela de swaps abaixo). |

### Swaps aplicados por fase

**Draw** (versão mesão, 6 swaps):

| Sai | Entra | US$ | Faixa |
|---|---|---|---|
| Voyage Home | Reckoner Bankbuster | 0.61 | mesão (alt. torneio: Bident of Thassa 0.43) |
| Marketback Walker | Thopter Spy Network | 0.38 | torneio |
| Memory Guardian | Whirlwind of Thought | 0.22 | torneio |
| Frogmyr Enforcer | Padeem, Consul of Innovation | 3.24 | mesão (alt. torneio: Reconnaissance Mission 0.35) |
| Ethersworn Sphinx | Reverse Engineer | 0.19 | torneio |
| Solar Array | Midnight Clock | 0.38 | torneio |

**Ramp** (swaps a–d; Chainsaw mantida por decisão do usuário):

| Sai | Entra | US$ | Faixa |
|---|---|---|---|
| Astral Cornucopia | Arcane Signet | 0.36 | torneio |
| Zookeeper Mechan | Sphere of the Suns | 0.20 | torneio |
| Cargo Ship | Talisman of Progress | 0.26 | torneio |
| Captain Storm, Cosmium Raider | Empowered Autogenerator | 0.30 | torneio |

**Interação** (3 swaps):

| Sai | Entra | US$ | Faixa |
|---|---|---|---|
| Diversion Unit | Disruption Protocol | 0.26 | torneio |
| Phyrexian Revoker | Leonin Abunas | 3.06 | mesão |
| Thopter Fabricator | Fumigate | 0.34 | torneio |

**Manabase** (2 cortes obrigatórios + 1 swap opcional):

| Sai | Entra | US$ | Faixa |
|---|---|---|---|
| Boros Garrison | Blast Zone | 0.38 | torneio |
| Temple of Triumph | +1 Island | — | — |
| — | +2 Island (ajuste de básicos) | — | — |
| Voyager Quickwelder | *(corte, financia terreno extra)* | — | — |
| Emissary Escort | *(corte, financia terreno extra)* | — | — |
| Golem Foundry | Gilded Lotus | 0.46 | torneio |

**Custo total dos swaps**: ≈ US$ 3,58 na versão mesão / ≈ US$ 1,02 na versão 100% torneio (excluindo o swap de terreno/ramp da manabase, que soma mais ≈US$ 0,84 em ambas as faixas).

### Pendência aberta (opcional, não aplicada)

**Chief of the Foundry** (CMC 3, US$ 0.20, torneio) — segundo anthem para o Caminho A de vitória, sugerido pelo `wincon-tester` apenas **se o goldfishing mostrar vitórias tardias (turno 8+)**. Não é necessário agora — o deck já bate a meta de 3+ caminhos de vitória. Ver protocolo completo em `07-wincons.md`.

## Protocolo de Goldfishing

Ver `decks/inspirit-flagship-vessel/07-wincons.md` para o protocolo completo (preparação, sequência de turnos, tabela de registro para 5+ partidas, meta de vitória projetada ≤ turno 7, e guia de diagnóstico caso os testes revelem travas de mana, mão morta ou vitória tardia).
