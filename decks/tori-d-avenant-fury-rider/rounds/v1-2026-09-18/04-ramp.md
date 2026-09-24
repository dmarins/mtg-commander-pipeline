# Aceleração (Ramp) — v1 · Otharri, Suns' Glory

**Modo:** `improve` · **Comandante:** Otharri, Suns' Glory `{3}{R}{W}` (CMC 5) · **Identidade:** RW
**Passe 1:** 2026-09-19 (construção) · **Passe 2:** 2026-09-20 (meta de explosivo)
**Passe 3 (este):** **2026-09-23 — devolução após o teste em proxy**

> **Ramp padrão hoje: 11/10–11 · Explosivo: 2/2–3.** As metas estão batidas **no papel**.
> O que este passe mede é outra coisa: **as 11 fontes valem 9 na prática**, e a curva medida
> explica exatamente a frase do usuário.
>
> Todo texto oracle foi puxado nesta sessão via `bin/mtgdb` (regra 6). Preços: **LigaMagic
> (menor)**, via `bin/mtgdb prices`, com a data ao lado (regra 2). Nenhum número inventado.

---

# ⚠ O QUE MUDOU NESTE PASSE (2026-09-23)

O usuário imprimiu a v1 em proxy e jogou. Relato: *"senti lentidão para descer o comandante"*.
Nada foi comprado — não há custo afundado, e a lista pode mudar de graça.

Os passes 1 e 2 fecharam metas **de contagem**. Este passe **mede a curva** e conclui que a
contagem estava certa e a leitura estava errada. Quatro mudanças de veredito:

| # | O que os passes 1–2 disseram | O que a medição de hoje diz |
|---|---|---|
| 1 | "11 fontes de ramp padrão ✅" | **9 aceleram o comandante.** `Sword of the Animist` acelera em **0,58%** das partidas; `Loyal Warhound` traz terreno **virado** e é condicional |
| 2 | "rochas de mv2 põem o Otharri no T4" | **Verdade — e T4 sai em 42,6% das partidas.** A mediana é **T5**, e em **12,7%** ele não é conjurável nem no T7 |
| 3 | "`Sword of the Animist` é land ramp que sobrevive a wipe" | Custa **R$ 27,00 — a carta mais cara do deck inteiro** (16% do teto) e é o pior slot de ramp da lista pela métrica que o usuário reclamou |
| 4 | "36 terrenos, o deck está certo" (Fase 6) | A fórmula da própria Fase 6, **rodada sobre a lista final**, dá **36,7–37,6**, não 35,0. Os inputs dela eram da lista pré-corte de orçamento |

**Proposta:** 2 trocas dentro dos meus próprios slots, **zero slot novo**, que **devolvem
R$ 26,43** ao orçamento e levam Otharri-no-T4 de **42,6% → 47,9%** e Otharri-no-T3 de
**6,7% → 11,9%**. Com 1 ou 2 slots negociados, vai a **48,8%** / **52,3%** e a mediana cai para **T4**.

---

## 1. A medição — método e resultado

### 1.1 Método

Monte Carlo de **40.000 partidas por configuração**, metade na saída e metade no draw, sobre a
lista exata da §6 do `report.md`: **99 cartas, 36 terrenos** (14 Mountain, 17 Plains,
`Command Tower`, `Battlefield Forge`, `Furycalm Snarl`, `Rustvale Bridge`, `Blast Zone`).

Modelado carta a carta: custo, mana líquido, cor produzida, **entra virado**
(`Sphere of the Suns`, `Rustvale Bridge`), condicional (`Furycalm Snarl` desvira se houver básico
na mão), e o `Boros Signet`, que resolve `{R}` **e** `{W}` numa tapada só. Mulligan de Londres,
até 2, mantendo 2–5 terrenos. O critério de "conjurável" é o completo: **5 manas disponíveis
E fonte de `{R}` E fonte de `{W}` distintas**.

**O modelo é otimista de propósito**, e isso importa para ler os números:
- o piloto **sempre** prioriza ramp e **nunca** gasta mana em outra coisa;
- ninguém destrói suas rochas, ninguém contra-mágica;
- `Sphere of the Suns` é tratada como fonte **permanente** (na prática ela tem **3 cargas**).

Ou seja: **os números abaixo são o teto**. A experiência real do usuário é pior que isto.

### 1.2 O baseline

| Configuração | T3 | **T4** | **T5** | T6 | T7 | nunca até T7 | mediana |
|---|---|---|---|---|---|---|---|
| **v1 atual — 36 terrenos, 9 aceleradores reais** | 6,7% | **42,6%** | **71,9%** | 81,2% | 87,3% | **12,7%** | **T5** |

**A resposta direta às duas perguntas do prompt: Otharri sai no T4 em 42,6% das partidas e no
T5 em 71,9%. A mediana é T5.** E em **1 partida a cada 8** ele não é conjurável nem no turno 7 —
esse é o balde que produz a frase *"senti lentidão"*, porque é a partida que o jogador lembra.

### 1.3 Decomposição das falhas — onde exatamente o T4 se perde

Das 57,4% de partidas em que ele **não** sai no T4:

| Motivo | % de todas as partidas | % das falhas |
|---|---|---|
| Chegou ao T4 com **4 terrenos e nenhum acelerador em campo** | **29,2%** | **51%** |
| **Perdeu land drop** — menos de 4 terrenos no T4 | 22,4% | 39% |
| Tinha 5 manas mas **não conseguia produzir `{R}{W}`** | 5,2% | 9% |

No T5 o quadro vira: a falha dominante passa a ser **land drop perdido (23,5%)**, e a falha por
falta de rocha some (1,2%).

### 1.4 O que **não** é a causa — refutações medidas

