# Interação — Otharri, Suns' Glory (v1 · 2026-09-19 · modo `improve`)

Interação já no deck: **13 nominais / 4 eficazes** — meta ~10 · Wipes: **0** — meta 2–4

> Todo texto oracle desta página foi puxado de `bin/mtgdb` **nesta sessão** (regra 6). Ruling do
> Otharri conferido em `bin/mtgdb rulings`. Preços em US$ são **estimativa (Scryfall)** e **não**
> valem como régua (regra 2); valores em R$ marcados `LM` são LigaMagic (menor) com a data da
> captura. As cotações das compras propostas **ainda não existem** — captura manual pendente.

---

## 0. Auditoria do que já está no deck

O deck não tem *pouca* interação em número — tem interação **de cobertura errada**. As 13 peças
nominais se distribuem assim:

| Carta | CMC | Vel. | O que responde de fato | Veredito |
|---|---|---|---|---|
| **Celebrate the Mountain-king** | 4 | sorcery (ETB) | exila **1 não-terreno por oponente** até sair — pega criatura, artefato, encantamento, PW | **melhor interação do deck.** Fica |
| **Reduce to Memory** | 3 | sorcery | exila qualquer não-terreno; devolve um 3/2 ao dono | fica (o 3/2 é irrelevante contra quem faz 5 corpos por ataque) |
| **Swift Reckoning** | 2 | sorcery/flash | destrói criatura **virada** | fica, mas é o slot mais fraco: no seu turno o board do oponente costuma estar **desvirado** |
| **Gideon's Triumph** | 2 | instant | edict — o oponente escolhe a **pior** criatura; metade do texto lê um PW Gideon que não existe no deck | **sai** (já cortada em K17 pelo `theme-analyst`; confirmo) |
| **Disenchant** | 2 | instant | artefato/encantamento | fica — é a mais eficiente das quatro |
| **Seal of Cleansing** | 2 | ench. | artefato/encantamento, **pré-pago** (mana adiantado, resposta grátis) | fica |
| **Expose to Daylight** | 3 | instant | artefato/encantamento + scry 1 | **sai** — 3 manas pelo que Disenchant faz por 2 |
| **Invoke the Divine** | 3 | instant | artefato/encantamento + 4 vida | **sai** — estritamente pior que Disenchant |
| **Djeru's Renunciation** | 2 | instant | **tapa** 2 criaturas (não remove nada) + cycling `{W}` | **corte condicionado** — o cycling é função de draw |
| **Gods Willing** | 1 | instant | protege 1 criatura (prot. de cor: esquiva remoção **e** bloqueio) + scry | fica |
| **Feat of Resistance** | 2 | instant | protege 1 criatura + **contador +1/+1 permanente** | fica |
| **Adamant Will** | 2 | instant | protege 1 criatura (+2/+2 e indestrutível) | **sai** — a 3ª cópia funcional do mesmo efeito |
| **Duty Beyond Death** | 2 | instant | **indestrutível ao time inteiro** + contador em cada | fica — é a melhor resposta a wipe que o deck já tem |
| **Paladin Danse** | 3 | ativada | indestrutível a **artefatos ou Humanos** | **sai** — §1 abaixo |

Contagem honesta: **1 remoção flexível boa** (Celebrate), **2 medianas** (Reduce to Memory, Swift
Reckoning), **4 art/ench-only** (redundância pura), **4 proteções** (3 delas de alvo único e quase
idênticas), **0 remoção instantânea de criatura**, **0 resposta a combo**, **0 wipes**.

**Buraco nº 1 do deck inteiro:** não existe **uma única** carta capaz de matar uma criatura do
oponente em velocidade de instantâneo. Contra a mesa, o deck é cego.

---

## 1. O achado crítico — veredito sobre `Paladin Danse, Steel Maverick`

`{2}{W}` · 3/3 · Legendary **Artifact** Creature — Synth Knight · CMC 3

```
Vigilance, lifelink
Exile Paladin Danse: Each creature you control that's an artifact or Human gains
indestructible until end of turn.
```

### Ficha F1–F7 (regra 4), sob as **mesmas** condições das entradas (Otharri com 3–4 contadores, 6–10 fichas Rebel, um anthem estático em campo)

| Eixo | Leitura |
|---|---|
| **F1** | Duas linhas: (i) vigilância e vínculo com a vida; (ii) **exila a si mesma** para dar indestrutível a cada criatura sua **que seja artefato ou Humano**. A ativação é **grátis em mana** e não usa a pilha de tap — é resposta rápida e não pode ser respondida com remoção nela |
| **F2** | Corpo 3/3 por 3 **com vigilância** — ataca e continua de pé para bloquear. Com `The Circle of Loyalty` + `Jor Kadeen` é um **7/4 com vigilância e vínculo**. Não tapa para custo (o deck não tem crew/convoke), mas a vigilância a mantém disponível |
| **F3** | **Artifact Creature** — conta **duas vezes**: metalcraft do `Jor Kadeen` (↑↑ no re-pass) e os contadores de ETB do `Luxknight Breacher`. É legendária → dispara a ficha 2/2 do `The Circle of Loyalty` ao ser conjurada |
| **F4** | Recebe todos os anthems estáticos, contadores permanentes (`Basri's Solidarity`, `Duty Beyond Death`), equipamentos e proteções |
| **F5** | Concede **indestrutível em massa** — mas só a **artefatos ou Humanos** |
| **F6** | Turno 3 — entra antes do comandante e já é um corpo relevante |
| **F7** | **Atrito fatal e verificado no oracle:** as fichas do Otharri são `2/2 red **Rebel**` e o próprio Otharri é `Legendary Creature — **Phoenix**`. **Nem o enxame nem o comandante são artefato ou Humano.** A carta que o deck carregava como resposta a board wipe **não protege nada que importe** no eixo novo. Segundo atrito: a ativação **exila a própria carta**, então ela some do board *e* das contagens de artefato/criatura no exato momento em que o deck mais precisa de largura |

### Protocolo de corte (§2 do checklist) — quem cobre cada função

| Função | Quem cobre depois do corte |
|---|---|
| **Proteção em massa** (a dor 4) | **`Unbreakable Formation`** `{2}{W}` — *"Creatures you control gain indestructible"*, **sem restrição de tipo**: alcança as fichas Rebel **e** o Otharri Phoenix. Em fase principal, o addendum ainda põe **contador +1/+1 permanente em cada** e dá **vigilância** (que destrava a recursão do Otharri, §1.2 do 02-theme). Reforço: **`Boros Charm`** `{R}{W}` protege **permanentes**, não só criaturas — cobre também `Valor in Akros`, `The Circle of Loyalty` e os artefatos contra `Austere Command`/`Cleansing Nova`/`Vandalblast`. E `Duty Beyond Death`, que já está no deck, continua |
| **Contagem de artefato** (metalcraft do Jor Kadeen) | **`Glass Casket`** (**caixa**, `{1}{W}`, R$ 0,09 LM) — é **artefato** *e* remoção real. Repõe a contagem **e** fecha um buraco. Soma-se a `Hexgold Halberd`/`Barbed Batterfist` se a Fase 2 os trouxer |
| **Corpo 3/3 CMC 3 com vigilância/vínculo** | Fora da minha especialidade — é slot de criatura temática. `Skyclave Apparition` (§3) repõe corpo **e** remoção no mesmo custo; a decisão final é do `theme-analyst`/orquestrador |
| **Contagem de criatura / legendária (ficha do Circle of Loyalty)** | Coberta com folga pelo motor do comandante e pelas legendárias que ficam (`Jor Kadeen`, `Belladonna Took`, `Éomer`, `Syr Alin`, `Syr Carah`) |

