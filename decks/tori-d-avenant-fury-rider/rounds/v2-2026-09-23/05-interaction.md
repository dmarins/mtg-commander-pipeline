# Interação — Phlage, Titan of Fire's Fury (v2 · eixo de controle de atrito)

> **Natureza desta fase.** Não existe `deck.md`. A cota já vem reservada pela Fase 2
> (§5: **20–22 remoções + 3–4 wipes**, contra os ~10 + 2–4 do pipeline). Eu **seleciono dentro
> da cota**, não brigo por slots — e o resultado desta fase é que ela **devolve 2,5 slots** ao
> deck, que estava estourado.
>
> **Regra 5 — conferido.** `decisions.md` registra apenas movimentos de zona de comando
> (Otharri entra/sai, Tori sai). **Nenhuma carta do 99 tem histórico de corte.** Nenhuma proposta
> abaixo é reposição de corte anterior. As três travas do escape (não reanimar, não blinkar,
> 6 cartas no cemitério para o 1º escape) foram respeitadas, não redescobertas.
>
> **Regra 7 — caixa primeiro.** `bin/mtgdb collection -list` (116 cartas) e as 69 não-básicas de
> `lista.txt` foram varridas **antes** de qualquer busca. **14 das 22 peças titulares custam
> R$ 0** de aquisição. Toda dispensa de carta da caixa está nomeada e justificada na §7.
>
> **Regra 6 — oracle.** Todo texto citado foi puxado com `bin/mtgdb oracle` nesta sessão.
>
> **Regra 2 — preço.** Nenhum número do Scryfall. `bin/mtgdb prices` consultado para as 38
> candidatas; o que não tem cotação sai como **`a cotar`**, nunca estimado.

**Interação selecionada: 20/20–22 · Wipes: 4/3–4 · Proteção dedicada: 2**

---

## 1. Resposta ao pedido da Fase 3 — os 2 slots. **Aceito, com uma correção de contagem.**

A Fase 3 (§7, §9) pede que a cota caia de **20–22 → 19–20**, argumentando que `Palace Jailer` e
`Court of Ire` **são** remoção e já devem ser contados dentro dela. Puxei o oracle das duas:

| Carta | Oracle (conferido) | É remoção? | Quanto vale como slot de remoção |
|---|---|---|---|
| **Court of Ire** `{3}{R}{R}` | *ETB: você vira monarca. No seu upkeep, causa **2 de dano em qualquer alvo; 7 se você for monarca**.* | **Sim, e sem ressalva.** 7 por upkeep mata quase qualquer criatura do formato, **todo turno**, e ainda é alcance na cara. Fonte **vermelha** → 9 com `Torbran`. | **1,0 slot cheio.** Remoção repetível vale mais que remoção única; o preço é o CMC 5 e o fato de o gatilho ser no *seu* upkeep (não responde a ameaça que entrou depois). |
| **Palace Jailer** `{2}{W}{W}` | *ETB: você vira monarca. ETB: exila criatura de oponente **até que um oponente vire monarca**.* | **Sim, mas condicional.** Num pod de 4 a coroa é roubada — e no instante em que é, a ameaça **volta**. | **0,5 slot.** É a única forma de remoção do deck que se **desfaz sozinha exatamente quando você está perdendo**. Como corpo 2/2 e motor de coroa ela se paga; como remoção, meia. |

**Contabilidade final: 1,5 slot de remoção efetiva, não 2.** Por isso não aceito a faixa
**19–20** — aceito **20 no ponto**, que é o piso que a própria Fase 2 declarou:

- **20 remoções no deck**, das quais **2 são cartas da Fase 3** (`Court of Ire`, `Palace Jailer`).
- **Eu seleciono 18.** Os 2 slots que a Fase 3 pediu saem daqui, integralmente.
- **Não desço de 20.** A Fase 2 mediu P(≥1 remoção na mão de 7) = 80,5% com 20 e 76,6% com 18, e
  P(≥2 até o T4) = 64,3% com 20. Em 19 a segunda métrica cai abaixo do piso que ela mesma chamou
  de alvo, e este é o deck em que "não tenho resposta neste turno" é a derrota, não um contratempo.
- **Compensação da meia-peça do `Palace Jailer`:** exigi que **4 das minhas 18** sejam remoção
  *irrestrita* (`destroy/exile target permanent`), que é justamente o que o Jailer não é —
  `Chaos Warp`, `Generous Gift`, `Reduce to Memory`, `Celebrate the Mountain-king`.

**Consequência de orçamento, e é boa notícia.** A Fase 2 reservou 21 remoções + 3,5 wipes =
**24,5 slots**. Eu uso **22** (18 + 4). **Devolvo 2,5 slots.**

| Conta de slots (37 terrenos → 62 não-terreno) | Antes desta fase | Depois |
|---|---|---|
| Soma das Fases 2+3+4 (§7 da Fase 3) | 58,5 | 56,0 |
| + os 4 slots exclusivos que a Fase 3 pediu | 62,5 ⚠ **estourado** | **60,0** |
| Folga | **−0,5** | **+2,0** |

**O deck deixa de estar estourado por conta desta fase.** Consequência prática: a Fase 3 **não
precisa** do seu segundo pedido (cortar 2 dos 10 loots, que levaria P(escape no T6) de 52,4% para
~48%). O combustível fica em **14 (config D)**. A decisão é da Fase 6, mas o custo já não é
necessário.

---

## 2. Candidatas — remoção / proteção / counters

**Counters: zero.** RW não tem contramágica. Isso é buraco de cor declarado, não omissão — ver §5.

### 2.1 As 18 titulares

