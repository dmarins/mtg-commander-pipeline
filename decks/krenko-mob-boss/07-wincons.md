# Condições de Vitória e Testes — Krenko, Mob Boss (v3)

**Fase 7 · modo `improve` · 2026-08-22.** Todo texto oracle desta análise foi puxado nesta sessão
com `bin/mtgdb deck krenko-mob-boss -full` e `bin/mtgdb oracle`. Rulings via `bin/mtgdb rulings`.
Base analisada: `deck.md` **depois** das 26 trocas.

> **Achado que precede tudo: o deck tem 99 cartas, não 100.**
> `bin/mtgdb deck krenko-mob-boss` lê 69 entradas — 63 não-terrenos + Krenko + 5 entradas de terreno
> (35 terrenos). Soma: **1 + 63 + 35 = 99**. `Chandra, Torch of Defiance` foi **aprovada** na fase 5
> (`05-interaction.md` linha 31) e **registrada como entrada** em `decisions.md` linha 54, mas
> **não existe seção de Planeswalkers em `deck.md`** e ela não aparece em nenhuma outra.
> **Há um slot livre. A primeira recomendação desta fase não custa corte nenhum.**

---

## 1. Caminhos de vitória atuais

Legenda de alcance: **mesa** = mata os 3 oponentes na mesma janela · **1 opp** = mata um por vez.

| # | Caminho | Cartas envolvidas | Turno estimado (mesa de 4, 40 vidas) | Alcance | Passa por combate? | Consistência |
|---|---|---|---|---|---|---|
| **1** | **Dreno por ETB** — cada token que entra pinga *cada* oponente | Krenko + `Impact Tremors` e/ou `Purphoros, God of the Forge` (+ `Goblin Chieftain`/`Warchief`/`Hammer of Purphoros` p/ haste) | **T6–T7** com as duas · **T7** só Purphoros · **T8** só Impact Tremors | **mesa** | **NÃO** | média — 2 cartas em 99, sem tutor que as ache (~19–23% de ver ao menos uma até T6) |
| **2** | **Dreno por morte** — sacrifica o enxame na cara | `Goblin Bombardment` + `Boggart Shenanigans` (+ `Rundvelt Hordemaster` p/ cartas) | **T7–T8** para o 1º oponente | **1 opp** | **NÃO** | média — Bombardment é o gargalo; `Goblin Matron` **acha Boggart Shenanigans** (é card Goblin) |
| **3** | **Combo Umbral Mantle** (plano B aprovado pelo usuário) | `Umbral Mantle` + `Skirk Prospector` + Krenko **com haste** + 4+ Goblins + qualquer 1 de {`Impact Tremors`, `Purphoros`, `Goblin Bombardment`} | **T5–T7**, fecha no turno em que liga | **mesa** (com Tremors/Purphoros) · **1 opp** (com Bombardment) | **NÃO** | baixa isolada (~1–2% das 2 peças de graça), **mas tutorável**: `Goblin Matron`→Prospector, `Goblin Engineer`→Umbral Mantle |
| **4** | **Alfa-strike quadrático** | `Shared Animosity` + `Ferocity of the Wilds` + 10–12 Goblins atacando com haste | **T6–T7** | **mesa** (ver §2.5 — o bônus **não** depende de quem é atacado) | sim | boa — 2 encantamentos + o enxame que Krenko já faz |
| **5** | **Alfa-strike com double strike** | `Assault on Osgiliath` (X≥2) + 8+ Goblins + lordes | fecha **no turno em que resolve**, a partir de T6 | 1–2 opp | sim | média |
| **6** | **Reach / últimos 20** — não fecha sozinho | `Hobgoblin Bandit Lord`, `Massive Raid`, `Lightning Volley`, `Guttersnipe`, `Chandra` (se entrar), `Fanatical Firebrand`, `Ember Hauler`, `Mudbutton Torchrunner`, `Den of the Bugbear` | somam 15–30 de dano dirigível a partir do T6 | 1 opp por vez | **NÃO** | alta (9 peças) — mas é **módulo de acabamento**, não caminho |

### Resposta direta à queixa nº 1

**Caminhos que fecham SEM depender de combate: 3 (nº 1, 2 e 3).** Era **0** na lista antiga
(`02-theme.md` §4.5 listava 7 caminhos, e os 2 que fechavam de verdade — combate em massa e
alfa-strike — passavam ambos pelo combate; o único candidato não-combate, `Boggart Shenanigans`,
estava marcado **CAMINHO MORTO** por falta de sac outlet).

E **dois deles matam a mesa inteira ao mesmo tempo** (nº 1 e nº 3), porque `Impact Tremors` e
`Purphoros` causam dano **a cada oponente** sem alvo — bloqueador, proteção de criatura e
"escolha um oponente" não interferem. É matemática de 40 vidas ÷ 3, não de 120 de dano.

---

## 2. A matemática dos gatilhos — o que muda em relação ao combate

### 2.1 A conta base

Krenko taps: cria **X** tokens, X = nº de Goblins que você controla, **incluindo o próprio Krenko**
(ruling 2012-07-01). Os X tokens entram **simultaneamente**.

| Em campo | Dano **a cada oponente** por ativação com **G** Goblins | Dano total na mesa (3 opp) |
|---|---|---|
| `Impact Tremors` só | **G** | 3G |
| `Purphoros` só | **2G** | 6G |
| **As duas** | **3G** | **9G** |

