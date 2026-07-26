# Deck — Tori D'Avenant, Fury Rider (R/W Go-Wide Knights)

Fonte de verdade do deck. Atualizado pelo orquestrador após cada checkpoint aprovado.

## Status do pipeline (sessão pausada em 2026-07-25)

Modo `improve`. Progresso até aqui:

- **Fase 0 (intake)**: completa — `00-briefing.md`.
- **Fase 1 (auditoria)**: completa — `02-theme.md`. Diagnóstico aprovado: atacar **todas as áreas** (draw, ramp, interação/wipes, manabase, wincons), nessa ordem.
- **Fase 2 — draw**: completa e **aplicada** neste `deck.md` (3 swaps, ver histórico abaixo). Análise em `03-draw.md`.
- **Fase 2 — ramp**: proposta pronta em `04-ramp.md` (10 swaps: 7 padrão + 3 explosivos, ~US$3,51, todos `torneio`), **mas ainda NÃO aprovada nem aplicada**. Alerta do especialista: a versão completa troca 9 criaturas por só 4, saldo líquido de -5 corpos no plano de ataque.
  - **Próximo passo ao retomar**: perguntar ao usuário como aplicar os swaps de ramp — (a) os 10 completos, (b) versão conservadora de 7 (remove os swaps 3 Bogardan Lancer→Talisman of Conviction, 7 Youthful Knight→Solemn Simulacrum, 9 Knight of Sorrows→Big Score da lista), ou (c) ajuste manual. Depois disso, atualizar `deck.md` com os swaps aprovados.
- **Fase 2 — interação/wipes**: não iniciada (gap crítico: 0 wipes, meta 2–4).
- **Fase 2 — manabase**: não iniciada (gap crítico: fixação R/W fraca, 6 terrenos utilitários todos tapped, provavelmente precisa subir de 36 para ~38+ terrenos totais).
- **Fase 2 — wincons**: não iniciada (gap: só 1 caminho de vitória — combate —, meta 3+).
- **Fase 3 (revisão final)** e **Fase 4 (relatório + goldfishing)**: não iniciadas.

Para retomar, reabra `/improve-deck` neste projeto e continue a partir da decisão de ramp pendente acima.

## Comandante

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Tori D'Avenant, Fury Rider | 4 | Creature — Human Knight | RW | tema, wincon | Anthem no ataque (+1/+1 a outros atacantes), trample aos vermelhos, desvira brancos — núcleo do plano go-wide |

