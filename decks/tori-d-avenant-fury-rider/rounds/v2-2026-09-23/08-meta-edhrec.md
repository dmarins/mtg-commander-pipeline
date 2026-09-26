# 08 — Meta do EDHREC e Archidekt para o Phlage (2026-09-24)

> Análise do orquestrador, pedida pelo usuário depois de conferir a v2: *"me deu a impressão do
> deck ainda estar fraco"*. Fonte de **candidatas e calibragem**, não de veredito (regra 1): toda
> carta abaixo passa pela ficha F1–F7 e pelos 2+ pontos de sinergia na fase responsável antes de
> entrar. Oracle puxado no `mtgdb` na hora (regra 6). Preços: LigaMagic (menor), cotações de
> 2026-09-24, registradas com `mtgdb prices -add`.

## 1. Conclusão em uma linha

A v2 usa o Phlage em **um** dos seus dois modos. A comunidade o joga como **motor de reanimação**:
cada mágica de 1–2 manas que devolve o Phlage do cemitério é um [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix), e ele volta
sozinho para o cemitério, pronto para a próxima. A v2 **proibiu** esse pacote com base numa
leitura errada da carta.

## 2. O erro na trava mecânica 2

`decisions.md` (Fase 1) e `02-theme.md` §L1 registraram:

> *"Só vale conjurar por escape: reanimar ou blinkar faz o gatilho de sacrifício pegar.
> Nenhum pacote de reanimação ou blink entra neste deck."*

A premissa está certa, mas a conclusão não. O Phlage tem **duas** habilidades de entrada
independentes. Ruling oficial (WotC, 2024-06-07):

> *"Phlage's second ability triggers when it enters the battlefield, even if it didn't escape."*

Reanimado, ele faz **3 de dano + 3 de vida**, é sacrificado e **volta ao cemitério**. O sacrifício
não custa nada: é o que o deixa pronto para a próxima reanimação. A própria trava 3 registrou que
o dano dispara "escapado ou não". A trava 2 descartou o pacote apesar disso.

**Consequência para o ritmo (a queixa central):**

| Plano | 1º Helix repetível | Custo por Helix |
|---|---|---|
| v2: só escape | T6–T7 (6 cartas no cemitério + `{R}{R}{W}{W}`) | 4 manas + 5 cartas exiladas, uma vez |
| reanimação | **T4** (o Phlage já está no cemitério desde o T3) | **1–2 manas**, sem exilar nada |

Os dois planos convivem: a reanimação alimenta o cemitério para o escape, e o Phlage escapado que
morre volta a ser alvo de reanimação.

Probabilidade de ter ao menos um efeito de reanimação até o T4 (10 cartas vistas, hipergeométrica
sobre 99): **6 efeitos 48%** · **8 efeitos 59%** · **10 efeitos 67%** · **12 efeitos 74%**.

## 3. O que as listas mostram

