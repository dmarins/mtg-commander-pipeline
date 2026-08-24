# Template de Relatório de Otimização e Feedback do Deck

Preencha todas as seções. Nomes de cartas em inglês; análises em português (Brasil).

## Links de carta (obrigatório)

Todo nome de carta citado no relatório é um **link clicável para a ficha da carta na LigaMagic**:

    [**Krenko, Mob Boss**](https://www.ligamagic.com.br/?view=cards/card&card=Krenko%2C+Mob+Boss)

- **Base da URL**: `https://www.ligamagic.com.br/?view=cards/card&card=`
- **Encoding**: nome oficial em inglês; espaço vira `+` e todo caractere fora de `A–Z a–z 0–9 - _ . ~`
  vira percent-encoding UTF-8. Equivale a `urllib.parse.quote_plus(nome)`.
  - `Krenko, Mob Boss` → `Krenko%2C+Mob+Boss`
  - `Urza's Incubator` → `Urza%27s+Incubator`
  - `Grishnákh, Brash Instigator` → `Grishn%C3%A1kh%2C+Brash+Instigator`
  - `Goblin-town Flunkies` → `Goblin-town+Flunkies` (hífen é literal)
- **Onde aplicar**: *toda* menção a carta — tabela de mudanças (colunas "Sai" e "Entra"), diagnóstico,
  sinergias, combos, wincons, custos, pendências e a lista de cartas por tipo. Vale também para cartas
  que **saíram** do deck, que foram descartadas ou que estão só registradas para a próxima rodada.
- **Onde não aplicar**: dentro do bloco de exportação padrão MTG Online (tem de ser texto puro para
  importar em MTGO/Moxfield/Archidekt), em nomes de arquivo/caminho e nos rótulos de categoria.
- **Nunca use crase em nome de carta.** `` `Impact Tremors` `` vira link, não `code`.
- **Carta de dupla face**: a URL usa o nome completo (`Goblin Glasswright // Craft with Pride` →
  `Goblin+Glasswright+%2F%2F+Craft+with+Pride`); o texto do link pode ser só a face frontal.
- **Apelidos e formas curtas** também levam link para a carta completa: `Krenko` → `Krenko, Mob Boss`,
  `Purphoros` → `Purphoros, God of the Forge`.
- Terreno básico repetido leva um único link: `**31×** [**Mountain**](...)`.
- O link é **ficha da carta, não cotação**: ele não substitui a regra 2 do `CLAUDE.md` — preço continua
  sendo o menor valor do bloco "Preço Médio de Venda no Marketplace", conferido no navegador e
  registrado com data via `mtgdb prices -add`.

## Informações Gerais do Deck

- **Comandante**: [Nome do Comandante]
- **Cores do Deck**: [Cores]
- **Número de Cartas**: 100 (99 + Comandante)
- **Foco do Deck**: [Tema/Sinergia Principal, Palavras-chave]
- Quantidade de cartas por tipo:
  - **Comandante**: 1
  - **Subcomandantes**: [Quantidade]
  - **Planeswalkers**: [Quantidade]
  - **Criaturas**: [Quantidade]
  - **Artefatos**: [Quantidade]
  - **Instantâneos**: [Quantidade]
  - **Feitiços**: [Quantidade]
  - **Encantamentos**: [Quantidade]
  - **Terrenos Básicos por Cor**: [Quantidade] (cor), [Quantidade] (cor), ...
  - **Terrenos Não-Básicos**: [Quantidade]
  - **Tokens**: [Quantidade]
- Quantidade de Draw, Ramp e Interação:
  - **Draw**: [Quantidade]
  - **Ramp**: [Quantidade]
  - **Interação**: [Quantidade]
  - **Board Wipes**: [Quantidade]

## Curva de Mana

- **Número de Terrenos Básicos**: [Quantidade]
- **Número de Terrenos Não-Básicos**: [Quantidade]
- **CMC Médio** (sem terrenos): [Valor]
- **Distribuição de CMC**:
  - 0–1: [Quantidade]
  - 2: [Quantidade]
  - 3: [Quantidade]
  - 4: [Quantidade]
  - 5: [Quantidade]
  - 6+: [Quantidade]
- **Distribuição de CMC por Cor**:
  - [Cor]: [Distribuição]
- **Quantidade de Cartas por Cor**:
  - [Cor]: [Quantidade] ([percentual do total do deck])

## Sinergias e Estratégias

- **Estratégia Geral**: [Como o deck pretende vencer — agressão, controle, combo, etc.]
- **Principais Sinergias**: [Combos e interações mais importantes entre cartas]
- **Cartas-Chave**: [Cartas essenciais para a estratégia]
- **Pontos Fortes**: [Consistência, sinergias, velocidade, etc.]
- **Pontos Fracos**: [Vulnerabilidades, dependências, lacunas de interação, etc.]

## Combos e Condições de Vitória

- **Combos Principais**: [Como funcionam e quais cartas envolvem]
- **Condições de Vitória**: [Como o deck fecha o jogo — dano, combo, controle, etc.]

## Lista de Cartas por Tipo

Para cada carta: `[**Nome**](link-ligamagic) — CMC, Atributos, Habilidades resumidas`.

- **Comandante**: [...]
- **Subcomandantes**: [...]
- **Planeswalkers**: [...]
- **Criaturas**: [...]
- **Artefatos**: [...]
- **Instantâneos**: [...]
- **Feitiços**: [...]
- **Encantamentos**: [...]
- **Terrenos**: [...]

## Lista de Cartas para Exportação (padrão MTG Online)

**Sem links aqui** — este bloco é consumido por importadores e precisa ser texto puro.

```
1 [Nome da carta]
1 [Nome da carta]
...
[N] [Terreno básico]

1 [Comandante]
```

Formato: uma linha por carta (`quantidade nome-em-inglês`), terrenos básicos agrupados com a quantidade total, linha em branco e o comandante por último (slot de comandante/sideboard).
