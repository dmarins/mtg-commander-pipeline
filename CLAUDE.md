# MTG Commander Pipeline

Pipeline de subagentes para **construção e otimização de decks de Commander (EDH)**, baseado no processo de 7 passos: comandante → cartas temáticas → draw → ramp → interação → terrenos e cortes → condições de vitória e testes.

## Como usar

- `/build-deck [tema ou comandante]` — constrói um deck novo do zero.
- `/improve-deck [caminho da decklist ou link do Archidekt]` — audita e otimiza um deck existente.
- `/update-collection [lista de cartas]` — substitui a coleção pela lista informada; o que não estiver nela sai.

A sessão principal atua como **orquestradora** (especialista em Commander): coleta preferências, delega cada fase a um subagente especialista, apresenta os resultados para revisão do usuário e consolida o deck.

## Subagentes

| Agente | Fase | Meta quantitativa |
|---|---|---|
| `commander-scout` | 1. Escolha do comandante | 3–5 opções analisadas |
| `theme-analyst` | 2. Cartas temáticas | pool de 35–45 candidatas com sinergias sobrepostas |
| `draw-specialist` | 3. Vantagem de cartas | **12–13 fontes** de card advantage real |
| `ramp-specialist` | 4. Aceleração | **10–11 ramps padrão** + 2–3 explosivos |
| `interaction-specialist` | 5. Interação | **~10 peças** + **2–4 board wipes** |
| `manabase-engineer` | 6. Terrenos e cortes | **38 terrenos** (base) + cortes até 99 cartas |
| `wincon-tester` | 7. Vitória e testes | 3+ condições de vitória + protocolo de goldfishing |

Uma carta pode contar para mais de uma categoria (ex.: criatura temática que compra cartas), mas o deck final precisa bater todas as metas.

## Fluxo da rodada — do comando ao fechamento

```
comando inicial  →  pipeline (cada subagente escreve o relatório da sua fase)
                 →  CONSOLIDAÇÃO: o orquestrador funde tudo em UM report.md
                 →  validação do usuário
                 →  OK? fecha a rodada.  Não? devolve à fase responsável e reconsolida.
```

**O entregável da rodada é `rounds/v<N>-<data>/report.md`, e é um só.** Os arquivos de fase
(`02-theme.md` … `07-wincons.md`) são matéria-prima do orquestrador e material de consulta
pontual — **o usuário não deve precisar abrir nenhum deles para validar o deck**. Terminar o
pipeline entregando N arquivos e um resumo em chat é entregar pela metade.

O `report.md` precisa, obrigatoriamente:
- **nomes de carta clicáveis** para a ficha da LigaMagic (regra 8) — é assim que o usuário
  confere preço e texto carta a carta na hora de validar;
- **bloco de importação em texto puro** (padrão MTG Online), para colar no Moxfield/Archidekt e
  fazer playtest;
- a lista final completa, o que saiu com o motivo, o custo real em reais e as **pendências**;
- rodar `bin/linkify <report.md>` no fim — a linkagem é mecânica, não se faz à mão.

**Na devolução, refaça só o pedaço.** Se o usuário reprovar uma carta ou uma categoria, o
orquestrador reinvoca **apenas o especialista daquela fase**, com o motivo da reprovação, e
reconsolida o mesmo `report.md`. Não se reinicia o pipeline inteiro.

**Não produza painéis, dashboards, artifacts ou HTML** a menos que o usuário peça. O formato
do processo é markdown versionado no repositório, ao lado do deck. Link publicado não entra no
git, não sobrevive à rodada e não é onde o usuário trabalha.

## Estado persistente por deck

Cada deck vive em `decks/<slug>/` (slug = nome do comandante em kebab-case):

