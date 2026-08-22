# 04 — Aceleração (Ramp) · Krenko, Mob Boss

- **Modo:** `improve` · **Data:** 2026-08-22 · **Identidade:** R (mono-vermelho)
- **Fonte de todo texto oracle:** `bin/mtgdb oracle` / `bin/mtgdb rulings` / `bin/mtgdb deck krenko-mob-boss`, puxados **nesta sessão**. Nenhum veredito de memória.
- **Cota recebida:** 4 entradas (2 NÚCLEO + 2 COMPLEMENTAR).
- **Pool de cortes exclusivo:** `Shiny Impetus` · `Mycosynth Wellspring` · `Redcap Thief` · `Goblin Glasswright // Craft with Pride` · `Reckless Lackey` · `Ember Hauler`. **Nada fora dele foi proposto para corte.** Terrenos não foram tocados (lane do `manabase-engineer`).
- **Preços:** `estimativa (Scryfall)`, cotação de 2026-08-22. ⚠️ marca > US$ 10. A régua real continua sendo o menor valor da LigaMagic (regra 2) — nenhuma dessas cartas tem cotação registrada em `mtgdb prices`.

---

## 0. Posição sobre o alvo de 6–7 — **concordo, com uma emenda de composição**

A meta genérica de 10–11 ramps padrão do `CLAUDE.md` está errada para este deck, e os números do orquestrador se sustentam quando eu os recalculo na lista:

| Evidência | Número apurado |
|---|---|
| MV médio (não-terrenos) | **2,38** — `mtgdb deck` |
| Curva não-terreno | 1:14 · 2:20 · 3:21 · 4:7 · 5:1 · 6+:0 |
| Custo de ativação do comandante | **`{T}`** — zero mana |
| Terrenos | 36, dos quais 32 Mountain |
| Peças que consomem mana excedente de forma **repetível** | `Dragon Mantle` (`{R}`: +1/+0) — **1**. `Massive Raid` e `Assault on Osgiliath` são X **de uso único**. |

Um mana rock genérico neste deck entrega, no T6, um recurso para o qual existe **um** escoadouro repetível em 99 cartas. Ramp padrão aqui é carta em branco disfarçada.

**Emenda:** o alvo "6–7 **fontes de mana**" só fecha se for lido como "6–7 **peças que tocam a economia de mana**". Comprar 4 produtores de mana pioraria o problema medido. A composição que eu entrego é:

| Eixo | Antes | Depois |
|---|---|---|
| Produção real de mana | Sol Ring · Arcane Signet = **2** | + `Skirk Prospector` = **3** |
| Redução de custo (mana virtual) | Goblin Warchief = **1** | + `Urza's Incubator` = **2** |
| Sink **repetível** (mana → tabuleiro) | Dragon Mantle = **1** | + `Umbral Mantle` · `Hammer of Purphoros` = **3** |
| **Total de peças de economia de mana** | **4** | **8** |
| *Contagem estrita do orquestrador* | *2* | ***6*** |

Ou seja: **bate o 6–7 pela contagem estrita**, mas o ganho real vem de dois produtores a menos e dois escoadouros a mais do que o alvo literal sugeriria.

---

## 1. Coleção primeiro (regra 7) — veredito escrito das 7 cartas que o orquestrador nomeou

Oracle de todas puxado nesta sessão. Regra 7 exige avaliar **antes** e justificar por escrito a dispensa.