E o nº de Goblins **dobra** a cada ativação (G → 2G).

### 2.2 Tabela de clock — dano acumulado por oponente

| Ativações | G inicial | `Impact Tremors` só | `Purphoros` só | **As duas** |
|---|---|---|---|---|
| 1ª / 2ª / 3ª | **3** | 3 / 9 / 21 | 6 / 18 / 42 → **morre na 3ª** | 9 / 27 / **63 → morre na 3ª** |
| 1ª / 2ª / 3ª | **4** | 4 / 12 / 28 (4ª: 60) | 8 / 24 / **56 → 3ª** | 12 / 36 / **84 → 3ª** |
| 1ª / 2ª / 3ª | **6** | 6 / 18 / **42 → 3ª** | 12 / 36 / **84 → 3ª** | 18 / **54 → 2ª** |
| 1ª / 2ª | **8** | 8 / 24 (3ª: 56) | 16 / **48 → 2ª** | 24 / **72 → 2ª** |

Traduzindo para turnos, na linha mediana (Krenko no T4, haste disponível, 4 Goblins):
**ativações no T4/T5/T6 → mesa morta no T6–T7.** Com `Sol Ring` ou `Urza's Incubator` a linha
adianta um turno inteiro (Krenko no T3 → **T5–T6**).

### 2.3 O atalho que fecha um turno antes

Depois da 2ª ativação com as duas encantamentos você está em **36 por oponente** e com **16 Goblins**.
Faltam 4 em cada. `Goblin Bombardment` sacrificando 12 tokens (4 apontados em cada cara) **fecha ali**.
Com `Boggart Shenanigans` em campo são **2 de dano por Goblin sacrificado**, e bastam 6 sacrifícios.
→ **T6 é um turno letal realista quando o pacote completo está montado.**

### 2.4 Onde `Goblin Bombardment` / `Boggart Shenanigans` são MATEMÁTICA DIFERENTE

Os dois causam dano a **um alvo por gatilho** ("any target" / "target player or planeswalker"),
não a cada oponente. Sacrificar N Goblins com os dois em campo = **2N de dano distribuível**,
não 2N *em cada* oponente. Para matar 3 oponentes de 40 só por essa rota são **120 de dano** →
~60 Goblins. É rota de **um oponente por vez** (40 de dano ≈ 20 sacrifícios com Shenanigans).
**Não confunda os dois números no goldfishing.**

### 2.5 `Shared Animosity` — o achado que muda a avaliação do caminho de combate

O oracle é: *"Whenever a creature you control attacks, it gets +1/+0 until end of turn for each
**other attacking creature** that shares a creature type with it."* **Não há cláusula sobre quem é o
jogador defensor.** O ruling 2008-04-01 confirma que a habilidade conta **criaturas, não tipos**.

Consequência: você pode **dividir o ataque entre os 3 oponentes** e cada Goblin ainda recebe o bônus
cheio. Com **A** Goblins atacando, cada um vira `(base + A − 1)`:

| A atacantes | Cada Goblin (1/1 base) | Total | Dividido 4/4/4 |
|---|---|---|---|
| 8 | 8/1 | 64 | 24 / 24 / 16 |
| 10 | 10/1 | 100 | 40 / 30 / 30 |
| **12** | **12/1** | **144** | **48 / 48 / 48 → mesa morta** |

Com `Goblin Chieftain` (base 2) e `Ferocity of the Wilds` (+1/+0 e **trample**, e 100% das criaturas
do deck são não-Humanas) o limiar cai para **~10 atacantes**. Krenko sai de 3 Goblins para 12 em
duas ativações. **O caminho nº 4 é um kill de mesa, não de um oponente** — a §4.5 antiga o
classificava como "letal em 1 oponente por volta do T9–T11".

---

## 3. Verificação de regras — o que está confirmado e o que NÃO está

