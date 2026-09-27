# Manabase e Cortes — v5 · troca pontual (`/swap-card Hall of Echoes`)

> Modo `improve`, troca pontual. **Entrada fixa: `Hall of Echoes`** (decisão do orquestrador; nenhuma
> outra entrada avaliada, nenhuma busca no Scryfall). Deck **montado** (Status físico, 2026-09-18);
> lista viva = `deck.md` = lista final da v4. Oracle de todas as cartas citadas puxado nesta sessão
> (`bin/mtgdb oracle` / `bin/mtgdb deck krenko-mob-boss -full`, 2026-09-26). `Hall of Echoes` não tem
> rulings oficiais no banco — as leituras de regra abaixo citam as Comprehensive Rules.

## Cálculo

CMC médio (64 não-terrenos, sem o comandante): **2,33** · draws+ramps com mv ≤ 2 (rótulos do
`deck.md`): **17** · Fórmula: 31,42 + 3,13×2,33 − 0,28×17 = **33,9 → 34 terrenos**

A fórmula literal pede 34, mas 7 dos 17 não são fonte recorrente: 3 cantrips que só viram vantagem
com a Zada (`Expedite`, `Renegade Tactics`, `Ancestral Anger`), 2 rituais de uso único
(`Brightstone Ritual`, `Battle Hymn`), `Fists of Flame` (cantrip) e `Rundvelt Hordemaster`
(condicional). Contando só as 10 fontes recorrentes, a fórmula dá **35,9**. E a própria entrada
**puxa para cima**: `Hall of Echoes` é um mana sink de `{5}` que não se paga com a própria terra.
**Fico em 35** — troca terreno por terreno, sem devolver nem pedir slot de spell.

Listas de referência (Archidekt, bracket 3, comandante confirmado no cabeçalho, quantidades somadas):
#3047743 **35** (29 Mountain + 6) · #7582532 **35** (18 Mountain + 17; esta roda `Hall of Echoes`) ·
#2029118 **35** (19 Mountain + 16). As três batem com a nossa contagem.

## Coleção (regra 7)

`mtgdb collection "Hall of Echoes"` → não está nas sobressalentes: **compra**, R$ 19,98 (LigaMagic
menor, cotação de 2026-09-26 — edição nova com pré-venda, provável volatilidade; reconferir antes
de comprar). Única sobressalente que é terreno: `Blast Zone` — fora do escopo (entrada fixa), só
registrada para constar que a caixa foi olhada.

## Conferência da leitura do orquestrador

Oracle: *"{T}: Add {C}. / {5}: This land becomes a copy of target creature you control until end of
turn. The "legend rule" doesn't apply to permanents you control this turn."*

| Afirmação | Veredito | Detalhe |
|---|---|---|
| Vira cópia de Krenko, é Goblin e conta no X | **Correto** | A cópia copia os valores copiáveis (CR 707.2): Legendary Creature — Goblin Warrior 3/3 com `{T}: Create X…`. Conta para o X dela e do Krenko original. |
| Sem regra da lenda | **Correto** | Só neste turno. No cleanup (CR 514.2) o efeito de cópia e a isenção terminam juntos; a Hall volta a ser a terra, nada vai ao cemitério. |
| Pode usar o `{T}` porque está sob controle desde o início do turno | **Correto, com uma ressalva** | CR 302.6 olha o **objeto**, não o momento em que virou criatura: Hall baixada em turno anterior pode `{T}` como Krenko. Hall baixada **neste turno** precisa de ímpeto — e o deck tem 4 fontes que cobrem a cópia: `Goblin Warchief` e `Howlsquad Heavy`/`Goblin Chieftain` (Goblins/outros Goblins), `Hammer of Purphoros` (todas as criaturas). |
| Segunda ativação de Krenko, em velocidade de instantâneo | **Correto** | As duas habilidades são ativadas sem restrição de velocidade. Linha forte: end step do oponente, com `Brightstone Ritual`/`Battle Hymn` (ambos instantâneos) pagando o `{5}` com o enxame. |
| Tokens ao fim da cópia | **Ficam** | São objetos independentes; a cópia terminar não os afeta. |
| "Deixa de produzir mana no turno" | **Correto, e mais estrito** | (a) A Hall precisa estar **desvirada** — se foi tapada para mana, vira um Krenko tapado e não ativa; o `{5}` sai de **outras** fontes, então a 2ª ativação custa 6 mana-equivalentes. (b) Como cópia ela **não é terreno** (a cópia sobrescreve o tipo; não há "it's still a land") e perde o `{T}: Add {C}`. |
| Exposta a remoção instantânea | **Correto, e o custo é maior que parece** | Se a cópia morrer, o card `Hall of Echoes` vai ao cemitério: **perde-se um terreno**. `Goblin Chirurgeon` (sacrificar Goblin: regenerar criatura) protege a cópia. O Krenko da cópia, se ativado, resolve mesmo que ela morra em resposta (X contado na resolução). |
| Produz `{C}` num mono-R | **Correto** | Ver pips abaixo. |
| Não é Mountain | **Correto** | Afeta `Castle Embereth` e `Dwarven Mine` — ver recontagem. |