### Veredito

**Corte recomendado, condicionado a duas entradas nomeadas:** `Unbreakable Formation` (ou
`Rootborn Defenses`/`Make a Stand`, substitutas mais baratas) **e** `Glass Casket`. Com as duas,
o corte é limpo no eixo de interação e no eixo de contagem de artefato. **Sem elas, não cortar** —
seria trocar uma proteção em massa *errada* por proteção em massa *nenhuma*, que é pior.

O corpo 3/3 vigilância/vínculo é função **fora da minha especialidade** (§ "Escopo de corte"):
devolvo esse eixo ao orquestrador. Se o deck final ficar curto de criaturas brancas de CMC 3,
a carta pode ficar **como criatura**, com a ativação tratada como texto morto.

---

## 2. As proteções de alvo único × a dor 4

O briefing pede resiliência, não menos criaturas. As três proteções pontuais do deck fazem
**quase o mesmo**:

| Carta | CMC | O que dá | Alcance | Diferencial real |
|---|---|---|---|---|
| `Gods Willing` | 1 | proteção da cor escolhida | 1 criatura | **1 mana**; proteção de cor também **fura bloqueio** → garante o gatilho de ataque do Otharri; scry 1 |
| `Feat of Resistance` | 2 | proteção da cor + **contador +1/+1 permanente** | 1 criatura | o contador **fica** depois do turno |
| `Adamant Will` | 2 | +2/+2 e **indestrutível** | 1 criatura | cobre dano **incolor** e "destroy", que proteção de cor não cobre |

**Leitura correta (e ela é a do §1.4 do re-pass temático, que eu confirmo):** contra **remoção
pontual**, o alvo que importa é **um** — o Otharri. Proteção de alvo único é o slot certo para
isso, e `Gods Willing` a 1 mana é a melhor da mesa. **Contra board wipe, as três são inúteis**:
salvam 1 corpo de 10.

Por isso a arquitetura da dor 4 fica assim, em três camadas — e **nenhuma delas reduz criaturas**:

1. **Camada zero (grátis, já paga):** os contadores de experiência ficam **no jogador**. Wipe,
   exílio, bounce e roubo não os tocam. Com 5 contadores acumulados, **um** ataque pós-wipe
   devolve 5 fichas de uma vez. Isso muda o cálculo: o deck **não precisa temer wipes como um
   deck de tokens normal** — ele é o único na mesa que reconstrói o board inteiro num turno.
2. **Camada de proteção em massa (responde ao wipe):** `Duty Beyond Death` (deck) +
   `Unbreakable Formation` (compra) + `Boros Charm` (compra) — 3 cartas que salvam **tudo**,
   incluindo Rebel e Phoenix. Boros Charm salva até os encantamentos.
3. **Camada de proteção do motor (responde à remoção pontual):** `Gods Willing` +
   `Feat of Resistance`. **`Adamant Will` é a terceira cópia e sai** — a função "indestrutível
   contra dano incolor" passa a ser coberta, em **massa**, por `Unbreakable Formation` pelo
   mesmo custo de mana.

---

## 3. Candidatas — remoção / proteção / counters

Origem: `deck` = já em `lista.txt` (custo zero) · `caixa` = sobressalente (custo zero) ·
`compra`. Regra 3: 2+ pontos de sinergia declarados em cada linha.

### 3.1 Núcleo recomendado

| Carta | CMC | Tipo | Subcategoria | O que responde | Sinergias (2+) | Na coleção? |
|---|---|---|---|---|---|---|
| **Celebrate the Mountain-king** | 4 | Enchantment | remoção múltipla | **1 não-terreno por oponente** (criatura, artefato, ench., PW) | (a) 3-por-1 num pod; (b) recruit gera ficha → gatilho de `Belladonna Took`/`Valor in Akros`; (c) é encantamento, sobrevive a wipe de criatura | **deck** |
| **Glass Casket** | 2 | Artifact | remoção (criatura) | exila criatura MV ≤3 — a faixa de 90% dos mana dorks, hatebears e engines | (a) **artefato** → metalcraft do `Jor Kadeen` + contador do `Luxknight Breacher`; (b) exílio (não volta); (c) repõe a contagem que o corte do Paladin Danse abre | **caixa** (R$ 0,09 LM) |
| **Smite the Deathless** | 2 | Instant R | remoção (criatura) | 3 de dano, **remove indestrutível** e **exila** | (a) a única resposta do pool a criatura indestrutível/com recursão; (b) instant a 2 manas — o deck não tem nenhuma; (c) custo zero | **caixa** |
| **Flame Slash** | 1 | Sorcery R | remoção (criatura) | 4 de dano por **{R}** | (a) mata comandante 4/4 no turno 1–2 por 1 mana; (b) custo zero; (c) libera mana para o Otharri no mesmo turno | **caixa** |
| **Molten Blast** | 3 | Instant R | remoção modal | 2 de dano a criatura/PW **ou** destrói artefato | (a) **flexibilidade** — conta para as duas cotas de cobertura; (b) instant; (c) custo zero | **caixa** |
| **Generous Gift** | 3 | Instant W | remoção **flexível** | **destrói qualquer permanente** (terreno incluso) | (a) a resposta a "qualquer coisa" que RW quase não tem; (b) o 3/3 que ele devolve é irrelevante contra um deck que fabrica 5 corpos por ataque; (c) instant | compra · ~US$ 0,72 |
| **Chaos Warp** | 3 | Instant R | remoção **flexível** | **qualquer permanente** → topo do grimório | (a) única resposta **vermelha** a encantamento; (b) instant; (c) já cotado: R$ 3,81 LM (26 d) | compra · R$ 3,81 LM |
| **Unbreakable Formation** | 3 | Instant W | **proteção em massa** | board wipe, dano em massa, combate | (a) indestrutível a **todas** as criaturas — Rebel **e** Phoenix (a função que o Paladin Danse falha); (b) addendum: **contador +1/+1 permanente em cada** — escala com o enxame; (c) addendum dá **vigilância** → cria Rebel desvirado, que **destrava a recursão do Otharri** (ruling 2023-02-04) | compra · ~US$ 1,40 |
| **Boros Charm** | 2 | Instant RW | proteção em massa + alcance | wipe, `Vandalblast`/`Austere Command`, corrida de vida | (a) protege **permanentes** — salva `Valor in Akros`, `The Circle of Loyalty`, Sol Ring, não só criaturas; (b) 4 de dano ao rosto = alcance para os últimos pontos (dor 1); (c) double strike no Otharri (3/3 voador com vínculo + anthems) | compra · ~US$ 4,55 |
| **Duty Beyond Death** | 2 | Instant W | proteção em massa | board wipe | (a) indestrutível ao time **sem restrição de tipo**; (b) **+1/+1 permanente em cada criatura**; (c) o custo adicional (sacrificar criatura) se paga com **uma ficha 2/2 descartável** | **deck** (R$ 0,43 LM) |
| **Gods Willing** | 1 | Instant W | proteção do motor | remoção pontual no Otharri | (a) 1 mana; (b) proteção de cor **fura bloqueio** → garante o gatilho de ataque; (c) scry 1 | **deck** |
| **Feat of Resistance** | 2 | Instant W | proteção do motor | remoção pontual | (a) contador **permanente** (fica depois do turno); (b) proteção de cor | **deck** |
| **Disenchant** | 2 | Instant W | remoção art/ench | `Rhystic Study`, `Smothering Tithe`, equipamentos, prisões | (a) a mais eficiente das 4 art/ench do deck; (b) instant | **deck** |
| **Seal of Cleansing** | 2 | Enchantment W | remoção art/ench | idem, **pré-paga** | (a) mana adiantado → resposta **grátis** no turno crítico; (b) é permanente no board, não some para wipe de criatura | **deck** |

