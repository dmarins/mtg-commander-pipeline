# 02b — Fase 2a · Lordes e conversores (lacuna nº 1)

- **Data:** 2026-08-22 · **Modo:** `improve` · **Agente:** theme-analyst
- **Mandato:** 8 entradas (5 NÚCLEO + 3 COMPLEMENTAR) contra a lacuna nº 1 da auditoria (`02-theme.md` §7).
- **Pool de cortes:** exclusivo, recebido do orquestrador. **Não saí dele.**
- **Oracle:** todo texto citado foi puxado nesta sessão com `bin/mtgdb oracle` / `bin/mtgdb search`. Nada de memória.
- **`decisions.md` conferido:** o registro tem **uma única linha** (`baseline`, 2026-08-21). Nenhuma das 8 entradas é reposição de corte anterior. Regra 5 satisfeita por vacuidade — mas conferida.
- **Preços:** `mtgdb prices` não tinha cotação de nenhuma das 8. Os valores abaixo são **estimativa (Scryfall), 2026-08-22**, menor impressão. **Não** registrei nada em `prices.tsv` — a série é de observações da LigaMagic, e estimativa de proxy não entra lá (regra 2).

## Duas correções à minha própria auditoria
Puxando o oracle de novo para esta fase, dois tipos estavam errados em `02-theme.md` §3:
- `Ambitious Assault` é **Instant**, não Sorcery.
- `Warlord's Fury` é **Sorcery**, não Instant.
Nenhuma das duas muda o veredito (ambas seguem `marginal`), mas ficam registradas.

---

## 1. O que eu procurei, e o que existe

| Prioridade do orquestrador | Busca | Achado |
|---|---|---|
| **1. Lorde estático de Goblin** | `mtgdb search '"Goblins you control get" OR "Goblin creatures you control get"' -id R` | **7 no jogo inteiro**: `Goblin Chieftain`, `Goblin Trashmaster`, `Rundvelt Hordemaster`, `Hobgoblin Bandit Lord`, `Goblin General`, `Battle Cry Goblin`, `Dropkick Bomber`. Destes, **4 são anthem permanente de verdade** (Chieftain, Trashmaster, Hordemaster, Hobgoblin) — General e Battle Cry são pump temporário, Dropkick é condicional. |
| **2. Conversor permanente de largura em dano** | `mtgdb search '"creature enters" AND "deals" AND "each opponent"'` · `'"Sacrifice a creature:"'` · `'"creatures you control have trample"'` | Três classes distintas: **(i) dano por ETB** (`Impact Tremors`, `Purphoros, God of the Forge`) — ignora bloqueador por completo; **(ii) sac outlet grátis** (`Goblin Bombardment`, `Skirk Prospector`) — converte corpo em dano/mana e responde a wipe; **(iii) multiplicador exponencial de combate** (`Shared Animosity`). Trample coletivo permanente em mono-R praticamente **não existe** (4 resultados, todos ruins) — por isso a rota certa é *ignorar* o bloqueio, não atravessá-lo. |
| **3. Haste permanente** | coleção + busca | `Goblin Chieftain` resolve como **efeito secundário de um lorde**. Ver §4 sobre `Haunted Cloak`. |

**Nenhuma das 8 entradas está na coleção** (`mtgdb collection`, conferido carta a carta). Justificativa de dispensa da coleção em §4 — regra 7 atualizada.

---

## 2. Tabela de swaps

| # | Sai | CMC | Entra | CMC | US$ est. | Marca |
|---|---|---|---|---|---|---|
| 1 | `Crash Through` | 1 | **`Impact Tremors`** | 2 | 2,50 | **NÚCLEO** |
| 2 | `Warlord's Fury` | 1 | **`Goblin Bombardment`** | 2 | 2,51 | **NÚCLEO** |
| 3 | `Goblin Rabblemaster` | 3 | **`Goblin Chieftain`** | 3 | 4,64 | **NÚCLEO** |
| 4 | `Ambitious Assault` | 3 | **`Shared Animosity`** | 3 | 3,76 | **NÚCLEO** |
| 5 | `Bravado` | 2 | **`Purphoros, God of the Forge`** | 4 | **⚠️ 28,76** | **NÚCLEO** |
| 6 | `Tidings of War` | 1 | **`Rundvelt Hordemaster`** | 2 | 3,17 | complementar |
| 7 | `Akki Ember-Keeper` | 2 | **`Skirk Prospector`** | 1 | 0,34 | complementar |
| 8 | `Mob Mentality` | 1 | **`Hobgoblin Bandit Lord`** | 3 | 4,36 | complementar |

