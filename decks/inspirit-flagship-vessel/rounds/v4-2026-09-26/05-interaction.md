# Interação — troca pontual (`/swap-card`) · Inspirit, Flagship Vessel

> Modo `improve`, escopo de **uma** troca. Entrada fixada pelo orquestrador: **Requisition Raid**
> (sobressalente, `mtgdb collection`: tem 1). Este arquivo verifica a triagem, fecha a ficha da
> entrada e propõe **a saída** (duas candidatas ranqueadas). Sem busca de carta nova.
> Todos os oracles foram puxados em 2026-09-26 com `bin/mtgdb oracle` (regra 6).

Interação já no deck: **9/~10** · Wipes: **2/2–4** · Respostas a artefato/encantamento: **2** · Counterspells: **0**

Contagem conferida no oracle: Swords to Plowshares, Dispatch, Stone by Sunlight, Spring-Loaded
Sawblades, Warmaker Gunship, Dawnsire, Sunstar Dreadnought, Alibou, Ancient Witness (criatura/qualquer alvo) ·
Sunder the Gateway (só artefato/encantamento **nontoken** de oponente) · Perilous Snare (qualquer
permanente não-terreno de oponente — é a única flexível). Wipes: Chain Reaction, Organic Extinction.

---

## 1. Verificação da triagem do orquestrador

| Afirmação da triagem | Conferido no oracle | Veredito |
|---|---|---|
| Interação 9, art/enc só 2, 0 counters | lista acima | ✅ confere |
| Modo 3 casa com as 27 criaturas | 27 criaturas de verdade (19 criaturas-artefato + 8 não-artefato: Evangel, Seedshark, Malcator, Sai, Vraska, Jhoira, Padeem, Deepglow) + tokens (Thopters, Servos, Sculptures, Golems, Drones, Alien). O modo 3 dá contador a **toda criatura** do jogador-alvo, artefato ou não. | ✅ confere, **com limite**: não alcança o Inspirit abaixo de 8, spacecraft abaixo do limiar, veículo não tripulado, Incubator não transformado nem rocks. Tripular Vertibird/Recon Craft **antes** de lançar a Raid faz o veículo pegar o contador (fica nele depois) — nicho, custa os taps. |
| Modo 3 casa com o gatilho `1+` do comandante | O `1+` põe +1/+1 **ou** 2 charges em `other target artifact`. Não interage com a Raid: são duas fontes paralelas de contador. | ⚠️ **sinergia fraca/indireta** — não conto como ponto. |
| 5 fontes de proliferate | Metastatic Evangel, Surge Conductor, Kilo, Recon Craft Theta, Tezzeret's Gambit — todas proliferam; + Deepglow Skate **dobra** contadores no ETB. | ✅ confere: um contador em cada criatura vira N contadores ao longo do jogo. |
| Hangarback / Marketback Walker | contadores +1/+1 = Thopters / cartas quando morrem | ✅ confere; + **Crystalline Crawler** (contador = mana de qualquer cor), que a triagem não citou |
| Dispara Vraska, Seedshark, Whirlwind | as três dizem "whenever you cast a noncreature spell" | ✅ Vraska (Sculpture) e Whirlwind (1 carta) plenos. ⚠️ **Seedshark: Incubate 1**, não mais — o MV da Raid é sempre 1, qualquer que seja o spree (ruling 2024-04-12; Seedshark usa o MV da mágica). |
| — (não citado) | Saheeli, Sublime Artificer também dispara (Servo 1/1) | ✅ ponto a mais |
| — (não citado) | **Ordem de resolução**: os gatilhos de lançamento (Sculpture da Vraska, Servo da Saheeli) resolvem **antes** da Raid → os tokens criados por ela mesma **recebem** o contador do modo 3. O Incubator do Seedshark não (não é criatura). | ✅ ponto a mais |

---

## 2. Candidatas — remoção / proteção / counters