| Carta (coleção) | Oracle relevante | Veredito como peça de mana | Motivo |
|---|---|---|---|
| **`Urza's Incubator`** {3} | *Creature spells of the chosen type cost {2} less to cast.* | ✅ **ENTRA — NÚCLEO** | **33 das 63 cartas não-terreno são feitiços de criatura Goblin** (32 criaturas + o comandante da zona de comando). Ver §2. |
| `Liquimetal Coating` {2} | *{T}: Target permanent becomes an artifact until end of turn.* | ❌ dispensada | **Produz zero mana e reduz zero custo.** O único uso mana-adjacente é virar um terreno em artefato para o `Krenko, Baron of Tin Street` sacrificar — 2 cartas + `{T}` para converter um terreno em contadores, uma vez por turno. `Hammer of Purphoros` faz a mesma conversão terreno→permanente sozinha, e ainda dá haste. |
| `Golden Egg` {2} | *ETB: draw a card. {1},{T}, Sac: Add one mana of any color.* | ❌ dispensada **como ramp** | **Mana-negativa:** paga 2, depois paga mais 1 para receber 1. "Any color" é letra morta em mono-R com 32 Mountains. É um cantrip em corpo de artefato — argumento do `draw-specialist`, não meu. Fica como **reserva de combustível do Baron** (§6). |
| `Instant Ramen` {2} | *Flash. ETB: draw a card. {2},{T}, Sac: gain 3 life.* | ❌ dispensada | Produz **zero** mana. É o `Golden Egg` sem a habilidade de mana. Não é carta desta fase. |
| `Oracle's Vault` {4} | *{2},{T}: impulso + brick counter; com 3 bricks, {T}: impulso grátis.* | ❌ dispensada | **É** um mana sink, e por isso mereceu conta. Mas: CMC 4 num deck mandado a **descer** a curva; 3 turnos e 6 manas para ficar grátis; e o card impulsionado precisa ser jogado no turno — com 36 terrenos, uma fração alta dos impulsos vira terreno inútil. `Hammer of Purphoros` é sink por `{2}{R}` que devolve um **3/3 artefato sacrificável** e ainda dá haste, por US$ 0,32. Reserva (§6). |
| `Panic Spellbomb` {1} | *{T}, Sac: target creature can't block. Morre → pague {R}: compre.* | ❌ dispensada **como ramp** | Produz zero mana. **Mas é a emenda mais barata do buraco que meus cortes abrem** (artefato descartável para o Baron) — entregue como handoff em §7. |
| `Thaumaton Torpedo` {1} | *{6}, {T}, Sac: destroy target nonland permanent. Custa {3} a menos se atacou com Spacecraft.* | ❌ dispensada | O deck tem **zero Spacecraft**, logo a ativação é **sempre `{6}`**. Como escoadouro de mana é a pior taxa disponível na coleção: 6 manas + `{T}` + sacrifício para uma remoção. E remoção é lane do `interaction-specialist`. |

**Conclusão da regra 7:** 1 das 7 entra (e é a peça mais impactante da rodada, a custo R$ 0). As outras 6 foram dispensadas com motivo escrito e nome da carta que cobre a função. As 3 compras propostas somam **≈ US$ 13,18 — estimativa (Scryfall)**.

---

## 2. Candidatas — entradas propostas (4/4)

### 🔴 NÚCLEO 1 · `Urza's Incubator` — {3} · Artifact · **coleção (R$ 0)**

> *As this artifact enters, choose a creature type. Creature spells of the chosen type cost {2} less to cast.*

| Eixo | Ficha |
|---|---|
| **F1** | Escolhe um tipo ao entrar (**Goblin**). Todo **feitiço de criatura** daquele tipo custa `{2}` a menos. Simétrico (vale para oponentes, irrelevante numa mesa sem Goblins). |
| **F2** | Não é corpo. |
| **F3** | **Artefato** → conta para o `Krenko, Baron of Tin Street` e para `Kuldotha Rebirth` (se o orquestrador a trouxer da coleção). |
| **F4** | Não recebe nada. |
| **F5** | **Redutor de custo — a função central.** |
| **F6** | T3. A partir do T3, **13 das 21 cartas de CMC 3 do deck passam a custar 1**: `Clamor Shaman`, `Goblin Chainwhirler`, `Goblin Rabblemaster`, `Goblin Sky Raider`, `Goblin Warchief`, `Grishnákh`, `Grotag Night-Runner`, `Guttersnipe`, `Krenko, Baron of Tin Street`, `Mudbutton Torchrunner`, `Rummaging Goblin`, `Squee` (+ `Redcap Thief`, que estou cortando). `Gundabad Opportunist` e `Zada` caem de 4 → 2. `Misty Mountains Raider` de 5 → 3. E **`Krenko, Mob Boss` passa a custar `{R}{R}`** — cai no T2 com `Sol Ring`, no T3 sem nada. |
| **F7** | Só reduz a parte **genérica**; `{R}{R}` é o piso do comandante. Não ajuda os 30 feitiços não-criatura. Empilha com `Goblin Warchief` só até o piso colorido. |

**Sinergias (≥2):** (1) o deck é **32/32 Goblins** — nenhuma outra carta do formato acerta 52% da lista; (2) ataca diretamente o engarrafamento medido em CMC 3 (21 cartas, 33% do deck), a causa apontada de *"lanço um feitiço por turno do T3 ao T6"*; (3) adianta o comandante 1–2 turnos, que é literalmente o que a fase de ramp existe para fazer; (4) é artefato — recompõe a contagem que meus cortes reduzem.

**Por que é ramp e não "outra coisa":** um redutor de `{2}` sobre 33 cartas produz mais mana efetiva por partida do que qualquer rocha de 2 manas, **e não entrega um recurso sem destino** — ele diminui a demanda em vez de aumentar a oferta. Este é o ponto que o alvo de 10–11 rochas erra.

