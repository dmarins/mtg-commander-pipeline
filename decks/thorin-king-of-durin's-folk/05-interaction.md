# Interação — Thorin, King of Durin's Folk (improve · rodada de custo)

Data: 2026-09-16 · oracle puxado nesta sessão via `bin/mtgdb deck ... -full` e `mtgdb oracle`.
Preços: `LigaMagic (menor)` quando havia cotação em `mtgdb prices`; senão `estimativa (Scryfall)` = US$ × 5,5, **só peneira** — o orquestrador confere na LigaMagic.

Interação já no deck: **14 por rótulo** (8 remoção + 6 proteção), das quais **~9 são de fato eficientes** · Wipes: **1/2–4** (Easy Pickings; Ori é wipe só de artefato/encantamento)

## Diagnóstico do mix atual (conferido com o oracle)

| Peça | O que faz de verdade | Nota |
|---|---|---|
| Chainsaw | 3 de dano numa criatura ao entrar | boa, barata |
| Thorin, Mountain-king | dano = poder da criatura equipada, se anexar equipamento | condicional (precisa de Equipment em campo) |
| Magnificent End | 5 de dano; custa {W}{1} contra criatura virada | boa |
| Thorin's Last Stand | modo destruir artefato/encantamento (instant) | boa, 4 manas |
| Deconstruction Hammer | {3}, {T}, sacrificar: destruir artefato/encantamento | lenta (equipar + 3 + tap) |
| Ori, Plate Stacker | ETB destrói artefatos/encantamentos **só dos oponentes** | wipe unilateral de A/E, 7 manas |
| The Black Arrow | 1 de dano a qualquer alvo (mata Dragon) | marginal |
| Dire Flail | só vira remoção após craft de 5 manas (Blunderbuss sacrifica artefato — Treasure — ao atacar) | lenta, mas repetível |
| Glóin → Easy Pickings | 1 de dano em cada criatura dos oponentes | único wipe de criatura; só limpa fichas 1/1 |
| Bofur / Concerted Care | hexproof + indestrutível numa peça (instant) | boa |
| Swiftfoot Boots | hexproof + haste | boa |
| Thorin Oakenshield / Dwarven Mattock | ward {1} | taxa leve |
| Dáin, Lord of the Iron Hills | taxa de {1} por atacante (storied) | proteção de vida, não de board |
| The Eagles Are Coming! | salva 1 criatura (ou todas, kicked a 7 manas) | anti-wipe só no kicker |

**Buracos:**
1. **Wipes: 1/2–4.** E o único é 1 de dano.
2. **Remoção de criatura sem condição e barata:** só Chainsaw e Magnificent End. Nada exila.
3. **Proteção contra board wipe** (o que mais pune um deck largo de Anões equipados): só Eagles kicked (7 manas). Nenhuma proteção de time a 2 manas.
4. Planeswalkers: só dano de combate e The Black Arrow (1). Aceitável num deck agressivo — não é prioridade.
5. Combos/stax: RW sem counters; resposta é A/E (Ori, Last Stand, Hammer) + relógio rápido. Coberto o suficiente.

## Candidatas — remoção / proteção / counters

