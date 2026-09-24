# Condições de Vitória e Testes — Otharri, Suns' Glory (v1 · 2026-09-19 · modo `improve`)

> **Dor nº 1 do usuário:** *"monta board, ataca, não converte em vitória"*. Esta fase existe para
> essa frase. O diagnóstico é preciso: sob a Tori o deck **construía largura sem multiplicador e
> sem furo de bloqueio** — 20 corpos 2/2 contra 3 bloqueadores 4/4 causam zero de dano. Otharri
> resolve a *largura*; ele **não** resolve a *conversão*. Largura só vira vitória com três coisas
> que o deck hoje não tem: **multiplicador estático**, **furo de bloqueio** e **uma rota que não
> passe pelo combate**.

Todos os textos oracle desta fase foram puxados nesta sessão via `bin/mtgdb oracle`
(regra 6). Rulings conferidos via `bin/mtgdb rulings`.

---

## 0. A regra de ouro, aplicada carta a carta

Ruling de **Basri Ket** (2020-06-23), a referência canônica para o motor do Otharri:

> *"Although the Soldiers are attacking creatures, they **were never declared as attacking
> creatures**. This means that abilities that trigger whenever a creature attacks **won't
> trigger** when the Soldiers enter the battlefield attacking."*

Consequência mecânica, e é ela que separa carta boa de carta morta neste deck:

| Classe de habilidade | Alcança as fichas que **entram atacando**? | Quando liga |
|---|---|---|
| **Estática** (`Attacking tokens you control have double strike`, `Creatures you control have menace`, anthem `+X/+X`) | **Sim, sempre** | **1º ataque** |
| **Gatilho de "criatura entra"** (`Whenever a creature you control enters`) | **Sim — e dispara N vezes**, uma por ficha | **1º ataque** |
| **Efeito de feitiço do turno** conjurado na fase principal pré-combate (`creatures can't block this turn`, `tap all creatures your opponents control`, `creatures gain trample`) | **Sim** — o efeito é do turno, pega quem entrar depois | **1º ataque** |
| **Gatilho `whenever ~ attacks` de uma carta que você declarou atacante** (Tori, Hero of Bladehold, Combat Celebrant) | **Sim, por ordenação de pilha** — ponha o gatilho do Otharri por último na pilha (resolve primeiro); o efeito enxerga o board **na resolução** | **1º ataque** |
| **Gatilho `whenever a creature you control attacks`** (Shared Animosity, Hellrider) | **Não pelas fichas novas** — só pelas fichas **antigas**, que desviraram e foram declaradas | **2º ataque em diante** |
| **Gatilho `whenever one or more tokens you control attack a player`** (2ª habilidade da Neyali) | **Não pelas fichas novas** — idem | **2º ataque em diante** |
| **Melee, mentor, battle cry concedido** | **Nunca** | — |

**Teste aplicado a toda carta proposta nesta fase — coluna "Liga em qual ataque?" das tabelas
abaixo.** Nenhuma recomendação minha depende do 2º ataque para existir.

---

## 1. A matemática, refeita com os multiplicadores

Otharri: `Flying, lifelink, haste` · *Whenever Otharri attacks, you get an experience counter.
Then create a 2/2 red Rebel creature token that's tapped and attacking **for each experience
counter you have**.*

Acumulação **N(N+1)/2**. Contra **um** oponente de 40, atacando todo turno a partir do turno em
que Otharri entra (haste = ataca no mesmo turno):

| Ataque nº | Contadores | Fichas novas | Fichas no board | Dano do turno (sem multiplicador) | Acumulado |
|---|---|---|---|---|---|
| 1 | 1 | 1 | 1 | 3 (Otharri) + 2 = **5** | 5 |
| 2 | 2 | 2 | 3 | 3 + 6 = **9** | 14 |
| 3 | 3 | 3 | 6 | 3 + 12 = **15** | 29 |
| 4 | 4 | 4 | 10 | 3 + 20 = **23** | **52** ✅ |

**Sem nenhum multiplicador, o 4º ataque mata um oponente de 40.** Se Otharri entra no T4 (com
ramp de 2 manas no T2), o 4º ataque é no **T7** — exatamente o limite da meta. **Não há folga:**
qualquer turno perdido (Otharri no T5, um wipe, um bloqueador) empurra para T8+.

Com os multiplicadores, a mesma tabela colapsa dois turnos:

| Cenário | 3º ataque (T6) | Quando mata 1 oponente |
|---|---|---|
| Nenhum multiplicador | 15 (acum. 29) | **T7** |
| + `Neyali` (double strike nas fichas atacantes — **estática**) | 6 fichas × 2 × 2 = 24 + 3 = **27** (acum. 47) | **T6** |
| + `Jor Kadeen` metalcraft (+3/+0, fichas viram 5/2) | 6 × 5 = 30 + 6 + 8 = **44** | **T6** |
| + `Neyali` **e** `Jor Kadeen` (fichas 5/2 double strike) | 6 × 5 × 2 = 60 | **T5–T6** |

**Conclusão que orienta toda a fase:** o deck não precisa de *mais corpos* — precisa de **2 a 3
multiplicadores estáticos** e de **furo de bloqueio**. Um único slot de multiplicador vale mais
que cinco slots de criatura.