**Preço:** já é do usuário. `estimativa (Scryfall)` US$ 19,55 caso precisasse ser comprada — mais uma razão para usar a que existe.

---

### 🔴 NÚCLEO 2 · `Umbral Mantle` — {3} · Artifact — Equipment · **compra** · US$ 12,52 ⚠️ `estimativa (Scryfall)`

> *Equipped creature has "{3}, {Q}: This creature gets +2/+2 until end of turn." ({Q} is the untap symbol.) — Equip {0}*

| Eixo | Ficha |
|---|---|
| **F1** | Concede ao portador uma habilidade ativada de custo `{3}` + **desvirar o próprio portador** (`{Q}`), dando +2/+2 até o fim do turno. **Equip `{0}`.** |
| **F2** | Não é corpo. **Dá** ao corpo alheio um custo de ativação. |
| **F3** | **Artefato** (Baron / `Kuldotha Rebirth`) **e Equipamento**. |
| **F4** | — |
| **F5** | **É o mana sink sem fundo que o deck não tem.** Krenko tapa → cria X tokens → `{3}` + `{Q}` desvira Krenko → tapa de novo → cria **2X** tokens. Repetível enquanto houver mana. O tag do próprio Scryfall Tagger é literalmente `bottomless-mana-sink`. Também torna o portador **"modified"** (equipamento anexado). |
| **F6** | T3, equip `{0}` no mesmo turno. Do T6 em diante, com 8–9 manas desvirados, são **2 ativações extras de Krenko por turno** — o crescimento passa de N→2N para N→8N. |
| **F7** | **Rulings puxados (`mtgdb rulings`, 2008-05-01):** (a) `{Q}` só pode ser pago se a permanente estiver **virada** — então é obrigatório tapar Krenko primeiro, o que é exatamente a sequência natural; (b) uma criatura que não está sob seu controle desde o início do turno **não pode** ativar `{Q}` **a menos que tenha haste** — resolvido por `Goblin Warchief`, `Goro-Goro` e pela entrada COMPLEMENTAR 2; (c) o desvirar é **custo**, não pode ser respondido. Compete com o `{T}` de Krenko no sentido de que resolve o conflito, não o cria. |

**Sinergias (≥2):** (1) escoadouro direto da queixa medida *"do T6 em diante sobra mana e não há onde gastar"* — converte mana excedente na única moeda que o deck sabe gastar, corpos Goblin; (2) `equip {0}` = migra de graça para o próximo Krenko depois de uma remoção, e para `Krenko, Baron of Tin Street`/`Fanatical Firebrand`/`Rummaging Goblin` (os 4 corpos com `{T}` que a auditoria marcou como atrito — o Mantle **transforma o atrito em rendimento**); (3) preenche o pacote **"modified"** que a auditoria §5.3 marcou como *pela metade* (3 payoffs, **0 equipamentos**): liga `Goro-Goro`, `Ambitious Assault` e `Akki Ember-Keeper`; (4) artefato, recompondo a contagem.

⚠️ **Divulgação obrigatória ao orquestrador (lane do `wincon-tester`):** `Krenko, Mob Boss` + `Umbral Mantle` + `Skirk Prospector` (a COMPLEMENTAR 1) é um **motor de Goblins arbitrariamente grande** a partir de ~4 Goblins em campo — sacrifica 3 Goblins pelo `{R}{R}{R}`, paga `{3}`+`{Q}`, redobra. São 3 cartas, todas úteis isoladamente, e **não é um lock** (não impede ninguém de jogar — regra 11 ok); é uma condição de vitória. Ainda exige haste para fechar no mesmo turno. Sinalizo para que a decisão de power level seja do usuário, não minha.

---

### 🟡 COMPLEMENTAR 1 · `Skirk Prospector` — {R} · Creature — Goblin 1/1 · **compra** · US$ 0,34 `estimativa (Scryfall)`

> *Sacrifice a Goblin: Add {R}.*

