# 00 — Briefing · Tori D'Avenant, Fury Rider

- **Modo:** `improve` (rodada v1 aberta em 2026-09-18, **reaberta em 2026-09-19** com o comandante decidido)
- **Comandante atual:** **Otharri, Suns' Glory** — `{3}{R}{W}` · 3/3 · Legendary Creature — Phoenix
  — **decidido pelo usuário em 2026-09-19** (ver `decisions.md`).
- **Comandante anterior:** Tori D'Avenant, Fury Rider — `{1}{R}{R}{W}` · 3/3 · Human Knight.
  Sai da zona de comando; **é RW, então continua elegível para os 99** — a avaliação dela como
  carta do 99 é tarefa da Fase 2 (re-pass com Otharri), não decisão tomada.
- **Identidade de cor:** RW (inalterada)
- **Tema:** **ataque em massa com exército próprio** — Otharri gera fichas de Rebel que escalam
  com contadores de experiência. O eixo muda de *battalion/knights com pump temporário* para
  *go-wide com fichas + anthems permanentes + resiliência a wipe*. O pool temático precisa ser
  reavaliado sob esse eixo.
- **Tamanho atual:** 100 cartas exatas (comandante + 99), conferido em `lista.txt`
  (**36 terrenos**: 30 básicos — 15 Mountain + 15 Plains — + 6 não-básicos: Boros Guildgate,
  Command Tower, Fields of Strife, Sandstone Bridge, Stone Quarry, Wind-Scarred Crag)

  > ⚠ **Correção de inventário (v1, 2026-09-18).** A primeira leitura deste briefing anotou
  > 38 terrenos. Era erro de leitura do agregado `Land 8` do `mtgdb`, que conta **nomes
  > distintos** (Mountain e Plains entram como 1 cada), não cópias. `Esgaroth Garrison`
  > também foi confundida com terreno — é `Creature — Human Soldier */5`. O deck está
  > **−2 na base de terrenos**, não em 38.
- **Status físico:** **montado** — confirmado pelo usuário em **2026-09-18**.
  Lista de referência: `lista.txt`. Deck construído fora do pipeline.

## Intake da otimização (2026-09-18)

### Objetivo
Tornar o deck **extremamente competitivo no Commander 200**.

### Regras do formato (confirmadas pelo usuário)
- Teto de **R$ 200,00** para as **99 cartas** do deck.
- **O comandante NÃO conta na régua** (declarado pelo usuário em 2026-09-19). Consequência
  direta: a escolha do comandante é **puramente de potência** — nenhum candidato pode ser
  descartado por preço, e um comandante caro **libera** orçamento em vez de consumir.
- Régua: **menor valor** do bloco "Preço Médio de Venda no Marketplace" da **LigaMagic**.
- Banlist oficial de Commander, sem restrições extras e sem teto por carta.
- **Terrenos básicos:** o usuário não tem certeza se contam. **Assumimos que contam**
  (critério conservador; 30 básicos ≈ R$ 5, não inverte nenhuma decisão).

### Escopo do comandante
> *"Se para ele ficar competitivo for necessário trocar a Tori por outro comandante dentro do
> tema e cor, por mim OK. Mas verifique como fica com ela primeiro."*

Ordem obrigatória: **auditar com Tori no comando primeiro**. A troca só entra em pauta se a
auditoria mostrar que ela não sustenta o alvo de potência — e, se entrar, é **dentro de RW e
dentro do tema de ataque**, nunca uma reconstrução para outra cor ou outro arquétipo.

### Pool disponível
> **Atualizado em 2026-09-19:** o comandante mudou, então a cláusula condicional abaixo **entrou
> em vigor**. O usuário autorizou desmontar o Tori. As **69 não-básicas de `lista.txt` existem
> fisicamente** (deck montado, confirmado em 2026-09-18) e são pool disponível **sem custo de
> compra** — os especialistas devem varrê-las junto com as sobressalentes, antes do Scryfall.
> Compra-se apenas **Otharri (R$ 24,79)** e o que os gaps exigirem.

