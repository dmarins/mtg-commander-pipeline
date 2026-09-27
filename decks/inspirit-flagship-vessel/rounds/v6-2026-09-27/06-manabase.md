# Manabase e Cortes — v6 · troca pontual (`/swap-card Room of Refuge`)

> Modo `improve`, troca pontual. **Entra fixo: Room of Refuge.** Lista de referência: `deck.md` da
> raiz (montado, confirmado em 2026-09-18, com Vraska e Requisition Raid já dentro). A pasta
> `rounds/v5-2026-09-27/` (sessão concorrente) **não** foi lida. Todo oracle abaixo foi puxado nesta
> sessão com `bin/mtgdb oracle` (regra 6). Preços: LigaMagic (menor), via `mtgdb prices`.

## Cálculo

CMC médio: **3,08** (63 não-terrenos, 194 de MV total; X contado como 0 em Hangarback, Marketback,
Astral Cornucopia e Everflowing Chalice) · draws+ramps mv≤2: **9** (Sol Ring, Arcane Signet,
Talisman of Progress, Pentad Prism, Everflowing Chalice, Astral Cornucopia, Etherium Sculptor,
Enthusiastic Mechanaut; Marketback Walker) · Fórmula: 31,42 + 3,13×3,08 − 0,28×9 = **38,5 terrenos**

Listas de referência (Archidekt, bracket 3 — **o briefing não declara bracket**; usei 3 como
aproximação, o mais comum entre as listas do Inspirit): #14529919 **36** · #15081774 **34** ·
#15666449 **34** (somadas as quantidades; `2x Island`/`2x Plains` contam 2).

**Leitura.** A fórmula pede 38–39; as listas reais rodam 34–36; o deck está em 36. A divergência
tem causa conhecida: a fórmula só desconta aceleração de mv ≤ 2, e este deck tem **mais 5 fontes de
mana de mv 3–4** (Midnight Clock, Solar Array, Cloud Key, Empowered Autogenerator, Crystalline
Crawler), 3 redutores de mv 3 (Foundry Inspector, Voyager Quickwelder, Cloud Key), Powerstone da
Stern Lesson e Treasures da Vraska. E as listas de 34 compram isso com manabase quase toda destapada
(shocks, fetches, triomes, duais originais), coisa que este deck não tem.

**O 36 é decisão do usuário.** A v2 fechou em 38; entre 13/08 e 18/09 ele cortou por conta própria
**Temple of Epiphany e uma Mountain** (decisions.md, "Registro pós-v3"), motivo não declarado. O
goldfishing nunca rodou, então não há medição de travamento de mana a favor de subir a contagem.
Voltar a 37 é desfazer uma escolha dele — antes, é preciso perguntar (regra 5).

## Ficha da entrada — Room of Refuge (oracle desta sessão)

`Land · CMC 0 · identidade incolor` · R$ 0,39 (LigaMagic menor, 2026-09-27) · não está na coleção → compra.
`mtgdb collection -list` não traz nenhum terreno sobressalente (a caixa tem 1 carta, Wilderland Scrounger).

| # | Eixo | Room of Refuge |
|---|---|---|
| F1 | Texto | (1) entra virada; (2) ao entrar, escolhe **uma** cor; (3) `{T}`: 1 mana **dessa** cor, pelo resto do jogo; (4) `{5}, {T}, sacrifique`: dois contadores +1/+1 em criatura alvo, **só como feitiço**. |
| F2 | Corpo | Nenhum. Pode ser sacrificada, mas só pela própria habilidade (não há consumidor de sacrifício de terreno no deck). |
| F3 | Tipo | Land sem subtipo. **Não é artefato**: não entra nos 42 artefatos, nem em affinity (Thought Monitor, Ethersworn Sphinx, Voyage Home), metalcraft (Dispatch), improvise (Kappa, Organic Extinction) ou no poder do Master of Etherium/Brotherhood Vertibird; não recebe a proteção do Inspirit. Sem Plains/Island/Mountain → não liga Clifftop Retreat, Glacial Fortress nem Port Town. |
| F4 | Receptor | Nada. Não tem contador próprio (diferente da Vivid Crag, o proliferate não a alcança). |
| F5 | Facilitador | Fixação de **uma** cor escolhida na entrada. Os 2 contadores +1/+1: em Hangarback = +2 Thopters ao morrer; em Marketback = +2 cartas ao morrer (à taxa de 2,5 por contador, melhor que o `{4}` da própria Marketback); em Crystalline Crawler = 2 manas de qualquer cor devolvidas; em qualquer corpo = **+2 de poder = +2 charge counters por Station**; depois disso, Surge Conductor, Kilo, Metastatic Evangel, Recon Craft Theta e Tezzeret's Gambit proliferam em cima. Deepglow Skate só dobra se a Room for ativada antes (feitiço + feitiço no mesmo turno custa ~11 manas — raro). Não mira o Inspirit antes do 8+ (não é criatura) e não coloca charge. |
| F6 | Curva | Jogada ideal no T1 (virada sem custo). O sorvedouro é de T7+: 6 fontes de mana (as 5 pagas + ela mesma) e um terreno a menos, para sempre. |
| F7 | Atrito | Entra virada. Travada numa cor. O sacrifício tira terreno do campo — é inversão de ramp, aceitável só em flood. A ativação disputa mana com Marketback (`{4}`), Hangarback (`{1},{T}`), craft da Sawblades e com a Organic Extinction (improvise conta artefato, não terreno). |