Duas observações que o orquestrador não listou:
- **Cópia não "entra"**: `Goblin Matron`, `Mogg War Marshal`, `Volley Veteran` copiados não disparam
  ETB; `Metallic Mimic` não põe contador; `Impact Tremors`/`Purphoros` não disparam pela cópia.
  O alvo é quase sempre Krenko (ou `Purphoros` quando for criatura: duas instâncias do gatilho,
  4 de dano por token a cada oponente).
- **Equipamento e contadores não copiam**: o `Umbral Mantle` do combo continua no Krenko original.
  `Fable of the Mirror-Breaker` (cap III) só copia **não-lendária** — a Hall é a única peça do deck
  que duplica o comandante.

## Pips e fontes de `{R}`

72 pips `{R}` nas 65 não-terrenos (comandante incluso); **13 cartas com `{R}{R}` ou mais**, entre
elas `Goblin Chainwhirler` `{R}{R}{R}` e `Assault on Osgiliath` `{X}{R}{R}{R}`. O `{C}` da Hall só paga
genérico (`{2}` do Krenko, `{1}` do Castle, `{3}` do Den, `{3}` do Mantle, `{8}` do Idol).
P(Hall entre as 7 iniciais) ≈ 7%; nos ~10 primeiros cards ≈ 10%. O risco real é atrasar em um turno
um 2-drop `{R}{R}` (`Conspicuous Snoop`, `Ember Hauler`, `Goblin Wardriver`) quando a Hall é uma das
duas primeiras terras. Com 34 fontes de `{R}` em 35, aceito.

## Candidatas de saída — comparação terreno a terreno

| Terreno | Funções | Veredito |
|---|---|---|
| `Castle Embereth` | `{R}`, anthem `+1/+0` repetível, mana sink | **Fica** — núcleo desde a v3 |
| `Den of the Bugbear` | `{R}`, vira Goblin 3/2 que fabrica Goblins, atravessa wipe, sink | **Fica** — núcleo desde a v3 |
| `Dwarven Mine` | `{R}`, **é Mountain**, Anão 1/1 no ETB (gatilho de Tremors/Purphoros) | **Fica** — é o único não-básico que *soma* Mountains; tirar ele equivale a tirar uma Mountain e ainda perder o Anão |
| `Sokenzan, Crucible of Defiance` | `{R}` desvirado, Channel em instantâneo: 2 corpos com ímpeto (2 gatilhos de ETB, comida de Bombardment/Skullclamp), barato com 7 lendárias | **Fica** — faz tudo que a Mountain faz e mais o Channel; entrou comprado na v4 |
| **`Mountain` (1 de 30)** | `{R}`, subtipo Mountain, desvirada | **Candidata 1** |
| **`The Autonomous Furnace`** | `{R}`, entra virada, `{1}{R}`,`{T}`, sac: compra 1 | **Candidata 2** |

### Candidata 1 — `Mountain` (1 de 30) · **recomendada**

*"({T}: Add {R}.)"* — Basic Land — Mountain.

