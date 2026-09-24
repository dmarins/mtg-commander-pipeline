# 01 — Opções de Comandante · rodada v2 (2026-09-23)

**Fase 1 reaberta.** A v1 (Otharri, go-wide de fichas de Rebel) foi reprovada em teste de mesa.
As 7 trocas da §12 do `rounds/v1-2026-09-18/report.md` corrigiram as queixas *medidas* e o deck
ainda assim não convenceu — a causa é o **eixo**, não as peças.

**Restrições desta fase:** identidade RW · `legal:commander` · comandante **fora** do teto de
R$ 200 (escolha puramente de potência) · R$ 200 nas 99 · regra 11 (sem MLD/stax/lock) ·
eixos **go-wide de fichas** e **equipamento/voltron** excluídos pelo usuário.

> **Regra 2 — preço.** Nenhum número de preço do Scryfall entra aqui. `bin/mtgdb prices` foi
> consultado para os 12 comandantes cogitados e para 18 peças-chave: **nenhum dos comandantes
> tem cotação registrada** — todos saem como **`a cotar`**. Como o comandante está fora da régua,
> isso **não bloqueia a decisão**; a cotação só é necessária para as 99.
>
> **Regra 6 — oracle.** Todo texto abaixo foi puxado com `bin/mtgdb oracle` nesta sessão. Os
> rulings do Phlage foram puxados com `bin/mtgdb rulings` (ponto load-bearing da análise dele).

---

## 0. Varredura feita

Pool completo de comandantes RW consultado no Scryfall (regra 1): `is:commander legal:commander
id=rw order:edhrec` em duas fatias — `mv<=5` (135 resultados) e `mv>=6` (10 resultados). **145
lendárias RW exatas**, o pool inteiro, lido nome a nome. Descartados de saída os que caem nos
eixos vetados (Anim Pakal, Iroas, Winota, Neyali, Adeline, Aurelia the Law Above, Duke Ulder
Ravengard, Ertha Jo, Commander Mustard / Wyleth, Akiri ×2, Nahiri, Bruenor, Reyav, Astor, Alibou,
Osgir, Sami Wildcat Captain, Nori, Dwalin, Kolodin, Depala, Cass, Tiana, Mabel, Amy Rose).

Pool sem custo medido com `bin/mtgdb collection` + `lista.txt`: **85 cartas RW-legais na caixa**
(34 instantâneos/feitiços · 28 artefatos · 18 criaturas · 4 encantamentos · 1 terreno) + as 69
não-básicas do Tori físico. A caixa é, de fato, uma **caixa de mágica e queima**: 18 das 34
mágicas são remoção ou dano direto. Isso é informação de custo, **não** critério de escolha.

---

## 1. Tabela comparativa

