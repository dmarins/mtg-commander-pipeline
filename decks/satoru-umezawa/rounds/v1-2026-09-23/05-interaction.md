# Interação — Satoru Umezawa (v1 · build)

Fase 5 · `interaction-specialist` · 2026-09-27. Oracle e rulings puxados nesta sessão via `bin/mtgdb`.
Preços: **LigaMagic (menor), via `mtgdb prices`, cotações de 2026-08-12 a 2026-09-27**. Carta sem cotação
registrada vai como **`a cotar`**, sem estimativa (regra 2). Nesta sessão não havia navegador para capturar preço.

Interação já no deck: **0/~10** · Wipes: **0/2–4** (`deck.md` só tem o comandante)
No pool temático (Fase 2): **8 remoções recicláveis via ninjutsu** + 1 proteção (Diversion Unit) + 2 premium fora do teto.

---

## 0. Ponto de partida

### O que o pool temático já entrega

| Carta (pool Fase 2) | O que responde | Custo real para usar | Preço |
|---|---|---|---|
| Ravenous Chupacabra | criatura (destroy) | ninjutsu 4 **ou** conjurar 4 | R$ 1,49 |
| Hostage Taker | criatura **ou artefato** (exile + roubo) | ninjutsu 4 ou conjurar 4 | R$ 1,25 |
| Noxious Gearhulk | criatura (destroy) | ninjutsu 4 ou conjurar 6 | R$ 1,54 |
| Meteor Golem | **qualquer permanente não-terreno** | ninjutsu 4 ou conjurar 7 | R$ 0,04 |
| Shriekmaw | criatura não-preta, não-artefato | **evoke 2 (remoção de mão)** ou ninjutsu 4 | R$ 0,39 |
| Dream Eater | bounce de permanente não-terreno | ninjutsu 4 ou **flash** 6 | R$ 0,55 |
| Marang River Regent | bounce de **2** permanentes não-terreno | ninjutsu 4 ou conjurar 6 | R$ 8,74 |
| Dragonlord Silumgar | roubo de criatura/PW | ninjutsu 4 ou conjurar 6 | R$ 7,35 |
| Archon of Cruelty · Massacre Wurm | edict+dreno · wipe unilateral | ninjutsu 4 | R$ 39,90 · R$ 32,96 — **fora do envelope** |

**Leitura:** a cobertura "de ETB" é boa e barata, mas toda ela pede **Satoru vivo + atacante conectado + 4 manas**,
ou 4–7 manas conjurando. Isso não cobre três situações: (1) a remoção **no Satoru antes da ativação** (o turno morre,
ruling oficial: só o ninjutsu *já ativado* sobrevive à saída dele); (2) a ameaça que precisa morrer **no turno do
oponente**; (3) o bloqueador voador que impede o ativador de conectar **antes** do ninjutsu existir. A interação de
mão abaixo é desenhada para esses três buracos.

### Critério de sinergia usado nesta fase (regra 3)

Os pontos que conto para uma peça de interação neste deck:
- **(J) Janela** — protege o Satoru entre a declaração de bloqueadores e a ativação do ninjutsu, ou antes do combate.
- **(M) Mana** — custa 0–2 e deixa os 4 do ninjutsu no mesmo turno (T6: 2 + 4).
- **(B) Bloqueador** — tira da frente o bloqueador voador/alcance que impede o ativador de conectar.
- **(R) Reciclagem** — devolve o próprio bichão ou ativador à mão para repetir o ETB.
- **(C) Cobertura** — responde a um tipo de ameaça que o pool temático não cobre (encantamento, instant no turno do oponente).
- **(K) Caixa** — custo zero de compra (não é ponto de sinergia, é desempate de orçamento).

### Radar do EDHREC

O `02-theme.md` desta rodada **não tem** a seção `Radar do EDHREC`, então fiz a chamada única
(`get_edhrec_recommendations`, limit 20) e registro aqui o que é da minha fase. Inclusão veio quebrada (`NaN%`),
só a nota de sinergia vale. **Orquestrador: vale mover este bloco para o `02-theme.md` em rodadas futuras.**

