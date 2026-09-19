# 03 — Draw · Krenko, Mob Boss

> Fase executada **inline pelo orquestrador** em 2026-08-22: o `draw-specialist` caiu por limite
> de sessão antes de escrever. Mesmo protocolo: oracle puxado nesta sessão via `bin/mtgdb`,
> ficha F1–F7 em todo corte, cobertura de funções declarada, simetria de critério.

## Diagnóstico herdado (`02-theme.md` §4.1)

Draw real **5** · meta 12–13 · lacuna **−7**. Sete slots ocupados por loot/cantrip que não geram
carta. Queixa do usuário: *"esvazio a mão no turno 4–5 e passo o resto do jogo em topdeck"*.

## A conclusão que define a fase

Rodei `mtgdb search '"creature you control enters" AND draw' -id R` e
`'"created a token" OR "token you control enters"' -id R`: **um resultado cada**. Mono-vermelho
**não tem** draw que escale nativamente com criaturas. O caminho não é procurar a carta que
compra por Goblin — é **converter corpos em cartas**. Duas peças fazem isso, e as duas são
artefato incolor (logo, também combustível de `Krenko, Baron of Tin Street` e alvo de
`Goblin Engineer`).

## Entradas (7)

| # | Carta | CMC | Marca | Por quê |
|---|---|---|---|---|
| 1 | `Skullclamp` | 1 | **NÚCLEO** | *"Equipped creature gets +1/-1. Whenever equipped creature dies, draw two cards. Equip {1}"* — num token 1/1 é **equip {1} → morre → 2 cartas**. Krenko fabrica o combustível. Sinergia quádrupla: dispara `Boggart Shenanigans` (1 dano), `Rundvelt Hordemaster` (exila topo, Goblin pode ser jogado) e `Krenko, Baron of Tin Street` (é artefato). |
| 2 | `Idol of Oblivion` | 2 | **NÚCLEO** | *"{T}: Draw a card. Activate only if you created a token this turn."* Krenko cria token todo turno em que ativa → **1 carta por turno, sem custo de mana**. Artefato. |
| 3 | `Reckless Impulse` | 2 | **NÚCLEO** | *"Exile the top two cards. Until the end of your next turn, you may play those cards."* Duas cartas reais por 2 manas, janela larga. |
| 4 | `Light Up the Stage` | 3 | **NÚCLEO** | Mesmo efeito, **spectacle {R}**. Com `Impact Tremors` e `Purphoros` em campo, um oponente perde vida em quase todo turno → custa **{R} efetivo**. |
| 5 | `Goblin Engineer` | 2 | compl. | Corpo **Goblin** + tutor de artefato para o cemitério, e `{R},{T}, sac artefato: devolve artefato CMC≤3 do cemitério ao campo`. Busca `Skullclamp` (CMC 1) e o recompra. |
| 6 | `Goblin Matron` | 3 | compl. | Corpo **Goblin** + tutor de **qualquer Goblin card**: acha `Goblin Chieftain`, `Skirk Prospector`, `Rundvelt Hordemaster` ou o segundo Krenko. Consistência é o objetivo nº2 do usuário. |
| 7 | `Outpost Siege` | 4 | compl. | Khans: exila o topo **todo upkeep** e você pode jogar. É a única fonte de draw repetível que **sobrevive a board wipe** (encantamento) — cruza com a queixa nº3. |

## Cortes — ficha F1–F7 e protocolo §2

### `Rummaging Goblin` {2}{R} · Goblin Rogue 1/1 · CMC 3
*"{T}, Discard a card: Draw a card."*
- **F1** Loot repetível. **F2** Corpo 1, **usa `{T}` de verdade**. **F3** Goblin (soma ao X). **F4** Recebe anthems. **F5** — **F6** T3. **F7** `{T}` disputa com atacar e com `Lightning Volley`; **loot = net 0 cartas**.

| Função | Quem cobre |
|---|---|
| Corpo Goblin | `Goblin Engineer` + `Goblin Matron` (2 entradas são Goblin) |
| Seleção repetível | `Idol of Oblivion` (draw **real** repetível, e sem gastar `{T}` de criatura) |
| Descarte-outlet para `Squee` | **Descoberto.** Custo aceito: sobra `Tormenting Voice`. Squee continua se recomprando sozinho. |

### `Grotag Night-Runner` {2}{R} · Goblin Rogue 2/3 · CMC 3
*"Whenever this creature deals combat damage to a player, exile the top card of your library. You may play that card this turn."*
- **F1** Impulso 1/turno **ao conectar**. **F2** Corpo 2. **F3** Goblin. **F4** Anthems. **F6** T3. **F7** Exige conectar dano de combate — e o novo plano do deck **deixa de depender de combate**; a janela é só "this turn" (mais estreita que a das entradas).

| Função | Quem cobre |
|---|---|
| Impulso repetível | `Outpost Siege` (todo upkeep, **sem precisar conectar**), `Idol of Oblivion` |
| Corpo Goblin 2/3 | `Goblin Engineer` 1/2, `Goblin Matron` 1/1 — **perda de estatística declarada** |

### `Goblin Sky Raider` {2}{R} · Goblin Warrior 1/2 · CMC 3
*"Flying"*
- **F1** Só voar. **F2** Corpo 1. **F3** Goblin. **F4** Anthems. **F6** T3. **F7** 3 manas por 1 de dano evasivo.

