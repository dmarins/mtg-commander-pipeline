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

A Fase 3 (§7, §9) pede que a cota caia de **20–22 → 19–20**, argumentando que [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) e
[**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) **são** remoção e já devem ser contados dentro dela. Puxei o oracle das duas:

| Carta | Oracle (conferido) | É remoção? | Quanto vale como slot de remoção |
|---|---|---|---|
| [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) `{3}{R}{R}` | *ETB: você vira monarca. No seu upkeep, causa **2 de dano em qualquer alvo; 7 se você for monarca**.* | **Sim, e sem ressalva.** 7 por upkeep mata quase qualquer criatura do formato, **todo turno**, e ainda é alcance na cara. Fonte **vermelha** → 9 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell). | **1,0 slot cheio.** Remoção repetível vale mais que remoção única; o preço é o CMC 5 e o fato de o gatilho ser no *seu* upkeep (não responde a ameaça que entrou depois). |
| [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) `{2}{W}{W}` | *ETB: você vira monarca. ETB: exila criatura de oponente **até que um oponente vire monarca**.* | **Sim, mas condicional.** Num pod de 4 a coroa é roubada — e no instante em que é, a ameaça **volta**. | **0,5 slot.** É a única forma de remoção do deck que se **desfaz sozinha exatamente quando você está perdendo**. Como corpo 2/2 e motor de coroa ela se paga; como remoção, meia. |

**Contabilidade final: 1,5 slot de remoção efetiva, não 2.** Por isso não aceito a faixa
**19–20** — aceito **20 no ponto**, que é o piso que a própria Fase 2 declarou:

- **20 remoções no deck**, das quais **2 são cartas da Fase 3** ([**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire), [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer)).
- **Eu seleciono 18.** Os 2 slots que a Fase 3 pediu saem daqui, integralmente.
- **Não desço de 20.** A Fase 2 mediu P(≥1 remoção na mão de 7) = 80,5% com 20 e 76,6% com 18, e
  P(≥2 até o T4) = 64,3% com 20. Em 19 a segunda métrica cai abaixo do piso que ela mesma chamou
  de alvo, e este é o deck em que "não tenho resposta neste turno" é a derrota, não um contratempo.
