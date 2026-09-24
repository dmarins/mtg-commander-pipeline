# Condições de Vitória e Testes — Phlage, Titan of Fire's Fury (v2 · controle de atrito RW)

> **Natureza desta fase.** Não existe `deck.md`. Trabalhei sobre o pool das Fases 2 e 3 e sobre
> a `lista.txt`/caixa. O requisito duro recebido de `decisions.md` é literal: **entregar um fecho
> nomeado que não seja "atacar com o comandante"**, com turno medido e probabilidade de ter as peças.
>
> **Regra 6 — oracle.** Todo texto citado foi puxado com `bin/mtgdb oracle` nesta sessão.
> Rulings de [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), `Gisela`, [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition), [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate), [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power),
> [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun), [**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign) e `Phlage` puxados com `bin/mtgdb rulings`.
> **[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) não tem ruling oficial** — a leitura de regras dela está marcada como tal.
>
> **Regra 2 — preço.** Nenhum número do Scryfall. `bin/mtgdb prices` consultado para as 23
> candidatas; **13 têm cotação de 2026-09-23**, o resto sai como **`a cotar`**. As cotações de hoje
> **derrubaram duas recomendações da Fase 2** (§4.3) — é o item mais consequente deste relatório
> depois da simulação.
>
> **Regra 5 — conferido.** `decisions.md` registra apenas movimentos de zona de comando. Nenhuma
> carta proposta aqui é reposição de corte anterior. [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) está registrada
> em `decisions.md` como primeira candidata a finisher — ela entra, e o motivo está na §4.3.
>
> **Regra 11.** Nenhuma proposta é lock, stax ou infinito. A Rota 3 é um fecho de **duas cartas**;
> está declarada como tal na §5 para o usuário decidir (regra 9).

---

## 1. Resumo executivo — leia isto antes de comprar qualquer coisa

Três conclusões, todas medidas, todas desconfortáveis:

1. **A meta de saúde do pipeline (vencer até o T7) é inalcançável para este eixo.** Com a lista
   como ela está saindo das Fases 2–5, a **mesa inteira morre na mediana do T13–T14**, e em
   nenhuma configuração testada o deck mata 3 oponentes de 40 antes do T9. O **primeiro** oponente
   morre na mediana do **T10–T11**.
2. **O pacote de "5–7 pingadores + 3–4 multiplicadores" da §6 da Fase 2 não fecha o jogo.** Ele
   custa 9,5 slots e move a probabilidade de matar a mesa até o T13 de 18% para 23%. A premissa
   que o sustentava — *"o deck conjura 3–5 mágicas por turno no meio de jogo"* — está **errada por
   um fator de ~3**: a simulação dá **0,8 a 1,2 mágica por turno** no T8+ (§3.4). Um deck singleton
   de 99 cartas que saca 2,5 cartas por turno não sustenta storm-lite.
