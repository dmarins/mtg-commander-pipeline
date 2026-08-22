# 02 — Análise Temática · Krenko, Mob Boss

- **Modo:** `improve` (auditoria da lista de entrada)
- **Data:** 2026-08-22
- **Fonte de todo texto oracle:** `bin/mtgdb deck krenko-mob-boss -full` + `bin/mtgdb oracle` + `bin/mtgdb rulings`, puxados nesta sessão. Nenhum veredito de memória.
- **Escopo:** 67 linhas singleton (63 não-terrenos + 4 terrenos não-básicos) + 32× Mountain = 99.
- **Coleção:** a lista estendida do usuário (`sideboard.txt`, 47 cartas / **34 elegíveis**) chegou durante esta fase e está cruzada com as lacunas na **seção 9**.

---

## 1. Comandante — análise linha a linha

`Krenko, Mob Boss` · `{2}{R}{R}` · Legendary Creature — Goblin Warrior · **3/3**
`{T}: Create X 1/1 red Goblin creature tokens, where X is the number of Goblins you control.`

**Ruling oficial (2012-07-01, WotC):** *"The ability counts each Goblin you control, including Krenko itself, not just the tokens it creates."*

| Fragmento | O que **exige** | O que **destrava** | Eixo de sinergia que abre |
|---|---|---|---|
| `{T}:` | Krenko **desvirado** e **sem summoning sickness**. Sem haste, ele só ativa no turno seguinte ao que entrou. | Nada no turno em que resolve. | **E1 — Haste.** Haste não é "melhoria": é o que corta um turno inteiro do relógio. Também torna Krenko resistente a *sorcery-speed removal* (ativa antes do turno do oponente). |
| `{T}:` (custo, não gatilho) | Krenko **precisa não atacar** no turno em que ativa (atacar tapa? não — atacar não usa {T}, mas vigilância importa: sem vigilância ele ataca **ou** ativa, nunca os dois na mesma janela útil). | — | **E2 — Vigilância / desvirar.** Vigilance, untappers (`Thousand-Year Elixir`-likes) e "as though it had haste" dobram o rendimento por turno. |
| **Habilidade ativada**, não disparada | Passa pela pilha. Oponente pode responder removendo Krenko: você perde a carta **e** os tokens. | — | **E3 — Proteção com resposta instantânea** (hexproof/indestrutível/blink/redirecionar). Este eixo está em **0 cartas** na lista. |
| `Create X ... tokens` | Nada além da ativação. | Bodies. **Só bodies.** Tokens 1/1 vanilla, sem haste, sem evasão. | **E4 — Conversor de body em dano** (lorde, trample, "não pode bloquear", sac outlet, ping em massa). Tokens sem conversor não fecham jogo. |
| `X = the number of **Goblins** you control` | Conta **Goblins**, não criaturas. Conta o próprio Krenko (ruling). Conta Goblins **não-token**. Conta permanentes **Kindred — Goblin** que não são criaturas (ex.: `Boggart Shenanigans` é `Kindred Enchantment — Goblin` e soma +1 ao X). **Não** conta Orcs, Spirits, Armies que não sejam Goblin. | Crescimento exponencial: N → 2N por ativação. | **E5 — Densidade de corpos Goblin baratos.** Cada Goblin de 1 mana no campo antes de Krenko vale **o dobro** por ativação. Contrapartida: efeitos que criam **1 corpo grande** (amass) valem 1, não N. |
| `1/1 red` | — | Alvo de anthem, de contadores em massa, de convoke/sac. | **E6 — Anthems e contadores em massa.** Um lorde vale mais que qualquer criatura isolada: transforma 15 blanks em 15 ameaças. |
| `3/3` corpo, CMC 4 | 4 manas + um turno de espera. | — | **E7 — Aceleração para o turno 3.** Krenko no T3 com haste = ativação no T3. Krenko no T5 sem haste = ativação no T6. Isso é a diferença entre o deck existir e não existir. |

### Os 7 eixos, e como a lista atual pontua neles

| Eixo | Cartas na lista | Situação |
|---|---|---|
| **E1** Haste para Krenko | 5 (2 permanentes) | **crítico** |
| **E2** Vigilância / desvirar | **0** | **descoberto** |
| **E3** Proteção com resposta | **0** | **descoberto** |
| **E4** Conversor de body em dano | 5 reais | insuficiente |
| **E5** Densidade Goblin barata | 32 criaturas, **todas Goblin** | **o único eixo saudável** |
| **E6** Anthem / contador em massa | 3 (nenhum é lorde) | **crítico** |
| **E7** Aceleração para T3 | 2 | **crítico** |

> **Achado estrutural:** as 32 criaturas do deck são **32 Goblins** — nenhuma exceção. A base de corpos está certa. Tudo o que falta é o que transforma corpos em vitória (E4, E6) e o que faz Krenko ligar cedo e sobreviver (E1, E3, E7).

---

## 2. Termos de busca do tema

Derivados dos eixos acima, para as fases seguintes (sempre com `legal:commander id<=r`):

1. `t:goblin` + `mv<=2` — densidade de corpos (E5)
2. `o:"Goblins you control get"` / `o:"Goblin creatures you control get"` — **lordes** (E6)
3. `o:"creatures you control have haste"` / `o:"gain haste"` — E1
4. `o:"number of creatures you control"` + `o:damage` — conversores (E4)
5. `o:"Sacrifice a creature:"` sem custo de mana — sac outlets grátis (E4, liga `Boggart Shenanigans`)
6. `o:"can't be blocked"` / `kw:trample` em massa — E4
7. `o:"whenever a creature you control enters" o:draw` — draw que escala com tokens
8. `o:indestructible` / `o:hexproof` + `t:instant` `mv<=2` — E3
9. `o:"return" o:"from your graveyard to the battlefield"` + `t:creature` — recuperação pós-wipe
10. `otag:mana-rock mv<=2` + `t:land o:"Add {R}{R}"` — E7

---

## 3. Auditoria por categoria — as 67 linhas

**Legenda de `Pontos`:** contagem de sinergias reais (regra 3 exige ≥2). Um "ponto" é: (a) somar/criar Goblin para o X de Krenko; (b) dar haste/proteção a Krenko; (c) converter tabuleiro largo em dano ou cartas; (d) habilitar ou ser habilitado por outra carta **nomeada** da lista.

