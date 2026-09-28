# Aceleração (Ramp) — Satoru Umezawa (v1 · build)

Fase 4 · `ramp-specialist` · 2026-09-27. Oracle, tags e rulings puxados nesta sessão via `bin/mtgdb`.
Preços: **LigaMagic (menor), via `mtgdb prices`, cotações de 2026-08-12 a 2026-09-27** (data carta a carta
na tabela). Carta sem cotação vai como **`a cotar`**, sem estimativa (regra 2).

Ramp padrão já no deck: **0/10–11** · Explosivo: **0/2–3**
(o `deck.md` só tem o comandante; no pool temático, só **Silver-Fur Master** funciona como ramp. **Pilgrim's Eye**
não conta, porque põe o terreno na mão e não no campo. **Palinchron** está fora do teto: R$ 104,99.)

Proposta desta fase: **11 padrão + 2 explosivos**, sendo **4 da caixa** (Sol Ring, Commander's Sphere,
Hedron Crawler e Hydraulic Helper), **1 do pool temático** (Silver-Fur Master) e **8 compras**.

---

## O que o ramp precisa comprar neste deck

O CMC impresso dos bichões não importa. O ramp compra **tempo de Satoru** e **mana no combate**:

| Marco | Sem ramp (só terrenos) | Com 1 rocha de mv2 no T2 |
|---|---|---|
| T3 | Satoru (3), nada sobra | Satoru + ativador de 1 (Faerie Seer, Slither Blade, Ornithopter) |
| T4 | 1 ninjutsu (4), zero de sobra | ninjutsu + 1: segura contramágica/proteção, ou ninjutsu nativo de `{U}` |
| T5+ | ninjutsu + 1 | **dois ninjutsus no mesmo combate** (4 + 4, ou 3 + 3 com Silver-Fur) |

Três fatos de regra que orientaram a escolha. Todos conferidos no oracle ou nos rulings desta sessão:

1. **Mana restrita "não pode ser gasta para conjurar mágica que não seja artefato" paga ninjutsu.** Ninjutsu
   é **habilidade ativada**, não mágica. O ruling do Powerstone (Stern Lesson, 2022-10-14) diz: *"You can use
   the {C}… on anything that isn't a nonartifact spell. This includes paying costs to activate abilities."*
   O Hydraulic Helper tem a mesma restrição e **não tem ruling próprio**, então vale a leitura por analogia
   com o texto idêntico.
2. **Ninjutsu pode ser ativado na etapa de dano de combate e no fim do combate** (ruling do Prosperous Thief
   e do Silver-Fur Master, 2022-02-18). Um tesouro gerado **por dano de combate** paga um ninjutsu
   **no mesmo combate**, depois do dano. Isso serve para bichão de **ETB** (Chupacabra, Agent, Gearhulk),
   não para bichão de dano de combate (Lord of the Void), que já chegaria tarde.
3. **Rochas valem mais que dorks aqui.** Dork não ataca e não é ativador. As duas exceções propostas custam
   R$ 0 (caixa) e têm uma segunda função escrita na ficha: bloqueador que protege o Satoru, e ativador com Tetsuko.

---

## Sobressalentes — varredura da caixa (regra 7)

`mtgdb collection -list`: **182 cartas**, e **76 dentro de UB/incolor**. Varridas antes de qualquer busca.

**Entram (4):** Sol Ring, Commander's Sphere, Hedron Crawler e Hydraulic Helper. As fichas estão nas tabelas abaixo.

**Dispensadas, com o motivo escrito:**

