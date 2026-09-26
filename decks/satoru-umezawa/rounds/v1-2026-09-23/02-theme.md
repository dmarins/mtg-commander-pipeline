# Análise Temática — Satoru Umezawa (v1 · build)

Fase 2 · `theme-analyst` · 2026-09-23. Oracle e rulings puxados nesta sessão via `bin/mtgdb`.
Preços: **LigaMagic (menor), via `mtgdb prices`**. Das 45 candidatas do pool, só 3 já tinham
cotação, e as três estão na caixa. Todo o resto vai como **`a cotar`**, sem estimativa (regra 2).
Nesta sessão não havia ferramenta de navegador para capturar preços.

---

## Comandante — análise linha a linha

`{1}{U}{B}` · 2/4 · Legendary Creature — Human Ninja

| Linha/habilidade | Gatilho/termo | O que habilita |
|---|---|---|
| "Each creature card in your hand has ninjutsu {2}{U}{B}." | **ninjutsu**, *creature card in hand* | Todo bichão da mão entra por 4 manas, **sem ser conjurado**. O CMC impresso não pesa, mas **gatilhos "when you cast" e "if you cast it" não disparam**. Só funciona com o Satoru **no campo**: a habilidade é estática dele. |
| (ruling) ninjutsu só depois de declarar bloqueadores | **unblocked attacker** | É preciso uma criatura atacante **que passe sem bloqueio**. Por isso o deck vive de ativadores evasivos baratos (flying, unblockable, shadow, fear). |
| (ruling) o ninja entra atacante, mas **não foi declarado como atacante** | *whenever attacks* ✗ | Gatilhos de ataque (Archon, Grave Titan, anulação) **não disparam** no turno do ninjutsu. O bichão precisa valer pelo **ETB** ou pelo **dano de combate**. |
| (ruling) entra depois dos bloqueadores = **não bloqueado** | *deals combat damage to a player* ✓ | O gatilho de dano de combate fica **garantido** no turno da entrada: Lord of the Void, Ancient Brass/Silver Dragon e Fallen Shinobi disparam sem depender de evasão. |
| (ruling) ninjutsu **devolve o atacante à mão** | *return to hand* | Custo que vira recurso. Um ativador com ETB (Faerie Seer, Baleful Strix, Mulldrifter, Pilgrim's Eye) é reconjurado e repete o ETB. Um bichão que sobreviveu e atacou sem bloqueio no turno seguinte pode ser devolvido por outro ninjutsu e **reentrar**, repetindo o ETB. |
| "Whenever you activate a ninjutsu ability, look at the top three… put one into your hand… **only once each turn**." | **qualquer** habilidade de ninjutsu, não só a que o Satoru concede | Ninjas nativos baratos (Moon-Circuit Hacker `{U}`, Ninja of the Deep Hours `{1}{U}`, Silver-Fur Master `{U}{B}`) disparam o impulse de 3 pagando 1–2 manas. É seleção de topo 1×/turno. Um segundo ninjutsu no turno não gera carta extra por essa linha. |
| (ruling) se o Satoru sair em resposta, o ninjutsu já ativado resolve | resiliência | A remoção precisa vir **antes** da ativação. Proteção a Satoru no turno do ataque vale ouro (tema para a Fase 5). |
| Corpo 2/4 Human **Ninja** | tipo Ninja | Conta para Ingenious Infiltrator e recebe o anthem de Silver-Fur Master. Fica em casa como bloqueador: ele precisa **estar vivo** para conceder ninjutsu. |

**O que um bichão precisa ter para valer a troca** (4 manas + devolver um atacante):
1. **ETB forte e independente de ter sido conjurado**: remoção, roubo, tutor, compra, reanimação, board wipe unilateral. Ou
2. **gatilho de dano de combate a jogador**, que fica garantido na entrada. Ou
3. um **efeito estático** que já pague a partida por estar no campo (raro no orçamento).
4. Bônus: **evasão própria** (flying, intimidate). No turno seguinte ele ataca, passa e pode ser trocado de novo, **reciclando o ETB**.

**Anti-sinergias confirmadas no oracle** (ficam fora do pool):
- *"when you cast"*: Ulamog (ambos), Kozilek, Emrakul, Nulldrifter, Desolation Twin.
- *"if you cast it"* ou *"if {U}{U} was spent"*: Deceit, Shard of the Nightbringer, Bringer of the Last Gift, Deathbringer Regent, Sunderflock, Cyclone Summoner, Transcendent Dragon, Doomsday Excruciator.
- **Phage the Untouchable**: *"When Phage enters, if you didn't cast it from your hand, you lose the game."* Via ninjutsu, **perde o jogo**. Não pode entrar no deck em hipótese nenhuma.
- Annihilator e "whenever attacks" só disparam a partir do **segundo** turno em campo, e nunca na entrada.

## Termos de busca do tema

| # | Termo | Busca usada |
|---|---|---|
| 1 | ninjutsu / Ninja | `mtgdb tag ninjutsu -id UB`, `typal-ninja`, `synergy-ninjutsu` |
| 2 | unblockable | `mtgdb tag unblockable -id UB`, `gives-unblockable` |
| 3 | evasão barata + ETB | `mtgdb search '"enters" AND flying' -type Creature -cmc-max 2 -id UB` |
| 4 | ETB de criatura grande | SQL no banco local: `cmc>=6`, `id⊆UB`, oracle com `enters`, ordenado por EDHREC |
| 5 | dano de combate a jogador | mesma consulta com `combat damage to a player` (**sinergia dupla**: garantido via ninjutsu + reusável) |
| 6 | reciclar ETB (bounce próprio / blink) | Kaito, Dancing Shadow · Thassa, Deep-Dwelling · Palinchron · Conjurer's Closet |
| 7 | "não conjurado" | Satoru, the Infiltrator (`amount-spent-matters`) |
| 8 | manipular topo | ETB scry/surveil (Faerie Seer, Dream Eater) antes do impulse de 3 |

---

## Sobressalentes da coleção — varredura (regra 7)

`mtgdb collection -list`: 116 cartas, **39 dentro da identidade UB/incolor**. Varridas antes de qualquer busca.

**Entram no pool (4):**

| Carta | Por que entra | Preço |
|---|---|---|
| Diversion Unit | 2/1 flier por 2 (ativador) **+** sacrifício que taxa uma remoção instantânea no turno do ninjutsu. Artefato-criatura: reciclável pelo próprio ninjutsu. | na caixa (R$ 0,45 em 2026-08-12) |
| Pilgrim's Eye | 1/1 flier (ativador) **+** ETB que busca básico. Cada vez que o ninjutsu o devolve, ele é reconjurado e busca outro terreno. | na caixa (R$ 0,05 em 2026-08-12) |
| Reconnaissance Mission | Compra por **cada** criatura que conecta, e o deck ataca com várias criaturas pequenas evasivas **mais** o bichão (que conecta garantido). Tem cycling `{2}` quando sobra. | na caixa (sem cotação; custo 0) |
| Bident of Thassa | Mesmo motor de compra por dano de combate; é artefato-encantamento lendário. A ativação que força os oponentes a atacar abre a defesa deles para o ataque seguinte. | na caixa (R$ 5,80 em 2026-08-12) |

**Dispensadas, com motivo escrito:**

| Sobressalente | Motivo da dispensa (eixo da ficha) |
|---|---|
| Thrummingbird | F1/F3: é ativador válido (flier 1/1 por 2), mas o proliferate não tem contador para multiplicar. **1 ponto de sinergia só**, e Faerie Seer e Spectral Sailor ocupam o mesmo slot com 2. **Primeira reserva** se faltar ativador: custo zero. |
| Somber Hoverguard | F6: affinity para artefatos num deck com poucos artefatos custa ~5–6 manas por um 3/2 flier. Pelo ninjutsu vira um vanilla sem ETB. Falha nos dois papéis. |
| Cargo Ship / Thopter Fabricator | F7: veículos precisam de crew, que **tapa** a criatura que deveria atacar para o ninjutsu. O gatilho do Fabricator pede segunda compra no turno, e o Satoru põe na mão sem comprar. Atrito direto com o plano. |
| Elrond, Moon-Reader | F1: "ability **of a creature**" significa criatura no campo (regra 109.2). O ninjutsu é ativado de um **card** na mão e não dispara o Elrond. Sobra um 3/3 sem evasão. |
| Phyrexian Revoker, Lesser Masticore, Myr Convert, Hedron Crawler, Master's Councillors | F2: corpos sem evasão. Não conectam e, portanto, não ativam ninjutsu. Myr Convert e Hedron Crawler ficam anotados para a **Fase 4** (ramp). |
| Titan Forge | O Golem 9/9 é ficha. Devolvido à mão pelo ninjutsu, ele **deixa de existir**. Sem sinergia. |
| Darksteel Reactor, Lux Artillery, Lux Cannon, Moxite Refinery, Liquimetal Coating, Oracle's Vault, Bargaining Table, Thaumaton Torpedo, Leonin Bola, Seer's Lantern, Magnifying Glass | Fora do tema. Menos de 2 pontos de sinergia com ninjutsu ou evasão. |
| Sol Ring, Commander's Sphere, Prophetic Prism, Manalith, Sphere of the Suns, Mycosynth Wellspring | Função de ramp. **Encaminhadas à Fase 4**, não julgadas aqui. |
| Thirst for Knowledge, Experimental Augury | Draw. **Encaminhadas à Fase 3** (Augury é o mesmo "olhe 3, pegue 1" do Satoru). |
| Negate, Disruption Protocol, Stoic Rebuttal, Thranduil's Decree, Old Fat Spider Can't See Me, Culling Dais, Ratchet Bomb, Blast Zone | Interação e proteção. **Encaminhadas à Fase 5**. O Old Fat Spider dá hexproof ao Satoru e merece olhar. |

---

## Pool temático — 45 candidatas

Legenda F1–F7: **F2** corpo/tapável · **F3** tipo que alimenta · **F4** recebe · **F5** facilita · **F6** turno · **F7** atrito.
Tudo `a cotar` salvo indicação. **⚠ cotar primeiro** marca a carta que, pelo histórico de mercado, *pode* pesar
no teto. Isso **não é** estimativa de preço, é prioridade de cotação.

### (a) Ativadores — criaturas baratas que conectam (13)

| Carta | CMC | Tipo | Ficha F1–F7 resumida | Sinergias (2+) | Coleção? | Preço |
|---|---|---|---|---|---|---|
| Ornithopter | 0 | Artifact Creature — Thopter | F1 flying · F2 0/2 · F3 artefato · F4 Tetsuko (poder 0) · F6 T1, custo 0: reconjurar depois do ninjutsu é grátis · F7 nenhum | (1) ativador flier grátis; (2) volta à mão pelo ninjutsu e é reconjurado **sem gastar mana**, então o ninjutsu custa só os 4 do Satoru; (3) Tetsuko | não | a cotar |
| Changeling Outcast | 1 | Creature — Shapeshifter | F1 unblockable, não bloqueia · F3 **todo tipo** (Ninja, Rogue, Faerie) · F4 anthem Silver-Fur, conta para Ingenious Infiltrator · F6 T1 · F7 não defende | (1) ativador inbloqueável; (2) é **Ninja**: dispara Ingenious Infiltrator e recebe +1/+1 do Silver-Fur; (3) Faerie para Spellstutter Sprite, se entrar | não | a cotar |
| Faerie Seer | 1 | Creature — Faerie Wizard | F1 flying + ETB scry 2 · F6 T1 · F7 nenhum | (1) ativador flier; (2) ETB scry 2 **arruma o topo antes** do impulse de 3 do Satoru; (3) volta à mão e repete o scry | não | a cotar |
| Spectral Sailor | 1 | Creature — Spirit Pirate | F1 flash, flying, `{3}{U}`: draw · F6 T1, ou flash no fim do turno do oponente · F7 disputa mana com o ninjutsu no turno de ataque | (1) ativador flier com flash, que se esquiva de wipe de sorcery; (2) escoadouro de mana e compra **no turno do oponente** | não | a cotar |
| Slither Blade | 1 | Creature — Snake Rogue | F1 unblockable · F2 1/2 · F3 **Rogue** · F4 anthem Silver-Fur, Tetsuko · F6 T1 | (1) ativador inbloqueável; (2) Rogue recebe +1/+1 do Silver-Fur; (3) Tetsuko redundante | não | a cotar |
| Baleful Strix | 2 | Artifact Creature — Bird | F1 flying, deathtouch, ETB draw · F2 bloqueador deathtouch · F3 artefato · F6 T2 · F7 nenhum | (1) ativador flier; (2) ETB draw **repetível**: cada ninjutsu que o devolve vira +1 carta ao reconjurar; (3) defende o Satoru sozinho | não | a cotar ⚠ cotar primeiro |
| Dimir Infiltrator | 2 | Creature — Spirit | F1 unblockable + transmute `{1}{U}{B}` (busca CMC 2) · F2 1/3 · F6 T2 · F7 nenhum | (1) ativador inbloqueável; (2) **tutor** de CMC 2: Tetsuko, Satoru the Infiltrator, Silver-Fur Master, Baleful Strix, Moon-Circuit Hacker ou Invisible Stalker | não | a cotar |
| Invisible Stalker | 2 | Creature — Human Rogue | F1 hexproof + unblockable · F3 Rogue · F4 anthem Silver-Fur, Whispersilk dispensável · F6 T2 | (1) o ativador mais **resiliente** (hexproof), sobrevive a remoção pontual; (2) Rogue para o Silver-Fur; (3) Tetsuko | não | a cotar |
| Satoru, the Infiltrator | 2 | Legendary Creature — Human Ninja Rogue | F1 menace + "entra sem ser conjurado → draw" · F3 **Ninja + Rogue** · F4 Silver-Fur · F5 compra em **todo** ninjutsu · F6 T2 | (1) cada bichão que entra por ninjutsu **não foi conjurado** e dá +1 carta, somando ao impulse do Satoru; (2) dispara também com Sakashima's Student, reanimação de Gyruda/Brass Dragon e blink da Thassa; (3) é Ninja para o Ingenious. Nome diferente: sem conflito de lenda | não | a cotar ⚠ cotar primeiro |
| Mulldrifter | 5 | Creature — Elemental | F1 flying, ETB draw 2, evoke `{2}{U}` · F6 T3 (evoke) ou T5 · F7 nenhum | (1) **ativador e alvo ao mesmo tempo**: por ninjutsu (4) compra 2 e fica um flier; (2) como flier volta à mão pelo próximo ninjutsu e repete o ETB; (3) evoke cedo quando falta pressão | não | a cotar |
| Shriekmaw | 5 | Creature — Elemental | F1 fear, ETB destroy (não-artefato, não-preto), evoke `{1}{B}` · F6 T2 (evoke) | (1) ativador com **fear**; (2) remoção reciclável via ninjutsu, com o corpo 3/2 ficando; (3) evoke como remoção barata | não | a cotar |
| Diversion Unit | 2 | Artifact Creature — Robot | F1 flying, `{U}`+sac: counter instant/sorcery salvo `{3}` · F3 artefato · F4 Tetsuko (toughness 1) · F6 T2 | (1) ativador flier; (2) protege a janela de ninjutsu contra a remoção instantânea no Satoru | **sim** | R$ 0,45 (2026-08-12) |
| Pilgrim's Eye | 3 | Artifact Creature — Thopter | F1 flying, ETB busca básico · F3 artefato · F4 Tetsuko · F6 T3 · F7 compete com os 2-drops | (1) ativador flier; (2) ETB **repetível** de terreno a cada volta à mão: fixa cor e garante o 4º terreno do ninjutsu | **sim** | R$ 0,05 (2026-08-12) |

### (a') Ninjas nativos — ativam ninjutsu barato e disparam o Satoru (8)

O impulse do Satoru dispara com **qualquer** ninjutsu. Um ninja nativo de `{U}`/`{1}{U}`/`{U}{B}` compra a
seleção de topo por 1–2 manas e ainda sobra mana para outra ação no mesmo turno.

| Carta | CMC | Tipo | Ficha F1–F7 resumida | Sinergias (2+) | Coleção? | Preço |
|---|---|---|---|---|---|---|
| Moon-Circuit Hacker | 2 | Enchantment Creature — Human Ninja | F1 ninjutsu `{U}`, dano → draw (sem descarte no turno em que entra) · F3 Ninja/encantamento · F4 Silver-Fur · F6 T2 ou ninjutsu T3–4 | (1) ninjutsu de **`{U}`** dispara o Satoru por 1 mana; (2) compra ao conectar; (3) Ninja para o Ingenious | não | a cotar |
| Ninja of the Deep Hours | 4 | Creature — Human Ninja | F1 ninjutsu `{1}{U}`, dano → draw · F6 ninjutsu T3–4 | (1) Satoru por 2 manas; (2) compra por conexão; (3) Ninja/Silver-Fur/Ingenious | não | a cotar |
| Silver-Fur Master | 2 | Creature — Rat Ninja | F1 ninjutsu `{U}{B}`, **ninjutsu custa `{1}` a menos**, anthem +1/+1 Ninja/Rogue · F5 **redutor de custo** · F6 T2 | (1) o ninjutsu do Satoru cai para `{1}{U}{B}`, e o de Moon-Circuit para `{U}`; (2) anthem para Changeling, Slither Blade, Invisible Stalker, os dois Satorus e os ninjas; (3) dispara o Satoru | não | a cotar |
| Ingenious Infiltrator | 4 | Creature — Vedalken Ninja | F1 ninjutsu `{U}{B}`, **Ninja** que conecta → draw · F6 ninjutsu T3 | (1) motor de compra: Changeling, os dois Satorus e os nativos são Ninjas; (2) dispara o Satoru por 2 manas | não | a cotar |
| Thousand-Faced Shadow | 1 | Creature — Human Ninja | F1 flying; ninjutsu `{2}{U}{U}`; ao entrar da mão atacando, **token cópia** de outro atacante · F2 1/1 · F4 Tetsuko · F6 T1 como ativador | (1) ativador flier de 1 mana; (2) no turno grande, copia o **bichão** que acabou de entrar: o token repete o ETB (Agent, Chupacabra, Gearhulk); (3) Ninja | não | a cotar |
| Sakashima's Student | 4 | Creature — Human Ninja | F1 ninjutsu `{1}{U}`, entra como **cópia** de qualquer criatura (e vira Ninja) · F6 ninjutsu T3+ | (1) cópia do melhor bichão de qualquer campo, **com o ETB** dele, por 2 manas; (2) dispara o Satoru; (3) vira Ninja para o Ingenious. É o "bichão barato" do deck | não | a cotar ⚠ cotar primeiro |
| Fallen Shinobi | 5 | Creature — Zombie Ninja | F1 ninjutsu `{2}{U}{B}`, dano → oponente exila 2 do topo e você joga **de graça** · F2 5/4 | (1) dano de combate **garantido** na entrada = 2 cartas grátis; (2) Ninja para Silver-Fur/Ingenious | não | a cotar ⚠ cotar primeiro |
| Kaito, Bane of Nightmares | 4 | Legendary Planeswalker — Kaito | F1 ninjutsu `{1}{U}{B}`; no seu turno é Ninja 3/4 hexproof; +1 emblema anthem Ninja; 0 surveil 2 + draw por oponente que perdeu vida; −2 stun · F3 planeswalker | (1) ninjutsu próprio dispara o Satoru; (2) o 0 compra depois de um turno de conexões; (3) −2 trava um bloqueador; (4) emblema para Ninjas | não | a cotar ⚠ cotar primeiro |

### (b) Alvos de ninjutsu — bichões que valem sem ter sido conjurados (18)

Todos entram por `{2}{U}{B}` (`{1}{U}{B}` com Silver-Fur). **Nenhum depende de "cast".**

| Carta | CMC | Tipo | Ficha F1–F7 resumida | Sinergias (2+) | Coleção? | Preço |
|---|---|---|---|---|---|---|
| Ravenous Chupacabra | 4 | Creature — Beast Horror | F1 ETB destroy criatura do oponente · F2 2/2 · F6 conjurável T4 também · F7 sem evasão, difícil de reciclar | (1) remoção incondicional; (2) é barata também **sem** Satoru (plano B); (3) alvo favorito de Sakashima/Thousand-Faced | não | a cotar |
| Hostage Taker | 4 | Creature — Human Pirate | F1 ETB exila criatura/artefato até sair; você pode conjurar · F7 se o ninjutsu o devolver **antes** de você conjurar a carta roubada, ela volta ao dono | (1) remoção + **roubo**; (2) conjurável no T4 sem Satoru. Conjure o roubado no mesmo turno | não | a cotar ⚠ cotar primeiro |
| Noxious Gearhulk | 6 | Artifact Creature — Construct | F1 menace, ETB destroy + ganho de vida · F2 5/4 menace (reciclável) · F3 artefato | (1) remoção + vida; (2) **menace**: no turno seguinte ataca e pode ser trocado de novo, repetindo o ETB | não | a cotar |
| Massacre Wurm | 6 | Creature — Phyrexian Wurm | F1 ETB −2/−2 no campo dos oponentes, perda de 2 por criatura que morre · F2 6/5 · F7 `{B}{B}{B}` irrelevante via ninjutsu | (1) **wipe unilateral** instantâneo contra tokens e bloqueadores; (2) abre o ataque dos ativadores no mesmo combate; (3) dreno de vida (wincon parcial) | não | a cotar ⚠ cotar primeiro |
| Grave Titan | 6 | Creature — Giant | F1 deathtouch; ETB (e ataque) 2 Zombies · F2 6/6 · F7 o gatilho de **ataque** não dispara na entrada | (1) 3 corpos por 4 manas: bloqueadores que protegem o Satoru; (2) nos turnos seguintes gera mais 2 a cada ataque | não | a cotar ⚠ cotar primeiro |
| Dream Eater | 6 | Creature — Nightmare Sphinx | F1 flash, flying, ETB surveil 4 + bounce permanente do oponente · F6 conjurável em flash | (1) surveil 4 **arruma o topo** do impulse do Satoru; (2) bounce; (3) **flier**, reciclável | não | a cotar |
| Dragonlord Silumgar | 6 | Legendary Creature — Elder Dragon | F1 flying, deathtouch, ETB rouba criatura/PW enquanto ficar · F7 reciclá-lo pelo ninjutsu **devolve** o roubo | (1) roubo da melhor ameaça da mesa; (2) flier deathtouch que defende e ataca | não | a cotar ⚠ cotar primeiro |
| Gyruda, Doom of Depths | 6 | Legendary Creature — Demon Kraken | F1 ETB todos moem 4; você pega criatura de CMC **par** dos moídos · F2 6/6 | (1) pode trazer o bichão do oponente de graça; (2) a criatura trazida também "entra sem ser conjurada" (Satoru the Infiltrator) | não | a cotar |
| Marang River Regent // Coil and Catch | 6 | Creature — Dragon // Instant — Omen | F1 flying, ETB devolve **2** permanentes não-terreno; a face Omen compra 3 e descarta 1 · F7 quem entra por ninjutsu é o lado criatura | (1) bounce duplo; (2) flier reciclável; (3) a face instant é draw quando falta ativador | não | a cotar |
| Agent of Treachery | 7 | Creature — Human Rogue | F1 ETB **controle permanente** de qualquer permanente; fim de turno com 3+ alheios → compra 3 · F3 Rogue | (1) roubo **sem prazo**, então reciclar = mais roubos (sem o atrito do Silumgar); (2) com Sakashima/Thousand-Faced, 2–3 roubos liberam a compra de 3; (3) Rogue para o Silver-Fur | não | a cotar |
| Rune-Scarred Demon | 7 | Creature — Demon | F1 flying, ETB **tutor** qualquer carta · F2 6/6 flier | (1) busca a peça que falta (Kaito, Thassa, bichão, remoção); (2) flier reciclável | não | a cotar ⚠ cotar primeiro |
| Lord of the Void | 7 | Creature — Demon | F1 flying; dano a jogador → exila 7 do topo dele e põe uma criatura sob seu controle · F2 7/7 | (1) gatilho de dano **garantido** na entrada; (2) 7 de dano; (3) a criatura roubada entra sem ser conjurada (Satoru the Infiltrator) | não | a cotar |
| Ancient Brass Dragon | 7 | Creature — Elder Dragon | F1 flying; dano → d20, reanima criaturas de **qualquer** cemitério com CMC total ≤ X | (1) gatilho garantido; (2) reanima **vários bichões** do próprio cemitério (moídos por Gyruda, mortos por wipe); (3) cada um entra sem ser conjurado | não | a cotar ⚠ cotar primeiro |
| Sepulchral Primordial | 7 | Creature — Avatar | F1 intimidate, ETB um criatura de cada cemitério oponente · F2 5/4 | (1) até 3 criaturas por 4 manas; (2) **intimidate** ajuda a reciclar | não | a cotar |
| Meteor Golem | 7 | Artifact Creature — Golem | F1 ETB destroy **qualquer** permanente não-terreno do oponente · F3 artefato · F7 sem evasão | (1) remoção universal (encantamento, artefato, PW); (2) incolor, sai em qualquer ramp | não | a cotar |
| Archon of Cruelty | 8 | Creature — Archon | F1 flying; ETB/ataque: oponente sacrifica, descarta e perde 3; você compra e ganha 3 · F7 o ataque não dispara na entrada | (1) ETB **multifunção** (remoção + discard + dreno + draw); (2) flier: no turno seguinte ataca e dispara de novo | não | a cotar ⚠ cotar primeiro |
| Ancient Silver Dragon | 8 | Creature — Elder Dragon | F1 flying; dano → compra d20 e sem limite de mão · F2 8/8 | (1) gatilho **garantido**: média de 10,5 cartas; (2) enche a mão de bichões para os ninjutsus seguintes | não | a cotar ⚠ cotar primeiro |
| Palinchron | 7 | Creature — Illusion | F1 flying, ETB **desvira até 7 terrenos**, `{2}{U}{U}` volta à mão · F5 mana | (1) ninjutsu de **saldo positivo**: paga 4 e desvira 7, financiando um 2º ninjutsu no mesmo combate (Thousand-Faced, nativo); (2) **auto-bounce**: volta à mão e repete; (3) flier. Com Silver-Fur, cada ciclo Palinchron ↔ ativador fecha com +1 de mana. **Não é combo sozinho**: o impulse do Satoru é 1×/turno. A ser analisado pela Fase 7 | não | a cotar ⚠ cotar primeiro |

### (c) Habilitadores — evasão, reciclagem de ETB, compra por conexão (6)

| Carta | CMC | Tipo | Ficha F1–F7 resumida | Sinergias (2+) | Coleção? | Preço |
|---|---|---|---|---|---|---|
| Tetsuko Umezawa, Fugitive | 2 | Legendary Creature — Human Rogue | F1 criaturas com poder **ou** resistência ≤ 1 inbloqueáveis · F2 1/3, ela mesma inbloqueável · F3 Rogue · F5 **evasão em massa** | (1) quase todo ativador do pool tem poder/resistência ≤ 1 (Ornithopter, Seer, Sailor, Slither, Strix, Infiltrator, Stalker, Thousand-Faced, Diversion Unit, Pilgrim's Eye); (2) é ativadora ela mesma; (3) Rogue, recebe o Silver-Fur | não | a cotar ⚠ cotar primeiro |
| Whispersilk Cloak | 3 | Artifact — Equipment | F1 inbloqueável + shroud, equip `{2}` · F3 artefato · F5 evasão · F7 disputa mana com o ninjutsu | (1) torna **qualquer** criatura um ativador, inclusive o bichão da rodada anterior (reciclar ETB); (2) shroud protege o portador. Fica em campo quando o equipado volta à mão | não | a cotar |
| Kaito, Dancing Shadow | 4 | Legendary Planeswalker — Kaito | F1 criatura sua causa dano → **devolve à mão** (e Kaito ativa 2×); +1 trava criatura; 0 draw; −2 Drone | (1) **recicla o bichão**: depois do dano, volta à mão e é ninjutsado de novo no turno seguinte, repetindo o ETB; (2) o 0 é draw; (3) o +1 tira um bloqueador | não | a cotar ⚠ cotar primeiro |
| Thassa, Deep-Dwelling | 4 | Legendary Enchantment Creature — God | F1 indestrutível; no **fim do seu turno** exila e devolve outra criatura sua; `{3}{U}`: tap criatura · F2 só vira criatura com devoção 5 | (1) **re-ETB grátis todo turno** para o bichão (Agent, Archon, Chupacabra) ou ativador (Strix, Seer, Pilgrim's Eye); (2) quem volta pela Thassa **não foi conjurado** e dá draw com Satoru the Infiltrator; (3) o tap remove bloqueador | não | a cotar ⚠ cotar primeiro |
| Reconnaissance Mission | 4 | Enchantment | F1 criatura sua causa dano a jogador → draw; cycling `{2}` · F3 encantamento | (1) cada ativador e o bichão conectam, então 2–4 cartas por turno; (2) cycling quando a mão já está cheia | **sim** | na caixa (sem cotação) |
| Bident of Thassa | 4 | Legendary Enchantment Artifact | F1 igual à Mission + `{1}{U}`,T: criaturas oponentes atacam se puderem · F3 artefato **e** encantamento | (1) redundância do motor de compra por conexão; (2) forçar ataques esvazia a defesa dos oponentes para o seu turno | **sim** | R$ 5,80 (2026-08-12) |

**Contagem do pool:** 13 ativadores + 8 ninjas nativos + 18 alvos + 6 habilitadores = **45**.
4 vêm da caixa (Diversion Unit, Pilgrim's Eye, Reconnaissance Mission, Bident of Thassa).

### Reserva avaliada (fora do pool de 45, prontas se houver corte ou estouro de custo)

| Carta | Grupo | Por que está na reserva e não no pool |
|---|---|---|
| Kitesail Freebooter | ativador | Flier com ETB de discard reciclável, mas o efeito acaba quando ele volta à mão (F7): perde para Faerie Seer e Baleful Strix. |
| Siren Stormtamer | ativador | Flier de 1 que protege o Satoru. **Encaminhado à Fase 5** (proteção). |
| Ornithopter of Paradise | ativador | Flier 0/2 que gera mana de qualquer cor. **Encaminhado à Fase 4** (ramp). Se entrar, entra também como ativador. |
| Tormented Soul, Mist-Cloaked Herald | ativador | Inbloqueáveis de 1 com **um** ponto de sinergia. Reservas diretas de Slither Blade. |
| Yuriko, the Tiger's Shadow | ninja | Ninjutsu `{U}{B}` + dreno por MV revelado. Com bichões de CMC 6–8 no topo, o dreno é alto, mas só dispara com Ninjas (Changeling e os nativos, não os bichões). |
| Mistblade Shinobi, Moonsnare Specialist | ninja | Bounce de criatura com ninjutsu barato. Opções se o deck precisar de mais nativos. |
| Kaito Shizuki | habilitador | O −2 gera ficha Ninja 1/1 inbloqueável (ativador recorrente) e o +1 é draw. Perde slot para o Kaito, Dancing Shadow. |
| Lord of Change, Overseer of the Damned, Diluvian Primordial, Scourge of Fleets, Sphinx of Uthuun, Hoarding Broodlord, Silent-Blade Oni, Mindleech Mass | alvo | ETB ou dano válidos. São os **substitutos baratos** se os alvos marcados ⚠ estourarem o teto (tabela abaixo). |
| Conjurer's Closet | habilitador | Blink de fim de turno, igual à Thassa, mas custa 5 e não tem o tap. Substituto se a Thassa pesar. |
| Aqueous Form, Key to the City | habilitador | Evasão pontual. Reservas do Whispersilk Cloak. |
| Blightsteel Colossus | alvo/wincon | 11 de infect **garantido** mata um jogador por ninjutsu. Mas é **wincon de um alvo só** e o histórico de mercado sugere preço alto. **⚠ cotar primeiro** antes de a Fase 7 considerar. |
| Access Tunnel, Rogue's Passage | terreno | Evasão sob demanda em slot de terreno. **Encaminhados à Fase 6.** |
| Kederekt Leviathan | alvo | ETB devolve **todos** os não-terrenos, inclusive o **Satoru** e os seus. Sai do pool: reset simétrico que desmonta o próprio motor. |

---

## Tensão de orçamento — quantos bichões caros o teto comporta

O teto é **R$ 200 para as 99** (básicos contam). O tema pede **~10–12 alvos + 10–12 ativadores/ninjas + 3–5
habilitadores**, ou seja, ~25–29 slots temáticos disputando dinheiro com ~10 ramp, ~12 draw, ~10 interação e
~38 terrenos. **Sem cotação não dá para dizer quantos premium cabem**, e não vou estimar (regra 2).
A regra de decisão que proponho ao orquestrador:

1. **Cotar primeiro as 17 marcadas com ⚠** (lista no resumo). O resultado define o tamanho da cota premium.
2. **O que faz o tema funcionar custa pouco.** A troca via ninjutsu vale 4 manas por **qualquer** ETB forte,
   então a espinha pode ser montada com alvos cujo papel não depende de raridade: Ravenous Chupacabra, Meteor
   Golem, Noxious Gearhulk, Agent of Treachery, Gyruda, Marang River Regent, Dream Eater, Sepulchral
   Primordial, Lord of the Void, Mulldrifter e Shriekmaw. **Na cotação, estas devem ser a base.**
3. **Mapa de substituição por papel**, caso um premium estoure:

| Premium ⚠ | Papel | Substituto que cobre o papel |
|---|---|---|
| Archon of Cruelty | ETB de remoção + card + dreno, flier | Overseer of the Damned (remoção + fichas, flier) ou Noxious Gearhulk |
| Ancient Silver Dragon | compra massiva por dano garantido | Lord of Change (ETB compra 3, flier) + Mulldrifter |
| Ancient Brass Dragon | reanimação em massa | Sepulchral Primordial (reanima do cemitério dos oponentes) + Gyruda |
| Rune-Scarred Demon | tutor | Dimir Infiltrator (transmute CMC 2); tutor de verdade fica com a Fase 3/7 |
| Massacre Wurm / Grave Titan | wipe unilateral / corpos defensivos | Wipe fica com a Fase 5; Myr Battlesphere (4 corpos no ETB) como reserva |
| Dragonlord Silumgar / Hostage Taker | roubo | Agent of Treachery (roubo permanente, sem atrito de reciclagem) |
| Palinchron | ninjutsu de saldo positivo | **Sem substituto equivalente.** Se estourar, o deck só perde o 2º ninjutsu no mesmo turno |
| Thassa, Deep-Dwelling | re-ETB todo turno | Conjurer's Closet |
| Kaito, Dancing Shadow | reciclar bichão | Palinchron (auto-bounce) ou ninjutsu sobre o próprio bichão no turno seguinte |
| Tetsuko Umezawa | evasão em massa | Whispersilk Cloak + ativadores nativamente inbloqueáveis |
| Baleful Strix / Satoru, the Infiltrator / Sakashima's Student / Fallen Shinobi / Kaito, Bane of Nightmares | ativador de valor / ninja de valor | Faerie Seer, Moon-Circuit Hacker, Ninja of the Deep Hours, Ingenious Infiltrator |

4. **Critério de corte por custo**: um alvo premium só fica se o seu **ETB não tem equivalente** entre os
   baratos (Palinchron, Ancient Silver Dragon) **ou** se ele cobre duas categorias da meta ao mesmo tempo
   (Archon: remoção + draw; Rune-Scarred: tutor). Premium que só faz o mesmo que um barato faz "maior" é o
   primeiro a sair.

---

## Notas para as próximas fases

- **Fase 3 (draw):** o tema já gera compra: impulse do Satoru, Satoru the Infiltrator, Moon-Circuit, Ninja of the Deep Hours, Ingenious Infiltrator, Mulldrifter, Baleful Strix, Recon Mission, Bident, Lord of Change e Ancient Silver. Conte-as antes de buscar fora.
- **Fase 5 (interação):** vários alvos são remoção (Chupacabra, Gearhulk, Meteor Golem, Shriekmaw, Archon, Massacre Wurm). O ponto fraco é **proteger o Satoru antes da ativação**: Siren Stormtamer, Diversion Unit e Old Fat Spider (caixa) merecem olhar.
- **Fase 6:** Access Tunnel e Rogue's Passage cumprem a função de evasão em slot de terreno. Pilgrim's Eye reciclado busca básico.
- **Fase 7:** Palinchron + Silver-Fur Master geram mana líquida por ciclo, **sem sumidouro** no pool. Massacre Wurm drena. Blightsteel é wincon de um alvo. Nada disso é lock (regra 11 ok).