- **Pool sem custo:** 69 não-básicas de `lista.txt` + 30 terrenos básicos + 113 sobressalentes
  de `data/collection.tsv`.
- **Cláusula original (cumprida):** *"Se o comandante mudar, o usuário autoriza desmontar o Tori."*
- **Sobressalentes:** 113 cartas em `data/collection.tsv` (`bin/mtgdb collection -list`).
- **Modo:** caixa primeiro **com compras permitidas** (padrão). Não é modo restrito.

### Orçamento de compra
**Sem teto de compra.** O único limite é o R$ 200 do formato, **medido só sobre as 99 cartas**.
Monta-se o melhor deck possível dentro da régua e o usuário compra o que faltar.

### Dores relatadas (todas as quatro marcadas)
1. **Não fecha o jogo** — monta board, ataca, não converte em vitória.
2. **Mão morta / sem cartas** — esvazia a mão cedo e fica em topdeck.
3. **Sem respostas** — não lida com as ameaças da mesa.
4. **Morre para board wipe** — investe o board inteiro e perde tudo de uma vez.

> **Tradução em alvo, não em corte (Fase 0, item 3):** as quatro dores são *lacunas de
> categoria*, não reclamações contra peças específicas. Nenhuma carta do deck foi apontada
> como culpada, então nenhuma está sob acusação — mas nenhuma está sob proteção por essa via.
> A dor 4 ("morre para board wipe") é o alvo mais delicado: a resposta correta é **resiliência**
> (proteção, recursão, ameaças que sobrevivem), **não** reduzir a quantidade de criaturas —
> o deck é de ataque em massa e esvaziar o board destrói o objetivo.

### Cartas intocáveis
**Sol Ring**, **Arcane Signet**, **Command Tower** — permanecem no deck em qualquer cenário,
inclusive se o comandante mudar.

## Rodadas

| Rodada | Data | Estado | Relatório |
|---|---|---|---|
| v1 | 2026-09-18 | **encerrada sem aplicação** — consolidada (20/09), testada em proxy (23/09), 7 trocas aplicadas no report; **reprovada pelo usuário em 2026-09-23** (eixo, não peças). Nada comprado, `deck.md` não escrito | `rounds/v1-2026-09-18/report.md` |
| v2 | 2026-09-23 | **consolidada, aguardando validação** — Phlage; Fases 1–7 completas; 99 cartas a R$ 176,71 (46 compras, R$ 148,16); 4 decisões do usuário na §0 do report. Nada comprado, `deck.md` não escrito | `rounds/v2-2026-09-23/report.md` |

---
## ESTADO — 2026-09-19

A rodada **v1 está em andamento**. A decisão travante (o comandante) foi resolvida:
**Otharri, Suns' Glory**. Nenhum swap de carta foi proposto ainda, `deck.md` ainda não existe.

### O que já está feito
- Intake completo (acima).
- **Fase 1 concluída (2026-09-19):** comandante decidido. Análise dos três finalistas com texto
  oracle, imposto, cotação LigaMagic e cruzamento com as quatro dores no painel abaixo.
- **Fase 2 (primeiro passe, para Tori):** auditoria do `theme-analyst` em
  `rounds/v1-2026-09-18/02-theme.md` — 723 linhas: análise linha a linha da Tori, as 99 cartas
  classificadas, ficha F1–F7 dos cortes, varredura da coleção, pool de 42 candidatas.
  **A classificação das 99 cartas continua válida; a análise do comandante e o eixo temático não.**

### Entregável da rodada
`rounds/v1-2026-09-18/report.md` — **o consolidado das 6 fases**, com nomes de carta clicáveis
para a LigaMagic e bloco de importação em texto puro para Moxfield/MTGO. É o único arquivo que
precisa ser lido para validar o deck; os relatórios por fase ficam ao lado para consulta pontual.

### Correções aplicadas à auditoria pelo orquestrador
1. **36 terrenos, não 38** — o agregado `Land 8` do `mtgdb` conta nomes distintos;
   `Esgaroth Garrison` é criatura. Já corrigido no topo deste briefing.