| Sobressalente | Encaminhada pela | Motivo da dispensa (eixo da ficha) |
|---|---|---|
| Manalith | Fase 2 | **Reserva nº 2** (custo zero). F6: mv3 disputa o T3 com o próprio Satoru, e as rochas de mv2 compradas chegam um turno antes. Entra se a cotação das mv2 estourar. |
| Dependable Quinjet | — | **Reserva** (custo zero). Igual à Manalith (mv3, qualquer cor). F7: o crew 4 **tapa** criaturas que deveriam atacar para ativar ninjutsu, e o Satoru (poder 2) não tripula sozinho. O corpo voador não se aproveita. |
| Prophetic Prism | Fase 2 | F1: filtro (`{1},{T}`: 1 mana de qualquer cor) dá **saldo zero de mana**. Não é ramp. Num deck de 2 cores o filtro quase não se usa. Sobra um cantrip de 2. |
| Mycosynth Wellspring | Fase 2 | F1: põe básico **na mão**, sem mana extra. É seguro de land drop, não aceleração. Se a Fase 6 quiser consistência, Pilgrim's Eye (já no pool) faz o mesmo e voa. |
| Spider-Bot | pedido do orquestrador | F1: põe o básico **no topo**, o que troca a próxima compra em vez de somar mana. F7: **atrito com o impulse do Satoru**, porque 1 das 3 cartas vistas no próximo ninjutsu já é um terreno conhecido e a seleção piora. F2: 2/1 com alcance, sem evasão, não ativa ninjutsu. |
| H.E.R.B.I.E. Scout Unit | pedido do orquestrador | **Não conta na meta de ramp**: mv4 (fora de mv ≤ 3) e só ramp com terreno na mão. Tem ficha boa para **outra fase**: flier 2/1 (ativador), ETB compra 1 + terreno da mão, e **reciclável**, porque o ninjutsu o devolve e a reconjuração repete compra + terreno. **Encaminhada à Fase 3 / orquestrador** como ativador com draw. Custo zero. |
| Pilgrim's Eye | já no pool temático | Fica no pool pelo tema. Não entra na contagem de ramp (terreno na mão, não no campo). Ajuda a garantir o 4º terreno do ninjutsu, e isso é consistência. |
| Stern Lesson | — | Draw 2 / descarta 1 + **Powerstone tapado**, cuja mana paga ninjutsu (ruling acima). É uma carta de **draw com ramp de brinde**. **Encaminhada à Fase 3.** Se entrar lá, soma +1 de ramp parcial sem custar slot desta fase. Custo zero. |
| Cargo Ship | Fase 2 | F1: a mana só paga mágica de artefato ou habilidade **de fonte artefato**. O ninjutsu de um bichão não-artefato não se qualifica. F7: crew tapa atacante. |
| Magnifying Glass, Seer's Lantern | — | F1/F6: rochas **incolores de mv3**, que não fixam U/B para o Satoru no T3. As habilidades extras (`{4}` investigar, `{2}` scry 1) disputam a mana do ninjutsu. Commander's Sphere cobre o slot de mv3 **com cor e com draw**. |
| Omni-Cheese Pizza | — | F1: ovo de mana. Custa 2 e devolve 1 ao sacrificar, então o saldo é negativo. É cantrip, não ramp. |

Da caixa, **Sphere of the Suns** e **Myr Convert** saíram desde a Fase 2 e não estão mais disponíveis.

---

## Radar do EDHREC (consultado nesta fase)

O `02-theme.md` desta rodada **não tem** a seção `Radar do EDHREC`, então fiz a chamada única
(`get_edhrec_recommendations`, limit 20). Leitura pela nota de **Synergy**, com a inclusão quebrada ignorada.
Orquestrador: se quiser o radar completo para as Fases 5–7, as seções `Instants`, `Sorceries`,
`Utility Artifacts`, `Lands` e `Utility Lands` vieram no mesmo retorno. Reproduzo aqui só o que é de ramp.

- `Mana Artifacts`: Sol Ring (0.04), Arcane Signet (0.04), Dimir Signet (0.04), Talisman of Dominance (0.00),
  Thought Vessel (0.14), Fellwar Stone (0.03), Mind Stone (−0.02), Commander's Sphere (0.02), **Dimir Keyrune (0.08)**,
  Sky/Charcoal Diamond, Decanter of Endless Water, Lotus Petal, Chrome Mox.
- `Creatures` (ramp): **Prosperous Thief (0.31)**, Ornithopter of Paradise (0.24).
- `Instants` (ramp): Dark Ritual (0.02).
- `Utility Artifacts` (ramp): Sword of the Animist (0.06), Wayfarer's Bauble (−0.02).
- `Utility Lands` (ramp): Cabal Coffers, Urborg, Temple of the False God. **Encaminhados à Fase 6.**