**Total estimado: US$ 50,04** — `estimativa (Scryfall), 2026-08-22`. Sem Purphoros: **US$ 21,28**.
⚠️ Só `Purphoros, God of the Forge` passa de US$ 10. Confira na LigaMagic.

**Efeito na curva (declarado, não escondido):** saem 14 de CMC somado, entram 20 → **+6 em 8 slots**. O CMC 3 vai de **21 → 22**. Isso **piora marginalmente** o engarrafamento que eu mesmo apontei em `02-theme.md` §8.3. Aceito o custo por dois motivos: (a) 6 das 8 entradas são **permanentes que ficam em jogo**, contra 6 das 8 saídas que eram **one-shot que cicla**; (b) o desengarrafamento do CMC 3 é trabalho da fase 6 (manabase/cortes), que tem o pool certo para isso. **Alternativa se o orquestrador quiser a curva neutra:** trocar `Hobgoblin Bandit Lord` (3) por **`Battle Cry Goblin`** (2, US$ ~0,50) — perde-se o anthem estático, ganha-se um **mana sink repetível** (`{1}{R}`: Goblins +1/+0 e haste), que ataca a queixa nº 4.

---

## 3. Gate de swap — ficha F1–F7 de cada carta que **sai**

> **Simetria de critério (§3 do checklist):** todas as comparações abaixo assumem o **mesmo** estado de jogo nos dois lados — **10 Goblins no campo depois de uma ativação de Krenko**. É a condição que eu uso para elogiar a entrada, então é a condição em que a saída é medida.

### Swap 1 — `Crash Through` → `Impact Tremors`
`Crash Through` {R} · Sorcery — *"Creatures you control gain trample until end of turn. Draw a card."*
- **F1** Trample coletivo até o fim do turno + compra 1. **F2** Não é corpo. **F3** Sorcery — soma ao cemitério que alimenta `Blitz of the Thunder-Raptor` e à contagem de `Guttersnipe`. **F4** Não recebe nada. **F5** Concede trample. **F6** T1. **F7** **Não gatilha `Zada`** (sem alvo); cantrip = net 0 cartas.

| Função dela | Quem cobre depois |
|---|---|
| Trample coletivo | `Ferocity of the Wilds` (permanente, e ainda dá +1/+0) — **já no deck**. `Impact Tremors` torna a questão irrelevante: o dano deixa de passar pelo combate. |
| Cicla (net 0) | Fase 3 (draw). Não é vantagem de carta, é reposição. |
| Sorcery no cemitério (`Blitz`) / contagem de `Guttersnipe` | 19 outros instants/sorceries continuam na lista. **Perda: 1 de 20 (5%).** Custo aceito. |

**Sob 10 Goblins:** `Crash Through` dá trample a dez 1/1. Bloqueados por 2/2, o excesso é **0**. `Impact Tremors` (*"Whenever a creature you control enters, this enchantment deals 1 damage to each opponent"*) converte a **própria ativação** de Krenko em **10 de dano a cada oponente** — 30 numa mesa de quatro, sem combate, sem bloqueador, e **repete a cada ativação**.

### Swap 2 — `Warlord's Fury` → `Goblin Bombardment`
`Warlord's Fury` {R} · Sorcery — *"Creatures you control gain first strike until end of turn. Draw a card."*
- **F1** First strike coletivo + compra 1. **F2** — **F3** Sorcery (cemitério/`Guttersnipe`). **F4** — **F5** Concede first strike. **F6** T1. **F7** **Não gatilha `Zada`**; cantrip net 0.

| Função dela | Quem cobre depois |
|---|---|
| First strike coletivo | `Assault on Osgiliath` (double strike **contém** first strike e ainda dobra o dano) — **já no deck**. `Reckless Lackey` tem first strike próprio. |
| Cicla | Fase 3. |
| Sorcery no cemitério | 19 restantes. Custo aceito. |