| Carta | CMC | Tipo | Subcategoria | O que responde | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|---|---|
| **Fairgrounds Warden** | 3 | Creature — Dwarf Soldier 1/3 | remoção (exílio) | qualquer criatura de oponente (até sair de campo) | **é Anão** (queixa: faltam Anões) · Treasure do Thorin → anthem · Bifur dobra o ETB (**2 exílios**) · Depala/Magda lordes · Herald's Horn reduz e revela · Dáin's Company e Depala encontram · Kíli compra · Fíli gera ficha · Balin dispara · Cloudshift/Slip On the Ring reaproveitam o ETB | não — compra | ~R$ 0,66 estimativa (Scryfall) |
| **Duty Beyond Death** | 2 | Instant | proteção de time | wipes de "destroy" e de dano, remoção pontual (sacrifica o alvo como custo e ainda distribui contadores) | +1/+1 em **todas** as criaturas (soma com Thorin, Magda, Depala, Bifur) · sacrifica ficha de Anão (Fíli, Shortsword, Dwarven Mine, Lonely Mountain) · rev counter no Chainsaw | **sim — sideboard físico** | R$ 0 |
| **Mizzium Mortars** | 2 (6 overload) | Sorcery | remoção + wipe unilateral | 4 de dano a criatura que você não controla; overload = 4 em **cada** criatura dos oponentes | Treasures (Thorin, Magda, Dori, Gimli, Orcrist) pagam o overload · limpa bloqueadores antes do ataque com golpe duplo (Reyav, Dáin Ironfoot) · rev counters no Chainsaw · poupa o próprio board | não — compra | ~R$ 3,08 estimativa (Scryfall) |
| Stone by Sunlight *(reserva)* | 2 | Instant | remoção / proteção | criatura com poder ≥4 **ou** indestrutível a uma peça sua | modal cobre dois buracos; alvo vira artefato (storied) | maybeboard (posse incerta) | R$ 0,88 LigaMagic (menor), 2026-08-12 |
| Iroh's Demonstration *(reserva)* | 2 | Sorcery — Lesson | remoção / mini-wipe | 4 de dano numa criatura **ou** 1 em cada criatura dos oponentes | modal barato; limpa fichas antes do ataque | não — compra | ~R$ 1,70 estimativa (Scryfall) |
| Overwhelming Surge *(reserva)* | 3 | Instant | remoção flexível | 3 de dano numa criatura **e/ou** destruir artefato não-criatura | cobre criatura + artefato num slot | maybeboard (posse incerta) | R$ 0,03 LigaMagic (menor), 2026-08-12 |
| Shattered Acolyte *(reserva)* | 2 | Creature — Dwarf Warlock 2/2 | remoção A/E | artefato ou encantamento ({1}, sacrificar) | **é Anão** · Treasure do Thorin · lifelink | não — compra | sem cotação |
| Valorous Stance *(reserva)* | 2 | Instant | remoção / proteção | criatura com resistência ≥4 **ou** indestrutível | modal | não — compra | ~R$ 1,10 estimativa (Scryfall) |
| Lightning Strike *(reserva)* | 2 | Instant | remoção | 3 de dano a qualquer alvo (inclui planeswalker) | só o efeito — sem sobreposição | **sim — sideboard físico** | R$ 0 |

Counters: não se aplica (identidade RW).

## Candidatas — board wipes

| Carta | CMC | Simétrico? | Sinergias | Na coleção? | Preço |
|---|---|---|---|---|---|
| **Mizzium Mortars** (overload) | 6 | **não** — só criaturas que você não controla | Treasures pagam · abre caminho pro alfa com golpe duplo · também é remoção pontual a 2 | não | ~R$ 3,08 estimativa (Scryfall) |
| **Cosmotronic Wave** | 4 | **não** — só oponentes | 1 de dano **+ criaturas dos oponentes não bloqueiam** = turno de alfa com Reyav/Dáin Ironfoot/Dragon Throne · mata fichas e dorks · rev counters no Chainsaw | não | ~R$ 0,94 estimativa (Scryfall) |
| Iroh's Demonstration *(reserva)* | 2 | não | mesma função da Easy Pickings, mas modal | não | ~R$ 1,70 estimativa (Scryfall) |
| Radiating Lightning *(descartada)* | 4 | não, mas **um só oponente** | 3 no jogador + 1 nas criaturas dele | sim | R$ 0 |
| Seismic Wave *(descartada)* | 3 | não, mas **um só oponente** e só não-artefato | 2 a qualquer alvo + 1 | sim | R$ 0 |
| Ratchet Bomb *(descartada)* | 2 | **sim** | com 0 contadores destrói nossos Treasures, Axes e fichas de Anão; com 2, acerta 29 cartas nossas de MV 2 | maybeboard | R$ 1,79 LigaMagic (menor), 2026-08-12 |
| Delayed Blast Fireball *(descartada)* | 3 | não | a melhor da lista (instant, 2 ou 5 de dano) | não | ~R$ 90 estimativa (Scryfall) — **fora do objetivo de custo** |
| Winds of Abandon *(descartada)* | 6 | não | exila tudo dos oponentes, mas dá terrenos a eles | não | ~R$ 29 estimativa (Scryfall) — cara para a sessão |