Toda carta daqui passou pelo mesmo filtro das outras (oracle, ficha, 2+ sinergias) e leva `origem: EDHREC`.

---

## Candidatas — ramp padrão (11)

Legenda: **F2** corpo · **F3** tipo · **F4** recebe · **F5** facilita · **F6** turno · **F7** atrito.

| # | Carta | CMC | Tipo | O que gera | Ficha F1–F7 resumida | Sinergias (mín. 2) | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|---|
| 1 | Sol Ring | 1 | Artifact | `{C}{C}` | F3 artefato · F6 **T1 → Satoru no T2** (2 terrenos + CC) · F7 nenhum | (1) Satoru **no T2**, ativador no T3 e ninjutsu no T3–4; (2) CC paga o genérico `{2}` de todo ninjutsu; (3) artefato para Disruption Protocol/Stoic Rebuttal (caixa, Fase 5). origem: EDHREC | **sim** | na caixa (R$ 7,95 em 2026-09-18) |
| 2 | Arcane Signet | 2 | Artifact | U ou B | F3 artefato · F6 T2 → Satoru + 1-drop no T3 · F7 nenhum | (1) **fixa as duas cores** de `{1}{U}{B}` e `{2}{U}{B}`; (2) T2 → T3 com 4 manas: Satoru + ativador no mesmo turno. origem: EDHREC | não | **R$ 4,00** (2026-09-18) |
| 3 | Talisman of Dominance | 2 | Artifact | `{C}` ou U/B (1 de dano) | F3 artefato · F6 T2 · F7 1 de vida por mana colorida, irrelevante | (1) fixação U/B sem entrar tapado; (2) a saída `{C}` paga o genérico do ninjutsu sem dano. origem: EDHREC | não | **a cotar** |
| 4 | Dimir Signet | 2 | Artifact | `{1}` → `{U}{B}` | F3 artefato · F6 T2 · F7 precisa de 1 mana de entrada (saldo +1) | (1) devolve **exatamente U+B**, as duas cores do ninjutsu e do Satoru; (2) com Sol Ring vira `{C}` → `{U}{B}`. origem: EDHREC | não | **a cotar** |
| 5 | Mind Stone | 2 | Artifact | `{C}` | F1 `{1},{T}`, sacrifica: compra 1 · F3 artefato · F6 T2 · F7 incolor | (1) T2 → T3 Satoru + 1; (2) **vira carta** no fim de jogo, quando a mana sobra e o deck precisa de bichão na mão. origem: EDHREC | não | **R$ 2,25** (2026-08-26, 32 dias: reconferir) |
| 6 | Dimir Keyrune | 3 | Artifact | U ou B | F1 `{U}{B}`: vira 2/2 Horror **inbloqueável** até o fim do turno · F2 **vira criatura** · F3 artefato · F7 usá-la como ativador faz o ninjutsu devolvê-la à mão (recast 3) | (1) rocha de cor; (2) **ativador de emergência que sobrevive a wipe de criaturas**: depois de um Damnation/Toxic Deluge, a Keyrune ainda liga o ninjutsu no turno seguinte; (3) inbloqueável garante o dano para o ninja que entra. origem: EDHREC | não | **a cotar** |
| 7 | Commander's Sphere | 3 | Artifact | U ou B | F1 sacrifica: compra 1 · F3 artefato · F6 T3 disputa com o Satoru, então vai no T4+ ou no T3 depois de Sol Ring · F7 mv3 | (1) fixação U/B; (2) vira carta no fim de jogo. **Custo zero** | **sim** | na caixa (R$ 0,50 em 2026-09-23) |
| 8 | Silver-Fur Master | 2 | Creature — Rat Ninja | **−`{1}` em todo ninjutsu** | F1 ninjutsu `{U}{B}`, redutor, anthem +1/+1 Ninja/Rogue · F2 2/2 · F3 Ninja · F5 **redutor de custo** · F6 T2 · F7 sem evasão (não é ativador); a redução não vale para o próprio ninjutsu dele (a estática só funciona no campo) | **Conta como ramp, e a justificativa é esta:** o ninjutsu do Satoru cai para `{1}{U}{B}` e o do Moon-Circuit Hacker para `{U}`. Com 2 ninjutsus por turno, rende **+2 de mana por turno**, o mesmo que um Sol Ring enquanto estiver em campo. (2) anthem nos Rogues/Ninjas ativadores; (3) dispara o impulse. **Já está no pool temático**: não é slot nem compra nova. origem: EDHREC alta sinergia (0.75) | não | R$ 0,60 (2026-09-27; já contado no pool) |
| 9 | Prosperous Thief | 3 | Creature — Human Ninja | **Tesouro** a cada combate em que Ninja/Rogue conecta | F1 ninjutsu `{1}{U}`; Ninja ou Rogue causa dano de combate a jogador → 1 Treasure · F2 3/2 (4/3 com Silver-Fur) · F3 **Ninja** · F4 Silver-Fur · F5 Treasure (mana de qualquer cor, **instantânea**, e artefato) · F6 **ninjutsu no T2** (ativador de 1 no T1) → tesouro → **Satoru + 1 no T3** · F7 só conta dano de Ninja/Rogue | (1) **ninja nativo barato**: ninjutsu de 2 dispara o impulse do Satoru; (2) ramp **recorrente** alimentado por Slither Blade, Invisible Stalker, Changeling Outcast, Tetsuko, os dois Satorus, Agent of Treachery e todos os ninjas; (3) o tesouro paga um ninjutsu **pós-dano** no mesmo combate (ruling); (4) é Ninja para Ingenious Infiltrator. origem: EDHREC (0.31) | não | **a cotar** |
| 10 | Hedron Crawler | 2 | Artifact Creature — Construct | `{C}` | F2 **0/1** · F3 artefato · F4 **Tetsuko** (resistência ≤ 1 → inbloqueável) · F6 T2 → **paga o `{1}` do Satoru no T3** · F7 enjoo de invocação; tapar para mana e atacar são excludentes; morre em wipe | (1) T2 → T3 com 4 manas (a C paga o genérico do Satoru); (2) **com Tetsuko em campo vira ativador inbloqueável** e volta à mão pelo ninjutsu por 2. **Custo zero** | **sim** | na caixa (R$ 0,10 em 2026-08-12) |
| 11 | Hydraulic Helper | 2 | Artifact Creature — Robot | `{U}` (só para artefato ou habilidade) | F1 Defender; `{T}`: U que não conjura mágica não-artefato · F2 **2/3 defensor** · F3 artefato · F6 T2, com mana útil **a partir do T4** · F7 **não** paga o Satoru nem instantâneas; paga ninjutsu, a ativação da Keyrune, o draw do Spectral Sailor e a conjuração de Ornithopter, Baleful Strix, Diversion Unit, Pilgrim's Eye, Meteor Golem e Noxious Gearhulk | (1) o `{U}` paga o **U do ninjutsu** (ruling do Powerstone, texto idêntico); (2) **bloqueador 2/3 que protege o Satoru** em casa. O ponto fraco do deck, segundo a Fase 2, é o Satoru morrer antes da ativação; (3) recast barato dos ativadores-artefato devolvidos pelo ninjutsu. **Custo zero** | **sim** | na caixa (sem cotação; custo 0) |