| Carta | CMC | Tipo | Subcategoria | O que responde | Sinergias (≥2, regra 3) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| **Flame Slash** | 1 | Sorcery R | queima | criatura (4 dano) | (1) 4 de dano por `{R}` é a melhor taxa do pool e cabe **junto com um loot** no mesmo turno; (2) fonte **vermelha** → 6 com `Torbran` | **sim (caixa)** | R$ 2,48 (12/08) |
| **Magma Spray** | 1 | Instant R | queima + exílio | criatura pequena, **recursão alheia** | (1) exila — desliga Phoenix/persist/reanimação do oponente, que é o que mais castiga um deck de atrito; (2) 1 mana, fonte vermelha | **sim (caixa)** | `a cotar` |
| **Unholy Heat** | 1 | Instant R | queima escalável | criatura **ou planeswalker** | (1) **delirium: 6 de dano** — o único cartão do pool cuja qualidade *cresce* com o cemitério cheio, que é o recurso que o deck fabrica de propósito; (2) fonte vermelha → 8 com `Torbran`; (3) `{R}` instantâneo deixa mana aberto | não | **R$ 0,75** (12/08) |
| **Requisition Raid** | 1 +Spree | Sorcery W | artefato **e** encantamento | ódio de cemitério, rocha, anthem, equipamento | (1) **o único card do pool que mata artefato E encantamento no mesmo lançamento** (`{3}{W}` pelos dois modos) — o cenário de blowout é `Rest in Peace` **mais** `Grafdigger's Cage`; (2) feitiço = combustível + gatilho de pingador | **sim (caixa)** | `a cotar` |
| **Lightning Strike** | 2 | Instant R | queima flexível | criatura, planeswalker, **jogador** | (1) 3 em **qualquer alvo** — é remoção quando precisa e é relógio quando a mesa está limpa; (2) fonte vermelha → 5 com `Torbran` | **sim (caixa)** | `a cotar` |
| **Smite the Deathless** | 2 | Instant R | queima + exílio | criatura **indestrutível**, comandante recorrente | (1) **remove indestrutível e exila** — é a resposta a comandante que volta e a `Avacyn`/`Blightsteel`; (2) instantâneo vermelho | **sim (caixa)** | `a cotar` |
| **Mizzium Mortars** | 2 / 6 | Sorcery R (Overload) | queima + **wipe unilateral** | criatura · **enxame alheio** | (1) `{1}{R}`: 4 de dano numa criatura que **você não controla** — nunca erra de lado; (2) **Overload `{3}{R}{R}{R}`: 4 em cada criatura que você não controla** = wipe 100% assimétrico num slot de remoção; (3) fonte vermelha → 6/6 com `Torbran` | não | **R$ 1,10** (16/09) |
| **Abrade** | 2 | Instant R | modal criatura/artefato | criatura (3) **ou** artefato | (1) modal — cobre dois tipos de ameaça num slot, que é o câmbio certo num deck estourado de slots; (2) instantâneo vermelho | não | `a cotar` |
| **Lightning Helix** | 2 | Instant **RW** | queima + vida | criatura, planeswalker, jogador | (1) é **o gatilho do comandante em carta** — 3 de dano + 3 de vida; (2) os 3 de vida são o que faz o deck chegar ao T7 sob a coroa (a Fase 3 avisou que ser monarca convida ataque); (3) fonte vermelha → 5 com `Torbran` | não | `a cotar` |
| **Glass Casket** | 2 | Artifact W | exílio permanente | criatura MV ≤3 | (1) exila **permanentemente** o que 3 de dano não resolve (indestrutível, regeneração, ETB abusivo); (2) trava as ameaças de T1–T3, que é a janela em que este deck é mais lento. **F7: fica em campo — não abastece** | **sim (caixa)** | **R$ 0,09** (12/08) |
| **Swift Reckoning** | 2 | Sorcery W (flash) | destruição incondicional | criatura **virada** | (1) **spell mastery**: com 2+ mágicas no cemitério vira instantâneo — o deck satisfaz a condição em quase todo turno a partir do T3, então na prática é remoção instantânea de 2 manas; (2) destrói **incondicionalmente**, sem régua de resistência — mata o 8/8 que a queima não pega | **sim (lista)** | `a cotar` |
| **Disenchant** | 2 | Instant W | artefato/encantamento | `Rest in Peace`, equipamento, anthem | (1) resposta **instantânea** ao ódio de cemitério — pode ser guardada e usada no turno do oponente; (2) 2 manas deixa espaço para um loot no mesmo turno | **sim (lista)** | `a cotar` |
| **Seal of Cleansing** | 2 | Enchantment W | artefato/encantamento pré-pago | idem | (1) **pré-pago**: você resolve o custo num turno morto e detona de graça no turno em que precisa do mana para o escape `{R}{R}{W}{W}`; (2) **sai do campo para o cemitério quando usada = combustível que não custou uma carta de mão** | **sim (lista)** | `a cotar` |
| **Chaos Warp** | 3 | Instant R | **irrestrita** | **qualquer permanente** | (1) a única resposta RW a permanente com hexproof/indestrutível/proteção fora de um wipe; (2) instantânea — responde à peça de combo no turno do oponente; (3) fonte vermelha (sem dano, mas dentro do plano) | não | **R$ 3,81** (24/08) |
| **Generous Gift** | 3 | Instant W | **irrestrita** | **qualquer permanente** | (1) segunda cópia funcional do `Chaos Warp`, e sem o risco de revelar um permanente melhor; (2) o 3/3 Elefante que ela concede é alvo legítimo do Helix do Phlage e morre a qualquer wipe do deck | não | **R$ 5,00** (19/09) |
| **Reduce to Memory** | 3 | Sorcery W | **irrestrita, exílio** | **qualquer não-terreno** | (1) **exila** — a resposta mais limpa do pool sem custo; (2) o 3/2 concedido é alvo do Helix e some no `Fumigate`; (3) feitiço = combustível | **sim (lista)** | `a cotar` |
| **Wear // Tear** | 3 (fuse) | Instant // Instant RW | artefato **e** encantamento | ódio de cemitério, rocha, anthem | (1) **fuse**: `{1}{R}{W}` mata artefato **e** encantamento no mesmo lançamento, **em instantâneo** — é a `Requisition Raid` com o timing que a `Requisition Raid` não tem; (2) as duas metades são conjuráveis isoladas (`{W}` só Tear) — nunca é carta morta | não | `a cotar` |
| **Celebrate the Mountain-king** | 4 | Enchantment W | **irrestrita ×3** | 1 não-terreno **por oponente** | (1) **três remoções irrestritas num card** num pod de 4 — a maior densidade de resposta do pool inteiro; (2) `recruit` = compra-e-descarta = **combustível** + um corpo 1/1 que conta no piso da §4 da Fase 2 | **sim (lista)** | `a cotar` |

**+ 2 da Fase 3, contadas aqui (§1):** `Court of Ire` (remoção repetível de 7/turno) e
`Palace Jailer` (exílio condicional). **Total no deck: 20.**

### 2.2 Proteção — o dimensionamento (dor 4 do intake)