- **F1** Produz `{R}`. **F2** Não é corpo. **F3** Land — **subtipo Mountain** (conta para
  `Castle Embereth` "unless you control a Mountain" e `Dwarven Mine` "three or more other Mountains").
  **F4** —. **F5** Fonte de `{R}`. **F6** Sempre desvirada. **F7** Nenhum.

| Função | Quem cobre |
|---|---|
| Land drop | `Hall of Echoes` (desvirada, sempre) |
| Fonte de `{R}` | **Não coberta pela Hall** (`{C}`) → 34 fontes restantes (29 Mountain + Castle, Den, Mine, Sokenzan, Furnace). Custo aceito: ver pips |
| Subtipo Mountain | 29 Mountain + `Dwarven Mine` = **30** Mountains. `Castle Embereth` precisa de **uma**; `Dwarven Mine` de **três outras** — diferença de 31→30 é desprezível para as duas |

**Histórico (`decisions.md`)**: Mountains já saíram duas vezes pelo mesmo caminho — v3 (36→35, MV
2,45 + Incubator) e v4 (1:1 por `Dwarven Mine`, "terreno com upside"). Esta é a terceira, pelo mesmo
motivo da v4. Não há defesa registrada de basic específico.

**Simetria**: elogiei a Hall assumindo Krenko em campo e 5 mana sobrando. Nas mesmas condições a
Mountain rende 1 `{R}` e nada mais; sem Krenko, a Mountain rende 1 `{R}` e a Hall rende 1 `{C}` —
perda de uma cor que o deck tem de sobra. Nenhuma condição favorece a Mountain.

**Por que ela e não a Furnace**: é o único corte em que **nada fica descoberto além de redundância**,
e preserva a única carta que a v3 defendeu nominalmente na manabase (ver abaixo).

### Candidata 2 — `The Autonomous Furnace` · alternativa

*"This land enters tapped. {T}: Add {R}. {1}{R}, {T}, Sacrifice this land: Draw a card."* — Land — Sphere.

- **F1** Entra virada; `{R}`; converte a si mesma em uma carta por `{1}{R}`. **F2** Não é corpo.
  **F3** Land — Sphere (nenhuma carta do deck conta Sphere; **não** é Mountain). **F4** —.
  **F5** Fonte de `{R}`. **F6** **Única terra do deck que entra virada sempre** (Castle, Den e Mine
  são condicionais). **F7** Disputa `{1}{R}` com as outras ativações; ao sacrificar, reduz a contagem
  de terras (irrelevante no late).

| Função | Quem cobre |
|---|---|
| Land drop | `Hall of Echoes` — e **desvirada**, o que é ganho |
| Fonte de `{R}` | Não coberta pela Hall → mesmas **34** fontes da candidata 1 (30 Mountain + Castle, Den, Mine, Sokenzan) |
| Anti-flood no slot de terreno (terra → carta) | **Parcial.** Com criatura em campo, a Hall é um sink muito maior (`{5}` → uma ativação inteira de Krenko). **Com o campo vazio pós-wipe, a Hall não faz nada** e a Furnace ainda compra. Resta no slot de terreno: `Den of the Bugbear` (mana → Goblin, sem precisar de board), `Sokenzan` (da mão); fora dele, `Hammer of Purphoros` (terra → Golem 3/3, exige o Hammer). **Custo declarado: 1 carta de flood relief incondicional no pós-wipe — queixa nº 3 do intake** |
| Fonte de draw (`deck.md`: `draw`) | Contava entre os 3 "filtros net-zero" da v4; as **12 fontes reais** não mudam (18 → 17 no total) |

**Histórico (`decisions.md` e `v3/06-manabase.md`)**: veio na lista de entrada (baseline). A v3
**defendeu-a por escrito** — *"as duas únicas cartas do deck que resolvem flood no próprio slot de
terreno. Cortá-las por 'entra virado' seria julgar por um eixo só"*. Na v4 a irmã
(`Forgotten Cave`) saiu com "cycling → NINGUÉM cobre (perda declarada)", o que tornou a Furnace a
**única** que restou. **O que mudou desde a v3**: só a própria Hall — uma segunda terra-sink, mas
condicionada a ter criatura. O motivo da v3 (flood relief incondicional) continua **parcialmente
válido**; por isso ela fica em segundo.

