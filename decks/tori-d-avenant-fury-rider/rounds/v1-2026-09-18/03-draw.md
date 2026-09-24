# Vantagem de Cartas — Otharri, Suns' Glory (v1 · 2026-09-19 · modo `improve`)

Fonte de todo texto oracle: `bin/mtgdb` (dump Scryfall), puxado nesta sessão (regra 6).
Rulings de `bin/mtgdb rulings`. Preços em US$ são **estimativa (Scryfall)** e **não valem como
régua** — a régua é o menor valor da LigaMagic (regra 2), captura manual pendente.
Cotações em R$ marcadas são LigaMagic (menor) com a data da observação.

**Fontes reais já no deck: 3/12.**

---

## 0. A meta, argumentada — por que **12** e não 13, e por que a *forma* importa mais que o número

A régua do pipeline é 12–13. Para este deck eu proponho **12**, e sustento que a distribuição
importa mais que o total. Três fatos do eixo Otharri mudam a conta:

1. **Otharri fabrica board sem gastar carta.** Ataque nº 4 = 10 fichas 2/2 acumuladas, todas
   pagas com **zero cartas da mão**. Num deck normal, boa parte do saque existe para **repor
   presença de mesa**; aqui, presença de mesa é de graça. O saque deste deck existe para achar
   **interação, terreno e anthem** — funções que ficha nenhuma cobre. Isso **reduz a demanda**.
2. **O recurso escasso é mana, não slot.** O comandante custa 5 e quer atacar todo turno; o deck
   está **−9 em ramp**. Fonte de saque que cobra `{2},{T}` por carta compete diretamente com
   conjurar a ameaça. Por isso a forma correta aqui é **gatilho passivo ligado ao que o deck já
   faz de graça** (atacar, criar ficha), não *mana sink*. Isso **muda o tipo**, não o total.
3. **Dor 2 é de fluxo, não de volume.** "Esvazia a mão cedo e fica em topdeck" se resolve com
   **1–2 cartas por turno, todo turno**, e não com um *burst* de 3 que volta ao topdeck no turno
   seguinte. Motor recorrente > efeito único, sempre, neste deck.

**Forma alvo para as 12 fontes:**

| Critério | Mínimo | Por quê |
|---|---|---|
| Permanentes **não-criatura** (artefato/encantamento) | **6** | dor 4 — sobrevivem ao wipe e continuam sacando enquanto o Otharri reconstrói com os contadores de experiência |
| Disparam **sozinhas** do ataque/ficha (custo marginal ≤ `{1}`) | **5** | mana é o gargalo (ponto 2) |
| Efeito único / impulse de 1 uso | **≤ 3** | não resolvem topdeck |
| mv 5+ | **≤ 1** | a curva já trava em 4–5 (15 cartas em mv 4, 9 em mv 5) |

A proposta abaixo entrega **7 permanentes não-criatura**, **6 de gatilho passivo**, **1 efeito
único** e **1 em mv 5+**. Se o orçamento apertar na fase de manabase, a 12ª fonte
(`Light Up the Stage`) é a primeira a sair sem quebrar a forma.

---

## 1. O que já existe — auditoria das fontes atuais (modo `improve`)

Critério: **acesso a cartas extras**. Loot/rummage (compra + descarte forçado) **não conta** —
o deck não tem tema de cemitério, então o descarte é perda líquida. Cantrip isolado (troca 1-por-1)
**não conta**.

| Carta | Oracle (resumido, puxado nesta sessão) | Conta? | Leitura |
|---|---|---|---|
| **Belladonna Took** `{1}{W}` | *Whenever a token you control enters* — 1ª vez no turno ganha 1 vida, **2ª compra uma carta**, 3ª põe +1/+1 em cada criatura | **SIM — forte** | Otharri cria **N fichas num único ataque**, e cada ficha é um gatilho separado. Do **2º ataque em diante** (2 fichas) ela **compra todo turno**; do 3º em diante compra **e** vira anthem permanente. É a melhor fonte que o deck já tem e não custa nada |
| **Syr Carah, the Bold** `{3}{R}{R}` | quando ela **ou um instante/feitiço seu** causa dano a um jogador, exila o topo e você pode jogá-lo no turno; `{T}`: 1 de dano a qualquer alvo | **SIM — condicional** | impulse repetível **todo turno** via o `{T}`, mesmo sem atacar. Atrito: mv 5 na faixa congestionada, e o `{T}` **compete com atacar** (ela é um 3/3 que recebe todos os anthems). Fica, mas não é resposta à dor 2 |
| **Recruitment Officer** `{W}` | `{3}{W}`: olhe as 4 do topo, revele e compre **uma criatura mv ≤ 3** | **SIM — fraca** | vantagem real (carta da biblioteca para a mão) e repetível, mas `{3}{W}` por uso é caríssimo num deck sem ramp, e só acha **criatura** — que é exatamente o recurso que o comandante dá de graça. Corpo 2/1 mv 1 justifica o slot; o saque não |
| **Esgaroth Garrison** `{4}{W}` | ETB **recruit** = *draw 1, discard 1*, + ficha se descartar não-terreno | **NÃO** | **loot**, não saque. O corpo (poder = nº de criaturas → 7/5–11/5 com o enxame) é função **temática**, não minha — ver §5, corte condicionado |
| **Celebrate the Mountain-king** `{3}{W}` | ETB exila 1 não-terreno **por oponente** + ETB **recruit** | **NÃO** (como draw) | é a melhor **remoção** do deck; o recruit é loot de brinde. Permanece pelo eixo de interação |
| **Elite Interceptor // Rejoinder** `{W}` | entra *prepared*; pode conjurar cópia de `Rejoinder` (`{1}{W}`: tapa/destapa criatura + **compra 1**) | **NÃO** | você paga `{1}{W}` para comprar 1 — troca de mana por carta, não vantagem. O valor real é o *destapar* (destrava a recursão Rebel do Otharri) e o corpo mv 1. Não é fonte |
| **Crash Through** · **Warlord's Fury** · **Shoulder to Shoulder** | cantrips (efeito + *draw a card*) | **NÃO** | trocam 1-por-1. `Crash Through` fica por outro motivo (trample em massa, ↑↑ no re-pass temático) |
| **Djeru's Renunciation** `{1}{W}` | tapa até 2 criaturas; **cycling `{W}`** | **NÃO** | ciclar é trocar a carta por outra — neutro |
| **Lake-town Lookout** `{W}` | ao morrer, **recruit** (loot) | **NÃO** | loot condicionado à própria morte. Já é corte limpo K14 do re-pass temático |