**Pergunta:** a recursão do comandante basta, ou o deck precisa de proteção dedicada?
**Resposta: 2 slots, e apenas 1 deles é compra.** A conta:

| O que morre num wipe | Volta sozinho? | Custo de reconstruir |
|---|---|---|
| **Phlage** | **Sim** — escape `{R}{R}{W}{W}`, custo fixo, sem imposto de comandante | 4 manas + 5 cartas do cemitério. E o wipe **encheu** o cemitério. |
| 4 motores de combustível (`Lorehold Excavation`, `Perpetual Timepiece`, `Millikin`, `Tablet of Discovery`) | não | 2–3 manas cada, mas **o motor é o que paga o 2º e o 3º escape** |
| 5–7 pingadores (o wincon da §6 da Fase 2) | não | 2–3 manas cada — **é o plano de vitória inteiro** |
| 3–4 multiplicadores (`Torbran`, `Solphim`) | não | 4 manas |

A leitura honesta: **a recursão do comandante resolve a peça que menos precisava de resolução.**
O que o deck de fato perde num wipe alheio é o pacote de alcance e os motores — e nenhum deles
volta. Mas três fatos comprimem o tamanho da resposta:

1. **Você é quem dá o wipe.** Com 4 wipes na lista, o deck escolhe o turno da limpeza.
2. **O que morre é barato.** O pacote inteiro está em CMC 2–4; reconstruir é 2–3 turnos, não a
   partida. A v1 perdia o *plano*; a v2 perde *tempo*.
3. **A ameaça real não é wipe, é ódio de cemitério.** `Rest in Peace` desliga o escape, o
   `Palantír`, o `Sin Prodder`, o `Case of the Crimson Pulse` e o delirium da `Unholy Heat` —
   quatro fases de uma vez. **O orçamento de "proteção" deste deck está gasto em remoção de
   encantamento** (9 cartas, §4), não em hexproof.

| Carta | CMC | Tipo | Subcategoria | O que responde | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| **Boros Charm** | 2 | Instant RW | **proteção em massa** | wipe alheio (destroy), remoção pontual, `Cyclonic Rift` não | (1) *"**Permanents** you control gain indestructible"* — **todos os permanentes, não só criaturas**: salva motores, pingadores e multiplicadores de uma vez, por 2 manas em instantâneo; (2) é o que faz o **seu** `Fumigate`/`Chain Reaction` ser unilateral de verdade; (3) modo 2 = 4 na cara (alcance) e modo 3 = double strike no Phlage 6/6 = 12 de combate. **Três modos, três funções.** | não | `a cotar` |
| **Gods Willing** | 1 | Instant W | proteção pontual | remoção mirada no `Torbran`/pingador-chave | (1) `{W}` para salvar o multiplicador, que é a peça cujo corte custa o fecho do jogo; (2) **scry 1** — nunca é carta morta num deck que precisa achar o 4º terreno; (3) custo **R$ 0**, já está fisicamente na `lista.txt` | **sim (lista)** | `a cotar` |

**Por que não um 3º e 4º slot de proteção:** cada slot adicional compra seguro para permanentes
que custam 2 manas, num deck já estourado de slots, e *não* cobre o único desastre irrecuperável
(ódio de cemitério), que se responde com remoção. `Feat of Resistance`, `Adamant Will` e
`Duty Beyond Death` (todas na `lista.txt`, R$ 0) ficam de fora por isso — ficha na §7.

---

## 3. Candidatas — board wipes

**4 wipes, dos quais 2 são assimétricos por construção e 1 é seletivo.** A tabela mede o que
cada um leva do **meu** board típico de T6–T8 (Phlage 6/6 escapado + 2–3 pingadores de poder 0–2,
resistência 3–4 + `Torbran` 2/4 + `Millikin` 0/1):

| Carta | CMC | Simétrico? | Quanto leva do **meu** board | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|
| **Slaughter the Strong** | 3 | **NÃO — unilateral por construção do deck** | *Cada jogador escolhe criaturas com **poder total ≤ 4** e sacrifica o resto.* Meu pacote é `Electrostatic Field` 0/4, `Thermo-Alchemist` 0/3, `Millikin` 0/1, `Guttersnipe` 2/2, `Torbran` 2/4 — **poder total 0+0+0+2 = 2**: eu fico com **quatro deles**. Perco só o Phlage 6/6, **que quer estar no cemitério**. | (1) **é sacrifício, não destruição** — ignora indestrutível, hexproof e proteção; (2) a assimetria é *o piso de corpos da §4 da Fase 2*: o deck foi construído com criaturas de poder 0–2, e esta carta cobra exatamente por poder; (3) CMC 3 = cabe no mesmo turno em que sobra mana | não | `a cotar` |
| **Ratchet Bomb** | 2 | **NÃO — seletivo** | Com **0 contadores**: mata **todas as fichas da mesa** e nada meu (nenhuma das minhas peças é MV 0). Com 2: leva meus 2-drops junto. Você escolhe. | (1) **a resposta a enxame alheio sem perder uma carta** — a lacuna que a v1 fechou declarando que não tinha; (2) o deck vive em CMC 1–3 e os oponentes não: você aponta para a curva deles; (3) artefato barato que entra no T2 e espera | **sim (caixa)** | **R$ 1,79** (12/08) |
| **Fumigate** | 5 | Sim no texto, **não no custo** | Leva tudo, inclusive Phlage. Perco 3–4 permanentes de 2–3 manas; a mesa perde o investimento inteiro dela. | (1) **ganha 1 de vida por criatura destruída** — num wipe de 8 criaturas são 8 de vida, e estabilizar é metade do plano de atrito; (2) os meus mortos **vão para o meu cemitério** = 3–4 cartas de combustível de escape entregues no mesmo turno; (3) `Boros Charm` o transforma em unilateral | **sim (caixa)** | **R$ 1,37** (12/08) |
| **Chain Reaction** | 4 | Sim, mas **escalável** | X = nº de criaturas na mesa. Com 4 criaturas leva meus x/3 e x/4; com 8 leva tudo. Contra mesa pequena é barato e não dispara. | (1) **fonte vermelha** — com `Torbran` mata +2 e com `Fiery Emancipation` triplica: o único wipe do pool que os multiplicadores do deck melhoram; (2) escala com a ameaça: quanto maior o enxame alheio, mais ele mata; (3) feitiço = combustível | não | **R$ 0,90** (18/09) |