**Os três argumentos da triagem, conferidos:**

1. **Fixação — procede, com ressalva.** R é a cor mais magra (12 pips em 10 cartas; 4 Mountains;
   ~15 terrenos que dão R contando os flexíveis, contra ~22 de W e ~22 de U). A Room vira a fonte
   da cor que falta **na hora em que entra** — mas fica presa nela; não é "qualquer cor a cada tap".
2. **Contadores +1/+1 — procede como 2º ponto de sinergia, fraco.** É um sorvedouro de flood de
   6 manas + 1 terreno por 2 contadores; os pagadores (Hangarback, Marketback, Crawler, Station)
   existem e as 5 fontes de proliferate multiplicam. A Deepglow depende de sequência improvável.
3. **"Lacuna de 36 contra 38" — contesto.** O 36 cabe na faixa das listas reais (34–36), a fórmula
   ignora 5 fontes de mana de mv 3–4 e o usuário **escolheu** sair de 38 para 36. Não é lacuna
   medida; é distância para uma referência que o próprio dono do deck rejeitou.

## Simetria com a triagem da Evolving Wilds (Tori v2, §T)

A Wilds ficou fora do Inspirit com **1 ponto** (fixação), porque "o deck já tem a Perilous Landscape
e 8 terrenos virados, e a Wilds não é artefato". Aplicando o mesmo critério à Room:

- **"Não é artefato"** vale igual para a Room. Se ela sair de uma Bridge, o deck perde artefato e
  ela é vetada; se sair de outro terreno não-artefato, o critério é neutro na troca.
- **"8 terrenos virados"** (7 sempre virados — Mystic Monastery, Irrigated Farmland, Temple of
  Enlightenment, Vivid Crag e as 3 Bridges — mais a Perilous Landscape, que busca básico virado)
  vale igual: **qualquer rota que leve a contagem de virados de 8 para 9 falha pelo mesmo motivo que
  eliminou a Wilds.** Isso elimina a rota (b) e a rota (a) contra básico. Só sobra trocar a Room
  por um terreno que **já** entra virado.
- O que distingue a Room da Wilds é o 2º ponto (o sorvedouro de contadores). Ele não compensa um
  9º terreno virado; basta para desempatar contra uma tapland que tenha menos a oferecer.

## Rota (a) — terreno por terreno

| Terreno que sairia | O que ele faz (oracle desta sessão) | Contra a Room | Veredito |
|---|---|---|---|
| **Temple of Enlightenment** | virado; scry 1 ao entrar; `{T}`: W **ou** U | fixa W/U, as duas cores já com ~22 fontes; a Room pode ser a fonte de R. Virados: 8 → 8. Perde scry 1; ganha o sorvedouro | **1ª candidata** |
| Mystic Monastery | virado; U, R **ou** W a cada tap | fixa **as três** cores a todo turno — domina a Room em fixação | fica |
| Irrigated Farmland | virado; W ou U; **cycling {2}**; tipos **Plains Island** (liga Clifftop Retreat e Glacial Fortress, pode ser revelada para a Port Town) | o cycling já é o seguro contra flood (e melhor: uma carta contra 2 contadores) e o tipo alimenta 3 terrenos condicionais | fica |
| Vivid Crag | virado com 2 charge; R, ou qualquer cor gastando charge — **recarregável por proliferate** (5 fontes) | já é R nativo + qualquer cor, no eixo de contadores | fica |
| Rustvale / Razortide / Silverbluff Bridge | virado; indestrutível; **Artifact Land** | tirar artefato é exatamente o critério que vetou a Wilds | ficam |
| Perilous Landscape | destapado dando `{C}`; sac → básico virado; cycling `{U}{R}{W}` | mesma fixação "uma cor escolhida", mais o `{C}` no T1 (Sol Ring) e o cycling | fica |
| Básico (Island, o mais sobrante) | destapado; tipo Island liga Glacial Fortress e Port Town; alvo da Perilous Landscape | virados 8 → 9: **falha na simetria com a Wilds** | 2ª, **não recomendada** |

