# 00 — Briefing · Tori D'Avenant, Fury Rider

- **Modo:** `registro` (otimização ainda não iniciada)
- **Data do registro:** 2026-09-18
- **Comandante:** Tori D'Avenant, Fury Rider — `{1}{R}{R}{W}` · 3/3 · Legendary Creature — Human Knight
- **Identidade de cor:** RW
- **Tamanho atual:** 100 cartas exatas (comandante + 99), conferido em `lista.txt`
  (30 básicos: 15 Mountain + 15 Plains)
- **Status físico:** **montado** — deck construído pelo usuário **fora do pipeline**.
  `lista.txt` é a lista real das cartas que estão na caixa, não uma proposta.
  Carta deste deck **não está disponível** para outros decks (regra 7).

## Rodadas

Nenhuma. A primeira otimização abrirá a `v1` em `rounds/`.

## Situação

O deck está montado e jogável. A otimização está **bloqueada por dependência**, não por
falta de interesse:

> O usuário tem cartas sobressalentes que **ainda não constam** em `data/collection.tsv`.
> Otimizar antes de registrá-las faria os especialistas buscarem no Scryfall o que já está
> na caixa — o oposto da regra 7 (coleção primeiro), e o deck sairia mais caro do que precisa.

**Ordem correta:** `/update-collection <lista de sobressalentes>` → depois
`/improve-deck decks/tori-d-avenant-fury-rider/lista.txt`.

## Pendente para o intake de otimização

Nada foi coletado ainda. Quando a otimização começar, a Fase 0 do `/improve-deck` precisa
levantar: objetivo da otimização, o que incomoda nas partidas, orçamento para novas cartas,
uso das sobressalentes (padrão: caixa primeiro com compras permitidas) e cartas intocáveis.