**+ 1 efeito de massa que não gasta slot de wipe:** **`Mizzium Mortars`** com Overload
(`{3}{R}{R}{R}`) = 4 de dano em **cada criatura que você não controla**. Assimetria total, e ele
já está contado como remoção na §2.1. Na prática o deck tem **5 respostas a enxame**.

**Reservas de wipe (Fase 6 decide):**

| Carta | CMC | Por que é reserva | Na coleção? | Preço |
|---|---|---|---|---|
| **Blast Zone** | — (Land) | Mesmo efeito do `Ratchet Bomb` **sem gastar slot de mágica** — entra na conta da Fase 6, não na minha | **sim (caixa)** | **R$ 0,89** (12/08) |
| **Blasphemous Act** | 9 → `{R}` | 13 de dano a cada criatura, custando `{R}` numa mesa cheia; fonte vermelha. **Upgrade direto sobre `Chain Reaction`** se a cotação fechar abaixo de ~R$ 3 | não | `a cotar` |
| **Conflagrate** | X | Wipe cirúrgico **e** finisher **e** combustível (flashback descarta X). Já está no pacote da Fase 2 — não duplico o slot | não | `a cotar` |
| **Winds of Abandon** | 2 / 6 | `{1}{W}` exila 1 criatura; Overload `{4}{W}{W}` **exila todas que você não controla**. É o melhor wipe unilateral que RW tem. **Fora por preço: R$ 19,41 = 9,7% do teto** num deck que a Fase 2 já declarou "majoritariamente compra". Upgrade sobre `Slaughter the Strong` **se o total fechar com folga** | não | **R$ 19,41** (19/09) |

---

## 4. Cobertura de ameaças

| Tipo de ameaça | Quantas respostas | Quais |
|---|---|---|
| **Criaturas** | **17** (+ o comandante, todo turno) | 12 das 18 titulares + os 4 wipes + `Court of Ire`. O Helix do Phlage são 3 de dano garantidos a cada entrada **e** a cada ataque. |
| **Artefatos** | **10** | `Requisition Raid` · `Abrade` · `Disenchant` · `Seal of Cleansing` · `Wear // Tear` · `Reduce to Memory` · `Generous Gift` · `Chaos Warp` · `Celebrate the Mountain-king` · `Ratchet Bomb` |
| **Encantamentos** | **9** | `Requisition Raid` · `Disenchant` · `Seal of Cleansing` · `Wear // Tear` · `Reduce to Memory` · `Generous Gift` · `Chaos Warp` · `Celebrate the Mountain-king` · `Ratchet Bomb` |
| **Planeswalkers** | **10** | `Unholy Heat` · `Lightning Strike` · `Lightning Helix` · `Generous Gift` · `Chaos Warp` · `Reduce to Memory` · `Celebrate` · `Ratchet Bomb` · `Court of Ire` (7/turno) · Helix do Phlage |
| **Ameaça em massa (enxame alheio)** | **5 titulares + 3 reservas** | `Mizzium Mortars` (Overload, unilateral) · `Ratchet Bomb` (seletivo) · `Slaughter the Strong` (unilateral) · `Chain Reaction` · `Fumigate` — reservas: `Blast Zone`, `Blasphemous Act`, `Conflagrate`. **A lacuna declarada da v1 está fechada, e nominalmente.** |
| **Flexível — "destroy/exile target permanent"** | **4** | `Chaos Warp` · `Generous Gift` · `Reduce to Memory` · `Celebrate the Mountain-king` (×3 alvos) |
| **Indestrutível / hexproof / proteção** | **6** | `Slaughter the Strong` (sacrifício, sem alvo) · `Ratchet Bomb` / `Blast Zone` (sem alvo) · `Fumigate` (sem alvo) · `Smite the Deathless` (remove indestrutível) · `Chaos Warp` (embaralha) |
| **Ódio de cemitério** (`Rest in Peace`, `Bojuka Bog`, `Soulless Jailer`, `Grafdigger's Cage`) | **9** | as 9 respostas a encantamento acima, das quais **4 em instantâneo** (`Disenchant`, `Wear // Tear`, `Generous Gift`, `Chaos Warp`). A Fase 2 exigiu 2; a Fase 3 elevou a obrigatórias porque o pacote de saque dela também morre. **Entreguei 9, e é de propósito** — ver §2.2. |
| **Combos / mágicas na pilha** | **0 — DESCOBERTO** | RW **não tem contramágica**. Ver §5.1. |

### 4.1 Mana aberto — distribuição instantâneo/feitiço

De **22 peças** (18 remoções + 4 wipes):

| Velocidade | Nº | % | Quais |
|---|---|---|---|
| **Responde no turno do oponente** | **13** | **59%** | 10 instantâneos verdadeiros (`Magma Spray`, `Unholy Heat`, `Lightning Strike`, `Smite the Deathless`, `Abrade`, `Lightning Helix`, `Disenchant`, `Chaos Warp`, `Generous Gift`, `Wear // Tear`) + `Swift Reckoning` (spell mastery ligado a partir do T3, na prática sempre) + `Seal of Cleansing` (sacrifício em instantâneo) + `Ratchet Bomb` (sacrifício em instantâneo) |
| Só no seu turno | 9 | 41% | `Flame Slash`, `Requisition Raid`, `Mizzium Mortars`, `Glass Casket`, `Reduce to Memory`, `Celebrate`, `Fumigate`, `Slaughter the Strong`, `Chain Reaction` |

**Leitura honesta:** 59% é aceitável, não excelente — um controle puro miraria 65–70%. A defesa é
que **os feitiços são os mais baratos da lista** (`Flame Slash` `{R}`, `Mizzium Mortars` `{1}{R}`,
`Requisition Raid` `{W}`), então a perda de tempo por feitiço é de 1–2 manas, não de um turno.
Empurrar acima de 59% exigiria trocar `Flame Slash` (4 dano por `{R}`, R$ 0) por instantâneos de
taxa pior e de compra — câmbio ruim num deck estourado de slots e de compra.

### 4.2 Curva

