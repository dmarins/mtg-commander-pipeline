# 00 — Briefing · Krenko, Mob Boss

- **Modo:** `improve`
- **Data do intake:** 2026-08-21
- **Comandante:** Krenko, Mob Boss — `{2}{R}{R}` · 3/3 · Legendary Creature — Goblin Warrior
  - `{T}: Create X 1/1 red Goblin creature tokens, where X is the number of Goblins you control.`
- **Identidade de cor:** R (mono-vermelho)
- **Tema declarado pela lista:** Goblin tribal / enxame de tokens + burn
- **Tamanho atual:** 100 cartas exatas (comandante + 99), todas dentro da identidade

## Respostas do intake

### Objetivo da otimização
O usuário marcou **três objetivos ao mesmo tempo**:
1. **Fechar o jogo mais rápido** — converter os goblins em dano letal antes do turno 8.
2. **Consistência e resiliência** — mais draw real, mais ramp, recuperação pós-board wipe.
3. **Subir o power level** — aceita trocar filler por peças de nível mais alto e mudar bastante a lista.

Leitura do orquestrador: os três convergem no mesmo diagnóstico — a lista atual é larga em
cartas de efeito único e baixo impacto. Não há conflito entre eles nesta rodada.

### Queixas nas partidas (todas as quatro marcadas)
| Queixa | Tradução em alvo |
|---|---|
| Não fecho o jogo | Falta payoff que converta N goblins em dano — dano não passa em bloqueadores |
| Fico sem cartas na mão | Falta card advantage real (deck esvazia no turno 4–5) |
| Apanho de board wipe | Falta recuperação/proteção — Krenko + tokens morrem juntos |
| Mana / curva travando | 36 terrenos quase todos Mountain, ramp quase inexistente |

### Orçamento
**Sem teto rígido.** O usuário mantém uma **lista própria de cartas à parte** (várias cores)
que deve ser **consultada primeiro** — funciona como coleção estendida para esta rodada.
Só as cartas mono-R / incolores dessa lista são elegíveis (identidade do comandante).

**Lista recebida em 2026-08-22** — `sideboard.txt`, 47 cartas. Registradas em
`data/collection.tsv` com a nota `sideboard fisico`. Filtradas por identidade:
**34 elegíveis** (mono-R ou incolores) e **13 fora** (todas brancas).
Ficha completa das 34 em `00b-colecao-elegivel.md`.

Regra 7 em vigor: essas 34 têm **prioridade sobre qualquer compra**. Um especialista só pode
propor carta nova depois de dizer por que a peça equivalente da coleção não serve.
Regra 2 continua valendo para o que for comprado: preço = **menor valor da LigaMagic**,
nunca o USD da Scryfall.

### Cartas intocáveis
**Nenhuma.** Só Krenko é fixo. Qualquer outra carta pode sair se a ficha F1–F7 completa
justificar o corte.

### Alvos protegidos (Fase 0, item 3)
Nenhum alvo sob proteção explícita nesta rodada — o usuário abriu a lista inteira.
Zada, Hedron Grinder e Conspicuous Snoop foram **oferecidas** como intocáveis e **recusadas**;
elas seguem o protocolo de corte normal (ficha completa antes de qualquer veredito).

## Agregados da lista atual (`bin/mtgdb deck krenko-mob-boss`, 2026-08-21)

```
criaturas de verdade      32
veículos / spacecraft      0
artefatos (qualquer tipo)  3
corpos com poder ≥3        6
corpos que usam o {T}      7   (atrito com a ativação do próprio Krenko)
tipos    Artifact 3 · Creature 32 · Enchantment 9 · Instant 10 · Land 5 (36 cartas) · Sorcery 10
curva (não-terrenos)  1:14  2:20  3:21  4:7  5:1  6:0  7+:0
```

Observação de curva: **top-end zero**. Nenhuma carta acima de CMC 5, e a única CMC 5 é
Misty Mountains Raider. Isso é coerente com "não fecho o jogo": não há peça que transforme
o enxame em vitória.

## Decklist de entrada

Ver `lista.txt` (formato normalizado, sem códigos de set). 32× Mountain agrupados.
