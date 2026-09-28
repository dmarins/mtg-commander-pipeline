# Manabase e Cortes — Satoru Umezawa (v1 · build)

Fase 6 · `manabase-engineer` · 2026-09-27. Oracle puxado nesta sessão via `bin/mtgdb` (regra 6).
Preços: **LigaMagic (menor entre edições normais legais), via `mtgdb prices`, cotações de 2026-08-12 a
2026-09-27**. Carta sem cotação = **`a cotar`**, sem estimativa (regra 2). Não havia navegador nesta sessão.
EDHREC **não** foi chamado de novo (instrução do orquestrador): as seções `Lands`/`Utility Lands` não foram
transcritas pelas Fases 3–5; uso o que a Fase 4 registrou (Cabal Coffers, Urborg, Temple of the False God) mais os encaminhamentos da Fase 2 e do orquestrador.

Ponto de partida: **pool consolidado de 63 não-terrenos, R$ 173,86** (conferido: soma de `mtgdb prices`
das 63 cartas = R$ 173,86, todas cotadas).

---

## Cálculo

Duas leituras de CMC, porque o comandante muda o custo real dos bichões:

| Leitura | Como | CMC médio (62 não-terrenos, já sem o corte) | Fórmula | Resultado |
|---|---|---|---|---|
| **Impressa** | CMC do oracle | 3,27 | 31,42 + 3,13×3,27 − 0,28×11 | **38,6 → 39** |
| **Efetiva** | bichões de CMC ≥ 5 contam **4** (ninjutsu `{2}{U}{B}`, warp de Starwinder); Shriekmaw conta **2** (evoke `{1}{B}`) | 2,74 | 31,42 + 3,13×2,74 − 0,28×11 | **36,9 → 37** |

**Draws + ramps com mv ≤ 2 (N = 11):** draw — Spectral Sailor, Satoru, the Infiltrator e Moon-Circuit Hacker;
ramp — Sol Ring, Arcane Signet, Talisman of Dominance, Dimir Signet, Mind Stone, Silver-Fur Master, Hedron
Crawler e Hydraulic Helper.

**Recomendação: 37 terrenos** (1 abaixo da base de 38). O motivo do desvio:
- A leitura que manda é a **efetiva**. Com o Satoru em campo, 14 das 62 não-terrenos entram por 4, e a curva
  real termina em 4 (tabela da curva no fim). A fórmula efetiva dá 37 cravado.
- A leitura impressa (39) descreve o cenário **sem Satoru**. Nesse cenário o deck não quer mais terreno, quer
  **o Satoru de volta** (3 + 2 de taxa): o plano B conjurável por 4 já existe (Ravenous Chupacabra, Hostage
  Taker, Kaito, Dancing Shadow, Ingenious Infiltrator), e o resto espera.
- Além dos 11 ramps padrão (9 com mv ≤ 2) e do Peregrine Drake, o deck tem **3 peças de consistência de
  terreno** fora da conta de ramp: Pilgrim's Eye (básico na mão, repetível pelo ninjutsu), H.E.R.B.I.E. Scout
  Unit (terreno da mão para o campo + compra) e Evolving Wilds. Mão travada em 2 terrenos é menos provável do
  que a fórmula base supõe.
- O deck tem **sumidouros** para o excesso: segundo ninjutsu no turno (4), Spectral Sailor (`{3}{U}`), Rogue's
  Passage (`{4}`), Access Tunnel (`{3}`), Sinister Hideout (`{4}`), Recon Mission (cycling `{2}`), Mind Stone e
  Commander's Sphere (sacrifício compra). Se inundar, a mana vira carta ou evasão.

**Se o orquestrador preferir 38** ("na dúvida, 38"): a troca é **Hedron Crawler → Island** (ficha e custo
na seção de cortes, variante B). Não é minha recomendação, mas é a variante de menor atrito.

**Listas de referência (Archidekt, quantidades somadas, comandante confirmado no cabeçalho):**