## Criaturas

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Adriana, Captain of the Guard | 5 | Creature — Human Soldier | RW | tema, wincon | Melee dá +1/+1 por oponente atacado, estende a todos — combina com o anthem de Tori |
| Blazing Blade Askari | 3 | Creature — Human Knight | R | tema | Flanking; tribal Knight |
| Bogardan Lancer | 2 | Creature — Human Knight | R | tema | Bloodthirst + flanking; tribal |
| Cavalry Drillmaster | 2 | Creature — Human Knight | W | tema | Combat trick em corpo; tribal |
| Dawnstrike Vanguard | 6 | Creature — Angel Soldier | W | tema | Bônus com 2+ criaturas tapped |
| Embereth Paladin | 4 | Creature — Human Knight | R | tema | Haste + adamant contador |
| Éomer of the Riddermark | 5 | Creature — Human Knight | RW | tema | Token condicional (maior poder do board) |
| Éomer, Marshal of Rohan | 4 | Creature — Human Knight | RW | tema, wincon | Extra combat ao morrer lendária — combina com Adriana, Kwende, Syr Alin, Syr Carah, Tori |
| Fervent Cathar | 3 | Creature — Human Knight | R | tema, remoção | Haste + "não pode blocar" — remove bloqueador antes do alpha strike |
| Fireborn Knight | 4 | Creature — Human Knight | R | tema | Double strike dobra o valor de qualquer anthem |
| Hanweir Lancer | 3 | Creature — Human Knight | R | tema | First strike a si e ao par — combina com Kwende |
| Inspiring Captain | 4 | Creature — Human Soldier | W | tema | ETB pump de time; empilha com Tori no mesmo turno |
| Inspiring Veteran | 2 | Creature — Elf Warrior | W | tema | Anthem Knight permanente — peça central da tribal |
| Irreverent Gremlin | 2 | Creature — Gremlin | R | draw, tema | Menace + loot recorrente a cada criatura poder ≤2; alimenta o pacote de recursão |
| Knight Luminary | 4 | Creature — Human Knight | W | tema | Token ETB + Warp; feed para Valor in Akros, Vigilante Justice, Sanctuary Lockdown |
| Knight of Sorrows | 5 | Creature — Human Knight | W | tema | Corpo defensivo + afterlife; valor de token |
| Knight of Sursi | 4 | Creature — Human Knight | W | tema | Evasão (flanking) + suspend; tribal |
| Kwende, Pride of Femeref | 4 | Creature — Human Knight | RW | tema | Converte first strike em double strike no time — combo com Hanweir Lancer, Syr Alin, Youthful Knight, Warlord's Fury |
| Luxknight Breacher | 4 | Creature — Human Knight | W | tema | Contadores escalando com board wide |
| Mentor of the Meek | 3 | Creature — Human Soldier | W | draw, tema | Motor recorrente: compra a cada criatura poder ≤2 que entra — dispara com tokens do go-wide |
| Parhelion Patrol | 4 | Creature — Human Knight | W | tema | Mentor bota contador em atacante de menor poder |
| Relentless Rohirrim | 4 | Creature — Human Knight | W | tema | Ring tempts (pacote pequeno com Rohirrim Lancer) |
| Resistance Squad | 3 | Creature — Human Soldier | W | draw, tema | ETB "se controla outro Human, compre uma carta" — quase garantido no deck |
| Rohirrim Lancer | 1 | Creature — Human Knight | W | tema | Knight barato de 1cc, ajuda curva |
| Sheriff of Safe Passage | 3 | Creature — Human Soldier | W | tema | Contadores escalam com o board — payoff do go-wide |
| Syr Alin, the Lion's Claw | 5 | Creature — Human Knight | R | tema, wincon | Segundo "Tori" — anthem no ataque, empilha com Tori e dispara Éomer Marshal ao morrer |
| Syr Carah, the Bold | 5 | Creature — Human Knight | R | tema, draw, wincon | Impulse draw + ping repetível — motor de vantagem e finalização, lendária |
| Veteran Soldier | 2 | Creature — Human Soldier | W | tema | Concede ao comandante "cria tokens ao atacar se nenhum oponente tem mais vida" |
| Youthful Knight | 2 | Creature — Human Knight | W | tema | Knight vanilla barato; funciona com Kwende |

## Artefatos

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Boros Locket | 3 | Artifact | C | ramp, draw | Fixa R/W e sacrifica por 2 cartas |
| Commander's Sphere | 3 | Artifact | C | ramp, draw | Fixa qualquer cor da identidade + sac por carta |
| Manalith | 3 | Artifact | C | ramp | Fixa qualquer cor |
| Prying Blade | 1 | Artifact — Equipment | C | tema, ramp | +1/0 barato + treasure; ramp incidental |
| Seer's Lantern | 3 | Artifact | C | ramp | Rock incolor, não fixa R/W |
| The Circle of Loyalty | 6 | Artifact — Legendary | C | tema, wincon | Hub tribal: affinity Knight, anthem, token por lendária conjurada (7 lendárias no deck) |
| True-Faith Censer | 2 | Artifact — Equipment | C | tema | Vigilance + bônus Human — combina com Sanctuary Lockdown |

## Encantamentos

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Sanctuary Lockdown | 3 | Enchantment | W | tema | Anthem de Human — maioria dos Knights é "Human Knight" |
| Valor in Akros | 4 | Enchantment | W | tema | Pump a cada ETB de criatura — excelente com geradores de token |
| Vigilante Justice | 4 | Enchantment | W | tema | Ping por Human que entra — dispara com tokens Human Soldier |