- **Instants:** An Offer You Can't Refuse, Arcane Denial, Infernal Grasp (0.13), Pongify (0.12), Deadly Rollick,
  Mana Drain, Swan Song, Go for the Throat, Withering Torment, Reality Shift, Rapid Hybridization, Aetherize,
  Spell Pierce, Murder, Force of Negation.
- **Sorceries:** Feed the Swarm, Toxic Deluge, Damnation, Blasphemous Edict, Enter the Enigma (não é interação).
- **Utility Artifacts:** Whispersilk Cloak (já no pool), Mithril Coat, Smoke Bomb, Dowsing Dagger.
- **Enchantments:** Witness Protection, Smoke Shroud, Cunning Evasion (evasão, não interação).
- **Top Cards:** Counterspell, Negate, Lightning Greaves (R$ 25,00), Swiftfoot Boots (R$ 14,00), Siren Stormtamer (0.41).
- **Game Changers:** Cyclonic Rift, Fierce Guardianship, Force of Will — **fora** por orçamento e cota de bracket.

---

## 1. Varredura da caixa (regra 7) — antes de qualquer busca

`mtgdb collection -list`: **182 sobressalentes**, das quais ~80 dentro da identidade UB/incolor. Todas as que tocam
interação foram fichadas com oracle desta sessão.

### Entram (8, custo zero)

| Carta | Por que entra (resumo da ficha — detalhe nas tabelas abaixo) |
|---|---|
| Theorix Charm | modal instant de 2: counter soft de não-criatura (J), −2/−2 no bloqueador voador x/2 (B), cantrip quando nada disso importa |
| Negate | hard counter de não-criatura por 2: a remoção típica no Satoru é instant/sorcery (J); responde artefato/encantamento/combo **no stack** (C) |
| Disruption Protocol | hard counter por `{U}{U}` + **virar um artefato** — equipamento parado (Winged Boots, Whispersilk, Costume) paga o custo sem perder nada (J, M) |
| Into the Roil | bounce instant de **qualquer** não-terreno (C), devolve o próprio bichão à mão para reciclar o ETB (R), salva o Satoru de exílio; kicker compra |
| Vraska's Final Mercy | `{B}{B}` destroy criatura/PW; sorcery-speed casa com o uso real: mata o bloqueador no 1º main e sobra mana para o ninjutsu (B, M) |
| Reasonable Doubt | counter soft (J) **+ suspect**: a criatura suspeita **não pode bloquear** enquanto estiver suspeita — o voador inimigo sai da defesa de forma permanente (B) |
| Old Fat Spider Can't See Me | hexproof ao Satoru **enquanto a saga estiver em campo** (3 turnos) (J); cap. II neutraliza o dano de uma criatura; III e IV compram 1 cada |
| Diversion Unit | já está no pool da Fase 2 como ativador; **conto aqui como proteção** (counter soft de instant/sorcery com `{U}` + sac) |

### Dispensadas, com motivo escrito

