# Análise Temática — Tori D'Avenant, Fury Rider (v1 · 2026-09-18 · modo `improve`)

Fonte de todos os textos oracle: `bin/mtgdb` (dump Scryfall 2026-08-12), puxados nesta sessão.
MCP do Scryfall indisponível. Nenhuma carta foi julgada de memória.

> **Correção de inventário antes de tudo.** O `00-briefing.md` registra "38 terrenos (30 básicos
> + 8 não-básicos)". A contagem real de `lista.txt` é **36 terrenos**: 30 básicos (15 Mountain +
> 15 Plains) + **6** não-básicos (Boros Guildgate, Command Tower, Fields of Strife, Sandstone
> Bridge, Stone Quarry, Wind-Scarred Crag). O "8" do briefing é o número de **nomes distintos de
> terreno** que o `mtgdb` reporta (inclui Mountain e Plains como 2 nomes). Total confirmado:
> 100 cartas (72 entradas, 100 unidades). **O deck está 2 terrenos abaixo da base de 38** —
> insumo para o `manabase-engineer`.

---

## 1. Comandante — análise linha a linha

`Tori D'Avenant, Fury Rider` · `{1}{R}{R}{W}` · 3/3 · Legendary Creature — Human Knight · CMC 4

```
Vigilance, trample
Whenever Tori D'Avenant attacks, all other attacking creatures you control get +1/+1 until
end of turn. Other red attacking creatures you control gain trample until end of turn.
Untap each other white attacking creature you control.
```

| Linha/habilidade | Gatilho/termo | O que exige | O que habilita | O que **não** faz |
|---|---|---|---|---|
| `Vigilance` | estática, própria | — | Tori ataca e fica de pé para bloquear | não estende vigilância a ninguém |
| `Trample` | estática, própria | — | 3 de dano passa por chump-block | corpo 3/3 — trample num 3/3 é quase decorativo |
| `Whenever Tori attacks` | **attack-trigger** | Tori **precisa atacar** todo turno | é o motor inteiro da carta | não dispara em combate extra? dispara sim, mas exige Tori atacando de novo |
| `all other attacking creatures get +1/+1` | **anthem temporário em massa** | N outros atacantes | +N de dano no combate | não é permanente, não é contador, não sobrevive ao turno, **não conta Tori** |
| `Other red attacking creatures gain trample` | **gives-trample**, só vermelhas | criaturas **vermelhas** atacando | fura chump-block das vermelhas | **não dá trample às brancas** — e o deck é 21 branco × 11 vermelho |
| `Untap each other white attacking creature` | **pseudo-vigilância**, só brancas | criaturas **brancas** atacando | ficam de pé para bloquear o crack-back; destravam custos de `{T}` | **não remove de combate** (o dano ainda é dado), **não gera combate extra**, **não gera mana**, **não compra nada** |

### O que o deck atual realmente extrai do gatilho

Contagem feita sobre as 27 criaturas não-comandante de `lista.txt`:

| Público-alvo da cláusula | Quantas criaturas | % do deck |
|---|---|---|
| Recebem `+1/+1` (qualquer atacante) | 27 | 100% |
| Recebem `untap` (**brancas**) | **21** | 78% |
| Recebem `trample` (**vermelhas**) | **11** | 41% |
| Recebem as **três** (RW) | **5** — Adriana, Fireborn Knight, Honored Crop-Captain, Inspiring Veteran, Jor Kadeen | 19% |
| Só vermelhas (trample, sem untap) | 6 — Fervent Cathar, Relentless Rohirrim, Rohirrim Lancer, Syr Carah, Éomer of the Riddermark, Éomer Marshal | 22% |

**Quanto isso vale por combate real.** Com um board típico de Tori + 4 outros atacantes:
- `+1/+1` → **+4 de dano** naquele combate (não conta Tori);
- `trample` → relevante em ~1,6 dos 4 atacantes (41% são vermelhos), e só quando o corpo supera o bloqueador;
- `untap` → ~3,1 dos 4 atacantes ficam de pé. **Isso é defesa, não dano.**

Ou seja: **o gatilho da Tori é, em dano, um `Honored Crop-Captain` (custo `{R}{W}`, +1/+0 aos
outros atacantes) num corpo de 4 manas**, acrescido de pseudo-vigilância e de um trample que só
alcança 41% do board. A parcela *ofensiva* do texto é a mais fraca das três e a parcela *branca*
(78% do deck) é puramente defensiva.

### Os dois lugares onde o `untap` poderia virar vantagem real — e o que o deck tem hoje

1. **Custos de `{T}` pós-combate.** Só duas cartas no deck aproveitam: `Sanctuary Lockdown`
   (`{2}`, tap two untapped Humans: tap target creature) e `Syr Carah, the Bold` (`{T}`: 1 dano).
   Duas cartas em 99.
2. **Combates extras** (atacar de novo com o mesmo board). Só uma carta no deck: `Éomer, Marshal
   of Rohan`, e o gatilho dele exige que **uma legendária atacante sua morra** — condição que o
   deck não controla e não quer. **Zero combates extras confiáveis.** Este é o buraco mais
   gritante: a única cláusula diferenciada da Tori (untap) não tem pagamento no deck.
3. Ela **não** tem proteção, **não** tem hexproof, **não** tem haste e **não** tem recursão. Num
   pod competitivo ela é removida no primeiro ataque e o imposto de comandante sobe a cada volta.

### Teto competitivo — honestidade obrigatória

- **Pseudo-vigilância é a habilidade menos valiosa num deck de ataque em massa.** Quem ataca em
  massa não quer bloquear; quer atacar de novo. A cláusula seria excelente com *extra combat* ou
  com *untapper payoffs* — nenhum dos dois está no deck, e nem seria a linha mais eficiente de RW.
- **O `+1/+1` é temporário.** Não empilha entre turnos, não vira contador, não sobrevive a wipe,
  não muda a matemática de bloqueio no turno do oponente.
- **O trample está na cor errada.** 78% das criaturas do deck são brancas e nenhuma delas recebe
  trample da Tori.
- **O gatilho exige Tori atacando** — ou seja, exige que o pior alvo de remoção da mesa se exponha
  todo turno, sem proteção embutida.

Veredito preliminar (detalhado na §8): **Tori é um comandante de mesa casual competente e um
comandante competitivo medíocre.** Ela não gera cartas, não gera corpos, não gera mana e não
fecha o jogo — as quatro coisas que separam um comandante de topo num teto de R$200.

---

## 2. Termos de busca do tema

Derivados do oracle, na ordem de força:

1. `attack-trigger` / `attacking-matters` — o deck inteiro gira em torno de "quando ataca".
2. `token` / `repeatable-creature-tokens` — o `+1/+1` escala linearmente com número de atacantes.
3. `anthem` / `keyword-anthem` — empilha com o `+1/+1` temporário.
4. `extra-combat-phase` — **o pagamento que falta** para a cláusula de untap.
5. `Knight` (typal) — 21 Knights no deck; `Inspiring Veteran` e `The Circle of Loyalty` já leem.
6. `Human` (typal) — **25 das 27 criaturas são Humanas** (exceções: Belladonna Took = Halfling
   Citizen; Paladin Danse = Synth Knight). `Sanctuary Lockdown` e `Vigilante Justice` têm base real.
7. `combat damage → draw / impulse` — a cura da dor 2 sem sair do tema.
8. `indestructible` / `mass protection` / `recursão de board` — a cura da dor 4.

Buscas efetivamente rodadas: `mtgdb tag attacking-matters -id RW`, `tag extra-combat-phase -id RW`,
`search '"deals combat damage to a player" AND draw' -id RW`, `search '"attacks" AND "draw a card"'
-id RW -cmc-max 5`, mais confirmação oracle carta a carta.

---

## 3. Auditoria por categoria — as 99 cartas

Rótulos do `CLAUDE.md`: `tema`, `draw`, `ramp`, `remoção`, `proteção`, `counter`, `wipe`, `wincon`,
`terreno`. Acrescento `recursão` como sub-rótulo de `proteção` (resposta a wipe).

### 3.1 Criaturas (27)