| Carta | CMC | Tipo | Categorias | Pontos de sinergia com Krenko/deck | Veredito |
|---|---|---|---|---|---|
| Akki Ember-Keeper | 2 | Ench. Creature — Goblin Warrior 2/1 | tema | **1** — (a) corpo Goblin. Gatilho pede criatura **nontoken modificada** morrendo e cria um **Spirit** (não-Goblin): num deck de tokens sem equipamento, ≈0 disparos. | marginal |
| Ambitious Assault | 3 | Sorcery | tema | **1** — (c) +2/+0 no time. Draw só com "modified"; fontes de modificação = 4 Auras + poucos contadores, **zero equipamento**. | marginal |
| Arcane Signet | 2 | Artifact | ramp | **2** — (E7) acelera Krenko para o T3; (d) artefato = combustível de `Krenko, Baron of Tin Street`. | sólida |
| Assault on Osgiliath | X+3 | Sorcery | tema, wincon | **3** — (c) **double strike** em todos os Goblins = dobra o dano do enxame; (b) dá **haste** aos tokens recém-criados; (a) amass Orcs (1 corpo). Único alfa-strike verdadeiro da lista. | núcleo |
| Blitz of the Thunder-Raptor | 2 | Instant | remoção | **2** — (d) escala com os **20 instants/sorceries** do deck; exila (mata recursivos). | sólida |
| Boggart Shenanigans | 3 | **Kindred Ench. — Goblin** | tema, wincon | **2** — (a) **conta como Goblin para o X de Krenko mesmo sem ser criatura** (ruling de kindred confirmado); (c) dreno quando Goblins morrem. **Mas o deck não tem nenhum sac outlet repetível e gratuito** — a segunda metade fica inerte. | marginal |
| Bravado | 2 | Aura | tema | **1** — (c) escala com o enxame (+1/+1 por outra criatura). Aura em deck sem proteção = 2-por-1. **Não funciona com Zada** (encantamento). | marginal |
| Burn, Burn, Tree and Fern | 4 | Ench. — Saga | remoção, ramp | **1** — remoção de criatura (6 dano) + artefato. Nenhuma interação com Goblins/enxame. Cap. III–IV adicionam {R}. | marginal |
| Clamor Shaman | 3 | Creature — Goblin Shaman 1/1 | tema | **2** — (a) corpo Goblin; (c) tira um bloqueador do caminho ao atacar; riot dá contador (modificação) ou haste. | sólida |
| Conspicuous Snoop | 2 | Creature — Goblin Rogue 2/2 | tema, draw | **4** — (a) corpo Goblin de 2; (c) **card advantage virtual** (lança Goblins do topo); (d) **copia habilidades ativadas do Goblin no topo — incluindo o próprio `{T}` de Krenko**; (d) 32 Goblins no deck = topo quase sempre "vivo". | núcleo |
| Cosmotronic Wave | 4 | Instant | wipe, tema | **2** — (c) **"can't block"** = alfa strike com 1/1s; mini-wipe assimétrico em x/1. | sólida |
| Crash Through | 1 | Sorcery | tema, draw | **1** — (c) trample. Trample num 1/1 sem lorde entrega ~0 excesso. Cantrip (net 0 cartas). **Não funciona com Zada** (sem alvo). | marginal |
| Dragon Mantle | 1 | Aura | draw | **1** — cantrip + `{R}: +1/+0`. Não escala com o enxame. **Não funciona com Zada** (encantamento). | marginal |
| Ember Hauler | 2 | Creature — Goblin 2/2 | tema, remoção | **2** — (a) corpo Goblin; (c) converte corpo em 2 de dano; (d) morre → dispara `Boggart Shenanigans`. | sólida |
| Fanatical Firebrand | 1 | Creature — Goblin Pirate 1/1 | tema, remoção | **2** — (a) Goblin de 1 mana **com haste** (soma ao X já no turno em que entra); (c) `{T}`, sac: 1 dano; (d) `Boggart Shenanigans`. | sólida |
| Ferocity of the Wilds | 3 | Enchantment | tema, wincon | **3** — (c) **+1/+0 e trample a todos os atacantes não-Humanos** — as 32 criaturas do deck são não-Humanas, e os tokens também. Único anthem permanente que também resolve evasão. | núcleo |
| Fissure Wizard | 2 | Creature — Goblin Wizard 2/1 | tema | **2** — (a) corpo Goblin; (d) loot casa com `Squee, Goblin Nabob` (descarte livre). **Loot ≠ draw.** | sólida |
| Fists of Flame | 2 | Instant | draw, wincon | **3** — (d) **um dos dois únicos habilitadores reais de `Zada, Hedron Grinder`**; com Zada, cada cópia **compra uma carta** e o pump escala com as compras; (c) trample. | núcleo |
| Goblin Chainwhirler | 3 | Creature — Goblin Warrior 3/3 | tema, wipe, remoção | **2** — (a) corpo Goblin com **poder 3** (raro nesta lista); (c) wipe **assimétrico** de x/1 — limpa bloqueadores sem tocar nos seus tokens. | sólida |
| Goblin Cratermaker | 2 | Creature — Goblin Warrior 2/2 | tema, remoção | **3** — (a) corpo Goblin; (c) 2 dano a criatura **ou destrói permanente incolor não-terreno** (única resposta a Mana Rock/Eldrazi/veículo do deck); (d) `Boggart Shenanigans`. | núcleo |
| Goblin Glasswright // Craft with Pride | 2 | Creature — Goblin Sorcerer 2/2 | tema, ramp | **2** — (a) corpo Goblin; (d) Treasure alimenta `Krenko, Baron of Tin Street` (sac artefato) e o próprio Baron cria Goblin quando artefato vai pro cemitério. | sólida |
| Goblin Rabblemaster | 3 | Creature — Goblin Warrior 2/2 | tema | **2 pontos e 1 anti-sinergia grave** — (a) cria 1 Goblin **com haste** a cada combate; (c) cresce ao atacar. **F7: "Other Goblin creatures you control attack each combat if able" força TODO o enxame a atacar** — com 15 tokens 1/1 e sem lorde, isso é suicídio em massa e remove todos os bloqueadores. | marginal |
| Goblin Sky Raider | 3 | Creature — Goblin Warrior 1/2 | tema | **1** — (a) corpo Goblin. 3 manas por um 1/2 voador; a evasão entrega 1 de dano. | marginal |
| Goblin Surprise | 3 | Instant | tema | **2** — (a) 2 tokens Goblin (dobram no próximo Krenko); (c) modo +2/+0 no time. Flexível, mas 3 manas por 2 corpos é caro. | sólida |
| Goblin War Party | 4 | Sorcery | tema | **2** — (a) 3 tokens; (b) modo com **haste** ao time (liga Krenko no turno); entwine {2}{R} faz os dois. | sólida |
| Goblin Warchief | 3 | Creature — Goblin Warrior 2/2 | tema, ramp | **4** — (a) corpo Goblin; (b) **haste a todos os Goblins, permanentemente — inclusive Krenko**; (E7) reduz **{1}** de todo Goblin spell (Krenko por {1}{R}{R}); (d) liga `Lightning Volley` (tokens novos podem tapar). **A carta mais importante da lista depois do comandante.** | núcleo |
| Goblin Wardriver | 2 | Creature — Goblin Warrior 2/2 | tema | **2** — (a) corpo Goblin de 2; (c) battle cry = +1/+0 a cada outro atacante (escala com o enxame). | sólida |
| Goblin-town Flunkies | 2 | Creature — Goblin Soldier 1/1 | tema | **2** — (a) corpo Goblin **com haste**; (a) amass Goblins 1 (o Army **é** Goblin, soma ao X). | sólida |
| Goro-Goro, Disciple of Ryusei | 2 | Leg. Creature — Goblin Samurai 2/2 | tema | **3** — (a) corpo Goblin de 2; (b) **`{R}`: haste ao time, repetível** — segundo habilitador permanente de haste; (c) 5/5 voador, mas exige atacante **modificado** (pacote pela metade). | núcleo |
| Grishnákh, Brash Instigator | 3 | Leg. Creature — Goblin Soldier 1/1 | tema, remoção | **2** — (a) corpo Goblin; (c) rouba temporariamente a melhor criatura do oponente. **Amass Orcs 2 cria um ORC, que NÃO conta para o X de Krenko** (ruling confirmado). | marginal |
| Grotag Night-Runner | 3 | Creature — Goblin Rogue 2/3 | tema, draw | **2** — (a) corpo Goblin; (c) impulso 1 carta/turno ao conectar; (d) `Cosmotronic Wave`/`Clamor Shaman` garantem a conexão. | sólida |
| Gundabad Opportunist | 4 | Creature — Goblin Rogue 4/2 | tema, draw | **2** — (a) corpo Goblin **poder 4**; (c) impulso +1 carta com janela de 2 turnos. | sólida |
| Guttersnipe | 3 | Creature — Goblin Shaman 2/2 | tema, wincon | **2** — (a) corpo Goblin; (c) 2 dano a **cada** oponente por instant/sorcery — o deck tem **20**. Wincon secundário real, mas exige sobreviver. | sólida |
| Hordeling Outburst | 3 | Sorcery | tema | **2** — (a) **3 corpos Goblin de uma vez** (o melhor "combustível" pré-Krenko da lista); (d) enche `Quest for the Goblin Lord` e `Massive Raid`. | sólida |
| Innocent Bystander | 2 | Creature — Goblin Citizen 2/1 | tema | **1** — (a) corpo Goblin. Investigate exige **levar 3+ de dano** — nenhuma carta do deck consegue disparar isso de propósito (`Goblin Chainwhirler` só atinge oponentes). | marginal |
| Kick in the Door | 1 | Sorcery | tema | **3** — (b) **haste em alvo único = liga Krenko no turno em que ele entra por {R}**; (d) **um dos dois habilitadores reais de `Zada`**; (d) contador +1/+1 = modificação para `Goro-Goro`/`Ambitious Assault`. | núcleo |
| Krenko's Command | 2 | Sorcery | tema | **2** — (a) 2 Goblins por 2 no T2 → dobram na primeira ativação; (d) alimenta `Quest for the Goblin Lord`/`Massive Raid`. Também é o *redeploy* mais barato pós-wipe. | sólida |
| Krenko, Baron of Tin Street | 3 | Leg. Creature — Goblin 3/3 | tema, wincon | **3** — (a) corpo Goblin poder 3 **com haste**; (c) **`{T}`, sac artefato: +1/+1 em CADA Goblin — o único anthem permanente e cumulativo do deck**; (a) artefato no cemitério → token Goblin com haste. **F7: só há 3 artefatos + ~4 Treasures como combustível.** | núcleo |
| Lightning Volley | 4 | Instant | wincon | **2** — (c) **converte N Goblins em N pontos de dano dirigível** (ignora bloqueadores); (d) depende de `Goblin Warchief`/`Goro-Goro` para os tokens novos poderem tapar. Segundo caminho de vitória mais rápido do deck. | sólida |
| Massive Raid | 3 | Instant | wincon, remoção | **2** — (c) dano = nº de criaturas, a qualquer alvo; (d) escala diretamente com Krenko. Finalizador do último terço. | sólida |
| Misty Mountains Raider | 5 | Creature — Goblin Soldier 4/4 | tema | **2** — (a) corpo Goblin 4/4; (a) amass Goblins 2 por ataque — mas **empilha no MESMO Army**: 1 corpo, não N. Único 5-drop do deck. | marginal |
| Mob Mentality | 1 | Aura | tema | **2** — (c) trample + `+X/+0` onde X = atacantes; (d) casa com `Goblin Rabblemaster` (que força todos a atacar) e com `Tin Street Dodger` (indesbloqueável). **Não funciona com Zada.** Aura sem proteção = 2-por-1. | marginal |
| Mogg Sentry | 1 | Creature — Goblin Warrior 1/1 | tema | **1** — (a) Goblin de 1 mana. `+2/+2` no turno do oponente não converte nada. | marginal |
| Mudbutton Clanger | 1 | Creature — Goblin Warrior 1/1 | tema | **1** — (a) Goblin de 1 mana (isso importa: dobra na primeira ativação). Kinship = `+1/+1` até o fim do turno, irrelevante. | marginal |
| Mudbutton Torchrunner | 3 | Creature — Goblin Warrior 1/1 | tema, remoção | **2** — (a) corpo Goblin; (c) morre → 3 dano a qualquer alvo (remoção que sobrevive a wipe); (d) `Boggart Shenanigans`. 3 manas por um 1/1 é caro. | sólida |
| Mycosynth Wellspring | 2 | Artifact | ramp | **1** — busca **terreno básico para a mão** (não acelera, não compra). (d) artefato = combustível de `Krenko, Baron of Tin Street`, e volta a gatilhar ao ir pro cemitério. | marginal |
| Pinecone Strike | 2 | Instant | remoção | **2** — 3 dano + exílio a criatura; modo alternativo destrói **token de artefato** (Treasure/Clue/Food adversário). Modal = flexível. | sólida |
| Punishing Fire | 2 | Instant | remoção | **2** — 2 dano a qualquer alvo; (c) volta do cemitério quando **oponente ganha vida** — em mesa com lifegain vira motor. Sem lifegain na mesa, é um Shock. | sólida |
| Quest for the Goblin Lord | 1 | Enchantment | tema, wincon | **3** — (c) **+2/+0 a todas as criaturas** de forma permanente; (d) **uma única ativação de Krenko carrega as 5 contadores de uma vez**; (d) `Hordeling Outburst`/`Krenko's Command`/`Goblin Surprise` também carregam. Anthem de 1 mana que triplica o dano do enxame. | núcleo |
| Raging Goblin | 1 | Creature — Goblin Berserker 1/1 | tema | **1** — (a) Goblin de 1 mana com haste **apenas para si**. Não dá haste a Krenko, não converte nada. Vanilla. | marginal |
| Reckless Lackey | 1 | Creature — Goblin Pirate 1/2 | tema, draw, ramp | **3** — (a) Goblin de 1 mana com haste e first strike; (c) sac → 1 carta + Treasure; (d) Treasure alimenta `Krenko, Baron of Tin Street`. | sólida |
| Redcap Thief | 3 | Creature — Goblin Rogue 2/3 | tema, ramp | **2** — (a) corpo Goblin; (d) Treasure alimenta `Krenko, Baron of Tin Street`. 3 manas por um 2/3 + 1 Treasure é a taxa errada. | marginal |
| Rummaging Goblin | 3 | Creature — Goblin Rogue 1/1 | tema | **2** — (a) corpo Goblin; (d) loot repetível casa com `Squee, Goblin Nabob`. **F7: usa o `{T}`, disputando com atacar e com `Lightning Volley`. Loot ≠ draw.** | marginal |
| Shiny Impetus | 3 | Aura | ramp | **1** — Treasure por ataque da criatura goada (normalmente do oponente). Dá **+2/+2 a uma criatura adversária**. Não soma Goblin, não converte enxame. | marginal |
| Skullcrack | 2 | Instant | wincon | **1** — 3 de dano na cara + trava lifegain. Alcance para os últimos pontos; zero interação com o enxame. | marginal |
| Smite the Deathless | 2 | Instant | remoção | **2** — 3 dano a criatura, **remove indestrutível** e **exila** — a resposta mais "premium" da lista contra comandantes recursivos. | sólida |
| Sol Ring | 1 | Artifact | ramp | **3** — (E7) Krenko no T3; (d) artefato = combustível de `Krenko, Baron of Tin Street`; melhor carta de aceleração do deck. | núcleo |
| Squee, Goblin Nabob | 3 | Leg. Creature — Goblin 1/1 | tema, draw | **3** — (a) corpo Goblin; (c) **volta do cemitério todo upkeep = card advantage repetível e recuperação pós-wipe**; (d) descarte grátis para `Rummaging Goblin`/`Tormenting Voice`/`Fissure Wizard`. | sólida |
| The Autonomous Furnace | — | Land | terreno, draw | **1** — `{1}{R}`,`{T}`, sac: compra 1. Entra **virado**. Troca terreno por carta. | marginal |
| Tidings of War | 1 | Sorcery | tema | **1** — amass Goblins 1 (flashback: 3). **Cinco cartas de amass na lista empilham no MESMO Army: 5 cartas produzem 1 corpo Goblin, não 5.** O flashback é a única recursão pós-wipe além de Squee. | marginal |
| Tin Street Dodger | 1 | Creature — Goblin Rogue 1/1 | tema | **2** — (a) Goblin de 1 mana com haste; (c) `{R}`: **indesbloqueável** — portador ideal de `Bravado`/`Mob Mentality` (kill de 2 cartas com o enxame no campo). | sólida |
| Tormenting Voice | 2 | Instant | draw | **2** — **card advantage real (net +1)**; (d) o descarte é grátis com `Squee, Goblin Nabob`. Uma das pouquíssimas fontes de draw verdadeiro. | sólida |
| Warlord's Fury | 1 | Instant | tema, draw | **1** — first strike no time + cantrip (net 0). First strike em 1/1 só ganha de 1/1. **Não funciona com Zada** (sem alvo). | marginal |
| Zada, Hedron Grinder | 4 | Leg. Creature — Goblin Ally 3/3 | tema, draw | **2** — (a) corpo Goblin poder 3; (d) copia instant/sorcery que mira **só** Zada. **Habilitadores reais na lista inteira: 2 (`Fists of Flame`, `Kick in the Door`).** | marginal |
| **Forgotten Cave** | — | Land | terreno, draw | **1** — cycling `{R}` (substitui-se). Entra **virado**. | marginal |
| **Looming Spires** | — | Land | terreno | **0** — `+1/+1` e first strike **até o fim do turno** em um alvo (não é contador → não é "modified"). Entra **virado**. Efeito de 1 turno, custo de 1 turno. | sem sinergia |
| **Memorial to War** | — | Land | terreno, remoção | **0** — destrói terreno por `{4}{R}` + `{T}` + sac. Entra **virado**. **Violação da regra 11 (MLD)** e mana-negativa: paga 5 manas e um terreno para atrasar um oponente de três. | sem sinergia |
| **Mountain** ×32 | — | Basic Land | terreno | — | base |