**Composição:** 7 rochas (5 de mv ≤ 2), 2 ninjas de ramp (redutor + tesouro) e 2 dorks da caixa com segunda
função. Das 11, **9 têm mv ≤ 2**. As 7 rochas que produzem U ou B fixam as cores; Sol Ring e Mind Stone são
incolores e pagam o genérico.

## Candidatas — ramp explosivo (2)

| # | Carta | CMC | Tipo | O que gera | Ficha F1–F7 resumida | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|---|
| E1 | Peregrine Drake | 5 | Creature — Drake | **desvira até 5 terrenos** | F1 flying; ETB desvira até 5 terrenos · F2 2/3 flier (ativador no turno seguinte; não serve para Tetsuko) · F6 via ninjutsu por 4 a partir do T4 · F7 desvira **só terrenos**, não rochas; o impulse do Satoru segue 1×/turno | **É o substituto do Palinchron que a Fase 2 achava não existir.** (1) Ninjutsu do Drake com 4 terrenos, que ele desvira; aí o **próprio Drake, atacante não bloqueado, é devolvido** por um 2º ninjutsu que traz o bichão. Resultado: **dois ninjutsus pelo preço de um, todo turno**, e o Drake volta à mão pronto para repetir; (2) flier: se ficar em campo, ataca e ativa o ninjutsu no turno seguinte; (3) o ETB não depende de "cast" | não | **a cotar** |
| E2 | Grim Hireling | 4 | Creature — Tiefling Rogue | **2 Treasures** por combate com dano a jogador | F1 qualquer criatura sua causa dano de combate a jogador → 2 Treasures; `{B}` + sacrificar X tesouros: −X/−X (feitiço) · F2 3/2 (4/3 com Silver-Fur), sem evasão, fica em casa · F3 **Rogue** · F5 tesouros · F6 conjurada no T4 **ou** ninjutsu (entra não bloqueada, conecta e já gera 2 no mesmo combate) · F7 remoção só em velocidade de feitiço | (1) **qualquer** ativador que conecta gera 2 tesouros, e o deck conecta todo turno por design; (2) os tesouros pagam **um 2º ninjutsu pós-dano** no mesmo combate (ruling); (3) Rogue: dispara o Prosperous Thief e recebe o Silver-Fur; (4) escoadouro de tesouro vira **remoção** (−X/−X) | não | **a cotar** |