2. **Adeline não é a troca mais barata — é a mais cara.** O relatório a classificou como
   "custo mínimo"; a conferência carta a carta dá **34 slots** (19 cartas com vermelho + 15
   Mountain), porque ela é mono-branca. Descartada.
3. **A auditoria não trouxe os candidatos vencedores.** Varredura sem filtro de preço
   (possível porque o comandante não conta na régua) encontrou `Otharri, Suns' Glory` e
   `Aurelia, the Law Above` — o primeiro virou o comandante escolhido.

### Diagnóstico em uma linha
43 dos 64 slots não-terreno são tema (pump temporário de Limited) e ~10 são função. Gaps:
draw −9 · ramp −9/−3 · interação −6 · wipes −4 · wincons −3 · terrenos −2. **~33 slots mudam
em qualquer cenário.** Otharri zera o gap de wincon de graça (não conta na régua).

### O que muda no eixo temático com Otharri
O pump temporário da Tori (`Basri's Solidarity`, `Inspiring Roar`, `Pride of Conquerors`,
`Shoulder to Shoulder`, `Zealous Display`…) era sinérgico com um gatilho que pumpava no ataque.
Otharri **não pumpa** — ele *cria corpos*. Consequências a testar na Fase 2:
- **Anthems permanentes** (`Valor in Akros`, `The Circle of Loyalty`) sobem de valor: buffam
  fichas que continuam existindo. **Pump temporário de um turno** desce.
- **Rebels: resolvido em 2026-09-19 — o deck NÃO precisa de pacote de Rebels.** As fichas entram
  viradas, mas **desviram no turno seguinte** e pagam o custo da recursão a partir daí. O que vale
  é dar **vigilância às fichas** (Radiant Destiny nomeando Rebel, Intangible Virtue) ou usar
  equipamentos *For Mirrodin!*, que trazem um Rebel desvirado **e** um artefato para o metalcraft.
- **Contadores de experiência** são o motor de resiliência. Nada no deck atual interage com eles.
- **Vigilância / desvirar** e **combate extra** valem mais do que valiam (mais gatilhos de ataque
  = mais contadores). `Aurelia, the Warleader` está fora por preço (R$ 58,50 = 29% do teto).
- **Tori D'Avenant** passa a ser candidata a carta do 99, não peça obrigatória.
  > ⚠ **Correção 2026-09-19 (hipótese minha, refutada pela Fase 2).** Eu havia escrito que a Tori
  > "destrava um segundo ataque parcial". **Falso:** a cláusula de untap dela é para criaturas
  > **brancas** atacantes, e as fichas do Otharri são **vermelhas**. Ela não destrava o enxame.
  > O que a mantém no pool é a outra cláusula — **trample às criaturas vermelhas atacantes**, que
  > pega as fichas — somada ao +1/+1 do mesmo gatilho. Quem destrava segundo ataque de verdade é
  > `Combat Celebrant` / `Great Train Heist` / `Éomer, Marshal of Rohan`.

- **Metalcraft do Jor Kadeen é um multiplicador maior do que parecia.** `Creatures you control get
  +3/+0` com 3+ artefatos, num deck de enxame, é dano em cada ficha. Isso dá **segunda função aos
  artefatos de mana** e foi passado como sinal cruzado ao `ramp-specialist`.

- **`Paladin Danse` não protege nada que importa** — a indestrutibilidade dele é só para
  **artefato ou Humano**; as fichas são Rebel e o comandante é Phoenix. Devolvido ao
  `interaction-specialist` como corte condicionado (a função de proteção em massa é a dor 4 e
  precisa de substituto nomeado).

### Próximos passos
1. **Fase 2 — re-pass temático com Otharri** (`theme-analyst`), destino
   `rounds/v1-2026-09-18/02-theme.md`. Reaproveita a classificação das 99 cartas, refaz o eixo.
