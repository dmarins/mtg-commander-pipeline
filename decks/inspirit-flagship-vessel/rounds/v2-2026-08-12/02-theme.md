# Análise Temática — v2 (segunda auditoria, cética, do zero)

> Esta auditoria reavalia as 99 cartas atuais (pós-v1) sem assumir que nenhuma escolha anterior está correta. Onde a conclusão diverge da v1, isso é sinalizado explicitamente com o motivo.

## Comandante — análise linha a linha

| Linha/habilidade | Gatilho/termo | O que exige do deck |
|---|---|---|
| `{U}{R}{W}`, Legendary Artifact — Spacecraft, CMC 3 | artefato barato | É **artefato, não criatura**, até virar 8+. Remoção de artefato o atinge desde o T3; recast é barato mesmo após a 1ª morte. |
| **Station** (tap outra criatura sua → charge counters = poder dela nele; só sorcery-speed; 8+ = criatura-artefato 5/5) | `station`, poder alto por custo baixo | Quer criaturas com poder desproporcional ao custo (Master of Etherium, Brotherhood Vertibird, Kilo). Bônus real: criaturas com **gatilho de "fica tapada"** (Kilo) disparam ao serem usadas para Station — combo gratuito. |
| **1+**: início do combate → +1/+1 counter **ou** 2 charge counters em outro artefato alvo | `charge counter`, `+1/+1 counter`, `proliferate` | Alimenta diretamente Lux Cannon, Lux Artillery, Everflowing Chalice, Sphere of the Suns, Empowered Autogenerator, Reckoner Bankbuster, Pentad Prism, Crystalline Crawler — qualquer artefato com contadores vira um "banco" que o comandante reabastece sozinho todo turno. |
| **8+**: Flying | evasão | Plano B: "comandante 5/5 voador" como wincon independente. |
| **Estático: "Other artifacts you control have hexproof and indestructible"** | proteção assimétrica | Define o plano central: board de artefatos imune a remoção pontual e à maioria dos wipes → **wipes simétricos viram assimétricos a nosso favor** (qualquer "destroy all creatures" mata só os não-artefatos e os oponentes). **Ponto cego confirmado**: o próprio Inspirit fica de fora ("*other* artifacts") — é o alvo prioritário óbvio da mesa, e criaturas não-artefato do deck (Jhoira, Sai, Malcator, Third Path Iconoclast, Saheeli-planeswalker) também ficam expostas. |

**Termos do tema**: `t:artifact`, charge counter, +1/+1 counter, proliferate, "whenever you cast an artifact spell", "whenever you cast a noncreature spell" (cluster de gatilho), token de artefato (thopter/servo/golem/soldier), affinity/improvise, redução de custo, hexproof/indestructible.

**Achado novo desta auditoria**: quatro cartas do deck compartilham o **mesmo gatilho** — "whenever you cast a noncreature spell" — Third Path Iconoclast, Chrome Host Seedshark, Saheeli (Sublime Artificer) e Whirlwind of Thought. Isso não é redundância ruim: é uma sinergia sobreposta real, porque cada spell não-criatura (a maioria do deck, dado o volume de artefatos/instants/sorceries) dispara 4 efeitos simultâneos (token + token com counters + token + draw). Vale citar isso ao usuário como "motor" do deck.

**Achado novo #2 (correção da v1)**: **Fumigate** ("destroy all creatures") é, como Chain Reaction e Organic Extinction, **assimétrico via estático do comandante** — nossas criaturas-artefato (a maioria do board) sobrevivem por indestructible; só morrem as não-artefato (nossas e dos oponentes) e o board rival inteiro. A v1 não tinha esse wipe na lista; agora o deck tem **3 wipes coerentes com o plano**, todos vantajosos para nós.

**Achado novo #3 (correção da v1)**: **Chainsaw** — a v1 cortou esta carta assumindo que "nossas criaturas não morrem" tornava o rev counter inútil. Releitura cuidadosa do oracle: *"Whenever **one or more creatures die**"* não tem restrição de controle — conta mortes de qualquer jogador. Em mesa de 4, isso dispara com frequência (combate, remoção alheia, e principalmente **nossos próprios 3 wipes**, que geram uma rajada de rev counters no mesmo turno). Ainda assim, `Equip {3}` é caro e a carta continua marginal — não é lixo, mas também não é destaque. Reclassificada de "corte" para "mantém, com ressalva".

---

## Auditoria carta a carta (99 não-comandante)

