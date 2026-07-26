# Análise Temática — Auditoria (modo improve)

## Comandante — análise linha a linha

**Tori D'Avenant, Fury Rider** — {1}{R}{R}{W}, mv4, Lendária Humano Cavaleiro (Human Knight), 3/3.

| Linha/habilidade | Gatilho/termo | O que habilita |
|---|---|---|
| Vigilance | keyword `vigilance` | Ataca sem travar bloqueadores/defesa; empilha com "untap white attackers" (redundância parcial) |
| Trample | keyword `trample` | Ela mesma ignora chump blocks |
| "Whenever Tori attacks, all other attacking creatures you control get +1/+1 until end of turn" | gatilho `attacks` + anthem `+1/+1 until end of turn` | Núcleo do plano go-wide: quanto mais criaturas atacando, mais dano total. Premia tokens e presença ampla no board |
| "Other red attacking creatures you control gain trample until end of turn" | gatilho `attacks` + cor `red` | Remove necessidade de evasão nas criaturas vermelhas; sinergiza com pump (mais poder = mais trample relevante) |
| "Untap each other white attacking creature you control" | gatilho `attacks` + cor `white` | Pseudo-vigilance para brancos; permite blocar de volta e (com combos de extra combat) reatacar |

**Termos de busca do tema**: `o:"whenever ~ attacks"` / `o:"other attacking creatures"`, anthem `o:"+1/+1 until end of turn"`, `o:trample`, `kw:vigilance`, tribal `t:knight`/`t:human`, tokens `o:"create a"`, `+1/+1 counter` (`o:"+1/+1 counter"`), extra combat `o:"additional combat phase"`.

## Auditoria por categoria — classificação das 69 cartas não-básicas

Comandante à parte. Total do deck atual: 30 terrenos básicos (15 Mountain / 15 Plains) + **6 terrenos utilitários** (não 5 — havia um esquecido: Memorial to War também é terreno) + 63 mágicas não-terrestres = 99 + comandante = 100.

### Curva de mana (63 cartas não-terrestres)

| CMC | 1 | 2 | 3 | 4 | 5 | 6 |
|---|---|---|---|---|---|---|
| Qtde | 4 | 17 | 17 | 14 | 9 | 2 |

CMC médio das não-terrestres: **3.21**. Curva razoável no meio (2–3), mas **só 4 one-drops** para um plano de ataque que quer pressão cedo, e **25 cartas (40%) custam 4+** — pesado para um agro go-wide.

### Tabela completa