## Instantâneos

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Adamant Will | 2 | Instant | W | proteção, tema | Combat trick pontual, protege 1 atacante |
| Bond of Discipline | 5 | Instant | W | proteção, wincon | Taps bloqueadores adversários — fog + habilita alpha strike |
| Disenchant | 2 | Instant | W | remoção | Interação genérica necessária |
| Djeru's Renunciation | 2 | Instant | W | proteção, remoção, draw(cycling) | Taps 2 bloqueadores pré-combate + cycling de fallback |
| Expose to Daylight | 3 | Instant | W | remoção | 2ª de 3 cópias quase idênticas de "destroy artifact or enchantment" — candidata a corte |
| Feat of Resistance | 2 | Instant | W | proteção, tema | Contador permanente + proteção de cor |
| Gideon's Triumph | 2 | Instant | W | remoção | Edict de combate |
| Invoke the Divine | 3 | Instant | W | remoção | 3ª cópia redundante de "destroy artifact or enchantment" — candidata a corte |
| Joust | 2 | Instant | W | tema, remoção | Fight com bônus para Knight |
| Miraculous Recovery | 5 | Instant | W | tema(recursão) | Reanimação a instant speed + contador |
| Reduce to Memory | 3 | Instant | W | remoção | Exila mas devolve corpo 3/2 ao oponente — remoção com downside |
| Swift Reckoning | 2 | Instant | W | remoção | Remove criatura tapped; flash com spell mastery |
| Warlord's Fury | 1 | Instant | W | tema, draw | First strike no time + cantrip — habilita Kwende |

## Feitiços

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Ambitious Assault | 3 | Sorcery | R | tema, draw(cond.) | +2/+0 no time + draw se controlar criatura "modified" |
| Basri's Solidarity | 2 | Sorcery | W | tema | Contador em todo o time — alimenta Ambitious Assault, Fireborn Knight, Kwende |
| Cosmotronic Wave | 4 | Sorcery | R | remoção(mini-wipe), wincon | 1 dano a tudo do oponente + "não pode blocar" |
| Crash Through | 1 | Sorcery | R | tema, draw | Trample no time inteiro + cantrip |
| Destructive Tampering | 3 | Sorcery | R | remoção, wincon | Remove artefato OU impede bloqueio |
| Emerge from the Cocoon | 5 | Sorcery | W | tema(recursão) | Parte do pacote de reanimação |
| Knight Watch | 5 | Sorcery | W | tema | 2 tokens Knight vigilance |
| Late to Dinner | 4 | Sorcery | W | tema(recursão) | Reanimação + Food |
| Remember the Fallen | 3 | Sorcery | W | draw, tema(recursão) | Devolve criatura/artefato à mão |
| Shoulder to Shoulder | 3 | Sorcery | W | tema, draw | Contadores em 2 alvos + cantrip |
| Tormenting Voice | 2 | Sorcery | R | draw | Discard+draw2 — habilita o pacote de recursão |

## Planeswalkers

*(nenhum)*

## Terrenos

| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
|---|---|---|---|---|---|
| Boros Guildgate | — | Land | RW | terreno | Dual tapped, sem upside |
| Looming Spires | — | Land | RW | terreno | Dual tapped com pump pontual no ETB |
| Memorial to War | — | Land | RW | terreno | ETB tapped, LD situacional caro |
| Sandstone Bridge | — | Land | RW | terreno | Dual tapped com pump pontual (vigilance) |
| Stone Quarry | — | Land | RW | terreno | Dual tapped, sem upside |
| Wind-Scarred Crag | — | Land | RW | terreno | Dual tapped com 1 de vida |
| Mountain | — | Land | R | terreno | 15x — fonte básica |
| Plains | — | Land | W | terreno | 15x — fonte básica |

## Histórico de swaps aplicados

| Fase | Sai | Entra | Motivo |
|---|---|---|---|
| Draw | Magnifying Glass | Mentor of the Meek | Rock incolor sem fixação → motor de draw recorrente com tokens |
| Draw | Angelic Gift | Irreverent Gremlin | Aura de risco → loot recorrente que alimenta a recursão |
| Draw | Culling Dais | Resistance Squad | Sac outlet isolado → draw ETB quase garantido (Human tribal) |