| Carta | CMC | Tipo | Subcategoria | O que responde | Sinergias | Na coleção? |
|---|---|---|---|---|---|---|
| Requisition Raid | 1 (paga-se `{1}{W}` a `{3}{W}`) | Sorcery | remoção de artefato **e** encantamento (dois alvos numa mágica) + modo de contadores em massa | artefato (inclusive **token** e de qualquer jogador — Sunder só pega nontoken), encantamento; nada instantâneo, nada que seja indestrutível | (1) modo 3 × 5 proliferates + Deepglow; (2) Hangarback/Marketback/Crawler convertem o contador em Thopter/carta/mana; (3) +1 de poder em cada corpo = +1 charge por criatura tapada em Station; (4) mágica não-criatura de MV 1: Vraska, Saheeli, Whirlwind (+ Seedshark, só Incubate 1); (5) os tokens da Vraska/Saheeli nascem a tempo de pegar o contador | **sim (1)** — entrada fixada |

### Ficha F1–F7 — Requisition Raid (entra)

| # | Eixo | Leitura |
|---|---|---|
| F1 | Texto | Spree: `+{1}` destrói artefato alvo; `+{1}` destrói encantamento alvo; `+{1}` um +1/+1 em cada criatura do jogador-alvo. Modos resolvem na ordem escrita. MV = 1 sempre. |
| F2 | Corpo | Nenhum. |
| F3 | Tipo | Feitiço → conta como mágica não-criatura (Vraska, Saheeli, Whirlwind, Seedshark). **Não é artefato**, **não é histórica** (Jhoira não compra). |
| F4 | Receptor | Nada. |
| F5 | Facilitador | Modo 3: +1/+1 permanente no time (independe de Chief/Master vivos, ao contrário dos anthems); alimenta Hangarback/Marketback/Crawler; +1 charge por corpo tapado em Station. |
| F6 | Curva | Castável no turno 2 com um modo; no fim de jogo, `{3}{W}` para os três. |
| F7 | Atrito | Velocidade de feitiço (não responde a nada na pilha nem no turno do oponente). Destrói, não exila — artefato indestrutível sobrevive. Pede só `{W}` (6 Plains + fontes W) — tira pressão de U, a cor mais pedida. |

---

## 3. Candidatas — board wipes

Nenhuma — escopo desta troca é uma entrada fixada. Wipes seguem em **2** (no piso).

---

## 4. Saída proposta (ranqueada)

Critério de busca da saída: o deck está **acima da meta de draw (14 contra 12–13)** e abaixo em
interação (9 contra ~10). A saída natural é uma **fonte de draw de mágica não-criatura e não-artefato**:
sai sem mexer em artefatos (42), corpos (27), mágicas não-criatura (Vraska/Saheeli/Whirlwind
continuam com o mesmo número de gatilhos) nem na base de mana.

Descartadas antes da ficha completa, por exercer função que a Raid não cobre:
Stern Lesson (instantâneo + Powerstone = ramp e artefato), Tezzeret's Gambit (1 dos 5 proliferates),
Thought Monitor / Ethersworn Sphinx / Midnight Clock / Uthros (artefatos → contagem de 42 e/ou corpo),
qualquer peça de interação/proteção (anularia o ganho), terrenos (36, já abaixo de 38), Vraska e o
comandante (protegidos).

### Candidata 1 (recomendada) — **Reverse Engineer**

`{3}{U}{U}` · Sorcery · CMC 5 · Improvise · Draw three cards. · R$ 0,20 (LigaMagic menor, 2026-09-18)

| # | Eixo | Leitura |
|---|---|---|
| F1 | Texto | Improvise (tapar artefatos paga genérico); compra 3. |
| F2 | Corpo | Nenhum. |
| F3 | Tipo | Feitiço, não-criatura, não-artefato, não-histórico. |
| F4 | Receptor | Nada. |
| F5 | Facilitador | Nada para outras cartas. Improvise transforma artefatos ociosos (inclusive tokens com enjoo de invocação, Golem Foundry, Cloud Key, spacecraft abaixo do limiar) em mana; **tapar artefatos antes do combate aumenta o X da Alibou** — relevante com a Vraska em campo, porque atacantes com vigilância não contam para o X. |
| F6 | Curva | Nominal 5; na prática `{U}{U}` + 3 artefatos tapados a partir do meio de jogo. |
| F7 | Atrito | Os artefatos tapados por improvise **deixam de estacionar/tripular** naquele turno (disputa de tap com Station/crew). Pede `{U}{U}`. |

Gatilhos que ela dispara hoje: Vraska (Sculpture), Saheeli (Servo), Whirlwind (+1 carta) e
**Seedshark com Incubate 5**.