### 1.1 O que a matemática acima **esconde** — e é a dor do usuário

A tabela assume **todo o dano passando**. Fichas 2/2 terrestres são a coisa mais fácil de parar
do jogo: três bloqueadores 4/4 absorvem três fichas **e as matam**, e um único `Fog` zera o
turno. A Rota 1 sozinha é uma rota **com um único ponto de falha**. É por isso que as Rotas 2 e 3
existem — e é por isso que elas são o conteúdo novo desta fase, não a Rota 1.

---

## 2. Caminhos de vitória atuais

Avaliação do estado **hoje** — `lista.txt` (99 cartas) com Otharri no comando, **sem** nenhuma
entrada das fases 2–5 aprovada.

| Caminho | Cartas envolvidas (hoje) | Turno estimado | Consistência |
|---|---|---|---|
| **1 · Dano de combate largo** | Otharri (motor) + `The Circle of Loyalty` (+1/+1 estático) + `Jor Kadeen` (+3/+0 metalcraft) + `Valor in Akros` (N gatilhos de entrada) + `Tori D'Avenant` (trample + pump) + `Crash Through` (trample) | **T7–T8** | **Média-baixa.** O motor é o comandante — sem ele, nada acontece. `Jor Kadeen` depende de **3 artefatos** e o deck tem 7 (dos quais `Sol Ring`/`Signet`/`Circle` são os confiáveis). Trample vem de 2 cartas, uma delas cantrip de uso único |
| **2 · Dano não-combate** | **nenhuma carta funcional.** `Vigilante Justice` lê **Human** e o enxame é **Rebel vermelho** — nunca dispara | — | **Zero.** É o buraco que esta fase preenche |
| **3 · Alfa desbloqueável** | `Bond of Discipline` (tapa todas as criaturas dos oponentes) — **1 carta** | T6+ | **Baixa.** Uma cópia, CMC 5, feitiço. Nota: o `theme-analyst` propôs cortá-la em **K16** — ver §6 |
| **4 · Combate extra** | `Éomer, Marshal of Rohan` (desvira tudo + combate extra quando uma legendária **atacante sua morre**) | T6+ | **Baixa.** Condicionado à morte de uma legendária atacante |
| **5 · Alt-win** | nenhuma | — | Zero |

**Veredito:** o deck tem **uma** rota real e **quatro quinze-avos de uma segunda**. A queixa do
usuário está mecanicamente correta e não é impressão.

---

## 3. Caminhos de vitória propostos

| Rota | Falha quando… | Peças-chave | Turno-alvo |
|---|---|---|---|
| **1 · Alfa em massa com multiplicador estático** | há bloqueadores grandes, `Fog`, ou o board é varrido | **Neyali** · `Jor Kadeen` · `The Circle of Loyalty` · `Warleader's Call` (anthem) · `Valor in Akros` | **T6–T7** |
| **2 · Dano agregado não-combate** | nunca — **ignora bloqueador, voo, trample, `Fog` e a fase de combate inteira** | **Warleader's Call** · **Molten Gatekeeper** · **Witty Roastmaster** | **T7–T8** sozinha; **T6** somada à Rota 1 |
| **3 · Alfa desbloqueável (limpar/furar a defesa)** | o oponente tem proteção contra remoção em massa | **Goblin War Drums** (menace estático) · **Barrage of Boulders** · `Winds of Abandon` overload · `Bond of Discipline` · `Hour of Reckoning` | **T6–T7** |
| **4 · Combate extra (multiplicador das rotas 1 e 3)** | não há board para atacar | `Éomer, Marshal of Rohan` · `Combat Celebrant` · `Great Train Heist` | **T5–T6** quando aparece |

**Por que estas três (1, 2, 3) são rotas de verdade e não a mesma rota pintada de outra cor:**
elas falham por motivos **diferentes**. Um `Fog` zera a 1 e a 3 e **não toca** a 2. Um muro de
bloqueadores para a 1 e **não toca** a 2 nem a 3. Um board wipe atrasa as três — mas os
contadores de experiência ficam no jogador e `Warleader's Call` (encantamento) e
`Molten Gatekeeper` (**unearth `{R}`**) voltam sozinhos.

### 3.1 A matemática da Rota 2 — a que o usuário nunca teve

`Warleader's Call` · `Molten Gatekeeper` · `Witty Roastmaster` disparam **em "criatura entra"**,
não em ataque. Cada ficha que entra = **1 de dano a cada oponente, por pingador**.

| Ataque nº | Fichas novas | Dano a **cada** oponente com 3 pingadores | Acumulado a **cada** oponente |
|---|---|---|---|
| 1 | 1 | 3 | 3 |
| 2 | 2 | 6 | 9 |
| 3 | 3 | 9 | 18 |
| 4 | 4 | 12 | **30** |
| 5 | 5 | 15 | **45** ✅ |

**30 de dano a cada um dos três oponentes até o 4º ataque, sem declarar um único bloqueio
resolvido.** Num pod, são **90 de dano total** que nenhuma das defesas usuais toca. Somada aos
5+9+15+23 = 52 de combate da Rota 1 contra um alvo, a mesa acaba no **T6–T7** mesmo com metade
do combate bloqueado. E a Rota 2 **não pede que o Otharri conecte** — só que ele ataque.

