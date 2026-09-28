# Condições de Vitória e Testes — Satoru Umezawa (v1 · build)

Fase 7 · `wincon-tester` · 2026-09-28. Lista analisada: rascunho consolidado pelo orquestrador (99 + comandante),
**já com o swap Hedron Crawler → High Tide**, total R$ 179,81 (LigaMagic, menor entre edições legais, cotações de
2026-09-27). Folga no teto de R$ 200: **R$ 20,19**.
Oracle e rulings puxados nesta sessão com `bin/mtgdb` (todas as 66 cartas não básicas citadas). Regras de jogo
conferidas na Comprehensive Rules (`get_rule`), com o número na coluna "Confirmado por".
Game Changers (radar do EDHREC, Fases 3 e 5: Cyclonic Rift, Fierce Guardianship, Force of Will, Rhystic Study):
**nenhuma na lista, nenhuma proposta aqui.** A cota do bracket está zerada.

---

## Caminhos de vitória atuais

Meta de goldfish: **120 de dano** (3 oponentes × 40) até o turno 7.

| Caminho | Cartas envolvidas | Turno estimado (goldfish) | Consistência |
|---|---|---|---|
| **A. Pressão de bichões evasivos via ninjutsu** (plano principal) | Satoru Umezawa + 16 ativadores evasivos (Slither Blade, Changeling Outcast, Invisible Stalker, Dimir Infiltrator, Faerie Seer, Ornithopter etc.) → bichões de 5–7 de poder por `{2}{U}{B}`: Rune-Scarred Demon 6/6 voador, Marang River Regent 6/7 voador, Archfiend of Depravity 5/4 voador, Noxious Gearhulk 5/4 menace, Sepulchral Primordial 5/4 intimidate, Grave Titan 6/6 (+2 Zombies no ETB e **a cada ataque**), Starwinder 7/7, Gyruda 6/6, Scourge of Fleets 6/6. Evasão para reciclar: Tetsuko, Whispersilk Cloak, Access Tunnel, Rogue's Passage. Mana: High Tide (2 ninjutsus no turno), Silver-Fur Master (−1 por ninjutsu), Prosperous Thief (tesouros). | **T8–T9** para os 120. **T6** para matar **um** oponente de 40. Curva boa com Grave Titan no T4: ~9 / 26 / 54 / 99 de dano acumulado nos T4–T7, kill no T8 | **Alta** para montar a mesa (16 ativadores, 15 alvos, Satoru por 3). O relógio é que é lento: **não bate o T7 sozinho**. |
| **B. Loop Drake + Thousand-Faced Shadow + Silver-Fur Master** (infinito determinístico) | Satoru Umezawa (campo) + Silver-Fur Master (campo) + Peregrine Drake + Thousand-Faced Shadow (mão) + **2 atacantes não bloqueados** + 5 terrenos + 3 manas para começar | T5–T6 quando montado | **Baixa**: peças na mão até o T7 em **~7%** dos jogos sem tutor e **~17%** contando Rune-Scarred Demon e Dimir Infiltrator (→ Silver-Fur) (Monte Carlo de 200 mil mãos, compra + impulse do Satoru; não conferiu mana nem terrenos, então é teto) |
| **C. Loop Drake + High Tide** (ETB infinito do parceiro) | Satoru Umezawa (campo) + Peregrine Drake + High Tide + **qualquer** criatura na mão como parceira + 1 atacante não bloqueado + 5 terrenos (≥3 de tipo Island, ≥2 que gerem B) | T5–T6 quando montado | **Baixa**, dentro dos ~17% acima. A saída depende do parceiro: **Thousand-Faced Shadow** = fichas infinitas (kill no combate); **Grave Titan** = Zombies infinitos (kill no turno seguinte); **Gyruda** = todos os grimórios moídos (os oponentes perdem ao comprar, antes do seu próximo turno). Com a entrada proposta abaixo, **Gray Merchant** = dreno infinito, **kill imediato** |

