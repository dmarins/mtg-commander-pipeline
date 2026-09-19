# Registro de decisões — Inspirit, Flagship Vessel

> Arquivo aberto em 2026-09-18, na reconstrução do deck a partir do histórico do git.
> As rodadas v1 e v2 foram conduzidas **antes** de a regra 5 existir e não produziram registro
> cronológico de cortes e entradas. **Nada foi inventado aqui**: o histórico delas é o que está
> escrito nos relatórios de cada rodada.

- **v1 (2026-07-25)** — trocas e justificativas em [`rounds/v1-2026-07-25/report.md`](rounds/v1-2026-07-25/report.md).
- **v2 (2026-07-25 → 2026-08-12)** — trocas e justificativas em [`rounds/v2-2026-08-12/report.md`](rounds/v2-2026-08-12/report.md);
  divergência torneio × mesão em [`rounds/v2-2026-08-12/08-versoes.md`](rounds/v2-2026-08-12/08-versoes.md).
- **v3 (2026-08-13)** — 6 trocas aplicadas. O relatório original **se perdeu** (a pasta existiu em
  disco e nunca foi commitada); resumo reconstruído em
  [`rounds/v3-2026-08-13/NOTAS-RECONSTRUIDAS.md`](rounds/v3-2026-08-13/NOTAS-RECONSTRUIDAS.md).
  Ficou **aberta** a troca `Rip Apart` → `Sunder the Gateway` (custo declarado: Rip Apart é o único
  removedor dedicado de planeswalker do deck).
- **pós-v3** — o usuário evoluiu o deck por conta própria. O que ele montou **não está documentado**
  e será registrado quando ele passar a lista atual.

⚠️ **Consequência para a regra 5:** ao propor uma carta que já esteve neste deck, o "quem cortou e
por quê" precisa ser procurado nos dois relatórios acima — e, para o período pós-v2, **perguntado ao
usuário**, porque não há registro. A partir da v3 este arquivo passa a ser append-only por rodada.

---

## Registro pós-v3 — lista montada informada em 2026-09-18

O usuário passou a lista do deck **físico**. Não houve rodada de pipeline: é **registro**, não
otimização. As mudanças abaixo foram feitas por ele, fora do processo, entre 2026-08-13 e
2026-09-18. **O motivo de cada uma não foi declarado** — o que está aqui é o *fato* medido, não a
justificativa. Para a regra 5, tratar como "cortada pelo usuário, motivo desconhecido": ao
repropor qualquer uma delas, **pergunte a ele** antes.

> ⚠️ O diff é **v2 → hoje**, não v3 → hoje: o relatório da v3 se perdeu e não existe lista dela em
> disco. As 6 trocas da v3 estão dentro deste intervalo e aparecem misturadas às do usuário.

### Saíram (15 cartas + 1 Mountain)

| Carta | Função que exercia na v2 | Observação |
|---|---|---|
| Blast Zone | wipe escalável em terreno | só proliferate a alimentava (o gatilho 1+ mira artefato) |
| Boros Signet | fixação R/W | entrou na v2 justamente para socorrer o R; hoje o R é a cor mais magra (4 Mountain) |
| Disruption Protocol | counter | **o deck ficou com 0 counterspells** |
| Stoic Rebuttal | counter | idem |
| Fumigate | wipe assimétrico | wipes caíram de 3 (com Blast Zone) para 2 |
| Glass Casket | remoção por exílio | — |
| Leonin Abunas | 2ª camada de hexproof no comandante | sobrou só Padeem cobrindo o Inspirit |
| Lux Artillery | wincon B (queima aos 30 counters) | o caminho B de vitória deixou de existir |
| Lux Cannon | remoção recorrente por charge counters | — |
| Reckoner Bankbuster | draw recarregável pelo gatilho 1+ | — |
| Rip Apart | única remoção dedicada de planeswalker | a troca aberta da v3 **foi resolvida**: `Sunder the Gateway` está no deck e `Rip Apart` não |
| Sphere of the Suns | ramp com charges | corte da v3 (→ Solar Array) |
| Thirst for Knowledge | draw 3 descartando artefato | — |
| Unwanted Remake | remoção instantânea por {W} | corte da v3 (→ Paladin Danse, que **não** está na lista final) |
| Temple of Epiphany | terreno U/R | — |
| Mountain (5 → 4) | fonte de R | — |

### Entraram (16 cartas)

| Carta | Função | Observação |
|---|---|---|
| Astral Cornucopia | ramp escalável em charge counters | — |
| Chief of the Foundry | anthem de criaturas-artefato | era a peça da versão **torneio** da v2; virou fixa |
| Dawnsire, Sunstar Dreadnought | remoção/wincon (100 de dano no 10+) | 4º spacecraft — aumenta a disputa por corpos para Station |
| Emissary Escort | corpo grande e barato para Station | corte→entrada da v3 (saiu Cloud Key) — mas **Cloud Key voltou** |
| Ethersworn Sphinx | draw via cascade + 4/4 voador por affinity | — |
| Golem Foundry | fábrica de Golems 3/3 por charge counters | entrada da v3 (saiu Organic Extinction) — mas **Organic Extinction voltou** |
| Marketback Walker | corpo X + draw ao morrer | — |
| Metastatic Evangel | proliferate por criatura nontoken | entrada da v3 (saiu Crystalline Crawler) — mas **Crystalline Crawler voltou** |
| Recon Craft Theta | proliferate ao atacar + traz o próprio tripulante | 3ª fonte repetível de contadores |
| Solar Array | mana + sunburst (arranque frio de contadores) | entrada da v3 (saiu Sphere of the Suns) |
| Spring-Loaded Sawblades | remoção de criatura virada em flash; craft → veículo | — |
| Stone by Sunlight | remoção modal / indestrutível + artificializa | entrada da v3 (saiu Swords to Plowshares) — mas **Swords to Plowshares voltou** |
| Sunder the Gateway | remoção de artefato/encantamento + Incubate 2 | resolve a questão aberta de 2026-08-13 |
| Vivid Crag | terreno de qualquer cor com charge counters | — |
| Voyage Home | draw 3 + 3 de vida por affinity | — |
| Voyager Quickwelder | redutor de custo nº 5 | — |

### Cartas que voltaram depois de cortadas na v3

`Cloud Key`, `Crystalline Crawler`, `Organic Extinction` e `Swords to Plowshares` foram cortadas na
v3 e **estão de volta** na lista montada. Nenhuma das entradas da v3 que as substituiu foi desfeita
— elas convivem. `Paladin Danse, Steel Maverick` é a **única** entrada da v3 que não chegou à lista
final. Motivos não declarados; **perguntar antes de tratar esses cortes da v3 como válidos** (regra 5).
