# Análise Temática — Otharri, Suns' Glory (v1 · modo `improve`)

> **Passe 3 — devolução de 2026-09-23 (densidade de corpos).** O usuário imprimiu a lista v1 em
> proxy e jogou. Relato verbatim: *"fiquei com muita mágica instantânea e encantamentos na mão e
> poucas criaturas, senti lentidão para descer o comandante e quando desço e o perco, acabo não
> usando o custo fixo por ficar sem tokens de rebeldes"*. Nada foi comprado — **não há custo
> afundado e os slots são livres**.
>
> Esta devolução **não** refaz o eixo temático: a análise do comandante (§1), os termos de busca
> (§2), a reclassificação das 99 cartas (§3) e o pool (§4) do passe de 2026-09-19 **continuam
> válidos** e estão preservados abaixo, a partir da §1. O que muda está concentrado na **§0**,
> nova: correção da leitura de vigilância, meta de densidade de corpos com a conta
> hipergeométrica, pacote *For Mirrodin!*, varredura da caixa sob o novo critério e a lista
> entra/sai com ficha F1–F7 dos dois lados.
>
> Histórico: §1–§7 = passe de 2026-09-19 (re-pass com Otharri). O passe original, escrito para
> Tori D'Avenant, está em `02-theme--passe-tori.md`.

Fonte de todo texto oracle: `bin/mtgdb` (dump Scryfall), puxado nesta sessão. Rulings de
`bin/mtgdb rulings`. Preços em US$ são **estimativa (Scryfall)** e **não valem como régua** —
a régua é o menor valor da LigaMagic (regra 2).

---

# §0 — Devolução de 2026-09-23: densidade de corpos

## 0.1 A correção técnica que motiva o resto

A §11.1 do `report.md` afirma que *"vigilância na ficha é o que paga a recursão do Otharri"* e
mantém a `Intangible Virtue` por esse motivo. **A afirmação está errada como escrita.**

Texto do custo (oracle puxado nesta sessão):
`{2}{R}{W}, Tap an untapped Rebel you control: Return this card from your graveyard to the battlefield tapped.`

E o gatilho: `create a 2/2 red Rebel creature token that's tapped and attacking for each experience counter you have.`

Vigilância significa *"atacar não faz esta criatura virar"*. Ela **não desvira** nada. As fichas
do Otharri **entram viradas** — não foram viradas por atacar, já chegam assim —, então:

| Cenário | Há Rebel desvirado para pagar a recursão? |
|---|---|
| Otharri morre em combate no **primeiro** ataque | **Não.** As fichas daquele ataque entraram viradas; não existe ficha de turno anterior |
| Otharri morre num **wipe** com o board zerado | **Não** |
| Otharri morre a partir do **segundo** ataque, com fichas sobreviventes do ataque anterior | **Sim**, e é aqui — **e só aqui** — que a vigilância ajuda: as fichas antigas atacam sem virar e continuam disponíveis como custo |

Ou seja: a `Intangible Virtue` **melhora** a recursão a partir do 2º ataque, mas **não** cobre o
caso que o usuário viveu. E a única fonte de Rebel do deck além do comandante é a
`Hexgold Halberd`. **Consequência:** o deck precisa de Rebels **desvirados que não dependam de o
Otharri ter atacado antes** — é a função que a §0.4 preenche.

> A `Intangible Virtue` **continua no deck**, mas por outro motivo (ver §0.6): ela é o anthem mais
> barato do deck e alcança **todas** as fichas, não só as Rebel.

## 0.2 A conta: quantos corpos o deck precisa

**Definição usada.** *Fonte de corpo* = carta que, ao resolver, **põe pelo menos uma criatura no
campo**. Isso inclui criaturas, equipamentos *For Mirrodin!* (a ficha 2/2 Rebel vem junta e
desvirada), `Ancestral Blade`, `Staff of the Storyteller` (ETB cria um 1/1 voador) e
`Raise the Alarm`. **Não** inclui motores de turno seguinte (`Assemble the Legion`,
`Wedding Announcement`) nem fábricas ativadas (`The Circle of Loyalty`) — eles contam como
reconstrução pós-wipe, não como corpo na mão de abertura.

Hipergeométrica sobre 99 cartas, mão de 7 (o comandante está na zona de comando):

`P(0 corpos) = C(99−K, 7) / C(99, 7)`

| K (fontes de corpo) | P(mão de 7 **sem** corpo) | P(≥1) | P(≥2) | Corpos esperados/mão |
|---|---|---|---|---|
| **15** (só as criaturas da v1) | **30,42 %** | 69,58 % | 28,62 % | 1,06 |
| **19** (v1 pela definição acima) | **21,34 %** | 78,66 % | 40,31 % | 1,34 |
| 22 | 16,15 % | 83,85 % | 48,81 % | 1,56 |
| 24 | 13,33 % | 86,67 % | 54,21 % | 1,70 |
| **26 (meta proposta)** | **10,94 %** | **89,06 %** | **59,32 %** | **1,84** |
| 27 | 9,90 % | 90,10 % | 61,77 % | 1,91 |
| 30 | 7,25 % | 92,75 % | 68,60 % | 2,12 |

**Meta adotada: 26 fontes de corpo**, das quais **18 são cartas do tipo criatura**. Justificativa
de parar em 26 e não em 30:

- 26 já cumpre o alvo pedido (`P(sem corpo) < ~12 %`) com margem: **10,94 %**. Com o mulligan de
  Londres, a chance de duas mãos seguidas sem corpo cai para **1,2 %** (0,1094²).
- Passar de 26 exigiria cortar peça de função (draw 13 / ramp 11+2 / interação 14 / wipes 2), que
  a devolução proíbe explicitamente e que são as metas do pipeline.
- O outro lado da conta também melhora: instantâneos + feitiços + encantamentos caem de **30 para
  25**, ou seja, de **2,12 para 1,77** por mão. A razão corpo:mágica sai de **0,63 : 1** e vai para
  **1,04 : 1** — a inversão é o que o usuário pediu em palavras.

> **Se o slot da `Basri's Solidarity` for liberado** (ver §0.7), entra a `Patrolling Peacemaker` e
> a conta vai para **K = 27 → 9,90 %**.

## 0.3 Varredura da caixa **antes** do Scryfall (regra 7)

`bin/mtgdb collection -list` → 116 cartas. Filtro: identidade ⊆ RW **e** produz corpo. O critério
mudou desde o passe de 2026-09-19: lá, "corpo sem gatilho" reprovava na regra 3 por ter 1 ponto de
sinergia. **Agora o corpo em si é a função escassa** — o que muda o veredito de algumas cartas, e
isso está declarado carta a carta abaixo.

| Carta da caixa | CMC | Tipo | Veredito | Por quê |
|---|---|---|---|---|
| **Goblin Rabblemaster** | 3 | Creature — Goblin Warrior 2/2 | **ENTRA** | ficha §0.5 |
| **Breeches, Eager Pillager** | 3 | Leg. Creature — Goblin Pirate 3/3 | **ENTRA** | ficha §0.5 |
| **Zookeeper Mechan** | 2 | Artifact Creature — Robot 1/3 | **ENTRA** | ficha §0.5 |
| **Patrolling Peacemaker** | 3 | Artifact Creature — Robot Soldier 2/2 | **reserva nº 1** | ficha §0.5; só falta slot |
| **Goblin Gathering** | 3 | Sorcery | reserva nº 2 | 2 fichas 1/1 por 3 manas. `Raise the Alarm` faz o mesmo por **2** e em **instant**; entra se abrir um 8º slot |
| **Knight Watch** | 5 | Sorcery | reserva nº 3 | 2 corpos 2/2 com vigilância por **5 manas** — é o ponto da curva que o usuário chamou de lento, e os Knights **não são Rebel**, então não pagam a recursão. Foi listada como reserva já no passe anterior (§4.8 nº 42) e continua reserva |
| **Redcap Thief** | 3 | Creature — Goblin Rogue 2/3 | reserva nº 4 | corpo + Treasure (ramp + artefato temporário). Perde para Breeches no mesmo custo: Breeches tem 3/3, first strike e escolhe Treasure/impulse/"não pode bloquear" **todo ataque** |
| **Lesser Masticore** | 2 | Artifact Creature 2/2 | dispensada | corpo + artefato + persist (resiliência a wipe, dor 4). **Custo adicional: descartar uma carta.** O deck acabou de ser diagnosticado com excesso de mágica situacional na mão, então o descarte tem upside — mas é desvantagem de carta real num deck cujo gap nº 1 é saque. Fica fora enquanto houver corpo sem custo |
| **Leonin Abunas** | 4 | Creature — Cat Cleric 2/5 | dispensada | corpo + hexproof aos 21 artefatos. 4 manas por um 2/5 que não ataca e não faz ficha; o deck é de ataque. 2 pontos, mas ambos defensivos |
| **Mudbutton Clanger** | 1 | Creature — Goblin Warrior 1/1 | dispensada | 1/1 por `{R}`; kinship só olha Goblin (o deck tem 2 após as entradas). **1 ponto** — reprova na regra 3 mesmo com o critério afrouxado |
| **Youthful Knight**, **Embereth Paladin**, **Bogardan Lancer** | 2–4 | Creatures | dispensadas (**veredito mantido**) | eram candidatas por contagem de Knight e por *first strike → Kwende*; os dois eixos caíram em 2026-09-19. Como **corpo puro** perdem para as entradas propostas: nenhuma traz Rebel desvirado, artefato para metalcraft, mana ou saque junto. 1 ponto cada |
| **Thopter Fabricator**, **Master's Councillors**, **Elrond, Moon-Reader**, **Somber Hoverguard**, **Thrummingbird** | — | — | fora da identidade | identidade **U** — ilegais no deck |
| **Wilderland Scrounger**, **Grumgully**, **Zhur-Taa Goblin** | — | — | fora da identidade | identidade **G/RG** |

**Resultado da varredura:** 3 entradas de custo zero (`Goblin Rabblemaster`, `Breeches`,
`Zookeeper Mechan`) + 4 reservas. Tudo o que for **comprado** na §0.4 tem, abaixo, a carta
equivalente da caixa nomeada e o motivo escrito de ela não cobrir a função.

## 0.4 O pacote *For Mirrodin!* — por que ele é a resposta certa, e quais escolher

Todo equipamento *For Mirrodin!* resolve **três funções num slot**:

1. **Rebel 2/2 desvirado** — chega desvirado e já equipado; é o custo da recursão do Otharri
   disponível **no turno em que ele morre**, sem depender de ataque anterior (§0.1);
2. **artefato permanente** — alimenta o metalcraft do `Jor Kadeen` (3+ artefatos → `+3/+0` a
   **todas** as criaturas) e a contagem do `Luxknight Breacher`;
3. **corpo na mão** — conta na hipergeométrica da §0.2.

Ruling confirmado (`mtgdb rulings "Hexplate Wallbreaker"`, 2023-02-04): *"The Rebel enters the
battlefield as a 2/2 creature, then the Equipment becomes attached to it. Abilities that trigger
when a creature enters the battlefield see that a 2/2 creature entered."* → **a ficha dispara
`Valor in Akros`, `Warleader's Call`, `Belladonna Took`, `Staff of the Storyteller`,
`Tocasia's Welcome` e `Molten Gatekeeper`.** Quarta função.

Outro ruling do mesmo bloco: *"If the Rebel is destroyed, the Equipment stays on the battlefield"*
— o equipamento sobrevive ao wipe de criaturas e pode ser realocado.

### Avaliação dos 10 candidatos (não adotados em bloco)

| Carta | Custo | O que o equipado ganha | Veredito |
|---|---|---|---|
| **Barbed Batterfist** | `{1}{R}` | +1/−1 · equip `{1}` | **ENTRA** — o mais barato do pacote em mana e o equip mais barato do deck. Atrito declarado: a ficha vira **3/1** e morre para qualquer ping; com `Intangible Virtue` fica 4/2 |
| **Bladehold War-Whip** | `{1}{R}{W}` | **double strike** · equip de **outros** equipamentos custa `{1}` a menos | **ENTRA** — ficha 2/2 com double strike = 4 de dano por 3 manas, e cada anthem conta **duas vezes** nela. A redução de equip destrava `Sword of the Animist`/`Ancestral Blade`/`Hexgold Halberd` no mesmo turno. Atrito: **redundante com a `Neyali`** (que já dá double strike a *todas* as fichas atacantes) — mas cobre o que a Neyali não cobre (não-fichas: Otharri, Jor Kadeen) e funciona sem ela |
| **Glimmer Lens** | `{1}{W}` | — · **"sempre que o equipado e pelo menos outra criatura atacam, compre um card"** | **ENTRA** — é **saque repetível de graça** num deck que ataca com 4+ criaturas todo turno, embrulhado num corpo. Cobre simultaneamente a dor 2 e a densidade de corpo |
| **Hexplate Wallbreaker** | `{3}{R}{R}` | +2/+2 · **ao atacar (1ª fase de combate): desvira todos os atacantes e há um combate adicional** | **ENTRA (com ressalva de preço)** — combate extra **repetível todo turno**, e ele traz a criatura que o carrega. Com Otharri: ataca → 1º gatilho (contador + leva de fichas) → desvira **todos os atacantes, inclusive as fichas que entraram viradas** → 2º combate → **2º gatilho do Otharri**, com o contador já maior. É a rota R4 em modo motor. Atritos: CMC 5 disputa o turno do Otharri; o ruling de 2023-02-04 confirma que **não há fase principal entre combates** (não dá para equipar no meio); e US$ 7,72 de estimativa é o maior risco de preço da lista |
| Vulshok Splitter | `{3}{R}` | +2/+0 · equip `{2}{R}` | fora — 4 manas pelo mesmo corpo que o Batterfist dá por 2; só +2/+0 de bônus |
| Dragonwing Glider | `{3}{R}{R}` | +2/+2, voar e haste | fora — 5 manas; o efeito é bom mas disputa o slot com o Wallbreaker, que dá **combate extra** pelo mesmo custo |
| Goldwarden's Helm | `{2}{W}` | +0/+1 | fora — 3 manas por um corpo 2/3; o Batterfist entrega corpo por 2 |
| Kemba's Banner | `{3}{W}` | +1/+1 **por criatura que você controla** | fora — o bônus é enorme num enxame, mas **concentra tudo numa criatura só**, que vira alvo único de remoção. 4 manas. Reserva nº 5 |
| Mirran Bardiche | `{4}{W}` | +2/+1 e **vigilância** | fora — 5 manas; a vigilância é a função que a §0.1 acabou de mostrar ser secundária |
| Hexgold Hoverwings | `{3}{W}` | voar · **criaturas equipadas** ganham +1/+0 | fora — o bônus coletivo só vale com muitos equipamentos equipados, e as fichas do Otharri entram já atacando (não dá para equipá-las antes do combate) |

**Já no deck:** `Hexgold Halberd` `{1}{R}` (For Mirrodin! + first strike e trample no seu turno).
Com as quatro entradas, o pacote vai a **5 equipamentos *For Mirrodin!*** = 5 Rebels desvirados
independentes do comandante.

### Dispensa da caixa para estas 4 compras (regra 7, por escrito)

