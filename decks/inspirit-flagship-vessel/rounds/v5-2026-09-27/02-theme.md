# Análise Temática — v5 · troca pontual (`/swap-card Traxos, Academy Guardian`)

> Modo `improve`, escopo **troca pontual**: a entrada é fixa (`Traxos, Academy Guardian`) e a tarefa é
> achar **uma** saída (no máximo duas, ranqueadas). Lista de referência: `deck.md` da raiz (montado,
> confirmado em 2026-09-18; já inclui Vraska, Soul of Stone e Requisition Raid).
> Oracle de **todas** as 64 cartas não-terreno puxado nesta sessão com `bin/mtgdb oracle` (2026-09-27);
> `mtgdb deck inspirit-flagship-vessel` resolveu as 84 entradas sem faltantes.

## Comandante — análise linha a linha

| Linha/habilidade | Gatilho/termo | O que habilita |
|---|---|---|
| `Station (Tap another creature you control: Put charge counters equal to its power…)` | **tap de criatura**, **poder** | todo contador do comandante vem de corpo virado → a moeda do deck é **poder disponível em fase principal** (a "economia de poder" achada na v3) |
| `1+ \| At the beginning of combat… +1/+1 counter or two charge counters on up to one other target artifact` | **+1/+1 / charge**, **outro artefato** | alimenta naves, rocks de charge e criaturas-artefato; desligado com 0 contadores e nunca se alimenta (ruling Adagia) |
| `8+ \| Flying` | voar | o comandante vira 5/5 voador |
| `Other artifacts you control have hexproof and indestructible` | **artefato** | todo artefato vira permanente protegido; wipes de destruição do próprio deck (Chain Reaction, Organic Extinction) ficam assimétricos |

## Termos de busca do tema

`artifact` (contagens, affinity, improvise, metalcraft) · `artifact creature` (anthems de Chief/Master) ·
`power` / `tap another creature` (Station, crew) · `+1/+1 counter` / `charge counter` / `proliferate` ·
`noncreature spell` (Vraska, Saheeli, Whirlwind, Seedshark — e agora prowess) · `artifact spell`
(Pinnacle, Sai, Golem Foundry, Uthros, Jhoira, redutores).

Sem busca de cartas novas nesta rodada — a entrada é fixa.

## Radar do EDHREC

**Não consultado nesta rodada**, por instrução do orquestrador: troca pontual com entrada fixa, sem
busca de candidatas. As fases seguintes, se forem invocadas, não têm seção de radar para ler aqui.

---

## Entrada fixa — ficha F1–F7 de `Traxos, Academy Guardian`

`{3}{U}` · 1/5 · Legendary Artifact Creature — Dragon Construct · CMC 4 · identidade U · **R$ 1,89**
(LigaMagic menor, 2026-09-27) · **não está nas sobressalentes** → compra. Sem rulings oficiais.
Nunca esteve neste deck (`decisions.md` e rodadas sem ocorrência) — sem histórico a declarar.