**Subtotal recomendado: 14 peças** (as duas art/ench e as duas proteções pontuais contam;
`Reduce to Memory` e `Swift Reckoning` continuam no deck como 15ª/16ª e são os primeiros slots
que o `manabase-engineer` pode reclamar para os 2 terrenos faltantes).

### 3.2 Anulação (counters) — a limitação da cor, declarada

**RW não tem counterspell.** Não existe carta de contramágica genérica na identidade; a meta de
"anulação" do briefing só pode ser cumprida por substitutos. Os dois que existem:

| Carta | CMC | O que faz | Por que não recomendo como núcleo |
|---|---|---|---|
| `Rebuff the Wicked` | 1 | contra mágica **que tem como alvo um permanente seu** | funciona (protege o Otharri de `Swords to Plowshares`, `Beast Within`, etc.), mas **~US$ 7,21** por um efeito que `Gods Willing` faz por 1 mana branco e R$ 0 — **reserva**, não núcleo |
| `Grand Abolisher` | 2 | no **seu** turno, oponentes não conjuram nada nem ativam artefato/criatura/encantamento | protege o alfa inteiro. **Regra 11:** trava só o *seu* turno, não impede ninguém de jogar o próprio jogo — é aceitável. Preço a conferir; **reserva** |

**Encaminhamento:** declarar a meta de counters como **coberta por proteção**, não por
contramágica. É a resposta certa dentro da cor, e não custa slot extra.

---

## 4. Candidatas — board wipes

### 4.1 A pergunta que o briefing levantou: um deck que *é* o board mais largo deve rodar wipe?

**Sim — mas quase nunca simétrico.** Três fatos mudam o cálculo:

1. **Os contadores de experiência ficam no jogador.** Depois de qualquer wipe, o primeiro ataque
   do Otharri devolve **N fichas de uma vez**, onde N é o total acumulado. Nenhum oponente
   reconstrói assim. O deck é o que **menos** perde num wipe.
2. **Mas a conta de mana é dura.** Wipe simétrico de 5 manas + Otharri de volta pela zona de
   comando (imposto: **7 manas**) = dois turnos inteiros sem board. A recursão do cemitério
   (`{2}{R}{W}` + virar um Rebel desvirado) **fica bricada** pelo wipe, porque o Rebel morreu
   junto. `Late to Dinner` é o que salva esse cenário.
3. **Existem wipes que simplesmente não matam fichas.** É onde o deck deve gastar os slots.

### 4.2 Recomendados

| Carta | CMC | Simétrico? | O que faz | Sinergias | Na coleção? |
|---|---|---|---|---|---|
| **Hour of Reckoning** | 7 (convoke) | **não** — poupa **todas** as fichas | `Destroy all nontoken creatures` | (a) o enxame inteiro do Otharri **sobrevive**; o board do oponente (majoritariamente não-ficha) some; (b) **convoke**: com 6 fichas custa ~`{W}` real; (c) não toca `Valor in Akros`, `The Circle of Loyalty` nem os artefatos. *Atrito F7 a declarar:* mata os seus **não-fichas** (Otharri, Jor Kadeen, Belladonna Took) — mas os contadores ficam e a recursão funciona, porque **as fichas sobreviveram e fornecem o Rebel desvirado**. Convoke exige fichas **desviradas**: conjurar na fase principal **pré**-combate, ou ter vigilância em massa | compra · ~US$ 0,27 |
| **Winds of Abandon** | 2 / overload 6 | **não** — unilateral puro | spot: exila 1 criatura que você não controla · overload: **exila todas as criaturas que você não controla** | (a) o wipe perfeito: **seu board inteiro fica de pé** e ainda ataca no mesmo turno; (b) **duas cartas em uma** — remoção pontual a 2 manas no early, wipe no late; (c) exílio (comandante não volta pro cemitério). *Custo declarado:* cada oponente busca um básico — é rampa para eles, e o deck aceita porque está correndo | compra · ~US$ 4,84 |
| **Vandalblast** | 1 / overload 5 | **não** — unilateral | spot: destrói 1 artefato que você não controla · overload: **destrói todos os artefatos que você não controla** | (a) resposta em massa a artefatos **sem tocar** Sol Ring, Arcane Signet, `The Circle of Loyalty` e `Glass Casket`; (b) a 1 mana é remoção pontual barata; (c) cobre a cota "2+ respostas a artefato/encantamento" com folga | compra · ~US$ 1,33 |
| **Austere Command** | 6 | **assimétrico por escolha** | escolha **dois**: destrói todos os artefatos · todos os encantamentos · todas as criaturas **MV ≤3** · todas as criaturas **MV ≥4** | (a) o modo **"MV 4 ou maior"** poupa **todas as fichas** (MV 0) e quase todo o board do deck, e mata exatamente o que RW não consegue matar sozinho: os fatties do oponente; (b) o segundo modo cobre artefato **ou** encantamento em massa; (c) é a única carta do pool que responde a uma mesa de encantamentos | compra · ~US$ 0,35 |
| **Blast Zone** | 0 (land) | modular | `{3},{T}, sac`: destrói cada não-terreno com MV igual aos contadores | (a) **entra com 1 contador** — conferido no oracle: ela **nunca pode ser ajustada para MV 0**, então **é fisicamente incapaz de matar as suas fichas**. O risco que se supõe nela não existe neste deck; (b) é **terreno** → cobre parte do −2 da manabase sem gastar slot de feitiço; (c) responde a prisões e a enxames de 1–2 drops | **caixa** (R$ 0,89 LM) |
| **Fumigate** | 5 | **simétrico** | destrói todas as criaturas, ganha 1 vida por criatura | ver §4.3 — **custo zero**, mas é o **último** slot de wipe, não o primeiro | **caixa** (R$ 1,37 LM) |

### 4.3 Fumigate — os dois lados, como pedido

**A favor:**
- Custo **zero** (já na caixa, R$ 1,37 pago).
- Os contadores de experiência sobrevivem → **você é quem menos perde**. Numa mesa em que todo
  mundo reconstrói do zero, você reconstrói com 5 fichas de uma vez.
- Vira **assimétrico** de verdade combinada com `Duty Beyond Death`, `Unbreakable Formation` ou
  `Boros Charm` — você conjura o wipe e responde ao próprio wipe com indestrutível em massa.
  É um `Plague Wind` de 7 manas em duas cartas, e as três proteções já estariam no deck por
  outro motivo.
- Ganha 1 vida por criatura destruída — num wipe de 15 corpos, são 15 de vida numa mesa que já
  te odeia por ser o board mais largo.

**Contra:**
- O deck **é** o board mais largo da mesa. Na maioria das mesas o Fumigate é uma carta que você
  segura e nunca joga, porque o estado de jogo em que ele é bom (você atrás, mesa à frente) é
  exatamente o estado que o Otharri resolve sozinho atacando.
- Ele **brica a recursão barata** do Otharri: sem Rebel desvirado sobrevivente, o comandante só
  volta pela zona de comando, com imposto.
- 5 manas de sorcery que **não adicionam corpo** num deck cuja tese é largura de board.
- Já existem **quatro** wipes melhores para este deck que não têm nenhum desses problemas.

**Veredito:** **incluir como o 4º wipe, e só se a régua de R$ 200 apertar as compras.** Prioridade
de slot: `Hour of Reckoning` > `Winds of Abandon` > `Vandalblast` > `Austere Command` > `Fumigate`.
Se todas as compras passarem, Fumigate fica de fora — não por ser ruim, mas por ser a única das
cinco que pode virar contra você.

