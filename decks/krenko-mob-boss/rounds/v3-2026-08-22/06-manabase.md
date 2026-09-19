# 06 — Manabase · Krenko, Mob Boss

> Fase executada **inline pelo orquestrador** em 2026-08-22: o `manabase-engineer` caiu por limite
> de sessão antes de escrever qualquer coisa. Mesmo protocolo: oracle puxado nesta sessão,
> ficha F1–F7 em todo corte, cobertura declarada.

## Diagnóstico herdado (`02-theme.md` §4.6 e §8)

36 terrenos: 32 `Mountain` + 4 utility, **4 de 4 entrando virados**. Zero terrenos que produzam
mais de um mana, zero com relevância tribal, zero que deem haste. Veredito: *errado nos dois
sentidos ao mesmo tempo* — terreno demais para MV 2,38, e 0% dos slots livres viraram utility
num deck **mono-vermelho**, onde o custo de cor é trivial.

## Coleção (regra 7)

Conferido em `00b-colecao-elegivel.md`: das 34 cartas elegíveis, **nenhuma é terreno**. Não há o
que priorizar aqui; as duas entradas são compra.

## Número final: 35 terrenos (de 36)

| Fator | Efeito no número |
|---|---|
| MV 2,38 → 2,50 após as fases 2a/2b | pressiona para cima |
| `Urza's Incubator`: `{2}` a menos em **33** feitiços de criatura Goblin | pressiona forte para baixo |
| `Sol Ring`, `Arcane Signet`, `Skirk Prospector` (mana recorrente) | para baixo |
| Krenko custa **`{T}`**, não mana — o motor não consome mana | para baixo |
| `Hammer of Purphoros`, `Umbral Mantle`, `Purphoros`, `Castle Embereth` = 4 mana sinks | tolera excedente |

**35** é o número: um a menos que hoje, e o excedente deixou de ser carta morta porque agora há
onde gastá-lo — inclusive `Hammer of Purphoros`, que **converte terreno em corpo 3/3**.
Devolvo **1 slot de spell** ao orquestrador.

## Trocas

| Sai | Entra | Motivo |
|---|---|---|
| `Memorial to War` | `Den of the Bugbear` | MLD sai por **regra 11**, não por avaliação |
| `Looming Spires` | `Castle Embereth` | 0 pontos de sinergia na auditoria |
| `Mountain` (1 de 32) | — | 36 → 35; slot devolvido |

### `Den of the Bugbear` — o achado da fase
*"If you control two or more other lands, this land enters tapped. {T}: Add {R}. {3}{R}: Until end
of turn, this land becomes a 3/2 red Goblin creature with 'Whenever this creature attacks, create
a 1/1 red Goblin creature token that's tapped and attacking.' It's still a land."*

Quatro funções de uma vez: (a) **entra desvirado nos turnos 1–3**, justamente quando o deck quer
velocidade; (b) produz `{R}`; (c) é **mana sink**; (d) vira **Goblin 3/2 que fabrica Goblins** —
e, sendo terreno, **atravessa board wipe**. É recuperação pós-wipe que não ocupa slot de feitiço,
ligando as queixas nº1 e nº3 num terreno.

### `Castle Embereth`
*"This land enters tapped unless you control a Mountain. {T}: Add {R}. {1}{R}{R}, {T}: Creatures
you control get +1/+0 until end of turn."*

Com **31 Mountains** restantes, entra **desvirado** em praticamente 100% dos jogos — resolve o
defeito que a auditoria apontou nos 4 utility atuais. Anthem ativável que atinge o enxame inteiro
e mais um mana sink.

## Cortes — ficha F1–F7

### `Memorial to War` · Land
*"This land enters tapped. {T}: Add {R}. {4}{R}, {T}, Sacrifice this land: Destroy target land."*
- **F1** Entra virado; `{R}`; MLD por 5 manas + sacrifício. **F2** Não é corpo. **F3** Land (land drop). **F5** Produz `{R}` — irrelevante com 32 Mountains. **F6** **Entra virado**: custa meio turno em toda partida. **F7** **Viola a regra 11** (MLD); mana-negativa (5 manas e um terreno para atrasar 1 de 3 oponentes).

| Função | Quem cobre |
|---|---|
| Land drop + fonte de `{R}` | `Den of the Bugbear` |
| Destruir terreno | **Descoberta de propósito** — é a regra da diversão, não um custo lamentado |

### `Looming Spires` · Land
*"This land enters tapped. When this land enters, target creature gets +1/+1 and gains first strike until end of turn. {T}: Add {R}."*
- **F1** Entra virado; pump de **um turno** em um alvo; `{R}`. **F2** — **F3** Land. **F5** Pump temporário — e **não é contador**, logo não gera "modified" para `Goro-Goro`. **F6** Entra virado. **F7** Efeito de 1 turno pelo custo de 1 turno.

| Função | Quem cobre |
|---|---|
| Land drop + `{R}` | `Castle Embereth` (que entra **desvirado** com Mountain) |
| Pump pontual | `Castle Embereth` faz `+1/+0` no time **todo turno**, repetível |
| First strike | **Descoberta.** Declarado: `Assault on Osgiliath` dá double strike, que contém first strike. |

### `Mountain` ×1
- **F1/F5** Produz `{R}`. **F6** Sempre desvirado. **F7** Nenhum — é o corte mais limpo possível.
- Função: land drop. Coberta pelos 31 restantes. Sai porque 36 é um terreno a mais do que a curva pede.

## Por que `Forgotten Cave` e `The Autonomous Furnace` **ficam**

Ambas entram viradas, e eu quase as cortei por isso. A ficha diz outra coisa: as duas **trocam
terreno por carta** (cycling `{R}` e `{1}{R},{T}`, sac: compra 1). O usuário reclamou
explicitamente de **flood**; estas são as duas únicas cartas do deck que resolvem flood no próprio
slot de terreno. Cortá-las por "entra virado" seria julgar por um eixo só — exatamente o erro que
o checklist §1 existe para impedir. Ficam.

## Estado final

| | Antes | Depois |
|---|---|---|
| Terrenos | 36 | **35** |
| Básicos | 32 | **31** |
| Utility | 4 | **4** (`Forgotten Cave`, `The Autonomous Furnace`, `Den of the Bugbear`, `Castle Embereth`) |
| Utility que entram **virados** | 4 de 4 (100%) | **2 de 4** (50%) — e as duas que ficam viradas são as anti-flood |
| Terrenos com relevância tribal Goblin | 0 | **1** (`Den of the Bugbear`) |
| Terrenos que são **mana sink** | 0 | **2** |
| Terrenos que sobrevivem ao wipe **e refazem tabuleiro** | 0 | **1** |
| Slots de spell devolvidos | — | **1** |
