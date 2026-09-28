# Vantagem de Cartas — Satoru Umezawa (v1 · build)

Fase 3 · `draw-specialist` · 2026-09-27. Oracle e rulings puxados nesta sessão via `bin/mtgdb` (regra 6).
Preços: **LigaMagic (menor), via `mtgdb prices`, cotações de 2026-09-27** (caixa: cotações de 2026-08-12 a
2026-09-24). O que não tem cotação vai como **`a cotar`**, sem estimativa (regra 2). Não havia navegador nesta sessão.

Fontes já no deck: **1/12–13** em `deck.md` (só o comandante).
Contando o pool temático da Fase 2 (só as peças baratas que devem sobreviver ao teto): **10/12–13**.
Com as 4 entradas propostas abaixo: **13/12–13** (14 se a Marang contar inteira; ver nota).

---

## 1. O que o tema já entrega (contado antes de buscar)

Critério: só conta carta que dá acesso a carta **extra**. Premium acima do teto (Baleful Strix R$ 59,99, Ninja of
the Deep Hours R$ 29,00, Fallen Shinobi R$ 30,60, Ancient Silver Dragon R$ 263,99, Archon R$ 39,90,
Kaito, Bane of Nightmares R$ 92,00) **fica fora da conta**, como o orquestrador pediu.

| # | Carta | CMC | Tipo de fonte | Como gera vantagem (oracle de hoje) | Preço |
|---|---|---|---|---|---|
| 1 | Satoru Umezawa (comandante) | 3 | **repetível**, 1×/turno | Toda ativação de ninjutsu: olha 3, põe 1 na mão. +1 carta por turno de ninjutsu, de graça sobre a mana do ninjutsu. | comandante (fora da régua) |
| 2 | Satoru, the Infiltrator | 2 | **repetível** | Criatura que entra sem ter sido conjurada → compra. Todo ninjutsu (do comandante ou nativo) dispara. | R$ 2,63 |
| 3 | Moon-Circuit Hacker | 2 | **repetível** | Conecta → compra; no turno em que entra por ninjutsu `{U}`, compra **sem** descartar. | R$ 3,95 |
| 4 | Ingenious Infiltrator | 4 | **repetível** | Todo **Ninja** que conecta → compra (os dois Satorus, Moon-Circuit, Changeling, ficha do Kaito Shizuki). | R$ 5,62 |
| 5 | Spectral Sailor | 1 | **repetível** (sumidouro) | `{3}{U}`: compra. Flash: usa a mana do turno do oponente. | R$ 0,49 |
| 6 | Kaito, Dancing Shadow | 4 | **repetível** | 0: compra. O estático devolve quem conectou à mão e libera duas ativações no turno. | R$ 1,85 |
| 7 | Reconnaissance Mission | 4 | **repetível** (motor) | Toda criatura que conecta → compra. Cycling `{2}` quando sobra. | caixa |
| 8 | Bident of Thassa | 4 | **repetível** (motor) | Mesmo motor da Mission; artefato **e** encantamento. | caixa (R$ 5,80 em 08-12) |
| 9 | Mulldrifter | 5 (evoke 3) | pontual **reciclável** | ETB compra 2; como flier volta à mão pelo próximo ninjutsu e repete. | R$ 1,62 |
| 10 | Marang River Regent // Coil and Catch | 6 / 4 | pontual (meia fonte) | A face Omen é **instant**: compra 3, descarta 1 (+1 líquido) e embaralha de volta. A face criatura não compra. | R$ 8,74 |

**Fora da conta, com motivo:** Agent of Treachery (R$ 17,09) compra 3 só com 3+ permanentes alheias em campo —
condicional demais para meta; conta como bônus se a Fase 6 o mantiver. Rune-Scarred Demon é tutor, não compra.
Dream Eater e Faerie Seer são seleção (surveil/scry), não carta extra.