### 4.4 Wipes avaliados e **dispensados**, com o motivo

| Carta | Por que não |
|---|---|
| `Ratchet Bomb` (**caixa**) | Começa com **0 contadores** — o estado padrão dela mata exatamente as suas fichas (MV 0). Diferente da `Blast Zone`, o risco aqui é real. Continua **jogável** (você escolhe quando sacrificar e pode subir os contadores), mas é lenta e exige disciplina. **Reserva**, não núcleo |
| `Retribution of the Meek` (US$ 8,90) | "Poder 4 ou maior" parece poupar fichas 2/2 — **não poupa**: com `Jor Kadeen` em campo as fichas são **5/2**, e com `The Circle of Loyalty` + qualquer segundo anthem passam de 4. Anti-sinergia com o próprio plano, e cara |
| `The Battle of Bywater` (US$ 3,06) | Mesmo problema com corte ainda mais baixo (poder 3): `The Circle of Loyalty` sozinha já mata o enxame todo. O upside (um Food por criatura sua, que dispara `Belladonna Took`) não paga o risco |
| `Split Up` (US$ 2,56) | O modo "destrói todas as desviradas" no seu pós-combate é quase unilateral (seus atacantes estão virados). Mas é **anti-sinérgico com a vigilância em massa** que o `theme-analyst` propõe (`Radiant Destiny`, `Intangible Virtue`), que deixaria o seu próprio board desvirado. Conflito de plano — **reserva**, e só se a vigilância não entrar |
| `Cleansing Nova` (US$ 0,29) | Modal e barata, mas o modo de criatura é **simétrico** (mesmo problema do Fumigate) e o modo art/ench é pior que `Austere Command`. **Reserva de orçamento** |
| `Blasphemous Act` | 13 de dano a **cada** criatura — mata o seu board inteiro e fica mais barata justamente quando você tem mais corpos. Anti-sinergia direta |
| `Fell the Mighty` (US$ 0,28) | Destrói criaturas com poder maior que o de uma criatura **alvo sua** — é boa, mas exige um alvo pequeno seu sobrevivendo e é sorcery de 5. Perde para `Austere Command` no mesmo slot |
| `Settle the Wreckage` | Exila os atacantes de **um** jogador — defensiva, e o deck é o agressor da mesa. Baixa aplicação |
| `Sunfall`, `Comeuppance` | Simétricos/defensivos e sem cotação favorável; nenhum ganho sobre os cinco acima |

---

## 5. Cobertura de ameaças (com o núcleo recomendado)

- **Criaturas — pontual:** `Glass Casket` (exílio MV≤3) · `Smite the Deathless` (instant, fura
  indestrutível, exila) · `Flame Slash` (4 de dano por {R}) · `Molten Blast` · `Swift Reckoning`
  (condicional) · `Winds of Abandon` (modo spot) — **6**, sendo **3 em velocidade de instantâneo**
  (hoje são **0**).
- **Criaturas — massa:** `Hour of Reckoning` (poupa fichas) · `Winds of Abandon` overload
  (unilateral) · `Austere Command` modo MV≥4 (poupa fichas) · `Fumigate` (simétrico, reserva) —
  **3 + 1**.
- **Artefatos / Encantamentos:** `Disenchant` · `Seal of Cleansing` · `Molten Blast` (modo
  artefato) · `Vandalblast` (**massa, unilateral**) · `Austere Command` (**massa**) ·
  `Celebrate the Mountain-king` · `Generous Gift` · `Chaos Warp` — **8**, com **2 respostas em
  massa**. Cota do briefing (2+) cumprida com folga.
- **Flexível ("qualquer permanente"):** `Generous Gift` · `Chaos Warp` ·
  `Celebrate the Mountain-king` · `Reduce to Memory` · `Blast Zone` — **5**. Cota (1–2) cumprida.
- **Combos / Spells:** **cobertura indireta.** RW não tem contramágica; a resposta do deck é
  (i) matar a peça-chave com remoção flexível em **instant** (`Generous Gift`, `Chaos Warp`),
  (ii) `Vandalblast` overload contra combos de artefato, (iii) **pressão** — o Otharri mata a mesa
  em 4 ataques, o que encurta a janela do combo. `Grand Abolisher` e `Rebuff the Wicked` ficam
  como reservas se o meta do usuário for pesado em interação no turno dele.
- **Proteção do motor (remoção pontual no Otharri):** `Gods Willing` · `Feat of Resistance` — **2**.
- **Proteção em massa (board wipe):** `Duty Beyond Death` · `Unbreakable Formation` ·
  `Boros Charm` — **3**, todas alcançando **Rebel e Phoenix** (hoje: **1**).

**Meta:** ~10 interação + 2–4 wipes → **proposta: 14 interação + 4 wipes (3 unilaterais/assimétricos
+ 1 simétrico opcional) + Blast Zone no slot de terreno.**

---

## 6. Cortes propostos — ficha F1–F7 (regra 4)

Consulta ao `decisions.md` (regra 5): **nenhuma** das cartas abaixo tem histórico de corte — o
deck foi montado fora do pipeline e esta é a primeira rodada. Nenhuma entrada proposta é
reposição de corte anterior.

### I1 · `Paladin Danse, Steel Maverick` — **corte condicionado** → ficha completa no §1

Sai: Paladin Danse — funções: [proteção em massa, corpo 3/3 vigilância+vínculo CMC 3, contagem de
artefato, contagem de criatura/legendária] → proteção em massa **coberta e ampliada** por
`Unbreakable Formation` + `Boros Charm`; contagem de artefato coberta por `Glass Casket`;
corpo → **devolvido ao orquestrador** (fora da minha especialidade); contagem de legendária
coberta pelas legendárias que ficam. **Condição: só sai se `Unbreakable Formation` e `Glass Casket`
entrarem.**

### I2 · `Adamant Will` — `{1}{W}` Instant — **corte limpo**
- **F1** `Target creature gets +2/+2 and gains indestructible until end of turn.`
- **F2** Sem corpo. **F3** Instantâneo — o deck não tem prowess, magecraft nem contagem de
  feitiços. **F4** Não recebe nada. **F5** Dá +2/+2 e indestrutível a **uma** criatura.
- **F6** Turno 2 — barata e disponível cedo.
- **F7** É a **terceira** carta funcionalmente idêntica (`Gods Willing`, `Feat of Resistance`), e
  a pior das três: `Gods Willing` faz por 1 mana e ainda **fura bloqueio**; `Feat of Resistance`
  deixa **contador permanente**. O único diferencial real dela — indestrutível cobre dano
  **incolor** e efeitos de "destroy" que proteção de cor não cobre — é **exatamente** o texto do
  `Unbreakable Formation`, que faz isso para **o time inteiro** por 1 mana a mais.
- **Protocolo:** funções = [proteção de alvo único, buff temporário +2/+2, cobertura de dano
  incolor]. Proteção → coberta por `Gods Willing` e `Feat of Resistance`. Buff → coberto e
  superado por `Pride of Conquerors`/`Basri's Solidarity` (em massa). Dano incolor → **coberto e
  ampliado** por `Unbreakable Formation` e `Duty Beyond Death`. **Nenhuma função descoberta.**

### I3 · `Invoke the Divine` — `{2}{W}` Instant — **corte limpo**
- **F1** `Destroy target artifact or enchantment. You gain 4 life.`
- **F2–F4** Sem corpo, tipo sem uso, não recebe nada. **F5** Não dá nada.
- **F6** Turno 3. **F7** `Disenchant` (no mesmo deck) faz a parte relevante por **1 mana a menos**.
  As 4 vidas são irrelevantes num deck com vínculo no comandante e num plano de corrida.