**Sob 10 Goblins:** first strike em dez 1/1 contra bloqueadores 2/2 não muda um único resultado de combate — os 1/1 morrem igual e dão 1 de dano igual. Ganho ≈ **0**. `Goblin Bombardment` (*"Sacrifice a creature: This enchantment deals 1 damage to any target"*) converte os mesmos 10 corpos em **10 de dano dirigível**, e como `Boggart Shenanigans` já está na lista (*"Whenever another Goblin you control is put into a graveyard from the battlefield… 1 damage to target player"*), cada Goblin sacrificado vale **2**: **20 de dano**. É também a única carta da lista que **responde a board wipe e a exílio** — você sacrifica em resposta e transforma a perda em dano. **Ela ressuscita sozinha o caminho de vitória nº 7 que eu declarei morto** em `02-theme.md` §4.5.

### Swap 3 — `Goblin Rabblemaster` → `Goblin Chieftain`
`Goblin Rabblemaster` {2}{R} · Creature — Goblin Warrior 2/2 — *"Other Goblin creatures you control attack each combat if able. At the beginning of combat on your turn, create a 1/1 red Goblin creature token with haste. Whenever this creature attacks, it gets +1/+0 until end of turn for each other attacking Goblin."*
- **F1** Ataque compulsório dos outros Goblins; 1 token com haste por combate; auto-pump ao atacar. **F2** Corpo 2/2. **F3** **É Goblin** — soma ao X de Krenko. **F4** Recebe anthems, contadores, Auras. **F5** Gera corpo por turno. **F6** T3. **F7** **O ataque compulsório é anti-sinergia de primeira ordem** e piora com as entradas desta fase: com `Goblin Bombardment` você quer *sacrificar*, não atacar; com `Impact Tremors`/`Purphoros` você não precisa atacar.

| Função dela | Quem cobre depois |
|---|---|
| Corpo Goblin CMC 3 (soma ao X) | `Goblin Chieftain` — **também** Goblin, **também** CMC 3. Troca 1-por-1, sem perda de contagem. |
| Gerador de 1 Goblin com haste por combate | `Krenko, Mob Boss` (o motor de tokens do deck), `Krenko, Baron of Tin Street`, `Hordeling Outburst`, `Krenko's Command`, `Goblin Surprise`, `Goblin War Party`. **Perda real: ~1 token/turno.** Custo aceito — `Goblin Chieftain` devolve mais que isso em dano. |
| Anthem-ao-atacar (só nela mesma) | `Shared Animosity` e `Goblin Chieftain` fazem isso **em todas** as criaturas. |
| Ataque compulsório | **Perder isto é ganho, não custo.** |

**Sob 10 Goblins:** `Rabblemaster` acrescenta 1 corpo (+1 de dano) e **obriga** dez 1/1 a atacar contra qualquer tabuleiro, zerando seus bloqueadores e destruindo o próprio X da ativação seguinte. `Goblin Chieftain` (*"Other Goblin creatures you control get +1/+1 and have haste"*) dá **+10 de dano** aos mesmos 10 corpos, torna cada token um 2/2 (sobrevive a ping e a `Goblin Chainwhirler` alheio) e — o ponto decisivo — **dá haste a Krenko**, que é Goblin: Krenko passa a **ativar no turno em que entra**. Resolve as lacunas nº 1 e nº 5 num único card de 3 manas.

### Swap 4 — `Ambitious Assault` → `Shared Animosity`
`Ambitious Assault` {2}{R} · **Instant** — *"Creatures you control get +2/+0 until end of turn. If you control a modified creature, draw a card."*
- **F1** +2/+0 coletivo; compra 1 **se** houver criatura modificada. **F2** — **F3** Instant (cemitério/`Guttersnipe`). **F4** — **F5** Pump coletivo. **F6** T3. **F7** A cláusula de draw é quase morta (0 equipamentos no deck); **não gatilha `Zada`** (sem alvo); mais uma carta no CMC 3 já engarrafado.