O prompt listou candidatos sem veredito prévio. Três morreram na medição:

| Hipótese | Teste | Resultado |
|---|---|---|
| "Excesso de rochas de **mv 2** empurra o comandante para o T4 em vez do T3" | Verdade parcial, mas **irremediável**: para sair no T3 é preciso **+2 de mana líquido até o T2**, e só existem 3 cartas assim em RW dentro do teto — `Sol Ring` (já no deck), `Mana Vault` (US$ 120) e `Pentad Prism`. Duas rochas de mv2 **não** dão T3 (T2 rocha, T3 3 terrenos + 1 = 4) | Não é defeito de escolha, é **falta de uma carta específica** |
| "As rochas de **mv 3** (`Commander's Sphere`, `Boros Locket`) são lentas" | Troquei as duas por mv2 na simulação | **+0,3 pp no T4.** Estatisticamente nada. Uma rocha de mv3 jogada no T3 entrega T4 igual a uma de mv2 |
| "`{R}{W}` com só 4 não-básicos é o gargalo" | Troquei **todas** as rochas incolores por coloridas | **+0,5 pp no T4.** Com 18 fontes de `{R}` e 21 de `{W}`, cor é 9% do problema, não o problema |
| "`Everflowing Chalice` / `Sphere of the Suns` são mana que não acelera" | **Falso para as duas.** `Everflowing Chalice` por `{2}` = 1 mana líquido a partir do T3, igual ao `Mind Stone`. `Sphere of the Suns` entra virada mas ainda entrega o T4 (T2 Sphere → T4 = 4 terrenos + 1) | Ficam. Ressalva declarada na §4.5 |
| "`Oracle's Vault` é mana que não acelera" | **Verdade, mas ela não é ramp e nunca foi contada como tal.** É saque por impulso de mv **4** (não 3) que disputa o turno do comandante. Devolvida à Fase 3 na §8.1 | Fora da minha cota |

---

## 2. Diagnóstico — a causa dominante, em uma frase

> **A causa dominante é densidade de aceleradores: 51% das falhas de T4 são partidas em que o
> deck chegou ao turno 4 com os 4 terrenos na mesa e nenhuma rocha em campo.** A segunda causa,
> com 39%, é land drop perdido. Cor é 9%.

E a razão da densidade baixa **não é a contagem** — é que **duas das 11 fontes não aceleram**:

**`Sword of the Animist` (mv2, R$ 27,00).** Para ela pôr um terreno em campo a tempo de destravar
o T4, é preciso: criatura no T1, Sword no T2, equipar (`{2}`) no T3 e atacar. **O deck tem
exatamente uma criatura de mv1 em 99 cartas** — `Recruitment Officer`. A chance de abrir com as
duas nas 8 primeiras cartas é **56/9702 = 0,58%**. Em qualquer outra linha (Sword no T3, equipar
no T4) o terreno chega **virado no T4** e não acelera nada. **É um slot de ramp que, medido,
acelera em menos de 1% das partidas — e é a carta mais cara do deck.**

**`Loyal Warhound` (mv2, R$ 2,59).** Busca **Plains virada**, e só *se um oponente controlar mais
terrenos que você*. Ela não perde valor — é 3/1 com vigilância, recebe os anthems estáticos e serve à
reclamação de "poucas criaturas". Mas **como acelerador do T4 ela vale perto de zero**, porque o
terreno entra virado no mesmo turno em que foi buscado. **Não proponho cortá-la**: ela paga o
slot pelo corpo, e o corpo é o que a Fase 2 está sendo chamada a resolver.

**Aceleradores reais hoje: 9.** `Sol Ring`, `Arcane Signet`, `Boros Signet`,
`Talisman of Conviction`, `Mind Stone`, `Sphere of the Suns`, `Everflowing Chalice`,
`Commander's Sphere`, `Boros Locket`.

### 2.1 A causa secundária: os 36 terrenos não são o que a Fase 6 calculou

A Fase 6 rodou a própria fórmula e chegou a 35,0–35,6 → 36. **Rodei a mesma fórmula sobre a lista
que de fato existe** (§5/§6 do `report.md`, 63 não-terrenos, contados um a um nesta sessão):

| Inputs | CMC médio | N (ramp/draw mv≤2) | Fórmula | Terrenos |
|---|---|---|---|---|
| **O que a Fase 6 declarou** | 2,667 | 17 | 31,42 + 3,13·2,667 − 0,28·17 | **35,01** |
| **A lista real da v1** | **2,857** | 13 | 31,42 + 3,13·2,857 − 0,28·13 | **36,72** |
| **A lista real da v1** | **2,857** | 10 (só aceleradores) | idem, N=10 | **37,56** |

**Por que os inputs divergem.** A soma de CMC dos 63 não-terrenos é **180**, não 168 — média
**2,857**, não 2,667 (conferida carta a carta no script desta sessão). E dos 17 itens que a Fase 6
contou como N, **quatro foram cortados por preço no fechamento e nunca repostos**:
`Fellwar Stone` (R$ 14,50), `Knight of the White Orchid` (R$ 10,99), `Idol of Oblivion` (R$ 8,94)
e `Great Train Heist` (R$ 22,80). A conta foi feita sobre a lista **antes** do corte de orçamento.

**Não é erro de método — é input desatualizado.** Pela própria régua da Fase 6, o alvo é **37**.
Isso não é um corte meu: é o achado que eu devolvo ao orquestrador para negociar 1 slot (§5).

---

## 3. Varredura da caixa — regra 7, refeita hoje