| Compra | Equivalente na caixa | Por que a caixa não cobre |
|---|---|---|
| Barbed Batterfist · Bladehold War-Whip · Hexplate Wallbreaker · Glimmer Lens | **`Knight Watch`** e **`Goblin Gathering`** são as únicas fontes de ficha da caixa | Nenhuma das duas produz **Rebel** — as fichas são Knight e Goblin —, então **nenhuma paga o custo `Tap an untapped Rebel`**, que é exatamente a função que o usuário reportou faltando. Além disso as duas são **feitiço**: não deixam permanente, não contam para metalcraft e não sobrevivem ao wipe. As 4 compras cobrem as três coisas |
| Hexplate Wallbreaker (combate extra) | **nenhuma** — a caixa não tem nenhuma fonte de combate extra em RW | verificado na varredura completa das 116 |
| Glimmer Lens (saque) | **nenhuma** — a caixa não tem saque dentro da identidade RW (`Reconnaissance Mission` e `Bident of Thassa` são azuis; registrado em `decisions.md`) | o gap de saque é o maior do deck e a caixa não o toca |

## 0.5 ENTRADAS — ficha F1–F7

Condições assumidas (idênticas para quem entra e para quem sai, §3 do checklist): Otharri em campo
com 2–4 contadores, 4–10 fichas Rebel no board, `Intangible Virtue` e/ou `Warleader's Call` em
campo, 3+ artefatos (metalcraft do `Jor Kadeen` ligado).

**Consulta ao `decisions.md` (regra 5):** nenhuma das 7 entradas tem histórico de corte. O registro
contém apenas duas linhas — entrada do Otharri e saída da Tori da zona de comando. `Barbed Batterfist`
constava do **pool** do passe de 2026-09-19 (§4.5 nº 27) sem ter sido escolhida; `Knight Watch` constava
como reserva (§4.8 nº 42) e **continua reserva**. Nenhuma reposição de corte anterior.

### 1. Goblin Rabblemaster — `{2}{R}` · Creature — Goblin Warrior 2/2 · **caixa (R$ 0)**

- **F1** `Other Goblin creatures you control attack each combat if able.` / `At the beginning of combat on your turn, create a 1/1 red Goblin creature token with haste.` / `Whenever this creature attacks, it gets +1/+0 until end of turn for each other attacking Goblin.`
- **F2** corpo 2/2 no turno 3; a ficha de Goblin tem **haste** e ataca no turno em que nasce
- **F3** criatura (convoke do `Hour of Reckoning`, contagem do `Esgaroth Garrison`/`Luxknight Breacher`)
- **F4** recebe `Intangible Virtue` (as fichas), `Warleader's Call`, `Wedding Festivity`, `Jor Kadeen` (+3/+0), contadores de `Duty Beyond Death`/`Requisition Raid`
- **F5** dá +1/+0 a si por Goblin atacante
- **F6** turno 3 — e a partir do turno 4 **fabrica um corpo por turno sem board nenhum**, que é exatamente o cenário de mão morta relatado
- **F7** `Other Goblin creatures attack each combat if able` obriga as próprias fichas a atacar, esvaziando a defesa. Num deck cujo plano é atacar em massa, o atrito é pequeno; declarado
- **Sinergias (≥2):** (a) motor de corpo **independente do comandante** — resolve a dor "poucas criaturas" e reconstrói depois do wipe; (b) cada ficha que entra dispara `Valor in Akros`, `Warleader's Call` (1 de dano a cada oponente), `Belladonna Took`, `Staff of the Storyteller`, `Tocasia's Welcome` e `Molten Gatekeeper`; (c) as fichas entram **desviradas** e podem bloquear no crack-back

### 2. Breeches, Eager Pillager — `{2}{R}` · Leg. Creature — Goblin Pirate 3/3 · **caixa (R$ 0)**

- **F1** `First strike` / `Whenever a Pirate you control attacks, choose one that hasn't been chosen this turn — • Create a Treasure token. • Target creature can't block this turn. • Exile the top card of your library. You may play it this turn.` **Breeches é ele próprio um Pirate**, então o gatilho gira sempre que ele ataca
- **F2** corpo 3/3 com first strike por 3 manas — sobrevive ao bloqueio de 2/2 e mata
- **F3** criatura; a Treasure que ele cria é **artefato** (metalcraft) até ser sacrificada
- **F4** recebe todos os anthems e o metalcraft
- **F5** concede "não pode bloquear" a um bloqueador inimigo — furo de bloqueio dirigido, que é a rota R3
- **F6** turno 3; do turno 4 em diante entrega **ramp ou saque ou evasão, à escolha, todo turno**
- **F7** os três modos são exclusivos por turno (só um por ataque). Nenhuma competição de tap ou mana com outra peça
- **Sinergias (≥2):** (a) corpo agressivo cedo na curva que o deck não tinha; (b) **impulse draw ligado ao ataque** — reforça a categoria de saque sem gastar slot dela; (c) **Treasure** = ramp que ajuda a descer o Otharri no turno 4 (dor "lentidão para descer o comandante"); (d) "não pode bloquear" remove o chump-blocker que trava o enxame de 2/2

### 3. Zookeeper Mechan — `{1}{R}` · Artifact Creature — Robot 1/3 · **caixa (R$ 0)**

- **F1** `{T}: Add {R}.` / `{6}{R}: Target creature you control gets +4/+0 until end of turn. Activate only as a sorcery.`
- **F2** corpo 1/3 no **turno 2** — resistência 3 bloqueia 2/2 e sobrevive
- **F3** **artefato** (metalcraft do Jor Kadeen, contagem do Luxknight Breacher) **e** criatura (convoke do `Hour of Reckoning`, contagem do `Esgaroth Garrison`)
- **F4** recebe anthems e metalcraft — com `Jor Kadeen` em campo é um 4/3 que ataca
- **F5** **fixa mana `{R}`** e é ramp de turno 2; a ativação `{6}{R}` é dreno de mana tardio
- **F6** turno 2 — o ponto mais vazio da curva atual
- **F7** disputa o próprio `{T}` entre "gerar mana" e "atacar"; com `Intangible Virtue` ele não ganha vigilância (não é ficha). Atrito real e declarado
- **Sinergias (≥2):** (a) corpo barato **e** rocha de mana no mesmo slot — cobre a densidade de corpo **sem custar nada à contagem de ramp**, que é o que permite propor o corte do `Sword of the Animist`; (b) **artefato** — liga o metalcraft do `Jor Kadeen` já no turno 2 e conta para o `Luxknight Breacher`; (c) acelera o Otharri de T5 para **T4**, que é a reclamação literal de "lentidão para descer o comandante". *Atrito honesto: por ser artefato-**criatura**, ele morre nos board wipes de criatura — o metalcraft pós-wipe continua dependendo das rochas não-criatura.*

### 4. Barbed Batterfist — `{1}{R}` · Artifact — Equipment · **compra** (pendência de cotação)

- **F1** `For Mirrodin! (When this Equipment enters, create a 2/2 red Rebel creature token, then attach this to it.)` / `Equipped creature gets +1/−1.` / `Equip {1}`
- **F2** corpo 2/2 **desvirado** no turno 2 (3/1 depois de equipado)
- **F3** **artefato** (metalcraft, Luxknight) + a ficha é **criatura** e **Rebel**
- **F4** a ficha recebe `Intangible Virtue` (+1/+1 e vigilância → 4/2), `Warleader's Call`, `Jor Kadeen`, `Wedding Festivity`, e **double strike da `Neyali`** quando atacar
- **F5** equip `{1}` — o mais barato do deck; move o bônus para o Otharri ou para uma ficha de turno anterior
- **F6** turno 2
- **F7** **−1 de resistência** deixa a ficha em 3/1 e vulnerável a qualquer ping (`Molten Blast`, `Barrage of Boulders` do oponente). Mitigado por qualquer anthem
- **Sinergias (≥2):** (a) **Rebel desvirado que não depende de ataque anterior** — paga `{2}{R}{W}, Tap an untapped Rebel` no mesmo turno em que o Otharri morre (§0.1); (b) artefato que liga o metalcraft já no turno 2; (c) ficha que entra = gatilho de `Valor in Akros`/`Warleader's Call`/`Belladonna`/`Staff`/`Tocasia's Welcome`/`Molten Gatekeeper`

### 5. Bladehold War-Whip — `{1}{R}{W}` · Artifact — Equipment · **compra** (pendência de cotação)

- **F1** `For Mirrodin!` / `Equip abilities you activate of other Equipment cost {1} less to activate.` / `Equipped creature has double strike.` / `Equip {3}{R}{W}`
- **F2** corpo 2/2 Rebel **desvirado** com **double strike** = 4 de dano por 3 manas
- **F3** artefato + a ficha é criatura e Rebel
- **F4** cada anthem conta **duas vezes** nela: com `Intangible Virtue` + `Jor Kadeen` ela é um 6/2 com double strike = **12 de dano**
- **F5** **reduz o equip de todos os outros equipamentos em `{1}`** — `Sword of the Animist` `{2}`→`{1}`, `Hexgold Halberd`, `Ancestral Blade`, `Barbed Batterfist` `{1}`→`{0}`. É o único redutor de custo do deck
- **F6** turno 3
- **F7** o **equip próprio** é `{3}{R}{W}` (proibitivo — na prática nunca se realoca). **Redundância parcial com a `Neyali`**, que já dá double strike a todas as fichas atacantes; declarada. Cobre o que ela não cobre: criaturas **não-ficha** (Otharri, Jor Kadeen, Hero of Bladehold)
- **Sinergias (≥2):** (a) Rebel desvirado; (b) double strike que **multiplica todo anthem do deck**; (c) redutor de custo de equip — o deck passa a ter 5 equipamentos; (d) artefato para metalcraft

### 6. Glimmer Lens — `{1}{W}` · Artifact — Equipment · **compra** (pendência de cotação)

- **F1** `For Mirrodin!` / `Whenever equipped creature and at least one other creature attack, draw a card.` / `Equip {1}{W}`
- **F2** corpo 2/2 Rebel **desvirado** no turno 2
- **F3** artefato + criatura + Rebel
- **F4** recebe todos os anthems e o double strike da Neyali
- **F5** —
- **F6** turno 2; **do turno 3 em diante compra uma carta por ataque**, porque o Otharri sozinho já satisfaz "pelo menos outra criatura"
- **F7** o saque exige que a **ficha equipada** ataque — ou seja, ela não pode ficar de bloqueadora; e se a ficha morrer, é preciso pagar `{1}{W}` para reequipar
- **Sinergias (≥2):** (a) **saque repetível e incondicional** num deck cujo gap nº 1 é card advantage — sem gastar um slot da categoria draw; (b) Rebel desvirado para a recursão; (c) artefato para metalcraft; (d) corpo no turno 2

### 7. Hexplate Wallbreaker — `{3}{R}{R}` · Artifact — Equipment · **compra** (pendência de cotação, maior risco)

- **F1** `For Mirrodin!` / `Equipped creature gets +2/+2.` / `Whenever equipped creature attacks, if it's the first combat phase of the turn, untap each attacking creature. After this phase, there is an additional combat phase.` / `Equip {3}{R}`
- **F2** corpo **4/4** Rebel desvirado (2/2 + 2/+2), já equipado
- **F3** artefato + criatura + Rebel
- **F4** recebe anthems; com `Jor Kadeen` é um 7/4
- **F5** **desvira todos os atacantes** — inclusive as fichas do Otharri que entraram viradas — e concede **combate adicional**, repetível **todo turno**
- **F6** turno 5
- **F7** três atritos declarados: (i) CMC 5 **disputa o turno do Otharri** — é o pior ponto da curva para este deck; (ii) ruling 2023-02-04: *"There's no main phase between your combat phases"* — não dá para equipar nem conjurar nada entre os combates; (iii) redundância parcial com `Combat Celebrant` (que faz o mesmo **uma vez**, por exert) — mas o Wallbreaker é **repetível** e traz o próprio corpo
- **Sinergias (≥2):** (a) **segundo gatilho de ataque do Otharri no mesmo turno** = +1 contador de experiência **e** uma leva inteira de fichas a mais, todo turno — é o multiplicador do motor; (b) desvira as fichas que entraram viradas → elas atacam **de novo** no combate extra **e** ficam disponíveis como Rebel desvirado; (c) Rebel desvirado próprio; (d) artefato para metalcraft

### Reserva (se abrir um 8º slot) — Patrolling Peacemaker `{2}{W}` · Artifact Creature 2/2 · **caixa (R$ 0)**

- **F1** `This creature enters with two +1/+1 counters on it.` / `Whenever an opponent commits a crime, proliferate.`
- **F2** corpo 2/2 no turno 3 · **F3** artefato **e** criatura · **F4** recebe anthems e contadores
- **F5** **proliferate** — e o contador mais valioso da mesa é o **de experiência do Otharri**: +1 contador = **+1 ficha por ataque para o resto do jogo**
- **F6** turno 3 · **F7** o gatilho depende de oponente cometer crime (mirar qualquer jogador, permanente ou carta de cemitério de um oponente) — num pod de 4 isso acontece várias vezes por rodada, mas **não é garantido**
- **Sinergias (≥2):** (a) é a **única fonte de proliferate de custo zero** — as três do pool (§4.6) eram todas compra, e a `Karn's Bastion` foi cortada por preço (R$ 9,85); (b) corpo + artefato para metalcraft; (c) também engorda `Sheriff of Safe Passage`, `Luxknight Breacher` e `Basri Ket`

## 0.6 SAÍDAS — ficha F1–F7 e protocolo de corte

Mesmas condições da §0.5. Nenhuma das cartas abaixo tem histórico em `decisions.md`.

### Corte limpo 1 — `Pride of Conquerors` `{1}{W}` · Instant (já é seu, R$ 0)

- **F1** `Ascend` / `Creatures you control get +1/+1 until end of turn. If you have the city's blessing, those creatures get +2/+2 instead.`
- **F2** nenhum corpo · **F3** instantâneo, não alimenta contagem nenhuma · **F4** não recebe nada · **F5** pump de massa **de um turno** · **F6** turno 2, mas **morta até haver board** · **F7** disputa o mana do turno do alfa com `Duty Beyond Death`/`Unbreakable Formation`, que são a resposta ao wipe — e o deck só tem um turno de mana por vez
- **Cobertura:** pump em massa fica com `Valor in Akros` (dispara **por ficha que entra**, de graça), `Jor Kadeen` (+3/+0 estático), `Warleader's Call`, `Intangible Virtue`, `Wedding Festivity` e o modo 3 da `Requisition Raid` (contadores **permanentes**). **Nada fica descoberto** — é a 6ª cópia funcional do mesmo efeito.

### Corte limpo 2 — `Radiant Destiny` `{2}{W}` · Enchantment (compra, **devolve R$ 1,58**)

- **F1** `Ascend` / `As this enters, choose a creature type. Creatures you control of the chosen type get +1/+1. As long as you have the city's blessing, they also have vigilance.`
- **F2** nenhum corpo · **F3** encantamento · **F4** — · **F5** +1/+1 e vigilância **só ao tipo escolhido** · **F6** turno 3, morta sem board · **F7** obriga a escolher **um** tipo; nomeando Rebel ela **ignora** as fichas de Goblin do Rabblemaster, o Soldier do `Ancestral Blade`, o Spirit do `Staff of the Storyteller`, o Human do `Wedding Announcement` e os Soldiers do `Assemble the Legion`
- **Cobertura:** `Intangible Virtue` faz **o mesmo por 1 mana a menos**, sem escolher tipo (`Creature tokens you control get +1/+1 and have vigilance`) e **com a vigilância incondicional** — a `Radiant Destiny` só a concede com *city's blessing*. Depois das entradas da §0.5 a base de fichas fica **heterogênea** (Rebel, Goblin, Soldier, Spirit, Human), e a `Radiant Destiny` passa a cobrir a minoria delas. **Corte estritamente dominado. Nada descoberto.**