**Por que as da coleção foram dispensadas como wipe (regra 7):** Radiating Lightning e Seismic Wave só atingem **um** oponente numa mesa de quatro. Fazem o papel de remoção, não de reset. Ratchet Bomb é simétrica e, justamente neste deck, o número "seguro" (0) destrói os nossos Treasures — o recurso que alimenta o anthem do comandante. Nenhuma das três cobre a função "wipe unilateral que atinge a mesa toda".

## Cobertura de ameaças (depois dos swaps propostos)

- **Criaturas:** Chainsaw, Mizzium Mortars, Fairgrounds Warden (exílio), Magnificent End, Thorin Mountain-king, Dire Flail (Blunderbuss), The Black Arrow (x/1) · em massa: Mizzium overload, Cosmotronic Wave, Easy Pickings
- **Artefatos/Encantamentos:** Thorin's Last Stand (instant), Deconstruction Hammer, Ori (em massa, só dos oponentes) — 3 respostas, meta ≥2 atendida
- **Combos/Spells:** sem counter (RW). A/E instant (Last Stand) + proteção de peça (Concerted Care, Swiftfoot) + relógio agressivo
- **Flexível:** Fairgrounds Warden (qualquer criatura por exílio). Não há "destroy target permanent" barato na coleção; Chaos Warp (R$ 3,81 LigaMagic, 2026-08-24) fica como opção se o orquestrador quiser — ver reservas
- **Proteção de board contra wipe:** Duty Beyond Death (2 manas, instant) + Eagles kicked · proteção de peça: Concerted Care, Swiftfoot Boots, Slip On the Ring/Cloudshift, ward (Oakenshield, Mattock)

## Swaps propostos (improve)

Condições assumidas **dos dois lados** de cada troca: Thorin em campo, 1–3 artifact tokens (Treasures), Magda e Depala em campo, Reyav em campo, Bifur com storied ativo, 18 Equipments na lista atual.

### Swap 1 — sai Smaug's Fury → entra Mizzium Mortars · **corte limpo**

**Ficha de Smaug's Fury** ({1}{R}, Instant)
- F1: alvo ganha +3/+0, alcance e iniciativa até o fim do turno. Mais nada.
- F2: não é corpo.
- F3: instant — nada no deck conta instants ou sorceries (Erebor Flamesmith e Radiant Scrollwielder não estão na lista).
- F4: não recebe nada.
- F5: +3/+0, iniciativa e alcance numa criatura por um turno (vence um combate ou segura um voador).
- F6: 2 manas, útil do turno 2 em diante, mas não destrava nada.
- F7: a iniciativa é **redundante** com o golpe duplo de Reyav e Dáin Ironfoot (golpe duplo já inclui dano de iniciativa); compete com Vow to Erebor pelo slot de truque de combate.

**Mapa de cobertura:** +3/+0 → Vow to Erebor (+2/+2, desvira e anexa equipamento), Dragon Throne, Dwarven Provisioner · iniciativa → Reyav/Dáin Ironfoot (golpe duplo), Nori, Dwalin, Looming Spires · alcance → The Black Arrow (equipado dá alcance), Iron Hills Stalwart · "ganhar um combate contra um bloqueador" → **Mizzium Mortars faz melhor** (mata o bloqueador por 2 manas). Nenhuma função descoberta.