3. **As cotações de hoje mataram o plano de multiplicadores da Fase 2.** [**Solphim, Mayhem Dominus**](https://www.ligamagic.com.br/?view=cards/card&card=Solphim%2C+Mayhem+Dominus)
   custa **R$ 118,90 = 59,5% do teto** e [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) **R$ 54,54 = 27,3%**. Dentro do
   Commander 200, RW tem exatamente **dois** multiplicadores de dano viáveis: `Torbran, Thane of
   Red Fell` (R$ 6,98) e `Gisela, Blade of Goldnight` (R$ 15,00).

**O que eu entrego mesmo assim:** quatro caminhos de vitória nomeados, um deles capaz de matar a
mesa **num único turno** ([**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + wipe), dentro do orçamento de slots já reservado — na
verdade **devolvendo 0,5 slot** à Fase 3. E o número que o usuário precisa para decidir se quer
este deck: **ele começa a ameaçar no T6–T7 e só fecha no T13**. São 6 a 7 turnos de "estou na
frente e não consigo terminar". É exatamente a sensação que reprovou a v1.

---

## 2. Caminhos de vitória atuais

| Caminho | Cartas envolvidas | Turno estimado | Consistência |
|---|---|---|---|
| **R1 · Relógio focado — Phlage + Court of Ire** | `Phlage` (zona de comando/escape) · [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) · [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | **1º oponente morto: mediana T10–T11**; 42% até T10 | **Alta.** O Phlage é garantido (zona de comando); [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) está no pacote de draw da Fase 3, custo zero de slot |
| **R2 · Dreno agregado multiplicado** | 4 pingadores · [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) · `Gisela` · [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) · [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) | **mesa: mediana T13–T14**; 23% até T13 | **Baixa.** P(≥1 pingador **e** ≥1 multiplicador **vistos** até o T8) = **21,8%**; em jogo, menos |
| **R3 · [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + wipe** | [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) (ou [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition)) · [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) · `Gisela` | **mesa num único turno**, a partir do **T7–T9** | **Média-baixa.** Depende de a mesa ter criaturas e de o **seu** board estar pequeno (§5) |
| **R4 · Alt-win — [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun)** | 1 carta, conjurada **duas vezes** | **T11–T13** (7 mana × 2 + o reshuffle 7 cartas abaixo) | **Média.** Não depende de tabuleiro, de vida da mesa nem de aritmética; depende de mana e de sacar de novo |
| **R5 · reserva — [**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign)** | [**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign) + ganho de vida do Helix/[**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) | goldfish: **T7–T8**; jogo real: muito mais tarde | **Enganosa** — ver a advertência na §4.5 |

**Redundância e competição declaradas:**

- **R1 e R2 são complementares, não redundantes.** R1 concentra dano num alvo (combate + Helix +
  [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire)); R2 espalha em todos ao mesmo tempo. O único ponto de disputa é o slot de mana
  do T4, onde [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) briga com uma remoção.
- **R2 e R3 compartilham [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) e `Gisela`.** As duas rotas se pagam com os mesmos 2 slots de
  multiplicador — é o melhor negócio de slot do relatório.
- **R3 e R4 competem pelo mesmo turno de mana grande (T7+)**, e R3 canibaliza um slot de wipe que
  a Fase 5 já ia gastar. R4 não canibaliza nada e não compete com R3 por peça nenhuma.
- **R4 e R5 são substitutos, não complementos.** As duas são alt-win; rodar as duas é diluir.
  Se entrar uma, a outra é reserva.
- **O que NÃO é caminho de vitória:** "atacar com o Phlage". 21 de dano de comandante = 4 ataques
  num único oponente, e é o padrão que o requisito desta fase proíbe nomear como fecho.

---

## 3. Turno de vitória medido

### 3.1 Método

Monte Carlo, **10.000 partidas por configuração**, `on the play`, 3 oponentes de 40 pontos de vida,
mulligan simples (mantém mão de 7 com 2–5 terrenos, até 2 embaralhadas). Modelo de 99 cartas com
37 terrenos, 10 rochas, o pacote de saque da Fase 3 (motores dando +0,5 carta/turno, teto 3 ativos),
os 14 de combustível da Fase 2, Phlage conjurado da zona de comando no T3 e escapado quando houver
5 cartas no cemitério + `{R}{R}{W}{W}`. Dano contabilizado carta a carta, com a ordem de
multiplicadores **escolhida pelo oponente** (é o que a regra manda — §3.3).

Segue o padrão que as fases anteriores usaram: a Fase 2 mediu P(escape) por turno, a Fase 3 mediu
taxa de realização de impulse. Onde meu modelo diverge do da Fase 2, digo (§3.5).

### 3.2 O resultado

**Configuração da Fase 2 (6 pingadores + 4 multiplicadores, sem fecho dedicado):**

| Turno | T8 | T9 | T10 | T11 | T12 | T13 | T14 | T15 |
|---|---|---|---|---|---|---|---|---|
| **1º oponente morto** | 3,6% | 19,1% | **49,5%** | 74,0% | 87,4% | 93,5% | 96,3% | — |
| **MESA INTEIRA morta** | 0,0% | 0,3% | 1,7% | 5,0% | 10,6% | 18,0% | 27,3% | 36,5% |

**Configuração recomendada (4 pingadores + 2 multiplicadores + 3 fechos dedicados):**

| Turno | T8 | T9 | T10 | T11 | T12 | T13 | T14 | T15 |
|---|---|---|---|---|---|---|---|---|
| **1º oponente morto** | 3,3% | 14,9% | 42,4% | **69,5%** | 85,1% | 92,3% | 95,2% | — |
| **MESA INTEIRA morta** | 0,5% | 1,3% | 3,0% | 6,8% | 13,2% | **22,8%** | 33,7% | 44,3% |

**Medianas:** 1º oponente **T10–T11** · mesa inteira **T13–T14**.

**Sensibilidade (o que eu testei para não entregar um número frágil):**

| Variação | Mesa morta até T13 | Mediana da mesa |
|---|---|---|
| Sem nenhum multiplicador | 3,9% | T14+ |
| 2 multiplicadores | 11,4% | T14 |
| 4 multiplicadores | 18,0–19,1% | T13 |
| 8 pingadores em vez de 6 | 23,4% | T13 |
| Sem [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | 19,5% | T13 |
| Taxa de saque otimista (+1,17 mágica/turno no T8+, o topo do que a Fase 3 promete) | 46,1% | T13 |
| **Recomendada (com `Crackle`+[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion)+`Approach`)** | **22,8%** | **T13** |

**Nenhuma variação chega perto do T7.** Duplicar pingadores compra +5 pp; tirar todos os
multiplicadores custa −14 pp; o saque otimista é a alavanca maior de todas e ainda para no T13.
Isso é a assinatura de um problema estrutural, não de seleção de carta.

### 3.3 Por que os multiplicadores rendem menos do que a Fase 2 calculou

Ruling de [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) (2019-10-04) e de `Gisela` (2018-03-16), idênticos no ponto:

> *"If multiple replacement effects would modify how damage would be dealt, the player being dealt
> damage (or the controller of the permanent being dealt damage) chooses the order in which to
> apply those effects."*

**Quem escolhe a ordem é o oponente**, e ele escolhe a pior ordem para você: primeiro os
multiplicativos, depois os aditivos. Consequência concreta no Helix do Phlage (base 3):

| Peças em jogo | Conta ingênua | **Conta correta** |
|---|---|---|
| [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | 5 | 5 |
| `Gisela` | 6 | 6 |
| [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) + `Gisela` | 3+2=5 → ×2 = **10** | ×2 = 6 → +2 = **8** |

A Fase 2 não erra nenhum número isolado, mas a soma dos dois **não** é multiplicativa como a
leitura otimista sugere. Registrado porque isso reaparece no goldfishing.

**Segunda correção de escopo, pedida explicitamente na invocação:** [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) lê *"a **red source**
you control"*. O que fica **de fora**:

| Fica dentro (fonte vermelha) | Fica **fora** |
|---|---|
| `Phlage` (carta vermelha e branca → é fonte vermelha), os 4 pingadores, [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire), [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation), [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder), [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act), [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction), [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate), [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash), [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike), [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) | **Toda a remoção branca** ([**Swords to Plowshares**](https://www.ligamagic.com.br/?view=cards/card&card=Swords+to+Plowshares), [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate), [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket), [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer)) · **[**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) e [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc)**, que causam **perda de vida, não dano** — nenhum multiplicador do jogo os afeta |

Na lista como ela está saindo, isso significa que **~7 das 21 peças de remoção e 2 das 14 fontes
de saque ficam fora do [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)** — cerca de um terço do que o deck faz. A restrição "preferir
`{R}` a `{W}` no empate" que a Fase 2 passou à Fase 5 está certa e **sobe de prioridade**.

### 3.4 A premissa que quebra o pacote de pingadores

A §6.1 da Fase 2 diz: *"Um deck com 21 remoções + 14 loots conjura **3–5 mágicas por turno** no
meio de jogo."* Medido:

| Turno | Mana média | **Mágicas conjuradas/turno** |
|---|---|---|
| T5 | 4,9 | 1,04 |
| T7 | 6,4 | 0,76 |
| T9 | 7,9 | 0,78 |
| T11 | 9,5 | 0,80 |

O gargalo **não é mana** (sobram 7–9 no T9–T11) — é **carta na mão**. Um deck que saca 1 natural
+ ~1,5 dos motores da Fase 3 = 2,5 cartas/turno, das quais ~37% são terrenos, tem ~1,6 não-terreno
por turno, e parte dela é rocha, criatura ou permanente. Mesmo na taxa otimista da Fase 3 (+3
cartas/turno com 3 motores), a medição dá **1,17 mágica/turno**.

Com 1 mágica por turno, um [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) entrega **2 por oponente por turno** (4 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)).
Não 8–10. É essa a diferença entre o pacote da Fase 2 no papel e no relógio.

### 3.5 Divergência declarada com a Fase 2 — turno do escape

| Fonte | P(1º escape até T6) | até T7 | até T8 |
|---|---|---|---|
| Simulação da Fase 2 (§3.1, config D) | **52,4%** | 64,1% | — |
| **Minha simulação** | **17,5%** | 43,0% | 63,9% |

A diferença é de modelagem de combustível: a Fase 2 modela os 10 loots e 4 motores explicitamente
(cada loot despeja 2–3 cartas); eu uso `cemitério += mágicas + 0,6×mágicas`, que é mais conservador
e está acoplado à taxa medida de mágicas/turno da §3.4 — que é baixa. **A verdade está entre os
dois: o 1º escape é um evento de T6–T7.** Não resolvo isso aqui; é uma das coisas que o
goldfishing mede diretamente (§6, métrica 2), e é a métrica que decide qual dos dois modelos está
certo.

---

## 4. Finishers recomendados

**Orçamento:** a Fase 2 reservou **9,5 slots** para o fecho (6 pingadores + 3,5 multiplicadores).
Eu uso **9** e **devolvo 0,5 slot** à Fase 3, que pediu 4 exclusivos e tinha contrapartida só para
3,5. **Não peço nenhum slot novo.**

| Carta | CMC | Como fecha o jogo | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | 4 | +2 por evento de dano de **fonte vermelha** a cada oponente. Num [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe), é +200% sobre a base 2; num pingador de 1, +200%. Leva o Helix de 3 para 5 e o [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) de 7 para 9 | (1) multiplica os 4 pingadores **e** o Helix do comandante — duas rotas com um slot; (2) multiplica [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act)/[**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate)/[**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction), que já estão no orçamento de wipe; (3) 2/4 bloqueia de verdade num deck que precisa de bloqueador. **F7: `{1}{R}{R}{R}` numa base que precisa de `{R}{R}{W}{W}` no T6** | não | **R$ 6,98** (23/09) |
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | 7 | **Dobra todo dano a oponentes, inclusive combate** — e **previne metade do dano que você recebe**. É a única peça que faz [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + wipe ser sobrevivível (§5) | (1) dobra as duas rotas de dano ao mesmo tempo (agregado e focado); (2) a metade defensiva é sobrevivência real num deck que precisa chegar ao T11; (3) 5/5 voadora com iniciativa cobre o céu, que é o buraco de RW; (4) **é a candidata registrada em `decisions.md`** — entra, não é dispensada. **F7: CMC 7, a carta mais cara em mana do deck** | não | **R$ 15,00** (23/09) |
| [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) | 3 | 2 a **cada** oponente por instantâneo/feitiço — 4 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), 4 com `Gisela`, 6 com os dois | (1) a maior taxa por gatilho do pool; (2) corpo 2/2 que bloqueia; (3) o único pingador cujo dano-base sobrevive a não ter multiplicador | não | **R$ 5,40** (23/09) |
| [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) | 2 | `{T}`: 1 a cada oponente; **destapa a cada instantâneo/feitiço** → N+1 pings por turno | (1) **0/3 defender** — o bloqueador barato que um deck de controle sem board precisa no T2; (2) é o único pingador que dispara **sem** você conjurar nada (o primeiro tap do turno); (3) R$ 0,55 | não | **R$ 0,55** (23/09) |
| [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) | 2 | 1 a cada oponente por instantâneo/feitiço | (1) mesmo gatilho do [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) por **1/30 do preço**; (2) corpo 2/1 que ataca se a mesa esvaziar; (3) redundância — o pacote precisa de 4 para P(≥1 visto até T8) = 46% | não | **R$ 0,33** (23/09) |
| [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer) | 2 | 1 a cada oponente por mágica **não-criatura** — pega também as 10 rochas, os artefatos de saque e os encantamentos | (1) o gatilho mais largo do pacote (não-criatura, não instantâneo/feitiço); (2) 2/1; (3) R$ 0,99 | não | **R$ 0,99** (23/09) |
| [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) | 3 | *"Whenever a creature is dealt damage, this enchantment deals that much damage to that creature's controller."* Converte **as 21 remoções e os 4 wipes** em dano na cara. Com [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) (13 a cada criatura), mata a mesa **num turno** | (1) transforma cada queima em criatura num Lightning Helix na cara do dono — o deck tem 21 delas; (2) o dano é de **[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion)**, um permanente **vermelho** → [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) soma +2 **por instância**; (3) fecha com um wipe que a Fase 5 já ia comprar. **F7 grave: é simétrico — ver §5** | não | **`a cotar`** |
| [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) | X | `{X}{X}{X}{R}{R}`: **cinco vezes X** de dano a cada um de até X alvos. X=3 custa 11 mana e dá **15 a cada um dos 3 oponentes** (ruling 2021-04-16 confirma a conta) | (1) é o único dreno simultâneo do pool que escala com o mana tardio que este deck acumula (7–9 sobrando no T9); (2) com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) vira 17 cada, com `Gisela` 32 cada; (3) o deck tem 10 rochas + 37 terrenos — o mana existe e hoje não tem onde ir | não | **`a cotar`** |
| [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) | 7 | Conjurada **duas vezes** = você vence. Ignora inteiramente a aritmética de 3×40 | (1) é o **único** fecho do pool que não depende de tabuleiro, de vida da mesa nem de multiplicador — é a resposta literal a "controla 12 turnos e não fecha"; (2) ganha 7 de vida por conjura, o que sustenta o plano de atrito até lá; (3) volta 7 cartas abaixo do topo, e as 14 fontes de saque da Fase 3 são o que a traz de volta. **F7: 14 mana no total; a 2ª tem de ser conjurada da mão (ruling 2017-04-18)** | não | **`a cotar`** |

**Total cotado dos 9 slots: R$ 29,25** (6 cartas) **+ 3 `a cotar`**.

### 4.1 Fechos que entram sem gastar slot novo

| Carta | Onde já está no orçamento | O que acrescenta ao fecho | Preço |
|---|---|---|---|
| [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | pacote de saque da Fase 3 (§4.1) | **7 de dano em qualquer alvo por upkeep**, 9 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), 14 com `Gisela`. É o **maior relógio focado do deck** e não passa por combate. Tirá-lo custa −9 pp na morte do 1º oponente até T10 (medido) | `a cotar` |
| [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) | um dos 3–4 slots de wipe da Fase 2 (§9.4) | 13 a cada criatura. Sozinho é wipe; com [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) é **o fecho da mesa** | **R$ 12,10** (23/09) |
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | um dos slots de wipe (§9.4) | X dividido entre qualquer número de alvos; **flashback `{R}{R}` descartando X**. Com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), cada alvo leva +2 — dividir 3 entre 3 oponentes vira 5/5/5. É wipe, alcance e combustível num card **a R$ 0,07** | **R$ 0,07** (23/09) |
| [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition) | slot de wipe (alternativa ao [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate)) | Phlage 6/6 causa **6 a cada outra criatura e a cada oponente** — 8 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) (ruling 2015-06-22: a **criatura** é a fonte, e ela é vermelha). Com [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion), um oponente com 3 criaturas leva 8 + 24 = **32**. **F7: mata os seus 4 pingadores** — é botão de reset, não peça de motor | `a cotar` |
| [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) · [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) · [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) · [**Syr Carah**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | pacotes de combustível e saque das Fases 2 e 3 | dano/perda de vida incidental a cada oponente, todo turno. Somados valem ~4 por oponente por turno no meio de jogo. **`Monument` e `Palantír` causam perda de vida — não multiplicam** | `a cotar` |

### 4.2 O que sai do pacote da Fase 2 — com ficha (regra 4)

Nenhuma destas está no deck (não existe `deck.md`): são **dispensas de pool com ficha**, não cortes
de carta montada. Mas a ficha vem completa, porque é o que impede o vai-e-vem.

**[**Solphim, Mayhem Dominus**](https://www.ligamagic.com.br/?view=cards/card&card=Solphim%2C+Mayhem+Dominus)** `{2}{R}{R}` 5/4 — **FORA por preço**
**F1** dobra dano **não-combate** de qualquer fonte sua a oponentes (inclui a remoção **branca**,
o que [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) não faz); `{1}{R/P}{R/P}`, descarte 2: contador de indestrutível — que é
**combustível de escape + proteção contra o próprio wipe**. **F2** 5/4, corpo real. **F3** Creature,
lendária (liga [**Jaya's Immolating Inferno**](https://www.ligamagic.com.br/?view=cards/card&card=Jaya%27s+Immolating+Inferno)). **F4** recebe contadores e anthems. **F5** multiplica
todo o deck. **F6** CMC 4, `{R}{R}` — **melhor custo de cor que o [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)**. **F7** não dobra
combate.
**Veredito:** funcionalmente é a **melhor** das quatro para este deck — e custa **R$ 118,90 =
59,5% do teto de R$ 200**. Funções e quem as cobre: *multiplicar dano não-combate* → [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)
(parcial: só fonte vermelha) e `Gisela` (integral, 3 manas mais cara); *descarte como combustível*
→ os 10 loots da Fase 2; *corpo 5/4* → `Gisela` 5/5 e `Phlage` 6/6; *indestrutível* → [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm)
(**R$ 33,17** — sinal à Fase 5: também é caro). **Função descoberta declarada:** multiplicar a
remoção **branca**, que [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) não pega. Custo aceito.

**[**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation)** `{3}{R}{R}{R}` — **FORA por preço e por curva**
**F1** triplica todo dano de fonte sua. **F2** sem corpo. **F3** Enchantment. **F4/F5** nada.
**F6** **CMC 6 com `{R}{R}{R}`** numa base que precisa de `{R}{R}{W}{W}` — os dois requisitos de
cor brigam pelo mesmo turno. **F7** não faz nada sozinha; precisa de um pingador ou do Phlage vivo.
**Veredito:** **R$ 54,54 = 27,3% do teto**, por uma carta que só liga se outra já estiver ligada.
O aviso da invocação estava certo. Funções cobertas por [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) + `Gisela`. **Nada descoberto.**

**[**Kessig Flamebreather**](https://www.ligamagic.com.br/?view=cards/card&card=Kessig+Flamebreather)** `{1}{R}` 1/3 — **FORA por preço**
**F1** 1 a cada oponente por mágica não-criatura. **F2** 1/3, sobrevive a queima de 2. **F3**
Creature. **F4** recebe anthems (o deck não tem). **F5** nada. **F6** T2. **F7** nenhum.
**Veredito:** **R$ 24,79 = 12,4% do teto** por um gatilho **idêntico** ao do [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer)
(**R$ 0,99**). Funções: *gatilho de não-criatura* → [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer); *corpo que sobrevive a 2*
→ [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) 0/3 e [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) 0/4, se a Fase 5 quiser a parede. **Nada
descoberto.** Dispensa por taxa de preço, não por qualidade.

**[**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field)** `{1}{R}` 0/4 defender — **corte condicionado, devolvido à Fase 5**
**F1** 1 a cada oponente por instantâneo/feitiço. **F2** **0/4 com defender** — o melhor bloqueador
barato do pool inteiro; para agressão de T2–T5 melhor que qualquer 2/2. **F3** Creature. **F4**
recebe anthems (inúteis, defender). **F5** nada. **F6** T2. **F7** **não ataca** — não converte
board em dano nem carrega equipamento.
**Veredito:** **R$ 9,90 = 5,0% do teto** por um gatilho que o [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) entrega a
**R$ 0,33**. O que a diferencia é **só o corpo 0/4**, e corpo/bloqueio **não é da minha
especialidade** — é decisão da Fase 5 sobre a dor 3 e a política da coroa ([**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer)/`Court
of Ire` fazem a mesa atacar você). **Eu não a corto.** Devolvo ao orquestrador: se a Fase 5 quiser
a parede, ela entra **no slot de bloqueador**, não no slot de wincon, e o [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith)
permanece como o 4º pingador.

**[**Wild Guess**](https://www.ligamagic.com.br/?view=cards/card&card=Wild+Guess)** `{R}{R}` — **o meio slot que eu não precisei tomar**
A Fase 3 já a nomeou como corte condicionado, com ficha completa (§7 do `03-draw.md`) e todas as
funções cobertas. Eu uso 9 dos 9,5 slots reservados; **ela não precisa sair por minha causa**.
Registro apenas que, se a Fase 6 precisar de meio slot, esta é a fila.

### 4.3 A inversão que as cotações de hoje forçaram

A Fase 2 ranqueou [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) em 1º e `Gisela` em 2º, e pôs `Solphim` e [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) no
mesmo pool como se fossem opções reais. **Com as cotações de 2026-09-23, o pool de multiplicadores
de RW dentro do Commander 200 tem exatamente dois nomes:**

| Carta | Preço (23/09) | % do teto | Status |
|---|---|---|---|
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | R$ 6,98 | 3,5% | **entra** |
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | R$ 15,00 | 7,5% | **entra** |
| [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) | R$ 54,54 | 27,3% | fora |
| [**Solphim, Mayhem Dominus**](https://www.ligamagic.com.br/?view=cards/card&card=Solphim%2C+Mayhem+Dominus) | R$ 118,90 | 59,5% | fora |

A consequência é direta e vai para o `report.md`: o alvo de "3–4 multiplicadores" da Fase 2
**não existe nesta faixa de preço**. São 2, e a simulação com 2 multiplicadores dá **11,4% de
mesa morta até o T13** contra 18–19% com 4. **Metade do plano de multiplicação do deck é um
plano que o orçamento não comporta**, e isso precisa ser dito ao usuário antes de ele comprar.

`Gisela` sobe de "nº 2 com ressalva de CMC 7" para **peça obrigatória**: além de dobrar, ela é a
única carta do pool que **previne metade do dano que você recebe**, e é isso que torna a Rota 3
sobrevivível (§5). O registro de `decisions.md` que a colocou como primeira candidata a finisher
está confirmado.

### 4.4 Varredura da coleção (regra 7) — feita antes de qualquer busca

`bin/mtgdb collection -list` (116 cartas) + as 69 não-básicas de `lista.txt`, filtradas por fecho.

| Carta | Origem | Ficha resumida | Por que **não** cobre a função de fecho |
|---|---|---|---|
| **[**Darksteel Reactor**](https://www.ligamagic.com.br/?view=cards/card&card=Darksteel+Reactor)** (R$ 5,44) | caixa | **F1** indestrutível; 1 contador de carga por upkeep; **com 20 contadores você vence o jogo**. **F3** Artifact. **F6** CMC 4. **F7** usa o upkeep | **É literalmente um alt-win na caixa, a custo zero, e mesmo assim não serve.** Sem proliferate ele leva **20 turnos** — o deck já mata a mesa no T13–T14 por dano. A caixa tem [**Courage in Crisis**](https://www.ligamagic.com.br/?view=cards/card&card=Courage+in+Crisis) e [**Thrummingbird**](https://www.ligamagic.com.br/?view=cards/card&card=Thrummingbird) como proliferate, mas o [**Thrummingbird**](https://www.ligamagic.com.br/?view=cards/card&card=Thrummingbird) é **azul** (fora da identidade) e o [**Courage in Crisis**](https://www.ligamagic.com.br/?view=cards/card&card=Courage+in+Crisis) prolifera uma vez. Dispensa por **velocidade**, com a carta nomeada — não por "carta ruim" |
| **[**Lux Artillery**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Artillery)** (R$ 1,22) | caixa | **F1** 10 de dano a cada oponente no end step **com 30+ contadores** entre artefatos e criaturas suas; dá sunburst a criaturas-artefato | 30 contadores exige um deck de contadores; este tem 4 criaturas de 0–2 de poder e nenhuma fonte de contador. **0 pontos de sinergia** (regra 3). Dispensa por eixo |
| **[**Titan Forge**](https://www.ligamagic.com.br/?view=cards/card&card=Titan+Forge)** (R$ 0,49) | caixa | **F1** `{3},{T}`: contador; remover 3: ficha 9/9. **F3** Artifact | 9 manas e 3 turnos para um 9/9 **sem evasão** num deck que não ataca. É rota de combate, que é o que o requisito proíbe. 1 ponto |
| **[**Smaug, the Great Calamity // Spew Flame**](https://www.ligamagic.com.br/?view=cards/card&card=Smaug%2C+the+Great+Calamity+%2F%2F+Spew+Flame)** | caixa | **F1** 5/5 voadora CMC 7 // Adventure: 5 de dano numa criatura | O lado Adventure é remoção (Fase 5); o corpo é CMC 7 sem gatilho de dano a jogador. Como **fecho**, 0 pontos |
| **[**Seismic Wave**](https://www.ligamagic.com.br/?view=cards/card&card=Seismic+Wave)**, **[**Lightning Volley**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Volley)**, **[**Radiating Lightning**](https://www.ligamagic.com.br/?view=cards/card&card=Radiating+Lightning)**, **[**Searing Barrage**](https://www.ligamagic.com.br/?view=cards/card&card=Searing+Barrage)**, **[**Stonefury**](https://www.ligamagic.com.br/?view=cards/card&card=Stonefury)**, **[**Molten Blast**](https://www.ligamagic.com.br/?view=cards/card&card=Molten+Blast)** | caixa | queima de Limited, 2–4 de dano, CMC 3–5 | Todas disparam os pingadores e enchem o cemitério — **elas contam nas 21 remoções da Fase 5**, e por isso não as dispenso lá. Como **fecho** não fazem nada: 2–4 num alvo não muda um total de 40. Dispensa por função, dentro desta fase apenas |
| **[**Punishing Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Punishing+Fire)**, **[**Skullcrack**](https://www.ligamagic.com.br/?view=cards/card&card=Skullcrack)** | caixa | já dispensadas com ficha na §8.2 da Fase 2 | [**Skullcrack**](https://www.ligamagic.com.br/?view=cards/card&card=Skullcrack) desliga os 3 de vida do **seu próprio** Helix (cláusula simétrica). [**Punishing Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Punishing+Fire) tira carta do cemitério a cada retorno. Mantenho as duas dispensas |
| **[**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold)** | **lista** | **F1** exila o topo e deixa jogar sempre que uma mágica sua fere um **jogador**; `{T}`: 1 de dano, que dispara ela mesma | **Aproveitada** — mas como **motor de cartas**, não como fecho (já contada pela Fase 3). Ela multiplica a Rota 2: cada ping em jogador vira carta, e cada carta vira mágica, e cada mágica vira ping. **Custo zero** |
| **[**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate)**, **[**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb)**, **[**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone)** | caixa | wipes | Aproveitados pela Fase 5. Relevante para mim: **[**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) ganha vida por criatura** — é a peça que liga a Rota 5 ([**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign)) |

**Conclusão da varredura:** a caixa **não tem nenhum fecho**. Tem um alt-win ([**Darksteel Reactor**](https://www.ligamagic.com.br/?view=cards/card&card=Darksteel+Reactor))
que é lento demais e uma bomba de contadores ([**Lux Artillery**](https://www.ligamagic.com.br/?view=cards/card&card=Lux+Artillery)) que não tem deck para ligar. Os 9
slots de fecho são **integralmente de compra**, e 3 deles estão `a cotar`. Isso está declarado.

### 4.5 Advertência sobre [**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign) (Rota 5, reserva)

`{4}{W}{W}` 4/6 vigilância, vínculo com a vida: *"At the beginning of your upkeep, if you have 40
or more life, you win the game."* (ruling 2015-08-25: checa no início do upkeep **e** de novo ao
resolver).

**Ela ganha o goldfishing e pode não ganhar a partida.** Você começa com 40. Num goldfishing sem
oponentes, qualquer ganho de vida do Helix já a liga — ela "vence no T7–T8" e o número fica lindo.
Numa mesa de 4, você toma dano e a conta muda. Os pontos de sinergia são reais (o Helix dá 3 de
vida a **cada** entrada e **cada** ataque; [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) ganha em bloco; o corpo 4/6 com vigilância
é um bloqueador de verdade), mas **eu não a coloco nos titulares** justamente porque ela
distorceria a métrica que esta fase existe para medir. Fica como reserva nomeada, `a cotar`, e
o goldfishing **não deve contá-la como vitória** (§6, regra de registro).

---

## 5. Combos

| Combo | Peças | Resultado | Confirmado por rulings? |
|---|---|---|---|
| **[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act)** | 2 cartas · **1 a 2 manas** quando a mesa está cheia ([**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) custa `{1}` a menos por criatura em jogo) | [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) dá 13 a cada criatura. [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) dispara **uma vez por criatura**: cada oponente leva 13 × (nº de criaturas dele). Mesa com 3 criaturas por jogador = **39 a cada oponente**. Com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) é fonte **vermelha** → +2 por instância = **45 cada**. **Mata a mesa num turno** | **Parcialmente.** [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) e [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) têm rulings; **[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) não tem nenhum ruling oficial** (`bin/mtgdb rulings` retorna 0). A leitura vem do texto: o gatilho é *"whenever a creature is dealt damage"*, sem restrição de controlador, e a fonte do segundo dano é o próprio [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) |
| **[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition) (mirando `Phlage`)** | 2 cartas + Phlage em campo · 5 manas | `Phlage` 6/6 causa 6 a cada outra criatura **e** a cada oponente. Oponente com 3 criaturas: 6 (direto) + 18 ([**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion)) = **24**; com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), 8 + 24 = **32** | Sim — ruling de [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition) (2015-06-22) confirma que **a criatura é a fonte**, o que faz [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) aplicar |
| **`Gisela` + [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) X=3** | 2 cartas · 11 manas | 15 × 2 = **30 a cada um dos 3 oponentes**, num único feitiço. Com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) também: 32 cada | Sim — ruling de [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) (2021-04-16) confirma a conta de custo e dano; ruling de `Gisela` (2018-03-16) confirma a ordem de replacement |

### ⚠ O atrito que torna a Rota 3 perigosa — e a carta que o resolve

**[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) é simétrico.** *"that creature's controller"* inclui **você**. Com
[**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) e 3 criaturas suas em campo, **você leva 39** e morre junto com a mesa.

Três mitigações, em ordem de qualidade:

1. **[**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight)** — *"If a source would deal damage to you or a permanent you
   control, prevent half that damage, rounded up."* Os 13 de [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) viram **6** em você.
   Com 3 criaturas suas, você leva 18 em vez de 39. **É por isso que `Gisela` é obrigatória e não
   opcional.**
2. **Jogar o wipe com o seu board vazio.** É o estado natural de um deck de controle com 4
   criaturas em 99, e é a razão de eu ter cortado o pacote de pingadores de 6 para 4: menos corpos
   seus = menos dano refletido em você.
3. **Não é mitigável por [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm).** Indestrutível **não previne dano** — as criaturas
   sobrevivem e o [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) dispara igual. Registro isso porque a §7 da Fase 2 trata o
   [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) como a peça que "faz o seu wipe ser assimétrico", e contra [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) ele
   não faz nada. ([**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) custa **R$ 33,17 = 16,6% do teto** — sinal separado à Fase 5.)

**Regra 11 e regra 9.** Isto é um fecho de duas cartas numa mesa que o briefing descreve como
"Commander 200 competitivo". Não é lock, não é infinito, não impede ninguém de jogar, e as duas
peças são individualmente jogáveis ([**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) é um dos 4 wipes que a Fase 5 ia comprar de
qualquer jeito; [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) transforma 21 remoções em dano na cara). **Mas é uma decisão de
power level, e ela é do usuário** — vai ao `report.md` como pergunta, não como fato consumado.

---

## 6. Protocolo de goldfishing

**Este deck nunca foi goldfished. A v1 morreu sem isso e foi reprovada em mesa.** O protocolo
abaixo está pronto para rodar em proxy, antes de qualquer compra.

### 6.1 Execução

1. Embaralhe, compre 7. **Mulligan como numa partida real** (Londres: embaralhe, compre 7, ponha
   N no fundo). Anote quantos mulligans.
2. Jogue **sozinho, turnos consecutivos, sem oponente**: terreno, ramp, comandante na curva,
   desenvolva. Sem oponente não há alvo para remoção — **aponte todo Helix e toda queima "na cara"
   e some**, é assim que se mede o relógio.
3. Pare no **T14** ou quando a projeção de letal estiver completa, o que vier primeiro.
4. **10 partidas.** (O padrão do pipeline é 5+; peço 10 porque a variância aqui é alta — a
   diferença entre ter e não ter multiplicador muda o resultado inteiro.)

### 6.2 O que registrar, por partida

| # | Campo | Por que |
|---|---|---|
| 1 | Mulligans | separa mão morta de azar |
| 2 | **Turno em que o Phlage saiu da zona de comando** | tem de ser T3 em quase toda partida |
| 3 | **Turno do 1º escape** (`{R}{R}{W}{W}` + 5 no cemitério) | é a métrica que decide a divergência da §3.5 |
| 4 | Turno do 2º escape | mede se os **motores** de combustível estão funcionando ou se só os one-shots estão |
| 5 | **Turno em que o 1º multiplicador entrou** ([**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)/`Gisela`) — ou "nunca" | é a peça de que tudo depende |
| 6 | **Turno em que o 1º pingador entrou** | idem |
| 7 | **Mágicas conjuradas por turno, do T6 ao T12** | testa diretamente a premissa quebrada da §3.4 |
| 8 | **Turno da vitória projetada — 1º oponente (40 de vida)** | métrica principal |
| 9 | **Turno da vitória projetada — 3 oponentes de 40** | métrica do pipeline |
| 10 | **Turno em que o deck "passa a ameaçar"** — defina como: o turno em que a mesa vê ≥8 de dano por turno saindo do seu lado | §7 |
| 11 | Travas de mana: turnos com ≤3 terrenos; turnos em que faltou `{W}{W}` ou `{R}{R}` especificamente | devolve à Fase 6 |
| 12 | Mãos mortas: mãos sem **nenhum permanente jogável** nos 3 primeiros turnos | devolve à Fase 3 |
| 13 | Se [**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign) estiver na lista: **não conte a vitória dela** — anote separado | §4.5 |

### 6.3 Quais números derrubam quais decisões

| Se acontecer em… | …o número é | Decisão que cai | Para onde volta |
|---|---|---|---|
| **≥6 de 10** | vitória projetada sobre **1 oponente** depois do **T12** | O pacote de fecho falhou. Não adianta trocar peça: é o eixo | **Orquestrador** — decisão de eixo, com o usuário |
| **≥5 de 10** | **1º escape depois do T8** | Os 14 slots de combustível da Fase 2 estão mal dimensionados **ou** a base não faz `{R}{R}{W}{W}` | **Fase 2** (combustível) e **Fase 4** (cor) — o registro 11 diz qual |
| **≥7 de 10** | nenhum multiplicador em campo até o **T9** | Com 2 multiplicadores em 99, o plano de multiplicação é ficção. Corte para 1 ([**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell)) e ponha o slot numa 2ª coroa/`Court` | **Fase 7** (eu), com o slot devolvido |
| **≥6 de 10** | **< 1,5 mágica por turno** do T8 ao T12 | Confirma a §3.4: os pingadores não pagam 4 slots. Corte para 2 ([**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) + [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist)) e ponha 2 slots em remoção ou em [**Court of Embereth**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Embereth) | **Fase 7** (eu) e **Fase 5** |
| **≥3 de 10** | mão sem **nenhum permanente jogável** até o T3 | Curva/densidade de saque | **Fase 3** |
| **≥3 de 10** | ≤3 terrenos no T4, **ou** falta específica de `{W}{W}`/`{R}{R}` no turno do escape | Manabase | **Fase 6 · manabase-engineer** |
| **≥4 de 10** | Phlage **não** sai no T3 | Rampa e fontes de cor | **Fase 4** |
| **≤2 de 10** | [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) foi útil (no goldfishing não há criaturas inimigas) | **Esperado, não é falha.** [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) é a única peça que o goldfishing **não consegue medir** — ela precisa de mesa. Avalie-a em jogo real, não aqui | — |

### 6.4 Métrica de saúde honesta para este eixo

A meta do pipeline — **vencer até o T7** — foi escrita para decks que convertem tabuleiro em dano.
Este não é um. Pela simulação, **nenhuma configuração dentro do orçamento chega ao T7**, e fingir
que chega seria repetir o erro da v1.

**Proponho ao orquestrador duas linhas de corte, não uma:**

| Métrica | Alvo | Justificativa |
|---|---|---|
| **1º oponente morto** | **≤ T10 em 6 de 10 partidas** | é o que a simulação diz ser alcançável (49,5% até T10) e é o que muda a política da mesa |
| **Mesa inteira (projeção de goldfish)** | **≤ T13 em 5 de 10** | é a mediana medida; abaixo disso o deck não é um deck de Commander competitivo, é um deck que sobrevive |

Se o usuário não aceitar essas linhas, **a conclusão correta não é ajustar cartas — é rever o
eixo**, e isso é decisão dele (regra 9), com os números desta seção à vista.

---

## 7. A pergunta da paciência

O usuário reprovou a v1 em parte por **"jeito de jogar"**. A pergunta certa não é "quando o deck
mata", é **"quando o deck deixa de ser espectador"**. Medido:

| Turno | O que o deck está fazendo | Probabilidade |
|---|---|---|
| **T1–T2** | Terreno, rocha, no máximo uma remoção de 1–2 manas. **Nada acontece na sua direção** | — |
| **T3** | **`Phlage` da zona de comando: 3 de dano + 3 de vida, em toda partida, sem depender de compra.** Ele se sacrifica e vai para o cemitério, onde quer estar | **100%** |
| **T4–T5** | Remoção, loot, rocha. Você está reagindo. O tabuleiro é do oponente | — |
| **T6–T7** | **O deck passa a ameaçar.** 1º escape: um 6/6 em campo que dá 9 por turno num alvo (6 de combate + 3 de Helix). É o primeiro momento em que a mesa precisa responder a você | **17,5%–43%** até T7 (meu modelo) · **52%–64%** (modelo da Fase 2) |
| **T7–T8** | O dano agregado fica visível: ≥8 por turno **a cada** oponente | mediana **T8** |
| **T10–T11** | **1º oponente morre** | mediana |
| **T13–T14** | **Mesa morre** | mediana |

**A resposta honesta, em uma frase:** o deck **nunca tem um turno morto** (o T3 é garantido em
100% das partidas, o que a v1 não tinha), **passa a ameaçar no T6–T7**, e **fecha no T13**. São
**6 a 7 turnos** entre "estou claramente na frente" e "o jogo acabou".

Isso é o perfil clássico de um deck de controle, e é uma **escolha de gosto, não um defeito**.
Mas é exatamente o padrão que produz a frase *"monta board, ataca, não converte em vitória"* — só
que com remoção no lugar do board. **Se o incômodo do usuário foi de ritmo, e não de potência,
este eixo troca um tipo de espera por outro.** Ele precisa ver esta seção antes de comprar as
99 cartas, e não depois.

**Contrapeso justo, para não pintar só o lado ruim:**
- Diferente da v1, **você está fazendo alguma coisa todo turno desde o T1** — remoção é interação,
  e interação é participação na mesa.
- Diferente da v1, **um board wipe não encerra o seu jogo**: o comandante volta do cemitério pelo
  mesmo preço na quinta vez, e o wipe **enche** o seu cemitério.
- O [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) dá o que a v1 não tinha: um botão que aponta 7 (9 com [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), 14 com
  `Gisela`) em qualquer alvo, todo upkeep, sem passar por combate.

---

## 8. Orçamento de slots — o que entra, o que sai, quanto custa

**Reservado pela Fase 2 para o fecho: 9,5 slots.** **Usado: 9.** **Devolvido: 0,5.**

| # | Entra | Slot vem de | Preço |
|---|---|---|---|
| 1 | [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | multiplicador 1 (reservado) | R$ 6,98 |
| 2 | [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | multiplicador 2 (reservado) | R$ 15,00 |
| 3 | [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) | pingador 1 (reservado) | R$ 5,40 |
| 4 | [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) | pingador 2 (reservado) | R$ 0,55 |
| 5 | [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) | pingador 3 (reservado) | R$ 0,33 |
| 6 | [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer) | pingador 4 (reservado) | R$ 0,99 |
| 7 | [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) | pingador 5 (reservado) — **trocado por fecho dedicado** | `a cotar` |
| 8 | [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) | pingador 6 (reservado) — idem | `a cotar` |
| 9 | [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) | multiplicador 3 (reservado) — idem | `a cotar` |
| — | **0,5 slot devolvido** | multiplicador 4 (era 3,5) → vai para a Fase 3 | — |

**Sem slot novo (dentro de orçamentos de outras fases):** [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) (R$ 12,10) e
[**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) (R$ 0,07) nos slots de wipe · [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition) (`a cotar`) como alternativa ao
[**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) · [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (`a cotar`) no pacote de saque da Fase 3.

**Custo cotado dos 9 slots: R$ 29,25** — mas **3 dos 9 estão `a cotar`**, e são os três fechos
dedicados. **O custo do fecho é desconhecido, não baixo.**

### Pendência dura de cotação (bloqueia a Fase 6)

[**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) · [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) · [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) · [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition) ·
[**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) · [**Felidar Sovereign**](https://www.ligamagic.com.br/?view=cards/card&card=Felidar+Sovereign) · [**Aetherflux Reservoir**](https://www.ligamagic.com.br/?view=cards/card&card=Aetherflux+Reservoir) · [**Comet Storm**](https://www.ligamagic.com.br/?view=cards/card&card=Comet+Storm) ·
[**Jaya's Immolating Inferno**](https://www.ligamagic.com.br/?view=cards/card&card=Jaya%27s+Immolating+Inferno) · [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) · [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) · [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold).

Precedente registrado hoje: `Solphim` apareceu a **R$ 118,90** e [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) a **R$ 54,54**
— duas cartas que a Fase 2 tratava como candidatas normais. **Não assuma viabilidade de nenhuma
carta `a cotar`.** Fila de substituição se os fechos estourarem, em ordem:
[**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) (R$ 0,07) → [**Chain Reaction**](https://www.ligamagic.com.br/?view=cards/card&card=Chain+Reaction) (R$ 0,90) → [**Comet Storm**](https://www.ligamagic.com.br/?view=cards/card&card=Comet+Storm)/[**Jaya's Immolating Inferno**](https://www.ligamagic.com.br/?view=cards/card&card=Jaya%27s+Immolating+Inferno)
(X-spells, `a cotar`) → [**Darksteel Reactor**](https://www.ligamagic.com.br/?view=cards/card&card=Darksteel+Reactor) (R$ 5,44, **caixa**, com a ressalva de velocidade da §4.4).

---

## 9. Ajustes pós-teste

*(A preencher depois do goldfishing. A tabela de gatilhos está na §6.3 — cada linha já diz qual
troca ou qual devolução ela dispara.)*

| Sai | Entra | Motivo |
|---|---|---|
| — | — | aguardando as 10 partidas |

---

## 10. Sinais cruzados

| Destino | Sinal |
|---|---|
| **Fase 2 · tema** | (a) A premissa de **3–5 mágicas/turno** da §6.1 está medida em **0,8–1,2** (§3.4) — o dimensionamento do pacote de pingadores depende dela e cai junto. (b) O alvo de "3–4 multiplicadores" **não existe no Commander 200**: são 2 (§4.3). (c) A divergência de P(escape) entre os dois modelos (§3.5) é a métrica nº 3 do goldfishing |
| **Fase 4 · ramp** | O fecho precisa de **11 manas** para [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) X=3 e **7** para [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun), **duas vezes**. Isso muda o alvo de rampa: não é só chegar a `{R}{R}{W}{W}` no T6, é ter **9–11 manas no T9–T11**. [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) e [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) (Treasures) ganham peso; rocha que só destrava o escape e para por aí, perde |
| **Fase 5 · interação** | (a) [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) (**R$ 12,10**) sobe de "candidata a wipe" para **peça de fecho** — é a parceira do [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion). (b) [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) a **R$ 0,07** é wipe + fecho + combustível: recomendo travar o slot. (c) **[**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) custa R$ 33,17 = 16,6% do teto** — e **não protege contra [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion)** (indestrutível não previne dano, §5). (d) Devolvo [**Electrostatic Field**](https://www.ligamagic.com.br/?view=cards/card&card=Electrostatic+Field) como **corte condicionado**: R$ 9,90 por um gatilho de R$ 0,33; o que a justifica é o corpo **0/4 defender**, e corpo é decisão de vocês, não minha |
| **Fase 6 · manabase** | [**Torbran**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) é `{1}{R}{R}{R}` e [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) é `{X}{X}{X}{R}{R}` — o deck tem **dois requisitos de cor pesada em direções opostas** (`{R}{R}{R}` do Torbran contra `{W}{W}` do escape). Isso é mais exigente do que a Fase 2 dimensionou. Além disso: [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) pede **7 manas duas vezes**, o que empurra a base para cima, não para baixo |
| **Orquestrador** | (1) **A meta de T7 é inalcançável neste eixo** — proponho duas linhas de corte na §6.4 e peço decisão do usuário. (2) **A pergunta da paciência (§7) precisa ir ao `report.md` como pergunta**, não como nota de rodapé: o deck ameaça no T6–T7 e fecha no T13. (3) **[**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) + [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) é um fecho de 2 cartas** — power level é decisão do usuário (regra 9, §5). (4) **3 dos 9 slots de fecho estão `a cotar`** e o precedente do `Solphim` (R$ 118,90) mostra que o risco de preço é real |

---

## 11. Riscos declarados

1. **O modelo é um modelo.** Ele não conta: política de mesa, oponentes se matando, o Phlage
   voltando ao cemitério a cada morte depois do 1º escape, remoção adversária nas suas peças, e
   [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) (que precisa de criaturas inimigas e por isso é **invisível ao goldfishing**).
   Erra para os dois lados; o goldfishing é o que o corrige.
2. **Ódio de cemitério desliga a Rota 1 e metade da Rota 2.** [**Rest in Peace**](https://www.ligamagic.com.br/?view=cards/card&card=Rest+in+Peace) apaga o escape, o
   `Palantír`, o [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) e o [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse). As 2 respostas a encantamento que
   a Fase 2 pediu e a Fase 3 tornou obrigatórias **são requisito do fecho também**.
3. **Remoção na `Gisela` custa a Rota 3.** Sem a metade defensiva dela, [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) +
   [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) pode matar você junto. É preciso saber contar o próprio board antes de puxar
   o gatilho — é a única jogada deste deck que tem risco de auto-kill.
4. **[**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) é telegrafada.** A primeira conjura anuncia a segunda, e a mesa
   tem ~4 turnos para te matar antes. É o preço de um alt-win em RW, que não tem counterspell para
   proteger nada.
5. **Nenhum fecho está na caixa.** Os 9 slots são compra integral, 3 deles `a cotar`.
