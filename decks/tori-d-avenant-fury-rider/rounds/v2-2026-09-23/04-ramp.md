# Aceleração (Ramp) — Phlage, Titan of Fire's Fury (v2 · eixo de controle de atrito)

**Ramp padrão proposto: 9 fontes** (a meta do pipeline é 10–11 — proponho **menos**, com a conta na §2)
· **Explosivo: 2–3**, redefinido (§6), e a **custo zero de slot**.

> **Natureza desta fase.** Não existe `deck.md`. O deck v1 (Otharri) foi reprovado e descartado.
> Isto é um `build` na prática: não há ramp "já no deck" para contar — há um pool sem custo
> (caixa + `lista.txt`) e um pacote a montar do zero.
>
> **Regra 5 — conferido.** `decisions.md` registra **três** movimentos, todos de zona de comando
> (Otharri entra/sai, Tori sai, Phlage entra). **Nenhuma carta do 99 tem histórico de corte.**
> Nenhuma proposta abaixo é reposição de corte anterior; nenhuma carta está barrada.
>
> **Regra 7 — a caixa veio antes do Scryfall.** `bin/mtgdb collection -list` (116 cartas) e as 69
> não-básicas de `lista.txt` foram varridas primeiro (§3). **6 das 9 titulares saem do pool sem
> custo.** As dispensas estão justificadas por escrito, com ficha, na §3.2.
>
> **Regra 2 — preço.** Nenhum número do Scryfall, em lugar nenhum. Cotações vêm de
> `bin/mtgdb prices` (LigaMagic, menor). O que não tem cotação sai como **`a cotar`** — nunca com
> valor estimado.
>
> **Regra 6 — oracle.** Todo texto citado foi puxado com `bin/mtgdb oracle` nesta sessão.
>
> **Travas mecânicas do escape respeitadas** (`decisions.md`, 2026-09-23): 6 cartas no cemitério
> para o primeiro escape · só o **escape** evita o sacrifício (nada de reanimar/blinkar) · o Helix
> acontece na conjura da zona de comando. Nenhuma peça abaixo interage com reanimação ou blink.

---

## 1. Por que a meta do pipeline não serve a este deck

A meta de **10–11 ramps padrão + 2–3 explosivos** existe para um deck que precisa **chegar a um
alvo de mana alto**: comandante de CMC 5–7, bombas no topo da curva, um enxame para sustentar.
Era a conta certa para o Otharri (`{3}{R}{W}`, CMC 5). **Este deck não tem alvo alto nenhum:**

| | v1 (Otharri) | v2 (Phlage) |
|---|---|---|
| Comandante da zona de comando | **CMC 5** (`{3}{R}{W}`) | **CMC 3** (`{1}{R}{W}`) |
| Recasts posteriores | imposto 5 → 7 → 9 | **fixo em 4** (`{R}{R}{W}{W}`, escape — sem imposto) |
| Topo real da curva | 5 + imposto crescente | **4** (Torbran 4, Solphim 4, escape 4) |
| Perfil das outras 60 cartas | anthems, corpos, pump | **20–22 remoções + 14 combustível, curva 1–3** |
| O que o mana precisa comprar | **chegar ao comandante** | **dobrar ações por turno** e **pagar 4 pips coloridos** |

