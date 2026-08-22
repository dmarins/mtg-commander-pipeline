# 05 — Proteção e recuperação · Krenko, Mob Boss

> Fase executada **inline pelo orquestrador** em 2026-08-22: o `interaction-specialist` caiu por
> limite de sessão antes de escrever. Mesmo protocolo: oracle puxado nesta sessão, ficha F1–F7 em
> todo corte, cobertura declarada, simetria de critério.

## Diagnóstico herdado (`02-theme.md` §4.4) — o buraco nº 1 do deck

**Proteção: 0.** Recuperação: 2,5. A auditoria isolou o fato que define a fase: a habilidade de
Krenko é **ativada** — ela passa pela pilha, e qualquer remoção em resposta apaga o turno inteiro.
São **dois** problemas distintos: (1) fazer a ativação resolver; (2) voltar a existir depois do wipe.

## Correção de rota antes das entradas

`Mudbutton Torchrunner` estava na minha lista de cortes. **Retirei-o.** Oracle: *"When this
creature dies, it deals 3 damage to any target."* Na auditoria ele era "um 1/1 por 3 manas" porque
o deck **não tinha sac outlet nenhum**. Nesta rodada entraram três — `Goblin Bombardment`,
`Skirk Prospector` e `Goblin Chirurgeon`. Ele deixou de depender de morrer por acaso e virou
**3 de dano dirigível a qualquer momento**, mais os gatilhos de `Boggart Shenanigans` e
`Rundvelt Hordemaster`. O contexto mudou; o veredito muda com ele. `Redcap Thief` (pool de ramp,
vago desde a colisão do `Skirk Prospector`) entrou no lugar dele na lista de cortes.

## Entradas (5)

| # | Carta | CMC | Marca | Por quê |
|---|---|---|---|---|
| 1 | `Deflecting Swat` | 3 | **NÚCLEO** | *"If you control a commander, you may cast this spell without paying its mana cost. You may choose new targets for target spell or ability."* **Custo real zero** com Krenko em campo. Não só salva Krenko da remoção — **devolve** o Swords ao comandante do oponente. Também redireciona *habilidade*, o que cobre ativações. |
| 2 | `Goblin Chirurgeon` | 1 | **NÚCLEO** | *"Sacrifice a Goblin: Regenerate target creature."* Proteção **repetível e grátis** — o custo é um token que Krenko refaz. Regenera contra remoção por destruição **e contra board wipe de destroy**. É corpo Goblin, sac outlet, e CMC 1. Não salva de exílio nem de −X/−X: declarado. |
| 3 | `Swiftfoot Boots` | 2 | **NÚCLEO** | Hexproof **e** haste. Hexproof (não shroud) é o certo aqui: seus próprios `Kick in the Door` e `Fists of Flame` continuam podendo mirar Krenko. Artefato → combustível do Baron e alvo do `Goblin Engineer`. |
| 4 | `Lightning Greaves` | 2 | compl. | Haste + shroud, **equip {0}**: protege Krenko **no mesmo turno em que ele entra**, sem gastar mana. Complementar às Boots, não redundante — equip {0} contra hexproof-que-aceita-seus-alvos são coberturas diferentes. |
| 5 | `Chandra, Torch of Defiance` | 4 | compl. | **Recuperação:** planeswalker sobrevive a board wipe e, sozinha, refaz vantagem — `+1` exila e joga (ou 2 dano a cada oponente), `+1` adiciona `{R}{R}` (mana sink), `−3` mata criatura. Cobre draw, ramp, remoção e resiliência num permanente que o Wrath não toca. |

## Como a recuperação pós-wipe fica coberta

O usuário reclamou de perder tudo num Wrath. A resposta desta rodada **não** é uma carta de
recursão — é o deck ter deixado de guardar todo o seu valor em criaturas:

| Permanente que sobrevive a um wipe de criaturas | Origem |
|---|---|
| `Purphoros, God of the Forge` (indestructible, e não-criatura com devoção <5) | fase 2a |
| `Impact Tremors`, `Shared Animosity`, `Goblin Bombardment` (encantamentos) | fase 2a |
| `Hammer of Purphoros`, `Urza's Incubator`, `Umbral Mantle`, `Skullclamp`, `Idol of Oblivion` | fases 2b e 3 |
| `Chandra, Torch of Defiance` (planeswalker) | esta fase |
| `Den of the Bugbear` (terreno que vira Goblin 3/2 e gera tokens) | fase 6 |
| `Squee, Goblin Nabob` (volta todo upkeep) | já no deck |
| Redeploy barato: `Krenko's Command`, `Hordeling Outburst`, `Goblin Surprise`, `Goblin War Party` | já no deck |

Antes: 2,5 fontes. Depois: **12 permanentes atravessam o wipe**, e três deles (`Impact Tremors`,
`Purphoros`, `Goblin Bombardment`) transformam o próprio wipe em dano — os Goblins que morrem
disparam `Boggart Shenanigans` e `Rundvelt Hordemaster` na saída.

## Cortes — ficha F1–F7 e protocolo §2

