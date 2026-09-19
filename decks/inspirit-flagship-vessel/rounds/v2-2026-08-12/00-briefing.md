# Briefing — Inspirit, Flagship Vessel (v2 — segunda auditoria)

## Modo
`improve` — segunda rodada de auditoria e otimização, agora em cima do deck **já otimizado na v1** (`decks/inspirit-flagship-vessel/`), que ainda não foi testado em partidas/goldfishing. O usuário quer uma auditoria completa do zero antes de investir dinheiro nas cartas da v1.

> Esta é uma pasta separada da v1 — **não sobrescreve** `decks/inspirit-flagship-vessel/` (briefing, análises e relatório da primeira rodada permanecem intactos como registro histórico).

## Comandante
- **Inspirit, Flagship Vessel** — {U}{R}{W} (CMC 3)
- Legendary Artifact — Spacecraft, 5/5 quando estacionado (8+ ganha Flying)
- **Station** (vira outra criatura sua: põe marcadores de carga iguais ao poder dela; só como feitiço; vira criatura-artefato com 8+)
- **1+**: no início do combate no seu turno, põe um marcador +1/+1 **ou** dois marcadores de carga em até um outro artefato alvo
- **Efeito estático central: "Other artifacts you control have hexproof and indestructible"** — protege tudo, menos ele próprio.
- Identidade de cor: **W/U/R (Jeskai)** → buscas Scryfall: `legal:commander id<=wur`

## Tema
Artefatos em massa (go-wide de thopters/golems/constructs) + marcadores de carga/+1/+1 + proteção estática do comandante. Subtemas: proliferate, token-makers, redução de custo, affinity. (Ver `decks/inspirit-flagship-vessel/02-theme.md` da v1 para a análise linha a linha original do comandante — ainda válida, não precisa refazer essa parte se não houver mudança de entendimento.)

## Objetivo da otimização (v2)
**Auditoria completa do zero**, com olhar cético, sobre as 100 cartas já resultantes da v1 — confirmar que nenhuma troca da v1 ficou frouxa, identificar sinergias adicionais não vistas antes, e validar se o deck está pronto para investimento antes do usuário comprar as cartas pendentes.

## Dores relatadas pelo usuário
Nenhuma nova — o deck da v1 ainda **não foi jogado/testado** (goldfishing pendente). Esta auditoria é preventiva, não reativa a partidas reais.

## Orçamento (REGRA IMPORTANTE — dupla faixa, mantida da v1)
- **Torneio**: custo **total do deck ≤ R$ 200,00** (preços de referência: LigaMagic).
- **Mesão/1x1**: teto total de **R$ 350,00**.
- **Toda troca sugerida deve ser rotulada `torneio` ou `mesão`.** Sugestões de torneio: preferir `usd<0.5`, idealmente `usd<0.25` (conversão aproximada US$1 ≈ R$5,5).
- Nas buscas Scryfall com orçamento: incluir `usd<X` conforme a faixa.

## Coleção / restrições
- Sem arquivo de coleção. **Sem cartas intocáveis** — qualquer uma das 99 pode sair de novo com justificativa, inclusive as que acabaram de entrar na v1.

## Decklist atual (v1 otimizada — 99 + comandante)

```
Commander
1 Inspirit, Flagship Vessel

Deck
1 Hangarback Walker
1 Coretapper
1 Etherium Sculptor
1 Third Path Iconoclast
1 Enthusiastic Mechanaut
1 Brotherhood Vertibird
1 Foundry Inspector
1 Kilo, Apogee Mind
1 Malcator, Purity Overseer
1 Master of Etherium
1 Pinnacle Emissary
1 Sai, Master Thopterist
1 Surge Conductor
1 Chrome Host Seedshark
1 Uthros Research Craft
1 Warmaker Gunship
1 Jhoira, Weatherlight Captain
1 Padeem, Consul of Innovation
1 Leonin Abunas
1 Crystalline Crawler
1 Alibou, Ancient Witness
1 Deepglow Skate
1 Kappa Cannoneer
1 Cyberdrive Awakener
1 Thought Monitor
1 Everflowing Chalice
1 Sol Ring
1 Arcane Signet
1 Sphere of the Suns
1 Talisman of Progress
1 Chainsaw
1 Empowered Autogenerator
1 Glass Casket
1 Pentad Prism
1 Azorius Signet
1 Cloud Key
1 Perilous Snare
1 Reckoner Bankbuster
1 Midnight Clock
1 Gilded Lotus
1 Lux Artillery
1 Lux Cannon
1 Thousand Moons Smithy
1 Thopter Spy Network
1 Whirlwind of Thought
1 Saheeli, Sublime Artificer
1 Dispatch
1 Swords to Plowshares
1 Loran's Escape
1 Blacksmith's Skill
1 Restoration Magic
1 Unwanted Remake
1 Disruption Protocol
1 Thirst for Knowledge
1 Rip Apart
1 Stern Lesson
1 Tezzeret's Gambit
1 Chain Reaction
1 Fumigate
1 Reverse Engineer
1 Organic Extinction
1 Command Tower
1 Spire of Industry
1 Exotic Orchard
1 Battlefield Forge
1 Rugged Prairie
1 Skycloud Expanse
1 Port Town
1 Glacial Fortress
1 Clifftop Retreat
1 Irrigated Farmland
1 Mystic Monastery
1 Temple of Epiphany
1 Temple of Enlightenment
1 Blast Zone
1 Rustvale Bridge
1 Razortide Bridge
1 Silverbluff Bridge
1 Perilous Landscape
6 Plains
9 Island
5 Mountain
```

Contagem verificada (herdada do `report.md` da v1): **38 terrenos** (18 não básicos + 20 básicos: 6 Plains, 9 Island, 5 Mountain) + **61 não-terrenos** = 99 cartas + comandante = 100. ✓

## Referência à v1

- Registro completo de todas as trocas já aplicadas: `decks/inspirit-flagship-vessel/deck.md` (seção "Registro de mudanças") e `decks/inspirit-flagship-vessel/report.md` (seção "Mudanças Aplicadas").
- Se esta v2 propuser desfazer ou re-trocar algo que veio da v1, justifique explicitamente por que a escolha anterior não se sustenta na auditoria cética.