**Leitura do plano:** a mana do turno de ataque vai para o ninjutsu. As 10 acima já cobrem bem o "draw que vem de
conexão/ETB". O que falta é (a) **draw instantâneo** para a mana do turno do oponente, (b) **um pico** de compra
que não dependa de 3 criaturas conectando, e (c) redundância barata caso a Fase 6 corte peças do pool.

---

## 2. Sobressalentes — varredura da caixa (regra 7)

`mtgdb collection -list`: **182 cartas**, 79 dentro da identidade UB/incolor. Varri todas as que têm texto ou tag de
compra/seleção **antes** de qualquer busca. Ficha F1–F7 em cada linha; a dispensa diz o eixo.

**Entram como candidatas (2):**

| Carta | Por que entra | Preço |
|---|---|---|
| Thirst for Knowledge | Instant de 3: compra 3, descarta 2 (ou 1 artefato). +1 líquido mínimo, +2 descartando artefato. | caixa (R$ 0,05 em 08-12) |
| H.E.R.B.I.E. Scout Unit | Flier artefato 2/1 com ETB "compra + terreno da mão". Ativador cujo ETB se repete a cada volta à mão pelo ninjutsu. | caixa (a cotar) |

**Reservas da caixa (custo zero, entram se faltar slot barato):** Sphinx's Approach, Stern Lesson — ver §4.

**Encaminhadas a outras fases (servem, mas o papel principal não é draw):**

| Sobressalente | Fase | Por quê |
|---|---|---|
| Old Fat Spider Can't See Me | 5 (proteção) | Cap. I dá **hexproof ao Satoru** enquanto a saga durar; III e IV compram 1 cada. Se a Fase 5 a escolher, **conta +1 fonte de draw** aqui. |
| Vraska's Final Mercy | 5 (remoção) | Modal: destrói criatura/PW **ou** Jace com 6 de lealdade (−3 compra, dois turnos). Remoção que vira 2 cartas quando não há alvo. |
| Into the Roil | 5 (remoção) | Bounce + compra com kicker. Também devolve o **próprio bichão** à mão para repetir o ETB. |
| Theorix Charm | 5 (counter/remoção) | O modo "mill 3 + compra 1" é cantrip (troca 1 por 1): não conta para a meta. O valor está nos outros dois modos. |
| Trickster's Stratagem | 5 (remoção) | Tuck de criatura + connive (loot). O connive não é carta extra. |
| Commander's Sphere | 4 (ramp) | Sacrifício compra 1, mas é rocha. |

**Dispensadas, com motivo escrito:**

| Sobressalente | Motivo (eixo da ficha) |
|---|---|
| Anticipate | F1: olha 3, pega 1 — **troca 1 por 1**. É o mesmo efeito que o Satoru já dá de graça a cada ninjutsu. Não conta para a meta. |
| Experimental Augury | F1: igual ao Anticipate + proliferate sem contador no deck para multiplicar (o Kaito Shizuki é o único alvo, e isso é marginal). Troca 1 por 1. A Fase 2 já a tinha encaminhado; o veredito é **dispensa**. |
| Quiet Contemplation | F1: não compra carta; tapa criatura quando se conjura **não-criatura**. O deck é de criaturas e ninjutsu (que é ativação, não conjuração). 0 pontos. |
| Way of the Necromancer | F1: Jace com 2 de lealdade só faz −1 surveil; o −3 compra exige criaturas **morrendo**, e o deck não sacrifica. F6: 2 manas por 2 surveil. Menos de 2 pontos. |
| Mindseeker Oculus | F2: 2/1 **sem evasão** — não conecta, logo não ativa ninjutsu nem é reciclado por ele. F6: 3 manas por uma compra adiada (Jace −3 no turno seguinte). 1 ponto (corpo). |
| Reverse Engineer | F6/F7: feitiço de 5 manas no turno em que a mana vai para o ninjutsu. Improvise precisa de artefatos, e o pool tem ~5. Perde para Thirst (instant, 3) e para Sphinx's Approach. Última reserva. |
| Culling Dais | F1/F7: o deck não tem motor de sacrifício. Sacrificar ativador para carregar contador é pior do que deixá-lo atacar e ser reciclado. |
| Oracle's Vault | F6/F7: `{2}`+tap por impulse de 1 carta, **disputando a mana do ninjutsu** todo turno; 3 turnos até o modo grátis. |
| Bargaining Table | F6: 5 manas para entrar + X por carta. Disputa mana com o ninjutsu. |
| Seer's Lantern | F1: rocha com scry 1. Seleção, não carta. Rampa ruim (3 manas por 1). Eventual papel é da Fase 4. |
| Elrond, Moon-Reader | Já dispensado pela Fase 2 (ninjutsu é ativado de card na mão, não de criatura — regra 109.2). Mantido. |

