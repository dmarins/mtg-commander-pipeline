# Interação — v2 (escopo estreito: indestructible para o próprio Inspirit)

> Esta fase **não é uma auditoria geral de interação** — o pipeline de v1 já cobriu remoção/wipes/proteção com boa profundidade (ver `02-theme.md`: 9 remoções, 3 wipes assimétricos, 5 fontes de proteção, 1 counter isolado). O escopo aqui é resolver o ponto cego específico identificado: **o próprio Inspirit não tem indestructible, só hexproof (via Padeem/Leonin Abunas)**, e a estática dele ("*other* artifacts") não se aplica a si mesmo.

## Contexto do swap

**Corte aprovado**: Gilded Lotus (5 mana, rock sem contadores/corpo, redundante com os outros 10 rocks do deck — ver diagnóstico em `02-theme.md`).

**Vaga a preencher**: 1 carta que conceda **indestructible** a um artefato-alvo ou permanente-alvo. Restrição-chave: Inspirit é **artefato não-criatura** até acumular 8+ marcadores via Station — logo cartas restritas a "target creature" (a maioria dos equipamentos indestructible do formato: Darksteel Plate, Hammer of Nazahn, Kaldra Compleat, Shield of Kaldra etc. — todos exigem criatura equipada) **não funcionam nele** na maior parte do jogo. Busquei apenas o subconjunto que funciona com "target artifact" ou "target permanent".

## Achado prévio relevante (transparência)

Ao revisar `02-theme.md`, notei que o deck **já tem 3 instantâneos** que tecnicamente concedem indestructible a permanente-alvo (incluindo o próprio Inspirit) — mas todos **reativos e "até o final do turno"**, não permanentes:
- **Loran's Escape** (W) — "Target artifact or creature gains hexproof and indestructible until end of turn."
- **Blacksmith's Skill** (W) — "Target permanent gains hexproof and indestructible until end of turn."
- **Restoration Magic** (W, tiered) — mesmo efeito, com ganho de vida opcional.

O diagnóstico da v2 ("nenhuma fonte dá indestructible") ignorou essas três — na prática, **já existe uma resposta pontual reativa**, desde que a carta certa esteja na mão e o jogador tenha mana disponível no momento exato (o que falha contra wipes surpresa sem mana aberta, ou quando a carta certa não é comprada a tempo). A vaga do Gilded Lotus deve, portanto, ser usada para **reforçar essa camada** com algo que cubra mais alvos por cópia e/ou seja mais barato — não apenas duplicar o que já existe.

## Candidatas — indestructible para artefato/permanente-alvo

| Carta | CMC | Tipo | O que faz | Sinergias | Faixa/preço |
|---|---|---|---|---|---|
| **Invisible Force Field** ⭐ | {1}{W} = 2 | Instant | "Up to four target permanents you control gain indestructible until end of turn." + **Rebound** (recasta de graça no próximo upkeep, sem pagar custo) | Protege **até 4 permanentes simultaneamente** — cobre Inspirit **e** o cluster de criaturas não-artefato expostas identificado em `02-theme.md` (Jhoira, Sai, Malcator, Third Path Iconoclast, Chrome Host Seedshark, Saheeli) num único cast; Rebound dá 2 janelas de proteção pelo preço de 1 carta; instant speed = pode ser jogada em resposta a um wipe ou remoção antes de resolver | torneio — **US$0,44** |
| Reroute Systems | {W} = 1 | Instant | Modal: "Target artifact or creature gains indestructible until end of turn" **ou** 2 dano a criatura tapada-alvo | 1 mana, mais barata; modo secundário serve como remoção leve de emergência; só 1 alvo por cópia (menos cobertura que Invisible Force Field) | torneio — **US$0,15** |
| Slobad, Goblin Tinkerer | {1}{R} = 2 | Legendary Creature | "Sacrifice an artifact: Target artifact gains indestructible until end of turn." — **repetível**, sem limite de usos por turno | Motor de token-makers do deck (Sai, Third Path Iconoclast, Saheeli, Chrome Host Seedshark, Thopter Spy Network) fornece fodder farto para sacrificar sem perder valor real; mas Slobad é criatura não-artefato — não protegido pela estática do Inspirit, e some para qualquer remoção/wipe | torneio — **US$0,60** |
| Aegis Angel (referência, não recomendada) | {4}{W}{W} = 6 | Creature — Angel | ETB: "another target permanent gains indestructible for as long as you control this creature" — efeito **contínuo**, não "até o final do turno" | Proteção permanente enquanto sobreviver; mas custo alto, corpo fora do tema (sem interação com artefatos), e a proteção cai junto se a Angel for removida — competindo justamente com o problema que resolve | torneio (US$0,49) mas descartada por custo de curva e desalinhamento temático |
| Darksteel Forge (referência, não recomendada) | {9} | Artifact | "Artifacts you control have indestructible" — cobre **todos** os artefatos, incluindo Inspirit, permanentemente, sem exigir alvo | Seria a solução mais completa em teoria (nenhuma outra carta do pool cobre isso) | **US$39–45** — muito acima de qualquer faixa (torneio ≤R$200 total; mesão ≤R$350 total); inviável |