| Sobressalente | Ficha resumida | Motivo da dispensa (eixo) |
|---|---|---|
| Disdainful Stroke | F1 counter de MV ≥ 4 | F1: a remoção que mira o Satoru é quase toda MV ≤ 3 (Swords, Path, Infernal Grasp, Chaos Warp, Beast Within). Não cobre o ponto fraco declarado. 1 ponto (protege contra wipe MV ≥ 4). **Reserva 5.** |
| Extended Absence | F1 exile criatura/PW + 1 de dano a cada oponente, instant, 4 | F7: 4 manas disputam exatamente o orçamento do ninjutsu. **Deadly Rollick faz o mesmo exílio por 0** com o Satoru em campo. **Reserva 1** — é o substituto de custo zero se a Rollick estourar a cotação. |
| We Say Thee Nay! | F1 counter soft 2, ou 4 com teamwork 2 | Teamwork pode tapar o **próprio Satoru** (poder 2, não precisa estar desvirado para dar ninjutsu) — 2 pontos (J + corpo ocioso). Perde a vaga para Reasonable Doubt, que além de counter **remove um bloqueador**. **Reserva 3.** |
| Stoic Rebuttal | F1 hard counter 3, metalcraft −1 | F3: metalcraft pede 3 artefatos, contagem que o deck só atinge às vezes. Disruption Protocol faz o mesmo por 2 + equipamento parado. **Reserva 4.** |
| Vapor Snag | F1 `{U}` bounce de criatura, controlador perde 1 | 2 pontos (B + R). Perde para Into the Roil, que alcança não-criatura e tem kicker. **Reserva 2.** |
| Void Snare | F1 `{U}` sorcery bounce não-terreno | 2 pontos (B, R), mas sorcery: não protege o Satoru. Into the Roil cobre o mesmo alvo em instant. Reserva de fundo. |
| Thranduil's Decree | F1 hard counter 6 + conjurar o permanente | F6/F7: 6 manas no turno do oponente é o turno inteiro; não sobra mana de resposta no seu turno. |
| Trickster's Stratagem | F1 sorcery 4: o **dono** escolhe 2ª do topo ou fundo | F1: o dono põe em 2º e recompra; 4 manas sorcery. Não é remoção de verdade. |
| Culling Dais | F1 sac outlet + compra | Como interação, só salva de roubo/exílio sacrificando a própria peça. O deck recicla pela **mão**, não pelo cemitério. < 2 pontos. |
| Fog Bank | F2 defender voador, previne dano | Não ataca (não ativa ninjutsu); só defende o Satoru 2/4, que raramente morre em combate. 1 ponto. |
| Quiet Contemplation | F1 tapa criatura ao conjurar não-criatura | F3: o deck tem poucos não-criatura (~20); gatilho pouco frequente. 1 ponto. |
| Leonin Bola | F1 equipado tapa para tapar criatura | O Satoru ocioso poderia tapar o bloqueador (2 pontos: B + corpo ocioso), mas cada uso custa re-equip `{1}` sorcery. Anotado para o orquestrador; fora do top 10. |
| Lux Cannon · Thaumaton Torpedo | remoção de qualquer permanente | F6: 3 turnos carregando / 7 manas totais. Lento demais para bracket 3–4. |
| I Am Iron Man | F1 vira 4/4 voador + compra | Não é interação. Dá voo a um bichão de chão para reconectar (tema), mas precisa ser antes dos bloqueadores. Encaminhada ao orquestrador como habilitador, 1,5 ponto. |
| Dream Twist · Stern Lesson | mill / draw | Não são interação (Stern Lesson é da Fase 3). |
| Phyrexian Revoker | trava ativadas | Corpo sem evasão (Fase 2 já dispensou); como resposta a combo é 1 ponto. |
| **Ratchet Bomb · Blast Zone** | — | **Saíram da caixa**; não estão disponíveis. |

---

## Candidatas — remoção / proteção / counters

Legenda: J janela · M mana · B bloqueador · R reciclagem · C cobertura (critério na seção 0).
Ficha F1–F7 resumida em cada linha. **Nenhuma** delas foi cortada antes (`decisions.md` só tem o comandante).

### A. Interação de mão — 10 peças