2 explosivos fecham a meta (2–3). O 3º fica como reserva (Crypt Ghast, abaixo) porque o deck tem **teto de uso
de mana baixo**: cada ninjutsu exige um atacante não bloqueado, e mana de dobrador de Swamp sem escoadouro
é mana perdida. Os dois escolhidos geram mana **no combate**, que é onde o deck gasta.

---

## Reservas (em ordem de entrada)

| Reserva | CMC | Entra no lugar de / se… | Ficha e sinergias (2+) | Na coleção? | Preço |
|---|---|---|---|---|---|
| Everflowing Chalice | 0/2 | uma rocha a cotar estourar, ou a Fase 6 cortar um dork | F1 multikicker `{2}`, `{C}` por contador · F3 artefato · F6 T2 como rocha de 1; no fim de jogo, 4 manas → 2 contadores. (1) T2 → T3 Satoru + 1; (2) **escala**, o único ramp da lista que cresce com o jogo; (3) sobrevive a wipe de criaturas (ao contrário dos dorks) | não | **R$ 0,99** (2026-09-18) |
| Manalith | 3 | Talisman/Signet estourar | ver dispensa acima: fixa cor, mv3. **Custo zero** | **sim** | na caixa (sem cotação) |
| Ornithopter of Paradise | 2 | Hedron Crawler (upgrade pago) | F1 flying, `{T}`: qualquer cor · F2 0/2 · F4 Tetsuko. (1) **dork que é ativador voador**; (2) qualquer cor; (3) artefato de recast 2. F7: tapar e atacar são excludentes, então cumpre um papel por turno. origem: EDHREC (0.24), encaminhado pela Fase 2 | não | R$ 6,79 (2026-09-23) |
| Crypt Ghast | 4 | 3º explosivo, se a Fase 6 fechar com ≥ 12 fontes de tipo Swamp | F1 extort; tapar Swamp → +`{B}` · F2 2/2 · F6 ninjutsu por 4. (1) entra por 4 via ninjutsu, independente do CMC; (2) extort drena a cada mágica conjurada. F7: só Swamps (as duais com tipo Swamp contam), **sem escoadouro** de B em massa no pool. Alternativa de corpo maior: **Nirkana Revenant** (4/4, `{B}`: +1/+1 como escoadouro próprio) | não | **a cotar** |
| Wayfarer's Bauble | 1 | um slot de mv ≤ 2 a mais | F1 `{2}`, sacrifica: básico **no campo** tapado. (1) terreno a mais, que sobrevive a remoção de artefato; (2) mais um básico para Crypt Ghast/Cabal Coffers. F7 custa 3 no total, lento. origem: EDHREC | não | **a cotar** |
| Dependable Quinjet | 3 | Commander's Sphere | ver dispensa acima. **Custo zero** | **sim** | na caixa (sem cotação) |
| Thought Vessel / Decanter of Endless Water | 2 / 3 | Mind Stone | `{C}` (Vessel) ou qualquer cor (Decanter) + **sem limite de mão**. O impulse + draw do deck pode passar de 7 cartas, e bichão na mão é munição de ninjutsu. origem: EDHREC | não | **a cotar** |