| # | Pergunta | Veredito | Fonte |
|---|---|---|---|
| **R1** | Krenko cria N tokens de uma vez — `Impact Tremors` dispara N vezes? | **SIM, confirmado.** *"If multiple creatures you control enter at the same time, Impact Tremors will trigger once for each of those creatures."* | ruling `Impact Tremors` 2024-11-08 |
| **R2** | E `Purphoros`? | **SIM, mas SEM ruling específico da carta.** Os 12 rulings de Purphoros tratam só de devoção e de God/criatura. O gatilho tem redação idêntica (*"Whenever another creature you control enters"*), então a regra do R1 se aplica pelo mesmo mecanismo. **Registro que não há ruling na carta.** | oracle + ruling R1 por analogia de template |
| **R3** | Ordem dos 2N gatilhos | **Você escolhe a ordem** (são todos seus, vão para a pilha juntos). **Não muda nada:** nenhum gatilho depende do resultado do outro. Se um oponente chega a 0 no meio da pilha, ele sai do jogo por ação baseada em estado e os gatilhos restantes que o teriam atingido simplesmente não fazem nada. **O que importa é a ordem em relação às ATIVAÇÕES:** `Hobgoblin Bandit Lord` conta *"Goblins that entered the battlefield this turn"* — **ative-o DEPOIS de Krenko**. E o ruling 2021-07-23 confirma: *"It doesn't matter if those Goblins are still on the battlefield as it resolves"* → você pode sacrificar tudo antes e o dano do Bandit Lord continua cheio. | ruling `Hobgoblin Bandit Lord` 2021-07-23 |
| **R4** | `Boggart Shenanigans` dispara por **cada** token sacrificado ao `Goblin Bombardment`? | **SIM — mas SEM ruling na carta.** Os 3 rulings de Boggart Shenanigans só tratam de *kindred* (ex-*tribal*). A base é: (a) ficha sacrificada **vai ao cemitério** antes de deixar de existir, satisfazendo *"put into a graveyard from the battlefield"*; (b) é um gatilho por evento. **Cada gatilho é opcional (*"you may"*) e exige um ALVO** (`target player or planeswalker`) — 12 sacrifícios = 12 gatilhos, 12 alvos escolhidos um a um. `Goblin Bombardment` **não tem ruling nenhum**. | rulings `Boggart Shenanigans` 2024-06-07 (só kindred); `Goblin Bombardment` = 0 rulings |
| **R5** | `Rundvelt Hordemaster` dispara por cada Goblin sacrificado? | **SIM.** Ruling: *"If Rundvelt Hordemaster and one or more other Goblins you control die at the same time, its ability will trigger once for each of them."* Sacrificar é morrer. **Cada gatilho exila o topo da sua biblioteca** — ver Risco 4. | ruling `Rundvelt Hordemaster` 2022-09-09 |
| **R6** | `Goblin Chirurgeon` (regenerar) vs board wipe | **NÃO há ruling da carta sobre wipes** (os 2 rulings dizem só que ela pode se sacrificar e se regenerar). A base é o **texto lembrete impresso de "Regenerate"**: *"The next time this creature would be **destroyed** this turn, instead **tap it**, **remove it from combat**, and heal all damage on it."* → **funciona** contra `Wrath of God`, `Damnation`, `Blasphemous Act`, `Chain Reaction` (destruição, inclusive por dano letal). → **NÃO funciona** contra **exílio** (`Farewell`, `Merciless Eviction`), **−X/−X** (`Toxic Deluge`, `Languish` — resistência 0 vai ao cemitério por ação baseada em estado, que **não é destruição**), **sacrifício forçado** (edicts) nem **bounce** (`Cyclonic Rift`). **Dois efeitos colaterais que decidem jogo:** regenerar **VIRA a criatura** (regenerar Krenko antes de ativá-lo custa a ativação do turno) e **a remove do combate**. | texto lembrete oficial + rulings `Goblin Chirurgeon` 2004/2007 |
| **R7** | `Umbral Mantle` no combo — Krenko precisa de haste? | **SIM, confirmado por ruling.** *"If a creature with an {Q} ability hasn't been under your control since your most recent turn began, you can't activate that ability, unless the creature has haste."* E: *"If the permanent is already untapped, you can't activate its {Q} ability"* → a sequência obrigatória é **tapar Krenko pela habilidade dele PRIMEIRO**, depois pagar `{3}`+`{Q}`. O untap em si **não pode ser respondido**. | 3 rulings `Umbral Mantle` 2008-05-01 |
| **R8** | `Goblin Matron` acha `Boggart Shenanigans`? | **SIM.** Matron busca *"a Goblin card"*, não *"a Goblin creature card"*. `Boggart Shenanigans` é **Kindred Enchantment — Goblin**, e o ruling confirma: *"Kindred is a card type that allows noncreature cards to have creature types."* Mesma lógica: `Conspicuous Snoop` pode **lançá-la do topo da biblioteca** (*"You may cast Goblin spells from the top"*). | ruling `Boggart Shenanigans` 2024-06-07 |
| **R9** | `Urza's Incubator` desconta `Boggart Shenanigans`? | **NÃO** — desconta *"Creature spells of the chosen type"*, e ela é encantamento. **Mas `Goblin Warchief` SIM** (*"Goblin spells you cast cost {1} less"* — sem restrição a criatura). Assimetria fácil de errar na mesa. | oracle das duas + R8 |
| **R10** | `Assault on Osgiliath` amassa Orc — o token conta pro X de Krenko? | **NÃO** se você não tiver Army; o ruling confirma que o token criado é um **0/0 black Orc Army**. Mas **se você já tiver o Army Goblin de `Goblin-town Flunkies`**, os contadores vão nele e ele vira **Orc Goblin Army** — continua Goblin, continua contando. E o 0/0 Orc entrando **dispara Impact Tremors e Purphoros** de qualquer forma. | 9 rulings `Assault on Osgiliath` 2023-06-16 |

---

## 4. Combos

