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