**Contagem: 3 fontes reais** (1 forte, 1 condicional, 1 fraca). **Gap = −9**, exatamente como o
briefing declarou. Nenhuma das três é permanente não-criatura → **o deck hoje tem 0 fontes de
saque que sobrevivem a um board wipe**, o que amarra a dor 2 na dor 4.

---

## 2. Varredura da caixa — 113 sobressalentes (regra 7), **antes** de qualquer busca

Rodei `bin/mtgdb collection -list` e cruzei com `bin/mtgdb oracle`. **Confirmo o achado do
orquestrador: não há nenhuma fonte de saque recorrente dentro de RW na caixa.** As únicas
candidatas de custo zero são **artefatos incolores**. Dispensas por escrito abaixo.

| Carta da caixa | Oracle real | Veredito e **por que não cobre a função** |
|---|---|---|
| **Reconnaissance Mission** `{2}{U}{U}` · **Bident of Thassa** `{2}{U}{U}` | saque por criatura que conecta | **Fora da identidade** — azuis. Não são elegíveis (regra 10), não é dispensa de mérito |
| Thirst for Knowledge · Experimental Augury · Master's Councillors · Diversion Unit · Thranduil's Decree · Stoic Rebuttal · Negate | — | azuis; mesma razão |
| Courage in Crisis · Troll Negotiations · Beorn's Hospitality · Smell Fear · Wilderland Scrounger · Ulvenwald Mysteries | — | verdes; mesma razão |
| **Oracle's Vault** `{4}` | `{2},{T}`: exila o topo, pode jogá-lo no turno, põe contador de tijolo; com **3 contadores**, `{T}` exila e joga **sem pagar o custo** | **Engine real e de custo zero — mantida como reserva forte (R1).** Dispensada da titularidade por **custo de mana e curva**: 4 para entrar + `{2}` por ativação = 6 manas para a primeira carta, num deck com ramp −9 cujo turno 5 já está comprometido com o comandante. `Outpost Siege` faz impulse **automático no upkeep**, sem gastar mana, por 1 a menos e **R$ 0,90** (cotado). A troca custa noventa centavos e devolve ~4 manas por turno |
| **Magnifying Glass** `{3}` | `{T}`: adiciona `{C}`; `{4},{T}`: investigar (Clue: `{2}`, sac: compre 1) | **Dispensada por custo/curva.** `{4}` + `{2}` = **6 manas por carta**, e o `{T}` do saque compete com o `{T}` do ramp — a mesma carta não pode acelerar e sacar no mesmo turno. **Devolvo ao `ramp-specialist`**: como rocha mv 3 com *mana sink* tardio ela tem mérito; como fonte de saque, não |
| **Culling Dais** `{2}` | `{T}`, sacrifique uma criatura: contador; `{1}`, sac o artefato: compre 1 por contador | **Dispensada por anti-sinergia.** As fichas **são o dano** deste deck — sacrificá-las é converter vitória em carta. Pior: o `{T}` limita a **1 sacrifício por turno**, então o *burst* só chega no turno 4 de acumulação, e chega **quebrando o board**. 1 ponto de sinergia, reprova na regra 3 |
| **Bargaining Table** `{5}` | `{X},{T}`: compre 1, onde X = cartas na mão de **um oponente** | **Dispensada por custo e por inversão de incentivo.** Custa 5 para entrar e cobra caro exatamente quando é útil (mesão com oponentes de mão cheia); fica barata quando os oponentes já esvaziaram a mão — ou seja, quando você já está ganhando. 1 ponto |
| **Seer's Lantern** · **Prophetic Prism** · **Mycosynth Wellspring** · **Panic Spellbomb** · **Sunset Pyramid**(n/a) | filtro/cantrip/scry | **Não são fontes** — scry e cantrip não dão acesso a carta extra. `Prophetic Prism`/`Seer's Lantern` são assunto do `ramp-specialist` |
| **Thrill of Possibility** · **Seize Opportunity** | rummage (descarta, depois compra) | **Rummage puro — excluído por definição.** O deck **não tem tema de cemitério**: `Late to Dinner` e `Remember the Fallen` recuperam do cemitério, mas escolher o que descartar não é o gargalo. Troca 1-por-1 com perda de escolha |
| **Mad Ratter** `{3}{R}` | *Whenever you draw your **second** card each turn, create two 1/1 Rats* | **Não é fonte de saque** (é *token-maker* que lê saque). Nota de contexto: com `Idol of Oblivion` + `Tome of Legends` no deck, você compra 2+ por turno rotineiramente e ele viraria 2 fichas/turno — mas é um **1/2 por 4 manas** na pior faixa da curva. Devolvo como curiosidade, **não recomendo** |
| **Redcap Thief** `{2}{R}` | ETB: cria um Tesouro | ramp, não saque — `ramp-specialist` |
| **Bogardan Lancer** · **Youthful Knight** · **Embereth Paladin** · **Ori** · demais brancas/vermelhas da caixa | — | nenhuma tem linha de saque. Já dispensadas no re-pass temático §4.9 por outros eixos |

**Conclusão da varredura:** a caixa contribui com **1 reserva** (`Oracle's Vault`) e **zero
titulares**. As 9 entradas propostas são **compras**, e a justificativa escrita de cada dispensa
está na tabela acima (regra 7). Não é modo restrito (briefing), então comprar é a decisão correta.