---

## 4. Contagem por categoria — o diagnóstico

### 4.1 Draw / card advantage **real** (loot e cantrip não contam)

| Fonte | Tipo de vantagem | Real? |
|---|---|---|
| Conspicuous Snoop | motor — lança Goblins do topo (32 Goblins no deck) | ✅ motor |
| Squee, Goblin Nabob | +1 carta/upkeep, do cemitério | ✅ motor lento |
| Grotag Night-Runner | impulso 1/turno ao conectar | ✅ condicional |
| Tormenting Voice | net +1 | ✅ |
| Gundabad Opportunist | impulso +1 (one-shot) | ✅ |
| The Autonomous Furnace | +1 carta, −1 terreno | ⚠️ meio-ponto |
| Fists of Flame | cantrip sozinho; **draw em massa com Zada** | ⚠️ 2-card |

**Total: 5 fontes reais (6,5 sendo generoso). Meta: 12–13. Lacuna: −7.**

**Loot / rummage puro (net 0 cartas) — não conta como draw:** `Fissure Wizard`, `Rummaging Goblin`, `Forgotten Cave` (cycling) = **3 loots**.
**Cantrips (net 0):** `Crash Through`, `Warlord's Fury`, `Dragon Mantle`, `Fists of Flame` = **4**.
**Conditional/quase-morto:** `Ambitious Assault` (precisa de "modified"), `Innocent Bystander` (precisa levar 3+ dano), `Reckless Lackey` (troca corpo por carta) = **3**.

> **Diagnóstico:** o deck **parece** ter 12 fontes de carta e tem 5. Sete cartas ocupam o slot de "draw" sem entregar vantagem. É exatamente por isso que a mão esvazia no T4–T5 e nunca se recompõe.

### 4.2 Ramp

| Tipo | Cartas | Total |
|---|---|---|
| **Ramp padrão** | `Sol Ring`, `Arcane Signet` | **2** (meta 10–11 → **lacuna −8**) |
| **Ramp explosivo** | — | **0** (meta 2–3 → **lacuna −3**) |
| Redução de custo | `Goblin Warchief` ({1} em Goblin spells) | 1 |
| Treasure pontual | `Redcap Thief`, `Goblin Glasswright`, `Reckless Lackey`, `Shiny Impetus` | 4 (não-repetíveis) |
| Ritual marginal | `Burn, Burn, Tree and Fern` (cap. III–IV: `{R}`) | 1 |
| **Falso ramp** | `Mycosynth Wellspring` — põe básico **na mão**, não em jogo | — |

### 4.3 Interação

