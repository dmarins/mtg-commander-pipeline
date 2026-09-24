# Análise Temática — Phlage, Titan of Fire's Fury (v2 · eixo de controle de atrito)

> **Natureza desta fase.** Não existe `deck.md`. O deck v1 (Otharri) foi reprovado e descartado,
> e a v1 nunca foi comprada. Portanto isto é um **`build` na prática**: pool temático do zero,
> com a `lista.txt` (69 não-básicas do Tori físico) e a caixa de sobressalentes entrando como
> pool sem custo, não como deck a auditar.
>
> **O que reaproveitei da v1 e o que descartei.** As fichas F1–F7 carta a carta de
> `rounds/v1-2026-09-18/02-theme.md` continuam válidas como *levantamento de fatos* e foram
> usadas. **Todos os vereditos de entrada/corte daquele arquivo foram descartados** — foram
> emitidos para um eixo go-wide que morreu.
>
> **Regra 5 — conferido.** `decisions.md` registra **apenas** dois movimentos, ambos de zona de
> comando (Otharri entra/sai, Tori sai). **Nenhuma carta do 99 tem histórico de corte.** Logo
> nenhuma proposta abaixo é reposição de corte anterior, e nenhuma carta está barrada.
>
> **Regra 2 — preço.** Nenhum número do Scryfall em lugar nenhum. `bin/mtgdb prices` foi
> consultado para as 48 candidatas principais: **11 têm cotação registrada**, as demais saem
> como **`a cotar`**.
>
> **Regra 6 — oracle.** Todo texto citado foi puxado com `bin/mtgdb oracle` nesta sessão. Os
> 9 rulings do Phlage foram puxados com `bin/mtgdb oracle -rulings`.

---

## 1. Comandante — análise linha a linha