| Lista | Bracket | Terrenos | Observação |
|---|---|---|---|
| #6488308 "Ninjitsu And Big Stuff" | 3 | **36** | mesmo plano (bichões); 14 Island + 7 Swamp |
| #3427781 "Ninjutsu" | 4 | **38** | 15 Island + 15 Swamp |
| #2380572 "Ninjas" | 3 | **32** | fetch + shocks + Otawara/Takenuma, muito ramp rápido; fora do nosso orçamento |
| #12430195 "30 dollar budget" | 3 | **40** | orçamento baixo, pouco ramp de mv 2 |

Mediana 37. As duas pontas se explicam pelo ramp: a de 32 compensa com terrenos que viram mágica (channel) e
fetches; a de 40 tem pouco ramp barato. Nosso deck fica no meio (11 ramps, 9 de mv ≤ 2) e bate com a mediana.
É contexto, não molde.

---

## Varredura da caixa (regra 7) — antes de qualquer busca

`mtgdb collection -list`: 182 sobressalentes, **7 terrenos**. Os 6 que o usuário separou para o Satoru + Room
of Refuge. Oracle de todos lido nesta sessão.

| Terreno da caixa | Veredito | Motivo (eixo da ficha) |
|---|---|---|
| Evolving Wilds | **entra** | F1 busca básico virado: busca **Island** (conta para Scourge of Fleets) ou Swamp conforme a mão. R$ 0,09 |
| Dimir Guildgate | **entra** | F1 dual U/B virado. Fixação sem custo real (caixa). Tapland 1/4 |
| Dismal Backwater | **entra** | F1 dual U/B virado + 1 de vida. Tapland 2/4 |
| Sinister Hideout | **entra** | F1 dual U/B virado + `{4},{T}`: surveil 1. É **sumidouro** de fim de jogo que arruma o topo antes do impulse do Satoru. Tapland 3/4 |
| Theorix Annex | **entra** | F1 dual U/B, **desvirado se você controla planeswalker**: o pool tem Kaito Shizuki e Kaito, Dancing Shadow. Na maioria das vezes entra virado. Tapland 4/4 |
| Hexhaven Dueling Arena | **dispensada** | F1: as duas ativações só tornam "prepared" criaturas **com prepare spells**, e o deck não tem nenhuma. Sobra um terreno que só gera `{C}`. O slot incolor rende mais com **Access Tunnel** (evasão para ativador) ou com uma **Island** (conta para Scourge e paga o `{U}`). O usuário separou a carta para este deck; a dispensa é por F1, não por preço |
| Room of Refuge | **dispensada** | F1: entra virado e gera uma cor só. O `{5}`, sacrifício, põe 2 contadores e **tira ativadores da faixa da Tetsuko** (poder/resistência ≤ 1). F7: seria a 5ª tapland. Uma dual da caixa já cobre a função com 2 cores |

Com isso a caixa entrega **5 terrenos** (4 duais + Wilds). Quatro entram virados, e é o teto que aceito:
mais que isso atrasa o Satoru no T3.

---

## Terrenos recomendados (37)

| Terreno | Produz | Entra virado? | Sinergia/Utilidade | Na coleção? | Preço |
|---|---|---|---|---|---|
| Command Tower | U ou B | não | dual perfeita | não | R$ 1,50 (09-18) |
| Sunken Hollow | U ou B | só com < 2 básicos (com 26 básicos, desvirada do T3 em diante) | **tipo Island Swamp**: conta para Scourge of Fleets | não | a cotar |
| Morphic Pool | U ou B | **não, em mesa com 2+ oponentes** (sempre, em pod de 4) | dual desvirada. ⚠ Se o torneio tiver mesa 1×1, entra virada | não | a cotar |
| Drowned Catacomb | U ou B | só sem Island/Swamp (quase nunca) | dual desvirada | não | a cotar |
| Dimir Guildgate | U ou B | sim | fixação | **sim** | a cotar (caixa) |
| Dismal Backwater | U ou B | sim | fixação + 1 de vida | **sim** | a cotar (caixa) |
| Sinister Hideout | U ou B | sim | `{4}`: surveil 1, sumidouro que arruma o topo antes do impulse | **sim** | a cotar (caixa) |
| Theorix Annex | U ou B | sim, salvo com Kaito em campo | fixação; desvira com os 2 Kaitos | **sim** | a cotar (caixa) |
| Evolving Wilds | (básico) | busca virado | busca Island para a Scourge | **sim** | R$ 0,09 (09-26) |
| Access Tunnel | C | não | `{3}`: criatura de poder ≤ 3 **inbloqueável**. Todo ativador do deck cabe, e Hostage Taker (2/3) também, para reciclar o ETB | não | a cotar |
| Rogue's Passage | C | não | `{4}`: **qualquer** criatura inbloqueável. Serve para o bichão de chão (Chupacabra, Grave Titan, Starwinder, Meteor Golem) conectar de novo e reciclar o ETB ou o gatilho de dano | não | a cotar |