`bin/mtgdb collection -list` rodado nesta sessão: **116 sobressalentes** (eram 113 no passe 1).
Conferi as 116 e puxei oracle das que eu não conhecia (`Ori, Keeper of Songs`,
`Mudbutton Clanger`, `Wilderland Scrounger`, `Zookeeper Mechan`, `Bargaining Table`,
`Wedgelight Rammer`, `Steel Wrecking Ball`, `Thopter Fabricator`, `Lesser Masticore`,
`Leonin Abunas`, `Mycosynth Wellspring`, `Titan Forge`, `Culling Dais`, `Lux Artillery`,
`Melded Moxite`, `Moxite Refinery`, `Thaumaton Torpedo`, `Diversion Unit`, `Cargo Ship`).

**Nenhuma das 3 cartas novas produz mana.** A lista de produtores de mana da caixa é a mesma:

| Sobressalente | Custo / o que gera | Veredito hoje | Por escrito |
|---|---|---|---|
| **Myr Convert** | `{2}` 2/1 art. creature · `{T}` + 2 vida: 1 mana de **qualquer cor** | **ENTRA** (§4.4) | mv2 · **artefato** (metalcraft) · **criatura** que recebe os anthems estáticos (`Warleader's Call`, `Wedding Festivity`, `Jor Kadeen`) — o único produtor de mana da caixa que também responde à outra reclamação do usuário |
| **Hedron Crawler** | `{2}` 0/1 art. creature · `{T}`: `{C}` | **reserva alta** — entra no Nível 2 (§5) | mesma curva do Myr, mas só `{C}` e corpo 0/1. Vale +3,5 pp no T4 se houver slot |
| **Boros Locket** | `{3}` · `{T}`: `{R}`/`{W}` · 4 manas + sac: compra 2 | **SAI** (§4.3) | está no deck; dominado pelo `Commander's Sphere` no mesmo mv3 |
| **Commander's Sphere** | `{3}` · 1 mana + sac: compra 1 | **fica** | já no deck |
| **Sphere of the Suns** | `{2}` virada, 3 cargas | **fica** | já no deck; ressalva na §4.5 |
| **Manalith** | `{3}` · 1 mana de qualquer cor | **não** | **cortada por mim em 2026-09-20** (passe 2, §D.3). O motivo — terceira cópia do mesmo efeito a mv3 — continua válido. **Não reproponho** (regra 5) |
| **Seer's Lantern** `{3}` · **Magnifying Glass** `{3}` | 1 `{C}` a mv3 | **não** | pior taxa da caixa; perdem para o `Hedron Crawler`, que é mv2 e ainda é corpo |
| **Prophetic Prism** `{2}` · **Pilgrim's Eye** `{3}` | filtro / terreno **para a mão** | **não são ramp** | mana líquido 0 e negativo. Seguem elegíveis por outras funções (Fases 2/3) |
| **Moss Diamond** | `{2}` virada · `{G}` | **ilegal** | identidade G, fora de RW |
| **Sol Ring** (2ª cópia) | — | **inutilizável** | singleton |
| Demais 106 | conferidas | **não** | nenhuma adiciona mana, reduz custo ou busca terreno |

**Saldo da caixa: `Myr Convert` entra de graça, `Hedron Crawler` fica de prontidão.** A única
compra que proponho é justificada nominalmente contra elas na §4.2.

### 3.1 As 69 não-básicas de `lista.txt`
Reconferidas: produzem mana **`Sol Ring`** e **`Arcane Signet`**, só. Inalterado desde o passe 1.

---

## 4. As trocas — ficha F1–F7 completa (regras 4 e 5)

**Consulta ao `decisions.md` (regra 5), feita nesta sessão.** O registro tem duas linhas, ambas de
2026-09-19 e ambas sobre o comandante. **Nenhuma** das quatro cartas abaixo tem histórico de corte
e retorno. A única carta com histórico na minha categoria é `Manalith` (cortada por mim em
2026-09-20) — e ela **não** volta.

### 4.1 SAI · `Sword of the Animist` — `{2}` Legendary Artifact — Equipment · **R$ 27,00**

*(LigaMagic menor, 2026-09-19 — 4 dias)*

| Eixo | Leitura (oracle puxado nesta sessão) |
|---|---|
| **F1** | Criatura equipada recebe **+1/+1** · Sempre que a equipada **atacar**, procure um **terreno básico** e ponha-o em jogo **virado** · **Equip `{2}`** |
| **F2** | **Sem corpo.** Não pode ser tapada para custo algum — o deck não tem crew, station, convoke, improvise nem exert |
| **F3** | **É artefato** → metalcraft do `Jor Kadeen` (3+). Também **lendária**, mas nada no deck lê "legendary matters" |
| **F4** | **Não recebe nada** — não é criatura, não recebe contador, anthem, nem gatilho do comandante |
| **F5** | **Dá** +1/+1 a **uma** criatura e land ramp condicionado a ataque |
| **F6** | mv2 no papel; **4 manas reais** (2 + equip 2) antes de fazer qualquer coisa, e ainda exige uma criatura viva que ataque. **Medido: a linha que destrava o T4 (criatura no T1 + Sword no T2 + equip no T3) sai em 0,58% das partidas**, porque o deck tem **uma única criatura de mv1** em 99 cartas |
| **F7** | Três atritos, todos contra a queixa do usuário: (i) o terreno entra **virado**, então nunca acelera o turno em que aparece; (ii) o equip de `{2}` **compete pelo mana do turno em que se quer conjurar o Otharri**; (iii) depois de um board wipe ela é **carta morta** no topdeck — não há o que equipar, que é exatamente o cenário da dor 4 que ela foi comprada para resolver |

**Protocolo de corte (§2 do checklist).** Funções = [land ramp repetível, +1/+1 a um portador,
contagem de artefato, equip barato… **não, equip `{2}` não é barato**]:

- **Land ramp repetível** → **parcialmente descoberta, custo declarado**. Sobra `Loyal Warhound`
  (um básico, uma vez, condicional) e o 37º terreno do Nível 1. **O deck perde sua única fonte
  *repetível* de terreno.** Aceito o custo porque a função entregava terreno **virado** e
  **gated em combate**, ou seja, nunca resolvia o problema medido — e porque a alternativa que a
  cobriria (`Knight of the White Orchid`, terreno **desvirado**) já foi descartada a R$ 10,99.
- **+1/+1 no portador** → coberta por `Ancestral Blade` (equip `{2}`, e ainda traz um corpo 1/1),
  `Hexgold Halberd`, e pelos anthems que alcançam qualquer criatura sua — `Warleader's Call` e
  `Wedding Festivity` (+1/+1 estáticos), `Jor Kadeen` (+3/+0 com metalcraft) e `Valor in Akros`
  (+1/+1 até o fim do turno a cada criatura que entra). *(`Radiant Destiny` só pega o tipo
  escolhido e `Intangible Virtue` só pega fichas — oracle conferido nesta sessão.)*
- **Contagem de artefato (metalcraft)** → **coberta, contagem inalterada**. Saem 2 artefatos
  (Sword, Locket), entram 2 artefatos (Pentad Prism, Myr Convert). O deck segue com **17
  artefatos + `Rustvale Bridge`** = 18 permanentes artefato. `Jor Kadeen` precisa de 3.
- **Lendária** → nenhuma carta do deck se importa. Dispensável.

**Veredito: corte limpo, com uma função parcialmente descoberta e declarada.** E é o corte
**dentro da minha especialidade**: `Sword of the Animist` não exerce função de tema, saque,
interação, wincon ou manabase — é um acelerador que não acelera.

### 4.2 ENTRA · `Pentad Prism` — `{2}` Artifact · **R$ 0,57** *(LigaMagic menor, 2026-09-18 — 5 dias)*

> **Sunburst** (entra com um marcador de carga para cada cor de mana gasta para conjurá-la)
> **Remova um marcador de carga: adicione um mana de qualquer cor.**

| Eixo | Leitura |
|---|---|
| **F1** | Conjurada com Mountain + Plains, entra com **2 cargas** = **2 manas de qualquer cor**, disponíveis **no turno seguinte, de uma vez** |
| **F2** | Sem corpo |
| **F3** | **É artefato** → metalcraft. **E continua artefato depois de esvaziar**, exatamente como a `Sphere of the Suns` — metalcraft permanente por 2 manas |
| **F4** | **Recebe marcadores de carga** — e o deck tem `Karn's Bastion` na fila de reservas de terreno da Fase 6 (proliferate). Se ela entrar, o Prism vira recarregável |
| **F5** | Dá **mana de qualquer cor** — resolve `{R}`, `{W}` e o genérico |
| **F6** | **mv2, e é a única carta dentro do teto que produz o salto de +2.** É o que separa o T3 do T4 |
| **F7** | **Atrito declarado: as cargas acabam.** São 2 manas no total da partida, não por turno. É um ritual guardado, não uma rocha |

**Por que ela, com 2+ pontos de sinergia (regra 3):**
1. **É a única carta afordável do formato que põe o Otharri no T3.** Para sair no T3 é preciso
   +2 de mana líquido até o T2. Em RW, dentro do teto, existem exatamente três cartas: `Sol Ring`
   (já no deck), `Mana Vault` (US$ 120, fora) e `Pentad Prism` (R$ 0,57).
   **Medido: T3 vai de 6,7% para 11,9% — quase dobra** — e T4 de 42,6% para 47,7%.
2. **As 2 cargas casam com o imposto do comandante.** Otharri sai a 5 → 7 → 9. Guardar as duas
   cargas para o turno do recast paga exatamente a diferença entre 5 e 7.
3. **Artefato permanente** (metalcraft do `Jor Kadeen`) que **sobrevive ao board wipe**, que era
   o critério declarado no passe 1 (§0.2) para preferir rocha a dork.

**Regra 7 — por que compra, tendo equivalente na caixa.** Nomeando as equivalentes:
`Manalith`, `Seer's Lantern`, `Magnifying Glass` e `Boros Locket` (todas da caixa) produzem
**1 mana por tap**. **Nenhuma carta da caixa produz 2 manas de uma vez**, e é exatamente isso que
compra o turno 3. `Hedron Crawler` e `Myr Convert` (caixa, mv2) produzem 1 e **não** dão T3 —
medido: com `Hedron Crawler` no lugar do Prism, T3 fica em 6,8% em vez de 11,9%.

**Regra 7, segunda parte — a cópia física não está disponível.** Existe um `Pentad Prism` no
`decks/inspirit-flagship-vessel/`, que está **montado**. Deck montado é coleção fechada: a carta
não migra. **Compra-se a segunda cópia, R$ 0,57.** Declarado.

### 4.3 SAI · `Boros Locket` — `{3}` Artifact · **caixa, R$ 0,00**

| Eixo | Leitura (oracle puxado nesta sessão) |
|---|---|
| **F1** | `{T}`: `{R}` ou `{W}` · `{R/W}{R/W}{R/W}{R/W}`, `{T}`, sacrifique: **compre 2 cartas** |
| **F2** | Sem corpo. Não pode ser tapada para nenhum custo além da própria habilidade |
| **F3** | **É artefato** → metalcraft |
| **F4** | **Não recebe nada** |
| **F5** | **Fixação de cor** — `{R}` ou `{W}` |
| **F6** | **mv3** — o turno mais disputado do deck (`Radiant Destiny`, `Warleader's Call`, `Goblin War Drums`, `Basri Ket`, `Combat Celebrant`, `Molten Gatekeeper`, `Tocasia's Welcome`, `Unbreakable Formation`) |
| **F7** | **Dominância declarada:** `Commander's Sphere` custa o **mesmo mv3**, produz o **mesmo mana** (num deck bicolor "qualquer cor" e "qualquer cor da identidade" são idênticos) e saca 1 carta **de graça** ao sacrificar. O saque do Locket custa **4 manas + tap + sac** — 5 de investimento num turno que o deck sempre quer usar para desenvolver |