**[**Phlage, Titan of Fire's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Phlage%2C+Titan+of+Fire%27s+Fury)** — `{1}{R}{W}` · CMC 3 · Legendary Creature — Elder Giant · **6/6**

| # | Linha do oracle | Gatilho/termo | O que habilita |
|---|---|---|---|
| L1 | *When Phlage enters, **sacrifice it unless it escaped**.* | `enters`, `sacrifice`, `escaped` | Conjurado da zona de comando ele **vai para o cemitério no mesmo turno**. Isso não é custo — é **a primeira carta do combustível**. Termo: **auto-abastecimento**. |
| L2 | *Whenever Phlage **enters or attacks**, it deals **3 damage to any target** and you **gain 3 life**.* | `enters`, `attacks`, `damage any target`, `lifegain` | Lightning Helix recorrente: **remoção** (3 numa criatura), **alcance** (3 na cara) ou **estabilização** (3 de vida). Dispara **sempre**, inclusive na conjura da zona de comando (ruling 2024-06-07: *"triggers when it enters the battlefield, even if it didn't escape"*). Termos: **remoção pontual**, **dano direto**, **multiplicador de dano**. |
| L3 | *Escape—**{R}{R}{W}{W}**, Exile **five other cards** from your graveyard.* | `escape`, `graveyard`, `exile`, `custo fixo` | Reconjura **do cemitério** por preço **fixo** — o imposto de comandante não incide (ruling: escape é permissão de conjuração de outra zona; o imposto só existe para a zona de comando). Termos: **combustível de cemitério**, **loot**, **self-mill**, **auto-descarte**. |
| L4 | Corpo **6/6** (sem palavras-chave impressas) | `corpo grande`, `bloqueio` | 6/6 por 4 manas efetivos no escape. É o maior corpo que o deck vai ter e o único bloqueador que importa. Termo: **receptor de buff/evasão** (trample e lifelink **não** são impressos — a memória popular erra aqui; conferido no oracle). |

### 1.1 As três travas mecânicas (dadas pelo orquestrador — respeitadas, não redescobertas)

1. **6 cartas no cemitério** para o primeiro escape (5 exiladas **além** do Phlage).
2. **Só o escape evita o sacrifício.** Reanimar ou blinkar dispara L1. → **Nenhum pacote de
   reanimação ou blink neste deck.** Consequência direta e concreta: [**Late to Dinner**](https://www.ligamagic.com.br/?view=cards/card&card=Late+to+Dinner),
   [**Miraculous Recovery**](https://www.ligamagic.com.br/?view=cards/card&card=Miraculous+Recovery) e [**Remember the Fallen**](https://www.ligamagic.com.br/?view=cards/card&card=Remember+the+Fallen) (as três na `lista.txt`) estão **fora do pool**
   — as duas primeiras devolvem permanente ao campo (dispara o sacrifício se mirarem o Phlage) e
   a terceira devolve o Phlage **à mão**, de onde ele só pode ser conjurado normalmente e
   também se sacrifica. As três ainda **tiram carta do cemitério**, que é combustível.
3. **O Helix acontece na conjura da zona de comando.** O T3 nunca é um turno morto: 3 de dano +
   3 de vida garantidos, e o Phlage cai no cemitério onde ele quer estar.

### 1.2 O que o escape custa de verdade — o número que reordena o deck

O ponto que a análise da Fase 1 não dimensionou: **cada escape exila 5 cartas**. O cemitério
**zera a cada reconjura**. Isso significa que combustível não é um problema de abertura — é um
problema **de fluxo contínuo**. Um pacote de 10 *one-shots* resolve o primeiro escape e não
resolve o terceiro; motores repetíveis resolvem os dois. Daí a divisão que proponho na §3.

---

## 2. Termos de busca do tema

Os 7 termos que definem o eixo, derivados da §1:

| # | Termo | Busca (sempre `legal:commander id<=rw`) |
|---|---|---|
| T1 | **combustível de cemitério — motor repetível** | `o:mill -o:"target player"` · `o:"Retrace"` · `otag:self-mill` |
| T2 | **combustível de cemitério — auto-descarte + saque** | `o:discard o:draw` (fatiado por `mv<=2` e `mv=3`) |
| T3 | **remoção/queima barata** (o eixo *é* isto) | `otag:removal mv<=3 (t:instant OR t:sorcery)` · `otag:burn mv<=3` |
| T4 | **wipe assimétrico** | `otag:boardwipe mv<=6` |
| T5 | **alcance sem combate** (mágica → dano a cada oponente) | `o:"damage to each opponent"` + `o:"whenever you cast"` |
| T6 | **multiplicador de dano** | `o:"deals double"` / `o:"triple that damage"` / `o:"plus 2 instead"` |
| T7 | **auto-recursão de mágica** (carta que volta do cemitério sozinha) | `o:Flashback` · `o:Retrace` · `o:Escape` |

**Cruzamentos que produziram as melhores cartas** (regra 3 — sinergia sobreposta): **T2×T3**
([**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre): remoção **ou** loot no mesmo card) · **T1×T5** ([**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation): moe 1
por turno **e** queima cada oponente) · **T3×T7** ([**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab): ping repetível **e** despeja uma
carta no cemitério a cada uso) · **T2×T3** ([**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore): custo adicional de descarte +
ping repetível + persist) · **T5×T6** (pingador + Torbran).

---

## 3. Combustível de cemitério — o pacote que define o eixo

> **Este pacote não existia em nenhuma lista anterior.** Nem a `lista.txt` do Tori nem a v1 do
> Otharri têm uma única carta cuja função seja encher o próprio cemitério.

### 3.1 A conta

Simulação de Monte Carlo (30.000 mãos por configuração, `on the play`, modelo com 36 terrenos,
9 rochas de mana, curva real das queimas, loots que compram e descartam, motores moendo a partir
do turno seguinte, Phlage conjurado da zona de comando assim que houver 3 manas). A métrica é
**P(≥5 cartas *além* do Phlage no cemitério)** = P(escape disponível).

| Config | burn | loot | motores | **T4** | **T5** | **T6** | **T7** |
|---|---|---|---|---|---|---|---|
| **A** — sem pacote (perfil da v1) | 20 | 2 | 0 | 1,7% | **7,1%** | 12,4% | 17,8% |
| **B** — leve | 17 | 6 | 2 | 6,9% | 22,2% | 34,8% | 45,0% |
| **C** — médio (**piso**) | 15 | 8 | 3 | 8,5% | 27,7% | 43,6% | 54,5% |
| **D** — pesado (**meta**) | 14 | 10 | 4 | 11,1% | **35,5%** | **52,4%** | 64,1% |
| **E** — máximo (**teto**) | 13 | 12 | 5 | 13,9% | 41,7% | 59,8% | 71,7% |

**O número: 14 cartas de combustível** — 10 de auto-descarte/loot + **4 motores repetíveis**.

**Leitura honesta, e ela corrige a premissa do briefing:** com 14 fontes, **o primeiro escape
confiável é o T6, não o T5.** No T5 ele sai em pouco mais de um terço das partidas. Quem promete
escape no T5 está prometendo o percentil 35. O que o pacote de fato compra é a **diferença entre
12,4% e 52,4% no T6** — um fator de 4,2× — e é isso que separa "comandante que volta" de
"comandante que ficou no cemitério a partida inteira".

Por que **não** ir até a config E: os 3 slots extras saem de 1 queima e 2 criaturas, e compram
+6,2 pp no T5. Com o piso de corpos apertado (§4), esse câmbio é ruim. **D é o ponto de parada.**

O modelo é **conservador de propósito** — não conta: (a) o Phlage voltando ao cemitério a cada
morte após o primeiro escape, (b) criaturas mortas em combate e em wipe, (c) terrenos com
cycling, (d) descarte de terreno excedente pelos loots. O número real fica **acima** da tabela;
eu prefiro errar para baixo no dimensionamento.

### 3.2 Por que 4 motores e não 14 *one-shots*

Cada escape **exila 5 cartas**: o cemitério volta a zero. Um *one-shot* abastece uma vez; um
motor abastece todo turno, para sempre.

| Cenário (pós-escape #1) | Abastecimento/turno | Turnos até o escape #2 |
|---|---|---|
| Só *one-shots* (0 motores em campo) | ~1,2 carta (mágicas conjuradas) | **4–5 turnos** |
| 1 motor em campo | ~2,5 cartas | **~2,5 turnos** |
| 2 motores em campo | ~4 cartas | **~1,5 turno** |

P(≥1 motor em jogo até o T6) com 4 motores em 99 = **40,9%**; com 6 = **54,9%**. Por isso o pool
abaixo traz **6 candidatos a motor para 4 slots** — a Fase 6 corta, mas precisa cortar de uma
reserva.

### 3.3 Lista do combustível (o pacote nominal)

**Motores repetíveis (6 candidatos → 4 slots)**

| Carta | CMC | Tipo | Cartas/turno no cemitério | Origem |
|---|---|---|---|---|
| [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) | 2 | Enchantment RW | 1 (end step, automático) | compra — `a cotar` |
| [**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece) | 2 | Artifact | 2 (só tapar) | compra — `a cotar` |
| [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) | 1 | Sorcery R (Retrace) | 1 (o terreno descartado), a pedido | compra — `a cotar` |
| [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) | 2 | Artifact Creature 0/1 | 1 **+ 1 mana** | compra — **R$ 0,17** (12/08) |
| [**Tablet of Discovery**](https://www.ligamagic.com.br/?view=cards/card&card=Tablet+of+Discovery) | 3 | Artifact R | 1 no ETB + rocha de 2 manas p/ mágicas | compra — `a cotar` |
| [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) | 2 | Artifact Creature 2/2 | 1 no cast (custo adicional) + ping repetível + persist | **caixa** |

**Auto-descarte / loot (10 slots)** — todos são **simultaneamente fontes de saque**, e é por isso
que este pacote cabe no deck: ele resolve a dor 2 no mesmo slot em que resolve o escape.

| Carta | CMC | Cartas p/ o cemitério | Compra | Origem |
|---|---|---|---|---|
| [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) | 2 (instant) | 2 (ela + 1 descarte) | +2 | **caixa** |
| [**Cathartic Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Reunion) | 2 | 3 (ela + 2) | +3 | compra — `a cotar` |
| [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) | 1 | 3 (ela + 2) | +2 | compra — `a cotar` |
| [**Tormenting Voice**](https://www.ligamagic.com.br/?view=cards/card&card=Tormenting+Voice) | 2 | 2 | +2 | compra — `a cotar` |
| [**Wild Guess**](https://www.ligamagic.com.br/?view=cards/card&card=Wild+Guess) | 2 | 2 | +2 | compra — `a cotar` |
| [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) | 2 (instant) | 1 ou 3 (modal) | +2 no modo loot | compra — `a cotar` |
| [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers) | 2 (instant) | 2 | +2 | compra — `a cotar` |
| [**Electric Revelation**](https://www.ligamagic.com.br/?view=cards/card&card=Electric+Revelation) | 3 (instant) | 2 + flashback | +2 (×2 com flashback) | compra — `a cotar` |
| [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) | 4 (instant) | 2 | +2 **e 2 Treasures** | compra — `a cotar` |
| [**Bitter Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Bitter+Reunion) | 2 (enchantment) | 2 | +2, e sacrifica p/ haste | compra — `a cotar` |
| [**Faithless Salvaging**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Salvaging) | 2 (instant) | 2, **duas vezes** (Rebound) | +1 ×2 | compra — `a cotar` (reserva) |
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | X | 3 no flashback (descarta 3) | — (é queima) | compra — `a cotar` |

### 3.4 Atrito interno do pacote (F7) — três coisas que precisam estar escritas

1. **Flashback exila.** [**Firebolt**](https://www.ligamagic.com.br/?view=cards/card&card=Firebolt), [**Electric Revelation**](https://www.ligamagic.com.br/?view=cards/card&card=Electric+Revelation), [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate), [**Devil's Play**](https://www.ligamagic.com.br/?view=cards/card&card=Devil%27s+Play),
   [**Strike It Rich**](https://www.ligamagic.com.br/?view=cards/card&card=Strike+It+Rich) **saem do cemitério** quando você usa o flashback. Elas são combustível
   *antes* de serem flashbackeadas, nunca depois. Não empilhar mais do que 3 delas.
2. **[**Containment Construct**](https://www.ligamagic.com.br/?view=cards/card&card=Containment+Construct) está fora do pool** apesar de parecer feita para o deck: ela exila
   a carta descartada. É anti-combustível puro.
3. **[**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece) tem uma segunda habilidade que embaralha o cemitério de volta.** É
   opcional e alvo escolhido — serve como **resposta a ódio de cemitério**, nunca como rotina.
4. **Ódio de cemitério** ([**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace), [**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog), [**Soulless Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Soulless+Jailer)) desliga o escape.
   Mitigação estrutural: o Phlage continua acessível pela zona de comando — você perde o corpo
   mas **não** perde o Helix. Sinal cruzado à Fase 5: 1–2 respostas a encantamento já cobrem.

---

## 4. Piso de corpos — o número, e por que ele é diferente do da v1

### 4.1 O contexto que não pode ser ignorado

A queixa nº 1 do teste de mesa da v1 foi, verbatim: *"fiquei com muita mágica instantânea e
encantamentos na mão e poucas criaturas"*. **O eixo do Phlage produz exatamente esse tipo de
mão por construção.** Fingir que isso desapareceu com a troca de comandante seria desonesto.
O que mudou é **o que uma mão sem criatura significa**.

| | v1 (Otharri) | v2 (Phlage) |
|---|---|---|
| Mão sem corpo no T1 | **jogo perdido** — sem board não há motor, e o motor só liga no T5 | **mão de controle normal** — você segura remoção e tem jogada garantida no T3 |
| Jogada garantida no T3 | nenhuma (comandante CMC 5) | **Phlage da zona de comando**: 3 de dano + 3 de vida, todo jogo, sem depender de compra |
| Corpos são o plano de vitória? | **sim** — o dano vinha do enxame | **não** — o dano vem de mágica multiplicada (§6) |
| Corpos são necessários para quê? | tudo | **bloquear** até o T6 e **carregar o pacote de alcance** |

### 4.2 O número

**Piso: 14 criaturas. Meta: 16–17.**

| Criaturas | P(mão de 7 **sem** criatura) | P(≥1 criatura até o T4) | P(≥2 até o T4) |
|---|---|---|---|
| 12 | 39,3% | 74,3% | 34,8% |
| **14 (piso)** | **33,2%** | 79,9% | 42,9% |
| **16–17 (meta)** | **27,9% – ~25,6%** | **84,4%** | **50,6%** |
| 20 | 19,5% | 90,8% | 64,3% |
| 25 (a v1 corrigida) | 12,1% | — | — |

**Defesa do 14 como piso — quatro razões, nenhuma delas "o eixo não precisa de criatura":**

1. **A métrica certa para "mão morta" não é criatura, é permanente jogável.** Com 14 criaturas +
   9 rochas/ramp + 4 motores + 3–5 outros permanentes = **~31 permanentes não-terreno**, a
   **P(mão de 7 sem nenhum permanente) é 6,5%** e **P(≥2 permanentes) é ~71%**. Isso é
   *melhor* que os 12,1% que a v1 alcançou depois das 7 correções — medindo a coisa certa.
2. **O T3 é garantido.** O Phlage sai da zona de comando **em toda partida**, independente de
   compra. A v1 não tinha nenhuma jogada garantida.
3. **Cada slot de criatura é dois slots.** 5–7 das 14–17 criaturas são o **pacote de alcance**
   ([**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) 0/4, [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) 0/3, [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer), `Kessig
   Flamebreather`, `Erebor Flamesmith`, `Guttersnipe`, `Syr Carah`): são simultaneamente o
   bloqueador barato **e** a condição de vitória. Um deck de controle com 17 criaturas em que
   7 delas são o fecho não é o mesmo deck que um go-wide com 17 criaturas.
4. **Os bloqueadores certos são baratos e largos.** [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) (0/4 defender por
   `{1}{R}`) e [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) (0/3 defender) param agressão de T2–T5 melhor do que os 2/2
   da `lista.txt` — e a v1 não tinha nenhum dos dois.

**Recomendação operacional às Fases 5–7:** mirem **16–17**, e usem o **piso de 14** como linha
que não se cruza em nenhum corte. Não subam para 20+ — cada criatura acima de 17 sai de um slot
de remoção, e a remoção é o eixo.

**⚠ Risco declarado ao orquestrador.** 27,9% de mãos sem criatura é **mais do que o dobro** dos
12,1% da v1 corrigida. Se o incômodo do usuário for **sensação de mão**, e não perda de partida,
este eixo **piora** o que a v1 consertou. Isso precisa entrar no `report.md` como pergunta
explícita ao usuário, não como nota de rodapé — ele já reprovou uma rodada inteira por
"jeito de jogar".

---

## 5. Densidade de remoção — o alvo numérico

O pipeline (`CLAUDE.md`) pede **~10 peças de interação + 2–4 wipes**. **Este deck não é esse
deck.** A remoção não é a resposta dele — é o corpo dele.

**Alvo: 20–22 peças de remoção/queima + 3–4 wipes.**

| Remoção no deck | P(≥1 na mão de 7) | P(≥2 até o T4) |
|---|---|---|
| 15 | 69,6% | 46,8% |
| 18 | 76,6% | 57,8% |
| **20–22 (alvo)** | **80,5% – 83,8%** | **64,3% – 70,2%** |
| 25 | 87,9% | 77,7% |

Abaixo de 20, o deck tem turnos em que não responde — e um deck sem tabuleiro próprio que não
responde perde. Acima de 22, ele fica sem fecho e sem combustível.

**Sobreposição declarada (não é dupla contagem):** as 20–22 remoções também caem no cemitério
quando conjuradas, e a simulação da §3.1 **já conta isso** (coluna `burn`). As 14 do pacote de
combustível são cartas cuja **função** é encher o cemitério — categoria diferente, slots
diferentes. Soma: 14 + 21 + 3,5 wipes ≈ **38,5 slots** dos 63 não-terreno.

**Seleção fina é da Fase 5.** O que passo como restrição dura:
- **Preferir instantâneo/feitiço a permanente** — mágica vai para o cemitério; artefato de
  remoção fica em campo e não abastece nada.
- **Preferir `{R}` a `{W}` onde empatar** — [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) só multiplica **fonte vermelha** (§6.2).
- **Curva:** ao menos 8 das 21 em CMC ≤ 2, para caber junto com o loot no mesmo turno.
- **2 respostas a encantamento** obrigatórias (ódio de cemitério, §3.4).

---

## 6. Finisher — o eixo **não** resolve a dor 1, e aqui está o pacote que resolve

O `01-commander.md` já declarou: *"Risco: relógio lento. O deck pode controlar a mesa por 12
turnos e não fechar — que é literalmente a dor 1."* O pool abaixo é a resposta.

### 6.1 A engrenagem: **pingadores de mágica** (5–7 slots, dentro dos 16–17 corpos)

Um deck com 21 remoções + 14 loots conjura **3–5 mágicas por turno** no meio de jogo. Cada
pingador converte isso em dano **em todos os oponentes de uma vez**, sem combate, sem
depender de tabuleiro. É o mesmo slot que já é bloqueador.

| Carta | CMC | Corpo | Dano por mágica, por oponente |
|---|---|---|---|
| [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) | 3 | 2/2 | **2** (instantâneo/feitiço) |
| [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) | 2 | 0/3 defender | 1 por tap, **destapa a cada mágica** → 1×N no turno |
| [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer) | 2 | 2/1 | 1 (qualquer não-criatura) |
| [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) | 2 | **0/4 defender** | 1 (instantâneo/feitiço) |
| [**Kessig Flamebreather**](https://www.ligamagic.com.br/?view=cards/card&card=Kessig+Flamebreather) | 2 | 1/3 | 1 (qualquer não-criatura) |
| [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) | 2 | 2/1 | 1 (instantâneo/feitiço) |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | 5 | 3/3 | **não pinga — compra**: exila o topo e deixa jogar sempre que uma mágica sua fere um jogador. É o motor de cartas do pacote. **Já está na `lista.txt`** |

P(≥1 pingador até o T6) com 6 no deck = **54,9%**; com 7 = 60,7%.

### 6.2 O multiplicador (3–4 slots)

P(≥1 multiplicador até o T7) com 3 = 34,8%; com 4 = **43,6%**. Recomendo **3–4**.

| Carta | CMC | Efeito | Helix do Phlage vira | Com 3 pingadores e 4 mágicas/turno |
|---|---|---|---|---|
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | **4** | fonte **vermelha** → +2 de dano a oponente/permanente dele | **5** | cada pingador de 1 vira **3** → 3×3×4 = **36 por turno, por oponente** |
| [**Solphim, Mayhem Dominus**](https://www.ligamagic.com.br/?view=cards/card&card=Solphim%2C+Mayhem+Dominus) | **4** | dobra dano **não-combate** a oponentes | **6** | dobra tudo do pacote; corpo 5/4 |
| [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) | 6 | **triplica** todo dano de fonte sua | **9** | Phlage atacando = **18 com um gatilho de 9 antes** |
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | 7 | **dobra** todo dano a oponentes (inclui combate) **e reduz pela metade o que você recebe** | **6** | Phlage = 6 (Helix) + 12 (combate). A metade defensiva é valor real num deck de atrito |

**Veredito sobre a `Gisela` (a candidata registrada em `decisions.md`):** **fica no pool, como
finisher nº 2, não nº 1.**
- **A favor:** é a única das quatro que dobra **também o dano de combate** (o Phlage 6/6 vira 12)
  **e** tem uma metade defensiva — reduzir pela metade tudo que você recebe é, num deck de
  controle que precisa chegar ao T7+, uma carta de sobrevivência disfarçada de finisher. Corpo
  5/5 voador com first strike bloqueia o céu, que é o buraco de RW.
- **Contra, e é o que a tira do nº 1:** **CMC 7**. Com 36 terrenos + 9 rochas ela é jogada real
  no T7–T8; é a carta mais cara do pool inteiro em mana. E o efeito é **simétrico para a mesa
  toda no ataque** — dobra o dano de *qualquer* fonte sua, mas a defesa só protege você.
- **Por que [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) é o nº 1:** mesmo tipo de efeito, **3 manas mais barato**, entra no mesmo
  turno em que o Phlage escapa, e a interação com o pacote de pingadores é multiplicativa
  (+2 sobre um dano-base de 1 é **+200%**, contra +100% da Gisela). **Ressalva registrada:**
  [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) lê **"red source"** — ele **não** multiplica remoção branca nem o dano de
  [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate)-likes. O Phlage **é** fonte vermelha (carta vermelha e branca), então o Helix
  conta. Isso é o motivo da restrição "preferir `{R}` a `{W}`" que passei à Fase 5.

**Alternativas de fecho fora do pacote de dano:**
- [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) — `{X}{X}{R}` dividido entre qualquer número de alvos, **flashback `{R}{R}`
  descartando X**. Com [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) é 3X. É wipe, finisher e combustível no mesmo card.
- [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) — 4 na cara **ou** indestrutível em todos os seus permanentes (é a resposta ao
  **seu próprio** wipe, §7) **ou** double strike no Phlage (12 de combate). Três modos, três
  funções.
- [**Sunforger**](https://www.ligamagic.com.br/?view=cards/card&card=Sunforger) — tutor de instantâneo RW ≤4 a cada `{R}{W}`: transforma o pacote de remoção em
  toolbox. **Ressalva F7:** ele **não** abastece cemitério (a mágica tutorada vai para o
  cemitério, ok, mas o Sunforger fica em campo) e o corpo equipado do deck é magro.

---

## 7. As quatro dores do intake — o que o eixo resolve

| Dor | Status sob o eixo Phlage | Quem fecha |
|---|---|---|
| **3. Sem respostas** | **Resolvida estruturalmente.** 20–22 remoções + 3–4 wipes é o deck inteiro. | Fase 5 (seleção fina) |
| **4. Morre para wipe** | **Resolvida estruturalmente.** Poucos permanentes a perder; **você** é quem dá o wipe; o comandante volta **do cemitério** e o wipe **enche** o cemitério. Ressalva honesta: os 4 motores e os 6 pingadores **morrem** no wipe — [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) (indestrutível) é a peça que faz o seu wipe ser assimétrico de verdade. | Fase 5 |
| **2. Mão morta** | **Resolvida no mesmo slot do combustível.** Os 10 loots compram 2–3 cada; [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) transforma dano de mágica em cartas grátis. Sinal duro à Fase 3: **8 a 10 das 12–13 fontes de draw devem ser loot ou impulse que também abastece o cemitério** — draw "limpo" (que não descarta) é slot desperdiçado neste deck. | Fase 3 |
| **1. Não fecha o jogo** | **NÃO resolvida pelo eixo.** Resolvida pelo pacote da §6: 5–7 pingadores + 3–4 multiplicadores. Sem esse pacote, o deck controla e não vence — e essa é a crítica nº 1 do usuário desde o intake. | **Fase 7** |

---

## 8. Varredura da caixa e da `lista.txt` (regra 7) — **antes** do Scryfall

Pool sem custo conferido com `bin/mtgdb collection -list` (116 cartas) + as 69 não-básicas de
`lista.txt`. Filtrei pelos 7 termos da §2.

### 8.1 Aproveitado — 21 cartas sem custo de compra

| Carta | CMC | Tipo | Função no eixo | Origem |
|---|---|---|---|---|
| [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) | 1 | Artifact | ramp · **intocável** | caixa + lista |
| [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet) | 2 | Artifact | ramp · **intocável** | lista |
| [**Command Tower**](https://www.ligamagic.com.br/?view=cards/card&card=Command+Tower) | — | Land | fixação · **intocável** | lista |
| [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) | 2 | Instant | **combustível (2 cartas) + saque 2, em instantâneo** | caixa |
| [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) | 2 | Art. Creature 2/2 | **combustível no cast** + ping repetível `{4}` + persist + corpo | caixa |
| [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) | 5 | Sorcery | **wipe assimétrico** (você tem poucos corpos) + vida por criatura | caixa |
| [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) | 1 | Sorcery | 4 de dano numa criatura por `{R}` — a remoção mais eficiente da caixa | caixa |
| [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) | 2 | Artifact | exila criatura MV≤3 — trava ameaças de T1–T3 | caixa |
| [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) | 2 | Instant | 3 em qualquer alvo (fonte **vermelha** → +2 com Torbran) | caixa |
| [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) | 1 | Instant | 2 + **exila** (mata recursão alheia) | caixa |
| [**Seal of Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Fire) | 1 | Enchantment | 2 pré-pagos; sai do campo e vai ao cemitério = combustível | caixa |
| [**Twin Bolt**](https://www.ligamagic.com.br/?view=cards/card&card=Twin+Bolt) | 2 | Instant | 2 divididos — mata dois x/1 | caixa |
| [**Fateful End**](https://www.ligamagic.com.br/?view=cards/card&card=Fateful+End) | 3 | Instant | 3 + scry 1 | caixa |
| [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) | 2 | Instant | 3 numa criatura, **remove indestrutível e exila** | caixa |
| [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) | 1 | Sorcery (Spree) | artefato **e** encantamento no mesmo card — resposta a ódio de cemitério | caixa |
| [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) | 2 | Artifact | wipe escalável por CMC — o deck tem curva baixa, os oponentes não | caixa |
| [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) | — | Land | wipe escalável **no slot de terreno** | caixa |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | 5 | Leg. Creature 3/3 | **motor de cartas do pacote de alcance** + ping `{T}` | lista |
| [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) | 2 | Sorcery | **spell mastery**: com 2+ mágicas no cemitério vira instantâneo — o deck satisfaz sempre | lista |
| [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) | 2 | Enchantment | remoção de artefato/encantamento pré-paga | lista |
| [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) | 2 | Instant | idem, em instantâneo | lista |
| [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) | 4 | Enchantment | exila 1 não-terreno **por oponente** + `recruit` (compra-e-descarta = **combustível**) | lista |
| [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) | 3 | Sorcery | exila **qualquer** não-terreno (a remoção mais ampla do pool sem custo) | lista |
| [**Expose to Daylight**](https://www.ligamagic.com.br/?view=cards/card&card=Expose+to+Daylight) | 3 | Instant | artefato/encantamento + scry (reserva do [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant)) | lista |

Rochas de mana da caixa ([**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket), [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere), [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns), [**Manalith**](https://www.ligamagic.com.br/?view=cards/card&card=Manalith),
[**Prophetic Prism**](https://www.ligamagic.com.br/?view=cards/card&card=Prophetic+Prism), [**Moss Diamond**](https://www.ligamagic.com.br/?view=cards/card&card=Moss+Diamond), [**Seer's Lantern**](https://www.ligamagic.com.br/?view=cards/card&card=Seer%27s+Lantern), [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler)) ficam para a **Fase 4** —
registro apenas que [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) e [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) **se sacrificam para comprar**, o que
os coloca no cemitério e os torna preferíveis às demais neste eixo.

### 8.2 Dispensas da caixa e da `lista.txt` — justificadas por escrito (regra 7)

| Carta | Origem | Ficha resumida | Por que **não** cobre a função |
|---|---|---|---|
| [**Late to Dinner**](https://www.ligamagic.com.br/?view=cards/card&card=Late+to+Dinner), [**Miraculous Recovery**](https://www.ligamagic.com.br/?view=cards/card&card=Miraculous+Recovery) | lista | F1 reanimação · F6 CMC 4/5 | **Trava mecânica 2.** Devolvem permanente ao campo → se mirarem o Phlage, disparam o sacrifício. Não há pacote de reanimação neste deck e não há criatura no pool que valha 4–5 manas de reanimação. Ainda **tiram carta do cemitério**. Dispensa por eixo, não por qualidade. |
| [**Remember the Fallen**](https://www.ligamagic.com.br/?view=cards/card&card=Remember+the+Fallen) | lista | F1 devolve criatura/artefato do cemitério à mão · F6 CMC 3 | Devolveria o Phlage **à mão**, de onde ele só é conjurável normalmente — e se sacrifica. Anti-sinergia direta com L1/L3. |
| [**Punishing Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Punishing+Fire) | caixa | F1 2 de dano; volta do cemitério **quando um oponente ganha vida** · F6 CMC 2 | O gatilho é **oponente** ganhar vida — o lifelink do Phlage é **seu**. Em pod de 4 dispara às vezes, mas cada retorno **tira uma carta do cemitério**: é anti-combustível. 1 ponto de sinergia, e negativo no outro eixo. |
| [**Skullcrack**](https://www.ligamagic.com.br/?view=cards/card&card=Skullcrack) | caixa | F1 3 num jogador; **ninguém ganha vida neste turno** · F6 CMC 2 | A cláusula é **simétrica**: desliga os 3 de vida do seu próprio Helix, exatamente no turno em que você mais quer estabilizar. Atrito F7 com o comandante. |
| [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) | caixa | F1 impulse 2 **ou** +2/+1 em dois alvos · F6 CMC 3 | Impulse **não abastece o cemitério** (exila). 1 ponto (saque). [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) faz mais por 1 mana a menos. |
| [**Oracle's Vault**](https://www.ligamagic.com.br/?view=cards/card&card=Oracle%27s+Vault) (R$ 0,25) | caixa | F1 impulse repetível; grátis com 3 contadores · F6 CMC 4 | Mesmo problema: exila em vez de moer. CMC 4 e 3 turnos para ligar num deck que precisa dos T2–T4 para remoção. |
| [**Magnifying Glass**](https://www.ligamagic.com.br/?view=cards/card&card=Magnifying+Glass), [**Seer's Lantern**](https://www.ligamagic.com.br/?view=cards/card&card=Seer%27s+Lantern) | caixa | F1 rocha + investigate/scry · F6 CMC 3 | Rochas de 3 manas com upside de 4–2 manas. [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) (R$ 0,17) faz rocha **e** combustível por 2. Dispensa por curva e por eixo. |
| [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais) | caixa | F1 sacrifica criatura → contador; saca por contador | Exige sacrificar criaturas; o deck tem 14–17 corpos e 6 deles **são** o wincon. Anti-sinergia com a §6.1. |
| [**Panic Spellbomb**](https://www.ligamagic.com.br/?view=cards/card&card=Panic+Spellbomb) | caixa | F1 criatura não bloqueia; compra ao ir ao cemitério | O deck não ataca com enxame — "não pode bloquear" não faz nada aqui. 1 ponto. |
| [**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon) | caixa | F1 `{T}`: contador; 3 contadores → **destrói qualquer permanente** · F6 CMC 4 | **Corte condicionado, não dispensa.** 2 pontos reais (remoção repetível irrestrita + artefato). O que o desqualifica é **velocidade**: 4 manas + 3 turnos = resposta no T7 para um problema do T4. **Devolvo à Fase 5** — se ela achar que o deck precisa de remoção *inevitável* e não só rápida, ele é o candidato sem custo. |
| [**Searing Barrage**](https://www.ligamagic.com.br/?view=cards/card&card=Searing+Barrage), [**Radiating Lightning**](https://www.ligamagic.com.br/?view=cards/card&card=Radiating+Lightning), [**Stonefury**](https://www.ligamagic.com.br/?view=cards/card&card=Stonefury), [**Seismic Wave**](https://www.ligamagic.com.br/?view=cards/card&card=Seismic+Wave), [**Lightning Volley**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Volley), [**Molten Blast**](https://www.ligamagic.com.br/?view=cards/card&card=Molten+Blast), [**Chandra's Outrage**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Outrage), [**Bombard**](https://www.ligamagic.com.br/?view=cards/card&card=Bombard), [**Pinecone Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Pinecone+Strike), [**Smaug's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Smaug%27s+Fury), [**Mercadia's Downfall**](https://www.ligamagic.com.br/?view=cards/card&card=Mercadia%27s+Downfall) | caixa | queima de Limited, CMC 3–5, 3–4 de dano | **Dispensa por taxa, não por "carta fraca".** Todas cabem no termo T3 e todas abastecem o cemitério — mas o deck só tem 21 slots de remoção, e [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) (4 dano por `{R}`), [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) e [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) dominam por curva. Ficam como **reserva de orçamento**: se a Fase 5 estourar o teto, elas substituem compras a custo zero. |
| [**Tori D'Avenant, Fury Rider**](https://www.ligamagic.com.br/?view=cards/card&card=Tori+D%27Avenant%2C+Fury+Rider) | lista | F1 pump a outros atacantes, trample a vermelhas, destapa brancas atacantes · F2 3/3 vigilância e trample · F3 Human Knight · F5 buffa atacantes · F6 CMC 4 · F7 disputa o T4 com Torbran/Solphim | **Fora do pool. 0 pontos de sinergia com o eixo.** Todas as cláusulas exigem **múltiplos atacantes**; este deck tem 14–17 corpos e ~6 deles têm **defender**. O gatilho não vê o tabuleiro que o deck vai ter. Não é corte (nunca esteve sob Phlage) — é dispensa de pool com ficha. |
| [**Otharri, Suns' Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Otharri%2C+Suns%27+Glory) | comandante v1 | F1 fichas de Rebel por contador de experiência; recursão exige **destapar um Rebel** · F2 3/3 voador lifelink haste · F6 CMC 5 | **Fora do pool.** A recursão dele é inoperante sem pacote de Rebel (já comprovado no teste de mesa da v1), e CMC 5 num deck cuja curva precisa caber remoção + loot no mesmo turno. Sobra um 3/3 voador com lifelink por 5 — 1 ponto. |
| 60 criaturas de combate da `lista.txt` ([**Honored Crop-Captain**](https://www.ligamagic.com.br/?view=cards/card&card=Honored+Crop-Captain), [**Fireborn Knight**](https://www.ligamagic.com.br/?view=cards/card&card=Fireborn+Knight), [**Adriana**](https://www.ligamagic.com.br/?view=cards/card&card=Adriana%2C+Captain+of+the+Guard), [**Kwende**](https://www.ligamagic.com.br/?view=cards/card&card=Kwende%2C+Pride+of+Femeref), [**Inspiring Veteran**](https://www.ligamagic.com.br/?view=cards/card&card=Inspiring+Veteran), [**Parhelion Patrol**](https://www.ligamagic.com.br/?view=cards/card&card=Parhelion+Patrol), [**Éomer**](https://www.ligamagic.com.br/?view=cards/card&card=%C3%89omer%2C+Marshal+of+Rohan) ×2, [**Syr Alin**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Alin%2C+the+Lion%27s+Claw), [**Relentless Rohirrim**](https://www.ligamagic.com.br/?view=cards/card&card=Relentless+Rohirrim), [**Knight of Sorrows**](https://www.ligamagic.com.br/?view=cards/card&card=Knight+of+Sorrows), [**Cavalry Drillmaster**](https://www.ligamagic.com.br/?view=cards/card&card=Cavalry+Drillmaster), [**Rohirrim Lancer**](https://www.ligamagic.com.br/?view=cards/card&card=Rohirrim+Lancer), [**Lake-town Lookout**](https://www.ligamagic.com.br/?view=cards/card&card=Lake-town+Lookout), [**Sheriff of Safe Passage**](https://www.ligamagic.com.br/?view=cards/card&card=Sheriff+of+Safe+Passage), [**Luxknight Breacher**](https://www.ligamagic.com.br/?view=cards/card&card=Luxknight+Breacher), [**Knight Luminary**](https://www.ligamagic.com.br/?view=cards/card&card=Knight+Luminary), [**Dawnstrike Vanguard**](https://www.ligamagic.com.br/?view=cards/card&card=Dawnstrike+Vanguard), [**Esgaroth Garrison**](https://www.ligamagic.com.br/?view=cards/card&card=Esgaroth+Garrison), [**Belladonna Took**](https://www.ligamagic.com.br/?view=cards/card&card=Belladonna+Took), [**Jor Kadeen**](https://www.ligamagic.com.br/?view=cards/card&card=Jor+Kadeen%2C+the+Prevailer), [**Paladin Danse**](https://www.ligamagic.com.br/?view=cards/card&card=Paladin+Danse%2C+Steel+Maverick), …) + os pumps ([**Basri's Solidarity**](https://www.ligamagic.com.br/?view=cards/card&card=Basri%27s+Solidarity), [**Inspiring Roar**](https://www.ligamagic.com.br/?view=cards/card&card=Inspiring+Roar), [**Pride of Conquerors**](https://www.ligamagic.com.br/?view=cards/card&card=Pride+of+Conquerors), [**Shoulder to Shoulder**](https://www.ligamagic.com.br/?view=cards/card&card=Shoulder+to+Shoulder), [**Zealous Display**](https://www.ligamagic.com.br/?view=cards/card&card=Zealous+Display), [**Crash Through**](https://www.ligamagic.com.br/?view=cards/card&card=Crash+Through), [**Warlord's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Warlord%27s+Fury), [**Bond of Discipline**](https://www.ligamagic.com.br/?view=cards/card&card=Bond+of+Discipline), [**Djeru's Renunciation**](https://www.ligamagic.com.br/?view=cards/card&card=Djeru%27s+Renunciation), [**Gideon's Triumph**](https://www.ligamagic.com.br/?view=cards/card&card=Gideon%27s+Triumph)) + anthems ([**Valor in Akros**](https://www.ligamagic.com.br/?view=cards/card&card=Valor+in+Akros), [**The Circle of Loyalty**](https://www.ligamagic.com.br/?view=cards/card&card=The+Circle+of+Loyalty), [**Vigilante Justice**](https://www.ligamagic.com.br/?view=cards/card&card=Vigilante+Justice), [**Sanctuary Lockdown**](https://www.ligamagic.com.br/?view=cards/card&card=Sanctuary+Lockdown)) + equipamentos ([**True-Faith Censer**](https://www.ligamagic.com.br/?view=cards/card&card=True-Faith+Censer), [**Squire's Lightblade**](https://www.ligamagic.com.br/?view=cards/card&card=Squire%27s+Lightblade), [**Ancestral Blade**](https://www.ligamagic.com.br/?view=cards/card&card=Ancestral+Blade), [**Trailblazer's Torch**](https://www.ligamagic.com.br/?view=cards/card&card=Trailblazer%27s+Torch)) + [**Basri Ket**](https://www.ligamagic.com.br/?view=cards/card&card=Basri+Ket) | lista | fichas completas em `rounds/v1-2026-09-18/02-theme.md` §3.1–§3.4, que **continuam válidas como levantamento** | **Dispensa em bloco, por eixo.** Todas foram desenhadas para **atacar em massa**: pump temporário condicionado a board, anthem estático que precisa de enxame, typal [**Knight**](https://www.ligamagic.com.br/?view=cards/card&card=Knight)/[**Human**](https://www.ligamagic.com.br/?view=cards/card&card=Human), equipamento que quer um atacante. Este deck tem 14–17 corpos, ~6 com defender, e **não vence por combate**. Nenhuma delas alcança 2 pontos de sinergia (regra 3). Exceções já resgatadas na §8.1: [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold), [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning), [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing), [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant), [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Expose to Daylight**](https://www.ligamagic.com.br/?view=cards/card&card=Expose+to+Daylight). |

**Consequência de custo, declarada:** a mudança de eixo **descarta ~60 das 69 não-básicas do
Tori físico**. O deck v2 é majoritariamente de compra. Isso não é surpresa nova — o briefing de
2026-09-20 já registrou que "mudar de eixo descarta quase todas as 69 não-básicas da Tori" e que
"nenhum custo de eixo alternativo está medido". **Ele continua não medido**, porque 37 das 48
candidatas principais estão **`a cotar`**. Sinal duro ao orquestrador: a cotação da LigaMagic
precisa ser feita **antes** de a Fase 6 fechar os cortes, não depois.

---

## 9. Pool temático — 42 candidatas

Regra 3: mínimo **2 pontos de sinergia**, explicitados. Preço: `bin/mtgdb prices`; sem cotação =
**`a cotar`** (regra 2). Origem: `caixa` / `lista` (custo zero) · `compra`.

### 9.1 Combustível — motores repetíveis (6)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) | 2 | Enchantment RW | (1) **moe 1/turno automático** = combustível sem gastar carta nem mana; (2) **1 de dano a cada oponente** por turno = alcance que fecha (§6.1); (3) se moer terreno, ganha vida — sinergiza com o plano de atrito | não | `a cotar` |
| [**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece) | 2 | Artifact | (1) **2 cartas/turno só tapando** — a maior taxa de combustível do pool; (2) segunda habilidade **embaralha o cemitério de volta** = resposta a [**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog)/[**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace), o ponto cego do escape | não | `a cotar` |
| [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) | 1 | Sorcery R (Retrace) | (1) **repetível para sempre** — cada uso despeja um terreno no cemitério; (2) ping repetível que converte terrenos excedentes em dano (com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), 3 por uso); (3) **nunca sai do cemitério** (retrace, não flashback) | não | `a cotar` |
| [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) | 2 | Art. Creature 0/1 | (1) moe 1/turno; (2) **é rocha de mana** — acelera o `{R}{R}{W}{W}` do escape; (3) **é corpo** e conta para o piso da §4 | não | **R$ 0,17** (12/08) |
| [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) | 2 | Art. Creature 2/2 | (1) **descarte no custo de conjura** = combustível imediato; (2) `{4}`: 1 de dano numa criatura, **repetível**; (3) **persist** — volta sozinha de wipe, e a segunda morte a manda de volta ao cemitério | **sim (caixa)** | R$ 0 |
| [**Tablet of Discovery**](https://www.ligamagic.com.br/?view=cards/card&card=Tablet+of+Discovery) | 3 | Artifact R | (1) ETB moe 1 **e deixa jogar**; (2) rocha que faz `{R}{R}` **para instantâneo/feitiço** — o deck é 35+ mágicas; (3) ajuda no `{R}{R}` do escape | não | `a cotar` |

### 9.2 Combustível + saque — auto-descarte (11)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) | 2 | Instant R | (1) 2 cartas ao cemitério; (2) saca 2 (dor 2); (3) **instantâneo** — abastece no turno do oponente, escape no seu | **sim (caixa)** | R$ 0 |
| [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) | 1 | Sorcery R | (1) **3 cartas ao cemitério por `{R}`** — melhor taxa do pool; (2) saca 2; (3) flashback dá um segundo uso (custo: ela se exila) | não | `a cotar` |
| [**Cathartic Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Reunion) | 2 | Sorcery R | (1) **3 ao cemitério**; (2) saca **3** — maior refil do pool | não | `a cotar` |
| [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) | 2 | Instant R | (1) **modal remoção/loot** — conta nos dois alvos numéricos (§3 e §5); (2) 3 de dano cobre a maioria das criaturas do formato; (3) instantâneo | não | `a cotar` |
| [**Tormenting Voice**](https://www.ligamagic.com.br/?view=cards/card&card=Tormenting+Voice) | 2 | Sorcery R | (1) 2 ao cemitério; (2) saca 2 | não | `a cotar` |
| [**Wild Guess**](https://www.ligamagic.com.br/?view=cards/card&card=Wild+Guess) | 2 | Sorcery R | (1) 2 ao cemitério; (2) saca 2; (3) `{R}{R}` puro ajuda a exercitar a base vermelha que o escape exige | não | `a cotar` |
| [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers) | 2 | Instant R | (1) 2 ao cemitério **ou** sacrifica artefato (converte rocha gasta em carta); (2) saca 2; (3) instantâneo | não | `a cotar` |
| [**Electric Revelation**](https://www.ligamagic.com.br/?view=cards/card&card=Electric+Revelation) | 3 | Instant R | (1) 2 ao cemitério; (2) saca 2; (3) **flashback** = a mesma carta duas vezes (F7: exila na segunda) | não | `a cotar` |
| [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) | 4 | Instant R | (1) 2 ao cemitério; (2) saca 2; (3) **2 Treasures** = rampa explosiva que paga o `{R}{R}{W}{W}` no mesmo turno | não | `a cotar` |
| [**Bitter Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Bitter+Reunion) | 2 | Enchantment R | (1) 2 ao cemitério + saca 2 no ETB; (2) **sacrifica-se para dar haste** — o Phlage escapado ataca no mesmo turno = Helix duplo | não | `a cotar` |
| [**Faithless Salvaging**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Salvaging) | 2 | Instant R | (1) loot; (2) [**Rebound**](https://www.ligamagic.com.br/?view=cards/card&card=Rebound) — dispara de novo no seu próximo upkeep, de graça | não | `a cotar` |

### 9.3 Remoção/queima que também é combustível (10 — a Fase 5 escolhe as 21 finais)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) | 1 | Sorcery R | (1) 4 de dano por `{R}` — mata quase tudo até o T5; (2) fonte **vermelha** (Torbran → 6); (3) 1 mana cabe junto com um loot no T3 | **sim (caixa)** | R$ 0 |
| [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) | 2 | Instant R | (1) 3 em **qualquer alvo** — flexível entre criatura e cara; (2) fonte vermelha; (3) instantâneo | **sim (caixa)** | R$ 0 |
| [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) | 1 | Instant R | (1) 2 + **exila** — desliga recursão alheia; (2) 1 mana | **sim (caixa)** | R$ 0 |
| [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) | 2 | Instant R | (1) 3 + **remove indestrutível** + exila; (2) resposta a comandante recorrente | **sim (caixa)** | R$ 0 |
| [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) | 2 | Artifact W | (1) exila criatura MV≤3 **permanentemente**; (2) trava as ameaças de T1–T3 que o deck lento sofre | **sim (caixa)** | R$ 0 |
| [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) | 2 | Sorcery W | (1) destrói criatura virada; (2) **spell mastery** — com 2+ mágicas no cemitério vira instantâneo, e este deck satisfaz a condição em quase todo turno | **sim (lista)** | R$ 0 |
| [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) | 3 | Sorcery W | (1) exila **qualquer não-terreno** — a remoção mais ampla sem custo; (2) o 3/2 que ela dá ao oponente é alvo do Helix | **sim (lista)** | R$ 0 |
| [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) | 4 | Enchantment W | (1) exila 1 não-terreno **por oponente** (3 remoções num card); (2) `recruit` = compra-e-descarta = **combustível** | **sim (lista)** | R$ 0 |
| [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) | 1 | Sorcery W (Spree) | (1) artefato **e** encantamento no mesmo card; (2) **resposta a ódio de cemitério** (§3.4) por 1 mana base | **sim (caixa)** | R$ 0 |
| [**Firebolt**](https://www.ligamagic.com.br/?view=cards/card&card=Firebolt) | 1 | Sorcery R | (1) 2 por `{R}`; (2) **flashback** = duas remoções num slot (F7: exila) | não | `a cotar` |

Complementos de compra que a Fase 5 deve cotar por serem o padrão do formato (cada um com 2+
pontos: remoção incondicional + fonte para o cemitério): [**Swords to Plowshares**](https://www.ligamagic.com.br/?view=cards/card&card=Swords+to+Plowshares) (**R$ 13,99**,
18/09) · [**Path to Exile**](https://www.ligamagic.com.br/?view=cards/card&card=Path+to+Exile) (`a cotar`) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) (**R$ 5,00**, 19/09) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp)
(**R$ 3,81**, 24/08) · [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) (`a cotar`) · [**Deflecting Palm**](https://www.ligamagic.com.br/?view=cards/card&card=Deflecting+Palm) (`a cotar`).
**[**Vandalblast**](https://www.ligamagic.com.br/?view=cards/card&card=Vandalblast) (R$ 39,99, 19/09) = 20% do teto — recomendo excluir por preço**, não por função.

### 9.4 Wipes (5 candidatos → 3–4 slots)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) | 5 | Sorcery W | (1) **wipe assimétrico** — você tem 14–17 corpos baratos, a mesa tem os caros; (2) **ganha vida por criatura** = estabiliza junto com o Helix; (3) enche seu cemitério com os seus mortos | **sim (caixa)** | R$ 0 |
| [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) | 9→~1 | Sorcery R | (1) custa `{R}` com 8 criaturas na mesa; (2) **fonte vermelha** — com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) mata +2 e com [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) triplica | não | `a cotar` |
| [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) | 5 | Sorcery R | (1) dano = nº de criaturas, escala com a mesa; (2) fonte vermelha (multiplicadores) | não | **R$ 0,90** (18/09) |
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | X | Sorcery R | (1) X dividido = wipe cirúrgico **ou** finisher na cara; (2) **flashback descartando X = X cartas ao cemitério**; (3) fonte vermelha | não | `a cotar` |
| [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) | 2 | Artifact | (1) wipe **escalável por CMC** — a curva do deck é 1–3, a dos oponentes não; (2) artefato barato que não é carta morta | **sim (caixa)** | R$ 0 |

Reserva sem custo no slot de terreno: **[**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone)** (caixa) — mesmo efeito do [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb)
sem gastar slot de mágica. Fase 6.

### 9.5 Alcance / pingadores — o fecho, dentro do orçamento de corpos (7)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) | 3 | Creature 2/2 | (1) **2 a cada oponente por mágica** — com 35+ mágicas é o maior relógio do pool; (2) corpo que bloqueia; (3) multiplica com Torbran/Solphim | não | `a cotar` |
| [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) | 2 | Creature 0/3 defender | (1) 1 a cada oponente por tap, **destapa a cada mágica** = N pings por turno; (2) **0/3 defender** — o bloqueador que o deck lento precisa no T2 | não | `a cotar` |
| [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) | 2 | Creature 0/4 defender | (1) 1 a cada oponente por mágica; (2) **0/4** — para praticamente toda agressão inicial do formato | não | `a cotar` |
| [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer) | 2 | Creature 2/1 | (1) 1 a cada oponente por **não-criatura** (pega também as rochas); (2) corpo que ataca se a mesa esvaziar | não | `a cotar` |
| [**Kessig Flamebreather**](https://www.ligamagic.com.br/?view=cards/card&card=Kessig+Flamebreather) | 2 | Creature 1/3 | (1) 1 por não-criatura; (2) 1/3 sobrevive a queima de 2 | não | `a cotar` |
| [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) | 2 | Creature 2/1 | (1) 1 por instantâneo/feitiço; (2) redundância de slot — o pacote precisa de 5–7 para P(≥1 até T6) ≥ 55% | não | `a cotar` |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | 5 | Leg. Creature 3/3 | (1) **exila o topo e deixa jogar** sempre que uma mágica sua fere um **jogador** — o pacote de pingadores dispara isso 3–5×/turno = vantagem de cartas explosiva; (2) `{T}`: 1 de dano, que é ela mesma disparando o próprio gatilho; (3) **já existe fisicamente** | **sim (lista)** | R$ 0 |

### 9.6 Multiplicadores de dano (4 candidatos → 3–4 slots)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | 4 | Leg. Creature 2/4 | (1) Helix do Phlage **3→5**; (2) **+200% em cada pingador** (1→3); (3) 2/4 bloqueia; (4) entra no mesmo turno do primeiro escape. **F7: só fonte vermelha** — não multiplica remoção branca | não | `a cotar` |
| [**Solphim, Mayhem Dominus**](https://www.ligamagic.com.br/?view=cards/card&card=Solphim%2C+Mayhem+Dominus) | 4 | Leg. Creature 5/4 | (1) **dobra dano não-combate** a oponentes — Helix 3→6, pingadores ×2; (2) corpo 5/4 real; (3) a habilidade de indestrutível custa **descartar 2 cartas** = combustível. **F7: não dobra combate** | não | `a cotar` |
| [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) | 6 | Enchantment R | (1) **triplica todo dano de fonte sua** — Helix 9, Phlage atacando 18; (2) com [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) vira finisher de um card. **F7: CMC 6** | não | `a cotar` |
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | 7 | Leg. Creature 5/5 | (1) **dobra todo dano a oponentes, inclusive combate** — Phlage = 6 + 12; (2) **reduz pela metade o dano recebido** = sobrevivência real até o T7; (3) 5/5 voadora com first strike cobre o céu, que é o buraco de RW. **F7: CMC 7, o mais caro do pool** | não | `a cotar` |

### 9.7 Proteção do plano e utilidade (5)

| Carta | CMC | Tipo | Sinergias (≥2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) | 2 | Instant RW | (1) **indestrutível em todos os seus permanentes** — faz o *seu* wipe ser assimétrico e salva os motores/pingadores; (2) 4 na cara = alcance; (3) double strike no Phlage = 12 de combate | não | `a cotar` |
| [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) | 2 | Enchantment W | (1) remoção pré-paga de artefato/encantamento; (2) **sai do campo para o cemitério** quando usada = combustível | **sim (lista)** | R$ 0 |
| [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) | 2 | Instant W | (1) resposta a [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace)/[**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog) em instantâneo; (2) resposta a equipamento/anthem alheio | **sim (lista)** | R$ 0 |
| [**Sunforger**](https://www.ligamagic.com.br/?view=cards/card&card=Sunforger) | 3 | Artifact — Equipment RW | (1) tutora instantâneo RW ≤4 a cada `{R}{W}` = transforma as 21 remoções em toolbox; (2) +4/+0 no Phlage 6/6 = 10/6. **F7: equip `{3}` é caro e o corpo do deck é magro** | não | `a cotar` |
| [**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) | 2→`{R}` | Sorcery R | (1) impulse 2 por `{R}` com spectacle — o Helix do Phlage **liga o spectacle sozinho**; (2) saque barato (dor 2). **F7: impulse exila, não abastece** | **não — está dentro do Krenko montado** (regra 7: compra-se a segunda cópia) | **R$ 0,74** (22/08) |

**Total do pool: 6 + 11 + 10 + 5 + 7 + 4 + 5 = 48 candidatas**, das quais **6 são reserva
explícita** ([**Faithless Salvaging**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Salvaging), [**Firebolt**](https://www.ligamagic.com.br/?view=cards/card&card=Firebolt), [**Expose to Daylight**](https://www.ligamagic.com.br/?view=cards/card&card=Expose+to+Daylight), [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone), [**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon),
[**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage)) → **42 no pool principal**, para ~35 slots temáticos. O excedente é a
reserva de cortes da Fase 6, como o processo pede.

---

## 10. Ficha de funções — as 12 cartas que carregam o eixo

| Carta | Categorias | Corpo tapável (F2) | Tipo alimenta (F3) | Recebe (F4) | Facilita (F5) | Entra no turno (F6) | Atritos (F7) |
|---|---|---|---|---|---|---|---|
| [**Phlage, Titan of Fire's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Phlage%2C+Titan+of+Fire%27s+Fury) | tema, remoção, wincon | 6/6, sem tap | Creature (convoke/anthem) | anthems, equipamentos, multiplicadores de dano | nada | **T3** (zona de comando) / T6 escape | disputa cemitério com flashback; sacrifica se não escapar |
| [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) | tema, wincon parcial | não | Enchantment | — | **combustível + alcance** | T2 | moe **você**; sem escolha de alvo |
| [**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece) | tema | não (só tap p/ moer) | Artifact (metalcraft/improvise) | — | combustível | T2 | a 2ª habilidade **exila-se**; é uso único |
| [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) | tema, remoção, wincon parcial | não | Sorcery | — | converte terreno morto em dano | T1, repetível a partir daí | consome terrenos da mão (compete com land drop) |
| [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) | tema, **ramp** | **0/1 — tapa para mana** | **Artifact Creature** (conta no piso de corpos **e** em metalcraft) | anthems | mana + combustível | T2 | 0/1 não bloqueia nada |
| [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) | tema, remoção, proteção | **2/2** | **Artifact Creature** | anthems | combustível no cast | T2 | `{4}` por 1 de dano é caro; custo adicional de descarte em mão vazia dói |
| [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) | **draw**, tema | não | Instant | — | combustível | T2, instantâneo | descarte forçado com mão boa |
| [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) | **remoção**, draw, tema | não | Instant | — | modal — cobre dois buracos | T2 | escolhe um dos dois; não faz os dois |
| [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) | **wincon**, tema | 2/2 | Creature | anthems, multiplicadores | alcance | T3 | 2/2 morre a qualquer queima |
| [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) | wincon, **proteção** | **0/4 defender** | Creature | anthems (inúteis — defender) | alcance + bloqueio | T2 | **não ataca**; não converte board em dano |
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | **wincon** | 2/4 | Creature | anthems | **+2 em cada fonte vermelha** | T4 | **só vermelho** — remoção branca fica de fora |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | **draw**, wincon parcial | 3/3, **tapa para 1 de dano** | Creature | anthems, multiplicadores | vantagem de cartas | T5 | CMC 5 na curva mais apertada do deck |

---

## 11. Sinais cruzados para as Fases 3–7

| Fase | Sinal |
|---|---|
| **3 · draw** | Meta do pipeline: 12–13 fontes. **8 a 10 delas devem ser loot/auto-descarte** (§3.3) — draw que **não** descarta é slot desperdiçado aqui. [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) (lista, custo zero) é a maior fonte não-loot do pool. Impulse ([**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage), [**Reckless Impulse**](https://www.ligamagic.com.br/?view=cards/card&card=Reckless+Impulse) R$ 9,78, [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) R$ 0,90) **exila e não abastece** — no máximo 2 slots. |
| **4 · ramp** | O gargalo não é o CMC 3 do Phlage — é o **`{R}{R}{W}{W}` do escape**, que exige **duas fontes vermelhas e duas brancas no mesmo turno**, e ainda sobrar mana para remoção. Preferir rochas de cor ([**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket), [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere), [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet), [**Fellwar Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Fellwar+Stone)) a rochas incolores. [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) e [**Tablet of Discovery**](https://www.ligamagic.com.br/?view=cards/card&card=Tablet+of+Discovery) são ramp **e** combustível — contem nas duas metas. [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) faz 2 Treasures no turno do escape. |
| **5 · interação** | Alvo revisado: **20–22 remoções + 3–4 wipes** (não os ~10 do pipeline). Restrições: preferir instantâneo/feitiço a permanente (vai ao cemitério); preferir `{R}` a `{W}` no empate (Torbran); ≥8 em CMC ≤2; **2 respostas a encantamento obrigatórias** (ódio de cemitério). Devolvo [**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon) (caixa) como **corte condicionado** — 2 pontos reais, reprovado só por velocidade. |
| **6 · manabase** | **36 terrenos não bastam** para uma base que precisa de RRWW no T5–T6 junto com mágicas de 1–2. Avaliar **terrenos com cycling** ([**Forgotten Cave**](https://www.ligamagic.com.br/?view=cards/card&card=Forgotten+Cave), [**Secluded Steppe**](https://www.ligamagic.com.br/?view=cards/card&card=Secluded+Steppe), [**Smoldering Crater**](https://www.ligamagic.com.br/?view=cards/card&card=Smoldering+Crater), [**Desert of the Fervent**](https://www.ligamagic.com.br/?view=cards/card&card=Desert+of+the+Fervent), [**Desert of the True**](https://www.ligamagic.com.br/?view=cards/card&card=Desert+of+the+True), [**Glittering Massif**](https://www.ligamagic.com.br/?view=cards/card&card=Glittering+Massif)): eles são **combustível de graça no slot de terreno** e resolvem inundação. [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) (caixa) e [**Sequestered Stash**](https://www.ligamagic.com.br/?view=cards/card&card=Sequestered+Stash) como utilidade. |
| **7 · wincon** | O eixo **não** fecha o jogo — §6 é o pacote. Requisito duro: **5–7 pingadores + 3–4 multiplicadores**. [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) é o nº 1 por curva e por multiplicar +200% sobre base 1. `Gisela` fica como **nº 2** (dobra combate + metade defensiva, mas CMC 7). Goldfishing precisa medir **turno do primeiro escape** (previsão: T6, P=52%) e **turno de morte do primeiro oponente**. |
| **Orquestrador** | (a) **37 das 48 candidatas estão `a cotar`** — a captura na LigaMagic precisa acontecer **antes** de a Fase 6 fechar cortes. (b) O eixo **descarta ~60 das 69 não-básicas do Tori**: o deck v2 é majoritariamente compra, e o custo total é **desconhecido**, não "baixo". (c) **Risco de percepção**: 27,9% de mãos sem criatura é o dobro dos 12,1% da v1 corrigida — o usuário reclamou exatamente disso, e isto precisa ser pergunta explícita no `report.md`. |
