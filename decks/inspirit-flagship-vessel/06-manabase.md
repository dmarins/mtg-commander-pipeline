# Manabase e Cortes — Inspirit, Flagship Vessel

## Cálculo

CMC médio dos 63 não-terrenos atuais: **3,03** (X-spells contados como CMC 0: Hangarback Walker, Everflowing Chalice).
Draws + ramps com mv ≤ 2: **10** (Reckoner Bankbuster [draw, mv2]; Etherium Sculptor, Enthusiastic Mechanaut, Everflowing Chalice, Sol Ring, Arcane Signet, Sphere of the Suns, Talisman of Progress, Azorius Signet, Pentad Prism [ramp, mv≤2]).

Fórmula: 31,42 + 3,13×3,03 − 0,28×10 = **38,1 → 38 terrenos**

O deck está hoje em **36 terrenos**, exatamente 2 abaixo do recomendado pela fórmula — o que bate com a dor nº 4 relatada ("deck lento / sensação de pouco terreno"): não é sensação, é déficit real de 2 fontes. Recomendo fechar em **38**, financiado por 2 cortes de não-terreno (ver Parte 3).

## Diagnóstico da manabase atual (36 terrenos)

- **Básicos**: 6 Island, 5 Mountain, 6 Plains (17) — praticamente empatados, mas a demanda de pips do deck **não é equilibrada**: contando símbolos coloridos nos 63 não-terrenos, U=24 (46%), W=16 (31%), R=12 (23%). Os básicos atuais subrepresentam Island frente à real demanda azul.
- **Não-básicos (19)**: mistura de painland/filterland/checkland untapped (Battlefield Forge, Rugged Prairie, Skycloud Expanse, Glacial Fortress, Clifftop Retreat, Port Town — bons, majoritariamente destravados) + fixadores tricolor sempre úteis (Command Tower, Spire of Industry — excelente aqui, quase sempre há artefato em jogo —, Exotic Orchard) + **9 terrenos que entram virados**: Irrigated Farmland, Mystic Monastery, Boros Garrison, as 3 Temples e as 3 Bridges (Rustvale/Razortide/Silverbluff). Quase metade (47%) do pacote não-básico entra tapped — isso é o principal motor real da "sensação de deck lento", mais até do que a contagem bruta de terrenos.
- **Pior oferensor**: **Boros Garrison** — tapland *e* devolve um terreno seu à mão (perda líquida de tempo, pior que um tapland comum). Não tem função alguma que os outros fixadores WR (Battlefield Forge, Rugged Prairie, Clifftop Retreat, Rustvale Bridge) não cubram sem o penalty extra.
- **Redundância evitável**: as 3 Temples (Epiphany UR, Enlightenment WU, Triumph WR) empilham tapland+scry nos mesmos pares de cor já cobertos por painlands/filterlands/checklands untapped. Manter todas as 3 é redundância pura na dupla WR (coberta 4x sem contar a Temple).

## Terrenos recomendados (swap)