---

## 3. Radar do EDHREC

A seção `Radar do EDHREC` **não existe** no `02-theme.md` desta rodada, então fiz a chamada uma vez
(`get_edhrec_recommendations("Satoru Umezawa", limit 20)`). Inclusão veio quebrada (`0 decks`); só a nota de
sinergia foi lida. Recorte das seções da minha fase, classificado pela oracle:

| Carta (seção) | Synergy | O que é pela oracle | Destino |
|---|---|---|---|
| Kaito Shizuki (Planeswalkers) | 0,42 | +1 compra (sem descarte se atacou), −2 ficha Ninja inbloqueável | **candidata** |
| Bident of Thassa (Utility Artifacts) | 0,28 | motor por conexão | já no pool |
| Reconnaissance Mission (Enchantments) | 0,30 | motor por conexão | já no pool |
| Coastal Piracy (Enchantments) | 0,15 | 3ª cópia do motor | reserva |
| Rogue Class (Enchantments) | 0,12 | exila o topo de quem levou dano; nível 3 deixa jogar | reserva fraca (7 manas até virar carta) |
| Enter the Enigma (Sorceries) | 0,20 | inbloqueável + compra 1: cantrip | não conta (troca 1 por 1) |
| Sign in Blood / Night's Whisper / Read the Bones (Sorceries) | ≤ 0,04 | compra 2 genérica | dispensadas: sinergia genérica, feitiço (ver §4) |
| Mystic Remora / Phyrexian Arena (Enchantments) | 0,02 / −0,01 | motores genéricos | dispensados: 1 ponto só (compra), sem conversa com ninjutsu |
| Brainstorm / Ponder / Preordain | — | cantrips | não contam para a meta |
| Starwinder (Creatures, "and 30 more") | — | — | veio da **busca por termo**, não do EDHREC |

`Rhystic Study` aparece em Game Changers: fora (cota de bracket e preço de mercado; não é pedido do plano).

---

## 4. Candidatas recomendadas