---

## 4. Finishers recomendados

Origem: `lista.txt` = já existe fisicamente (custo zero) · `caixa` = sobressalente (custo zero) ·
`compra` · `fase N` = **já proposta por outra fase** (não consome slot novo nem dinheiro novo).
Preços US$ são **estimativa (Scryfall)** — **não são régua de orçamento** (regra 2).

### 4.1 Entradas que eu peço de verdade — **4 cartas novas**

O deck está com **mais candidatas do que slots**: as fases 2–5 já propuseram ~40 entradas para
~33 slots. Peço **quatro**, e as quatro cobrem funções que **nenhuma** das 40 cobre.

| Carta | CMC | Como fecha o jogo | Liga em qual ataque? | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|
| **Molten Gatekeeper** `{2}{R}` **Artifact Creature — Golem 2/3** | 3 | **Rota 2.** *Whenever another creature you control enters, this creature deals 1 damage to **each opponent***. N fichas = **N de dano em cada um dos 3 oponentes**. Dano que não passa por bloqueio, voo, trample nem `Fog` | **1º ataque** — gatilho de *enters*, dispara N vezes | (a) **é artefato** → 3ª peça do **metalcraft do `Jor Kadeen`** (+3/+0 no enxame) e contador do `Luxknight Breacher`; (b) **Unearth `{R}`** → volta do cemitério por **1 mana** depois do wipe, ataca a **dor 4** no mesmo slot; (c) corpo 2/3 que **recebe todos os anthems** e bloqueia o crack-back; (d) dispara também com as criaturas conjuradas, não só com fichas. **4 eixos** | **não** — a caixa não tem nenhum pingador de *enters* (ver §5) | US$ 0,25 *(est. Scryfall)* |
| **Witty Roastmaster** `{2}{R}` **Creature — Devil Citizen 3/2** | 3 | **Rota 2.** Texto **idêntico** ao do Gatekeeper. Dois pingadores dobram a curva da Rota 2 (de 2 para 3 de dano por ficha, com o `Warleader's Call`) | **1º ataque** | (a) **poder 3** → com qualquer anthem vira **poder 4+**, que é a condição *ferocious* do `Barrage of Boulders` abaixo; (b) recebe todos os anthems e a **double strike da Neyali não** (não é ficha) mas recebe `Jor Kadeen`; (c) corpo 3/2 vermelho que ataca junto | **não** | US$ 1,37 *(est.)* |
| **Goblin War Drums** `{2}{R}` Enchantment | 3 | **Rota 3.** *Creatures you control have **menace*** — **estática**. Cada ficha 2/2 passa a exigir **dois** bloqueadores. Com 10 fichas o oponente precisa de **20 bloqueadores** para segurar o alfa. É o conserto direto da frase *"monta board, ataca, não converte"* | **1º ataque** — estática, alcança as fichas que entram atacando | (a) o único mass-menace barato dentro de RW (busca `bin/mtgdb search '"creatures you control have menace"' -id RW`: só ele, `Iroas` e um Sliver); (b) **não é criatura** → sobrevive ao board wipe e continua valendo quando o Otharri reconstrói; (c) multiplica com trample da `Tori`/`Crash Through`: menace **e** trample no mesmo corpo é quase inbloqueável | **não** | US$ 1,47 *(est.)* |
| **Barrage of Boulders** `{2}{R}` Sorcery | 3 | **Rota 3, botão de fechar.** *Deals 1 damage to each creature you don't control.* **Ferocious — se você controla uma criatura de poder 4+, as criaturas não podem bloquear neste turno.* É o turno em que **todo o dano passa** | **1º ataque** — feitiço da fase principal pré-combate; o efeito é **do turno** e alcança as fichas que entram depois | (a) o 1 de dano **varre os x/1** do oponente e **os enxames de fichas 1/1** adversários — remoção em massa unilateral de graça; (b) *ferocious* é **trivial** neste deck: `Jor Kadeen` (5/4), `Éomer` (4/4), `Relentless Rohirrim`, ou qualquer ficha 2/2 sob `Jor Kadeen` (5/2) / Otharri sob `Circle of Loyalty` (4/4); (c) **3 manas** contra os 5 do `Bond of Discipline` — sobra mana para uma proteção no mesmo turno | **não** | US$ 0,06 *(est.)* |

**Custo somado das 4: US$ 3,15 — `estimativa (Scryfall)`.** **Não afirmo que cabem no teto de
R$ 200.** Nenhuma das quatro tem cotação LigaMagic registrada (`bin/mtgdb prices` conferido em
2026-09-19: "sem cotação registrada" para as quatro). A captura é manual e precisa ser feita
antes de qualquer total (regra 2). Registro o alerta: erros do proxy medidos em 2026-08-12 foram
de **6,5× para mais** e **14× para menos** — `Barrage of Boulders` a US$ 0,06 pode perfeitamente
sair a R$ 8 na LigaMagic, e `Goblin War Drums` é carta de edições antigas, a faixa de maior risco.

### 4.2 Finishers que **já estão nas candidatas de outra fase** — endosso, custo zero de slot novo