### `Misty Mountains Raider` {4}{R} · Goblin Soldier 4/4 · CMC 5
*"Whenever you attack, amass Goblins 2."*
- **F1** Amass 2 por ataque. **F2** **Maior corpo do deck (4/4)**. **F3** Goblin. **F4** Anthems/contadores. **F6** T5 — **único 5-drop**. **F7** Amass empilha no **mesmo Army**: contribui **+1** ao X de Krenko, não +2 por ataque.

| Função | Quem cobre |
|---|---|
| Maior corpo / sobrevive a ping | `Goblin Chieftain` torna **todo** token 2/2; `Hobgoblin Bandit Lord` 2/3; `Purphoros` 6/5 |
| Amass Goblins (tornava o Army um Goblin, salvando `Assault on Osgiliath`) | **Descoberto.** Custo aceito: `Goblin-town Flunkies` continua amassando **Goblins**, então o Army segue Goblin quando existir. |
| Único 5-drop | Perder é ganho: o deck precisava descer a curva, não subir. |

### `Punishing Fire` {1}{R} · Instant · CMC 2
*"Deals 2 damage to any target. Whenever an opponent gains life, you may pay {R}: return this from your graveyard to your hand."*
- **F1** 2 dano; recursão **condicionada a lifegain adversário**. **F6** T2. **F7** Sem lifegain na mesa é um Shock de 2 num formato de 40 vidas.

| Função | Quem cobre |
|---|---|
| Remoção pontual barata | `Blitz of the Thunder-Raptor`, `Smite the Deathless`, `Goblin Cratermaker`, `Ember Hauler`, `Fanatical Firebrand`, `Mudbutton Torchrunner` (agora sacrificável à vontade) — a categoria continua com **10 peças** |
| Recursão do cemitério | **Descoberta.** Declarado: dependia de um oponente ganhar vida, o que é fora do seu controle. |

### `Pinecone Strike` {1}{R} · Instant · CMC 2
*"Choose one or both — deals 3 damage to target creature (exile if it would die); destroy target artifact token."*
- **F1** Modal: 3 dano com exílio, e/ou destrói **token de artefato**. **F6** T2. **F7** O segundo modo é situacional (só Treasure/Clue/Food alheios).

| Função | Quem cobre |
|---|---|
| 3 dano + exílio | `Smite the Deathless` (3 dano, remove indestrutível **e** exila) — cobertura melhor, já no deck |
| Destruir token de artefato | **Descoberta.** Custo declarado: `Goblin Cratermaker` cobre artefato **não-token** incolor; token de artefato fica sem resposta. |

### `Cosmotronic Wave` {3}{R} · Sorcery · CMC 4
*"Deals 1 damage to each creature your opponents control. Creatures your opponents control can't block this turn."*
- **F1** Mini-wipe assimétrico de 1 + "não podem bloquear". **F6** T4. **F7** 1 de dano não responde nada depois do turno 4; e o modo "can't block" perde valor num plano que **não passa pelo combate**.

| Função | Quem cobre |
|---|---|
| Mini-wipe assimétrico x/1 | `Goblin Chainwhirler` (permanece, e é corpo Goblin) |
| Alfa-strike / remover bloqueadores | `Impact Tremors`, `Purphoros`, `Goblin Bombardment`, `Hobgoblin Bandit Lord` — o dano deixa de precisar atravessar bloqueador. `Assault on Osgiliath` segue como alfa-strike. |

### `Redcap Thief` {2}{R} · Goblin Rogue 2/3 · CMC 3
*"When this creature enters, create a Treasure token."*
- **F1** 1 Treasure ao entrar. **F2** Corpo 2/3. **F3** Goblin; Treasure é artefato (Baron). **F4** Anthems. **F6** T3 — no CMC mais engarrafado do deck. **F7** 3 manas por um 2/3 e um Treasure de uso único.

| Função | Quem cobre |
|---|---|
| Corpo Goblin | `Goblin Chirurgeon` (CMC 1, Goblin) — troca **para baixo** na curva |
| Treasure / combustível do Baron | `Hammer of Purphoros` (Golem 3/3 **todo turno**, renovável), `Reckless Lackey`, `Sol Ring`/`Arcane Signet`, `Skullclamp`, `Idol of Oblivion` — de one-shot para renovável |
| Ramp pontual | `Skirk Prospector` (recorrente), `Urza's Incubator` (redução em 33 cartas) |

## Contagem resultante

| | Antes | Depois |
|---|---|---|
| **Proteção** | **0** | **4** (`Deflecting Swat`, `Goblin Chirurgeon`, `Swiftfoot Boots`, `Lightning Greaves`) |
| Recuperação / permanentes que atravessam wipe | 2,5 | **12** |
| Remoção pontual | 10 | **10** (mesma contagem: saem 2, `Mudbutton Torchrunner` vira remoção repetível e `Chandra −3` entra) |
| Wipes assimétricos | 2 | **1** (`Goblin Chainwhirler`) — declarado |
| Resposta a encantamento | 0 | **0** — **lacuna que permanece aberta.** Mono-vermelho paga caro por isso; as opções (`Chaos Warp`, `Vandalblast`) não couberam na cota de 5. Fica registrado para a próxima rodada. |