**Protocolo de corte.** Funções = [1 mana R/W a mv3, contagem de artefato, saída tardia de 2 cartas]:
- **Produção de mana + fixação de cor** → coberta por `Myr Convert` (**qualquer cor**, e a mv**2**),
  mais 8 outras fontes. Fontes de `{W}` e `{R}` permanecem 21 e 18.
- **Contagem de artefato** → coberta: contagem final **inalterada** (saem 2 artefatos, entram 2).
- **Saída tardia de 2 cartas** → **parcialmente descoberta, custo declarado**. Sobram
  `Commander's Sphere` (sac → 1 carta) e `Mind Stone` (`{1}`, sac → 1 carta), além das 13 fontes
  de saque da §1 do `report.md`. O Locket era a única rocha que sacava **duas**, mas ao custo de
  5 manas — a pior taxa por carta do pacote de saque. Aceito.

**Veredito: corte limpo, uma função parcialmente descoberta e declarada.**

**Honestidade sobre a magnitude (§3 do checklist — mesma régua dos dois lados):** medido isolado,
`Boros Locket` → `Myr Convert` vale **+0,3 pp no T4**. **Esta troca não é de velocidade.** Ela se
justifica por três eixos que não são o meu, e eu declaro isso em vez de inflar o número: o
**corpo** (a Fase 2 está sendo chamada a resolver "poucas criaturas" — esta troca entrega +1
criatura de graça), o **receptor de anthems** (um 2/1 vira 4/3 com `Warleader's Call` + `Wedding Festivity`,
e 7/3 com o metalcraft do `Jor Kadeen` ligado — enquanto o Locket é inerte) e a **curva** (mv3 → mv2 num turno superlotado).

### 4.4 ENTRA · `Myr Convert` — `{2}` Artifact Creature — Phyrexian Myr 2/1 · **caixa, R$ 0,00**

| Eixo | Leitura (oracle puxado nesta sessão) |
|---|---|
| **F1** | **Toxic 1** · `{T}`, **pague 2 vidas**: adicione 1 mana de **qualquer cor** |
| **F2** | **Corpo 2/1** — ataca, bloqueia, e pode ser tapado (a própria habilidade de mana) |
| **F3** | **Artifact Creature** → conta para **metalcraft** do `Jor Kadeen` **e** para a contagem de criaturas do deck |
| **F4** | **Recebe os anthems que não são condicionais de tipo** (oracle dos 6 conferido nesta sessão): `Warleader's Call` **+1/+1** e `Wedding Festivity` **+1/+1** (estáticos) → **4/3**; `Jor Kadeen` **+3/+0** com metalcraft → **7/3**; `Valor in Akros` +1/+1 até o fim do turno a cada criatura que entra. **Não recebe** `Radiant Destiny` (só o tipo escolhido, que é Rebel) nem `Intangible Virtue` (só fichas) |
| **F5** | Dá **qualquer cor** — cobre `{R}`, `{W}` e genérico |
| **F6** | **mv2** — entra no turno que a curva pede |
| **F7** | **Dois atritos declarados:** (i) **morre no board wipe**, ao contrário de uma rocha — e o passe 1 (§0.2) usou exatamente esse critério para rebaixá-la a reserva; (ii) **2 vidas por mana**, e o corpo disputa o próprio `{T}` entre atacar e produzir mana |

**Regra 5 — o que mudou desde que eu a rebaixei.** Em 2026-09-19 eu escrevi que `Myr Convert` cai
para reserva porque "morre no wipe, que é exatamente o turno em que o mana é necessário". Três
coisas mudaram, e nenhuma delas é "reavaliei e agora gostei":
1. **O usuário jogou.** A dor 4 (board wipe) era hipótese de intake; **a lentidão do comandante é
   observação de mesa**. A prioridade entre as duas mudou por dado, não por opinião.
2. **A troca é contra o `Boros Locket`, não contra uma rocha que fica.** A contagem de artefatos
   permanece 17 + `Rustvale Bridge`, então a resiliência a wipe do pacote de mana **não cai**:
   continuam 8 artefatos de mana em campo depois de um wipe de criaturas.
3. **A Fase 2 foi reaberta pedindo criaturas.** Em 2026-09-19 o corpo do Myr era ruído; hoje é
   parte da resposta a uma queixa declarada. Ele entra **sem consumir slot da Fase 2**.

### 4.5 As que **ficam**, com a ressalva escrita

- **`Everflowing Chalice`** (R$ 0,99) — conjurada por `{2}` entrega 1 mana líquido a partir do T3,
  taxa idêntica à do `Mind Stone`. É a única rocha do deck que **escala no late** (multikicker).
  **Fica.** É a primeira a sair se aparecer uma mv2 melhor cotada.
- **`Sphere of the Suns`** (caixa) — **ressalva F7 que a simulação não penaliza**: são **3 cargas**,
  não mana infinito. O modelo a trata como permanente, então **o T4 real é ligeiramente pior que
  os 42,6% medidos**. Ela fica porque (i) continua artefato depois de esvaziar (metalcraft de
  graça) e (ii) as 3 cargas cobrem exatamente as 3 conjurações do Otharri (5 → 7 → 9).