### Corte limpo 3 — `The Circle of Loyalty` `{4}{W}{W}` · Leg. Artifact (já é seu, R$ 0)

- **F1** `Affinity for Knights` / `Creatures you control get +1/+1.` / `Whenever you cast a legendary spell, create a 2/2 white Knight creature token with vigilance.` / `{3}{W}, {T}: Create a 2/2 white Knight creature token with vigilance.`
- **F2** não é corpo na resolução; **é fábrica de corpo ativada** · **F3** artefato (metalcraft) · **F4** — · **F5** anthem +1/+1 a todas · **F6** **turno 6** na prática: o deck tem 3 Knights (`Tori D'Avenant`, `Éomer, Marshal of Rohan`, `Hero of Bladehold`) e raramente há 2 em campo antes do turno 5, então a affinity quase nunca desce o custo abaixo de 5 · **F7** a `{T}` compete com nada, mas os `{3}{W}` por ficha competem com descer o Otharri, protegê-lo e recastá-lo
- **Cobertura, função a função:**
  - anthem +1/+1 → `Warleader's Call`, `Intangible Virtue` (fichas), `Wedding Festivity`, `Jor Kadeen` (+3/+0). **Coberto com folga**;
  - artefato para metalcraft → o deck sai de 18 para **21 permanentes-artefato** (§0.8). **Coberto**;
  - ficha por lendária conjurada → o deck tem 7 lendárias; efeito perdido, **custo declarado e aceito** (é o menor dos quatro);
  - **fábrica de corpo repetível imune a wipe de criaturas** → `Assemble the Legion` (encantamento, cria fichas **com haste** em quantidade crescente **todo upkeep**, sem gastar mana) e `Wedding Announcement` (1/1 por end step enquanto você não ataca com 2+). **Coberto por duas peças que fazem melhor e mais barato.**
- **Por que é este o slot:** é o **anthem mais caro do deck** num deck que o usuário descreveu como lento, e o 5º efeito de +1/+1 numa lista que já tem 9 fontes de pump. Alternativa se o orquestrador preferir preservar o dreno de mana: trocar este corte por `Dawnstrike Vanguard` (CMC 6) — mas o Vanguard **é corpo** e a §0.2 está justamente comprando corpo, então a troca piora a conta.

### Corte limpo 4 — `Barrage of Boulders` `{2}{R}` · Sorcery (compra, **devolve R$ 0,05**)

- **F1** `deals 1 damage to each creature you don't control.` / `Ferocious — If you control a creature with power 4 or greater, creatures can't block this turn.`
- **F2** nenhum corpo · **F3** feitiço · **F4** — · **F5** Falter condicional a *ferocious* · **F6** turno 3, **morta sem board dos dois lados** · **F7** o modo Falter exige poder 4+; as fichas são 2/2 e só chegam lá **com** o `Jor Kadeen` ou anthems — ou seja, ele só liga quando você já estava ganhando
- **Cobertura:** furo de bloqueio (rota R3) fica com `Goblin War Drums` (menace **estático** a todas — contra um enxame de 10 exige 20 bloqueadores), `Crash Through` (trample a todas, e **compra um card**), `Hexgold Halberd` (trample) e o modo "target creature can't block" do **`Breeches`**, que entra nesta rodada e é **repetível todo ataque**. O 1 de dano a cada criatura do oponente fica **descoberto**: é varredura parcial de X/1, parcialmente coberta por `Blast Zone`, `Celebrate the Mountain-king` e `Molten Blast`. **Custo declarado e aceito** — o `report.md` já registra em §11.2 que o deck aceita não ter resposta em massa a enxame alheio.

### Corte limpo 5 — `Late to Dinner` `{3}{W}` · Sorcery (já é seu, R$ 0)

- **F1** `Return target creature card from your graveyard to the battlefield. Create a Food token.`
- **F2** não é corpo na mão de abertura — **depende de já haver criatura no cemitério** · **F3** feitiço; a Food é artefato (metalcraft) · **F4** — · **F5** — · **F6** turno 4, e **morta em toda mão de abertura** · **F7** compete com o `{2}{R}{W}` da recursão do próprio Otharri pelo mesmo mana e pelo mesmo turno
- **Cobertura:** a recursão do comandante **desvirado** era o argumento único a favor. Depois desta rodada o deck tem **5 Rebels desvirados** independentes do comandante (`Hexgold Halberd` + as 4 entradas *For Mirrodin!*), então a habilidade própria do Otharri — que **não gasta carta** — passa a estar disponível de forma confiável, inclusive no turno em que ele morre (§0.1). Fica descoberto: trazer de volta `Hero of Bladehold`/`Neyali`/`Jor Kadeen` depois de um wipe. **Custo declarado**; a reconstrução pós-wipe fica com `Assemble the Legion`, `Wedding Announcement`, `Goblin Rabblemaster` e os 5 equipamentos, que **sobrevivem ao wipe de criaturas** (ruling: *"If the Rebel is destroyed, the Equipment stays on the battlefield"*).

### Corte condicionado 1 — `Sword of the Animist` `{2}` · Leg. Artifact — Equipment · **R$ 27,00** → **ramp-specialist**

**Não corto: a função principal é ramp, fora da minha especialidade.** Devolvo ao orquestrador com
a ficha pronta, porque é **o maior alavanque de orçamento do deck** (16 % do teto num slot só).

- **F1** `Equipped creature gets +1/+1.` / `Whenever equipped creature attacks, you may search your library for a basic land card, put it onto the battlefield tapped, then shuffle.` / `Equip {2}`
- **F2** nenhum corpo · **F3** artefato (metalcraft) · **F4** — · **F5** +1/+1 e **ramp por terreno básico** ao atacar · **F6** turno 2, mas o ciclo completo (equipar `{2}` + atacar) só fecha no turno 4 · **F7** **as fichas do Otharri entram já atacando** — equipar é velocidade de feitiço, então elas **nunca** podem receber a espada antes do combate. Só criaturas de turnos anteriores a carregam. Além disso disputa o mana de equip com `Ancestral Blade`, `Hexgold Halberd` e as 4 entradas *For Mirrodin!*
- **Se sair, quem cobre:** ramp padrão cai de **11 para 10** (`Sol Ring`, `Arcane Signet`, `Boros Signet`, `Talisman of Conviction`, `Mind Stone`, `Everflowing Chalice`, `Sphere of the Suns`, `Boros Locket`, `Commander's Sphere`, `Loyal Warhound`) — **ainda dentro da meta de 10–11** — e **volta a 11** com o `Zookeeper Mechan` (`{T}: Add {R}`), que entra nesta rodada; somam-se as Treasures do `Breeches`. Metalcraft: 18 → **21** artefatos mesmo sem ela. Fixação de cor: 6 rochas coloridas + 4 terrenos não-básicos. **Nada fica descoberto** — mas a decisão é do orquestrador, com o ramp-specialist.
- **O que muda se ficar:** a folga cai de R$ 62,40 para R$ 35,40, o que aperta as 4 compras pendentes de cotação (§0.8).

### Corte condicionado 2 — `Feat of Resistance` `{1}{W}` · Instant (já é seu, R$ 0) → **interaction-specialist**

- **F1** `Put a +1/+1 counter on target creature you control. It gains protection from the color of your choice until end of turn.`
- **F2** nenhum corpo · **F3** instantâneo · **F4** põe **contador permanente** no alvo · **F5** proteção de cor a **um** alvo · **F6** turno 2 · **F7** é a 2ª de duas proteções de alvo único (`Gods Willing` faz o mesmo por **1 mana**, com scry 1); somadas a `Duty Beyond Death` e `Unbreakable Formation`, o deck carrega **4 instantâneos reativos** — o excesso exato de que o usuário reclamou
- **Cobertura:** proteção de alvo único → `Gods Willing` (1 mana, protection from a color + scry). Proteção em massa → `Duty Beyond Death` e `Unbreakable Formation`. O **contador permanente** fica descoberto por esta via, mas é replicado por `Requisition Raid` (modo 3), `Basri's Solidarity`, `Basri Ket` (+1) e `Duty Beyond Death`.
- **Por que condicionado:** proteção é categoria do `interaction-specialist` e a contagem de interação cai de **14 para 13**. **Não corto sozinho.**

### NÃO CORTO — `Basri's Solidarity` `{1}{W}` · Sorcery

O `report.md` §11.2 registra que **o usuário pediu explicitamente para manter a `Basri's Solidarity`
junto com a `Requisition Raid`**. Pela ficha isolada ela seria o candidato mais óbvio a corte (o modo
3 da Raid é superconjunto funcional e a Raid ainda mata artefato/encantamento), mas o item 5 da §5 do
checklist é explícito: **nenhum corte atinge a peça que o próprio usuário disse querer usar**.
Registro aqui como **pendência do orquestrador**, não como recomendação minha. Se ele liberar o slot,
a fila é `Patrolling Peacemaker` (§0.5) e a conta da §0.2 vai para **K = 27 → 9,90 %**.

## 0.6.1 Ficha de funções — tabela compacta das 14 cartas que se movem

| Carta | Mov. | Categorias | Corpo tapável (F2) | Tipo alimenta (F3) | Recebe (F4) | Facilita (F5) | Entra no turno (F6) | Atritos (F7) |
|---|---|---|---|---|---|---|---|---|
| Goblin Rabblemaster | **entra** | tema, wincon | 2/2 + 1/1 haste por turno; convoke do `Hour of Reckoning` | criatura (Esgaroth, Luxknight) | anthems, metalcraft, contadores | +1/+0 a si por Goblin atacante | 3 | força os próprios Goblins a atacar |
| Breeches, Eager Pillager | **entra** | tema, draw, ramp | 3/3 first strike | criatura; Treasure = artefato | anthems, metalcraft | "não pode bloquear" a um alvo | 3 | um modo por ataque |
| Zookeeper Mechan | **entra** | tema, ramp | 1/3 que **usa o próprio {T}** para mana | **artefato** + criatura | anthems, metalcraft | fixa `{R}`; dreno `{6}{R}` | 2 | tap disputado entre mana e ataque; morre a wipe de criatura |
| Barbed Batterfist | **entra** | tema | ficha 2/2 Rebel **desvirada** (3/1 equipada) | **artefato** + criatura + Rebel | anthems, double strike da Neyali | equip `{1}` | 2 | −1 de resistência deixa a ficha em 3/1 |
| Bladehold War-Whip | **entra** | tema, wincon | ficha 2/2 Rebel desvirada com **double strike** | **artefato** + criatura + Rebel | anthems contam em dobro | **reduz equip de outros equipamentos em `{1}`** | 3 | equip próprio `{3}{R}{W}`; redundante com a Neyali nas fichas |
| Glimmer Lens | **entra** | tema, **draw** | ficha 2/2 Rebel desvirada | **artefato** + criatura + Rebel | anthems, Neyali | — | 2 | saque exige que a ficha equipada **ataque** |
| Hexplate Wallbreaker | **entra** | tema, wincon | ficha **4/4** Rebel desvirada | **artefato** + criatura + Rebel | anthems, metalcraft | **desvira todos os atacantes + combate adicional** | 5 | CMC 5 disputa o turno do Otharri; sem fase principal entre combates |
| *(reserva)* Patrolling Peacemaker | reserva | tema | 2/2 | **artefato** + criatura | anthems, contadores | **proliferate** → contador de experiência | 3 | gatilho depende de crime do oponente |
| Pride of Conquerors | **sai** | tema | — | — | — | pump de massa de 1 turno | 2 (morta sem board) | disputa o mana do turno do alfa com a proteção |
| Radiant Destiny | **sai** | tema | — | encantamento | — | +1/+1 e vigilância **a um tipo só** | 3 (morta sem board) | escolhe 1 tipo; ignora Goblin/Soldier/Spirit/Human |
| The Circle of Loyalty | **sai** | tema, wincon | fábrica ativada `{3}{W},{T}` | artefato | — | anthem +1/+1; affinity for Knights | ~6 real | 6 manas num deck que o usuário chamou de lento |
| Barrage of Boulders | **sai** | tema, remoção parcial | — | feitiço | — | Falter condicional a *ferocious* | 3 (morta sem board) | exige poder 4+; só liga quando já está ganhando |
| Late to Dinner | **sai** | recursão | — (depende de cemitério) | feitiço; Food = artefato | — | — | 4 (morta na abertura) | compete com a recursão própria do Otharri |
| Feat of Resistance | **sai** (condicionado) | proteção | — | instantâneo | põe contador permanente | proteção de cor a **um** alvo | 2 | 2ª de duas proteções de alvo único |
| Sword of the Animist | **sai** (condicionado) | ramp | — | artefato | — | +1/+1 e busca terreno básico ao atacar | 2 (ciclo fecha no 4) | **as fichas entram já atacando e nunca podem ser equipadas** |

## 0.7 Resumo entra/sai

| Sai | R$ devolvido | Entra | R$ | Motivo em uma linha |
|---|---|---|---|---|
| `Pride of Conquerors` | — (é seu) | `Goblin Rabblemaster` | **caixa** | troca o 6º pump de um turno por um motor que fabrica corpo **sem board** |
| `Radiant Destiny` | **1,58** | `Breeches, Eager Pillager` | **caixa** | anthem estreito (só Rebel) → corpo 3/3 que dá ramp/saque/evasão todo ataque |
| `The Circle of Loyalty` | — (é seu) | `Zookeeper Mechan` | **caixa** | anthem de CMC 6 → corpo de CMC 2 que também é rocha de mana e artefato |
| `Barrage of Boulders` | **0,05** | `Barbed Batterfist` | *pendência* | Falter condicional → Rebel **desvirado** no turno 2 + artefato |
| `Late to Dinner` | — (é seu) | `Bladehold War-Whip` | *pendência* | recursão que exige cemitério → Rebel desvirado com **double strike** + redutor de equip |
| `Feat of Resistance` **(condicionado → interação)** | — (é seu) | `Glimmer Lens` | *pendência* | 2ª proteção de alvo único → Rebel desvirado + **saque repetível por ataque** |
| `Sword of the Animist` **(condicionado → ramp)** | **27,00** | `Hexplate Wallbreaker` | *pendência* | equipamento que as fichas nunca podem usar → Rebel 4/4 + **combate extra repetível** |
| *(pendência do usuário)* `Basri's Solidarity` | — | *(reserva)* `Patrolling Peacemaker` | caixa | libera o 8º slot e leva a conta a 9,90 % |

**7 saem · 7 entram · 99 cartas mantidas.** Nenhum slot fica aberto para outra fase.
Se o orquestrador recusar os dois cortes condicionados, a fila de entrada por prioridade é:
**1.** `Goblin Rabblemaster` · **2.** `Barbed Batterfist` · **3.** `Bladehold War-Whip` ·
**4.** `Glimmer Lens` · **5.** `Breeches` — as cinco primeiras cabem nos 5 cortes limpos e levam
a conta a **K = 24 → 13,33 %**, ainda acima do alvo de 12 %. **Cortar ao menos um dos dois
condicionados é o que fecha a meta.**