**CMC ≤ 2: 13 das 18** (`Flame Slash` 1, `Magma Spray` 1, `Unholy Heat` 1, `Requisition Raid` 1,
`Lightning Strike` 2, `Smite the Deathless` 2, `Mizzium Mortars` 2, `Abrade` 2, `Lightning Helix` 2,
`Glass Casket` 2, `Swift Reckoning` 2, `Disenchant` 2, `Seal of Cleansing` 2)
**CMC 3: 4** (`Chaos Warp`, `Generous Gift`, `Reduce to Memory`, `Wear // Tear`)
**CMC 4: 1** (`Celebrate the Mountain-king`)
**Wipes: 2, 3, 4, 5** (`Ratchet Bomb`, `Slaughter the Strong`, `Chain Reaction`, `Fumigate`)

A Fase 2 pediu **≥8 em CMC ≤ 2**. Entreguei **13**. Isso é o que permite conjurar remoção e loot
no mesmo turno até o T4, e remoção + segurar o `{R}{R}{W}{W}` do escape a partir do T5.

### 4.3 Remoção que também é combustível (item 2 do briefing desta fase)

Toda mágica de remoção vai ao cemitério **uma vez** e vira 1 de 5 cartas do escape. As que fazem
**mais** que isso, declaradas:

| Carta | Dupla função de cemitério |
|---|---|
| **Unholy Heat** | **Relação inversa** — quanto mais cheio o cemitério, melhor a carta. Delirium (4 tipos entre Instant/Sorcery/Artifact/Creature/Enchantment/Land) fica ligado a partir do T5 neste deck: os loots despejam terreno, o `Flame Jab` despeja terreno por retrace, `Seal of Cleansing` despeja encantamento, `Millikin`/`Perpetual Timepiece` despejam artefato. **2 de dano vira 6.** |
| **Swift Reckoning** | Idem — spell mastery (2+ instantâneos/feitiços no cemitério) transforma feitiço em instantâneo. É a única remoção do pool que **melhora de velocidade** pelo cemitério. |
| **Seal of Cleansing** | Sai do **campo** para o cemitério ao ser usada: é combustível que **não custou uma carta da mão**. |
| **Celebrate the Mountain-king** | `recruit` = compra 1, descarta 1 → **+1 carta no cemitério além dela mesma**, e um corpo 1/1. |
| **Conflagrate** (reserva) | Flashback custa *descartar X cartas* — a única remoção do pool que é um **motor** de combustível, não uma carga. |
| **Flame Jab** (pacote da Fase 2) | Retrace: **nunca sai do cemitério** e cada uso descarta um terreno. É remoção repetível contada pela Fase 2 — declaro a sobreposição para ninguém contar duas vezes. |

**Anti-combustível declarado (F7):** `Glass Casket` e `Ratchet Bomb` são permanentes — ficam em
campo e não abastecem nada até morrerem. São os **únicos 2** dos 22, e ambos entram por função
que nenhuma mágica cobre (exílio permanente de MV≤3 e wipe seletivo por MV). Custo aceito e
nomeado.

---

## 5. Buracos declarados

### 5.1 Contramágica: **zero, e sem mitigação plena**

RW não tem counterspells jogáveis. O deck **não responde a** `Thassa's Oracle`, `Ad Nauseam`,
`Approach of the Second Sun`, nem a um combo que ganhe na pilha. As mitigações reais, nomeadas:

- **Peça de combo que é permanente** (a maioria em pod casual/mid): `Chaos Warp` e
  `Generous Gift` respondem **em instantâneo**, e `Ratchet Bomb`/`Blast Zone` respondem por MV
  sem alvo. Isso cobre `Isochron Scepter`, `Dramatic Reversal`, `Kiki-Jiki`, `Food Chain`.
- **Combo de dano letal:** **`Deflecting Palm`** `{R}{W}` previne o dano e o devolve ao
  controlador da fonte. **Deixo como reserva nomeada**, não titular — é situacional demais para
  um dos 20 slots, mas é a única carta de RW que "contra" um `Comet Storm`.
- **O que fica sem resposta:** vitória alternativa e combo que resolve na pilha. **Declarado.**
  Este é o imposto da identidade de cor escolhida na Fase 1, não uma falha de seleção.

### 5.2 Céu

O deck bloqueia com `Electrostatic Field` 0/4, `Thermo-Alchemist` 0/3 e Phlage 6/6 — **nenhum com
alcance ou voo**. Contra um deck de voadores, a resposta é remoção pontual e wipe, não bloqueio.
`Gisela, Blade of Goldnight` (candidata nº 2 da Fase 7) é 5/5 voadora com first strike e cobriria
isso — registro a sobreposição para a Fase 7, sem pedir o slot.

### 5.3 Risco de percepção — carregado da Fase 2

A Fase 2 declarou que **27,9% das mãos abrem sem criatura**, contra 12,1% da v1 corrigida, e que o
usuário reprovou a v1 exatamente por *"muita mágica instantânea e encantamento na mão e poucas
criaturas"*. **Esta fase é a que mais agrava esse número**: 18 das minhas 22 peças são mágicas ou
encantamentos. Não tenho como reduzi-lo sem abandonar o eixo. **Sinalizo de novo ao orquestrador**
— é a pergunta que precisa ir ao usuário no `report.md`, não uma nota de rodapé.

---

## 6. Pendências devolvidas pela Fase 2 — decididas

### 6.1 `Lux Cannon` (caixa, R$ 1,25) — **NÃO ENTRA**

Ficha F1–F7 completa (regra 4), oracle puxado nesta sessão:

> `{T}`: Put a charge counter on this artifact.
> `{T}`, Remove three charge counters from this artifact: Destroy target permanent.

| Eixo | Conteúdo |
|---|---|
| **F1 texto** | Duas ativadas. A segunda é **remoção irrestrita e repetível** — destrói *qualquer* permanente, incluindo terreno, e ignora hexproof? **Não**: ela mira. Ignora indestrutível? Não, destrói. É irrestrita quanto ao **tipo**, não quanto à proteção. |
| **F2 corpo** | Não é criatura, não vira criatura. Tapa **só para si mesmo** — o deck não tem crew, station, convoke nem improvise, então o `{T}` não é recurso compartilhado. Não bloqueia. Não se sacrifica por valor. |
| **F3 tipo como recurso** | Artefato. **O pool da v2 não conta artefatos para nada** — o metalcraft do `Jor Kadeen` morreu com o eixo go-wide, e nenhuma carta das 48 candidatas da Fase 2 lê "number of artifacts". Valor: **zero**. |
| **F4 receptor** | Recebe charge counters **de si mesma**. Não há proliferate nem adicionador de contadores no pool. Não recebe anthem (não é criatura). Valor: zero. |
| **F5 facilitador** | Não dá nada a nenhuma outra carta. |
| **F6 curva** | CMC 4. Conjurada no T4, tapa no T4/T5/T6 → **primeira destruição no T7**. E o T4 é o turno do primeiro escape (`{R}{R}{W}{W}`) e do `Torbran` — ela **disputa o turno mais congestionado do deck**. |
| **F7 atrito** | (a) Nunca vai ao cemitério sozinha — **num deck cujo recurso escasso é "cartas no cemitério", um artefato que fica em campo é anti-combustível**; (b) disputa o T4 com escape e multiplicador; (c) como é permanente conhecido e lento, é o alvo óbvio da remoção da mesa nos 3 turnos em que não faz nada. |