```
decks/<slug>/
├── 00-briefing.md      # VIVO — modo (build/improve/registro), comandante, identidade de cor,
│                       # tema, palavras-chave, orçamento, power level, uso da coleção
│                       # (caixa primeiro × só sobressalentes), decklist atual,
│                       # **status físico** (montado × só no papel + data + lista de referência)
│                       # e o **índice de rodadas** (tabela: rodada, data, estado, link)
├── deck.md             # VIVO — lista consolidada (fonte de verdade; é o que `mtgdb deck` lê)
├── decisions.md        # VIVO — registro cronológico de cortes e entradas (regra 5)
└── rounds/
    └── v<N>-<AAAA-MM-DD>/   # uma pasta por rodada, criada na Fase 0 e fechada no relatório
        ├── 01-commander.md  # opções de comandante e escolha (se aplicável)
        ├── 02-theme.md      # análise linha a linha + pool temático
        ├── 03-draw.md       # candidatas de draw
        ├── 04-ramp.md       # candidatas de ramp
        ├── 05-interaction.md# candidatas de interação e wipes
        ├── 06-manabase.md   # terrenos + plano de cortes
        ├── 07-wincons.md    # condições de vitória + protocolo de goldfishing
        └── report.md        # relatório da rodada
```

**A linha é a natureza do arquivo, não a rodada.** O que é *estado* fica na raiz e é único; o que é
*retrato de um momento* vai para `rounds/v<N>-<data>/` e nunca mais é tocado. Em particular,
`decisions.md` **não** se divide por rodada: a regra 5 exige **um** lugar para consultar antes de
repropor uma carta, e conferência espalhada por N arquivos é conferência que não acontece.
A data da pasta é a do **fechamento** da rodada. Não existe `report.md` na raiz — o relatório
corrente é o da última pasta, e o índice de rodadas do briefing aponta para ele.

Subagentes **leem** `00-briefing.md` e `deck.md`, **escrevem** apenas o arquivo da sua fase — dentro da pasta da rodada corrente, cujo caminho o orquestrador passa na invocação — e devolvem um resumo curto. Só o orquestrador atualiza `deck.md`, `decisions.md` e o índice de rodadas do briefing.

### Formato de `deck.md`

Tabela por seção (Comandante, Criaturas, Artefatos, Encantamentos, Instantâneos, Feitiços, Planeswalkers, Terrenos):

```
| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
```

`Categorias` usa os rótulos: `tema`, `draw`, `ramp`, `remoção`, `proteção`, `counter`, `wipe`, `wincon`, `terreno`.

## Regras não negociáveis

1. **Toda busca no Scryfall** segue `references/scryfall-search-guide.md` — leia antes de buscar. Sempre inclua `legal:commander` e `id<=<identidade do comandante>`. **Nunca use filtro de preço na busca** (`usd<`, `eur<`, `tix<`): o Scryfall é repositório de *cartas*, não de preço (regra 2). Para controlar volume, use `order:edhrec` e os limites do guia — a triagem por custo vem depois, com cotação da LigaMagic.

   **EDHREC e Archidekt** (pelo MCP `mtg`) são fonte de **candidatas e de calibragem, nunca de veredito**: carta que vem de lá passa pelas regras 3, 4 e 6 como qualquer outra. O `theme-analyst` consulta o EDHREC **uma vez por rodada** e grava o `Radar do EDHREC` no `02-theme.md`, que as fases seguintes leem; o `manabase-engineer` compara a contagem de terrenos com listas do Archidekt no bracket do briefing. Moxfield e `get_edhrec_combos` estão quebrados na v2.2.1, e `validate_deck` não confere identidade nem banimento — estado de cada ferramenta e regras de uso em "Fontes de meta" do guia de busca.

2. **Preço vem da LigaMagic; do Scryfall vem a carta.** Os dois papeis não se misturam: Scryfall (e o banco local `mtgdb`) responde o que a carta é — oracle, tipo, CMC, cores, tags, rulings; preço é **exclusivamente** o **menor valor da LigaMagic**.

   **Nenhum número de preço do Scryfall entra em lugar nenhum** — nem em filtro de busca, nem em tabela, nem em total, nem como "estimativa" ou "ordem de grandeza". A ferramenta `mcp__mtg__get_card_price` (USD do Scryfall convertido para BRL por câmbio) está **fora de uso** (e bloqueada em `.claude/settings.local.json`), os operadores `usd<`/`eur<`/`tix<` não entram em query, e se um campo de preço vier junto de outro retorno, ignore-o. Os erros medidos em 2026-08-12 chegaram a **6,5× para mais** (Restoration Magic: proxy R$ 1,65 × real R$ 10,75) e **14× para menos** (Chief of the Foundry: proxy R$ 1,16 × real R$ 0,08) — nos dois sentidos, sem fator de correção possível. Um proxy que erra para menos também **descarta carta barata**: por isso o filtro sai da busca, e não só do total.

   **Como cotar**: consulte primeiro o que já foi capturado com `mtgdb prices <nomes...>`. O que faltar, confira em `https://www.ligamagic.com.br/?view=cards/card&card=<Nome+Em+Ingles>` e use o **primeiro** dos três números do bloco "Preço Médio de Venda no Marketplace" (menor / médio / maior). A página carrega preço via JS — WebFetch não funciona, use as ferramentas de browser. Registre toda cotação nova com `mtgdb prices -add "<carta>" <valor>` — preço é observação datada, nunca um valor que se sobrescreve. `mtgdb prices -volatile` mostra quais cartas de fato oscilam, que são as únicas que precisam ser reconferidas antes de um torneio.

   **Como rotular**: todo total leva origem **e** idade — `LigaMagic (menor), cotações de <data>`. Carta ainda não cotada aparece como **`a cotar`**, nunca com número estimado; um relatório pode fechar com pendência de cotação, mas não com preço inventado.