| # | Carta | CMC | Tipo | Subcategoria | O que responde | Ficha F1–F7 + Sinergias (2+) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|---|
| 1 | Deadly Rollick | 4 (0) | Instant | remoção | criatura (**exile**) | F1 exila criatura; **grátis se você controla um comandante** (ruling: ninguém pode remover o Satoru em resposta para cobrar o custo). F6 joga do T3 em diante. F7 sem Satoru custa 4. **(M)** custo 0 com o Satoru em campo, que o plano já exige: remove e ainda ninjutsa no mesmo turno; **(B)** instant na etapa de ataque, antes dos bloqueadores, tira o voador da frente; exílio passa por indestrutível e comandante. `origem: EDHREC Instants` | não | a cotar ⚠ cotar primeiro |
| 2 | Infernal Grasp | 2 | Instant | remoção | criatura (destroy) | F1 destroy + perde 2 de vida. F6 T2. **(M)** 2 manas deixa os 4 do ninjutsu no T6; **(B)** mata o bloqueador no passo de ataque; incondicional (sem restrição de cor/artefato). `origem: EDHREC Instants 0.13` | não | a cotar |
| 3 | Withering Torment | 3 | Instant | remoção | criatura **ou encantamento** | F1 destroy criatura/encantamento, perde 2. **(C)** a única resposta de mão a **encantamento** do pacote: o pool temático não tem nenhuma além do Meteor Golem; **(B)** também mata bloqueador em instant. `origem: EDHREC Instants` | não | a cotar |
| 4 | Counterspell | 2 | Instant | counter | qualquer spell | F1 hard counter `{U}{U}`. **(J)** responde a remoção no Satoru, wipe e combo; **(M)** 2 manas no seu turno = ninjutsu + resposta com 6 terrenos. `origem: EDHREC Top Cards` | não | a cotar |
| 5 | Negate | 2 | Instant | counter | não-criatura | F1 hard counter de não-criatura. **(J)** a remoção no Satoru é spell não-criatura; **(C)** artefato/encantamento/combo **no stack** | **sim** | a cotar (custo 0) |
| 6 | Disruption Protocol | 2+ | Instant | counter | qualquer spell | F1 `{U}{U}` + tapar artefato desvirado **ou** `{1}`. F7: tapar rocha custa a mana dela. **(J)** hard counter; **(M)** equipamento parado (Winged Boots, Silver Shroud Costume, Whispersilk Cloak) paga o custo adicional de graça. | **sim** | R$ 0,20 (2026-08-12) |
| 7 | Theorix Charm | 2 | Instant | counter/remoção | não-criatura (soft) · criatura x/2 · cantrip | F1 três modos. **(J)** counter soft de 2 contra remoção no Satoru; **(B)** −2/−2 mata o voador x/2 (Faeries, Birds, Thopters) que segura os ativadores; modo 3 nunca fica morto | **sim** | a cotar (custo 0) |
| 8 | Into the Roil | 2 (4) | Instant | remoção flexível | **qualquer não-terreno** (bounce) | F1 bounce; kicker `{1}{U}` compra. **(C)** única resposta de mão a artefato **e** encantamento; **(R)** devolve o próprio bichão (Chupacabra, Gearhulk) para ninjutsá-lo de novo; **(J)** em último caso devolve o Satoru à mão (sai sem taxa, reconjura por 3) | **sim** | a cotar (custo 0) |
| 9 | Vraska's Final Mercy | 2 | Sorcery | remoção | criatura/PW (destroy) | F1 destroy criatura/PW, perde 2; modo 2 Jace +6 (dois `−3: draw`). **(B)** sorcery no 1º main é exatamente quando o bloqueador precisa morrer; **(M)** `{B}{B}` + 4 do ninjutsu no T6 | **sim** | a cotar (custo 0) |
| 10 | Reasonable Doubt | 2 | Instant | counter | qualquer spell (soft 2) + suspect | F1 counter unless `{2}`; suspect até uma criatura (menace + **não bloqueia**, enquanto suspeita). **(J)** contra a remoção no Satoru; **(B)** suspect no voador inimigo o tira da defesa **de forma permanente**; alternativa: dá menace a um ativador próprio | **sim** | a cotar (custo 0) |

### B. Proteção ao Satoru — contada à parte (5)

O ponto fraco declarado é a remoção **antes** da ativação. Os counters 4–7 e 10 já cobrem remoção por spell; as
peças abaixo cobrem o que counter não pega: **habilidade ativada/disparada** (Siren, ward) e o **turno em que não há
mana aberta** (hexproof/ward estáticos).