| # | Comandante | CMC | Eixo (jeito de jogar) | Como fecha | Dependência do comandante | Aproveitamento do pool | Risco principal |
|---|---|---|---|---|---|---|---|
| 1 | [**Phlage, Titan of Fire's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Phlage%2C+Titan+of+Fire%27s+Fury) | 3 | Controle de atrito: remoção densa, poucas criaturas, ameaça que **volta sozinha do cemitério** | Inevitabilidade — Helix recorrente + 6/6 **pelado** (⚠ corrigido em 2026-09-23 pela Fase 2: sem trample, sem lifelink — o ganho de vida vem do gatilho) + queima final | **Baixa** (única RW que reconjura **sem imposto**) | **Alto** — 18 remoções/queimas + Fumigate + Syr Carah na caixa | Fecha devagar; precisa de finisher nomeado |
| 2 | [**Feather, the Redeemed**](https://www.ligamagic.com.br/?view=cards/card&card=Feather%2C+the+Redeemed) | 3 | Spellslinger reativo: ~30 mágicas baratas que miram criatura sua e **voltam para a mão** | Uma criatura protegida e inflada, ou queima em loop | **Muito alta** — sem ela as mágicas são cartas mortas | Médio — ~6 truques da `lista.txt`, quase nada da caixa | É a falha da v1 repetida: tudo depende de 1 permanente |
| 3 | [**Hofri Ghostforge**](https://www.ligamagic.com.br/?view=cards/card&card=Hofri+Ghostforge) | 5 | Recursão: suas criaturas **morrem e voltam melhores**; wipe vira lucro | Exército de Spirits +1/+1 com trample e haste | **Alta** | **Baixo** — a caixa tem 3 criaturas com ETB que valem cópia | Precisa comprar quase todas as criaturas boas |
| 4 | [**Quintorius, Loremaster**](https://www.ligamagic.com.br/?view=cards/card&card=Quintorius%2C+Loremaster) | 5 | Cemitério como segunda mão: conjura suas melhores mágicas **de graça**, um por turno | Motor de valor + free-cast de wipe/queima | **Alta** | Médio-alto — as 34 mágicas da caixa viram munição reciclável | Lento; morre para exílio de cemitério |
| 5 | [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | 7 | Anjo de controle: **dano dobrado** para eles, **metade** para você | Qualquer queima vira letal; 2–3 ameaças voadoras enormes | Média (deck funciona sem ela, mas ela é o fecho) | **Alto** — as mesmas 18 queimas, com o dobro do dano | CMC 7 com imposto 7→9→11 num deck de R$ 200 |

---

## 2. Reavaliação dos 5 candidatos levantados em 2026-09-20

O briefing pede confirmação ou refutação de cada risco anotado. Com oracle na mão:

| Candidato | Risco anotado em 20/09 | Veredito com oracle |
|---|---|---|
| [**Feather, the Redeemed**](https://www.ligamagic.com.br/?view=cards/card&card=Feather%2C+the+Redeemed) | "depende de manter criatura viva" | **CONFIRMADO e agravado.** O gatilho exige *instantâneo ou feitiço que mire criatura que você controla*. Sem criatura no campo, **nenhuma** mágica do pacote é conjurável de forma útil, e sem a Feather nenhuma delas volta. **Segue finalista** (nº 2) porque o motor de vantagem de cartas é real, mas é exatamente a forma de dependência que reprovou a v1. |
| [**Hofri Ghostforge**](https://www.ligamagic.com.br/?view=cards/card&card=Hofri+Ghostforge) | "exige criaturas boas com ETB — nenhuma na caixa" | **CONFIRMADO por contagem.** Das 18 criaturas RW-legais da caixa, só [**Redcap Thief**](https://www.ligamagic.com.br/?view=cards/card&card=Redcap+Thief) (Treasure), [**Pilgrim's Eye**](https://www.ligamagic.com.br/?view=cards/card&card=Pilgrim%27s+Eye) (terreno) e [**Phyrexian Revoker**](https://www.ligamagic.com.br/?view=cards/card&card=Phyrexian+Revoker) têm ETB, e nenhuma justifica uma cópia Spirit. **Segue finalista** (nº 3), mas com custo de compra alto declarado. |
| [**Commander Liara Portyr**](https://www.ligamagic.com.br/?view=cards/card&card=Commander+Liara+Portyr) | "'mágica grande boa' briga com R$ 200" | **PARCIALMENTE REFUTADO — e substituído por um risco pior.** O oracle **não** exige mágica grande: `spells you cast from exile this turn cost {X} less`, com X = jogadores atacados. Atacando os 3, qualquer mágica custa {3} a menos, e o gatilho exila 3 cartas jogáveis no turno. É um motor de *impulse* com desconto, não de bombas. **Mas:** o gatilho é `whenever you attack`, com um corpo **5/3** — morre para qualquer bloqueio ou queima, o motor inteiro é a comandante, e as cartas exiladas **somem no fim do turno** se não forem usadas (desperdício em mão travada de mana). **NÃO segue** — é dependência total do comandante num corpo frágil, que é o critério de eliminação nº 1 desta rodada. |
| [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) | "CMC 7; precisa sobreviver até lá" | **CONFIRMADO, mas mal enquadrado.** O problema não é *sobreviver* — é *chegar*. Ela é o único candidato cujo deck **precisa funcionar sozinho por 7 turnos**, e isso é uma virtude disfarçada: o deck não pode ser dependente dela por construção. O risco real é o **imposto 7→9→11**, não a fragilidade. **Segue finalista** (nº 5). |
| [**Cadric, Soul Kindler**](https://www.ligamagic.com.br/?view=cards/card&card=Cadric%2C+Soul+Kindler) | "pool de lendárias com ETB dentro do teto" | **CONFIRMADO e eliminatório.** Duas travas somadas: (a) o token é **sacrificado no início do próximo end step**, então só copia ETB e ataque — é efeito de rajada, não de tabuleiro; (b) exige um deck de lendárias RW com ETB relevante, e o pool barato disso em RW é raso (a caixa tem 4 lendárias RW-legais: `Breeches`, `Ori`, [**Goblin Rabblemaster**](https://www.ligamagic.com.br/?view=cards/card&card=Goblin+Rabblemaster) não é lendária, [**Smaug**](https://www.ligamagic.com.br/?view=cards/card&card=Smaug)). Sem Cadric, o deck é uma pilha de lendárias medianas. **NÃO segue.** |

---

## 3. Análise detalhada

### 1. Phlage, Titan of Fire's Fury — {1}{R}{W} — RW

- **Tipo**: Legendary Creature — Elder Giant · 6/6 · CMC 3
- **Texto (oracle)**:
  > When Phlage enters, sacrifice it unless it escaped.
  > Whenever Phlage enters or attacks, it deals 3 damage to any target and you gain 3 life.
  > Escape—{R}{R}{W}{W}, Exile five other cards from your graveyard.

**Ficha F1–F7**

- **F1 (texto, linha a linha)** — (i) *sacrifice unless it escaped*: conjurado da zona de comando ele **morre no mesmo turno**; (ii) *enters or attacks → 3 de dano em qualquer alvo + 3 de vida*: um [**Lightning Helix**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Helix) que dispara tanto ao entrar quanto a cada ataque — remoção, alcance ou estabilização, à escolha; (iii) *escape {R}{R}{W}{W}, exilar 5 do cemitério*: reconjura **do cemitério**, e o custo de escape é **fixo**.
- **F2 (corpo)** — 6/6 por **4 de mana** (`{R}{R}{W}{W}`) + exilar 5 outras cartas, quando escapado. Bloqueia qualquer coisa. **Sem trample e sem lifelink** — é um 6/6 pelado; o ganho de vida vem do gatilho de entrada/ataque, não de lifelink. Não tem habilidade de tap. *(⚠ duas correções aplicadas em 2026-09-23: o texto original dizia "3 de mana efetivo" e a tabela §1 dizia "trample-lifelink".)*
- **F3 (tipo como recurso)** — Elder Giant; nenhum tema tribal relevante em RW. Conta como criatura para convocar/anthem.
- **F4 (recebe)** — anthems, equipamentos, contadores. Um 6/6 já é alvo de qualquer buff.
- **F5 (facilita)** — não dá nada a outras cartas. É ameaça e remoção, não facilitador.
- **F6 (curva)** — **CMC 3**: entra no T3 já como Helix. Escape a partir do T4–T5 com 5 cartas no cemitério.
- **F7 (atrito)** — compete com o cemitério de qualquer outro pacote de recursão (não haverá). O sacrifício obrigatório na primeira conjura é **feature, não bug**: coloca o card no cemitério onde ele quer estar.

**Como o deck joga**
- **T1–T3**: rampa barata ([**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring), [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet), [**Mind Stone**](https://www.ligamagic.com.br/?view=cards/card&card=Mind+Stone)) e remoção pontual nas ameaças da mesa. **T3**: Phlage da zona de comando = Lightning Helix (3 dano + 3 vida) e ele se sacrifica direto para o cemitério — você *quer* isso.
- **T4–T6**: você é o deck de respostas. Queima, remoção, um board wipe quando a mesa se encher. A cada abertura, escape do Phlage por {R}{R}{W}{W}: Helix de novo + um 6/6 lifelink para segurar o combate.
- **Fecho**: por atrito e inevitabilidade. Cada turno em que o Phlage está vivo é 3 de vida sua, 3 de vida do adversário e 6 de dano no combate. Os finishers dedicados (multiplicador de dano, uma ou duas ameaças evasivas) ficam para o `wincon-tester`.

**Por que é competitivo no Commander 200**
1. **Resiliência que dinheiro não compra.** O escape é uma *permissão de conjuração do cemitério*, e o custo é **fixo em {R}{R}{W}{W}** — `bin/mtgdb rulings` confirma: o imposto de comandante só incide sobre conjuras da **zona de comando**. Você escolhe deixar o Phlage no cemitério e reconjurá-lo pelo mesmo preço na 5ª vez que ele morrer. **Nenhum outro comandante RW tem isso.**
2. **Duas cartas em uma.** O corpo *é* a remoção. Num deck com teto de R$ 200, onde cada slot precisa render, um comandante que é simultaneamente Lightning Helix recorrente e finisher 6/6 economiza slots de wincon e de remoção.
3. **Custo 3.** Sai no T3 com ramp mediano. Isso mata a queixa de "lentidão para descer o comandante" que apareceu no teste da v1 (Otharri, CMC 5).

**Dependência do comandante** — **A mais baixa dos cinco.** O deck sob ele é um deck de remoção e queima que funciona sem tabuleiro próprio. Se o Phlage morrer duas vezes, a resposta correta é deixá-lo no cemitério e escapá-lo pela terceira vez pelo mesmo custo. É o único candidato em que "ele morreu duas vezes" **não é um problema**.

**As quatro dores**
- **3 (sem respostas)** — **resolvida estruturalmente**: o eixo *é* remoção.
- **4 (morre para wipe)** — **resolvida estruturalmente**: você tem pouco board a perder, você é quem dá o wipe, e o comandante volta do cemitério pelo custo cheio.
- **1 (não fecha o jogo)** — **parcialmente**: 3+6 de dano por turno é inevitabilidade, mas lenta em pod de 4. Fica como tarefa explícita do `wincon-tester`.
- **2 (mão morta)** — **fica para o `draw-specialist`**. Atenuada pelo custo baixo das peças e por impulse barato já cotado ([**Light Up the Stage**](https://www.ligamagic.com.br/?view=cards/card&card=Light+Up+the+Stage) R$ 0,74 · [**Outpost Siege**](https://www.ligamagic.com.br/?view=cards/card&card=Outpost+Siege) R$ 0,90, cotações de 2026-08-22).

**Aproveitamento do pool** — **Alto.** Das 85 cartas RW-legais da caixa, **~30** servem diretamente: 18 remoções/queimas, [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (wipe), [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) e [**Seize Opportunity**](https://www.ligamagic.com.br/?view=cards/card&card=Seize+Opportunity) (enchem o cemitério e compram), mais os artefatos de mana. Da `lista.txt`: [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold), [**Swift Reckoning**](https://www.ligamagic.com.br/?view=cards/card&card=Swift+Reckoning), [**Reduce to Memory**](https://www.ligamagic.com.br/?view=cards/card&card=Reduce+to+Memory), [**Seal of Cleansing**](https://www.ligamagic.com.br/?view=cards/card&card=Seal+of+Cleansing), [**Sol Ring**](https://www.ligamagic.com.br/?view=cards/card&card=Sol+Ring), [**Arcane Signet**](https://www.ligamagic.com.br/?view=cards/card&card=Arcane+Signet), [**Command Tower**](https://www.ligamagic.com.br/?view=cards/card&card=Command+Tower).
**As 5 mais relevantes**: [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) (exila o topo e deixa jogar **sempre que uma mágica sua causa dano a um jogador** — motor de cartas grátis, já no deck físico) · [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (wipe que você aguenta e eles não) · [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) (compra 2 **e** põe carta no cemitério = combustível de escape) · [**Flame Slash**](https://www.ligamagic.com.br/?view=cards/card&card=Flame+Slash) (4 de dano por {R}) · [**Punishing Fire**](https://www.ligamagic.com.br/?view=cards/card&card=Punishing+Fire) (dano recorrente barato).

**Viabilidade no teto** — **A melhor dos cinco.** O eixo não depende de nenhuma carta individualmente cara: remoção barata, wipes baratos, e o único pedaço tradicionalmente caro de RW — vantagem de cartas — tem saídas já cotadas a menos de R$ 1. Peças a cotar: [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act), [**Wrath of God**](https://www.ligamagic.com.br/?view=cards/card&card=Wrath+of+God), [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm), [**Solemn Simulacrum**](https://www.ligamagic.com.br/?view=cards/card&card=Solemn+Simulacrum), [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation), [**Sunforger**](https://www.ligamagic.com.br/?view=cards/card&card=Sunforger) — **a cotar**.

**Risco declarado** — **Relógio lento.** Sem um multiplicador de dano ou uma segunda ameaça evasiva, o deck pode controlar a mesa por 12 turnos e não fechar — que é literalmente a dor 1. O `wincon-tester` precisa entregar fecho nomeado, não "ataca com o comandante". Risco secundário: ódio a cemitério (Rest in Peace, Bojuka Bog) desliga o escape — mitigado por o Phlage continuar acessível pela zona de comando.

---

### 2. Feather, the Redeemed — {R}{W}{W} — RW

- **Tipo**: Legendary Creature — Angel · 3/4 · CMC 3
- **Texto (oracle)**:
  > Flying
  > Whenever you cast an instant or sorcery spell that targets a creature you control, exile that card instead of putting it into your graveyard as it resolves. If you do, return it to your hand at the beginning of the next end step.

**Ficha F1–F7**

- **F1** — (i) *voar*: evasão no corpo; (ii) o gatilho exige **instantâneo ou feitiço** que **mire criatura sua**; a carta é exilada em vez de ir ao cemitério e **volta para a mão** no próximo end step. Cada mágica do pacote vira reutilizável a cada turno.
- **F2** — 3/4 voador. Bloqueia bem, é alvo legítimo das próprias mágicas (protege-se sozinha).
- **F3** — Angel; irrelevante em RW barato.
- **F4** — recebe contadores e buffs; é a receptora natural de todo o pacote de truques.
- **F5** — **facilita muito**: transforma mágicas descartáveis em permanentes funcionais. É o único dos cinco cujo valor está inteiramente em *habilitar outras cartas*.
- **F6** — CMC 3; o motor liga no T3–T4.
- **F7** — **atrito grave**: o gatilho precisa de *outra* criatura sua para mirar quando a Feather não está em campo, e precisa da *própria* Feather em campo para devolver a carta. As duas condições falham juntas num board wipe.

**Como o deck joga**
- **T1–T3**: criaturas baratas e resilientes (alvos), rampa. T3 Feather.
- **T4–T6**: a cada turno você conjura 1–3 mágicas de 1 mana mirando sua própria criatura — proteção ([**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing)), contadores ([**Moment of Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Moment+of+Glory)), ou cantrips —, elas voltam para a mão no end step e você repete. A mão **não esvazia**, que é exatamente a dor 2.
- **Fecho**: uma criatura gigante com proteção/evasão passando 20+ por combate, ou o pacote de queima redirecionado.

**Por que é competitivo no Commander 200** — É o arquétipo **mais barato de todos** em RW: o pacote da Feather é feito de comuns de 1 mana. Vantagem de cartas real (a mesma carta usada N vezes) sem pagar por draw, que é o buraco caro de RW. E joga **no turno dos oponentes** — resposta a remoção, a combate, a gatilhos. É o candidato com o "jeito de jogar" mais distante da v1.

**Dependência do comandante** — **A pior dos cinco, e é eliminatória se o usuário der peso a esse critério.** Sem a Feather em campo: as ~30 mágicas do pacote são truques de Limited que se usam uma vez e vão para o cemitério. Sem criatura em campo: metade delas **não é conjurável**. Se ela morrer duas vezes, o imposto (3→5→7) num deck com curva de 1–2 mana dói, e cada turno sem ela é um turno de cartas mortas. **É estruturalmente a mesma falha que reprovou a v1** — o tabuleiro inteiro do deck depende do comandante — só que com outra cara.

**As quatro dores**
- **2 (mão morta)** — **resolvida estruturalmente** (a mão se reabastece sozinha).
- **3 (sem respostas)** — **resolvida em parte**: proteção e truques instantâneos, sim; remoção de permanente alheio, não — as mágicas da Feather miram *criatura sua*.
- **1** — depende do wincon-tester (dano de combate concentrado).
- **4 (morre para wipe)** — **NÃO resolvida; agravada.** É a dor original do usuário, no pior formato possível.

**Aproveitamento do pool** — **Médio.** Da `lista.txt`, servem [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing), [**Feat of Resistance**](https://www.ligamagic.com.br/?view=cards/card&card=Feat+of+Resistance), [**Adamant Will**](https://www.ligamagic.com.br/?view=cards/card&card=Adamant+Will), [**Shoulder to Shoulder**](https://www.ligamagic.com.br/?view=cards/card&card=Shoulder+to+Shoulder) (Support 2 + compra uma carta — reciclável!), [**Smaug's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Smaug%27s+Fury), [**Duty Beyond Death**](https://www.ligamagic.com.br/?view=cards/card&card=Duty+Beyond+Death). Da caixa: [**Moment of Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Moment+of+Glory), [**Joust**](https://www.ligamagic.com.br/?view=cards/card&card=Joust), [**Ambitious Assault**](https://www.ligamagic.com.br/?view=cards/card&card=Ambitious+Assault). **As 18 queimas da caixa não servem** — miram criatura *do oponente*, não sua, e não disparam a Feather. O aproveitamento é menor do que a intuição sugere.
**As 5 mais relevantes**: [**Gods Willing**](https://www.ligamagic.com.br/?view=cards/card&card=Gods+Willing) · [**Feat of Resistance**](https://www.ligamagic.com.br/?view=cards/card&card=Feat+of+Resistance) · [**Shoulder to Shoulder**](https://www.ligamagic.com.br/?view=cards/card&card=Shoulder+to+Shoulder) · [**Adamant Will**](https://www.ligamagic.com.br/?view=cards/card&card=Adamant+Will) · [**Moment of Glory**](https://www.ligamagic.com.br/?view=cards/card&card=Moment+of+Glory).

**Viabilidade no teto** — **Excelente.** É o eixo mais barato dos cinco; nenhuma peça-chave é individualmente cara. [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm) e [**Sejiri Shelter**](https://www.ligamagic.com.br/?view=cards/card&card=Sejiri+Shelter+%2F%2F+Sejiri+Glacier)-likes: **a cotar**.

**Risco declarado** — **Um board wipe encerra o jogo.** Você perde a Feather e todos os alvos ao mesmo tempo, e a mão vira papel. Em mesa com 3 oponentes, isso acontece.

---

### 3. Hofri Ghostforge — {3}{R}{W} — RW

- **Tipo**: Legendary Creature — Dwarf Cleric · 4/5 · CMC 5
- **Texto (oracle)**:
  > Spirits you control get +1/+1 and have trample and haste.
  > Whenever another nontoken creature you control dies, exile it. If you do, create a token that's a copy of that creature, except it's a Spirit in addition to its other types and it has "When this token leaves the battlefield, return the exiled card to its owner's graveyard."

**Ficha F1–F7**

- **F1** — (i) *anthem condicional*: Spirits +1/+1 com **trample e haste** — e todo token que ele cria é Spirit, então o anthem se aplica a tudo que ele produz; (ii) o gatilho é em **criatura não-token sua que morre**: exila e devolve como **cópia**, o que **redispara ETBs** e devolve o corpo com haste, pronto para atacar no mesmo turno.
- **F2** — 4/5. Corpo sólido, sobrevive a queima pequena.
- **F3** — Dwarf Cleric; Spirits como tipo interno criado por ele próprio.
- **F4** — recebe anthems e proteção; é o alvo nº 1 de remoção da mesa.
- **F5** — **facilita**: dá haste e trample ao que ele fabrica, o que converte morte em dano imediato.
- **F6** — CMC 5. Liga tarde, e é o pior ponto dele.
- **F7** — **anti-sinergia clara com token**: só funciona com criaturas **não-token**. Um deck de fichas (a v1) seria literalmente incompatível — o que é bom para a exigência de eixo novo, e ruim para o aproveitamento do pool.

**Como o deck joga**
- **T1–T3**: rampa e criaturas de ETB valioso (Solemn-likes, tutores de terreno, remoção-em-corpo).
- **T4–T6**: Hofri no T5. A partir daí, **cada criatura sua que morre volta maior, com trample e haste**: bloco favorável é lucro, remoção do oponente é lucro, e um board wipe do oponente devolve **todo** o seu tabuleiro como Spirits +1/+1 com haste — atacando no turno seguinte.
- **Fecho**: combate com um time que não acaba, mais um outlet de sacrifício para converter as mortes em valor a pedido.

**Por que é competitivo no Commander 200** — É a **única resposta estrutural à dor 4** entre os cinco: seu deck *lucra* com o wipe que os outros temem. E, como cada criatura vale duas (o corpo original mais a cópia Spirit maior), cada real gasto em criatura rende o dobro — que é exatamente o tipo de alavanca que um teto de R$ 200 pede.

**Dependência do comandante** — **Alta, mas não terminal.** As criaturas do deck são escolhidas por ETB e por corpo próprio; sem o Hofri, o deck é um midrange de valor que funciona, só sem o multiplicador. Se ele morrer duas vezes, o imposto 5→7→9 é pesado num deck sem ramp verde — **este é o ponto fraco real**, e é o mesmo problema de curva que o usuário já reclamou no Otharri (CMC 5).

**As quatro dores**
- **4 (morre para wipe)** — **resolvida estruturalmente e em grau máximo.**
- **1 (não fecha)** — **resolvida em boa parte**: trample + haste em corpos que crescem é conversão de board em dano.
- **2 (mão morta)** — atenuada (as criaturas de ETB compram), mas fica para o `draw-specialist`.
- **3 (sem respostas)** — não resolvida pelo eixo; fica para o `interaction-specialist`.

**Aproveitamento do pool** — **Baixo, e é o custo declarado deste eixo.** Das 18 criaturas RW-legais da caixa, **três** têm ETB e nenhuma justifica uma cópia Spirit ([**Redcap Thief**](https://www.ligamagic.com.br/?view=cards/card&card=Redcap+Thief), [**Pilgrim's Eye**](https://www.ligamagic.com.br/?view=cards/card&card=Pilgrim%27s+Eye), [**Phyrexian Revoker**](https://www.ligamagic.com.br/?view=cards/card&card=Phyrexian+Revoker)). As 69 não-básicas do Tori físico são criaturas de combate de Limited sem ETB — servem como corpo genérico, não como munição. **Risco de 20/09 confirmado por contagem.**
**As 5 mais relevantes**: [**Redcap Thief**](https://www.ligamagic.com.br/?view=cards/card&card=Redcap+Thief) · [**Pilgrim's Eye**](https://www.ligamagic.com.br/?view=cards/card&card=Pilgrim%27s+Eye) · [**Late to Dinner**](https://www.ligamagic.com.br/?view=cards/card&card=Late+to+Dinner) e [**Miraculous Recovery**](https://www.ligamagic.com.br/?view=cards/card&card=Miraculous+Recovery) (recursão que redispara o mesmo motor) · [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (wipe simétrico que é assimétrico a seu favor).

**Viabilidade no teto** — **Apertada, mas não estruturalmente inviável.** O eixo pede ~20 criaturas com ETB relevante, praticamente todas compradas. Não há peça individualmente caríssima e insubstituível — há **volume** de compra. [**Solemn Simulacrum**](https://www.ligamagic.com.br/?view=cards/card&card=Solemn+Simulacrum), `Dockside`-likes não existem em RW barato; alternativas incolores e brancas comuns cobrem. Todas **a cotar**.

**Risco declarado** — **Curva.** 5 de mana para ligar o motor, num deck que precisa ter criaturas em campo *antes*. Se o Hofri for removido no turno em que sai, você gastou o turno 5 em nada — e o imposto sobe. É o mesmo padrão de lentidão que o usuário reprovou no Otharri.

---

### 4. Quintorius, Loremaster — {3}{R}{W} — RW

- **Tipo**: Legendary Creature — Elephant Cleric · 3/5 · CMC 5
- **Texto (oracle)**:
  > Vigilance
  > At the beginning of your end step, exile target noncreature, nonland card from your graveyard. Create a 3/2 red and white Spirit creature token.
  > {1}{R}{W}, {T}, Sacrifice a Spirit: Choose target card exiled with Quintorius. You may cast that card this turn without paying its mana cost. If that spell would be put into a graveyard, put it on the bottom of its owner's library instead.

**Ficha F1–F7**

- **F1** — (i) *vigilância*: ataca e ainda bloqueia, e continua desvirado para a habilidade ativada — **as duas coisas no mesmo turno**; (ii) *end step*: exila uma mágica do seu cemitério **e** cria um 3/2 — o token é combustível, não plano; (iii) a ativada troca um Spirit por **conjurar de graça** qualquer mágica exilada por ele, e a mágica vai para o **fundo da biblioteca** em vez do cemitério (não recicla infinitamente — corta o loop).
- **F2** — 3/5 com vigilância; sobrevive à maior parte da queima e é ele mesmo o outlet de tap.
- **F3** — Elephant Cleric; os Spirits que ele cria são recurso interno.
- **F4** — recebe anthems e proteção.
- **F5** — dá **free-cast**, que é a forma mais cara de vantagem que existe.
- **F6** — CMC 5, e a ativada só gira **um turno depois** do primeiro exílio. O motor completo só roda no T6–T7.
- **F7** — a ativada usa `{T}` do próprio Quintorius: **atacar com ele e ativar no mesmo turno só é possível graças à vigilância** — leia-se, a vigilância é peça estrutural, não enfeite. Compete por `{1}{R}{W}` com as suas próprias mágicas.

**Como o deck joga**
- **T1–T3**: rampa e remoção barata, gastando mágicas de propósito — cada mágica usada é munição futura.
- **T4–T6**: Quintorius no T5; end step exila a sua melhor mágica já usada e te dá um 3/2. A partir do T6, **um free-cast por turno**: o [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) que você já jogou, a remoção grande, a queima que fechou uma ameaça.
- **Fecho**: por acúmulo — a mesa fica sem recursos e você reconjura o seu melhor efeito todo turno, com 3/2 sobrando para pressionar.

**Por que é competitivo no Commander 200** — Ele **desfaz a diferença entre uma carta de R$ 2 e uma carta de R$ 60**: num deck barato, a limitação não é ter *uma* boa mágica, é ter *muitas*. Quintorius transforma a sua melhor mágica barata em uma carta por turno, de graça. É a maior inevitabilidade por real gasto dos cinco candidatos, e reaproveita o perfil real da caixa (34 mágicas).

**Dependência do comandante** — **Alta.** O deck sem ele é "um monte de remoção", que funciona, mas sem motor. Imposto 5→7→9. Mitigação parcial: o corpo 3/5 com vigilância não morre para queima de 3.

**As quatro dores**
- **2 (mão morta)** — **resolvida estruturalmente**: o cemitério vira segunda mão.
- **1 (não fecha)** — parcialmente: free-cast repetido de queima/wipe é inevitabilidade, mas lenta.
- **3 (sem respostas)** — atenuada (você reusa a mesma remoção), mas o pacote inicial fica para o `interaction-specialist`.
- **4 (morre para wipe)** — parcialmente: os Spirits morrem, mas as mágicas exiladas **não estão no campo** e sobrevivem ao wipe. Você reconstrói mais rápido que a mesa.

**Aproveitamento do pool** — **Médio-alto.** As 34 mágicas RW-legais da caixa são munição direta, mais [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) e as remoções da `lista.txt`. Ressalva honesta: munição reciclada de qualidade de Limited continua sendo de Limited — free-cast de [**Radiating Lightning**](https://www.ligamagic.com.br/?view=cards/card&card=Radiating+Lightning) não ganha jogo.
**As 5 mais relevantes**: [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) (free-cast de wipe todo turno é brutal) · [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) (Spree modal, destrói artefato **e** encantamento) · [**Chandra's Outrage**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Outrage) · [**Searing Barrage**](https://www.ligamagic.com.br/?view=cards/card&card=Searing+Barrage) · [**Thrill of Possibility**](https://www.ligamagic.com.br/?view=cards/card&card=Thrill+of+Possibility) (enche o cemitério que ele consome).

**Viabilidade no teto** — **Boa.** Nenhuma peça-chave cara: o eixo pede remoção e wipes baratos, que é o que RW tem de mais barato. [**Wrath of God**](https://www.ligamagic.com.br/?view=cards/card&card=Wrath+of+God), [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act): **a cotar**.

**Risco declarado** — **Velocidade.** O motor completo roda no T6–T7; contra uma mesa rápida você perde antes de ligar. E ele morre para ódio de cemitério de forma **terminal** (diferente do Phlage, que ainda tem a zona de comando).

---

### 5. Gisela, Blade of Goldnight — {4}{R}{W}{W} — RW

- **Tipo**: Legendary Creature — Angel · 5/5 · CMC 7
- **Texto (oracle)**:
  > Flying, first strike
  > If a source would deal damage to an opponent or a permanent an opponent controls, that source deals double that damage to that player or permanent instead.
  > If a source would deal damage to you or a permanent you control, prevent half that damage, rounded up.

**Ficha F1–F7**

- **F1** — (i) voar + iniciativa: 5/5 que não morre em combate nem passa; (ii) **dobra o dano de qualquer fonte** contra oponentes e permanentes deles — inclui suas queimas, seu combate, e também **wipes de dano** ([**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) vira 26); (iii) **previne metade do dano** que você e seus permanentes recebem, arredondando para cima — um [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) seu mata o board da mesa e deixa metade do seu vivo.
- **F2** — 5/5 voador com iniciativa. Fecha jogo sozinha em pod que já tomou dano.
- **F3** — Angel; sem tribo relevante barata.
- **F4** — recebe anthems; qualquer +1/+0 nela vale o dobro no dano.
- **F5** — **facilita em massa**: é multiplicador de **todo** o resto do deck, sem exigir construção específica.
- **F6** — **CMC 7**. É o eixo em que o comandante é o *último* passo, não o primeiro.
- **F7** — a metade-prevenida é **simétrica ao seu favor**, sem atrito interno. O atrito é só de mana: 7 de custo com imposto 7→9→11.

**Como o deck joga**
- **T1–T3**: rampa agressiva (rochas) e remoção. Você joga como controle puro.
- **T4–T6**: wipes e remoção pontual, segurando a mesa. Uma ou duas ameaças voadoras intermediárias. Rampa acumulada para chegar aos 7.
- **Fecho**: Gisela entra e a matemática da mesa muda de uma vez. Seu [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act) limpa a mesa e sobra metade do seu board; qualquer queima de 3 vira 6; ela mesma bate 10 no ar por combate. **Dois ataques dela matam um jogador de 40.**

**Por que é competitivo no Commander 200** — Porque é um **multiplicador universal que não custa slot**. Ela está de graça na zona de comando (fora da régua) e dobra o output de **todas** as 99 cartas. Um deck de R$ 200 não consegue comprar cartas melhores — mas consegue dobrar as que tem. É o uso mais literal da regra "comandante caro libera orçamento".

**Dependência do comandante** — **Média, e é o candidato com o perfil mais saudável nesse eixo.** Por construção, o deck **tem de funcionar sozinho até o T7** — é um deck de controle completo antes de ela existir. Se ela morrer duas vezes, o deck continua sendo um deck de controle; o que se perde é o fecho rápido. **Mas** o imposto 7→9→11 significa, na prática, que você a joga **uma ou duas vezes por partida**, não mais.

**As quatro dores**
- **1 (não fecha o jogo)** — **resolvida estruturalmente e melhor que qualquer outro candidato.** Dano dobrado é o fecho.
- **4 (morre para wipe)** — **resolvida em parte, por um caminho diferente**: você tem poucas criaturas a perder, e a prevenção de metade faz você **sobreviver a wipes de dano** que matam a mesa.
- **2 e 3** — inteiramente nas mãos do `draw-specialist` e do `interaction-specialist`. O eixo não ajuda.

**Aproveitamento do pool** — **Alto**, e pelo mesmo motivo do Phlage: as 18 queimas/remoções da caixa são munição direta, e **com o dobro do dano** várias delas saem do "fraco para Commander" (uma [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) de 6, um [**Seismic Wave**](https://www.ligamagic.com.br/?view=cards/card&card=Seismic+Wave) de 4+2). [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) e as remoções da `lista.txt` entram inteiras.
**As 5 mais relevantes**: [**Fumigate**](https://www.ligamagic.com.br/?view=cards/card&card=Fumigate) · [**Lightning Strike**](https://www.ligamagic.com.br/?view=cards/card&card=Lightning+Strike) · [**Chandra's Outrage**](https://www.ligamagic.com.br/?view=cards/card&card=Chandra%27s+Outrage) (8 na criatura, 4 no dono) · [**Seismic Wave**](https://www.ligamagic.com.br/?view=cards/card&card=Seismic+Wave) · [**Syr Carah, the Bold**](https://www.ligamagic.com.br/?view=cards/card&card=Syr+Carah%2C+the+Bold) (o ping de 1 vira 2, e cada mágica que acerta jogador compra).

**Viabilidade no teto** — **Boa, com uma exigência clara**: o deck precisa de **ramp acima da média** para chegar aos 7 de mana com o deck ainda vivo, e ramp em RW é rocha (barata). Nenhuma peça estruturalmente cara. A cotar: [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act), [**Wrath of God**](https://www.ligamagic.com.br/?view=cards/card&card=Wrath+of+God), [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation).

**Risco declarado** — **Os 7 de mana.** Se a rampa não vier, ela nunca chega; se ela chegar e for removida, o segundo lançamento custa 9. Um deck de R$ 200 não tem tutores nem proteção de sobra para garantir isso. Este é o candidato com a maior variância dos cinco.

---

## 4. Banco de reservas (analisados, não finalistas)

Todos com oracle puxado, listados caso o usuário queira um sexto ângulo:

- **[**Taii Wakeen, Perfect Shot**](https://www.ligamagic.com.br/?view=cards/card&card=Taii+Wakeen%2C+Perfect+Shot)** {R}{W} 2/3 — *"cada queima sua compra uma carta"*: compra quando uma fonte sua causa dano não-combate igual à resistência da criatura, e `{X}, {T}` aumenta esse dano em X. Vira o eixo de queima da caixa em motor de cartas e mata qualquer coisa. **Comandante de 2 mana = imposto mínimo.** Fora da lista só porque o fecho é pouco claro e o eixo se sobrepõe ao do Phlage.
- **[**Bell Borca, Spectral Sergeant**](https://www.ligamagic.com.br/?view=cards/card&card=Bell+Borca%2C+Spectral+Sergeant)** {2}{R}{W} */5 — impulse grátis todo upkeep, sem atacar, sem condição. Resolve a dor 2 de graça. Fora porque não fecha jogo nenhum.
- **[**Brion Stoutarm**](https://www.ligamagic.com.br/?view=cards/card&card=Brion+Stoutarm)** {2}{R}{W} 4/4 lifelink — arremessa criaturas na cara, ignorando bloqueadores. Eixo genuinamente distinto, mas exige criaturas grandes, que é o que RW não tem barato.
- **[**Lorehold, the Historian**](https://www.ligamagic.com.br/?view=cards/card&card=Lorehold%2C+the+Historian)** {3}{R}{W} 5/5 — todo instantâneo/feitiço na mão ganha **miracle {2}**. Conjurar wipe por {2} é espetacular; depender de ser a primeira carta comprada no turno é aleatório demais para um deck que precisa parecer competitivo.
- **[**Nelly Borca, Impulsive Accuser**](https://www.ligamagic.com.br/?view=cards/card&card=Nelly+Borca%2C+Impulsive+Accuser)** {2}{R}{W} — eixo político de goad: a mesa briga entre si. "Jeito de jogar" radicalmente diferente; fora porque não fecha e porque o suporte a goad mora em preto.
- **[**Aurelia, the Warleader**](https://www.ligamagic.com.br/?view=cards/card&card=Aurelia%2C+the+Warleader)** — combate extra. **Fora por eixo, não por preço**: mesmo com o comandante isento da régua, combate extra com exército é a pegada Krenko/Otharri que o usuário reprovou. Registro da decisão de 2026-09-19 (regra 5) mantido: o motivo original — imposto 6→8→10 — continua válido e agora soma-se o veto de eixo.

---

## 5. Recomendação (a decisão é do usuário — regra 9)

**Minha recomendação é [**Phlage, Titan of Fire's Fury**](https://www.ligamagic.com.br/?view=cards/card&card=Phlage%2C+Titan+of+Fire%27s+Fury)**, e o argumento é um só: ele é o único
candidato que **não repete a falha que reprovou a v1**.

O diagnóstico do teste de mesa foi "o tabuleiro inteiro do deck depende do comandante". Dos cinco
finalistas, quatro reproduzem esse padrão em algum grau — a Feather de forma aguda (sem ela, 30
cartas mortas), o Hofri e o Quintorius de forma clássica (motor de 5 mana com imposto crescente),
a Gisela de forma invertida mas cara (7 de mana, uma ou duas conjuras por jogo). O Phlage é o
**único comandante de RW que se reconjura sem imposto**, do cemitério, pelo mesmo custo fixo na
quinta vez que morrer — `bin/mtgdb rulings` confirma. Some-se a isso: CMC 3 (mata a queixa de
lentidão do Otharri), corpo que *é* remoção (economiza slots num teto de R$ 200), eixo de controle
que resolve as dores 3 e 4 **estruturalmente**, e o maior aproveitamento da caixa entre os cinco.

**Se o critério do usuário for "jeito de jogar mais diferente possível"**, a resposta é
[**Feather, the Redeemed**](https://www.ligamagic.com.br/?view=cards/card&card=Feather%2C+the+Redeemed) — jogar no turno dos outros, com a mão sempre cheia, é o oposto exato do
que ele fez nas duas listas anteriores. Mas ele precisa aceitar, de olhos abertos, que um board
wipe encerra o jogo dele.

**Se o critério for "quero um deck que fecha"**, a resposta é [**Gisela, Blade of Goldnight**](https://www.ligamagic.com.br/?view=cards/card&card=Gisela%2C+Blade+of+Goldnight) — e o
preço é chegar aos 7 de mana.

**O que o Phlage ainda deve ao deck e não resolve sozinho**: a dor 1 (fecho). Se ele for o
escolhido, o `wincon-tester` recebe como requisito explícito entregar **um fecho nomeado que não
seja "atacar com o comandante"**.

### Pendências desta fase
- Cotação LigaMagic dos comandantes: **não necessária** para decidir (fora da régua), mas nenhum
  dos 12 avaliados tem cotação registrada — todos **`a cotar`** se o usuário quiser saber o gasto.
- Peças-chave sem cotação, comuns a mais de um eixo: [**Blasphemous Act**](https://www.ligamagic.com.br/?view=cards/card&card=Blasphemous+Act), [**Wrath of God**](https://www.ligamagic.com.br/?view=cards/card&card=Wrath+of+God),
  [**Faithless Looting**](https://www.ligamagic.com.br/?view=cards/card&card=Faithless+Looting), [**Boros Charm**](https://www.ligamagic.com.br/?view=cards/card&card=Boros+Charm), [**Solemn Simulacrum**](https://www.ligamagic.com.br/?view=cards/card&card=Solemn+Simulacrum), [**Fiery Emancipation**](https://www.ligamagic.com.br/?view=cards/card&card=Fiery+Emancipation), [**Sunforger**](https://www.ligamagic.com.br/?view=cards/card&card=Sunforger) —
  **a cotar** antes de fechar as 99.
- [**Requisition Raid**](https://www.ligamagic.com.br/?view=cards/card&card=Requisition+Raid) (pendência aberta em 2026-09-20, adiada até o eixo estar resolvido):
  **sobrevive a qualquer um dos cinco eixos** como remoção modal de artefato/encantamento por
  `{2}{W}`, e já está na caixa. Decisão devolvida ao `interaction-specialist` da v2.