**Remoção pontual (10 peças — a meta numérica está batida):** `Blitz of the Thunder-Raptor`, `Goblin Cratermaker`, `Pinecone Strike`, `Punishing Fire`, `Smite the Deathless`, `Ember Hauler`, `Fanatical Firebrand`, `Mudbutton Torchrunner`, `Burn, Burn, Tree and Fern`, `Massive Raid`. (+`Grishnákh` como roubo temporário, +`Memorial to War` como MLD.)

**Mas a qualidade não bate:** *toda* a remoção é **dano**. O maior número num único card é 6 (`Burn, Burn, Tree and Fern`); a mediana é 3.
- Contra criatura com resistência ≥4: só 1 carta responde.
- Contra **encantamento**: **0 respostas**.
- Contra artefato: 3 (`Goblin Cratermaker` — só incolor, `Pinecone Strike` — só token, `Burn, Burn, Tree and Fern`).
- Contra **hexproof / protection**: 0.
- Contra **counterspell / stax**: 0.

**Board wipes:** **0 simétricos.** Mini-wipes assimétricos: `Goblin Chainwhirler` (1 dano nas criaturas dos oponentes) e `Cosmotronic Wave` (1 dano + can't block) = **2**, ambos matando só x/1.
Meta 2–4. Ressalva legítima: um deck de tokens **não** quer wipe simétrico — quer wipe assimétrico. O problema não é o tipo, é o **tamanho**: 1 de dano não responde nada acima de turno 4.

### 4.4 Proteção e recuperação pós-wipe — **a categoria crítica**

| Subcategoria | Cartas | Total |
|---|---|---|
| **Proteção** (hexproof, indestrutível, regenerar, flicker, redirecionar, contra-remoção) | — | **0** |
| **Recuperação real** (devolve cartas/permanentes ao jogo depois do wipe) | `Squee, Goblin Nabob` (a si mesmo), `Tidings of War` (flashback, sobre o mesmo Army), `Punishing Fire` (só se um oponente ganhar vida) | **2,5** |
| **Redeploy barato** (não é recuperação: gasta cartas novas) | `Krenko's Command` (2), `Hordeling Outburst` (3), `Goblin Surprise` (3), `Goblin War Party` (4) | 4 |

> **Este é o buraco número 1 do deck.** Krenko é um alvo de 4 manas, sem proteção nenhuma, cuja habilidade **passa pela pilha** — qualquer oponente com um Swords/Bolt em aberto desfaz o turno inteiro. E quando o wipe cai, a lista perde 15 corpos, o comandante e a mão de uma vez, com **zero** cartas capazes de reconstruir. Corresponde 1-para-1 à queixa *"apanho de board wipe"*.

### 4.5 Wincons — caminhos distintos e clock real

| # | Caminho | Peças necessárias | Clock realista (mesa de 4, 40 vidas) | Veredito |
|---|---|---|---|---|
| 1 | **Combate em massa** | Krenko + `Ferocity of the Wilds` ou `Quest for the Goblin Lord` | Krenko T4–5 → 1ª ativação T5–6 → 2ª T6–7 (4 gob.) → 3ª T7–8 (8 gob.). Letal em **1 oponente** por volta do **T9–T11** | real, **lento demais** |
| 2 | **Alfa-strike com double strike** | `Assault on Osgiliath` + 8+ Goblins + `X`≥2 | fecha **1 oponente** no turno em que resolve, a partir do T7 | real |
| 3 | **Burn dirigido** | `Lightning Volley` + `Goblin Warchief`/`Goro-Goro` + 12+ Goblins | ignora bloqueadores; ~12–20 dano num turno, T7+ | real, **combo de 2–3 cartas** |
| 4 | **`Massive Raid`** | 12+ criaturas | 12 dano num alvo — **finalizador do último terço**, não fecha sozinho | parcial |
| 5 | **`Guttersnipe`** | sobreviver + lançar dos 20 instants/sorceries | 2 dano por spell a **cada** oponente; ~4 gatilhos/jogo = 8 dano | parcial, lento |
| 6 | **`Fists of Flame` + `Zada`** | 2 cartas específicas + tabuleiro | compra N cartas e mata 1 oponente com trample | real, **frágil** |
| 7 | **`Boggart Shenanigans` dreno** | sac outlet repetível | — | **CAMINHO MORTO: o deck não tem nenhum sac outlet repetível e gratuito** |

**Caminhos que fecham de verdade e sem exigir combo de 2+ cartas: 2 (nº 1 e nº 2).** Meta: 3+.
Todos os 7 caminhos passam por **"ter tabuleiro largo e vivo"** — não há nenhuma rota alternativa se o enxame for respondido. Corresponde à queixa *"não fecho o jogo"*.

### 4.6 Terrenos

| | |
|---|---|
| **Total** | **36** |
| Básicos (`Mountain`) | 32 (89%) |
| **Utility** | **4** — `Forgotten Cave`, `Looming Spires`, `Memorial to War`, `The Autonomous Furnace` |
| Utility que entram **virados** | **4 de 4 (100%)** |
| Terrenos que produzem >1 mana | **0** |
| Terrenos com relevância tribal Goblin | **0** |
| Terrenos que dão evasão/haste | **0** |

### 4.7 Lordes de Goblin e habilitadores de haste — a contagem que decide o deck

**Lordes de Goblin (anthem estático que engorda Goblins): 0. Zero. Nenhum.**
O comandante fabrica 1/1 vanilla e **não existe uma única carta na lista que os torne permanentemente maiores** — exceto `Krenko, Baron of Tin Street`, que precisa de artefato para cada ativação (há 3 artefatos + ~4 Treasures no deck inteiro).

Anthems permanentes que se aproximam (nenhum é lorde):
1. `Ferocity of the Wilds` — +1/+0 e trample, **só atacando** (mas atinge 100% das criaturas: todas são não-Humanas)
2. `Quest for the Goblin Lord` — +2/+0 a **todas** as criaturas, após 5 ETBs de Goblin (Krenko liga numa ativação)
3. `Krenko, Baron of Tin Street` — +1/+1 **contador** em cada Goblin, custa 1 artefato por uso
4. `Goblin Wardriver` — battle cry, +1/+0 aos outros atacantes

**Habilitadores de haste:**

| Permanentes (afetam Krenko e os tokens) | 2 | `Goblin Warchief`, `Goro-Goro, Disciple of Ryusei` |
|---|---|---|
| One-shot (afetam o time num turno) | 3 | `Goblin War Party`, `Assault on Osgiliath`, `Kick in the Door` |
| **Haste só para si (NÃO ajudam Krenko)** | 7 | `Fanatical Firebrand`, `Goblin-town Flunkies`, `Krenko, Baron of Tin Street`, `Raging Goblin`, `Reckless Lackey`, `Tin Street Dodger`, `Clamor Shaman` (riot) |

> Com **2** habilitadores permanentes em 99 cartas, a chance de ter haste disponível quando Krenko entra é ≈15%. Nos outros 85% dos jogos, Krenko fica um turno inteiro exposto e ativa uma vez pela metade do valor.

### 4.8 Correção do agregado do briefing — "corpos que usam o `{T}`"

O `mtgdb` reportou **7**; o número real é **4**. Os falsos positivos vêm de **texto lembrete**:
- `Raging Goblin`, `Reckless Lackey` — reminder de **haste** ("can attack and `{T}` as soon as...")
- `Redcap Thief`, `Goblin Glasswright` — reminder de **Treasure** ("`{T}`, Sacrifice this token")

Corpos com custo `{T}` de verdade: `Krenko, Mob Boss`, `Krenko, Baron of Tin Street`, `Fanatical Firebrand`, `Rummaging Goblin`. O atrito interno é **leve** — e só aparece de fato quando `Lightning Volley` está em jogo.

---

## 5. Os quatro sub-temas: quais têm massa crítica e quais são pacote pela metade

O usuário pediu números. Aqui estão.

### 5.1 Pacote **Zada, Hedron Grinder** — 2 de 12 · **17% · PACOTE PELA METADE**

Ruling confirmado nesta sessão: *"Zada's ability triggers whenever you cast an instant or sorcery spell that targets **only** Zada and no other object or player."*

| Carta que "parece" habilitar Zada | Funciona? | Por quê |
|---|---|---|
| `Fists of Flame` | ✅ | instant, alvo único; **cada cópia também compra uma carta** |
| `Kick in the Door` | ✅ | sorcery, alvo único |
| `Crash Through` | ❌ | **não tem alvo** |
| `Warlord's Fury` | ❌ | **não tem alvo** |
| `Ambitious Assault` | ❌ | **não tem alvo** |
| `Goblin Surprise` | ❌ | **não tem alvo** |
| `Cosmotronic Wave` | ❌ | **não tem alvo** |
| `Goblin War Party` | ❌ | **não tem alvo** |
| `Tidings of War` | ❌ | **não tem alvo** |
| `Bravado` | ❌ | **Aura — encantamento, não instant/sorcery** |
| `Dragon Mantle` | ❌ | **Aura** |
| `Mob Mentality` | ❌ | **Aura** |
| `Shiny Impetus` | ❌ | **Aura** |
| `Massive Raid`, `Punishing Fire`, `Smite the Deathless`, `Pinecone Strike`, `Blitz…` | ❌ na prática | miram, mas **copiar dano mata o próprio enxame** |

> **Veredito:** o pacote Zada não é "meio" — é **um sexto**. Quatro Auras foram incluídas achando que Zada as copiaria (ela não copia encantamento) e seis mágicas de pump foram incluídas achando que Zada as copiaria (elas não têm alvo, então nem gatilham). **Zada em si não é o problema — é a carta que fez 10 outras cartas entrarem por um motivo que não existe.**

### 5.2 Pacote **amass** — 5 cartas, **1 corpo** · **ANTI-SINÉRGICO COM O COMANDANTE**

Ruling confirmado: *"To amass Goblins N, **if you don't control an Army creature**, create a 0/0 black Goblin Army token. **Then you choose an Army creature you control** and put N +1/+1 counters on it."*

Cartas: `Assault on Osgiliath` (Orcs), `Goblin-town Flunkies` (Goblins 1), `Grishnákh` (Orcs 2), `Misty Mountains Raider` (Goblins 2 por ataque), `Tidings of War` (Goblins 1/3).

- Todos os cinco empilham contadores no **mesmo Army**. Cinco cartas produzem **um** corpo.
- Krenko conta **corpos**, não poder. Um Army 9/9 vale exatamente **+1** no X.
- `Grishnákh` e `Assault on Osgiliath` amassam **Orcs** — o Army vira Orc e **não conta como Goblin** (a menos que já tenha sido feito Goblin por outra carta antes).
- O Army é um alvo único e gordo: morre para uma remoção e leva os contadores de cinco cartas junto.

> **Veredito:** amass é o oposto do que Krenko quer. É um mecanismo de **concentrar** poder num corpo; Krenko é um mecanismo de **espalhar** corpos. `Assault on Osgiliath` sobrevive à crítica pelo texto da segunda metade (double strike + haste em **Goblins e Orcs**), não pelo amass. `Misty Mountains Raider` e `Goblin-town Flunkies` sobrevivem pelo corpo. `Tidings of War` e `Grishnákh` não sobrevivem por nada.

### 5.3 Pacote **"modified"** — 3 payoffs, 0 equipamentos · **PACOTE PELA METADE**

Payoffs: `Akki Ember-Keeper` (nontoken **modificada** morrendo → Spirit), `Ambitious Assault` (draw se controlar modificada), `Goro-Goro` (Dragão 5/5 exige **atacante modificado**).
Fontes de modificação: 4 Auras (`Bravado`, `Dragon Mantle`, `Mob Mentality`, `Shiny Impetus` — esta normalmente vai numa criatura **adversária**), contadores de `Kick in the Door`, riot de `Clamor Shaman`, `Krenko, Baron of Tin Street`, contadores de amass. **Equipamentos: 0.**

> **Veredito:** metade. E `Akki Ember-Keeper` é pior que metade — pede uma criatura **nontoken** **modificada** morrendo, num deck onde ~80% dos corpos são tokens, e o token que produz é um **Spirit**, que não conta para Krenko.

### 5.4 Pacote **"ataque em massa" / pump coletivo** — 15 cartas · **TEM MASSA CRÍTICA (24%), MAS ESTÁ INVERTIDO**

`Ambitious Assault`, `Assault on Osgiliath`, `Bravado`, `Cosmotronic Wave`, `Crash Through`, `Ferocity of the Wilds`, `Fists of Flame`, `Goblin Surprise`, `Goblin War Party`, `Goblin Wardriver`, `Lightning Volley`, `Massive Raid`, `Mob Mentality`, `Quest for the Goblin Lord`, `Warlord's Fury` = **15 de 63 não-terrenos**.

- **11 dos 15 são one-shot** (instant/sorcery/Aura de um turno). Só 3 são anthem permanente (`Ferocity of the Wilds`, `Quest for the Goblin Lord`, `Goblin Wardriver`).
- Todos exigem **tabuleiro pronto**. Nenhum ajuda a construir o tabuleiro.
- O deck tem **0 lordes** e **15 cartas de pump temporário** — está pagando o preço de um pacote de anthem sem receber a permanência.

> **Veredito:** este sub-tema **tem** massa. O erro é de **forma**, não de quantidade: 15 slots de pump de um turno onde deveriam estar 4–5 lordes permanentes + conversores repetíveis. Trocar 6–8 desses one-shots por lordes/sac outlets **aumenta** o dano total e **reduz** a dependência de ter a carta certa na hora certa.

---

## 6. Candidatas a corte — ficha F1–F7 completa

> Regra 4: nenhuma destas é "carta fraca". Cada uma exerce funções reais, listadas abaixo, para que o protocolo de corte (§2 do checklist) possa rodar nas fases seguintes. **Eu não corto nada aqui** — entrego a ficha ao orquestrador.

### 6.1 · 0 pontos de sinergia

**`Memorial to War`** — Land
- **F1** Entra virado. `{T}`: Add `{R}`. `{4}{R}`, `{T}`, Sacrifice: destrói terreno alvo.
- **F2** Não é corpo. **F3** Tipo Land, conta para land drop. **F4** Não recebe nada.
- **F5** Fixação: produz `{R}` (irrelevante em mono-R com 32 Mountains). **F6** Entra virado — custa meio turno em toda partida.
- **F7** **Regra 11 (MLD)**; mana-negativo (5 manas + 1 terreno para atrasar 1 de 3 oponentes); entra virado num deck que quer Krenko no T3–T4.
- *Funções a cobrir num corte:* land drop + fonte de `{R}`. Cobertas por qualquer terreno.

**`Looming Spires`** — Land
- **F1** Entra virado. ETB: criatura alvo ganha `+1/+1` e first strike **até o fim do turno**. `{T}`: Add `{R}`.
- **F2** Não é corpo. **F3** Land. **F4** — **F5** Concede pump de 1 turno; **não é contador**, logo **não gera "modified"** para `Goro-Goro`/`Ambitious Assault`/`Akki Ember-Keeper`.
- **F6** Entra virado; atrasa Krenko. **F7** Compete com o land drop desvirado que o deck precisa.
- *Funções a cobrir:* land drop + `{R}` + um pump irrelevante. Cobertas por Mountain.

### 6.2 · 1 ponto de sinergia

**`Akki Ember-Keeper`** {1}{R} — Enchantment Creature — Goblin Warrior 2/1
- **F1** Criatura **nontoken modificada** sua morre → cria **Spirit** 1/1 incolor. **F2** Corpo 2/1, bloqueia, poder 2. **F3** É **Goblin** (soma ao X) **e** encantamento. **F4** Recebe anthems, contadores, Auras. **F5** — **F6** T2, corpo cedo (dobra na 1ª ativação). **F7** O gatilho exige nontoken+modificada (≈0 no deck); o token gerado é **Spirit**, não Goblin — não realimenta Krenko.
- *Funções:* corpo Goblin T2 · tipo encantamento · receptor de anthem. **A função de gatilho já está descoberta hoje.**

**`Ambitious Assault`** {2}{R} — Sorcery
- **F1** `+2/+0` no time; compra 1 **se** controlar modificada. **F2/F3** — **F4** — **F5** Pump coletivo. **F6** T3. **F7** Concorre com 20 outros CMC-3; a cláusula de draw é morta sem "modified"; **não gatilha Zada** (sem alvo).
- *Funções:* pump coletivo (F5) · draw condicional. Pump coberto por `Goblin Surprise`/`Goblin War Party`; draw **não coberto** (categoria já deficitária).

**`Bravado`** {1}{R} — Aura
- **F1** Encantada ganha `+1/+1` por **cada outra criatura** sua. **F2** Não é corpo. **F3** Encantamento. **F4** — **F5** Escala com o enxame; em `Tin Street Dodger` (indesbloqueável) é kill de 2 cartas. **F6** T2. **F7** 2-por-1 contra remoção (deck tem 0 proteção); **não gatilha Zada**; morre no wipe junto com o portador.
- *Funções:* conversor de largura em dano (F5) · combo com `Tin Street Dodger`/`Mob Mentality`.

**`Burn, Burn, Tree and Fern`** {3}{R} — Enchantment — Saga
- **F1** I: 6 dano a criatura de oponente. II: destrói artefato de oponente. III–IV: Add `{R}`. **F2** — **F3** Encantamento. **F4** Recebe contadores de lore. **F5** Ritual `{R}` (2×). **F6** T4, resolve ao longo de 4 turnos. **F7** Lento; a remoção sai só no turno seguinte; alvo obrigatoriamente de oponente.
- *Funções:* **maior remoção pontual do deck (6 dano)** · **única remoção de artefato não-condicional** · ramp marginal. Se sair, a resposta a criaturas de resistência 4–6 fica **descoberta**.

**`Crash Through`** {R} — Sorcery
- **F1** Trample no time + compra 1. **F2/F3/F4** — **F5** Concede trample. **F6** T1. **F7** Cantrip = net 0; trample em 1/1 sem lorde ≈ 0 dano em excesso; **não gatilha Zada**.
- *Funções:* trample coletivo (F5) · cicla. Trample coberto — e melhor — por `Ferocity of the Wilds` (permanente).

**`Dragon Mantle`** {R} — Aura
- **F1** ETB: compra 1. Encantada tem `{R}`: `+1/+0`. **F2** — **F3** Encantamento. **F4** — **F5** Mana sink repetível (o deck não tem outro), gera "modified". **F6** T1. **F7** Cantrip net 0; **não gatilha Zada**; 2-por-1 no wipe.
- *Funções:* cicla · **único mana sink repetível da lista** · fonte de "modified". O mana sink fica descoberto num corte — e mana sink é parte da queixa de flood.

**`Innocent Bystander`** {1}{R} — Creature — Goblin Citizen 2/1
- **F1** Ao levar **3+ de dano**, investigate (Clue). **F2** Corpo 2/1, bloqueia. **F3** **Goblin** (soma ao X); o Clue é artefato → combustível de `Krenko, Baron of Tin Street`. **F4** Recebe anthems. **F5** — **F6** T2. **F7** Nenhuma carta do deck consegue disparar de propósito (`Goblin Chainwhirler` só atinge oponentes); levar 3 dano num 2/1 = morrer.
- *Funções:* corpo Goblin T2 · receptor de anthem · gerador incidental de artefato.

**`Mogg Sentry`** {R} — Creature — Goblin Warrior 1/1
- **F1** Oponente lança mágica → `+2/+2` até o fim do turno. **F2** Corpo 1/1. **F3** **Goblin de 1 mana** (o slot mais valioso do deck: dobra em toda ativação). **F4** Recebe anthems e contadores. **F5** — **F6** T1. **F7** O buff acontece no turno **do oponente** (defensivo), não no seu ataque.
- *Funções:* **corpo Goblin de 1 mana** · bloqueador que sobe · receptor de anthem.

**`Mudbutton Clanger`** {R} — Creature — Goblin Warrior 1/1
- **F1** Kinship: olha o topo, se compartilha tipo, `+1/+1` até o fim do turno. **F2** Corpo 1/1. **F3** **Goblin de 1 mana**. **F4** Recebe anthems/contadores. **F5** Informação do topo (sinergia leve com `Conspicuous Snoop`). **F6** T1. **F7** Kinship é irrelevante (buff de 1 turno).
- *Funções:* **corpo Goblin de 1 mana** · receptor de anthem.

**`Mycosynth Wellspring`** {2} — Artifact
- **F1** Ao entrar **ou ir para o cemitério**: busca terreno **básico para a mão**. **F2** — **F3** **Artefato** → combustível de `Krenko, Baron of Tin Street` (sac artefato: +1/+1 em cada Goblin) e gatilha a criação de Goblin do Baron ao ir para o cemitério — **e dispara duas vezes**. **F4** — **F5** Fixação (irrelevante em mono-R). **F6** T2 sem impacto. **F7** Não acelera (terreno vai para a **mão**); num deck com 36 terrenos, buscar mais terreno agrava o flood.
- *Funções:* **artefato-combustível duplo do Baron** · anti-flood reverso. A função de artefato **fica descoberta** se sair sem reposição (restam Sol Ring + Signet, que você não quer sacrificar).

**`Raging Goblin`** {R} — Creature — Goblin Berserker 1/1
- **F1** Haste. Só isso. **F2** Corpo 1/1 com haste. **F3** **Goblin de 1 mana**; entra e **já soma ao X no mesmo turno** (relevante se Krenko ativar naquele turno). **F4** Recebe anthems/contadores; com `Ferocity of the Wilds` vira 2/1 trample. **F5** — **F6** T1. **F7** Nenhum.
- *Funções:* **corpo Goblin de 1 mana com haste** · receptor de anthem. É a carta mais "vanilla" da lista, mas exerce a função E5 exatamente como o comandante pede.

**`Shiny Impetus`** {2}{R} — Aura
- **F1** Encantada ganha `+2/+2` e é **goaded**; ao atacar, você cria Treasure. **F2** — **F3** Encantamento; produz **artefato** (Treasure) → `Krenko, Baron of Tin Street`. **F4** — **F5** Treasure repetível = **a única fonte de ramp recorrente do deck**; goad é interação política. **F6** T3. **F7** Dá `+2/+2` a criatura **adversária**; não soma Goblin; 3 manas que não constroem tabuleiro.
- *Funções:* ramp recorrente (Treasure) · combustível do Baron · interação política (goad).

**`Skullcrack`** {1}{R} — Instant
- **F1** Ninguém ganha vida no turno; dano não pode ser prevenido; 3 dano em jogador/PW. **F2/F3/F4** — **F5** **Desliga fog e prevenção de dano** — a única carta que protege um alfa-strike de `Fog`/`Holy Day`/lifegain. **F6** T2. **F7** 3 dano isolado num formato de 40 vidas.
- *Funções:* alcance final (3 dano) · **anti-lifegain** · **anti-prevenção de dano num alfa-strike**. A terceira função é real e não está coberta por mais nada.

**`Tidings of War`** {R} — Sorcery
- **F1** Amass Goblins 1; se lançada do cemitério, amass Goblins 3. Flashback `{3}{R}`. **F2** — **F3** — **F4** O Army recebe contadores. **F5** — **F6** T1. **F7** **Empilha no mesmo Army que outras 4 cartas: 1 corpo, não N** — anti-sinérgico com o X de Krenko.
- *Funções:* 1 corpo Goblin (Army) · **recursão pós-wipe (flashback)** — uma das únicas 2 do deck. A função de recursão está **numa categoria em déficit total**.

**`Warlord's Fury`** {R} — Instant
- **F1** First strike no time + compra 1. **F2/F3/F4** — **F5** Concede first strike coletivo. **F6** T1. **F7** Cantrip net 0; first strike em 1/1 só ganha de 1/1; **não gatilha Zada**.
- *Funções:* first strike coletivo (F5) · cicla. Substituível por `Assault on Osgiliath` (double strike, que **inclui** first strike e dobra o dano).

**`The Autonomous Furnace`** — Land
- **F1** Entra virado. `{T}`: Add `{R}`. `{1}{R}`, `{T}`, Sac: compra 1. **F2** — **F3** Land. **F5** `{R}`. **F6** Entra virado. **F7** Converte terreno em carta (bom contra flood) mas entra virado (ruim contra a curva).
- *Funções:* land drop · **mana sink que vira carta** — anti-flood, diretamente ligado à queixa de mana.

**`Forgotten Cave`** — Land
- **F1** Entra virado; `{T}`: `{R}`; cycling `{R}`. **F6** Entra virado. **F7** Entra virado num deck que quer curva limpa.
- *Funções:* land drop · **flexibilidade terreno/carta** (anti-flood e anti-screw ao mesmo tempo).

### 6.3 · 2 pontos, mas com atrito grave o suficiente para revisão

**`Goblin Rabblemaster`** {2}{R} — Creature — Goblin Warrior 2/2
- **F1** *Outras* Goblins suas **atacam a cada combate se puderem**; início do combate: 1 Goblin com haste; ao atacar, `+1/+0` por outro Goblin atacante. **F2** 2/2. **F3** Goblin. **F4** Anthems. **F5** Gera corpo/turno. **F6** T3. **F7** **A cláusula de ataque forçado é uma anti-sinergia de primeira ordem com Krenko**: com 15 tokens 1/1 e zero lorde, o ataque compulsório joga o enxame inteiro contra bloqueadores maiores, esvazia os bloqueadores e **destrói o próprio X de Krenko para a ativação seguinte**. Também força tokens a atacar em vez de ficarem disponíveis para `Lightning Volley`.
- *Funções:* gerador de Goblin com haste por turno · corpo Goblin · anthem-ao-atacar. **Nota:** se ficar, casa com `Mob Mentality` (que quer todos atacando).

**`Grishnákh, Brash Instigator`** {2}{R} — Legendary Creature — Goblin Soldier 1/1
- **F1** ETB: amass **Orcs** 2; então rouba criatura não-lendária de oponente com poder ≤ poder do Army, desvira, haste até o fim do turno. **F2** 1/1. **F3** Goblin (o **Army é Orc**, não conta). **F4** Anthems. **F5** Haste à criatura roubada. **F6** T3. **F7** Amass Orcs pode **converter seu Army Goblin em Goblin Orc** (ok) ou criar um Orc puro (não conta para Krenko); o roubo é temporário e sem sac outlet não vira vantagem permanente.
- *Funções:* corpo Goblin · **roubo temporário (única interação "tempo" do deck)** · amass.

**`Misty Mountains Raider`** {4}{R} — Creature — Goblin Soldier 4/4
- **F1** Sempre que você atacar, amass Goblins 2. **F2** **Corpo 4/4 — o maior do deck.** **F3** Goblin. **F4** Anthems/contadores. **F5** — **F6** T5, **único 5-drop do deck**. **F7** Amass empilha num único Army: contribui **+1** ao X de Krenko, não +2/ataque.
- *Funções:* **maior corpo do deck** · atacante/bloqueador que sobrevive a `Goblin Chainwhirler` alheio · amass Goblins (torna o Army um Goblin, o que salva `Assault on Osgiliath`).

**`Redcap Thief`** {2}{R} — Creature — Goblin Rogue 2/3
- **F1** ETB: cria Treasure. **F2** 2/3. **F3** Goblin + gera artefato (`Krenko, Baron of Tin Street`). **F4** Anthems. **F5** Ramp pontual/fixação. **F6** T3 (mais um no engarrafamento de 21 cartas em CMC 3). **F7** Taxa cara: 3 manas por 2/3 + 1 Treasure.
- *Funções:* corpo Goblin · Treasure (ramp pontual + combustível do Baron).

**`Rummaging Goblin`** {2}{R} — Creature — Goblin Rogue 1/1
- **F1** `{T}`, descarte 1: compre 1. **F2** 1/1, **usa `{T}` de verdade**. **F3** Goblin. **F4** Anthems. **F5** Seleção repetível. **F6** T3. **F7** Loot = **net 0 cartas** (não resolve a queixa de mão vazia); `{T}` compete com atacar e com `Lightning Volley`.
- *Funções:* corpo Goblin · **seleção repetível** · descarte-outlet para `Squee, Goblin Nabob` (loop de seleção infinito e grátis, 1×/turno).

**`Zada, Hedron Grinder`** {3}{R} — Legendary Creature — Goblin Ally 3/3
- **F1** Instant/sorcery que mira **só** Zada é copiado para cada outra criatura sua que pudesse ser alvo. **F2** **Corpo 3/3** — um dos 6 corpos com poder ≥3. **F3** Goblin (soma ao X). **F4** Anthems/contadores/Auras. **F5** Multiplica efeitos direcionados. **F6** T4 — **compete diretamente com o slot de Krenko**. **F7** **Só 2 habilitadores em 99 cartas**; 4 manas para um 3/3 que não faz nada sozinho; o próprio Zada precisa sobreviver um turno.
- *Funções:* corpo Goblin poder 3 · motor de cópia (2 habilitadores) · receptor de anthem. **Nota do briefing:** foi **oferecida como intocável e recusada** pelo usuário — segue protocolo normal.

**`Boggart Shenanigans`** {2}{R} — Kindred Enchantment — Goblin
- **F1** Outro **Goblin** seu vai do campo para o cemitério → 1 dano a jogador/PW. **F2** Não é corpo. **F3** **É um Goblin que você controla sem ser criatura — soma +1 ao X de Krenko** (kindred confirmado por ruling). Encantamento. **F4** — **F5** Converte morte de token em dano (imune a bloqueador e a wipe: **o wipe adversário vira sua wincon**). **F6** T3. **F7** **Não há um único sac outlet repetível e gratuito no deck** — depende de mortes involuntárias.
- *Funções:* **+1 permanente no X de Krenko** · **conversor de wipe em dano** · payoff de aristocratas. É a carta que **mais ganha** se o deck receber um sac outlet grátis.

---

## 7. Lacunas em ordem de gravidade, amarradas às queixas

| # | Lacuna | Queixa que explica | Números |
|---|---|---|---|
| **1** | **Nenhum lorde de Goblin e nenhum conversor permanente** | *"não fecho o jogo"* | **0 lordes**. 3 anthems permanentes, dos quais 2 só valem atacando. 15 cartas de pump **temporário**. 20 tokens 1/1 = 20 de dano bruto que qualquer bloqueador 2/2 anula. |
| **2** | **Zero proteção; recuperação pós-wipe quase nula** | *"apanho de board wipe"* | **0** cartas de proteção. **2,5** de recuperação. Krenko é ativado, não disparado → morre **em resposta** e você perde o turno. |
| **3** | **Card advantage real é 5, não 12** | *"fico sem cartas na mão"* | 5 fontes reais vs meta 12–13. **7 cartas ocupam o slot de draw sem gerar carta** (3 loots + 4 cantrips). Nenhuma fonte escala com o número de Goblins. |
| **4** | **Ramp de 2 e zero mana sink** | *"mana/curva travando"* | 2 rocks. 36 terrenos com top-end 5. Do T6 em diante sobra mana e **não há onde gastar** (`Massive Raid` e `Assault on Osgiliath` são os únicos X/sinks). |
| **5** | **Só 2 habilitadores permanentes de haste** | *"não fecho o jogo"* + *"curva travando"* | ≈15% de chance de ter haste quando Krenko entra. Nos outros 85%, o relógio atrasa **um turno inteiro** e Krenko passa exposto. |
| **6** | **Três pacotes pela metade ocupando ~20 slots** | todas as quatro | Zada 2/12 · amass 5 cartas → 1 corpo · modified 3 payoffs / 0 equipamentos. |
| **7** | **Remoção é só dano; 0 resposta a encantamento; wipes de 1 ponto** | *"não fecho o jogo"* | 10 peças de interação, mediana 3 de dano. Um `Ghostly Prison`/`Propaganda`/`Blood Moon` na mesa **encerra o deck**. |
| **8** | **Anti-sinergia ativa: `Goblin Rabblemaster` força ataque compulsório** | *"apanho de board wipe"* | Com 15 tokens 1/1 sem lorde, o ataque forçado destrói o próprio X de Krenko e zera os bloqueadores. |

---

## 8. Curva e mana — o formato está errado nos **dois** sentidos ao mesmo tempo

**Sim. E é essa a resposta exata.**

Números apurados: 36 terrenos · curva 1:14 · 2:20 · 3:21 · 4:7 · 5:1 · 6+:0 · **soma de CMC 150 / 63 cartas = MV médio 2,38**.

### 8.1 Terreno demais para a curva que existe
Para MV médio 2,38 com top-end 5, a referência de mono-vermelho é **33–35 terrenos**. Com 36, **o 6º terreno em diante é uma carta em branco**, porque:
- Krenko custa `{T}`, **não mana** — ele não é mana sink.
- Os únicos escoadouros de mana excedente são `Massive Raid` e `Assault on Osgiliath` (**2 cartas**), mais `Dragon Mantle` e `The Autonomous Furnace` como sinks marginais.
- Resultado: da metade da partida em diante você desvira 8–10 manas e passa o turno. **A queixa "mana travando" neste deck não é falta de mana nem de cor — é excesso de mana sem destino.** Com 32 Mountains, screw de cor é matematicamente impossível.

### 8.2 Terreno de menos do tipo que importa
Ao mesmo tempo, o deck **não tem nenhum terreno que faça algo**:
- 0 que produzem mais de 1 mana; 0 tribais de Goblin; 0 que dão evasão ou haste; 0 desvirados com efeito.
- Dos 4 utility, **4 entram virados** — ou seja, o único desvio da base básica **piora** a curva em vez de melhorar.
- Em mono-vermelho, a exigência de cor é trivial (`{R}{R}{R}` de `Goblin Chainwhirler` é o teto). **100% dos slots de terreno estavam livres para virar utility, e 0% viraram.**

### 8.3 O engarrafamento no 3 é o pior formato possível para este plano
**21 das 63 cartas não-terreno (33%) custam 3.** Contra 14 de custo 1.
- Com 36 terrenos, do T3 ao T6 você lança **exatamente um feitiço por turno**. Um deck que ganha por **quantidade de corpos** está desdobrando **um permanente por turno**.
- A forma correta para Krenko é **fundo-pesada**: muitos 1s e 2s (que dobram em toda ativação) e poucos 3s de altíssimo impacto. A forma atual é **barrigudinha no 3**, que é a curva de um deck de valor midrange — não de um deck de enxame.
- Consequência direta no relógio: Krenko (CMC 4) chega no T4–T5 com o tabuleiro tendo **2–3 Goblins**, então a primeira ativação rende 3–4 tokens em vez de 6–8.

### 8.4 O "top-end zero" **não é** o problema que parece ser
Faltar carta de CMC 6–7 **não é o defeito**. `Krenko, Mob Boss` **é** o top-end: ele é a bomba que o deck compra todo jogo (zona de comando). Adicionar 7-drops só pioraria a curva.
O que falta no topo não é **custo** — é **escala**: peças baratas cujo efeito cresce com o número de Goblins (lordes, sac outlets grátis, drenos por criatura, mana sinks repetíveis). Hoje o deck tem 3 delas (`Ferocity of the Wilds`, `Quest for the Goblin Lord`, `Massive Raid`).

### 8.5 Direção para a fase 6 (manabase)
1. **36 → 34–35 terrenos**, com os 2 slots liberados indo para ramp de 2 manas (que resolve E7 e não fica morto no late).
2. **Trocar os 4 utility que entram virados** por utility **desvirados** e/ou terrenos que escoem mana.
3. Reduzir a barriga de CMC 3: alvo `1:16–18 · 2:20–22 · 3:12–14 · 4:6–8`.
4. Todo slot novo de 3 manas precisa ser **anthem permanente, sac outlet ou motor de carta** — não mais pump de um turno.

---

## 9. Coleção do usuário × as 8 lacunas (regra 7 — coleção antes de compra)

A lista estendida chegou durante esta fase: `sideboard.txt`, **47 cartas**, das quais **34 são elegíveis**
(mono-R ou incolores) — ficha em `00b-colecao-elegivel.md`. Oracle das peças abaixo puxado nesta sessão.
O maybeboard antigo do deck *Inspirit* (43 cartas azuis/artefato) permanece **irrelevante** para Krenko.

> **Resultado geral: 6 das 8 lacunas têm resposta parcial ou total dentro da coleção, a custo zero.**
> As duas que **não** têm são justamente a nº 2 (proteção) e a nº 8 (anti-sinergia do Rabblemaster).

### 9.1 As peças que mudam o deck (achados de primeira ordem)

| Carta | CMC | Lacuna que ataca | Por que é um achado (≥2 pontos de sinergia) |
|---|---|---|---|
| **`Haunted Cloak`** | 3 (equip **{1}**) | **#5 haste** + **#1** + **#6 modified** | Concede **vigilance + trample + haste**. Em Krenko: haste = ativa **no turno em que entra**; **vigilance = ataca E ativa no mesmo turno** (resolve o eixo E2, que estava em zero); trample resolve E4. Equip {1} permite mover para o próximo Krenko depois de uma remoção. Ainda torna o portador **"modified"**, ligando `Goro-Goro`, `Ambitious Assault` e `Akki Ember-Keeper`. **Cinco funções — é a melhor carta da coleção para este comandante.** |
| **`Dragon Throne of Tarkir`** | 4 (equip {3}) | **#1 conversor permanente** | `{2}`,`{T}`: *outras* criaturas ganham trample e `+X/+X`, X = poder do portador. Num 3/3 (`Goblin Chainwhirler`, `Zada`) ou no 4/4 (`Misty Mountains Raider`), é um **Overrun repetível todo turno** — exatamente o "conversor permanente" que o deck tem **zero**. Também é mana sink (lacuna #4). **F7:** o portador ganha **defender** e usa `{T}` → **não equipar em Krenko** (conflita com a ativação dele). |
| **`Urza's Incubator`** | 3 | **#4 ramp** + **#5** + **#8 curva** | Escolha "Goblin": **todo feitiço de criatura Goblin custa `{2}` menos**. 32 das 63 cartas não-terreno do deck são Goblins. Krenko passa a custar `{R}{R}`; com `Goblin Warchief` junto, `{R}`. Desengarrafa o CMC 3 inteiro e é o único "ramp" que escala com o tema. |
| **`Kuldotha Rebirth`** | **1** | **#1** + **#8 curva** + E5 | **Três Goblins por `{R}`** — a melhor taxa de corpos possível para o X de Krenko. Sacrifica um artefato: o deck já tem `Sol Ring`, `Arcane Signet`, `Mycosynth Wellspring` e Treasures, e a coleção acrescenta 4 artefatos de 1–2 CMC. **Também transforma `Mycosynth Wellspring` (hoje um "1 ponto") num combo real**, e o artefato sacrificado gatilha `Krenko, Baron of Tin Street`. |
| **`Panic Spellbomb`** | **1** | **#3 draw** + **#1 evasão** + combustível | Três funções por 1 mana: (a) **artefato** de 1 CMC = combustível de `Kuldotha Rebirth` **e** de `Krenko, Baron of Tin Street`; (b) `{T}`, sac: criatura alvo **não pode bloquear** — abre caminho para o alfa-strike; (c) ao ir pro cemitério, pague `{R}`: **compre uma carta**. Cobre lacuna 3 e 1 no mesmo slot. |

### 9.2 Cobertura lacuna a lacuna

| # | Lacuna | Coberta pela coleção? | Peças |
|---|---|---|---|
| **1** | Nenhum lorde / conversor permanente | **parcial — boa** | `Dragon Throne of Tarkir` (Overrun repetível), `Haunted Cloak` (trample), `Kuldotha Rebirth` (largura), `Panic Spellbomb`/`Leonin Bola` (remoção de bloqueador), `Mercadia's Downfall` (Overrun contra manabases não-básicas). **Continua faltando um lorde estático de verdade** — nenhuma das 34 é `Goblin Chieftain`/`Goblin King`/`Coat of Arms`. |
| **2** | **Zero proteção / recuperação pós-wipe** | **NÃO COBERTA** | Nenhuma das 34 dá hexproof, indestrutível, flicker ou devolve permanentes do cemitério. `Haunted Cloak` mitiga (haste faz Krenko ativar antes de morrer) mas não protege. **Esta lacuna exige compra.** |
| **3** | Card advantage real = 5 | **parcial — boa** | `Thrill of Possibility` (net +1, instant), `Seize Opportunity` (impulso 2, modal com pump), `Oracle's Vault` (impulso repetível → **grátis** com 3 brick counters), `Bargaining Table` (draw repetível **e** mana sink), `Panic Spellbomb`, `Golden Egg`/`Instant Ramen` (cantrip **em corpo de artefato** = combustível de Kuldotha/Baron), `Mad Ratter` (**é Goblin**, soma ao X, e transforma draw em 2 corpos/turno). Chega perto de fechar o déficit de 7. |
| **4** | Ramp 2 e **zero mana sink** | **metade coberta** | Mana sink: `Bargaining Table`, `Oracle's Vault`, `Dragon Throne of Tarkir`, `Thaumaton Torpedo` — **4 escoadouros novos**, que era o que a queixa de mana realmente pedia. Ramp de verdade: **só `Urza's Incubator`** (redução). Não há um único mana rock nas 34 — **a metade "ramp" exige compra.** |
| **5** | 2 habilitadores permanentes de haste | **coberta em 1** | `Haunted Cloak` — permanente, móvel, e ainda entrega vigilância. Passa de 2 → 3. |
| **6** | Três pacotes pela metade | **coberta — e é a decisão da rodada** | A coleção tem **6 equipamentos elegíveis** (`Haunted Cloak`, `Dragon Throne of Tarkir`, `Armory of Iroas`, `Leonin Bola`, `Dire Flail`, `Trailblazer's Torch`). O pacote "modified" falhava por ter **0 equipamentos**; agora dá para **completá-lo de graça** em vez de cortá-lo. **Decisão para o orquestrador:** completar o pacote modified com 2–3 equipamentos (que já valem sozinhos) ou cortar `Akki Ember-Keeper`/`Ambitious Assault`. O pacote **Zada** e o pacote **amass** continuam sem reparo possível na coleção. |
| **7** | Remoção só de dano; 0 resposta a encantamento | **parcial** | Sweepers assimétricos maiores: `Seismic Wave` (2 em qualquer alvo **+ 1 em cada criatura não-artefato de um oponente**), `Radiating Lightning` (3 no jogador + 1 nas criaturas dele). Remoção pontual mais eficiente: `Magma Spray` (exila), `Lightning Strike`, `Bombard` (4), `Chandra's Pyrohelix`/`Twin Bolt` (2 divididos = mata 2 bloqueadores x/1), `Searing Barrage` (5), `Stonefury` (dano = terrenos; com 34–36 terrenos, mata quase tudo), `Fateful End`, `Seal of Fire`. Artefato: `Steel Wrecking Ball`, `Smashing Success`. **Encantamento: 1 resposta em 34 — `Thaumaton Torpedo` (`{6}`,`{T}`, sac: destrói permanente não-terreno)**, cara mas é a única; `Liquimetal Coating` + `Smashing Success`/`Steel Wrecking Ball` é uma segunda rota de 2 cartas. |
| **8** | `Goblin Rabblemaster` força ataque compulsório | **NÃO COBERTA** | Nenhuma carta resolve; é decisão de corte, não de reposição. (Nota: `Haunted Cloak`/vigilância **mitiga** — atacar deixa de custar os bloqueadores.) |

### 9.3 Peças da coleção que **não** passam no filtro de 2 pontos

`Goblin Gathering` (3 CMC por 2 tokens — `Hordeling Outburst` é estritamente melhor e o deck já está engarrafado no 3) · `Dire Flail` (equipamento de +2/+0 sem função sobreposta) · `Armory of Iroas` (contador por ataque num deck de 1/1 sem proteção) · `Trailblazer's Torch` (4 CMC, punir bloqueio é win-more) · `Liquimetal Coating` (só serve como metade de um combo de 2 cartas) · `Golden Egg`/`Instant Ramen` (só entram se `Kuldotha Rebirth` entrar — sinergia **dependente**, não própria) · `Chandra's Outrage`, `Searing Barrage`, `Stonefury` (CMC 4–5 num deck que precisa **descer** a curva).

### 9.4 O que ainda exige compra (para as fases 3–5)

1. **Proteção de comandante** — o eixo E3 está em **0** e a coleção não tem nada. Prioridade máxima.
2. **Lorde estático de Goblin** — a coleção não tem nenhum; é o item nº 1 da lacuna nº 1.
3. **Mana rocks de 2 CMC** — a coleção só oferece redução de custo, não aceleração real.
4. **Sac outlet repetível e gratuito** — destrava `Boggart Shenanigans`, `Mudbutton Torchrunner`, `Ember Hauler` e converte o board wipe adversário em dano (ataca as lacunas 1 **e** 2 ao mesmo tempo).
