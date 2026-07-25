# Condições de Vitória e Testes — Inspirit, Flagship Vessel

## Diagnóstico

O deck já saiu das quatro fases anteriores (draw, ramp, interação, manabase) com **6 cartas** rotuladas `wincon` na auditoria temática (`02-theme.md`) e no `deck.md`: Master of Etherium, Alibou Ancient Witness, Kappa Cannoneer, Cyberdrive Awakener, Lux Artillery e Thousand Moons Smithy. Depois de reler o oracle text atualizado de cada uma (todas impressas em EOC/agosto-2025, fora do meu conhecimento "de memória", por isso conferidas uma a uma no Scryfall), concluo que isso **não é um punhado de cartas boas soltas** — elas se organizam em **3 caminhos de vitória genuinamente distintos**, todos nascidos do próprio mecanismo do comandante (artefatos em massa + marcadores + proteção estática). Não vejo lacuna que exija swap obrigatório. Proponho **um único reforço opcional** (não obrigatório) abaixo.

Ponto estrutural que vale destacar: **quase todos os finishers são, eles mesmos, artefatos** (Master of Etherium, Alibou, Kappa Cannoneer, Cyberdrive Awakener e Thousand Moons Smithy — face frontal — são todos `Artifact`/`Artifact Creature`). Isso significa que, enquanto Inspirit estiver vivo, a estática **"Other artifacts you control have hexproof and indestructible"** protege os próprios finishers, não só os tokens de suporte. É a razão pela qual esses 6 wincons fazem mais sentido aqui do que em qualquer outro deck "artifacts matter" — eles são auto-protegidos pelo comandante que os habilita.

## Caminhos de vitória atuais

| Caminho | Cartas envolvidas | Turno estimado | Consistência |
|---|---|---|---|
| **A. Enxame de combate (anthem + evasão)** | Master of Etherium (lord, P/T = nº artefatos), token-makers (Sai Master Thopterist, Third Path Iconoclast, Pinnacle Emissary, Chrome Host Seedshark, Saheeli Sublime Artificer), Kappa Cannoneer (unblockable, ward 4, cresce a cada artefato ETB), Cyberdrive Awakener (voo **permanente** para todos os artefatos-criatura + zumbificação pontual de rocks), Alibou (haste ao time) | 6–8 | **Alta** — quase todo o deck (35+ cartas) alimenta este caminho; é o "plano A" natural do tema |
| **B. Queima/alt-win via marcadores** | Lux Artillery (30+ marcadores entre artefatos/criaturas → 10 dano/oponente no end step), alimentado por Tezzeret's Gambit, Deepglow Skate, Surge Conductor, Kilo, Crystalline Crawler, Hangarback Walker, Empowered Autogenerator, Everflowing Chalice, Sphere of the Suns, Pentad Prism, Coretapper, o próprio trigger 1+ do Inspirit | 7–9 | **Média** — depende de acumular 30 marcadores no total (não numa carta só), mas o subtema de proliferate já existente no deck (8+ peças) alimenta isso de graça; **não depende de combate nem de os oponentes terem bloqueadores** |
| **C. Conversão explosiva de rocks** | Cyberdrive Awakener (transforma cada artefato não-criatura em 4/4 por 1 turno, todos protegidos pelo próprio Inspirit) + Alibou (haste, dano extra por artefatos virados) + pacote de 10+ mana rocks já otimizado na fase ramp | 6–8 | **Média-alta** — precisa de 5+ rocks em jogo (realista dado o ramp package), mas é a virada "de graça" que transforma uma mesa de suporte numa ofensiva letal num único turno |

Os caminhos A e C se sobrepõem fortemente (Cyberdrive/Alibou empurram A); B é o único plano independente de combate, útil se a mesa travar a via terrestre com bloqueadores/vigilância. Isso já configura **3+ caminhos de vitória com 2+ pontos de sinergia cada** (meta cumprida) sem necessidade de troca.

## Finishers recomendados (reforço opcional — sem obrigatoriedade de swap)

| Carta | CMC | Como fecha o jogo | Sinergias (mín. 2) | Na coleção? |
|---|---|---|---|---|
| Chief of the Foundry | 3 | "Other artifact creatures you control get +1/+1" — segundo anthem, empilha com Master of Etherium (time fica +2/+2 combinado) | (1) é `Artifact Creature`, logo protegido pela estática do próprio Inspirit assim como Master of Etherium; (2) reforça diretamente o Caminho A; (3) colorless, cabe sem stress na manabase já ajustada | Não — `torneio`, US$ 0,20 (≈ R$ 1,10), bem dentro do teto de R$ 200 |

Não recomendo forçar esse swap agora — o deck já bate a meta de 3+ caminhos. Se o usuário quiser mais redundância no caminho A (por exemplo, se os testes abaixo mostrarem vitórias "por pouco" no turno 7–8), Chief of the Foundry é o primeiro candidato de entrada, financiado por um corte entre as cartas menos essenciais (candidato natural: Warmaker Gunship, que já cumpre seu papel de remoção pontual mas é o corpo autônomo menos necessário ao plano de vitória).