### Ficha da saída nº 1 — Temple of Enlightenment

`Land · CMC 0 · identidade WU` · R$ 0,50 (LigaMagic menor, 2026-09-18) · no deck desde a v1.

| # | Eixo | Temple of Enlightenment | Quem cobre depois do corte |
|---|---|---|---|
| F1a | entra virado | — | a Room também entra virada: a contagem de virados fica em 8 |
| F1b | scry 1 ao entrar | seleção no T1–T2 | **descoberto** — custo aceito: 1 scry uma vez por jogo. Seleção restante: Loran's Escape (scry 2), Alibou (scry por artefato virado), Stern Lesson (loot), cycling da Irrigated Farmland e da Perilous Landscape |
| F1c | `{T}`: W ou U | fonte dupla de W/U | W e U caem de ~22 para ~21 fontes cada (6 Plains, 9 Island, Command Tower, Spire, Orchard, Monastery, Landscape, Skycloud, Port Town, Glacial, Irrigated, Razortide) — e a Room pode ser qualquer uma das duas. **Perda real**: a Temple escolhe W *ou* U a cada tap; a Room escolhe uma vez. Custo aceito, porque W/U são as cores sobrando e R a que falta |
| F2 | corpo | nenhum | — |
| F3 | tipo | Land sem subtipo, não artefato | igual à Room: nenhuma contagem muda (artefatos 42, affinity, metalcraft, improvise, checklands) |
| F4 | receptor | nada | — |
| F5 | facilitador | só a fixação (F1c) | ver F1c |
| F6 | curva | T1 virada | idem Room |
| F7 | atrito | virada | idem Room |