Os *outliers* de curva do pool da Fase 2 são três: [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (5), [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) (5),
[**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) (6) e [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) (7) — e os dois últimos estão declarados
como **finishers opcionais**, não como alvo de rampa. [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) custa `{R}` na prática.

**Consequência: a pergunta "quantos turnos o ramp adianta o comandante" é a pergunta errada aqui.**
As perguntas certas são as três que a §2 mede.

---

## 2. As três medições

> **Método.** Monte Carlo próprio, 20–30 mil mãos por configuração, `on the play`, mulligan
> londrino simplificado (mantém 2–5 terrenos, devolve terreno excedente ao fundo), 99 cartas +
> comandante. Política de gasto: **Phlage da zona de comando tem prioridade no T3**, depois rochas
> da mais barata para a mais cara, depois motores, depois loots, depois **no máximo 1 mágica
> reativa por turno** (deck de controle segura remoção; conjurar tudo seria trapacear a favor do
> combustível). Base de terrenos assumida: 36 terrenos, com o número de fontes duplas variando
> onde indicado. Rochas que entram viradas são modeladas como tal.
>
> **Divergência declarada com a Fase 2.** Minha simulação conjunta (§2.3) é **mais otimista** que a
> da §3.1 do `02-theme.md`: ela conjura mágicas com mais agressividade e conta o refil dos loots,
> e a Fase 2 disse explicitamente que errava para baixo de propósito. **Não uso meus números
> absolutos para revisar os dela** — uso as **diferenças entre configurações de ramp**, que é o que
> esta fase precisa responder, e que é robusto a essa escolha de modelo.

### 2.1 Medição 1 — P(escapar no T5 e no T6) em função do ramp

**Disponibilidade de mana pura** — P(conseguir pagar `{R}{R}{W}{W}`), 36 terrenos / 6 duais,
rochas **coloridas**:

| Fontes de ramp | T4 | **T5** | **T6** | T7 | T6 **com ≥2 de mana sobrando** |
|---|---|---|---|---|---|
| 0 | 42,4% | 55,8% | 64,8% | 72,0% | 25,4% |
| 2 (só os intocáveis) | 46,7% | 59,8% | 68,5% | 75,4% | 33,6% |
| 4 | 52,9% | 65,6% | 74,0% | 80,1% | 40,4% |
| 6 | 59,4% | 71,4% | 79,2% | 84,6% | 46,3% |
| 8 | 65,7% | 76,9% | 83,6% | 88,2% | 52,7% |
| **9 (proposto)** | **66,9%** | **78,4%** | **85,1%** | **89,7%** | **55,9%** |
| 11 (meta do pipeline) | 70,1% | 81,1% | 87,2% | 91,2% | 61,6% |

Isolado, o mana melhora monotonicamente e **parece** justificar 11. Mas o mana não é a única
restrição — o combustível é a outra, e **eles competem pelo mesmo slot**.

### 2.2 Medição 1b — o mesmo, com o combustível junto (a conta que decide)

Simulação conjunta: combustível **fixo** no alvo da Fase 2 (14 = 10 loot + 4 motores) e cada rocha
adicional saindo de um slot de queima/remoção. P(**primeiro escape** até o turno):

| Fontes de ramp | T5 | **T6** | T7 | T8 | T6 **com ≥2 de mana sobrando** |
|---|---|---|---|---|---|
| 0 | **49,5%** | 63,6% | 79,7% | 87,7% | 6,4% |
| 2 | 50,4% | 65,3% | 81,5% | 88,9% | 15,5% |
| 4 | 49,1% | 67,3% | 82,0% | 90,1% | 19,3% |
| 6 | 49,2% | 69,4% | 83,4% | 91,5% | 25,1% |
| 8 | 47,9% | 70,0% | 83,9% | 91,9% | 30,7% |
| **9 (proposto)** | 47,1% | **70,9%** | **84,7%** | **92,2%** | **32,6%** |
| 10 | 45,9% | 70,9% | 84,5% | 92,1% | 35,1% |
| 11 | 44,9% | 69,9% | 84,0% | 91,8% | 36,9% |
| 13 | 41,1% | 68,4% | 82,8% | 90,9% | 40,6% |

**Três leituras, e a terceira é a que fecha o número:**

1. **O T5 não melhora com ramp. Ele piora.** De 49,5% (zero ramp) para 44,9% (11 rochas). O motivo
   é mecânico e não tem volta: no T5 o gargalo do escape **não é mana, é cemitério** — são precisas
   5 cartas além do Phlage, e cada rocha que entra no deck é uma mágica que **não** foi conjurada e
   **não** caiu no cemitério. Ramp compra mana; o escape do T5 precisa de *lixo*.
2. **A curva do T6/T7 satura entre 8 e 10.** De 8 para 9 rochas: **+0,9 pp**. De 9 para 10: **0,0 pp**.
   De 10 para 11: **−1,0 pp** — a 11ª rocha é **negativa**. A meta de 10–11 do pipeline, aplicada
   literalmente aqui, custa probabilidade de escape em vez de comprar.
3. **O que cresce sem saturar é a coluna da direita** — escapar **e ainda ter 2 de mana livre no
   mesmo turno**: 6,4% → 32,6% → 36,9%. **Fator de 5× entre zero ramp e 9 rochas.** É esse o eixo
   que o ramp compra neste deck, e é o assunto da Medição 3.

**Nota de correção à Fase 2 (§4.2).** O `02-theme.md` afirma: *"O T3 é garantido. O Phlage sai da
zona de comando **em toda partida**, independente de compra."* Medido: **75,4%** exatamente no T3
sem nenhum ramp (36/6), e **85,0%** com o pacote de 9 (**87,5%** com 10 duais, **89,8%** com 38
terrenos e 12 duais). Até o T4: 84,2% sem ramp, **92,1%** com o pacote. A afirmação é falsa porque
o gargalo do T3 não é compra — é **land drop + os pips `{R}` e `{W}` juntos**. A defesa do piso de
14 corpos continua de pé, mas com 85% no lugar de 100%, e **10 pontos percentuais dessa garantia
são comprados pelo ramp**. Esse é, sozinho, o argumento mais forte a favor de rochas coloridas
baratas neste deck.

### 2.3 Medição 2 — os pips: `{R}{R}{W}{W}` não aceita mana incolor

O escape é **quatro pips coloridos, dois de cada cor, e zero genérico**. Nenhuma fração dele é
pagável com `{C}`. Isolando (36 terrenos / 10 duais, 8 rochas coloridas de base):

| Nona peça adicionada | T4 | T5 | T6 | T6 **com ≥2 sobrando** | mana médio no T6 |
|---|---|---|---|---|---|
| — (só as 8 coloridas) | 72,2% | 82,3% | 87,9% | 52,0% | 5,48 |
| **[**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring)** (`{C}{C}`) | 72,5% | 82,3% | **88,2%** | **56,8%** | 5,72 |
| **1 rocha colorida mv 2** | **74,9%** | **84,3%** | **89,6%** | 55,5% | 5,61 |
| **[**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin)** (`{C}`) | 72,2% | 82,3% | 88,2% | 55,0% | 5,60 |

E o contraste grosso — **9 rochas incolores contra 9 coloridas**, 36/6:

| Pacote | T4 | T5 | T6 | mana médio no T6 |
|---|---|---|---|---|
| 9 **coloridas** | 66,9% | **78,4%** | **85,1%** | 5,72 |
| 9 **incolores** | 49,2% | **62,7%** | **71,7%** | **5,80** |

**O pacote incolor produz mais mana total (5,80 × 5,72) e mesmo assim escapa 15,7 pp menos vezes
no T5.** Não há como ser mais explícito: neste deck, **mana incolor não é a mesma mercadoria que
mana colorido**, e a régua de "quantidade de rochas" mente se não separar as duas.

Três consequências operacionais:

- **[**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) paga zero pip do escape.** Ele é intocável por decisão do briefing e continua
  correto no deck — mas por **outro** motivo: ele soma **+0,3 pp** no eixo velocidade e **+4,8 pp**
  no eixo *escapar e agir no mesmo turno*. É a melhor carta do pacote para a Medição 3 e uma das
  piores para a Medição 1.
- **[**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) é ramp de mentirinha e combustível de verdade.** `{T}, moer 1: adiciona {C}` — zero
  pip, +3,0 pp no duplo-gasto. Ele fica no deck porque a Fase 2 já o conta como **motor** (moe 1
  por turno, em velocidade de instantâneo); como rocha ele é a mais fraca do pacote. Slot
  compartilhado, não slot de ramp.
- **Toda rocha nova deve produzir `{R}` ou `{W}` à escolha.** Rocha de uma cor só ([**Iron Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Iron+Myr),
  [**Gold Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Gold+Myr), [**Zookeeper Mechan**](https://www.ligamagic.com.br/?view=cards/card&card=Zookeeper+Mechan), [**Fire Diamond**](https://www.ligamagic.com.br/?view=cards/card&card=Fire+Diamond), [**Marble Diamond**](https://www.ligamagic.com.br/?view=cards/card&card=Marble+Diamond)) resolve metade do problema e
  vai para reserva. Rocha incolor nova ([**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone), [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler), [**Seer's Lantern**](https://www.ligamagic.com.br/?view=cards/card&card=Seer%27s+Lantern),
  [**Magnifying Glass**](https://www.ligamagic.com.br/?view=cards/card&card=Magnifying+Glass), [**Worn Powerstone**](https://www.ligamagic.com.br/?view=cards/card&card=Worn+Powerstone)) **não entra** — ficha na §3.2.

### 2.4 Medição 3 — repetição: um mana a mais para **escapar e ainda responder**

A pergunta do orquestrador: *um mana que adianta o escape em um turno, ou um mana que permite
escapar e jogar remoção no mesmo turno?* A Medição 1b já respondeu qual dos dois o ramp
**consegue** comprar: **o segundo**. A Medição 2 diz **quanto** e a §6 diz **com o quê**.

Por que o segundo vale mais **neste** deck, e não é só uma consolação:

1. **O turno do escape é o turno mais vulnerável do deck.** Você gastou 4 manas e exilou 5 cartas —
   o cemitério **zerou**. Se o escape consome o turno inteiro, você passa com mana zero e sem
   cemitério para a resposta seguinte. Com 2 de sobra, o mesmo turno ainda cabe [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) +
   [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray), ou um [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) em resposta.
2. **O segundo escape leva 4–5 turnos** (tabela da §3.2 da Fase 2). Ou seja: escapar mais cedo e
   morrer no mesmo turno **não é uma troca melhor** do que escapar no T6 protegido. O tempo perdido
   não se recupera pela velocidade — recupera-se pela sobrevivência do que escapou.
3. **O deck não tem nada para fazer com mana no T3–T4 além de conjurar o Phlage e remoção barata.**
   O pico de demanda de mana é exatamente o turno do escape. É onde o ramp precisa estar.

Pelo número: o pacote de 9 leva a probabilidade de "escapar **com** 2 de mana livre no T6" de
**6,4% para 32,6%** — e a de "escapar" pura de 63,6% para 70,9%. **O ramp deste deck compra 26,2 pp
de ação dupla e 7,3 pp de velocidade.** É por isso que ele existe, e é por isso que ele para em 9.

---

## 3. Varredura da caixa e da `lista.txt` (regra 7) — antes do Scryfall

Fontes de mana encontradas: **15 na caixa** + **3 em `lista.txt`** ([**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring), [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet),
[**Command Tower**](https://www.ligamagic.com.br/?view=cards/card&card=Command+Tower)). Oracle de todas puxado nesta sessão.

### 3.1 Aproveitado — 6 das 9 titulares saem do pool sem custo de compra

| Carta | CMC | Produz | Por que entra | Origem |
|---|---|---|---|---|
| [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) | 1 | `{C}{C}` | **intocável** (briefing). 2 manas no T1–T2; paga a **segunda ação** do turno do escape (+4,8 pp na Medição 3). Zero pip. | caixa + lista |
| [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet) | 2 | `{R}` ou `{W}` | **intocável** (briefing). Rocha de 2 manas que cobre os dois pips — exatamente o perfil que a Medição 2 pede. | lista |
| [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) | 2 | qualquer cor (3 usos) | mv 2, qualquer cor, entra virada (custo medido: **~1 pp no T4, zero no T6**). **F7 declarado:** só 3 contadores — ela cobre a janela T3–T5 e depois morre. Em compensação, **vira custo adicional de [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers)** (sacrificar artefato) quando esgota. | caixa |
| [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert) | 2 | qualquer cor (paga 2 de vida) | rocha mv 2 de qualquer cor **e corpo 2/1** — conta no piso de 14–17 da Fase 2. O Phlage ganha 3 de vida por gatilho; 2 de vida por mana é troca que este deck aguenta. | caixa |
| [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) | 3 | `{R}` ou `{W}` | rocha de cor **que se sacrifica para comprar 1** — quando não precisa mais de mana, vira carta **e** vai para o cemitério (combustível). Tripla função. | caixa |
| [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) | 3 | `{R}` ou `{W}` | idem, **saca 2** por `{R/W}×4`. Mais caro de ativar, dobro de cartas. Mesma tripla função. | caixa |

### 3.2 Dispensas da caixa — justificadas por escrito, com ficha (regra 7 + regra 4)

| Carta | Origem | Ficha F1–F7 (resumida) | Por que **não** cobre a função |
|---|---|---|---|
| [**Manalith**](https://www.ligamagic.com.br/?view=cards/card&card=Manalith) | caixa | F1 `{T}`: 1 mana de qualquer cor · F3 Artifact · F6 **CMC 3** · F7 usa `{T}` | Faz **exatamente** o que [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) e [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) fazem, pelo mesmo CMC 3, e **não** se sacrifica para comprar. Dispensa por **dominância estrita** dentro da própria caixa — as duas outras estão disponíveis, de graça, e fazem mais. |
| [**Seer's Lantern**](https://www.ligamagic.com.br/?view=cards/card&card=Seer%27s+Lantern) | caixa | F1 `{T}`: `{C}`; `{2}`,`{T}`: scry 1 · F6 CMC 3 · F7 usa `{T}` | **Zero pip** (Medição 2: pacote incolor escapa 15,7 pp menos no T5) e CMC 3 (Sim: rochas mv 3 perdem 3,9 pp no T5 para mv 2). O scry por `{2}` compete com o mana do próprio escape. |
| [**Magnifying Glass**](https://www.ligamagic.com.br/?view=cards/card&card=Magnifying+Glass) | caixa | F1 `{T}`: `{C}`; `{4}`,`{T}`: investigate · F6 CMC 3 | Idem, com upside de **4 manas** num deck cujo turno crítico já custa 4. |
| [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler) | caixa | F1 `{T}`: `{C}` · **F2 corpo 0/1** · F3 Artifact Creature · F4 recebe anthems · F6 CMC 2 | **Corte condicionado à Fase 2, não dispensa minha.** Zero pip (não serve ao meu eixo), mas é **corpo** e o deck tem piso de 14 e 27,9% de mãos sem criatura. Um 0/1 não bloqueia nada — mas quem decide se o piso de corpos aceita um 0/1 é a Fase 2, não eu. **Se entrar, entra como corpo, e não deve ser contado nas 9 fontes de ramp.** Minha alternativa, no mesmo CMC 2: [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert) (2/1, qualquer cor) e [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise) (0/2 voador, qualquer cor) fazem as duas coisas. |
| [**Prophetic Prism**](https://www.ligamagic.com.br/?view=cards/card&card=Prophetic+Prism) | caixa | F1 ETB compra 1; `{1}`,`{T}`: 1 mana de qualquer cor · F3 Artifact · F6 CMC 2 | **Não é ramp — é filtro de mana líquido zero** (paga 1, devolve 1). Como fixação ela é real (converte o `{C}{C}` do [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) em pip colorido), mas num orçamento de 63 slots não-terreno uma carta que **não adiciona mana** perde para uma rocha colorida que adiciona **e** fixa. O cantrip do ETB é 1 carta, e a Fase 3 tem 12–13 fontes melhores. |
| [**Pilgrim's Eye**](https://www.ligamagic.com.br/?view=cards/card&card=Pilgrim%27s+Eye) | caixa | F1 voar; ETB busca terreno básico **para a mão** · F2 1/1 voador · F6 CMC 3 | **Não é ramp.** Terreno para a mão não acelera nada — ele só substitui o land drop que você já ia fazer. É fixação de cor a 3 manas por um corpo 1/1. [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise) custa 1 a menos, voa igual, tem 1 de resistência a mais e **produz mana de verdade**. |
| [**Moss Diamond**](https://www.ligamagic.com.br/?view=cards/card&card=Moss+Diamond) | caixa | F1 entra virada; `{T}`: `{G}` | **Fora da identidade de cor** (identidade G). Ilegal no deck. Registrado para não reaparecer. |
| [**Melded Moxite**](https://www.ligamagic.com.br/?view=cards/card&card=Melded+Moxite) | caixa | F1 ETB: **pode descartar 1 para comprar 2**; `{3}`, sac: ficha 2/2 virada · F6 CMC 2 (`{1}{R}`) | **Não é rocha de mana — não produz mana nenhum.** Mas o ETB é **loot de 1-por-2 num permanente barato**: 2 cartas para o cemitério (ela não, só o descarte) e refil de 2. **Devolvo à Fase 3 como candidata de draw/combustível**, fora do meu orçamento. Foi catalogada errado como "rocha" em rodadas anteriores. |
| [**Moxite Refinery**](https://www.ligamagic.com.br/?view=cards/card&card=Moxite+Refinery) | caixa | F1 `{2}`,`{T}`, remover X contadores: modal · F6 CMC 2 | **Não produz mana.** O nome engana. Fora do escopo. |
| [**Fellwar Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Fellwar+Stone) | **não está na caixa** | F1 `{T}`: mana de qualquer cor que um terreno de um **oponente** possa produzir · F6 CMC 2 · **R$ 14,50** (19/09) | Dispensa por **confiabilidade + preço**, não por função. Num pod aleatório ela cobre `{R}`/`{W}` com frequência, mas é **aposta na mesa**, e o pip que este deck precisa é `{W}{W}` — a cor menos comum de se encontrar. R$ 14,50 = **7,3% do teto** por uma rocha condicional, contra R$ 1,50 do [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet) que dá **os dois pips de uma vez**, sempre. |
| [**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone) | não está na caixa | F1 `{T}`: `{C}`; `{1}`,`{T}`, sac: compra 1 · F6 CMC 2 · **R$ 2,25** (26/08) | **Reserva, não dispensa.** Zero pip (Medição 2), mas o auto-saque a coloca no cemitério. Só entra se a Fase 6 cortar [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket)/[**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) por outro motivo. |

---

## 4. Candidatas — ramp padrão (9 titulares + 4 reservas)

| # | Carta | CMC | Tipo | O que gera | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| 1 | [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) | 1 | Artifact | `{C}{C}` | (1) **intocável** por decisão do briefing; (2) +4,8 pp na métrica "escapar **e** ter 2 de mana livre" (Medição 3) — é a melhor carta do pacote nesse eixo; (3) T1–T2 desbloqueia rocha + remoção no mesmo turno. **F7 declarado: paga zero dos 4 pips do escape.** | **sim (caixa + lista)** | **R$ 7,95** (18/09) |
| 2 | [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet) | 2 | Artifact | `{R}` ou `{W}` | (1) **intocável** por decisão do briefing; (2) mv 2 de cor, o perfil exato da Medição 2; (3) T2 → Phlage no T3 com 1 de sobra | **sim (lista)** | **R$ 4,00** (18/09) |
| 3 | [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet) | 2 | Artifact RW | `{1}`,`{T}` → `{R}{W}` | (1) **é a única peça do pool que entrega os dois pips numa ativação só** — converte 1 genérico em metade do escape; (2) mv 2; (3) transforma o `{C}` do [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) em `{R}{W}`, que é o buraco da Medição 2 | não | **R$ 1,50** (12/08) |
| 4 | [**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction) | 2 | Artifact RW | `{C}`, ou `{R}`/`{W}` por 1 de vida | (1) mv 2 de cor à escolha; (2) o modo incolor cobre o genérico das mágicas e o colorido cobre o escape — **flexível dentro do mesmo turno**; (3) 1 de vida é irrelevante num deck que ganha 3 por gatilho do comandante | não | **R$ 3,80** (19/09) |
| 5 | [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) | 2 | Artifact | qualquer cor, **3 usos** | (1) mv 2 de qualquer cor; (2) **já está na caixa**; (3) quando esgota os contadores vira **custo adicional de [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers)** (sacrificar artefato) — converte rocha morta em 2 cartas. **F7: entra virada (custo medido ≈1 pp no T4, 0 no T6) e expira em 3 ativações.** | **sim (caixa)** | R$ 0,11 (12/08) |
| 6 | [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert) | 2 | Artifact Creature 2/1 | qualquer cor (2 de vida) | (1) mv 2 de qualquer cor; (2) **corpo 2/1 — conta no piso de 14–17 da Fase 2** e ataca se a mesa esvaziar; (3) artefato + criatura num deck com 27,9% de mãos sem corpo. **F7: disputa o próprio `{T}` entre mana e ataque — na prática nunca ataca.** | **sim (caixa)** | R$ 0,15 (12/08) |
| 7 | [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) | 3 | Artifact | `{R}` ou `{W}` | (1) mana de cor; (2) **sacrifica-se para comprar 1** — resolve a dor 2 quando o mana já não falta; (3) **o sacrifício a manda para o cemitério = combustível de escape**. Tripla função num slot. **F7: mv 3 (−3,9 pp no T5 contra uma mv 2).** | **sim (caixa)** | `a cotar` |
| 8 | [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) | 3 | Artifact RW | `{R}` ou `{W}` | (1) mana de cor; (2) **saca 2** e vai para o cemitério; (3) a Fase 2 já o marcou como preferível às demais rochas da caixa exatamente por isso. **F7: mv 3; a ativação de saque custa `{R/W}×4`, que é o turno inteiro.** | **sim (caixa)** | `a cotar` |
| 9 | [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) | 2 | Artifact Creature 0/1 | `{C}` (moendo 1) | (1) **slot compartilhado com o pacote de combustível da Fase 2** — é 1 dos 4 motores; (2) moe **em velocidade de instantâneo**, então abastece no end step do oponente e o escape sai no seu turno; (3) corpo que conta no piso. **F7: zero pip (Medição 2) — ele está aqui pelo cemitério, não pelo mana.** | não | **R$ 0,17** (12/08) |
| **R1** | [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise) | 2 | Artifact Creature 0/2 | **qualquer cor** | (1) mv 2 de qualquer cor — perfil ideal da Medição 2; (2) **0/2 voador**: bloqueia o céu, que a Fase 2 nomeou como o buraco de RW ao defender a `Gisela`; (3) corpo no piso de 14–17. **É a minha primeira troca se a Fase 6 quiser trocar um mv 3 por um mv 2** (entra no lugar de [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket): +2,7 pp no T4, +2,0 pp no T5). | não | `a cotar` |
| **R2** | [**Zookeeper Mechan**](https://www.ligamagic.com.br/?view=cards/card&card=Zookeeper+Mechan) | 2 | Artifact Creature 1/3 | `{R}` | (1) mv 2, corpo **1/3** — bloqueia agressão de T2–T4 melhor que qualquer outra rocha-criatura do pool; (2) **já está na caixa**. **F7: só `{R}` — cobre metade do pip.** Entra se a Fase 2 pedir corpo antes de cor. | **sim (caixa)** | R$ 0,09 (12/08) |
| **R3** | [**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone) | 2 | Artifact | `{C}` | (1) mv 2; (2) auto-saque que a manda ao cemitério. **F7: zero pip.** Reserva de orçamento. | não | R$ 2,25 (26/08) |
| **R4** | [**Gold Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Gold+Myr) / [**Iron Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Iron+Myr) | 2 | Artifact Creature 1/1 | `{W}` / `{R}` | (1) mv 2 + corpo; (2) **[**Gold Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Gold+Myr) é a peça específica para `{W}{W}`**, que é a metade mais escassa do escape. Reserva se a Fase 6 entregar uma base vermelha forte e branca fraca. | não | `a cotar` |

**Valor do pacote titular (régua = valor das 99, LigaMagic menor):** [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) 7,95 + `Arcane
Signet` 4,00 + `Boros Signet` 1,50 + `Talisman of Conviction` 3,80 + `Sphere of the Suns` 0,11 +
[**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert) 0,15 + [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) 0,17 = **R$ 17,68 em 7 das 9**, cotações de 12/08 a 19/09.
[**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) e [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) ficam **`a cotar`** — são as duas últimas do pacote e estão
**na caixa**, então nenhuma delas é compra. **Compra real do pacote: 3 cartas** ([**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet),
[**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction), [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin)) = **R$ 5,47**.
Dos R$ 17,68, **R$ 11,95 (68%) são os dois intocáveis**, sobre os quais não houve escolha.

---

## 5. Curva e composição do pacote — o que a simulação exige

| Restrição | Medido | Regra para a Fase 6 |
|---|---|---|
| **mv 2 × mv 3** | 9 rochas todas mv 2 → T5 **84,0%**; 5 mv2 + 4 mv3 → 82,2%; todas mv 3 → **80,1%** | **Mínimo 6 das 9 em mv ≤ 2.** O pacote proposto tem **7**. |
| **colorida × incolor** | 9 coloridas T5 78,4% × 9 incolores **62,7%** (com *mais* mana total) | **Mínimo 6 das 9 produzindo `{R}` ou `{W}` à escolha.** O pacote tem **7** ([**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) e [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) são as duas incolores, e ambas estão lá por outro motivo declarado). |
| **rocha que entra virada** | 1 virada custa ~1 pp no T4 e **0 no T6**; 4 viradas custam ~1,4 pp no T4 | **Até 2 viradas é de graça.** O pacote tem 1 ([**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns)). Não é um critério de corte. |
| **dupla função** | — | 5 das 9 fazem outra coisa: 2 são **corpo** ([**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert), [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin)), 2 **sacam carta e viram combustível** ([**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere), [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket)), 1 é **motor de cemitério** ([**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin)). Só [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring), [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet), [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet) e `Talisman` são rocha pura. |

---

## 6. Candidatas — ramp explosivo (2–3), redefinido

**O explosivo clássico não serve a este deck.** Dobradores de mana e rituais grandes existem para
pagar uma carta que custa 7–9. A carta mais cara que este deck quer pagar é **4** (o escape), e o
pico real de demanda é **4 + 2 ou 3 de remoção no mesmo turno**. O que serve é **burst de mana
colado num loot**: ele paga o turno do escape, enche o cemitério e compra carta — três metas num
slot, e o slot **já está no orçamento de 10 loots da Fase 2**.

| Carta | CMC | Tipo | O que gera | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|
| [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) | 4 | Instant R | **2 Treasures** + saca 2 (descarta 1) | (1) **já está no pool da §3.3 da Fase 2** como loot — slot compartilhado, custo líquido **zero**; (2) os 2 Treasures fazem **qualquer cor**: cobrem `{W}{W}` do escape num turno em que só sobrou Mountain; (3) 2 cartas ao cemitério; (4) **instantâneo** — joga no end step do oponente e o mana fica para o seu turno | não | `a cotar` |
| [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) | 3 | Sorcery R | **1 Treasure** + saca 2 (descarta 1) | (1) mesma função por 1 mana a menos — **a peça de curva** do trio; (2) 2 cartas ao cemitério; (3) o Treasure paga o 4º pip do escape no turno seguinte | não | `a cotar` |
| [**Unexpected Windfall**](https://www.ligamagic.com.br/?view=cards/card&card=Unexpected+Windfall) | 4 | Instant R | **2 Treasures** + saca 2 (descarta 1) | (1) idêntica ao [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) em efeito; (2) instantânea. **F7 declarado: custa `{2}{R}{R}` — dois pips vermelhos num deck que já briga por `{W}{W}`.** Terceira na fila, atrás de [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) (`{3}{R}`, 1 pip) | não | `a cotar` |
| **R** | [**Strike It Rich**](https://www.ligamagic.com.br/?view=cards/card&card=Strike+It+Rich) | 1 | Sorcery R | 1 Treasure (+ haste) | (1) 1 mana para adiantar o escape em um turno; (2) **flashback** = segundo uso. **F7: o flashback a exila — ela deixa de ser combustível depois de usada** (a Fase 2 já limitou o deck a 3 cartas com flashback) | não | `a cotar` |

**Medido** (36 terrenos / 10 duais, 9 rochas, convertendo slots de loot barato em loot-com-tesouro):

| Loot-com-tesouro | T5 | T6 | T7 | T6 **com ≥2 de mana sobrando** |
|---|---|---|---|---|
| 0 | **49,6%** | 72,2% | 85,4% | 32,2% |
| 2 | 44,6% | 70,5% | 85,6% | 36,6% |
| **3 (recomendado)** | 42,2% | 69,2% | 85,0% | **38,5%** |
| 4 | 39,5% | 68,0% | 85,0% | 40,4% |

A troca é explícita e não é grátis: cada loot-com-tesouro custa **~2,5 pp de escape no T5** (porque
é mv 3–4 no lugar de um loot mv 1–2) e compra **~2 pp de ação dupla no T6**. **Parar em 3** —
na quarta, o T5 cai abaixo de 40% e o retorno já é marginal.

### 6.1 Explosivos dispensados, com ficha

| Carta | Ficha F1–F7 | Por que não |
|---|---|---|
| [**Mana Geyser**](https://www.ligamagic.com.br/?view=cards/card&card=Mana+Geyser) (**R$ 12,35**, 19/09) | F1 `{3}{R}{R}`: adiciona `{R}` por terreno virado dos oponentes · F3 Sorcery · F6 **CMC 5** · F7 depende do tabuleiro alheio | **Responde a pergunta que o briefing deixou aberta na v1** (*"o [**Mana Geyser**](https://www.ligamagic.com.br/?view=cards/card&card=Mana+Geyser) produz só `{R}`"*): ele produz **exclusivamente `{R}`** e **não paga nenhum dos dois `{W}` do escape**. Um monte de mana da cor que não está faltando. Somando: CMC 5 num deck cujo topo é 4, R$ 12,35 = **6,2% do teto**, e é feitiço (não segura mana aberto). Dispensa em todos os eixos. |
| [**Jeska's Will**](https://www.ligamagic.com.br/?view=cards/card&card=Jeska%27s+Will) | F1 modal; **ambos os modos só "se você controla um comandante"**; (a) `{R}` por carta na mão de um oponente, (b) exila 3 do topo e deixa jogar · F6 CMC 3 | **A cláusula não dispara neste deck.** O Phlage passa a maior parte do jogo **no cemitério ou na zona de comando** — em nenhuma das duas você "controla um comandante". Sem os dois modos ela é um ritual só de `{R}` (mesmo problema do [**Mana Geyser**](https://www.ligamagic.com.br/?view=cards/card&card=Mana+Geyser)) **ou** um impulse de 3 que **exila** — anti-combustível (§3.4 da Fase 2). Preço **`a cotar`** e historicamente alto. |
| [**Smothering Tithe**](https://www.ligamagic.com.br/?view=cards/card&card=Smothering+Tithe) | F1 oponente que compra paga `{2}` ou você faz Treasure · F3 Enchantment · F6 CMC 4 | **Dispensa por preço, não por função** — a função é excelente (Treasures de qualquer cor, escalando com a mesa). Preço **`a cotar`**, e a carta é notoriamente das mais caras de branco em EDH. **Devolvo ao orquestrador como corte condicionado à cotação:** se vier abaixo de ~R$ 20, ela é a melhor peça desta seção e entra no lugar do [**Unexpected Windfall**](https://www.ligamagic.com.br/?view=cards/card&card=Unexpected+Windfall). Acima disso, não cabe no teto ao lado de 21 remoções. |
| [**Solemn Simulacrum**](https://www.ligamagic.com.br/?view=cards/card&card=Solemn+Simulacrum) | F1 ETB busca básico **para o campo virado**; morre → compra 1 · F2 **2/2** · F3 Artifact Creature · F6 **CMC 4** | **Corte condicionado, devolvido à Fase 2.** Como ramp é ruim aqui: **CMC 4 é o custo do escape** — no turno em que você o joga, você não escapa. Mas ele é **corpo 2/2 + fixação de básico (escolhe a cor que falta) + carta ao morrer + vai para o cemitério**. Quatro funções fora do meu eixo. **Não o corto**: a função descoberta seria corpo + saque, e quem decide o piso de corpos é a Fase 2. |
| [**Wayfarer's Bauble**](https://www.ligamagic.com.br/?view=cards/card&card=Wayfarer%27s+Bauble) / [**Burnished Hart**](https://www.ligamagic.com.br/?view=cards/card&card=Burnished+Hart) | F1 sacrificam-se para buscar básico(s) **para o campo, virado(s)** · F6 1+2 e 3+2 manas | Dispensa por **taxa**. 3 manas por um terreno virado (`Bauble`) ou 5 por dois (`Hart`) é lento demais para um deck cujo turno decisivo é o T5–T6. Elas **fixam pip de verdade** (busca-se a cor que falta) e viram combustível ao sacrificar — por isso ficam registradas, não esquecidas. Se a Fase 6 não conseguir 10 fontes duplas, [**Burnished Hart**](https://www.ligamagic.com.br/?view=cards/card&card=Burnished+Hart) volta à mesa. |
| Dobradores ([**Mana Reflection**](https://www.ligamagic.com.br/?view=cards/card&card=Mana+Reflection)-likes), [**Gilded Lotus**](https://www.ligamagic.com.br/?view=cards/card&card=Gilded+Lotus) (5), [**Thran Dynamo**](https://www.ligamagic.com.br/?view=cards/card&card=Thran+Dynamo) (4) | F6 CMC 4–6, produzem `{C}` | **Duplamente errados aqui:** mana **incolor** (Medição 2) e **CMC acima do topo real da curva** (§1). Pagam por uma carta que este deck não tem. |

---

## 7. Slots — quantos eu preciso e de onde saem

Não há corte a fazer: **o deck v2 é montado do zero**. O que existe é o orçamento de 63 slots
não-terreno (a 36 terrenos) que a Fase 2 já repartiu. Minha conta:

| Categoria | Slots | Fonte |
|---|---|---|
| Combustível (10 loot + 4 motores) | 14 | Fase 2 §3 |
| Remoção (21) + wipes (3,5) | 24,5 | Fase 2 §5 |
| Pingadores | 6 | Fase 2 §6.1 |
| Multiplicadores de dano | 3,5 | Fase 2 §6.2 |
| **Subtotal Fase 2** | **48** | |
| **Ramp padrão (meu pedido)** | **9**, dos quais **1 compartilhado** ([**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) já é motor) → **8 novos** | esta fase |
| **Ramp explosivo** | **3**, **todos dentro dos 10 loots** → **0 novos** | esta fase |
| **Total comprometido** | **56** | |
| **Livre** | **7** | para draw não-loot (Fase 3), proteção ([**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm)), [**Sunforger**](https://www.ligamagic.com.br/?view=cards/card&card=Sunforger), finisher opcional |

**Peço 8 slots novos.** É o que cabe. A meta do pipeline (10–11) pediria **10 a 11 slots novos** e
comeria 2–3 dos 7 livres — que são exatamente os slots da proteção e do draw não-loot. A Medição
1b mostra que essa troca é **negativa**: a 10ª rocha vale 0,0 pp e a 11ª vale −1,0 pp no T6.

**Ordem de cedência, se a Fase 6 precisar de slot** (e a justificativa de cada cedência):

1. **[**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket)** — a peça com a pior curva do pacote (mv 3) e função redundante com
   [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) (mesma cor, mesmo auto-saque). Funções descobertas: nenhuma — as três
   (mana de cor, saque, combustível) continuam cobertas pelo [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere).
2. **[**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin)** — se ceder, **quem perde é a Fase 2, não eu**: ele é 1 dos 4 motores, e o pacote
   de ramp perde só a sua fonte mais fraca (zero pip). A função descoberta é
   **moagem automática em velocidade de instantâneo** — reposição nomeada: [**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece)
   (2 cartas/turno, mv 2), que a Fase 2 já tem no pool.
3. **Nada mais.** Abaixo de 8 fontes, o T4 cai de 66,9% para 59,4% e o Phlage no T3 sai de 85% para
   ~80%. **O pacote não cede a 7.**

**O que eu recomendo que ceda ANTES do meu 9º slot:** a **22ª remoção**. A Fase 2 declarou a faixa
como 20–22 e mediu que a diferença entre 20 e 22 é **+3,3 pp** em P(≥1 na mão de 7). Meu 9º slot
vale **+0,9 pp de escape no T6 e +1,9 pp de ação dupla**, e o 8º vale **+6,3 pp no T4**. A
comparação é entre grandezas diferentes e eu a declaro como tal — **a decisão é do orquestrador**,
não minha.

---

## 8. O que eu espero da manabase (Fase 6) — medido, não opinado

Terrenos são da Fase 6. Estes são os números que o meu pacote assume, e todos os três movem mais a
agulha do que a 10ª rocha:

### 8.1 Fontes duplas valem mais que terrenos a mais

P(pagar `{R}{R}{W}{W}`), pacote de 9 rochas coloridas fixo:

| Terrenos | Fontes duplas | T4 | T5 | T6 |
|---|---|---|---|---|
| 36 | 4 | 64,2% | 75,9% | 82,9% |
| 36 | **6** (a `lista.txt` atual) | 67,8% | **79,0%** | 85,7% |
| 36 | 10 | 72,3% | **82,1%** | 87,8% |
| 36 | 14 | 76,6% | **85,1%** | 90,6% |
| 37 | 10 | 73,7% | 83,4% | 88,9% |
| 38 | 10 | 75,0% | 84,3% | 89,7% |
| 38 | 14 | 78,7% | **87,0%** | 91,6% |

**Passar de 6 para 14 fontes duplas, sem mexer na contagem de terrenos, vale +6,1 pp no T5** — mais
do que **três rochas**, e a custo de slot **zero**. Passar de 36 para 38 terrenos vale +2,2 pp.
**Pedido: mínimo 10 fontes duplas, alvo 12–14.** Hoje há 6 ([**Command Tower**](https://www.ligamagic.com.br/?view=cards/card&card=Command+Tower), [**Boros Guildgate**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Guildgate),
[**Sandstone Bridge**](https://www.ligamagic.com.br/?view=cards/card&card=Sandstone+Bridge), [**Stone Quarry**](https://www.ligamagic.com.br/?view=cards/card&card=Stone+Quarry), [**Wind-Scarred Crag**](https://www.ligamagic.com.br/?view=cards/card&card=Wind-Scarred+Crag), [**Fields of Strife**](https://www.ligamagic.com.br/?view=cards/card&card=Fields+of+Strife)).

### 8.2 Terreno que entra virado custa pouco — mas só até ~4

36 terrenos, 9 rochas coloridas:

| Fontes duplas | quantas entram viradas | T4 | T5 | T6 |
|---|---|---|---|---|
| 6 | 0 | 67,2% | 78,3% | 84,8% |
| 10 | **0** | **71,5%** | **81,6%** | **87,6%** |
| 10 | 4 | 69,3% | 81,2% | 87,5% |
| 10 | 7 | 67,5% | 80,2% | 86,8% |
| 10 | 10 | 65,0% | 79,2% | 86,0% |
| 14 | 0 | 76,0% | 84,6% | 89,9% |
| 14 | 6 | 70,8% | 83,3% | 88,9% |

Leitura: **10 duais com 4 viradas (69,3% no T4) ainda ganham de 6 duais desviradas (67,2%)** — ou
seja, **acrescentar tapland dupla é positivo**. Mas 10 duais **todas viradas** (65,0%) perdem para
6 desviradas. Cada dual desvirada vale ~+1,1 pp no T4; cada dual virada, ~+0,5 pp.

**Pedido: no máximo 4–5 das fontes duplas entrando viradas.** Prioridade para as que entram
desviradas — e o motivo é o desta fase, não estético: **um deck que quer mana aberto no turno do
oponente não pode pagar o turno inteiro por uma fonte de cor.** Gates e tri-lands são o último
recurso, não o primeiro.

### 8.3 O 37º e o 38º terreno ganham do 10º e do 11º ramp

+2 terrenos (36→38, 10 duais) = **+2,2 pp** no T5. +2 rochas (9→11) = **+2,1 pp** no T5 no modelo
isolado, e **−2,2 pp** no modelo conjunto (porque a rocha come slot de combustível). E há um
desempate que só existe neste deck: **terreno excedente aqui não é carta morta** — ele é descartado
ao loot (combustível) e é munição do [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) (retrace). **Recomendo 37–38 terrenos**, tirados
do excedente de queima, não do ramp.

---

## 9. Pendências desta fase

1. **3 cartas do pacote titular estão `a cotar`**: [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) e [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) (ambas
   **na caixa** — cotação serve só para o valor das 99, não é compra) e [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise)
   (reserva R1, seria compra). Os 3 explosivos ([**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score), [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils),
   [**Unexpected Windfall**](https://www.ligamagic.com.br/?view=cards/card&card=Unexpected+Windfall)) também estão `a cotar`. **Nenhum total desta fase está fechado** — o que
   está fechado é **R$ 17,68 em 7 das 9 titulares** (LigaMagic menor, cotações de 12/08 a 19/09),
   dos quais **R$ 11,95 são os dois intocáveis** e a **compra real é de R$ 5,47** em 3 cartas.
2. **[**Smothering Tithe**](https://www.ligamagic.com.br/?view=cards/card&card=Smothering+Tithe) — corte condicionado à cotação.** É funcionalmente a melhor peça da §6 e
   está fora só por preço presumido, que eu não tenho. Se a cotação vier baixa, ela entra no lugar
   do [**Unexpected Windfall**](https://www.ligamagic.com.br/?view=cards/card&card=Unexpected+Windfall).
3. **[**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler) e [**Solemn Simulacrum**](https://www.ligamagic.com.br/?view=cards/card&card=Solemn+Simulacrum) — cortes condicionados devolvidos**, à Fase 2 e ao
   orquestrador. Os dois exercem função **fora da minha especialidade** (corpo no piso de 14–17,
   e no caso do `Solemn` também saque e fixação de básico). Eu os reprovo **como ramp**; não os
   corto do deck.
4. **[**Melded Moxite**](https://www.ligamagic.com.br/?view=cards/card&card=Melded+Moxite) foi catalogada errado como rocha** em rodadas anteriores — ela **não produz
   mana**. O ETB é loot 1-por-2. **Devolvida à Fase 3** como candidata de draw/combustível.
5. **Correção à §4.2 da Fase 2**, para o `report.md`: o Phlage no T3 **não é garantido**. É
   **85,0%** com este pacote (75,4% sem ramp; 89,8% com 38 terrenos e 12 duais). O gargalo é land
   drop + pip, não compra.

---

## Revisão 2026-09-24 — pacote de reanimação

**Ramp na lista revisada: 9 padrão / 2 explosivos. Proposta: 10 / 1, no mesmo número de slots.**
Uma troca, com a carta que sai marcada como **corte condicionado à Fase 2**:
[**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) (caixa, R$ 0,11) entra e [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) (R$ 7,65) sai. Delta **−R$ 7,54**.
Não proponho subir o ramp tirando slot de outra categoria. A §R4 mostra que o 11º slot de ramp compra
menos de 0,5 pp em cada marco novo.

> **Por que esta seção existe.** A Fase 2 (§12 do `02-theme.md`) trocou o eixo para **reanimação
> em loop + escape concedido + controle**, e o orquestrador (`08-meta-edhrec.md`) pediu os marcos
> refeitos. A conclusão das §2.2 e §7 ("a 10ª rocha vale 0,0 pp, a 11ª −1,0 pp") foi medida para o
> **escape próprio no T6**, com cada rocha nova tirando slot de **combustível**. Os dois pressupostos
> caíram: o escape próprio virou plano B, e o combustível desceu de 14 para ~11. O que está acima desta
> seção continua valendo como registro. **Onde houver conflito, vale esta seção.**
>
> **Regra 5.** `decisions.md` não tem linha para nenhuma rocha, nem para [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score), nem para
> [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns). A [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns), porém, **saiu da v2** na Fase 6 (`06-manabase.md`
> §5.2, C10; `report.md` §7.2) para dar o slot ao [**Zookeeper Mechan**](https://www.ligamagic.com.br/?view=cards/card&card=Zookeeper+Mechan). O que mudou desde então está na §R6.
>
> **Regra 6.** Oracle puxado nesta sessão com `bin/mtgdb oracle` para as 62 não-terrenos da lista
> revisada, os 15 terrenos não-básicos e todas as candidatas citadas.
>
> **Regra 7.** A caixa (`mtgdb collection -list`, 116 cartas) foi varrida antes de qualquer busca.
> Fontes de mana na caixa e fora do deck: [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns), [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler), [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket),
> [**Manalith**](https://www.ligamagic.com.br/?view=cards/card&card=Manalith), [**Seer's Lantern**](https://www.ligamagic.com.br/?view=cards/card&card=Seer%27s+Lantern), [**Magnifying Glass**](https://www.ligamagic.com.br/?view=cards/card&card=Magnifying+Glass) e [**Prophetic Prism**](https://www.ligamagic.com.br/?view=cards/card&card=Prophetic+Prism). [**Moss Diamond**](https://www.ligamagic.com.br/?view=cards/card&card=Moss+Diamond) está
> fora da identidade. [**Mycosynth Wellspring**](https://www.ligamagic.com.br/?view=cards/card&card=Mycosynth+Wellspring) e [**Pilgrim's Eye**](https://www.ligamagic.com.br/?view=cards/card&card=Pilgrim%27s+Eye) põem terreno na **mão**, então não são ramp.
> As quatro primeiras foram medidas (§R7). **A troca proposta não tem compra.**
>
> **Regra 2.** Os preços saem de `bin/mtgdb prices` (LigaMagic, menor), com cotações de 2026-08-12
> a 2026-09-24. [**Gold Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Gold+Myr), [**Star Compass**](https://www.ligamagic.com.br/?view=cards/card&card=Star+Compass) e [**Manalith**](https://www.ligamagic.com.br/?view=cards/card&card=Manalith) estão `a cotar`, e nenhuma delas faz parte
> da troca proposta.
>
> **EDHREC.** Não chamei. O `08` §3 transcreveu só `High Synergy Cards` e `Top Cards`. A seção
> `Mana Artifacts` **não foi transcrita**, então nenhuma candidata desta revisão vem do meta. As
> rochas da lista-modelo Budget Phlag ([**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet), [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet), [**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction)) já estão no deck.

### R1. O que mudou na pergunta

| | Medição anterior (§2) | Plano novo (§12 do `02`) |
|---|---|---|
| Turno-alvo | escape próprio no T6 | Phlage no T3 → reanimação + outra jogada no T4–T5 → motor de 4 no T4–T5 → Phlage escapado pelo Dial no T5 |
| Custo decisivo | `{R}{R}{W}{W}`, **zero genérico** | `{1}{R}{W}` (Phlage, escape do Dial ou da Desdemona) · `{W}` a `{2}{W}` (reanimação) · `{2}{X}{Y}` (motores) · `{3}` (Dial) |
| Onde a rocha nova tirava slot | combustível, com o escape dependente de cemitério | remoção ou payoff, porque o combustível desceu para ~11 e o Dial faz vidência 3 |
| 4-drops | 5 | **9**: [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer), Torbran, [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king), [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) + Teshar, [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth), Desdemona, [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer), [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) |

**Consequência para a §2.3 (pips).** Aquela medição mostrou que mana incolor não pagava o escape de
quatro pips. Os custos do plano novo têm parte genérica, então **o incolor voltou a pagar quase tudo**.
Na §R7, uma rocha incolor de 2 manas fica a 0,7 pp da colorida no T3 e empata nos outros marcos.

### R2. Método

- **Monte Carlo com 60 mil partidas por configuração, na jogada.** As 99 cartas são a lista revisada
  (v2 do `report.md` §6 com as 15 trocas da §12 aplicadas). A base é a real: 37 terrenos, 12 Mountain e
  12 Plains, 11 fontes duplas (Clifftop, Snarl e Summit com a condição de entrada, 5 que entram viradas)
  e [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) + Throne produzindo `{C}`. Mulligan londrino simplificado.
- **Números aleatórios comuns.** A carta trocada ocupa a mesma posição no baralho, então duas
  configurações jogam as mesmas partidas e só a carta muda. **Diferença abaixo de ~0,3 pp é empate.**
- **Política de jogo:**
  1. T1–T2: rochas e dorks primeiro.
  2. Phlage da zona de comando assim que couber, com prioridade no T3.
  3. Na ordem: escape pelo Dial → conjurar o Dial (e escapar no mesmo turno, se houver 6 manas) →
     escape pela Desdemona, se ela atacou → **um** motor de 4 → **uma** reanimação, se o Phlage está
     no cemitério → rochas → [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) / [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) → resto do mais caro para o mais barato.
- **Detalhes do modelo:**
  - Os loots estão modelados: [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse), [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre), Light Up
    the Stage, [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils), [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) e a ativação do Bankbuster.
  - Dorks têm enjoo.
  - A vidência 3 do Dial manda 3 cartas ao cemitério.
  - A [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) perde um contador sempre que a mana dela é necessária.
- **Fora do modelo:** compras da [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) e o impulse do Ark; gatilhos do [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan); combate.
  O custo de cemitério da Desdemona é dado como pago, porque o modelo só conta mágicas e descartes no
  cemitério e subestima o real. O número dela é teto.
- **Os números absolutos não se comparam com os da §2.** Mudaram a lista, a base e a política. A
  comparação vale **dentro** das tabelas abaixo.

### R3. Os marcos novos com o ramp atual (9 / 2)

| Marco | P absoluta | P(ter a carta a tempo) | P(marco \| tem a carta) |
|---|---|---|---|
| Phlage da zona de comando até o T3 | **83,8%** | — | — |
| **Motor de 4 até o T4** | **23,2%** | 39,4% (um dos 5 visto até o T4) | **58,8%** |
| Motor de 4 até o T5 | 33,3% | | |
| Reanimação + outra jogada no mesmo turno, **no T4** | 14,7% | | |
| **Reanimação + outra jogada no mesmo turno, no T4 ou no T5** | **23,1%** | 41,4% (uma das 5 vista até o T5, com o Phlage conjurado até o T4) | **55,8%** |
| **Phlage escapado pelo Dial até o T5** | **8,1%** | 10,9% (Dial visto até o T5) | **74,2%** |
| Phlage escapado pelo Dial **ou** pela Desdemona até o T5 | 12,9% (teto) | | |
| Motor em campo **e** reanimação conjurada até o T5 | 9,2% | | |
| [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) pagável no T6 (`{4}{W}{W}`) | 53,9% | | |
| Gisela pagável no T7 (`{4}{R}{W}{W}`) | 47,2% | | |
| Mana média disponível no T6 / T8 | 5,7 / 7,0 | | |

**Leitura: o gargalo dos três marcos é a carta, não a mana.** Pela hipergeométrica (99 cartas), a
chance de ver um dos 5 motores em 10 cartas é 42%, e a de ver o Dial em 11 cartas é 11,1%. **Com a
carta na mão, o ramp atual entrega o marco em 56–74% dos casos.** Rocha nenhuma muda a coluna do
meio. Quem muda é o draw (Fase 3) e a redundância de peças (Fase 2).

### R4. Quanto cada rocha move os marcos

| Configuração | Phlage T3 | Motor T4 (\| tem) | Reanim. + outra T4–T5 (\| tem) | Dial T5 (\| tem) | Titan T6 | Gisela T7 |
|---|---|---|---|---|---|---|
| 8 (sem [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert)) | 82,4 | 22,5 (57,1) | 22,7 (55,2) | 8,0 (72,7) | 51,2 | 44,2 |
| **9 / 2 atual** | **83,8** | **23,2 (58,8)** | **23,1 (55,8)** | **8,1 (74,2)** | **53,9** | **47,2** |
| **10 / 1: Big Score → Sphere of the Suns (proposta)** | **84,8** | **23,6 (60,2)** | **23,7 (56,9)** | **8,2 (75,7)** | **55,7** | **48,6** |
| 10 / 2: + [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) no lugar de uma remoção | 84,8 | 23,8 (60,4) | 23,4 (56,2) | 8,3 (75,6) | 56,1 | 49,7 |
| 11 / 1: a proposta + [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler) no lugar de uma remoção | 85,1 | 24,0 (61,1) | 24,0 (57,4) | 8,3 (76,6) | 57,9 | 51,4 |
| 11 / 2: + Sphere + [**Gold Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Gold+Myr) no lugar de duas remoções | 85,6 | 24,2 (61,6) | 23,7 (56,5) | 8,4 (76,6) | 58,6 | 52,6 |

**Três leituras:**

1. **Ao contrário da medição anterior, toda rocha agora soma.** Uma rocha nova vale +1,0 pp de
   Phlage no T3, +0,6 pp de motor no T4 e +2,2 pp de [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) no T6. Antes eram 0,0 e −1,0, porque
   a rocha nova tirava combustível de um escape que dependia de cemitério. Agora ela tira remoção, e a
   curva ficou mais pesada.
2. **Mesmo assim, o ganho é pequeno: +0,3 a +1,0 pp por marco.** O 11º slot de ramp (linha 11 / 1
   contra a proposta) compra +0,3 pp no T3, +0,4 no motor, +0,3 na reanimação e +0,1 no Dial, e custa
   uma remoção ou um payoff. **Não proponho.** A troca seria fora da minha categoria, e o ganho fica
   no nível do empate.
3. **A conversão proposta captura quase todo o ganho da 10ª rocha sem gastar slot.** A razão está na §R5.

### R5. O que cada peça atual faz nos marcos novos

Cada peça foi trocada por uma carta em branco do mesmo tipo de custo, e a coluna mostra o que se perde:

| Peça substituída por carta em branco | Phlage T3 | Motor T4 | Reanim. + outra | Dial T5 | Titan T6 | Gisela T7 | Leitura |
|---|---|---|---|---|---|---|---|
| [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) (→ 2-drop em branco) | −0,1 | −0,1 | +0,2 | 0,0 | −0,7 | −1,2 | **≈ 0 do T3 ao T5.** Com CMC 4, disputa o T4 com os 9 quatro-drops e quase nunca é a jogada do T4 |
| [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) (→ 2-drop em branco) | −0,1 | −0,1 | 0,0 | −0,1 | −1,0 | −1,4 | ≈ 0 do T3 ao T5. **Fica**: é loot da Fase 2 por R$ 0,06, e o Treasure é bônus |
| [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) (→ 3-drop em branco) | −0,1 | 0,0 | 0,0 | 0,0 | **−1,9** | **−2,7** | É a **rocha do jogo longo**. **Fica** |
| [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert) (→ 2-drop em branco) | **−1,4** | −0,7 | −0,4 | −0,1 | −2,7 | −3,0 | As rochas de 2 manas sustentam o começo. **Fica** |

**O "explosivo" deste deck não faz trabalho de explosivo.** A §6 já tinha redefinido a categoria como
"burst de mana colado num loot" para pagar o turno do escape de 4 pips. Esse turno virou plano B.
Nos marcos novos, os dois Treasures do [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) chegam no T4+ e não mudam nada até o T5. No topo da
curva ([**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) 6, Gisela 7, Approach 7), a rocha permanente de 2 manas rende mais que o burst:
+1,8 pp no Titan e +1,4 pp na Gisela quando ela entra no lugar dele.
**A meta de 2–3 explosivos não se aplica a este plano, e declaro isso aqui.** Fica 1 nominal ([**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils)).

### R6. A troca proposta

| Entra | Sai | R$ |
|---|---|---|
| [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) (caixa) | [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) — **corte condicionado à Fase 2** (combustível) | +0,11 − 7,65 = **−7,54** |

**Ficha da que sai: Big Score** (`{3}{R}` instantâneo, oracle conferido)

| Eixo | Função | Quem cobre depois do corte |
|---|---|---|
| F1 | Custo adicional de descartar 1, compra 2 e cria **2 Treasures**, em velocidade de instantâneo | Treasures: [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) (3 manas de qualquer cor, espalhadas em 3 turnos), [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) (1 Treasure) e [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) (1 Treasure). Loot: [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) (mesmo texto, feitiço de 3), [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) e [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) |
| F1 | **Saldo de +1 carta** | **Descoberto.** O [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) **não** está entre as 12 fontes de draw da Fase 3, mas a carta existia. [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) e o `{T}` do [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) são os motores que ficam. Reserva que devolve a carta: [**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone) (§R7) |
| F2 | Sem corpo | — |
| F3 | Instantâneo que vai ao cemitério: 2 cartas de **combustível** (ele e o descarte) e um tipo para o delirium da [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | Combustível dedicado cai de **11 para 10**, mais a vidência 3 do [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial). **Função da Fase 2 → condicionado.** O tipo instantâneo continua com outros 10 no deck |
| F4 | Não recebe nada | — |
| F5 | Fixação em burst (os Treasures fazem `{W}{W}`), e o descarte dispara o **Inti** (impulse) | Fixação: [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns), 6 rochas de cor à escolha e 11 duplas. Inti: [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils), [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre), [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore), [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) e o descarte de ataque do próprio Inti |
| F6 | CMC 4, **um dos 9 quatro-drops**. Conjurado no fim do turno do oponente com a mana de remoção guardada | **Medido na §R5: ≈ 0 do T3 ao T5.** O uso no fim de turno fica **descoberto** (é o único loot instantâneo além do [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre)). Custo aceito |
| F7 | Disputa a mana do T4 com os 5 motores, Torbran, Jailer e Celebrate. Um pip vermelho | Com o corte, os 4-drops caem de 9 para 8 |

**Simetria (seção 3 do checklist).**
- **Haliya:** a Fase 3 está propondo a [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light), que ganha 1 de vida por artefato
  que entra. Os 2 Treasures do [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) davam +2, e a Sphere dá +1. Pesei isso dos dois lados, e não inverte.
- **Teshar:** o critério também vale ao contrário. A Sphere é artefato, então conta como mágica
  histórica para o [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) (21 → **22**), e o [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) não conta. Medido:
  P(Helix do Teshar até o T8) fica igual (8,3% → 8,3%), porque as rochas são conjuradas **antes** de
  o Teshar chegar. **É critério de desempate, não motivo.**

**Ficha da que entra: Sphere of the Suns** (`{2}` artefato)
- **F1:** entra virada com 3 contadores; `{T}` e remover um contador: 1 mana de qualquer cor.
- **F2:** sem corpo.
- **F3:** artefato, que conta como histórico para o Teshar. Ela nunca vai sozinha ao cemitério, então
  [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) e [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) não a reutilizam.
- **F4:** não recebe nada.
- **F5:** 3 manas de qualquer cor, para `{W}{W}` (Broodmoth, [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan), Jailer, [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory)) e
  `{R}{R}` (Bearer, Torbran).
- **F6:** conjurada no T2, paga o Phlage no T3.
- **F7:** **expira depois de 3 usos**, o que o modelo contabiliza, e entra virada. Custo medido: nenhum
  (é a melhor linha da §R7 no T3).

**O que mudou desde a saída da Sphere na Fase 6 (regra 5).** Ela saiu para dar o slot ao **Zookeeper
Mechan**, pelo **corpo** (piso de criaturas 16–17), e a Fase 6 mediu o custo em −0,5 pp num escape
de `{R}{R}{W}{W}`. Três coisas mudaram:
1. Ela volta no slot do [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score), e o **Zookeeper fica**. O motivo da saída continua respeitado: as criaturas seguem 17.
2. O custo foi medido para o escape próprio. Nos marcos novos, ela é a melhor rocha de 2 manas
   disponível (§R7), e o slot que ela ocupa rende ≈ 0 hoje (§R5).
3. O Teshar torna artefatos mágicas históricas.

Uma função dela **não volta**: a de sacrifício para o [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers), que saiu na §12.

**Se a Fase 2 recusar** (combustível abaixo de 11), a troca não acontece e o ramp fica em **9 / 2**.
Nada mais depende dela.

### R7. Alternativas medidas (todas no lugar do [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score)) e reservas

| Entra | Phlage T3 | Motor T4 | Reanim. + outra | Titan T6 | Gisela T7 | Na coleção? | R$ | Papel |
|---|---|---|---|---|---|---|---|---|
| [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) | **84,8** | **23,6** | **23,7** | 55,7 | 48,6 | **sim** | 0,11 | **proposta** |
| [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler) | 84,1 | 23,4 | 23,6 | 55,8 | 49,0 | **sim** | 0,10 | **1ª reserva, sem compra.** Produz `{C}`, e isso custa 0,7 pp no T3. Em troca é **permanente** e é **criatura de MV 2**: dispara a [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome), volta com a Broodmoth e leva o piso de 17 para 18. **O que mudou desde que eu a dispensei (§3.2):** o motivo era zero pip para `{R}{R}{W}{W}`, e agora os custos do plano têm parte genérica. **Se a Fase 2 quiser o corpo, ela é a escolha** |
| [**Star Compass**](https://www.ligamagic.com.br/?view=cards/card&card=Star+Compass) | 84,8 | 23,6 | 23,7 | 56,0 | 49,1 | não | `a cotar` | empata com a Sphere e não expira. É compra, então fica atrás de duas opções da caixa |
| [**Gold Myr**](https://www.ligamagic.com.br/?view=cards/card&card=Gold+Myr) | 84,4 | 23,6 | 23,6 | 56,0 | 49,0 | não | `a cotar` | `{W}` e corpo 1/1. É compra |
| [**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone) | 84,1 | 23,5 | 23,6 | 55,8 | 49,0 | não | 2,25 | **Reserva se a Fase 3 pedir a carta de volta**: é a única que mantém "sacrifica: compra 1" |
| [**Manalith**](https://www.ligamagic.com.br/?view=cards/card&card=Manalith) / [**Boros Locket**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Locket) (caixa) | 83,8 | 23,1 | 23,1 | 55,0 | 48,7 | sim | `a cotar` / 0,18 | **CMC 3 não mexe no começo** (igual ao [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) do T3 ao T5). Dispensa por curva (F6) |
| carta em branco de 2 manas | 83,7 | 23,1 | 23,3 | 53,2 | 46,0 | — | — | referência |

**Trocas dentro do ramp, medidas e recusadas:**
- **Commander's Sphere → Sphere of the Suns:** no começo empata com a proposta (84,9 / 23,8 / 23,7),
  mas perde no jogo longo (Titan 54,5 × 55,7; Gisela 47,2 × 48,6). Perde também o "sacrifica:
  compra 1", que manda a rocha ao cemitério, onde [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan), [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) e **Sevinne's
  Reclamation** a devolvem. Recusada.
- **Zookeeper Mechan → Gold Myr:** +0,3 pp no T3, que é empate. O Zookeeper é da caixa, é 1/3 e
  tem o `{6}{R}` de dreno de mana. Fica.

### R8. Contagem e custo depois da troca

| Item | Lista revisada | Com a troca |
|---|---|---|
| Ramp padrão | 9 | **10**: [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring), [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet), [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet), [**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction), [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere), [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin), [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert), [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise), [**Zookeeper Mechan**](https://www.ligamagic.com.br/?view=cards/card&card=Zookeeper+Mechan), [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) |
| Ramp explosivo | 2 | **1** ([**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils), nominal; §R5) |
| Slots de ramp | 11 | **11** |
| Curva (62 não-terrenos) | 2: 26 · 4: 9 · média 2,74 | 2: **27** · 4: **8** · média **2,71** |
| Cor (não-terrenos) | vermelho 22 · incolor 11 | vermelho **21** · incolor **12** |
| Mágicas históricas (Teshar) | 21 | **22** |
| Custo das 99 | R$ 162,12 | **R$ 154,58** (folga R$ 37,88 → **R$ 45,42**) |

Preços: LigaMagic (menor). [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) 7,65 (2026-09-23), [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) 0,11 (2026-09-23).

### R9. Sinais para as outras fases

| Fase | Sinal |
|---|---|
| **2 · tema** | **Decide o corte condicionado.** Com a troca, o combustível dedicado cai de 11 para 10, mais a vidência 3 do Dial. Se recusar, a troca cai inteira. Se quiser corpo, a [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler) substitui a Sphere (§R7) |
| **3 · draw** | O [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) não está nas 12 fontes, mas o saldo de +1 carta dele sai. A reserva que devolve a carta é a [**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone) (R$ 2,25). Para a **Haliya**: a Sphere dá +1 de vida ao entrar, e os 2 Treasures davam +2 |
| **5 · interação** | Nada. Não tirei nenhuma remoção. As linhas da §R4 que tiram remoção são **medidas**, não propostas |
| **6 · manabase** | Phlage no T3 medido em **83,8%** com a base real (37/11). Nada na troca depende da base |
| **7 · wincons** | Com a troca, [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) no T6 vai de 53,9% para 55,7%, e Gisela no T7 de 47,2% para 48,6%. Os marcos do T3–T5 estão na §R3, para a re-simulação |
| **Orquestrador** | Delta **−R$ 7,54**, sem compra. A seção `Mana Artifacts` do EDHREC nunca foi transcrita. Se quiser o radar de rochas, é a única chamada que faltou |