---

## 3. Candidatas recomendadas — 9 entradas, ordenadas por custo-benefício

Todas com 2+ pontos de sinergia explicitados (regra 3). **Rótulo de mecânica**: `saque` = carta
para a mão · `impulse` = exila do topo e permite jogar por tempo limitado (**você não pode
segurar a carta**).

### 3.1 Inegociáveis (5) — resolvem a dor 2 e sobrevivem à dor 4

| # | Carta | CMC | Tipo | Como gera vantagem | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| **D1** | **Tome of Legends** | 2 | Artifact — Book | **saque.** Entra com 1 contador de página; **+1 contador toda vez que seu comandante entra ou ataca**; `{1},{T}`, remova: compre 1 | (a) Otharri **ataca todo turno** por desenho — o contador é automático e **não custa mana**; (b) a recursão dele (`{2}{R}{W}` + virar Rebel) o faz **entrar de novo**, gerando contador extra a cada volta; (c) **artefato, não-criatura** → o wipe que mata o enxame não o toca; (d) mv 2 = entra antes do comandante | **não** — nada equivalente na caixa (`Oracle's Vault` custa 4 + `{2}`/uso) | US$ 0,29 *(estimativa Scryfall)* |
| **D2** | **Idol of Oblivion** | 2 | Artifact | **saque.** `{T}`: compre 1. **Só pode ativar se você criou uma ficha neste turno** | (a) a condição é **o texto do comandante** — todo ataque do Otharri cria ficha, então é **saque grátis (`{T}` apenas, zero mana) todo turno**, a única fonte de custo marginal **zero** do deck; (b) `Raise the Alarm`, `Ancestral Blade`, `Knight Luminary`, `Legion's Landing` ligam a condição **fora** do combate, inclusive pós-wipe; (c) artefato não-criatura → resiste ao wipe; (d) `{8}`: ficha 10/10 como *mana sink* de jogo longo | **não** | **R$ 8,94** (LigaMagic menor, 2026-08-22 — 28 dias) |
| **D3** | **Wedding Announcement** | 3 | Enchantment (transforma) | **saque.** No seu end step: contador de convite; **se você atacou com 2+ criaturas, compre 1**; senão, crie uma ficha 1/1. Com 3 contadores **transforma em `Wedding Festivity`: criaturas que você controla ganham +1/+1** | (a) **os dois modos servem** — no ataque nº 1 (Otharri sozinho) ela dá **ficha**, que alimenta Belladonna/Idol/Valor in Akros; do nº 2 em diante dá **carta**; (b) vira **anthem estático permanente**, que o re-pass temático mostrou valer N pontos de dano (+10 no 4º ataque); (c) encantamento → sobrevive ao wipe **e** continua fazendo ficha para reconstruir. **Duas funções do deck num slot** | **não** | US$ 0,41 *(estimativa)* |
| **D4** | **Palace Jailer** | 4 | Creature — Human Soldier 2/2 | **saque.** ETB: **você se torna o monarca** (compre 1 no fim de cada turno seu, para sempre) + ETB: **exile uma criatura de um oponente** até que um oponente vire monarca | (a) **o monarca fica no jogador**, igual aos contadores de experiência — o wipe mata o Jailer e **você continua comprando**; é a segunda fonte de vantagem imune a wipe que RW tem; (b) o deck é o **melhor defensor de coroa da mesa**: board mais largo + Otharri **voador com vínculo** bloqueando; (c) **atende a dor 3** no mesmo slot (remoção); (d) Human Soldier recebe todos os anthems. **Atrito declarado:** as fichas criadas entram **viradas** e não bloqueiam no turno seguinte ao ataque — a defesa da coroa depende das fichas **antigas** (que desviram) e sobe muito com vigilância em massa (`Radiant Destiny`/`Intangible Virtue`, §4.2 do temático). Se a coroa passar, a criatura exilada **volta** | **não** | US$ 0,25 *(estimativa)* |
| **D5** | **Tocasia's Welcome** | 3 | Enchantment | **saque.** *Whenever one or more creatures you control with **mana value 3 or less** enter, draw a card* (1×/turno) | (a) **ficha tem mana value 0** — a condição é satisfeita por definição, todo ataque; (b) **imune ao atrito que mata os concorrentes**: `Mentor of the Meek` e `Welcoming Vampire` leem **poder ≤ 2**, e qualquer anthem estático do deck (`The Circle of Loyalty`, `Jor Kadeen` +3/+0, `Radiant Destiny`, `Warleader's Call`) transforma as fichas em poder 3+ e **desliga os dois**. Mana value **não muda com anthem**; (c) encantamento → resiste ao wipe; (d) também dispara com as criaturas mv ≤ 3 conjuradas | **não** | US$ 0,81 *(estimativa)* |

### 3.2 Alto valor (4) — fecham as 12

| # | Carta | CMC | Tipo | Como gera vantagem | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| **D6** | **Staff of the Storyteller** | 2 | Artifact | **saque.** ETB: ficha 1/1 voadora; **contador de história toda vez que você cria uma ou mais fichas de criatura**; `{W},{T}`, remova: compre 1 | (a) o gatilho é **criar ficha**, não atacar — funciona no combate **e** fora dele (`Raise the Alarm` no turno do oponente); (b) o ETB **já é uma ficha**, que dispara `Belladonna Took`, `Valor in Akros` e o próprio `Idol of Oblivion` no turno em que entra; (c) artefato não-criatura → resiste ao wipe; (d) mv 2 com saque a `{W}` — o custo marginal mais barato da lista depois do Idol | **não** | US$ 0,50 *(estimativa, edição SOC)* |
| **D7** | **Outpost Siege** *(modo Khans)* | 4 | Enchantment | **impulse recorrente.** No seu upkeep, exile o topo; **você pode jogá-la neste turno** | (a) **custo marginal zero** — dispara sozinha, não gasta mana nem `{T}`, que é exatamente a forma que o §0 exige; (b) encantamento → **sobrevive ao wipe e continua cavando** enquanto o Otharri volta pelos contadores de experiência; (c) **R$ 0,90 cotado** é o melhor preço/efeito do deck inteiro. **Custo do rótulo impulse:** a carta exilada **é perdida se você não a jogar no turno** — terreno e feitiço barato aproveitam bem; `Duty Beyond Death` e as proteções de instante ficam desconfortáveis (você precisa gastá-las no seu turno) | **não** | **R$ 0,90** (LigaMagic menor, 2026-08-22 — 28 dias) |
| **D8** | **Neyali, Suns' Vanguard** | 4 | Leg. Creature — Human **Rebel** 3/3 | **impulse recorrente.** *Attacking tokens you control have **double strike*** (estática) · *Whenever one or more tokens you control **attack a player**, exile o topo; você pode jogá-la em **qualquer turno em que atacou com ficha*** | (a) **correção de mecânica importante:** o gatilho pede fichas **declaradas atacantes** — as fichas que o Otharri põe *entrando atacando* **nunca foram declaradas** (ruling Basri Ket 2020-06-23). Quem dispara a Neyali são as fichas **dos ataques anteriores**, que desviram e são declaradas normalmente. Resultado prático: **do 2º ataque em diante, dispara todo turno**; num mesão, o ruling 2023-02-04 confirma **um gatilho por jogador atacado com ficha** — até 3 cartas por combate; (b) a estática **dobra o dano do enxame inteiro** (6 fichas 2/2 = 24 em vez de 12); (c) **é Rebel não-ficha** → destrava a recursão `{2}{R}{W}` do Otharri mesmo depois de um wipe que leve as fichas; (d) ruling: o direito de jogar a carta exilada **persiste mesmo se a Neyali sair de campo**. Já é a candidata nº 1 do pool temático — **1 slot, 3 funções** | **não** | US$ 2,69 *(estimativa)* |
| **D9** | **Light Up the Stage** | 3 *(spectacle `{R}`)* | Sorcery | **impulse, efeito único.** Exile as **2** do topo; você pode jogá-las **até o fim do seu próximo turno** | (a) *spectacle* está **ligado por padrão** — o deck causa perda de vida todo turno desde o ataque nº 1, então o custo real é **`{R}` por 2 cartas**, o melhor preço/carta disponível a RW; (b) a janela de **dois turnos** (diferente de `Reckless Impulse` em corpo, mesma janela) torna o impulse quase saque para terrenos e peças baratas; (c) **R$ 0,74 cotado**. É a **12ª fonte e a primeira a sair** se a manabase precisar do slot — é o único efeito único da lista | **não** | **R$ 0,74** (LigaMagic menor, 2026-08-22 — 28 dias) |

**Total das 9:** 3 com cotação LigaMagic somam **R$ 10,58**; as outras 6 somam **US$ 4,95 em
estimativa (Scryfall)**. **Não afirmo que cabem no teto** — as 6 precisam de captura manual na
LigaMagic antes de qualquer total (regra 2). Erros medidos do proxy chegaram a 6,5× para mais.

---

## 4. Reservas e swaps

### 4.1 Reservas (4) — entram se alguma titular estourar a cotação real

| # | Carta | CMC | Tipo | Como gera vantagem | Sinergias | Origem | Preço |
|---|---|---|---|---|---|---|---|
| **R1** | **Oracle's Vault** | 4 | Artifact | **impulse recorrente.** `{2},{T}`: exile o topo, jogue neste turno + contador; com **3 contadores**, `{T}`: exile e jogue **de graça** | (a) **custo zero de compra — já está na caixa**; (b) artefato → resiste ao wipe; (c) o modo grátis ignora a curva travada em 4–5. Atrito: 6 manas até a 1ª carta | **caixa** | **R$ 0,00** |
| **R2** | **Case of the Crimson Pulse** | 3 | Enchantment — Case | **saque.** ETB: descarte 1, **compre 2**. *Solve*: estar sem cartas na mão. **Resolvido: no seu upkeep, descarte a mão e compre 2** | (a) **é a resposta literal à dor 2** — o deck esvazia a mão por desenho, então o *solve* é automático e vira **2 cartas por turno, para sempre**, por 3 manas; (b) encantamento → resiste ao wipe. **Atrito grave e declarado:** descartar a mão no upkeep **proíbe segurar resposta de instante** — briga de frente com `Gods Willing`/`Feat of Resistance`/`Duty Beyond Death`, que são o plano da dor 4. Por isso é reserva, não titular | compra | US$ 0,24 *(est.)* |
| **R3** | **Mask of Memory** | 2 | Artifact — Equipment | **saque líquido +1.** Quando a criatura equipada causa dano de combate a um jogador: **compre 2, descarte 1** | (a) o portador natural é o **Otharri — voador com ímpeto**, que conecta no turno em que entra e quase sempre depois; (b) **é artefato** → repõe a contagem de metalcraft do `Jor Kadeen` que os cortes K3/K15 do temático colocam em risco; (c) equip `{1}`. **Rótulo honesto:** tem componente de loot, mas o saldo é **+1 carta por conexão** — é vantagem real, não rummage | compra | US$ 0,41 *(est.)* |
| **R4** | **Palace Sentinels** | 4 | Creature — Human Soldier 2/4 | **saque.** ETB: **você se torna o monarca** | (a) segunda fonte de coroa — se um oponente tomar a coroa do `Palace Jailer`, ela **volta** por 4 manas; (b) corpo **2/4** bloqueia melhor que o Jailer e **defende a própria coroa**; (c) US$ 0,17. Só entra **junto** com D4, como pacote de monarca | compra | US$ 0,17 *(est.)* |

### 4.2 Swaps propostos (modo `improve`)

Cada corte com ficha F1–F7 na §5. Os slots das 9 entradas saem **majoritariamente dos cortes já
protocolados pelo `theme-analyst`** (K1, K2, K7, K8, K10, K11, K12, K13, K16, K17 — 10 cortes
limpos); listo aqui só os que pertencem à **minha** lente.

| Sai | Entra | Justificativa |
|---|---|---|
| **Shoulder to Shoulder** *(K9 do temático — endosso)* | **Tome of Legends** | K9 declarou o cantrip como "função descoberta, custo aceito — é exatamente o slot que a fase de draw troca por saque real". Cumprido: mesmo slot mv 2–3, saque **recorrente** no lugar de 1-por-1 |
| **Lake-town Lookout** *(K14 — endosso)* | **Idol of Oblivion** | K14 deixou o *loot* descoberto. Trocado por saque real e **grátis** por turno |
| **Warlord's Fury** | **Tocasia's Welcome** | ficha **D1** na §5 — cantrip cujo efeito (first strike em massa) morreu junto com `Kwende` |
| **Djeru's Renunciation** | **Staff of the Storyteller** | ficha **D2** na §5 |
| **Esgaroth Garrison** | *(a definir)* | **corte condicionado — não corto.** Ficha **D3** na §5: o *recruit* é loot e não é fonte, mas o **corpo** é função temática. Devolvo ao orquestrador |

---

## 5. Fichas F1–F7 dos cortes que eu proponho (regra 4 · protocolo §2 do checklist)

**Consulta ao `decisions.md` (regra 5):** nenhuma das três tem histórico de corte — o deck foi
montado fora do pipeline e esta é a primeira rodada. Nenhuma é reposição de carta cortada antes.

Condições assumidas **idênticas** para quem sai e quem entra (§3 do checklist): Otharri em campo
com 3–4 contadores de experiência, 6–10 fichas Rebel no board, um anthem estático em jogo.

### D1 · Warlord's Fury — `{R}` Sorcery — **corte limpo**
- **F1** `Creatures you control gain first strike until end of turn.` / `Draw a card.` — duas linhas, ambas lidas nesta sessão.
- **F2** Sem corpo. Não tapa, não é sacrificável, não bloqueia.
- **F3** Feitiço — o deck não tem contagem de instantâneo/feitiço (`Syr Carah` premia **dano** de instante/feitiço, e esta carta não causa dano).
- **F4** Não recebe nada — anthems, contadores e equipamentos não a alcançam.
- **F5** Concede **first strike ao time inteiro** por 1 mana. Sob o eixo antigo isso ligava com `Kwende, Pride of Femeref` (first strike → **double strike**) e com `Squire's Lightblade`.
- **F6** Turno 1. Não destrava nada a partir dali — é efeito de um turno.
- **F7** **Atrito fatal:** o pagamento do first strike **sai do deck**. `Kwende` é corte condicionado K4, `Squire's Lightblade` é K15, `Cavalry Drillmaster` é K13. Sem eles, first strike em massa num enxame de 2/2 é quase letra morta: 2/2 com first strike ainda morre para qualquer bloqueador 2/2, e o que o enxame precisa é **trample/menace** (§4.7 do temático), não iniciativa.
- **Protocolo de corte:** funções = [cantrip, first strike em massa, custo 1 mana].
  - *cantrip* → **substituída e ampliada** por `Tocasia's Welcome` (saque recorrente, todo turno, no lugar de uma carta só).
  - *first strike em massa* → **coberta e superada** por `Neyali, Suns' Vanguard` (**double strike** estático em todas as fichas atacantes — estritamente melhor e permanente).
  - *1 mana* → `Crash Through`, que **fica** (↑↑ no temático), ocupa o mesmo espaço de curva com trample, que é o keyword de que o enxame precisa.
  - **Nenhuma função descoberta.** Corte limpo.

