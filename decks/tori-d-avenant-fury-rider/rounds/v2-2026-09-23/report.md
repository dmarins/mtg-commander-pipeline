# Relatório da rodada v2 — Phlage, Titan of Fire's Fury (revisão de 2026-09-24)

> **Estado:** reconsolidado em 2026-09-24 e **aguardando sua validação**. Nada foi comprado, e
> `deck.md`/`decisions.md` ainda não têm as cartas: só entram quando você aprovar (regra 9).
> Este é o único arquivo que você precisa ler. Os relatórios por fase ficam ao lado para consulta;
> a revisão de cada fase está na seção `Revisão 2026-09-24` no fim do arquivo dela.
>
> **O que mudou desde a versão de 2026-09-23:** você achou o deck fraco e pediu para olhar o
> EDHREC. Ele mostrou que a v2 estava usando o Phlage pela metade (§1). As 7 fases revisaram o deck
> em torno disso: **19 trocas**, custo de R$ 176,71 para **R$ 166,98**.
>
> **Próximo passo sugerido:** responder às duas perguntas da §0, imprimir a lista da §9 em proxy e
> jogar antes de comprar. Trocar agora não custa nada; depois da compra, custa.

---

## 0. Duas decisões suas antes de validar

### 0.1 O ritmo: o deck ficou mais ativo, mas não fecha mais cedo

A Fase 7 simulou a v2 e a lista nova com o mesmo código, na mesma régua (goldfish, modo mesa,
3 oponentes com 40 de vida):

| Métrica | v2 (23/09) | **Lista nova** |
|---|---|---|
| Phlage da zona de comando até o T3 | 86% | **86%** |
| **1ª reanimação do Phlage** até o T4 · T6 · T8 | — | **33% · 57% · 67%** (mediana T6) |
| Helix por turno, T4 → T10 | 0,08 → 0,80 | **0,40 → 0,96** |
| Phlage permanente em campo até o T6 · T8 | 25% · 58% | **24% · 53%** (mediana T8 nas duas) |
| **1º oponente morto** até o T9 | 35% | **48%** (mediana T10 nas duas) |
| **Mesa inteira morta** até o T13 | 50% | **55%** (mediana **T13** nas duas) |
| Cartas extras por partida (compra) | 4,5 | **5,8** |

**O que isso quer dizer, sem suavizar:**
- O deck **faz alguma coisa** do T4 ao T6, coisa que a v2 não fazia: Helix barato, carta comprada,
  corpo em campo. O 1º oponente cai mais cedo com mais frequência.
- **A mesa continua morrendo no T13.** Cada volta do loop é 3 de dano em **um** oponente, contra 120
  de vida na mesa. Os pingadores da v2 batiam nos três ao mesmo tempo e saíram para dar lugar ao loop.
- **Correção do que eu disse ontem:** eu escrevi "Helix repetido a partir do T4 em ~67% das
  partidas". A simulação dá **33%**. Os 67% contavam ter a carta na mão; faltou contar ter a mana e o
  Phlage no cemitério no mesmo turno. E o "Phlage permanente no T5 com o Confession Dial" acontece em
  **6%** das partidas: é uma cópia só.