| Sai | Entra | Produz | Entra virado? | Sinergia/Utilidade | Faixa | US$ |
|---|---|---|---|---|---|---|
| Boros Garrison | Blast Zone | C | Não | Wipe-em-terreno: entra com 1 charge counter; `{X}{X},{T}` acumula mais; sacrifica para destruir tudo com MV = counters. Escala com o pacote de proliferate do deck (Kilo, Surge Conductor, Tezzeret's Gambit, Deepglow Skate) — cada proliferate empurra o gatilho para uma faixa de MV diferente sem gastar mana extra. Nota: o trigger 1+ do Inspirit **não** alimenta (é terreno, não artefato-alvo). 4º wipe do deck (meta 2–4, já atendida com 3 — este é bônus de utilidade, não contagem obrigatória) | torneio | 0,38 |
| Temple of Triumph | — (some, vira básico) | WR | tapped+scry | Corte de redundância: WR já fixado por Battlefield Forge, Rugged Prairie, Clifftop Retreat e Rustvale Bridge, todos untapped ou quase-sempre-untapped. Cortar a Temple reduz 1 terreno tapped sem perder cor. | — | — |
| — | +3 Island | U | Não | Ajusta a base para a real demanda de pips (U=46% dos símbolos coloridos, hoje subrepresentado nos básicos) | — | 0 |

**Resultado**: não-básicos 19 → 18 (−Boros Garrison, −Temple of Triumph, +Blast Zone); básicos 17 → 20 (Island 6→9, Mountain 5, Plains 6). Total: **38 terrenos** (18 não-básicos + 20 básicos), proporção de pips U:W:R ≈ 9:6:5 (~46/31/23%, batendo com a contagem de símbolos).

Terrenos tapped caem de 9/19 (47%) para 7/18 (39%) do pool não-básico — mantidas Irrigated Farmland e Mystic Monastery (cycling de segurança contra flood e fixação das 3 cores num só terreno, respectivamente) e as 2 Temples restantes (Epiphany, Enlightenment) e as 3 Bridges (artefatos — hexproof/indestrutível do próprio Inspirit, alimentam Spire of Industry e Cyberdrive Awakener). Se o usuário quiser reduzir ainda mais o tapped-count, a próxima candidata natural de corte seria Irrigated Farmland por mais um básico.

**Na coleção?**: sem arquivo de coleção informado no briefing — não verificável; preços acima são referência Scryfall (USD) para conversão manual em R$ (≈ ×5,5) contra o teto de cada faixa.

## Plano de cortes (deck em 99 → mantém 99, abrindo espaço para +2 terrenos)

| Corte proposto | CMC | Motivo |
|---|---|---|
| Voyager Quickwelder | 3 | 5º efeito redundante de "artefatos custam {1} a menos" (junto com Etherium Sculptor, Enthusiastic Mechanaut, Foundry Inspector, Cloud Key) — redundância de categoria acima do necessário; corpo 2/4 sem impacto de tabuleiro por si só. |
| Emissary Escort | 2 | 0/4 base que só cresce se **outro** artefato de MV alto já estiver em jogo — carta clássica de "não faz nada sozinha na mão" (dor nº 2 relatada pelo usuário); sinergia única (poder alto alimenta Station) sem 2º ponto de sinergia sólido. |

Esses 2 cortes financiam integralmente os 2 terrenos extras (63 não-terrenos → 61; 36 terrenos → 38; 61+38=99 ✓).

### Swap opcional (resolve a pendência de ramp explosivo — separado do financiamento acima, líquido zero em cartas)

| Sai | Entra | CMC | US$ | Faixa | Motivo |
|---|---|---|---|---|---|
| Golem Foundry | Gilded Lotus | 3→5 | 0,46 | torneio | Golem Foundry é um motor lento (3 charges → Golem 3/3) com sinergia real (recebe charge do trigger 1+ do Inspirit, além do cast-trigger), mas o deck já tem fartura de motores de charge counter (Coretapper, Everflowing Chalice, Sphere of the Suns, Pentad Prism, Empowered Autogenerator, Lux Cannon/Artillery). Gilded Lotus fecha a meta de ramp explosivo (1→2, meta 2–3) com {T}: 3 mana de uma cor — direto, sem depender de acumular counters, e resolve peças de curva alta (Alibou, Kappa Cannoneer, Cyberdrive Awakener, Thought Monitor). |

Esse swap é **opcional** e independente da manabase (1 corte + 1 entrada = líquido zero em não-terrenos); repasso para aprovação do usuário junto com os 2 cortes obrigatórios.

## Curva final projetada (61 não-terrenos, após os 2 cortes obrigatórios)

0–1: 9 · 2: 14 · 3: 18 · 4: 12 · 5: 4 · 6+: 4

(Se o swap opcional Golem Foundry → Gilded Lotus for aprovado, desloca 1 carta do bucket 3 para o bucket 5: 3→17, 5→5, mantendo os demais.)