Legenda: ✅ 2+ pontos de sinergia (mantém) · ⚠️ marginal/redundante (candidata a revisão) · ❌ sem sinergia sobreposta (corte).

### Criaturas e artefatos-criatura (25)

| Carta | CMC | Categorias | Sinergias (2+) | Veredito |
|---|---|---|---|---|
| Hangarback Walker | XX | tema | +1/+1 counters (trigger 1+, proliferate); artefato protegido; bom combustível de Station | ✅ |
| Coretapper | 2 | tema | 1 charge counter/turno (ou 2 no sac) em qualquer artefato-alvo — reabastece Lux Cannon/Sphere/Autogenerator/Reckoner Bankbuster | ✅ |
| Etherium Sculptor | 2 | tema, ramp | Redutor de custo (~40 artefatos); artefato protegido | ✅ |
| Third Path Iconoclast | 2 | tema | Cluster "noncreature spell" (ver acima); gera Soldiers-artefato para Station/affinity; não-artefato (exposta) | ✅ |
| Enthusiastic Mechanaut | 2 | tema, ramp | Redutor de custo; artefato; corpo voador | ✅ |
| Brotherhood Vertibird | 3 | tema | Poder = nº artefatos → Station gigante; flying; protegido | ✅ |
| Foundry Inspector | 3 | tema, ramp | Redutor de custo; artefato protegido | ✅ |
| Kilo, Apogee Mind | 3 | tema | **Proliferate sempre que fica tapado** — inclui ser usado para Station (combo gratuito); haste; artefato | ✅ forte |
| Malcator, Purity Overseer | 3 | tema | Golems 3/3 (bom poder p/ Station); recompensa 3 artefatos/turno; não-artefato | ✅ |
| Master of Etherium | 3 | tema, wincon | Poder = nº artefatos (Station enorme); lord +1/+1 para todos artefatos | ✅ |
| Pinnacle Emissary | 3 | tema | Drone por artifact spell; warp barateia; artefato | ✅ |
| Sai, Master Thopterist | 3 | tema, draw | Thopter por artifact spell; sac 2 artefatos → draw (fodder farto com todos os token-makers); não-artefato | ✅ |
| Surge Conductor | 3 | tema | Proliferate a cada artefato nontoken entrando — motor central do pacote de contadores | ✅ forte |
| Chrome Host Seedshark | 3 | tema | Cluster "noncreature spell"; incubate = tokens com +1/+1 counters (proliferate escala); flying; não-artefato | ✅ |
| Uthros Research Craft | 3 | tema, draw | Draw por artifact spell + Spacecraft (recebe 2 charges do trigger 1+) | ✅ forte |
| Warmaker Gunship | 3 | tema, remoção | ETB dano = nº artefatos; Spacecraft (charges do trigger 1+) | ✅ |
| Jhoira, Weatherlight Captain | 4 | draw | Draw por spell histórica (~85% do deck); não-artefato, alvo prioritário sem proteção própria | ✅ forte |
| Padeem, Consul of Innovation | 4 | proteção, draw | **"Artifacts you control have hexproof" inclui o próprio Inspirit** — fecha o ponto cego do estático; draw condicional (maior CMV) | ✅✅ — resolve dor real |
| Leonin Abunas | 4 | proteção | Mesmo efeito de Padeem (hexproof em artefatos, incluindo Inspirit); corpo 2/5 vanilla, sem draw/token | ⚠️ redundante com Padeem — ver diagnóstico |
| Crystalline Crawler | 4 | tema, ramp | +1/+1 counters viram mana de qualquer cor; recarrega via trigger/proliferate | ✅ |
| Alibou, Ancient Witness | 5 | tema, remoção, wincon | Haste time inteiro; dano+scry escala com artefatos tapados — **inclui Station** tapando outra criatura | ✅ forte |
| Deepglow Skate | 5 | tema | Dobra qualquer contador (charge/+1+1) — sinergiza com quase todo o pacote de contadores | ✅ forte |
| Kappa Cannoneer | 6 | tema, wincon | Cresce com artefato entrando; unblockable; ward 4; improvise | ✅ forte |
| Cyberdrive Awakener | 6 | tema, wincon | Anima mana rocks/Bridges em 4/4 hexproof+indestructible (dobrado com o estático) → alfa-strike | ✅ |
| Thought Monitor | 7 | tema, draw | Affinity — custa 1–3 na prática; draw 2 + corpo voador | ✅ |

### Artefatos não-criatura (18)