## Proposta de swap

**Gilded Lotus → Invisible Force Field** (torneio, US$0,44 — mais barata que a carta cortada)

Justificativa (2+ pontos de sinergia):
1. Resolve diretamente o ponto cego do comandante: pode ter Inspirit como um dos até 4 alvos, concedendo indestructible exatamente quando um wipe não-direcionado, "-X/-X para todos" ou "destroy all artifacts" ameaçar matá-lo (a estática dele não protege a si mesmo).
2. Cobertura múltipla incomum: como aceita até 4 alvos simultâneos, na mesma resolução protege também Jhoira, Sai, Malcator, Third Path Iconoclast, Chrome Host Seedshark e/ou Saheeli — exatamente o cluster de criaturas não-artefato que `02-theme.md` apontou como exposto (elas não se beneficiam da estática "other artifacts" por não serem artefatos).
3. Rebound dobra o valor por carta sem custo adicional — uma única cópia entrega 2 janelas de proteção (turno atual + próximo upkeep), o que compensa bem a natureza "reativa" da carta.
4. Custo de mana baixíssimo ({1}{W}) para instant speed — cabe fácil no orçamento de mana do turno junto com os outros protetores one-shot já existentes (Loran's Escape, Blacksmith's Skill, Restoration Magic), sem competir por slots de curva alta como o Gilded Lotus cortado.

## Alternativas (para o usuário escolher, caso prefira outro perfil)

1. **Reroute Systems** — se preferir a opção mais barata (US$0,15) e não se importar em proteger só 1 alvo por cópia; ganha um modo de remoção leve como bônus (2 dano em criatura tapada), o que ajudaria levemente a cobertura de "remoção" também.
2. **Slobad, Goblin Tinkerer** — se preferir um motor **repetível** ao longo do jogo (em vez de um efeito pontual), aproveitando o volume de tokens de artefato que o deck já gera; ressalva: Slobad é criatura não-artefato, portanto ele mesmo fica exposto a remoção/wipe (inclusive aos 3 wipes do próprio deck, que precisaria proteger com cuidado ou aceitar perder).

## Cobertura de ameaças (apenas para esta vaga específica)

Indestructible para Inspirit/permanentes-chave: **0 → 1** fonte dedicada (mais robusta que as 3 já existentes, que seguem no deck sem alteração) — cobre wipes simétricos/assimétricos que usam "-X/-X", "destroy all artifacts" ou "destroy all permanents", e complementa (não substitui) o hexproof permanente de Padeem/Leonin Abunas.

## Reservas / Swaps propostos

Nenhum swap adicional foi avaliado nesta fase — escopo estreito conforme solicitado. Não avaliar Talisman of Progress/Azorius Signet, Chainsaw, ou a lacuna de counterspells aqui; esses pontos já estão registrados em `02-theme.md` ("Candidatas a corte/revisão") para uma eventual fase de interação mais ampla, caso o usuário decida reabrir esse escopo depois.

---

## Fase 2 — reabertura do escopo: lacuna de counterspell (Chainsaw → contramágica)

> Escopo desta seção: fechar a lacuna identificada em `02-theme.md` ("Counterspells: 1 — Disruption Protocol isolado") e formalizada no briefing desta rodada. O corte abaixo já foi discutido e aprovado pelo usuário antes desta busca.

### Corte aprovado: Chainsaw