| Eixo | Ficha |
|---|---|
| **F1** | Sacrifica um Goblin (**qualquer**, inclusive ela mesma): adiciona `{R}`. Sem custo de mana, sem `{T}`, **sem limite por turno**, velocidade de instantâneo. |
| **F2** | Corpo 1/1 por 1 mana. Bloqueia. É sacrificável como recurso. |
| **F3** | **É Goblin** → soma ao X de Krenko. Cada Goblin de 1 mana em campo antes de Krenko vale **o dobro** por ativação (eixo E5 da auditoria). |
| **F4** | Recebe anthems (`Ferocity of the Wilds`, `Quest for the Goblin Lord`) e os contadores +1/+1 do `Krenko, Baron of Tin Street`. |
| **F5** | **Ramp explosivo:** com 15 Goblins em campo, é um ritual de até 15 `{R}`. Também é o **único sac outlet repetível e gratuito** que o deck passaria a ter. |
| **F6** | T1 — o slot que a auditoria pediu engordar (14 cartas em CMC 1, alvo 16–18). |
| **F7** | Cada mana custa um corpo, que é o recurso do X de Krenko. **Ruling puxado (2022-12-08):** se for sacrificar vários Goblins incluindo ela, sacrifique a `Skirk Prospector` **por último**. Não usa `{T}` — zero atrito com Krenko. |

**Sinergias (≥2):** (1) corpo Goblin de 1 mana = o slot de maior alavancagem do comandante; (2) ramp **explosivo** de verdade, que só existe porque o deck tem largura — é ramp que **só este deck** consegue jogar; (3) destrava `Boggart Shenanigans` (a auditoria §4.5 marcou o caminho de vitória nº 7 como **MORTO** por falta exatamente disto), `Mudbutton Torchrunner` e `Ember Hauler` sob demanda; (4) resposta a board wipe: converte o enxame condenado em mana e em dano antes que o wipe resolva — ataca a queixa *"apanho de board wipe"* sem gastar slot de proteção.

---

### 🟡 COMPLEMENTAR 2 · `Hammer of Purphoros` — {1}{R}{R} · Legendary Enchantment Artifact · **compra** · US$ 0,32 `estimativa (Scryfall)`

> *Creatures you control have haste. — {2}{R}, {T}, Sacrifice a land: Create a 3/3 colorless Golem enchantment artifact creature token.*

| Eixo | Ficha |
|---|---|
| **F1** | Duas linhas: (a) **todas** as suas criaturas têm haste — permanente, não é one-shot; (b) `{2}{R}`, `{T}`, sacrifica um **terreno**: cria um **3/3 Golem enchantment artifact creature token**. |
| **F2** | Não é corpo, mas **fabrica** corpo 3/3 — o deck tem só 6 corpos com poder ≥3. |
| **F3** | **Artefato e Encantamento.** O Golem que ela cria também é **artefato**. |
| **F4** | — |
| **F5** | (a) Haste: `Krenko, Mob Boss` **ativa no turno em que entra** — corta um turno inteiro do relógio. A auditoria mediu **2** habilitadores permanentes de haste em 99 cartas (≈15% de disponibilidade); isto vai para **3**. Também libera o `{Q}` do `Umbral Mantle` num Krenko recém-lançado (ruling acima). (b) **Mana sink que come terreno**: é o escoadouro desenhado para a queixa exata — *36 terrenos, MV 2,38, o 6º terreno em diante é carta em branco*. Converte terreno excedente em 3/3. |
| **F6** | T3. A partir daí, todo turno de flood vira um 3/3. |
| **F7** | Usa o próprio `{T}` → **um Golem por turno**, não é ilimitado. Sacrificar terreno é irreversível — só ative com 7+ terrenos em campo. Lendária (irrelevante, singleton). O Golem **não é Goblin**: não soma ao X de Krenko. A haste é parcialmente redundante com `Goblin Warchief`/`Goro-Goro` — mas 3 fontes em 99 cartas ainda é ≈22%, contra 15%. |

**Sinergias (≥2):** (1) o `{T}`+sac é o **único** escoadouro do deck que consome o recurso que sobra de verdade (terreno), e não o que falta (cartas); (2) o Golem é **artefato descartável renovável** → alimenta `Krenko, Baron of Tin Street` (`{T}`, sac artefato: +1/+1 em cada Goblin) **e** dispara a segunda linha dele (*"Whenever an artifact is put into a graveyard from the battlefield, you may pay {R}: create a 1/1 red Goblin token with haste"*) — ou seja, terreno excedente → 3/3 → contador em todo Goblin **+** um Goblin novo com haste, todo turno; (3) haste resolve o eixo E1, o mais crítico da auditoria; (4) casa com o `manabase-engineer`: se o deck for para 34–35 terrenos, esta carta é o que torna aceitável manter os que sobrarem.

---

## 3. Cortes propostos — ficha F1–F7 completa e protocolo da §2 do checklist