| Função dela | Quem cobre depois |
|---|---|
| Pump coletivo | `Shared Animosity`, `Goblin Chieftain`, `Quest for the Goblin Lord` (mantida), `Ferocity of the Wilds` (mantida), `Goblin Surprise`, `Goblin War Party` (ambas fora do meu escopo de corte, seguem no deck). |
| Draw condicional | **Fica descoberta.** Custo aceito e declarado: era draw de 3 manas que só funciona com "modified", pacote que eu medi em `02-theme.md` §5.3 como 3 payoffs / **0 equipamentos**. A fase 3 (draw) cobre a categoria com fontes incondicionais. |
| Instant no cemitério | 19 restantes. |

**Sob 10 Goblins atacando:** `Ambitious Assault` = +2/+0 em cada → **+20 de dano**, uma vez, por 3 manas. `Shared Animosity` (*"Whenever a creature you control attacks, it gets +1/+0 until end of turn for each other attacking creature that shares a creature type with it"*) — como **as 32 criaturas do deck são Goblin e todos os tokens de Krenko são Goblin**, cada um dos 10 recebe **+9/+0**: dez 10/1 = **100 de dano**, e **toda vez que você atacar**, para sempre. Mesmo tabuleiro, mesmo custo de 3 manas: +20 uma vez contra +90 todo turno.

### Swap 5 — `Bravado` → `Purphoros, God of the Forge` ⚠️
`Bravado` {1}{R} · Enchantment — Aura — *"Enchant creature. Enchanted creature gets +1/+1 for each other creature you control."*
- **F1** +1/+1 no portador por cada outra criatura sua. **F2** Não é corpo. **F3** Encantamento. **F4** — **F5** Escala com o enxame; em `Tin Street Dodger` (indesbloqueável) é kill de 2 cartas. **F6** T2. **F7** **2-por-1 contra qualquer remoção** num deck com 0 proteção; morre junto no wipe; **não gatilha `Zada`** (é encantamento, não instant/sorcery).

| Função dela | Quem cobre depois |
|---|---|
| Conversor de largura em dano | `Shared Animosity` (mesma escala, mas em **todas** as criaturas, e sem risco de 2-por-1), `Goblin Chieftain`, `Hobgoblin Bandit Lord`, `Impact Tremors`, `Purphoros`. |
| Combo com `Tin Street Dodger` / voadores | **Fica parcialmente descoberta.** Custo declarado: perde-se a linha "um indesbloqueável gigante". `Mob Mentality` também sai (swap 8), então essa linha some **inteira**. Aceito: ela dependia de 2 cartas específicas e de nenhuma remoção no meio; as entradas fazem mais dano sem depender de conectar. |
| Encantamento (contagem) | Sobram 8 encantamentos; nenhuma carta do deck conta encantamentos. Sem perda. |

**Sob 10 Goblins:** `Bravado` transforma **uma** criatura em ~11/11 → **+10 de dano**, se ela conectar, e é 2-por-1 se responderem. `Purphoros, God of the Forge` (*"Indestructible. As long as your devotion to red is less than five, Purphoros isn't a creature. Whenever another creature you control enters, Purphoros deals 2 damage to each opponent. {2}{R}: Creatures you control get +1/+0 until end of turn."*) converte a ativação de Krenko que **cria** esses 10 tokens em **20 de dano a cada oponente** (60 numa mesa de quatro), sem combate. Além disso: (a) **indestructible** e normalmente **não é criatura** — é o permanente mais resistente a wipe que existe em mono-R, atacando diretamente a queixa "apanho de board wipe"; (b) `{2}{R}: +1/+0` é **mana sink repetível**, atacando a queixa "mana travando" (`02-theme.md` §8.1: o deck não tinha onde gastar mana do T6 em diante); (c) com devoção ≥5 vira um **6/5**.
**⚠️ US$ 28,76 (menor impressão, estimativa Scryfall).** É a única carta cara do pacote. **Se o usuário recusar o preço:** `Impact Tremors` (swap 1) já cobre o eixo "dano por ETB" a 1 por criatura; o que se perde é a indestrutibilidade e o mana sink — e aí a queixa nº 2 fica **inteiramente** para a fase 5.