3. **Sinergia sobreposta**: só recomende carta que tenha **2+ pontos de sinergia** com o comandante e/ou com outras cartas já escolhidas. Justifique cada recomendação. Evite cartas que só funcionam isoladas.

4. **Ficha completa antes de qualquer veredito** — leia `references/card-evaluation-checklist.md` antes de recomendar ou cortar qualquer carta. Um deck de Commander é multifacetado: a mesma carta costuma exercer 3–5 funções ao mesmo tempo (efeito escrito, corpo que pode ser tapado para crew/station/convoke/improvise, tipo que alimenta contagens, receptor de contadores e anthems, redutor de custo, fixação de cor). **Julgar por um aspecto só é a causa raiz do vai-e-vem de cartas entre rodadas.** Nenhum corte é proposto sem enumerar por escrito todas as funções da carta e nomear quem cobre cada uma. Frases como "criatura fraca" ou "corpo pequeno" não são justificativa de corte — são sinal de que a ficha não foi feita. O mesmo critério usado para defender a entrada vale para a carta que sai.

5. **Registro de decisão** — todo corte e toda entrada vão para `decks/<slug>/decisions.md`. Antes de propor carta que já esteve no deck, consulte o registro: a proposta precisa dizer quem cortou, por quê, e **o que mudou desde então**. Se o motivo original continua válido, a carta não volta.

6. **Puxe o texto oracle na hora, sempre.** Nunca julgue carta de memória — nem as do próprio deck. Use **`bin/mtgdb`** (banco local com o bulk data do Scryfall — ver `references/mtgdb.md`): `mtgdb oracle "<nome>" ...` para cartas e `mtgdb deck <slug>` para o deck inteiro. Caia para o MCP `mtg` só quando a carta for mais nova que o último dump. Se o banco não existir, rode `make db` (~15 s).