> Só cartas do meu pool exclusivo. Nenhum corte foi decidido antes de a ficha estar escrita. As **mesmas** condições assumidas para defender as entradas (`Urza's Incubator` em campo, anthems em campo, 8+ manas no T6) foram aplicadas às cartas que saem.

### CORTE 1 · `Shiny Impetus` {2}{R} — Enchantment — Aura

- **F1** *Enchant creature. Encantada ganha +2/+2 e é **goaded** (ataca a cada combate se puder, e ataca um jogador que não seja você se puder). Sempre que a encantada ataca, você cria um Treasure.*
- **F2** Não é corpo. **F3** Encantamento; **produz artefato** (Treasure). **F4** Não recebe nada. **F5** Ramp recorrente (1 Treasure por combate) + goad (interação política). **F6** T3, no meio do engarrafamento de 21 cartas.
- **F7** **Correção à auditoria (§6.2), que afirmou que ela dá `+2/+2` a uma criatura *adversária*:** o oracle diz apenas *"Enchant creature"* — ela **pode** ser colocada numa criatura sua, e o goad ("ataca um jogador que não seja você") é automático nesse caso. Mas nas duas leituras ela falha: **na criatura do oponente**, você entrega `+2/+2` permanente por 1 Treasure/turno e a peça morre com a criatura deles; **na sua própria**, você importa exatamente a anti-sinergia de ataque compulsório que a auditoria §6.3 usou para condenar o `Goblin Rabblemaster` — 1/1 forçado a atacar contra bloqueadores maiores **destrói o X da próxima ativação de Krenko**. Em ambos os casos: 3 manas, **zero corpo Goblin**, 2-por-1 contra remoção num deck com 0 proteção.

**Protocolo de corte:**

| Função | Quem cobre depois do corte |
|---|---|
| Ramp recorrente (1 Treasure/combate, condicional a atacar) | **`Skirk Prospector`** — mana recorrente, velocidade de instantâneo, ilimitada por turno, do **seu** lado do campo, por 1 mana em vez de 3. **Cobertura superior.** |
| Artefato para `Krenko, Baron of Tin Street` | **`Hammer of Purphoros`** (Golem artefato renovável, 1/turno) + `Urza's Incubator` + `Umbral Mantle` na contagem. **Coberta e melhorada** (de one-shot para renovável). |
| Goad / interação política | **DESCOBERTA — custo aceito.** Resta `Grishnákh, Brash Instigator` como única interação de tempo. Numa mesa de 4 com um deck que já é ameaça visível, goad tem valor marginal; e a auditoria §4.3 registra a interação em 10 peças (meta batida). |

---

### CORTE 2 · `Mycosynth Wellspring` {2} — Artifact

- **F1** *When this artifact enters **or is put into a graveyard from the battlefield**, you may search your library for a **basic land card**, reveal it, put it **into your hand**, then shuffle.*
- **F2** Não é corpo. **F3** **Artefato — combustível duplo do `Krenko, Baron of Tin Street`** (dispara ao entrar **e** ao ir para o cemitério, e o próprio ir-para-o-cemitério dispara a linha 3 do Baron). **F4** — **F5** Fixação de cor (letra morta: 32 Mountains). **F6** T2 sem nenhum impacto no tabuleiro.
- **F7** **Não é ramp:** o terreno vai para a **mão**, não para o campo. Num deck com **36 terrenos e MV médio 2,38**, buscar mais um terreno básico **alimenta o flood** — é a carta que agrava a queixa nº 4 (*"mana travando"*) em vez de resolvê-la. É a única carta do deck cujo efeito escrito trabalha **contra** o diagnóstico da auditoria.

**Protocolo de corte:**

| Função | Quem cobre depois do corte |
|---|---|
| Artefato sacrificável para o Baron (gatilho duplo) | **`Hammer of Purphoros`** — o Golem 3/3 é `enchantment artifact creature token`, **renovável todo turno**, e ao ser sacrificado ao Baron dá +1/+1 em cada Goblin **e** dispara a 3ª linha do Baron (pague `{R}`: Goblin 1/1 com haste). Trocamos **1 gatilho duplo de uso único** por **um artefato descartável por turno**. **Cobertura superior.** |
| Contagem de artefatos | 3 → **5** na lista (`Sol Ring`, `Arcane Signet`, `Urza's Incubator`, `Umbral Mantle`, `Hammer of Purphoros`), +Golem/turno, +Treasure de `Reckless Lackey` (mantida). **Coberta.** |
| Busca de terreno básico / anti-screw | **DESCOBERTA — custo aceito e desejado.** Com 32 Mountains em 36 terrenos, a probabilidade de screw de cor é matematicamente nula; a auditoria §8.1 pede *menos* terreno, não mais. |
| Fixação de cor | Irrelevante em mono-R. |