- **Protocolo:** funções = [remoção art/ench, ganho de vida]. Remoção → coberta por `Disenchant`,
  `Seal of Cleansing`, `Molten Blast`, `Vandalblast`, `Austere Command`, `Generous Gift`,
  `Chaos Warp`. Ganho de vida → coberto pelo vínculo do Otharri e pelo `Fumigate`/`Dawnstrike
  Vanguard`. **Nenhuma função descoberta.**

### I4 · `Expose to Daylight` — `{2}{W}` Instant — **corte limpo**
- **F1** `Destroy target artifact or enchantment. Scry 1.`
- **F2–F5** Idem I3, com scry 1 no lugar das 4 vidas.
- **F6** Turno 3. **F7** Mesma redundância: é a **quarta** carta art/ench-only de um deck que
  precisa de resposta a **criatura**. Quatro cartas cobrindo um tipo de ameaça e zero cobrindo
  outro é má distribuição, não profundidade.
- **Protocolo:** funções = [remoção art/ench, scry 1]. Remoção → mesma lista de I3. Scry →
  coberto por `Gods Willing` (que fica) e **declarado dispensável**. **Nenhuma função descoberta.**

### I5 · `Djeru's Renunciation` — `{1}{W}` Instant — **corte condicionado (escopo de draw)**
- **F1** `Tap up to two target creatures.` / `Cycling {W}`.
- **F2–F4** Sem corpo, tipo sem uso, não recebe nada.
- **F5** Tapa 2 criaturas — **não remove nada**; é Falter parcial/fog parcial.
- **F6** Turno 2, ou **turno 1 como cycling**.
- **F7** Num deck que resolve bloqueio por **volume** (6–10 corpos) e que vai ganhar trample
  (`Crash Through`, `Tori D'Avenant`) ou menace (`Iroas`), tapar dois bloqueadores é o efeito
  menos relevante do pacote. Defensivamente, tapa dois atacantes — mas o deck tem vigilância e
  corpos de sobra para bloquear.
- **Protocolo:** funções = [tapar 2 criaturas, **cycling `{W}` = ciclagem/filtro de mão**].
  Tapar → coberto e superado por remoção de verdade (`Glass Casket`, `Smite the Deathless`) e pela
  evasão em massa. **Cycling → função de draw, fora da minha especialidade.** **Não corto
  sozinho:** devolvo ao `draw-specialist`/orquestrador. Se o pacote de draw entregar 12–13 fontes
  reais, o cycling aqui é redundante e o corte fica limpo.

### I6 · `Gideon's Triumph` — já cortada em **K17** pelo `theme-analyst` — **confirmo**
Função única = [remoção condicional de baixa qualidade: edict em que o **oponente** escolhe, e
metade do texto lê um planeswalker Gideon que o deck não tem]. **Coberta e superada** por
`Glass Casket` (caixa, exila de verdade, e é artefato) e por `Smite the Deathless`/`Flame Slash`.
Nenhuma função descoberta.

### Slots que **não** corto, mas sinalizo como os mais fracos da categoria
`Swift Reckoning` (só pega criatura **virada**) e `Reduce to Memory` (sorcery, devolve um 3/2)
continuam sendo interação real. São os dois primeiros candidatos se o `manabase-engineer`
precisar de espaço para os 2 terrenos faltantes — **mas o corte é dele, com ficha**, não meu.

---

## 7. Reservas / swaps propostos

### 7.1 Tabela de swaps (6 saídas nominais → 6 entradas)

| Sai | Entra | Justificativa em uma linha |
|---|---|---|
| `Paladin Danse, Steel Maverick` | **Unbreakable Formation** (compra) | proteção em massa que **alcança Rebel e Phoenix**, + contadores permanentes + vigilância que destrava a recursão |
| `Adamant Will` | **Glass Casket** (**caixa**, R$ 0,09) | troca a 3ª proteção pontual redundante por **remoção real + artefato** (repõe o metalcraft do Jor Kadeen) |
| `Invoke the Divine` | **Smite the Deathless** (**caixa**) | troca a 4ª carta art/ench-only pela **primeira remoção de criatura em instant** do deck |
| `Expose to Daylight` | **Generous Gift** (compra) | troca resposta estreita por **"destrói qualquer permanente"** em instant |
| `Gideon's Triumph` | **Chaos Warp** (compra, R$ 3,81 LM) | troca edict ruim pela única resposta **vermelha** a encantamento, em instant |
| `Djeru's Renunciation` *(condicionado ao draw)* | **Boros Charm** (compra) | troca "tapa 2" por proteção de **permanentes** + 4 de alcance + double strike no comandante |

**Wipes entram em slots novos** (o deck tem **zero**): `Hour of Reckoning`, `Winds of Abandon`,
`Vandalblast` (+ `Austere Command` ou `Fumigate` como o 4º). Os slots vêm dos cortes temáticos
já aprovados na Fase 2 (K1, K2, K7–K17), não da interação.
`Blast Zone` entra no **slot de terreno** (`manabase-engineer`), cobrindo parte do −2.

### 7.2 Reservas, por ordem de prioridade

| # | Carta | Origem | Quando entra |
|---|---|---|---|
| R1 | **Rootborn Defenses** (~US$ 0,15) | compra | substituta barata do `Boros Charm`: indestrutível em massa **+ populate** (copia uma ficha 2/2 Rebel → dispara `Belladonna Took`/`Valor in Akros`/`Warleader's Call`). 3 pontos de sinergia por US$ 0,15 |
| R2 | **Skyclave Apparition** (~US$ 0,32) | compra | remoção **com corpo**: exila não-terreno não-ficha MV ≤4, é criatura que recebe os anthems e **dispara `Valor in Akros`/`Warleader's Call` ao entrar**. Repõe o corpo do Paladin Danse |
| R3 | **Austere Command** (~US$ 0,35) | compra | 4º wipe, se o orçamento couber — modo MV≥4 poupa **todas** as fichas |
| R4 | **Return to Dust** (~US$ 0,33) | compra | exila **até 2** artefatos/encantamentos — substitui `Disenchant` se a mesa for pesada em permanentes |
| R5 | **Flame Slash** / **Molten Blast** / **Seismic Wave** | **caixa** | remoção vermelha de custo zero; `Seismic Wave` é mini-wipe unilateral (1 de dano em cada criatura não-artefato **de um** oponente) contra mesas de tokens |
| R6 | **Fumigate** | **caixa** | 4º wipe se as compras não couberem na régua (§4.3) |
| R7 | **Anointed Peacekeeper** (~US$ 0,35) | compra | 3/3 vigilância que **taxa {2}** toda mágica/habilidade que mira você ou seus permanentes — proteção preventiva do Otharri **com corpo** |
| R8 | **Grand Abolisher** · **Rebuff the Wicked** | compra | o substituto de counter, se o meta do usuário for pesado em interação no turno dele. `Rebuff` a ~US$ 7,21 é caro para o efeito |
| R9 | **Ratchet Bomb** · **Lux Cannon** | **caixa** | `Ratchet Bomb` é usável (nunca sacrifique com 0 contadores); `Lux Cannon` destrói **qualquer permanente** repetidamente e é artefato, mas leva 3 turnos para carregar |
| R10 | **Swords to Plowshares** (R$ 13,99 LM) / **Path to Exile** (~US$ 1,00) | compra | upgrade de qualidade da remoção pontual. StP é a melhor do jogo, mas a **R$ 13,99 consome 7% do teto** — só se sobrar orçamento |