| # | Eixo | Conteúdo (conferido contra o oracle) |
|---|---|---|
| F1 | Texto | (a) `costs {2} less if you've cast a noncreature spell this turn` — **só genérico**: `{3}{U}` → `{1}{U}`; com **um** redutor de artefato (Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector, Voyager Quickwelder, Cloud Key em "artifact") → **`{U}`**, que é o piso: o 2º redutor em diante não rende nada se o desconto próprio ligou. Sem mágica não-criatura no turno: `{2}{U}` com 1 redutor, `{U}` com 3. **A alegação da triagem está correta.** (b) Flying. (c) Vigilance. (d) Prowess: +1/+1 até o fim do turno por mágica não-criatura **lançada com a Traxos já em campo** — a mágica que liga o desconto dela **não** a bombeia. |
| F2 | Corpo | 1/5 voador. Tapável para **Station** (inclusive no turno em que entra — o `{T}` é custo da nave) e **crew** (Vertibird crew 2, Recon Craft crew 2, Bladewheel Chariot crew 1). Não usa o próprio `{T}`. Bloqueador voador de resistência 5 (7 com anthems), indestrutível. **Poder na hora de estacionar:** 1 sem nada · 2 com um anthem · 3 com Chief + Master · +1 por mágica não-criatura lançada antes, no mesmo turno, depois de ela entrar. Faixa realista a partir do turno seguinte ao da entrada: **2–4**. |
| F3 | Tipo | **Artefato** (42 → 43): Master of Etherium (P/T), Brotherhood Vertibird (poder), Kappa Cannoneer (+1/+1 e imbloqueável ao entrar), Malcator (conta para "3 artefatos entraram"), affinity de Thought Monitor / Ethersworn Sphinx / Voyage Home, metalcraft do Dispatch, improvise de Kappa / Organic Extinction, Gnome do Thousand Moons Smithy. **Artifact spell**: Pinnacle Emissary (Drone), Sai (Thopter), Golem Foundry (charge), Uthros 3+ (compra + charge), Jhoira (compra — histórica também por ser lendária), e os 5 redutores. **Artefato nontoken que entra**: Surge Conductor (proliferate). **Criatura nontoken que entra**: Metastatic Evangel (proliferate). Criaturas 27 → 28. **É creature spell** → **não** dispara Vraska, Saheeli, Whirlwind nem Seedshark. |
| F4 | Receptor | hexproof + indestrutível do comandante; hexproof do Padeem; anthems de Chief of the Foundry e Master of Etherium (→ 3/7); +1/+1 do gatilho `1+`, da Requisition Raid (modo 3) e dos 5 proliferates + Deepglow Skate; haste da Alibou (pode atacar no turno em que entra); +2/+2 do Blacksmith's Skill; prowess de **35 mágicas não-criatura** (inclusive as 7 instantâneas de proteção/remoção). |
| F5 | Facilitador | **Nada para outras cartas**: o desconto é só dela; vigilância e voar são só dela. Facilita indiretamente pelo tipo (F3). |
| F6 | Curva | CMC nominal 4; custo efetivo **`{1}{U}`/`{U}`** no turno em que se lança outra mágica não-criatura antes → encaixa como 2ª mágica do turno a partir do T3–T4. Não acelera nada. |
| F7 | Atrito | (1) **Vigilância redundante com a Vraska** quando ela está em campo (Vraska dá vigilância a toda criatura-artefato). Vraska é 1-of, não-artefato, morre nos wipes do próprio deck — a vigilância da Traxos só vale sozinha nos jogos sem Vraska, e mesmo aí o ganho é pequeno: atacar por 1–3 no ar (dispara Thopter Spy Network e o gatilho da Alibou) e ainda estacionar na 2ª principal. **Reduz o valor, sim** — a vigilância não é argumento de entrada. (2) Atacante com vigilância **não conta** no X da Alibou (mesmo atrito da Vraska). (3) Prowess é até o fim do turno: exige ordenar o turno (mágicas antes de estacionar). (4) Poder base 1 é o pior eixo para a dor do deck. (5) `{U}` — azul já é a cor mais pedida; neutro contra a saída proposta, que também pede um `{U}`. |

**Veredito da entrada:** a Traxos **não é peça de poder** (base 1), mas é **corpo permanente, protegido e
quase de graça**, cujo poder escala com o recurso mais abundante do deck (35 mágicas não-criatura) e que
dispara todo o pacote de *artifact spell* e de *nontoken que entra*. Ela ganha de uma carta que não põe
corpo em campo; **não ganha de nenhuma criatura do deck** (ver triagem abaixo).

---

## Triagem das saídas possíveis

Filtros prévios, por instrução e histórico: **naves, veículos e receptores de contador** não são
candidatos (memória "não resolver a dor cortando o alvo"); **Vraska e Requisition Raid** entraram em
2026-09-26; **Voyage Home, Cloud Key, Crystalline Crawler, Organic Extinction, Swords to Plowshares**
voltaram pelo usuário sem motivo declarado — só com ressalva; **Loran's Escape** o usuário recusou cortar
na v3 (quer camada extra de proteção do comandante).

### Criatura por criatura (preservaria as 36 mágicas não-criatura) — nenhuma perde para a Traxos