| Carta | CMC | Tipo | Como gera vantagem | Sinergias (mín. 2) | Na coleção? | Preço (LigaMagic menor) |
|---|---|---|---|---|---|---|
| Thirst for Knowledge | 3 | Instant | Compra 3, descarta 2 (ou 1 artefato). **+1 a +2 líquido, em instant.** | (1) **instant**: sai no fim do turno do oponente e não toca na mana do ninjutsu; (2) o pool tem artefatos que viram descarte de 1 só (Ornithopter, Diversion Unit, Pilgrim's Eye, H.E.R.B.I.E., Meteor Golem, Noxious Gearhulk); (3) põe terreno excedente no cemitério e mantém bichões na mão, que com o Satoru são todos ninjutsu disponíveis | **sim** | caixa (R$ 0,05 em 2026-08-12) |
| H.E.R.B.I.E. Scout Unit | 4 | Artifact Creature — Robot Scout | ETB: compra 1 **e** põe terreno da mão no campo. Fica o corpo flier: +1 carta, +1 terreno. | (1) **ativador flier**; (2) o ninjutsu o devolve à mão e cada reconjuração repete compra + terreno (mesmo papel do Mulldrifter e do Pilgrim's Eye); (3) resistência 1 → **inbloqueável com Tetsuko**; (4) artefato para o descarte do Thirst; (5) por ninjutsu entra atacando e compra, se faltar bichão | **sim** | caixa (a cotar) |
| Kaito Shizuki | 3 | Legendary Planeswalker — Kaito | +1: compra; descarta **só se você não atacou** no turno. Como o deck ataca todo turno, é compra pura toda rodada. | (1) motor **repetível** que casa com o plano: ataque, ninjutsu, e o +1 no main 2 sai sem descarte; (2) −2: ficha **Ninja 1/1 inbloqueável** = ativador recorrente que também dispara **Ingenious Infiltrator** e recebe o anthem do Silver-Fur Master; (3) entra e **faz phasing out** no fim do turno, escapando da resposta dos oponentes; (4) UB, 3 manas | não | **a cotar** · origem: EDHREC (synergy 0,42) |
| Starwinder | 7 (warp 4) | Creature — Leviathan | "Criatura sua causa dano de combate a jogador → compre **essa quantidade**." Por ninjutsu entra **sem bloqueio**: 7 de dano = **7 cartas**, mais uma compra por ponto de dano de cada outro ativador que conectou. | (1) **alvo de ninjutsu** cujo valor é o gatilho de dano, que fica **garantido** na entrada (ruling da Fase 2); (2) soma com os motores do deck: Recon Mission, Bident e Ingenious também disparam no mesmo dano; (3) o 7/7 fica como bloqueador do Satoru; (4) **warp `{2}{U}{U}`**: sem Satoru, conjura-se por 4 no turno de ataque e volta do exílio depois | não | **a cotar** |

### Fichas F1–F7 das entradas

**Thirst for Knowledge** — F1: compra 3; descarta 2 salvo se descartar 1 artefato. F2: não é corpo. F3: instant (nada no pool
conta instant). F4: —. F5: —. F6: 3 manas, do T3 em diante, fora do próprio turno. F7: nenhum atrito de mana (é
jogada de fim de turno); risco de descartar bichão se a mão estiver só com criaturas, mitigado por escolher o artefato.

**H.E.R.B.I.E. Scout Unit** — F1: flying; ETB compra 1 e pode pôr terreno da mão virado. F2: 2/1 flier, conecta.
F3: artefato **e** criatura. F4: Tetsuko (resistência 1), anthem do Silver-Fur não (Robot Scout). F5: rampa (terreno
extra). F6: 4 manas — é o ponto fraco: no T4 compete com o primeiro ninjutsu. Melhor jogado no T4 **sem** ninjutsu
disponível, ou reconjurado com mana que sobrou. F7: disputa o turno 4 com Kaito, Dancing Shadow, Recon Mission e
Bident. **Caixa, custo zero** — por isso fica apesar do F6.

**Kaito Shizuki** — F1: phasing no fim do turno em que entra; +1 compra (descarta se não atacou); −2 ficha Ninja 1/1
inbloqueável; −7 emblema que busca criatura azul/preta para o campo a cada conexão (os bichões entram sem ser
conjurados → Satoru, the Infiltrator). F2: não é criatura; a ficha é. F3: planeswalker; a ficha é **Ninja**.
F4: —. F5: gera **ativador** inbloqueável (Tetsuko redundante). F6: T3, e no T4 já tem ativador ou carta extra.
F7: disputa o T3 com o próprio Satoru (ambos 3 manas); precisa de defesa, que os bloqueadores do Satoru (2/4)
e os bichões dão.

**Starwinder** — F1: gatilho de dano de combate de **qualquer** criatura sua → compra X = dano; warp `{2}{U}{U}`.
F2: 7/7 **sem evasão**: não se recicla sozinho pelo ninjutsu no turno seguinte (precisaria de Whispersilk Cloak ou de
Kaito, Dancing Shadow; a Tetsuko não serve, poder 7). F3: criatura. F4: —. F5: —. F6: entra por 4 (ninjutsu ou warp).
F7: estourar a mão — 7 cartas + impulse do Satoru passam de 7 no cleanup. **Nota à Fase 6: Reliquary Tower**
(não está na caixa) evita o descarte. Sem ruling oficial; o texto é direto.

### Nota sobre a Marang River Regent

Conto a Marang como **meia fonte**: a face instant é +1 líquido e reutilizável, mas quem joga a face criatura
(bounce duplo) abre mão da compra. Se a Fase 6 a mantiver como alvo de ninjutsu, a meta fecha em 13 **sem** ela
precisar comprar.

---

## 5. Reservas (entram se a Fase 6 cortar peça de draw do pool ou se a cotação estourar)

Em ordem de entrada. Todas com oracle puxado hoje.

| # | Carta | CMC | Tipo | Como gera vantagem | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| R1 | Grazilaxx, Illithid Scholar | 3 | Legendary Creature — Horror | 1 compra por combate em que alguma criatura sua conecta | (1) motor por conexão, como Mission/Bident, mas em corpo; (2) **ativador bloqueado volta à mão** em vez de morrer — salva o ativador e repete o ETB (Mulldrifter, H.E.R.B.I.E., Pilgrim's Eye) | não | a cotar |
| R2 | Chart a Course | 2 | Sorcery | Compra 2; descarta 1 **só se não atacou** | (1) jogada no main 2 depois do ataque: compra 2 pura por 2 manas; (2) preenche a faixa mv 1–2 da curva de draw, que está magra | não | a cotar |
| R3 | Sphinx's Approach | 3 | Instant | Compra 2 (+1 líquido) | (1) instant de fim de turno, preserva a mana do ninjutsu; (2) com o Satoru toda criatura comprada é um ninjutsu disponível. O modo tutor de Sphinx é irrelevante | **sim** | caixa (a cotar) |
| R4 | Sphinx of Enlightenment | 6 | Creature — Sphinx | ETB: você compra 3, um oponente compra 1 | (1) **alvo de ninjutsu** de ETB puro: +3 por 4 manas; (2) flier 5/5 — ataca no turno seguinte, passa, e é trocado de novo repetindo o ETB. Alternativa: Lord of Change (7, compra 3 sem presentear ninguém, ward 3), que a Fase 2 já listou na reserva | não | a cotar |
| R5 | Coastal Piracy | 4 | Enchantment | Motor por conexão (3ª cópia) | (1) redundância exata de Mission/Bident; (2) cada ativador **e** o bichão conectam | não | a cotar · origem: EDHREC (0,15) |
| R6 | Stern Lesson | 3 | Instant | Compra 2, descarta 1 + ficha Powerstone | (1) instant de fim de turno; (2) **a mana da Powerstone paga ninjutsu** (ninjutsu é habilidade ativada, não conjuração de feitiço não-artefato). **Não conta para a meta** (loot: troca neutra de cartas); é reserva de utilidade | **sim** | caixa (R$ 0,10 em 2026-09-18) |

**Olhadas e não propostas:** Jin-Gitaxias, Core Augur (alvo que compra 7 no fim do turno — sinergia enorme, mas
mítica de 10 CMC; **⚠ cotar primeiro** se o orquestrador quiser testar o teto) · Enduring Curiosity (motor com flash
e resiliente; **⚠ cotar primeiro**) · Impaler Shrike (flier 3/1 que se sacrifica por 3 cartas ao conectar; perde para
o Mulldrifter, que não se sacrifica) · Shadowmage Infiltrator (fear + compra, mas o ninjutsu tira-o do combate
antes do dano) · Phyrexian Gargantua (ETB compra 2 sem evasão; Mulldrifter faz o mesmo e voa) ·
Psychic Frog e Overlord of the Floodpits (mercado costuma pesar; sem cotação, não entram antes das acima) ·
Mask of Memory / Curious Obsession (equipamento e aura saem do ativador quando o ninjutsu o devolve — F7).

---

## 6. Lista final proposta (13 fontes inteiras + 1 meia)

| # | Fonte | Origem | Repetível? | Preço |
|---|---|---|---|---|
| 1 | Satoru Umezawa | comandante | sim | — |
| 2 | Satoru, the Infiltrator | pool temático | sim | R$ 2,63 |
| 3 | Moon-Circuit Hacker | pool temático | sim | R$ 3,95 |
| 4 | Ingenious Infiltrator | pool temático | sim | R$ 5,62 |
| 5 | Spectral Sailor | pool temático | sim | R$ 0,49 |
| 6 | Kaito, Dancing Shadow | pool temático | sim | R$ 1,85 |
| 7 | Reconnaissance Mission | pool temático (caixa) | sim | caixa |
| 8 | Bident of Thassa | pool temático (caixa) | sim | caixa |
| 9 | Mulldrifter | pool temático | reciclável | R$ 1,62 |
| 10 | Marang River Regent // Coil and Catch | pool temático | pontual (meia) | R$ 8,74 |
| 11 | **Thirst for Knowledge** | **Fase 3 — caixa** | pontual | caixa |
| 12 | **H.E.R.B.I.E. Scout Unit** | **Fase 3 — caixa** | reciclável | caixa |
| 13 | **Kaito Shizuki** | **Fase 3 — compra** | sim | a cotar |
| 14 | **Starwinder** | **Fase 3 — compra** | pico (reciclável com ajuda) | a cotar |

13 inteiras + a Marang como meia. **9 repetíveis**, o que é o ponto forte: a compra do deck vem quase toda de
conexão e ninjutsu, e o que eu acrescentei preenche o que faltava (instant fora do turno, pico, ativador que compra).

## Curva das fontes de draw

mv 1–2: 3 (Spectral Sailor, Satoru the Infiltrator, Moon-Circuit Hacker) ·
mv 3–4: 8 (Satoru Umezawa, Kaito Shizuki, Thirst, Ingenious, Kaito Dancing Shadow, Recon Mission, Bident,
H.E.R.B.I.E.) · mv 5+: 3 (Mulldrifter [evoke 3], Marang [Omen 4], Starwinder [ninjutsu/warp 4])

Na prática nada custa mais que 4 para entrar, mas a faixa 3–4 está **concentrada** e disputa o T3–T4 com o
primeiro ninjutsu. Se a Fase 6 precisar aliviar, **Chart a Course (R2)** é a troca de curva.

## Custo desta fase

- **Compras novas:** Kaito Shizuki (**a cotar**) + Starwinder (**a cotar**). Total cotado: **R$ 0,00 + 2 a cotar**.
- **Caixa:** Thirst for Knowledge e H.E.R.B.I.E. — custo zero (reservas R3 e R6 também da caixa).
- As 9 fontes do pool temático já estão no orçamento da Fase 2 (R$ 24,90 somando as cotadas; Recon e Bident na caixa).
- **Regra de envelope (~R$ 20–25):** se Kaito Shizuki + Starwinder passarem do envelope, sai primeiro a mais cara
  e entra **Grazilaxx (R1)**; se ainda estourar, **Sphinx's Approach (R3, caixa)** fecha a meta sem compra.

## Pendências

- **Cotar na LigaMagic:** Kaito Shizuki, Starwinder, H.E.R.B.I.E. Scout Unit (caixa, só para registro), e as reservas
  Grazilaxx, Chart a Course, Sphinx of Enlightenment, Coastal Piracy.
- **Fase 5:** Old Fat Spider Can't See Me (hexproof ao Satoru + 2 compras) — se entrar, conta +1 fonte aqui.
- **Fase 6:** Reliquary Tower (Starwinder + impulse do Satoru estouram a mão). Cortes de peças do pool que compram
  (Moon-Circuit, Ingenious, Mulldrifter, Marang) exigem reposição pelas reservas desta fase, em ordem.
- **Orquestrador:** o `Radar do EDHREC` não estava gravado no `02-theme.md`; o recorte da §3 foi feito aqui. As
  seções de `Mana Artifacts`, `Lands`/`Utility Lands` e `Game Changers` vieram no mesmo retorno, mas não foram
  transcritas — as Fases 4 e 6 vão precisar delas.