- O que encurta o jogo são **multiplicadores**. A Gisela sozinha pesa tanto quanto o pacote inteiro de
  reanimação (§4). O multiplicador que a Fase 7 mediu como melhor, [**Twinflame Tyrant**](https://www.ligamagic.com.br/?view=cards/card&card=Twinflame+Tyrant), custa
  **R$ 121,67**, e o 2º, [**Agate Instigator**](https://www.ligamagic.com.br/?view=cards/card&card=Agate+Instigator), **R$ 55,99**. Os dois estão fora do teto.
- Vencer até o T7 continua fora de alcance em RW sob R$ 200, em qualquer configuração que a Fase 7 testou.

> **Pergunta:** o deck mais ativo cedo, com a mesa morrendo por volta do T13, resolve a sensação de
> "fraco"? Se a resposta for não, o problema **não** é mais de carta: é o teto de preço contra este
> comandante, e aí a conversa é outra (subir o teto ou rever o comandante).

### 0.2 Um multiplicador que cabe: [**Fiendish Duo**](https://www.ligamagic.com.br/?view=cards/card&card=Fiendish+Duo) no lugar do [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate)

| | Custo | Efeito medido (Fase 7) |
|---|---|---|
| [**Fiendish Duo**](https://www.ligamagic.com.br/?view=cards/card&card=Fiendish+Duo) no lugar do [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | **+R$ 9,88** (folga R$ 33,02 → R$ 23,14) | mesa morta até o T13 **+6,0 pp** |

- O [**Fiendish Duo**](https://www.ligamagic.com.br/?view=cards/card&card=Fiendish+Duo) (`{4}{R}{R}`, 5/5, iniciativa) dobra **todo** dano a oponentes, inclusive o que
  eles causam uns aos outros. Em você, não dobra nada.
- O que se perde com a saída do [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate): o último feitiço X (wipe cirúrgico ou fecho) e uma
  fonte de combustível (10 → 9). Por isso a Fase 7 deixou a troca condicionada, e não a apliquei na lista.
- Também cabe o [**Dictate of the Twin Gods**](https://www.ligamagic.com.br/?view=cards/card&card=Dictate+of+the+Twin+Gods) (R$ 4,50, +6,3 pp), mas ele é **simétrico**: dobra o
  dano que você recebe também. Não recomendo.

> **Pergunta:** entra o [**Fiendish Duo**](https://www.ligamagic.com.br/?view=cards/card&card=Fiendish+Duo)? Minha recomendação é **sim**: é a única alavanca de relógio
> que cabe no teto.

### Das quatro perguntas da versão de 23/09, o que aconteceu

- **Ritmo:** virou a 0.1 acima.
- **Mão sem criatura:** resolvida pela revisão. As criaturas foram de 15 para **17**, e a mão inicial
  sem criatura cai de 30,4% para **25,5%**.
- **Fecho barato (**Repercussion**):** não se aplica mais. O [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) saiu e não
  há combo de 2 cartas nem loop infinito na lista nova (Fase 7, rulings conferidos).
- **Proteção:** entrou [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) (do Tori físico). O [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) (R$ 33,17) segue
  fora: consumiria a folga inteira e não segura exílio nem ódio de cemitério.

---

## 1. Por que esta revisão: o que o EDHREC mostrou

**O erro da v2.** A Fase 1 registrou uma trava: *"reanimar o Phlage dispara o sacrifício, então nenhum
pacote de reanimação entra"*. A premissa está certa, mas a conclusão não. O Phlage tem **duas**
habilidades de entrada, e o ruling oficial diz que o Helix dispara mesmo sem escape (*"Phlage's second
ability triggers when it enters the battlefield, even if it didn't escape"*, WotC 2024-06-07).
Reanimado, ele causa 3 de dano, você ganha 3 de vida, ele se sacrifica e volta ao cemitério, pronto
para a próxima reanimação. A trava foi revogada no `decisions.md`.

**O que a comunidade faz com ele.** No EDHREC, as cartas de maior sinergia com o Phlage são quase
todas de reanimação: [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand), [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation), [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan),
[**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth), [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness), [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission). No Archidekt (bracket 3),
as duas listas mais próximas do nosso teto, *Usain Bolt* e *Budget Phlag*, são loops de reanimação.
Descartei uma lista de hatebears e stax (regra 11) e uma turbo com One Ring e Smothering Tithe (preço).
A v2 coincidia com elas no esqueleto (rochas, terrenos, loots, Gisela, Inti) e em **nenhuma** peça do
motor. Análise completa em `08-meta-edhrec.md`.

**Achado da Fase 2, fora do meta:** [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial) e [**Desdemona, Freedom's Edge**](https://www.ligamagic.com.br/?view=cards/card&card=Desdemona%2C+Freedom%27s+Edge) dão ao
Phlage um escape de `{1}{R}{W}`. Pelos rulings, escape concedido por outra carta também conta como
escape, e o 6/6 **fica em campo**.

---

## 2. Resumo em números

| Item | Valor |
|---|---|
| Comandante | [**Phlage, Titan of Fire's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Phlage%2C+Titan+of+Fire%27s+Fury) `{1}{R}{W}` · escape `{R}{R}{W}{W}` + exilar 5 do cemitério |
| Eixo | **reanimação em loop do Phlage** + escape (concedido ou próprio) + controle |
| Cartas | **100** (99 + comandante), todas RW e legais |
| Terrenos | **37**: 24 básicos (12/12) + 13 não-básicos, com 11 fontes duplas |
| Criaturas / artefatos / encantamentos / instantâneos / feitiços | 17 / 12 / 7 / 11 / 15 |
| **Custo das 99** | **R$ 166,98**, folga de **R$ 33,02** no teto de R$ 200 · LigaMagic (menor), cotações de 2026-08-12 a 2026-09-24 |
| **A comprar** | **52 cartas · R$ 141,85** |
| Já existem | 47 cartas (R$ 25,13 na régua): 12 não-básicas do Tori físico + 11 da caixa + 24 básicos |
| Cartas sem cotação | **nenhuma** |

### Contra as metas do pipeline

| Categoria | v2 | **Nova** | Meta | Nota |
|---|---|---|---|---|
| Reanimação do Phlage | 0 | **10** (+2 dobradores) | — | 5 mágicas avulsas, 3 motores, 2 escapes concedidos |
| Draw | 12 | **12** | 12–13 | no piso. Entram Tocasia's Welcome, Ark of Hunger e Haliya |
| Ramp padrão / explosivo | 9 / 2 | **10 / 1** | 10–11 / 2–3 | a Fase 4 mediu que o limite é acesso à carta, não mana |
| Remoção | 20 | **15** (+2,5 via Helix) | ~10 | a Fase 5 conta o Helix como 2,5 remoções, só contra resistência ≤ 3 |
| Respostas a artefato/encantamento | 9 | **9** | — | inclui o 2-por-1 [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) |
| Wipes | 3 + 1 terreno | **3 + 1 terreno** | 2–4 | + overload do [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) e o [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) |
| Proteção | 1 | **2 + 2 parciais** | 2 | Gods Willing e Duty Beyond Death; parciais: Angelic Renewal e Broodmoth |
| Combustível de cemitério | 14 | **10** + vidência 3 do Dial | ~11 | abaixo do alvo depois que o Big Score saiu (§11) |
| Caminhos de vitória | 4 | **4** | 3+ | §4 |
| Terrenos | 37 | **37** | — | Fase 6: divisão 12/12 mantida |

---

## 3. Como o deck joga

1. **T3 — Phlage da zona de comando** (86%): 3 de dano, 3 de vida, e ele se sacrifica.
   **Deixe-o no cemitério** (a volta à zona de comando é opcional, CR 903.9a). Só mande à zona de
   comando se ele for exilado.
2. **T4 em diante — reanimação.** Cada [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand), [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant),
   [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness), [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) ou [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) é um Helix de
   1–2 manas. [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) e [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) transformam cada volta em carta.
   A 1ª reanimação sai até o T4 em 1/3 das partidas e até o T6 em mais da metade.
3. **Motores.** [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) (a cada artefato ou lendária conjurada),
   [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) (ao entrar e ao atacar) e [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) (ao conectar) repetem o Helix
   sem gastar carta. [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) dobra cada reanimação.
4. **Phlage permanente.** Pelo [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial) ou pela [**Desdemona, Freedom's Edge**](https://www.ligamagic.com.br/?view=cards/card&card=Desdemona%2C+Freedom%27s+Edge) (escape
   de 3 manas) ou pelo escape próprio (plano B). Aí ele é um 6/6 que dá Helix a cada ataque.
5. **Fecho.** Multiplicadores sobre o Helix e o combate: [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) (Helix de 6),
   [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight). Somam-se [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire),
   [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) e [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation).

**Regras de jogo que valem para qualquer troca futura:**
- Reanimação, Broodmoth e Angelic Renewal devolvem o Phlage **sem** escape: Helix e sacrifício. Só o
  escape (próprio ou concedido) mantém o 6/6 em campo.
- Blink ([**Cloudshift**](https://www.ligamagic.com.br/?view=cards/card&card=Cloudshift), [**Ephemerate**](https://www.ligamagic.com.br/?view=cards/card&card=Ephemerate)) fica fora: só funciona em resposta ao sacrifício, e no
  Phlage escapado tira o escape.
- [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness) põe contador de voar, e isso desliga a Broodmoth naquela volta. Com
  ela em campo, use outra reanimação.
- [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) + [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell): quem sofre o dano ordena os efeitos, e o
  resultado é **8**, não 10.
- Não exile o [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) pagando escape antes de usar o flashback dele.

**Pontos fortes:** jogada útil do T3 ao T6; resiliência a wipe (Broodmoth, Duty e o próprio escape);
compra ligada ao motor; 15 remoções, 9 respostas a artefato/encantamento.

**Pontos fracos:**
- relógio de mesa na mediana do T13 (§0.1);
- **ódio de cemitério instantâneo** ([**Relic of Progenitus**](https://www.ligamagic.com.br/?view=cards/card&card=Relic+of+Progenitus), [**Bojuka Bog**](https://www.ligamagic.com.br/?view=cards/card&card=Bojuka+Bog)) não tem resposta em RW.
  A saída é mandar o Phlage exilado à zona de comando e reconjurá-lo com imposto;
- os motores chegam tarde: só se vê um deles até o T4 em ~39% das partidas;
- nenhuma contramágica.

---

## 4. Caminhos de vitória

| # | Caminho | Peças |
|---|---|---|
| **R1** | Phlage escapado + monarca | Phlage 6/6 atacando (Helix a cada ataque) + [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) (7 por upkeep, 9 com Torbran) |
| **R2** | Helix em loop multiplicado | reanimações e motores × [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer), [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell), [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) |
| **R3** | Dano espalhado | [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) (cada saída do cemitério) + [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) |
| **R4** | Vitória alternativa | [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun), 9% das vitórias simuladas |

**Sem combo de 2 cartas e sem loop infinito.** A cadeia Broodmoth + Renewal é finita, e a cópia do
Sevinne's não mira o Phlage de novo (Fase 7, rulings conferidos).

**O que mais pesa no relógio**, medido carta a carta pela Fase 7 (quanto a mesa morta até o T13 cai sem
ela): Gisela −10,4 pp · Court of Ire −6,6 · Calamity Bearer −5,9 · Sun Titan −5,0 · Ark of Hunger −4,5 ·
Approach −4,2 · Torbran −4,0. As reanimações avulsas valem ~0 no fechamento, mas cada uma vale −5 a
−6 pp na reanimação até o T4. São elas que fazem o deck jogar cedo.

---

## 5. Base de mana

- **37 terrenos**, 13 não-básicos intocados; a fórmula recalculada pela Fase 6 dá 36,4.
- **Básicos 12 Mountain / 12 Plains, mantidos.** As mágicas agora pedem cor quase meio a meio
  (33 símbolos de `{R}`, 34 de `{W}`), e há 23 fontes de cada cor em terreno. Passar para 11/13 ajudaria
  a Broodmoth em 1,4 pp e tiraria 2,8 pp do Torbran.
- Marcos na simulação da Fase 6: Helping Hand + outra carta de 2 manas no T4 em **93,6%**; Broodmoth
  (`{W}{W}`) no T4 em **68,4%**; Dial no T4 e escape de `{1}{R}{W}` no T5 em **92,2%** (desde que as
  cartas estejam na mão).
- Pendência opcional: trocar [**Boros Guildgate**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Guildgate) por [**Sundown Pass**](https://www.ligamagic.com.br/?view=cards/card&card=Sundown+Pass) (R$ 10,00) dá +0,8 a
  +1,3 pp. Não recomendo com a folga atual.

---

## 6. O deck final — 100 cartas

**Comandante (fora da régua de preço):** [**Phlage, Titan of Fire's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Phlage%2C+Titan+of+Fire%27s+Fury) — `{1}{R}{W}` ·
Legendary Creature — Elder Giant 6/6. Ao entrar, sacrifica-se se não escapou. Sempre que entra ou ataca,
causa 3 de dano em qualquer alvo e você ganha 3 de vida. Escape: `{R}{R}{W}{W}` + exilar cinco outras
cartas do seu cemitério.

### Criaturas (17)

| Carta | CMC | Custo | Categorias | O que faz neste deck | Origem | R$ |
|---|---|---|---|---|---|---|
| [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) | 2 | `{1}{R}` | draw, tema | todo descarte vira impulse (janela longa) — multiplica os loots; corpo 2/2 | compra | 4,90 |
| [**Lesser Masticore**](https://www.ligamagic.com.br/?view=cards/card&card=Lesser+Masticore) | 2 | `{2}` | tema, remoção | descarte no custo = combustível; ping {4} repetível; persist; histórico para o Teshar | caixa | 0,09 |
| [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) | 2 | `{2}` | ramp, tema | motor: moe 1/turno em velocidade de instantâneo + mana + corpo | compra | 0,18 |
| [**Myr Convert**](https://www.ligamagic.com.br/?view=cards/card&card=Myr+Convert) | 2 | `{2}` | ramp, tema | rocha de qualquer cor + corpo 2/1; histórico para o Teshar | caixa | 0,15 |
| [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise) | 2 | `{2}` | ramp, tema | mana de qualquer cor + 0/2 voador; artefato histórico para o Teshar | compra | 6,79 |
| [**Zookeeper Mechan**](https://www.ligamagic.com.br/?view=cards/card&card=Zookeeper+Mechan) | 2 | `{1}{R}` | ramp, tema | rocha {R} + corpo 1/3 que segura T2–T4; {6}{R}: +4/+0 no Phlage (Fase 4 R2) | caixa | 0,09 |
| [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) | 3 | `{2}{W}` | draw | compra no fim do turno se você ganhou 3+ de vida (todo Helix basta); lendária MV 3, reanimável | compra | 11,88 |
| [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) | 3 | `{2}{R}` | draw, tema | carta ou moinho+dano todo upkeep; 3/2 menace; fonte vermelha | compra | 1,00 |
| [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) | 3 | `{1}{R}{W}` | tema | causou 3+ de dano de combate → devolve o Phlage; 3/3 vigilância e atropelar | compra | 0,49 |
| [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) | 4 | `{2}{R}{R}` | wincon | o Phlage é Giant: Helix de 6, combate do escapado 12; 3/4 | compra | 0,24 |
| [**Desdemona, Freedom's Edge**](https://www.ligamagic.com.br/?view=cards/card&card=Desdemona%2C+Freedom%27s+Edge) | 4 | `{2}{R}{W}` | tema | ao atacar, o Phlage ganha escape de {1}{R}{W} + 2 cartas: o 6/6 **fica**; 3/4 vigilância | compra | 1,45 |
| [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) | 4 | `{2}{W}{W}` | tema, proteção | o Phlage sacrificado volta com voar = 2 Helix por reanimação; devolve seu time no wipe | compra | 6,74 |
| [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) | 4 | `{2}{W}{W}` | draw, remoção | monarca + exila criatura; corpo 2/2 | compra | 0,30 |
| [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) | 4 | `{3}{W}` | tema | devolve o Phlage (MV 3) a cada mágica histórica: 16 artefatos + lendárias; 2/2 voador | compra | 0,60 |
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | 4 | `{1}{R}{R}{R}` | wincon | +2 por fonte vermelha: Helix 5 (8 com o Calamity Bearer, na ordem pior para você), Court 9 | compra | 6,98 |
| [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) | 6 | `{4}{W}{W}` | tema, wincon | ao entrar e a cada ataque devolve o Phlage, rocha, Casket ou Seal; é Giant (Bearer dobra) | compra | 2,25 |
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | 7 | `{4}{R}{W}{W}` | wincon | dobra dano a oponentes (inclui combate) e corta pela metade o recebido; 5/5 voadora | compra | 15,00 |

### Artefatos (12)

| Carta | CMC | Custo | Categorias | O que faz neste deck | Origem | R$ |
|---|---|---|---|---|---|---|
| [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring) | 1 | `{1}` | ramp | intocável; paga reanimação + outra jogada no mesmo turno | Tori físico | 7,95 |
| [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet) | 2 | `{2}` | ramp | intocável; {R}/{W} no T2 → Phlage T3 | Tori físico | 4,00 |
| [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet) | 2 | `{2}` | ramp | {1}→{R}{W}: metade do escape numa ativação; converte o {C} do Sol Ring | compra | 1,50 |
| [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) | 2 | `{1}{W}` | remoção | exílio permanente MV≤3; Recommission e Sun Titan a recompram | caixa | 0,09 |
| [**Ratchet Bomb**](https://www.ligamagic.com.br/?view=cards/card&card=Ratchet+Bomb) | 2 | `{2}` | wipe | seletivo por MV; mata fichas com 0 contadores | caixa | 1,00 |
| [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) | 2 | `{2}` | draw | 3 cartas por {2} cada; Treasure no fim | compra | 1,90 |
| [**Sphere of the Suns**](https://www.ligamagic.com.br/?view=cards/card&card=Sphere+of+the+Suns) | 2 | `{2}` | ramp | 3 manas de qualquer cor (entra virada); histórico para o Teshar | caixa | 0,11 |
| [**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction) | 2 | `{2}` | ramp | mv2 de cor à escolha; 1 de vida pago pelo Helix | compra | 3,80 |
| [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) | 2 | `{2}` | draw | página a cada entrada ou ataque do comandante; **cada reanimação é uma entrada** | compra | 1,20 |
| [**Commander's Sphere**](https://www.ligamagic.com.br/?view=cards/card&card=Commander%27s+Sphere) | 3 | `{3}` | ramp | {R}/{W}; sacrifica para comprar; histórico para o Teshar | caixa | 0,50 |
| [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial) | 3 | `{3}` | tema | {T}: o Phlage ganha escape de {1}{R}{W} + 3 cartas: o 6/6 **fica**; ETB vidência 3 | compra | 1,42 |
| [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) | 4 | `{2}{R}{W}` | draw, wincon, tema | 1 a cada oponente + 1 de vida sempre que carta sai do cemitério; {T}: moe 1 e pode jogá-la | compra | 0,45 |

### Encantamentos (7)

| Carta | CMC | Custo | Categorias | O que faz neste deck | Origem | R$ |
|---|---|---|---|---|---|---|
| [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) | 2 | `{1}{W}` | tema, proteção | pré-paga: devolve o Phlage (+1 Helix) ou salva Torbran, Gisela ou Sun Titan | compra | 0,80 |
| [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) | 2 | `{R}{W}` | tema, wincon | motor: moe 1/turno automático + 1 a cada oponente/turno (fonte vermelha → Torbran) | compra | 0,08 |
| [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing) | 2 | `{1}{W}` | remoção | artefato/encantamento pré-pago; o Sun Titan o recompra | Tori físico | 0,30 |
| [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) | 3 | `{2}{R}` | draw, tema | ETB descarta 1/compra 2; resolvido: todo upkeep descarta a mão e compra 2 (motor) | compra | 0,90 |
| [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) | 3 | `{2}{W}` | draw | 1 carta por turno em que criatura sua de MV≤3 entra: cada reanimação do Phlage | compra | 7,88 |
| [**Celebrate the Mountain-king**](https://www.ligamagic.com.br/?view=cards/card&card=Celebrate+the+Mountain-king) | 4 | `{3}{W}` | remoção, tema | 1 não-terreno por oponente + recruit (loot) | Tori físico | 0,23 |
| [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | 5 | `{3}{R}{R}` | draw, remoção, wincon | monarca + 7 de dano/upkeep em qualquer alvo (9 c/ Torbran) | compra | 4,95 |

### Instantâneos (11)

| Carta | CMC | Custo | Categorias | O que faz neste deck | Origem | R$ |
|---|---|---|---|---|---|---|
| [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) | 1 | `{W}` | proteção | protege o Phlage escapado, Torbran, Gisela ou Sun Titan por {W}; scry 1 | Tori físico | 0,45 |
| [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | 1 | `{R}` | remoção | delirium → 6; cresce com o cemitério | compra | 0,50 |
| [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) | 2 | `{1}{R}` | remoção | 3 em criatura OU destrói artefato | compra | 0,30 |
| [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) | 2 | `{1}{R}` | tema, remoção | modal: 3 de dano em criatura/PW OU loot 2 — instantâneo | compra | 0,84 |
| [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death) | 2 | `{1}{W}` | proteção | sacrifica o Phlage reanimado (que já ia sair) → seu time indestrutível + contador; Fumigate unilateral | Tori físico | 0,43 |
| [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) | 2 | `{R}{W}` | remoção, wincon | 3 dano + 3 vida; fonte vermelha | compra | 0,45 |
| [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) | 2 | `{1}{R}` | remoção, wincon | 3 em qualquer alvo; fonte vermelha | caixa | 0,10 |
| [**Smite the Deathless**](https://www.ligamagic.com.br/?view=cards/card&card=Smite+the+Deathless) | 2 | `{1}{R}` | remoção | 3 + remove indestrutível + exila | caixa | 0,12 |
| [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) | 3 | `{2}{R}` | remoção | irrestrita, instantânea | compra | 4,12 |
| [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) | 3 | `{2}{W}` | remoção | irrestrita, instantânea | compra | 5,00 |
| [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) | 3 | `{1}{R} // {W}` | remoção | fuse: artefato E encantamento em instantâneo — responde a Rest in Peace e Grafdigger's Cage (Fase 5) | compra | 2,91 |

### Feitiços (15)

| Carta | CMC | Custo | Categorias | O que faz neste deck | Origem | R$ |
|---|---|---|---|---|---|---|
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | 1 | `{X}{X}{R}` | tema, wipe, wincon | X dividido (wipe cirúrgico/fecho); flashback descarta X = combustível; fonte vermelha | compra | 0,07 |
| [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) | 1 | `{R}` | tema | 3 ao cemitério por {R}; flashback | compra | 5,40 |
| [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) | 1 | `{R}` | tema, remoção | retrace: nunca sai do cemitério, converte terreno excedente em 1 de dano (3 c/ Torbran) | compra | 0,05 |
| [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) | 1 | `{W}` | tema, remoção | reanima o Phlage: Helix por 1 mana | compra | 0,45 |
| [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness) | 2 | `{1}{W}` | tema, remoção | reanima o Phlage: Helix (o contador de voar desliga a Broodmoth nessa volta) | compra | 0,15 |
| [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) | 2 | `{1}{R}` | remoção, wipe | 4 numa criatura alheia; overload = wipe unilateral | compra | 1,10 |
| [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) | 2 | `{1}{W}` | tema, remoção | reanima o Phlage **ou** um artefato (Casket, Dial, Sol Ring) | compra | 1,00 |
| [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant) | 2 | `{1}{W}` | tema, remoção | reanima o Phlage: Helix | compra | 0,20 |
| [**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) | 3 | `{2}{R}` | draw | impulse 2 por {R}: o Helix liga o spectacle | compra | 0,74 |
| [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory) | 3 | `{1}{W}{W}` | remoção | exila qualquer não-terreno | Tori físico | 0,13 |
| [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) | 3 | `{2}{R}` | tema, ramp | loot + 1 Treasure; mv3 | compra | 0,06 |
| [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) | 3 | `{2}{W}` | tema, remoção | reanima o Phlage; flashback {4}{W} = 2º Helix e a cópia devolve terreno ou rocha | compra | 3,50 |
| [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) | 3 | `{1}{W}{W}` | wipe | sacrifício por poder: poupa os pingadores 0–2 | compra | 0,20 |
| [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) | 5 | `{3}{W}{W}` | wipe | vida por criatura; com Broodmoth ou Duty Beyond Death fica unilateral | caixa | 1,37 |
| [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) | 7 | `{6}{W}` | wincon | alt-win: conjurada 2× = vitória; 7 de vida por conjura | compra | 12,02 |

### Terrenos (37)

| Carta | CMC | Custo | Categorias | O que faz neste deck | Origem | R$ |
|---|---|---|---|---|---|---|
| [**Battlefield Forge**](https://www.ligamagic.com.br/?view=cards/card&card=Battlefield+Forge) | 0 | — | terreno | dual desvirada; {C} grátis; 1 de dano pago pelo Helix · origem: EDHREC/Archidekt | compra | 3,79 |
| [**Blast Zone**](https://www.ligamagic.com.br/?view=cards/card&card=Blast+Zone) | 0 | — | terreno, wipe | wipe por MV no slot de terreno | caixa | 0,89 |
| [**Boros Guildgate**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Guildgate) | 0 | — | terreno | dual virada; substitui Sunbillow Verge (R$ 56,00) | Tori físico | 0,23 |
| [**Clifftop Retreat**](https://www.ligamagic.com.br/?view=cards/card&card=Clifftop+Retreat) | 0 | — | terreno | dual desvirada com Mountain/Plains (22 básicos + 2 Mountain Plains) · origem: EDHREC/Archidekt | compra | 2,90 |
| [**Command Tower**](https://www.ligamagic.com.br/?view=cards/card&card=Command+Tower) | 0 | — | terreno | intocável; {R}/{W} desvirado | Tori físico | 1,50 |
| [**Fields of Strife**](https://www.ligamagic.com.br/?view=cards/card&card=Fields+of+Strife) | 0 | — | terreno | dual virada + surveil 1 por {2}{R}{W} | Tori físico | 0,01 |
| [**Furycalm Snarl**](https://www.ligamagic.com.br/?view=cards/card&card=Furycalm+Snarl) | 0 | — | terreno | dual; desvirada revelando básico | compra | 1,10 |
| [**Glittering Massif**](https://www.ligamagic.com.br/?view=cards/card&card=Glittering+Massif) | 0 | — | terreno, tema | dual virada + cycling {2}: combustível, tipo Land no cemitério p/ delirium da Unholy Heat, seguro contra inundação | compra | 2,48 |
| **12×** [**Mountain**](https://www.ligamagic.com.br/?view=cards/card&card=Mountain) | 0 | — | terreno | básico | Tori físico | 2,40 |
| **12×** [**Plains**](https://www.ligamagic.com.br/?view=cards/card&card=Plains) | 0 | — | terreno | básico | Tori físico | 2,64 |
| [**Radiant Summit**](https://www.ligamagic.com.br/?view=cards/card&card=Radiant+Summit) | 0 | — | terreno | desvirada com 2+ básicos (22 no deck) · origem: EDHREC/Archidekt | compra | 5,46 |
| [**Rugged Prairie**](https://www.ligamagic.com.br/?view=cards/card&card=Rugged+Prairie) | 0 | — | terreno | filtro {R/W}→{R}{R}/{R}{W}/{W}{W}: Mountain + Prairie = {W}{W} do escape · origem: EDHREC/Archidekt | compra | 2,73 |
| [**Stone Quarry**](https://www.ligamagic.com.br/?view=cards/card&card=Stone+Quarry) | 0 | — | terreno | dual virada | Tori físico | 0,25 |
| [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) | 0 | — | terreno, draw | 3ª entrada para a coroa (Jailer, Court); conta nas 12 de draw da Fase 3 | compra | 0,65 |
| [**Wind-Scarred Crag**](https://www.ligamagic.com.br/?view=cards/card&card=Wind-Scarred+Crag) | 0 | — | terreno | dual virada + 1 de vida | Tori físico | 0,10 |

---

## 7. O que ficou fora

### 7.1 Saíram da v2 nesta revisão (19)

| Sai | R$ | Fase | Motivo | Quem cobre |
|---|---|---|---|---|
| [**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) | 5,40 | 2 e 7 | pingador: a Fase 7 mediu ~1 mágica por turno, contra as 3–5 que a rota supunha | Helix multiplicado, Ark of Hunger. Voltar vale ≤ +2,7 pp |
| [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist) | 0,55 | 2 e 7 | idem | Lorehold Excavation, Ark. **1ª reserva** se a Fase 7 medir mais mágicas por turno |
| [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith) | 0,33 | 2 e 7 | idem | Ark, Excavation |
| [**Firebrand Archer**](https://www.ligamagic.com.br/?view=cards/card&card=Firebrand+Archer) | 0,99 | 2 e 7 | idem | Ark, Excavation |
| [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) | 26,40 | 2 e 7 | fecho de 11 manas. Voltar rende +0,8 pp no T11 | Helix multiplicado, Court of Ire |
| [**Perpetual Timepiece**](https://www.ligamagic.com.br/?view=cards/card&card=Perpetual+Timepiece) | 0,65 | 2 | combustível; a reanimação não precisa de cemitério cheio | Ark, Dial (vidência 3), Millikin, Excavation |
| [**Cathartic Reunion**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Reunion) | 0,14 | 2 | combustível | loots que ficam + Tocasia |
| [**Demand Answers**](https://www.ligamagic.com.br/?view=cards/card&card=Demand+Answers) | 2,80 | 2 | combustível | loots que ficam |
| [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) | 0,50 (caixa) | 2 | combustível | Cathartic Pyre. Voltar vale +1,2 pp (§11) |
| [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) | 0,05 (caixa) | 2 e 3 | draw avulso | Tocasia's Welcome |
| [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) | 0,90 | 2 e 3 | draw | Ark of Hunger (o mesmo impulse, mais alcance) |
| [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) | 0,09 (Tori) | 3 | 5 manas, fora do alcance de toda reanimação, e o Helix não a dispara | Haliya |
| [**Magma Spray**](https://www.ligamagic.com.br/?view=cards/card&card=Magma+Spray) | 0,05 (caixa) | 2 e 5 | dano pequeno em criatura é o que o Helix mais repete | Unholy Heat, Smite the Deathless |
| [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning) | 0,10 (Tori) | 2 e 5 | idem | Chaos Warp, Generous Gift, Reduce to Memory |
| [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) | 0,44 (caixa) | 2 e 5 | 1 mana a mais por modo que o Wear // Tear, e é feitiço | Wear // Tear |
| [**Disenchant**](https://www.ligamagic.com.br/?view=cards/card&card=Disenchant) | 0,25 (Tori) | 5 | dá lugar ao Wear // Tear, que responde a Rest in Peace **e** Grafdigger's Cage em instantâneo | Wear // Tear |
| [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) | 2,48 (caixa) | 5 | feitiço, só criatura, só dano: o que o Helix já faz | Helix; entra Duty Beyond Death |
| [**Big Score**](https://www.ligamagic.com.br/?view=cards/card&card=Big+Score) | 7,65 | 4 | custo 4 disputando o T4 com oito 4-drops; ≈0 do T3 ao T5 | Sphere of the Suns |

O [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) foi cortado pela Fase 2 e **devolvido** pela Fase 5, por isso não está na lista acima.

As cartas que saíram e vieram da caixa ou do Tori físico **continuam sobressalentes**. Fichas F1–F7 de
cada corte: `02-theme.md` §12.6, `03-draw.md`, `04-ramp.md` e `05-interaction.md` (seções de revisão).

### 7.2 Por preço

| Carta | R$ | Função |
|---|---|---|
| [**Smothering Tithe**](https://www.ligamagic.com.br/?view=cards/card&card=Smothering+Tithe) | 233,91 | Treasures |
| [**Twinflame Tyrant**](https://www.ligamagic.com.br/?view=cards/card&card=Twinflame+Tyrant) | 121,67 | o melhor multiplicador para o loop (+8,2 pp na mesa até o T13) |
| [**Solphim, Mayhem Dominus**](https://www.ligamagic.com.br/?view=cards/card&card=Solphim%2C+Mayhem+Dominus) | 118,90 | multiplicador |
| [**Palantír of Orthanc**](https://www.ligamagic.com.br/?view=cards/card&card=Palant%C3%ADr+of+Orthanc) | 117,98 | draw + auto-moinho |
| [**Celestine, the Living Saint**](https://www.ligamagic.com.br/?view=cards/card&card=Celestine%2C+the+Living+Saint) | 110,38 | reanimação a cada fim de turno |
| [**Monument to Endurance**](https://www.ligamagic.com.br/?view=cards/card&card=Monument+to+Endurance) | 110,00 | draw por descarte |
| [**Lightning, Army of One**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning%2C+Army+of+One) | 73,59 | multiplicador |
| [**Arena of Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Arena+of+Glory) | 69,12 | ímpeto para o Phlage escapado |
| [**The Gaffer**](https://www.ligamagic.com.br/?view=cards/card&card=The+Gaffer) | 63,75 | compra com 3+ de vida (a Haliya faz o mesmo) |
| [**Agate Instigator**](https://www.ligamagic.com.br/?view=cards/card&card=Agate+Instigator) | 55,99 | 1 a cada oponente por criatura que entra |
| [**Sunbillow Verge**](https://www.ligamagic.com.br/?view=cards/card&card=Sunbillow+Verge) | 56,00 | dupla desvirada |
| [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation) | 54,54 | multiplicador |
| [**Repercussion**](https://www.ligamagic.com.br/?view=cards/card&card=Repercussion) | 47,99 | fecho de 1 turno |
| [**Spectator Seating**](https://www.ligamagic.com.br/?view=cards/card&card=Spectator+Seating) | 38,90 | dupla |
| [**Chandra's Ignition**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Ignition) | 37,90 | wipe/fecho |
| [**Panharmonicon**](https://www.ligamagic.com.br/?view=cards/card&card=Panharmonicon) | 34,45 | dobraria o Helix |
| [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) | 33,17 | proteção em massa |

---

## 8. Custo — LigaMagic (menor), cotações de 2026-08-12 a 2026-09-24

Régua: **R$ 200,00 nas 99 cartas**. O comandante não conta; os básicos contam (critério
conservador do briefing). Toda cotação é o **primeiro** número do bloco "Preço Médio de Venda no
Marketplace" e está registrada com `mtgdb prices -add`.

| Bloco | Cartas | R$ |
|---|---|---|
| **A comprar** | 52 | **141,85** |
| Já existem (Tori físico + caixa), contam na régua | 47 | 25,13 |
| **Total das 99** | **99** | **166,98** |
| **Folga até R$ 200,00** | | **33,02** |
| Com o [**Fiendish Duo**](https://www.ligamagic.com.br/?view=cards/card&card=Fiendish+Duo) da §0.2 | | total 176,86 · folga 23,14 |

### Lista de compra (52), da mais cara para a mais barata

| Carta | R$ (menor) |
|---|---|
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | 15,00 |
| [**Approach of the Second Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Approach+of+the+Second+Sun) | 12,02 |
| [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) | 11,88 |
| [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) | 7,88 |
| [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell) | 6,98 |
| [**Ornithopter of Paradise**](https://www.ligamagic.com.br/?view=cards/card&card=Ornithopter+of+Paradise) | 6,79 |
| [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) | 6,74 |
| [**Radiant Summit**](https://www.ligamagic.com.br/?view=cards/card&card=Radiant+Summit) | 5,46 |
| [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) | 5,40 |
| [**Generous Gift**](https://www.ligamagic.com.br/?view=cards/card&card=Generous+Gift) | 5,00 |
| [**Court of Ire**](https://www.ligamagic.com.br/?view=cards/card&card=Court+of+Ire) | 4,95 |
| [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun) | 4,90 |
| [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp) | 4,12 |
| [**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction) | 3,80 |
| [**Battlefield Forge**](https://www.ligamagic.com.br/?view=cards/card&card=Battlefield+Forge) | 3,79 |
| [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) | 3,50 |
| [**Wear // Tear**](https://www.ligamagic.com.br/?view=cards/card&card=Wear+%2F%2F+Tear) | 2,91 |
| [**Clifftop Retreat**](https://www.ligamagic.com.br/?view=cards/card&card=Clifftop+Retreat) | 2,90 |
| [**Rugged Prairie**](https://www.ligamagic.com.br/?view=cards/card&card=Rugged+Prairie) | 2,73 |
| [**Glittering Massif**](https://www.ligamagic.com.br/?view=cards/card&card=Glittering+Massif) | 2,48 |
| [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) | 2,25 |
| [**Reckoner Bankbuster**](https://www.ligamagic.com.br/?view=cards/card&card=Reckoner+Bankbuster) | 1,90 |
| [**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet) | 1,50 |
| [**Desdemona, Freedom's Edge**](https://www.ligamagic.com.br/?view=cards/card&card=Desdemona%2C+Freedom%27s+Edge) | 1,45 |
| [**Confession Dial**](https://www.ligamagic.com.br/?view=cards/card&card=Confession+Dial) | 1,42 |
| [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) | 1,20 |
| [**Furycalm Snarl**](https://www.ligamagic.com.br/?view=cards/card&card=Furycalm+Snarl) | 1,10 |
| [**Mizzium Mortars**](https://www.ligamagic.com.br/?view=cards/card&card=Mizzium+Mortars) | 1,10 |
| [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) | 1,00 |
| [**Sin Prodder**](https://www.ligamagic.com.br/?view=cards/card&card=Sin+Prodder) | 1,00 |
| [**Case of the Crimson Pulse**](https://www.ligamagic.com.br/?view=cards/card&card=Case+of+the+Crimson+Pulse) | 0,90 |
| [**Cathartic Pyre**](https://www.ligamagic.com.br/?view=cards/card&card=Cathartic+Pyre) | 0,84 |
| [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) | 0,80 |
| [**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) | 0,74 |
| [**Throne of the High City**](https://www.ligamagic.com.br/?view=cards/card&card=Throne+of+the+High+City) | 0,65 |
| [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) | 0,60 |
| [**Unholy Heat**](https://www.ligamagic.com.br/?view=cards/card&card=Unholy+Heat) | 0,50 |
| [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) | 0,49 |
| [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) | 0,45 |
| [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) | 0,45 |
| [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) | 0,45 |
| [**Abrade**](https://www.ligamagic.com.br/?view=cards/card&card=Abrade) | 0,30 |
| [**Palace Jailer**](https://www.ligamagic.com.br/?view=cards/card&card=Palace+Jailer) | 0,30 |
| [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) | 0,24 |
| [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant) | 0,20 |
| [**Slaughter the Strong**](https://www.ligamagic.com.br/?view=cards/card&card=Slaughter+the+Strong) | 0,20 |
| [**Millikin**](https://www.ligamagic.com.br/?view=cards/card&card=Millikin) | 0,18 |
| [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness) | 0,15 |
| [**Lorehold Excavation**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold+Excavation) | 0,08 |
| [**Conflagrate**](https://www.ligamagic.com.br/?view=cards/card&card=Conflagrate) | 0,07 |
| [**Seize the Spoils**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+the+Spoils) | 0,06 |
| [**Flame Jab**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Jab) | 0,05 |

---

## 9. Lista para playtest (padrão MTG Online)

Texto puro, sem links: cole no Moxfield, no Archidekt ou no MTGO. O comandante vem por último, depois
da linha em branco. **Não inclui o **Fiendish Duo** da §0.2**; se você disser sim, troco pelo
Conflagrate.

```
1 Inti, Seneschal of the Sun
1 Lesser Masticore
1 Millikin
1 Myr Convert
1 Ornithopter of Paradise
1 Zookeeper Mechan
1 Haliya, Guided by Light
1 Sin Prodder
1 Venerable Warsinger
1 Calamity Bearer
1 Desdemona, Freedom's Edge
1 Luminous Broodmoth
1 Palace Jailer
1 Teshar, Ancestor's Apostle
1 Torbran, Thane of Red Fell
1 Sun Titan
1 Gisela, Blade of Goldnight
1 Sol Ring
1 Arcane Signet
1 Boros Signet
1 Glass Casket
1 Ratchet Bomb
1 Reckoner Bankbuster
1 Sphere of the Suns
1 Talisman of Conviction
1 Tome of Legends
1 Commander's Sphere
1 Confession Dial
1 Ark of Hunger
1 Angelic Renewal
1 Lorehold Excavation
1 Seal of Cleansing
1 Case of the Crimson Pulse
1 Tocasia's Welcome
1 Celebrate the Mountain-king
1 Court of Ire
1 Gods Willing
1 Unholy Heat
1 Abrade
1 Cathartic Pyre
1 Duty Beyond Death
1 Lightning Helix
1 Lightning Strike
1 Smite the Deathless
1 Chaos Warp
1 Generous Gift
1 Wear // Tear
1 Conflagrate
1 Faithless Looting
1 Flame Jab
1 Helping Hand
1 Call a Surprise Witness
1 Mizzium Mortars
1 Recommission
1 Return Triumphant
1 Light Up the Stage
1 Reduce to Memory
1 Seize the Spoils
1 Sevinne's Reclamation
1 Slaughter the Strong
1 Fumigate
1 Approach of the Second Sun
1 Battlefield Forge
1 Blast Zone
1 Boros Guildgate
1 Clifftop Retreat
1 Command Tower
1 Fields of Strife
1 Furycalm Snarl
1 Glittering Massif
12 Mountain
12 Plains
1 Radiant Summit
1 Rugged Prairie
1 Stone Quarry
1 Throne of the High City
1 Wind-Scarred Crag

1 Phlage, Titan of Fire's Fury
```

---

## 10. Protocolo de goldfishing

Nenhuma versão deste deck passou por goldfishing. Protocolo completo em `07-wincons.md` §R9.

**Execução:** 10 partidas, sozinho, com mulligan de verdade (Londres). Aponte todo Helix e toda queima
para o oponente com menos vida. Só causam dano de combate o Phlage escapado, o Sun Titan, a Gisela, as
voadoras, o Warsinger e a Desdemona. Aplique os multiplicadores na ordem pior para você. Pare no T14 ou
quando a mesa morrer.

**O que anotar, com o que a simulação espera:**

| # | Registro | Esperado em 10 partidas |
|---|---|---|
| 1 | mulligans | — |
| 2 | turno do Phlage da zona de comando | T3 em ~9 |
| 3 | **turno da 1ª reanimação** e qual carta | até o T4 em ~3; até o T6 em ~6 |
| 4 | **Helix por turno, do T4 ao T10** | ~0,4 no T4–T5; ~1 do T9 em diante |
| 5 | **turno do Phlage permanente** e por qual via | até o T8 em ~5; pelo Dial em ~2 |
| 6 | turno do 1º multiplicador | — |
| 7 | cartas compradas por Tocasia, Haliya, Tome e monarca | ~6 por partida |
| 8 | turno em que você causa 8+ de dano num turno | T7–T8 |
| 9 | **turno do 1º oponente morto** | até o T9 em ~5 |
| 10 | **turno da mesa morta**, e se foi pela Approach | até o T13 em ~5–6 |
| 11 | ≤ 3 terrenos no T4, ou faltou `{W}{W}`/`{R}{R}` | — |
| 12 | mão sem permanente jogável até o T3 | — |

**O que cada resultado derruba:**

| Se acontecer em… | …isto | Volta para |
|---|---|---|
| ≥ 6 de 10 | 1ª reanimação depois do T6 | Fase 2 |
| ≥ 6 de 10 | < 0,5 Helix por turno no T6–T10 | Fases 2 e 3 |
| ≥ 5 de 10 | Phlage permanente depois do T9 | Fases 2 e 4 |
| ≥ 5 de 10 | nenhum multiplicador até o T8 | Fase 7 |
| ≥ 4 de 10 | 1º oponente morto depois do T11 | Fase 7 |
| ≥ 5 de 10 | mesa morta depois do T14 | **você**: é o eixo, não a peça |
| ≥ 3 de 10 | Phlage não sai no T3 | Fase 4 |
| ≥ 3 de 10 | falta de terreno ou de `{W}{W}`/`{R}{R}` | Fase 6 |
| ≥ 7 de 10 | nenhum de Tocasia, Haliya e Tome até o T6 | Fase 3 |

Quando tiver os registros, é só trazer: eu os passo ao `wincon-tester` em modo `post-goldfish`.

---

## 11. Pendências

1. **As duas decisões da §0** (ritmo e Fiendish Duo).
2. **Goldfishing e teste em proxy** antes de comprar (§10).
3. **Combustível em 10**, abaixo do alvo de ~11 da Fase 2, depois que a Fase 4 tirou o Big Score. A
   reserva sem compra é o [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) (caixa, +1,2 pp). O registro 5 do goldfishing
   mede se falta.
4. **Reservas sem compra, já avaliadas:** [**Culling Dais**](https://www.ligamagic.com.br/?view=cards/card&card=Culling+Dais) (caixa, R$ 0,08: o Phlage sacrificado
   nela vira carta; a v2 a descartou por engano como "não compra"), [**Emerge from the Cocoon**](https://www.ligamagic.com.br/?view=cards/card&card=Emerge+from+the+Cocoon) e
   [**Hedron Crawler**](https://www.ligamagic.com.br/?view=cards/card&card=Hedron+Crawler) (caixa).
5. **Ódio de cemitério instantâneo** sem resposta em RW (§3). Risco declarado, não resolvido.
6. **Básicos na régua:** assumido que contam. Se não contarem, a folga sobe R$ 5,04.
7. **Cartas que sobram do Tori** quando o deck for montado: só entram em `data/collection.tsv` via
   `/update-collection`.
8. **Arquivos de fase desatualizados:** a §2.2 e a §9 do `06-manabase.md` ainda falam em 11/11 e em
   15 não-básicos, e a contagem de cor da §12.8 do `02-theme.md` conta Boros Signet e Talisman como RW.
   Nada disso muda a lista.

---

## 12. Como validar

1. **Responda às duas perguntas da §0.**
2. Confira as cartas na §6: cada nome abre a ficha da LigaMagic, com texto e preço.
3. Cole o bloco da §9 no Moxfield ou no Archidekt, ou imprima em proxy.
4. Aprovou → escrevo `deck.md` e `decisions.md` e fecho a rodada. Reprovou uma carta ou uma
   categoria → devolvo só à fase responsável e reconsolido este arquivo.
