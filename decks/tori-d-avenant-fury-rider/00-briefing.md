# 00 — Briefing · Tori D'Avenant, Fury Rider

- **Modo:** `improve` (rodada v1 aberta em 2026-09-18)
- **Comandante atual:** Tori D'Avenant, Fury Rider — `{1}{R}{R}{W}` · 3/3 · Legendary Creature — Human Knight
- **Identidade de cor:** RW
- **Tema:** ataque em massa / battalion / knights — gatilho de ataque que pumpa e destapa
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
- **Enquanto Tori for o comandante:** o deck continua montado. Cartas cortadas voltam para a
  caixa como sobressalentes (existem fisicamente — o deck está montado).
- **Se o comandante mudar:** o usuário autoriza **desmontar o Tori**. As 69 não-básicas de
  `lista.txt` passam a contar como pool disponível, sem custo de compra.
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
| v1 | 2026-09-18 | **pausada** — aguarda escolha do comandante | `rounds/v1-2026-09-18/02-theme.md` |

---

## PONTO DE PARADA — 2026-09-19

A rodada **v1 está aberta e pausada** a pedido do usuário, aguardando **uma única decisão: o
comandante**. Nenhum swap foi proposto, nenhuma carta entrou ou saiu, `deck.md` ainda não existe
e `decisions.md` está vazio. Nada a desfazer.

### O que já está feito
- Intake completo (acima).
- **Fase 1 concluída:** auditoria do `theme-analyst` em `rounds/v1-2026-09-18/02-theme.md`
  (723 linhas: análise linha a linha da Tori, as 99 cartas classificadas, ficha F1–F7 dos
  cortes, varredura da coleção, pool de 42 candidatas).
- **Painel de diagnóstico publicado:** https://claude.ai/artifact/2xU9fQUXZdHVkHaQdLmb1X
  (gaps por categoria, curva, as 100 cartas com função, comparação de comandantes).

### Correções aplicadas à auditoria pelo orquestrador
1. **36 terrenos, não 38** — o agregado `Land 8` do `mtgdb` conta nomes distintos;
   `Esgaroth Garrison` é criatura. Já corrigido no topo deste briefing.
2. **Adeline não é a troca mais barata — é a mais cara.** O relatório a classificou como
   "custo mínimo"; a conferência carta a carta dá **34 slots** (19 cartas com vermelho + 15
   Mountain), porque ela é mono-branca. E a afirmação de que "a Tori continua no 99 em
   qualquer troca" é falsa no caso dela: Tori é RW e ficaria fora da identidade.
3. **A auditoria não trouxe dois candidatos melhores.** Varredura sem filtro de preço
   (possível porque o comandante não conta na régua) encontrou `Otharri, Suns' Glory` e
   `Aurelia, the Law Above`, ambos superiores à recomendação original.

### Diagnóstico em uma linha
43 dos 64 slots não-terreno são tema (pump temporário de Limited) e ~10 são função. Gaps:
draw −9 · ramp −9/−3 · interação −6 · wipes −4 · wincons −3 · terrenos −2. **~33 slots mudam
em qualquer cenário**, com ou sem troca de comandante.

### A decisão pendente
O usuário confirmou que a régua é **competir no Commander 200**, e portanto a Tori sai do
comando (ela não gera carta, corpo, mana nem fechamento). Ele **gostou da Aurelia** mas
levantou o imposto de comandante: *"se ela voltar pra zona de comando já são 2 a mais de
custo, indo pra 8"* — a conta de 6→8 é da **Warleader**.

**Pergunta em aberto, não respondida:** qual Aurelia ele quis dizer — `the Warleader` (CMC 6,
combate extra) ou `the Law Above` (CMC 5, compra carta sempre que *qualquer* jogador ataca
com 3+ criaturas)?

| Candidato | CMC | Imposto | Argumento |
|---|---|---|---|
| **Otharri, Suns' Glory** | 5 | **nenhum** — volta do cemitério por `{2}{R}{W}` fixo, sem passar pela zona de comando | recomendação; contadores de experiência ficam no jogador e sobrevivem a wipe |
| Aurelia, the Law Above | 5 | 5→7→9 | resolve o gap de draw (−9) no próprio comando |
| Aurelia, the Warleader | 6 | 6→**8**→10 | o efeito que o usuário gostou; a mais exposta ao receio dele |
| Anim Pakal | 3 | 3→5→7 | mais rápida; perde todo o progresso se morrer |
| Agrus Kos, Spirit of Justice | 4 | 4→6→8 | remoção repetível; resolve a dor 3 |

Mitigações levantadas para o receio do imposto: `Command Beacon` (terreno, põe o comandante na
mão da zona de comando — **preço ainda não capturado**), `Swiftfoot Boots` (R$ 14,00),
`Lightning Greaves` (R$ 25,00 — 12,5% do teto numa carta só, desaconselhado).

### Ao retomar
1. Perguntar apenas o comandante — todo o resto do intake está coletado.
2. Disparar os especialistas em modo `improve` na ordem: **ramp → draw → interação →
   resiliência a wipe → wincons → manabase**, passando `rounds/v1-2026-09-18/` como destino.
3. **Preços:** só 5 de 70 cartas têm cotação. Nenhum total em reais foi calculado ainda. A
   captura na LigaMagic é manual (navegador) e deve ser feita quando a lista de entradas
   fechar — a régua incide **só sobre as 99 cartas**.