| # | Carta | CMC | Tipo | Subcategoria | O que responde | Ficha F1–F7 + Sinergias (2+) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|---|
| P1 | Winged Boots | 2 | Artifact — Equipment | proteção | alvo no Satoru (ward 4) | F1 flying + **ward {4}**, equip `{1}`. F3 artefato (metalcraft, Disruption Protocol). **(J)** ward 4 no Satoru é estático: protege mesmo sem mana aberta, e 4 manas a mais costuma inviabilizar a remoção; **(R)** equip 1 dá **voo** a um bichão de chão (Chupacabra, Hostage Taker, Meteor Golem): ele ataca, conecta e é trocado por ninjutsu, repetindo o ETB; **(M)** Disruption Protocol tapa a bota | não | a cotar |
| P2 | Silver Shroud Costume | 2 | Artifact — Equipment | proteção | alvo no Satoru (shroud instant) | F1 **flash**, entra anexando e dá **shroud até o fim do turno**; equipado inbloqueável; equip `{3}`. F3 artefato. **(J)** resposta **instantânea** à remoção no Satoru que o mantém em campo (o ninjutsu continua possível); **(R)** depois, equip 3 torna um bichão inbloqueável para reciclar o ETB; **(M)** paga Disruption Protocol | não | a cotar |
| P3 | Siren Stormtamer | 1 | Creature — Siren Pirate Wizard | proteção | spell **ou habilidade** que mira você/sua criatura | F1 flying 1/1; `{U}`, sac: counter. F2 corpo voador de 1. F4 Tetsuko (poder 1 → inbloqueável). **(J)** única peça que counta **habilidade** (ETB de remoção, ativada) mirando o Satoru, por `{U}`; **(tema)** é ativador voador de 1 mana. F7: se for o atacante devolvido pelo ninjutsu, deixa de proteger; use outro ativador. `origem: EDHREC Top Cards 0.41` | não | a cotar |
| P4 | Old Fat Spider Can't See Me | 3 | Enchantment — Saga | proteção | alvo no Satoru (hexproof) | F1 I: hexproof enquanto a saga estiver em campo (3 turnos seus); II: previne o dano de uma criatura; III e IV: compra 1 cada. F7 sorcery-speed, T3–4. **(J)** hexproof de vários turnos sem mana aberta; **(draw)** +2 cartas | **sim** | R$ 0,88 (2026-08-12) |
| P5 | Diversion Unit | 2 | Artifact Creature — Robot | proteção | instant/sorcery (soft 3) | Já no pool da Fase 2 (ativador voador). Aqui conta como proteção: `{U}` + sac counta a remoção no Satoru a menos que paguem 3. | **sim** | R$ 0,45 (2026-08-12) |

Somando os counters 4–7 e 10, o Satoru tem **10 peças que o protegem**. Mais de uma cobre cada tipo de ataque:
spell (counters, Costume), habilidade (Stormtamer, ward), turno sem mana (Spider, Boots).

---

## Candidatas — board wipes

O deck vive de criaturas **pequenas e baratas** + bichões **que voltam à mão**. Os três wipes propostos são
**unilaterais** ou **bounce** (que devolve os bichões para serem ninjutsados de novo). Nenhum destrói o próprio campo.

| # | Carta | CMC | Simétrico? | Ficha F1–F7 + Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|
| W1 | Scourge of Fleets | 7 (ninjutsu 4) | **Não**: só criaturas dos oponentes | F1 ETB: devolve **toda criatura dos oponentes** com resistência ≤ nº de **Islands** que você controla (X no resolve). F2 6/6 Kraken. **(tema)** alvo de ninjutsu: o ETB não depende de ter sido conjurado; **(B)** na etapa de dano de combate da mesma entrada, a mesa fica sem bloqueador para o turno seguinte; **(sinergia W3)** é Kraken, sobrevive à Whelming Wave. **Dependência da Fase 6:** precisa de Islands (básicas ou duais com tipo Island). Com ~8+ ela varre quase tudo. Estava na reserva da Fase 2 | não | a cotar |
| W2 | Archfiend of Depravity | 5 (ninjutsu 4) | **Não** | F1 flying 5/4; **no end step de cada oponente**, ele escolhe até 2 criaturas e sacrifica o resto. **(tema)** alvo de ninjutsu que não depende de "cast"; voador, então fica atacando e conectando; **(B)** cada oponente fica com ≤ 2 bloqueadores. Regra 11: não é lock (o oponente segue jogando e mantém 2 criaturas), mas pune mesas de fichas. **Sinalizo ao orquestrador**; é padrão em bracket 3–4. | não | a cotar ⚠ cotar primeiro |
| W3 | Whelming Wave | 4 | Sim (bounce), **exceto Krakens/Leviathans/Octopuses/Serpents** | F1 sorcery: devolve todas as criaturas à mão, menos as dos 4 tipos. **(R)** os bichões voltam à mão = **munição de ninjutsu** (4 manas cada); os ativadores custam 0–2 para recolocar; **(tema)** Gyruda e Scourge of Fleets (Krakens) **ficam**; fichas dos oponentes somem de vez. O Satoru volta à mão (não é conjurado da zona de comando, então **sem taxa**) e é reconjurado por 3. F7: custa o turno de reconstrução, então use quando estiver atrás na mesa. | não | a cotar |