**Entrada:** Mizzium Mortars — sinergias: (1) Treasures de Thorin/Magda/Dori/Gimli/Orcrist pagam o overload de 6; (2) wipe unilateral que abre caminho para o ataque com golpe duplo (Reyav, Dáin Ironfoot) sem tocar nos nossos Anões; (3) rev counters no Chainsaw. Cobre dois buracos (wipe + remoção barata).
**Coleção considerada:** Bombard (4 de dano, 3 manas, sideboard físico) faz a parte pontual, mas não tem overload — não fecha o buraco de wipe. Radiating Lightning/Seismic Wave: só um oponente (ver acima).
**Preço:** ~R$ 3,08 estimativa (Scryfall). Smaug's Fury já é do usuário (não está na lista de faltantes) → sem economia, custo líquido ~R$ 3.
**Simetria:** Fury avaliada com Reyav em campo (é aí que a iniciativa dela fica redundante); Mortars avaliada com os mesmos Treasures que Fury não aproveita.

### Swap 2 — sai Moment of Glory → entra Duty Beyond Death · **corte limpo** · coleção

**Ficha de Moment of Glory** ({W}, Sorcery)
- F1: contador +1/+1 numa criatura sua; se conjurada do cemitério, também um contador em cada **outra** criatura sua. Flashback {4}{W}.
- F2: não é corpo.
- F3: sorcery — nada conta.
- F4: não recebe nada.
- F5: contador permanente (1 alvo a 1 mana; o time todo a 5 manas via flashback).
- F6: 1 mana, mas feitiço de 1 contador no turno 1–2 tem impacto mínimo; o flashback só vale do turno 5 em diante.
- F7: velocidade de feitiço — não reage a remoção nem a wipe; disputa mana do turno 5 com Balin, Gandalf e o comandante.

**Mapa de cobertura:** contador em 1 criatura → Bifur (ETB/ataque), Armory of Iroas, Iron Hills (terreno), Bagel and Schmear · contador no time todo → **Duty Beyond Death faz o mesmo por 2 manas, instant** · uso duplo (cast + flashback) → descoberto (custo aceito: Duty é de uso único, mas o efeito de time chega 3 manas mais barato e ainda protege). Nenhuma função relevante descoberta.

**Entrada:** Duty Beyond Death — sinergias: (1) +1/+1 em todas as criaturas, somando com Thorin, Magda e Depala; (2) indestrutível no time todo em resposta a wipe — o único anti-wipe barato do deck; (3) o custo de sacrifício usa fichas de Anão (Fíli, Dwarven Shortsword, Dwarven Mine, The Lonely Mountain) ou a própria criatura alvo de remoção por exílio; (4) rev counter no Chainsaw.
**Coleção:** **é do usuário (sideboard físico)** — R$ 0. Alternativa de compra considerada e não necessária: Boros Charm (~R$ 25 estimativa (Scryfall), sem custo de sacrifício) — a coleção cobre a função.
**Simetria:** as duas avaliadas com board largo e anthems em campo; o contador de Moment of Glory recebeu o mesmo crédito dado aos contadores de Duty.

### Swap 3 — sai Long-Lost Lances → entra Fairgrounds Warden · **corte condicionado** (densidade de Equipment 18→17, função de tema)

**Ficha de Long-Lost Lances** ({2}, Artifact — Equipment, equipar {2})
- F1: equipada ganha +2/+0; **no seu turno**, todas as suas criaturas equipadas têm iniciativa e **vigilância**.
- F2: não é corpo.
- F3: Equipment/artefato → conta para Sram (compra ao conjurar), Kíli (compra quando Equipment entra), Dwalin (hone counter), The Lonely Mountain (entra desvirada e reduz custo), Thorin Mountain-king e Iron Hills Stalwart (anexar), Vow to Erebor, storied (artefato), Gandalf (Equipment entrando dispara Kíli uma vez a mais).
- F4: recebe hone counters (Dwalin).
- F5: +2/+0; iniciativa e vigilância em massa para o time equipado.
- F6: 2 manas + equipar 2; efeito de time só com vários equipados (turno 4+).
- F7: **atrito forte** — a vigilância em massa impede que Anões equipados virem ao atacar e **anula o Treasure da Magda** e **o gatilho da Depala** (as duas dependem de "becomes tapped"). Foi o mesmo motivo que tirou Ori, Keeper of Songs hoje. A iniciativa é redundante com o golpe duplo de Reyav e Dáin Ironfoot.