## 0.8 Efeito nas contagens do deck

| Métrica | v1 | Proposta | Meta do pipeline |
|---|---|---|---|
| **Fontes de corpo** (def. §0.2) | 19 | **26** | novo alvo: P(sem corpo) < 12 % ✔ **10,94 %** |
| Cartas do tipo criatura | 15 | **18** | — |
| Instantâneos + feitiços + encantamentos | 30 | **25** | razão corpo:mágica de 0,63 : 1 → **1,04 : 1** |
| Permanentes-artefato (metalcraft ≥ 3) | 18 | **21** | ✔ sobe |
| Rebels desvirados independentes do comandante | **1** | **5** | ✔ resolve a dor relatada |
| Combate extra | 1 (`Combat Celebrant`) | **2** (+ `Hexplate Wallbreaker`, repetível) | rota R4 ✔ |
| Fontes de pump/anthem | 9 | **6** | ✔ corta a redundância |
| Saque | 13 | **15** (+`Glimmer Lens`, +`Breeches`) | 12–13 ✔ |
| Ramp padrão | 11 | **11** (−`Sword`, +`Zookeeper Mechan`) | 10–11 ✔ |
| Ramp explosivo | 2 | 2 | 2–3 ✔ |
| Interação | 14 | **13** (se o corte condicionado 2 passar) | ~10 ✔ |
| Board wipes | 2 | 2 | 2–4 ✔ |
| Terrenos | 36 | 36 | 38 ✘ (pendência da Fase 6, inalterada) |

**Composição final proposta dos 99:** 18 criaturas · 19 artefatos não-criatura · 10 encantamentos ·
8 instantâneos · 7 feitiços · 1 planeswalker · 36 terrenos = **99**.

## 0.9 Orçamento — delta e pendências

Régua: **menor valor da LigaMagic** (regra 2). Base: R$ 166,23 nas 99 (cotações de 19–20/09/2026).

| Movimento | R$ |
|---|---|
| Base v1 | 166,23 |
| − `Radiant Destiny` | −1,58 |
| − `Barrage of Boulders` | −0,05 |
| − `Sword of the Animist` *(condicionado)* | −27,00 |
| + `Goblin Rabblemaster`, `Breeches`, `Zookeeper Mechan` (caixa) | 0,00 |
| **Subtotal antes das 4 compras** | **137,60** |
| **Folga contra o teto de R$ 200** | **62,40** |

**Se o `Sword of the Animist` ficar:** subtotal 164,60 · folga **35,40**.

### Pendências de cotação — 4 cartas, nenhuma com preço registrado

`bin/mtgdb prices` não tem nenhuma das quatro. **Não invento número.** A estimativa em US$ abaixo é
da Scryfall e serve **só para priorizar a captura manual na LigaMagic** — o guia registra erros
medidos de 6,5× para mais e 14× para menos, nos dois sentidos.

| Carta | Estimativa Scryfall (US$) | Risco | Plano B se estourar |
|---|---|---|---|
| `Barbed Batterfist` | 0,19 | baixo (comum de ONE; a `Hexgold Halberd`, do mesmo ciclo, saiu a **R$ 0,10**) | — |
| `Bladehold War-Whip` | 0,27 | baixo (mesmo ciclo) | — |
| `Glimmer Lens` | 4,46 | **médio** | `Goldwarden's Helm` (mesmo custo de corpo, sem o saque) ou `Patrolling Peacemaker` da caixa |
| `Hexplate Wallbreaker` | 7,72 | **alto** — é a única compra que pode estourar sozinha | `Dragonwing Glider` / `Vulshok Splitter` (mesmo corpo, sem combate extra) ou `Kemba's Banner`; todos do mesmo ciclo de commons/uncommons |

Com a folga de **R$ 62,40** (cenário com o `Sword` cortado), as quatro cabem mesmo no pior caso
plausível. Com o `Sword` mantido (**R$ 35,40**), só cabem se o `Hexplate Wallbreaker` sair por
menos de ~R$ 25 — **é este o ponto de decisão que devolvo ao orquestrador**.

---

# Passe de 2026-09-19 — re-pass com Otharri (preservado; continua válido salvo onde a §0 corrige)


## 1. Comandante — análise linha a linha

`Otharri, Suns' Glory` · `{3}{R}{W}` · 3/3 · Legendary Creature — Phoenix · CMC 5

```
Flying, lifelink, haste
Whenever Otharri attacks, you get an experience counter. Then create a 2/2 red Rebel creature
token that's tapped and attacking for each experience counter you have.
{2}{R}{W}, Tap an untapped Rebel you control: Return this card from your graveyard to the
battlefield tapped.
```

| Linha/habilidade | Gatilho/termo | O que **exige** do deck | O que **oferece** |
|---|---|---|---|
| `Flying` | evasão própria | nada | o gatilho de ataque é **quase incondicional** — poucos decks bloqueiam um 3/3 voador |
| `Lifelink` | — | nada | com anthems, o corpo dele vira dreno relevante; segura o *crack-back* do go-wide |
| `Haste` | — | nada | **ataca no turno em que entra** → o motor começa no mesmo turno. É o que torna o CMC 5 tolerável e o que faz o ramp valer (T3 Sol Ring → T4 Otharri → ataque no T4) |
| `Whenever Otharri attacks, you get an experience counter` | **attack-trigger**, contador **no jogador** | Otharri **atacando** todo turno (não precisa conectar) | recurso que **não está no campo de batalha**: board wipe, exílio, bounce e remoção **não tiram os contadores** |
| `Then create ... for each experience counter you have` | **token-maker escalonado** | nada além do ataque | 1, 2, 3, 4… fichas 2/2 **vermelhas** por ataque |
| fichas entram `tapped and attacking` | **enters-attacking** | nada | dano **no mesmo turno**; **mas** não foram *declaradas atacantes* — ver §1.3 |
| fichas são `Rebel` **vermelhas** | tipo/cor | nada | cor **vermelha** e tipo **Rebel** são os dois eixos que mais mudam a avaliação do deck atual (§3) |
| `{2}{R}{W}, Tap an untapped Rebel you control: Return this card from your graveyard to the battlefield tapped` | **recursão do cemitério** | **um Rebel desvirado** | recompra o comandante por **4 manas fixos, sem imposto de comandante**, quantas vezes forem necessárias |

### 1.1 Matemática de acumulação

Contadores de experiência crescem **1 por ataque de Otharri** (não por criatura, não por dano).
As fichas criadas em cada ataque são **cumulativas** — as antigas ficam.

| Ataque nº | Contadores | Fichas novas | Fichas no total | Poder atacando (2/2, sem anthem) | Poder com +1/+1 estático |
|---|---|---|---|---|---|
| 1 | 1 | 1 | 1 | 3 (Otharri) + 2 = **5** | 4 + 3 = **7** |
| 2 | 2 | 2 | 3 | 3 + 6 = **9** | 4 + 9 = **13** |
| 3 | 3 | 3 | 6 | 3 + 12 = **15** | 4 + 18 = **22** |
| 4 | 4 | 4 | 10 | 3 + 20 = **23** | 4 + 30 = **34** |
| 5 | 5 | 5 | 15 | 3 + 30 = **33** | 4 + 45 = **49** |

Total de fichas após N ataques = **N(N+1)/2**. Dano acumulado a **um** jogador em 4 ataques,
sem nenhum anthem: **52**. Com **um único** anthem estático de +1/+1: **76**.

**Consequência de construção nº 1:** cada ponto de anthem **estático** vale N pontos de dano,
onde N é o tamanho do enxame — e o enxame cresce quadraticamente. **Um +1/+1 estático no
ataque nº 4 vale +10 de dano; no ataque nº 5, +15.** Nenhum efeito de alvo único chega perto.

**Consequência de construção nº 2:** o que mais escala **não** é anthem — é **combate extra**.
Cada combate adicional é **outro gatilho de ataque**: +1 contador *e* uma leva nova de fichas
dimensionada pelo contador já incrementado. Dois combates no turno 3 do Otharri produzem
3 + 4 = 7 fichas em vez de 3.

### 1.2 A recursão — o deck precisa de uma fonte de Rebel desvirado?

O custo é `{2}{R}{W}` **+ virar um Rebel desvirado**. Ruling oficial (2023-02-04): virar o Rebel
é **parte do custo**, então o oponente não consegue responder removendo o Rebel.

As fichas do próprio Otharri entram **viradas**. Elas **não** pagam o custo no turno em que
nascem — mas **desviram na sua próxima etapa de desvirar**. A partir daí são Rebels desvirados.

Cenário a cenário:

| Situação | Há Rebel desvirado? | Veredito |
|---|---|---|
| Otharri atacou ≥1 vez e depois morreu (remoção pontual) | **Sim** — as fichas do turno anterior já desviraram | recursão funciona, **4 manas fixos**, sem imposto |
| Otharri removido **antes** do primeiro ataque (counter, remoção no ETB, exílio) | **Não** — nenhuma ficha foi criada | zona de comando, com imposto |
| **Board wipe**: Otharri e todas as fichas morrem juntos | **Não** — as fichas morreram com ele | recursão **bricada**; um Rebel não-ficha teria morrido no mesmo wipe |
| Wipe + um Rebel **não-ficha** que sobreviva (indestrutível/proteção) | Sim | recursão funciona |

**Veredito:** **a zona de comando basta como plano B, e as próprias fichas bastam como plano A
na esmagadora maioria dos casos.** Um *pacote* de Rebels (Ramosian Sergeant, Lin Sivvi,
Ramosian Lieutenant…) **não** se justifica: são corpos 1/1–2/2 sem outra função, e no caso em
que a recursão falha de verdade (o board wipe) eles morrem junto e não resolvem nada.

**Mas há duas cartas que entregam um Rebel desvirado como *efeito colateral* de algo que o deck
já quer** — e essas sim entram no pool, porque somam 3 eixos cada:

- **Barbed Batterfist** / **Hexgold Halberd** (`{1}{R}`, *For Mirrodin!*) — ao entrar, **criam uma
  ficha 2/2 vermelha Rebel** (desvirada) e se prendem a ela. São **artefato** (metalcraft do Jor
  Kadeen, contagem do Luxknight Breacher), **corpo vermelho 2/2** e **Rebel de plantão** para a
  recursão. Custam 2 manas.
- **Neyali, Suns' Vanguard** (`{2}{R}{W}`) — é **Human Rebel**, e o resto do texto dela é a melhor
  sinergia disponível para o deck (§4).

E um terceiro eixo que resolve a mesma coisa sem gastar carta: **vigilância nas fichas**
(`Radiant Destiny` nomeando Rebel, `Intangible Virtue`). Ficha com vigilância **não vira ao
atacar** → há sempre um Rebel desvirado, e ele ainda bloqueia.

**Nota tática:** a recursão devolve Otharri **virado**. Ele não ataca no turno em que volta.
`Late to Dinner` (já no deck) devolve **desvirado** — e com haste ele ataca imediatamente.
A recursão própria é *resiliência barata e infinita*; Late to Dinner é *velocidade*.
As duas não competem, se complementam.

### 1.3 A regra que reordena o deck inteiro — fichas que entram atacando

Ruling de Basri Ket (2020-06-23), idêntico em mecânica: *"Although the Soldiers are attacking
creatures, they were never declared as attacking creatures. This means that abilities that
trigger whenever a creature attacks won't trigger when the Soldiers enter the battlefield
attacking."* Ruling de Adriana (2016-08-23) diz o mesmo para melee.

Disso saem **duas** conclusões opostas, e o passe anterior teria errado as duas:

**(a) O que NÃO alcança as fichas — e não tem conserto:**
- gatilhos que exigem que **a própria ficha** tenha sido declarada atacante: **melee**
  (`Adriana`), *battle cry* concedido, `Hellrider`-likes, **mentor** por comparação de poder
  (`Parhelion Patrol`: mentor exige poder **menor**, e as fichas são poder 2 como ela);
- qualquer efeito que leia **tipo de criatura errado**: as fichas são **Rebel vermelho**, não
  Human, não Knight, não Soldier. `Inspiring Veteran`, `Sanctuary Lockdown`, `Vigilante Justice`,
  `True-Faith Censer` **nunca** veem o enxame.

**(b) O que ALCANÇA as fichas — por ordenação de gatilhos:**
Todo gatilho `whenever ~ attacks` que você controla vai para a pilha **junto** com o de Otharri,
no passo de declarar atacantes, e **você escolhe a ordem**. Colocando o gatilho de Otharri por
**último na pilha**, ele **resolve primeiro**: as fichas passam a existir **antes** dos demais
gatilhos resolverem, e efeitos do tipo *"outras criaturas atacantes ganham X"* as incluem,
porque esses efeitos enxergam o board **no momento em que resolvem**.

Isso salva `Honored Crop-Captain`, `Syr Alin, the Lion's Claw`, `Hero of Bladehold` (battle cry
dela própria) — **e a Tori D'Avenant** (§6). Não salva melee nem mentor.

**(c) O que alcança sem nenhuma condição:** **estáticas** (`The Circle of Loyalty`, `Jor Kadeen`,
`Radiant Destiny`, `Warleader's Call`) e **gatilhos de "criatura entra"** (`Valor in Akros`,
`Belladonna Took`, `Cathars' Crusade`, `Impact Tremors`) — e estes últimos **disparam N vezes**,
uma por ficha. Um ataque com 4 contadores dispara `Valor in Akros` **4 vezes** = +4/+4 no time
inteiro; dispara `Belladonna Took` 4 vezes = vida + **carta** + **contador em cada criatura**.

### 1.4 Contadores de experiência e resiliência a wipe

Os contadores ficam **no jogador**, não em permanente. Wrath of God, Farewell, exílio, bounce e
roubo **não os tocam**. Depois de um wipe com 5 contadores acumulados, o primeiro ataque do
Otharri recolocado devolve **5 fichas (10 de poder) de uma vez**. É a única resiliência
estrutural a wipe que RW tem em qualquer faixa de preço — e ela **já está paga**, porque o
comandante não conta na régua de R$ 200.

**Corolário para a dor 4:** a resposta correta não é só proteção — é **proteger o Otharri**
(para que o ataque aconteça) e **deixar o wipe passar** (porque os contadores ficam). O que
precisa de proteção é **uma** criatura, não oito. Isso **inverte** a leitura do passe anterior,
que tratava as 4 proteções de alvo único como excesso.

---

## 2. Termos de busca do novo eixo

| # | Termo | Por que | Busca usada |
|---|---|---|---|
| 1 | **anthem estático** (`creatures you control get +X/+X`) | multiplica por N o enxame, permanentemente | `search '"creatures you control get"' -id RW -cmc-max 4` |
| 2 | **creature-enters / creaturefall** | N fichas = **N gatilhos** por ataque | oracle carta a carta (`Valor in Akros`, `Warleader's Call`, `Impact Tremors`, `Cathars' Crusade`, `Belladonna Took`) |
| 3 | **extra combat** | cada combate = +1 contador **e** uma leva maior de fichas | `tag extra-combat-phase -id RW` |
| 4 | **token-maker com gatilho de ataque** | empilha corpos com o motor principal | `search '"Whenever you attack"' -id RW -cmc-max 5` |
| 5 | **Rebel** (tipo) / **vigilância em ficha** | destrava a recursão `{2}{R}{W}` | `search "Rebel" -id RW` |
| 6 | **proliferate** | **contador de experiência extra** = +1 ficha por ataque para sempre | `search "proliferate" -id RW -cmc-max 5` |
| 7 | **trample / menace em massa** | fichas 2/2 terrestres travam em chump-block | `Crash Through`, `Hexgold Halberd`, `Iroas` |
| 8 | **token doubler** | dobra a curva quadrática | `Anointed Procession`, `Mondrak` |