**Descartadas por custo ou atrito** (com cotação ou ficha): Fellwar Stone (R$ 14,50; faz o que o Arcane Signet
faz por R$ 4,00) · Sword of the Animist (R$ 27,00; o equip `{2}` disputa a mana do ninjutsu, e o +1/+1 tira
ativadores da faixa do Tetsuko) · Springleaf Drum (F7: tapa a criatura que ataca ou o Satoru que bloqueia) ·
Gilded Lotus, Thran Dynamo, Hedron Archive (F6: rocha de 4–5 conjurada, enquanto o explosivo por ninjutsu custa 4
sempre e ainda é corpo) · Mana Vault, Grim Monolith, Chrome Mox (histórico de mercado alto, `a cotar`; não
precisam de cotação porque as mv2 cotadas já cobrem a função) · Charcoal/Sky Diamond (entram tapados, mono-cor,
piores que Talisman/Signet).

---

## Encaminhamentos a outras fases

- **Fase 2 / orquestrador (ativadores):** **Cloud of Faeries** (`{1}{U}`, flier 1/1, ETB desvira 2 terrenos,
  cycling `{2}`). Devolvida pelo ninjutsu, **a reconjuração sai de graça** (paga 2, desvira 2). É o ativador
  reciclável de custo líquido zero, e com Tetsuko fica inbloqueável. Não é ramp (saldo 0), mas é a melhor carta
  que achei para a curva do ninjutsu. `a cotar`.
- **Fase 3 (draw):** Stern Lesson (caixa; draw + Powerstone que paga ninjutsu) e H.E.R.B.I.E. Scout Unit
  (caixa; flier com ETB compra + terreno, reciclável).
- **Fase 5 (interação):** **Snap** (`{1}{U}`, bounce + desvira 2 terrenos, de graça) e Dark Ritual (mana de
  combate, instantânea). Snap pode ser a remoção temporária mais barata do deck. Ambas `a cotar`.
- **Fase 6 (terrenos):** Cabal Coffers, Urborg, Tomb of Yawgmoth e Temple of the False God vieram no radar.
  Se entrar Crypt Ghast, a contagem de fontes de tipo Swamp decide.
- **Fase 7 (wincons):** Peregrine Drake + Silver-Fur Master: cada ciclo Drake ↔ ativador fecha **positivo**
  (paga 3 e desvira 5). **Não é combo infinito**: exige um atacante não bloqueado por ativação e o Drake só
  volta ao campo por ninjutsu. Regra 11 ok.

---

## Custo desta fase

| Grupo | Cartas | Custo |
|---|---|---|
| Caixa | Sol Ring, Commander's Sphere, Hedron Crawler, Hydraulic Helper | R$ 0,00 |
| Pool temático (já contado lá) | Silver-Fur Master | (R$ 0,60) |
| Compras cotadas | Arcane Signet 4,00 · Mind Stone 2,25 | **R$ 6,25** |
| Compras **a cotar** | Talisman of Dominance, Dimir Signet, Dimir Keyrune, Prosperous Thief, Peregrine Drake, Grim Hireling | **a cotar** |

**Total cotado: R$ 6,25** — LigaMagic (menor), cotações de 2026-08-26 e 2026-09-18. Para caber no envelope de
R$ 20–25, as **6 cartas a cotar** precisam somar no máximo **~R$ 14–19**. Se estourarem, a ordem de troca é esta
(cada troca preserva a função):
1. Dimir Keyrune → Dependable Quinjet (caixa). Perde o ativador anti-wipe, mantém a rocha de cor.
2. Talisman of Dominance → Manalith (caixa). Perde 1 turno de curva.
3. Dimir Signet → Everflowing Chalice (R$ 0,99). Perde a fixação de cor.
4. Grim Hireling / Peregrine Drake: **sem substituto de mesma função.** Se os dois estourarem, o explosivo
   fica só com Crypt Ghast/Nirkana (a cotar), com o atrito escrito acima.