**Leitura honesta:** o deck tem **um** caminho consistente (A), que fecha um turno ou dois **depois** da meta, e
**um** combo de verdade (B/C, mesma espinha), que é bônus de ~1 jogo em 6. "Mesa grande de bichões" não é wincon por
si só. O que a converte em vitória é o volume de dano evasivo, e ele só chega aos 120 no T8. A lacuna é **alcance**
(dano que não depende de atacar com tudo), e é ela que a entrada proposta cobre.

---

## Motor Peregrine Drake — análise do loop (item 2 do orquestrador)

**A Fase 4 acertou o caso isolado e errou a conclusão geral.** Drake + Silver-Fur Master sozinhos **não** fecham
loop: um ciclo completo exige **dois** ninjutsus (o do Drake e o do parceiro que o devolve à mão).

| Configuração | Custo por ciclo | Mana desvirada por ciclo | Saldo | Infinito? |
|---|---|---|---|---|
| Drake + parceiro, sem nada | 4 + 4 = 8 | 5 | **−3** | não |
| Drake + parceiro + **Silver-Fur Master** | 3 + 3 = 6 | 5 | **−1** | não (é o que a Fase 4 viu) |
| Drake + **Thousand-Faced Shadow** + parceiro + **Silver-Fur** | 3 + 3 + 3 = 9 | 5 + 5 (Drake + ficha-cópia do Drake) | **+1** | **sim** |
| Drake + parceiro + **High Tide** (3 Islands + Swamp + Sunken Hollow desvirados) | 4 + 4 = 8 | 3×2 + 1 + 2 = 9 (2 de B) | **+1** | **sim** |
| Drake + parceiro + **High Tide** (3 Islands + 2 Swamps) | 8 | 6 + 2 = 8 | **0** | **sim** (repete sem fim, sem mana extra) |
| Drake + parceiro + High Tide + Silver-Fur | 6 | 8–10 | **+2 a +4** | **sim**, mana infinita |

### Sequência B (Shadow + Silver-Fur), toda na etapa de declarar bloqueadores

Estado inicial: dois atacantes não bloqueados (A1, A2), Drake e Shadow na mão, Satoru e Silver-Fur em campo.

1. Ninjutsu do **Drake** (`{1}{U}{B}`), devolvendo A1. ETB: desvira 5 terrenos (**+2**).
2. Ninjutsu da **Shadow** (`{1}{U}{U}` nativo, com desconto), devolvendo A2. A Shadow entrou da mão atacando e
   cria **ficha-cópia do Drake**, que entra atacando e desvira mais 5 (**+2**).
3. Ninjutsu de **A1** (`{1}{U}{B}` do Satoru), devolvendo o **Drake** à mão (**−3**). O ETB de A1 dispara.
4. Repete: Drake devolvendo a Shadow, Shadow devolvendo A1 ou A2, a outra carta devolvendo o Drake.

Cada volta: **+1 mana, +1 ficha de Drake 2/3 voadora atacando e não bloqueada, +1 ETB do parceiro.** A cada ficha
você escolhe qual oponente ela ataca (ruling da Shadow). Com 60 fichas = 120 de dano no mesmo combate.

### Sequência C (High Tide), mais simples

Satoru em campo, High Tide resolvida no turno, um atacante A não bloqueado, Drake na mão:
ninjutsu do Drake devolvendo A (desvira 5; as Islands rendem 2 cada) → ninjutsu de A devolvendo o Drake → repete.
Pede **só 1 atacante**, e o parceiro pode ser qualquer criatura da mão, trocado a cada volta.
Exemplo: 10 voltas de Meteor Golem limpam os permanentes não-terreno dos oponentes, e depois vêm N voltas de
Grave Titan ou a Shadow (se estiver na mão). **Com Spectral Sailor em campo** e saldo positivo (Sunken Hollow entre
os desvirados, ou Silver-Fur), a mana infinita vira **compra à vontade** (`{3}{U}`: compre) e acha a peça que falta.

### Travas e cuidados (anotar no goldfish)