| Carta | Por que a Traxos não a substitui |
|---|---|
| Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector, Voyager Quickwelder | redutores de todo artifact spell (F5) que a Traxos não cobre; Inspector tem poder 3 e Quickwelder 2 — **mais** poder base que ela para Station |
| Emissary Escort | 0/4 + maior MV (Sphinx 9, Monitor 7) = o corpo de mais poder por mana do deck |
| **Thought Monitor** | parecia o par natural (voador-artefato barato), mas é **compra 2** e é **MV 7** — o 2º maior MV de artefato: alimenta o **+X/+0 da Emissary Escort** e a condição do Padeem. Cortá-lo tiraria poder de Station, não daria |
| Coretapper, Hangarback, Marketback, Crystalline Crawler, Kappa | põem/recebem contador — fora por instrução (e Crawler com ressalva do usuário) |
| Metastatic Evangel, Surge Conductor, Kilo | fontes de proliferate — o motor de contadores |
| Chief of the Foundry, Master of Etherium | os anthems que a própria Traxos precisa para chegar a poder 3 |
| Pinnacle Emissary, Sai, Malcator, Chrome Host Seedshark | fábricas de corpo — oferta de poder repetível, maior que a da Traxos |
| Jhoira, Padeem | motores de compra; Padeem é o único hexproof do comandante |
| Alibou, Cyberdrive, Deepglow, Ethersworn Sphinx | wincon/haste, voar em massa, dobra de contadores, cascade |

### Mágica não-criatura por criatura — aqui está a saída

| Carta | Situação |
|---|---|
| **Stern Lesson** | **candidata nº 1** — a fonte de "draw" mais fraca em vantagem real (líquido 0 carta) e a única do deck que **não põe corpo nem motor em campo** |
| **Talisman of Progress** | candidata nº 2 — a rocha menos flexível (W/U, sem R, 1 de dano) |
| Midnight Clock | conta em ramp **e** draw (recarga de 7 por hour counters, que proliferate acelera) — cortar tira de duas metas |
| Tezzeret's Gambit | compra 2 **+ proliferate** — motor de contadores |
| Thopter Spy Network, Golem Foundry, Saheeli, Whirlwind | fábricas de corpo / motores |
| Everflowing Chalice, Pentad Prism, Astral Cornucopia, Solar Array | receptores de charge / arranque frio — fora por instrução |
| Blacksmith's Skill, Restoration Magic, Invisible Force Field, Loran's Escape, Stone by Sunlight | proteção (7) — o usuário já recusou reduzir camadas |
| Remoções e wipes | interação em 10/~10 e wipes em 2/2–4 (no piso) |
| Voyage Home | compra 3 real (líquido +2) — melhor draw que Stern Lesson; e voltou pelo usuário sem motivo |

---

## Saída recomendada (nº 1) — `Stern Lesson` · corte condicionado (draw/ramp)

`{2}{U}` · Instant · CMC 3 · **R$ 0,10** (LigaMagic menor, 2026-09-18) · não está nas sobressalentes.

**Ficha F1–F7**

| # | Eixo | Conteúdo |
|---|---|---|
| F1 | Texto | (a) `Draw two cards, then discard a card` — **loot**: gasta 1 carta, compra 2, descarta 1 → **líquido 0 carta**, só seleção. (b) `Create a tapped Powerstone token` — artefato token; `{T}: Add {C}`, que **não paga mágica não-artefato**. |
| F2 | Corpo | Nenhum. (O Powerstone vira 4/4 até o fim do turno sob o ETB do Cyberdrive Awakener — situacional.) |
| F3 | Tipo | **Mágica não-criatura, instantânea** → Vraska (Sculpture), Saheeli (Servo), Whirlwind (compra), Seedshark (Incubate 3), prowess; em instant speed os tokens nascem no fim do turno do oponente e chegam destapados. **Powerstone** = +1 artefato (Master, Vertibird, Kappa, Malcator, affinity, metalcraft, improvise) — mas é **token**: não dispara Surge Conductor. Não é histórica (não dispara Jhoira). |
| F4 | Receptor | Nenhum. |
| F5 | Facilitador | Mana do Powerstone para artifact spells e ativações. |
| F6 | Curva | CMC 3, instantâneo — flexível, sem pressão de curva. |
| F7 | Atrito | Descarte sem sinergia de cemitério; Powerstone entra virado e é mana restrita. |

**Protocolo de corte (§2 do checklist)**