- **Compensação da meia-peça do [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer):** exigi que **4 das minhas 18** sejam remoção
  *irrestrita* (`destroy/exile target permanent`), que é justamente o que o Jailer não é —
  [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king).

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
| [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) | 1 | Sorcery R | queima | criatura (4 dano) | (1) 4 de dano por `{R}` é a melhor taxa do pool e cabe **junto com um loot** no mesmo turno; (2) fonte **vermelha** → 6 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | **sim (caixa)** | R$ 2,48 (12/08) |
| [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) | 1 | Instant R | queima + exílio | criatura pequena, **recursão alheia** | (1) exila — desliga Phoenix/persist/reanimação do oponente, que é o que mais castiga um deck de atrito; (2) 1 mana, fonte vermelha | **sim (caixa)** | `a cotar` |
| [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | 1 | Instant R | queima escalável | criatura **ou planeswalker** | (1) **delirium: 6 de dano** — o único cartão do pool cuja qualidade *cresce* com o cemitério cheio, que é o recurso que o deck fabrica de propósito; (2) fonte vermelha → 8 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell); (3) `{R}` instantâneo deixa mana aberto | não | **R$ 0,75** (12/08) |
| [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) | 1 +Spree | Sorcery W | artefato **e** encantamento | ódio de cemitério, rocha, anthem, equipamento | (1) **o único card do pool que mata artefato E encantamento no mesmo lançamento** (`{3}{W}` pelos dois modos) — o cenário de blowout é [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) **mais** [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage); (2) feitiço = combustível + gatilho de pingador | **sim (caixa)** | `a cotar` |
| [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) | 2 | Instant R | queima flexível | criatura, planeswalker, **jogador** | (1) 3 em **qualquer alvo** — é remoção quando precisa e é relógio quando a mesa está limpa; (2) fonte vermelha → 5 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | **sim (caixa)** | `a cotar` |
| [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) | 2 | Instant R | queima + exílio | criatura **indestrutível**, comandante recorrente | (1) **remove indestrutível e exila** — é a resposta a comandante que volta e a [**Avacyn**](https://www.ligamagic.com.br/?view=cards/card&card=Avacyn)/`Blightsteel`; (2) instantâneo vermelho | **sim (caixa)** | `a cotar` |
| [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) | 2 / 6 | Sorcery R (Overload) | queima + **wipe unilateral** | criatura · **enxame alheio** | (1) `{1}{R}`: 4 de dano numa criatura que **você não controla** — nunca erra de lado; (2) **Overload `{3}{R}{R}{R}`: 4 em cada criatura que você não controla** = wipe 100% assimétrico num slot de remoção; (3) fonte vermelha → 6/6 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | não | **R$ 1,10** (16/09) |
| [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) | 2 | Instant R | modal criatura/artefato | criatura (3) **ou** artefato | (1) modal — cobre dois tipos de ameaça num slot, que é o câmbio certo num deck estourado de slots; (2) instantâneo vermelho | não | `a cotar` |
| [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) | 2 | Instant **RW** | queima + vida | criatura, planeswalker, jogador | (1) é **o gatilho do comandante em carta** — 3 de dano + 3 de vida; (2) os 3 de vida são o que faz o deck chegar ao T7 sob a coroa (a Fase 3 avisou que ser monarca convida ataque); (3) fonte vermelha → 5 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | não | `a cotar` |
| [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) | 2 | Artifact W | exílio permanente | criatura MV ≤3 | (1) exila **permanentemente** o que 3 de dano não resolve (indestrutível, regeneração, ETB abusivo); (2) trava as ameaças de T1–T3, que é a janela em que este deck é mais lento. **F7: fica em campo — não abastece** | **sim (caixa)** | **R$ 0,09** (12/08) |
| [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) | 2 | Sorcery W (flash) | destruição incondicional | criatura **virada** | (1) **spell mastery**: com 2+ mágicas no cemitério vira instantâneo — o deck satisfaz a condição em quase todo turno a partir do T3, então na prática é remoção instantânea de 2 manas; (2) destrói **incondicionalmente**, sem régua de resistência — mata o 8/8 que a queima não pega | **sim (lista)** | `a cotar` |
| [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) | 2 | Instant W | artefato/encantamento | [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace), equipamento, anthem | (1) resposta **instantânea** ao ódio de cemitério — pode ser guardada e usada no turno do oponente; (2) 2 manas deixa espaço para um loot no mesmo turno | **sim (lista)** | `a cotar` |
| [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) | 2 | Enchantment W | artefato/encantamento pré-pago | idem | (1) **pré-pago**: você resolve o custo num turno morto e detona de graça no turno em que precisa do mana para o escape `{R}{R}{W}{W}`; (2) **sai do campo para o cemitério quando usada = combustível que não custou uma carta de mão** | **sim (lista)** | `a cotar` |
| [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) | 3 | Instant R | **irrestrita** | **qualquer permanente** | (1) a única resposta RW a permanente com hexproof/indestrutível/proteção fora de um wipe; (2) instantânea — responde à peça de combo no turno do oponente; (3) fonte vermelha (sem dano, mas dentro do plano) | não | **R$ 3,81** (24/08) |
| [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) | 3 | Instant W | **irrestrita** | **qualquer permanente** | (1) segunda cópia funcional do [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), e sem o risco de revelar um permanente melhor; (2) o 3/3 Elefante que ela concede é alvo legítimo do Helix do Phlage e morre a qualquer wipe do deck | não | **R$ 5,00** (19/09) |
| [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) | 3 | Sorcery W | **irrestrita, exílio** | **qualquer não-terreno** | (1) **exila** — a resposta mais limpa do pool sem custo; (2) o 3/2 concedido é alvo do Helix e some no [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate); (3) feitiço = combustível | **sim (lista)** | `a cotar` |
| [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) | 3 (fuse) | Instant // Instant RW | artefato **e** encantamento | ódio de cemitério, rocha, anthem | (1) **fuse**: `{1}{R}{W}` mata artefato **e** encantamento no mesmo lançamento, **em instantâneo** — é a [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) com o timing que a [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) não tem; (2) as duas metades são conjuráveis isoladas (`{W}` só Tear) — nunca é carta morta | não | `a cotar` |
| [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) | 4 | Enchantment W | **irrestrita ×3** | 1 não-terreno **por oponente** | (1) **três remoções irrestritas num card** num pod de 4 — a maior densidade de resposta do pool inteiro; (2) `recruit` = compra-e-descarta = **combustível** + um corpo 1/1 que conta no piso da §4 da Fase 2 | **sim (lista)** | `a cotar` |

**+ 2 da Fase 3, contadas aqui (§1):** [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (remoção repetível de 7/turno) e
[**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) (exílio condicional). **Total no deck: 20.**

### 2.2 Proteção — o dimensionamento (dor 4 do intake)

**Pergunta:** a recursão do comandante basta, ou o deck precisa de proteção dedicada?
**Resposta: 2 slots, e apenas 1 deles é compra.** A conta:

| O que morre num wipe | Volta sozinho? | Custo de reconstruir |
|---|---|---|
| **Phlage** | **Sim** — escape `{R}{R}{W}{W}`, custo fixo, sem imposto de comandante | 4 manas + 5 cartas do cemitério. E o wipe **encheu** o cemitério. |
| 4 motores de combustível ([**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation), [**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece), [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin), [**Tablet of Discovery**](https://www.ligamagic.com.br/?view=cards/card&card=Tablet+of+Discovery)) | não | 2–3 manas cada, mas **o motor é o que paga o 2º e o 3º escape** |
| 5–7 pingadores (o wincon da §6 da Fase 2) | não | 2–3 manas cada — **é o plano de vitória inteiro** |
| 3–4 multiplicadores ([**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), `Solphim`) | não | 4 manas |

A leitura honesta: **a recursão do comandante resolve a peça que menos precisava de resolução.**
O que o deck de fato perde num wipe alheio é o pacote de alcance e os motores — e nenhum deles
volta. Mas três fatos comprimem o tamanho da resposta:

1. **Você é quem dá o wipe.** Com 4 wipes na lista, o deck escolhe o turno da limpeza.
2. **O que morre é barato.** O pacote inteiro está em CMC 2–4; reconstruir é 2–3 turnos, não a
   partida. A v1 perdia o *plano*; a v2 perde *tempo*.
3. **A ameaça real não é wipe, é ódio de cemitério.** [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) desliga o escape, o
   `Palantír`, o [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder), o [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) e o delirium da [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) —
   quatro fases de uma vez. **O orçamento de "proteção" deste deck está gasto em remoção de
   encantamento** (9 cartas, §4), não em hexproof.

| Carta | CMC | Tipo | Subcategoria | O que responde | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) | 2 | Instant RW | **proteção em massa** | wipe alheio (destroy), remoção pontual, [**Cyclonic Rift**](https://www.ligamagic.com.br/?view=cards/card&card=Cyclonic+Rift) não | (1) *"**Permanents** you control gain indestructible"* — **todos os permanentes, não só criaturas**: salva motores, pingadores e multiplicadores de uma vez, por 2 manas em instantâneo; (2) é o que faz o **seu** [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate)/[**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) ser unilateral de verdade; (3) modo 2 = 4 na cara (alcance) e modo 3 = double strike no Phlage 6/6 = 12 de combate. **Três modos, três funções.** | não | `a cotar` |
| [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) | 1 | Instant W | proteção pontual | remoção mirada no [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)/pingador-chave | (1) `{W}` para salvar o multiplicador, que é a peça cujo corte custa o fecho do jogo; (2) **scry 1** — nunca é carta morta num deck que precisa achar o 4º terreno; (3) custo **R$ 0**, já está fisicamente na `lista.txt` | **sim (lista)** | `a cotar` |

**Por que não um 3º e 4º slot de proteção:** cada slot adicional compra seguro para permanentes
que custam 2 manas, num deck já estourado de slots, e *não* cobre o único desastre irrecuperável
(ódio de cemitério), que se responde com remoção. [**Feat of Resistance**](https://www.ligamagic.com.br/?view=cards/card&card=Feat+of+Resistance), [**Adamant Will**](https://www.ligamagic.com.br/?view=cards/card&card=Adamant+Will) e
[**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) (todas na `lista.txt`, R$ 0) ficam de fora por isso — ficha na §7.

---

## 3. Candidatas — board wipes

**4 wipes, dos quais 2 são assimétricos por construção e 1 é seletivo.** A tabela mede o que
cada um leva do **meu** board típico de T6–T8 (Phlage 6/6 escapado + 2–3 pingadores de poder 0–2,
resistência 3–4 + [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) 2/4 + [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) 0/1):

| Carta | CMC | Simétrico? | Quanto leva do **meu** board | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|
| [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) | 3 | **NÃO — unilateral por construção do deck** | *Cada jogador escolhe criaturas com **poder total ≤ 4** e sacrifica o resto.* Meu pacote é [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) 0/4, [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) 0/3, [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) 0/1, [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) 2/2, [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) 2/4 — **poder total 0+0+0+2 = 2**: eu fico com **quatro deles**. Perco só o Phlage 6/6, **que quer estar no cemitério**. | (1) **é sacrifício, não destruição** — ignora indestrutível, hexproof e proteção; (2) a assimetria é *o piso de corpos da §4 da Fase 2*: o deck foi construído com criaturas de poder 0–2, e esta carta cobra exatamente por poder; (3) CMC 3 = cabe no mesmo turno em que sobra mana | não | `a cotar` |
| [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) | 2 | **NÃO — seletivo** | Com **0 contadores**: mata **todas as fichas da mesa** e nada meu (nenhuma das minhas peças é MV 0). Com 2: leva meus 2-drops junto. Você escolhe. | (1) **a resposta a enxame alheio sem perder uma carta** — a lacuna que a v1 fechou declarando que não tinha; (2) o deck vive em CMC 1–3 e os oponentes não: você aponta para a curva deles; (3) artefato barato que entra no T2 e espera | **sim (caixa)** | **R$ 1,79** (12/08) |
| [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) | 5 | Sim no texto, **não no custo** | Leva tudo, inclusive Phlage. Perco 3–4 permanentes de 2–3 manas; a mesa perde o investimento inteiro dela. | (1) **ganha 1 de vida por criatura destruída** — num wipe de 8 criaturas são 8 de vida, e estabilizar é metade do plano de atrito; (2) os meus mortos **vão para o meu cemitério** = 3–4 cartas de combustível de escape entregues no mesmo turno; (3) [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) o transforma em unilateral | **sim (caixa)** | **R$ 1,37** (12/08) |
| [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) | 4 | Sim, mas **escalável** | X = nº de criaturas na mesa. Com 4 criaturas leva meus x/3 e x/4; com 8 leva tudo. Contra mesa pequena é barato e não dispara. | (1) **fonte vermelha** — com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) mata +2 e com [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) triplica: o único wipe do pool que os multiplicadores do deck melhoram; (2) escala com a ameaça: quanto maior o enxame alheio, mais ele mata; (3) feitiço = combustível | não | **R$ 0,90** (18/09) |

**+ 1 efeito de massa que não gasta slot de wipe:** **[**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars)** com Overload
(`{3}{R}{R}{R}`) = 4 de dano em **cada criatura que você não controla**. Assimetria total, e ele
já está contado como remoção na §2.1. Na prática o deck tem **5 respostas a enxame**.

**Reservas de wipe (Fase 6 decide):**

| Carta | CMC | Por que é reserva | Na coleção? | Preço |
|---|---|---|---|---|
| [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) | — (Land) | Mesmo efeito do [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) **sem gastar slot de mágica** — entra na conta da Fase 6, não na minha | **sim (caixa)** | **R$ 0,89** (12/08) |
| [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) | 9 → `{R}` | 13 de dano a cada criatura, custando `{R}` numa mesa cheia; fonte vermelha. **Upgrade direto sobre [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction)** se a cotação fechar abaixo de ~R$ 3 | não | `a cotar` |
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | X | Wipe cirúrgico **e** finisher **e** combustível (flashback descarta X). Já está no pacote da Fase 2 — não duplico o slot | não | `a cotar` |
| [**Winds of Abandon**](https://www.ligamagic.com.br/?view=cards/card&card=Winds+of+Abandon) | 2 / 6 | `{1}{W}` exila 1 criatura; Overload `{4}{W}{W}` **exila todas que você não controla**. É o melhor wipe unilateral que RW tem. **Fora por preço: R$ 19,41 = 9,7% do teto** num deck que a Fase 2 já declarou "majoritariamente compra". Upgrade sobre [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) **se o total fechar com folga** | não | **R$ 19,41** (19/09) |

---

## 4. Cobertura de ameaças

| Tipo de ameaça | Quantas respostas | Quais |
|---|---|---|
| **Criaturas** | **17** (+ o comandante, todo turno) | 12 das 18 titulares + os 4 wipes + [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire). O Helix do Phlage são 3 de dano garantidos a cada entrada **e** a cada ataque. |
| **Artefatos** | **10** | [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) · [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) · [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) · [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) · [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) · [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) · [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) |
| **Encantamentos** | **9** | [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) · [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) · [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) · [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) · [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) · [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) |
| **Planeswalkers** | **10** | [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) · [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) · [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) · [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) · `Celebrate` · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) · [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (7/turno) · Helix do Phlage |
| **Ameaça em massa (enxame alheio)** | **5 titulares + 3 reservas** | [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) (Overload, unilateral) · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) (seletivo) · [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) (unilateral) · [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) · [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) — reservas: [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone), [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act), [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate). **A lacuna declarada da v1 está fechada, e nominalmente.** |
| **Flexível — "destroy/exile target permanent"** | **4** | [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) · [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) · [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) (×3 alvos) |
| **Indestrutível / hexproof / proteção** | **6** | [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) (sacrifício, sem alvo) · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) / [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) (sem alvo) · [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (sem alvo) · [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) (remove indestrutível) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) (embaralha) |
| **Ódio de cemitério** ([**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace), [**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog), [**Soulless Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Soulless+Jailer), [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage)) | **9** | as 9 respostas a encantamento acima, das quais **4 em instantâneo** ([**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant), [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp)). A Fase 2 exigiu 2; a Fase 3 elevou a obrigatórias porque o pacote de saque dela também morre. **Entreguei 9, e é de propósito** — ver §2.2. |
| **Combos / mágicas na pilha** | **0 — DESCOBERTO** | RW **não tem contramágica**. Ver §5.1. |

### 4.1 Mana aberto — distribuição instantâneo/feitiço

De **22 peças** (18 remoções + 4 wipes):

| Velocidade | Nº | % | Quais |
|---|---|---|---|
| **Responde no turno do oponente** | **13** | **59%** | 10 instantâneos verdadeiros ([**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray), [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat), [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless), [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade), [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant), [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear)) + [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) (spell mastery ligado a partir do T3, na prática sempre) + [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) (sacrifício em instantâneo) + [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) (sacrifício em instantâneo) |
| Só no seu turno | 9 | 41% | [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash), [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid), [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars), [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), `Celebrate`, [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate), [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong), [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) |

**Leitura honesta:** 59% é aceitável, não excelente — um controle puro miraria 65–70%. A defesa é
que **os feitiços são os mais baratos da lista** ([**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) `{R}`, [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) `{1}{R}`,
[**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) `{W}`), então a perda de tempo por feitiço é de 1–2 manas, não de um turno.
Empurrar acima de 59% exigiria trocar [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) (4 dano por `{R}`, R$ 0) por instantâneos de
taxa pior e de compra — câmbio ruim num deck estourado de slots e de compra.

### 4.2 Curva

**CMC ≤ 2: 13 das 18** ([**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) 1, [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) 1, [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) 1, [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) 1,
[**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) 2, [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) 2, [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) 2, [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) 2, [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) 2,
[**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) 2, [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) 2, [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) 2, [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) 2)
**CMC 3: 4** ([**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear))
**CMC 4: 1** ([**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king))
**Wipes: 2, 3, 4, 5** ([**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb), [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong), [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction), [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate))

A Fase 2 pediu **≥8 em CMC ≤ 2**. Entreguei **13**. Isso é o que permite conjurar remoção e loot
no mesmo turno até o T4, e remoção + segurar o `{R}{R}{W}{W}` do escape a partir do T5.

### 4.3 Remoção que também é combustível (item 2 do briefing desta fase)

Toda mágica de remoção vai ao cemitério **uma vez** e vira 1 de 5 cartas do escape. As que fazem
**mais** que isso, declaradas:

| Carta | Dupla função de cemitério |
|---|---|
| [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | **Relação inversa** — quanto mais cheio o cemitério, melhor a carta. Delirium (4 tipos entre Instant/Sorcery/Artifact/Creature/Enchantment/Land) fica ligado a partir do T5 neste deck: os loots despejam terreno, o [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) despeja terreno por retrace, [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) despeja encantamento, [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin)/[**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece) despejam artefato. **2 de dano vira 6.** |
| [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) | Idem — spell mastery (2+ instantâneos/feitiços no cemitério) transforma feitiço em instantâneo. É a única remoção do pool que **melhora de velocidade** pelo cemitério. |
| [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) | Sai do **campo** para o cemitério ao ser usada: é combustível que **não custou uma carta da mão**. |
| [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) | `recruit` = compra 1, descarta 1 → **+1 carta no cemitério além dela mesma**, e um corpo 1/1. |
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) (reserva) | Flashback custa *descartar X cartas* — a única remoção do pool que é um **motor** de combustível, não uma carga. |
| [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) (pacote da Fase 2) | Retrace: **nunca sai do cemitério** e cada uso descarta um terreno. É remoção repetível contada pela Fase 2 — declaro a sobreposição para ninguém contar duas vezes. |

**Anti-combustível declarado (F7):** [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) e [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) são permanentes — ficam em
campo e não abastecem nada até morrerem. São os **únicos 2** dos 22, e ambos entram por função
que nenhuma mágica cobre (exílio permanente de MV≤3 e wipe seletivo por MV). Custo aceito e
nomeado.

---

## 5. Buracos declarados

### 5.1 Contramágica: **zero, e sem mitigação plena**

RW não tem counterspells jogáveis. O deck **não responde a** [**Thassa's Oracle**](https://www.ligamagic.com.br/?view=cards/card&card=Thassa%27s+Oracle), [**Ad Nauseam**](https://www.ligamagic.com.br/?view=cards/card&card=Ad+Nauseam),
[**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun), nem a um combo que ganhe na pilha. As mitigações reais, nomeadas:

- **Peça de combo que é permanente** (a maioria em pod casual/mid): [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) e
  [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) respondem **em instantâneo**, e [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb)/[**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) respondem por MV
  sem alvo. Isso cobre [**Isochron Scepter**](https://www.ligamagic.com.br/?view=cards/card&card=Isochron+Scepter), [**Dramatic Reversal**](https://www.ligamagic.com.br/?view=cards/card&card=Dramatic+Reversal), [**Kiki-Jiki**](https://www.ligamagic.com.br/?view=cards/card&card=Kiki-Jiki%2C+Mirror+Breaker), [**Food Chain**](https://www.ligamagic.com.br/?view=cards/card&card=Food+Chain).
- **Combo de dano letal:** **[**Deflecting Palm**](https://www.ligamagic.com.br/?view=cards/card&card=Deflecting+Palm)** `{R}{W}` previne o dano e o devolve ao
  controlador da fonte. **Deixo como reserva nomeada**, não titular — é situacional demais para
  um dos 20 slots, mas é a única carta de RW que "contra" um [**Comet Storm**](https://www.ligamagic.com.br/?view=cards/card&card=Comet+Storm).
- **O que fica sem resposta:** vitória alternativa e combo que resolve na pilha. **Declarado.**
  Este é o imposto da identidade de cor escolhida na Fase 1, não uma falha de seleção.

### 5.2 Céu

O deck bloqueia com [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) 0/4, [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) 0/3 e Phlage 6/6 — **nenhum com
alcance ou voo**. Contra um deck de voadores, a resposta é remoção pontual e wipe, não bloqueio.
[**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) (candidata nº 2 da Fase 7) é 5/5 voadora com first strike e cobriria
isso — registro a sobreposição para a Fase 7, sem pedir o slot.

### 5.3 Risco de percepção — carregado da Fase 2

A Fase 2 declarou que **27,9% das mãos abrem sem criatura**, contra 12,1% da v1 corrigida, e que o
usuário reprovou a v1 exatamente por *"muita mágica instantânea e encantamento na mão e poucas
criaturas"*. **Esta fase é a que mais agrava esse número**: 18 das minhas 22 peças são mágicas ou
encantamentos. Não tenho como reduzi-lo sem abandonar o eixo. **Sinalizo de novo ao orquestrador**
— é a pergunta que precisa ir ao usuário no `report.md`, não uma nota de rodapé.

---

## 6. Pendências devolvidas pela Fase 2 — decididas

### 6.1 [**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon) (caixa, R$ 1,25) — **NÃO ENTRA**

Ficha F1–F7 completa (regra 4), oracle puxado nesta sessão:

> `{T}`: Put a charge counter on this artifact.
> `{T}`, Remove three charge counters from this artifact: Destroy target permanent.

| Eixo | Conteúdo |
|---|---|
| **F1 texto** | Duas ativadas. A segunda é **remoção irrestrita e repetível** — destrói *qualquer* permanente, incluindo terreno, e ignora hexproof? **Não**: ela mira. Ignora indestrutível? Não, destrói. É irrestrita quanto ao **tipo**, não quanto à proteção. |
| **F2 corpo** | Não é criatura, não vira criatura. Tapa **só para si mesmo** — o deck não tem crew, station, convoke nem improvise, então o `{T}` não é recurso compartilhado. Não bloqueia. Não se sacrifica por valor. |
| **F3 tipo como recurso** | Artefato. **O pool da v2 não conta artefatos para nada** — o metalcraft do [**Jor Kadeen**](https://www.ligamagic.com.br/?view=cards/card&card=Jor+Kadeen%2C+the+Prevailer) morreu com o eixo go-wide, e nenhuma carta das 48 candidatas da Fase 2 lê "number of artifacts". Valor: **zero**. |
| **F4 receptor** | Recebe charge counters **de si mesma**. Não há proliferate nem adicionador de contadores no pool. Não recebe anthem (não é criatura). Valor: zero. |
| **F5 facilitador** | Não dá nada a nenhuma outra carta. |
| **F6 curva** | CMC 4. Conjurada no T4, tapa no T4/T5/T6 → **primeira destruição no T7**. E o T4 é o turno do primeiro escape (`{R}{R}{W}{W}`) e do [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) — ela **disputa o turno mais congestionado do deck**. |
| **F7 atrito** | (a) Nunca vai ao cemitério sozinha — **num deck cujo recurso escasso é "cartas no cemitério", um artefato que fica em campo é anti-combustível**; (b) disputa o T4 com escape e multiplicador; (c) como é permanente conhecido e lento, é o alvo óbvio da remoção da mesa nos 3 turnos em que não faz nada. |

**Protocolo de corte — funções e quem as cobre:**

| Função | Quem cobre depois da dispensa |
|---|---|
| remoção **irrestrita** (qualquer tipo de permanente) | [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) — **4 peças, todas em 3–4 manas, duas delas instantâneas**, contra 4 manas + 3 turnos |
| remoção **repetível** | [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (7 de dano por upkeep, funciona **no turno seguinte** ao ETB) e [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) (retrace, infinito) — a função é coberta com folga e com velocidade |
| artefato para contagens | **nenhuma contagem existe** — função inexistente, nada fica descoberto |

**Nenhuma função fica descoberta.** A Fase 2 reprovou por velocidade; a ficha completa mostra que
velocidade é o **melhor** dos seus eixos: os outros seis são zero ou negativos. Não é "carta
fraca" — é carta cuja única função já tem 4 substitutos mais rápidos e mais baratos em mana, num
deck onde o slot é mais escasso que o dinheiro. **R$ 0 abrem mão; aceito.**

### 6.2 [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) (caixa, R$ 0) — **ENTRA** (titular)

Aberta desde 2026-09-20 e adiada até o eixo fechar. O eixo fechou. Oracle:

> Spree — {W} base · +{1} Destroy target artifact · +{1} Destroy target enchantment ·
> +{1} Put a +1/+1 counter on each creature target player controls.

| Eixo | Conteúdo |
|---|---|
| **F1** | Modal aditivo: paga-se `{2}{W}` por um dos dois primeiros modos e `{3}{W}` **pelos dois**. |
| **F2** | Sem corpo. |
| **F3** | **Sorcery** — alimenta o spell mastery da [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning), dispara [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe)/[**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field)/[**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer)/[**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) e vai ao cemitério = combustível de escape. |
| **F4** | Não recebe nada. |
| **F5** | O 3º modo dá +1/+1 a todas as criaturas de um jogador — **marginal aqui** (4 criaturas pequenas, algumas com defender). Não é o motivo da entrada. |
| **F6** | T3 para um modo, T4 para os dois. |
| **F7** | **Feitiço** — não responde no turno do oponente, que é a preferência declarada do eixo. Disputa função com [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant), [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) e [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear). |

**Por que entra (regra 3 — 2 pontos):** (1) é a **única** carta do pool que mata artefato **e**
encantamento no mesmo lançamento, e o cenário que mata este deck não é uma peça de ódio de
cemitério, são **duas** ([**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) + [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage)/[**Soulless Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Soulless+Jailer)); (2) é feitiço
que abastece o cemitério e dispara o pacote de pingadores, que é o fecho do jogo.

**A ressalva de 2026-09-20 caiu junto com o eixo.** Ela dizia *"contadores morrem com as
criaturas, então não ajuda na dor 4"* — o 3º modo não é por que ela entra, e a troca proposta lá
([**Basri's Solidarity**](https://www.ligamagic.com.br/?view=cards/card&card=Basri%27s+Solidarity) → [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid)) não existe mais: [**Basri's Solidarity**](https://www.ligamagic.com.br/?view=cards/card&card=Basri%27s+Solidarity) está fora do
pool desde a §8.2 da Fase 2. **A ressalva de velocidade continua válida e está coberta**:
[**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) e [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) respondem em instantâneo quando o timing importa.

---

## 7. Varredura da caixa e da `lista.txt` — dispensas justificadas (regra 7)

`bin/mtgdb collection -list` (116 cartas) + 69 não-básicas de `lista.txt`. **14 das 22 titulares
saíram daí, a R$ 0 de aquisição.** O que ficou de fora, com ficha:

| Carta | Origem | Ficha resumida | Por que **não** cobre a função |
|---|---|---|---|
| [**Twin Bolt**](https://www.ligamagic.com.br/?view=cards/card&card=Twin+Bolt) | caixa | F1 2 de dano dividido entre 1–2 alvos · F6 CMC 2 · F3 instantâneo = combustível | **Taxa, não qualidade.** 2 de dano por 2 manas não mata nada relevante em Commander a partir do T4. [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) ocupa a mesma linha da curva com 3 de dano **ou** destruição de artefato. Fica como **reserva de orçamento**: se alguma compra estourar, ela volta a R$ 0. |
| [**Fateful End**](https://www.ligamagic.com.br/?view=cards/card&card=Fateful+End) | caixa | F1 3 de dano em qualquer alvo + scry 1 · F6 CMC 3 | Mesma função do [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) **por 1 mana a mais**, e o scry 1 não paga a diferença num deck que precisa dos T2–T4 livres para remoção + loot. **Reserva de orçamento.** |
| [**Seal of Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Fire) | caixa | F1 2 de dano pré-pagos; sai do campo para o cemitério = combustível · F6 CMC 1 | 2 pontos reais (pré-pago + combustível), mas **2 de dano** é a menor taxa do pool e o slot compete com [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) (2 + **exílio**, instantâneo, mesmo custo). **Reserva.** |
| [**Chandra's Pyrohelix**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Pyrohelix) · [**Searing Barrage**](https://www.ligamagic.com.br/?view=cards/card&card=Searing+Barrage) · [**Radiating Lightning**](https://www.ligamagic.com.br/?view=cards/card&card=Radiating+Lightning) · [**Stonefury**](https://www.ligamagic.com.br/?view=cards/card&card=Stonefury) · [**Seismic Wave**](https://www.ligamagic.com.br/?view=cards/card&card=Seismic+Wave) · [**Lightning Volley**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Volley) · [**Molten Blast**](https://www.ligamagic.com.br/?view=cards/card&card=Molten+Blast) · [**Chandra's Outrage**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Outrage) · [**Bombard**](https://www.ligamagic.com.br/?view=cards/card&card=Bombard) · [**Pinecone Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Pinecone+Strike) · [**Smaug's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Smaug%27s+Fury) · [**Mercadia's Downfall**](https://www.ligamagic.com.br/?view=cards/card&card=Mercadia%27s+Downfall) | caixa | queima de Limited, CMC 2–5, 2–4 de dano, todas abastecem o cemitério | **Dispensa por taxa (F6), com a ficha feita** — cada uma cabe no eixo e nenhuma é "carta fraca". Mas são **20 slots de remoção**, e [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) (4 por `{R}`), [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) (6 com delirium por `{R}`) e [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) (4 por `{1}{R}`, que não erra de lado) dominam todas em dano-por-mana. A Fase 2 já as classificou como reserva de orçamento; confirmo. |
| [**Punishing Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Punishing+Fire) | caixa | F1 2 de dano; volta do cemitério quando **um oponente** ganha vida · F6 CMC 2 | Confirmo a dispensa da Fase 2 e acrescento o eixo que faltava: cada retorno **tira uma carta do cemitério**. Numa lista cujo recurso escasso é *cartas no cemitério*, recursão de mágica é **anti-combustível** — é a mesma razão pela qual a Fase 2 limitou flashback a 3 cartas. |
| [**Skullcrack**](https://www.ligamagic.com.br/?view=cards/card&card=Skullcrack) | caixa | F1 3 num jogador; **ninguém ganha vida neste turno** · F6 CMC 2 | Confirmo: a cláusula é **simétrica** e desliga os 3 de vida do próprio Helix e o ganho do [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) — os dois estabilizadores do deck. Atrito F7 direto com o comandante. |
| [**Lux Artillery**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Artillery) | caixa | F1 artefato CMC 4; converte contadores de carga em dano | Mesma ficha do [**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon) (§6.1): F2–F5 zerados, CMC 4, não vai ao cemitério, e o deck não tem nenhum outro gerador de contadores de carga para alimentá-la. Pior que o [**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon) porque nem remoção irrestrita entrega. |
| [**Destructive Tampering**](https://www.ligamagic.com.br/?view=cards/card&card=Destructive+Tampering) · [**Smashing Success**](https://www.ligamagic.com.br/?view=cards/card&card=Smashing+Success) · [**Steel Wrecking Ball**](https://www.ligamagic.com.br/?view=cards/card&card=Steel+Wrecking+Ball) · [**Thaumaton Torpedo**](https://www.ligamagic.com.br/?view=cards/card&card=Thaumaton+Torpedo) | caixa | remoção de artefato, CMC 1–5 | Função **já coberta 10×** (§4), e por cartas que também pegam encantamento. [**Steel Wrecking Ball**](https://www.ligamagic.com.br/?view=cards/card&card=Steel+Wrecking+Ball) CMC 5 e [**Smashing Success**](https://www.ligamagic.com.br/?view=cards/card&card=Smashing+Success) CMC 4 são caras para função de resposta; [**Thaumaton Torpedo**](https://www.ligamagic.com.br/?view=cards/card&card=Thaumaton+Torpedo) `{1}` só pega artefato. Nenhuma alcança 2 pontos de sinergia. |
| [**Return to Nature**](https://www.ligamagic.com.br/?view=cards/card&card=Return+to+Nature) · [**Smell Fear**](https://www.ligamagic.com.br/?view=cards/card&card=Smell+Fear) · [**Road // Ruin**](https://www.ligamagic.com.br/?view=cards/card&card=Road+%2F%2F+Ruin) · [**Ancient Animus**](https://www.ligamagic.com.br/?view=cards/card&card=Ancient+Animus) · [**Ambitious Assault**](https://www.ligamagic.com.br/?view=cards/card&card=Ambitious+Assault) | caixa | — | **Fora da identidade de cor.** Todas verdes ou com metade verde. |
| [**Negate**](https://www.ligamagic.com.br/?view=cards/card&card=Negate) · [**Stoic Rebuttal**](https://www.ligamagic.com.br/?view=cards/card&card=Stoic+Rebuttal) · [**Disruption Protocol**](https://www.ligamagic.com.br/?view=cards/card&card=Disruption+Protocol) | caixa | contramágica | **Fora da identidade de cor** (azuis). São a razão pela qual a §5.1 existe: a caixa *tem* counters, e nenhum é legal aqui. |
| [**Leonin Bola**](https://www.ligamagic.com.br/?view=cards/card&card=Leonin+Bola) | caixa | F1 equipamento `{1}`, `{T}` + unequip: vira criatura-alvo · F2 dá função de tap ao equipado | Pseudo-remoção (só vira, não mata) que **exige uma criatura desvirada** para funcionar. Este deck tem 14–17 corpos e vários com `defender`; e o corpo que importa (Phlage) quer estar atacando. 1 ponto. |
| [**Joust**](https://www.ligamagic.com.br/?view=cards/card&card=Joust) · [**Moment of Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Moment+of+Glory) · [**Phyrexian Revoker**](https://www.ligamagic.com.br/?view=cards/card&card=Phyrexian+Revoker) | caixa | luta / +1/+1 / trava ativada | [**Joust**](https://www.ligamagic.com.br/?view=cards/card&card=Joust) exige criatura maior que a do oponente (o deck tem poder 0–2 fora do Phlage). [**Moment of Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Moment+of+Glory) é pump, categoria morta neste eixo. [**Phyrexian Revoker**](https://www.ligamagic.com.br/?view=cards/card&card=Phyrexian+Revoker) é stax pontual — **regra 11**, e o deck não pediu. |
| [**Feat of Resistance**](https://www.ligamagic.com.br/?view=cards/card&card=Feat+of+Resistance) · [**Adamant Will**](https://www.ligamagic.com.br/?view=cards/card&card=Adamant+Will) · [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) | lista | proteção pontual de 2 manas | Função coberta por [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) (**todos** os permanentes, não só criaturas) e [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) (1 mana, com scry). [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) exige **sacrificar uma criatura** como custo adicional — num deck de 14–17 corpos em que 6 **são** o wincon, é anti-sinergia direta com a §6.1 da Fase 2. Ver o dimensionamento de proteção na §2.2: 2 slots, não 5. |
| [**Invoke the Divine**](https://www.ligamagic.com.br/?view=cards/card&card=Invoke+the+Divine) · [**Expose to Daylight**](https://www.ligamagic.com.br/?view=cards/card&card=Expose+to+Daylight) | lista | destrói artefato/encantamento (+4 vida / +scry) | **Reservas legítimas, a R$ 0.** Ficam fora porque já selecionei 4 respostas a artefato/encantamento e as duas custam 3 manas contra os 2 do [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant). Se a Fase 6 quiser uma 5ª resposta sem gastar dinheiro, é daqui que ela sai — [**Invoke the Divine**](https://www.ligamagic.com.br/?view=cards/card&card=Invoke+the+Divine) primeiro (os 4 de vida valem no plano de atrito). |
| [**Wedgelight Rammer**](https://www.ligamagic.com.br/?view=cards/card&card=Wedgelight+Rammer) · [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais) · [**Panic Spellbomb**](https://www.ligamagic.com.br/?view=cards/card&card=Panic+Spellbomb) · [**Liquimetal Coating**](https://www.ligamagic.com.br/?view=cards/card&card=Liquimetal+Coating) · [**Titan Forge**](https://www.ligamagic.com.br/?view=cards/card&card=Titan+Forge) · [**Darksteel Reactor**](https://www.ligamagic.com.br/?view=cards/card&card=Darksteel+Reactor) | caixa | — | Nenhuma é interação. Fora do escopo desta fase. |
| **[**Swords to Plowshares**](https://www.ligamagic.com.br/?view=cards/card&card=Swords+to+Plowshares)** | compra | F1 exila criatura por `{W}` em instantâneo; controlador ganha vida igual ao poder | **Dispensa por preço, e a função está coberta.** R$ 13,99 = **7,0% do teto** por uma das 20 remoções, num deck em que 14 das 22 são R$ 0. O ganho de vida também trabalha contra um plano que fecha por dano. Exílio de criatura fica com [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray), [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless), [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) e `Celebrate`. **Recomendo cotar [**Path to Exile**](https://www.ligamagic.com.br/?view=cards/card&card=Path+to+Exile)** (mesma função, cotação desconhecida) antes de desistir da linha. |
| **[**Vandalblast**](https://www.ligamagic.com.br/?view=cards/card&card=Vandalblast)** | compra | wipe unilateral de artefatos | R$ 39,99 = **20% do teto**. Confirmo a exclusão da Fase 2 — por preço, não por função. [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) a 0 contadores e [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) cobrem parte disso. |

---

## 8. Reservas — a fila, na ordem

Se alguma cotação estourar ou a Fase 6 precisar de slot, esta é a ordem de substituição:

| Ordem | Entra | Sai | Por quê | Preço |
|---|---|---|---|---|
| 1 | [**Twin Bolt**](https://www.ligamagic.com.br/?view=cards/card&card=Twin+Bolt) (caixa) | a compra mais cara entre [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) / [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) | R$ 0 na mesma linha da curva; perde 1–2 de dano | R$ 0 |
| 2 | [**Invoke the Divine**](https://www.ligamagic.com.br/?view=cards/card&card=Invoke+the+Divine) (lista) | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) | R$ 0; perde a metade em instantâneo do artefato + 1 mana | R$ 0 |
| 3 | [**Fateful End**](https://www.ligamagic.com.br/?view=cards/card&card=Fateful+End) / [**Seal of Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Fire) (caixa) | qualquer compra `a cotar` que estoure | R$ 0, taxa pior | R$ 0 |
| 4 | [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) | [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) | **upgrade**, se cotar abaixo de ~R$ 3 | `a cotar` |
| 5 | [**Winds of Abandon**](https://www.ligamagic.com.br/?view=cards/card&card=Winds+of+Abandon) | [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) | **upgrade** para o melhor wipe unilateral de RW — só se o total fechar com folga | R$ 19,41 |
| 6 | [**Fateful Absence**](https://www.ligamagic.com.br/?view=cards/card&card=Fateful+Absence) | [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) | destruição **incondicional** de criatura/PW por `{1}{W}`; troca 3 de vida por um Clue | `a cotar` |
| 7 | [**Deflecting Palm**](https://www.ligamagic.com.br/?view=cards/card&card=Deflecting+Palm) | — (entra só se a mesa tiver combo de dano) | única resposta RW a [**Comet Storm**](https://www.ligamagic.com.br/?view=cards/card&card=Comet+Storm)/[**Fireball**](https://www.ligamagic.com.br/?view=cards/card&card=Fireball) letal | `a cotar` |
| 8 | [**Reckless Rage**](https://www.ligamagic.com.br/?view=cards/card&card=Reckless+Rage) | — | 4 de dano por `{R}`; **exige criatura sua** (2 de dano nela), e 27,9% das mãos abrem sem corpo | `a cotar` |
| 9 | [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) (caixa) | — (slot de **terreno**, Fase 6) | 5º efeito de massa sem gastar slot de mágica | R$ 0,89 |

---

## 9. Custo

**Origem: LigaMagic (menor), cotações de 12/08 a 19/09/2026.** Nenhum número do Scryfall.

**Aquisição (o que precisa ser comprado) — 8 cartas:**

| Carta | Preço |
|---|---|
| [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | R$ 0,75 (12/08) |
| [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) | R$ 0,90 (18/09) |
| [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) | R$ 1,10 (16/09) |
| [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) | R$ 3,81 (24/08) |
| [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) | R$ 5,00 (19/09) |
| [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) | **a cotar** |
| [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) | **a cotar** |
| [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) | **a cotar** |
| [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) | **a cotar** |
| [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) | **a cotar** |
| **Subtotal cotado** | **R$ 11,56** · **5 cartas a cotar** |

**Valor no deck das 14 que já existem fisicamente** (regra 2 mede o *valor do deck*, não o gasto):
[**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) R$ 2,48 · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) R$ 1,79 · [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) R$ 1,37 · [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) R$ 0,09 =
**R$ 5,73 cotados**; [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray), [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless), [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid),
[**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning), [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant), [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory),
[**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king), [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) — **10 a cotar** (todas commons/uncommons).

**Total desta fase: R$ 17,29 cotados + 15 cartas a cotar.** As 15 pendentes são majoritariamente
commons e uncommons; a exposição real está em [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) e [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), que são as duas
únicas com histórico de preço alto no formato. **Cotar as 15 antes de a Fase 6 fechar cortes.**

---

## 10. Sinais cruzados

| Fase | Sinal |
|---|---|
| **2 · tema** | Restrições cumpridas: **13** em CMC ≤2 (pedido: ≥8) · **9** respostas a encantamento (pedido: 2) · preferência por instantâneo/feitiço sobre permanente respeitada (só 2 dos 22 são permanentes) · preferência por `{R}` sobre `{W}` respeitada onde empatou. **[**Lux Cannon**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Cannon) não entra** (§6.1, ficha completa); **[**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) entra** (§6.2). Sobreposição a declarar para não haver dupla contagem: **[**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) e [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) já estão no pacote de combustível da Fase 2** e são também remoção — eu **não** os contei nos meus 18. |
| **3 · draw** | **Pedido aceito, com correção**: [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) vale 1,0 slot de remoção, [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) vale 0,5 (o exílio se desfaz quando a coroa é roubada — oracle conferido). Cota fixada em **20 no ponto**, não 19–20. Os 2 slots que você pediu estão pagos. **E você não precisa do segundo pedido**: como devolvo 2,5 slots, o deck fecha em 60,0/62 e os **10 loots ficam intactos** — a config D (14 fontes, P(escape T6) 52,4%) não precisa descer. Atendido também o pedido de guardar remoção instantânea para defender a coroa: **13 das 22 peças respondem no turno do oponente.** |
| **4 · ramp** | Curva de interação concentrada em 1–2 manas (13 de 18) **de propósito**, para caber junto com o `{R}{R}{W}{W}` do escape a partir do T5. O pico de mana da fase é [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (`{3}{W}{W}`) — **dois brancos**, que reforça o gargalo que a Fase 2 apontou: o branco duplo, não o vermelho. [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) (`{1}{W}{W}`) idem. **Priorizem fixação de branco.** |
| **6 · manabase** | **Devolvo 2,5 slots** — a conta está na §1. [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) (caixa, R$ 0,89) é o 5º efeito de massa **no slot de terreno**; recomendo forte. Os terrenos com cycling que vocês vão avaliar também são combustível para o **delirium da [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat)** (o tipo Land no cemitério), o que sobe o valor deles acima do que a Fase 2 estimou. Se precisarem cortar da minha fase, a fila está na §8 e **nenhum corte meu desce de 20 remoções**. |
| **7 · wincon** | Sobreposições que já contam como alcance e não precisam de slot novo: [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat), [**Fateful End**](https://www.ligamagic.com.br/?view=cards/card&card=Fateful+End) (reserva) e [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) são **fonte vermelha em qualquer alvo** → contam no cálculo do [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) e do `Solphim`. [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) com Overload + [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) é 12 em cada criatura alheia. **Buraco que vocês herdam: o céu** (§5.2) — nenhum bloqueador com voo ou alcance; `Gisela` cobriria isso além do papel de finisher. |
| **Orquestrador** | (a) **15 das 22 estão `a cotar`**; a exposição real é [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) e [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix). Cotar **antes** dos cortes da Fase 6. (b) **Buraco declarado e sem conserto em RW: contramágica = 0** (§5.1) — combo que resolve na pilha ganha, e isso é imposto da cor escolhida na Fase 1, não erro de seleção. (c) **Repito o risco de percepção da Fase 2**: 18 das minhas 22 peças são mágicas e encantamentos, e o usuário reprovou a v1 dizendo *"muita mágica instantânea e encantamento na mão e poucas criaturas"*. Esta fase é a que mais agrava o número. **Tem que ser pergunta explícita no `report.md`.** (d) [**Swords to Plowshares**](https://www.ligamagic.com.br/?view=cards/card&card=Swords+to+Plowshares) (R$ 13,99) e [**Winds of Abandon**](https://www.ligamagic.com.br/?view=cards/card&card=Winds+of+Abandon) (R$ 19,41) ficaram fora **por preço**; se o total fechar com folga, são os dois melhores upgrades disponíveis. |

---

## Revisão 2026-09-24 — pacote de reanimação

> **Devolução do orquestrador, modo `improve`.** Lista de partida: v2 do `report.md` §6 com as 15
> trocas do `02-theme.md` §12 aplicadas. Pedido: confirmar ou contestar os 4 cortes condicionados a
> esta fase ( [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning), [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid), [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear), [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray)), dizer
> quanto do Helix condicional conta como remoção, avaliar a cobertura contra ódio de cemitério e
> revisar a proteção.
>
> **Regra 6.** Todo oracle desta seção foi puxado com `bin/mtgdb oracle` hoje, e os rulings do
> Phlage e do [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial) com `bin/mtgdb rulings`. A CR 903.9a é citada a partir do §12.1
> do `02` (o MCP `get_rule` não está disponível nesta sessão); o texto confere com a regra.
> **Regra 2.** Preços só de `bin/mtgdb prices`, LigaMagic (menor), cotações de 2026-09-16 a
> 2026-09-24. **Regra 5.** Nenhuma das cartas que movo aparece no `decisions.md`, que só registra
> movimentos de zona de comando. [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) foi dispensada **nesta fase**, no §7 acima,
> e declaro abaixo o que mudou. **EDHREC** não foi chamado de novo. As seções `Instants`/`Sorceries`/`Utility Artifacts`
> não foram transcritas no `02` §12.3, então não há candidata de interação vinda do meta nesta revisão.

### R.1 Quanto do Helix conta como remoção

O Helix do Phlage é **3 de dano em qualquer alvo** (5 com [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), 6 com
[**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer), porque o Phlage é Elder Giant). Quatro propriedades definem quanto ele vale
como remoção:

| Propriedade | Consequência |
|---|---|
| **Correlação total.** As 10 fontes dependem do mesmo cartão no mesmo cemitério | Um [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) ou um [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage) desliga as 10 **de uma vez**. Não são 10 sorteios independentes, e a hipergeométrica não se aplica a elas |
| **Velocidade de feitiço.** As 5 reanimações de uma vez são feitiços. Dial e Desdemona conjuram criatura. Sun Titan e Warsinger dependem de combate | Só [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) e [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) disparam o Helix fora do seu turno principal, e ambos dependem de o Phlage morrer |
| **Só criatura e planeswalker.** Dano não toca artefato nem encantamento | Não reduz em nada a necessidade de resposta a artefato/encantamento, que é justamente a que protege o motor |
| **Disputa o alvo com o fecho.** O mesmo Helix é o relógio da R2′ | Cada Helix numa criatura é um Helix a menos na cara |

**Como contei:**
- **As 5 reanimações de uma vez** ( [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand), [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant), [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness),
  [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission), [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation)) valem **0,5 cada = 2,5 slots**. Pelo lado bom, são
  baratas e cada uma é uma remoção de 3 de dano completa. Pelo lado ruim, a correlação e a
  velocidade de feitiço cortam o valor pela metade.
- **Os 5 motores** (Teshar, Sun Titan, Warsinger, Dial, Desdemona) valem **0**. Eles são ameaça e
  precisam sobreviver na mesa para repetir o Helix. Contá-los como remoção seria contar duas vezes
  a mesma carta.
- [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) e [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) também valem **0** aqui, porque já estão contados
  na proteção (R.5).

**Total aceito: +2,5 slots condicionais.** Remoção efetiva = formal + 2,5, e **só contra
criatura/planeswalker de resistência ≤ 3** (≤ 6 com o Bearer).

**Consequência para a composição.** O Helix faz em velocidade de feitiço o trabalho de *dano em
criatura pequena*. A remoção formal que sobra deve, portanto, puxar para os três eixos que o
Helix não cobre: **instantâneo**, **artefato/encantamento** e **criatura grande**. É esse o
critério dos vereditos abaixo.

### R.2 Os 4 cortes condicionados — veredito

| Carta | Veredito | Ficha F1–F7 (todas as funções) → quem cobre |
|---|---|---|
| [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) (caixa, R$ 0,05) | **CONFIRMO** | **F1** 2 de dano em criatura, instantâneo, e **exila** se ela morrer → dano pequeno: 10 fontes de Helix + [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade), [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless), [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre). Exílio de criatura recorrente: [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) (3 + exílio + tira indestrutível) e [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory). **F2** sem corpo. **F3** instantâneo → combustível e tipo do delirium → qualquer instantâneo. Fonte vermelha → [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) (4) → as outras queimas vermelhas. **F4/F5** nada. **F6** T1 → [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) (`{R}`, instantâneo). **F7** nada. **Descoberto:** a 2ª resposta instantânea de 1 mana. **Custo aceito**: é o papel mais redundante do deck depois que o Helix entrou. Continua sobressalente |
| [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) (Tori físico, R$ 0,10) | **CONFIRMO** | **F1** destrói criatura **virada**. É feitiço, mas com spell mastery (2+ instantâneos/feitiços no cemitério) ganha flash, e as 5 reanimações de feitiço deixam isso quase sempre ligado → criatura grande: [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) (1 por oponente), [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer), [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) (6 com delirium), [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (7 por upkeep), [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) e [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong); o Helix com o Bearer chega a 6. **F2** sem corpo. **F3** feitiço → combustível e delirium: as 5 reanimações. **F4/F5** nada. **F6** T2. **F7** só pega criatura **virada**, ou seja, depois do ataque, quando o gatilho de ataque já resolveu. **Descoberto:** matar em instantâneo, por 2 manas, uma criatura de resistência 7+. **Custo aceito**: sobram 4 respostas irrestritas, 2 delas instantâneas |
| [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) (caixa, R$ 0,44) | **CONFIRMO** | **F1** Spree: `{1}{W}` artefato, `{1}{W}` encantamento, `{3}{W}` os dois; 3º modo: +1/+1 em cada criatura de um jogador → o 2-por-1 volta com [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) (R.3). **F2** sem corpo. **F3** feitiço → combustível. **F5** o 3º modo agora teria alvo (17 criaturas, e o +1 ajuda o [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) a causar 3), mas continua marginal. **F6** T3 para um modo, T4 para os dois. **F7** feitiço. **Por que o 2-por-1 da caixa perde para o comprado:** os dois modos custam `{3}{W}` contra `{1}{R}{W}`, e um modo sozinho custa `{1}{W}` contra `{W}` do **Tear**. No T4, [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) fundido + [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) cabem em 4 manas; [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) nos dois modos + [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) pedem 5. Continua sobressalente |
| [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) (R$ 2,91) | **CONTESTO: volta, no lugar do Disenchant** | Ver R.3. O ódio de cemitério passou a desligar o motor inteiro, e as duas peças que mais o desligam são **uma de cada tipo** (Rest in Peace é encantamento, Grafdigger's Cage é artefato). O 2-por-1 em instantâneo tem agora o maior valor desde a v1, e não o menor |

### R.3 Troca proposta 1 — entra [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear), sai [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant)

**Slot:** a resposta instantânea a artefato/encantamento de 2 manas.

**Entra: Wear // Tear** — `{1}{R}` // `{W}`, instantâneo, fusível. R$ 2,91 (2026-09-23).
- **F1** Wear: destrói artefato. Tear: destrói encantamento. Fundido, a partir da mão, faz os dois por `{1}{R}{W}`.
- **F3** instantâneo → combustível e delirium. **F6** Tear na curva do T1, fundido no T3.
- **F7** o fusível só funciona a partir da mão, e no cemitério ela é um instantâneo só.
- **Sinergias (regra 3):**
  1. **O cenário que desliga o motor são duas peças**: [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) (encantamento) e
     [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage) (artefato; oracle: *creature cards in graveyards can't enter the battlefield;
     players can't cast spells from graveyards*). Juntas, desligam reanimação, escape próprio,
     Dial, Desdemona, flashback, [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) e [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth). [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) é a única carta do deck que
     responde às duas com um card só, em instantâneo.
  2. **Tear por `{W}` deixa mana para a reanimação no mesmo turno.** A sequência do T4 é Tear no
     [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) + [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) (2 manas) + 2 manas livres. Com [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant), sobra 1.
  3. **Instantâneo**, e é isso que o Helix não cobre (R.1).

**Sai: Disenchant** — Tori físico, R$ 0,25.

| Eixo | Função | Quem cobre depois do corte |
|---|---|---|
| **F1** | destrói artefato **ou** encantamento, instantâneo, `{1}{W}` | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear): cada metade custa igual ou menos e o fusível faz os dois |
| **F2** | sem corpo | — |
| **F3** | instantâneo → combustível e delirium | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) (instantâneo) |
| **F4 / F5** | nada | — |
| **F6** | T2 | Tear no T1–T2 |
| **F7** | nenhum | — |

**Descoberto:** destruir artefato **sem** fonte vermelha. É um caso marginal, porque o deck tem 22
vermelhas e 21 brancas nas mágicas. **Corte limpo.**

**Regra 7.** A caixa tem [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid), que faz o 2-por-1, e a dispensa está no R.2: custa
1 mana a mais por modo e é feitiço. A `lista.txt` tem [**Invoke the Divine**](https://www.ligamagic.com.br/?view=cards/card&card=Invoke+the+Divine) e **Expose to
Daylight**: são 3 manas por um alvo só e sem o 2-por-1. Ficam como reservas sem custo.

**Delta: +R$ 2,66.**

### R.4 Ódio de cemitério — cobertura

**O que desliga o quê.** Com [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace), [**Leyline of the Void**](https://www.ligamagic.com.br/?view=cards/card&card=Leyline+of+the+Void), [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage),
[**Relic of Progenitus**](https://www.ligamagic.com.br/?view=cards/card&card=Relic+of+Progenitus), [**Tormod's Crypt**](https://www.ligamagic.com.br/?view=cards/card&card=Tormod%27s+Crypt) ou [**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog), o motor inteiro para (10 fontes de
Helix, escape próprio, Dial, Desdemona, Renewal, Broodmoth, flashback). O plano B é a CR 903.9a:
o comandante que acabou de chegar ao cemitério ou ao exílio *pode* ir para a zona de comando.
A escolha é feita toda vez. **Quando o Phlage se sacrifica normalmente, deixe-o no cemitério**,
e mande à zona de comando **só** quando ele for exilado ou o cemitério estiver travado.

| Peça | Tipo | Efeito no motor | Respostas no deck | Sem resposta: mitigação |
|---|---|---|---|---|
| [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) / [**Leyline of the Void**](https://www.ligamagic.com.br/?view=cards/card&card=Leyline+of+the+Void) | encantamento, estático | tudo que iria ao cemitério vai ao exílio; o Phlage sacrificado é exilado | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) (Tear `{W}`) · [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) (pré-pago, instantâneo) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) · [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) · [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) (2 contadores para o RiP) · [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) (terreno) = **8** | 903.9a: o Phlage exilado vai à zona de comando e é reconjurado com imposto. Cada conjuração é **um Helix taxado** (5, depois 7…), e ele volta ao exílio e à zona de comando. O motor vira um Helix caro por turno até a resposta chegar |
| [**Grafdigger's Cage**](https://www.ligamagic.com.br/?view=cards/card&card=Grafdigger%27s+Cage) | artefato, estático | nenhuma criatura entra do cemitério e nada se conjura de lá | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) (Wear) · [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) · [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) · [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) · [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) · [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) · [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) · [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) (1) · [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) (entra com 1 contador; mata o nosso [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) junto) = **9** | a Cage não impede conjurar da zona de comando. Mande o Phlage sacrificado para lá (903.9a) e reconjure com imposto |
| [**Relic of Progenitus**](https://www.ligamagic.com.br/?view=cards/card&card=Relic+of+Progenitus) / [**Tormod's Crypt**](https://www.ligamagic.com.br/?view=cards/card&card=Tormod%27s+Crypt) / [**Soul-Guide Lantern**](https://www.ligamagic.com.br/?view=cards/card&card=Soul-Guide+Lantern) | artefato, ativação instantânea | exila o cemitério em resposta a qualquer coisa | **nenhuma efetiva**: o dono ativa em resposta à remoção | 903.9a: o Phlage exilado vai à zona de comando. **Uma** conjuração com imposto (5 manas) faz o Helix, ele se sacrifica, cai no cemitério limpo e **o motor religa**. Perde-se o combustível do escape, que o plano de reanimação não exige |
| [**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog) | terreno, gatilho de entrada | idem | **nenhuma** | idem. O [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) dispara uma vez (ruling: uma vez para cartas que saem juntas) |
| **Remoção que exila no Phlage escapado** ([**Swords to Plowshares**](https://www.ligamagic.com.br/?view=cards/card&card=Swords+to+Plowshares), [**Path to Exile**](https://www.ligamagic.com.br/?view=cards/card&card=Path+to+Exile), [**Farewell**](https://www.ligamagic.com.br/?view=cards/card&card=Farewell)) | — | tira o 6/6 | [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) (pontual, antes da resolução) | 903.9a: zona de comando → reconjura (imposto) → Helix → sacrifica → cemitério → a reanimação volta. **O exílio do 6/6 custa ao oponente uma remoção e custa a você 2 de imposto.** [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) e efeitos que embaralham caem na 903.9b (mão/grimório → zona de comando) |
| Criatura que come cemitério ([**Scavenging Ooze**](https://www.ligamagic.com.br/?view=cards/card&card=Scavenging+Ooze), [**Dauthi Voidwalker**](https://www.ligamagic.com.br/?view=cards/card&card=Dauthi+Voidwalker)) | criatura | exila cartas pontualmente | 13 remoções formais de criatura + o Helix: a Ooze morre para 3 de dano antes de crescer, e a Voidwalker (3/2) também | 903.9a para o Phlage que ela exilar |

**Leitura.**
- **Ódio estático (RiP, Leyline, Cage):** **8 a 9 respostas**, sendo 7 mágicas mais **Ratchet
  Bomb** e **Blast Zone**. P(≥1 até o T6, 12 cartas vistas) = **70,4%** com 9 e **60,7%** só com
  as 7 mágicas, antes dos loots. Enquanto a resposta não chega, o Phlage continua disparando o
  Helix da zona de comando, com imposto. **Coberto.**
- **Ódio instantâneo (Relic, Crypt, Bog):** **sem resposta em RW, declarado.** A mitigação é
  estrutural e sai barata: uma conjuração com imposto religa o motor, porque ele só precisa **do
  Phlage** no cemitério, não de cemitério cheio. O que se perde de verdade é o combustível do
  escape próprio (plano B) e do Dial.
- **Nenhuma carta dedicada entra por causa disso.** Uma 10ª resposta estática subiria o T6 de
  70,4% para ~75% e custaria um slot de remoção de criatura. E a 903.9a já transforma o pior caso
  em atraso, não em derrota.

### R.5 Proteção — revisão

**O que o deck precisa proteger mudou.** Na v2, eram os pingadores e o 6/6 escapado. Agora são:
(a) os **motores-criatura** (Teshar, Sun Titan, Warsinger, Desdemona, Bearer, Broodmoth, Torbran,
Gisela); (b) os **motores não-criatura** ([**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome), [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger), **Confession
Dial**, **Court of Ire**); (c) **o cemitério**, que é o R.4 e se protege com remoção de
artefato/encantamento, não com carta de proteção. O Phlage **reanimado** não precisa de proteção:
ele morre de propósito.

| Peça | Protege | Não protege |
|---|---|---|
| [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) | **o único que mantém o 6/6 escapado** contra remoção mirada; qualquer motor-criatura; scry 1 | wipe, exílio em massa |
| [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) | 1 motor-criatura contra qualquer morte (o "may" deixa guardar para Torbran, Gisela, Titan, Teshar) | o Phlage escapado (volta sem ter escapado e se sacrifica, **correção da Fase 2 confirmada**: vira +1 Helix); exílio; ele é bloqueado pela Cage |
| [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) | toda criatura **sem voar** num wipe de destruição, inclusive quando ela morre junto (ruling 2020-04-17) | as **voadoras**: [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle), a própria Broodmoth, [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight), [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise); o Phlage escapado (mesma correção); não-criaturas; exílio |

**O buraco:** num **Wrath** do oponente, a Broodmoth salva o time, mas **não salva a si mesma
nem o Teshar**, que são as duas peças que mais repetem o Helix. E, se ela já morreu antes, nada
salva o time.

#### Troca proposta 2 — entra [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death), sai [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash)

**Slot:** remoção de criatura por dano, em velocidade de feitiço, que é o trabalho do Helix.

**Entra: Duty Beyond Death** — `{1}{W}`, instantâneo, Tori físico (`lista.txt`), R$ 0,43
(2026-09-16).
- **F1** custo adicional: sacrificar uma criatura. Suas criaturas ganham indestrutível até o fim
  do turno, e cada uma recebe um contador +1/+1.
- **F3** instantâneo → combustível e delirium. **F5** indestrutível para o time e +1/+1
  permanente: o [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) vira 4/4 e passa a causar 3 com folga.
- **F6** T2, e se guarda na mão.
- **F7**: (1) exige uma criatura para sacrificar; há 17. (2) Não protege não-criaturas. (3) Não
  segura exílio em massa nem -X/-X. (4) Os contadores sobem o poder do time contra o seu próprio
  [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong).
- **Sinergias (regra 3):**
  1. **O custo virou recurso.** A criatura sacrificada de MV ≤ 3 ( [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert), **Zookeeper
     Mechan**, **Millikin**, **Ornithopter of Paradise**, **Venerable Warsinger**) é alvo de
     [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand), do Teshar e do Sun Titan. Com a Broodmoth em campo, a que não voa **volta na
     hora**, e a **Tocasia's Welcome** compra quando ela reentra.
  2. **Com o Phlage reanimado, o custo sai de graça.** Em resposta ao gatilho de sacrifício dele,
     conjure [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) sacrificando o próprio Phlage, que ia para o cemitério de
     qualquer jeito. Em seguida, [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) vira unilateral e ainda devolve vida por criatura
     alheia. É a linha proativa de wipe que o deck não tinha sem a Broodmoth.
  3. **Cobre o buraco da Broodmoth:** protege as voadoras (Teshar, a própria Broodmoth, Gisela) e o
     **Phlage escapado**, que fica indestrutível desde que não seja o sacrificado.
- **Regra 5 / o que mudou:** foi dispensada no §7 desta fase (v2) com o argumento *"6 dos corpos
  são o wincon, e sacrificar um é anti-sinergia direta"*. Os 4 pingadores saíram no `02` §12.6, e a
  reanimação transformou o sacrifício em recurso (sinergias 1 e 2). O motivo original não vale
  mais.

**Sai: Flame Slash** — caixa (maybeboard), R$ 2,48 (2026-08-12).

| Eixo | Função | Quem cobre depois do corte |
|---|---|---|
| **F1** | 4 de dano em criatura, feitiço, `{R}` | Helix (3; 5 com Torbran; 6 com Bearer) em 10 fontes · [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) (4, `{1}{R}`, só criatura alheia) · [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) (2, ou 6 com delirium, instantâneo `{R}`) · [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) / [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) / [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) / [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) / [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) (3, instantâneo) |
| **F2** | sem corpo | — |
| **F3** | feitiço → combustível e tipo do delirium. Fonte vermelha → Torbran (6 em criatura) | as 5 reanimações são feitiços; [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars), [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate), [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) e [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) também. Fonte vermelha: as outras 7 queimas |
| **F4 / F5** | nada | — |
| **F6** | T1, e cabe junto de outra jogada | [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) (T1, instantâneo). No T3, o Phlage da zona de comando já é um Helix |
| **F7** | feitiço de dano em criatura: **disputa exatamente o papel do Helix** | — |

**Descoberto:** matar uma criatura de resistência 4 por `{R}` no T1–T2, antes de o Phlage chegar
ao cemitério. **Custo aceito.** O T3 do Phlage já cobre a janela, e a [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) mata
resistência 2 (ou 6 com delirium) no mesmo custo. Das 16 remoções formais, ela é **a que mais se
sobrepõe ao Helix**: feitiço, só criatura, só dano. Pelo critério do R.1, é a primeira a sair.
Continua sobressalente.

**Por que Flame Slash e não Lightning Strike:** [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) é instantâneo e acerta
jogador. Ela também tem papel de alcance (categoria `wincon`), então cortá-la seria **corte
condicionado** à Fase 7. **Flame Slash** é só `remoção`, portanto da minha especialidade, e sai limpa.

**Delta: −R$ 2,05.** Para a Fase 6, a cor das mágicas muda: vermelho −1 e branco +1 (`{R}` sai,
`{1}{W}` entra).

#### [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) — não recomendo, fica como reserva

R$ 33,17 (2026-09-23) consome **88% da folga de R$ 37,88**, e as Fases 3, 4 e 6 estão disputando
essa mesma folga em paralelo. O que ele faz a mais que Duty + Broodmoth: protege **não-criaturas**
(Tocasia, Ark, Dial, Court, rochas) contra wipe de destruição, não pede sacrifício e tem os modos
4 na cara e golpe duplo (Phlage escapado com Bearer = 24 de combate, uma questão da Fase 7). O que
ele **não** faz, igual aos outros dois: segurar exílio e ódio de cemitério, que são a ameaça
número 1 (R.4). Se, depois da consolidação, sobrarem ≥ R$ 33 de folga, ele entra **no lugar do
Duty Beyond Death** (mesmo slot, +R$ 32,74).

### R.6 Contagem final (com as duas trocas)

| Categoria | v2 | `02` §12 | **Esta revisão** | Meta |
|---|---|---|---|---|
| **Remoção formal** | 20 | 16 | **15** | ~10 (pipeline) |
| + Helix condicional (R.1) | — | "10 fontes" | **+2,5** (só criatura/PW ≤ 3) | — |
| **Remoção efetiva** | 20 | — | **17,5** | — |
| Respostas a artefato/encantamento | 9 (+2-por-1 ×2) | 7 (sem 2-por-1) | **7 mágicas (2-por-1 ×1) + Ratchet Bomb + Blast Zone = 9** | ≥ 2 |
| Instantâneas entre as formais | 13/22 | — | **9/15 = 60%** | — |
| **Wipes** | 3 + terreno | 3 + terreno | **3** ( [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong), [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate), [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb)) **+ Blast Zone** · efeitos de massa extras: overload do [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) e [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | 2–4 |
| **Proteção** | 1 | 1 + 2 parciais | **2** ( [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) pontual, [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) em massa) **+ 2 parciais** (Renewal, Broodmoth) | 2 |

**As 15 formais:** [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat), [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless), [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade),
[**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear), [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing)
(instantâneas ou ativáveis em instantâneo, 9) · [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars), [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket), **Reduce to
Memory**, **Celebrate the Mountain-king**, **Court of Ire**, **Palace Jailer** (6). Remoção
extraoficial, não contada: [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre), [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab), [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) e [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate).

**Por que 15 formais bastam.** Com 15, P(≥1 remoção formal na mão de 7) = 69,6%, contra 80,5%
com 20. A diferença é pequena por três motivos: (1) o Helix cobre o que saiu (dano pequeno em
feitiço); (2) o Phlage da zona de comando é um Helix garantido no T3, sem carta; (3) o piso de 20
foi medido para o eixo de controle puro da v2, em que a remoção **era** o plano. Agora o plano é o
motor, e a remoção o protege. Eu não desço abaixo de 15.

**Cobertura de ameaças:**
- **Criaturas:** 13 formais + 10 fontes de Helix + 3 wipes.
- **Artefato:** 8 (7 mágicas + Ratchet Bomb) + Blast Zone.
- **Encantamento:** 7 (6 mágicas + Ratchet Bomb) + Blast Zone.
- **Flexível** ("destroy/exile target permanent"): [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift), **Reduce to
  Memory**, **Celebrate the Mountain-king**.
- **Combos e mágicas na pilha:** 0 (RW não tem contramágica; buraco declarado no §5.1, sem mudança).

### R.7 Custo

| Troca | Entra | Sai | Δ |
|---|---|---|---|
| 1 | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) R$ 2,91 (compra) | [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) R$ 0,25 (Tori físico) | **+2,66** |
| 2 | [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) R$ 0,43 (Tori físico) | [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) R$ 2,48 (caixa) | **−2,05** |
| **Total** | | | **+R$ 0,61** |

LigaMagic (menor), cotações de 2026-08-12 a 2026-09-23. **Compra nova:** só [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear), que
na régua do `02` §12.9 já estava contado como corte, então ele volta a pesar. Contra o total do
§12 (Dial R$ 1,42 e Desdemona R$ 1,45, já cotados): **+R$ 0,61** sobre a folga de R$ 37,88.

### R.8 Reservas (ordem)

| # | Entra | Sai | Quando | R$ |
|---|---|---|---|---|
| 1 | [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) | [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) | se sobrarem ≥ R$ 33 de folga depois da consolidação | 33,17 |
| 2 | [**Selfless Spirit**](https://www.ligamagic.com.br/?view=cards/card&card=Selfless+Spirit) (2/1 voadora, sacrifica-se: indestrutível para o time) | [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) | se o orquestrador quiser o 18º corpo: é criatura de MV 2, o Teshar e a [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) a trazem de volta e a Tocasia compra quando ela entra. Voa, então a Broodmoth não a devolve | `a cotar` |
| 3 | [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) (caixa) | [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) | se a compra for vetada: mantém o 2-por-1, mas custa 1 mana a mais por modo e é feitiço | 0,44 |
| 4 | [**Invoke the Divine**](https://www.ligamagic.com.br/?view=cards/card&card=Invoke+the+Divine) (lista) | qualquer resposta a artefato/encantamento | 10ª resposta estática, se a mesa local jogar muito ódio de cemitério | `a cotar` |
| 5 | [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) (caixa) | [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) | se o goldfishing mostrar que a Broodmoth sozinha segura os wipes | 2,48 |

### R.9 Sinais

| Fase | Sinal |
|---|---|
| **2 · tema** | Aceito os 3 cortes (Magma Spray, Swift Reckoning, Requisition Raid). Contesto 1: [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) volta pelo [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant). O argumento *"sobreposição de artefato/encantamento"* inverteu com o eixo novo: é a resposta que protege o motor. O [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) usa as criaturas de MV ≤ 3 que o §12.6 manteve como fodder reanimável |
| **6 · manabase** | Mágicas: vermelho −1 ([**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash)), branco 0 (sai [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) `{1}{W}`, entra [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) `{1}{W}`), RW +1 ([**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear)). O Tear pede `{W}` no T1–T2. [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) conta como resposta à Cage (1 contador ao entrar) e mata o [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) junto |
| **7 · wincons** | O Helix tem um dono só: cada disparo vai para criatura **ou** para a cara. Para a simulação: ~2,5 Helix por jogo desviados para remoção. [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) fica, e o papel de alcance segue com vocês. O modo de golpe duplo do [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) (24 com Bearer) fica registrado, mas ele não entra por preço |
| **Orquestrador** | (a) Regra operacional para o `report.md`: **o Phlage sacrificado fica no cemitério**. Ele só vai à zona de comando (903.9a) se for exilado ou se o cemitério estiver travado. (b) Ódio de cemitério instantâneo (Relic, Crypt, Bog) **não tem resposta em RW**. A mitigação é uma reconjuração com imposto, e isso deve ir para o §11 (pendências) como risco declarado. (c) [**Selfless Spirit**](https://www.ligamagic.com.br/?view=cards/card&card=Selfless+Spirit) está `a cotar` |