---

### CORTE 3 · `Redcap Thief` {2}{R} — Creature — Goblin Rogue 2/3

- **F1** *When this creature enters, create a Treasure token.*
- **F2** Corpo 2/3 — bloqueia bem para o tamanho. **F3** **É Goblin** (soma ao X) **e** gera artefato. **F4** Recebe anthems e os contadores do Baron. **F5** Ramp pontual de 1 mana + fixação. **F6** **T3 — mais uma das 21 cartas do engarrafamento**; com `Urza's Incubator` cairia para T1, o que é uma defesa legítima e por isso a avalio sob a mesma condição que usei para as entradas.
- **F7** Taxa nua: **3 manas por um 2/3 e um único Treasure**. Sob `Urza's Incubator` vira 1 mana por 2/3 + Treasure — bom, mas depende de outra carta estar em campo; sem ela, é o pior corpo-por-mana do deck.

**Protocolo de corte:**

| Função | Quem cobre depois do corte |
|---|---|
| Corpo Goblin (soma ao X de Krenko) | **`Skirk Prospector`** — Goblin por **1** mana em vez de 3, no tier que a auditoria mandou engordar (CMC 1: 14 → alvo 16–18). Perde-se 1 ponto de poder e 2 de resistência; **custo aceito**, porque o X de Krenko conta cabeças, não estatísticas. |
| Treasure (mana pontual) | **`Skirk Prospector`** (mana recorrente) + **`Urza's Incubator`** (redução permanente). **Cobertura superior.** |
| Artefato para o Baron | **`Hammer of Purphoros`** (Golem renovável). **Coberta.** |
| Receptor de anthem / contadores | `Skirk Prospector` e os outros 30 Goblins. **Coberta.** |
| Bloqueador de resistência 3 | **Parcialmente descoberta — custo aceito.** Restam `Grotag Night-Runner` (2/3), `Goblin Chainwhirler` (3/3), `Misty Mountains Raider` (4/4), `Zada` (3/3), `Krenko` (3/3), `Krenko, Baron` (3/3). |

---

### CORTE 4 · `Goblin Glasswright // Craft with Pride` {1}{R} // {R} — Creature — Goblin Sorcerer 2/2 // Sorcery