| Função | Quem cobre |
|---|---|
| Corpo Goblin | as 2 entradas Goblin |
| Evasão | `Impact Tremors`/`Purphoros`/`Goblin Bombardment`/`Hobgoblin Bandit Lord` — o dano passou a **ignorar** o bloqueio em vez de sobrevoá-lo |

### `Gundabad Opportunist` {3}{R} · Goblin Rogue 4/2 · CMC 4
*"When this creature enters, exile the top card. Until the end of your next turn, you may play that card."*
- **F1** Impulso +1, **uma vez**. **F2** Corpo **poder 4** (um dos 6 com poder ≥3). **F3** Goblin. **F4** Anthems. **F6** T4 — disputa o slot com Krenko. **F7** One-shot por 4 manas.

| Função | Quem cobre |
|---|---|
| Impulso de 1 carta | `Reckless Impulse` faz **2** por 2 manas; `Light Up the Stage` faz 2 por `{R}` |
| Corpo poder 4 | **Descoberto.** Custo aceito e declarado: o deck passa a ganhar por quantidade de gatilhos, não por tamanho de corpo. Restam `Goblin Chieftain`, `Hobgoblin Bandit Lord`, `Krenko, Baron of Tin Street`, `Zada` e `Purphoros` (6/5 com devoção ≥5). |

### `Fissure Wizard` {1}{R} · Goblin Wizard 2/1 · CMC 2
*"When this creature enters, you may discard a card. If you do, draw a card."*
- **F1** Loot **uma vez**, ao entrar. **F2** Corpo 2. **F3** Goblin. **F4** Anthems. **F6** T2. **F7** Net 0 cartas.

| Função | Quem cobre |
|---|---|
| Corpo Goblin T2 | `Goblin Engineer` (CMC 2, Goblin) — troca 1-por-1 na curva **e** no tipo |
| Loot | Não precisa de cobertura: **loot nunca foi draw**. É o slot que a fase existe para recuperar. |

### `Innocent Bystander` {1}{R} · Goblin Citizen 2/1 · CMC 2
*"Whenever this creature is dealt 3 or more damage, investigate."*
- **F1** Clue ao levar 3+ de dano. **F2** Corpo 2. **F3** Goblin; o Clue é artefato (combustível do Baron). **F4** Anthems. **F6** T2. **F7** **Nenhuma carta do deck dispara isso de propósito**, e levar 3 num 2/1 é morrer.

| Função | Quem cobre |
|---|---|
| Corpo Goblin T2 | `Goblin Engineer` |
| Gerador incidental de artefato | `Skullclamp`, `Idol of Oblivion` (2 artefatos permanentes, não um Clue que talvez nunca venha) |
| Draw via Clue | `Skullclamp` — mesma ideia, **sem depender de levar dano** |

### `Dragon Mantle` {R} · Aura · CMC 1
*"When this Aura enters, draw a card. Enchanted creature has '{R}: +1/+0'."*
- **F1** Cantrip + mana sink no portador. **F2** — **F3** Encantamento. **F5** Mana sink repetível. **F6** T1. **F7** Cantrip = **net 0**; 2-por-1 contra remoção num deck com 0 proteção.

| Função | Quem cobre |
|---|---|
| Cantrip | `Reckless Impulse`/`Light Up the Stage` (2 cartas, não 0 líquidas) |
| **Único mana sink repetível** (era, na auditoria) | `Umbral Mantle`, `Hammer of Purphoros`, `Purphoros` (`{2}{R}`) e `Castle Embereth` — **a função saiu de 1 para 4 fontes nesta rodada**. Cobertura folgada. |

## Contagem resultante

| | Antes | Depois |
|---|---|---|
| Draw **real** | 5 | **12** — `Conspicuous Snoop`, `Squee`, `Tormenting Voice`, `Rundvelt Hordemaster`, `Reckless Lackey`, `Skullclamp`, `Idol of Oblivion`, `Reckless Impulse`, `Light Up the Stage`, `Goblin Engineer`, `Goblin Matron`, `Outpost Siege` ✅ meta |
| Loot puro (net 0) | 3 | **1** (`Forgotten Cave`) |
| Cantrip (net 0) | 4 | **1** (`Fists of Flame`, que também é habilitador de Zada) |
| Fontes que **escalam com o enxame** | 0 | **3** (`Skullclamp`, `Idol of Oblivion`, `Rundvelt Hordemaster`) |
| Artefatos (combustível do Baron) | 5 | **7** |

## Custo declarado

**Goblins de corpo: −6 / +2 = −4 líquidos.** Cada Goblin a menos é **−1 token por ativação de
Krenko**. Aceito com o seguinte argumento de simetria: os seis que saem somam 6 pontos de poder
em corpos que o deck não usa para atacar, enquanto `Skullclamp` sozinho **transforma tokens em
cartas indefinidamente** — e `Goblin Matron` recompra um Goblin por vez. A queixa que esta fase
atende é mão vazia, não largura de tabuleiro; largura é o que o deck já tinha de sobra.

**Curva:** CMC 3 cai de 24 para 23, CMC 2 sobe para 22. A melhora real de curva vem de
`Urza's Incubator` (13 das cartas de CMC 3 são feitiços de criatura Goblin e passam a custar 1).