| Carta | CMC | Tipo | Categorias | Sinergia com o gatilho da Tori |
|---|---|---|---|---|
| Adriana, Captain of the Guard | 5 | Leg. Cre. Human Knight 4/4 | tema, wincon | RW: recebe as 3 cláusulas · dá **melee** ao time (escala com nº de oponentes atacados) |
| Belladonna Took | 2 | Leg. Cre. Halfling Citizen 2/2 | tema, draw | branca (untap) · 2º token do turno = comprar; deck faz poucos tokens |
| Cavalry Drillmaster | 2 | Cre. Human Knight 2/1 | tema | branca (untap) · ETB dá +2/+0 e first strike (liga com Kwende) |
| Dawnstrike Vanguard | 6 | Cre. Human Knight 4/5 | tema | branca (untap) · **anti-sinergia**: o gatilho quer *criaturas tapadas* no end step, e a Tori destapa as brancas |
| Elite Interceptor // Rejoinder | 1 | Cre. Human Wizard 1/2 | tema, draw | branca (untap) · Rejoinder destapa/tapa + compra |
| Esgaroth Garrison | 5 | Cre. Human Soldier */5 | tema, draw | branca (untap) · poder = nº de criaturas (escala com go-wide) · recruit = loot + token |
| Fervent Cathar | 3 | Cre. Human Knight 2/1 | tema | vermelha (trample) · haste · ETB tira um bloqueador |
| Fireborn Knight | 4 | Cre. Human Knight 2/3 | tema | **RW: recebe as 3** · double strike dobra o `+1/+1` da Tori |
| Honored Crop-Captain | 2 | Cre. Human Warrior 3/2 | tema | **RW: recebe as 3** · +1/+0 aos outros atacantes empilha com a Tori |
| Inspiring Captain | 4 | Cre. Human Knight 3/3 | tema | branca (untap) · ETB +1/+1 no time |
| Inspiring Veteran | 2 | Cre. Human Knight 2/2 | tema | **RW: recebe as 3** · lorde de Knight (21 Knights no deck) |
| Jor Kadeen, the Prevailer | 5 | Leg. Cre. Human Warrior 5/4 | tema, wincon | **RW: recebe as 3** · metalcraft +3/+0 no time (8 artefatos no deck) |
| Knight Luminary | 4 | Cre. Human Knight 3/2 | tema | branca (untap) · ETB token branco (também destapável) · warp `{1}{W}` |
| Knight of Sorrows | 5 | Cre. Human Knight 3/3 | tema | branca (untap) · **defensiva** (bloqueia 2, afterlife) num deck de ataque |
| Kwende, Pride of Femeref | 4 | Leg. Cre. Human Knight 2/2 | tema | branca (untap) · first strike → double strike (lê Cavalry Drillmaster, Warlord's Fury, Syr Alin, Squire's Lightblade) |
| Lake-town Lookout | 1 | Cre. Human Scout 1/1 | tema, draw | branca (untap) · recruit ao morrer |
| Luxknight Breacher | 4 | Cre. Human Knight 2/2 | tema | branca (untap) · contadores = nº de criaturas+artefatos |
| Paladin Danse, Steel Maverick | 3 | Leg. Art. Cre. Synth Knight 3/3 | tema, proteção | branca (untap) · **artefato** (metalcraft do Jor Kadeen) · exile: indestrutível a artefatos/Humanos = 1 resposta a wipe |
| Parhelion Patrol | 4 | Cre. Human Knight 2/3 | tema | branca · **já tem vigilance** (untap da Tori é redundante nela) · mentor + voar |
| Recruitment Officer | 1 | Cre. Human Soldier 2/1 | tema, draw | branca (untap) · `{3}{W}`: cava 4 por criatura ≤3 — **vantagem de carta real e repetível** |
| Relentless Rohirrim | 4 | Cre. Human Knight 4/3 | tema | vermelha (trample) · ETB só "the Ring tempts you" |
| Rohirrim Lancer | 1 | Cre. Human Knight 1/1 | tema | vermelha (trample) · menace |
| Sheriff of Safe Passage | 3 | Cre. Human Knight 0/0 | tema | branca (untap) · contadores = nº de criaturas · plot `{1}{W}` (esquiva de wipe pela mão) |
| Syr Alin, the Lion's Claw | 5 | Leg. Cre. Human Knight 4/4 | tema | branca (untap) · attack-trigger +1/+1 ao time **empilha com a Tori** |
| Syr Carah, the Bold | 5 | Leg. Cre. Human Knight 3/3 | tema, draw | vermelha (trample) · **impulse ao causar dano** · `{T}`: 1 dano — um dos 2 usos reais do untap… mas ela é vermelha, **não é destapada pela Tori** |
| Éomer of the Riddermark | 5 | Leg. Cre. Human Knight 5/4 | tema | vermelha (trample) · haste · token branco ao atacar se tiver o maior poder |
| Éomer, Marshal of Rohan | 4 | Leg. Cre. Human Knight 4/4 | tema | vermelha (trample) · **único combate extra do deck**, condicionado à morte de legendária atacante |

> **Atrito estrutural descoberto aqui:** `Syr Carah` tem o único `{T}` ofensivo do deck e é
> **vermelha** — a Tori destapa só as brancas. `Sanctuary Lockdown` precisa de Humanos
> **destapados** e a Tori entrega isso; mas o efeito é tapar 1 criatura por `{2}`.
> `Dawnstrike Vanguard` quer criaturas **tapadas** no end step e a Tori as destapa.
> A única cláusula diferenciada do comandante está mal servida pelo próprio deck.

### 3.2 Artefatos (8) e Encantamentos (5)

| Carta | CMC | Tipo | Categorias | Sinergias |
|---|---|---|---|---|
| Sol Ring | 1 | Artifact | **ramp**, tema(metalcraft) | intocável · conta para Jor Kadeen e Luxknight |
| Arcane Signet | 2 | Artifact | **ramp**, tema(metalcraft) | intocável |
| Ancestral Blade | 2 | Artifact — Equipment | tema | token branco 1/1 (destapável) + equip |
| Squire's Lightblade | 1 | Artifact — Equipment | tema | flash, +1/+0, first strike → liga com Kwende |
| True-Faith Censer | 2 | Artifact — Equipment | tema | +1/+1 + **vigilance** (redundante com o untap da Tori) · +1/+0 extra em Humano (25 Humanos) |
| Trailblazer's Torch | 4 | Artifact — Equipment | tema, remoção parcial | **initiative** (Undercity) · 2 de dano a cada bloqueador |
| The Circle of Loyalty | 6 | Leg. Artifact | tema, wincon parcial | affinity for Knights (21 Knights → costuma custar 2–3) · anthem +1/+1 permanente · token 2/2 por legendária conjurada (deck tem 10 legendárias) + `{3}{W},{T}` |
| Paladin Danse | 3 | Art. Creature | (ver criaturas) | 8º artefato |
| Sanctuary Lockdown | 3 | Enchantment | tema, remoção parcial | anthem de Humano (**25 Humanos**) · `{2}` + tapar 2 Humanos destapados: tapa 1 criatura — **consome o untap da Tori** |
| Valor in Akros | 4 | Enchantment | tema | +1/+1 ao time por criatura que entra (tokens contam) |
| Vigilante Justice | 4 | Enchantment | tema, wincon parcial | 1 dano por Humano que entra (25 Humanos + tokens Human Soldier) |
| Seal of Cleansing | 2 | Enchantment | **remoção** (art/ench) | sacrifica-se; deixa o board |
| Basri Ket | 3 | Planeswalker | tema, proteção | +1: contador + indestrutível · −2: token por atacante não-token |

### 3.3 Instantâneos (13)

| Carta | CMC | Categorias | Leitura honesta |
|---|---|---|---|
| Adamant Will | 2 | proteção | +2/+2 e indestrutível **a 1 criatura**. Não responde a wipe. |
| Disenchant | 2 | remoção (art/ench **apenas**) | morta contra deck de criaturas |
| Djeru's Renunciation | 2 | tema (tap 2), draw(cycling) | tapar 2 = pseudo-Falter; cycling `{W}` é a válvula |
| Duty Beyond Death | 2 | **proteção em massa**, tema | indestrutível ao time + contador em cada — **custo: sacrificar uma criatura**. Uma das 2 respostas reais a wipe |
| Feat of Resistance | 2 | proteção | contador + proteção de cor a 1 criatura (também fura bloqueio) |
| Gideon's Triumph | 2 | remoção (edict) | **o oponente escolhe** — sacrifica a pior criatura dele |
| Gods Willing | 1 | proteção | proteção de cor a 1 criatura + scry 1 |
| Invoke the Divine | 3 | remoção (art/ench **apenas**) | pior que Disenchant em 1 mana |
| Squire's Lightblade | — | (artefato) | — |
| Swift Reckoning | 2 | **remoção** (criatura tapada) | boa no turno do oponente; ruim contra vigilance/untapped |
| Warlord's Fury | 1 | tema, draw(cantrip) | first strike ao time (liga com Kwende) + compra 1 |
| Zealous Display | 3 | tema | +2/+0 ao time; o untap só vale **fora do seu turno** |
| Adamant Will / Feat / Gods Willing | — | — | três efeitos de **alvo único** ocupando 3 slots |

### 3.4 Feitiços (11) + recursão