### Swap 6 — `Tidings of War` → `Rundvelt Hordemaster`
`Tidings of War` {R} · Sorcery — *"Amass Goblins 1. If this spell was cast from a graveyard, amass Goblins 3 instead. Flashback {3}{R}."*
- **F1** Amass Goblins 1 (3 pelo flashback). **F2** — **F3** Sorcery. **F4** O Army recebe os contadores. **F5** — **F6** T1. **F7** **Empilha no mesmo Army das outras 4 cartas de amass: 5 cartas produzem 1 corpo** — anti-sinérgico com um comandante que conta **corpos**.

| Função dela | Quem cobre depois |
|---|---|
| 1 corpo Goblin (Army) | `Goblin-town Flunkies`, `Misty Mountains Raider`, `Assault on Osgiliath` continuam amassando o mesmo Army — **o corpo não se perde**, só deixa de receber 1–3 contadores. |
| **Recursão pós-wipe (flashback)** | **Fica descoberta, e o custo é real:** era 1 das 2,5 fontes de recuperação do deck (`02-theme.md` §4.4). Mitigação: `Purphoros` é indestructible e `Impact Tremors`/`Goblin Bombardment` são encantamentos que **sobrevivem a wipe de criatura** — a resiliência muda de "recuperar corpos" para "não depender de corpos". Se Purphoros for recusado, **esta função precisa ser recuperada na fase 5**. |
| Sorcery no cemitério | 19 restantes. |

**Sob 10 Goblins:** `Tidings of War` acrescenta contadores a um Army que já existe → **+1 a +3 de dano**. `Rundvelt Hordemaster` (*"Other Goblins you control get +1/+1. Whenever this creature or another Goblin you control dies, exile the top card of your library. If it's a Goblin creature card, you may cast that card until the end of your next turn."*) dá **+10 de dano** permanente e, com `Goblin Bombardment` em campo, **cada Goblin sacrificado vira uma carta impulsionada** — é simultaneamente lorde (lacuna 1) e motor de carta (lacuna 3), num corpo Goblin de CMC 2 que ainda melhora a curva.

### Swap 7 — `Akki Ember-Keeper` → `Skirk Prospector`
`Akki Ember-Keeper` {1}{R} · Enchantment Creature — Goblin Warrior 2/1 — *"Whenever a nontoken modified creature you control dies, create a 1/1 colorless Spirit creature token."*
- **F1** Criatura **nontoken modificada** sua morre → **Spirit** 1/1. **F2** Corpo 2/1; bloqueia. **F3** **É Goblin** (soma ao X) **e** encantamento. **F4** Recebe anthems e contadores — com `Goblin Chieftain` vira 3/2. **F5** — **F6** T2. **F7** O gatilho exige *nontoken* **e** *modified* num deck com **0 equipamentos** e ~80% de tokens; o token que ela cria é **Spirit**, que **não** realimenta Krenko.

| Função dela | Quem cobre depois |
|---|---|
| Corpo Goblin CMC 2 | `Skirk Prospector` é Goblin de CMC **1** — a contagem para o X de Krenko **não muda**. |
| Corpo de poder 2 (bloqueio/dano) | **Perda declarada: −1 de poder** (2/1 vira 1/1). Sob as mesmas condições com `Goblin Chieftain` em campo, é 3/2 contra 2/2. Custo aceito: 1 ponto de poder por um sac outlet grátis e 1 mana de curva. |
| Gatilho de Spirit | **Já estava descoberto hoje** — o pacote "modified" tem 3 payoffs e 0 equipamentos. Não é perda nova. |
| Encantamento (contagem) | Ninguém conta encantamentos no deck. |

**Sob 10 Goblins:** `Akki Ember-Keeper` é um 3/2 (com Chieftain) cujo gatilho não dispara. `Skirk Prospector` (*"Sacrifice a Goblin: Add {R}"*) é o **segundo sac outlet grátis**, e converte os 10 corpos em **10 de mana vermelho** — o que (a) dispara `Boggart Shenanigans` 10 vezes, (b) permite recastar Krenko no mesmo turno depois de uma remoção, (c) é resposta a wipe (sacrifica em resposta e usa o mana), (d) é o corpo de **1 mana** que o comandante mais valoriza, pois dobra em toda ativação.
**Nota de coordenação:** o efeito é *mana*, mas a função aqui é **sac outlet/conversor**, não ramp. Se o especialista de ramp também estiver contando `Skirk Prospector`, **não contem duas vezes**.