**Protocolo de corte (seção 2)**

`Sai: Reverse Engineer — funções: [compra 3, gatilho de mágica não-criatura, Incubate 5 via Seedshark, improvise → X da Alibou]`

| Função | Quem cobre depois |
|---|---|
| Compra 3 (burst) | **Voyage Home** (compra 3 por affinity — sai por `{W}{U}` com 5+ artefatos, **sem tapar** artefatos, portanto sem o atrito de Station que o improvise tem), Tezzeret's Gambit, Stern Lesson, Thought Monitor, Ethersworn Sphinx (cascade), Midnight Clock; motores: Whirlwind, Jhoira, Sai, Uthros, Thopter Spy Network, Padeem, Marketback. Draw **14 → 13**, dentro da meta 12–13. |
| Gatilho de mágica não-criatura (Vraska, Saheeli, Whirlwind) | **Requisition Raid** — troca feitiço por feitiço; a contagem de mágicas não-criatura não muda. Com a Raid, a Sculpture e o Servo ainda ganham um +1/+1 cada. |
| Incubate 5 via Seedshark | **Descoberta em parte.** A Raid dá Incubate 1. Continuam com MV alto: Voyage Home (Incubate 7), Organic Extinction (10), Tezzeret's Gambit (4), Whirlwind (4). **Custo aceito**: perde-se um Incubator 5/5 (Station de 5, 7 com os dois anthems) nas vezes em que Seedshark e Reverse Engineer se encontram. |
| Improvise → X da Alibou antes do combate | **Descoberta em parte.** Rocks tapados para pagar a Raid também ficam virados e contam; o que se perde é tapar artefatos que **não** produzem mana. Organic Extinction e Kappa Cannoneer continuam com improvise. **Custo aceito**: pequeno — é uma fonte ocasional de X, não um plano. |

Corte **não é 100% limpo**: duas funções ficam parcialmente descobertas, ambas de nicho e declaradas acima.

**Simetria de critério (seção 3)**

- **Vraska em campo**: os dois lados disparam a Sculpture; a da Raid ainda ganha o contador do modo 3. Mesma condição para os dois.
- **Seedshark em campo**: aplicada aos dois lados — Reverse Engineer gera Incubate 5, a Raid Incubate 1. Registrado como perda, não omitido.
- **Anthems (Chief + Master)**: nenhum dos dois lados tem corpo, então não entram em nenhum dos lados. O modo 3 da Raid soma **por cima** dos anthems e fica mesmo se eles morrerem.
- **Contagem de artefatos**: nenhum dos dois é artefato → **42 permanece**; Master of Etherium, metalcraft do Dispatch, affinity (Thought Monitor, Sphinx, Voyage Home), improvise, Malcator e Chief of the Foundry não perdem nada.
- **Proteção do comandante**: nenhum dos dois protege; Padeem, Loran's Escape, Blacksmith's Skill, Invisible Force Field e Restoration Magic continuam.
- **Alibou com Vraska**: aplicada aos dois lados — o improvise do Reverse Engineer era uma fonte de artefatos virados; os rocks usados para pagar a Raid também são.

**Histórico (regra 5)**: entrou na **v1 pelo pipeline** (fase draw, no slot da Ethersworn Sphinx:
"mv9 morto na mão → draw 3 barato via improvise") e seguiu na v2. **Não está** entre as 16 entradas
do usuário no diff pós-v3. O motivo original da entrada (substituir a Sphinx) não vale mais: o
próprio usuário trouxe a Sphinx de volta e as duas convivem. Cortá-la não desfaz escolha declarada
dele.

### Candidata 2 (reserva, **pergunte ao usuário**) — **Voyage Home**

`{5}{W}{U}` · Sorcery · CMC 7 · Affinity for artifacts · You draw three cards and gain 3 life. · R$ 0,20 (LigaMagic menor, 2026-09-18)

| # | Eixo | Leitura |
|---|---|---|
| F1 | Texto | Affinity for artifacts; compra 3 e ganha 3 de vida. |
| F2 | Corpo | Nenhum. |
| F3 | Tipo | Feitiço, não-criatura, não-artefato. |
| F4 | Receptor | Nada. |
| F5 | Facilitador | Nada para outras cartas (o "reduz custo" do `mtgdb` é a affinity dela mesma — falso positivo). |
| F6 | Curva | Morta no início (turno 3 com 3 artefatos = `{2}{W}{U}`); a partir de 5 artefatos sai por `{W}{U}`. |
| F7 | Atrito | Nenhum de tap (affinity não tapa artefatos). Pede `{W}{U}`. |