**Protocolo de corte — funções e quem as cobre:**

| Função | Quem cobre depois da dispensa |
|---|---|
| remoção **irrestrita** (qualquer tipo de permanente) | `Chaos Warp`, `Generous Gift`, `Reduce to Memory`, `Celebrate the Mountain-king` — **4 peças, todas em 3–4 manas, duas delas instantâneas**, contra 4 manas + 3 turnos |
| remoção **repetível** | `Court of Ire` (7 de dano por upkeep, funciona **no turno seguinte** ao ETB) e `Flame Jab` (retrace, infinito) — a função é coberta com folga e com velocidade |
| artefato para contagens | **nenhuma contagem existe** — função inexistente, nada fica descoberto |

**Nenhuma função fica descoberta.** A Fase 2 reprovou por velocidade; a ficha completa mostra que
velocidade é o **melhor** dos seus eixos: os outros seis são zero ou negativos. Não é "carta
fraca" — é carta cuja única função já tem 4 substitutos mais rápidos e mais baratos em mana, num
deck onde o slot é mais escasso que o dinheiro. **R$ 0 abrem mão; aceito.**

### 6.2 `Requisition Raid` (caixa, R$ 0) — **ENTRA** (titular)

Aberta desde 2026-09-20 e adiada até o eixo fechar. O eixo fechou. Oracle:

> Spree — {W} base · +{1} Destroy target artifact · +{1} Destroy target enchantment ·
> +{1} Put a +1/+1 counter on each creature target player controls.

| Eixo | Conteúdo |
|---|---|
| **F1** | Modal aditivo: paga-se `{2}{W}` por um dos dois primeiros modos e `{3}{W}` **pelos dois**. |
| **F2** | Sem corpo. |
| **F3** | **Sorcery** — alimenta o spell mastery da `Swift Reckoning`, dispara `Guttersnipe`/`Electrostatic Field`/`Firebrand Archer`/`Erebor Flamesmith` e vai ao cemitério = combustível de escape. |
| **F4** | Não recebe nada. |
| **F5** | O 3º modo dá +1/+1 a todas as criaturas de um jogador — **marginal aqui** (4 criaturas pequenas, algumas com defender). Não é o motivo da entrada. |
| **F6** | T3 para um modo, T4 para os dois. |
| **F7** | **Feitiço** — não responde no turno do oponente, que é a preferência declarada do eixo. Disputa função com `Disenchant`, `Seal of Cleansing` e `Wear // Tear`. |

**Por que entra (regra 3 — 2 pontos):** (1) é a **única** carta do pool que mata artefato **e**
encantamento no mesmo lançamento, e o cenário que mata este deck não é uma peça de ódio de
cemitério, são **duas** (`Rest in Peace` + `Grafdigger's Cage`/`Soulless Jailer`); (2) é feitiço
que abastece o cemitério e dispara o pacote de pingadores, que é o fecho do jogo.

**A ressalva de 2026-09-20 caiu junto com o eixo.** Ela dizia *"contadores morrem com as
criaturas, então não ajuda na dor 4"* — o 3º modo não é por que ela entra, e a troca proposta lá
(`Basri's Solidarity` → `Requisition Raid`) não existe mais: `Basri's Solidarity` está fora do
pool desde a §8.2 da Fase 2. **A ressalva de velocidade continua válida e está coberta**:
`Disenchant` e `Wear // Tear` respondem em instantâneo quando o timing importa.

---

## 7. Varredura da caixa e da `lista.txt` — dispensas justificadas (regra 7)

`bin/mtgdb collection -list` (116 cartas) + 69 não-básicas de `lista.txt`. **14 das 22 titulares
saíram daí, a R$ 0 de aquisição.** O que ficou de fora, com ficha:

| Carta | Origem | Ficha resumida | Por que **não** cobre a função |
|---|---|---|---|
| `Twin Bolt` | caixa | F1 2 de dano dividido entre 1–2 alvos · F6 CMC 2 · F3 instantâneo = combustível | **Taxa, não qualidade.** 2 de dano por 2 manas não mata nada relevante em Commander a partir do T4. `Abrade` ocupa a mesma linha da curva com 3 de dano **ou** destruição de artefato. Fica como **reserva de orçamento**: se alguma compra estourar, ela volta a R$ 0. |
| `Fateful End` | caixa | F1 3 de dano em qualquer alvo + scry 1 · F6 CMC 3 | Mesma função do `Lightning Strike` **por 1 mana a mais**, e o scry 1 não paga a diferença num deck que precisa dos T2–T4 livres para remoção + loot. **Reserva de orçamento.** |
| `Seal of Fire` | caixa | F1 2 de dano pré-pagos; sai do campo para o cemitério = combustível · F6 CMC 1 | 2 pontos reais (pré-pago + combustível), mas **2 de dano** é a menor taxa do pool e o slot compete com `Magma Spray` (2 + **exílio**, instantâneo, mesmo custo). **Reserva.** |
| `Chandra's Pyrohelix` · `Searing Barrage` · `Radiating Lightning` · `Stonefury` · `Seismic Wave` · `Lightning Volley` · `Molten Blast` · `Chandra's Outrage` · `Bombard` · `Pinecone Strike` · `Smaug's Fury` · `Mercadia's Downfall` | caixa | queima de Limited, CMC 2–5, 2–4 de dano, todas abastecem o cemitério | **Dispensa por taxa (F6), com a ficha feita** — cada uma cabe no eixo e nenhuma é "carta fraca". Mas são **20 slots de remoção**, e `Flame Slash` (4 por `{R}`), `Unholy Heat` (6 com delirium por `{R}`) e `Mizzium Mortars` (4 por `{1}{R}`, que não erra de lado) dominam todas em dano-por-mana. A Fase 2 já as classificou como reserva de orçamento; confirmo. |
| `Punishing Fire` | caixa | F1 2 de dano; volta do cemitério quando **um oponente** ganha vida · F6 CMC 2 | Confirmo a dispensa da Fase 2 e acrescento o eixo que faltava: cada retorno **tira uma carta do cemitério**. Numa lista cujo recurso escasso é *cartas no cemitério*, recursão de mágica é **anti-combustível** — é a mesma razão pela qual a Fase 2 limitou flashback a 3 cartas. |
| `Skullcrack` | caixa | F1 3 num jogador; **ninguém ganha vida neste turno** · F6 CMC 2 | Confirmo: a cláusula é **simétrica** e desliga os 3 de vida do próprio Helix e o ganho do `Fumigate` — os dois estabilizadores do deck. Atrito F7 direto com o comandante. |
| `Lux Artillery` | caixa | F1 artefato CMC 4; converte contadores de carga em dano | Mesma ficha do `Lux Cannon` (§6.1): F2–F5 zerados, CMC 4, não vai ao cemitério, e o deck não tem nenhum outro gerador de contadores de carga para alimentá-la. Pior que o `Lux Cannon` porque nem remoção irrestrita entrega. |
| `Destructive Tampering` · `Smashing Success` · `Steel Wrecking Ball` · `Thaumaton Torpedo` | caixa | remoção de artefato, CMC 1–5 | Função **já coberta 10×** (§4), e por cartas que também pegam encantamento. `Steel Wrecking Ball` CMC 5 e `Smashing Success` CMC 4 são caras para função de resposta; `Thaumaton Torpedo` `{1}` só pega artefato. Nenhuma alcança 2 pontos de sinergia. |
| `Return to Nature` · `Smell Fear` · `Road // Ruin` · `Ancient Animus` · `Ambitious Assault` | caixa | — | **Fora da identidade de cor.** Todas verdes ou com metade verde. |
| `Negate` · `Stoic Rebuttal` · `Disruption Protocol` | caixa | contramágica | **Fora da identidade de cor** (azuis). São a razão pela qual a §5.1 existe: a caixa *tem* counters, e nenhum é legal aqui. |
| `Leonin Bola` | caixa | F1 equipamento `{1}`, `{T}` + unequip: vira criatura-alvo · F2 dá função de tap ao equipado | Pseudo-remoção (só vira, não mata) que **exige uma criatura desvirada** para funcionar. Este deck tem 14–17 corpos e vários com `defender`; e o corpo que importa (Phlage) quer estar atacando. 1 ponto. |
| `Joust` · `Moment of Glory` · `Phyrexian Revoker` | caixa | luta / +1/+1 / trava ativada | `Joust` exige criatura maior que a do oponente (o deck tem poder 0–2 fora do Phlage). `Moment of Glory` é pump, categoria morta neste eixo. `Phyrexian Revoker` é stax pontual — **regra 11**, e o deck não pediu. |
| `Feat of Resistance` · `Adamant Will` · `Duty Beyond Death` | lista | proteção pontual de 2 manas | Função coberta por `Boros Charm` (**todos** os permanentes, não só criaturas) e `Gods Willing` (1 mana, com scry). `Duty Beyond Death` exige **sacrificar uma criatura** como custo adicional — num deck de 14–17 corpos em que 6 **são** o wincon, é anti-sinergia direta com a §6.1 da Fase 2. Ver o dimensionamento de proteção na §2.2: 2 slots, não 5. |
| `Invoke the Divine` · `Expose to Daylight` | lista | destrói artefato/encantamento (+4 vida / +scry) | **Reservas legítimas, a R$ 0.** Ficam fora porque já selecionei 4 respostas a artefato/encantamento e as duas custam 3 manas contra os 2 do `Disenchant`. Se a Fase 6 quiser uma 5ª resposta sem gastar dinheiro, é daqui que ela sai — `Invoke the Divine` primeiro (os 4 de vida valem no plano de atrito). |
| `Wedgelight Rammer` · `Culling Dais` · `Panic Spellbomb` · `Liquimetal Coating` · `Titan Forge` · `Darksteel Reactor` | caixa | — | Nenhuma é interação. Fora do escopo desta fase. |
| **`Swords to Plowshares`** | compra | F1 exila criatura por `{W}` em instantâneo; controlador ganha vida igual ao poder | **Dispensa por preço, e a função está coberta.** R$ 13,99 = **7,0% do teto** por uma das 20 remoções, num deck em que 14 das 22 são R$ 0. O ganho de vida também trabalha contra um plano que fecha por dano. Exílio de criatura fica com `Magma Spray`, `Smite the Deathless`, `Glass Casket`, `Reduce to Memory` e `Celebrate`. **Recomendo cotar `Path to Exile`** (mesma função, cotação desconhecida) antes de desistir da linha. |
| **`Vandalblast`** | compra | wipe unilateral de artefatos | R$ 39,99 = **20% do teto**. Confirmo a exclusão da Fase 2 — por preço, não por função. `Ratchet Bomb` a 0 contadores e `Mizzium Mortars` cobrem parte disso. |

---

## 8. Reservas — a fila, na ordem

Se alguma cotação estourar ou a Fase 6 precisar de slot, esta é a ordem de substituição:

| Ordem | Entra | Sai | Por quê | Preço |
|---|---|---|---|---|
| 1 | `Twin Bolt` (caixa) | a compra mais cara entre `Abrade` / `Lightning Helix` | R$ 0 na mesma linha da curva; perde 1–2 de dano | R$ 0 |
| 2 | `Invoke the Divine` (lista) | `Wear // Tear` | R$ 0; perde a metade em instantâneo do artefato + 1 mana | R$ 0 |
| 3 | `Fateful End` / `Seal of Fire` (caixa) | qualquer compra `a cotar` que estoure | R$ 0, taxa pior | R$ 0 |
| 4 | `Blasphemous Act` | `Chain Reaction` | **upgrade**, se cotar abaixo de ~R$ 3 | `a cotar` |
| 5 | `Winds of Abandon` | `Slaughter the Strong` | **upgrade** para o melhor wipe unilateral de RW — só se o total fechar com folga | R$ 19,41 |
| 6 | `Fateful Absence` | `Lightning Helix` | destruição **incondicional** de criatura/PW por `{1}{W}`; troca 3 de vida por um Clue | `a cotar` |
| 7 | `Deflecting Palm` | — (entra só se a mesa tiver combo de dano) | única resposta RW a `Comet Storm`/`Fireball` letal | `a cotar` |
| 8 | `Reckless Rage` | — | 4 de dano por `{R}`; **exige criatura sua** (2 de dano nela), e 27,9% das mãos abrem sem corpo | `a cotar` |
| 9 | `Blast Zone` (caixa) | — (slot de **terreno**, Fase 6) | 5º efeito de massa sem gastar slot de mágica | R$ 0,89 |