### Swap 8 — `Mob Mentality` → `Hobgoblin Bandit Lord`
`Mob Mentality` {R} · Enchantment — Aura — *"Enchant creature. Enchanted creature has trample. Whenever all non-Wall creatures you control attack, enchanted creature gets +X/+0 until end of turn, where X is the number of attacking creatures."*
- **F1** Trample no portador; se **todas** as criaturas não-Wall atacarem, +X/+0 (X = atacantes). **F2** — **F3** Encantamento. **F4** — **F5** Concede trample e pump escalável a **um** alvo. **F6** T1. **F7** Exige **ataque total** — que é exatamente o que você deixa de fazer quando o plano vira `Goblin Bombardment`/`Impact Tremors`; 2-por-1 contra remoção; **não gatilha `Zada`**; a única carta que a habilitava confortavelmente era `Goblin Rabblemaster` (ataque compulsório), que sai no swap 3.

| Função dela | Quem cobre depois |
|---|---|
| Trample | `Ferocity of the Wilds` (permanente, coletivo) — **já no deck**. |
| Pump escalável com o enxame | `Shared Animosity` (em **todos**, não em um), `Hobgoblin Bandit Lord`, `Goblin Chieftain`. |
| Combo com `Tin Street Dodger` (indesbloqueável gigante) | **Descoberta** — mesma linha que `Bravado` cobria. Custo declarado e aceito no swap 5. |
| Sinergia com `Goblin Rabblemaster` | O parceiro sai no swap 3. Sem custo residual. |

**Sob 10 Goblins:** `Mob Mentality` dá **+10/+0** a **uma** criatura, e só se **todas** atacarem — deixando você sem bloqueador nenhum. `Hobgoblin Bandit Lord` (*"Other Goblins you control get +1/+1. {R}, {T}: This creature deals damage equal to the number of Goblins that entered the battlefield under your control this turn to any target."*) dá **+10/+0 e +10 de resistência distribuídos por todo o enxame**, e a habilitada converte a ativação de Krenko — que fez **10 Goblins entrarem naquele turno** — em **10 de dano a qualquer alvo, todo turno, sem combate**. Lorde estático (item 1 do mandato) **e** conversor repetível (item 2) na mesma carta.
- **F7 da entrada, declarado:** usa `{T}` (entra na contagem de corpos que disputam o tap, que eu corrigi para 4 em `02-theme.md` §4.8 — passa a 5) e é a terceira carta em CMC 3.

---

## 4. Coleção — o que eu considerei e por que dispensei (regra 7 atualizada)

Regra 7 agora é **prioridade de análise, não obrigação de uso**. Olhei as 34 elegíveis de `00b-colecao-elegivel.md` antes de propor qualquer compra. Nomeando as peças e o motivo:

| Carta da coleção | Função que ela tentaria cobrir | Por que **não** cobre |
|---|---|---|
| **`Haunted Cloak`** (vigilance, trample, haste; equip {1}) | Item 3 do mandato — **haste permanente** | **Meia-solução, e piorou depois que os lordes entraram.** Ela dá haste a **uma** criatura por vez; `Goblin Chieftain` dá a **todos os Goblins, inclusive Krenko**, e ainda soma +1/+1 e um corpo Goblin — pelo mesmo CMC 3, sem os {1} de equipar. O que a Cloak tem de exclusivo é **vigilância** (Krenko ataca e ativa) e **trample**; mas Krenko é um 3/3 que quer ficar vivo, não atacar, e trample deixa de importar quando o dano passa a vir de `Impact Tremors`/`Purphoros`/`Goblin Bombardment`, que **ignoram bloqueio em vez de atravessá-lo**. **Recomendação: não nesta fase.** Ela volta a fazer sentido se — e só se — o orquestrador decidir **completar** o pacote "modified" (`Goro-Goro`, `Ambitious Assault`, `Akki Ember-Keeper`), porque aí ela é o equipamento que faltava. Como este pacote está sendo **desmontado** aqui (swaps 4 e 7), a Cloak fica sem a segunda perna. |
| **`Dragon Throne of Tarkir`** (`{2}`,`{T}`: outras criaturas ganham trample e +X/+X, X = poder do portador) | Item 2 — Overrun repetível | Genuinamente forte, e **quase entrou**. Barrada por três atritos: o portador ganha **defender** e a habilidade usa **`{T}`** — logo **não pode ir em Krenko** (o `{T}` dele é o motor do deck) nem em nenhum dos 4 corpos que já disputam tap; X = poder do portador, e o deck tem **6 corpos com poder ≥3** apenas; e o pacote total é 4 (lançar) + 3 (equipar) + 2 por ativação = **9 manas** para o primeiro Overrun. `Shared Animosity` faz mais dano, por 3 manas, sem tap e sem alvo. **Vale reconsiderar na fase 6** se sobrar slot — é a melhor peça da coleção para o item 2. |
| **`Armory of Iroas`**, **`Leonin Bola`**, **`Dire Flail`**, **`Trailblazer's Torch`** | anthem / conversor | Todas afetam **uma** criatura. Um deck que gera 10–20 corpos por ativação precisa de efeitos **coletivos**; equipamento de alvo único é a curva de valor errada aqui. |
| **`Mercadia's Downfall`** (atacantes +1/+0 por terreno não-básico do defensor) | Item 2 — Overrun | One-shot, e o valor depende da **manabase do oponente**, não da sua. `Shared Animosity` faz o mesmo escalando na **sua** contagem de Goblins, e é permanente. |
| **`Goblin Gathering`** | corpos | 3 manas por 2 tokens; `Hordeling Outburst` (já no deck) faz 3 pelo mesmo custo. Piora o CMC 3. |
| **`Mad Ratter`** | payoff que vira corpos | É Goblin e soma ao X, mas os tokens são **Rats** (não realimentam Krenko) e o gatilho é de **draw**, não de tema — é carta da fase 3, não desta. |

**Conclusão da regra 7:** a coleção tem **0 lordes de Goblin** e **0 conversores coletivos permanentes**. A lacuna nº 1 é, por construção, impossível de fechar com o que o usuário já tem. **A compra é a decisão correta**, e as 8 escolhidas custam US$ 50,04 estimados (US$ 21,28 sem Purphoros).

---

## 5. Decisão pendente — `Zada, Hedron Grinder`: completar ou cortar?

**Não decido — o usuário decide.** `Zada` não está no meu pool de cortes, e eu medi o pacote em **2 habilitadores reais entre 12 candidatas (17%)**. As duas opções, lado a lado, com o custo de cada uma:

### Opção (a) — **completar** o pacote

Para `Zada` virar um motor confiável, o deck precisa de instants/sorceries que (i) mirem **um único** alvo, (ii) sejam bons quando copiados N vezes. Em mono-R, a lista real é curta e barata:

| Carta | Custo | Efeito com `Zada` + 10 Goblins |
|---|---|---|
| `Expedite` | {R} | 11 criaturas ganham haste **e você compra 11 cartas** |
| `Crimson Wisps` | {R} | idem (haste + 11 cartas) |
| `Renegade Tactics` | {R} | 11 cartas + 11 bloqueadores não podem bloquear |
| `Reckless Charge` | {R} | +33/+0 distribuídos + haste, com flashback |
| `Titan's Strength` | {R} | +33/+22 + scry 11 |
| `Assault Strobe` | {R} | **double strike em 11 criaturas** — letal na mesa |

- **Custo em slots: 4 a 6.** Sairia de 2 → 6–8 habilitadores (17% → ~50%).
- **Custo escondido, que é o decisivo:** cada uma dessas cartas é **um blank sem `Zada` em campo**. "Target creature gains haste, draw a card" por {R} é filler. Você estaria pondo 4–6 cartas fracas para tornar **uma** carta forte — num deck cuja queixa nº 1 é *"a lista é larga em cartas de efeito único e baixo impacto"* (leitura do próprio briefing).
- **Custo de curva:** 4–6 cartas de CMC 1, o que **ajuda** o engarrafamento do CMC 3 — é o único argumento estrutural a favor.
- **Custo de consistência:** a linha exige `Zada` viva + 1 habilitador + tabuleiro largo = **três** condições simultâneas, num deck com **0 proteção**.