- **F1** *This creature enters **prepared**. (While it's prepared, you may cast a copy of its spell. Doing so unprepares it.)* — o feitiço preparado é **`Craft with Pride`** `{R}`: *Create a Treasure token.* **Rulings puxados (2026-03-20):** a cópia fica no exílio enquanto a criatura estiver em campo e preparada; ao lançá-la, a criatura deixa de estar preparada; só o **controlador atual** pode lançá-la; se a criatura sair do campo, a cópia deixa de existir.
- **F2** Corpo 2/2. **F3** **Goblin** (soma ao X) + tipo Sorcery na cara de trás + **gera artefato** (Treasure). **F4** Recebe anthems e contadores. **F5** Mana (1 Treasure, adiado, por `{R}` extra). **F6** T2 — o tier saudável.
- **F7** O Treasure **não é grátis**: custa `{R}` a mais e uma janela de feitiço, e some se a criatura morrer antes de você pagar. Total real: **3 manas em duas parcelas por um 2/2 Goblin + 1 Treasure** — a mesma taxa do `Redcap Thief`, só que parcelada. É a peça do meu pool cujas funções são todas duplicadas por outras.

**Protocolo de corte:**

| Função | Quem cobre depois do corte |
|---|---|
| Corpo Goblin CMC 2 | O deck fica com **18** cartas em CMC 2 (era 20) e **31 Goblins**. `Skirk Prospector` repõe uma cabeça a CMC 1. `Akki Ember-Keeper`, `Conspicuous Snoop`, `Ember Hauler`, `Fissure Wizard`, `Goblin Cratermaker`, `Goblin Wardriver`, `Goblin-town Flunkies`, `Goro-Goro`, `Innocent Bystander` seguram o tier 2. **Coberta.** |
| Treasure / mana adiado | **`Skirk Prospector`** + **`Urza's Incubator`**. **Cobertura superior.** |
| Artefato para o Baron | **`Hammer of Purphoros`** (Golem renovável). **Coberta.** |
| Receptor de anthem | Os outros 30 Goblins. **Coberta.** |
| Flexibilidade "spell no exílio" (cara de trás como recurso) | **DESCOBERTA — custo aceito.** Nenhuma outra carta do deck usa a mecânica; não há payoff de "lançar do exílio" na lista. |

---

## 4. As 2 cartas do meu pool que eu **NÃO** corto — e por quê

Regra 4: corte só sai depois da ficha. Estas duas **passaram** na ficha e ficam.

### ✅ FICA · `Reckless Lackey` {R} — Goblin Pirate 1/2

- **F1** *First strike, haste. `{2}{R}`, Sacrifice this creature: **Draw a card and create a Treasure token**.*
- **F2** Corpo 1/2 **com haste** → entra e **já soma ao X de Krenko no mesmo turno**. First strike faz um 1/2 trocar com 2/2 sem morrer. **F3** Goblin + gera artefato. **F4** Anthems/contadores. **F5** Mana + **carta**. **F6** **T1** — o tier que a auditoria mandou engordar. **F7** A ativação custa 3 manas e o corpo.
- **Veredito:** ela é, sozinha, **três** das coisas que esta fase foi encarregada de comprar — corpo Goblin de 1 mana, **mana sink** (`{2}{R}` de sobra vira carta) e gerador de artefato descartável. Cortá-la seria cortar a solução para pagar pela solução.

### ✅ FICA · `Ember Hauler` {R}{R} — Goblin 2/2

- **F1** *`{1}`, Sacrifice this creature: It deals 2 damage to any target.*
- **F2** Corpo 2/2. **F3** Goblin (soma ao X); ao morrer é **um Goblin indo para o cemitério** → dispara `Boggart Shenanigans`. **F4** Anthems/contadores. **F5** — **F6** T2. **F7** `{R}{R}` é trivial com 32 Mountains; sacrificar tira uma cabeça do X.
- **Veredito:** cinco funções — corpo Goblin CMC 2, **remoção sob demanda**, alcance à face, **mana sink de `{1}`** e gatilho de `Boggart Shenanigans` (que fica muito melhor com `Skirk Prospector` entrando). A remoção dela **não está no pool de corte de nenhum outro especialista** — cortá-la aqui removeria uma peça de interação sem ninguém cobrindo. Fora do meu escopo de corte por essa razão também.

---

## 5. Tabela de swaps

| Sai | → | Entra | CMC | Classe | Motivo em uma linha | Preço `estimativa (Scryfall)` |
|---|---|---|---|---|---|---|
| `Shiny Impetus` (3) | → | **`Urza's Incubator`** | 3 | **NÚCLEO** · redução de custo | Troca 1 Treasure/combate condicional por `{2}` de desconto em **33 cartas** (52% do deck) e Krenko a `{R}{R}` | **R$ 0 — coleção** |
| `Mycosynth Wellspring` (2) | → | **`Umbral Mantle`** | 3 | **NÚCLEO** · mana sink sem fundo | Troca a carta que **agrava** o flood pelo escoadouro que converte o flood em ativações extras de Krenko | US$ 12,52 ⚠️ |
| `Redcap Thief` (3) | → | **`Skirk Prospector`** | 1 | COMPLEMENTAR · ramp explosivo + corpo | Mesmo Goblin por 1 mana em vez de 3, com mana **recorrente** e o 1º sac outlet grátis do deck | US$ 0,34 |
| `Goblin Glasswright // Craft with Pride` (2) | → | **`Hammer of Purphoros`** | 3 | COMPLEMENTAR · mana sink + haste | Troca 1 Treasure adiado por: haste permanente para todos + terreno excedente → 3/3 artefato **todo turno** | US$ 0,32 |

**Total de compra: ≈ US$ 13,18 — `estimativa (Scryfall)`, 2026-08-22.** Uma carta acima de US$ 10 (`Umbral Mantle` ⚠️). **A régua real é o menor valor da LigaMagic (regra 2)** — nenhuma das quatro tem cotação em `mtgdb prices`; recomendo capturar `Umbral Mantle` antes de fechar.

### Impacto agregado

| Métrica | Antes | Depois |
|---|---|---|
| Fontes de mana (produção + redução) | 3 | **5** |
| Mana sinks **repetíveis** | 1 (`Dragon Mantle`) | **3** (`Umbral Mantle` ilimitado · `Hammer of Purphoros` 1/turno · `Dragon Mantle`) |
| Mana sinks totais (contando os de uso único) | 6 | **8** |
| Peças que tocam a economia de mana (contagem estrita do orquestrador) | **2** | **6** ✅ alvo batido |
| Artefatos na lista | 3 | **5** (+1 Golem/turno) |
| Artefatos **descartáveis** para o Baron | 4 Treasures de uso único + Wellspring | **1 Golem por turno** + Treasure do `Reckless Lackey` + `Signet`/`Sol Ring` no late — **de one-shot para renovável** |
| Criaturas / Goblins | 32 | **31** |
| Habilitadores **permanentes** de haste | 2 (≈15%) | **3** (≈22%) |
| Curva não-terreno | 1:14 · 2:20 · 3:21 · 4:7 · 5:1 | 1:**15** · 2:**18** · 3:21 · 4:7 · 5:1 |
| Curva **efetiva** com `Urza's Incubator` em campo | — | 13 das 21 cartas de CMC 3 passam a CMC **1**; Krenko de 4 → **2** |

> **Nota honesta sobre a curva:** o número **bruto** de cartas em CMC 3 não muda (saem `Shiny Impetus` e `Redcap Thief`, entram `Umbral Mantle` e `Hammer of Purphoros`). O ganho de curva desta fase é **efetivo**, via `Urza's Incubator`, não estrutural — e as duas cartas de CMC 3 que entram são **deploys únicos**, não peças que competem por turno. Reduzir o CMC 3 bruto de 21 para 12–14 (alvo da auditoria §8.5) continua sendo trabalho das outras fases e do `manabase-engineer`.

---

## 6. Reservas (não consomem cota)

| Carta | CMC | Função | Preço | Quando usar |
|---|---|---|---|---|
| `Herald's Horn` | 3 | Feitiços de Goblin `{1}` a menos + olha o topo todo upkeep e põe Goblin na mão | US$ 5,14 | Se o orquestrador quiser um **2º redutor**. Empilha mal com `Urza's Incubator` no piso colorido, mas o gatilho de upkeep é card advantage real (lacuna nº 3). |
| `Magewright's Stone` | 2 | `{1}`,`{T}`: desvira criatura com habilidade de `{T}` | US$ 10,08 ⚠️ | Substituta de `Umbral Mantle` se o combo arbitrário for **recusado** pelo usuário: 1 desvirar por turno, sem loop, e é CMC 2. |
| `Treasure Nabber` | 3 | **Goblin** 2/2; rouba as rochas de mana dos oponentes quando as tapam | US$ 3,40 | Se a mesa for pesada em artefatos. Corpo Goblin + ramp por roubo, mas depende do metagame. |
| `Golden Egg` (coleção) | 2 | Cantrip + artefato descartável | R$ 0 | **Não como ramp** (é mana-negativa). Só se o orquestrador quiser reforçar o combustível do Baron a custo zero. |
| `Oracle's Vault` (coleção) | 4 | Mana sink `{2}`/turno → impulso; grátis com 3 bricks | R$ 0 | Se `Hammer of Purphoros` for recusado. Pior taxa e pior curva, mas custo zero. |

---

## 7. Handoffs ao orquestrador

1. **⚠️ Motor arbitrário divulgado:** `Krenko, Mob Boss` + `Umbral Mantle` + `Skirk Prospector` produz Goblins arbitrariamente grandes com ~4 Goblins em campo. Não é lock (regra 11 ok), é wincon — mas **muda o power level** e a decisão é do usuário. Se ele recusar, a substituta é `Magewright's Stone` (§6), que mantém o mana sink sem o loop.
2. **Combustível do Baron:** eu **cubro** o buraco com o Golem renovável do `Hammer of Purphoros`, mas se o `Hammer` for recusado, `Panic Spellbomb` (coleção, R$ 0, CMC 1: artefato descartável + "não pode bloquear" + compra ao morrer) é a emenda mais barata — carta do slot do `draw-specialist`, não do meu.
3. **Sobreposição com outras lanes:** a haste do `Hammer of Purphoros` pode colidir com uma proposta de `Goblin Chieftain`/`Fervor` do especialista de tema. Se colidir, o valor único do `Hammer` continua sendo o **sink terreno→3/3**, que nenhum lorde de haste entrega.
4. **`Goblin Anarchomancer` foi descartada por ilegalidade, não por avaliação:** o oracle puxado nesta sessão mostra `{R}{G}` — **identidade GR**, fora da identidade do comandante. Registro porque é o candidato óbvio de "Goblin redutor de custo" e alguém vai propô-la.
5. **Não toquei em terreno nenhum**, mas o `Hammer of Purphoros` conversa diretamente com o corte de 36 → 34–35 terrenos: ele é o que torna aceitável o terreno excedente que sobrar.
6. **Registro em `decisions.md`:** 4 cortes e 4 entradas, todos com ficha F1–F7 nesta página. Nenhuma das 4 entradas é reposição de corte anterior — `decisions.md` só contém o baseline de 2026-08-21.