| Carta | CMC | Categorias | Leitura |
|---|---|---|---|
| Basri's Solidarity | 2 | tema | +1/+1 contador em cada criatura, one-shot, sorcery |
| Bond of Discipline | 5 | tema (Falter) | tapa todas as criaturas **deles** + lifelink — 5 manas por um ataque limpo |
| Celebrate the Mountain-king | 4 | **remoção**, draw | exila 1 nãoterreno **por oponente** enquanto ficar em campo + recruit (loot+token). Melhor carta de interação do deck |
| Crash Through | 1 | tema, draw(cantrip) | trample ao time — **cobre exatamente o buraco branco da Tori** + compra 1 |
| Expose to Daylight | 3 | remoção (art/ench **apenas**) | + scry 1 |
| Inspiring Roar | 4 | tema | **idêntico a Basri's Solidarity por 2 manas a mais** |
| Late to Dinner | 4 | **recursão** | 1 criatura do cemitério ao campo + Food |
| Miraculous Recovery | 5 | **recursão** | 1 criatura + contador |
| Pride of Conquerors | 2 | tema, wincon parcial | +2/+2 ao time com ascend (10 permanentes é trivial aqui) |
| Raise the Alarm | 2 | tema | 2 tokens brancos (destapáveis pela Tori) |
| Reduce to Memory | 3 | **remoção** (qualquer nãoterreno) | dá um 3/2 RW ao dono — a remoção mais ampla do deck |
| Remember the Fallen | 3 | recursão (à mão) | criatura e/ou artefato de volta à **mão** |
| Knight Watch / etc | — | — | não está no deck |

### 3.5 Terrenos (36)

| Carta | Qtd | Categorias | Observação |
|---|---|---|---|
| Mountain / Plains | 15 / 15 | terreno | 30 básicos |
| Command Tower | 1 | terreno | intocável |
| Boros Guildgate | 1 | terreno | entra tapado, sem upside |
| Stone Quarry | 1 | terreno | entra tapado, sem upside |
| Wind-Scarred Crag | 1 | terreno | entra tapado, 1 de vida |
| Fields of Strife | 1 | terreno | entra tapado · `{2}{R}{W},{T}`: surveil 1 (custo proibitivo) |
| Sandstone Bridge | 1 | terreno, tema | entra tapado · ETB: +1/+1 e **vigilance** a uma criatura (redundante com Tori) |

**4 dos 6 não-básicos entram tapados e não fazem nada relevante.** Nenhum terreno de utilidade
(saque, remoção, token, haste, unblockable). Zero fontes de mana além de terreno + Sol Ring +
Arcane Signet.

---

## 4. Contagem por categoria × metas do pipeline

| Categoria | Meta (`CLAUDE.md`) | Deck atual | Diferença | Leitura |
|---|---|---|---|---|
| **draw** (vantagem real) | **12–13** | **4** — Recruitment Officer, Syr Carah, Belladonna Took (condicional), Rejoinder (marginal) | **−8 a −9** | Cantrips (Crash Through, Warlord's Fury, Shoulder to Shoulder) e loots (3× recruit, Djeru's cycling) são **neutros em cartas**, não contam |
| **ramp padrão** | **10–11** | **2** — Sol Ring, Arcane Signet | **−8 a −9** | nenhum rock além dos dois, nenhum land ramp, nenhum Treasure |
| **ramp explosivo** | **2–3** | **0** | **−2 a −3** | — |
| **interação pontual** | **~10** | **4 reais** — Reduce to Memory, Swift Reckoning, Gideon's Triumph, Celebrate the Mountain-king | **−6** | **+4 cartas art/ench-only** (Disenchant, Expose to Daylight, Invoke the Divine, Seal of Cleansing) que são mortas contra metade das mesas |
| **board wipe** | **2–4** | **0** | **−2 a −4** | deck sem nenhum sweeper |
| **counter** | (não é meta fixa em RW) | 0 | — | RW não tem counters; a substituição correta é proteção + velocidade |
| **proteção (alvo único)** | — | 4 — Adamant Will, Feat of Resistance, Gods Willing, Basri Ket +1 | excesso | 4 slots gastos salvando **uma** criatura |
| **proteção em massa (anti-wipe)** | 3–4 num deck de criaturas | **2** — Duty Beyond Death (custa sacrificar), Paladin Danse (exílio, 1 uso, só artefatos/Humanos) | **−2** | dor 4 |
| **recursão pós-wipe** | — | 3 lentas — Late to Dinner (4), Miraculous Recovery (5), Remember the Fallen (3) | qualitativa | cada uma devolve **1** criatura; wipe tira 8 |
| **wincon** | **3+** | **0 fechadores reais** | **−3** | Jor Kadeen / Adriana / Pride of Conquerors são *aumentos de dano*, não fechadores. Zero dano não-combate, zero combate extra confiável, zero alt-win |
| **terrenos** | **38** (base) | **36** | **−2** | e 4 dos 6 não-básicos entram tapados sem upside |
| **tema** | — | 43 | ok em quantidade, **fraco em qualidade** | maioria é pump temporário de Limited |

**Diagnóstico numérico:** o deck tem ~43 cartas de tema e ~10 cartas de função. A proporção
saudável para "extremamente competitivo" é próxima do inverso: ~25–28 de tema e ~35 de função
(draw/ramp/interação/proteção), porque o tema é sustentado **pelo comandante e por 8–10 payoffs
densos**, não por 43 cartas de pump.

---

## 5. Cartas sem sinergia sobreposta — candidatas a corte (ficha F1–F7 + protocolo §2)

Regra 3: mínimo 2 pontos de sinergia. Todas as fichas abaixo assumem **as mesmas condições** que
uso para defender entradas (§3 do checklist): board de 4–5 atacantes, Tori atacando, os anthems
do deck (Inspiring Veteran, Circle of Loyalty, Sanctuary Lockdown) em campo.

---

### C1 · Inspiring Roar — `{3}{W}` Sorcery
- **F1** Põe um contador +1/+1 em cada criatura que você controla. Uma linha, uma vez.
- **F2** Sem corpo. Não tapa, não sacrifica, não bloqueia.
- **F3** Tipo Sorcery — o deck não conta feitiços (sem prowess, sem magecraft, sem spell mastery
  relevante além de Swift Reckoning, que só quer 2 cards no cemitério).
- **F4** Não recebe nada (anthems não afetam feitiço).
- **F5** Dá contadores permanentes ao time — **a única função real**, e ela *escala com nº de
  criaturas*, portanto tem 1 ponto de sinergia com go-wide.
- **F6** Turno 4+. Não destrava nada a partir dali.
- **F7** **Atrito direto:** `Basri's Solidarity` (`{1}{W}`) faz **exatamente o mesmo efeito por 2
  manas a menos**. Duas cópias funcionais do mesmo card, uma estritamente pior.
- **Protocolo de corte:** função única (contadores em massa) → **coberta integralmente por
  Basri's Solidarity**, que fica. Nenhuma função descoberta. **Corte limpo.**

### C2 · Shoulder to Shoulder — `{2}{W}` Sorcery
- **F1** Support 2 (contador em até **duas** criaturas alvo) + compre 1 carta.
- **F2/F3/F4** Sem corpo, sem tipo relevante, não recebe nada.
- **F5** 2 contadores permanentes + reposição de carta. Card-neutra (gasta 1, compra 1).
- **F6** Turno 3.
- **F7** Compete em slot com `Basri's Solidarity` (contador em **todas**, 1 mana a menos) e não
  gera vantagem de carta (cantrip ≠ draw).
- **Protocolo:** funções = [contador, cantrip]. Contador → coberto por Basri's Solidarity e
  Valor in Akros. Cantrip → **descoberto**, mas é justamente o que a fase de draw vai substituir
  por saque real (custo aceito e declarado). **Corte limpo.**

### C3 · Adamant Will — `{1}{W}` Instant
- **F1** Uma criatura alvo ganha +2/+2 e indestrutível até o fim do turno.
- **F2/F3** Sem corpo, sem tipo relevante.
- **F4** Não recebe nada.
- **F5** Protege **uma** criatura; salva de remoção pontual e ganha combates.
- **F6** Turno 2 (interação barata — ponto a favor).
- **F7** **Atrito de slot severo:** o deck tem **quatro** efeitos de proteção de alvo único
  (Adamant Will, Feat of Resistance, Gods Willing, Basri Ket +1) e **nenhum** deles responde ao
  problema que o usuário relatou (board wipe). Proteger 1 de 8 criaturas contra um Wrath é
  perder 7.
- **Protocolo:** funções = [proteção alvo único, combat trick]. Proteção → coberta por Gods
  Willing (1 mana, proteção de cor, também fura bloqueio) e Feat of Resistance (deixa contador).
  Combat trick → coberto por Feat of Resistance. Nenhuma função descoberta. **Corte limpo.**
  *Observação:* se a fase de interação trouxer proteção em massa (Unbreakable Formation / Boros
  Charm / Flawless Maneuver), o corte fica ainda mais confortável.

### C4 · Warlord's Fury — `{R}` Instant
- **F1** Criaturas que você controla ganham first strike até o fim do turno. Compre 1 carta.
- **F2/F3/F4** Sem corpo, sem tipo relevante, não recebe nada.
- **F5** First strike em massa → **liga com Kwende, Pride of Femeref** (first strike vira double
  strike) e com o `+1/+1` da Tori (dano dobrado nos que têm double strike). Isso são **2 pontos
  de sinergia reais** — a carta **passa** na regra 3.