2. Fan-out dos especialistas em `improve`: **ramp → draw → interação → wincons → manabase**,
   passando `rounds/v1-2026-09-18/` como destino.
3. **Preços:** só 8 de 70 cartas têm cotação. Nenhum total em reais foi calculado. A captura na
   LigaMagic é manual (navegador) e deve ser feita quando a lista de entradas fechar — a régua
   incide **só sobre as 99 cartas**, e o Otharri fica fora dela.

---
## DECISÃO EM ABERTO — eixo do deck (2026-09-20)

O usuário observou que o deck, **tanto na lista antiga quanto na v1 com Otharri**, joga na mesma
pegada do **Krenko** (enxame + anthem + dano que ignora combate) e, em menor grau, do **Thorin**
(RW, tribo com anthem que ataca). A observação procede: a semelhança **não** foi criada pela troca
de comandante — a lista da Tori já era go-wide de combate, só que sem motor de fichas (2 geradores
em 99 cartas) e com pump de um turno.

**Fatos levantados para a decisão:**
- Sobreposição de cartas entre os três decks é baixa: 5 com o Krenko, 5 com o Thorin, quase tudo
  staple. Mas 3 das 5 do Krenko não são staple — `Battle Hymn`, `Light Up the Stage` e
  `Outpost Siege` já estão dentro dele montado (regra 7: comprar-se-ia a segunda cópia).
- **Nenhum dos outros dois cabe no Commander 200.** Krenko acumulou R$ 543,14 em 34 compras até a
  v4 e carrega Purphoros (R$ 129,99) e Fable of the Mirror-Breaker (R$ 55,00). Thorin foi
  construído sem teto (R$ 178,68 só nas faltantes). A v1 do Otharri fecha as 99 em R$ 155,62.
- **A caixa de sobressalentes favorece eixo de mágica, não de criatura:** das 82 cartas jogáveis em
  RW, só **25 são criaturas** — 44 são instantâneos/feitiços e 31 são artefatos.
- Mudar de eixo descarta quase todas as 69 não-básicas da Tori física (criaturas pequenas de
  combate). **Nenhum custo de eixo alternativo está medido**; o R$ 155,62 da v1 só é baixo porque
  reaproveita 27 cartas do deck montado.

**Opções apresentadas (2026-09-20):**

| Eixo | Comandante | Como joga | Risco |
|---|---|---|---|
| Spellslinger de tricks | `Feather, the Redeemed` | ~30 mágicas baratas que miram criatura sua e voltam para a mão | depende de manter criatura viva |
| Recursão de cemitério | `Hofri Ghostforge` | criatura que morre volta como cópia Spirit +1/+1; wipe vira lucro | exige criaturas boas com ETB — nenhuma na caixa |
| Mágicas grandes do exílio | `Commander Liara Portyr` | atacar é gatilho, não plano; conjura do exílio com desconto | "mágica grande boa" briga com R$ 200 |
| Anjo de controle | `Gisela, Blade of Goldnight` | dano dobrado, metade do recebido; 2–3 ameaças grandes | CMC 7; precisa sobreviver até lá |
| Legendary copy | `Cadric, Soul Kindler` | copia cada lendária não-ficha que entra | pool de lendárias com ETB dentro do teto |

Descartados de saída: **equipamento/artefato** (Wyleth, Akiri, Nahiri, Bruenor, Reyav, Astor,
Alibou) por ser o Thorin; **go-wide** (Anim Pakal, Iroas, Winota, Aurelia the Law Above, Adeline)
por ser o Krenko e a própria v1.

### Pendência atrelada a esta decisão
- **`Requisition Raid`** entrou na coleção em 2026-09-20 e é candidata à v1 **se o eixo continuar
  sendo Otharri**. Troca natural: `Basri's Solidarity` → `Requisition Raid` (custo zero nos dois
  lados). Motivo: por `{2}{W}` mata artefato **e** encantamento no mesmo card (Disenchant e Seal of
  Cleansing escolhem um), e o terceiro modo do Spree é anthem em **contadores permanentes**, que é
  exatamente o que o Basri's faz — só que o Basri's não faz mais nada. Ressalvas: é feitiço, não
  segura combate; e contadores morrem com as criaturas, então não ajuda na dor 4.
  **Decisão adiada a pedido do usuário até o eixo estar resolvido.**

