# 00 — Briefing · Inspirit, Flagship Vessel

- **Modo:** `registro` — **lista atual registrada em 2026-09-18** (ver "Situação")
- **Comandante:** Inspirit, Flagship Vessel — `{U}{R}{W}` · Legendary Artifact — Spacecraft
  (5/5 estacionado; 8+ ganha Flying)
  - Estático central: `Other artifacts you control have hexproof and indestructible` — protege tudo, menos ele próprio.
  - Gatilho `1+`: **desligado com 0 charge counters** (ruling Adagia, 2025-07-25) e mira
    `other target artifact` — **nunca se alimenta**.
- **Identidade de cor:** WUR (Jeskai)
- **Tema:** artefatos em massa (go-wide de thopters/golems/constructs) + marcadores de carga/+1/+1
  + proteção estática do comandante. Subtemas: proliferate, token-makers, redução de custo, affinity.
- **Status físico:** **montado** — confirmado pelo usuário em **2026-09-18**, quando ele passou a
  lista do deck físico. **Lista de referência: [`deck.md`](deck.md)** (raiz), transcrita dessa lista.
- **Orçamento:** ≤ R$200 (torneio) / ≤ R$350 (mesão), medido pelo **menor valor da LigaMagic**.
  **Custo da lista montada: R$ 255,80** — LigaMagic (menor), cotações de **2026-09-18**, informadas
  pelo usuário. Cabe no mesão; **estoura o teto de torneio em R$ 55,80**. As 80 cotações estão em
  `data/prices.tsv` (`mtgdb prices <nomes...>`); as 4 entradas sem preço são o comandante e os
  básicos.
- **Uso da coleção:** caixa primeiro, compras permitidas (padrão).

## Situação

A pendência que este briefing carregava desde 2026-09-18 — "a lista real não está no repositório" —
**está resolvida**. O usuário informou a lista do deck montado e ela virou
[`deck.md`](deck.md), a fonte de verdade. `mtgdb deck inspirit-flagship-vessel` resolve as 84
entradas sem faltantes.

Esta lista **não é saída de rodada do pipeline**: descende da v2 (versão mesão), passou pela v3 e
foi evoluída pelo próprio usuário depois. O diff medido **v2 → hoje** (16 entradas, 15 saídas +
1 Mountain) está em [`decisions.md`](decisions.md), com a ressalva de que **os motivos das mudanças
do usuário não foram declarados** — ao repropor qualquer carta cortada nesse intervalo, pergunte a
ele (regra 5).

## Retrato do registro (2026-09-18)

Contagem do que a lista montada tem hoje. É **medição, não diagnóstico** — nenhuma troca foi
proposta nem aprovada nesta passagem.

| Categoria | Atual | Meta | Situação |
|---|---|---|---|
| Draw | 14 | 12–13 | acima da meta |
| Ramp (rocks + explosivo) | 11 + 1 | 10–11 + 2–3 | explosivo abaixo (só Empowered Autogenerator) |
| Redutores de custo | 5 | — | Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector, Voyager Quickwelder, Cloud Key |
| Interação (remoção) | 9 | ~10 | perto da meta |
| Counterspells | **0** | — | Disruption Protocol e Stoic Rebuttal saíram; o deck ficou sem counter |
| Board wipes | 2 | 2–4 | no piso (Chain Reaction, Organic Extinction) |
| Proteção | 7 | — | 1 hexproof ao comandante (Padeem) após a saída de Leonin Abunas |
| Terrenos | **36** | base 38 | 2 abaixo; R é a cor mais magra (4 Mountain) |
| Wincons | 6 peças / 2 caminhos | 3+ caminhos | o caminho "queima aos 30 counters" sumiu com Lux Artillery |
| Fontes de proliferate | 5 | — | Metastatic Evangel, Surge Conductor, Kilo, Recon Craft Theta, Tezzeret's Gambit |
| Spacecraft / Veículos | 4 / 3 | — | 7 permanentes disputando corpos para station/crew |
| Custo (LigaMagic menor, 2026-09-18) | R$ 255,80 | ≤200 torneio / ≤350 mesão | R$ 55,80 acima do teto de torneio |