- **`Mind Stone`** (R$ 2,25) — incolor, mas a medição mostrou que cor vale 0,5 pp e o saque tardio
  por `{1}` é a melhor taxa de conversão mana→carta do deck. **Fica.**
- **`Loyal Warhound`** (R$ 2,59) — **não acelera** (Plains **virada**, condicional), mas é 3/1 com
  vigilância que recebe os mesmos anthems estáticos do Myr (`Warleader's Call`,
  `Wedding Festivity`, `Jor Kadeen`). **Fica pelo corpo**, e eu conto ela na cota de ramp com a
  ressalva por escrito de que ela é a 11ª fonte no papel e a 10ª na prática.

---

## 5. Níveis de correção — o que cada slot compra, medido

| Nível | Composição | Slots novos | Δ R$ | T3 | **T4** | **T5** | nunca ≤T7 | mediana |
|---|---|---|---|---|---|---|---|---|
| — | **v1 hoje** | — | — | 6,7% | **42,6%** | **71,9%** | 12,7% | T5 |
| **N0** | Sword→**Pentad Prism** · Locket→**Myr Convert** | **0** | **−26,43** | 11,9% | **47,9%** | **75,0%** | 10,8% | T5 |
| **N1** | N0 + **37º terreno** | **1** | −26,23 | 11,6% | **48,8%** | **76,7%** | 9,7% | T5 |
| **N2a** | N1 + **`Hedron Crawler`** (caixa) | **2** | −26,23 | 11,6% | **52,3%** | **78,6%** | 8,7% | **T4** |
| **N2b** | N0 + **38 terrenos** | **2** | −26,03 | 12,1% | **50,5%** | **78,6%** | 8,4% | **T4** |

**Taxa de câmbio medida, por slot:** +1 terreno = **+1,4 pp no T4** e **+2,0 pp no T5**;
+1 rocha de mv2 = **+3,5 pp no T4** e **+1,9 pp no T5**.
Ou seja: **rocha compra o T4, terreno compra o T5 e derruba a cauda.**

### 5.1 O que eu peço ao orquestrador

**N0 é minha e não precisa de nada.** Duas trocas dentro dos meus slots, saldo de **+R$ 26,43**
devolvidos à folga. **Recomendo aplicar independentemente de qualquer negociação.**

**Peço 1 slot** para o Nível 1 (o 37º terreno). Justificativa: não é preferência minha, é a
**fórmula da própria Fase 6 rodada sobre a lista que existe** (§2.1) — **36,7 a 37,6**. O slot
sai de fora da minha categoria, então é negociação sua com a Fase 2.

**Nível 2 é opcional e eu ordeno assim:**
- **N2b (38 terrenos)** se a prioridade for **cauda** (8,4% × 8,7%) e manter o ramp padrão em
  **11**, dentro da meta 10–11.
- **N2a (37 terrenos + `Hedron Crawler`)** se a prioridade for o **T4** (52,3%, o melhor número da
  tabela) e a mediana T4. **Custo declarado: leva o ramp padrão a 12, acima da meta 10–11** —
  a decisão é sua, não minha.

**Se a Fase 2 precisar dos 2 slots para criaturas, eu cedo.** O N0 sozinho já entrega +5,3 pp no
T4, +3,1 pp no T5 e **paga a conta de todo mundo**: R$ 26,43 é mais que a folga inteira de
R$ 33,77 que estava em disputa.

---

## 6. Candidatas — ramp padrão

**Alvo 10–11 · No deck depois do N0: 11.** *(a coluna "Na coleção?" segue a regra 7)*

| Carta | CMC | Tipo | O que gera | Sinergias (mín. 2) | Na coleção? |
|---|---|---|---|---|---|
| **Sol Ring** | 1 | Artifact | `{C}{C}` | (a) único T3-Otharri do deck hoje; (b) artefato/metalcraft | **já é seu** · intocável |
| **Arcane Signet** | 2 | Artifact | 1 mana da identidade | (a) cor exata; (b) metalcraft | **já é seu** · intocável |
| **Boros Signet** | 2 | Artifact | `{1}`,`{T}`: `{R}{W}` | (a) **resolve as duas cores numa tapada** — é a única fonte que dispensa dois permanentes; (b) metalcraft | compra · R$ 1,50 |
| **Talisman of Conviction** | 2 | Artifact | `{C}`, ou `{R}`/`{W}` por 1 vida | (a) modo incolor grátis paga o `{3}` do Otharri; (b) metalcraft | compra · R$ 3,80 |
| **Mind Stone** | 2 | Artifact | `{C}` · `{1}`,sac: compra 1 | (a) melhor taxa mana→carta do deck; (b) metalcraft | compra · R$ 2,25 |
| **Sphere of the Suns** | 2 | Artifact | 3 cargas · 1 mana de qualquer cor | (a) **continua artefato depois de esvaziar**; (b) as 3 cargas = as 3 conjurações do Otharri | **CAIXA** |
| **Everflowing Chalice** | 0/2/4 | Artifact | `{C}` por carga | (a) única rocha que escala no late; (b) metalcraft; (c) sobrevive a wipe | compra · R$ 0,99 |
| **▲ Pentad Prism** | 2 | Artifact | **2 manas de qualquer cor** (sunburst) | (a) **única carta do teto que entrega Otharri no T3** — T3 6,7%→11,9%; (b) artefato permanente para metalcraft mesmo vazia; (c) recebe marcadores (alvo de `Karn's Bastion`) | **compra · R$ 0,57** |
| **Commander's Sphere** | 3 | Artifact | 1 mana da identidade · sac: compra 1 | (a) cor exata; (b) sac→carta; (c) metalcraft | **CAIXA** |
| **▲ Myr Convert** | 2 | Art. Creature 2/1 | 1 mana de **qualquer cor** (2 vidas) | (a) **artefato E criatura** — metalcraft + contagem de criaturas; (b) recebe os anthems estáticos (`Warleader's Call`, `Wedding Festivity`, `Jor Kadeen`) e ataca; (c) mv2 na curva certa | **CAIXA** |
| **Loyal Warhound** | 2 | Creature 3/1 | Plains **virada**, condicional | (a) corpo 3/1 com vigilância que recebe os anthems estáticos; (b) land ramp que sobrevive a wipe. **Ressalva: não acelera o T4** | compra · R$ 2,59 |