7. **Coleção pessoal primeiro — prioridade de análise, não obrigação de uso.** Toda carta que o usuário já possui é avaliada **antes** de qualquer compra: consulte com `mtgdb collection <nomes...>` (fonte: `data/collection.tsv`, atualizada por `/update-collection`). Mas possuir a carta não a torna elegível: se a peça da coleção não servir ao deck, **comprar é a decisão correta**. O que a regra exige é que a coleção seja *considerada primeiro* e que a dispensa seja *justificada por escrito* — o especialista que propõe uma compra precisa nomear a carta equivalente da coleção e dizer por que ela não cobre a função. Nunca force uma carta ruim no deck só porque ela já está na caixa, e nunca proponha compra sem ter olhado a caixa.

   **Posse é declarada, nunca inferida.** Se o deck está montado fisicamente está no campo
   `Status físico` do `00-briefing.md`, com a data da confirmação e qual arquivo é a lista real
   (`report.md` ou `lista.txt`). `lista.txt` e `deck.md` registram o deck **aprovado** — sozinhos
   não dizem se as cartas foram compradas. Deck sem o campo: **pergunte ao usuário**, não deduza.

   **Deck montado é coleção fechada.** Carta não migra de um deck para outro: se o Thorin usa um Sol Ring e o Tori também precisa de um, são duas cópias compradas. Por isso `data/collection.tsv` lista **só as sobressalentes** — cartas que saíram de um deck durante a otimização, ou que foram compradas e não entraram em nenhum. Uma carta que está dentro de outro deck **não está disponível** e não conta como "já possuo"; para usá-la aqui, compra-se outra. O que sai de um deck numa otimização vira sobressalente **se a carta existir fisicamente**; o que o usuário vende sai da coleção (`/update-collection`).

   **Não está num deck e não está nas sobressalentes = não existe.** Corte de otimização não entra na coleção automaticamente: boa parte das cartas que saem de um deck nunca foi comprada — foram sugestões aprovadas no papel. Registrá-las como sobressalentes inventaria cartas físicas e faria `mtgdb collection` mentir justamente na pergunta que ele existe para responder. **Nenhum processo infere posse**; a única fonte é a lista que o usuário passa em `/update-collection`.

   **Vale para construção e otimização, nas duas pontas.** Em `/build-deck` e em `/improve-deck`, as sobressalentes são varridas **antes** de qualquer busca no Scryfall — é o que faz o deck sair mais barato. A dispensa precisa ser **justificada por escrito**, nomeando a sobressalente e dizendo por que ela não cobre a função. Os motivos mais frequentes são falta de sinergia (menos de 2 pontos, regra 3) e régua de custo, mas **a lista não é fechada**: curva, cor, velocidade, tipo que não alimenta as contagens do deck — qualquer eixo da ficha F1–F7 serve, desde que dito. O que não vale é o veredito sem ficha ("carta fraca", "corpo pequeno" — regra 4). Havendo justificativa, **comprar é a decisão correta** — a caixa é o ponto de partida da busca, não o teto dela.

   **Modo restrito é opt-in.** Deck montado só com sobressalentes — sem nenhuma compra — só quando o usuário **pedir explicitamente**. O intake pergunta; o `00-briefing.md` registra. Sem pedido explícito, o padrão é caixa primeiro **com compras permitidas**, e nenhum especialista deve se auto-restringir às sobressalentes.

8. **Nomes de cartas sempre em inglês** (nome oficial do Scryfall). Textos, análises e conversa em **português (Brasil)**.
   Em **todo entregável que o usuário lê** — `report.md` à frente —, todo nome de carta é
   **link clicável para a ficha da LigaMagic**. Não faça à mão: rode **`bin/linkify <arquivo>`**,
   que valida cada nome contra o banco local, usa o nome canônico completo na URL mesmo quando o
   texto mostra o apelido, e preserva o bloco de exportação em texto puro. `--check` lista sem
   escrever. O formato é: `[**Nome**](https://www.ligamagic.com.br/?view=cards/card&card=<nome+percent-encoded>)`, com espaço virando `+` e o resto em percent-encoding UTF-8 (`Krenko%2C+Mob+Boss`). Única exceção: o bloco de exportação padrão MTG Online, que é texto puro. Formato completo em `references/deck-report-template.md`.

9. **O usuário decide**: nenhuma carta entra em `deck.md` sem aprovação dele. A revisão acontece
   **uma vez, sobre o `report.md` consolidado** — não fase a fase. Parar a cada especialista para
   pedir seleção de cartas soltas é pedir decisão sem contexto: o usuário precisa ver o deck
   inteiro, com custo e lista importável, para conseguir julgar qualquer parte.

10. O deck final tem exatamente **100 cartas** (comandante + 99), todas dentro da identidade de cor e legais no formato.

11. Regra da diversão: evite pacotes que impedem os oponentes de jogar (MLD, stax pesado, lock infinito) a menos que o usuário peça explicitamente.

## Referências

- `references/card-evaluation-checklist.md` — **ficha de funções F1–F7, protocolo de corte e registro de decisão** (regras 4 e 5). Leitura obrigatória antes de recomendar ou cortar carta.
- `references/mtgdb.md` — **banco local de cartas, tags e rulings** (`bin/mtgdb`). É por onde passam oracle, busca, tags, rulings, preços e coleção.
- `references/scryfall-search-guide.md` — sintaxe, tags confirmadas, receitas de busca, controle de volume, **estado das ferramentas do MCP `mtg`** e **fontes de meta** (EDHREC, Archidekt).
- `bin/linkify` — **linkagem automática de nomes de carta para a LigaMagic** (regra 8).
  `bin/linkify report.md` escreve; `bin/linkify --check report.md` só reporta.
- `references/deck-report-template.md` — template do relatório final, incluindo o **formato dos links de carta para a LigaMagic** (regra 8).
