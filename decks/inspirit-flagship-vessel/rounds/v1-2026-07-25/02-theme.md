# Análise Temática — Inspirit, Flagship Vessel (modo improve / auditoria)

## Comandante — análise linha a linha

| Linha/habilidade | Gatilho/termo | O que pede do deck |
|---|---|---|
| `{U}{R}{W}`, Legendary Artifact — Spacecraft, CMC 3 | artefato, barato | Desce cedo; recast barato após remoção (commander tax dói pouco até a 2ª morte). É **artefato mas não criatura** até 8+ — remoção de artefato o atinge desde o turno 3. |
| **Station** (vira outra criatura: charge = poder dela; 8+ = criatura-artefato 5/5) | `station`, `pow>=3` | Criaturas com **poder alto por custo baixo** para estacionar rápido (Master of Etherium, Brotherhood Vertibird, golems 3/3). Tokens 1/1 estacionam mal. |
| **1+**: início do combate → +1/+1 counter **ou** 2 charge counters em outro artefato | `charge counter`, `+1/+1 counter`, `proliferate` | Artefatos que **escalam com marcadores**: Lux Cannon, Everflowing Chalice, Golem Foundry, outros Spacecraft (Uthros, Warmaker), Kappa Cannoneer, Crystalline Crawler. Proliferate multiplica tudo. |
| **8+**: Flying | — | Evasão para o plano "comandante 5/5 voador bate". |
| **"Other artifacts you control have hexproof and indestructible"** | proteção estática | Define o plano: board de artefatos imune a remoção pontual e à maioria dos wipes → **wipes assimétricos são nossos** (Chain Reaction, Organic Extinction). **Ponto cego: o próprio Inspirit fica sem proteção** (dor nº 3) — ele é o alvo óbvio da mesa. Criaturas não-artefato (Sai, Jhoira, Third Path Iconoclast, Malcator, Seedshark, Deepglow Skate) também ficam de fora. |

**Termos do tema**: `t:artifact`, charge counter, +1/+1 counter, proliferate, token de artefato (thopter/golem/servo/construct), affinity/improvise, redução de custo de artefatos, "whenever you cast an artifact spell", poder proporcional a artefatos.

---

## Auditoria carta a carta (63 não-terrenos)

Legenda de sinergia: ✅ 2+ pontos (fica) · ⚠️ marginal (1–2 pontos fracos) · ❌ 0–1 ponto (candidata a corte).

### Criaturas e artefatos-criatura