| Combo | Peças | Resultado | Confirmado por rulings? |
|---|---|---|---|
| **Umbral Mantle / Krenko** (aprovado pelo usuário como plano B) | `Umbral Mantle` + `Skirk Prospector` + Krenko **com haste** + **≥4 Goblins** | Loop: tapa Krenko → G tokens (G→2G) → sacrifica 3 Goblins ao Prospector por `{R}{R}{R}` → paga `{3}`,`{Q}` e desvira Krenko. Crescimento `G → 2G−3`, **positivo para G > 3** (confere com o "~4 Goblins" do usuário). Sequência real a partir de 4: 4→5→7→11→19→35→67 | **Sim, na parte que importa**: os 3 rulings de `Umbral Mantle` confirmam (a) haste/enjoo de invocação se aplica ao `{Q}`, (b) a criatura precisa estar **tapada** para pagar o `{Q}`, (c) o untap não pode ser respondido |
| **Saída A do combo** | + `Impact Tremors` e/ou `Purphoros` | **Mata a mesa.** Com as duas e G=4: 12 → 15 (27) → 21 (48). **3 iterações e acabou** — só 9 cartas exiladas se Rundvelt estiver em campo | Sim (R1); R2 sem ruling na carta |
| **Saída B do combo** | + `Goblin Bombardment` | Dano arbitrário, **um alvo por sacrifício** — mata a mesa, mas exige ~120 sacrifícios | Sem rulings (`Goblin Bombardment` = 0 rulings) |
| **Saída C do combo** | sem nenhuma das acima | Só um enxame enorme **com enjoo de invocação** — precisa de `Goblin Chieftain`/`Warchief`/`Hammer of Purphoros` para atacar no mesmo turno. **É a saída ruim: não fecha** | — |
| **Motor de artefato** (não é combo infinito) | `Hammer of Purphoros` → Golem 3/3 → sacrifica ao `Krenko, Baron of Tin Street` | Golem entrando dispara Tremors/Purphoros; sacrificá-lo dá +1/+1 em cada Goblin **e** dispara a 2ª habilidade do Baron (`{R}`: mais um Goblin com haste → mais um gatilho). Custa um terreno por ciclo | Sem rulings (`Hammer of Purphoros` = 0 rulings) |

> **Regra de operação do combo:** o loop **não é obrigatório** — você conta as iterações e para.
> Isso importa por causa do Risco 4 abaixo.

---

## 5. Finishers recomendados

Coleção consultada primeiro (regra 7): `bin/mtgdb collection` confirma que **nenhuma** peça do
pacote de wincon atual está na coleção; das 34 elegíveis de `00b-colecao-elegivel.md`, uma serve.

| Carta | CMC | Como fecha o jogo | Sinergias (mín. 2) | Na coleção? | Slot |
|---|---|---|---|---|---|
| **`Chandra, Torch of Defiance`** | 4 | **Já aprovada — só não foi transcrita.** `+1` = **2 de dano a CADA oponente** quando você não lança a carta exilada: é uma 4ª fonte de dreno sem combate. `+1` = `{R}{R}` alimenta `Umbral Mantle`, `Purphoros`, `Castle Embereth`. Atravessa board wipe | (1) 4º dreno "each opponent"; (2) mana sink para o combo; (3) permanente que sobrevive a Wrath; (4) `−3` = remoção | não | **slot livre (99→100) — CUSTO ZERO** |
| **`Kuldotha Rebirth`** | 1 | **Três Goblins por `{R}`.** Com Tremors+Purphoros = **9 de dano a cada oponente por 1 mana**. Com Purphoros só = 6. É o melhor dano-por-mana do deck inteiro | (1) 3 corpos que entram no X de Krenko no mesmo turno; (2) o artefato sacrificado **dispara `Krenko, Baron of Tin Street`** (mais um Goblin com haste); (3) 3 contadores em `Quest for the Goblin Lord`; (4) alvos de sacrifício já existem: Golem do `Hammer`, Treasure do `Reckless Lackey`, `Arcane Signet` gasto | **SIM — `tem 1`, R$ 0** | precisa de corte → **condicionado** |
| **`General Kreat, the Boltbringer`** | 3 | **Terceiro `Impact Tremors`, num corpo Goblin.** *"Whenever another creature you control enters, deals 1 damage to each opponent."* Sobe a coluna "as duas" da §2.2 em +G por ativação. Tags do Scryfall Tagger: `impact-effect`, `group-slug` — as mesmas de `Impact Tremors` | (1) redundância do caminho nº 1 (hoje 2 cartas em 99); (2) **é Goblin** → conta no X, recebe os 3 lordes, `Shared Animosity`, `Ferocity`; (3) **`Goblin Matron` acha** e **`Conspicuous Snoop` lança do topo**; (4) `Urza's Incubator` o reduz para `{R}`; (5) faz um token por ataque de Goblins | não · **sem cotação registrada** (`mtgdb prices`) | CMC 3 já em 23 → **condicionado** |
| **`Goblin Sharpshooter`** | 3 | Dano repetível **sem combate**: `{T}`: 1 dano em qualquer alvo, e **desvira sempre que uma criatura morre** — inclusive as suas. Com `Goblin Bombardment` cada Goblin sacrificado passa a valer **2 de dano** sem gastar `Boggart Shenanigans` | (1) dobra a saída B do combo; (2) é Goblin (X de Krenko, lordes, Matron, Snoop, Incubator); (3) varre x/1 dos oponentes | não · **sem cotação registrada** | CMC 3 → **condicionado** |

### Cortes condicionados (fora da minha alçada — devolvo ao orquestrador)

Fichas F1–F7 completas, com as **mesmas condições** assumidas para as entradas (lordes em campo,
enxame montado). Nenhuma dessas eu corto sozinho: as três exercem função fora da minha lente.