- **Mana some entre etapas** (CR 106.4): o loop inteiro acontece dentro da etapa de declarar bloqueadores.
- **Satoru, the Infiltrator em campo obriga a comprar** ("draw a card", sem "may") a cada criatura não-ficha que
  entra sem ser conjurada, ou seja, **3 cartas por volta** na sequência B. Com ele em campo, o loop fica limitado
  ao tamanho do grimório, e o combo precisa parar antes de o deck acabar. Gyruda como parceira mói o **seu**
  grimório também (4 por ETB). Moon-Circuit Hacker, Reconnaissance Mission, Bident e Starwinder têm "may", sem
  problema.
- **O Satoru é peça viva do loop.** Remover o Satoru não anula o ninjutsu que já foi ativado (ruling), mas corta os
  seguintes que ele concede (Drake e parceiro). Shadow e Silver-Fur têm ninjutsu próprio e seguem funcionando.
- **Cores:** cada ninjutsu do Satoru pede `{B}`. Na sequência C sem Sunken Hollow, escolha **2 terrenos de B** entre
  os 5 desvirados, senão o B acaba antes do U.

### Existe algum outro loop na lista? (regra 11)

Oracle varrido: nenhum outro. Kaito, Dancing Shadow devolve **uma** criatura por combate. Whelming Wave e Marang
são pontuais. Hostage Taker e Silumgar não repetem sozinhos. **Nenhuma peça de lock, stax ou MLD.**
O loop B/C é **combo de vitória, não lock**: termina o jogo naquele combate (ou na volta seguinte, com Grave Titan
ou Gyruda). A regra 11 não o proíbe.

**Ponto para o usuário (bracket):** com High Tide, o combo passa a ser **2 cartas das 99 + o comandante** (Drake +
High Tide), a partir do T5. Pela minha leitura das diretrizes de bracket da WotC, bracket 3 não aceita combo
infinito de duas cartas **cedo** no jogo, e bracket 4 aceita. O briefing permite 3–4, então **cabe**, mas coloca o
deck no **topo** da faixa. A versão B (Shadow + Silver-Fur) são 3 cartas + comandante + 2 atacantes, dentro de
bracket 3 sem discussão. Se o usuário quiser ficar firme em bracket 3, a saída é trocar a High Tide (volta o
Hedron Crawler da caixa, ou Everflowing Chalice a R$ 0,99). O loop B continua existindo. **Decisão do usuário,
não minha.**

---

## Finishers recomendados

Varredura das sobressalentes (regra 7, `mtgdb collection -list`, 182 cartas, 39 não-terrenos em UB que não estão
no deck): **nenhum finisher.** Dispensas por eixo da ficha:
- **Wall of Lost Thoughts, Dream Twist:** moem 3–4 por uso. Com 3 grimórios de ~85 cartas, não é alcance (F1).
- **Trailblazer's Torch:** a iniciativa é valor por masmorra, não fecha jogo. O equip disputa a mana do ninjutsu (F7).
- **Passenger Ferry, Dependable Quinjet:** o crew tapa o atacante que deveria ativar o ninjutsu (F7, mesmo motivo
  da Fase 2).
- **I Am Iron Man, Rampart Hunter:** 1 ponto de sinergia só (corpo/voo temporário). Não convertem mesa em vitória.
- **Culling Dais, Lux Cannon:** compra e remoção lenta, fora da função de finisher.

Por isso a entrada é compra. Busca: `mtgdb search` por dreno "each opponent loses" em `id<=UB`, cruzada com a regra
do Satoru (ETB que **não** depende de "cast"). A candidata é uma só:

| Carta | CMC | Como fecha o jogo | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|
| **Gray Merchant of Asphodel** | 5 (ninjutsu 4) | ETB: cada oponente perde X, onde X = sua devoção a preto, e você ganha o total. Numa mesa típica do T5–T6 (Satoru 1 + Silver-Fur ou Satoru, the Infiltrator 1 + Gary 2 + 1 bichão preto 2), **X ≈ 6 → 18 de vida tirada por ETB**, sem precisar de dano de combate | (1) ETB sem "cast": o oracle é "When this creature enters", gatilho de entrada e não de conjuração, então entra por ninjutsu a 4 e dispara igual; (2) **fecha o loop C na hora**: com High Tide + Drake, dreno infinito que ganha **antes** do dano de combate, sem Shadow; (3) Thousand-Faced Shadow copia o Gary: a ficha tem o custo `{3}{B}{B}`, então soma +2 de devoção e dispara um 2º dreno maior; (4) reciclável: entra não bloqueado, conecta, e Kaito, Dancing Shadow o devolve à mão para o ninjutsu do turno seguinte (Access Tunnel e Rogue's Passage também, poder 2); (5) devoção vem de Satoru, Silver-Fur, Satoru the Infiltrator, Changeling Outcast, Dimir Infiltrator, os dois Kaitos e os bichões pretos (Chupacabra, Gearhulk, Grave Titan, Rune-Scarred, Archfiend, Gyruda híbrido) | não | **a cotar** |

**Efeito no relógio:** um Gary no T5 e outro reciclado no T7 somam ~36 de dano "fora do combate". A curva boa do
caminho A (99 no T7) passa de 120, e a mediana (T8–T9) cai **um turno**. Não garante o T7 sozinho, mas é a única
carta que encosta o deck na meta sem mudar o plano.

### Swap proposto (o deck já tem 99)

| Sai | Entra | Status |
|---|---|---|
| **Sepulchral Primordial** (R$ 0,75) | **Gray Merchant of Asphodel** (a cotar) | **corte condicionado**: a função que fica descoberta ("alvo de ninjutsu com ETB de reanimação do cemitério adversário") é **temática** (Fase 2), fora da minha lente. Decisão do orquestrador/usuário |

**Ficha da que sai — Sepulchral Primordial** (oracle desta sessão):
- **F1 Texto:** Intimidate. ETB: para cada oponente, põe até 1 criatura do cemitério dele no campo, sob seu controle.
- **F2 Corpo:** 5/4 com intimidate (só é bloqueado por criatura-artefato ou preta). Conecta e pode ser reciclado
  pelo ninjutsu no turno seguinte.
- **F3 Tipo:** Creature — Avatar. Soma **BB = 2 de devoção a preto**.
- **F4 Receptor:** alvo de cópia da Thousand-Faced Shadow (até 6 reanimações).
- **F5 Facilitador:** indireto: as criaturas roubadas podem ser qualquer coisa.
- **F6 Curva:** 4 via ninjutsu a partir do T4; 7 conjurada.
- **F7 Atrito:** o ETB depende do cemitério dos oponentes, que está vazio no T4–T5 e costuma ter 2–3 alvos do T6
  em diante. **No goldfish vale zero**, e o teste subestima a carta. Isso fica declarado.

**Protocolo §2:**

| Função | Quem cobre depois do corte |
|---|---|
| alvo de ninjutsu com ETB forte | 14 alvos seguem (Chupacabra, Hostage Taker, Gearhulk, Grave Titan, Dream Eater, Silumgar, Gyruda, Marang, Rune-Scarred, Meteor Golem, Scourge, Archfiend, Starwinder, Drake) + o próprio Gary |
| corpo evasivo 5+ reciclável | Rune-Scarred Demon, Marang River Regent, Archfiend (voadores) · Noxious Gearhulk (menace) |
| devoção BB | Gary traz os mesmos BB quando está em campo |
| alvo de cópia da Shadow | Gary copiado = 2º dreno, maior que o 1º |
| **reanimação/roubo do cemitério dos oponentes** | **descoberta** em parte. Gyruda cobre um pedaço (mói 4 de cada jogador e pega uma criatura de MV par entre **todas** as moídas). Custo aceito: perde-se o pico de valor do late game em troca de alcance |

**Simetria (§3):** as duas avaliadas nas mesmas condições: Satoru em campo, entrada por ninjutsu a 4, sem proteção
própria, ETB dependente de estado de jogo (cemitério adversário × devoção própria). A Sepulchral tem evasão
(intimidate) e o Gary não. A diferença está declarada, e o Gary depende de Kaito, Tunnel, Passage ou Whispersilk
para reciclar. O Gary vence o teste porque **avança o relógio em toda partida**, e a Sepulchral **gera mesa**, que
o deck já gera de sobra (lacuna do caminho A = alcance, não valor).
Registro: nenhuma das duas consta de `decisions.md`.

**Alternativa de saída avaliada e recusada:** Dream Eater. É a peça de bounce flexível que a Fase 5 conta para
artefato, encantamento e planeswalker (buraco parcial em artefatos já declarado lá), então não corto.

**Orçamento:** Gary precisa caber em **R$ 20,19** (a Sepulchral devolve R$ 0,75 → folga de R$ 20,94).

---

## Combos (se houver)

| Combo | Peças | Resultado | Confirmado por (ruling / regra CR) |
|---|---|---|---|
| Loop B: Drake + Shadow + Silver-Fur | Satoru Umezawa, Silver-Fur Master (campo) · Peregrine Drake, Thousand-Faced Shadow (mão) · 2 atacantes não bloqueados · 5 terrenos | Mana infinita (U/B, só na etapa), fichas de Drake 2/3 voadoras infinitas **atacando e não bloqueadas**, ETB infinito de um parceiro → kill de mesa no mesmo combate | CR **702.49a/c** (custo e entrada do ninjutsu) · CR **508.4d** (criatura que entra atacando depois dos bloqueadores é "unblocked", então a ficha pode ser devolvida ou causar dano) · ruling da Shadow (a ficha copia o Drake e o ETB dispara; você escolhe o oponente atacado pela ficha) · ruling do Drake (desvira terrenos sem alvo) · ruling do Silver-Fur (o desconto vale para todo ninjutsu que você ativa) · CR **106.4** (mana esvazia entre etapas) |
| Loop C: Drake + High Tide | Satoru Umezawa (campo) · Peregrine Drake + High Tide · 1 atacante · 5 terrenos (≥3 tipo Island, ≥2 de B) | ETB infinito do parceiro (≥0 de mana por volta; +1 com Sunken Hollow; +2 a +4 com Silver-Fur). Kill com Shadow (fichas), Gary (dreno, se entrar), Grave Titan (turno seguinte) ou Gyruda (grimórios) | CR **702.49a/c** · ruling da High Tide (vale para todo terreno com tipo Island, inclusive Sunken Hollow, e para Island que entrou depois) · ruling do Drake · CR **106.4** |
| Gary + Shadow (não é infinito) | Gray Merchant (entrada proposta) + Thousand-Faced Shadow | 2 drenos no mesmo combate; o 2º conta a ficha (+2 de devoção) | ruling da Shadow (a ficha copia o custo de mana impresso) · ruling do Gary (conta a devoção na resolução; símbolo híbrido conta, ou seja, Gyruda soma 2) · CR **700.5** |

---

## Fragilidades (item 3)

| Fragilidade | Gravidade | O que o deck tem | Leitura |
|---|---|---|---|
| **Satoru removido** | **a maior** | Protegem **antes** da ativação: Siren Stormtamer, Old Fat Spider (hexproof), Diversion Unit, Counterspell, Negate, Disruption Protocol, Reasonable Doubt, Theorix Charm. Depois de ativado, remover o Satoru não anula o ninjutsu (ruling). Plano B sem Satoru: 5 ninjas nativos (Silver-Fur, Moon-Circuit, Ingenious, Prosperous Thief, Shadow) + os bichões de 4 conjuráveis (Chupacabra, Hostage Taker). Recolocar custa 5, depois 7 | Sem o Satoru, os bichões de 6–7 ficam mortos na mão. **É a principal fonte de mão morta**, e o goldfish precisa medir isso (partidas de estresse no protocolo) |
| Mesa de voadores (bloqueadores no ar) | média | Inbloqueáveis de verdade: Slither Blade, Changeling Outcast, Invisible Stalker, Dimir Infiltrator, Dimir Keyrune, a ficha do Kaito Shizuki. **Tetsuko** torna inbloqueáveis todos os ≤1 (Ornithopter, Seer, Sailor, Siren, Shadow, Diversion Unit, Pilgrim's Eye, H.E.R.B.I.E.). Whispersilk, Access Tunnel e Rogue's Passage. Limpeza: Archfiend (deixa 2 por oponente), Scourge of Fleets | ok. **Tetsuko é a carta que mais pesa aqui** |
| Falta de ativador | baixa | 16 evasivos + 5 situacionais | ok |
| Relógio lento (caminho A) | média | Grave Titan (dano cresce a cada ataque), High Tide (2 ninjutsus), Shadow (cópia de bichão) | é a lacuna do T7. Gary encurta 1 turno |
| Wipe de criaturas | média | Dimir Keyrune e Kaito Shizuki (−2) geram ativador pós-wipe; Access Tunnel e Rogue's Passage dão evasão a qualquer um; Whelming Wave poupa Gyruda, Scourge e Starwinder | ok |

---

## Archfiend of Depravity e a regra 11 (item 5)

Oracle: *"At the beginning of each opponent's end step, that player chooses up to two creatures they control, then
sacrifices the rest."* Rulings: a escolha é do oponente, na resolução, sem alvo.

**Leitura: aceitável em bracket 3–4. Não é lock nem stax pesado.**
- O oponente **joga normalmente**: compra, baixa terreno, conjura mágicas e planeswalkers, conjura criaturas e
  **usa todas no próprio turno**. Só no fim do turno dele fica com as 2 que escolheu. É teto de mesa, não trava.
- É **uma** criatura 5/4, sem hexproof nem indestrutível, que morre para qualquer remoção pontual. Não está na lista
  de Game Changers. Não tem recursão em loop no deck: o gatilho é de fim de turno, e reciclar o Archfiend pelo
  ninjutsu não o repete.
- Pune de verdade **decks de fichas** (go-wide). É a função pela qual a Fase 5 o contou como wipe recorrente, e é
  uma das respostas à "mesa de voadores" acima.
- **Recomendo informar o usuário** como "board control unilateral forte" (o tipo de carta que se avisa na conversa
  de pré-jogo de bracket 3), sem cortar.

---

## Protocolo de goldfishing

Para o usuário rodar com proxies (histórico dele: testar antes de comprar).

**Volume:** **10 partidas** (mínimo 5), sendo 8 normais e 2 de estresse. Planilha simples, uma linha por partida.

**Mulligan** (multiplayer: o 1º é grátis, CR 103.5c. Compre 7 de novo sem pôr carta no fundo):
- **Fica:** 3–5 terrenos com ≥1 fonte de U **e** ≥1 de B (Command Tower, Arcane Signet, Talisman, Dimir Signet
  contam), e pelo menos **1 ativador de mv ≤ 2** ou rocha de mv ≤ 2 que ponha o Satoru no T2–T3.
- **Volta:** 0–2 terrenos sem Sol Ring; 6+ terrenos; mão só de bichões sem ativador (o ninjutsu precisa de um
  atacante que conecte).

**Regras do goldfish para este deck:**
1. Sem oponente: **todo atacante passa sem bloqueio**. O ninjutsu é ativado na etapa de declarar bloqueadores.
2. Impulse do Satoru: **1× por turno**, mesmo com dois ninjutsus.
3. Mana esvazia entre etapas: High Tide e Drake geram mana que precisa ser gasta **na mesma etapa** (bloqueadores).
4. Bichões com gatilho "whenever attacks" (Grave Titan) **não** disparam no turno em que entram por ninjutsu, só a
   partir do turno seguinte, quando são declarados atacantes.
5. ETB de remoção sem alvo no goldfish (Chupacabra, Meteor Golem, Hostage Taker, Sepulchral) conta só o corpo.
   Anote "ETB vazio" para não inflar a mesa.
6. Combo B ou C montado **com as condições de mana conferidas** (5 terrenos; para o C, ≥3 Islands de tipo e ≥2 de B)
   = vitória naquele turno. Anote se veio de tutor (Rune-Scarred, Dimir Infiltrator).
7. Satoru, the Infiltrator em campo durante o loop: limite as voltas ao grimório (3 compras por volta no loop B).

**Registrar por partida:**

| Campo | O quê |
|---|---|
| Mulligans | quantos, e por quê |
| T do Satoru | turno em que ele resolveu |
| T do 1º ninjutsu | e qual bichão entrou |
| "Mesa boa" | 1º turno com Satoru + ≥2 ativadores + ≥5 manas |
| Dano acumulado T4 / T5 / T6 / T7 | total nos 3 oponentes |
| T de 40 de dano | 1º oponente morto (referência de jogo real) |
| **T de 120 de dano** | vitória projetada |
| Combo montado? | turno, qual (B/C), veio de tutor? |
| Trava de mana | T3 com < 3 terrenos, ou sem U ou B no T3 |
| Mão morta | nº de bichões presos na mão no T6 sem mana ou sem ativador |

**Partidas de estresse (2 das 10):**
- **S1 — Satoru removido:** quando ele resolver, suponha que morra na 1ª ativação de ninjutsu do turno seguinte, antes
  de resolver. Recoloque pela taxa (5, depois 7). Mede o plano B e a mão morta.
- **S2 — mesa de voadores:** suponha que cada oponente tenha 1 bloqueador voador 2/2 a partir do T3. Só passam
  inbloqueáveis, Tetsuko e seus ≤1, ou terreno de evasão. Mede a dependência da Tetsuko.

**Critérios de sucesso** (sobre as 8 normais):

| Marco | Meta |
|---|---|
| Satoru em campo até o T3 | ≥ 6/8 |
| 1º ninjutsu até o T4 | ≥ 5/8 |
| 40 de dano (1 oponente) até o T6 | ≥ 5/8 |
| **120 de dano até o T7** | **≥ 4/8** (meta do pipeline). Expectativa honesta com a lista atual: 1–2/8, e 3–4/8 com o Gary |
| Trava de mana | ≤ 2/8 |
| Mão morta (≥2 bichões presos no T6) | ≤ 2/8 |

**Como ler o resultado** (para o modo `post-goldfish`):
- 120 depois do T8 na maioria, sem travas → faltam alcance/finishers: Gary primeiro, e depois Fallen Shinobi
  (R$ 19,99, reserva da Fase 6) ou um 2º dreno.
- Trava de mana em ≥ 3/8 → devolver ao `manabase-engineer` (variante de 38 terrenos da Fase 6).
- Mão morta em ≥ 3/8 → excesso de bichões ou falta de ativador. Trocar um alvo por ativador ou ninja nativo
  (Ninja of the Deep Hours R$ 7,50; Cloud of Faeries R$ 0,83, encaminhada pela Fase 4).
- S1 sem nenhuma ação relevante em 3 turnos → reforçar proteção do Satoru (Fase 5).

---

## Ajustes pós-teste (quando aplicável)

Pré-teste (proposta desta fase, depende de aprovação):

| Sai | Entra | Motivo |
|---|---|---|
| Sepulchral Primordial (**corte condicionado**: função temática de reanimação do cemitério adversário fica parcialmente descoberta; Gyruda cobre em parte) | Gray Merchant of Asphodel (a cotar) | alcance fora do combate (~18 por ETB), fecha o loop C sem Shadow, encurta o caminho A em ~1 turno |

Pós-teste: preencher depois das 10 partidas.

---

## Pendências

1. **Cotar Gray Merchant of Asphodel** na LigaMagic (teto R$ 20,94 com a saída da Sepulchral). Sem cotação, o swap
   não fecha.
2. **Decisão de bracket do usuário:** High Tide + Drake é combo de 2 cartas das 99 + comandante a partir do T5.
   Aceitar (topo de 3–4) ou trocar a High Tide (volta Hedron Crawler da caixa ou Everflowing Chalice a R$ 0,99).
3. **Corrigir a nota da Fase 4** no report: "Drake + Silver-Fur não é infinito" vale só para a dupla. Com
   Thousand-Faced Shadow **ou** High Tide, o loop é infinito e determinístico.
4. **Archfiend of Depravity:** informar o usuário (aceitável, não é lock).
5. **Goldfishing ainda não rodado:** relógio acima é estimativa de curva (caminho A) e Monte Carlo de peças (combo),
   não resultado de teste. Rodar o protocolo antes da compra.