| Carta | CMC | Como fecha o jogo | Liga em qual ataque? | Proposta em | Prioridade |
|---|---|---|---|---|---|
| **Neyali, Suns' Vanguard** | 4 | **A carta mais importante do deck depois do comandante.** *Attacking tokens you control have **double strike*** é **estática** → **dobra o dano do enxame inteiro**, incluindo as fichas que entram atacando. Sozinha antecipa a morte de um oponente do T7 para o **T6** | **1º ataque** (a estática). A 2ª habilidade — saque — é `whenever one or more tokens attack a player` e só liga do **2º ataque** em diante (fichas declaradas) | **Fase 2 (#1)** e **Fase 3 (D8)** | **Titular inegociável.** Se sair uma carta desta fase toda, que não seja esta |
| **Warleader's Call** | 3 | **Chave da Rota 2 e anthem da Rota 1 no mesmo slot.** `Creatures you control get +1/+1` (estático) **+** *whenever a creature you control **enters**, deals 1 damage to **each opponent*** | **1º ataque**, N vezes | **Fase 2 (#12)** | **Titular inegociável.** É `Impact Tremors` + anthem por 1 mana a mais — e `Impact Tremors` está cotado a **R$ 17,78** (LigaMagic, 2026-08-22) |
| **Winds of Abandon** (overload `{4}{W}{W}`) | 2 / 6 | **Rota 3 no modo nuclear.** Exila **todas** as criaturas que você **não** controla; o seu board fica inteiro de pé e ataca no mesmo turno. É o `Plague Wind` unilateral | Feitiço pré-combate — **1º ataque** | **Fase 5** | **Titular.** É o único wipe do pool que é literalmente um finisher |
| **Hour of Reckoning** | 7 (convoke) | **Rota 3.** `Destroy all nontoken creatures` — o enxame **sobrevive inteiro**, os bloqueadores não. Convoke pago pelas fichas dos ataques anteriores | Exige fichas **desviradas** para o convoke → **do 2º ataque em diante**. Mata Otharri/Neyali/Jor Kadeen (custo declarado; os contadores ficam) | **Fase 5** | **Titular se couber**; é a resposta a um board travado |
| **Combat Celebrant** | 3 | **Rota 4.** Exert ao atacar → *untap all **other** creatures you control* + combate adicional. Ruling 2017-04-18: *"untaps **all** of your creatures, not just the ones that are attacking"* → as fichas que acabaram de entrar viradas **desviram**, saem de combate no fim da fase e **podem ser declaradas atacantes no combate extra** — com a `Neyali` em campo, elas atacam com **double strike duas vezes no mesmo turno** | Ela é declarada atacante → **liga no 1º ataque dela**. Ordene o gatilho do Otharri por **último** na pilha para que as fichas existam quando o untap resolver | **Fase 2 (#22)** | **Alta.** É o maior salto de dano por slot depois da Neyali |
| **Great Train Heist** | 1+spree | **Rota 4 em instantâneo.** Modo `{2}{R}`: desvira todas + combate adicional. Modo `{2}`: +1/+0 e first strike ao time | Feitiço/instantâneo — o combate extra liga no turno em que é conjurado | **Fase 2 (#23)** | Média-alta (alternativa ao Celebrant, não somam bem) |
| **Boros Charm** | 2 | **Alcance final.** 4 de dano ao rosto fecha os últimos pontos; ou double strike no Otharri; ou indestrutível aos **permanentes** | Instantâneo | **Fase 5** | Média |

### 4.3 Já em `lista.txt` — finishers de custo **zero** que eu peço para **manter**

| Carta | Papel de fechamento | Liga em qual ataque? |
|---|---|---|
| **Jor Kadeen, the Prevailer** | Metalcraft `+3/+0` **estático**: 6 fichas 2/2 viram **5/2**. É o maior multiplicador gratuito do deck. Depende de 3 artefatos — `Molten Gatekeeper` (§4.1) é a 3ª peça, e a Fase 4 traz 5+ rochas | **1º ataque** (estática) |
| **The Circle of Loyalty** | `+1/+1` estático a todas **e** é artefato (metalcraft). Affinity for Knights derruba o custo | **1º ataque** (estática) |
| **Valor in Akros** | `Whenever a creature you control enters, creatures you control get +1/+1 until end of turn` — **N fichas = N gatilhos**. Um ataque com 4 contadores é **+4/+4 no time inteiro** naquele combate | **1º ataque**, N vezes |
| **Éomer, Marshal of Rohan** | Combate extra quando uma legendária atacante sua morre — e **a morte do Otharri é desejável** (ele volta do cemitério por `{2}{R}{W}` + virar um Rebel) | Condicionado à morte |
| **Crash Through** | Trample a todas por 1 mana, **com cantrip**. Menace + trample juntos = o enxame praticamente não é parável | **1º ataque** (efeito do turno) |
| **Tori D'Avenant, Fury Rider** | +1/+1 a todos os outros atacantes **e trample às criaturas vermelhas atacantes** — as fichas **são** vermelhas. Por ordenação de pilha, alcança as fichas novas | **1º ataque** (ordenando o gatilho do Otharri por último) |
| **Bond of Discipline** | Tapa **todas** as criaturas dos oponentes + lifelink ao time. Rota 3 — ver a contestação do corte K16 em §6 | **1º ataque** (feitiço pré-combate) |

---

## 5. Varredura do pool sem custo, **antes** do Scryfall (regra 7)

### 5.1 As 69 não-básicas de `lista.txt`
Varridas na íntegra. **Sete** exercem função de fechamento e estão em §4.3. As demais são pump
temporário do eixo antigo (`Basri's Solidarity`, `Inspiring Roar`, `Pride of Conquerors`,
`Shoulder to Shoulder`, `Zealous Display`) ou typal de Human/Knight que **não enxerga o enxame
Rebel** — já protocoladas para corte pelas Fases 2–5 (K1, K2, K8, K9, K10…). Não reabro.

### 5.2 As 113 sobressalentes (`bin/mtgdb collection -list`) — veredito por escrito

| Carta da caixa | Função de fechamento avaliada | Veredito (regra 7 — dispensa justificada) |
|---|---|---|
| **Crash Through** (2ª cópia) | Trample em massa por 1 mana com cantrip | **Aproveitável.** Redundância barata da Rota 1/3. **Reserva** — entra se o `Goblin War Drums` estourar a cotação |
| **Mercadia's Downfall** | *Each attacking creature gets +1/+0 for **each nonbasic land** defending player controls* — com 10 atacantes e 4 não-básicos no oponente, são **+40 de dano por 3 manas**, em instantâneo, e alcança as fichas que entraram atacando | **Reserva com variância declarada.** Contra um oponente de manabase básica (comum no Commander 200, que é justamente um formato de orçamento) ela faz **zero**. Não pode ser rota, mas é um kill grátis. Custo R$ 0 |
| **Lightning Volley** | *Creatures you control gain "{T}: deals 1 damage to any target"* | **Dispensada — anti-sinergia direta.** As fichas do Otharri **entram viradas** e o resto do board está tapado atacando. A carta pede exatamente o recurso que o deck gasta. 1 ponto de sinergia, reprova na regra 3 |
| **Darksteel Reactor** | **Alt-win real:** 20 contadores de carga → *you win the game* | **Dispensada por velocidade.** Entra no T4, ganha 1 contador por upkeep → vitória no **T24**. Mesmo com as 3 fontes de proliferate do pool temático (`Karn's Bastion`, `Grateful Apparition`, `Contagion Clasp`, todas **compras**), são ~3 contadores/turno = **T11+**, quatro turnos depois da Rota 1. Um alt-win mais lento que a rota principal não é rota, é slot morto |
| **Lux Artillery** | 30 contadores entre artefatos/criaturas → 10 de dano a cada oponente | **Dispensada.** O deck não tem motor de contadores (só `Cathars' Crusade`, que é compra de US$ 10,78 e nem é titular). Sem ele, é um artefato de 4 manas que não faz nada |
| **Titan Forge** | `{3},{T}` por contador; 3 contadores → ficha 9/9 | **Dispensada.** **9 manas e 3 turnos** por um corpo, num deck que fabrica 4 corpos por ataque de graça. F7: consome o `{T}` e o mana que o enxame precisa |
| **Goblin Gathering** | 2 fichas 1/1 por 3 manas | **Dispensada como finisher** (2 corpos não fecham nada), mas registro: com os pingadores da Rota 2 são **6 de dano a cada oponente** por 3 manas. **Reserva de baixa prioridade** |
| **Skullcrack**, **Seismic Wave**, **Punishing Fire**, **Searing Barrage** | Alcance direto ao rosto | **Dispensadas.** 3–5 de dano a **um** alvo num formato de 3 oponentes com 40 de vida = 2,5% de uma vitória por slot. `Boros Charm` (Fase 5) já cobre o alcance com 3 funções no mesmo slot |
| **Smaug, the Great Calamity // Spew Flame** | 5/5 voador CMC 7 / remoção de 5 de dano | **Dispensada por curva.** CMC 7 num deck cuja curva já trava em 4–5 e que quer atacar no T4. A face de remoção é trabalho do `interaction-specialist`, não meu |
| **Knight Watch**, **Moment of Glory**, **Courage in Crisis** | Corpos / pump pontual | **Dispensadas.** Largura e pump pontual é o que o deck **já tem em excesso** e é justamente o que **não** converte (§2) |

**Conclusão da varredura:** a caixa **não tem nenhum pingador de "criatura entra"** nem nenhum
concessor de evasão em massa dentro de RW. As quatro compras de §4.1 são compras porque a função
não existe na caixa — não por preferência.

---

## 6. Escopo de corte — o que eu **não** corto, e a contestação de K16

Eu enxergo o deck por **uma** lente (fechamento). **Não proponho nenhum corte novo.** As Fases
2–5 já protocolaram 15+ cortes limpos (K1, K2, K7–K14, K16, K17, D1–D2, I2–I4) e os quatro slots
que eu peço cabem folgadamente neles. Registro apenas **um desacordo**, com ficha:

### K16 · `Bond of Discipline` — o `theme-analyst` propôs **corte limpo com custo declarado**

Ele escreveu: *"Falter → **descoberto**; custo aceito e declarado: a evasão passa a vir de
trample/menace permanentes (`Crash Through`, `Hexgold Halberd`, `Iroas`)"*.

**Minha posição — concordo com o corte, mas só porque as minhas entradas cobrem a função que ele
deixou descoberta.** Aplicando os **mesmos** critérios aos dois lados (§3 do checklist —
condições idênticas: Otharri com 3–4 contadores, 6–10 fichas, um anthem em campo):

| Eixo | **Sai:** `Bond of Discipline` `{4}{W}` | **Entra:** `Barrage of Boulders` `{2}{R}` |
|---|---|---|
| **F1** | Tapa **todas** as criaturas dos oponentes; suas criaturas ganham lifelink | 1 de dano a **cada criatura que você não controla**; *ferocious*: **criaturas não podem bloquear** neste turno |
| **F2** | Sem corpo | Sem corpo |
| **F3** | Feitiço; não alimenta contagem | Feitiço; não alimenta contagem |
| **F4** | Não recebe nada | Não recebe nada |
| **F5** | Falter total **+ lifelink em massa** (com 10 fichas, ~20 de vida) | Falter total **+ varredura de x/1 e de enxames de fichas adversários** |
| **F6** | **Turno 5** — compete de frente com o turno em que o Otharri quer entrar | **Turno 3** — e sobram 2 manas para `Gods Willing`/`Feat of Resistance` no mesmo turno |
| **F7** | 5 manas que **não somam board**; a condição é incondicional (nunca falha) | 3 manas; **a condição *ferocious* pode falhar** se o board estiver só com fichas 2/2 sem anthem. Risco real, mas pequeno: `Jor Kadeen`, `Éomer`, `Relentless Rohirrim` e qualquer anthem ligam-na |

**Protocolo (§2):** funções do `Bond of Discipline` = [Falter em massa, lifelink em massa].
- **Falter em massa** → **coberto** por `Barrage of Boulders` (mais barato, mais amplo, mesmo
  turno) e **ampliado permanentemente** por `Goblin War Drums` (menace estático, todo turno, não
  só uma vez).
- **Lifelink em massa** → **parcialmente coberto** pelo próprio Otharri (lifelink) e por
  `Guide of Souls` se ela entrar. **Fica descoberta a ordem de grandeza** (≈20 de vida num alfa).
  **Custo aceito e declarado:** o deck está correndo para vencer no T7, não para sobreviver ao T12.

> ⚠ **Condição explícita ao orquestrador:** se o usuário **rejeitar** `Barrage of Boulders` e
> `Goblin War Drums`, o corte K16 **volta a ser um corte com função descoberta** e minha
> recomendação passa a ser **manter `Bond of Discipline`**. Não existe cenário em que o deck fique
> com **zero** cartas de furo de bloqueio — é a dor nº 1 do usuário escrita em uma linha.

---

## 7. Combos

| Combo | Peças | Resultado | Confirmado por rulings? | Veredito |
|---|---|---|---|---|
| **Otharri + Combat Celebrant** | 2 (ambas já no pool) | Segundo combate no mesmo turno: **+1 contador de experiência** e uma segunda leva de fichas, **e** as fichas do 1º combate desviram e são declaradas atacantes (ganhando os gatilhos que "fichas entrando" não ganham). **Não é infinito** — exert impede o Celebrant de desvirar no próximo untap | ✅ `Combat Celebrant` 2017-04-18: *"untaps **all** of your creatures, not just the ones that are attacking"* · *"you get an additional combat phase even if Combat Celebrant doesn't survive"* | **Recomendado.** Sinergia forte, não é combo travante |
| **Otharri + Neyali + Jor Kadeen** | 3 | Fichas **5/2 com double strike** = 10 de dano por ficha. 6 fichas = 60. Não é infinito, é aritmética | ✅ estáticas, sem ruling necessário (§0) | **Recomendado** |
| **Fumigate / Hour of Reckoning + Duty Beyond Death / Unbreakable Formation** | 2 | Wipe **assimétrico**: você conjura o wipe e responde ao próprio wipe com indestrutível em massa | — (interação de regras trivial) | **Recomendado**; já mapeado pela Fase 5 |
| **Breath of Fury + qualquer ficha que conecte** | 2 | **Combates arbitrariamente muitos.** A aura sacrifica a criatura que conectou, se reanexa a outra, **desvira todas** e dá combate adicional. Com um enxame, o loop repete até acabarem os corpos ou os bloqueios | ✅ texto oracle puxado nesta sessão | **NÃO recomendado — e o motivo não é a regra 11.** Julgo e declaro: isto **não** é um lock (não impede ninguém de jogar; é um kill), então a regra 11 **não** o proíbe. Recuso por **função**: o loop exige uma criatura conectando **desbloqueada** — ou seja, ele só funciona no cenário em que a Rota 1 **já estava ganhando**. Não diversifica nada, e é uma aura que morre 2-por-1 para qualquer remoção. O deck chega no T6–T7 sem ele |
| **Aggravated Assault + ramp infinito** | 2+ | Combates infinitos | — | **Fora.** `Aggravated Assault` está a US$ 35,79 *(est.)* e o deck não tem o motor de mana; a Fase 2 já a marcou como provavelmente fora da régua |

**Declaração explícita (regra 11):** **o deck não leva nenhum combo infinito de duas cartas.**
Não por incapacidade — `Breath of Fury` está disponível e barato — mas porque as três rotas
propostas matam no T6–T7 por aritmética aberta, que é lido na mesa e permite resposta. Um deck
que ganha por combate largo com 4 rotas redundantes é mais difícil de odiar e mais divertido de
jogar contra do que um que ganha por loop.

---

## 8. Recusas por preço (registro, não veredito de função)

| Carta | Função | Preço | Por que fica fora |
|---|---|---|---|
| **Shared Animosity** | *Whenever a creature you control attacks, it gets +1/+0 for **each other attacking creature that shares a creature type***. Todas as fichas são **Rebel** → crescimento quadrático sobre o crescimento quadrático do Otharri | **R$ 25,99** (LigaMagic menor, 2026-08-22 — 28 d) | **13% do teto de R$ 200 em um slot.** E — crítico — é gatilho `whenever a creature you control **attacks**`: as fichas que **entram atacando não disparam**, só as dos ataques anteriores (§0). Liga só do **2º ataque** em diante. Potência real, mas paga caro por um efeito que chega tarde. **Reserva de luxo** |
| **Impact Tremors** | 1 de dano a cada oponente por criatura que entra | **R$ 17,78** (LigaMagic menor, 2026-08-22) | `Warleader's Call` faz **isto + anthem estático** por 1 mana a mais, e `Molten Gatekeeper`/`Witty Roastmaster` fazem isto **com corpo** por US$ 0,25/1,37 *(est.)*. Dispensa por preço/função combinados |
| **Iroas, God of Victory** | Menace em massa **+ previne todo dano aos seus atacantes** (o alfa nunca perde criatura) | US$ 12,38 *(est.)* | Função superior ao `Goblin War Drums` (menace **+** imunidade em combate **+** indestrutível). **Dispensado só por preço.** Se a cotação LigaMagic vier abaixo de ~R$ 20, ele **substitui** o War Drums, não soma |
| **Purphoros, God of the Forge** | 2 de dano a cada oponente por criatura que entra — o pingador perfeito | **R$ 129,99** (LigaMagic, 2026-08-22) | 65% do teto. Já dispensado pela Fase 2; reconfirmo |
| **Seize the Day** | Desvira uma criatura + combate adicional, **com flashback `{2}{R}`** = **dois** combates extras de uma carta | US$ 2,21 *(est.)* | **Não é recusa por preço — é por redundância.** `Combat Celebrant`, `Great Train Heist` e `Éomer` já cobrem a Rota 4 e duas delas desviram **todas** as criaturas, não uma. **Reserva** |

---

## 9. Protocolo de goldfishing

> O usuário **nunca rodou goldfishing em nenhum deck**. Este roteiro é escrito para ser executado
> na mesa da cozinha com o deck físico, papel e caneta, sem ferramenta nenhuma. Tempo estimado:
> **~15 minutos por partida**, ~2h30 no total. Rode em dois blocos de 5.

### 9.1 Preparo
1. Monte as 100 cartas. **Otharri fica na zona de comando, separado.**
2. Três oponentes imaginários, **40 de vida cada**. Anote os três na folha.
3. Você **não** simula as jogadas deles. Goldfishing é medir a **velocidade máxima** do deck.
4. Tenha 30 fichas/marcadores de algum tipo. Sem elas você vai errar a contagem — o deck chega a
   ter 15 fichas em campo no T7.

### 9.2 Roteiro por partida
1. **Embaralhe de verdade** (7 riffles ou 1 minuto de mash). Compre 7.
2. **Mulligan como numa partida real** (London: compre 7, embaralhe de volta, ponha N no fundo).
   Critério de manter: **2 a 5 terrenos** *e* ao menos uma de {fonte de ramp, peça de 2–3 manas
   jogável}. Mão de 1 ou 6+ terrenos: mulligan.
3. Jogue **sozinho, turnos consecutivos**: desvira, compra, terreno, ramp, comandante na curva,
   desenvolve, ataca. **Sempre ataque quando puder** — é o motor.
4. **Direcione todos os ataques ao Oponente A.** É a métrica que interessa ("dano letal em 1").
   Anote em paralelo o dano de Rota 2 (pingadores), que vai **nos três ao mesmo tempo**.
5. **Empilhe a pilha certo:** sempre que o Otharri e outro gatilho de ataque dispararem juntos,
   ponha o do **Otharri por último** (resolve primeiro). Se você empilhar errado, o resultado do
   teste fica errado — é a §0 deste documento.
6. Pare no **turno 10** mesmo sem vitória, e anote "não fechou".
7. **Três das dez partidas são de estresse** (marque antes de começar):
   - **Partidas 3 e 8 — teste da dor 4:** no **início do seu turno 6**, remova **todas** as
     criaturas do seu board (simula um `Wrath of God`) e continue jogando. Anote em que turno
     você volta a ter board bom e quanto atrasou a vitória.
   - **Partida 6 — teste da dor 1:** a partir do turno 5, o Oponente A tem **três bloqueadores
     4/4 desvirados** que bloqueiam sempre da forma mais eficiente. É o teste de **se as Rotas 2
     e 3 realmente existem**.

### 9.3 Ficha de registro (copie 10 vezes)

| Campo | O que anotar |
|---|---|
| **Mulligans** | 0, 1, 2… |
| **T do Otharri em campo** | turno em que ele foi conjurado |
| **T da 1ª ficha** | = turno do 1º ataque do Otharri |
| **T do 3º contador de experiência** | o marco de "o motor está girando" |
| **T da mesa boa** | Otharri + **4 ou mais fichas** + **1 multiplicador** (Neyali / Jor Kadeen / Warleader's Call / Circle) em campo |
| **T de 20 de dano no Oponente A** | metade da vida — o marco de meio de curva |
| **T da vitória projetada (1 oponente)** | turno em que o A chega a 0 |
| **T da vitória projetada (3 oponentes)** | turno em que os três chegam a 0, contando **o dano dos pingadores nos três** |
| **Travou mana?** | S/N — critério objetivo: **menos de 3 terrenos no T4**, ou **2 turnos consecutivos** sem baixar terreno tendo carta jogável presa na mão |
| **Mão morta?** | S/N — critério objetivo: **nenhuma jogada relevante até o fim do T4** (relevante = ramp, criatura, ou draw) |
| **Carta que ficou inerte na mão** | anote o nome. É esta coluna que vira a tabela de cortes da §10 |

### 9.4 Critérios objetivos de aprovação

| Métrica | ✅ Aprova | ⚠ Atenção | ❌ Reprova → encaminhamento |
|---|---|---|---|
| **Vitória projetada em 1 oponente** | **mediana ≤ T7** | T8 | **mediana ≥ T8** → trocar cartas "boas mas inertes" por finishers da §4.1/§4.2 |
| **Vitória projetada nos 3** | ≤ T9 | T10 | > T10 → faltam pingadores (Rota 2) |
| **Otharri em campo** | **≥ 7 de 10 partidas até o T5** | 5–6 de 10 | ≤ 4 de 10 → **devolver ao `ramp-specialist`** |
| **T do 3º contador** | ≤ T7 em 7 de 10 | T8 | ≥ T8 → o motor não gira; faltam proteção do Otharri e combate extra |
| **Taxa de mulligan** | **≤ 30%** (≤ 3 mãos em 10) | 40% | **≥ 50%** → **devolver ao `manabase-engineer`** (36 terrenos é −2 da base) |
| **Travas de mana** | ≤ 1 em 10 | 2 em 10 | **≥ 3 em 10** → **devolver ao `manabase-engineer`** |
| **Mãos mortas** | ≤ 1 em 10 | 2 em 10 | **≥ 3 em 10** → **devolver ao `draw-specialist`** + achatar a curva |
| **Partidas 3 e 8 (pós-wipe)** | vitória atrasa **≤ 2 turnos** | 3 turnos | **≥ 4 turnos** → a dor 4 não foi resolvida; mais peças não-criatura (`Assemble the Legion`, `Warleader's Call`, unearth do `Molten Gatekeeper`) |
| **Partida 6 (muro de 4/4)** | **vence até o T8 por Rota 2 ou 3** | T9–T10 | **não fecha** → **a dor nº 1 continua de pé.** Nenhuma outra métrica compensa esta |

> **A partida 6 é a mais importante do teste.** Ela é a reprodução literal da queixa do usuário.
> Se o deck construir 15 fichas e não conseguir matar um oponente com três 4/4 na frente, todo o
> resto do trabalho das fases 2–7 falhou, e a resposta é **mais Rota 2 e mais Rota 3**, não mais
> corpos.

### 9.5 O que fazer com os números
Preencha a tabela da §10 com as cartas que apareceram na coluna *"carta que ficou inerte na mão"*
em **3 ou mais** das 10 partidas. Essas são as candidatas reais a corte pós-teste — e elas vêm do
teste, não da opinião de nenhum agente. Traga a folha de volta ao orquestrador em modo
`post-goldfish`.

---

## 10. Ajustes pós-teste

*(A preencher pelo orquestrador em modo `post-goldfish`, com os registros da §9.3 em mãos.)*

| Sai | Entra | Motivo (métrica da §9.4 que falhou) |
|---|---|---|
| | | |

**Mapa de diagnóstico → ação, decidido antes do teste para não haver viés depois:**

| Sintoma medido | Diagnóstico | Ação |
|---|---|---|
| Vitória ≥ T8 **com** mesa boa no T5–6 | falta **multiplicador**, não largura | entram `Neyali` / `Warleader's Call` / `Jor Kadeen` ativo; saem cartas de largura pura |
| Vitória ≥ T8 **sem** mesa boa | falta **ramp** ou o Otharri não chega | devolver ao `ramp-specialist` |
| Partida 6 não fecha | falta **furo de bloqueio** | entram `Goblin War Drums`, `Barrage of Boulders`, `Winds of Abandon` |
| Partidas 3/8 atrasam ≥ 4 turnos | falta **resiliência não-criatura** | entram encantamentos/artefatos que reconstroem (`Assemble the Legion`, unearth) |
| Travas de mana ≥ 3/10 | manabase | `manabase-engineer` |
| Mãos mortas ≥ 3/10 | curva e draw | `draw-specialist` + achatar a curva de 4–5 |