| Carta | CMC | Categorias | Sinergias | Veredito |
|---|---|---|---|---|
| Everflowing Chalice | X | ramp, tema | Charge counters crescem com trigger/proliferate/Coretapper | ✅ |
| Sol Ring | 1 | ramp | Melhor rock do formato; protegido | ✅ |
| Arcane Signet | 2 | ramp | Fixa qualquer cor da identidade; artefato protegido barato | ✅ |
| Sphere of the Suns | 2 | ramp, tema | **Entra com 3 charge counters** — trigger 1+/Coretapper/proliferate a mantêm viva além do normal | ✅ |
| Talisman of Progress | 2 | ramp | Fixa W/U; protegido — mas **repete a mesma fixação de Azorius Signet** | ⚠️ redundante, ver diagnóstico |
| Chainsaw | 2 | remoção | ETB 3 dano; rev counters (ver achado #3) crescem com QUALQUER morte, inclusive as geradas pelos nossos 3 wipes | ✅ (corrigido da v1) |
| Empowered Autogenerator | 4 | ramp, tema | Charge counters escaláveis — 1+/proliferate/Coretapper/Deepglow podem turboalimentá-lo para mana explosiva | ✅ |
| Glass Casket | 2 | remoção, tema | Exílio ≤CMC3; artefato protegido = quase irremovível | ✅ |
| Pentad Prism | 2 | ramp, tema | Charge counters recarregáveis (trigger/proliferate/Coretapper) = mana colorida | ✅ |
| Azorius Signet | 2 | ramp | Fixa W/U; protegido — mesma redundância de Talisman of Progress | ⚠️ |
| Cloud Key | 3 | tema, ramp | Reduz custo por tipo (escolhendo "artifact", reduz ~40 cartas); protegido | ✅ forte |
| Perilous Snare | 3 | remoção, tema | Exílio de não-terreno; protegido = quase permanente | ✅ |
| Reckoner Bankbuster | 2 | draw, tema | 3 charge counters (trigger 1+ reabastece 2 de uma vez); draw repetido; termina em Treasure+Pilot (crew) | ✅✅ excelente overlap |
| Midnight Clock | 3 | ramp, draw | Rock de U; hour counters são "contadores" — **proliferate (Surge Conductor/Kilo) acelera o wheel de 12 cartas** | ✅ |
| Gilded Lotus | 5 | ramp | 3 mana de 1 cor — sem contadores, sem corpo, só ramp puro num deck que já tem ~10 rocks | ⚠️ candidata a corte — ver diagnóstico |
| Lux Artillery | 4 | wincon, tema | Sunburst + 10 dano/oponente aos 30 counters; proliferate acelera | ✅ forte |
| Lux Cannon | 4 | remoção, tema | 2 charges/turno via trigger 1+ + Deepglow/proliferate → remoção recorrente de qualquer permanente | ✅ forte |
| Thousand Moons Smithy | 4 | tema, wincon | Token proporcional a artefatos; vira terreno-fábrica de gnomos | ✅ |

### Encantamentos e planeswalker (3)

| Carta | CMC | Categorias | Sinergias | Veredito |
|---|---|---|---|---|
| Thopter Spy Network | 4 | draw, tema | Thopter grátis por upkeep (quase sempre há 1 artefato); draw quando artefato-criatura bate | ✅✅ |
| Whirlwind of Thought | 4 | draw, tema | Cluster "noncreature spell" — draw por quase toda a mão do deck | ✅✅ |
| Saheeli, Sublime Artificer | 3 | tema | Cluster "noncreature spell" (Servo); -2 copia artefato (pode clonar Sol Ring ou um lord) | ✅ |

### Instantâneos e feitiços (15)

| Carta | CMC | Categorias | Sinergias | Veredito |
|---|---|---|---|---|
| Dispatch | 1 | remoção | Metalcraft trivial (40+ artefatos) → StP nº2 | ✅ |
| Swords to Plowshares | 1 | remoção | Melhor remoção do jogo; 0 sinergia temática direta, mas eficiência universal justifica manter fora da régua de "2 pontos" (ver nota) | ✅* |
| Loran's Escape | 1 | proteção | Funciona no próprio Inspirit ("artifact or creature") + scry | ✅ |
| Blacksmith's Skill | 1 | proteção | Funciona no Inspirit ("target permanent") | ✅ |
| Restoration Magic | 1 | proteção | Funciona no Inspirit; modo Curaga recompra de wipe de exílio/-X/-X | ✅ |
| Unwanted Remake | 1 | remoção | Destroi criatura por {W} instant; manifest dread é custo aceitável | ✅ |
| Disruption Protocol | 2 | counter | Único counterspell do deck; custo extra (tap artefato ou {1}) é trivial com 40+ artefatos | ✅ mas isolado — ver lacuna |
| Thirst for Knowledge | 3 | draw | Draw 3 descartando artefato/terreno (deck tem 40+ artefatos para alimentar) | ✅ |
| Rip Apart | 2 | remoção | Flexível (criatura/PW ou artefato/encantamento) | ✅ |
| Stern Lesson | 3 | draw, ramp, tema | Draw 2 + Powerstone (mana só para spells de artefato — reforça o plano tribal) | ✅ forte |
| Tezzeret's Gambit | 4 | draw, tema | Draw 2 + proliferate (charge + +1/+1 em massa) | ✅ forte |
| Chain Reaction | 4 | wipe | Assimétrico — criaturas-artefato sobrevivem (indestructible via Inspirit) | ✅ |
| Fumigate | 5 | wipe | Assimétrico (ver achado novo acima) + ganho de vida | ✅ (recontextualizado) |
| Reverse Engineer | 5 | draw, tema | Improvise — os 40+ artefatos pagam boa parte do custo | ✅ |
| Organic Extinction | 10 | wipe, tema | Muito assimétrico (só não-artefatos morrem) + improvise | ✅ |

*Nota sobre Swords to Plowshares: é o único caso do pool em que aplico exceção à régra de "2+ pontos de sinergia" — remoção 1-mana incondicional é eficiência estrutural que qualquer deck branco deveria rodar, independente de tema. Sinalizado para transparência, não é um problema.

### Terrenos (38 — auditoria de composição, detalhe fica para a fase de manabase)

| Grupo | Cartas | Observação desta auditoria |
|---|---|---|
| Básicos (20) | 6 Plains, 9 Island, 5 Mountain | R segue a menos representada nos básicos; ver nota de cor abaixo |
| Sempre tapped (10) | Mystic Monastery, Irrigated Farmland, Rustvale/Razortide/Silverbluff Bridge, Temple of Epiphany, Temple of Enlightenment, + outros | Bridges são artefato-terreno protegido pelo estático — ficam. Temples/Farmland/Monastery continuam sendo o maior custo de velocidade do deck |
| Condicionais (3) | Port Town, Glacial Fortress, Clifftop Retreat | Boas do meio-jogo em diante |
| Untapped fixação (6) | Command Tower, Spire of Industry, Exotic Orchard, Battlefield Forge, Rugged Prairie, Skycloud Expanse | Boas — mantidas |
| Utilidade (2) | Blast Zone, Perilous Landscape | Blast Zone é remoção-terreno; Perilous Landscape sacrifica por fetch condicional |

---

## Diagnóstico por categoria vs. metas do pipeline

| Categoria | Meta | Atual (v2) | Situação |
|---|---|---|---|
| Draw (fontes reais) | 12–13 | **11 sólidas** (Jhoira, Uthros Research Craft, Sai, Thought Monitor, Thirst for Knowledge, Stern Lesson, Tezzeret's Gambit, Reckoner Bankbuster, Thopter Spy Network, Whirlwind of Thought, Reverse Engineer) + 2 condicionais (Padeem, Midnight Clock) | **Meta batida** — grande evolução vs. v1 (que tinha 8). Boa notícia a confirmar em goldfish. |
| Ramp padrão | 10–11 | **11 rocks** (Sol Ring, Arcane Signet, Sphere of the Suns, Talisman of Progress, Azorius Signet, Pentad Prism, Everflowing Chalice, Crystalline Crawler, Empowered Autogenerator, Midnight Clock, Gilded Lotus) + 4 redutores de custo (pseudo-ramp) | **Excesso**: 11 rocks dedicados + 4 redutores é mais do que a meta pede. Overlap de fixação: Azorius Signet e Talisman of Progress fixam exatamente as mesmas 2 cores (W/U) — redundante. Gilded Lotus (5 mana, sem contador, sem corpo) é o rock mais fraco do lote. |
| Ramp explosivo | 2–3 | **~2** (Empowered Autogenerator e Sphere of the Suns, se turbinados por 1+/Coretapper/Deepglow) | Na faixa baixa da meta, mas real — não é ramp nominal como na v1 (Astral Cornucopia cortada corretamente). |
| Interação pontual (remoção) | ~10 | **9** (Dispatch, StP, Unwanted Remake, Rip Apart, Glass Casket, Perilous Snare, Lux Cannon, Chainsaw, Warmaker Gunship-ETB) | Praticamente na meta. |
| Counterspells | — | **1** (Disruption Protocol) | **Lacuna real**: um único counter em 99 cartas é frágil contra combo/bomba alheia; considerar 1–2 adicionais na próxima fase de interação. |
| Proteção do comandante | — | **5** (Loran's Escape, Blacksmith's Skill, Restoration Magic — one-shot; Padeem, Leonin Abunas — permanente hexproof) | Melhora real vs. v1 (que tinha só 3 one-shots): agora há proteção **permanente** que cobre o próprio Inspirit. Mas **nenhuma fonte dá indestructible a ele** (só hexproof) — ele ainda morre para "destroy"/"-X/-X" não-direcionado. Considerar 1 equipamento tipo indestructible-granting na fase de interação. |
| Wipes | 2–4 | **3** (Chain Reaction, Organic Extinction, Fumigate — todos assimétricos a nosso favor) | Na meta, com boa qualidade. |
| Wincons | 3+ | **7 caminhos** (Master of Etherium, Kappa Cannoneer, Cyberdrive Awakener, Alibou, Lux Artillery, Thousand Moons Smithy, Inspirit 8+ voador) | Bem acima da meta — diversidade saudável, mas todos dependem de board estabelecido (nenhum wincon "solo"/combo alternativo). |
| Terrenos | 38 | **38** | Na meta. |

## Candidatas a corte/revisão (da mais clara à mais discutível)

1. **Gilded Lotus** — 5 mana por um rock sem contadores, sem corpo, sem overlap com o pacote de contadores; o deck já tem 10 outras fontes de ramp mais baratas ou mais sinérgicas. Sobra na curva alta exatamente onde o deck quer estar jogando ameaças, não rocks.
2. **Leonin Abunas** — duplica exatamente o efeito de Padeem (hexproof a artefatos, incluindo o próprio Inspirit) sem agregar draw, token ou qualquer segunda função; se for para ter 2 fontes de proteção permanente, melhor usar a segunda vaga para algo que dê **indestructible** ao Inspirit (que ainda falta) em vez de hexproof redundante.
3. **Talisman of Progress vs. Azorius Signet** — ambos fixam apenas W/U; um dos dois é redundante e não ajuda a escassez relativa de R. Cortar um libera espaço para um counterspell adicional ou um fixador que inclua R.
4. **Disruption Protocol isolado** — não é candidata a corte, mas o diagnóstico expõe que "counter" é praticamente zero-profundidade (1 carta); se algum dos cortes acima abrir espaço, a prioridade é aqui.
5. **Chainsaw** — reavaliada e mantida (ver achado #3), mas é a mais fraca do pacote de remoção; se a lacuna de counter/proteção exigir espaço, é a primeira candidata a sair antes de qualquer core do tema.

Nenhuma outra carta do pool (creio ter revisado as 99) ficou abaixo de 2 pontos de sinergia sobreposta — a v1 já havia feito boa faxina (Captain Storm, Memory Guardian, Frogmyr Enforcer, Zookeeper Mechan, Cargo Ship, Phyrexian Revoker, Astral Cornucopia, Marketback Walker, Solar Array, Ethersworn Sphinx, Voyage Home, Thopter Fabricator, Diversion Unit, Boros Garrison e Temple de sobra seguem fora, corretamente).

## Curva de mana e cores (observação, aprofundar na fase de manabase)

- Curva geral parece saudável: pico em CMC 2–3 (a maioria dos habilitadores de tema), com CMC4–6 reservado a payoffs/wincons — bem mais equilibrada que a v1 (que tinha 20 cartas em CMC3 sozinho).
- **11 rocks de ramp + 4 redutores de custo é provavelmente redundante** para 99 cartas; o excesso de mana-fixing pode estar "roubando" espaço de interação (só 1 counter, wipes na meta mas justos).
- Cor R segue a mais escassa nos básicos (5 Mountain vs. 6 Plains/9 Island) e nenhum rock novo (Sphere of the Suns, Arcane Signet, Empowered Autogenerator, Midnight Clock, Gilded Lotus, exceto Arcane Signet/Sphere/Autogenerator que são pentacolor-ish) corrige especificamente isso — Talisman/Azorius Signet (ambos W/U) pioram a assimetria ao invés de ajudar R. Vale revisar na fase de manabase se a demanda de pips R (Chainsaw, Rip Apart, Alibou, Kilo, Saheeli, Warmaker Gunship) justifica subir Mountain ou trocar 1 rock W/U por um fixador que inclua R.