### Opção (b) — **cortar** `Zada`

O que se perde, e quem cobre (mesmas condições: 10 Goblins em campo, `Goblin Chieftain` incluso):

| Função de `Zada` (F1–F7) | Quem cobre depois do corte |
|---|---|
| **F2/F3** Corpo Goblin **3/3** (1 dos 6 corpos com poder ≥3) e +1 no X de Krenko | `Goblin Chieftain` (2/2 que faz todos os outros +1/+1), `Hobgoblin Bandit Lord` (2/3), `Goblin Trashmaster` se entrar (3/3). Contagem de Goblins **não muda** — é troca 1-por-1. |
| **F5** Motor de cópia | **Fica descoberto.** Mas o que ele produzia — pump massivo e draw em rajada — passa a vir de `Shared Animosity` (dano) e da fase 3 (draw), sem exigir 3 condições simultâneas. |
| **F4** Receptor de anthem/Aura | Qualquer Goblin. |
| Os 2 habilitadores órfãos | **Nenhum vira blank:** `Kick in the Door` continua sendo haste + contador por {R} (eu o classifiquei `núcleo` por conta própria em `02-theme.md`), e `Fists of Flame` continua cantrip + trample. Perdem escala, não função. |
| **F6/F7** Slot de CMC 4 disputando com Krenko | **Ganho, não perda** — libera o turno 4. |

### Minha leitura (não é a decisão)

As entradas desta fase mudam o eixo do deck: com `Impact Tremors`, `Purphoros`, `Goblin Bombardment` e `Shared Animosity`, o dano deixa de precisar de **uma criatura gigante que conecta** e passa a vir de **quantidade de gatilhos**. `Zada` é a melhor carta possível para o plano antigo ("empilhe tudo numa criatura") e uma carta mediana para o novo. A opção (a) custa 4–6 slots de filler para consertar um pacote que ficou **fora do eixo**; a opção (b) custa um corpo 3/3 que é reposto no mesmo custo.

**Se o usuário gosta da carta** (e é uma carta divertida), a opção (a) é jogável — mas então `Expedite` + `Crimson Wisps` + `Assault Strobe` são o mínimo viável (3 slots), e eles têm de sair de outro pool, não do meu.

---

## 6. Resumo de cobertura das lacunas depois destes 8 swaps

| Lacuna (`02-theme.md` §7) | Antes | Depois dos 8 | Depois só do NÚCLEO (5) |
|---|---|---|---|
| **1. Lordes estáticos de Goblin** | **0** | **3** (Chieftain, Hordemaster, Hobgoblin) | **1** (Chieftain) |
| **1b. Conversores permanentes de largura em dano** | 0 repetíveis | **5** (Impact Tremors, Purphoros, Bombardment, Shared Animosity, Hobgoblin) | **4** |
| **2. Proteção / resiliência a wipe** | 0 / 2,5 | **+3 permanentes que sobrevivem a wipe de criatura** (2 encantamentos + Purphoros indestructible) + Bombardment como resposta em pilha | **+3** |
| **3. Draw real** | 5 | **6** (Rundvelt Hordemaster) | 5 — fase 3 |
| **4. Mana sink** | 2 | **4** (Purphoros `{2}{R}`, Skirk Prospector, Bombardment, Hobgoblin) | **3** |
| **5. Haste permanente** | 2 | **3** (+ Goblin Chieftain, que cobre **Krenko**) | **3** |
| **7. Wincons que fecham sozinhos** | 2 | **4** (+ dano por ETB; + sac em massa) | **4** |
| **8. Anti-sinergia do `Goblin Rabblemaster`** | ativa | **removida** | **removida** |

**Caminho morto ressuscitado:** `Boggart Shenanigans` sai de "2 pontos, um deles inerte" para conversor real — `Goblin Bombardment` e `Skirk Prospector` dão a ela o sac outlet grátis que faltava. Cada Goblin sacrificado passa a valer **2 de dano**.