| Carta | CMC | Categorias | Sinergia | Nota |
|---|---|---|---|---|
| Adamant Will | 2 | proteção, tema | média | Combat trick pontual; protege 1 atacante, mas não synergiza além disso |
| Adriana, Captain of the Guard | 5 | tema, wincon | forte | Melee dá +1/+1 por oponente atacado e estende a todos — combina direto com o anthem de Tori em multiplayer |
| Ambitious Assault | 3 | tema, draw(cond.) | forte | +2/+0 no time + draw se controlar criatura "modified" (contadores contam) — sinergiza com o pacote de contadores (Basri's Solidarity, Feat of Resistance, Miraculous Recovery) |
| Angelic Gift | 2 | tema, draw | média | Cantrip + evasão pontual; carta-aposta (perde valor se a criatura morrer) |
| Basri's Solidarity | 2 | tema | forte | Contador em todo o time — alimenta Ambitious Assault (modified), Fireborn Knight, Kwende (first strike→double strike) |
| Blazing Blade Askari | 3 | tema | média | Só tribal (Knight) + flanking; nenhuma sinergia adicional |
| Bogardan Lancer | 2 | tema | média | Bloodthirst + flanking; tribal apenas |
| Bond of Discipline | 5 | proteção, wincon | forte | Taps todos os bloqueadores adversários — fog + habilita alpha strike; funciona muito bem com o gatilho de ataque de Tori |
| Boros Guildgate | 0 | terreno | — | Dual tapped, sem upside |
| Boros Locket | 3 | ramp, draw | média | Fixa R/W e sacrifica por 2 cartas; um dos poucos rocks que ajuda a trava de mana |
| Cavalry Drillmaster | 2 | tema | média | Combat trick em corpo de criatura; tribal |
| Commander's Sphere | 3 | ramp, draw | média | Fixa qualquer cor da identidade + sac por carta; bom para fixação |
| Cosmotronic Wave | 4 | remoção(mini-wipe), wincon | forte | 1 dano a tudo do oponente + "can't block" — limpa X/1 e habilita alpha strike |
| Crash Through | 1 | tema, draw | forte | Trample no time inteiro + cantrip; barato e sinergiza com o pacote de trample de Tori |
| Culling Dais | 2 | draw | fraca | Sem outro sac outlet no deck além de tokens; motor de draw lento e isolado |
| Dawnstrike Vanguard | 6 | tema | média | Precisa de 2+ criaturas *tapped*; conflita parcialmente com vigilance/untap de Tori, mas funciona com atacantes não-vigilantes |
| Destructive Tampering | 3 | remoção, wincon | forte | Remove artefato OU impede bloqueio (sem voadoras) — habilita alpha strike do go-wide |
| Disenchant | 2 | remoção | média | Interação genérica necessária, mas é a 1ª de 3 cópias quase idênticas na lista (ver Expose to Daylight, Invoke the Divine) |
| Djeru's Renunciation | 2 | proteção, remoção, draw(cycling) | forte | Taps 2 bloqueadores pré-combate + cycling de fallback — ótimo enabler de alpha strike |
| Embereth Paladin | 4 | tema | média | Haste + adamant contador; tribal, sem sinergia cruzada |
| Emerge from the Cocoon | 5 | tema(recursão) | média | Parte do pacotinho de reanimação (com Late to Dinner, Miraculous Recovery, Remember the Fallen) |
| Éomer of the Riddermark | 5 | tema | média | Token condicional (maior poder do board); go-wide, mas condição nem sempre é fácil de garantir |
| Éomer, Marshal of Rohan | 4 | tema, wincon | forte | Extra combat ao morrer lendária — combina com Adriana, Kwende, Syr Alin, Syr Carah, a própria Tori; peça-chave de vitória |
| Expose to Daylight | 3 | remoção | fraca | 3ª cópia funcional de "destroy artifact or enchantment" — redundante, candidata a corte |
| Feat of Resistance | 2 | proteção, tema | forte | Contador permanente + proteção de cor — sinergiza com o pacote de contadores e protege atacante-chave |
| Fervent Cathar | 3 | tema, remoção | forte | Haste + "target can't block" — remove bloqueador antes do alpha strike |
| Fireborn Knight | 4 | tema | forte | Double strike dobra o valor de qualquer anthem (Tori, Syr Alin, Inspiring Captain, Basri's Solidarity) |
| Gideon's Triumph | 2 | remoção | média | Edict de combate; cláusula de "Gideon planeswalker" é morta (deck não tem planeswalkers) |
| Hanweir Lancer | 3 | tema | forte | First strike a si e ao par — combina com Kwende (first strike vira double strike) |
| Inspiring Captain | 4 | tema | forte | ETB pump de time; empilha com o gatilho de ataque de Tori no mesmo turno |
| Inspiring Veteran | 2 | tema | forte | Anthem Knight permanente — peça central da tribal |
| Invoke the Divine | 3 | remoção | fraca | 2ª cópia redundante de "destroy artifact or enchantment"; candidata a corte |
| Joust | 2 | tema, remoção | forte | Fight com bônus para Knight — remoção que usa o próprio time pumped |
| Knight Luminary | 4 | tema | forte | Token ETB + Warp (reaproveita o ETB); token feed para Valor in Akros, Vigilante Justice, Sanctuary Lockdown |
| Knight of Sorrows | 5 | tema | média | Corpo defensivo + afterlife; menos alinhado com "atacar", mas ainda gera valor de token |
| Knight of Sursi | 4 | tema | média | Evasão (flying/flanking) + suspend; tribal |
| Knight Watch | 5 | tema | forte | 2 tokens Knight vigilance — alimenta Inspiring Veteran, Circle of Loyalty, Valor in Akros |
| Kwende, Pride of Femeref | 4 | tema | forte | Converte first strike em double strike em todo o time — combo com Hanweir Lancer, Syr Alin, Youthful Knight, Warlord's Fury |
| Late to Dinner | 4 | tema(recursão) | média | Reanimação + Food; parte do pacote de recursão |
| Looming Spires | 0 | terreno | — | Dual tapped com pump pontual no ETB (ajuda um pouco de tempo, mas não fixa 2 cores) |
| Luxknight Breacher | 4 | tema | forte | Contadores escalando com board wide — cresce com o próprio plano go-wide |
| Magnifying Glass | 3 | ramp, draw | fraca | Rock incolor — não fixa R/W (agrava a trava de mana); investigate é draw lento |
| Manalith | 3 | ramp | média | Fixa qualquer cor — ajuda a trava de mana, mas é só 1 fonte |
| Memorial to War | 0 | terreno | — | ETB tapped, quase sem upside (LD situacional caro); é o 6º terreno utilitário — muitas vezes esquecido na contagem |
| Miraculous Recovery | 5 | tema(recursão) | forte | Reanimação a instant speed + contador — melhor peça do trio de recursão |
| Parhelion Patrol | 4 | tema | forte | Mentor bota contador em atacante de menor poder — ótimo com exército de tokens pequenos |
| Prying Blade | 1 | tema, ramp | média | +1/0 barato + treasure; ramp incidental fraco |
| Reduce to Memory | 3 | remoção | fraca | Exila mas devolve corpo 3/2 ao oponente — remoção com downside relevante |
| Relentless Rohirrim | 4 | tema | fraca | Ring tempts é pacote de 2 cartas só (com Rohirrim Lancer) — não desenvolve, é ruído temático |
| Remember the Fallen | 3 | draw(card advantage), tema(recursão) | média | Devolve criatura/artefato à mão — reabastece o pacote de recursão |
| Rohirrim Lancer | 1 | tema | média | Knight barato de 1cc, curva; Ring tempts é upside menor isolado |
| Sanctuary Lockdown | 3 | tema | forte | Anthem de Human — a maioria dos Knights do deck é "Human Knight", overlap real com a tribal |
| Sandstone Bridge | 0 | terreno | — | Dual tapped com pump pontual (vigilance) no ETB |
| Seer's Lantern | 3 | ramp | fraca | Rock incolor, não fixa — mesmo problema de Magnifying Glass |
| Sheriff of Safe Passage | 3 | tema | forte | Contadores escalam com o board — payoff direto do go-wide |
| Shoulder to Shoulder | 3 | tema, draw | forte | Contadores em 2 alvos + cantrip — dupla sinergia (contadores + draw) |
| Stone Quarry | 0 | terreno | — | Dual tapped, idêntico a Boros Guildgate |
| Swift Reckoning | 2 | remoção | média | Remove criatura tapped; flash com spell mastery (deck tem instants/sorceries suficientes) |
| Syr Alin, the Lion's Claw | 5 | tema, wincon | forte | Segundo "Tori" — anthem no ataque, empilha com Tori e dispara Éomer Marshal ao morrer |
| Syr Carah, the Bold | 5 | tema, draw, wincon | forte | Impulse draw + ping repetível — motor de vantagem e finalização, também lendária (Éomer Marshal) |
| The Circle of Loyalty | 6 | tema, wincon | forte | Hub tribal: affinity para Knight, anthem, token por lendária conjurada (7 lendárias no deck!) e ativada |
| Tormenting Voice | 2 | draw | forte | Discard+draw2 — habilita o pacote de recursão (descarta alvo para reanimar depois) |
| True-Faith Censer | 2 | tema | forte | Vigilance + bônus Human — combina com Sanctuary Lockdown e a tribal Human/Knight |
| Valor in Akros | 4 | tema | forte | Pump a cada ETB de criatura — excelente com todos os geradores de token (Knight Luminary, Sheriff, Circle of Loyalty) |
| Veteran Soldier | 2 | tema | forte | Concede ao comandante "cria tokens ao atacar se nenhum oponente tem mais vida" — sinergia direta com Tori atacando |
| Vigilante Justice | 4 | tema | forte | Ping por Human que entra — dispara com os vários tokens Human Soldier gerados no deck |
| Warlord's Fury | 1 | tema, draw | forte | First strike no time + cantrip — barato, habilita Kwende (first strike→double strike) |
| Wind-Scarred Crag | 0 | terreno | — | Dual tapped com 1 de vida; sem upside relevante |
| Youthful Knight | 2 | tema | média | Knight vanilla barato; funciona com Kwende (first strike) mas isolado fora disso |

## Resumo de contagens (vs. metas do pipeline)

| Categoria | Atual | Meta | Gap |
|---|---|---|---|
| Draw | ~10 (várias são cantrip único, não motor sustentado) | 12–13 | -2/-3, e qualidade baixa |
| Ramp | 6 fontes (só 3 fixam cor: Locket, Sphere, Manalith) | 10–11 padrão + 2–3 explosivo | -4/-5 padrão; **0 explosivo** (sem Sol Ring/Signets/fast mana) |
| Remoção | ~11 peças, mas 3 redundantes (Disenchant/Expose to Daylight/Invoke the Divine) e o resto é situacional (tapped, edict, fight) | ~10 | numericamente ok, qualidade fraca — falta remoção incondicional |
| Wipe | 0 (Cosmotronic Wave é mini-wipe de 1 dano, não conta) | 2–4 | -2/-4, gap crítico |
| Wincon | 1 caminho real (dano de combate via go-wide); Éomer Marshal e Circle of Loyalty reforçam esse único caminho | 3+ | faltam 2+ caminhos alternativos |
| Terrenos | 36 total (30 básicos + **6** utilitários, todos ETB tapped) | ~38 ajustado pela curva (CMC médio 3.21, pesado em 4+) | provavelmente precisa **subir**, não descer — poucas fontes tapped sem fixação real |

## Candidatas a corte (sinergia fraca/nula)

- **Expose to Daylight** e **Invoke the Divine** — 2ª/3ª cópia de "destroy artifact or enchantment" (redundante com Disenchant).
- **Reduce to Memory** — remoção com downside (devolve corpo 3/2 ao oponente).
- **Relentless Rohirrim** — pacote "Ring tempts" tem só 2 cartas, não desenvolve.
- **Culling Dais** — sac outlet isolado, sem outra sinergia de sacrifício no deck.
- **Magnifying Glass** e **Seer's Lantern** — rocks incolores que não resolvem a fixação R/W (agravam a trava de mana).
- **Boros Guildgate** e **Stone Quarry** — duals tapped puros sem upside; primeiros candidatos a substituir por fixação melhor.
- **Prying Blade**, **Gideon's Triumph** — impacto marginal, sinergia única.

## Problemas mais graves de manabase/curva

1. **Fixação insuficiente**: só 6 terrenos não-básicos, todos **entram tapped** e nenhum é untapped/shock/check/fast land — para um agro que quer atacar cedo, isso trava o tempo justamente nos turnos 2–3 mais importantes.
2. **Ramp quase inexistente e sem explosivo**: apenas 6 fontes de ramp, 3 delas incolores (não ajudam R/W), e zero fast mana (sem Sol Ring, Signets, Talisman) — a "trava de mana" relatada é estrutural, não azar.
3. **Terrenos totais provavelmente baixos para a curva**: CMC médio 3.21 com 40% do deck em 4+ pede algo perto de 38 terrenos ou mais ramp — hoje são 36, e a fixação fraca faz esse número render menos ainda.
4. **Curva com poucos 1-drops (só 4)** para um plano de pressão constante via ataques em massa — o deck é lento para começar a atacar mesmo quando a mana cresce.