---
## TESTE DE MESA — 2026-09-23

O usuário **imprimiu a lista v1 em proxy e jogou com ela**. Nenhuma das 31 compras (R$ 166,23) foi
feita — **não há custo afundado e os slots estão livres**. Este é o primeiro dado de jogo real da
rodada; o goldfishing da pendência 1 do `report.md` continua não rodado, mas foi superado por ele.

### Relato, verbatim
> "Fiz o teste com o deck do tori/otharri, fiquei com muita mágica instantânea e encantamentos na
> mão e poucas criaturas, senti lentidão para descer o comandante e quando desço e o perco, acabo
> não usando o custo fixo por ficar sem tokens de rebeldes"

### Diagnóstico do orquestrador — as três queixas têm uma raiz só
**O tabuleiro inteiro do deck depende do comandante.** Cada queixa é uma face disso:

1. **Poucas criaturas na mão.** Não é azar, é a composição: **15 criaturas** em 99 (17 artefatos,
   11 encantamentos, 10 instantâneos, 9 feitiços, 1 PW, 36 terrenos). Hipergeométrica da mão
   inicial de 7: **30,4% abrem sem nenhum corpo**, só **28,6%** abrem com 2+; esperado de 1,06
   criatura contra 2,12 instantâneo/feitiço/encantamento. E a maioria dos 11 encantamentos é
   **anthem estático** — carta que não faz nada sem board, num deck que não tem board sem o
   comandante.
2. **Lentidão para descer o comandante.** CMC 5 com 36 terrenos e o grosso do ramp em CMC 2.
   Devolvido à Fase 4 para medir P(Otharri no T4) e P(no T5) e nomear a causa dominante.
3. **Recursão travada por falta de Rebel.** Confirmado, e expõe **um erro técnico na §11.1 do
   `report.md`**: ela defende a `Intangible Virtue` dizendo que "vigilância na ficha é o que paga a
   recursão do Otharri". **Falso** — vigilância não desvira o que já entrou virado, e as fichas do
   Otharri entram **viradas e atacando**. Só fichas de turnos *anteriores*, sobreviventes ao
   combate, pagam o `{2}{R}{W}`. Se o Otharri morre no primeiro ataque ou num wipe, não há Rebel.
   A única outra fonte de Rebel do deck é a `Hexgold Halberd`, uma carta em 99.

### Devoluções — RESOLVIDAS em 2026-09-23, consolidadas na §12 do `report.md`
**7 trocas aplicadas.** Corpos 19 → 25 · mão sem corpo 30,4% → 12,1% · Otharri no T4 42,6% → 47,9%
· Rebel desvirado independente do comandante 1 → 3 · custo **R$ 166,23 → R$ 139,16** (folga
R$ 60,84). Aguardando aprovação do usuário para escrever `deck.md` e `decisions.md`.
Recusei o `Hexplate Wallbreaker` (R$ 49,90 — redundante com o `Combat Celebrant`, já na lista a
R$ 11,99) e o 37º terreno (custaria um slot de corpo). Devolvida ao usuário a troca
`Feat of Resistance` → `Glimmer Lens` (R$ 21,75), que troca proteção por saque.

### O que foi pedido a cada fase (histórico)
- **Fase 2 — `theme-analyst`:** densidade de corpos (meta proposta: criaturas na casa de 24–28, com
  P(mão sem corpo) < ~12%, justificada pela hipergeométrica); fontes de Rebel desvirado
  independentes do comandante — o pacote **`For Mirrodin!`** em RW (`Barbed Batterfist`,
  `Bladehold War-Whip`, `Hexplate Wallbreaker`, `Vulshok Splitter`…) entrega Rebel desvirado +
  artefato para o metalcraft do `Jor Kadeen` + corpo num slot só; cortes saem do excesso de mágica
  situacional e de anthem redundante, não das metas de função. `Goblin Rabblemaster` e
  `Knight Watch` estão **na caixa** e a v1 os ignorou.