**▲ = entra neste passe.** Fora: `Sword of the Animist` (§4.1) e `Boros Locket` (§4.3).

## 6.1 Candidatas — ramp explosivo

**Alvo 2–3 · No deck: 2. Sem mudança neste passe.**

| Carta | CMC | Tipo | O que gera | Sinergias | Na coleção? |
|---|---|---|---|---|---|
| **Battle Hymn** | 2 | Instant R | `{R}` **por criatura sua** | (a) escala com o enxame do Otharri; (b) instantâneo → gera mana **depois** do gatilho de ataque, quando as fichas já existem; (c) converte board condenado em mana em resposta a wipe | compra · R$ 6,89 |
| **Mana Geyser** | 5 | Sorcery R | `{R}` por terreno virado dos oponentes | (a) **não pede board** — cobre o estado oposto ao do Battle Hymn; (b) paga a recast de 7/9 com o campo vazio. **Atrito medido na §7** | compra · R$ 12,35 |

---

## 7. A pergunta do prompt: a base branca sustenta recast + proteção no mesmo turno?

`Mana Geyser` produz **só `{R}`**. O recast do Otharri custa `{5}{R}{W}` = **7**. Então o turno
exige: pagar `{3}{R}{R}` do próprio Geyser **e** manter uma fonte de `{W}` desvirada.

**Medi exatamente isso** — configuração de campo no turno N, exigindo ≥6 manas, ≥2 fontes de `{R}`
para o Geyser e ≥1 fonte de `{W}` **sobrando** depois:

| Configuração | T5 | T6 | **T7** | T8 |
|---|---|---|---|---|
| v1 atual | 30,2% | 44,9% | **55,9%** | 64,7% |
| **N0** | 35,3% | 49,9% | **60,6%** | 69,1% |
| N1 (37 terrenos) | 36,1% | 50,7% | **61,3%** | 69,7% |

**Decomposição das falhas no T7 (v1 atual):**

| | % das partidas |
|---|---|
| **OK** — Geyser + `{W}` sobrando | 55,9% |
| falha: **menos de 6 manas em campo** | **26,8%** |
| falha: fontes existem mas **não sobra `{W}`** depois do Geyser | **13,9%** |
| falha: **menos de 2 fontes de `{R}`** (o Geyser pede `{R}{R}`) | 3,3% |
| falha: nenhuma fonte de `{W}` | 0,1% |

**Resposta: sim, a base branca sustenta — o `{W}` não é o gargalo.** Com 21 fontes de `{W}`,
"nenhuma fonte de branco" acontece em **0,1%** das partidas. O atrito real é outro e tem dois
nomes:

1. **Densidade de mana (26,8%)** — no T7 o deck simplesmente não tem 6 permanentes de mana em
   metade das partidas. É a **mesma** causa raiz da §2, aparecendo de novo. O N0/N1 a ataca e o
   número sobe para 60,6%/61,3%.
2. **A tapada reservada (13,9%)** — o `{W}` existe, mas a fonte que o produziria é necessária para
   pagar o próprio Geyser. Isso **não é problema de cor da base**; é problema de **contagem de
   permanentes**. `Pentad Prism` ajuda aqui de duas maneiras: produz **qualquer cor** e entrega
   **2 manas de uma tapada**, o que é exatamente a folga que falta.

**Ressalva honesta: "recast + proteção no mesmo turno" é mais caro que isso.** Somar
`Feat of Resistance` (`{1}{W}`) ou `Gods Willing` (`{W}`) ao turno do Geyser exige **1 a 2 manas
brancos a mais desvirados**, não um. A medição acima cobre **recast + 1 `{W}` sobrando**. Se o
orquestrador quiser o número de "recast + `Feat of Resistance`", eu rodo — mas adianto a direção:
cai mais uns 10 pp, e o conserto continua sendo densidade, não cor.

**Achado cruzado para a Fase 6 (não é corte meu, §8.3):** a base é **62% branca / 38% vermelha**
(17 Plains / 14 Mountain), e o único ramp explosivo do deck pede `{R}{R}`. Simulei
**16 Mountain / 15 Plains**: a linha do Geyser sobe de **55,9% para 59,6%** no T7 e a curva do
Otharri **não muda nada** (T4 42,7% × 42,6%), porque ele pede um de cada. **Mas o meu modelo não
enxerga os pips duplos de branco** (`Hero of Bladehold` `{2}{W}{W}`, `Palace Jailer` `{2}{W}{W}`,
`Basri Ket` `{1}{W}{W}`, `The Circle of Loyalty` `{4}{W}{W}`, `Hour of Reckoning` `{4}{W}{W}{W}`),
que são o motivo declarado do 17/13 da Fase 6. **Entrego o número e não decido** — é slot dela.

---

## 8. Cortes condicionados e achados fora da minha lente — devolvo sem veredito

### 8.1 `Oracle's Vault` — `{4}` Artifact · caixa · **corte condicionado, Fase 3**