**`Mudbutton Clanger`** `{R}` 1/1 Goblin Warrior
- **F1** — Kinship: no seu upkeep, olha o topo; se compartilhar tipo de criatura, revela e ganha +1/+1 **até o fim do turno**. ~33% dos upkeeps (32 criaturas Goblin + `Boggart Shenanigans` em 99). O "olhar o topo" é informação grátis e **não** filtra nada.
- **F2** — corpo 1/1; combustível de `Goblin Bombardment` / `Skirk Prospector` / `Goblin Chirurgeon`; bloqueia.
- **F3** — **é Goblin**: entra no X de Krenko, em `Shared Animosity`, em `Massive Raid`.
- **F4** — recebe os 3 lordes (**4/4 com os três**), `Quest for the Goblin Lord` (+2/+0), `Ferocity` (+1/+0 e trample), `Skullclamp`, `Krenko Baron`.
- **F5** — não dá nada a ninguém.
- **F6** — **T1**: já está no X do Krenko do T4. Esse é o argumento real a favor dela.
- **F7** — nenhum.
- **Cobertura se sair:** corpo Goblin T1 → `Skirk Prospector`, `Fanatical Firebrand`, `Tin Street Dodger`, `Raging Goblin`, `Mogg Sentry`, `Goblin Chirurgeon` (6 outros 1-drops Goblin). **Nada fica descoberto.** É a única das 30 criaturas cujo F1 não produz efeito relevante nenhum.

**`Raging Goblin`** `{R}` 1/1 Goblin Berserker
- **F1** — só haste. **F2** — corpo 1/1, sac fodder, bloqueia. **F3** — Goblin (X, Shared Animosity, Massive Raid). **F4** — recebe tudo (4/4 com os 3 lordes). **F5** — nada. **F6** — T1, e a haste faz dele um atacante imediato. **F7** — nenhum.
- **Cobertura:** idêntica à do Clanger, e a haste dele fica **redundante** com `Goblin Chieftain`, `Goblin Warchief` e `Hammer of Purphoros` (3 fontes de haste coletiva entraram nesta rodada). **Nada descoberto** — mas é o corte de menor perda depois do Clanger.

**`Goblin Surprise`** `{2}{R}` Instant
- **F1** — modal: +2/+0 em todas as criaturas **OU** dois Goblins 1/1. **F2** — sem corpo. **F3** — **é Instant**: dispara `Guttersnipe` (2 a cada oponente) e conta para `Blitz of the Thunder-Raptor`. **F4** — nada. **F5** — pump coletivo instantâneo (trick de combate). **F6** — CMC 3, o slot com **23 cartas**. **F7** — **atrito real**: compete no turno 3 com `Goblin Chieftain`, `Goblin Warchief`, `Shared Animosity`, `Hordeling Outburst` e `Ferocity of the Wilds`, todos mais impactantes.
- **Cobertura:** dois tokens em instant → `Krenko's Command` (2 tokens, CMC 2) e `Hordeling Outburst` (3, CMC 3); pump coletivo → `Purphoros` `{2}{R}`, `Castle Embereth`, `Goblin War Party` (entwine). **Função descoberta: nenhuma. Custo real: perde-se a flexibilidade modal em instant** — declarado.

> **Não proponho corte de:** `Skullcrack` (é **proteção do kill de dreno** — *"Players can't gain life this turn.
> Damage can't be prevented this turn"* desliga fog e lifegain no turno em que você aponta 40+ em cada cara,
> e ainda soma 3), `Burn, Burn, Tree and Fern` (**única resposta a artefato colorido** do deck — `Goblin
> Cratermaker` só destrói permanente **incolor**), `Mogg Sentry` (+2/+2 por spell de oponente = bloqueador
> real numa mesa de 4) e `Zada, Hedron Grinder` (**decisão explícita do usuário** em `decisions.md`).

---

## 6. Protocolo de goldfishing

### 6.0 Preparação

- **8 partidas** (o mínimo do pipeline é 5; 8 porque há **3 caminhos distintos** para observar).
- Jogue **na jogada** (sem compra no T1) em 4 partidas e **na resposta** (compra no T1) nas outras 4.
- **Sem oponente.** Assuma 3 oponentes de 40 vidas que **não interagem**, com **uma exceção obrigatória**: o teste de wipe da §6.4.
- Deixe a mesa montada à vista: você vai precisar contar Goblins todo turno.
- Se `Chandra, Torch of Defiance` ainda não estiver na lista física, jogue com 99 e anote isso.

### 6.1 Critérios de mulligan **deste deck**

Mulligan de Londres (compre 7, devolva N ao fundo). Anote **quantos mulligans**.

**MANTENHA** se a mão tiver as três coisas:
1. **2 a 4 terrenos** (nunca 5+, nunca 1);
2. **pelo menos um Goblin de CMC 1–2** (o corpo do T1–T2 é o que dobra no X do Krenko do T4);
3. **pelo menos uma** de: aceleração (`Sol Ring`, `Arcane Signet`, `Urza's Incubator`), habilitador de haste (`Goblin Chieftain`, `Goblin Warchief`, `Hammer of Purphoros`, `Goro-Goro`, `Lightning Greaves`, `Swiftfoot Boots`) **ou** payoff de dreno (`Impact Tremors`, `Purphoros`, `Goblin Bombardment`, `Boggart Shenanigans`, `Hobgoblin Bandit Lord`).

**MANTENHA SEMPRE:** qualquer mão com **`Sol Ring` ou `Urza's Incubator` + 2 terrenos**.