Avaliei também Urza, Prince of Kroog (+2/+2, US$ 0,60) e Unctus, Grand Metatect (+1/+1, US$ 0,36) como alternativas de anthem — descartei ambos porque **não são artefatos** (Legendary Creature, não Artifact Creature), logo ficam de fora da proteção estática do Inspirit e seriam alvos fáceis de remoção pontual, quebrando a lógica de "finisher auto-protegido" que faz os 6 wincons atuais funcionarem tão bem aqui.

## Combos (se houver)

| Combo | Peças | Resultado | Confirmado por rulings? |
|---|---|---|---|
| Nenhum combo de 2 peças / infinito identificado | — | O deck não tem loop infinito nem trava de recursos — condizente com o power level "casual" do briefing e com a regra da diversão do pipeline | N/A |

Interações relevantes conferidas via `get_rulings` (não são combos, mas afetam o goldfishing):
- **Lux Artillery**: a contagem de 30+ marcadores é checada **duas vezes** — no início do end step (para disparar) e de novo na resolução (se caiu abaixo de 30 nesse meio-tempo, não faz nada). Ao testar, anote o total de marcadores **no fim do turno**, não durante o turno.
- **Alibou**: o X do dano/scry é recalculado **na resolução** do trigger e conta artefatos **tapados** no momento — se os atacantes tiverem vigilance ou forem removidos/destapados antes de resolver, X pode cair a 0. Vigilance não existe no deck atual, então isso não é risco, mas vale checar se algum finisher futuro trouxer vigilance.
- **Alibou** perde haste do time se for removido — atacantes que já estavam declarados permanecem no combate mesmo perdendo haste depois.

## Protocolo de goldfishing

1. **Preparação**: embaralhe as 100 cartas, compre 7. Aplique mulligan como faria numa partida real (Londres): mãos sem fonte de mana até o turno 3 ou sem nenhuma peça de tema/ramp jogável até o turno 4 são mulligan.
2. **Execução solo**: jogue turnos consecutivos sem oponente. Sequência-padrão: terreno → ramp disponível → Inspirit na curva (ideal: turno 3) → desenvolva o board (station no Inspirit quando fizer sentido, jogue token-makers e redutores de custo) → proliferate quando disponível.
3. **Registre por partida** (tabela abaixo):
   - Turno em que Inspirit resolve e fica em jogo.
   - Turno em que a mesa está "boa" (4+ artefatos + pelo menos 1 anthem/finisher em jogo).
   - Turno da vitória projetada: calcule dano letal combinado contra 3 oponentes de 40 de vida cada — some (a) dano de combate do time inteiro (via Caminho A/C) e (b) 10 de queima por oponente se Lux Artillery bateu 30 marcadores no end step (Caminho B).
   - Total de marcadores entre artefatos/criaturas no fim de cada turno a partir do turno 5 (para acompanhar o progresso do Caminho B).
   - Travas de mana (sem cor certa, preso em 2-3 terrenos).
   - Mãos mortas (7 cartas sem jogada relevante até o turno 4).
4. **Repita 5+ vezes.** Meta: vitória projetada **≤ turno 7** na maioria das partidas (3 de 5 ou mais).

### Tabela de registro sugerida

| Partida | Turno mulligan? | Turno Inspirit | Turno "mesa boa" | Turno vitória projetada | Caminho decisivo (A/B/C) | Travou mana? | Mão morta? |
|---|---|---|---|---|---|---|---|
| 1 | | | | | | | |
| 2 | | | | | | | |
| 3 | | | | | | | |
| 4 | | | | | | | |
| 5 | | | | | | | |

## Ajustes pós-teste (quando aplicável)

| Sai | Entra | Motivo |
|---|---|---|
| *(preencher após rodar o protocolo — ver diagnóstico abaixo)* | | |

**Guia de diagnóstico para o modo `post-goldfish`:**
- **Vitória tarde (turno 8+) na maioria das partidas** → considerar a entrada opcional de Chief of the Foundry (ver tabela acima) e revisar se o Caminho B está realmente disparando (30 marcadores é uma meta agressiva; se nunca bater, pode valer reforçar o pacote de proliferate em vez de contar com Lux Artillery como plano principal).
- **Trava de mana recorrente** → devolver ao `manabase-engineer` com os dados específicos (quais cores faltaram, em que turno).
- **Mão morta recorrente (mesmo após mulligan correto)** → revisar curva/draw com o `draw-specialist`; candidatos a corte são cartas que só fazem algo combadas com outra carta específica ainda não resolvida em jogo.
- **Comandante removido antes do turno 5 na maioria das partidas apesar da proteção** → sinal de que as 5 camadas de proteção (Padeem + Leonin Abunas + 3 one-shots, já implementadas na fase interação) precisam ser jogadas mais defensivamente (segurar mana aberta para Loran's Escape/Restoration Magic/Blacksmith's Skill antes de baixar o Inspirit em mesas com remoção instantânea suspeita).