### D2 · Djeru's Renunciation — `{1}{W}` Instant — **corte limpo com custo declarado**
- **F1** `Tap up to two target creatures.` / `Cycling {W}`.
- **F2** Sem corpo.
- **F3** Instante — sem contagem relevante no deck.
- **F4** Não recebe nada.
- **F5** Tapa **até dois** bloqueadores (Falter parcial) **ou** vira outra carta por `{W}`.
- **F6** Turno 2; em instante, pode ser guardada.
- **F7** O enxame resolve bloqueio **por volume** (10 fichas contra 2–3 bloqueadores), não removendo 2 bloqueadores. E ciclar por `{W}` é **troca 1-por-1**: não repõe, só filtra.
- **Protocolo de corte:** funções = [Falter parcial, ciclagem].
  - *Falter parcial* → o `theme-analyst` já declarou essa função dispensável no corte **K16** (`Bond of Discipline`), com o custo aceito por escrito: a evasão passa a vir de **trample/menace permanentes** (`Crash Through`, `Hexgold Halberd`, `Iroas`) e de **combate extra**, que produzem mais dano por mana. Aplico o mesmo critério aos dois lados (§3 do checklist).
  - *ciclagem* → **substituída** por `Staff of the Storyteller`, que dá carta **toda vez que fichas entram**, sem gastar a própria carta.
  - **Custo declarado:** o deck perde a única forma de tapar bloqueadores em **instante**. Aceito — mas registro que, se o `interaction-specialist` quiser a carta como truque de combate, o slot é dele para reclamar, não meu para fechar.

