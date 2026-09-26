# Registro de decisão · Tori D'Avenant, Fury Rider

Regra 5 do `CLAUDE.md`: todo corte e toda entrada são registrados aqui, em ordem cronológica.
Antes de propor carta que já esteve no deck, **consulte este arquivo** — a proposta precisa
dizer quem cortou, por quê, e o que mudou desde então.

## Estado inicial

O deck foi montado pelo usuário **fora do pipeline**, antes de 2026-09-18. Não existe histórico
de cortes anteriores: as 99 cartas de `lista.txt` são a composição original, nenhuma delas
passou por corte e retorno. **Qualquer carta que saia a partir da v1 passa a ter histórico aqui.**

## Decisões

| Data | Carta | Ação | Fase/agente | Motivo | Funções descobertas pelo corte |
|---|---|---|---|---|---|
| 2026-09-19 | **Otharri, Suns' Glory** | **entra — comandante** | Fase 1 / orquestrador | Escolhido pelo usuário entre 3 finalistas. Resolve as dores 1 e 4 e contorna a 2; é o único que *fabrica* o board de que o resto do deck depende. Contadores de experiência ficam no jogador e sobrevivem a board wipe — única resiliência a wipe disponível a RW em qualquer faixa de preço. Zera o gap de wincon (−3) sem gastar slot, porque comandante não conta na régua de R$ 200. | F1 wincon · F2 corpo 3/3 voador com vínculo e ímpeto · F3 Phoenix/Rebel-maker · F4 recebe contadores de experiência e anthems · F6 CMC 5 |
| 2026-09-19 | **Tori D'Avenant, Fury Rider** | **sai da zona de comando** | Fase 1 / orquestrador | Não gera carta, corpo, mana nem fechamento — o gatilho dela é pump temporário condicionado a já haver board. Incompatível com a régua "extremamente competitivo no Commander 200". **Não é corte do deck:** é RW e continua elegível para os 99; a avaliação como carta do 99 é tarefa da Fase 2 e exige ficha F1–F7 (regra 4). | F1 pump em massa + concede trample + destapa brancas · F2 corpo 3/3 com vigilância e trample · F3 Human Knight · F5 buffa outros · F6 CMC 4 |

### Nota de contexto (regra 5)

**Aurelia, the Law Above** (R$ 4,00) e **Aurelia, the Warleader** (R$ 58,50) foram consideradas e
**não escolhidas** em 2026-09-19 — não estão banidas do deck, mas repropô-las como comandante exige
dizer o que mudou:

- **Law Above** — descartada porque o motor de compra dela exige **3+ atacantes seus** para girar,
  que é exatamente o estado que o deck perde na dor 4: compra enquanto você ganha e para quando
  você precisa recuperar. O argumento a favor continua forte e registrado: draw é o gap maior
  (−9), o mais caro em RW, e **a caixa não tem nenhuma fonte de draw dentro da identidade**
  (`Reconnaissance Mission` e `Bident of Thassa` são azuis). Se o deck pronto não resolver o draw
  dentro do teto, ela volta à mesa.
- **Warleader** — descartada por imposto (6→8→10, com gap de ramp de −9) e por não resolver
  nenhuma das quatro dores. **Também está fora dos 99**: a R$ 58,50 ela consome 29% do teto de
  R$ 200. Só voltaria se o preço cair muito ou se a régua do formato mudar.

---
## 2026-09-23 — v1 reprovada no teste de mesa; Otharri sai da zona de comando

**Decisão do usuário**, após imprimir a v1 em proxy e jogar: *"Não gostando desse deck."*
Incômodos marcados: **não parece competitivo**, **o comandante (Otharri)**, **o jeito de jogar**.
Autorização concedida: **trocar o comandante, mantendo RW**.