- **F6** Turno 1, instant, cantrip.
- **F7** Depende de Kwende em campo para ser mais que "meu time ganha combates que já ganharia".
- **Veredito:** **NÃO é corte por falta de sinergia.** É uma carta fraca em potência absoluta,
  mas tem 2 pontos e é card-neutra por 1 mana. Marco como **corte condicionado**: só sai se a
  fase de draw/interação precisar do slot e Kwende também sair. **Decisão do orquestrador.**

### C5 · Crash Through — `{R}` Sorcery
- **F1** Criaturas que você controla ganham trample até o fim do turno. Compre 1 carta.
- **F5** **Cobre exatamente o buraco de cor da Tori**: ela dá trample só às vermelhas (41%);
  Crash Through dá às 78% brancas. Somado ao `+1/+1` dela, é o pacote de "alfa strike passa".
  2 pontos de sinergia. **Passa na regra 3.**
- **F7** Sorcery (não pode ser usada em resposta). Efeito nulo se o board estiver travado.
- **Veredito:** **corte condicionado**, não corte limpo. Se entrar um trample-anthem permanente
  (ex.: Rally the Ranks não dá trample; a opção real é manter), ela sai. Sozinha, tem 2 pontos.
  *Nota:* **há uma cópia sobressalente na coleção** — cortar aqui não gera economia nem custo.

### C6 · Basri's Solidarity — `{1}{W}` Sorcery
- **F1** Contador +1/+1 em cada criatura que você controla.
- **F5** Contadores permanentes → **sobrevivem ao board wipe de −X/−X parcial**, empilham com
  anthems, e escalam com go-wide. Com 6 criaturas é +6/+6 distribuídos por 2 manas.
- **F7** One-shot, sorcery.
- **Veredito:** **fica** (é o original de que Inspiring Roar é a cópia cara). 2 pontos: go-wide +
  permanência dos contadores.

### C7 · Djeru's Renunciation — `{1}{W}` Instant · Cycling `{W}`
- **F1** Tape até duas criaturas alvo. Cycling `{W}`.
- **F2/F3/F4** Sem corpo, sem tipo relevante.
- **F5** Tapar 2 bloqueadores antes do ataque = pseudo-Falter parcial; **e** tapar 2 atacantes
  deles no turno do oponente. O cycling garante que nunca é carta morta.
- **F6** Turno 1 (como cycling) ou 2.
- **F7** Compete com `Bond of Discipline` (mesma função, escala total, 5 manas) e é **muito** pior
  que remoção real.
- **Protocolo:** funções = [tap 2, cycling]. Tap → coberto por Bond of Discipline e Sanctuary
  Lockdown; se ambos saírem, fica **descoberto** — mas "tapar criaturas" não é função que o deck
  precise preservar (é o efeito mais fraco entre os disponíveis no slot de interação). Cycling →
  descoberto, substituído por saque real. **Corte limpo, com custo declarado.**

### C8 · Bond of Discipline — `{4}{W}` Sorcery
- **F1** Tape todas as criaturas dos oponentes. Suas criaturas ganham lifelink até o fim do turno.
- **F5** Falter total + lifelink em massa = um turno de alfa strike sem bloqueio e ganho de vida
  proporcional. **2 pontos** (go-wide + alfa).
- **F6** Turno 5, sorcery, pré-combate.
- **F7** **5 manas que não adicionam nada ao board**; num pod de 3 oponentes, num deck sem
  fechador, é "ganhei uma rodada de dano" e não "ganhei o jogo". Compete em custo com um wipe
  (que o deck não tem) e com ramp.
- **Protocolo:** funções = [Falter em massa, lifelink em massa]. Falter → parcialmente coberto por
  Fervent Cathar e Djeru's Renunciation; se ambos saírem, **descoberto** (custo aceito: a função
  "alfa garantido" é substituída por **combate extra** e **evasão permanente**, que a fase de
  wincon cobre melhor). Lifelink → descoberto, dispensável. **Corte limpo com custo declarado.**

### C9 · Gideon's Triumph — `{1}{W}` Instant
- **F1** O **oponente alvo sacrifica uma criatura à escolha dele** que atacou ou bloqueou neste
  turno. Dobra se você controlar um planeswalker Gideon.
- **F2/F3/F4** Sem corpo, sem tipo relevante, não recebe nada.
- **F5** Remoção — mas **de escolha do oponente**, e apenas entre criaturas que atacaram/bloquearam.
- **F6** Turno 2.
- **F7** **O deck não tem nenhum planeswalker Gideon** (tem Basri Ket) — metade do texto é letra
  morta. Como edict, ela remove a **pior** criatura do oponente, nunca a ameaça.
- **Protocolo:** função única = [remoção condicional de baixa qualidade] → **coberta e melhorada**
  por qualquer remoção-alvo do pool (Swords to Plowshares / Generous Gift / Glass Casket, esta
  última **já na coleção**). Nenhuma função descoberta. **Corte limpo.**

### C10 · Knight of Sorrows — `{4}{W}` 3/3
- **F1** Pode bloquear uma criatura adicional a cada combate. Afterlife 1 (token 1/1 WB voador
  ao morrer).
- **F2** Corpo 3/3 por 5 — **bloqueador**, e o token de afterlife é um corpo extra que ataca.
- **F3** Human Knight → conta para os 21 Knights (Inspiring Veteran, Circle of Loyalty affinity)
  e os 25 Humanos (Sanctuary Lockdown, Vigilante Justice, True-Faith Censer).
- **F4** Recebe untap da Tori (branca), o `+1/+1`, os anthems.
- **F5** Não dá nada a ninguém.
- **F6** Turno 5 — pior ponto da curva do deck (9 cartas em CMC 5).
- **F7** **Anti-tema declarado:** "bloqueia uma criatura adicional" é uma habilidade defensiva num
  deck cuja tese é atacar em massa; o afterlife premia a morte dela, não o ataque.
- **Protocolo:** funções = [corpo 3/3, contagem Knight, contagem Human, bloqueio duplo, token ao
  morrer]. Corpo/contagens → cobertas por qualquer substituta Knight/Human de custo menor (o pool
  tem Hero of Bladehold, Adeline, Benalish Marshal, Knight Exemplar — todas Humanas/Knights).
  Bloqueio duplo → **descoberto** (custo aceito: o deck não quer bloquear). Token ao morrer →
  descoberto (custo aceito: 1/1 voador irrelevante). **Corte limpo com custos declarados.**

### C11 · Dawnstrike Vanguard — `{5}{W}` 4/5
- **F1** Lifelink. No início do seu end step, **se você controlar duas ou mais criaturas tapadas**,
  ponha um +1/+1 em cada criatura **exceto ela**.
- **F2** Corpo 4/5 lifelink por 6.
- **F3** Human Knight (contagens de Knight/Human — sim, alimenta).
- **F4** Recebe untap (branca), `+1/+1`, anthems, contadores.
- **F5** Distribui contadores permanentes ao time — **efeito forte de fato**, e permanente
  (sobrevive a wipe de −X/−X e a fim de turno).
- **F6** **Turno 6.** Em deck com 2 fontes de ramp, isso é turno 7–8 real.
- **F7** **Anti-sinergia direta com o comandante:** o gatilho exige **criaturas tapadas** no end
  step; a Tori **destapa** todas as brancas atacantes (78% do board). Atacar com o comandante
  ativamente *desliga* esta carta. Também é o único CMC 6 não-Circle do deck.
- **Protocolo:** funções = [corpo 4/5 lifelink, Knight/Human, contadores em massa recorrentes].
  Corpo → coberto por qualquer 4-drop do pool. Contagens → cobertas. Contadores recorrentes →
  **cobertos por Cathars' Crusade** (5 manas, dispara por criatura que entra, sem exigir board
  tapado) se ela entrar; **descobertos** se não. **Corte condicionado** à entrada de Cathars'
  Crusade ou equivalente. Mesmo sem ela, a anti-sinergia com o comandante é argumento suficiente
  registrado por escrito.

### C12 · Miraculous Recovery — `{4}{W}` Instant · e · Late to Dinner — `{3}{W}` Sorcery
- **F1** (Miraculous) Devolve 1 criatura do cemitério ao campo com um contador.
  (Late to Dinner) Devolve 1 criatura do cemitério ao campo + cria um Food.
- **F2/F3** Sem corpo. Late to Dinner gera **um artefato Food** → conta para metalcraft do
  Jor Kadeen e para Luxknight Breacher. Esse é um ponto real que a leitura de memória perderia.