### D3 · Esgaroth Garrison — `{4}{W}` Creature — Human Soldier */5 — **CORTE CONDICIONADO · não corto**
- **F1** `Esgaroth Garrison's power is equal to the number of creatures you control.` / `When this creature enters, **recruit**` (*draw a card, then discard a card; se descartou não-terreno, crie uma ficha 1/1 Human Soldier*).
- **F2** **Corpo real e grande:** resistência 5 fixa e poder = nº de criaturas → com 6–10 fichas é um **7/5 a 11/5**; com `The Circle of Loyalty` + `Jor Kadeen` passa de 11/6. Bloqueia muito bem. Não tapa para custo (o deck não tem crew/convoke).
- **F3** Human Soldier — alimenta contagens de Human (que estão encolhendo) e o affinity do `The Circle of Loyalty` só por ser criatura, não por tipo.
- **F4** **Recebe tudo:** anthems estáticos, `Basri's Solidarity`, `Duty Beyond Death`, `Valor in Akros`, equipamentos, proteções.
- **F5** O *recruit* pode gerar **uma ficha**, que dispara `Belladonna Took`/`Valor in Akros`/`Idol of Oblivion`.
- **F6** Turno 5 — a mesma faixa do comandante, que é o pior gargalo da curva (9 cartas em mv 5).
- **F7** Atrito de curva com o Otharri; e o *recruit* é **loot**, que este deck não converte em valor (sem tema de cemitério).
- **Protocolo:** funções = [**corpo grande escalonado**, resistência 5, receptor de anthem, loot, ficha ocasional].
  - *loot* → coberto e superado por **qualquer** das 9 entradas.
  - *ficha ocasional* → coberta pelo motor do comandante.
  - **corpo grande escalonado + bloqueador 5 de resistência → FICA DESCOBERTO se eu cortar.** Isso é função **temática/de combate**, fora da minha especialidade.
  - **Não corto.** Devolvo ao orquestrador nomeando a função descoberta: *o deck perderia seu maior corpo defensivo, num eixo em que as fichas recém-criadas entram viradas e não bloqueiam.* Se o `theme-analyst`/`manabase-engineer` precisar do slot de mv 5, o corte é defensável — **com essa perda declarada**, não por "é loot".