`02-theme.md` havia "salvo" Chainsaw na primeira passada desta v2 (achado #3: o rev counter conta qualquer morte, não só as do jogador, então os 3 wipes do próprio deck o alimentam). Discussão adicional com o usuário revisou a leitura da regra oficial: um wipe que mata N criaturas de uma vez ainda dispara o gatilho "whenever one or more creatures die" **uma única vez** (evento simultâneo = um gatilho), não N vezes — portanto Chain Reaction/Fumigate/Organic Extinction rendem **1 rev counter por wipe**, não uma rajada. O motor é bem mais lento do que a leitura inicial sugeria. Combinado com `Equip {3}` (caro para a curva do deck) e nenhum outro corte "de graça" disponível abaixo de 2 pontos de sinergia no pool de 99, o usuário aceitou cortar Chainsaw especificamente para abrir a vaga de counterspell — é a peça mais fraca e mais isolada do pacote de remoção (ETB 3 dano é bom, mas o resto da carta não paga o Equip caro).

### Busca

`otag:counterspell id<=wur legal:commander mv<=3 order:edhrec` (329 resultados, fatiado) e `otag:counterspell id<=wur legal:commander o:artifact order:edhrec` (31 resultados, para o subconjunto com sinergia direta de artefato). Priorizei: eficiência (baixo CMC, instant), sinergia com o volume de 40+ artefatos (metalcraft/affinity/token de Thopter — o deck já usa metalcraft como argumento para Dispatch), e preço dentro da faixa torneio (`usd<0.5`) quando possível.

### Candidatas

| Carta | CMC | Tipo | O que faz | Sinergias | Faixa/preço |
|---|---|---|---|---|---|
| **Stoic Rebuttal** ⭐ | {1}{U}{U} = 3 (efetivo {U}{U} com metalcraft) | Instant | Metalcraft — custa {1} a menos com 3+ artefatos. "Counter target spell" (qualquer spell, sem restrição de tipo) | Metalcraft é trivial no deck (mesmo argumento já usado para Dispatch/StP nº2: 40+ artefatos garantem 3+ em jogo já no T3–4); vira efetivamente um Counterspell de 2 manas sem downside; hard counter cobre combo, bomba e remoção alheia igualmente | torneio — **US$0,32** |
| Access Denied | {3}{U}{U} = 5 | Instant | "Counter target spell. Create X 1/1 Thopter artifact creature tokens, X = mana value do spell contra-atacado" | Cria Thopters — encaixa no subtema de tokens de artefato já presente (Sai gera Thopters, Thopter Spy Network, Uthros Research Craft); contra spells caros (bombas, wipes alheios) pode gerar 5+ Thopters, que viram fodder de Station/Sai/proliferate; mas custo alto (5) e sem redução | mesão — **US$2,62** |
| Assert Authority (referência, não recomendada) | {5}{U}{U} = 7, Affinity for artifacts | Instant | Custa {1} a menos por artefato controlado; "Counter target spell", exila em vez de ir pro cemitério se contado assim | Tematicamente a mais "correta" (Affinity é literalmente uma palavra-chave do tema do deck) e pode ficar barata (~{U}{U} com 5+ artefatos); mas depende de ter artefatos suficientes em jogo cedo, o que nem sempre coincide com o momento em que se precisa segurar um combo no T2–3 | torneio (US$0,26) mas descartada por fragilidade de curva no early game |
| Defabricate | {1}{U} = 2 | Instant | Modal: conta artefato/encantamento (exila) OU conta ativada/gatilhada | Cobre metade da "cobertura de artefato/encantamento" pedida pelo guia, mas restrita a esses tipos — não segura uma criatura ou combo genérico | torneio — US$0,17 |
| Halt Order | {2}{U} = 3 | Instant | Conta spell de artefato + compra 1 carta | Só conta artefato — muito restrito para ser o único counter novo do deck | torneio — US$0,09 |

### Escolha recomendada: Stoic Rebuttal (torneio, US$0,32)

Justificativa (2+ pontos de sinergia):
1. **Sinergia direta e comprovada com o próprio deck**: o mesmo argumento de "metalcraft trivial" já usado em `02-theme.md` para justificar Dispatch (40+ artefatos garantem 3+ em jogo cedo) se aplica aqui — na prática, Stoic Rebuttal age como um Counterspell de 2 manas ({U}{U}) a partir do T3–4, sem nenhuma restrição de tipo de spell (ao contrário de Negate/Disruption Protocol que só pegam "noncreature" ou têm custo extra).
2. **Cobertura complementar a Disruption Protocol**: Disruption Protocol já é um hard counter de 2 manas com custo adicional trivial (tap artefato ou {1}); Stoic Rebuttal traz um segundo hard counter de perfil parecido, mas sem exigir tap de um artefato (relevante se o artefato já foi usado para mana/ativação naquele turno) — os dois juntos dobram a chance de ter contramágica disponível em qualquer turno, resolvendo a lacuna "1 counter isolado" identificada na auditoria.
3. Preço muito abaixo do teto torneio (US$0,32 vs. os US$0,90 aproximados de Chainsaw), sobrando margem de orçamento.
4. Instant speed, sem downside — não compete com o plano de curva do deck (a maioria dos habilitadores de tema já ocupa CMC 2–3; este counter só é relevante quando há algo para contrariar, não estufa a curva de "peças ativas").

### Alternativa para o usuário (perfil mesão / maior impacto)

**Access Denied** (mesão, US$2,62) — se o usuário preferir um counter que também gera valor tangível em vez de só neutralizar (relevante em mesas mais lentas/casuais onde segurar uma bomba cara de 5+ CMC rende 5+ Thopters, board wide que alimenta Station, Sai e o cluster de proliferate). Trade-off: 5 mana é tarde para segurar combos baratos, e não tem redução de custo como Stoic Rebuttal.

### Proposta de swap

**Chainsaw → Stoic Rebuttal** (torneio, US$0,32)

### Cobertura de ameaças (atualização desta vaga)

Counterspells: **1 → 2** (Disruption Protocol + Stoic Rebuttal) — ambos hard counters de qualquer spell, cobrindo combo, bomba e remoção alheia; lacuna da auditoria (`02-theme.md`, "Counterspells: 1 — lacuna real") tratada, ainda que modestamente (2 counters em 99 cartas segue abaixo do padrão típico de 3–5 num Commander competitivo — aceitável dado que o deck não tem mais vagas "de graça" para ceder sem tocar no core temático).

### Reservas

Se o usuário quiser uma 3ª contramágica no futuro (fora do escopo desta troca 1-para-1), Access Denied e Defabricate seguem como candidatas registradas acima, junto com a redundância Talisman of Progress/Azorius Signet já sinalizada em `02-theme.md` como possível próxima vaga "de graça".