**Básicos: 16 Island, 10 Swamp** — a cotar (Island e Swamp não têm cotação no `prices.tsv`).

**Proporção de pips** (62 não-terrenos): **U 48 / B 28** → 63% / 37%. Aplicado aos 26 básicos: 16,4 / 9,6 →
**16 / 10**. O ninjutsu `{2}{U}{B}` e o Satoru `{1}{U}{B}` pedem uma de cada, o que segura o B em 10 e não em 9.

**Fontes resultantes:** U = 16 Islands + 8 duais + Wilds = **25** · B = 10 Swamps + 8 duais + Wilds = **19** ·
incolores = 2 (Access Tunnel, Rogue's Passage). As rochas coloridas (Arcane Signet, Talisman, Dimir Signet,
Dimir Keyrune, Commander's Sphere) somam mais 5 fontes das duas cores.
**Islands para Scourge of Fleets:** 16 básicas + Sunken Hollow = **17 cartas com tipo Island** (Evolving Wilds
e Pilgrim's Eye buscam mais).
**Taplands:** 4 (as duais da caixa), + Sunken Hollow nos dois primeiros turnos.

### Dispensados, com motivo escrito

| Terreno | Origem | Motivo |
|---|---|---|
| Reliquary Tower (R$ 9,49) | Fase 3 / orquestrador | F1 "sem limite de mão", mas só incolor. O estouro de mão acontece no turno do Starwinder (e no do impulse + draw), e descartar até 7 com mão cheia de bichões ainda deixa munição. F7: 3º terreno incolor num deck de 2 cores com Scourge contando Islands. Custa **~metade do envelope** de terrenos por um efeito situacional. Se o orquestrador quiser, a função cabe numa rocha (Thought Vessel, R$ 9,95, reserva da Fase 4), não num slot de terreno |
| Cabal Coffers | Fase 4 (radar) / orquestrador | F1: não gera mana sozinha, só com `{2}` de entrada e conta Swamps. Com 10 Swamps + Sunken Hollow, o saldo real é +3 a +5 **no fim de jogo** e zero no início. F7: o teto de uso de mana do deck é o número de atacantes que conectam (cada ninjutsu pede um), então mana em massa sem escoadouro é perdida. E tira um slot de Island da Scourge |
| Urborg, Tomb of Yawgmoth | Fase 4 (radar) / orquestrador | F1: todo terreno vira Swamp, **inclusive os dos oponentes**. Aqui a única função é fixar B, e ela já está coberta por 19 fontes. Só faz sentido com Coffers (dispensada). **Reserva**: entra no lugar de 1 Swamp se a cotação for ≤ R$ 2 (é ela mesma um Swamp, então não perde fonte de B) |
| Temple of the False God | Fase 4 (radar) | F6: não gera mana até o 5º terreno. É exatamente o turno em que o deck já precisa de 4 coloridos para o ninjutsu |
| "Tomb of Yawgmoth" | Fase 4 (radar) | não é outra carta: `mtgdb` resolve o nome para **Urborg, Tomb of Yawgmoth** (linha acima). A Fase 4 listou a mesma carta duas vezes |
| Creeping Tar Pit | Archidekt #6488308 | F1: ativador **em slot de terreno**, 3/2 inbloqueável que sobrevive a wipe de feitiço. F6/F7: `{1}{U}{B}` + ninjutsu 4 = 7 manas, e o próprio terreno não gera mana no turno em que ataca. Entra virado. **Reserva 1** se a cotação for baixa, no lugar do Dimir Guildgate (mesma tapland, com função a mais) |
| Exotic Orchard (R$ 1,00) | radar/Archidekt | depende dos terrenos dos oponentes. **Reserva** se Morphic Pool ou Drowned Catacomb vierem caras |
| Watery Grave, Polluted Delta, Otawara, Takenuma | Archidekt #2380572 | as listas de bracket 3–4 usam, mas o histórico de mercado é de carta cara, e 1 carta dessas consome o envelope inteiro de terrenos. Não precisam de cotação: Command Tower, Morphic Pool e Drowned Catacomb já cobrem a função "dual desvirada" |

### Recuo por cotação (se uma dual comprada passar do envelope)

| Se estourar | Recua para | O que se perde |
|---|---|---|
| Drowned Catacomb | Shipwreck Marsh ou Choked Estuary (a cotar) · Exotic Orchard (R$ 1,00) · Island | nada relevante (as duas desviram quase sempre); com Island, 1 fonte de B |
| Sunken Hollow | Undercity Sewers ou Fetid Pools (tipo Island Swamp, **virados**) · Island | o desvirado. Com Island, mantém a contagem da Scourge |
| Morphic Pool | Shipwreck Marsh · Exotic Orchard | nada relevante |
| Access Tunnel ou Rogue's Passage | Island | a evasão em slot de terreno (a Whispersilk Cloak e a Tetsuko seguem no deck) |

---

## Plano de cortes (deck em 63 + 37 = 100 cartas → 99)

Excedente: **1 carta**. Orçamento: pool R$ 173,86 + terrenos. Com o corte abaixo, os não-terrenos ficam em
R$ 156,77 e sobram **R$ 43,23** para os 37 terrenos. Cada corte seguiu o protocolo §2 e a simetria §3 do
checklist. Nenhuma das cartas envolvidas consta de `decisions.md` (só o comandante está lá).

| Corte proposto | CMC | Preço | Motivo |
|---|---|---|---|
| **Agent of Treachery** | 7 | R$ 17,09 | peça mais cara do pool; o roubo e a remoção que ela faz já estão cobertos por Hostage Taker (R$ 0,50), Dragonlord Silumgar (R$ 2,89) e Meteor Golem (R$ 0,04). Ficha abaixo |

### Ficha e protocolo — Agent of Treachery (ponto (c) do orquestrador)

- **F1 Texto:** (1) ETB: ganha controle de **qualquer** permanente, sem prazo. (2) No seu end step, com 3+
  permanentes que não são seus, compra 3.
- **F2 Corpo:** 2/3 sem evasão. Não se recicla sozinha pelo ninjutsu no turno seguinte. Com **Access Tunnel**
  (poder 2 ≤ 3), Rogue's Passage ou Whispersilk Cloak, sim.
- **F3 Tipo:** Human **Rogue**. Dispara o Treasure do Prosperous Thief, recebe o anthem do Silver-Fur Master.
- **F4 Receptor:** anthem do Silver-Fur. Alvo de cópia da Thousand-Faced Shadow (dois roubos).
- **F5 Facilitador:** nenhum.
- **F6 Curva:** entra por 4 (ninjutsu). Conjurada, 7.
- **F7 Atrito:** nenhum interno.

| Função | Quem cobre depois do corte |
|---|---|
| roubo de criatura | Hostage Taker (exila criatura **ou artefato**, e você conjura) · Dragonlord Silumgar (criatura **ou PW**, enquanto ficar) |
| remoção de qualquer permanente não-terreno | Meteor Golem (destrói) · Dream Eater e Marang River Regent (bounce) · Withering Torment (encantamento) |
| alvo de ninjutsu com ETB forte | 14 outros alvos no pool (Chupacabra, Hostage Taker, Noxious Gearhulk, Grave Titan, Dream Eater, Silumgar, Gyruda, Marang, Rune-Scarred Demon, Sepulchral Primordial, Meteor Golem, Scourge of Fleets, Archfiend of Depravity, Starwinder, Peregrine Drake) |
| Rogue para o Prosperous Thief e o Silver-Fur | Tetsuko Umezawa, Invisible Stalker, Slither Blade, Satoru, the Infiltrator (e o Changeling Outcast, que tem todos os tipos). Todos **conectam mais** que a Agent, que não tem evasão |
| compra 3 no end step | **descoberta.** Custo aceito: a Fase 3 já não a contava na meta ("condicional demais"). A meta de draw segue em 13 inteiras + meia (Marang) |
| roubo **permanente** de encantamento, terreno ou rocha | **descoberta.** Custo aceito: é a única função exclusiva dela. Contra encantamento, a resposta passa a ser destruir (Meteor Golem, Withering Torment) ou devolver (Into the Roil, Dream Eater, Marang), não roubar |

**Simetria (§3):** a Agent foi avaliada nas mesmas condições das que ficam. Com o Satoru em campo, todas entram
por 4. Agent e Hostage Taker têm o mesmo corpo 2/3 sem evasão, e as duas se reciclam pelas mesmas peças
(Access Tunnel, Rogue's Passage, Whispersilk). A cópia da Thousand-Faced Shadow vale para as duas. Das três,
só o Silumgar tem o atrito de devolver o roubo ao ser reciclado, e ele compensa com o voo, que a Agent não tem.
**Na mesma condição**, a Agent faz o que a Hostage Taker faz, **mais** o roubo de não-criatura e a compra
condicional, por **R$ 16,59 a mais**. É o preço que não se paga no teto de R$ 200.

**Veredito (c):** sai. O corte é limpo, com as duas funções descobertas declaradas acima. A decisão fica com o
orquestrador, porque a carta é tema (alvo) e não terreno.

### Variante B — se o orquestrador fechar em 38 terrenos

| Corte | CMC | Preço | Entra |
|---|---|---|---|
| Hedron Crawler | 2 | R$ 0,10 (caixa) | Island (a cotar) |

**Ficha — Hedron Crawler:** F1 `{T}`: `{C}`. F2 0/1, tapável para mana, morre em wipe. F3 **artefato** (descarte
de 1 só no Thirst for Knowledge, custo do Disruption Protocol). F4 **Tetsuko** (resistência ≤ 1 → ativador
inbloqueável). F5 mana incolor. F6 T2 → Satoru + 1 no T3. F7 tapar para mana e atacar se excluem.

| Função | Quem cobre |
|---|---|
| mana no T2 → Satoru no T3 | a Island faz o mesmo no T2 **e** não morre em wipe, nem sofre enjoo |
| ramp (meta 10–11) | cai para **10**: ainda dentro da meta |
| ativador com Tetsuko | Ornithopter, Faerie Seer, Slither Blade, Siren Stormtamer, Thousand-Faced Shadow, Diversion Unit, Pilgrim's Eye e H.E.R.B.I.E. (todos ≤ 1 em poder ou resistência) |
| artefato para Thirst/Disruption Protocol | Ornithopter, Diversion Unit, Pilgrim's Eye, H.E.R.B.I.E., as 7 rochas |

Custo declarado: 1 ramp a menos (sai uma peça que acelera, entra uma que não acelera) e 1 ativador condicional
a menos. É a troca "land ↔ dork" dentro da própria lente de mana. Por isso é a variante de menor atrito, mas
**não** a recomendo: o 38º terreno só vale se o playtest mostrar mão travada.

---

## Pontos do orquestrador

**(a) Whelming Wave devolve o Satoru.** Oracle: *"Return all creatures to their owners' hands except for
Krakens, Leviathans, Octopuses, and Serpents."* O Satoru volta. Pela regra de comandante, o dono pode mandá-lo à
zona de comando no lugar da mão. **Mantê-lo na mão** é o certo: conjurado da mão ele custa `{1}{U}{B}` **sem
taxa**. Custo real do Wave:
- **Mana:** 4 do Wave + 3 do Satoru = **7 no mesmo turno** para não perder o comandante de vista. Com 6
  manas, o Satoru volta, mas os ativadores ficam na mão.
- **Tempo:** os ativadores voltam à mão também. Reconjurados, têm enjoo, e **o turno seguinte não tem ninjutsu**
  a menos que um atacante tenha ficado em campo. Ficam: **Starwinder (Leviathan)**, **Gyruda (Kraken)** e
  **Scourge of Fleets (Kraken)**, nenhum com evasão (Rogue's Passage resolve), e a **Dimir Keyrune**, que não
  é criatura no feitiço e vira 2/2 inbloqueável no turno seguinte. Access Tunnel e Rogue's Passage são o
  seguro disso.
- **Ganho:** fichas dos oponentes somem, comandantes voltam com taxa, e os bichões do deck vão para a mão como
  munição de ninjutsu (4 manas cada).
- **Leitura:** é botão de recuperação quando o deck está atrás, **não** jogada proativa. Custa ~1 turno de
  ninjutsu. Não proponho corte (é a única wipe conjurável sem Satoru, meta 2–4 = 3 com ela). O custo fica
  declarado para o `07-wincons.md` e para o goldfishing.

**(b) Kaito Shizuki (R$ 11,89) vs. reservas de draw mais baratas.** Oracle: +1 compra (só descarta se você
**não** atacou), −2 ficha Ninja 1/1 inbloqueável, phasing no turno em que entra.

| Candidata | Preço | Compra | Função que a Kaito tem e ela não |
|---|---|---|---|
| Kaito Shizuki | R$ 11,89 | +1 por turno, sem descarte num deck que ataca todo turno | — |
| Chart a Course | R$ 0,13 | 2 de uma vez (sorcery, sem descarte se atacou) | ativador recorrente, motor repetível, planeswalker (Theorix Annex) |
| Stern Lesson (caixa) | R$ 0,10 | loot + Powerstone (a Fase 3 não conta na meta) | idem |
| Sphinx's Approach (caixa) | a cotar | 2, instant | idem |
| Grazilaxx | R$ 7,24 | 1 por combate com conexão, e salva o ativador bloqueado | ativador (Grazilaxx é 3/2 sem evasão), planeswalker |

A Kaito tem uma função que nenhuma reserva cobre: **gera ativador inbloqueável sem gastar carta**, e a ficha é
Ninja (Ingenious Infiltrator, Silver-Fur, Prosperous Thief). É isso que liga o comandante depois de um wipe. Com
o corte da Agent, o orçamento comporta. **Veredito: fica.** Primeiro recuo, se a cotação dos terrenos apertar:
**Chart a Course** (R$ 0,13, economiza R$ 11,76). Custo do recuo: perde o ativador recorrente e o PW que desvira
a Theorix Annex.

**(c) Agent of Treachery:** sai. Ficha e protocolo acima.

### Cartas que o preço voltou a viabilizar

Nenhuma proposta nesta fase. O excedente é de 1 carta, e a folga de orçamento (abaixo) serve primeiro para
absorver os **35 terrenos `a cotar`**.

| Carta | Preço | Por que não entra agora |
|---|---|---|
| Reliquary Tower | R$ 9,49 | dispensada acima (slot de terreno incolor) |
| Ninja of the Deep Hours | R$ 7,50 | ninjutsu `{1}{U}` + compra. O Moon-Circuit Hacker já faz o papel por `{U}`, e o impulse do Satoru é 1×/turno. Primeira candidata se a Fase 7 pedir mais ninja nativo |
| Grazilaxx | R$ 7,24 | ver (b) |
| Fallen Shinobi | R$ 19,99 | alvo forte (dano garantido = 2 cartas grátis), mas pede slot e dinheiro. Com a folga, é a **primeira upgrade de tema** a considerar **depois** da cotação dos terrenos, no lugar de um alvo que a Fase 7 julgar mais fraco |
| Lord of the Void · Massacre Wurm | R$ 24,00 · R$ 24,99 | cada uma sozinha consome mais da metade da folga |

---

## Orçamento projetado (Commander 200)

| Bloco | Cartas | Cotado | A cotar |
|---|---|---|---|
| Não-terrenos (pool − Agent) | 62 | **R$ 156,77** | — |
| Terrenos cotados | Command Tower, Evolving Wilds | **R$ 1,59** | — |
| Terrenos a cotar | 16 Island, 10 Swamp, Sunken Hollow, Morphic Pool, Drowned Catacomb, Access Tunnel, Rogue's Passage, Dimir Guildgate, Dismal Backwater, Sinister Hideout, Theorix Annex | — | **35 cartas** |
| **Total** | 99 | **R$ 158,36 + 35 a cotar** | teto para os 35: **R$ 41,64** |

LigaMagic (menor), cotações de 2026-08-12 a 2026-09-27. Para ficar no envelope de terrenos (~R$ 20), os 35 a
cotar precisam somar **≤ R$ 18,41**. **Prioridade de cotação:** Drowned Catacomb, Sunken Hollow e Morphic Pool
(as três duais compradas), depois Rogue's Passage, Access Tunnel, Island e Swamp. Recuo de cada uma na tabela
de recuo. Cotações com 47 dias para reconferir antes do torneio: Bident of Thassa (R$ 5,80), Diversion Unit,
Old Fat Spider, Disruption Protocol, Thirst for Knowledge, Pilgrim's Eye, Hedron Crawler.

---

## Contagens depois do corte

| Categoria | Meta | Deck (62 não-terrenos + comandante) |
|---|---|---|
| Draw | 12–13 | **13 + meia** (lista da Fase 3; a Agent não contava) |
| Ramp padrão | 10–11 | **11** (10 na variante B) |
| Ramp explosivo | 2–3 | **1** (Peregrine Drake). ⚠ Abaixo da meta **desde a consolidação** (Grim Hireling R$ 86,53 saiu). Fora da minha lente; fica para o orquestrador |
| Interação | ~10 | **10 de mão** + 3 proteções (Siren Stormtamer, Diversion Unit, Old Fat Spider) + 8 ETBs de remoção |
| Wipes | 2–4 | **3** (Scourge of Fleets, Archfiend of Depravity, Whelming Wave) |
| **Ativadores** | — | **16 criaturas evasivas**: Ornithopter, Changeling Outcast, Faerie Seer, Spectral Sailor, Slither Blade, Siren Stormtamer, Thousand-Faced Shadow, Dimir Infiltrator, Invisible Stalker, Satoru, the Infiltrator, Tetsuko, Diversion Unit, Pilgrim's Eye, H.E.R.B.I.E., Mulldrifter, Shriekmaw. Mais **5 situacionais**: Peregrine Drake (voa no turno seguinte), Dimir Keyrune, a ficha do Kaito Shizuki, Hedron Crawler com Tetsuko, e qualquer criatura via Access Tunnel/Rogue's Passage |
| Ninjas nativos | — | 5 (Moon-Circuit Hacker, Silver-Fur Master, Ingenious Infiltrator, Thousand-Faced Shadow, Prosperous Thief) |
| **Alvos de ninjutsu** | — | **15** (a lista da ficha da Agent) + Mulldrifter e Shriekmaw como alvo e ativador ao mesmo tempo |

---

## Curva final projetada (62 não-terrenos, sem o comandante)

**Impressa:** 0–1: 8 · 2: 21 · 3: 9 · 4: 9 · 5: 4 · 6+: 11
**Efetiva (bichões por ninjutsu/warp a 4, Shriekmaw por evoke a 2):** 0–1: 8 · 2: 22 · 3: 9 · 4: 23 · 5: 0 · 6+: 0

A curva efetiva termina em 4, com 30 cartas até mv 2. É o que sustenta 37 terrenos em vez de 38–39.

---

## Pendências

1. **Cotar 35 terrenos** (lista no orçamento), começando pelas 3 duais compradas. Teto: R$ 41,64. Envelope:
   R$ 18,41.
2. **Formato da mesa do torneio:** Morphic Pool só desvira com 2+ oponentes. Se houver rodada 1×1, trocar por
   Shipwreck Marsh.
3. **Ramp explosivo em 1/2–3** desde a consolidação. Não é corte meu. Fica para o orquestrador decidir se aceita
   ou se busca um 2º dentro da folga.
4. **Variante 38 terrenos** (Hedron Crawler → Island) fica à disposição se o goldfishing (Fase 7) mostrar mão
   travada.
5. Radar do EDHREC: as seções `Lands`/`Utility Lands` nunca foram transcritas nesta rodada. Os terrenos
   utilitários vieram do recorte da Fase 4, do Archidekt e dos encaminhamentos.