**MULLIGAN** se:
- 0–1 terreno, ou 5+ terrenos;
- **nenhum** Goblin abaixo de CMC 3 (a mão só age no T3+);
- **3 ou mais cartas de CMC 3** e nenhuma aceleração — este é o teste direto da congestão de CMC 3 (23 cartas). **Anote toda vez que mulligar por esse motivo**;
- só Krenko de peça relevante, sem habilitador de haste e sem corpo barato. **Krenko sozinho no T4 não é mão boa.**

### 6.2 O que registrar — tabela de partida (preencha à mão)

```
PARTIDA #___    ( ) na jogada  ( ) na resposta        mulligans: ___
mulligou por congestão de CMC 3? ( ) sim ( ) não

terrenos na mão mantida: ___        1º land drop perdido: turno ___ (ou "nenhum")

Krenko lançado no turno: ___    tinha haste? ( ) sim ( ) não   fonte: ______________
1ª ATIVAÇÃO de Krenko: turno ___   X (Goblins em campo) = ___
payoffs em campo nessa hora: [ ]Impact Tremors [ ]Purphoros [ ]Goblin Bombardment
                             [ ]Boggart Shenanigans [ ]Hobgoblin Bandit Lord
                             [ ]Shared Animosity [ ]Ferocity of the Wilds  [ ]nenhum

1º dano relevante (≥5 num oponente): turno ___   por qual carta: ______________

┌──────┬─────────┬───────────┬──────────────┬───────────┬────────────────────┐
│ Turno│ terrenos│ Goblins   │ dano ACUM.   │ cartas na │ CMC-3 na mão que   │
│      │ em jogo │ em campo  │ no opp #1    │ mão (fim) │ não pude lançar    │
├──────┼─────────┼───────────┼──────────────┼───────────┼────────────────────┤
│  T4  │         │           │              │           │                    │
│  T5  │         │           │              │           │                    │
│  T6  │         │           │              │           │                    │
│  T7  │         │           │              │           │                    │
│  T8  │         │           │              │           │                    │
└──────┴─────────┴───────────┴──────────────┴───────────┴────────────────────┘

LETAL no 1º oponente: turno ___   caminho: ( )dreno ETB ( )Bombardment
                                            ( )combo Mantle ( )combate ( )não chegou
LETAL na MESA INTEIRA: turno ___  caminho: ______________  ( )não chegou

turnos com 0 cartas na mão: ___     mana sobrando sem uso (turnos): ___
```

### 6.3 Como contar o "letal projetado"

- **Não** assuma bloqueadores nos caminhos 1, 2, 3 e 6 — eles **não passam pelo combate**, é dano direto.
- Nos caminhos 4 e 5 (combate), assuma **1 bloqueador 2/2 por oponente** e lembre que `Ferocity of the Wilds` dá **trample** — o excesso passa.
- Lembre da §2.5: com `Shared Animosity` você **divide o ataque entre os 3 oponentes sem perder o bônus**.
- Conte os gatilhos **um a um** na primeira partida (para pegar a mão do cálculo); depois use a tabela da §2.2.

### 6.4 Teste de board wipe — **obrigatório, 1 por partida**

**No turno seguinte à 1ª ativação de Krenko**, destrua **todas as suas criaturas** (simule um `Wrath of God`)
e ponha Krenko no command zone. Antes de fazer isso, responda no papel:

- `Goblin Chirurgeon` estava em campo? Quantos Goblins você poderia sacrificar para regenerar quantos? (**vale contra destruição — ver R6**)
- `Deflecting Swat` estava na mão? (custo **zero** com o comandante em campo)
- `Lightning Greaves` estava **em** Krenko? → **atenção ao Risco 2**: com Greaves, `Goblin Chirurgeon` **não consegue mirar** Krenko.

Depois anote:
```
permanentes que SOBREVIVERAM: ______________________________________
turnos até voltar a ameaçar (Goblins ≥ 6): ___
carta que reconstruiu: ( )Krenko recastado ( )Den of the Bugbear ( )Outpost Siege
                       ( )Chandra ( )Goblin Matron ( )Squee ( )Rundvelt (impulso)
                       ( )Hordeling Outburst/Krenko's Command ( )nenhuma
```

### 6.5 As perguntas que os dados precisam responder

