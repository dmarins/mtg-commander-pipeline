---
description: Atualiza data/collection.tsv a partir de uma lista nova — o que não estiver na lista sai da coleção
argument-hint: [caminho do arquivo com a lista, ou cole a lista]
---

Você atualiza a **coleção física** do usuário. Argumento recebido: `$ARGUMENTS`.

A coleção são as **cartas sobressalentes**: o que saiu de um deck numa otimização,
ou foi comprado e não entrou em nenhum. Decks montados são coleções fechadas —
suas cartas não entram aqui, porque não migram entre decks (precisar da mesma
carta em dois decks significa comprar duas cópias).

A lista recebida é o **retrato atual das sobressalentes**, não um acréscimo: toda
carta que está em `data/collection.tsv` e **não** aparece na lista sai do arquivo.
É uma substituição, e é por isso que a fase de conferência abaixo não é opcional.
Vender cartas é o caminho normal de saída.

Só `data/collection.tsv` é escrito (o espelho em `data/scryfall.db` é recarregado
junto). `data/prices.tsv` **não** é tocado: preço é série append-only, e carta que
saiu da caixa continua sendo candidata a compra — a cotação dela segue valendo.

## 1 — Obtenha a lista

Em ordem de preferência:

- `$ARGUMENTS` é um caminho de arquivo → use direto.
- `$ARGUMENTS` é a lista colada → grave num arquivo temporário do scratchpad da
  sessão (nunca dentro do repositório) e use esse caminho.
- `$ARGUMENTS` está vazio → peça a lista ao usuário e pare até recebê-la. **Não
  invente, não reaproveite uma lista anterior e não sincronize sem lista** — uma
  lista errada aqui apaga a coleção.

Formato aceito é o mesmo das decklists, sem conversão manual: exportação MTGO
(`1 Bombard (EOE) 129`), `1x Nome`, um nome por linha, ou lista em markdown.
Linhas vazias e comentários com `#` são ignorados.

## 2 — Simule antes de escrever

```bash
bin/mtgdb collection -sync -file <caminho> -dry-run
```

Se o binário não existir ou reclamar de schema, `make db` (~15 s).

## 3 — Apresente o diff e espere aprovação

Mostre ao usuário, com os números do `-dry-run`:

- **Entram** (N) — nomes.
- **Saem** (N) — nomes **com a nota e a data originais**, que é o que ele perde
  ao confirmar. Essa lista vai inteira, nunca resumida como "e mais X cartas":
  o ponto da conferência é ele reconhecer cada carta que está saindo.
- **Quantidade alterada**, quando houver.
- **Avisos de nome não resolvido**, quando houver. Trate-os antes de aplicar:
  quase sempre é erro de grafia ou nome em português (regra 8 — nomes em inglês),
  e gravar assim cria uma carta fantasma que nenhuma busca futura encontra.

Se o número de saídas for grande em relação à coleção, diga isso explicitamente e
pergunte se a lista está completa — uma exportação truncada tem exatamente essa
cara. Espere a aprovação. **Nada é escrito antes dela** (regra 9).

## 4 — Aplique

```bash
bin/mtgdb collection -sync -file <caminho>
```

Cartas que permanecem mantêm `note` e `added_at` — são anotação manual do
usuário, e uma lista de nomes não as carrega. Quem entra recebe a data de hoje;
use `-note "<texto>"` se o usuário deu um rótulo para essa leva.

## 5 — Feche

Confirme o total final (`bin/mtgdb collection -list | tail -1`) e lembre que
`data/collection.tsv` é versionado — ofereça o commit, sem fazê-lo por conta
própria.