### 7.3 Dispensas da caixa com justificativa escrita (regra 7)

| Carta da caixa | Por que não cobre a função |
|---|---|
| `Joust` | Fight: a criatura **sua** leva dano igual ao poder da outra. Com fichas 2/2 é suicídio, e o bônus `+2/+1 se for Knight` não alcança as fichas (são **Rebel**). Reprova como remoção confiável |
| `Bombard`, `Searing Barrage`, `Stonefury`, `Fateful End`, `Radiating Lightning` | Queima de criatura a **3–5 manas**. A caixa já tem `Flame Slash` (1), `Magma Spray` (1) e `Smite the Deathless` (2) fazendo o mesmo por menos; num deck que quer o Otharri no T4–T5, dano a 4–5 manas compete com o próprio plano |
| `Chandra's Pyrohelix`, `Twin Bolt`, `Punishing Fire`, `Seal of Fire`, `Pinecone Strike` | 2–3 de dano dividido. Matam X/1 e X/2; em Commander a maioria das ameaças relevantes está acima disso. `Pinecone Strike` é a melhor delas (3 de dano + exílio **ou** destrói ficha de artefato) — sobe para **reserva**, as outras ficam de fora |
| `Skullcrack`, `Mercadia's Downfall`, `Seize Opportunity`, `Ambitious Assault`, `Smaug's Fury`, `Smashing Success`, `Destructive Tampering`, `Lightning Volley`, `Moment of Glory` | Não são interação: são queima ao **jogador**, pump ou destruição de terreno. `Smashing Success`/`Destructive Tampering` destroem artefato, mas a 3–4 manas perdem para `Molten Blast` (3, modal) e `Vandalblast` (1) |
| `Steel Wrecking Ball` | 5 manas por 5 de dano a uma criatura. `Flame Slash` faz 4 por 1 |
| `Thaumaton Torpedo` | Destrói não-terreno por **{6}** de ativação (ou {3} se atacou com Spacecraft — o deck não tem nenhuma). Inviável |
| `Leonin Abunas` | Dá hexproof aos **artefatos**. O deck reconstruído terá 4–6 artefatos e nenhum deles é o alvo que o oponente quer remover — o alvo é o **Otharri**, que não é artefato. 1 ponto de sinergia, reprova na regra 3 |
| `Phyrexian Revoker` | Trava a habilidade ativada de **um nome** escolhido. É stax pontual (aceitável na regra 11), soma artefato para metalcraft, mas responde a uma fatia muito estreita de ameaças; `Glass Casket` ocupa melhor o mesmo slot de artefato+interação |
| `Negate`, `Stoic Rebuttal`, `Disruption Protocol` | **Azuis — fora da identidade RW.** Não elegíveis |
| `Return to Nature`, `Smell Fear`, `Ancient Animus`, `Courage in Crisis`, `Road // Ruin` | **Verdes — fora da identidade** (`Road // Ruin` tem face verde, então a identidade de cor inclui G) |

---

## 8. Impacto no orçamento — **estimativa, não régua**

> **Regra 2:** os números abaixo são **estimativa (Scryfall)** e **não decidem** se algo cabe no
> teto de R$ 200. A régua é o **menor valor da LigaMagic**, e a captura é manual (navegador),
> a ser feita pelo orquestrador quando a lista de entradas fechar. Nenhuma das compras abaixo
> tem cotação LigaMagic registrada — só `Chaos Warp` (R$ 3,81, 26 dias) e as da caixa.

| Pacote | Cartas | Estimativa (Scryfall) |
|---|---|---|
| **Custo zero (caixa)** | Glass Casket · Smite the Deathless · Flame Slash · Molten Blast · Blast Zone · Fumigate | **R$ 0,00** (já pagas — R$ 0,09 + 0,89 + 1,37 registradas) |
| **Compras — núcleo econômico** | Unbreakable Formation · Rootborn Defenses · Hour of Reckoning · Vandalblast · Generous Gift · Austere Command · Skyclave Apparition · Return to Dust | ~**US$ 5,17** ≈ R$ 28 *(estimativa)* |
| **Compras — upgrades** | Boros Charm (~US$ 4,55) · Winds of Abandon (~US$ 4,84) · Chaos Warp (R$ 3,81 LM) | ~**US$ 9,39 + R$ 3,81** ≈ R$ 55 *(estimativa)* |

**Alerta de régua:** `Boros Charm` e `Winds of Abandon` são, sozinhas, ~70% do custo estimado da
fase. Se a cotação real da LigaMagic confirmar, os substitutos já estão nomeados e testados:
`Rootborn Defenses` (US$ 0,15) pelo Boros Charm e `Austere Command` (US$ 0,35) pelo Winds of
Abandon. A cobertura de ameaças da §5 **não muda** com a troca — muda a qualidade, não a função.

---

# Devolução · 2026-09-20 — `Requisition Raid` e o slot que ela ocupa

> **Pedido do usuário:** ele adquiriu `Requisition Raid` (coleção desde 2026-09-20, custo zero) e,
> diante da troca proposta `Basri's Solidarity` → `Requisition Raid`, respondeu **"Não podemos ter
> as duas no deck?"**. Portanto `Basri's Solidarity` **fica** e a Raid precisa de **slot novo**.
> `Sol Ring`, `Arcane Signet`, `Command Tower`, `Basri's Solidarity` e qualquer carta de ramp/mana
> (inclusive `Intangible Virtue`) estão **fora do meu alcance de corte** nesta devolução.
>
> Todo oracle e todo ruling desta seção foram puxados de `bin/mtgdb` **nesta sessão** (regra 6).

## D1 · Ficha F1–F7 — `Requisition Raid` (entra)

`{W}` · Sorcery · **MV 1** · identidade W · **caixa (R$ 0,00)**

```
Spree (Choose one or more additional costs.)
+ {1} — Destroy target artifact.
+ {1} — Destroy target enchantment.
+ {1} — Put a +1/+1 counter on each creature target player controls.
```

| Eixo | Leitura |
|---|---|
| **F1** | Três modos de Spree, **pelo menos um obrigatório**. Custos reais: 1 modo = `{1}{W}` (2 manas) · artefato **+** encantamento = `{2}{W}` (3 manas) · os três = `{3}{W}` (4 manas). **Ruling 2024-04-12 conferido:** *"The mana value of a spell with spree is determined only by its mana cost"* → a carta é **MV 1** independente dos modos. Outro ruling relevante: *"If a mode requires a target, you can select that mode only if there's a legal target available"* — o terceiro modo mira **um jogador**, e jogador sempre existe, então **a carta nunca fica sem modo legal** |
| **F2** | Sem corpo. Não tapa, não bloqueia, não é sacrificável |
| **F3** | Sorcery. O deck **não** tem payoff de contagem de mágicas — `Syr Carah, the Bold` só dispara com mágica que causa **dano**, e a Raid não causa. Tipo sem uso aqui: **declarado** |
| **F4** | Não recebe nada (é mágica) |
| **F5** | **É o eixo forte.** O 3º modo põe **+1/+1 permanente** em cada criatura de um jogador alvo — mirando você, é `Basri's Solidarity` pelo **mesmo custo** (`{1}{W}`). Os contadores ficam depois do turno, ao contrário de `Pride of Conquerors` |
| **F6** | Turno 2 em modo único; turno 3 como 2-por-1; turno 4 como 3-por-1. **Escala com o mana disponível** — barata cedo, grande tarde. Não tem turno em que seja cara demais nem pequena demais |
| **F7** | **Feitiço**: não segura combate, não responde a `Rhystic Study` ativado nem a equipar em resposta — `Disenchant` continua sendo necessária por isso. Disputa o slot de 2 manas com `Glass Casket` e `Seal of Cleansing`. **Sobreposição declarada:** em modo único de contador ela é um **superconjunto funcional** de `Basri's Solidarity` (mesmo custo, mesmo efeito, com opções a mais) — o usuário quis as duas, e a redundância de um efeito bom num deck que fabrica 1→3→6→10 corpos é defensável, mas é redundância, não é ganho novo |