| Queixa | Métrica | Meta | Se falhar, o diagnóstico é |
|---|---|---|---|
| **"Não fecho o jogo"** | Letal no 1º oponente **≤ T7** | **≥ 5 de 8 partidas** | payoff de dreno não aparece cedo: subir a densidade de `impact-effect` (hoje **2** cartas em 99) |
| **"Não fecho o jogo"** | Letal na **mesa inteira ≤ T8** | **≥ 4 de 8** | só o caminho de 1 oponente está ligando → falta redundância de "each opponent" |
| **"Não fecho o jogo"** | **Qual caminho** fechou, por partida | os 3 sem-combate aparecerem em **≥ 4 de 8** somados | se 8 de 8 fecharam por combate, a queixa nº 1 **não foi resolvida** |
| **"Não fecho o jogo"** | 1ª ativação de Krenko **≤ T5** com X ≥ 4 | **≥ 6 de 8** | haste ou aceleração insuficientes (6 habilitadores hoje; eram 2) |
| **"Fico sem cartas"** | Turnos com **0 cartas na mão** antes do T7 | **≤ 1 turno** em ≥ 6 de 8 | draw ainda curto **apesar** dos 6 motores novos; anote **qual** motor estava online |
| **"Fico sem cartas"** | `Skullclamp` gerou compra? | ≥ 1 vez em ≥ 4 de 8 | ver **Risco 3** — com lorde em campo, Skullclamp exige sac outlet |
| **"Apanho de wipe"** | Turnos até voltar a ameaçar (§6.4) | **≤ 2 turnos** em ≥ 5 de 8 | se ≥ 3 turnos: resiliência ainda insuficiente, apesar dos ~19 permanentes que sobrevivem |
| **"Mana travando"** | Land drops perdidos | **0** em ≥ 6 de 8 | 35 terrenos pode estar baixo → devolver ao `manabase-engineer` |
| **"Mana travando"** | Cartas de **CMC 3** presas na mão | **≤ 1 por turno** em ≥ 6 de 8 | confirma a congestão de 23 cartas no CMC 3 → cortar CMC 3 para CMC 1–2 |
| **"Mana travando"** | Mana sobrando sem uso | **≤ 2 turnos** | flood: os sinks (`Umbral Mantle`, `Purphoros`, `Hammer`, `Castle Embereth`, `Den of the Bugbear`) não estão aparecendo |

---

## 7. Riscos — além dos três já conhecidos

> Já sabidos e **não** repetidos aqui: CMC 3 em 23 · lordes em 3 (meta 4–5) · resposta a encantamento em 0.

### Risco 1 — **O deck tem 99 cartas** (crítico, mecânico)
`Chandra, Torch of Defiance` foi aprovada e registrada mas nunca transcrita para `deck.md`; não existe
seção de Planeswalkers no arquivo. Confirmado por `bin/mtgdb deck krenko-mob-boss` (63 não-terrenos +
Krenko + 35 terrenos = 99). **Efeito colateral bom:** há um slot de graça, e a primeira recomendação
desta fase não custa corte.

### Risco 2 — **`Lightning Greaves` dá SHROUD, e isso desliga três cartas suas** (anti-sinergia criada nesta rodada)
Shroud impede **qualquer** alvo, inclusive seu. Com Greaves em Krenko:
- **você não consegue equipar `Umbral Mantle` nele** (Equip mira) → **a peça-chave do plano B fica trancada**;
- **`Goblin Chirurgeon` não consegue regenerá-lo** (*"Regenerate target creature"*) → a proteção repetível não alcança o alvo mais importante;
- `Kick in the Door` e `Fists of Flame` também não o alcançam.
**Não é motivo para cortar Greaves** — o Equip `{0}` no turno em que Krenko entra é exatamente o
ponto, e duas peças de proteção valem mais que uma. Mas a **ordem de jogo é obrigatória**:
**`Umbral Mantle` primeiro, `Lightning Greaves` depois** (ativar a habilidade do Mantle **não** mira,
então continua funcionando com shroud). `Swiftfoot Boots` (hexproof) **não** tem esse problema — é a
peça certa quando o combo está na mesa. **Coloque isso no protocolo e observe no goldfishing.**

### Risco 3 — **`Skullclamp` e os 3 lordes entraram na MESMA rodada e brigam** (anti-sinergia criada nesta rodada)
`Skullclamp` dá **+1/−1**: num token 1/1 vanilla ele mata na hora e compra 2. Mas com **qualquer** um
dos três lordes novos em campo (`Goblin Chieftain`, `Rundvelt Hordemaster`, `Hobgoblin Bandit Lord`),
o token é **2/2** → vira **3/1 e NÃO morre**. Skullclamp deixa de ser motor automático.
**Não quebra:** o deck tem 3 sac outlets gratuitos (`Goblin Bombardment`, `Skirk Prospector`,
`Goblin Chirurgeon`) e sacrificar o token equipado ainda dispara *"Whenever equipped creature dies,
draw two cards"*. **Mas vira um passo extra**, e sem sac outlet em campo o Skullclamp fica inerte.
Registre no goldfishing (§6.5) — senão vai parecer que a carta é ruim quando o problema é sequência.

### Risco 4 — **`Rundvelt Hordemaster` transforma o combo num relógio de biblioteca**
Cada Goblin sacrificado ao `Skirk Prospector` **morre** → Rundvelt exila o topo (R5). O loop do
Umbral Mantle sacrifica **3 Goblins por iteração**: 4→5→7→11→19→35→67 são **7 iterações = 21 cartas
exiladas**. Com `Goblin Bombardment` como saída, muito mais. **O combo não é infinito de verdade —
é "arbitrariamente grande, limitado pela sua biblioteca" quando Rundvelt está em campo.**
Mitigação real: com `Impact Tremors` **+** `Purphoros` o kill sai em **3 iterações (9 cartas)**.
**Regra de mesa: conte as iterações e pare.** Não deixe o loop rodar "até acabar".