```
Sai: Stern Lesson — funções: [loot 2/1, Powerstone (ramp restrito + 1 artefato), mágica não-criatura instantânea]
 → loot: coberto pelas outras 12 fontes de draw (Whirlwind, Jhoira, Uthros, Sai, Padeem, Thopter Spy Network,
   Thought Monitor, Ethersworn Sphinx, Midnight Clock, Tezzeret's Gambit, Voyage Home, Marketback);
   a própria Traxos compra com Jhoira ou Uthros 3+ em campo. Draw 13 → 12, dentro da meta 12–13.
 → Powerstone como mana: coberto por Sol Ring, Arcane Signet, Talisman, Chalice, Pentad Prism, Astral Cornucopia,
   Midnight Clock, Solar Array, Crystalline Crawler, Vraska (+ Autogenerator e 5 redutores). Ramp 11 → 10 (+1),
   dentro da meta 10–11.
 → Powerstone como artefato: coberto pela Traxos (+1 artefato, e nontoken → dispara Surge Conductor, que o token não dispara).
 → mágica não-criatura: DESCOBERTA EM PARTE — 36 → 35 gatilhos para Vraska/Saheeli/Whirlwind/Seedshark/prowess,
   e perde-se a única fonte de draw que gera tokens em instant speed. Custo aceito: −1 de 36 (~3%).
```

**Simetria de critério (§3)** — mesmas condições para os dois lados:

| Condição | Traxos (entra) | Stern Lesson (sai) |
|---|---|---|
| Nenhum motor em campo | corpo 1/5 voador protegido → **1 de poder** para Station, todo turno | **0 de poder**; loot + Powerstone |
| Chief + Master em campo | 3/7 → **3** por turno, permanente | 0 (o Powerstone não é criatura) |
| Vraska + Saheeli + os dois anthems | 3 + prowess; **não** dispara Vraska/Saheeli (creature spell) | Sculpture + Servo = **2 corpos de 3/3 = 6 de poder**, permanentes |
| Pacote de *artifact spell* (Pinnacle, Sai, Golem Foundry, Uthros, Jhoira, redutores) | dispara os 5 motores + 5 redutores | não dispara nenhum |
| Pacote de *nontoken que entra* (Surge Conductor, Evangel) | dispara os 2 | não dispara (Powerstone é token) |
| Pacote de *mágica não-criatura* (Vraska, Saheeli, Whirlwind, Seedshark) | não dispara; **recebe** prowess delas | dispara os 4 |
| Artefatos | +1 (nontoken) | +1 (token) |

Lido com honestidade: **no melhor caso** (Vraska + Saheeli + anthems), Stern Lesson rende mais poder que a
Traxos. Mas esse caso pede três 1-ofs específicos; a Traxos entrega corpo **incondicional** e protegido, e
casa com 10 motores do deck contra 4 da Stern Lesson. Contra a dor declarada (economia de **poder** para
Station), um corpo permanente vence uma compra de líquido zero.

**Regra 5 (histórico):** Stern Lesson estava na lista original do usuário; a v1 a manteve em draw e em
ramp ("fica") quando o deck tinha **8/12–13 fontes de draw e ~7 ramps confiáveis**. **O que mudou desde
então:** draw está em 13 e ramp em 11 + 1 — os dois déficits que justificavam mantê-la acabaram, e ela
segue dentro das duas metas depois do corte. Nunca foi cortada nem reposta; sem vai-e-vem.

**Por ser draw/ramp, está fora da minha especialidade → `corte condicionado`**: a decisão final é do
orquestrador, sabendo que as funções que ficariam descobertas são *1 fonte de draw (seleção, líquido 0)*
e *1 ramp restrito*, ambas com as metas ainda cumpridas.

## Alternativa (nº 2) — `Talisman of Progress` · corte condicionado (ramp)

`{2}` · Artifact · CMC 2 · **R$ 4,60** (LigaMagic menor, 2026-09-18) · não está nas sobressalentes.

| # | Eixo | Conteúdo |
|---|---|---|
| F1 | Texto | `{T}: Add {C}`; `{T}: Add {W} or {U}`, 1 de dano em você |
| F2 | Corpo | nenhum (4/4 temporário sob o Cyberdrive) |
| F3 | Tipo | artefato nontoken **e** mágica não-criatura **e** artifact spell → dispara **os dois** pacotes (Vraska/Saheeli/Whirlwind/Seedshark **e** Pinnacle/Sai/Golem Foundry/Uthros/Jhoira), Surge Conductor, redutores |
| F4 | Receptor | proteção do comandante; nada útil além disso |
| F5 | Facilitador | fixa W/U (não R, a cor magra) |
| F6 | Curva | **2-drop de aceleração**: T2 rock → T3 4-drop |
| F7 | Atrito | 1 de dano por uso colorido |