Dispara Vraska, Saheeli, Whirlwind e **Seedshark com Incubate 7** — o maior Incubator do deck.

`Sai: Voyage Home — funções: [compra 3, 3 de vida, gatilho de mágica não-criatura, Incubate 7]` → compra 3 coberta por Reverse Engineer e demais; gatilho coberto pela Raid; **3 de vida fica descoberta** (é o único ganho de vida do deck, que paga vida em Talisman, Battlefield Forge e Tezzeret's Gambit); **Incubate 7 fica descoberto** (Raid dá 1).

**Por que fica em 2º**: com 42 artefatos ela é **mais barata e com menos atrito** que o Reverse
Engineer (não disputa tap com Station), ganha vida e dá o maior Incubate. E, pela regra 5:
⚠️ **foi cortada pelo pipeline na v1** ("nominal 7 parado na mão; win-more") e **o usuário a trouxe
de volta** depois da v3, sem motivo declarado. Não corte sem perguntar a ele.

---

## 5. Cobertura de ameaças (após Raid in / Reverse Engineer out)

- **Criaturas**: Swords to Plowshares, Dispatch, Stone by Sunlight, Spring-Loaded Sawblades, Warmaker Gunship, Dawnsire, Alibou + Perilous Snare. Sólido.
- **Artefatos/Encantamentos**: Sunder the Gateway, Perilous Snare, **Requisition Raid** → **3** (a Raid pega token e artefato de qualquer oponente, e resolve artefato + encantamento numa só mágica). Todas de velocidade feitiço ou ETB — nada instantâneo.
- **Combos/Spells**: **0 counterspells** — buraco que continua aberto. A única resposta instantânea a um combo em andamento é remoção de criatura (StP, Dispatch, Stone, Sawblades).
- **Planeswalker**: Perilous Snare, Dawnsire (10+), Alibou (qualquer alvo). Sem remoção dedicada desde a saída de Rip Apart.
- **Flexível**: só Perilous Snare.

---

## 6. Recontagem após a troca

| Métrica | Antes | Depois | Nota |
|---|---|---|---|
| Interação (remoção) | 9 | **10** | meta ~10 atingida |
| Respostas a artefato/encantamento | 2 | **3** | |
| Counterspells | 0 | 0 | buraco persiste |
| Wipes | 2 | 2 | no piso |
| Proteção | 7 | 7 | |
| Draw | 14 | **13** | dentro de 12–13 |
| Ramp | 11 + 1 | 11 + 1 | |
| Artefatos | 42 | **42** | nenhum dos dois é artefato |
| Criaturas de verdade / corpos para Station-crew | 27 | 27 | |
| Mágicas não-criatura (Vraska/Saheeli/Whirlwind/Seedshark) | inalterado | inalterado | feitiço por feitiço |
| Curva não-terrenos | 1:6 · 5:4 | **1:7 · 5:3** | resto igual (0:4 2:12 3:22 4:9 6:2 7+:4) |
| Pips | — | −`{U}{U}`, +`{W}` | alivia U, a cor mais pedida |
| Total | 100 | **100** | 1 entra, 1 sai |

## 7. Custo

- **Compra: R$ 0,00** — Requisition Raid está nas sobressalentes.
- Valor da lista (LigaMagic menor): sai Reverse Engineer **R$ 0,20** (cotação de 2026-09-18), entra Requisition Raid **R$ 0,44** (cotação de 2026-09-23) → **+R$ 0,24**. A lista passa de R$ 252,22 para **R$ 252,46 + Vraska, Soul of Stone (a cotar)**; segue R$ 52,46+ acima do teto de torneio. Se a saída for Voyage Home (candidata 2), o impacto é o mesmo (R$ 0,20, cotação de 2026-09-18).
- Reverse Engineer sai do deck montado, mas **não** entra em `data/collection.tsv` por inferência (regra 7): só se o usuário a declarar na próxima `/update-collection`.