### Risco 5 — **Concentração em encantamentos: 8 dos 10 encantamentos são payoff**
`Impact Tremors`, `Purphoros`, `Goblin Bombardment`, `Boggart Shenanigans`, `Shared Animosity`,
`Ferocity of the Wilds`, `Quest for the Goblin Lord`, `Outpost Siege` — **os 3 caminhos sem combate
E o caminho de combate mais forte moram todos em encantamento**. Um `Farewell` que exila
encantamentos, um `Back to Nature` ou um `Aura of Silence` apaga o deck inteiro de uma vez, e o
deck tem **0 respostas a encantamento** para revidar. Isso é **diferente** da lacuna já conhecida:
não é só "não respondo encantamento", é **"toda a minha ameaça é encantamento"**. Vale considerar,
na próxima rodada, mover 1 payoff para um **corpo** (é exatamente o que `General Kreat, the
Boltbringer` e `Goblin Sharpshooter` fazem) e não só comprar `Chaos Warp`/`Vandalblast`.

### Risco 6 — **A ordem de sacrifício custa a ativação seguinte** (operacional)
`Skirk Prospector`, `Goblin Bombardment` e `Goblin Chirurgeon` comem Goblins, e o X de Krenko conta
Goblins **no momento da ativação**. **Regra fixa: ative Krenko PRIMEIRO, sacrifique depois.**
Exceção: `Hobgoblin Bandit Lord` conta *"Goblins que entraram este turno"* e **não se importa** com
sacrifícios posteriores (ruling 2021-07-23) — ele pode ser ativado por último sem perda.
Segunda armadilha: **regenerar VIRA a criatura** (R6) — regenerar Krenko antes de ativá-lo custa a
ativação do turno inteiro.

### Risco 7 — **Devoção do `Purphoros`: bom que ele NÃO seja criatura**
Com devoção a vermelho < 5, Purphoros **não é criatura** — logo não é atingido por `Wrath`, `−X/−X`,
nem por exílio de criatura, e não pode ser bloqueado nem atacado. **É o permanente mais resiliente
do deck.** Ele **nunca** conta no X de Krenko (é God, não Goblin), com ou sem devoção. O risco é o
oposto do intuitivo: quando a devoção **chega a 5** (fácil — Krenko 2 + Chieftain 2 + Warchief 2 +
Hobgoblin Bandit Lord 2 + Chainwhirler 3 + Hammer 2) ele vira um **6/5 criatura** e passa a ser
alvo legal de **edicts e −X/−X**, contra os quais indestructible não protege. **Anote a devoção no
goldfishing quando ele estiver em campo.**

### Risco 8 — **A densidade de payoff é baixa para o caminho principal**
Estimativa hipergeométrica (sem contar dig), cartas vistas até o turno T na jogada:

| Pacote | T4 | T5 | T6 | T7 |
|---|---|---|---|---|
| `Impact Tremors` **ou** `Purphoros` (2 cartas) | 19% | 21% | 23% | 25% |
| pacote de dreno sem combate (4 cartas) | 35% | 38% | 41% | 44% |
| + `Hobgoblin Bandit Lord` e `Massive Raid` (6) | 48% | 52% | 55% | 58% |
| as **duas** peças do combo, sem tutor | ~1% | ~1% | ~2% | ~2% |

O deck **compensa parcialmente** com dig real (`Reckless Impulse`, `Light Up the Stage`,
`Outpost Siege`, `Idol of Oblivion`, `Skullclamp`, `Rundvelt Hordemaster`, `Conspicuous Snoop`) e
com **tutores que o combo aceita**: `Goblin Matron` → `Skirk Prospector`; `Goblin Engineer` →
**`Umbral Mantle` é CMC 3**, então ele o busca para o cemitério e depois o **devolve ao campo**
(`{R}`,`{T}`, sacrifica um artefato: retorna artefato de CMC ≤ 3 do cemitério). Mesmo assim,
**é a densidade que o goldfishing precisa validar** — a métrica "qual caminho fechou" da §6.5 é
exatamente esse teste.

---

## 8. Ajustes pós-teste (preencher no modo `post-goldfish`)

| Sai | Entra | Motivo |
|---|---|---|
| — | `Chandra, Torch of Defiance` | **slot 100 vazio** — entrada já aprovada na fase 5, nunca transcrita. Não é corte, é correção |
| *(a definir)* | `Kuldotha Rebirth` | 9 de dano a **cada** oponente por `{R}` com o pacote completo; **coleção, R$ 0**. Candidatos condicionados: `Mudbutton Clanger`, `Raging Goblin`, `Goblin Surprise` (fichas na §5) |
| *(a definir)* | `General Kreat, the Boltbringer` | só se o goldfishing mostrar que o caminho nº 1 **não aparece** (métrica "qual caminho fechou" < 4 de 8). Sobe a densidade de `impact-effect` de 2 → 3 e move um payoff de encantamento para **corpo** (Risco 5) |

**Gatilhos de rediagnóstico:**
- letal > T7 na maioria → **densidade de payoff**, não curva: entra `Kuldotha Rebirth` / `General Kreat`;
- letal chega mas **sempre por combate** → a queixa nº 1 continua viva: entram as duas acima;
- land drop perdido em ≥ 3 de 8 → **devolver ao `manabase-engineer`** (35 pode estar baixo para um deck com 5 mana sinks);
- 2+ cartas de CMC 3 presas na mão em ≥ 5 de 8 → cortar CMC 3 para CMC 1–2 (a congestão já conhecida vira acionável);
- 0 cartas na mão antes do T7 em ≥ 4 de 8 → **devolver ao `draw-specialist`**;
- ≥ 3 turnos para se reerguer do wipe em ≥ 5 de 8 → resiliência insuficiente apesar dos ~19 permanentes que sobrevivem.