**Termos que morreram com a troca:** `Knight` (typal), `Human` (typal), `first strike → double
strike`, `untap white attacking creature`. As fichas não são nenhum desses tipos.

---

## 3. Reclassificação das 99 cartas sob o eixo Otharri

Legenda de mudança: **↑↑** sobe muito · **↑** sobe · **=** inalterada · **↓** desce ·
**↓↓** desce muito (o texto deixou de funcionar). Categorias do `CLAUDE.md`.

### 3.1 Criaturas (27)

| Carta | CMC | Tipo | Categorias | Δ | Por que mudou sob Otharri |
|---|---|---|---|---|---|
| Belladonna Took | 2 | Leg. Halfling Citizen 2/2 | tema, **draw** | **↑↑** | `Whenever a token you control enters` — Otharri cria 1,2,3,4 fichas **de uma vez**: a 2ª resolução **compra uma carta** e a 3ª põe **+1/+1 em cada criatura**. Do 3º ataque em diante é **saque + anthem permanente todo turno**. Era marginal ("deck faz poucos tokens"); virou motor |
| Dawnstrike Vanguard | 6 | Human Knight 4/5 | tema | **↑↑** | `if you control two or more tapped creatures` no end step → as fichas entram **viradas e permanecem viradas** → condição **sempre satisfeita** → +1/+1 **permanente em todo o board todo turno**. O passe anterior a marcou como **anti-sinergia** (Tori destapava); **a Tori saiu da zona de comando, a anti-sinergia sumiu**. Reversão completa |
| The Circle of Loyalty (art.) | 6→~2 | Leg. Artifact | tema, **wincon** | **↑↑** | `Creatures you control get +1/+1` **estático** — o melhor anthem do deck para o enxame. Affinity for Knights ainda barateia |
| Jor Kadeen, the Prevailer | 5 | Leg. Human Warrior 5/4 | tema, **wincon** | **↑↑** | metalcraft `+3/+0` **estático a todas as criaturas** → 6 fichas viram 5/2 cada = **+18 de poder**. Era "aumento de dano"; virou fechador |
| Éomer, Marshal of Rohan | 4 | Leg. Human Knight 4/4 | tema, **wincon** | **↑** | `whenever one or more other attacking legendary creatures you control die → untap all + combate extra`. Otharri é **legendário e atacante**, e a morte dele agora é **desejável**: combate extra *e* ele cai no cemitério, de onde a recursão o traz. Era "condição que o deck não controla e não quer" |
| Esgaroth Garrison | 5 | Human Soldier */5 | tema, draw | **↑** | poder = nº de criaturas; com 6–10 fichas é 7/5–11/5 |
| Luxknight Breacher | 4 | Human Knight 2/2 | tema | **↑** | entra com contador por criatura **e** artefato; com o enxame no board é um 8/8+ |
| Sheriff of Safe Passage | 3 | Human Knight 0/0 | tema | **↑** | contadores = outras criaturas; plot `{1}{W}` guarda da vaza de wipe |
| Knight Luminary | 4 | Human Knight 3/2 | tema | **↑** | ETB **token** → alimenta Belladonna/Valor in Akros/Warleader's Call; warp recompra |
| Syr Alin, the Lion's Claw | 5 | Leg. Human Knight 4/4 | tema | **=** | `whenever attacks, other creatures get +1/+1` — **alcança as fichas** por ordenação (§1.3b). Em board largo vale mais em números absolutos; em qualidade de slot, igual |
| Honored Crop-Captain | 2 | Human Warrior 3/2 | tema | **=** | idem — `other attacking creatures +1/+0` alcança as fichas por ordenação. 2 manas, corpo 3/2 RW |
| Fireborn Knight | 4 | Human Knight 2/3 | tema | **=** | double strike próprio; recebe os anthems estáticos. Custo `{R/W}×4` é flexível |
| Adriana, Captain of the Guard | 5 | Leg. Human Knight 4/4 | tema | **↓** | **melee não dispara para as fichas** (ruling 2016-08-23: criaturas que entram atacando nunca foram declaradas atacantes). A cláusula "other creatures you control have melee" — que era o valor dela — **não alcança 80% do board futuro**. Continua um 4/4 RW com melee própria |
| Inspiring Veteran | 2 | Human Knight 2/2 | tema | **↓↓** | `Other **Knights** get +1/+1` — as fichas são **Rebel**. O anthem passa a valer só para os Knights não-ficha, e a contagem de Knights encolhe conforme o deck vira tokens |
| Paladin Danse, Steel Maverick | 3 | Leg. Art. Cre. Synth Knight 3/3 | tema, **proteção** | **↓↓** | indestrutível só a **artefatos ou Humanos** — as fichas Rebel **não são nem um nem outro**. A principal resposta a wipe do deck **não protege o board que o deck passa a ter**. Achado crítico para a dor 4 |
| Kwende, Pride of Femeref | 4 | Leg. Human Knight 2/2 | tema | **↓↓** | `first strike → double strike`: nenhuma ficha tem first strike, e a base de first-strikers encolhe com os cortes |
| Parhelion Patrol | 4 | Human Knight 2/3 | tema | **↓** | mentor exige alvo de **poder menor**; as fichas são poder 2 e ela é poder 2 → **mentor nunca as alcança**. Sobra um 2/3 voador com vigilância |
| Cavalry Drillmaster | 2 | Human Knight 2/1 | tema | **↓** | ETB one-shot +2/+0 e first strike (ligava com Kwende, que caiu) |
| Rohirrim Lancer | 1 | Human Knight 1/1 | tema | **↓** | menace num 1/1; o deck não precisa mais de corpos pequenos avulsos — ele os fabrica |
| Lake-town Lookout | 1 | Human Scout 1/1 | tema, draw | **↓** | idem; recruit ao morrer é loot, não saque |
| Knight of Sorrows | 5 | Human Knight 3/3 | tema | **↓** | bloqueio adicional é defesa; Otharri quer atacar. Anti-tema em CMC 5 |
| Relentless Rohirrim | 4 | Human Knight 4/3 | tema | **=** | corpo vermelho 4/3; ETB só "the Ring tempts you" |
| Fervent Cathar | 3 | Human Knight 2/1 | tema | **=** | haste + ETB remove um bloqueador (ajuda o alfa das fichas) |
| Recruitment Officer | 1 | Human Soldier 2/1 | tema, draw | **=** | dig repetível por `{3}{W}` continua sendo saque real |
| Syr Carah, the Bold | 5 | Leg. Human Knight 3/3 | tema, draw | **=** | impulse ao causar dano ao jogador; `{T}`: 1 de dano |
| Éomer of the Riddermark | 5 | Leg. Human Knight 5/4 | tema | **=** | token ao atacar **se** tiver o maior poder — Jor Kadeen e anthems ajudam a satisfazer |
| Elite Interceptor // Rejoinder | 1 | Human Wizard 1/2 | tema, draw | **=** | Rejoinder pode **desvirar** uma ficha Rebel (destrava a recursão) + compra |

### 3.2 Artefatos (7 não-criatura) e Encantamentos (5)