## Rodadas

| Rodada | Data | Estado |
|---|---|---|
| [v1](rounds/v1-2026-07-25/report.md) | 2026-07-25 | primeira auditoria completa (7 fases) |
| [v2](rounds/v2-2026-08-12/report.md) | 2026-07-25 → 2026-08-12 | segunda auditoria do zero sobre a v1; **duas listas** (torneio ≤ R$200 × mesão ≤ R$350) que compartilham 93 cartas e divergem em 7 slots — ver [`08-versoes.md`](rounds/v2-2026-08-12/08-versoes.md); repreço LigaMagic em [`09-precos-ligamagic.md`](rounds/v2-2026-08-12/09-precos-ligamagic.md) |
| [v3](rounds/v3-2026-08-13/NOTAS-RECONSTRUIDAS.md) | 2026-08-13 | 6 trocas aplicadas (pool fechado na coleção). Achou a **causa raiz** do deck: com 0 charge counters o gatilho `1+` do comandante está desligado e ele nunca se alimenta → as duas lentidões são a mesma. **O relatório original se perdeu**; o que há é um resumo reconstruído. |
| — (registro) | 2026-09-18 | **não é rodada**: transcrição da lista montada para [`deck.md`](deck.md) + diff v2→hoje em [`decisions.md`](decisions.md). Sem análise, sem trocas. |

Recuperadas do histórico do git (commit `84ddb87^`) em 2026-09-18 e reunidas num único deck: antes
viviam em dois diretórios irmãos (`inspirit-flagship-vessel` e `inspirit-flagship-vessel-v2`).
**Versão é pasta dentro do deck, nunca sufixo de slug** — sufixo parte o `decisions.md` e quebra
`mtgdb deck <slug>`.

## Pendências para a próxima rodada (v4)

1. **O goldfishing nunca foi rodado, desde a v1** — é a pendência mais antiga do deck. O protocolo
   estava em `07-wincons.md` da v3 e se perdeu junto com ela; terá de ser refeito.
2. **A lista estoura o teto de torneio: R$ 255,80 contra R$ 200.** Cotações de 2026-09-18 já
   registradas. Três cartas concentram R$ 58,43 (23% do deck): **Dispatch R$ 31,50**,
   **Cyberdrive Awakener R$ 14,94** e **Swords to Plowshares R$ 13,99**.
   ⚠️ **Dispatch subiu 1013% desde 12/08** (R$ 2,83 → 31,50) — é a maior oscilação já observada no
   banco. Reconferir na LigaMagic antes de usar esse número para cortar: se for estoque escasso e
   não preço real, o deck fica a ~R$ 227.
3. **Motivos das mudanças do usuário não estão registrados.** O primeiro item do intake da v4 é
   perguntar o porquê dos cortes de Leonin Abunas, Lux Artillery/Lux Cannon, dos dois counterspells
   e dos 2 terrenos — e por que 4 cortes da v3 foram desfeitos.
4. **`/update-collection` pendente.** As 15 cartas que saíram **não** foram lançadas na coleção:
   posse é declarada, nunca inferida (regra 7). O usuário informou em 2026-09-18 que **parte delas
   foi vendida** — a regra combinada é: **o que não voltar na próxima lista de `/update-collection`
   foi vendido**. Até lá, nenhuma delas conta como disponível.

## Alertas herdados da v2

- Todos os valores em **US$** e as conversões `× 5,5` do `report.md` da v2 são **estimativa por
  proxy da Scryfall e estão errados** (regra 2). A régua é o menor valor da LigaMagic.
- No fechamento da v2 o deck custava **≈ R$ 220** contra teto de **R$ 200** na versão torneio.
  Cotações de agosto/2026 — reconfira com `mtgdb prices` antes de usar.