| Carta | CMC | Categorias | Sinergias | Veredito |
|---|---|---|---|---|
| Hangarback Walker | XX | tema | +1/+1 counters (trigger 1+, Deepglow, proliferate); artefato; morte→thopters (raro, é indestrutível) | ✅ |
| Marketback Walker | XX | tema, draw | +1/+1 counters; **mas o draw exige que ele morra — indestrutibilidade do Inspirit quase impede**; {4}/counter é caro | ⚠️ |
| Captain Storm, Cosmium Raider | 2 | — | Gatilha com artefatos entrando, **mas só põe counter em Pirate — o deck tem ZERO outros Piratas**. Sinergia efetiva nula | ❌ corte |
| Coretapper | 2 | tema | Charge counters em Spacecraft/Lux Cannon/Everflowing/Golem Foundry; sac = 2 charges instantâneos | ✅ |
| Diversion Unit | 2 | counter, tema | Único counterspell do deck; artefato protegido; counter frágil (paga {3}) | ⚠️ manter até substituir |
| Emissary Escort | 2 | tema | Poder alto (X = maior MV entre artefatos) → ótimo combustível de Station; artefato protegido | ✅ |
| Etherium Sculptor | 2 | tema, ramp | Redução de custo p/ ~40 artefatos; artefato protegido | ✅ |
| Third Path Iconoclast | 2 | tema | Tokens artefato a cada spell não-criatura → affinity/improvise/Vertibird/crew; **não é artefato (desprotegida)** | ✅ |
| Enthusiastic Mechanaut | 2 | tema, ramp | Redutor de custo; artefato; flying | ✅ |
| Phyrexian Revoker | 2 | remoção | Hate stick isolado; único ponto é ser artefato. Não avança nenhum plano | ❌ corte |
| Zookeeper Mechan | 2 | ramp | Mana dork só de {R} (cor menos exigida); ativada de 7 mana irreal | ❌ corte |
| Brotherhood Vertibird | 3 | tema | Poder = nº de artefatos → Station gigante; flying; protegido | ✅ |
| Foundry Inspector | 3 | tema, ramp | Redutor de custo; artefato | ✅ |
| Kilo, Apogee Mind | 3 | tema | Proliferate **sempre que fica virado — inclusive ao estacionar o Inspirit**; haste | ✅ forte |
| Malcator, Purity Overseer | 3 | tema | Golems 3/3 (bom poder p/ Station); recompensa 3 artefatos/turno; não-artefato (desprotegido) | ✅ |
| Master of Etherium | 3 | tema, wincon | Poder = nº artefatos (Station enorme); lord de +1/+1 | ✅ |
| Pinnacle Emissary | 3 | tema | Drone por artifact spell; warp barateia; artefato | ✅ |
| Sai, Master Thopterist | 3 | tema, draw | Thopter por artifact spell; sac 2 artefatos → draw; não-artefato | ✅ |
| Surge Conductor | 3 | tema | Proliferate a cada artefato nontoken entrando — motor de counters do deck | ✅ forte |
| Chrome Host Seedshark | 3 | tema | Incubate = tokens artefato com +1/+1 counters; não-artefato | ✅ |
| Uthros Research Craft | 3 | tema, draw | **Draw por artifact spell** + charge; Spacecraft (recebe 2 charges do trigger 1+) | ✅ forte |
| Voyager Quickwelder | 3 | tema, ramp | 4º redutor de {1}; artefato | ✅ |
| Warmaker Gunship | 3 | tema, remoção | ETB: dano = nº artefatos; Spacecraft (charges do trigger) | ✅ |
| Jhoira, Weatherlight Captain | 4 | draw | Draw por spell histórica (~85% do deck) — melhor motor de draw da lista; não-artefato, alvo prioritário | ✅ forte |
| Crystalline Crawler | 4 | tema, ramp | +1/+1 counters (trigger/proliferate recarregam) = mana de qualquer cor | ✅ |
| Alibou, Ancient Witness | 5 | tema, remoção, wincon | Haste para o time; dano + scry escalando com artefatos virados (Station vira!) | ✅ forte |
| Deepglow Skate | 5 | tema | Dobra charge/+1/+1 (estaciona Inspirit quase sozinho; Lux Cannon on-line) | ✅ forte |
| Memory Guardian | 5 | — | 3/4 flying **sem texto relevante** além de affinity. Corpo vanilla | ❌ corte |
| Kappa Cannoneer | 6 | tema, wincon | Cresce com cada artefato entrando; unblockable; ward 4; improvise | ✅ forte |
| Cyberdrive Awakener | 6 | tema, wincon | Anima mana rocks/Bridges em 4/4 **hexproof indestrutíveis** → alfa-strike | ✅ |
| Frogmyr Enforcer | 7 | — | Corpo sem texto (prototype 2/2 ou 4/4); só affinity | ❌ corte |
| Thought Monitor | 7 | tema, draw | Affinity draw 2 + corpo voador; frequentemente custa 1–3 | ✅ |
| Ethersworn Sphinx | 9 | tema | Cascade aleatório; nominal 9 entope a mão cedo (dor nº 2) | ⚠️ corte provável |

### Artefatos não-criatura