- **Fase 4 — `ramp-specialist`:** velocidade e confiabilidade do mana até `{3}{R}{W}`, e se a base
  branca sustenta recast + proteção no mesmo turno (o `Mana Geyser` produz só `{R}`).

### Consequência de processo
A validação do `report.md` **não é gatilho de compra**. O fluxo real é: report aprovado → proxy →
teste de mesa → ajuste → compra. Trocas propostas antes da compra são de graça; depois, não.

---
## RODADA v2 — ABERTA EM 2026-09-23 · Fase 1 reaberta

O usuário **reprovou a v1** depois do teste de mesa. Marcou três incômodos, nesta ordem de
peso declarada: **não parece competitivo**, **o comandante (Otharri)** e **o jeito de jogar**.
Autorização dada: **trocar o comandante, mantendo RW**.

### O que isso resolve e o que não resolve
As 7 trocas da §12 do `report.md` da v1 consertaram as **três queixas medidas** do teste
(corpos, velocidade do comandante, Rebel desvirado) — e ainda assim o deck não convenceu. Isso
é o sinal de que o problema **não era de peças**: é a `DECISÃO EM ABERTO — eixo do deck`
levantada em 2026-09-20 e nunca fechada. Reabrir a Fase 1 é fechá-la.

### Restrições desta rodada (inalteradas onde não digo o contrário)
- **Identidade RW obrigatória.** Não é reconstrução para outra cor.
- **Comandante fora da régua de preço** (declarado em 2026-09-19): a escolha é puramente de
  potência, e um comandante caro **libera** orçamento em vez de consumir.
- **R$ 200,00 nas 99 cartas**, menor valor da LigaMagic.
- **Pool sem custo:** 69 não-básicas de `lista.txt` (Tori físico, autorizado a desmontar) +
  30 básicos + 113 sobressalentes de `data/collection.tsv`. Mas **aproveitamento do pool não é
  critério de escolha do comandante** — é consequência a medir depois, e a caixa favorece eixo
  de mágica (44 instantâneos/feitiços e 31 artefatos contra 25 criaturas jogáveis em RW).
- **Intocáveis:** `Sol Ring`, `Arcane Signet`, `Command Tower`.
- **Regra 11 (diversão):** sem MLD, stax pesado ou lock — não foi pedido.

### O que "jeito de jogar" exclui
O usuário já apontou em 2026-09-20 que a lista joga na mesma pegada do **Krenko** (enxame +
anthem + dano fora do combate) e, em menor grau, do **Thorin** (RW com anthem que ataca). A v1
com Otharri **é go-wide de fichas** — é exatamente esse padrão. Um comandante novo que repita
go-wide de fichas não resolve o incômodo, só troca o nome.

### Estado da v1
**Encerrada sem aplicação.** Nenhuma compra feita (R$ 139,16 nunca gastos), `deck.md` nunca
escrito, `decisions.md` sem as 7 trocas. O `rounds/v1-2026-09-18/` fica como está — é retrato
de um momento e não se toca. A análise linha a linha das 99 cartas em `02-theme.md` e a
classificação de funções continuam **matéria-prima válida** para a v2.

### Comandante da v2 — decidido em 2026-09-23
**`Phlage, Titan of Fire's Fury`** — `{1}{R}{W}` · 6/6 · Legendary Creature — Elder Giant.
Escolhido pelo usuário entre os 5 finalistas da Fase 1 (`rounds/v2-2026-09-23/01-commander.md`).
**Identidade RW mantida.** Eixo: **controle de atrito** — remoção densa, pouco tabuleiro próprio,
ameaça que volta do cemitério por custo fixo. Registro e travas mecânicas em `decisions.md`.
`Otharri, Suns' Glory` e `Tori D'Avenant, Fury Rider` seguem elegíveis para os 99 como qualquer RW.