| Eixo | Leitura (oracle puxado nesta sessão) |
|---|---|
| **F1** | `{2}`,`{T}`: exile a carta do topo; você pode jogá-la neste turno; ponha um marcador de tijolo · `{T}`: exile o topo e jogue **sem pagar o custo**, só com 3+ marcadores |
| **F2** | Sem corpo. Usa o próprio `{T}` |
| **F3** | **É artefato** → metalcraft |
| **F4** | **Recebe marcadores de tijolo** (e proliferate, se `Karn's Bastion` entrar) |
| **F5** | Dá acesso a cartas, não mana |
| **F6** | **mv4 — não mv3, como o `report.md` registrou.** Cai no turno em que o deck quer conjurar o comandante |
| **F7** | **Não é ramp e nunca contou na minha cota.** Mas é o item da lista que mais **piora a lentidão percebida**: 4 manas para não fazer nada no turno, e cada ativação custa `{2}` que sairia do desenvolvimento. São **3 turnos de `{2}`** antes do modo grátis ligar |

**Não corto: a função dela é saque por impulso, que é da Fase 3.** Devolvo ao orquestrador com o
número: se este slot virar terreno, o deck ganha o equivalente ao Nível 1 **e** para de gastar
4 manas no turno do comandante. Quem decide se a função de saque fica descoberta é a Fase 3.

### 8.2 `Rustvale Bridge` → `Rugged Prairie` — **achado da Fase 6, zero slot**

`Rustvale Bridge` **entra virada**. `Rugged Prairie` (`{T}`: `{C}`; `{R/W}`,`{T}`: `{R}{R}`,
`{R}{W}` ou `{W}{W}`) **entra desvirada** e é o melhor fixador barato de RW. **Medido: +1,2 pp no
T4 e no T5, sem custar slot nenhum** (N1 vai de 48,8% para 50,0%).
**Custo declarado:** `Rustvale Bridge` é **Artifact Land** — conta para metalcraft e é
indestrutível. A troca derruba a contagem de artefatos de 18 para 17 (`Jor Kadeen` precisa de 3;
sem risco). **Pendência de cotação LigaMagic** — estimativa Scryfall US$ 0,33, que não decide nada.
**É slot da Fase 6, não meu.**

### 8.3 Proporção 17 Plains / 14 Mountain — **Fase 6**
Número medido na §7. Entregue, não decidido.

---

## 9. Orçamento

**Régua: menor valor da LigaMagic (regra 2).** Nada aqui é estimativa da Scryfall.

| Movimento | R$ | Cotação |
|---|---|---|
| **−** `Sword of the Animist` (compra cancelada) | **−27,00** | LigaMagic menor, 2026-09-19 (4 d) |
| **+** `Pentad Prism` (compra — a cópia física está montada no Inspirit) | **+0,57** | LigaMagic menor, 2026-09-18 (5 d) |
| **−** `Boros Locket` (era caixa) | 0,00 | — |
| **+** `Myr Convert` (caixa) | 0,00 | R$ 0,15 registrado, 2026-08-12 |
| **Saldo do Nível 0** | **−26,43** | |
| **+** 1 terreno básico (Nível 1) | +0,20 | básicos ≈ R$ 0,20 |
| **+** `Hedron Crawler` (Nível 2a, caixa) | 0,00 | R$ 0,10 registrado, 2026-08-12 |

**Custo das faltantes: R$ 166,23 → R$ 139,80 (N0) ou R$ 140,00 (N1).**
**Folga sobre o teto de R$ 200,00: R$ 33,77 → R$ 60,20 (N0) / R$ 60,00 (N1).**

**A folga deixa de ser disputada.** Eu devolvo **R$ 26,43** — mais do que a folga inteira que
estava na mesa. O pacote inteiro da Fase 4 custa **R$ 0,57**, e os R$ 26,43 ficam disponíveis para
a Fase 2 resolver "poucas criaturas".

> **Nota de escopo, não de decisão.** Estou reportando o delta na mesma unidade do `report.md`
> ("custo das faltantes"). Se a régua do Commander 200 for o **valor das 99 cartas** e não o
> **gasto de compra**, as cartas da caixa também contam e o total é outro. Não relitigio isso
> aqui — sinalizo para o orquestrador porque o corte do `Sword of the Animist` melhora o número
> **nas duas leituras**.

---

## 10. Pendências

1. **`Rugged Prairie` — sem cotação LigaMagic.** Vale +1,2 pp no T4 por zero slot. Estimativa
   Scryfall US$ 0,33, que **não decide** (o `Mana Geyser` estimava US$ 1,08 e saiu a R$ 12,35).
   **Capturar antes de decidir.** Slot da Fase 6.
2. **`Irencrag Feat` — segue sem cotação** (pendência aberta no passe 2). O usuário já a
   descartou por função ("limita a uma mágica por turno"); mantenho fora e **não reproponho**.
   Registro só para a pendência não sumir.
3. **Rocha permanente de mv2 que produza cor** (`Coldsteel Heart`, `Fire Diamond`,
   `Marble Diamond`, `Star Compass`, `Guardian Idol`) — **nenhuma tem cotação LigaMagic**. Seriam
   um upgrade estrito sobre a `Sphere of the Suns` (que tem só 3 cargas). Com R$ 26,43 de folga
   devolvida, **vale capturar uma cotação** — mas sem o número eu não proponho a compra.
4. **A simulação é otimista.** Piloto perfeito, sem interação adversária, `Sphere of the Suns`
   modelada como permanente. Os números reais são **piores** que os da tabela — o que reforça,
   não enfraquece, o diagnóstico. **O goldfishing da Fase 7 (10 partidas) é o teste que fecha
   isso**, e ele continua sem rodar.
5. **`Oracle's Vault`** (§8.1) — corte condicionado devolvido à Fase 3, com a ficha F1–F7 feita.