**Simetria**: se a Hall é elogiada como sink de flood *com Krenko vivo*, a Furnace precisa ser julgada
no cenário oposto também — *sem board*, que é justamente onde ela é a única das duas que funciona.
A vantagem dela sobre a Mountain (1 carta no late) é real; a desvantagem (entra virada) também, e já
foi pesada e aceita na v3.

**Quando preferir a candidata 2**: se o usuário der mais peso a "fechar mais rápido / curva
travando" (0 terras sempre viradas, 31 Mountains) do que ao flood relief pós-wipe.

## Recontagem

| | Hoje (v4) | Sai Mountain (rec.) | Sai Furnace (alt.) |
|---|---|---|---|
| Terrenos | 35 | **35** | **35** |
| Básicos | 30 | 29 | 30 |
| Não-básicos | 5 | 6 | 5 |
| Fontes de `{R}` | 35 | **34** | **34** |
| Fontes de `{C}` (terra) | 0 | 1 | 1 |
| Mountains (subtipo) | 31 | **30** | **31** |
| Terras sempre viradas | 1 | 1 | **0** |
| Terra → carta (flood relief) | 1 | 1 | **0** |
| Terras que duplicam o comandante | 0 | 1 | 1 |
| Total de cartas | 100 | **100** | **100** |

## Terrenos recomendados

| Terreno | Produz | Entra virado? | Sinergia/Utilidade | Na coleção? |
|---|---|---|---|---|
| Hall of Echoes | `{C}` | Não | 2ª ativação de Krenko por turno, em instantâneo; duplica qualquer lenda | Não — compra, R$ 19,98 (LigaMagic menor, 2026-09-26) |
| Castle Embereth | `{R}` | Só sem Mountain | Anthem `+1/+0` repetível | No deck |
| Den of the Bugbear | `{R}` | A partir da 3ª terra | Goblin 3/2 que fabrica Goblins; atravessa wipe | No deck |
| Dwarven Mine | `{R}` | Com < 3 outras Mountains | É Mountain; Anão 1/1 no ETB | No deck |
| Sokenzan, Crucible of Defiance | `{R}` | Não | Channel: 2 corpos com ímpeto em instantâneo | No deck |
| The Autonomous Furnace | `{R}` | Sempre | Terra → carta | No deck (sai na alternativa) |

Básicos: **29 Mountain** na recomendada (30 na alternativa). Mono-R: 100% dos pips são `{R}`.

## Plano de cortes (deck em 101 com a entrada → 100)

| Corte proposto | CMC | Motivo |
|---|---|---|
| **Mountain** (1 de 30) | — | Única função não coberta é 1 fonte de `{R}` (34 restantes); preserva o flood relief incondicional da Furnace, defendido na v3 |
| *alternativa:* The Autonomous Furnace | — | Tira a única terra sempre virada; custo declarado: terra→carta no pós-wipe fica descoberta |

Nenhum corte condicionado: as duas candidatas só exercem funções de manabase (nenhuma é corpo,
tipo contado, redutor ou receptor).

## Curva final projetada

Inalterada (troca terreno por terreno), 65 não-terrenos com o comandante:
0–1: 14 · 2: 22 · 3: 21 · 4: 8 · 5: 0 · 6+: 0 — MV médio 2,33 (sem o comandante).

## Riscos

1. **Hall morre como cópia → perde-se um terreno.** Mitigação: ativar no end step do oponente (menos
   janelas de resposta), ter `Goblin Chirurgeon` com Goblin para sacrificar.
2. **Custo real de 6 mana-equivalentes** (`{5}` de fora + a Hall desvirada): antes dos turnos 6–7 só
   liga com `Sol Ring`, rituais ou Treasures.
3. **Carta morta sem criatura** — no pós-wipe ela é um `{C}`.
4. **Preço volátil** (pré-venda): reconferir a cotação antes de comprar.