| Carta | CMC | Categorias | Sinergias | Veredito |
|---|---|---|---|---|
| Everflowing Chalice | X | ramp, tema | Charge counters → trigger 1+/proliferate/Coretapper fazem-na crescer | ✅ |
| Astral Cornucopia | XXX | ramp, tema | Charge counters, mas taxa péssima (3 mana = 1 mana); só paga com proliferate pesado | ⚠️ corte provável |
| Sol Ring | 1 | ramp | Melhor rock do formato; protegido pelo Inspirit | ✅ |
| Cargo Ship | 2 | ramp | Mana restrita a artefatos, crew, impacto mínimo | ❌ corte |
| Chainsaw | 2 | remoção | Rev counters exigem criaturas morrendo (**as nossas não morrem**); equip {3} caro; 3 dmg ETB é o único valor | ❌ corte |
| Glass Casket | 2 | remoção | Exílio que, protegido pelo Inspirit, dificilmente devolve a criatura | ✅ |
| Pentad Prism | 2 | ramp, tema | Charge counters (recarregáveis via trigger/proliferate/Coretapper) = mana colorida | ✅ |
| Azorius Signet | 2 | ramp | Fixação W/U; protegido | ✅ |
| Cloud Key | 3 | tema, ramp | Redutor de custo; protegido | ✅ |
| Golem Foundry | 3 | tema | Charge via cast + trigger 1+ (2 charges!) + proliferate → golems 3/3 (bons p/ Station) | ✅ |
| Perilous Snare | 3 | remoção, tema | Exílio de qualquer não-terreno; artefato protegido = exílio quase definitivo | ✅ |
| Solar Array | 3 | ramp, tema | 3 mana p/ 1 mana é fraco; sunburst rende pouco (a maioria dos artefatos é incolor/1 pip) | ⚠️ corte provável |
| Thopter Fabricator | 3 | tema | Exige 2º draw no turno — o deck **não tem** motor de draw consistente hoje | ⚠️ |
| Lux Artillery | 4 | wincon, tema | Sunburst nos artefatos-criatura + 10 dmg/oponente com 30 counters (proliferate acelera) | ✅ |
| Lux Cannon | 4 | remoção, tema | Trigger 1+ dá 2 charges/turno + Deepglow/proliferate → destrói permanente recorrente | ✅ |
| Thousand Moons Smithy | 4 | tema, wincon | Token */* proporcional; transforma em terreno que fabrica gnomos; conta artefatos | ✅ |

### Instantâneos e feitiços

| Carta | CMC | Categorias | Sinergias | Veredito |
|---|---|---|---|---|
| Dispatch | 1 | remoção | Metalcraft trivial → StP nº 2 | ✅ |
| Swords to Plowshares | 1 | remoção | Melhor remoção branca | ✅ |
| Loran's Escape | 1 | proteção | **Funciona no Inspirit** ("target artifact or creature") + scry | ✅ |
| Blacksmith's Skill | 1 | proteção | **Funciona no Inspirit** ("target permanent") | ✅ |
| Restoration Magic | 1 | proteção | **Funciona no Inspirit**; modo Curaga salva o time de wipe de exílio/-X/-X | ✅ |
| Unwanted Remake | 1 | remoção | Destroi criatura por {W} a instant; manifest dread é custo aceitável | ✅ |
| Rip Apart | 2 | remoção | Flexível (criatura/PW ou artefato/encantamento); sorcery | ✅ |
| Stern Lesson | 3 | draw, ramp | Draw 2 + Powerstone (artefato p/ affinity/improvise/contagens) | ✅ |
| Thirst for Knowledge | 3 | draw | Draw 3 descartando artefato (barato num deck com 40+) | ✅ |
| Tezzeret's Gambit | 4 | draw, tema | Draw 2 + proliferate (charge + +1/+1 em massa) | ✅ |
| Chain Reaction | 4 | wipe | **Assimétrico**: criaturas-artefato ficam (indestrutíveis c/ Inspirit); mata Sai/Jhoira/etc. junto | ✅ |
| Voyage Home | 7 | draw | Draw 3 + 3 vidas; affinity barateia, mas nominal 7 entope a mão cedo | ⚠️ |
| Organic Extinction | 10 | wipe | **Muito assimétrico** (só não-artefatos morrem) + improvise | ✅ |

### Terrenos (36)

| Grupo | Cartas | Nota |
|---|---|---|
| Básicos (17) | 6 Island, 6 Plains, 5 Mountain | OK |
| Sempre tapped (10) | Boros Garrison, Mystic Monastery, Irrigated Farmland, Rustvale/Razortide/Silverbluff Bridge, Temple of Epiphany/Enlightenment/Triumph | **28% da manabase entra virada** — agrava a lentidão (dor nº 4). Bridges são artefato-terreno (affinity/metalcraft, hexproof c/ Inspirit): valem o tapped. Boros Garrison e 3 Temples são os piores. |
| Condicionais (3) | Port Town, Glacial Fortress, Clifftop Retreat | Boas no mid, medianas no T1 |
| Untapped fixação (6) | Command Tower, Spire of Industry, Exotic Orchard, Battlefield Forge, Rugged Prairie, Skycloud Expanse | Boas (Spire é excelente aqui) |
| Utilidade (1) | Perilous Landscape | Fetch lenta + cycling; OK |

---

## Diagnóstico

### 1. Contagem por categoria vs. metas do pipeline

| Categoria | Meta | Atual | Situação |
|---|---|---|---|
| Draw (fontes reais) | 12–13 | **8** (Jhoira, Uthros, Sai, Thought Monitor, Thirst, Stern Lesson, Tezzeret's Gambit, Voyage Home) | **Déficit de 4–5** — confirma dor nº 4 |
| Ramp padrão | 10–11 | **~7 confiáveis** (Sol Ring, Signet, Pentad Prism, Everflowing, Crystalline Crawler, Stern Lesson, + 5 redutores de custo como pseudo-ramp) — Astral Cornucopia/Solar Array/Cargo Ship/Zookeeper são de baixa qualidade | Déficit de qualidade; nº nominal ok, taxa ruim |
| Ramp explosivo | 2–3 | **0** | Lacuna |
| Interação pontual | ~10 | **8** (StP, Dispatch, Unwanted Remake, Rip Apart, Glass Casket, Perilous Snare, Lux Cannon, Diversion Unit) | Déficit de ~2; **counter: só 1 e frágil**; confirma dor nº 1 |
| Wipes | 2–4 | **2** (Chain Reaction, Organic Extinction — ambos assimétricos, ótimos) | No mínimo; cabe +1 barato |
| Proteção do comandante | — | **3** one-shots (Loran's Escape, Blacksmith's Skill, Restoration Magic — todas funcionam no Inspirit) | **Zero proteção permanente** (equipamento/estático). Dor nº 3 parcialmente coberta |
| Wincons | 3+ | **4** (alfa-strike Cyberdrive/Alibou; Kappa Cannoneer; Lux Artillery; Inspirit 8+ voador) | OK, mas todas dependem de board |
| Terrenos | 38 base | **36** | 1–2 abaixo, com 28% tapped e ramp fraco → dor nº 4 |

### 2. Candidatas a corte (da mais fraca à menos fraca)

1. **Captain Storm, Cosmium Raider** — só bota counter em Pirate; o deck não tem outro Pirate. Sinergia efetiva zero.
2. **Memory Guardian** — corpo vanilla com affinity; não avança nenhum plano.
3. **Frogmyr Enforcer** — idem: corpo sem texto, só affinity.
4. **Chainsaw** — depende de criaturas morrendo (as nossas são indestrutíveis); equip caro.
5. **Zookeeper Mechan** — dork só de {R} (cor menos exigida) com ativada de 7 mana.
6. **Cargo Ship** — mana restrita, crew, quase nenhum impacto.
7. **Phyrexian Revoker** — hate isolado, 1 ponto de sinergia (ser artefato).
8. **Astral Cornucopia** — taxa 3:1; só funciona com proliferate já estabelecido.
9. **Marketback Walker** — o draw exige a própria morte, que o Inspirit impede; {4}/counter.
10. **Solar Array** — 3 mana → 1 mana; sunburst rende pouco em artefatos incolores.
11. **Ethersworn Sphinx** — nominal 9 parado na mão (dor nº 2); cascade aleatório.
12. **Voyage Home** — nominal 7; win-more (só barata quando você já está ganhando).
13. **Thopter Fabricator** — precisa de motor de draw que o deck ainda não tem (reavaliar após fase de draw).
14. **Diversion Unit** — counter que paga {3} anula; trocar por counter real na fase de interação.
15. **Boros Garrison + 2–3 Temples** (manabase) — taplands sem sinergia; trocar por untapped/básicos.

### 3. Curva de mana (dor nº 2 — "mão cheia, não baixa nada")

- **CMC médio nominal ≈ 3,05** (real ~3,2+ contando X-spells) — alto para 36 terrenos e ramp de baixa qualidade.
- Distribuição: mv1 = 7 · mv2 = 15 · **mv3 = 20 (engarrafamento)** · mv4 = 7 · mv5+ = **10 nominal** (Alibou, Deepglow, Memory Guardian, Kappa, Cyberdrive, Frogmyr, Thought Monitor, Voyage Home, Ethersworn Sphinx, Organic Extinction).
- As cartas de affinity/improvise **só ficam baratas com board estabelecido**: se o Inspirit morre cedo ou o board não engrena, a mão trava — exatamente a dor relatada. Com 8 fontes de draw e 36 terrenos, a probabilidade de "3 terrenos e parar" é alta.
- Correção: cortar 3–4 dos top-ends fracos (Ethersworn Sphinx, Frogmyr, Memory Guardian, Voyage Home), subir para 37–38 terrenos e trocar ramp ruim por rocks de 2 de qualidade.

### 4. Manabase

- **Quantidade**: 36 < 38 da base; com curva 3,05 e ramp mediano, recomendo 37–38.
- **Taplands**: 10 sempre viradas (28%) + 3 condicionais — deck perde ~1 turno de desenvolvimento por jogo. Piores: Boros Garrison, 3 Temples. Bridges ficam (artefato + indestrutível nativo + affinity).
- **Cores**: pips aproximados U 24 > W 20 > R 13; fontes W ≈ 17, U ≈ 15, R ≈ 14 — **balance razoável**, R levemente sobre-representada nos básicos (5 Mountain) vs. demanda. O problema é velocidade, não fixação.

### 5. Lacuna de proteção do comandante (dor nº 3)

- O estático protege **"other artifacts"** — Inspirit é o único artefato nu. Removê-lo derruba a proteção de todo o board: alvo prioritário óbvio.
- Protegem-no hoje: **Loran's Escape** (artifact or creature ✓), **Blacksmith's Skill** (permanent ✓), **Restoration Magic** (permanent ✓) — 3 instantes one-shot de {W}. Funcionam, mas exigem mana aberta e acabam.
- **Falta proteção permanente**: nenhum equipamento/encantamento/estático que dê hexproof, indestrutível ou ward ao Inspirit (nota: equipamentos anexados a ele são "other artifacts" → o próprio equipamento fica hexproof/indestrutível — sinergia excelente para a fase de interação explorar; shroud deve ser evitado só se quisermos mirar o Inspirit com pumps — o trigger 1+ dele mira apenas *outros* artefatos, então shroud nele não quebra o trigger).
- Criaturas não-artefato chave (Jhoira, Sai, Surge Conductor é artefato ok, Third Path Iconoclast, Malcator) também ficam expostas — priorizar versões-artefato ao substituir.