**EDHREC** (`get_edhrec_recommendations`, limit 20; só a nota de sinergia vale, a inclusão vem
quebrada). O topo de `High Synergy Cards` é quase inteiro de reanimação ou de dobra de entrada:
[**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) 0,64 · [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) 0,61 · [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) 0,58 · [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) 0,57 ·
[**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness) 0,51 · [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) 0,50 · [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting) 0,48 · [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) 0,48 ·
[**Panharmonicon**](https://www.ligamagic.com.br/?view=cards/card&card=Panharmonicon) 0,47 · [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) 0,44. Nos `Top Cards`: [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) 0,44 · [**The Gaffer**](https://www.ligamagic.com.br/?view=cards/card&card=The+Gaffer) 0,42 ·
[**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) 0,41 · [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) 0,39.

**Archidekt, bracket 3** (37 listas; abri as 4 mais relevantes, comandante confirmado):

| Lista | Arquétipo | Serve para nós? |
|---|---|---|
| Usain Bolt (#8068510, 1.389 views) | loop de reanimação + dobra de ETB | **sim**: é o plano |
| Budget Phlag (#25832205) | loop de reanimação barato: [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle), [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth), [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan), [**Bishop of Rebirth**](https://www.ligamagic.com.br/?view=cards/card&card=Bishop+of+Rebirth) | **sim**: a referência mais próxima do teto |
| Forth Phlage! (#14466984, 1.620 views) | hatebears/stax (Drannith, Hushbringer, Grand Abolisher) | não: regra 11 |
| Phlage Boros Combo (#16911901) | turbo com One Ring, Smothering Tithe, Strip Mine, Wheel | não: preço e regra 11 |

**Onde a v2 já coincide com as listas de loop (15 cartas):** o esqueleto. Rochas ([**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet),
[**Boros Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Signet), [**Talisman of Conviction**](https://www.ligamagic.com.br/?view=cards/card&card=Talisman+of+Conviction)), terrenos ([**Battlefield Forge**](https://www.ligamagic.com.br/?view=cards/card&card=Battlefield+Forge), [**Clifftop Retreat**](https://www.ligamagic.com.br/?view=cards/card&card=Clifftop+Retreat), Command
Tower, [**Furycalm Snarl**](https://www.ligamagic.com.br/?view=cards/card&card=Furycalm+Snarl), [**Radiant Summit**](https://www.ligamagic.com.br/?view=cards/card&card=Radiant+Summit), [**Rugged Prairie**](https://www.ligamagic.com.br/?view=cards/card&card=Rugged+Prairie)), loots ([**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), Thrill of
Possibility), [**Chaos Warp**](https://www.ligamagic.com.br/?view=cards/card&card=Chaos+Warp), [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight), [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun), [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends).
**Nenhuma peça do motor.** [**Tome of Legends**](https://www.ligamagic.com.br/?view=cards/card&card=Tome+of+Legends) ganha página a cada entrada, então cada reanimação
passa a carregá-lo também.

## 4. Candidatas cotadas

Nenhuma está na caixa (`mtgdb collection`). A única peça de reanimação do pool sem custo é
[**Emerge from the Cocoon**](https://www.ligamagic.com.br/?view=cards/card&card=Emerge+from+the+Cocoon) (`{4}{W}`, devolve qualquer criatura + 3 de vida): cara para Helix, útil
para devolver [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) ou [**Torbran, Thane of Red Fell**](https://www.ligamagic.com.br/?view=cards/card&card=Torbran%2C+Thane+of+Red+Fell).

### 4.1 Reanimação barata (o motor)

| Carta | Custo | O que faz com o Phlage | R$ |
|---|---|---|---|
| [**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand) | `{W}` feitiço | Helix por 1 mana | 0,45 |
| [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness) | `{1}{W}` feitiço | Helix; o contador de voar **desliga** o [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) nessa volta | 0,15 |
| [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant) | `{1}{W}` feitiço | Helix | 0,20 |
| [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission) | `{1}{W}` feitiço | Helix; também devolve rocha | 1,00 |
| [**Patch Up**](https://www.ligamagic.com.br/?view=cards/card&card=Patch+Up) | `{2}{W}` feitiço | Helix + até 2 outras de MV total ≤ 3 (sobra 0 com o Phlage) | 0,23 |
| [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation) | `{2}{W}` feitiço, flashback `{4}{W}` | **2 Helix** por carta (uma de cada conjuração); devolve também rocha ou terreno | 3,50 |
| [**Brought Back**](https://www.ligamagic.com.br/?view=cards/card&card=Brought+Back) | `{W}{W}` instantâneo | só o que foi ao cemitério **neste turno**: Helix no turno em que ele se sacrifica | 3,00 |
| [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) | `{1}{W}` encantamento | pré-pago: devolve o Phlage quando ele morrer | 0,80 |

### 4.2 Motores repetíveis

| Carta | Custo | Como repete o Helix | R$ |
|---|---|---|---|
| [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle) | `{3}{W}` 2/2 voador | **cada mágica histórica** (artefato ou lendária) devolve o Phlage. A v2 tem 10 artefatos + 4 lendárias | 0,60 |
| [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) | `{4}{W}{W}` 6/6 | ao entrar **e a cada ataque**: Phlage, rocha, [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing), [**Glass Casket**](https://www.ligamagic.com.br/?view=cards/card&card=Glass+Casket) | 2,25 |
| [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) | `{2}{W}{W}` 3/4 voador | toda vez que o Phlage se sacrifica, volta com voar: **2 Helix por entrada** | 6,74 |
| [**Bishop of Rebirth**](https://www.ligamagic.com.br/?view=cards/card&card=Bishop+of+Rebirth) | `{3}{W}{W}` 3/4 | a cada ataque | 0,45 |
| [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) | `{1}{R}{W}` 3/3 trample | ao causar 3+ de dano de combate | 0,49 |
| [**Karmic Guide**](https://www.ligamagic.com.br/?view=cards/card&card=Karmic+Guide) | `{3}{W}{W}` | ao entrar (echo) | 2,66 |

### 4.3 Multiplicador e compra

| Carta | Custo | Função | R$ |
|---|---|---|---|
| [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) | `{2}{R}{R}` 3/4 | o Phlage é Giant: o Helix vira **6** de dano (a vida não dobra). [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan) também é Giant | 0,24 |
| [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome) | `{2}{W}` encantamento | compra 1 por turno em que o Phlage (MV 3) entra: o motor vira draw | 7,88 |
| [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) | `{2}{W}` 3/3 | compra no fim do turno se ganhou 3+ de vida: cada Helix basta | 11,88 |
| [**Pursue the Past**](https://www.ligamagic.com.br/?view=cards/card&card=Pursue+the+Past) | `{R}{W}` feitiço | loot + 2 de vida, com flashback | 0,39 |
| [**Ark of Hunger**](https://www.ligamagic.com.br/?view=cards/card&card=Ark+of+Hunger) | `{2}{R}{W}` | 1 a cada oponente sempre que carta sai do cemitério (escape, flashback, reanimação) | 0,45 |

### 4.4 Fora, com motivo

| Carta | Motivo |
|---|---|
| [**Celestine, the Living Saint**](https://www.ligamagic.com.br/?view=cards/card&card=Celestine%2C+the+Living+Saint) | R$ 110,38: 55% do teto |
| [**The Gaffer**](https://www.ligamagic.com.br/?view=cards/card&card=The+Gaffer) | R$ 63,75. [**Haliya, Guided by Light**](https://www.ligamagic.com.br/?view=cards/card&card=Haliya%2C+Guided+by+Light) faz a mesma coisa (compra com 3+ de vida) a R$ 11,88 |
| [**Panharmonicon**](https://www.ligamagic.com.br/?view=cards/card&card=Panharmonicon) | R$ 34,45: dobraria o Helix, mas custa 1,5× a folga inteira |
| [**Staff of the Storyteller**](https://www.ligamagic.com.br/?view=cards/card&card=Staff+of+the+Storyteller) | R$ 9,85; só conta fichas, e o deck não faz fichas |
| [**Cloudshift**](https://www.ligamagic.com.br/?view=cards/card&card=Cloudshift), [**Ephemerate**](https://www.ligamagic.com.br/?view=cards/card&card=Ephemerate), [**Flicker of Fate**](https://www.ligamagic.com.br/?view=cards/card&card=Flicker+of+Fate) | blink **só funciona em resposta ao gatilho de sacrifício** (o Phlage está em campo por um instante). No Phlage escapado, faz ele perder o escape e ser sacrificado. Pior que qualquer reanimação da §4.1 |
| pacote hatebear/stax, turbo de One Ring/Wheel | regra 11 e preço |

## 5. Orçamento

Um núcleo de 12 peças ([**Helping Hand**](https://www.ligamagic.com.br/?view=cards/card&card=Helping+Hand), [**Recommission**](https://www.ligamagic.com.br/?view=cards/card&card=Recommission), [**Return Triumphant**](https://www.ligamagic.com.br/?view=cards/card&card=Return+Triumphant), [**Call a Surprise Witness**](https://www.ligamagic.com.br/?view=cards/card&card=Call+a+Surprise+Witness),
[**Patch Up**](https://www.ligamagic.com.br/?view=cards/card&card=Patch+Up), [**Sevinne's Reclamation**](https://www.ligamagic.com.br/?view=cards/card&card=Sevinne%27s+Reclamation), [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle), [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan), [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth), [**Bishop of Rebirth**](https://www.ligamagic.com.br/?view=cards/card&card=Bishop+of+Rebirth),
[**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer), [**Tocasia's Welcome**](https://www.ligamagic.com.br/?view=cards/card&card=Tocasia%27s+Welcome)) soma **R$ 23,69**, contra **R$ 23,29 de folga**. Isso
**antes** de contar o que sai: cada corte devolve o preço da carta cortada.

**Onde a Fase 2 deve procurar os slots.** São hipóteses de busca, não vereditos: nenhum corte sai
daqui sem a ficha F1–F7 (regra 4).
- **Rota R2 de pingadores** ([**Guttersnipe**](https://www.ligamagic.com.br/?view=cards/card&card=Guttersnipe) R$ 5,40, [**Thermo-Alchemist**](https://www.ligamagic.com.br/?view=cards/card&card=Thermo-Alchemist), [**Erebor Flamesmith**](https://www.ligamagic.com.br/?view=cards/card&card=Erebor+Flamesmith), Firebrand
  Archer). A própria Fase 7 mediu 0,8–1,2 mágicas por turno, contra as 3–5 que a rota supunha.
- **Parte do combustível de escape** (14 slots entre loots e motores). A reanimação não precisa de
  cemitério cheio, só do Phlage lá. Os loots continuam valendo por filtro e pelo [**Inti, Seneschal of the Sun**](https://www.ligamagic.com.br/?view=cards/card&card=Inti%2C+Seneschal+of+the+Sun).
- [**Crackle with Power**](https://www.ligamagic.com.br/?view=cards/card&card=Crackle+with+Power) (R$ 26,40, a compra mais cara), se o fecho passar a ser Helix repetido
  multiplicado.

## 6. O que isso muda nas decisões da §0 do report

- **0.1 Ritmo:** é a pergunta que este pacote ataca de frente. O relógio da v2 foi medido com o
  primeiro Helix repetível no T6–T7; com reanimação, ele vem do T4. O `wincon-tester` precisa
  re-simular.
- **0.2 Mão sem criatura:** [**Teshar, Ancestor's Apostle**](https://www.ligamagic.com.br/?view=cards/card&card=Teshar%2C+Ancestor%27s+Apostle), [**Sun Titan**](https://www.ligamagic.com.br/?view=cards/card&card=Sun+Titan), [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth), [**Bishop of Rebirth**](https://www.ligamagic.com.br/?view=cards/card&card=Bishop+of+Rebirth), [**Venerable Warsinger**](https://www.ligamagic.com.br/?view=cards/card&card=Venerable+Warsinger) e [**Calamity Bearer**](https://www.ligamagic.com.br/?view=cards/card&card=Calamity+Bearer) são
  criaturas. O número sobe sem custar função.
- **0.4 Proteção:** [**Angelic Renewal**](https://www.ligamagic.com.br/?view=cards/card&card=Angelic+Renewal) e [**Luminous Broodmoth**](https://www.ligamagic.com.br/?view=cards/card&card=Luminous+Broodmoth) dão resiliência a remoção pontual no Phlage
  escapado.