**Contagem de wipes:** 3 propostos (meta 2–4). Massacre Wurm (pool, R$ 32,96) seria o 4º ideal (unilateral + dreno),
mas não cabe no envelope.

---

## Cobertura de ameaças

| Ameaça | De mão (proposta) | Via ninjutsu (pool) | Leitura |
|---|---|---|---|
| **Criaturas** | Deadly Rollick, Infernal Grasp, Withering Torment, Vraska's Final Mercy, Theorix Charm (x/2), Into the Roil (bounce), Reasonable Doubt (neutraliza bloqueio) | Chupacabra, Gearhulk, Shriekmaw (+ evoke 2), Hostage Taker, Silumgar, Meteor Golem, Dream Eater, Marang | **Sobra.** |
| **Artefatos** | Into the Roil (bounce); Negate/Counterspell/Disruption no stack | Hostage Taker, Meteor Golem, Dream Eater, Marang | ⚠ **Buraco parcial:** nenhuma resposta de mão **permanente** a artefato já em campo. Reserva 6 (Resculpt) fecha. |
| **Encantamentos** | Withering Torment, Into the Roil; counters no stack | Meteor Golem, Dream Eater, Marang | ok (2 de mão + 3 de ETB) |
| **Planeswalkers** | Vraska's Final Mercy; counters | Meteor Golem, Silumgar (rouba), Dream Eater/Marang (bounce) | ok |
| **Combos/spells** | Counterspell, Negate, Disruption Protocol (hard) · Theorix, Reasonable Doubt (soft) | — | ok: 3 hard + 2 soft; sem counter a habilidade ativada (Stormtamer só protege o seu lado) |
| **Flexível** ("qualquer não-terreno") | Into the Roil | Meteor Golem, Dream Eater, Marang River Regent | ok |
| **Remoção no Satoru** | 5 counters + Stormtamer + Costume (instant) · Boots + Spider (estáticos) + Diversion Unit | — | **10 peças**, cobrindo spell, habilidade e o turno sem mana |
| **Wipe dos oponentes** | Counterspell, Disruption, Negate (e Disdainful Stroke na reserva) | — | o deck se refaz rápido: ativadores são baratos e bichões na mão voltam por ninjutsu |

---

## Orçamento da fase

| Carta | Origem | Preço |
|---|---|---|
| Theorix Charm, Negate, Disruption Protocol, Into the Roil, Vraska's Final Mercy, Reasonable Doubt, Old Fat Spider Can't See Me, Diversion Unit | caixa | **R$ 0,00** |
| Deadly Rollick ⚠ | compra | a cotar |
| Infernal Grasp | compra | a cotar |
| Withering Torment | compra | a cotar |
| Counterspell | compra | a cotar |
| Winged Boots | compra | a cotar |
| Silver Shroud Costume | compra | a cotar |
| Siren Stormtamer | compra | a cotar |
| Scourge of Fleets | compra | a cotar |
| Archfiend of Depravity ⚠ | compra | a cotar |
| Whelming Wave | compra | a cotar |
| **Total das compras** | | **a cotar: 10 cartas, nenhuma com cotação no `prices.tsv`**. Não estimo (regra 2). |

**Ordem de cotação e de recuo, se o envelope de ~R$ 25–30 estourar.** Cada compra tem um substituto de custo zero
**na caixa**, ou um corte que não abre buraco:

| Compra | Prioridade | Se estourar, recua para | O que se perde |
|---|---|---|---|
| Deadly Rollick | 1 | Extended Absence (caixa) | o custo 0: passa a disputar 4 manas com o ninjutsu |
| Counterspell | 1 | Stoic Rebuttal (caixa) | 1 mana a mais sem metalcraft |
| Winged Boots | 1 | Whispersilk Cloak (pool, R$ 10,00) no Satoru | ward 4 → shroud; perde o equip barato de 1 |
| Scourge of Fleets | 1 | Aetherize (reserva) | o wipe via ninjutsu vira wipe defensivo de mão |
| Infernal Grasp | 2 | Vapor Snag (caixa) | a remoção vira tempo (bounce) |
| Whelming Wave | 2 | Evacuation (reserva, instant) ou nenhum (fica com 2 wipes) | — |
| Withering Torment | 2 | Void Snare (caixa) | a resposta **permanente** a encantamento vira temporária |
| Siren Stormtamer | 3 | Diversion Unit já cobre spell; Thrummingbird (caixa) cobre o slot de ativador | counter de **habilidade** |
| Silver Shroud Costume | 3 | We Say Thee Nay! (caixa) | proteção instant **fora** do counter |
| Archfiend of Depravity | 3 | nenhum: fica com 2 wipes (W1 + W3) | o wipe recorrente |

Com as 10 compras recuando para a caixa, a fase fecha em **R$ 0** com 10 peças + 5 proteções + 1 wipe. É o piso.

---

## Reservas

| # | Carta | Na coleção? | Preço | Papel |
|---|---|---|---|---|
| 1 | Extended Absence | **sim** | custo 0 | exílio instant de criatura/PW; recuo da Deadly Rollick |
| 2 | Vapor Snag | **sim** | custo 0 | bounce `{U}` de bloqueador / recicla o próprio bichão |
| 3 | We Say Thee Nay! | **sim** | custo 0 | counter soft 4 com teamwork tapando o Satoru |
| 4 | Stoic Rebuttal | **sim** | R$ 0,20 | hard counter 3 |
| 5 | Disdainful Stroke | **sim** | custo 0 | anti-wipe MV ≥ 4 |
| 6 | Resculpt | não | a cotar | **fecha o buraco de artefato**: exila artefato/criatura em instant por 2; a ficha 4/4 que dá é de chão e não bloqueia os voadores/inbloqueáveis |
| 7 | Aetherize | não | a cotar | wipe defensivo instant **unilateral** (só atacantes), no turno do oponente, com a mana que o seu turno não usa. `origem: EDHREC Instants` |
| 8 | Dive Down | não | a cotar | `{U}` hexproof +0/+3: mantém o Satoru em campo (o ninjutsu segue), 1 mana |
| 9 | Keep Safe | não | a cotar | counter de spell que mira permanente seu + compra, 2 |
| 10 | Spider-Sense | não | a cotar | counter de instant/sorcery/**gatilho**; web-slinging `{U}` devolve um bichão **tapado** que atacou (recicla o ETB) |
| 11 | Evacuation | não | a cotar | Whelming Wave em instant (5): também devolve os bichões à mão |
| — | Swiftfoot Boots | não | R$ 14,00 (2026-09-16) | hexproof no Satoru; perde para Winged Boots, que tem equip 1 e dá voo para reciclar |
| — | Lightning Greaves | não | R$ 25,00 (2026-09-16) | fora do envelope sozinha |
| — | Toxrill, the Corrosive · Toxic Deluge · Cyclonic Rift | não | a cotar ⚠ | premium: Toxrill é wipe unilateral via ninjutsu, mas o histórico sugere preço alto. Rift é Game Changer |

---

## Pendências para o orquestrador

1. **Cotar as 10 compras** na LigaMagic, começando pelas ⚠ (Deadly Rollick, Archfiend of Depravity). Recuo por carta na tabela de orçamento.
2. **Fase 6:** a Scourge of Fleets conta **Islands**. A manabase deve privilegiar básicas Island e duais com tipo Island.
3. **Sobreposição de slots:** Siren Stormtamer (ativador), Winged Boots e Silver Shroud Costume (evasão para reciclar), Scourge of Fleets e Archfiend of Depravity (alvos de ninjutsu) e Diversion Unit (pool) ocupam slots temáticos também. As 18 peças propostas custam ~12 slots líquidos.
4. **Regra 11:** Archfiend of Depravity é sacrifício recorrente unilateral. Não é lock, mas vale confirmar com o usuário.
5. **Radar do EDHREC** não estava no `02-theme.md`. Registrei aqui as seções de Instants/Sorceries/Utility Artifacts/Enchantments.