| Carta | CMC | Tipo | Categorias | Δ | Por quê |
|---|---|---|---|---|---|
| Valor in Akros | 4 | Enchantment | tema, **wincon** | **↑↑** | `Whenever a creature you control enters, creatures you control get +1/+1 until EOT` — **N fichas = N gatilhos**. Ataque com 4 contadores: +4/+4 no time **inteiro**, incluindo as fichas recém-criadas. Passa a ser das melhores cartas do deck |
| Sol Ring | 1 | Artifact | **ramp**, tema | **=** | intocável; acelera Otharri para o T3–T4 |
| Arcane Signet | 2 | Artifact | **ramp**, tema | **=** | intocável |
| Ancestral Blade | 2 | Artifact — Equipment | tema | **↑** | ETB **cria ficha** (alimenta Belladonna/Valor/Warleader's Call) + artefato (metalcraft) |
| Trailblazer's Torch | 4 | Artifact — Equipment | tema, remoção parcial | **=** | initiative + 2 de dano a cada bloqueador (bom com enxame que força bloqueios) |
| True-Faith Censer | 2 | Artifact — Equipment | tema | **↓↓** | +1/+1 e vigilância a **uma** criatura; o bônus extra lê **Human** — as fichas não são |
| Squire's Lightblade | 1 | Artifact — Equipment | tema | **↓** | first strike a uma criatura; ligava com Kwende |
| Celebrate the Mountain-king | 4 | Enchantment | **remoção**, draw | **=** | melhor interação do deck (exila 1 não-terreno **por oponente**); recruit gera ficha |
| Seal of Cleansing | 2 | Enchantment | **remoção** | **=** | — |
| Sanctuary Lockdown | 3 | Enchantment | tema | **↓↓** | `Humans you control get +1/+1` — fichas Rebel não recebem. Ativação exige **dois Humanos desvirados**, e o deck que Otharri quer não tem Humanos de sobra desvirados |
| Vigilante Justice | 4 | Enchantment | tema, wincon parcial | **↓↓** | `Whenever a **Human** you control enters` — o motor de fichas do Otharri **nunca** a dispara. Era uma "wincon parcial"; virou carta quase morta. **`Warleader's Call` é literalmente a mesma carta lendo "a creature"** |
| Basri Ket (PW) | 3 | Leg. Planeswalker | tema, **proteção** | **↑** | `−2`: cria ficha por **criatura não-ficha** atacante — Otharri é não-ficha. `+1`: contador **e indestrutível** — agora o alvo certo é óbvio (Otharri, o motor). Emblema põe contador em cada criatura todo combate |

### 3.3 Instantâneos (13)

| Carta | CMC | Categorias | Δ | Leitura sob Otharri |
|---|---|---|---|---|
| Duty Beyond Death | 2 | **proteção em massa**, tema | **↑↑** | indestrutível ao time + **+1/+1 permanente em cada**. O custo adicional (sacrificar uma criatura) era doloroso; agora se paga com **uma ficha 2/2 descartável**. De "resposta cara" para "resposta barata com upside permanente" |
| Raise the Alarm | 2 | tema | **↑** | 2 fichas em **instant** → 2 gatilhos de Belladonna/Valor/Warleader's Call, no fim do turno do oponente |
| Pride of Conquerors | 2 | tema, wincon parcial | **↑** | **instant**: cast depois que as fichas existem. Ascend é trivial → **+2/+2 no enxame inteiro** por 2 manas. Em board de 10 fichas é +20/+20 |
| Gods Willing | 1 | proteção | **↑** | 1 mana para salvar **o motor**. §1.4 inverte a leitura: proteção de alvo único deixou de ser excesso, porque o alvo que importa é **um** |
| Feat of Resistance | 2 | proteção | **↑** | idem + contador permanente |
| Adamant Will | 2 | proteção | **=** | idem, sem o contador. É a 3ª de três proteções de alvo único equivalentes |
| Swift Reckoning | 2 | **remoção** | **=** | (sorcery pelo oracle) remoção de criatura virada |
| Zealous Display | 3 | tema | **↓** | +2/+0 ao time; o `untap` só vale **fora do seu turno** e existia para a Tori. `Pride of Conquerors` faz mais por 1 mana a menos |
| Djeru's Renunciation | 2 | tema, draw(cycling) | **↓** | tapar 2 bloqueadores; o enxame resolve bloqueio por volume, não por Falter parcial |
| Gideon's Triumph | 2 | remoção (edict) | **↓** | oponente escolhe a pior criatura; metade do texto lê planeswalker Gideon, que o deck não tem |
| Disenchant | 2 | remoção art/ench | **=** | — |
| Invoke the Divine | 3 | remoção art/ench | **↓** | pior que Disenchant por 1 mana a mais |
| Miraculous Recovery | 5 | recursão | **↓** | devolve **1** criatura por 5; Otharri se devolve sozinho por 4 e sem gastar carta |

### 3.4 Feitiços (10)

| Carta | CMC | Categorias | Δ | Leitura sob Otharri |
|---|---|---|---|---|
| Crash Through | 1 | tema, draw(cantrip) | **↑↑** | **trample a todas as criaturas** — o enxame é de 2/2 **terrestres**, e chump-block é a única defesa que funciona contra ele. Cantrip por 1 mana. Havia uma cópia sobressalente na caixa |
| Late to Dinner | 4 | **recursão** | **↑↑** | devolve Otharri do cemitério **desvirado**; com haste ele ataca no mesmo turno. A recursão própria dele volta **virado** — as duas se complementam |
| Basri's Solidarity | 2 | tema | **↑** | contadores **permanentes** em cada criatura; escala com o enxame. Sorcery → só alcança as fichas de ataques anteriores |
| Inspiring Roar | 4 | tema | **↓** | mesmo efeito de Basri's Solidarity por **2 manas a mais** |
| Bond of Discipline | 5 | tema (Falter) | **↓** | 5 manas por um alfa; o enxame já não precisa disso |
| Warlord's Fury | 1 | tema, draw(cantrip) | **↓** | first strike ao time — ligava com Kwende (↓↓). Sem Kwende é quase um cantrip puro |
| Shoulder to Shoulder | 2 | tema, draw(cantrip) | **↓** | 2 contadores em 2 alvos; card-neutra |
| Reduce to Memory | 3 | **remoção** | **=** | remoção mais ampla do deck |
| Remember the Fallen | 3 | recursão à mão | **=** | — |
| Swift Reckoning | — | (contada acima) | — | — |

### 3.5 Terrenos (36 · 6 não-básicos + 30 básicos)

| Carta | Δ | Observação |
|---|---|---|
| Command Tower | **=** | intocável |
| Boros Guildgate · Stone Quarry · Wind-Scarred Crag | **=** | entram virados, sem upside; o `manabase-engineer` decide |
| Fields of Strife | **=** | surveil 1 por `{2}{R}{W},{T}` — custo proibitivo |
| Sandstone Bridge | **↓** | ETB dá +1/+1 e **vigilância** a **uma** criatura — o efeito era desenhado para a Tori |
| 15 Mountain + 15 Plains | **=** | base; **o deck está −2 na base de 38** (`manabase-engineer`) |

### 3.6 Resumo numérico da reclassificação

| Δ | Nº de cartas | Peso |
|---|---|---|
| **↑↑** (viraram motor) | 9 | Belladonna Took, Dawnstrike Vanguard, The Circle of Loyalty, Jor Kadeen, Valor in Akros, Duty Beyond Death, Crash Through, Late to Dinner, (+ Otharri) |
| **↑** | 13 | Éomer Marshal, Esgaroth Garrison, Luxknight Breacher, Sheriff, Knight Luminary, Ancestral Blade, Basri Ket, Raise the Alarm, Pride of Conquerors, Gods Willing, Feat of Resistance, Basri's Solidarity, (Tori — §6) |
| **=** | 24 | — |
| **↓** | 18 | — |
| **↓↓** | 6 | Inspiring Veteran, Paladin Danse, Kwende, True-Faith Censer, Sanctuary Lockdown, Vigilante Justice |

**Leitura:** o deck tem **mais** cartas que sobem do que se supunha. O passe anterior classificou
43 slots como "pump temporário de Limited"; sob Otharri, **22 deles viram motor ou melhoram**.
O que quebra não é o "pump" — é o **typal**: tudo que lê `Knight`, `Human` ou `first strike`
deixou de ver o board. Essa é a fratura real da troca de comandante, e ela é **mais estreita e
mais funda** do que a hipótese do briefing sugeria.

---

## 4. Pool temático sob o eixo Otharri — 42 candidatas

Regra 3: mínimo **2 pontos de sinergia** explicitados. Origem: `lista.txt` (já existe fisicamente,
custo zero) · `caixa` (sobressalente, custo zero) · `compra`.
Preços US$ = **estimativa (Scryfall)**, nunca régua. R$ marcados são LigaMagic com a data.

### 4.1 Motor de fichas e escalada de experiência (8)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 1 | **Neyali, Suns' Vanguard** | 4 | Leg. Human **Rebel** RW 3/3 | (a) `Attacking tokens you control have double strike` é **estática** → as fichas que entram atacando **têm double strike** = **dobra o dano do enxame inteiro** (6 fichas 2/2 = 24 em vez de 12); (b) saque: exila o topo quando fichas atacam; (c) **é Rebel não-ficha** → destrava a recursão do Otharri mesmo sem fichas; (d) legendária → ficha do Circle of Loyalty. **4 eixos** | compra | US$ 2,69 |
| 2 | **Adeline, Resplendent Cathar** | 3 | Leg. Human Knight W */4 | (a) `Whenever you attack` → 1 ficha **por oponente** (3/turno em pod) — empilha com o gatilho do Otharri no mesmo ataque; (b) poder = nº de criaturas → cresce com o enxame; (c) vigilância; (d) Human Knight alimenta Inspiring Veteran/Circle of Loyalty. **Regra 5:** foi *descartada como comandante* em 2026-09-19 (colapsaria a identidade para mono-W). **Como carta do 99 essa objeção não existe** — W ⊆ RW. Mudança declarada: o papel | compra | US$ 2,85 |
| 3 | **Hero of Bladehold** | 4 | Human Knight 3/4 | (a) 2 fichas **viradas e atacando** por ataque; (b) battle cry — ela **foi declarada atacante**, então dispara, e por ordenação (§1.3b) o bônus alcança as fichas do Otharri; (c) Human Knight | compra | US$ 0,61 |
| 4 | **Anim Pakal, Thousandth Moon** | 3 | Leg. Human Soldier RW 1/2 | (a) `Whenever you attack with one or more non-Gnome creatures` → X fichas viradas e atacando, **X cresce todo turno** (mesma curva quadrática do Otharri); (b) os contadores ficam **nela**, não no board → parcialmente resiliente; (c) legendária | compra | US$ 4,64 |
| 5 | **Assemble the Legion** | 5 | Enchantment RW | (a) fichas **com haste** em quantidade crescente todo upkeep, **sem depender de atacar**; (b) **não é criatura** → sobrevive a board wipe e **reconstrói sozinha** (dor 4); (c) fichas RW recebem todos os anthems | compra | US$ 0,34 |
| 6 | **Secure the Wastes** | X+1 | Instant W | (a) X corpos em **instant** — pós-wipe, no fim do turno do oponente; (b) escala com o ramp que a fase 4 vai trazer; (c) X fichas = X gatilhos de Valor in Akros/Belladonna | compra | s/ cotação |
| 7 | **Mobilization** | 3 | Enchantment W | (a) `{2}{W}`: ficha à vontade = mana vira corpo (reconstrução pós-wipe); (b) **não é criatura** → sobrevive a wipe. *Ressalva:* a cláusula de vigilância lê **Soldier** e não alcança as fichas Rebel | compra | US$ 0,75 |
| 8 | **Legion's Landing // Adanto** | 1 | Leg. Ench. → Land | (a) ficha T1; (b) **transforma ao atacar com 3+ criaturas** — trivial com Otharri — virando **terreno que gera fichas** (resiliência a wipe + cobre parte do −2 de terrenos) | compra | US$ 3,46 |

### 4.2 Anthems estáticos — o multiplicador quadrático (7)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 9 | **The Circle of Loyalty** | 6→~2 | Leg. Artifact | (a) `+1/+1` **estático a todas**; (b) artefato (metalcraft do Jor Kadeen); (c) ficha 2/2 por legendária conjurada | **lista.txt** | — |
| 10 | **Jor Kadeen, the Prevailer** | 5 | Leg. Human Warrior | (a) metalcraft `+3/+0` **estático a todas** — 6 fichas viram 5/2; (b) corpo 5/4 first strike; (c) RW | **lista.txt** | — |
| 11 | **Radiant Destiny** (nomeando **Rebel**) | 3 | Enchantment W | (a) `+1/+1` às fichas Rebel; (b) **vigilância** com city's blessing (10 permanentes é trivial) → as fichas **não viram ao atacar** → há **sempre um Rebel desvirado** para a recursão **e** bloqueadores contra o crack-back; (c) não é criatura | compra | US$ 0,31 |
| 12 | **Warleader's Call** | 3 | Enchantment RW | (a) `+1/+1` estático a todas; (b) **1 de dano a cada oponente por criatura que entra** → 4 fichas = 12 de dano num pod de 3, **ignorando bloqueadores**; (c) não é criatura. É literalmente o `Vigilante Justice` que lê "a creature" em vez de "a Human" | compra | US$ 6,75 |
| 13 | **Intangible Virtue** | 2 | Enchantment W | (a) `+1/+1` **e vigilância** a **fichas** — e quase todo o board futuro é ficha; (b) vigilância destrava a recursão (§1.2); (c) 2 manas | compra | US$ 0,32 |
| 14 | **Dictate of Heliod** | 5 | Enchantment W | (a) `+2/+2` estático — o maior anthem disponível; (b) **flash** → surpresa em combate; (c) não é criatura | compra | US$ 0,51 |
| 15 | **Legion's Initiative** | 2 | Enchantment RW | (a) **criaturas vermelhas +1/+0** — as fichas Rebel **são vermelhas**; (b) exílio-e-retorno salva as não-fichas de um wipe. **Ressalva grave a declarar:** fichas exiladas **deixam de existir** — a habilidade é anti-sinergia com o board que o deck constrói. Entra no pool só pelo anthem de 2 manas | compra | US$ 1,76 |

### 4.3 Gatilhos de "criatura entra" — N fichas = N disparos (6)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 16 | **Valor in Akros** | 4 | Enchantment | (a) N fichas = **N × +1/+1 no time inteiro** naquele combate; (b) não é criatura | **lista.txt** | — |
| 17 | **Belladonna Took** | 2 | Leg. Halfling Citizen | (a) 2ª ficha do turno = **compra uma carta** (dor 2); (b) 3ª ficha = **+1/+1 permanente em cada criatura**; (c) corpo 2/2 por 2 | **lista.txt** | — |
| 18 | **Dawnstrike Vanguard** | 6 | Human Knight 4/5 | (a) as fichas entram **viradas** → condição `2+ tapped creatures` sempre satisfeita → **+1/+1 permanente no board todo turno**; (b) lifelink 4/5. *Atrito:* CMC 6 na pior faixa da curva | **lista.txt** | — |
| 19 | **Cathars' Crusade** | 5 | Enchantment W | (a) N fichas = **N gatilhos**, cada um pondo contador em **cada** criatura → crescimento exponencial **permanente**; (b) contadores sobrevivem a wipes de −X/−X e ao fim do turno; (c) não é criatura | compra | US$ 10,78 |
| 20 | **Impact Tremors** | 2 | Enchantment R | (a) 1 dano a cada oponente por criatura que entra — dano que **não depende de combate nem de bloqueio**; (b) não é criatura. *Nota:* `Warleader's Call` faz isso **mais** anthem por 1 mana a mais e é mais barata | compra | R$ 17,78 (LM, 28 d) |
| 21 | **Guide of Souls** | 1 | Human Cleric 1/2 | (a) energia + vida por criatura que entra → o enxame enche o tanque; (b) o gasto de energia põe **2 contadores + voar** num atacante; (c) 1 mana | compra | US$ 2,57 |

### 4.4 Combate extra — o multiplicador do próprio motor (4)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 22 | **Combat Celebrant** | 3 | Human Warrior 4/1 | (a) exert ao atacar → **desvira todas as outras criaturas** (inclusive Otharri e as fichas recém-criadas) **+ combate adicional** → **segundo gatilho de ataque do Otharri no mesmo turno**, com o contador já incrementado; (b) corpo 4/1 vermelho; (c) as fichas desviradas atacam de novo | compra | US$ 2,32 |
| 23 | **Great Train Heist** | 1 (+spree) | Instant R | (a) modo `{2}{R}`: desvira todas as criaturas + combate adicional — **em instant**; (b) modo `{2}`: +1/+0 e first strike ao time; (c) modo `{R}`: Treasure por criatura que conectar = **ramp**. Três modos, escolhe-se quantos couberem | compra | US$ 5,15 |
| 24 | **Éomer, Marshal of Rohan** | 4 | Leg. Human Knight | (a) morte de legendária atacante → desvira tudo + combate extra; **Otharri é legendário atacante e a morte dele é desejável** (cai no cemitério de onde a recursão o traz); (b) corpo 4/4 com haste | **lista.txt** | — |
| 25 | **Aggravated Assault** | 3 | Enchantment R | (a) combate extra **repetível** por `{3}{R}{R}` = contadores de experiência sem teto; (b) não é criatura. **Alerta de preço:** US$ 35,79 (estimativa Scryfall) — provavelmente fora da régua; listada para o orquestrador decidir com a cotação real | compra | US$ 35,79 |

### 4.5 Rebel desvirado, recursão e proteção do motor (6)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 26 | **Hexgold Halberd** | 2 | Artifact — Equipment R | (a) *For Mirrodin!* → **ficha 2/2 vermelha Rebel desvirada** = Rebel de plantão para a recursão; (b) **artefato** (metalcraft Jor Kadeen, contagem Luxknight); (c) dá **first strike e trample** ao equipado no seu turno — evasão que o enxame precisa | compra | US$ 0,21 |
| 27 | **Barbed Batterfist** | 2 | Artifact — Equipment R | (a) mesma ficha Rebel desvirada; (b) artefato; (c) equip `{1}` (o mais barato do deck) | compra | US$ 0,19 |
| 28 | **Sparring Regimen** | 3 | Enchantment W | (a) `Whenever you attack`: contador **permanente** num atacante **e o desvira** → gera Rebel desvirado todo turno **e** engorda o board; (b) ETB learn (filtro); (c) não é criatura | compra | US$ 0,26 |
| 29 | **Late to Dinner** | 4 | Sorcery W | (a) devolve Otharri do cemitério **desvirado** → com haste ele ataca no mesmo turno; (b) Food é artefato (metalcraft) | **lista.txt** | — |
| 30 | **Duty Beyond Death** | 2 | Instant W | (a) indestrutível ao time por 2 manas; (b) **+1/+1 permanente em cada criatura**; (c) o custo de sacrificar se paga com **uma ficha descartável** — custo que era proibitivo e deixou de ser | **lista.txt** | R$ 0,43 (LM) |
| 31 | **Gods Willing** | 1 | Instant W | (a) 1 mana para salvar **o motor** (§1.4: o alvo que importa é um); (b) proteção de cor também fura bloqueio; (c) scry 1 | **lista.txt** | — |

### 4.6 Proliferate — contador de experiência extra (3)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 32 | **Karn's Bastion** | 0 | Land | (a) `{4},{T}`: proliferate → **+1 contador de experiência = +1 ficha por ataque para sempre**; (b) também engorda Sheriff/Luxknight/Basri Ket; (c) **é terreno** → ajuda no −2 da base sem custar slot de feitiço | compra | US$ 2,39 |
| 33 | **Grateful Apparition** | 2 | Spirit 1/1 W | (a) voa → conecta quase sempre → proliferate **todo turno**; (b) +1 experiência por conexão dobra a taxa de crescimento do enxame; (c) corpo evasivo barato | compra | US$ 0,37 |
| 34 | **Contagion Clasp** | 2 | Artifact | (a) proliferate repetível; (b) **artefato** (metalcraft); (c) ETB −1/−1 mata uma criatura X/1 | compra | s/ cotação |

### 4.7 Evasão para as fichas 2/2 e fechamento (4)

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem | Preço |
|---|---|---|---|---|---|---|
| 35 | **Crash Through** | 1 | Sorcery R | (a) **trample a todas** — o chump-block é a única defesa que funciona contra 2/2 terrestres; (b) cantrip por 1 mana | **lista.txt** (+1 cópia na **caixa**) | — |
| 36 | **Inti, Seneschal of the Sun** | 2 | Leg. Human Knight R | (a) `Whenever you attack`: descarta → **contador permanente + trample** num atacante; (b) o descarte **exila o topo para jogar** = vantagem de carta ligada ao ataque; (c) 2 manas | compra | US$ 2,04 |
| 37 | **Iroas, God of Victory** | 4 | Leg. Ench. Cre. RW 7/4 | (a) **menace no time inteiro** → o enxame precisa de 2 bloqueadores por ficha; (b) **previne todo dano aos seus atacantes** → o alfa nunca perde criatura; (c) indestrutível e só é criatura com devoção alta → dificílimo de remover | compra | US$ 12,38 |
| 38 | **Phoenix Chick** | 1 | Phoenix R 1/1 | (a) voa, haste, 1 mana; (b) **se devolve do cemitério** ao atacar com 3+ criaturas (`{R}{R}`) **virada e atacando com contador** → sobrevive a board wipe (dor 4); (c) Phoenix, como Otharri (sem typal, mas mesma linha de recursão) | compra | US$ 0,31 |

### 4.8 Sobressalentes da caixa que servem ao novo eixo (4)

Varredura completa das 113 (regra 7). Fora as já citadas, estas continuam valendo:

| # | Carta | CMC | Tipo | Sinergias (2+) | Origem |
|---|---|---|---|---|---|
| 39 | **Glass Casket** | 2 | Artifact | (a) remoção de criatura ≤3 MV — o deck tem 4 remoções reais; (b) **artefato** (metalcraft Jor Kadeen, contagem Luxknight) | **caixa** |
| 40 | **Fumigate** | 5 | Sorcery W | (a) o wipe que falta (0/2–4); (b) **assimétrico** combinado com `Duty Beyond Death` (já no deck) — e mesmo simétrico, **os contadores de experiência sobrevivem** e o Otharri reconstrói | **caixa** |
| 41 | **Blast Zone** | 0 | Land | (a) wipe modular contra fichas/prisões; (b) **terreno** → parte do −2 da base | **caixa** |
| 42 | **Knight Watch** | 5 | Sorcery W | (a) 2 corpos 2/2 com **vigilância**; (b) 2 gatilhos de Valor in Akros/Belladonna/Warleader's Call. *Marginal:* 5 manas por 2 corpos é caro para o novo eixo — listada como reserva | **caixa** |

**Total do pool: 42 candidatas** · origem: **10 de `lista.txt`** · **4 da caixa** · **28 de compra**.

### 4.9 Dispensas da caixa com justificativa escrita (regra 7)

Além das dispensas já registradas no passe anterior (azuis/verdes fora da identidade; pacote de
artefatos do Inspirit; filler de Limited), **estas mudaram de veredito com o novo eixo**:

| Carta da caixa | Por que **não** cobre a função |
|---|---|
| **Youthful Knight**, **Embereth Paladin**, **Bogardan Lancer** | Eram candidatas por **contagem de Knight** e por *first strike* → **Kwende**. Os dois eixos caíram (§3.1): as fichas não são Knights e Kwende desceu a ↓↓. São corpos 1/1–4/1 sem gatilho — **1 ponto de sinergia**, reprovam na regra 3 |
| **Leonin Bola**, **S.H.I.E.L.D. Spy Kit** | Existiam para consumir o `untap` da Tori (Bola) ou premiar **ataque solo** (Spy Kit). Otharri ataca **acompanhado por definição** e não destapa nada. **0–1 ponto** |
| **Ori, Keeper of Songs** | Entrava por ser branca (untap da Tori) e legendária. Sem o untap, sobra "legendária → ficha do Circle of Loyalty" = **1 ponto** |
| **Joust** | Remoção por fight com `+2/+1 se for Knight`. O bônus continua existindo (há Knights não-ficha), mas fight com um 2/2 Rebel é suicídio e o deck deixou de ser Knight-cêntrico. Continua **elegível como remoção pura** — devolvo ao `interaction-specialist`, não é chamada minha |
| **Ratchet Bomb** | Wipe modular que mata **por valor de mana** — as fichas do Otharri são MV 0, ou seja, **ele mata o próprio board** no ajuste mais comum. Anti-sinergia direta com o novo eixo |
| **Phyrexian Revoker** | Stax leve + artefato (metalcraft). Mantém 2 pontos, mas não é tema — devolvo ao `interaction-specialist` |
| **Purphoros, God of the Forge** *(não está na caixa; avaliado por completude)* | Seria perfeito (2 de dano por criatura que entra). **R$ 129,99** (LigaMagic, 28 d) = **65% do teto de R$ 200**. Dispensado por preço, não por função |

---

## 5. Candidatos a corte — ficha F1–F7 completa

Protocolo da §2 do `card-evaluation-checklist.md`. Condições assumidas **idênticas** para quem
sai e para quem entra (§3 do checklist): Otharri em campo com 3–4 contadores, 6–10 fichas Rebel
no board, `The Circle of Loyalty` e/ou `Jor Kadeen` em campo.

**Consulta ao `decisions.md` (regra 5):** nenhuma das cartas abaixo tem histórico de corte —
o deck foi montado fora do pipeline e esta é a primeira rodada. Nenhuma é reposição.

---

### K1 · Vigilante Justice — `{3}{R}` Enchantment — **corte limpo**
- **F1** `Whenever a **Human** you control enters, this enchantment deals 1 damage to any target.`
- **F2** Sem corpo. Não tapa, não sacrifica, não bloqueia.
- **F3** Encantamento — o deck não tem constelação nem contagem de encantamentos.
- **F4** Não recebe nada (anthems não afetam encantamento).
- **F5** Não dá nada a outras cartas.
- **F6** Turno 4. A partir dali só dispara se **Humanos** entrarem.
- **F7** **Atrito fatal:** o motor de fichas do comandante cria **Rebels vermelhos**. A carta
  **nunca** dispara pelo Otharri. Sob o eixo antigo (25 Humanos) era wincon parcial; sob o novo,
  os Humanos que entram são só as criaturas conjuradas — 1 a 2 por turno no melhor caso.
- **Protocolo:** funções = [dano recorrente por criatura que entra, alcance ao rosto].
  **Ambas cobertas e ampliadas por `Warleader's Call`** (lê `a creature`, atinge **cada
  oponente**, e ainda é anthem estático) e por `Impact Tremors`. Nenhuma função descoberta.

### K2 · Sanctuary Lockdown — `{2}{W}` Enchantment — **corte limpo**
- **F1** `Humans you control get +1/+1.` / `{2}, Tap two untapped Humans you control: Tap target creature an opponent controls.`
- **F2** Sem corpo. **Consome** corpos (dois Humanos desvirados por ativação).
- **F3** Encantamento; não alimenta contagem alguma do deck.
- **F4** Não recebe nada.
- **F5** Anthem **typal de Human** — não alcança as fichas Rebel.
- **F6** Turno 3.
- **F7** **Duplo atrito:** (i) o anthem erra o board; (ii) a ativação exige **dois Humanos
  desvirados**, e o deck perdeu a fonte de destapar (a Tori saiu da zona de comando) — os
  Humanos que atacam ficam virados.
- **Protocolo:** funções = [anthem, tap de criatura do oponente].
  Anthem → **coberto e ampliado** por `The Circle of Loyalty` (+1/+1 a **todas**), `Radiant
  Destiny`, `Warleader's Call`. Tap → **descoberto**; custo declarado e aceito: tapar uma criatura
  por `{2}` + dois corpos é o pior efeito de interação disponível no slot, e o
  `interaction-specialist` traz remoção real.

### K3 · True-Faith Censer — `{2}` Artifact — Equipment — **corte condicionado**
- **F1** `Equipped creature gets +1/+1 and has vigilance.` / `+1/+0 adicional se for Human.` / `Equip {2}`.
- **F2** Sem corpo próprio.
- **F3** **É artefato** — conta para metalcraft do `Jor Kadeen` (que subiu a ↑↑) e para os
  contadores do `Luxknight Breacher`. **Esta função é real e não é temática.**
- **F4** Não recebe nada.
- **F5** Dá +1/+1 e **vigilância** a **uma** criatura; o bônus extra lê **Human**.
- **F6** Turno 2, equip `{2}` a partir do turno 4 — lento.
- **F7** O cláusula de Human erra as fichas; vigilância em **uma** criatura é irrelevante quando
  `Radiant Destiny`/`Intangible Virtue` dão vigilância ao **enxame**.
- **Protocolo:** funções = [buff de alvo único, vigilância, **contagem de artefato**].
  Buff e vigilância → cobertos por `Radiant Destiny`/`Intangible Virtue`/`Circle of Loyalty`.
  **Contagem de artefato → fica descoberta se nada entrar no lugar.** `Hexgold Halberd` e
  `Barbed Batterfist` (§4.5) são artefatos **e** trazem ficha Rebel — cobrem a contagem com
  sobra. **Corte fica limpo *se* uma dessas entrar; caso contrário, é corte com função
  descoberta.** Decisão do orquestrador.

### K4 · Kwende, Pride of Femeref — `{3}{W}` 2/2 — **corte condicionado**
- **F1** `Double strike` próprio. `Creatures you control with first strike have double strike.`
- **F2** Corpo 2/2 por 4 — recebe os anthems (com Circle + Jor Kadeen vira 6/3 com double strike).
- **F3** Human Knight → alimenta `Inspiring Veteran` e affinity do `Circle of Loyalty`.
- **F4** Recebe anthems, contadores, equipamentos.
- **F5** Concede double strike — **a nenhuma ficha** (nenhuma tem first strike), e a base de
  first-strikers do deck é `Cavalry Drillmaster` (ETB, one-shot), `Fireborn Knight`,
  `Syr Alin`, `Warlord's Fury`, `Squire's Lightblade` — todas ↓ ou candidatas a corte.
- **F6** Turno 4.
- **F7** A habilidade dele **depende de cartas que também estão saindo** — é o centro de um
  subpacote (first strike) que perdeu a razão de existir.
- **Protocolo:** funções = [corpo 2/2 com double strike, contagem Knight/Human, concessão de
  double strike]. Corpo e contagem → cobertos por `Neyali` (que dá **double strike a todas as
  fichas atacantes**, ou seja, faz o trabalho do Kwende no board que importa) e por
  `Adeline`/`Hero of Bladehold` (Human Knight). Concessão → **coberta e ampliada por `Neyali`**.
  **Corte limpo *se* Neyali entrar.** Sem Neyali, o corte deixa "multiplicador de dano" descoberto.

### K5 · Inspiring Veteran — `{R}{W}` 2/2 — **corte condicionado (fora do meu escopo decidir)**
- **F1** `Other Knights you control get +1/+1.`
- **F2** Corpo 2/2 por 2, RW.
- **F3** Human Knight.
- **F4** Recebe anthems e contadores.
- **F5** Anthem **typal de Knight**.
- **F6** **Turno 2** — é das poucas cartas que o deck pode jogar cedo.
- **F7** As fichas Rebel não recebem o anthem. O valor dele depende de **quantos Knights
  sobreviverem à reconstrução** — hoje são 21; o deck reconstruído provavelmente terá 8–12.
- **Protocolo:** funções = [anthem typal, corpo 2 CMC RW, contagem Knight/Human].
  **O veredito depende de uma contagem que ainda não existe** (quantos Knights ficam depois dos
  cortes de todas as fases). **Não corto** — devolvo ao orquestrador com a regra: se o deck final
  tiver **≤10 Knights**, o anthem vale menos que um slot de `Radiant Destiny`; se tiver **≥14**,
  ele fica.

### K6 · Paladin Danse, Steel Maverick — `{2}{W}` 3/3 — **corte condicionado (escopo de proteção)**
- **F1** `Vigilance, lifelink.` / `Exile Paladin Danse: Each creature you control that's an artifact or Human gains indestructible until end of turn.`
- **F2** Corpo 3/3 com vigilância e lifelink por 3 — **bom corpo**, ataca sem virar.
- **F3** **Artifact Creature** — conta duplo: metalcraft do `Jor Kadeen` **e** criatura.
- **F4** Recebe anthems (com Circle + Jor Kadeen é 7/4 com vigilância e lifelink).
- **F5** Indestrutível a **artefatos ou Humanos** — **não alcança as fichas Rebel**.
- **F6** Turno 3.
- **F7** **A função pela qual ela estava no deck (resposta a wipe) deixou de cobrir o board que
  o deck passa a ter.** Salva Otharri? Não — Otharri é **Phoenix**, nem artefato nem Humano.
  **A carta não protege nem o motor nem o enxame.**
- **Protocolo:** funções = [corpo 3/3 vigilância+lifelink, contagem de artefato, contagem de
  criatura, **proteção em massa**]. Corpo e contagens → cobertas. **Proteção em massa → fica
  descoberta, e é a dor 4 declarada pelo usuário.** `Duty Beyond Death` cobre (indestrutível a
  **todas**, sem restrição de tipo), mas é 1 carta. **Não corto:** a função é de outra
  especialidade. **Devolvo ao `interaction-specialist`** com o achado: *a melhor proteção em
  massa que o deck tem hoje não protege o board novo nem o comandante.*

### K7 · Parhelion Patrol — `{3}{W}` 2/3 — **corte limpo**
- **F1** `Flying. Vigilance. Mentor.`
- **F2** Corpo 2/3 voador com vigilância — ataca e bloqueia.
- **F3** Human Knight.
- **F4** Recebe anthems.
- **F5** Mentor: contador em atacante de **poder menor**. As fichas são **poder 2**, ela é
  **poder 2** → **mentor nunca as alcança**. Com qualquer anthem estático, ela e as fichas sobem
  juntas e a comparação continua empatada. Mentor é **letra morta** neste deck.
- **F6** Turno 4.
- **F7** 4 manas por 2/3 numa curva que já trava em 4–5 (15 cartas em CMC 4).
- **Protocolo:** funções = [corpo voador, vigilância própria, contagem Knight/Human, mentor].
  Corpo voador → coberto por `Grateful Apparition` (1/1 voador que **proliferate**) e pelo próprio
  Otharri. Vigilância → coberta por `Radiant Destiny`. Contagens → cobertas. Mentor → **já estava
  descoberto** (nunca funcionou). Nenhuma função real perdida.

### K8 · Inspiring Roar — `{3}{W}` Sorcery — **corte limpo**
- **F1** `Put a +1/+1 counter on each creature you control.`
- **F2–F4** Sem corpo, tipo irrelevante, não recebe nada.
- **F5** Contadores permanentes ao time — escala com o enxame (sobe em valor absoluto sob Otharri).
- **F6** Turno 4.
- **F7** **`Basri's Solidarity` faz exatamente o mesmo por `{1}{W}`.** Duas cópias funcionais,
  uma estritamente pior em 2 manas.
- **Protocolo:** função única → **integralmente coberta por `Basri's Solidarity`**, que fica.

### K9 · Shoulder to Shoulder — `{2}{W}` Sorcery — **corte limpo com custo declarado**
- **F1** Support 2 (contador em até **duas** criaturas) + compre 1 carta.
- **F2–F4** Sem corpo, sem tipo relevante, não recebe nada.
- **F5** 2 contadores + reposição.
- **F6** Turno 3.
- **F7** Num board de 6–10 fichas, distribuir **2** contadores é irrelevante; `Basri's Solidarity`
  põe **um em cada** por 1 mana a menos.
- **Protocolo:** funções = [contadores, cantrip]. Contadores → cobertos por `Basri's Solidarity`,
  `Duty Beyond Death`, `Cathars' Crusade`. Cantrip → **descoberto**; custo aceito e declarado —
  é exatamente o slot que a fase de draw troca por saque real.

### K10 · Zealous Display — `{2}{W}` Instant — **corte limpo**
- **F1** `Creatures you control get +2/+0 until end of turn. If it's not your turn, untap those creatures.`
- **F2–F4** Sem corpo, sem tipo relevante, não recebe nada.
- **F5** Pump temporário em massa (escala com o enxame) + untap **fora do seu turno**.
- **F6** Turno 3.
- **F7** A cláusula de untap foi desenhada para o eixo da Tori (destapar para bloquear) e é
  **inútil no turno do próprio jogador**. **`Pride of Conquerors` dá +2/+2 (não +2/+0) por
  `{1}{W}`**, também em instant, com ascend trivial.
- **Protocolo:** funções = [pump em massa instant, untap defensivo]. Pump → **coberto e superado
  por `Pride of Conquerors`** (que fica). Untap → descoberto; custo aceito: a vigilância em massa
  de `Radiant Destiny`/`Intangible Virtue` cobre a mesma necessidade de forma permanente.

### K11 · Knight of Sorrows — `{4}{W}` 3/3 — **corte limpo com custo declarado**
- **F1** `Can block an additional creature each combat. Afterlife 1.`
- **F2** Corpo 3/3 por 5, **defensivo**; o token de afterlife é um corpo extra.
- **F3** Human Knight.
- **F4** Recebe anthems.
- **F5** Não dá nada.
- **F6** Turno 5 — pior faixa da curva (9 cartas em CMC 5).
- **F7** Anti-tema: bloqueio duplo num deck cuja tese é atacar; afterlife premia a morte dela.
- **Protocolo:** funções = [corpo, contagem Knight/Human, bloqueio duplo, token ao morrer].
  Corpo e contagens → cobertos por `Hero of Bladehold`/`Adeline` (mesmos tipos, geram fichas).
  Bloqueio duplo → **descoberto**, custo aceito (a defesa agora vem da **vigilância em massa**,
  que deixa 6–10 corpos de pé). Token ao morrer → coberto com folga pelo motor do comandante.

### K12 · Rohirrim Lancer — `{R}` 1/1 — **corte limpo**
- **F1** `Menace.` **F2** 1/1 por 1. **F3** Human Knight. **F4** recebe anthems. **F5** nada.
  **F6** turno 1. **F7** um 1/1 é o corpo que o deck agora **fabrica de graça**, aos pares e trios.
- **Protocolo:** funções = [corpo barato, menace próprio, contagem Knight/Human]. Corpo → coberto
  pelo motor. Menace → **descoberto**, e a reposição correta é `Iroas` (menace ao **time**) ou
  `Crash Through`/`Hexgold Halberd` (trample), listadas no pool. Contagens → cobertas.

### K13 · Cavalry Drillmaster — `{1}{W}` 2/1 — **corte limpo**
- **F1** ETB: criatura alvo ganha +2/+0 e first strike até o fim do turno.
- **F2** 2/1 por 2. **F3** Human Knight. **F4** recebe anthems. **F5** buff one-shot + first
  strike (ligava com `Kwende`, K4). **F6** turno 2. **F7** o pagamento de first strike sai do deck.
- **Protocolo:** funções = [corpo 2 CMC, contagem, buff ETB one-shot, first strike].
  Corpo e contagem → cobertos. Buff/first strike → cobertos e superados por `Neyali`
  (double strike **permanente** nas fichas) e `Inti` (contador permanente + trample).

### K14 · Lake-town Lookout — `{W}` 1/1 — **corte limpo**
- **F1** 1/1; ao morrer, recruit (loot + ficha se descartar não-terreno).
- **F2** 1/1. **F3** Human Scout — **não** é Knight; alimenta só contagens de Human (que saíram).
  **F4** recebe anthems. **F5** nada. **F6** turno 1. **F7** recruit é **loot**, card-neutro —
  não é saque; e exige a morte dela.
- **Protocolo:** funções = [corpo 1 CMC, loot ao morrer]. Corpo → coberto pelo motor de fichas.
  Loot → **descoberto**; custo aceito e declarado: a fase de draw substitui por saque real, e o
  deck já conta com `Belladonna Took` (↑↑) para saque ligado a ficha.

### K15 · Squire's Lightblade — `{W}` Artifact — Equipment — **corte condicionado**
- **F1** Flash; equipado ganha +1/+0 e first strike; equip `{1}`.
- **F2** Sem corpo. **F3** **artefato** (metalcraft). **F4** nada. **F5** first strike a um alvo.
  **F6** turno 1. **F7** o pacote de first strike sai (K4).
- **Protocolo:** funções = [buff de alvo único, first strike, **contagem de artefato**].
  Buff/first strike → cobertos por `Neyali`. **Contagem de artefato → descoberta**, e o deck vai
  de 8 para 6–7 artefatos se K3 e K15 saírem juntos — **metalcraft do `Jor Kadeen` (↑↑) fica em
  risco**. Reposição obrigatória caso as duas saiam: `Hexgold Halberd` **e** `Barbed Batterfist`
  (2 artefatos + 2 fichas Rebel). Decisão do orquestrador.

### K16 · Bond of Discipline — `{4}{W}` Sorcery — **corte limpo com custo declarado**
- **F1** Tape todas as criaturas dos oponentes; suas criaturas ganham lifelink.
- **F2–F4** Sem corpo, sem tipo relevante, não recebe nada.
- **F5** Falter total + lifelink em massa.
- **F6** Turno 5, sorcery, pré-combate.
- **F7** 5 manas que **não adicionam nada ao board** num deck cuja força é largura de board;
  compete em custo com `Assemble the Legion` (que gera corpos todo turno) e com `Fumigate`.
- **Protocolo:** funções = [Falter em massa, lifelink em massa]. Falter → **descoberto**; custo
  aceito e declarado: a evasão passa a vir de **trample/menace permanentes** (`Crash Through`,
  `Hexgold Halberd`, `Iroas`) e de **combate extra**, que produzem mais dano por mana.
  Lifelink → coberto pelo próprio Otharri (lifelink) e dispensável.

### K17 · Gideon's Triumph — `{1}{W}` Instant — **corte limpo**
- **F1** Oponente alvo sacrifica uma criatura **à escolha dele** que atacou/bloqueou; dobra com
  planeswalker Gideon.
- **F2–F5** Sem corpo, sem tipo relevante, não recebe nem dá nada.
- **F6** Turno 2. **F7** o deck **não tem nenhum Gideon** (tem `Basri Ket`) — metade do texto é
  letra morta; como edict, remove a **pior** criatura do oponente.
- **Protocolo:** função única = [remoção condicional de baixa qualidade] → **coberta e superada**
  por `Glass Casket` (**já na caixa**) e pelo que o `interaction-specialist` trouxer.

### K18 · Sandstone Bridge — Land — **corte condicionado (escopo de manabase)**
- **F1** Entra virada; ETB: +1/+1 e **vigilância** a uma criatura.
- **F2–F5** É terreno; o ETB foi desenhado para o eixo de untap da Tori.
- **F6** — **F7** entra virada, o que num deck que quer `{3}{R}{W}` no T4–T5 custa um turno.
- **Protocolo:** funções = [fonte RW, buff ETB de alvo único]. Fonte → o `manabase-engineer`
  decide (o deck está **−2 na base de 38**). Buff → coberto por `Radiant Destiny`/`Intangible
  Virtue` (vigilância ao **enxame**). **Não corto** — é slot de terreno. Devolvo com a sugestão
  de trocar por `Karn's Bastion` (proliferate = +1 contador de experiência) ou `Blast Zone`
  (**já na caixa**), que resolvem o −2 **e** somam função.

---

### 5.1 Resumo dos cortes

| # | Carta | Tipo de corte | Função que exige reposição | Quem cobre |
|---|---|---|---|---|
| K1 | **Vigilante Justice** | limpo | nenhuma | Warleader's Call (amplia) |
| K2 | **Sanctuary Lockdown** | limpo | tap de criatura (custo aceito) | Circle of Loyalty / Radiant Destiny (anthem) |
| K7 | **Parhelion Patrol** | limpo | nenhuma (mentor nunca funcionou) | Grateful Apparition (voador) |
| K8 | **Inspiring Roar** | limpo | nenhuma | Basri's Solidarity |
| K9 | **Shoulder to Shoulder** | limpo | cantrip (custo aceito) | fase de draw |
| K10 | **Zealous Display** | limpo | untap defensivo (custo aceito) | Pride of Conquerors + vigilância em massa |
| K11 | **Knight of Sorrows** | limpo | bloqueio duplo (custo aceito) | vigilância em massa |
| K12 | **Rohirrim Lancer** | limpo | menace (custo declarado) | Iroas / Crash Through |
| K13 | **Cavalry Drillmaster** | limpo | nenhuma | Neyali / Inti |
| K14 | **Lake-town Lookout** | limpo | loot (custo aceito) | Belladonna Took |
| K16 | **Bond of Discipline** | limpo | Falter (custo declarado) | trample/menace permanentes + combate extra |
| K17 | **Gideon's Triumph** | limpo | nenhuma | Glass Casket (**caixa**) |
| K3 | **True-Faith Censer** | **condicionado** | contagem de artefato | só sai se Hexgold Halberd/Barbed Batterfist entrar |
| K4 | **Kwende** | **condicionado** | multiplicador de dano | só sai se **Neyali** entrar |
| K15 | **Squire's Lightblade** | **condicionado** | contagem de artefato (metalcraft do Jor Kadeen) | idem K3 — **não cortar as duas sem repor 2 artefatos** |
| K5 | **Inspiring Veteran** | **condicionado — não corto** | anthem typal | depende da contagem final de Knights (≤10 sai / ≥14 fica) |
| K6 | **Paladin Danse** | **condicionado — não corto** | **proteção em massa (dor 4)** | `interaction-specialist` — a proteção atual não cobre nem o enxame nem o comandante |
| K18 | **Sandstone Bridge** | **condicionado — não corto** | fonte RW | `manabase-engineer` |

**Fora do meu escopo de corte** (exercem função de outra especialidade — devolvo ao orquestrador
sem veredito): `Disenchant`, `Invoke the Divine`, `Expose to Daylight`, `Seal of Cleansing`
(4 cartas art/ench-only — o `interaction-specialist` decide quantas ficam), `Djeru's Renunciation`,
`Miraculous Recovery`, `Remember the Fallen`, `Adamant Will`, `Warlord's Fury`.

---

## 6. Veredito sobre Tori D'Avenant, Fury Rider como carta do 99

`{1}{R}{R}{W}` · 3/3 · Legendary Creature — Human Knight · CMC 4

```
Vigilance, trample
Whenever Tori D'Avenant attacks, all other attacking creatures you control get +1/+1 until end
of turn. Other red attacking creatures you control gain trample until end of turn.
Untap each other white attacking creature you control.
```

### Ficha F1–F7 (avaliada sob as **mesmas** condições das entradas — §3 do checklist)

| Eixo | Leitura |
|---|---|
| **F1** | Três linhas num gatilho de ataque: (i) **+1/+1 a todos os outros atacantes**; (ii) **trample às suas criaturas vermelhas atacantes**; (iii) **destapa cada outra criatura branca atacante**. Mais vigilância e trample próprios |
| **F2** | Corpo 3/3 com **vigilância** — ataca **e** fica de pé. Com `The Circle of Loyalty` + `Jor Kadeen` é um **7/4 com vigilância e trample**. Não pode ser tapada para custo (o deck não tem crew/convoke); bloqueia bem |
| **F3** | **Human Knight** e **legendária** — alimenta `Inspiring Veteran`, affinity do `The Circle of Loyalty` e o gatilho de legendária do próprio Circle |
| **F4** | Recebe **todos** os anthems estáticos, contadores (`Basri's Solidarity`, `Duty Beyond Death`, `Cathars' Crusade`), equipamentos e proteções |
| **F5** | **Este é o ponto que mudou.** O gatilho dela e o do Otharri vão para a pilha **juntos**, no mesmo passo de declarar atacantes, e **você escolhe a ordem** (§1.3b). Ordenando o do Otharri para resolver **primeiro**, as fichas já existem quando o gatilho da Tori resolve — e o efeito dela enxerga o board **na resolução**. Resultado: **as fichas recebem +1/+1** e, por serem **Rebels vermelhos**, **recebem trample**. Com 6 fichas isso é **+6/+6 distribuídos e trample no enxame inteiro** — e trample é **exatamente** o que fichas 2/2 terrestres precisam contra chump-block. A cláusula de untap (branca) **não** alcança as fichas, mas alcança as demais criaturas brancas do deck |
| **F6** | **Turno 4** — entra **um turno antes** do Otharri (CMC 5) e já está em campo quando o motor liga |
| **F7** | (i) o gatilho exige **ela atacando**, expondo um corpo 3/3 a bloqueio e remoção; (ii) **compete com o Otharri pelo slot de 4–5 manas** numa curva que já trava ali; (iii) a cláusula de untap perde valor à medida que o deck troca criaturas brancas por fichas vermelhas; (iv) **a anti-sinergia com `Dawnstrike Vanguard` volta se ela ficar** — Vanguard quer criaturas **viradas** no end step e a Tori destapa as brancas atacantes (as fichas Rebel continuam viradas, então a condição `2+ tapped` segue satisfeita pelas fichas; o atrito existe mas é **parcial**, não fatal) |

### Veredito

**Tori fica no pool e é uma boa carta do 99 — mas não é obrigatória, e o motivo pelo qual ela
sobe é diferente do motivo pelo qual ela estava no comando.**

- Ela **não** destrava "um segundo ataque parcial" de forma relevante: a cláusula de untap atinge
  só as **criaturas brancas atacantes**, e o board que o Otharri constrói é de **Rebels vermelhos**.
  O untap **não** gera combate extra, **não** remove de combate e **não** permite atacar de novo —
  ele deixa corpos brancos de pé para **bloquear**. Como fonte de segundo ataque, a resposta certa
  é `Combat Celebrant`, `Great Train Heist` ou `Éomer, Marshal of Rohan` (§4.4), não a Tori.
- O que **de fato** a valoriza sob Otharri é a **segunda** cláusula, que o passe anterior chamou
  de "trample na cor errada": as fichas são **vermelhas**. A Tori é, hoje, **a fonte de trample
  em massa do deck para o enxame** — e acumula +1/+1 no mesmo gatilho.
- Ela tem **3 pontos de sinergia** (pump do enxame · trample do enxame · corpo Human Knight
  legendário que alimenta Circle of Loyalty/Inspiring Veteran) e passa folgadamente na regra 3.
- **Atrito real a pesar:** CMC 4 numa curva congestionada, e `Crash Through` (1 mana, cantrip,
  trample a **todas**, inclusive brancas) faz a parte mais importante por 1/4 do custo.

**Encaminhamento:** **não cortar por impressão.** Tori entra na lista de candidatas com prioridade
**média-alta**, competindo pelo slot com `Adeline` (3 manas, gera 3 corpos por ataque) e
`Hero of Bladehold` (4 manas, 2 corpos + battle cry). Se o deck final ficar com **≤2 fontes de
trample/menace em massa**, ela sobe para prioridade alta. Decisão do usuário via orquestrador
(regra 9).

**Registro para o `decisions.md` (regra 5):** a Tori saiu da **zona de comando** em 2026-09-19 por
não gerar carta/corpo/mana/fechamento. **O que mudou para ela como carta do 99:** o board passou a
ser composto de **criaturas vermelhas**, que é exatamente o público-alvo da cláusula de trample
dela — cláusula que era letra morta quando 78% do deck era branco. Motivo original do corte
(não sustenta a zona de comando) **continua válido**; a função nova (trample + pump em massa
num enxame vermelho) **não existia antes**.

---

## 7. Lacunas remanescentes — o que o eixo novo **não** resolve

| Dor | Situação sob Otharri | Ainda falta |
|---|---|---|
| **1 · Não fecha o jogo** | **Resolvido pelo comandante.** 52 de dano a um jogador em 4 ataques sem anthem; 76 com um. `Warleader's Call`/`Impact Tremors` adicionam dano que **ignora bloqueio** | nada crítico — falta só escolher 2–3 multiplicadores |
| **2 · Mão morta** | **Contornado, não resolvido.** `Belladonna Took` (↑↑) e `Neyali` dão saque ligado ao motor | o gap de **−9 em draw** continua; `draw-specialist` |
| **3 · Sem respostas** | **Inalterado.** 4 remoções reais + 4 cartas art/ench-only + **0 wipes** | `interaction-specialist` — a caixa tem `Fumigate`, `Glass Casket`, `Blast Zone` |
| **4 · Morre para board wipe** | **Parcialmente resolvido de graça.** Contadores de experiência ficam no jogador; recursão de 4 manas sem imposto | **achado crítico:** a proteção em massa atual (`Paladin Danse`) **não cobre nem o enxame Rebel nem o comandante Phoenix**. `Duty Beyond Death` é a única que cobre — e é 1 carta. `interaction-specialist` |
| **Ramp** | **piorou de importância** | Otharri custa 5 e quer entrar cedo; `ramp-specialist` (gap −9/−3, caixa tem 5–7 peças) |
| **Terrenos** | −2 na base de 38 | `Karn's Bastion` e `Blast Zone` (**caixa**) resolvem **e** somam função |