- **F5** Recursão — **a resposta nominal à dor 4**.
- **F6** Turnos 4 e 5.
- **F7** **Cada uma devolve UMA criatura.** Um board wipe leva 6–8. Recuperar 1 criatura por 4–5
  manas depois de perder 8 não é resiliência, é lentidão. Ambas exigem que a criatura **já esteja
  no cemitério** — não previnem nada.
- **Protocolo:** funções = [recursão de criatura, (Late) artefato para metalcraft].
  Recursão → **a função é legítima e a dor 4 do usuário a exige**; a substituição correta é
  *prevenção* (Unbreakable Formation, Boros Charm, Flawless Maneuver, Make a Stand) + *recursão em
  massa* (Faith's Reward, Brought Back), ambas mais baratas e mais amplas. Metalcraft → coberto se
  entrar mais um artefato barato. **Corte condicionado**: só corta se a fase de interação/proteção
  entregar a prevenção em massa. Se não entregar, **estas duas são o pouco que o deck tem** e não
  podem sair — nomeio a função que ficaria descoberta: **recuperação pós-wipe**.

### C13 · Remember the Fallen — `{2}{W}` Sorcery
- **F1** Escolha um ou ambos: devolva 1 criatura **à mão**; devolva 1 artefato **à mão**.
- **F2/F3/F4** Sem corpo.
- **F5** Recursão à mão (2-for-1 potencial se pegar criatura + artefato).
- **F6** Turno 3.
- **F7** À **mão**, não ao campo — depois de um wipe, ainda é preciso pagar o custo de novo. Num
  deck com 2 fontes de ramp, recomprar e reconjurar uma criatura de 5 é uma volta inteira perdida.
- **Protocolo:** função = [recursão lenta]. Coberta e superada por Faith's Reward / Brought Back
  (retornam **ao campo**, múltiplos permanentes, em instant). **Corte limpo condicionado à
  entrada de uma delas.**

### C14 · Squire's Lightblade — `{W}` Artifact — Equipment
- **F1** Flash. ETB: anexa a uma criatura sua, ela ganha first strike **até o fim do turno**.
  Equipada recebe +1/+0. Equip `{3}`.
- **F2** Sem corpo.
- **F3** **Artefato** → metalcraft do Jor Kadeen (8 artefatos), contagem do Luxknight Breacher.
  **Ponto real.**
- **F5** +1/+0 permanente; first strike só no turno em que entra → liga com **Kwende** (double
  strike). 2 pontos.
- **F6** Turno 1, flash (pode entrar como truque de combate).
- **F7** **Equip `{3}` para +1/+0** é um dos piores rates do formato; depois do primeiro turno a
  carta é um +1/+0 imóvel.
- **Protocolo:** funções = [artefato para metalcraft, +1/+0, first strike pontual, flash].
  Metalcraft → coberto se entrar outro artefato barato (Idol of Oblivion, Mask of Memory,
  Skullclamp — todos melhores). First strike → coberto por Warlord's Fury / Cavalry Drillmaster.
  +1/+0 → dispensável. **Corte limpo condicionado** a manter a contagem de artefatos ≥3.

### C15 · Disenchant / Expose to Daylight / Invoke the Divine / Seal of Cleansing (4 cartas)
- **F1** Todas destroem **artefato ou encantamento** e nada mais. Expose adiciona scry 1; Invoke
  adiciona 4 de vida; Seal fica em campo e sacrifica-se (instant-speed a custo zero depois).
- **F2** Sem corpo. Seal of Cleansing é **encantamento em campo** — ocupa a mesa e é sacrificável.
- **F3** Seal alimenta contagens de encantamento (o deck **não tem** constelação — não alimenta nada).
- **F5** Respostas a Rhystic Study, Smothering Tithe, equipamentos, prisões.
- **F6** Turnos 2/3/3/2.
- **F7** **Quatro slots** com exatamente a **mesma** função e o mesmo ponto cego: **não tocam em
  criatura, em planeswalker, em terreno nem em battle.** Num pod competitivo, metade das ameaças
  que matam você são criaturas.
- **Protocolo:** função = [remoção de artefato/encantamento]. **Manter 2**, não 4 — e trocar por
  versões modais que também pegam criatura (**Generous Gift**, **Beast Within** fora de cor,
  **Anguished Unmaking** fora de cor; em RW a resposta é **Boros Charm**/**Swords to Plowshares**
  para criatura + manter 2 dos 4 art/ench). Recomendo manter **Seal of Cleansing** (custo afundado,
  responde em instant sem mana) e **Expose to Daylight** (scry) ou trocar ambos por
  **Glass Casket** (já na coleção, pega criatura) + 1 art/ench. **Corte limpo de 2, funções
  cobertas pelos 2 remanescentes.**

### C16 · Fields of Strife — Land
- **F1** Entra tapado. `{T}`: `{R}` ou `{W}`. `{2}{R}{W}, {T}`: surveil 1.
- **F7** Surveil 1 por **4 manas + o tap do terreno** é inutilizável. Como dual, é estritamente
  um Boros Guildgate. **Corte condicionado** — decisão do `manabase-engineer`; a função
  (fonte RW) precisa ser reposta, e o deck já está 2 terrenos abaixo da base.

---

### Resumo do ranking de cortes (do mais limpo ao mais condicionado)

| # | Carta | Tipo de corte | Função que exige reposição |
|---|---|---|---|
| 1 | **Inspiring Roar** | limpo | nenhuma (Basri's Solidarity cobre) |
| 2 | **Gideon's Triumph** | limpo | nenhuma |
| 3 | **Shoulder to Shoulder** | limpo | cantrip → saque real |
| 4 | **Adamant Will** | limpo | nenhuma (3 proteções de alvo único ficam) |
| 5 | **Knight of Sorrows** | limpo | bloqueio duplo (custo aceito) |
| 6 | **Djeru's Renunciation** | limpo | tap 2 (custo aceito) |
| 7 | **Dawnstrike Vanguard** | limpo (anti-sinergia com o comandante) | contadores recorrentes → Cathars' Crusade |
| 8 | **Bond of Discipline** | limpo | Falter em massa → combate extra/evasão |
| 9 | **Squire's Lightblade** | limpo | 1 artefato para metalcraft |
| 10 | **Invoke the Divine** + **Disenchant** (2 dos 4 art/ench) | limpo | nenhuma (2 remanescentes cobrem) |
| 11 | **Remember the Fallen** | condicionado | recursão → Faith's Reward / Brought Back |
| 12 | **Miraculous Recovery** ou **Late to Dinner** (uma das duas) | condicionado | recuperação pós-wipe |
| 13 | **Fields of Strife** | condicionado (manabase) | fonte RW |
| 14 | **Warlord's Fury**, **Crash Through** | **corte condicionado — NÃO corto** | ambas têm 2 pontos de sinergia; Crash Through é a única fonte de trample para as 78% brancas. Devolvo ao orquestrador |

**Fora do meu escopo de corte** (exercem função de outra especialidade — devolvo ao orquestrador):
`Recruitment Officer` e `Syr Carah` (draw), `Reduce to Memory`, `Swift Reckoning`, `Celebrate the
Mountain-king` (interação — são 3 das 4 respostas reais do deck), `Duty Beyond Death` e
`Paladin Danse` (as 2 únicas proteções em massa), `Sol Ring`/`Arcane Signet`/`Command Tower`
(intocáveis declarados).

---

## 6. Lacunas — o que falta para cada uma das quatro dores

### Dor 1 — "Não fecha o jogo" → categoria **wincon** (0/3) e **extra-combat** (0)
O deck causa dano linear e nunca dá o passo não-linear. Faltam, em ordem de impacto:
1. **Combate extra** — é o pagamento que a cláusula de untap da Tori pede e o deck não tem
   (`Aurelia, the Warleader`, `Aggravated Assault`, `Breath of Fury`, `Seize the Day`,
   `Waves of Aggression`, `Combat Celebrant`).
2. **Dano não-combate escalando com o board** — `Impact Tremors`, `Hellrider`, `Vigilante Justice`
   (já tem), `Chandra's Ignition`, `Goblin Bombardment`.
3. **Multiplicador de dano por tipo** — `Shared Animosity` lê **21 Knights** e é o maior salto de
   dano disponível no tema.
4. **Evasão em massa** — `Iroas, God of Victory` (menace + previne dano aos atacantes: resolve
   dor 1 e dor 4 numa carta).

### Dor 2 — "Mão morta" → categoria **draw** (4/12–13)
O deck confunde cantrip com saque. Faltam **motores** que leem o próprio tema:
- dano de combate → carta (`Mask of Memory`, `Sword of Fire and Ice`, `Bill the Pony`… );
- criatura barata entra → carta (`Tocasia's Welcome`, `Welcoming Vampire`, `Mentor of the Meek`,
  `Bygone Bishop`) — **o deck tem 18 cartas em CMC ≤2**, esses gatilhos disparam muito;
- token entra → carta (`Idol of Oblivion`, `Skullclamp` com tokens 1/1);
- impulse recorrente (`Outpost Siege`, `Light Up the Stage`, `Neyali` se virar comandante).

### Dor 3 — "Sem respostas" → categoria **remoção** (4 reais /10) e **wipe** (0/2–4)
- 4 dos 8 slots de "interação" só pegam artefato/encantamento;
- **zero** remoção que responda a uma criatura grande em instant sem condição
  (`Swift Reckoning` exige que ela esteja tapada);
- **zero** board wipe. `Fumigate` **está na coleção** e é a correção mais barata possível
  — mas num deck de criaturas, o wipe certo é **assimétrico** ou vem com proteção
  (`Fumigate` + `Unbreakable Formation`/`Boros Charm` é a linha clássica de RW).

### Dor 4 — "Morre para board wipe" → categoria **proteção em massa** (2 frágeis) + **recursão** (3 lentas)
O deck responde a wipe *depois* do wipe, uma criatura por vez. O que falta:
- **prevenção instantânea e barata**: `Boros Charm` (`{R}{W}`, indestrutível a **todos os seus
  permanentes**), `Unbreakable Formation`, `Flawless Maneuver` (grátis com comandante em campo),
  `Make a Stand`, `Clever Concealment` (convoke — pagável com o board que está atacando);
- **recursão em massa**: `Faith's Reward`, `Brought Back`;
- **ameaças que não são criaturas**: `Assemble the Legion` e `Mobilization` reconstroem o board
  sozinhos depois de um wipe — resiliência sem ocupar o board.
- **Nota do briefing respeitada:** a resposta é resiliência, **não** reduzir o número de criaturas.

---

## 7. Varredura da coleção (regra 7) — 113 sobressalentes

Varridas **antes** de qualquer busca nova. Da caixa, servem a este deck:

### Aproveitáveis — entram na análise das fases seguintes

| Carta | CMC | Tipo | Função neste deck | Sinergias (2+) | Preço conhecido |
|---|---|---|---|---|---|
| **Fumigate** | 5 | Sorcery | **wipe** (0→1) | cura dor 3 · liga com Boros Charm/Unbreakable Formation (wipe assimétrico) | R$ 1,37 |
| **Glass Casket** | 2 | Artifact | **remoção** de criatura ≤3 | 1º remoção-de-criatura real e barata · **artefato** → metalcraft Jor Kadeen + Luxknight | R$ 0,09 |
| **Blast Zone** | 0 | Land | **wipe** modular + terreno | responde a tokens/prisões · cobre parte do −2 de terrenos | R$ 0,89 |
| **Ratchet Bomb** | 2 | Artifact | wipe modular | artefato (metalcraft) · limpa tokens dos oponentes sem matar seus 3-drops | R$ 1,79 |
| **Boros Locket** | 3 | Artifact | **ramp** + **draw** (sac: 2 cartas) | cobre 2 lacunas · artefato (metalcraft) | — |
| **Commander's Sphere** | 3 | Artifact | **ramp** + **draw** (sac: 1 carta) | idem · artefato | — |
| **Sphere of the Suns** | 2 | Artifact | **ramp** T2 | artefato · acelera os 15 quatro-drops | R$ 0,11 |
| **Prophetic Prism** | 2 | Artifact | ramp/fixação + **cantrip** | artefato · fixa `{R}{R}{W}` da Tori | R$ 0,13 |
| **Manalith** | 3 | Artifact | ramp | artefato | — |
| **Seer's Lantern** | 3 | Artifact | ramp + scry | artefato | — |
| **Magnifying Glass** | 3 | Artifact | ramp + draw lento | artefato · Clue é artefato extra | — |
| **Hedron Crawler** | 2 | Art. Creature 0/1 | ramp | **corpo** (conta para Esgaroth Garrison, Luxknight, Sheriff, Adriana melee) + **artefato** (metalcraft) — dois eixos simultâneos | — |
| **Pilgrim's Eye** | 3 | Art. Creature 1/1 | ramp (busca terreno) | corpo + artefato + terreno garantido | — |
| **Knight Watch** | 5 | Sorcery | tema (2× 2/2 Knight **vigilance**) | 2 corpos brancos (untap Tori) · **Knights** (Inspiring Veteran, Circle of Loyalty) | — |
| **Youthful Knight** | 2 | Human Knight 2/1 | tema | **Knight + Human** (3 contagens) · first strike → **Kwende** = double strike | — |
| **Embereth Paladin** | 4 | Human Knight 4/1 | tema | Knight + Human · **vermelha** (trample da Tori) · haste | — |
| **Bogardan Lancer** | 2 | Human Knight 1/1 | tema (marginal) | Knight + Human · vermelha · flanking | — |
| **Joust** | 2 | Sorcery | **remoção** por fight | **+2/+1 se for Knight** (21 Knights) · remoção que o deck não tem | — |
| **S.H.I.E.L.D. Spy Kit** | 1 | Equipment | draw marginal + tema | artefato (metalcraft) · untap + scry ao atacar sozinho — **anti-sinergia com go-wide** (só ataque solo) | — |
| **Leonin Bola** | 1 | Equipment | remoção parcial (tapa) | artefato · **consome o untap da Tori** (tap da criatura equipada para tapar a deles) — 2 pontos reais | — |
| **Ori, Keeper of Songs** | 3 | Leg. Cre. Dwarf Bard 3/3 | tema (marginal) | branca (untap) · legendária → token do Circle of Loyalty · storied fácil (8 artefatos + 10 legendárias) | — |
| **Lightning Strike / Magma Spray / Flame Slash / Seal of Fire / Chandra's Pyrohelix / Twin Bolt / Smite the Deathless / Bombard / Fateful End** | 1–3 | Instant/Sorcery | **remoção** barata | cobrem a lacuna de remoção-de-criatura a custo ~zero. Seal of Fire é encantamento em campo (resposta sem mana) | — |
| **Skullcrack** | 2 | Instant | anti-lifegain + 3 dano | protege o plano de dano contra lifegain/fog — relevante em mesa competitiva | — |
| **Phyrexian Revoker** | 2 | Art. Cre. 2/1 | stax leve | corpo + **artefato** (metalcraft) · desliga Sol Ring/tutores dos oponentes | R$ 1,25 |
| **Sol Ring** (2ª cópia) | 1 | Artifact | — | **já há uma no deck**; a sobressalente não tem uso aqui | — |
| **Crash Through**, **Warlord's Fury** (cópias) | 1 | — | — | **duplicatas do que já está no deck** — sem uso adicional | — |

### Dispensadas — com motivo por escrito (regra 7)

| Carta | Por que não serve |
|---|---|
| Bident of Thassa, Reconnaissance Mission, Mad Ratter, Thrummingbird, Somber Hoverguard, Negate, Stoic Rebuttal, Thirst for Knowledge, Disruption Protocol, Experimental Augury, Master's Councillors, Elrond Moon-Reader, Cargo Ship, Diversion Unit, Thopter Fabricator, Old Fat Spider Can't See Me, Thranduil's Decree | **Fora da identidade de cor** (azul). Reconnaissance Mission e Bident seriam saque perfeito para o tema — é a perda mais dolorosa da caixa |
| Ancient Animus, Courage in Crisis, Smell Fear, Beorn's Hospitality, Beorn Reluctant Host, Return to Nature, Troll Negotiations, Ulvenwald Mysteries, Wilderland Scrounger, Grumgully, Zhur-Taa Goblin, Invigorating Hot Spring, Road // Ruin, Moss Diamond | **Fora da identidade** (verde) |
| Leonin Abunas | Dá hexproof a **artefatos**; o deck tem 8 artefatos e nenhum é a ameaça — protege a coisa errada. Corpo 2/5 defensivo num deck de ataque. **1 ponto de sinergia** |
| Titan Forge, Lux Cannon, Lux Artillery, Darksteel Reactor, Oracle's Vault, Culling Dais, Steel Wrecking Ball, Moxite Refinery, Melded Moxite, Myr Convert, Zookeeper Mechan, Mycosynth Wellspring, Liquimetal Coating, Thaumaton Torpedo, Wedgelight Rammer, Patrolling Peacemaker, Lesser Masticore | Pacote de artefatos/charge counters do deck Inspirit. Contam para metalcraft (1 ponto), mas **não atacam, não compram e não protegem** — nenhuma chega a 2 pontos aqui |
| Breeches, Eager Pillager | Excelente carta, mas o gatilho lê **Pirate**; o deck tem 0 Pirates. 1 ponto (corpo 3/3 first strike vermelho) |
| Redcap Thief, Goblin Gathering, Dragon Mantle, Shiny Impetus, Ragged Short Spear, Angelic Gift, Eagle's Rescue, Moment of Glory, Emerge from the Cocoon, Pinecone Strike, Ambitious Assault, Smaug's Fury, Mercadia's Downfall, Seize Opportunity, Lightning Volley, Molten Blast, Seismic Wave, Smashing Success, Stonefury, Searing Barrage, Chandra's Outrage, Radiating Lightning, Destructive Tampering, Punishing Fire, Panic Spellbomb | Filler de Limited ou remoção de custo alto/eficiência baixa. Substituíveis pelas queimas de 1–2 manas **da própria caixa** listadas acima |
| Smaug, the Great Calamity | 7 manas numa curva que já trava em 4–5; identidade R ok, mas **corte de curva** |
| Bargaining Table, Smuggler's-tier artifacts do Inspirit | idem acima |

**Economia identificada na caixa:** ramp (5–7 peças), wipe (Fumigate + Blast Zone + Ratchet Bomb),
remoção-de-criatura (Glass Casket + 6–8 queimas) — ou seja, **as três lacunas mais caras do deck
saem parcialmente da caixa sem compra**. O que **não** sai da caixa é **saque real** e
**proteção em massa** — e é aí que o orçamento de compra deve ir.

---

## 8. Pool temático RW competitivo (42 candidatas)

Todas dentro de `id<=RW`, `legal:commander`, com **2+ pontos de sinergia** explicados.
Preço: só onde há cotação registrada (`mtgdb prices`) — **a régua LigaMagic é do orquestrador**.

### 8.1 Motores de corpo — o tema que escala (11)

| Carta | CMC | Tipo | Sinergias (2+) | Coleção? |
|---|---|---|---|---|
| **Hero of Bladehold** | 4 | Cre. Human Knight 3/4 | 2 tokens **tapados e atacando** por ataque → alvos do `+1/+1` e do untap da Tori · **battle cry** empilha com o gatilho dela · **Knight + Human** (3 contagens) | não |
| **Adeline, Resplendent Cathar** | 3 | Leg. Cre. Human Knight */4 | 1 token **por oponente** a cada ataque (3 corpos/turno em pod) · poder = nº de criaturas · vigilance · **Knight + Human** · legendária → token do Circle of Loyalty | não |
| **Myrel, Shield of Argive** | 4 | Leg. Cre. Human Soldier 3/4 | X tokens Soldier ao atacar · **stax no seu turno** (oponentes não conjuram nem ativam) = **protege contra wipe em resposta e contra remoção da Tori** · Human · legendária | não |
| **Assemble the Legion** | 5 | Enchantment RW | tokens **com haste** crescendo todo upkeep · **não é criatura** → sobrevive a wipe e reconstrói sozinha (dor 4) · tokens vermelho-e-brancos = recebem as 3 cláusulas da Tori | não |
| **Mobilization** | 3 | Enchantment | `{2}{W}`: token Soldier à vontade = mana em corpos (dor 4: reconstrói pós-wipe) · dá **vigilance** a Soldiers · sobrevive a wipe | não |
| **Hanweir Garrison** | 3 | Cre. Human Soldier 2/3 | 2 tokens tapados e atacando por ataque · **Human** · vermelha (trample da Tori) · **combo com Breath of Fury** | não |
| **Legion Warboss** | 3 | Cre. Goblin Soldier 2/2 | token com haste que **ataca obrigatoriamente** todo combate · mentor (contador em atacante menor) · vermelha | não |
| **Brimaz, King of Oreskos** | 3 | Leg. Cre. Cat Soldier 3/4 | token ao atacar **e** ao bloquear · vigilance · branca (untap) · legendária → Circle of Loyalty | não |
| **Anim Pakal, Thousandth Moon** | 3 | Leg. Cre. Human Soldier RW | X tokens tapados e atacando, **X cresce todo turno** · RW (recebe as 3 cláusulas) · Human · legendária | não |
| **Knight Watch** | 5 | Sorcery | 2× Knight 2/2 **vigilance** · Knights (Inspiring Veteran +1/+1, Circle affinity) | **SIM** |
| **Secure the Wastes** | X+1 | Instant | X corpos brancos em **instant** (pós-wipe, fim do turno do oponente) · brancos = untap da Tori · escala com o ramp | não |

### 8.2 Multiplicadores de dano / wincons (9)

| Carta | CMC | Tipo | Sinergias (2+) | Coleção? | Preço conhecido |
|---|---|---|---|---|---|
| **Shared Animosity** | 3 | Enchantment R | **21 Knights** + 25 Humanos: cada atacante ganha +1/+0 por atacante do mesmo tipo → com 5 Knights atacando são **+4/+0 em cada um** · empilha com o `+1/+1` da Tori · não é criatura (sobrevive a wipe) | não | R$ 25,99 |
| **Aurelia, the Warleader** | 6 | Leg. Cre. Angel RW | **destapa todas as criaturas + combate extra** — é literalmente o pagamento que falta para o untap da Tori · voa, vigilance, haste · legendária | não | — |
| **Aggravated Assault** | 3 | Enchantment R | combate extra repetível (destapa suas criaturas) · não é criatura · liga com o ramp que a fase 4 vai trazer | não | — |
| **Breath of Fury** | 4 | Aura R | **combate extra por criatura que conecta**, sacrificando a encantada → com Hanweir Garrison/Assemble the Legion vira loop de combates = **o fechador real do deck** | não | — |
| **Hellrider** | 4 | Cre. Devil 3/3 R | 1 dano ao jogador **por criatura que ataca** (não depende de conectar → ignora bloqueio e fog) · haste · vermelha (trample) | não | **R$ 3,49** |
| **Impact Tremors** | 2 | Enchantment R | 1 dano a cada oponente por criatura que entra — com Hero/Adeline/Assemble são 3–6 por turno · não é criatura | não | R$ 17,78 |
| **Iroas, God of Victory** | 4 | Leg. Ench. Cre. RW 7/4 | **menace no time inteiro** (evasão em massa = dor 1) · **previne todo dano aos seus atacantes** (dor 4 em combate) · indestrutível · legendária | não | — |
| **Cathars' Crusade** | 5 | Enchantment W | contador +1/+1 **em cada criatura** por criatura que entra → com token-makers o board explode **permanentemente** (sobrevive a wipe de −X/−X e a fim de turno) · não é criatura | não | — |
| **Chandra's Ignition** | 5 | Sorcery R | wipe assimétrico + dano a cada oponente, usando o maior corpo (Jor Kadeen 5/4, Éomer 5/4, Esgaroth */5) · **é wipe e wincon na mesma carta** | não | — |

### 8.3 Saque que lê o tema (8)

| Carta | CMC | Tipo | Sinergias (2+) | Coleção? | Preço |
|---|---|---|---|---|---|
| **Tocasia's Welcome** | 3 | Enchantment W | compra por turno com criatura ≤3 MV entrando — **o deck tem 18 cartas ≤2 MV** e tokens não disparam, mas as criaturas sim · não é criatura (sobrevive a wipe) | não | — |
| **Welcoming Vampire** | 3 | Cre. Vampire 2/3 W | compra por turno com criatura de poder ≤2 entrando — **tokens 1/1 disparam** · voa · branca (untap) | não | — |
| **Mentor of the Meek** | 3 | Cre. Human Soldier 2/2 | `{1}` por criatura poder ≤2 que entra = saque **sem limite por turno** com token-makers · **Human + Soldier** · branca | não | — |
| **Bygone Bishop** | 3 | Cre. Spirit Cleric 2/3 W | Clue por criatura ≤3 MV conjurada · **Clue = artefato** (metalcraft Jor Kadeen) · voa · branca | não | — |
| **Mask of Memory** | 2 | Equipment | 2 cartas por conexão de combate (descarta 1) = **saque ligado ao ataque** · **artefato** (metalcraft) · equip `{1}` | não | — |
| **Idol of Oblivion** | 2 | Artifact | `{T}`: compre 1 se criou token neste turno — com Hero/Adeline/Assemble é saque todo turno · **artefato** (metalcraft) · sobrevive a wipe | não | R$ 8,94 |
| **Skullclamp** | 1 | Equipment | equipar um token 1/1 → ele morre → **2 cartas**; o deck faz tokens 1/1 · artefato · equip `{1}` | não | R$ 26,89 |
| **Outpost Siege** | 4 | Enchantment R | impulse todo upkeep (vantagem recorrente) · não é criatura (sobrevive a wipe) · modo Dragons pinga quando criatura sai (inclui wipe) | não | **R$ 0,90** |

### 8.4 Resiliência a wipe — a dor 4 (7)

| Carta | CMC | Tipo | Sinergias (2+) | Coleção? |
|---|---|---|---|---|
| **Boros Charm** | 2 | Instant RW | **indestrutível a TODOS os seus permanentes** (não só criaturas) por 2 manas · ou 4 de dano à cara (fecha jogo) · ou double strike (dobra o `+1/+1` da Tori) — **três funções numa carta** | não |
| **Unbreakable Formation** | 3 | Instant W | indestrutível ao time · **addendum na main: contador +1/+1 em cada + vigilance** = pump permanente proativo · dupla função | não |
| **Flawless Maneuver** | 3 | Instant W | **grátis** enquanto a Tori estiver em campo — e o deck quer a Tori em campo de qualquer jeito · indestrutível ao time | não |
| **Make a Stand** | 3 | Instant W | indestrutível + `+1/+0` ao time → combate ganho **e** wipe sobrevivido | não |
| **Clever Concealment** | 4 | Instant W | **convoke** — pago com o board que já está tapado de atacar · hexproof + indestrutível ao time | não |
| **Faith's Reward** | 4 | Instant W | devolve **todos** os permanentes que foram ao cemitério **neste turno** — é a resposta em massa que Late to Dinner/Miraculous Recovery não são | não |
| **Brought Back** | 2 | Instant W | 2 permanentes de volta ao **campo** por 2 manas (inclui terrenos destruídos) | não |

### 8.5 Interação e utilidade que também lê o tema (7)

| Carta | CMC | Tipo | Sinergias (2+) | Coleção? |
|---|---|---|---|---|
| **Glass Casket** | 2 | Artifact | remoção de criatura ≤3 MV · **artefato** (metalcraft Jor Kadeen, contagem Luxknight) | **SIM** |
| **Joust** | 2 | Sorcery R | remoção por fight **+2/+1 se o seu for Knight** (21 Knights) · vermelha | **SIM** |
| **Fumigate** | 5 | Sorcery W | o wipe que falta · **combina com Boros Charm/Unbreakable Formation** para virar assimétrico | **SIM** |
| **Blast Zone** | 0 | Land | wipe modular + **terreno** (deck está −2 na base de 38) | **SIM** |
| **Sunforger** | 3 | Equipment RW | +4/+0 no atacante · **tutor de instantâneo R/W ≤4 grátis** a cada combate (busca Boros Charm, Joust, remoção) · artefato (metalcraft) | não |
| **Legion's Landing // Adanto** | 1 | Leg. Ench. → Land | token lifelink T1 · **transforma quando você ataca com 3+ criaturas** (o deck faz isso) → vira **terreno que gera tokens** (resiliência a wipe + cobre o −2 de terrenos) | não |
| **Sokenzan, Crucible of Defiance** / **Eiganjo, Seat of the Empire** | 0 | Leg. Land | terreno que vira 2 corpos com haste / 4 de dano a atacante-bloqueador — **10 legendárias no deck barateiam o channel** · resolvem o −2 de terrenos sem custo de slot | não |

**Total do pool: 42 candidatas.** Já na coleção: **5** (Knight Watch, Glass Casket, Joust,
Fumigate, Blast Zone) + o pacote de ramp e queima da §7 que não repeti aqui.

---

## 9. Veredito sobre a Tori

### Com o deck consertado (draw 12–13, ramp 10–11, interação 10 + 2–4 wipes, resiliência), ela sustenta "extremamente competitivo no Commander 200"?

**Não.** E a razão é o texto dela, não a qualidade do deck em volta:

1. **Ela não gera recursos de nenhum tipo.** Não compra, não faz corpo, não faz mana, não tutora.
   Num teto de R$200, o comandante é a única carta que você joga toda partida — desperdiçar essa
   consistência num anthem temporário custa o deck inteiro. Os quatro pontos de dor do usuário
   (fecha, compra, responde, resiste) são exatamente as quatro coisas que o texto dela **não** faz.
2. **A cláusula que a diferencia (untap de brancas) é a mais fraca do formato para o arquétipo.**
   Ela produz pseudo-vigilância — valor **defensivo** — num deck cuja tese é atacar. Só vira valor
   ofensivo com combate extra, e nesse caso o comandante correto é **quem dá o combate extra**
   (Aurelia), não quem destapa.
3. **A divisão de cores é auto-contraditória.** 78% das criaturas são brancas e recebem só o
   untap; 41% são vermelhas e recebem só o trample. O texto premia um deck RW **equilibrado em
   cor**, que é mais difícil de montar e mais frágil de manabase — e ainda assim só 19% do board
   receberia as três cláusulas.
4. **Ela é obrigada a atacar, sem proteção, sendo o alvo mais óbvio da mesa.** CMC 4, 3/3, sem
   hexproof, sem ward, sem haste, sem recursão. Numa mesa competitiva ela é removida no primeiro
   ataque e o imposto de comandante a torna cara demais para re-conjurar.
5. **Comparação direta de rendimento por ataque, mesmo board (5 atacantes, anthems em campo):**

| Comandante | O que o gatilho de ataque entrega | Dano incremental | Recurso permanente |
|---|---|---|---|
| **Tori D'Avenant** | +1/+1 nos outros 4, trample em ~1,6, untap em ~3,1 | **+4** | **0** |
| Adeline, Resplendent Cathar | 3 tokens tapados e atacando (pod de 3) + poder próprio = nº de criaturas | **+3 + o próprio poder crescente** | **3 corpos por turno** |
| Hero of Bladehold | battle cry (+1/+0 nos outros 4) + 2 tokens atacando | **+6** | **2 corpos por turno** |
| Neyali, Suns' Vanguard | double strike nos tokens + impulse ao atacar com token | **dobra os tokens** | **1 carta por turno** |
| Winota, Joiner of Forces | Humano do topo 6 direto ao campo, tapado e atacando, indestrutível, **por atacante não-Humano** | explosivo | **corpo + tutor por gatilho** |

Tori é a única linha da tabela com **zero** na coluna de recurso permanente.

### 3–5 comandantes RW de tema de ataque que superam o teto dela

**Não escolho — a decisão é do usuário.** Ordenados por quão pouco o deck precisaria mudar:

| Comandante | CMC | Por que supera a Tori | Custo de reconstrução | Alerta |
|---|---|---|---|---|
| **Adeline, Resplendent Cathar** | 3 `{1}{W}{W}` | Gera **1 corpo por oponente a cada ataque** — resolve dor 1 (mais dano), dor 4 (reconstrói board) e alimenta todos os anthems/contagens do deck. Vigilance própria. Poder escala com o board. **Human Knight** → as 21 contagens de Knight e 25 de Humano do deck continuam valendo | **Mínimo** — o deck atual é 78% branco e 21 Knights; quase nada muda | É mono-branca: perde as 11 vermelhas? Não — a identidade do deck vira só W, então Fervent Cathar, Éomer×2, Relentless Rohirrim, Rohirrim Lancer, Syr Carah, Hellrider, Shared Animosity **saem** |
| **Isshin, Two Heavens as One** | 3 `{R}{W}{B}` | Dobra **todos** os gatilhos de ataque do deck (Tori, Adriana, Syr Alin, Hero of Bladehold, Adeline, Hellrider, Legion Warboss) | — | **FORA DE ESCOPO: identidade Mardu (RWB).** Registrado só para o usuário saber que existe e por que não é opção |
| **Winota, Joiner of Forces** | 4 `{2}{R}{W}` | **25 das 27 criaturas do deck são Humanas** — Winota precisa do inverso (não-Humanos atacando para colocar Humanos em campo). O teto de potência dela é o mais alto de RW | **Alto** — exige um pacote de não-Humanos/tokens não-Humanos | Poder cEDH, mas reconstrói ~15 slots e atrai ódio de mesa |
| **Neyali, Suns' Vanguard** | 4 `{2}{R}{W}` | **Double strike em todos os tokens atacantes** (dobra o dano do go-wide) + **impulse toda vez que um token ataca** = resolve dor 1 **e** dor 2 no mesmo texto | Médio — exige virar o deck para tokens (Hero, Adeline, Assemble, Mobilization, Anim Pakal — tudo já no pool §8.1) | Depende de fazer tokens; sem token-makers ela é um 3/3 vanilla |
| **Aurelia, the Warleader** | 6 `{2}{R}{R}{W}{W}` | **Destapa todas as criaturas + combate extra** ao atacar — é a Tori com o pagamento embutido. Voa, vigilance, haste | Baixo — o deck já é o deck dela | CMC 6 com 2 fontes de ramp é tarde demais **hoje**; com o ramp corrigido para 10–11, é jogável no T5 |
| **Anim Pakal, Thousandth Moon** | 3 `{1}{R}{W}` | Tokens **tapados e atacando** em quantidade **crescente** todo turno, e ela guarda os contadores (sobrevive melhor). RW puro: recebe as 3 cláusulas se a Tori virar carta do 99 | Baixo | Frágil (1/2); morre fácil antes de acumular |

**Recomendação de encaminhamento (não é escolha):** se o usuário quiser **manter o deck como está
e só consertá-lo**, Tori fica e a rodada é sobre as lacunas da §4/§6 — o deck sobe de "casual
médio" para "casual forte", não para "extremamente competitivo". Se o objetivo declarado
("extremamente competitivo") for a régua, o menor custo de reconstrução com maior ganho de teto é
**Adeline** (mínimo, mono-W) ou **Aurelia** (baixo, mantém RW e o tema inteiro) — e a Tori
continua no deck como carta do 99, onde o gatilho dela custa 4 manas e não custa o comandante.
