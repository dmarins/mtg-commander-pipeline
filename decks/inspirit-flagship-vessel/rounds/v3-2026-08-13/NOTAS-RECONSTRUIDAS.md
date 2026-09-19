# v3 · 2026-08-13 — notas reconstruídas

> ⚠️ **Isto não é o relatório original.** A pasta `decks/inspirit-flagship-vessel-v3/` existiu em
> disco (com `report.md` e `07-wincons.md`) mas **nunca foi commitada**, e se perdeu. Não há como
> recuperá-la do git — `git rev-list --all` não tem nenhum objeto `*v3*`.
>
> O que segue foi reconstruído em 2026-09-18 a partir das anotações de sessão. É **resumo**, não o
> documento original: use como pista, confirme antes de tratar como registro.

## A causa raiz achada nesta rodada

Ruling oficial (Adagia, WotC 2025-07-25): *"A station symbol means 'As long as this permanent has
N or more charge counters on it, it has [abilities]'"*.

Com **0 contadores**, o Inspirit tem **só** a estática — o gatilho `1+` está **desligado**. E ele
mira `up to one **other** target artifact`, ou seja, **nunca se alimenta**. Proliferate e Deepglow
Skate só multiplicam onde já existe contador.

**Conclusão:** as duas dores do usuário (lento para acumular charge counters, lento para
tripular/estacionar as naves) são **a mesma dor**. O deck não tem economia de contadores; tem
economia de **poder**, porque quase todo contador vem de tapar criatura via Station.

## 6 trocas aplicadas (pool fechado na coleção, sem compra)

| Sai | Entra |
|---|---|
| Sphere of the Suns | **Solar Array** |
| Unwanted Remake | **Paladin Danse, Steel Maverick** |
| Organic Extinction | **Golem Foundry** |
| Crystalline Crawler | **Metastatic Evangel** |
| Cloud Key | **Emissary Escort** |
| Swords to Plowshares | **Stone by Sunlight** |

A última foi **proposta do próprio usuário**; ele rejeitou a contraproposta de cortar Loran's Escape,
porque quer manter uma camada extra de proteção do comandante.

**Fraqueza conhecida, registrada na época:** o arranque frio depende de **Solar Array, que é 1-of** —
esperada em campo antes da 1ª nave em 1–2 de cada 10 partidas. As outras 4 trocas (oferta de corpo)
são consistentes.

## Questão deixada aberta em 2026-08-13

**`Rip Apart` → `Sunder the Gateway`** — recomendada, **não aplicada**. Mesmo MV e mesma velocidade,
troca o modo de 3 dano por um corpo permanente 2/2 artefato (4/4 sob os anthems) via Incubate 2, e
nunca é carta morta. **Custo declarado:** Rip Apart é o **único removedor dedicado de planeswalker**
do deck.

Pode ter sido resolvida pelo usuário na versão que ele montou sozinho depois — **perguntar**.

## Achados de regra que valem para qualquer rodada futura

- **Sunburst conta as cores efetivamente gastas, inclusive no custo genérico** → Uthros `{2}{U}` pago com W+R+U entra com 3 charge counters, direto no limiar `3+`.
- **Uthros no `3+` se auto-carrega** (`draw a card. Put a charge counter on this Spacecraft` por artifact spell) — é a nave mais fácil de ligar, não a mais difícil. O `12+` é inalcançável; ignorar.
- **Veículo tripulado é criatura destapada** → Brotherhood Vertibird tripulado (poder = nº de artefatos, 10–14) pode em seguida ser tapado para Station. É a linha mais rápida para o `8+`. Crew e Station cabem na mesma fase principal, e Station ignora enjoo de invocação.
- **Blast Zone não recebe o gatilho `1+`** — o gatilho mira `target artifact` e ela é terreno. Só proliferate chega nela.
- **`Dispatch` é o Swords to Plowshares do deck** — metalcraft é trivial com 38 artefatos, e não dá vida ao oponente.
- **O gatilho `1+` mira UM artefato por combate** e há 9+ receptáculos disputando. Toda carta que só *recebe* contador agrava o gargalo — foi por isso que Long-Range Sensor e Ratchet Bomb foram rejeitadas.
- **Lux Artillery**: gastar contadores em efeito (Golem do Golem Foundry, Lux Cannon, Bankbuster, Pentad Prism) **subtrai** do limiar de 30 — os caminhos A e B competem pelo mesmo recurso.
- **`Pentad Prism` não tem `{T}` no custo** (`Sphere of the Suns` tem) — o Prism é armazenamento, não sorvedouro.

## Armadilhas de processo registradas na época

- Classificações que precisaram de correção em `deck.md`: **Rip Apart é feitiço**, **Thirst for Knowledge é instantâneo**, **Thousand Moons Smithy é MDFC cuja face traseira é Land** (quebra filtro por type line).
- `mtgdb deck <slug>` prioriza `deck.md` e **conta as cartas citadas na tabela de mudanças**, inflando os agregados. Para verificação automática confiável, use uma lista limpa (`lista.txt`).
