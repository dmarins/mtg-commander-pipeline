# Vantagem de Cartas — Phlage, Titan of Fire's Fury (v2 · controle de atrito RW)

> **Natureza desta fase.** Não existe `deck.md`. O deck v1 (Otharri) foi reprovado e nunca
> comprado. Isto é um `build` na prática: não há fontes de draw "já no deck" para contar —
> há **um pool** (`02-theme.md`), a `lista.txt` do Tori físico e a caixa de sobressalentes.
>
> **Regra 5 — conferido.** `decisions.md` registra **três** movimentos, todos de zona de comando
> (Otharri entra e sai, Tori sai). **Nenhuma carta do 99 tem histórico de corte.** Nenhuma
> proposta abaixo é reposição de corte anterior.
>
> **Regra 6 — oracle.** Todo texto citado foi puxado com `bin/mtgdb oracle` nesta sessão,
> inclusive o ruling do [**Bargaining Table**](https://www.ligamagic.com.br/?view=cards/card&card=Bargaining+Table) (que muda o veredito dela).
>
> **Regra 2 — preço.** Nenhum número do Scryfall. `bin/mtgdb prices` consultado para as 38
> candidatas: **8 têm cotação**, o resto sai como **`a cotar`**.
>
> **Travas do escape respeitadas:** nenhuma proposta é reanimação ou blink.

**Fontes já no deck: 1/12–13** — apenas [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold), que existe fisicamente na
`lista.txt`. É esse o tamanho do buraco: a dor nº 2 do intake não é impressão do usuário, é a
composição.

---

## 1. A distinção que organiza esta fase

O orquestrador está certo e o número precisa ficar escrito: **loot não é vantagem de cartas.**
[**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) compra 2 e descarta 1 — a contagem de cartas na mão não sobe, o que sobe
é a **qualidade** da mão e a **altura do cemitério**. As 14 cartas de combustível dimensionadas
na §3 da Fase 2 resolvem o escape do Phlage; elas **não** resolvem "mão morta".

Três categorias, com fronteira explícita:

| Categoria | Definição operacional | Quem entrega |
|---|---|---|
| **A · Vantagem real** | aumenta o número de cartas a que você tem acesso além do saque do turno | **esta fase** — alvo 12–13 |
| **B · Combustível/seleção** | troca cartas por cartas, enche cemitério, filtra | **Fase 2** — 14 cartas (10 loot + 4 motores) |
| **C · Sobreposição** | faz as duas coisas, e é declarada nas duas | 5 das minhas 14 (§4.3) |

**A conta honesta: 14 (Fase 2) + 14 (esta fase) − 5 de sobreposição = 23 slots**, não 28.
Onde a sobreposição incide, eu **devolvo slot** à Fase 2 em vez de somar (§7).

### 1.1 Correção mecânica à Fase 2 — impulse **não** é anti-combustível

O `02-theme.md` (§8.2 e §11) dispensa impulse dizendo que "exila em vez de moer" e limita-o a
2 slots. Metade disso está errada, e a metade errada é load-bearing:

**Uma mágica conjurada a partir do exílio vai para o cemitério quando resolve, igual a uma
conjurada da mão.** O impulse só destrói carta quando você **não consegue jogá-la** no prazo.
Ou seja: impulse realizado = combustível normal; impulse perdido = carta destruída. Não existe
"impulse que tira do cemitério".

O que **é** verdade é a taxa de realização, e ela é o que dimensiona a categoria:

| Componente | Peso | Realiza? |
|---|---|---|
| carta exilada é terreno | 37,4% (37 terrenos em 99) | só se o land drop do turno estiver livre → ~55% |
| carta exilada é mágica | 62,6% | janela "até o fim do seu próximo turno" ~80% · janela "só neste turno" ~65–70% |

**Taxa composta ≈ 0,374 × 0,55 + 0,626 × 0,75 ≈ 0,68.** Uso **2/3** como régua: um [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege)
rende ~0,65 carta por turno, não 1; um impulse-2 ([**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity)) rende ~1,3 carta.
Num deck que segura mana em resposta a janela encolhe mais — por isso **prefiro impulse de
janela longa** ("até o fim do seu **próximo** turno") e **impulse de upkeep** (mana intacto) a
impulse de meio de turno.

Consequência prática: **impulse entra, mas descontado**. Das 14 fontes abaixo, 5 são impulse e
valem ~3,3 fontes efetivas.

### 1.2 Contrapeso ao sinal "8 a 10 das 12–13 devem ser loot"

A Fase 2 pediu isso. **Não é possível atender sem anular a fase.** Se 8–10 das 13 fontes forem
loot, o deck fica com 3–5 fontes de vantagem real e a dor 2 continua exatamente onde está — o
usuário relatou "esvazia a mão cedo e fica em topdeck", e loot não resolve topdeck, ele só
escolhe qual carta você topdecka.

O que eu faço em vez disso: **caço a sobreposição pela outra ponta** — cartas cuja vantagem real
é *paga com descarte ou com moinho*, de modo que o mesmo slot alimenta o cemitério. É o que
[**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder), [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc), [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) e
[**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) fazem. Cinco das catorze. As 14 de combustível da Fase 2 ficam
intactas como categoria B (com o pedido de redução de 2 slots da §7, que tem contrapartida).

---

## 2. Quantas fontes este eixo precisa — defesa do número

**Proponho 14 nominais, que valem ~11,5 efetivas.** Fico no topo da faixa do pipeline, e a
distribuição é o que justifica:

| Motivo | Efeito no número |
|---|---|
| **RW é o pior par em saque.** Não há Rhystic Study, não há Sylvan Library, não há tutor barato. O que existe é coroa, impulse e punisher — tudo com desconto embutido | sobe |
| **Controle gasta carta 1-por-1.** 20–22 remoções + 3–4 wipes (§5 da Fase 2) é o deck todo trocando 1:1 com uma mesa de 3. Sem motor, a mesa vence por aritmética | sobe |
| **Impulse rende 2/3** e **a coroa não acumula** (§4.2) | sobe o nominal para manter o efetivo |
| **2 das 14 são slot de terreno** ([**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City), [**War Room**](https://www.ligamagic.com.br/?view=cards/card&card=War+Room)) — não disputam slot de mágica | sobe o nominal de graça |
| **1 das 14 já existe fisicamente** ([**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold)) | custo zero |
| O eixo **não pode** virar deck de motor: 20–22 remoções é a espinha e não se mexe | trava o teto |

**Efetivas por regime (estimativa, não simulação):** com 3 motores em campo por volta do T7, o
pacote entrega **+2,5 a +3 cartas por turno** acima do saque. É o suficiente para pagar 2–3
respostas por rodada, que é a taxa que um pod de 4 exige de um deck de controle.

---

## 3. Varredura da caixa e da `lista.txt` (regra 7) — **antes** do Scryfall

`bin/mtgdb collection -list` (116 cartas) + as 69 não-básicas de `lista.txt`, filtradas por
`draw-engine`, `repeatable-pure-draw`, `burst-draw`, `long-term-impulsive-draw`.

### 3.1 Aproveitado — 3 cartas sem custo de compra

| Carta | Origem | Por que serve |
|---|---|---|
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | lista | exila o topo e deixa jogar **sempre que uma mágica sua fere um jogador**; `{T}` pinga 1 e dispara o próprio gatilho. Com 20+ queimas é a única fonte não-loot que o deck já tem |
| [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) | caixa | impulse 2 **em instantâneo**, janela "até o fim do seu **próximo** turno" — a janela longa e o instantâneo são exatamente o perfil que a §1.1 pede. Substitui a compra de [**Reckless Impulse**](https://www.ligamagic.com.br/?view=cards/card&card=Reckless+Impulse) (**R$ 9,78**), que faz o mesmo em feitiço |
| [**Bargaining Table**](https://www.ligamagic.com.br/?view=cards/card&card=Bargaining+Table) | caixa | `{X}, {T}: compre` — **o ruling de 2004-10-04 muda a carta**: você escolhe o oponente *na anúncio*, **antes** de X ser determinado. Num jogo de atrito longo sempre há um oponente com 0–2 cartas → é saque repetível quase de graça no late |

> **Discordo da dispensa de [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) feita pela Fase 2** ("[**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility)
> faz mais por 1 mana a menos"). As duas cartas não fazem a mesma coisa: `Thrill` é categoria B
> (troca 1:1 + combustível) e `Seize` é categoria A (**+2 cartas**). Comparar as duas é o erro
> que esta fase existe para não cometer. Com o buraco de draw em 1/13 e [**Reckless Impulse**](https://www.ligamagic.com.br/?view=cards/card&card=Reckless+Impulse) a
> R$ 9,78, uma fonte de vantagem real a R$ 0 na caixa não se dispensa.

### 3.2 Dispensas da caixa — ficha e motivo (regra 7)

| Carta | Ficha resumida | Por que **não** cobre a função |
|---|---|---|
| [**Bident of Thassa**](https://www.ligamagic.com.br/?view=cards/card&card=Bident+of+Thassa), [**Reconnaissance Mission**](https://www.ligamagic.com.br/?view=cards/card&card=Reconnaissance+Mission), [**Thirst for Knowledge**](https://www.ligamagic.com.br/?view=cards/card&card=Thirst+for+Knowledge), [**Experimental Augury**](https://www.ligamagic.com.br/?view=cards/card&card=Experimental+Augury), [**Thranduil's Decree**](https://www.ligamagic.com.br/?view=cards/card&card=Thranduil%27s+Decree), [**Elrond, Moon-Reader**](https://www.ligamagic.com.br/?view=cards/card&card=Elrond%2C+Moon-Reader), [**Master's Councillors**](https://www.ligamagic.com.br/?view=cards/card&card=Master%27s+Councillors), [**Cargo Ship**](https://www.ligamagic.com.br/?view=cards/card&card=Cargo+Ship), [**Diversion Unit**](https://www.ligamagic.com.br/?view=cards/card&card=Diversion+Unit), [**Old Fat Spider Can't See Me**](https://www.ligamagic.com.br/?view=cards/card&card=Old+Fat+Spider+Can%27t+See+Me) | F1 saque real · F6 mv 2–6 | **Fora da identidade RW.** Todas azuis. É o fato que `decisions.md` já registrou em 2026-09-19: *"a caixa não tem nenhuma fonte de draw dentro da identidade"*. Continua verdade — a caixa tem **3** (§3.1), e nenhuma delas é motor |
| [**Oracle's Vault**](https://www.ligamagic.com.br/?view=cards/card&card=Oracle%27s+Vault) (**R$ 0,25**) | F1 `{2},{T}`: impulse 1 + contador; com 3 contadores, `{T}`: impulse 1 **grátis** · F3 Artifact · F6 **CMC 4** · F7 usa `{T}` | **Reserva, não dispensa.** 2 pontos reais (impulse repetível + artefato). O que a barra é curva: 4 manas para começar e **3 turnos de `{2}` cada** para virar grátis, num deck cujos T2–T5 são de remoção. [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) faz o mesmo de graça a partir do turno seguinte por 4 manas. Fica como **reserva de orçamento a R$ 0,25** |
| [**Magnifying Glass**](https://www.ligamagic.com.br/?view=cards/card&card=Magnifying+Glass) | F1 rocha `{C}` + `{4},{T}`: Clue · F3 Artifact · F6 CMC 3 | **4 manas por um Clue que ainda custa `{2}` para virar carta** = 9 manas por 1 carta. [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) faz 3 cartas a `{2}` cada, por 2 manas de entrada |
| [**Panic Spellbomb**](https://www.ligamagic.com.br/?view=cards/card&card=Panic+Spellbomb) | F1 `{T}`,sac: criatura não bloqueia; ao ir ao cemitério, `{R}`: compre 1 · F3 Artifact · F6 CMC 1 | **Cartucho, não vantagem**: 1 mana + `{R}` para repor a própria carta = troca 1:1. O primeiro modo não faz nada num deck que não ataca em massa |
| [**Mad Ratter**](https://www.ligamagic.com.br/?view=cards/card&card=Mad+Ratter) | F1 no **segundo saque de cada turno**, 2 fichas 1/1 · F2 corpo 1/2 · F3 Creature · F4 recebe anthems · F6 CMC 4 | **Não é fonte de draw — é pagador de draw.** O deck de fato saca 2+ por turno, mas o produto são fichas 1/1 pretas: é go-wide, o eixo que o usuário reprovou em 2026-09-23. Vale 1 ponto (corpo no piso da §4 da Fase 2) e por isso é **corte condicionado devolvido à Fase 2**, não dispensa minha: se o piso de 14 corpos ficar apertado, ela é um corpo a R$ 0 |
| [**Mycosynth Wellspring**](https://www.ligamagic.com.br/?view=cards/card&card=Mycosynth+Wellspring) | F1 ETB **e** ao morrer: busca um terreno básico para a mão · F3 Artifact · F6 CMC 2 | Duas cartas ao longo da vida dela, mas **as duas são terrenos básicos** — é fixação de mana travestida de saque. Devolvida à **Fase 4** como ramp/fixação, onde pode valer |
| [**S.H.I.E.L.D. Spy Kit**](https://www.ligamagic.com.br/?view=cards/card&card=S.H.I.E.L.D.+Spy+Kit), [**Courage in Crisis**](https://www.ligamagic.com.br/?view=cards/card&card=Courage+in+Crisis), [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais), [**Lux Artillery**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Artillery) | F1 scry / proliferate / sac-outlet / contador | Scry não é carta; as outras três não sacam. 0–1 ponto |

### 3.3 O que a `lista.txt` não tem

Das 69 não-básicas do Tori físico, **uma única** é fonte de vantagem de cartas: [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold).
Isso não é crítica ao deck antigo — é a medida do gap. **O pacote desta fase é quase todo de
compra**, e isso precisa aparecer no `report.md` junto com o custo.

---

## 4. Candidatas recomendadas

### 4.1 Tabela

| Carta | CMC | Tipo | Como gera vantagem | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|
| [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) | 4 | Creature — Human Soldier 2/2 (W) | ETB: você vira **monarca** (+1 carta em cada end step seu) **e** exila uma criatura de oponente | (1) **é remoção**, conta nas 20–22 da Fase 5 — dois slots num; (2) corpo 2/2 que conta no piso de 14–17 corpos; (3) a coroa recompensa justamente o deck que tem 20+ remoções para defendê-la | não | **R$ 0,30** (19/09) |
| [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) | 3 | Artifact | A cada descarte: **compre 1** / Treasure / cada oponente perde 3 — um modo diferente por turno | (1) **multiplica as 10 cartas de loot da Fase 2**: cada `Thrill`/[**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) vira loot **+ carta**; (2) Treasure paga o `{R}{R}{W}{W}` do escape (sinal da Fase 4); (3) "3 de vida a cada oponente" é alcance sem combate = pacote da §6 da Fase 2 | não | `a cotar` |
| [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) | 3 | Enchantment — Case (R) | ETB: descarte 1, **compre 2** (+1 líquido). Resolvido (mão vazia): **todo upkeep descarte a mão e compre 2** | (1) o gatilho de "solve" é **mão vazia** — a dor 2 é a condição de ligar o motor; (2) "descarte a mão" é a maior injeção de combustível do pool: pode habilitar o escape num turno; (3) enchantment barato que sobrevive a wipe de criatura | não | `a cotar` |
| [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) | 2 | Legendary Creature — Human Knight 2/2 (R) | **Sempre que você descarta uma ou mais cartas**, exila o topo e você pode jogá-lo **até o fim do seu próximo turno** | (1) converte **cada uma das 10 cartas de loot** em carta extra — é o maior multiplicador de sobreposição do deck; (2) corpo 2/2 no T2 para o piso de corpos; (3) janela longa = a taxa de realização mais alta do pacote (~80%) | não | `a cotar` |
| [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) | 3 | Creature — Devil 3/2 menace (R) | Todo upkeep revela o topo: o oponente escolhe entre **te dar a carta** ou **mandá-la ao cemitério e tomar dano igual ao CMC** | (1) **os dois modos servem**: carta, ou combustível de escape + dano — é o punisher sem metade ruim neste deck; (2) corpo 3/2 com menace (piso de corpos + ameaça real); (3) o dano é fonte **vermelha** → +2 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | não | `a cotar` |
| [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc) | 3 | Legendary Artifact | Todo end step: contador + **scry 2**; um oponente escolhe entre **te dar 1 carta** ou **você moer X e ele perder vida igual ao CMC total** | (1) mesma lógica do [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder), e o lado "ruim" é **auto-moinho escalável** — exatamente o motor repetível que a §3.2 da Fase 2 pede; (2) scry 2 por turno de graça num deck que precisa achar terreno **e** resposta; (3) perda de vida é alcance | não | `a cotar` |
| [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) | 2 | Artifact — Book | Entra com 1 página; **ganha 1 página sempre que seu comandante entra ou ataca**; `{1},{T}`, remove página: compre | (1) o Phlage **entra todo escape** e ataca todo turno depois disso — nenhum outro comandante RW liga esta carta tão bem; (2) já paga a si mesma no T3, quando o Phlage entra pela zona de comando; (3) 2 manas, não compete com a curva de remoção | não | **R$ 1,20** (19/09) |
| [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | 5 | Enchantment (R) | ETB: **monarca**; todo upkeep, **7 de dano em qualquer alvo** enquanto você for monarca (2 se não for) | (1) 7 de dano por turno é **remoção repetível** (Fase 5) **e** relógio de 3 turnos na cara de um oponente (Fase 7 — a dor 1); (2) segunda entrada para a coroa depois que ela é roubada; (3) fonte **vermelha** → 9 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | não | `a cotar` |
| [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) | 4 | Enchantment (R) | Modo Khans: todo upkeep exila o topo e você pode jogá-lo **neste turno** — com o mana do turno inteiro intacto | (1) impulse de upkeep é o de melhor taxa (§1.1); (2) o modo Dragons é reserva de alcance se a mesa for de criaturas; (3) permanente que sobrevive a wipe de criatura | **não — está dentro do Krenko montado** (regra 7: compra-se a 2ª cópia) | **R$ 0,90** (22/08) |
| [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) | 3 | Instant (R) | **Impulse 2** com janela até o fim do seu **próximo** turno | (1) **instantâneo**: cast no end step do oponente e você tem o turno inteiro para gastar (a janela que o §1.1 valoriza); (2) custo **R$ 0 — está na caixa**; (3) modal: o segundo modo (+2/+1 em dois alvos) salva um bloqueador num aperto | **sim (caixa)** | R$ 0 |
| [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) | 2 | Artifact — Vehicle 4/4 | 3 contadores; `{2},{T}`, remove: compre. No último, ganha Treasure + ficha Pilot | (1) **3 cartas de um slot de 2 manas**, sem depender de cor — o melhor preço-por-carta do pool; (2) Treasure no fim ajuda o `{R}{R}{W}{W}`; (3) veículo 4/4 crew 3 é bloqueador de emergência (⚠ **F7: crew 3 é caro** neste deck — a Fase 2 avisa que veículo é consumidor de corpo, não corpo) | não | **R$ 1,90** (12/08) |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | 5 | Leg. Creature 3/3 (R) | Toda mágica sua que fere um **jogador** exila o topo e deixa jogar neste turno; `{T}` pinga 1 e dispara ela mesma | (1) já existe fisicamente; (2) com o pacote de queima, dispara 1–3× por turno; (3) conta como corpo e como peça de alcance (Fase 7) | **sim (lista)** | R$ 0 |
| [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) | — | Land | `{4},{T}`, sacrifica: você vira **monarca** | (1) **não gasta slot de mágica** — entra na conta da Fase 6; (2) terceira entrada para a coroa, que é o motor que mais rende neste eixo; (3) produz `{C}` enquanto espera | não | `a cotar` |
| [**War Room**](https://www.ligamagic.com.br/?view=cards/card&card=War+Room) | — | Land | `{3},{T}`, pague 2 de vida: compre 1 | (1) saque repetível **em slot de terreno**; (2) o Phlage devolve 3 de vida a cada gatilho — o custo em vida é o recurso que este deck mais tem sobrando; (3) produz `{C}` | não | `a cotar` |

### 4.2 Duas ressalvas que precisam estar escritas

**A coroa é uma só.** [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) + [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) + [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) **não somam**
3 cartas por turno: são **3 entradas para 1 motor**. O valor delas é a *reentrada* — você
recupera a coroa depois que um oponente a rouba, o que num pod de 4 acontece. Contei as três
como 3 fontes nominais e **~1,3 carta/turno efetiva**. Quem contar 3 está mentindo para si mesmo.

**A coroa convida ataque.** Ser monarca faz a mesa atacar **você**, e este deck tem 14–17 corpos,
vários com `defender`. Três mitigações reais, todas já no plano: (a) os bloqueadores do pacote de
alcance são largos e baratos ([**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) 0/4, [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) 0/3); (b) o Phlage
6/6 bloqueia qualquer coisa; (c) o Helix devolve 3 de vida por gatilho. **F7 do [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer):**
a criatura exilada **volta** no instante em que um oponente vira monarca — perder a coroa custa
uma carta *e* devolve a ameaça. Sinal à Fase 5: guardar 1–2 remoções instantâneas para o maior
atacante enquanto a coroa estiver com você.

### 4.3 Sobreposição declarada — as 5 que contam nas duas colunas

| Carta | Conta em A (vantagem) | Conta também em | Por quê |
|---|---|---|---|
| [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) | sim | **B · combustível — motor repetível** | "descarte a mão, compre 2" todo upkeep é a maior taxa de abastecimento do pool inteiro |
| [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) | sim | **B · combustível** (o oponente escolhe moer) + corpo + alcance | metade das ativações vira carta no cemitério |
| [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc) | sim | **B · combustível — auto-moinho escalável** + alcance | mói X crescente quando o oponente nega a carta |
| [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) | sim | **corpo** (piso da §4) + **pagador do pacote de loot** | não abastece sozinho; multiplica o que a Fase 2 já pôs |
| [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) | sim | **ramp** (Treasure) + **alcance** (3 de vida a cada oponente) | idem — pagador, não fonte |

**Não contei como vantagem real**, embora a Fase 2 as trate como saque: [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility),
[**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), [**Cathartic Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Reunion), [**Tormenting Voice**](https://www.ligamagic.com.br/?view=cards/card&card=Tormenting+Voice), [**Wild Guess**](https://www.ligamagic.com.br/?view=cards/card&card=Wild+Guess), [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers),
[**Electric Revelation**](https://www.ligamagic.com.br/?view=cards/card&card=Electric+Revelation), [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score), [**Bitter Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Bitter+Reunion), [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre), [**Faithless Salvaging**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Salvaging),
[**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king). Todas são troca 1:1 ou pior em contagem de cartas. São categoria B
e continuam valendo — só não fecham a meta desta fase.

---

## 5. Reservas (para os cortes da Fase 6)

| Carta | CMC | Por que é reserva e não titular | Na coleção? | Preço |
|---|---|---|---|---|
| [**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) | 3 → `{R}` | Impulse 2 quase de graça — o Helix do Phlage liga o spectacle sozinho. Fica atrás de [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) só porque esta é R$ 0 e instantânea | **não — dentro do Krenko** | **R$ 0,74** (22/08) |
| [**Bargaining Table**](https://www.ligamagic.com.br/?view=cards/card&card=Bargaining+Table) | 5 | Saque repetível a custo quase zero no late (§3.1), **R$ 0**. Perde por curva: 5 manas num deck cujo T5 é escape + resposta | **sim (caixa)** | R$ 0 |
| [**Dawn of Hope**](https://www.ligamagic.com.br/?view=cards/card&card=Dawn+of+Hope) | 2 | "Sempre que você ganha vida, pague `{2}`: compre" — o Phlage ganha 3 de vida **a cada entrada e cada ataque**, e [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) ganha em bloco. Também fabrica fichas 1/1 com lifelink (corpos + mais gatilhos). Fica fora do titular por competir com [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) no mesmo slot de 2 manas e por ser saque **pago** | não | `a cotar` |
| [**Valakut Awakening // Valakut Stoneforge**](https://www.ligamagic.com.br/?view=cards/card&card=Valakut+Awakening+%2F%2F+Valakut+Stoneforge) | 3 | Instantâneo: devolve N cartas ao fundo e compra **N+1**. Converte mão travada em mão nova (+1 líquido) e **o verso é terreno** — ajuda a Fase 6 a chegar a 37–38 sem gastar slot. Não é motor; é conserto de mão | não | `a cotar` |
| [**Magus of the Wheel**](https://www.ligamagic.com.br/?view=cards/card&card=Magus+of+the+Wheel) | 3 | Corpo 3/3 + `{1}{R},{T}`,sac: todos descartam a mão e compram 7. Para você: **a mão inteira vira combustível** e você reabastece. Fora do titular por ser **simétrico** — reabastece três oponentes num deck de atrito | não | `a cotar` |
| [**Oracle's Vault**](https://www.ligamagic.com.br/?view=cards/card&card=Oracle%27s+Vault) | 4 | Impulse repetível a **R$ 0,25** (caixa). Reserva de orçamento pura | **sim (caixa)** | **R$ 0,25** (19/09) |
| [**Endless Atlas**](https://www.ligamagic.com.br/?view=cards/card&card=Endless+Atlas) | 2 | `{2},{T}`: compre, com 3+ terrenos de mesmo nome — os 30 básicos garantem a condição. Saque incolor repetível. Perde para [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) no preço-por-carta e para [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) na sinergia | não | `a cotar` |
| [**Court of Embereth**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Embereth) | 4 | Coroa 1 mana mais barata que [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) e **fabrica um 3/1 por turno para defender a coroa**. Fora do titular porque o dano dela escala com nº de criaturas (este deck tem poucas) e porque fichas por turno é o cheiro de go-wide que o usuário reprovou | não | `a cotar` |

**Dispensada por preço, não por função:** [**Staff of the Storyteller**](https://www.ligamagic.com.br/?view=cards/card&card=Staff+of+the+Storyteller) (**R$ 9,88** = 4,9% do teto)
e [**Reckless Impulse**](https://www.ligamagic.com.br/?view=cards/card&card=Reckless+Impulse) (**R$ 9,78**, texto idêntico ao de [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity), que está na caixa).
Se a Fase 6 precisar de um impulse-2 em feitiço, **[**Wrenn's Resolve**](https://www.ligamagic.com.br/?view=cards/card&card=Wrenn%27s+Resolve) tem texto idêntico ao de
[**Reckless Impulse**](https://www.ligamagic.com.br/?view=cards/card&card=Reckless+Impulse)** e é comum — cotar as duas e comprar a mais barata.

---

## 6. Curva das fontes de draw

**mv 1–2: 3** (`Inti` 2 · [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) 2 · [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) 2)
**mv 3–4: 7** ([**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) 3 · [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) 3 · [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) 3 · [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc) 3 · [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) 3 · [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) 4 · [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) 4)
**mv 5+: 2** ([**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) 5 · [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) 5)
**slot de terreno (mv —): 2** ([**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) · [**War Room**](https://www.ligamagic.com.br/?view=cards/card&card=War+Room))

Média das 12 não-terreno: **3,0**. Nenhuma concentração em 5+: só 2 cartas, e uma delas já
existe fisicamente. O miolo em mv 3 é intencional — é o turno em que o Phlage sai da zona de
comando e o deck tem mana sobrando antes de o escape começar a consumir tudo.

---

## 7. Orçamento de slots — quanto eu preciso e o que espero que saia

Com **37 terrenos** (a Fase 2 já disse que 36 não bastam), sobram **62 slots** não-terreno.
Somando os pedidos declarados até aqui: 14 combustível + 21 remoção + 3,5 wipes + 6 pingadores +
3,5 multiplicadores + 10,5 ramp = **58,5**. Restam **3,5 slots** — e eu preciso de mais que isso.
**O deck está estourado antes de esta fase começar**, e é honesto dizer isso em vez de fingir que
cabe.

**Minha conta de slots:**

| Tipo de slot | Quantas | Quais |
|---|---|---|
| **Terreno** (Fase 6, não disputa mágica) | 2 | [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City), [**War Room**](https://www.ligamagic.com.br/?view=cards/card&card=War+Room) |
| **Já existe** (custo zero de slot novo) | 1 | [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) |
| **Compartilhado** — já cabe dentro de outra meta | 7 | [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) e [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (dentro das 20–22 remoções) · [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder), [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc) (dentro dos 14 de combustível) · `Inti` e [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) (dentro dos corpos / do alcance) |
| **Exclusivo — é o que eu preciso que abra** | **4** | [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends), [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster), [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege), [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) |

**Peço 4 slots exclusivos.** De onde eu proponho que saiam — a decisão é da Fase 6, eu só nomeio
com contrapartida:

1. **2 slots das 20–22 remoções → 19–20.** Não é perda de função: [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) **é** remoção
   (exila uma criatura) e [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) **é** remoção repetível (7 por turno em qualquer alvo).
   A contagem de remoção efetiva não cai; ela muda de nome. A Fase 2 mediu P(≥1 remoção na mão)
   80,5% com 20 e 76,6% com 18 — em 19–20 o deck fica dentro da faixa que ela mesma chamou de
   alvo.
2. **2 slots dos 10 loots one-shot → 8** (os 4 motores ficam intactos). Contrapartida em três
   linhas: (a) [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) descarta a **mão inteira** todo upkeep — abastece mais
   que dois [**Tormenting Voice**](https://www.ligamagic.com.br/?view=cards/card&card=Tormenting+Voice) somados; (b) [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) e `Palantír` moem passivamente todo
   turno; (c) **toda carta extra sacada vira uma mágica conjurada, que vira uma carta no
   cemitério** — o pacote de vantagem é, ele mesmo, combustível indireto. O custo declarado: a
   config da §3.1 da Fase 2 desce de **D (14 fontes)** para algo entre **C e D**, o que pela
   tabela dela move P(escape no T6) de 52,4% para ~48%. **Custo aceito e declarado**, não
   escondido.

**Os dois loots que eu indicaria, com ficha (regra 4) — `corte condicionado`, decisão da Fase 6:**

- **[**Tormenting Voice**](https://www.ligamagic.com.br/?view=cards/card&card=Tormenting+Voice)** `{1}{R}` — **F1** custo adicional: descarte 1; compre 2. **F2** sem corpo.
  **F3** Sorcery (alimenta spell mastery de [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) e os pingadores). **F4** nada
  recebe. **F5** nada facilita. **F6** T2. **F7** disputa o turno 2 com remoção e com `Inti`.
  Funções e quem as cobre: *loot 2* → [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) (instantâneo, melhor) e
  [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) (3 ao cemitério por 1 mana); *gatilho de pingador* → sobram 30+ mágicas;
  *combustível* → [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse). **Nenhuma função fica descoberta.**
- **[**Wild Guess**](https://www.ligamagic.com.br/?view=cards/card&card=Wild+Guess)** `{R}{R}` — **texto idêntico ao de [**Tormenting Voice**](https://www.ligamagic.com.br/?view=cards/card&card=Tormenting+Voice)**, com custo de cor
  **pior** (`{R}{R}` numa base que precisa de `{R}{R}{W}{W}` no T5–T6). Mesmas funções, mesmos
  cobridores. Único argumento a favor: exercita a base vermelha — e ele vira **contra** a carta,
  porque o gargalo declarado pela Fase 2 é o branco duplo. **Nenhuma função fica descoberta.**

Se a Fase 6 preferir não mexer no combustível, a alternativa é cortar [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) e
[**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) (as duas fontes com menor sobreposição do meu pacote) e **aceitar 12 fontes
nominais** — resultado: ~10 efetivas, com o pacote de coroa intacto. É o corte que eu faria se
tivesse de escolher.

---

## 8. Riscos declarados

1. **Metade do pacote titular está `a cotar`** — 8 das 14. As cotadas somam **R$ 4,30**
   ([**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) 0,30 + [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) 1,20 + [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) 0,90 + [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster)
   1,90) e 2 são R$ 0 (caixa/lista). **As desconhecidas são justamente as melhores**
   ([**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance), [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc), [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder),
   `Inti`) — são raras de sets recentes ou de LTR, e pelo menos uma delas pode sozinha valer 10%
   do teto. **Cotar antes de a Fase 6 fechar cortes.** Se estourarem, a fila de substituição por
   ordem é: [**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) (R$ 0,74) → [**Oracle's Vault**](https://www.ligamagic.com.br/?view=cards/card&card=Oracle%27s+Vault) (R$ 0,25, caixa) → `Bargaining
   Table` (R$ 0, caixa) → `Endless Atlas` → `Dawn of Hope`.
2. **Ódio de cemitério bate duas vezes.** [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) desliga o escape **e** apaga o valor de
   `Palantír`, [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) e [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse). A §3.4 da Fase 2 já pede 2 respostas
   a encantamento; com este pacote elas passam de recomendáveis a **obrigatórias**.
3. **A coroa muda a política da mesa** (§4.2). É o único item deste relatório que altera como o
   usuário é tratado no pod, e ele deve saber disso antes de comprar.
4. **[**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) é veículo, não corpo.** Crew 3 num deck de 14–17 criaturas pequenas
   quase nunca é pago. Ele entra como "3 cartas por 2 manas", e o 4/4 é bônus improvável — a
   Fase 2 avisou que veículo é consumidor de corpo, e a advertência vale aqui contra mim.
5. **Impulse perdido é carta destruída**, não carta no cemitério (§1.1). Com 5 fontes de impulse
   e taxa de 2/3, o deck "queima" ~1 carta a cada 3 gatilhos. É o preço de jogar saque em RW e
   está embutido no número efetivo de 11,5.

---

## 9. Sinais cruzados

| Fase | Sinal |
|---|---|
| **2 · tema** | Duas correções: (a) **impulse não tira carta do cemitério** — mágica conjurada do exílio vai para o cemitério normalmente; o limite de "no máximo 2 slots de impulse" foi fixado sobre premissa errada e eu proponho 5, descontados a 2/3. (b) O pedido de "8 a 10 das 12–13 fontes serem loot" anula a fase: loot é categoria B. Aceito a sobreposição por outra via — 5 das 14 (§4.3) |
| **4 · ramp** | [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) e [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) produzem fichas de Treasure, que é exatamente o que a Fase 4 precisa para o `{R}{R}{W}{W}`. Contem nas duas metas. [**Mycosynth Wellspring**](https://www.ligamagic.com.br/?view=cards/card&card=Mycosynth+Wellspring) (caixa) devolvida a vocês como fixação, não como saque |
| **5 · interação** | [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) (exila criatura) e [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (7 de dano/turno em qualquer alvo) **devem ser contados dentro das 20–22 remoções** — é assim que os 2 slots que eu peço se pagam. Pedido adicional: manter 1–2 remoções **instantâneas** disponíveis enquanto a coroa estiver conosco (§4.2) |
| **6 · manabase** | **2 terrenos do meu pacote**: [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) e [**War Room**](https://www.ligamagic.com.br/?view=cards/card&card=War+Room). Ambos produzem só `{C}` — **não** contem como fonte de cor; entram como utilidade sobre os 37, não no lugar de um terreno que produz `R`/`W`. [**Valakut Awakening**](https://www.ligamagic.com.br/?view=cards/card&card=Valakut+Awakening+%2F%2F+Valakut+Stoneforge) (reserva) é MDFC e pode entrar como meio-terreno se a base ficar curta. Os 4 slots exclusivos e as duas fontes de corte estão na §7 |
| **7 · wincon** | O pacote de saque **é** meio pacote de alcance: [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (7/turno), [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) (3 a cada oponente por descarte), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) e `Palantír` (dano por CMC quando negam a carta), [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold). Todos são **fonte vermelha**, exceto [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) — contem no cálculo do [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) |
| **Orquestrador** | (a) **8 de 14 `a cotar`**, e são as melhores — a captura da LigaMagic precisa cobrir [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance), [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc), [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder), [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun), [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire), [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City), [**War Room**](https://www.ligamagic.com.br/?view=cards/card&card=War+Room). (b) A caixa tem **3** fontes de vantagem real em RW, e só 1 é motor — o pacote é de compra. (c) O deck pede 62 slots não-terreno e as fases somam 58,5 **antes** desta; a §7 nomeia os 4 que faltam e a contrapartida de cada um |

---

## Revisão 2026-09-24 — pacote de reanimação

> **Escopo.** Devolução leve à Fase 3, modo `improve`. A lista de partida é a v2 do `report.md` §6
> **com as 15 trocas do `02-theme.md` §12 aplicadas**. Não mexi em terrenos nem proponho trocas
> fora de draw. As Fases 4, 5 e 6 rodam em paralelo.
>
> **Regra 6.** Oracle puxado nesta sessão com `bin/mtgdb oracle` (Tocasia's Welcome, Ark of Hunger,
> Haliya, Seize Opportunity, Outpost Siege, Tome of Legends, Syr Carah, Inti, Light Up the Stage,
> Culling Dais, Lesser Masticore, Lorehold Excavation, Reckoner Bankbuster). Rulings com
> `bin/mtgdb rulings` (Ark, Haliya, Tome, Culling Dais, Phlage). **Regra 2.** Preços só de
> `bin/mtgdb prices`, LigaMagic (menor). **EDHREC não chamado**: o radar é o do `08` §3.
> **Regra 5.** `decisions.md` conferido: nenhuma das cartas desta seção (entra ou sai) tem registro
> de corte. Nenhuma é reposição.

**Fontes no deck (lista revisada da §12): 12/12–13**, no piso.

| # | Fonte | MV | Situação na §12 |
|---|---|---|---|
| 1 | [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) | 2 | fica · **enfraquecida** (3 loots saíram, e cada descarte a menos é um impulse a menos) |
| 2 | [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) | 2 | fica · **muito fortalecida** (ver 3.2) |
| 3 | [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) | 2 | fica · o crew 3 ficou pagável (Warsinger, Bearer, Broodmoth e Desdemona têm poder 3) |
| 4 | [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) | 3 | fica |
| 5 | [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) | 3 | fica |
| 6 | [**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) | 3 (`{R}`) | fica · o spectacle liga com qualquer Helix na cara |
| 7 | [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) | 3 | **entra** (§12) |
| 8 | [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) | 4 | fica |
| 9 | [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) | 4 | **entra** (§12) |
| 10 | [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | 5 | fica · **a mais fraca agora** (ver 3.3) |
| 11 | [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | 5 | fica |
| 12 | [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) | terreno | fica |

### 1. As duas trocas da Fase 2: **confirmadas**

#### 1.1 Sai Seize Opportunity → entra Tocasia's Welcome · **confirmada**

**Ficha da que sai: Seize Opportunity** (caixa, R$ 0,05 na régua)

| Eixo | Leitura (oracle de hoje) | Quem cobre depois do corte |
|---|---|---|
| F1 | modal: **impulse 2** até o fim do seu próximo turno, **ou** +2/+1 em até duas criaturas | impulse: [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) (repetível) e o `{T}` do [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger); pump: o [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) põe +1/+1 e atropelar no ataque |
| F2 | sem corpo | — |
| F3 | instantâneo: tipo no cemitério para o delirium da [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | sobram 10 instantâneos, e o delirium só pede 4 tipos |
| F4 | nada | — |
| F5 | pump em combate | Inti (parcial) |
| F6 | 3 manas, qualquer turno | — |
| F7 | disputava a mana guardada para remoção instantânea | — |

**Descoberto:** a compra **instantânea** de uma vez (lançada no fim do turno do oponente). Custo
aceito: é 1,3 carta **uma vez**, e a Tocasia rende ~0,75 carta **por turno** com o motor girando.
**Regra 7:** sai uma carta da caixa, que continua sobressalente.

**Por que a Tocasia é mais que a Seize neste eixo.** O oracle diz *"Whenever one or more creatures
you control with mana value 3 or less enter, draw a card. This ability triggers only once each
turn."* O Phlage tem MV 3, então **toda reanimação dispara a Tocasia**, e com ela a mágica de
reanimação de uma vez vira um **Helix que se paga em carta**: gasta 1, compra 1. Ela também dispara
no **turno do oponente** (ela não diz "seu turno"): o [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) e a [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth)
devolvem o Phlage na hora da morte, e a persistência da [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) faz o mesmo. Outras
entradas de MV ≤ 3 na lista: [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun), [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin), [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert),
[**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise), [**Zookeeper Mechan**](https://www.ligamagic.com.br/?view=cards/card&card=Zookeeper+Mechan), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) e [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger).

#### 1.2 Sai Outpost Siege → entra Ark of Hunger · **confirmada, e o Ark conta como impulse repetível**

**Ficha da que sai: Outpost Siege** (compra, R$ 0,90; a cópia física está no Krenko montado)

| Eixo | Leitura | Quem cobre |
|---|---|---|
| F1 Khans | todo upkeep, exila o topo e deixa jogar **neste turno** | `{T}` do [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger): moe 1 e deixa jogar neste turno |
| F1 Dragons | 1 de dano em qualquer alvo quando criatura sua sai do campo (cada sacrifício do Phlage) | o Ark causa 1 a cada oponente quando o Phlage sai do cemitério, e o Helix causa 3. O modo exclui o Khans, então como fonte de draw ele nunca foi escolhido |
| F2 | sem corpo | — |
| F3 | encantamento; o deck não conta encantamentos | — |
| F4 | fonte vermelha ([**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) +2 no modo Dragons) | o Ark é RW, então também é fonte vermelha |
| F6 | 4 manas, e rende a partir do upkeep **seguinte** | o Ark ativa **no turno em que entra** |
| F7 | impulse perdido é carta exilada para sempre | — |

**Descoberto:** nada que o deck use, dentro do modo Khans.

**O Ark pelo meu critério (§1.1).** O critério é: impulse conta, descontado a 2/3 de realização.
Aplicado igual aos dois lados (simetria), o Ark **ganha** do Siege em três pontos e perde em um:

| Ponto | Outpost Siege (Khans) | Ark of Hunger `{T}` |
|---|---|---|
| Taxa de realização | ~2/3, janela "neste turno", no upkeep | ~2/3, mesma janela, **mas você escolhe quando ativar**: depois da compra do turno e com a land drop à vista |
| Carta **não** jogada | exilada: **carta destruída** | **vai ao cemitério**: combustível para o escape, para o [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial) e para o delirium |
| Carta jogada | sai do exílio | **sai do cemitério e dispara o próprio Ark** (1 a cada oponente + 1 de vida). Ruling de 2026-03-20: vários eventos simultâneos disparam uma vez só |
| Primeira ativação | upkeep seguinte | no turno em que entra (artefato não tem enjoo) |
| Exposição | encantamento | **artefato**: fica mais exposto a remoção de artefato, que é mais comum em mesa. É o ponto em que perde |

**Veredito: conta.** É impulse repetível de ~0,65 carta por turno, igual ao Siege, e a terça parte
que não se realiza vira recurso em vez de sumir. Com a mágica de reanimação que o Ark moer
acontece o mesmo: jogada do cemitério, ela dispara o Ark duas vezes, uma quando sai de lá e outra
quando o Phlage sai.

### 2. Contagem

**12 → 12.** As trocas são uma por uma. O pacote fica no piso da faixa.

### 3. Doze bastam? A pergunta certa é de **taxa**, e o motor mexeu nela

A Fase 7 (`07` §3.4) mediu que falta **carta na mão**, não mana: 0,8 mágica por turno no T7–T11,
com 7–9 manas sobrando. O motor novo muda os dois lados da conta.

#### 3.1 O que o motor **consome**

Das 10 reanimações, só as **5 de uma vez** gastam carta ([**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand), [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant),
[**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness), [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission), [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation)). As 5 repetíveis
([**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle), [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan), [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger), [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial),
[**Desdemona, Freedom's Edge**](https://www.ligamagic.com.br/?view=cards/card&card=Desdemona%2C+Freedom%27s+Edge)) geram Helix **sem gastar carta**. E as de uma vez não são gasto novo:
entraram no lugar de 4 remoções formais, e remoção também custava uma carta cada. **O consumo por
turno fica igual ao da v2.** O que muda é que ele passa a ter como se pagar.

#### 3.2 O que o motor **devolve** (taxa por fonte **em campo**, T6–T10)

> Estimativa de mesa, com a mesma premissa dos dois lados: **~1 reanimação por turno seu** a partir
> do T4, que é a hipótese da §12. **Não é simulação.** Quem mede é o goldfishing da Fase 7.

| Fonte | v2 | Revisada (§12) | Por quê |
|---|---|---|---|
| [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) | ~0,4/turno | **~0,9/turno** | na v2 ganhava página no T3 e depois só com o Phlage escapado (T6–T7). Agora **toda reanimação é o comandante entrando** (*"Whenever your commander enters or attacks"*), e com a Broodmoth são 2 páginas. O limite passa a ser o `{T}`: uma compra por rodada |
| [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) → [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) | 1,3 uma vez | **~0,75/turno** | 1 por turno com entrada de MV ≤ 3, às vezes também no turno do oponente |
| [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) → [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) | ~0,65/turno | ~0,65/turno | e a carta perdida vira combustível |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | ~0,65/turno a partir do T6 | ~0,65/turno a partir do T6 | é o `{T}` dela mesma que dispara o impulse. As mágicas que acertam jogador ficaram poucas ([**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab), [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate)). O Helix do Phlage **não** dispara a Syr Carah, porque é habilidade de criatura, e ela só lê *"instant or sorcery spell"* |
| [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) | ~0,6/turno | ~0,45/turno | saíram [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility), [**Cathartic Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Reunion) e [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers) |
| Nas quatro linhas de cima | 1,7/turno + 1,3 uma vez | **2,95/turno** | +1,25 carta por turno com as fontes em campo, sem mudar a contagem |

**Conclusão: 12 bastam em número, mas não em redundância.** A taxa subiu. O ponto fraco é outro:
só **2 cartas** devolvem carta pela reanimação ([**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) e [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends)).
Sem nenhuma delas em campo, cada mágica de reanimação de uma vez é −1 carta. P(ver ao menos uma das
duas), hipergeométrica sobre 99: **19,3% até o T4 · 22,9% até o T6 · 26,4% até o T8**. Com uma
terceira, sobe para **27,6% · 32,4% · 37,0%**.

#### 3.3 A fonte mais fraca do pacote é a Syr Carah

Na nova lista ela custa 5 manas, rende a partir do T6 (o `{T}` sofre enjoo), não entra no motor
(MV 5 fica fora de todas as reanimações de MV ≤ 3) e perdeu as mágicas de dano na cara que a
alimentavam. Ela ainda é uma fonte real. Mas é a única das 12 que o eixo novo **enfraqueceu sem dar
nada em troca**. O Inti também perdeu, mas continua multiplicando os 5 loots que ficaram.

### 4. Troca proposta — em par, dentro de draw

**Por que não uma 13ª fonte.** Uma 13ª exige abrir slot em outra categoria, e os três candidatos
óbvios pertencem a fases que estão rodando agora: combustível (~11, no alvo), remoção formal (16,
abaixo do que a Fase 5 pediu) e ramp (9/2, abaixo da meta do pipeline). Trocar a pior fonte pela
melhor disponível rende mais taxa do que somar uma fonte média, e não toca em nenhuma outra fase.
**Se alguma fase abrir slot, a 13ª de custo zero de compra é a Culling Dais (§5).**

#### 4.1 Tabela

| Carta | CMC | Tipo | Como gera vantagem | Sinergias (mín. 2) | Na coleção? | R$ |
|---|---|---|---|---|---|---|
| [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) | 3 (Warp `{W}`) | Legendary Creature — Human Soldier 3/3 | No **fim do seu turno**, compra 1 se você ganhou 3+ de vida. A vida é conferida **na resolução** (ruling 2025-07-25) | (1) **todo Helix dá 3 de vida**, então todo turno em que o motor gira ela compra. Ela pega o que a Tocasia não pega: o Helix de **ataque** do Phlage escapado, o [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) e o Phlage escapado que volta pelo Dial ou pela Desdemona. Ela mesma soma 1 por criatura ou artefato que entra: Phlage reanimado = 4 de vida; (2) **lendária de MV 3**: dispara a [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) ao entrar, e conjurá-la é **mágica histórica** para o [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) (também pelo Warp, então são 2 gatilhos por carta); também é alvo de toda reanimação do deck; (3) **Warp `{W}`**: no T4, Warp + [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) = 2 manas para Helix + 2 cartas (Haliya e Tocasia). O gatilho de compra e o exílio do Warp disparam juntos no começo do end step, e a compra resolve mesmo com ela exilada (CR 113.7a). Depois ela volta do exílio e dispara a Tocasia de novo; (4) corpo 3/3 que bloqueia e tripula o [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) (crew 3) · origem: EDHREC Top Cards 0,39 (`08` §3) | não | **11,88** (2026-09-24) |

**Ficha F2–F7 da que entra:** F2 3/3, bloqueia, tripula crew 3 · F3 criatura (piso de 17), lendária
(Teshar), MV 3 (Tocasia, reanimável) · F4 recebe o +1/+1 do Inti; é devolvida pela Broodmoth e pelo
Renewal · F5 +1 de vida por criatura ou artefato, e isso completa os 3 de vida em turno sem Helix
(rocha + criatura + ficha de Treasure do [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) ou do [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils)) · F6 T3 conjurada,
ou `{W}` no turno do Helix · F7 o Warp tira ela de campo no fim do turno, então ele só compensa em
turno de Helix; é lendária, mas o deck só tem uma cópia.

#### 4.2 Sai: Syr Carah, the Bold — ficha completa (regra 4)

Origem: **Tori físico** (R$ 0,09 na régua). Posse não é obrigação de uso (regra 7). Se ela vai
para as sobressalentes é o usuário quem diz, na `/update-collection`. Nada aqui infere posse.

| Eixo | Função | Quem cobre depois do corte |
|---|---|---|
| F1a | impulse 1 sempre que ela ou uma mágica instantânea/feitiço sua causa dano a **jogador** | [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light): 1 carta **garantida** (não impulse) por turno de Helix, a partir do T4 e não do T6 |
| F1b | `{T}`: 1 de dano em qualquer alvo, **repetível e instantâneo** (3 com [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)) | em volume, o **Helix em loop** (3 em qualquer alvo por reanimação). Criatura: `{4}` da [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore). Jogador: [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) e [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger). **Descoberto:** o ping barato em velocidade de instantâneo, que mata X/1 e ficha e termina planeswalker |
| F2 | 3/3, bloqueia, tripula crew 3 (e aí não pinga, porque o crew disputa o `{T}`) | Haliya 3/3 faz o mesmo |
| F3 | criatura (piso de 17) · **lendária** (Teshar) · vermelha | Haliya: criatura, lendária. **O piso fica em 17.** **Descoberto:** uma fonte vermelha a menos para o Torbran |
| F4 | recebe Inti, Renewal, Broodmoth | Haliya recebe os mesmos, e **ainda é reanimável** (MV 3). A Syr Carah não era |
| F5 | nada | — |
| F6 | T5, rende do T6 | Haliya: T3–T4, rende no mesmo turno |
| F7 | o ping e o crew disputam o `{T}` · MV 5 fica fora do motor · disputa o slot de 5 com a [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | — |

**Funções descobertas, declaradas:** (a) ping instantâneo repetível; (b) uma fonte vermelha para o
Torbran. As duas são de outras especialidades (remoção → Fase 5; alcance com Torbran → Fase 7), então
o corte é **condicionado** nesses dois pontos. Se uma das duas fases disser que o ping é
insubstituível, a troca cai e a Syr Carah fica. Draw continua em 12 e a redundância segue em 2
conversores. **Cor:** a troca passa uma mágica do vermelho para o branco (22/21 → 21/22), com custo
`{2}{W}` no lugar de `{3}{R}{R}`. É sinal leve para a Fase 6.

**Simetria (checklist §3).** Avaliei os dois lados com ~1 reanimação por turno, sem anthem, com o
Torbran em campo como possibilidade dos dois lados (ele ajuda o ping da Syr Carah, e não ajuda a
compra da Haliya, e isso está escrito como perda), e com a mesma régua de impulse a 2/3.

#### 4.3 Efeito

| | Antes | Depois |
|---|---|---|
| Fontes de draw | 12 | **12** |
| Conversores de reanimação em carta | 2 (Tocasia, Tome) | **3** (Tocasia, Tome, Haliya): P(≥1 até o T6) 22,9% → **32,4%** |
| Taxa das fontes trocadas | ~0,65/turno a partir do T6 | **~0,85–1/turno a partir do T4** |
| Criaturas / lendárias (Teshar) | 17 / 5 | 17 / 5 |

### 5. Regra 7 — a caixa antes da compra

`bin/mtgdb collection -list` foi varrida de novo contra o motor novo, filtrando identidade RW e texto
de compra, ganho de vida e reanimação. Candidatas reais em RW: [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais), [**Bargaining Table**](https://www.ligamagic.com.br/?view=cards/card&card=Bargaining+Table),
[**Oracle's Vault**](https://www.ligamagic.com.br/?view=cards/card&card=Oracle%27s+Vault), [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) (a que sai). [**Melded Moxite**](https://www.ligamagic.com.br/?view=cards/card&card=Melded+Moxite), [**Ragged Short Spear**](https://www.ligamagic.com.br/?view=cards/card&card=Ragged+Short+Spear) e
[**Prophetic Prism**](https://www.ligamagic.com.br/?view=cards/card&card=Prophetic+Prism) repõem a própria carta (troca 1 por 1), e as auras que compram ([**Angelic Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Gift),
[**Dragon Mantle**](https://www.ligamagic.com.br/?view=cards/card&card=Dragon+Mantle)) também.

#### 5.1 Culling Dais — correção de leitura e o que mudou

A v2 desta fase (§3.2 acima) dispensou a [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais) junto com outras três dizendo que *"não
sacam"*. **Ela saca:** *"{1}, Sacrifice this artifact: Draw a card for each charge counter on this
artifact."* Foi erro de leitura. Ela nunca esteve no deck, então não há registro no `decisions.md`.
**O que mudou:**
1. A leitura foi corrigida.
2. O Phlage passou a **se sacrificar a cada reanimação**. Com o gatilho de sacrifício na pilha,
   `{T}`, sacrifique o Phlage na Dais: +1 contador **de graça**. O Helix resolve igual, porque a
   habilidade existe independente da fonte (CR 113.7a). Ele vai ao cemitério pronto para a próxima.
3. Ela vira saída de sacrifício que **protege o Phlage escapado** de exílio ou roubo. Sacrificado
   em resposta, ele vai ao cemitério e continua reanimável, sem imposto (CR 903.9a, §12.1).

**Ficha:** F1 `{T}` + sacrificar criatura: contador; `{1}` + sacrificar a Dais: compra 1 por contador ·
F2 sem corpo · F3 **artefato**: histórico para o Teshar, e a [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission), o [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) e a
[**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) a devolvem (MV 2) · F4 recebe contadores · F5 proteção do comandante ·
F6 T2, rende a partir do T4 · F7 é **um contador por rodada** (o `{T}`), a compra vem em bloco e
tardia (3 cartas por volta do T7), e trazê-la de volta **disputa** o mesmo efeito de reanimação com o
Phlage.

**Por que ela não substitui a Haliya nesta troca:** ela não cobre as funções da Syr Carah que a
Haliya cobre. Não tem corpo (o piso cai para 16), não é lendária criatura para a Tocasia e só compra
em bloco, tarde e pagando `{1}` mais a própria carta. A Haliya compra no **mesmo turno** do Helix, a
partir do T4. **Motivo da dispensa:** curva (F6) e corpo (F2), não "carta fraca".

**Onde ela serve:** é a **13ª fonte de custo zero de compra**, se alguma fase abrir slot, e a
**1ª reserva** se a folga for gasta antes da Haliya (aí fica a Syr Carah e a Dais entra como 13ª).
Preço na régua: **`a cotar`** (está na caixa, mas o teto conta o valor da carta).

#### 5.2 Outras reservas (sem mudança de veredito)

| Carta | Por que continua reserva | Na coleção? | R$ |
|---|---|---|---|
| [**Oracle's Vault**](https://www.ligamagic.com.br/?view=cards/card&card=Oracle%27s+Vault) | impulse repetível, mas 4 manas e 3 turnos de `{2}` para ficar grátis. O [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) faz o mesmo de graça a partir do turno em que entra | sim | 0,25 |
| [**Bargaining Table**](https://www.ligamagic.com.br/?view=cards/card&card=Bargaining+Table) | 5 manas; no late é compra barata, mas a curva de 5 já tem [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) e [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) | sim | a cotar |
| [**Dawn of Hope**](https://www.ligamagic.com.br/?view=cards/card&card=Dawn+of+Hope) | *"Whenever you gain life, you may pay {2}: draw"*: todo Helix vira compra paga, e o deck sobra 7–9 manas no T9–T11 (`07` §3.4). **Promovida a 2ª reserva de compra**: é o mesmo gatilho da Haliya, trocando o corpo por uma compra que não tem limite por turno | não | a cotar |

### 6. Custo — LigaMagic (menor), cotações de 2026-08-22 a 2026-09-24

| Item | R$ |
|---|---|
| Lista revisada (§12, com as 2 trocas de draw já dentro) | 162,12 |
| As 2 trocas confirmadas (já contadas acima): −Seize Opportunity 0,05 −Outpost Siege 0,90 +Tocasia's Welcome 7,88 +Ark of Hunger 0,45 | (+7,38, já incluído) |
| **Proposta: +Haliya, Guided by Light 11,88 − Syr Carah, the Bold 0,09** | **+11,79** |
| **Total com a proposta** | **173,91** |
| Folga no teto de R$ 200 | 37,88 → **26,09** |

A folga é disputada pelas Fases 4, 5 e 6 em paralelo. Se ela acabar antes, a troca cai e a Culling
Dais entra no lugar dela (§5.1), com compra zero e régua `a cotar`.

### 7. Curva das fontes de draw (com a proposta)

**mv 1–2: 3** (Inti 2 · Tome of Legends 2 · Reckoner Bankbuster 2)
**mv 3–4: 7** (Sin Prodder 3 · Case of the Crimson Pulse 3 · Light Up the Stage 3 · Tocasia's Welcome 3 · Haliya 3 · Palace Jailer 4 · Ark of Hunger 4)
**mv 5+: 1** (Court of Ire 5)
**slot de terreno: 1** (Throne of the High City)

Antes da proposta: 3 · 6 · 2 · 1. A fonte de 5 que sai era a única que rendia só a partir do T6.

### 8. Sinais

| Fase | Sinal |
|---|---|
| **5 · interação** | O corte da Syr Carah tira **um ping instantâneo repetível** (mata X/1, fichas e planeswalker ferido). Se vocês o contavam como remoção, o corte é condicionado a isso |
| **6 · manabase** | Mais uma mágica branca e uma vermelha a menos (`{2}{W}` no lugar de `{3}{R}{R}`) |
| **7 · wincons** | Uma fonte vermelha a menos para o Torbran. Métrica nova para o goldfishing: **cartas compradas por reanimação** (Tocasia, Tome, Haliya), e o turno em que o primeiro conversor aparece |
| **Orquestrador** | (a) Cotar [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais) na LigaMagic (está na caixa, mas o teto conta o valor). (b) A §3.2 original desta fase dispensou a Dais por erro de leitura. A correção está em §5.1. (c) Se a folga não comportar a Haliya, a alternativa de compra zero é manter a Syr Carah e pôr a Dais como 13ª fonte, se alguma fase abrir slot |