### Sinergias (regra 3 — 2+ exigidos; contei 5)

1. **Nunca é carta morta.** O motivo pelo qual `Invoke the Divine` e `Expose to Daylight` foram
   cortadas (§6, I3/I4) é serem **art/ench-only**: contra uma mesa só de criaturas viram papel.
   A Raid tem o 3º modo como piso — sempre há um jogador alvo legal (ruling acima). Ela resolve
   a *causa* do corte daquelas duas, não repete o defeito delas.
2. **Dois permanentes num card por `{2}{W}`.** `Disenchant`, `Seal of Cleansing` e `Molten Blast`
   escolhem **um**. Num deck com 13 fontes de saque mas **zero** recursão de mágica, o 2-por-1
   contra dupla de prisões (`Ghostly Prison` + `Rhystic Study`, equipamento + estática) é economia
   de carta que o deck não tem de outra forma.
3. **Contadores permanentes no enxame do Otharri.** As fichas entram 2/2 viradas e atacando, e
   **continuam existindo** — contador nelas é dano que se repete em cada ataque, ao contrário de
   `Pride of Conquerors`, que expira.
4. **Destrava Ferocious sem depender do `Jor Kadeen`.** `Barrage of Boulders` (rota R3) exige
   criatura com **poder 4+** para o *"creatures can't block this turn"*. Ficha 2/2 + contador da
   Raid + `Intangible Virtue` = **4/4**. Hoje esse gatilho depende de o metalcraft do Jor Kadeen
   estar ligado.
5. **Resiliência parcial a wipe de dano** (correção à sua ressalva, §D4).

### Veredito: **entra**

**Confirmo o seu enquadramento em 2 de 3 pontos e corrijo o terceiro:**

- ✅ *"Por `{2}{W}` mata artefato **e** encantamento no mesmo card"* — confere, e nenhuma outra
  carta do deck faz isso.
- ✅ *"É feitiço, não segura combate"* — confere, e é a razão de `Disenchant` **não** sair.
- ⚠️ *"Contadores morrem junto com as criaturas, então não responde à dor 4"* — **parcialmente
  falso.** Contra wipe de **destruição** (`Wrath of God`, `Damnation`, `Austere Command`), sim,
  o contador é irrelevante. Contra wipe de **dano**, não: ficha 2/2 vira **3/3** e sobrevive a
  `Pyroclasm`, `Whipflare`, `Fiery Cannonade` e `Bane of Progress`-lite — a faixa de 2 de dano
  é a mais comum em mesas de orçamento. É resiliência pequena, mas é **real** e é mais do que
  `Basri's Solidarity` oferecia sozinha. Com as duas no deck, o enxame passa a 4/4 e escapa
  também de `Anger of the Gods`.

---

## D2 · O slot — **sai `Fumigate`**

Consulta obrigatória ao `decisions.md` (regra 5): o registro tem **duas** linhas, ambas sobre o
comandante. `Fumigate` **não tem histórico de corte**, e `Requisition Raid` nunca esteve no deck
(conferido em `lista.txt`). Nenhuma das duas é reposição de corte anterior.

`Fumigate` `{3}{W}{W}` · Sorcery · CMC 5 · **caixa (R$ 1,37 já pagos)** — `Destroy all creatures.
You gain 1 life for each creature destroyed this way.`

### Ficha F1–F7 completa (regra 4)

| Eixo | Leitura |
|---|---|
| **F1** | Uma linha: destrói **todas** as criaturas — as suas inclusive — e devolve 1 vida por criatura destruída (as suas contam) |
| **F2** | Sem corpo. Não tapa para nada, não bloqueia, não é sacrificável. Não deixa nada no campo |
| **F3** | Sorcery. **Não** alimenta o metalcraft do `Jor Kadeen` (19 artefatos), não é permanente, não entra em nenhuma contagem do deck. `Syr Carah` exige dano — Fumigate não causa dano |
| **F4** | Não recebe nada: nem anthem, nem contador, nem gatilho do Otharri. Não é alvo de `Late to Dinner` (que devolve **criatura** do cemitério) |
| **F5** | Não concede nada a outra carta. O ganho de vida é só para você, e **vida não é recurso gasto neste deck** — não há custo Phyrexiano, não há pagamento de vida além de 1 ponto ocasional de `Battlefield Forge` |
| **F6** | Turno 5 com a curva atual — **exatamente o turno do Otharri** (CMC 5). Os dois disputam o mesmo turno, e o que Fumigate destrava a partir dali é **um board vazio dos dois lados**, sem pressão |
| **F7** | **Quatro atritos, todos verificados no oracle:** (a) o deck **é** o board mais largo da mesa — Fumigate mata 6–10 fichas suas, `Hero of Bladehold`, `Neyali`, `Jor Kadeen` e **o próprio Otharri**, cobrando imposto de comandante; (b) **brica a recursão barata do Otharri** — `{2}{R}{W}, Tap an untapped Rebel you control` exige um Rebel vivo, e o seu wipe matou todos, então a única volta é a zona de comando a 7+; (c) anti-sinergia com **todas** as seis estáticas que só produzem com board (`Warleader's Call`, `Radiant Destiny`, `Intangible Virtue`, `Valor in Akros`, `The Circle of Loyalty`, metalcraft do `Jor Kadeen`); (d) **escala ao contrário** — quanto melhor o seu jogo, pior a carta |

### Protocolo de corte (§2 do checklist) — quem cobre cada função

| Função de `Fumigate` | Quem cobre depois do corte |
|---|---|
| **Wipe de criaturas em massa** | `Hour of Reckoning` — *"Destroy all **nontoken** creatures"*, com **convoke**: suas 6–10 fichas ficam de pé, pagam o custo e **atacam no mesmo turno**. É estritamente melhor aqui. Reforçado por `Blast Zone` (modular, por MV), `Barrage of Boulders` (1 de dano em cada criatura **que você não controla** — mini-wipe unilateral) e `Celebrate the Mountain-king` (exila 1 não-terreno **por oponente**) |
| **Resposta a criatura grande / recorrente** | `Smite the Deathless` (instant, **remove indestrutível** e exila) · `Flame Slash` (4 de dano por `{R}`) · `Glass Casket` (exila MV≤3) · `Molten Blast` · `Palace Jailer` · `Hour of Reckoning` |
| **Ganho de vida em massa** | Lifelink do próprio Otharri (3+ por ataque, mais com anthem) e `Dawnstrike Vanguard`. **Declarado dispensável:** o deck corre, não estabiliza — 15 vidas num turno em que você ficou sem board não compram o turno seguinte |
| **Resposta a enxame de FICHAS do oponente** | **Fica descoberta.** `Hour of Reckoning` poupa fichas **dos dois lados**; `Barrage of Boulders` só mata X/1. **Custo aceito e declarado:** contra outro deck go-wide, o plano correto deste deck é **correr** — o motor do Otharri é quadrático (1→3→6→10), `Warleader's Call` pinga cada oponente a cada ficha que entra, `Goblin War Drums` dá menace e `Barrage of Boulders` tira os bloqueios. Trocar 1-por-1 num wipe que apaga o **maior** board da mesa (o seu) é perder a troca |
| **Combo "wipe unilateral"** (Fumigate + `Duty Beyond Death`/`Unbreakable Formation`) | **Perdido.** Custo aceito: exige **7–8 manas no mesmo turno** e as **duas** cartas certas na mão. As proteções continuam no deck cumprindo a função principal delas — responder ao wipe **alheio**, que é a dor 4 |