```
Sai: Talisman of Progress — funções: [ramp T2, fixação W/U, artefato, mágica não-criatura + artifact spell]
 → ramp: Sol Ring, Arcane Signet, Pentad Prism, Everflowing Chalice cobrem o T2; ramp 11 → 10, na meta.
 → fixação W/U: coberta com folga (9 Island, 6 Plains, 16 não-básicos que produzem W ou U, Arcane Signet).
 → artefato: coberto pela Traxos.
 → gatilhos: DESCOBERTO — a Traxos não dispara o pacote de mágica não-criatura (36 → 35).
 → aceleração de T2: DESCOBERTA EM PARTE — 2-drops 12 → 11, e rocks de CMC ≤ 2 (Sol Ring, Signet, Prism, Chalice, Talisman) 5 → 4. Custo real de tempo.
```

**Por que fica em 2º:** tira aceleração de turno 2 num deck com 22 cartas de CMC 3, e dispara mais
gatilhos do deck que a Stern Lesson. A única vantagem sobre a nº 1 é o custo: baixa o valor da lista em
R$ 2,71 em vez de subir R$ 1,79 (continua acima do teto de torneio de qualquer jeito). Histórico: entrou
na v1 pelo pipeline (no lugar de Cargo Ship) e foi mantida na v2 como fixação W/U — sem vai-e-vem.

---

## Recontagem (Stern Lesson → Traxos)

| Métrica | Antes | Depois |
|---|---|---|
| Total | 100 | 100 |
| Criaturas de verdade | 27 | **28** |
| Artefatos | 42 | **43** |
| Mágicas não-criatura (não-terreno) | 36 | **35** |
| Instantâneos | 8 | 7 |
| Draw | 13 | **12** (meta 12–13 ✓) |
| Ramp | 11 + 1 | **10 + 1** (meta 10–11 ✓) + 5 redutores |
| Interação / Wipes / Proteção | 10 / 2 / 7 | 10 / 2 / 7 |
| Corpos com poder ≥3 (base) | 12 (+1 dinâmico) | 12 (+1 dinâmico) — Traxos só chega a 3 com os dois anthems |
| Corpos que usam o próprio `{T}` | 4 | 4 |
| Voadores nativos (criaturas) | — | +1 |
| Curva não-terreno (CMC 3 / 4) | 22 / 9 | **21 / 10** (custo efetivo da Traxos: `{1}{U}` ou `{U}`) |
| Pips | — | `{U}` sai, `{U}` entra — neutro |

Com a nº 2 (Talisman): ramp 11 → 10, 2-drops 12 → 11, draw inalterado (13), mágicas não-criatura 36 → 35.

## Delta de custo — LigaMagic (menor), cotações de 2026-09-18 e 2026-09-27

| Troca | Entra | Sai | Compra | Valor da lista |
|---|---|---|---|---|
| **nº 1** | Traxos R$ 1,89 (27/09) | Stern Lesson R$ 0,10 (18/09) | **R$ 1,89** | R$ 252,95 → **R$ 254,74** + Vraska `a cotar` |
| nº 2 | Traxos R$ 1,89 (27/09) | Talisman of Progress R$ 4,60 (18/09) | R$ 1,89 | R$ 252,95 → R$ 250,24 + Vraska `a cotar` |

Base R$ 252,95 = valor da lista após a v4 (report da v4). Nenhuma das duas resolve o estouro do teto de
torneio (R$ 200).

---

## Ficha de funções — uma linha por carta não-terreno do deck (pós-v4, antes da troca)

Anthems = Chief of the Foundry + Master of Etherium (+1/+1 cada, só em criatura-artefato).
"Exposta" = não é artefato, fora da estática do comandante.