---

## 6. Rejeitadas com justificativa (para não voltarem sem contexto novo — regra 5)

| Carta | Preço est. | Por que **não** entra |
|---|---|---|
| **Mentor of the Meek** · **Welcoming Vampire** | US$ 4,76 · US$ 3,09 | Leem **poder ≤ 2**. Qualquer anthem estático do plano (`The Circle of Loyalty` +1/+1, `Jor Kadeen` +3/+0, `Radiant Destiny`, `Warleader's Call`, `Wedding Festivity`) faz as fichas **entrarem com poder 3+** e **desliga a carta**. É anti-sinergia com o §4.2 inteiro do pool temático. `Tocasia's Welcome` faz o mesmo lendo **mana value**, que anthem não altera — por **1/6 do preço** |
| **Caretaker's Talent** | US$ 11,99 | Funcionalmente perfeita (saque por evento de ficha + nível 3 dá **+2/+2 às fichas**). **Rejeitada só por preço**: sozinha custaria mais que D1+D3+D4+D5+D6 somadas. **Luxo** — reproponha se o orçamento sobrar depois da manabase |
| **Bennie Bracks, Zoologist** | US$ 11,43 | Mesmo efeito de D3/D6 (saque se criou ficha, em **cada** end step) com **convoke** pago pelo enxame. **Rejeitada por preço** e por ser **criatura** (morre no wipe, contra a forma do §0). **Luxo** |
| **Chivalric Alliance** | US$ 15,06 | "Ataque com 2+ criaturas → compre 1" é exatamente o deck. **Rejeitada por preço** — `Wedding Announcement` faz a mesma leitura por **1/36 do custo** e ainda vira anthem |
| **Battle Angels of Tyr** | US$ 29,79 | Myriad + saque condicional; preço proibitivo na régua de R$ 200 |
| **Smuggler's Share** · **Trouble in Pairs** | US$ 11,16 · ~US$ 10 | Saque **reativo ao que o oponente faz** — não é ligado ao ataque nem à ficha, e o deck quer fluxo garantido. Preço alto |
| **Esper Sentinel** | US$ 59,91 | 30% do teto num efeito que os oponentes **escolhem** pagar. Fora |
| **Skullclamp** | **R$ 26,89** (LM, 28 d) | O reflexo "deck de fichas → Skullclamp" **falha aqui**: as fichas são **2/2**, e +1/−1 as deixa **3/1 vivas**. Só funcionaria com fichas X/1. Preço alto por zero sinergia |
| **Reckless Impulse** | **R$ 9,78** (LM, 28 d) | Mesmo efeito de `Light Up the Stage`, **13× o preço** com spectacle ligado. Descartada por custo |
| **Magus of the Wheel** · **Reforge the Soul** · **Wheel of Misfortune** | US$ 1,38 · 6,47 · 7,58 | *Wheel* refila **sua** mão (ótimo para a dor 2), mas entrega **7 cartas a cada oponente** — num deck que já **morre para board wipe** (dor 4), você está comprando os wipes deles. Anti-sinergia com a dor mais grave. Não recomendo sem pedido explícito |
| **Endless Atlas** | US$ 2,02 | Sempre ligada (15 Mountain + 15 Plains), mas `{2},{T}` por carta é **mana sink puro** — a forma que o §0 rejeita. Fica atrás de D1/D2/D6 no mesmo slot de artefato mv 2 |
| **Mazemind Tome** · **Arcane Encyclopedia** · **Mind's Eye** | US$ ~0,3 a 5,68 | Saque genérico desconectado do tema: nenhum lê ficha, ataque ou comandante. **Menos de 2 pontos de sinergia** — reprovam na regra 3 |
| **Reckoner Bankbuster** | **R$ 1,90** (LM, 38 d) | Boa carta (3 cartas + corpo 4/4 + Tesouro + **veículo não é criatura**, escapa do wipe). **Não conta como fonte recorrente** — são 3 usos e acaba. Fica como **candidata do `ramp-specialist`/orçamento**, não da minha cota |
| **Inti, Seneschal of the Sun** *(pool temático #36)* | US$ 2,04 | **Não conto como fonte de saque**: descartar 1 para exilar 1 do topo é **conversão**, não vantagem — o número de cartas não sobe. Excelente como **tema** (contador permanente + trample no ataque) e como **válvula** para as cartas mortas que `Outpost Siege` deixar passar. Recomendação de tema, não de draw |
| **Dawn of Hope** | US$ 0,55 | Otharri tem **vínculo com a vida** → 1 gatilho por combate; `{2}` por carta. Honesta e barata, mas **1 carta/turno por 2 manas** perde para D2 (grátis) e D6 (`{W}`). **Reserva de 2ª linha** |
| **Throne of the High City** | US$ 0,22 | **Monarca dentro de um slot de terreno** — e o deck está **−2 na base de 38**. Não gasto slot de feitiço com ela: **devolvo ao `manabase-engineer`** como terreno que também é fonte de saque |
| **Bonders' Enclave** | US$ 0,93 | `{2},{T}`: compre 1 se controlar criatura de **poder 4+**. Com `Jor Kadeen`/`The Circle of Loyalty` é trivial; sem eles, não liga. Também é **slot de terreno** → `manabase-engineer` |
| **Forth Eorlingas!** · **Oath of Eorl** · **Court of Embereth** · **Court of Grace** | US$ 3,51 · 0,32 · 5,44 · 2,98 | Todas dão **monarca** e fichas. Boas, mas **D4 (`Palace Jailer`) faz monarca por US$ 0,25 e ainda remove uma criatura**. Devolvo `Oath of Eorl` e `Forth Eorlingas!` ao `theme-analyst` como **token-makers com monarca de brinde**, não como fonte de draw |

---

## 7. Placar final da categoria

| Fonte | mv | Tipo | Mecânica | Origem | Sobrevive a wipe? | Custo marginal |
|---|---|---|---|---|---|---|
| Belladonna Took | 2 | Creature | saque | **deck** | não | zero |
| Recruitment Officer | 1 | Creature | saque | **deck** | não | `{3}{W}` |
| Syr Carah, the Bold | 5 | Creature | impulse | **deck** | não | `{T}` |
| **Tome of Legends** (D1) | 2 | Artifact | saque | compra | **sim** | `{1}`+`{T}` |
| **Idol of Oblivion** (D2) | 2 | Artifact | saque | compra | **sim** | **zero** (só `{T}`) |
| **Wedding Announcement** (D3) | 3 | Enchantment | saque | compra | **sim** | zero |
| **Palace Jailer** (D4) | 4 | Creature | saque (monarca) | compra | **sim** (a coroa fica) | zero |
| **Tocasia's Welcome** (D5) | 3 | Enchantment | saque | compra | **sim** | zero |
| **Staff of the Storyteller** (D6) | 2 | Artifact | saque | compra | **sim** | `{W}`+`{T}` |
| **Outpost Siege** (D7) | 4 | Enchantment | impulse | compra | **sim** | zero |
| **Neyali, Suns' Vanguard** (D8) | 4 | Creature | impulse | compra | não | zero |
| **Light Up the Stage** (D9) | 3 *(`{R}`)* | Sorcery | impulse, único | compra | — | — |

**Total: 12 fontes** · permanentes não-criatura: **7/12** (meta ≥6 ✅) · gatilho passivo de custo
marginal ≤`{1}`: **8/12** (meta ≥5 ✅) · efeito único: **1** (meta ≤3 ✅) · mv 5+: **1** (meta ≤1 ✅).
**Rótulo de mecânica:** 8 são **saque** (carta para a mão, pode ser segurada) e **4 são impulse**
(D7, D8, D9 e `Syr Carah`) — nestas, a carta exilada **se perde se você não a jogar na janela**,
o que penaliza justamente as respostas de instante que a dor 4 pede. Por isso mantive o impulse
em **1/3 do total** e concentrei o saque verdadeiro nas inegociáveis.

## Curva das fontes de draw

**mv 1–2: 5** (Recruitment Officer 1 · Belladonna Took 2 · Tome of Legends 2 · Idol of Oblivion 2 · Staff of the Storyteller 2)
**mv 3–4: 6** (Wedding Announcement 3 · Tocasia's Welcome 3 · Light Up the Stage 3/`{R}` · Outpost Siege 4 · Palace Jailer 4 · Neyali 4)
**mv 5+: 1** (Syr Carah, the Bold 5)

---

## 8. Pendências que devolvo ao orquestrador

1. **Cotação LigaMagic obrigatória** para 6 das 9 entradas (D1, D3, D4, D5, D6, D8) — sem elas
   **nenhuma afirmação de orçamento é válida** (regra 2). As 3 já cotadas somam **R$ 10,58**.
2. **`Esgaroth Garrison`** — corte condicionado, função de corpo descoberta (§5 · D3).
3. **`Throne of the High City`** e **`Bonders' Enclave`** ao `manabase-engineer`: são **fontes de
   saque em slot de terreno**, e o deck está −2 na base. Resolvem duas lacunas com um slot.
4. **`Magnifying Glass`** (caixa) ao `ramp-specialist` — dispensada como draw, tem mérito como rocha.
5. **`Djeru's Renunciation`** ao `interaction-specialist`, caso queira reclamá-la como truque de combate.
6. **Pacote de monarca:** se o usuário aprovar D4, avaliar **R4 `Palace Sentinels`** (US$ 0,17) como
   seguro contra a perda da coroa — decisão de mesa, não de construção.