**Mapa de cobertura:** +2/+0 → Dire Flail, Dwarven Mattock, Orcrist, Grafted Wargear · iniciativa → Reyav/Dáin Ironfoot (golpe duplo), Nori, Dwalin · vigilância → **dispensável de propósito** (é atrito) · defesa após o ataque que a vigilância dava → Dáin, Lord of the Iron Hills (taxa de ataque), Dáin's Company (lifelink) · gatilhos de Equipment (Sram, Kíli, Dwalin, Lonely Mountain, Vow, Mountain-king, Stalwart, Reyav) → **17 Equipments restantes**; a densidade cai 1/18 (~6%) — **custo declarado, função de tema fora da minha especialidade → decisão do orquestrador**.

**Entrada:** Fairgrounds Warden (Dwarf Soldier 1/3) — sinergias: (1) **é Anão**, então ataca a queixa de faltar Anão (+1 na contagem); (2) Treasure do Thorin ao entrar → anthem; (3) com Bifur e storied, o ETB dispara duas vezes → **exila duas criaturas**; (4) Depala +1/+1 e Magda +1/+0 (ao atacar vira e gera Treasure — sem a vigilância da Lances no caminho); (5) Herald's Horn reduz o custo e revela; Dáin's Company e Depala conseguem encontrá-la; Kíli compra quando ela entra; Fíli cria ficha; Balin dispara; Cloudshift/Slip On the Ring reaproveitam o ETB.
**Coleção considerada:** Bombard, Flame Slash (maybeboard), Magma Spray, Stone by Sunlight (maybeboard) — todas removem criatura, mas **nenhuma é Anão**: gastar o slot com elas não ajuda na queixa principal nem aproveita Thorin/Bifur/Herald's Horn. Por isso a compra se justifica aqui.
**Preço:** ~R$ 0,66 estimativa (Scryfall). **Long-Lost Lances está na lista de faltantes** (R$ 5,90 LigaMagic (menor), 2026-09-16) → **economia líquida ~R$ 5,24**.
**Simetria:** Lances avaliada com Magda e Depala em campo (onde a vigilância custa Treasures e cartas) e com Reyav (onde a iniciativa é redundante); Warden avaliada com os mesmos lordes e com Bifur — sem o bônus de Gandalf, que não se aplica (Warden não é lendária).
**Risco:** se Warden sair de campo, a criatura exilada volta. Com corpo 1/3 (3/4 com Magda + Depala), ela morre para wipe — Duty Beyond Death (Swap 2) protege.

### Swap 4 (opcional) — sai S.H.I.E.L.D. Spy Kit → entra Cosmotronic Wave · **corte condicionado** (densidade de Equipment 17→16, função de tema)

**Ficha de S.H.I.E.L.D. Spy Kit** ({W}, Artifact — Equipment, equipar {1})
- F1: +1/+1; sempre que a criatura equipada ataca **sozinha**, desvira e scry 1.
- F2: não é corpo.
- F3: Equipment/artefato a 1 mana → Sram compra no turno 2–3, Kíli, Dwalin, Lonely Mountain, Reyav (equipado + atacando = golpe duplo), storied.
- F4: recebe hone counter.
- F5: +1/+1; desvirar + scry só atacando sozinho.
- F6: 1 mana + equipar 1: o equipamento mais barato de ligar do deck, junto com Spatula/Hammer/Dire Flail.
- F7: "ataca sozinha" **contradiz o plano de atacar em massa**; a desvirada depois do ataque não gera novo Treasure da Magda (o gatilho já aconteceu).