**Histórico (regra 5):** Temple of Enlightenment nunca foi cortada nem reposta. Foi apontada na v1
(`02-theme.md`: "Sempre tapped... agrava a lentidão") e na v2 (`02-theme.md`: "Temples/Farmland/
Monastery continuam sendo o maior custo de velocidade") como tapland substituível. A irmã dela,
**Temple of Epiphany, foi cortada pelo próprio usuário** entre 13/08 e 18/09. Room of Refuge nunca
esteve no deck. Nenhuma peça que o usuário trouxe ou pediu para fazer funcionar é tocada.

**Balanço da troca:** tempo igual (virada por virada), total de terrenos igual, artefatos iguais;
troca "2 cores dobradas, as que já sobram" + scry 1 por "1 cor flexível, que pode ser o R" +
sorvedouro de contadores no flood. É **ganho pequeno, mas líquido**, e não passa por cima de
nenhuma decisão do usuário (36 terrenos e 8 virados preservados).

## Rota (b) — terreno por não-terreno (Room como 37º terreno)

**Não recomendo.** Três motivos, qualquer um suficiente:

1. **Simetria:** vai de 8 para 9 terrenos virados — o critério que tirou a Evolving Wilds do Inspirit.
2. **Desfaz decisão do usuário:** ele cortou uma Temple e uma Mountain para chegar a 36, sem motivo
   declarado. Subir para 37 com um terreno virado é repor exatamente o que ele tirou (uma tapland).
   Pela regra 5, isso é pergunta ao usuário, não proposta.
3. **Não há slot de mágica sobrando:** interação em 10 (meta ~10), wipes em 2 (piso), draw em 13
   (teto da meta, mas as peças são multifunção: Stern Lesson dá artefato + mágica não-criatura; Thought
   Monitor e Ethersworn Sphinx são artefatos; Tezzeret's Gambit é uma das 5 fontes de proliferate;
   Voyage Home foi cortada pelo pipeline na v1 e **o usuário a trouxe de volta**). Proteção (7, sem
   meta) protege a única peça que a estática não cobre — o próprio Inspirit — e as quatro de 1–2
   manas também disparam Vraska, Saheeli e Whirlwind. Não há corte limpo que financie um terreno
   cuja necessidade não foi medida.

Se o goldfishing (pendência nº 1 desde a v1) mostrar travamento de mana, a rota (b) volta a ser
discutível, **com o usuário perguntado antes**.

## Terrenos recomendados (pós-troca, 36)

| Terreno | Produz | Entra virado? | Sinergia/Utilidade | Na coleção? |
|---|---|---|---|---|
| **Room of Refuge** (entra) | 1 cor escolhida | sim | fonte flexível (R, na prática) + sorvedouro de +1/+1 no flood | não — compra R$ 0,39 |
| ~~Temple of Enlightenment~~ (sai) | W/U | sim | scry 1 | no deck (montado) |
| demais 35 | inalterados | — | — | no deck (montado) |

Básicos: **6 Plains, 9 Island, 4 Mountain** (inalterados) — proporção de pips: W 24 (38%) · U 27
(43%) · R 12 (19%), contando o híbrido da Saheeli nas duas cores e o `{U/P}` do Gambit como U;
básicos em 32% / 47% / 21%. R segue coberto por 11 terrenos não-básicos + rocks de qualquer cor.

## Plano de cortes (deck em 100 → 100)

| Corte proposto | CMC | Motivo |
|---|---|---|
| **Temple of Enlightenment** (1ª) | — | tapland por tapland: mesma velocidade, mesmo total, mesmas contagens; troca fixação W/U redundante + scry 1 por fonte flexível (R) + sorvedouro de contadores. Scry 1 fica descoberto (custo aceito) |
| Island (2ª, **não recomendada**) | — | só se o usuário preferir manter a Temple; custa 1 terreno virado a mais (8 → 9) — falha no critério que vetou a Evolving Wilds |

## Recontagem pós-troca (Temple of Enlightenment → Room of Refuge)

| Métrica | Hoje | Depois |
|---|---|---|
| Total | 100 (comandante + 63 não-terrenos + 36 terrenos) | **100** (idem) |
| Terrenos | 36 | **36** |
| Terrenos que entram virados | 8 (7 sempre + Perilous Landscape) | **8** (Room no lugar da Temple) |
| Condicionais | 3 (Clifftop, Glacial, Port Town) | 3 — nenhum perde habilitador (Temple não tem tipo básico) |
| Fontes de W (terrenos) | ~22 | ~21 + Room |
| Fontes de U (terrenos) | ~22 | ~21 + Room |
| Fontes de R (terrenos) | ~15 (4 Mountain + CT, Spire, Orchard, Monastery, Landscape, Forge, Prairie, Clifftop, Vivid Crag, Rustvale, Silverbluff) | ~15 + Room |
| Artefatos | 42 | **42** |
| Categorias afetadas | — | só `terreno` (1 por 1); draw 13, ramp 11+1, interação 10, wipes 2, proteção 7 inalterados. Perde-se 1 scry (sem categoria) |
| Curva (não-terrenos) | 0–1: 11 · 2: 12 · 3: 22 · 4: 9 · 5: 3 · 6+: 6 | inalterada |

> "~" porque Spire of Industry (vida + artefato), Exotic Orchard (depende dos oponentes), Perilous
> Landscape (busca básico) e o "qualquer cor" da Vivid Crag (2 usos, recarregáveis) são contados
> como fonte; em contagem estrita, R tem 12.

**Custo:** compra de **R$ 0,39** (Room of Refuge, LigaMagic menor, cotação de 2026-09-27). Valor da
lista: R$ 252,95 − 0,50 (Temple, cotação de 2026-09-18) + 0,39 = **R$ 252,84 + Vraska `a cotar`**
— continua acima do teto de torneio. A Temple só vira sobressalente se o usuário a listar na próxima
`/update-collection` (regra 7).

## Slots tocados (para reconciliar com a v5 concorrente)

- **Terrenos → Temple of Enlightenment** (sai) e **Room of Refuge** (entra). Nenhum não-terreno é tocado.
- A proposta depende de a contagem de terrenos continuar em **36** e a de virados em **8**. Se a v5
  mexer em qualquer terreno (sobretudo em Temple of Enlightenment, Mystic Monastery, Irrigated
  Farmland, Vivid Crag ou num básico), refazer a comparação da rota (a).

---

## Revisão após reprovação (2026-09-27)

**Contexto novo.** (1) O usuário reprovou a saída da Temple of Enlightenment: *"faz scry 1, isso ajuda
o deck"*. A Temple fica **protegida** (decisions.md, seção v6), e a seleção no começo do jogo deixa
de ser custo aceitável. (2) A **v5 foi aplicada**: entrou Traxos, Academy Guardian e saiu Stern
Lesson. `mtgdb deck inspirit-flagship-vessel` agora dá **43 artefatos, 28 criaturas, 7
instantâneos**; curva 0:4 1:7 2:12 3:21 4:10 5:3 6:2 7+:4. Segundo o registro da v5, draw está em
**12** e ramp em **10 + 1**. A Stern Lesson era uma das coberturas de seleção citadas acima e saiu.

**Critérios que continuam valendo:** 36 terrenos (escolha do usuário), 8 virados como teto
(simetria com a Evolving Wilds), Bridges vetadas por serem artefato. **Critério novo, por
simetria:** se o scry 1 da Temple vale o slot, qualquer terreno com scry, cycling ou filtro tem o
mesmo peso de seleção/suavização e **não sai** para dar lugar à Room.

**Fórmula, recalculada:** CMC médio 195/63 = 3,10; mv ≤ 2 inalterado (9; a Stern Lesson tinha mv 3)
→ 31,42 + 3,13×3,10 − 0,28×9 = **38,6**. A leitura da seção "Cálculo" não muda: 36 fica.

**Seleção de início de jogo que resta no deck** (oracle desta sessão): scry 1 da Temple of
Enlightenment, cycling `{2}` da Irrigated Farmland, cycling `{U}{R}{W}` da Perilous Landscape, scry 2
da Loran's Escape e o scry da Alibou (5 manas; não é início de jogo). Sem a Stern Lesson, os
terrenos com seleção passaram a ser **a maior parte** da suavização do deck, o que reforça a proteção.

### Varredura dos 35 terrenos restantes (Temple fora da mesa)

| Terreno | Entra virado? | Seleção (scry/cycling/filtro)? | Com a Room no lugar | Veredito |
|---|---|---|---|---|
| Plains ×6, Island ×9, Mountain ×4 | não | — | virados 8 → **9**; também tira o habilitador de Clifftop/Glacial/Port Town e o alvo da Perilous Landscape | **falha** (critério dos 8) |
| Command Tower, Spire of Industry, Exotic Orchard, Battlefield Forge | não | — | virados 8 → 9, e cada um dá 2–3 cores contra 1 | **falha** |
| Rugged Prairie, Skycloud Expanse | não | **filtro** | virados 8 → 9 **e** perde filtro | **falha** (duas vezes) |
| Clifftop Retreat, Glacial Fortress, Port Town | condicional (quase sempre destapada) | — | virados 8 → 9 | **falha** |
| Perilous Landscape | `{C}` destapada; busca básico virado | **cycling** + afina o grimório | perde seleção (simetria com a Temple) | **falha** |
| Irrigated Farmland | sim | **cycling `{2}`** | perde seleção; perde ainda os tipos Plains Island (liga Clifftop e Glacial, pode ser revelada para a Port Town) | **falha** |
| Rustvale / Razortide / Silverbluff Bridge | sim | — | tira artefato (43 → 42), o critério que vetou a Wilds | **falha** |
| Mystic Monastery | sim | — | ver ficha abaixo | **falha** |
| Vivid Crag | sim | — | ver ficha abaixo | **falha** (troca lateral) |

Só dois terrenos passam pelos filtros de tempo, seleção e artefato: **Mystic Monastery e Vivid Crag**.
Nenhum dos dois perde para a Room.

### Ficha — Mystic Monastery (oracle desta sessão)

`Land · CMC 0 · identidade RUW` · R$ 0,20 (LigaMagic menor, 2026-09-18) · no deck desde a v1; nunca cortada.

| # | Eixo | Mystic Monastery | Com a Room no lugar |
|---|---|---|---|
| F1 | Texto | entra virada; `{T}`: U, R **ou** W | a Room entra virada igual, mas dá **uma** cor, travada na entrada |
| F2 | Corpo | nenhum | — |
| F3 | Tipo | Land sem subtipo, não artefato | igual |
| F4 | Receptor | nada | igual |
| F5 | Facilitador | fonte das **três** cores a cada tap; é uma das 12 fontes estritas de R e paga sozinha `{U}{R}{W}` em sequência de turnos (Kilo, Vraska, Whirlwind of Thought, Alibou) | **descoberto**: W, U **e** R perdem uma fonte cada, e a Room só devolve uma delas. É justamente o eixo (R) que a triagem queria reforçar |
| F6 | Curva | tapland de T1 | igual |
| F7 | Atrito | virada | igual |

A Monastery vence a Room em toda mão possível. O único ganho seria o sorvedouro de contadores, que
só vale no flood e custa um terreno. **Não passa.**

### Ficha — Vivid Crag (oracle desta sessão)

`Land · CMC 0 · identidade R` · R$ 0,49 (LigaMagic menor, 2026-09-18) · **entrou pela mão do
usuário** (diff v2 → hoje, decisions.md "Registro pós-v3"); nunca foi cortada.

| # | Eixo | Vivid Crag | Com a Room no lugar |
|---|---|---|---|
| F1 | Texto | entra virada **com 2 charge counters**; `{T}`: R; `{T}`, remova um charge: 1 mana de **qualquer** cor | a Room dá 1 cor escolhida; nada de contadores |
| F2 | Corpo | nenhum | — |
| F3 | Tipo | Land sem subtipo, não artefato | igual |
| F4 | Receptor | **charge counters**: recarregada pelas 5 fontes de proliferate (Surge Conductor, Kilo, Metastatic Evangel, Recon Craft Theta, Tezzeret's Gambit) e dobrada pela Deepglow Skate | **descoberto**: sai a única terra do deck no eixo de contadores |
| F5 | Facilitador | R garantido (a cor que falta) + 2 usos de qualquer cor, renováveis | a Room escolhida como R iguala só a 1ª metade; escolhida como W/U, o deck perde uma fonte de R |
| F6 | Curva | tapland de T1 | igual |
| F7 | Atrito | virada; os charges se esgotam sem proliferate | a Room não esgota, mas também não cresce |

Na fixação, Crag e Room empatam (R fixo + flexível limitado contra 1 cor livre). No tema, a Crag
recebe contadores (proliferate, Deepglow); a Room dá contadores (sorvedouro de 6 manas + 1 terreno).
É **troca lateral**: não melhora o deck, custa R$ 0,39 e mais uma carta rodando, e desfaz uma
escolha do usuário sem ganho a mostrar. **Não passa.**

### Rota (b) — revista

Continua reprovada, agora com mais força. A Room como 37º terreno leva os virados a 9 e desfaz o 36
que o usuário escolheu. E, depois da v5, **draw está em 12 e ramp em 10, os dois no piso da meta**.
Interação está em 10 e wipes em 2 (piso). Não existe mágica que saia sem derrubar uma categoria
abaixo da meta.

### Conclusão

**A Room of Refuge não entra no Inspirit.** Com a Temple protegida pelo scry, o mesmo peso de
seleção protege Irrigated Farmland, Perilous Landscape, Rugged Prairie e Skycloud Expanse. Os
básicos e os terrenos destapados levariam os virados de 8 para 9, o mesmo critério que tirou a
Evolving Wilds do deck. As Bridges são artefato. Sobram Mystic Monastery, que domina a Room em
fixação, e Vivid Crag, que empata na fixação, ganha no eixo de contadores e foi escolha do usuário.
Nenhuma saída melhora o deck.

O ganho que a Room trazia era **pequeno desde o começo**: uma fonte flexível de R e um sorvedouro
de flood. Num deck cujos terrenos virados já foram todos escolhidos por uma segunda função (scry,
cycling, tipo básico, artefato, contadores ou três cores), não sobra slot em que esse ganho pague a
troca. Se o orquestrador quiser outro destino para a carta, é uma nova triagem por cor, fora deste
deck.

### Recontagem (sem troca)

| Métrica | Valor |
|---|---|
| Total | 100 (comandante + 63 não-terrenos + 36 terrenos) |
| Terrenos | 36 |
| Terrenos que entram virados | 8 (Mystic Monastery, Irrigated Farmland, Temple of Enlightenment, Vivid Crag, 3 Bridges + Perilous Landscape) |
| Fontes por cor (terrenos) | W ~22 · U ~22 · R ~15 (estritas de R: 12) |
| Artefatos | 43 |
| Categorias | inalteradas (draw 12, ramp 10 + 1, interação 10, wipes 2, proteção 7) |
| Custo | nenhuma compra; lista em R$ 254,74 + Vraska `a cotar` (LigaMagic menor, cotações de 2026-09-18 a 2026-09-27, conforme o registro da v5) |

### Slots tocados

**Nenhum.** A análise foi feita contra o `deck.md` pós-v5.