---

## 9. Custo

**Origem: LigaMagic (menor), cotações de 12/08 a 19/09/2026.** Nenhum número do Scryfall.

**Aquisição (o que precisa ser comprado) — 8 cartas:**

| Carta | Preço |
|---|---|
| `Unholy Heat` | R$ 0,75 (12/08) |
| `Chain Reaction` | R$ 0,90 (18/09) |
| `Mizzium Mortars` | R$ 1,10 (16/09) |
| `Chaos Warp` | R$ 3,81 (24/08) |
| `Generous Gift` | R$ 5,00 (19/09) |
| `Abrade` | **a cotar** |
| `Lightning Helix` | **a cotar** |
| `Wear // Tear` | **a cotar** |
| `Slaughter the Strong` | **a cotar** |
| `Boros Charm` | **a cotar** |
| **Subtotal cotado** | **R$ 11,56** · **5 cartas a cotar** |

**Valor no deck das 14 que já existem fisicamente** (regra 2 mede o *valor do deck*, não o gasto):
`Flame Slash` R$ 2,48 · `Ratchet Bomb` R$ 1,79 · `Fumigate` R$ 1,37 · `Glass Casket` R$ 0,09 =
**R$ 5,73 cotados**; `Magma Spray`, `Lightning Strike`, `Smite the Deathless`, `Requisition Raid`,
`Swift Reckoning`, `Disenchant`, `Seal of Cleansing`, `Reduce to Memory`,
`Celebrate the Mountain-king`, `Gods Willing` — **10 a cotar** (todas commons/uncommons).

**Total desta fase: R$ 17,29 cotados + 15 cartas a cotar.** As 15 pendentes são majoritariamente
commons e uncommons; a exposição real está em `Boros Charm` e `Lightning Helix`, que são as duas
únicas com histórico de preço alto no formato. **Cotar as 15 antes de a Fase 6 fechar cortes.**

---

## 10. Sinais cruzados

| Fase | Sinal |
|---|---|
| **2 · tema** | Restrições cumpridas: **13** em CMC ≤2 (pedido: ≥8) · **9** respostas a encantamento (pedido: 2) · preferência por instantâneo/feitiço sobre permanente respeitada (só 2 dos 22 são permanentes) · preferência por `{R}` sobre `{W}` respeitada onde empatou. **`Lux Cannon` não entra** (§6.1, ficha completa); **`Requisition Raid` entra** (§6.2). Sobreposição a declarar para não haver dupla contagem: **`Flame Jab` e `Conflagrate` já estão no pacote de combustível da Fase 2** e são também remoção — eu **não** os contei nos meus 18. |
| **3 · draw** | **Pedido aceito, com correção**: `Court of Ire` vale 1,0 slot de remoção, `Palace Jailer` vale 0,5 (o exílio se desfaz quando a coroa é roubada — oracle conferido). Cota fixada em **20 no ponto**, não 19–20. Os 2 slots que você pediu estão pagos. **E você não precisa do segundo pedido**: como devolvo 2,5 slots, o deck fecha em 60,0/62 e os **10 loots ficam intactos** — a config D (14 fontes, P(escape T6) 52,4%) não precisa descer. Atendido também o pedido de guardar remoção instantânea para defender a coroa: **13 das 22 peças respondem no turno do oponente.** |
| **4 · ramp** | Curva de interação concentrada em 1–2 manas (13 de 18) **de propósito**, para caber junto com o `{R}{R}{W}{W}` do escape a partir do T5. O pico de mana da fase é `Fumigate` (`{3}{W}{W}`) — **dois brancos**, que reforça o gargalo que a Fase 2 apontou: o branco duplo, não o vermelho. `Slaughter the Strong` (`{1}{W}{W}`) idem. **Priorizem fixação de branco.** |
| **6 · manabase** | **Devolvo 2,5 slots** — a conta está na §1. `Blast Zone` (caixa, R$ 0,89) é o 5º efeito de massa **no slot de terreno**; recomendo forte. Os terrenos com cycling que vocês vão avaliar também são combustível para o **delirium da `Unholy Heat`** (o tipo Land no cemitério), o que sobe o valor deles acima do que a Fase 2 estimou. Se precisarem cortar da minha fase, a fila está na §8 e **nenhum corte meu desce de 20 remoções**. |
| **7 · wincon** | Sobreposições que já contam como alcance e não precisam de slot novo: `Lightning Strike`, `Lightning Helix`, `Unholy Heat`, `Fateful End` (reserva) e `Court of Ire` são **fonte vermelha em qualquer alvo** → contam no cálculo do `Torbran` e do `Solphim`. `Mizzium Mortars` com Overload + `Fiery Emancipation` é 12 em cada criatura alheia. **Buraco que vocês herdam: o céu** (§5.2) — nenhum bloqueador com voo ou alcance; `Gisela` cobriria isso além do papel de finisher. |
| **Orquestrador** | (a) **15 das 22 estão `a cotar`**; a exposição real é `Boros Charm` e `Lightning Helix`. Cotar **antes** dos cortes da Fase 6. (b) **Buraco declarado e sem conserto em RW: contramágica = 0** (§5.1) — combo que resolve na pilha ganha, e isso é imposto da cor escolhida na Fase 1, não erro de seleção. (c) **Repito o risco de percepção da Fase 2**: 18 das minhas 22 peças são mágicas e encantamentos, e o usuário reprovou a v1 dizendo *"muita mágica instantânea e encantamento na mão e poucas criaturas"*. Esta fase é a que mais agrava o número. **Tem que ser pergunta explícita no `report.md`.** (d) `Swords to Plowshares` (R$ 13,99) e `Winds of Abandon` (R$ 19,41) ficaram fora **por preço**; se o total fechar com folga, são os dois melhores upgrades disponíveis. |