**Mapa de cobertura:** Equipment de 1 mana e equipar 1 para Sram/Kíli/Reyav no início do jogo → Well-Worn Spatula, Deconstruction Hammer, Dire Flail (os três com custo {1}, equipar {1}) · +1/+1 → Spatula, Hammer · scry/desvirar → dispensável (depende de atacar sozinho) · densidade de Equipment → 16 — **custo declarado, decisão do orquestrador**.

**Entrada:** Cosmotronic Wave — sinergias: (1) wipe unilateral de 1 de dano (fichas, dorks, x/1) que **poupa todo o nosso board**; (2) **criaturas dos oponentes não bloqueiam neste turno** → turno de alfa com Reyav/Dáin Ironfoot (golpe duplo), Dragon Throne e o anthem do Thorin — funciona também como finalizador; (3) rev counters no Chainsaw.
**Coleção considerada:** Radiating Lightning (sideboard físico) é o mais parecido — 1 de dano nas criaturas + 3 no jogador —, mas atinge **um** oponente e não impede bloqueio. Seismic Wave idem. Nenhuma das duas cumpre o papel de finalizador.
**Preço:** ~R$ 0,94 estimativa (Scryfall). Spy Kit já é do usuário → sem economia.
**Por que opcional:** é o segundo Equipment cortado na rodada. Se o orquestrador (ou o tema) achar que 16 Equipments deixa Sram/Reyav fracos, fique só com os swaps 1–3: o deck chega a 2 wipes e bate a meta mínima.

### Resultado após os swaps

| | Antes | Swaps 1–3 | Swaps 1–4 |
|---|---|---|---|
| Remoção (rótulo) | 8 | **10** (+Mortars, +Warden) | 10 |
| Wipes de criatura | 1 | **2** (Easy Pickings, Mortars overload) | **3** (+Cosmotronic Wave) |
| Wipe de A/E unilateral | 1 (Ori) | 1 | 1 |
| Proteção | 6 | **7** (+Duty Beyond Death) | 7 |
| Anões | 26 | **27** | 27 |
| Equipments | 18 | 17 | 16 |
| Terrenos | 36 | 36 | 36 |
| Custo | — | ~R$ 3,74 compra − R$ 5,90 economizados (Lances) | + ~R$ 0,94 |
| Curva (1/2/3/4) | 11/29/10/7 | 10/29/11/7 | 9/29/11/8 |

Nenhuma das cartas que saem ou entram consta em `decisions.md` como corte ou entrada anterior. Nenhuma é intocável nem Anão.

## Reservas

1. **Iroh's Demonstration** (~R$ 1,70 estimativa (Scryfall)) — substitui o Swap 4 se o orquestrador preferir 2 manas modais (4 de dano ou 1 em massa) a um finalizador de 4 manas.
2. **Stone by Sunlight** (maybeboard, R$ 0,88 LigaMagic 2026-08-12) — remoção de poder ≥4 ou proteção indestrutível; confirmar posse física.
3. **Overwhelming Surge** (maybeboard, R$ 0,03 LigaMagic 2026-08-12) — criatura + artefato num slot; confirmar posse.
4. **Shattered Acolyte** (compra, sem cotação) — Anão de 2 manas que destrói A/E; entra se alguém quiser trocar Deconstruction Hammer mantendo a contagem de Anões.
5. **Duergar Hedge-Mage** (R$ 2,17 LigaMagic 2026-09-16) — Anão, A/E duplo (dobrado pela Bifur). Histórico: foi oferecido **no lugar de Ori** e o usuário recusou; não foi cortado nem rejeitado como carta. Só volta como adição, nunca contra Ori.
6. **Chaos Warp** (R$ 3,81 LigaMagic 2026-08-24) — a única resposta a "qualquer permanente" barata em vermelho; aleatória, só se aparecer um buraco de encantamento/planeswalker.
