# 00 — Briefing · Inspirit, Flagship Vessel

- **Modo:** `registro` — **lista atual pendente** (ver "Situação")
- **Data do registro:** 2026-09-18
- **Comandante:** Inspirit, Flagship Vessel — `{U}{R}{W}` · Legendary Artifact — Spacecraft
  (5/5 estacionado; 8+ ganha Flying)
  - Estático central: `Other artifacts you control have hexproof and indestructible` — protege tudo, menos ele próprio.
- **Identidade de cor:** WUR (Jeskai)
- **Tema:** artefatos em massa (go-wide de thopters/golems/constructs) + marcadores de carga/+1/+1
  + proteção estática do comandante. Subtemas: proliferate, token-makers, redução de custo, affinity.
- **Status físico:** **montado** — o usuário confirmou em 2026-09-18 que o deck existe e
  "melhorou muito". **A lista real não está neste repositório.**

## Situação

⚠️ **Não existe `deck.md` nem `lista.txt` na raiz, e isso é proposital.** O usuário evoluiu o deck
por conta própria depois da v2; a lista da `rounds/v2-2026-08-12/` **não** corresponde ao deck físico
de hoje. Gravá-la na raiz faria `mtgdb deck inspirit-flagship-vessel` devolver uma lista errada —
a mesma armadilha do `lista.txt` (regra 7: posse é declarada, nunca inferida).

O usuário disse que vai registrar a lista atual em breve. Até lá, o que existe aqui é **histórico**.

## Rodadas

| Rodada | Data | Estado |
|---|---|---|
| [v1](rounds/v1-2026-07-25/report.md) | 2026-07-25 | primeira auditoria completa (7 fases) |
| [v2](rounds/v2-2026-08-12/report.md) | 2026-07-25 → 2026-08-12 | segunda auditoria do zero sobre a v1; **duas listas** (torneio ≤ R$200 × mesão ≤ R$350) que compartilham 93 cartas e divergem em 7 slots — ver [`08-versoes.md`](rounds/v2-2026-08-12/08-versoes.md); repreço LigaMagic em [`09-precos-ligamagic.md`](rounds/v2-2026-08-12/09-precos-ligamagic.md) |
| [v3](rounds/v3-2026-08-13/NOTAS-RECONSTRUIDAS.md) | 2026-08-13 | 6 trocas aplicadas (pool fechado na coleção). Achou a **causa raiz** do deck: com 0 charge counters o gatilho `1+` do comandante está desligado e ele nunca se alimenta → as duas lentidões são a mesma. **O relatório original se perdeu** (nunca foi commitado); o que há é um resumo reconstruído. |

Recuperadas do histórico do git (commit `84ddb87^`) em 2026-09-18 e reunidas num único deck: antes
viviam em dois diretórios irmãos (`inspirit-flagship-vessel` e `inspirit-flagship-vessel-v2`).
**Versão é pasta dentro do deck, nunca sufixo de slug** — sufixo parte o `decisions.md` e quebra
`mtgdb deck <slug>`.

## Próximo passo

1. Usuário fornece a lista atual → vira `deck.md` na raiz (ou `lista.txt`).
2. `/improve-deck decks/inspirit-flagship-vessel/...` abre a **v4** em `rounds/`.

Duas coisas a perguntar no intake dessa rodada:

- **O goldfishing nunca foi rodado, desde a v1** — é a pendência mais antiga do deck. O protocolo
  estava em `07-wincons.md` da v3 e se perdeu junto com ela; terá de ser refeito.
- **`Rip Apart` → `Sunder the Gateway`** ficou como questão aberta em 2026-08-13. Ele pode ter
  resolvido sozinho — confirmar antes de repropor (regra 5).

## Alertas herdados da v2

- Todos os valores em **US$** e as conversões `× 5,5` do `report.md` da v2 são **estimativa por
  proxy da Scryfall e estão errados** (regra 2). A régua é o menor valor da LigaMagic.
- No fechamento da v2 o deck custava **≈ R$ 220** contra teto de **R$ 200** na versão torneio.
  Cotações de agosto/2026 — reconfira com `mtgdb prices` antes de usar.