### Simetria de critério (§3 do checklist)

Avaliei as duas cartas **sob o mesmo estado de jogo** que usei para defender a entrada: 6–10 fichas
Rebel, 1–2 anthems estáticos, Otharri vivo.

- Nesse estado, `Requisition Raid` põe **+1/+1 permanente em 6–10 corpos** e ainda mata dois
  permanentes por 4 manas.
- Nesse **mesmo** estado, `Fumigate` apaga exatamente esse board — é uma carta de valor **negativo**.
- No estado inverso (board seu vazio, mesa à frente), `Fumigate` é boa e `Raid` é pequena — mas
  **não morta** (mata artefato ou encantamento por `{1}{W}`).

A assimetria decide: a pior hora da Raid é "carta pequena"; a pior hora do Fumigate é "carta que
joga contra você". E, como o §4.3 desta mesma página já havia registrado, **Fumigate sempre foi o
5º de 5 wipes** — entrou por régua de orçamento, não por mérito. O orçamento não mudou; o que
mudou é que agora existe um uso melhor para o slot.

### Escopo (§ "Escopo de corte")

`Fumigate` é **board wipe puro** — função inteiramente dentro da minha especialidade, sem corpo,
sem tipo que alimente contagem, sem receber nada. **Não há função fora do meu alcance**, então o
corte é meu e é **limpo**, com a única lacuna (enxame de fichas alheio) declarada acima.

### Cartas que considerei e **não** cortei, com o motivo

| Carta | Por que não é o slot |
|---|---|
| `Molten Blast` (caixa) | Parecia o corte óbvio pela sobreposição de artefato com a Raid. **A ficha barrou:** é uma das **duas** remoções de criatura em **instant** do deck (a outra é `Smite the Deathless`). O "buraco nº 1" diagnosticado no §0 desta página era *zero remoção de criatura em instant*; cortá-la regride de 2 para 1 o número que a rodada inteira existiu para consertar |
| `Pride of Conquerors` | Parecia o pump temporário redundante com `Basri's Solidarity` + Raid. **A ficha barrou:** tem **Ascend**, e o deck estoura 10 permanentes com facilidade → é **+2/+2 em todo o time, em instant**. Com 10 fichas são **20 de dano extra** no turno do alfa, a preço de 2 manas. Não é o mesmo efeito das outras duas: é finalizador e truque de combate |
| `Seal of Cleansing` | O sacrifício é em **velocidade de instantâneo** e o mana já está **pré-pago** — é a resposta que você tem com o mana todo gasto no Otharri. A Raid, sendo feitiço, não cobre essa janela |
| `Disenchant` | Única resposta art/ench **em instant** do deck depois que `Chaos Warp` e `Generous Gift` saíram por preço. Ela é justamente o que a Raid **não** faz |
| `Basri Ket` | O `+1` é **proteção de alvo único recorrente e grátis, todo turno**, no Otharri — vale mais que qualquer proteção de uso único; o `-2` faz corpo; o emblema é rota de vitória. Fora disso, é o único planeswalker e absorve ataques |
| `Barrage of Boulders` (R$ 0,05) | É **mini-wipe unilateral** *e* peça da rota R3 (*"creatures can't block"*). Duas funções por 5 centavos |

---

## D3 · Efeito líquido na contagem

| Métrica | Antes (v1 fechada) | Depois da troca |
|---|---|---|
| Peças de interação eficazes | 14 | **14** (Raid entra, Fumigate sai) |
| Remoção em **massa** | 3 — `Hour of Reckoning`, `Fumigate`, `Blast Zone` (+ `Barrage of Boulders`) | **3** — `Hour of Reckoning`, `Blast Zone`, `Barrage of Boulders` |
| Wipes **simétricos** (que matam o seu board) | **1** | **0** |
| Cartas que respondem a artefato **e** encantamento no mesmo card | 0 | **1** |
| Cartas de interação que podem ser **mortas na mão** contra mesa sem art/ench | 2 (`Disenchant`, `Seal of Cleansing`) | 2 (a Raid **não** entra nessa conta — sempre tem modo legal) |
| Anthems permanentes (contadores) | 1 (`Basri's Solidarity`) | **2** |

**Meta da fase: ~10 interação + 2–4 wipes.** Depois da troca: **14 interação + 3 efeitos de
remoção em massa, todos assimétricos ou unilaterais**. As duas metas continuam cumpridas, e a
qualidade dos wipes **sobe** — 100% deles poupam o seu board, contra 67% antes.

### Cobertura de ameaças — recalculada

- **Criaturas (pontual):** `Glass Casket` · `Smite the Deathless` (instant) · `Flame Slash` ·
  `Molten Blast` (instant) · `Palace Jailer` · `Celebrate the Mountain-king` — **6**, sendo **2 em
  instant**. *Inalterado pela troca.*
- **Criaturas (massa):** `Hour of Reckoning` (poupa fichas) · `Barrage of Boulders` (unilateral) ·
  `Blast Zone` (modular) — **3, todos assimétricos**. *Buraco declarado:* nenhum deles mata um
  **enxame de fichas** do oponente.
- **Artefatos / Encantamentos:** `Disenchant` (instant) · `Seal of Cleansing` (pré-pago, instant) ·
  **`Requisition Raid` (os dois no mesmo card)** · `Molten Blast` (artefato) ·
  `Celebrate the Mountain-king` · `Blast Zone` — **6**, cota (2+) cumprida com folga. **Melhora.**
- **Flexível ("qualquer permanente"):** `Celebrate the Mountain-king` · `Blast Zone` — **2, no
  piso da cota (1–2), e nenhuma em instant.** É o buraco que sobrou de `Generous Gift` (R$ 5,00)
  e `Chaos Warp` (R$ 3,81) terem saído por preço no fechamento da v1.
- **Proteção do motor:** `Gods Willing` · `Feat of Resistance` · `Basri Ket` `+1` — **3**.
- **Proteção em massa (dor 4):** `Duty Beyond Death` · `Unbreakable Formation` — **2** (`Boros
  Charm` ficou fora por preço). Piso aceitável, mas é o segundo ponto mais magro.
- **Combo / counters:** **zero direto** — limitação de RW já declarada no §3.2. Resposta é pressão
  + remoção da peça-chave.

## D4 · Orçamento

**A troca é neutra: R$ 0,00 → R$ 0,00.** As duas cartas são da caixa (`Requisition Raid` adquirida
em 2026-09-20, custo zero; `Fumigate` já paga, R$ 1,37, **volta para as sobressalentes** quando o
orquestrador aplicar o corte). O total das faltantes **continua R$ 156,72** e a folga **continua
R$ 43,28** — LigaMagic (menor), cotações de 2026-09-19/20.

**Recomendação de uso da folga (não é parte desta troca, é encaminhamento):** o buraco mais caro
da §D3 é *remoção flexível em instant*, hoje em 2 e nenhuma instantânea. `Generous Gift` (R$ 5,00,
*destroy any permanent*, instant) resolve por **3,5% do teto** e é a melhor carta por real de toda
a fase. `Chaos Warp` (R$ 3,81) é a segunda. As duas juntas custam R$ 8,81 e cabem com folga — mas
**exigem 2 slots**, e os slots vieram do meu escopo só uma vez. Fica registrado para o orquestrador
decidir junto com a pendência 1 (ramp explosivo).