| Carta | Movimento | Motivo |
|---|---|---|
| `Otharri, Suns' Glory` | **sai da zona de comando** | Reprovado em jogo real. O deck inteiro dependia dele para ter tabuleiro; as 7 trocas da §12 da v1 corrigiram as três queixas *medidas* (corpos 19→25, mão sem corpo 30,4%→12,1%, T4 42,6%→47,9%) e ainda assim o deck não convenceu — sinal de que a causa é o **eixo**, não as peças. Continua elegível para os 99 como qualquer RW. |

**Nada foi comprado** — as 31 compras da v1 (R$ 139,16) nunca saíram do papel, então não há custo
afundado e todos os slots estão livres. `deck.md` segue inexistente; as 7 trocas da §12 da v1
**não** entram neste registro porque nunca foram aprovadas.

**Consequência:** fica fechada a `DECISÃO EM ABERTO — eixo do deck` de 2026-09-20. A v2 reabre a
Fase 1 com o critério declarado: potência real no Commander 200 e **eixo de jogo diferente de
go-wide de fichas** (que é a pegada do Krenko, já apontada pelo usuário).

## 2026-09-23 — comandante da v2: Phlage, Titan of Fire's Fury

**Decisão do usuário**, sobre as 5 opções da Fase 1 (`rounds/v2-2026-09-23/01-commander.md`):
*"Vou testar com Phlage."*

| Carta | Movimento | Motivo |
|---|---|---|
| `Phlage, Titan of Fire's Fury` | **entra na zona de comando** | Único dos 5 finalistas que não repete nenhuma das duas queixas do Otharri: sai no **T3** (contra CMC 5) e **se reconjura do cemitério por custo fixo** `{R}{R}{W}{W}`, sem imposto de comandante (rulings conferidos). Maior aproveitamento do pool (58 cartas). O corpo *é* a remoção — dois slots em um, sob teto de R$ 200. |

**Eixo adotado:** controle de atrito RW — remoção densa, pouco tabuleiro próprio, ameaça recorrente.
Substitui o go-wide de fichas da v1.

**Descartados com motivo** (detalhe em `01-commander.md`): `Feather, the Redeemed` (dependência do
comandante pior que a do Otharri) · `Quintorius, Loremaster` e `Hofri Ghostforge` (CMC 5, imposto
5→7→9 — a mesma lentidão reprovada) · `Gisela, Blade of Goldnight` (CMC 7) · `Commander Liara
Portyr` e `Cadric, Soul Kindler` (refutados na Fase 1).

**Pendência aberta junto com a escolha:** o eixo **não** resolve a dor 1 (não fecha o jogo). O
`wincon-tester` recebe como requisito duro um fecho nomeado que não seja "atacar com o comandante".
**`Gisela, Blade of Goldnight` fica como primeira candidata a finisher nos 99** — em RW e sem
imposto de comandante quando é carta do deck. Preço **a cotar**, junto com `Fiery Emancipation`.

**Travas mecânicas do escape, a respeitar na construção:**
1. O custo exila **5 cartas *além* dele** — são 6 no cemitério para escapar a primeira vez.
2. Só vale **conjurar por escape**: reanimar ou blinkar faz o gatilho de sacrifício pegar.
   **Nenhum pacote de reanimação ou blink entra neste deck.**
   > ⚠ **Correção 2026-09-24 (orquestrador): a conclusão está errada.** O sacrifício pega, mas o
   > Helix também dispara. Ruling oficial de 2024-06-07: *"Phlage's second ability triggers when
   > it enters the battlefield, even if it didn't escape."* Reanimar o Phlage é um Helix por 1–2
   > manas, e ele volta ao cemitério pronto para a próxima reanimação. É o motor das listas do
   > EDHREC e do Archidekt. A trava fica **revogada para reanimação**. Blink continua fraco: só
   > funciona em resposta ao sacrifício e tira o escape do Phlage escapado. Análise em
   > `rounds/v2-2026-09-23/08-meta-edhrec.md`.
3. Conjurado da zona de comando, o Helix (3 de dano + 3 de vida) **acontece assim mesmo** — são
   duas habilidades separadas, e o dano dispara sempre que ele entra, escapado ou não.