| Carta | Categorias | Corpo tapável (F2) | Tipo alimenta (F3) | Recebe (F4) | Facilita (F5) | Entra no turno (F6) | Atritos (F7) |
|---|---|---|---|---|---|---|---|
| Inspirit, Flagship Vessel | tema, proteção | spacecraft (consome corpo); 5/5 voa a 8+ | artefato | charge via Station/proliferate/Deepglow; hexproof do Padeem; one-shots | hexproof+indestrutível aos outros artefatos; `1+` em outro artefato | T3 | desprotegido; `1+` desligado a 0 |
| Hangarback Walker | tema | 0/0 +X; usa `{T}` para crescer | art. criatura | +1/+1, anthems | Thopters ao morrer (sac) | flexível | `{T}` próprio × Station |
| Marketback Walker | tema, draw | 0/0 +X | art. criatura | +1/+1, anthems | compra ao morrer | flexível | — |
| Coretapper | tema | 1/1 (3/3); usa `{T}` | art. criatura | anthems | charge em artefato; sac = 2 | T2 | `{T}` próprio × Station |
| Emissary Escort | tema | 0/4 + maior MV (até 9) | art. criatura | anthems, +1/+1 | — | T2 | depende de artefato caro |
| Etherium Sculptor | tema, ramp | 1/2 (3/4) | art. criatura | anthems | −{1} artifact spells | T2 | — |
| Enthusiastic Mechanaut | tema, ramp | 2/2 voa (4/4) | art. criatura | anthems | −{1} artifact spells | T2 | `{U}{R}` |
| Metastatic Evangel | tema | 3/1 | criatura nontoken | +1/+1 | proliferate por criatura nontoken | T2 | exposta |
| Chief of the Foundry | tema | 2/3 (3/4) | art. criatura | anthem do Master | +1/+1 às art. criaturas | T3 | — |
| Chrome Host Seedshark | tema | 2/4 voa | criatura | +1/+1 | Incubate X por não-criatura | T3 | exposta |
| Foundry Inspector | tema, ramp | 3/2 (5/4) | art. criatura | anthems | −{1} artifact spells | T3 | — |
| Kilo, Apogee Mind | tema | 3/3 haste (5/5) | art. criatura lendária | anthems | proliferate ao virar | T3 | com vigilância não prolifera no ataque |
| Malcator, Purity Overseer | tema | 1/1 voa + Golem 3/3 | criatura (Golems art.) | — | Golems 3/3 | T3 | exposta |
| Master of Etherium | tema, wincon | */* = nº de artefatos | art. criatura | anthem do Chief | +1/+1 às art. criaturas | T3 | — |
| Pinnacle Emissary | tema | 3/3 + Drones 1/1 voa | art. criatura | anthems | Drone por artifact spell | T3 (warp T2) | — |
| Sai, Master Thopterist | tema, draw | 1/4 + Thopters | criatura | — | Thopter por artifact spell; sac 2 → compra | T3 | exposta |
| Surge Conductor | tema | 3/2 (5/4) | art. criatura | anthems | proliferate por artefato nontoken | T3 | — |
| Voyager Quickwelder | tema, ramp | 2/4 (4/6) | art. criatura | anthems | −{1} artifact spells | T3 | — |
| Vraska, Soul of Stone | tema, ramp | 3/3 + Sculptures 1/1; Sculpture usa `{T}` para mana | criatura | — | Sculpture por não-criatura; vigilância às art. criaturas | T3 | exposta; reduz X da Alibou |
| Crystalline Crawler | tema, ramp | 1/1 + converge (até 4/4); usa `{T}` | art. criatura | +1/+1, anthems | mana de qualquer cor | T4 | `{T}` próprio × Station; voltou sem motivo declarado |
| Jhoira, Weatherlight Captain | draw | 3/3 | criatura lendária | — | compra por histórica | T4 | exposta |
| Padeem, Consul of Innovation | draw, proteção | 1/4 | criatura lendária | — | hexproof aos artefatos (inclui o comandante); compra | T4 | exposta |
| Alibou, Ancient Witness | tema, remoção, wincon | 4/5 (6/7) | art. criatura lendária | anthems | haste às art. criaturas; dano X | T5 | vigilância reduz X |
| Deepglow Skate | tema | 3/3 | criatura | — | dobra contadores | T5 | exposta |
| Cyberdrive Awakener | tema, wincon | 4/4 voa | art. criatura | anthems | voar às art. criaturas; anima artefatos | T6 | — |
| Kappa Cannoneer | tema, wincon | 4/4 + contadores | art. criatura | +1/+1, anthems | — | T4–6 (improvise) | improvise × tap |
| Thought Monitor | tema, draw | 2/2 voa (4/4) | art. criatura **MV 7** (Escort, Padeem) | anthems | compra 2 | T4–5 (affinity) | — |
| Ethersworn Sphinx | tema, draw | 4/4 voa (6/6) | art. criatura **MV 9** | anthems | cascade | T5+ (affinity) | — |
| Astral Cornucopia | ramp, tema | — | artefato | charge | mana escalável | flexível | — |
| Everflowing Chalice | ramp, tema | — | artefato | charge | mana | T2 | — |
| Sol Ring | ramp | — | artefato | — | mana | T1 | — |
| Arcane Signet | ramp | — | artefato | — | fixa WUR | T2 | — |
| Pentad Prism | ramp, tema | — | artefato | charge | mana armazenada | T2 | — |
| Spring-Loaded Sawblades | remoção, tema | veículo crew 1 (craft) | artefato | — | 5 de dano (flash) | T2 | só criatura virada |
| Talisman of Progress | ramp | — | artefato | — | fixa W/U | T2 | 1 de dano |
| Brotherhood Vertibird | tema | veículo crew 2; poder = nº de artefatos | artefato | anthems quando tripulado | vira corpo gigante para Station | T3 | consome corpo |
| Cloud Key | tema, ramp | — | artefato | — | −{1} ao tipo escolhido | T3 | voltou sem motivo declarado |
| Golem Foundry | tema | — (Golems 3/3) | artefato | charge | Golem a cada 3 charges | T3 | consome contadores |
| Midnight Clock | ramp, draw, tema | — | artefato | hour counters | `{U}`; compra 7 | T3 | embaralha a mão |
| Perilous Snare | remoção, tema | — | artefato | — | exílio; max speed +1/+1 | T3 | — |
| Solar Array | ramp, tema | — | artefato | — | sunburst (arranque frio) | T3 | — |
| Uthros Research Craft | tema, draw | spacecraft (3+) | artefato | charge | compra por artifact spell | T3 | consome corpo |
| Warmaker Gunship | tema, remoção | spacecraft (6+) | artefato | charge | dano no ETB | T3 | consome corpo |
| Empowered Autogenerator | ramp | — | artefato | charge | mana explosiva | T4 | entra virado |
| Recon Craft Theta | tema | veículo crew 2 + Alien 1/1 | artefato | +1/+1, anthems tripulado | proliferate ao atacar | T4 | consome corpo |
| Thousand Moons Smithy | tema, wincon | Gnome */* | artefato lendário / terreno | — | Gnome; vira terreno | T4 | virar 5 × Station |
| Dawnsire, Sunstar Dreadnought | remoção, wincon | spacecraft (10+/20+) | artefato lendário | charge | 100 de dano | T5 | consome muito corpo |
| Thopter Spy Network | draw, tema | Thopter por upkeep | encantamento | — | compra por dano de art. criatura | T4 | `{U}{U}` |
| Whirlwind of Thought | draw | — | encantamento, não-criatura | — | compra por não-criatura | T4 | `{U}{R}{W}` |
| Saheeli, Sublime Artificer | tema | Servos 1/1 | planeswalker | — | Servo por não-criatura; −2 cópia | T3 | — |
| Blacksmith's Skill | proteção | — | instant, não-criatura | — | hexproof+indestrutível; +2/+2 em art. criatura | T1+ | — |
| Dispatch | remoção | — | instant | — | exílio (metalcraft) | T1+ | — |
| Loran's Escape | proteção | — | instant | — | hexproof+indestrutível + scry | T1+ | usuário quer manter |
| Restoration Magic | proteção | — | instant | — | tiered; Curaga em massa | T1+ | — |
| Swords to Plowshares | remoção | — | instant | — | exílio | T1+ | voltou sem motivo declarado |
| Invisible Force Field | proteção | — | instant | — | indestrutível a 4 + rebound | T2+ | — |
| Stone by Sunlight | remoção, proteção | — | instant | — | destrói poder 4+ / artificializa | T2+ | — |
| **Stern Lesson** | draw, ramp | — | instant, não-criatura; Powerstone (token) | — | loot + mana p/ artefato | T3+ | **saída nº 1** |
| Requisition Raid | remoção, tema | — | feitiço, não-criatura | — | +1/+1 no time (modo 3) | T1+ | entrou em 26/09 |
| Sunder the Gateway | remoção, tema | Incubate 2 → 2/2 art. | feitiço | — | — | T2+ | — |
| Chain Reaction | wipe | — | feitiço | — | assimétrico | T4 | mata não-artefatos do próprio lado |
| Tezzeret's Gambit | draw, tema | — | feitiço | — | compra 2 + proliferate | T4 | — |
| Voyage Home | draw | — | feitiço (affinity) | — | compra 3 + 3 vida | T4+ | voltou sem motivo declarado |
| Organic Extinction | wipe | — | feitiço (improvise) | — | destrói não-artefatos | T5+ | voltou sem motivo declarado |

Terrenos (36): sem alteração nesta troca; os 3 Bridges contam como artefato (F3).