## Pendências

- **Cotar na LigaMagic (6 da proposta):** Talisman of Dominance, Dimir Signet, Dimir Keyrune, Prosperous Thief,
  Peregrine Drake e Grim Hireling. **Das reservas:** Crypt Ghast, Nirkana Revenant, Wayfarer's Bauble,
  Thought Vessel, Decanter of Endless Water e Cloud of Faeries (encaminhada).
- **Reconferir** Mind Stone (cotação de 32 dias).
- **Hydraulic Helper:** a leitura "a mana paga ninjutsu" vem do ruling do **Powerstone**, que tem texto de
  restrição idêntico. A carta não tem ruling próprio.
- O `Radar do EDHREC` desta rodada foi consultado **nesta fase**, porque não existia no `02-theme.md`. Se o
  orquestrador quiser, pode movê-lo para lá.

---

## Ajuste pós-consolidação (2026-09-27)

Retorno do orquestrador: **Grim Hireling caiu** (R$ 86,53, fora do teto). O explosivo ficou em **1/2–3**
(só Peregrine Drake). O rascunho tem 99 cartas e R$ 175,91, sem slot livre, então a resposta é **um swap**.
Cotações do orquestrador, pela régua nova da regra 2 (menor preço entre edições legais): Crypt Ghast R$ 34,00 e
Nirkana Revenant R$ 24,00 (as duas acima de ~R$ 10) · Cloud of Faeries 0,83 · Wayfarer's Bauble 1,00.

### Candidatas a Y (explosivo até ~R$ 10)

| Candidata | Explosivo? | Veredito |
|---|---|---|
| **Great Whale** | **sim** | **Proposta.** Mesmo padrão do Drake, com teto maior (desvira até 7). Ficha abaixo. **a cotar** |
| High Tide | sim, uma vez só | **Plano B** se o Whale passar de R$ 10. Instantânea `{U}`: cada Island tapada dá +`{U}` até o fim do turno, e o deck tem 17 fontes de tipo Island (16 Island + Sunken Hollow). Conjurada depois dos bloqueadores, com 4–5 Islands, rende +3/+4 no combate. F7: gera **só U** (o ninjutsu pede B, que vem das Swamps), é de uso único e fica morta no início do jogo. **a cotar** |
| Cloud of Faeries | **não** | Saldo de mana zero (paga 2, desvira 2). É ativador reciclável grátis, e não ramp explosivo. Não fecha a meta. R$ 0,83 |
| Dark Ritual | não | +2 uma vez. É aceleração pontual, não mana de meio ou fim de jogo. **a cotar** |
| Snap | não | Saldo zero (bounce grátis). É interação. **a cotar** |
| Treachery | sim | Roubo + desvira 5, de graça, mas precisa ser **conjurado** (5), e o histórico de mercado é alto. **a cotar**, não proposto |
| Magus of the Coffers | parcial | Entra tapado pelo ninjutsu e só gera no turno seguinte (`{2}`: B por Swamp; 11 fontes de tipo Swamp). Mais lento e pior que o Whale. **a cotar** |

### Swap proposto: sai **Hedron Crawler** → entra **Great Whale**

**Entra — Great Whale** (`{5}{U}{U}` · 5/5 · Creature — Whale) · **a cotar**

| Eixo | Ficha |
|---|---|
| F1 | "When this creature enters, untap up to seven lands." (ruling 2022-12-08: sem alvo; você escolhe os terrenos na resolução) |
| F2 | 5/5 **sem evasão**; bloqueia bem se ficar em campo |
| F3 | Criatura (o CMC 7 é irrelevante via ninjutsu) |
| F4 | Alvo de Thousand-Faced Shadow e de Kaito, Dancing Shadow |
| F5 | **Mana**: desvira até 7 terrenos no meio do combate |
| F6 | Ninjutsu por 4 a partir do T4–5. Com 5 terrenos: paga 4, desvira 4, e o **próprio Whale** (atacante não bloqueado) é devolvido por um 2º ninjutsu que traz o bichão. Com Silver-Fur, 3 + 3 e sobra mana |
| F7 | Só desvira **terrenos** (não rochas). Sem evasão, então se ficar em campo não reataca com garantia. O padrão é devolvê-lo à mão no mesmo combate. O impulse do Satoru segue 1×/turno. O ETB não depende de "cast" |

Sinergias: (1) **redundância do padrão Peregrine Drake** — "dois ninjutsus pelo preço de um", repetível todo
turno porque o Whale volta à mão; 2 cópias do efeito em 99 dobram a chance de tê-lo; (2) **teto maior**: se
os terrenos já foram gastos antes (interação, conjuração no 1º main), desvira até 7 e financia 2 ninjutsus extras
quando há atacantes para isso; (3) **Thousand-Faced Shadow** (na lista) copia o Whale atacante, e o token desvira
mais 7; (4) Kaito, Dancing Shadow devolve o Whale à mão depois do dano.

**Sai — Hedron Crawler** (`{2}` · 0/1 · Artifact Creature — Construct) · caixa (R$ 0,10 em 2026-09-27). Volta a
ser sobressalente.

| Função (F1–F6) | Quem cobre depois do corte |
|---|---|
| F1 ramp mv2 (`{C}`) | 10 ramps padrão continuam, **dentro da meta 10–11**: Sol Ring, Arcane Signet, Talisman, Dimir Signet, Mind Stone, Dimir Keyrune, Commander's Sphere, Hydraulic Helper, Silver-Fur Master e Prosperous Thief |
| F6 acelerar o Satoru para o T3 (a C paga o `{1}`) | Sol Ring, Arcane Signet, Talisman, Dimir Signet e Mind Stone (5 peças). **Custo aceito:** a densidade de aceleradores de Satoru no T3 cai de 6 para 5 em 99 |
| F2 corpo 0/1 (chump block que protege o Satoru) | Hydraulic Helper (2/3 defensor) e o próprio Satoru (2/4) |
| F4 ativador inbloqueável com Tetsuko | 12+ ativadores na lista (Ornithopter, Changeling Outcast, Faerie Seer, Spectral Sailor, Slither Blade, Dimir Infiltrator, Invisible Stalker, Diversion Unit, Pilgrim's Eye, Siren Stormtamer, Thousand-Faced Shadow, e a Dimir Keyrune animada) |
| F3 artefato (Disruption Protocol tapa artefato; Thirst for Knowledge descarta artefato) | 8 rochas e Keyrune, Ornithopter, Diversion Unit, Pilgrim's Eye, H.E.R.B.I.E., Whispersilk, Meteor Golem, Noxious Gearhulk e Bident |
| F5 | nenhuma |

**Simetria (checklist §3):** os dois lados foram avaliados com peças condicionais da lista. O Crawler levou crédito
pelo Tetsuko; o Whale, pelo Thousand-Faced Shadow e pelo Kaito. Sem essas peças, o Crawler é um dork incolor de mv2
e o Whale segue sendo o 2º ninjutsu grátis por turno. A vantagem do Whale se mantém.

**Por que o Crawler e não outra peça:** entre os 11 padrão, ele é o único **dork incolor sem segunda função
exclusiva** (as duas que tinha estão cobertas acima). Commander's Sphere fica porque o deck tem só 10 Swamps e
precisa de fixação de B. O Hydraulic Helper fica porque é o bloqueador 2/3. Talisman e Signets são a curva de mv2
com cor.

**Resultado:** padrão **10/10–11** · explosivo **2/2–3** (Peregrine Drake + Great Whale). Custo: R$ 175,91 +
Great Whale (**a cotar**). Cabe no teto de R$ 200 se o Whale sair por até ~R$ 24.
**Se o Whale passar de ~R$ 10:** entra **High Tide** (a cotar), com o atrito escrito acima. **Se os dois
passarem:** o gap de 1 explosivo fica como **custo aceito**. Nenhuma outra candidata barata rende mais que o
Hedron Crawler, e o swap não deve ser feito.
