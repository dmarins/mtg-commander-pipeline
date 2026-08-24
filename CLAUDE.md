# MTG Commander Pipeline

Pipeline de subagentes para **construção e otimização de decks de Commander (EDH)**, baseado no processo de 7 passos: comandante → cartas temáticas → draw → ramp → interação → terrenos e cortes → condições de vitória e testes.

## Como usar

- `/build-deck [tema ou comandante]` — constrói um deck novo do zero.
- `/improve-deck [caminho da decklist]` — audita e otimiza um deck existente.

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

## Estado persistente por deck

Cada deck vive em `decks/<slug>/` (slug = nome do comandante em kebab-case):

```
decks/<slug>/
├── 00-briefing.md      # modo (build/improve), comandante, identidade de cor, tema,
│                       # palavras-chave, orçamento, power level, coleção, decklist atual
├── 01-commander.md     # opções de comandante e escolha (se aplicável)
├── 02-theme.md         # análise linha a linha + pool temático
├── 03-draw.md          # candidatas de draw
├── 04-ramp.md          # candidatas de ramp
├── 05-interaction.md   # candidatas de interação e wipes
├── 06-manabase.md      # terrenos + plano de cortes
├── 07-wincons.md       # condições de vitória + protocolo de goldfishing
├── deck.md             # lista consolidada viva (fonte de verdade)
├── decisions.md        # registro cronológico de cortes e entradas (regra 5)
└── report.md           # relatório final
```

Subagentes **leem** `00-briefing.md` e `deck.md`, **escrevem** apenas o arquivo da sua fase e devolvem um resumo curto. Só o orquestrador atualiza `deck.md`.

### Formato de `deck.md`

Tabela por seção (Comandante, Criaturas, Artefatos, Encantamentos, Instantâneos, Feitiços, Planeswalkers, Terrenos):

```
| Carta | CMC | Tipo | Cores | Categorias | Sinergias |
```

`Categorias` usa os rótulos: `tema`, `draw`, `ramp`, `remoção`, `proteção`, `counter`, `wipe`, `wincon`, `terreno`.

## Regras não negociáveis

1. **Toda busca no Scryfall** segue `references/scryfall-search-guide.md` — leia antes de buscar. Sempre inclua `legal:commander` e `id<=<identidade do comandante>`; com orçamento definido, inclua `usd<X`.

2. **Orçamento é medido pelo menor valor da LigaMagic — nunca pelo preço da Scryfall.** O `usd<X` serve só para peneirar candidatas na busca; ele **não** decide se a carta ou o deck cabem no teto. Antes de afirmar que algo cabe no orçamento, confira em `https://www.ligamagic.com.br/?view=cards/card&card=<Nome+Em+Ingles>` e use o **primeiro** dos três números do bloco "Preço Médio de Venda no Marketplace" (menor / médio / maior). A página carrega preço via JS — WebFetch não funciona, use as ferramentas de browser. Erros do proxy medidos em 2026-08-12 chegaram a **6,5× para mais** (Restoration Magic: proxy R$1,65 × real R$10,75) e **14× para menos** (Chief of the Foundry: proxy R$1,16 × real R$0,08), nos dois sentidos — não há fator de correção possível. Ao apresentar totais, rotule sempre a origem **e a idade**: `estimativa (Scryfall)` ou `LigaMagic (menor), cotações de <data>`. Consulte o que já foi capturado com `mtgdb prices <nomes...>` antes de recapturar, e registre toda cotação nova com `mtgdb prices -add "<carta>" <valor>` — preço é observação datada, nunca um valor que se sobrescreve. `mtgdb prices -volatile` mostra quais cartas de fato oscilam, que são as únicas que precisam ser reconferidas antes de um torneio.

3. **Sinergia sobreposta**: só recomende carta que tenha **2+ pontos de sinergia** com o comandante e/ou com outras cartas já escolhidas. Justifique cada recomendação. Evite cartas que só funcionam isoladas.

4. **Ficha completa antes de qualquer veredito** — leia `references/card-evaluation-checklist.md` antes de recomendar ou cortar qualquer carta. Um deck de Commander é multifacetado: a mesma carta costuma exercer 3–5 funções ao mesmo tempo (efeito escrito, corpo que pode ser tapado para crew/station/convoke/improvise, tipo que alimenta contagens, receptor de contadores e anthems, redutor de custo, fixação de cor). **Julgar por um aspecto só é a causa raiz do vai-e-vem de cartas entre rodadas.** Nenhum corte é proposto sem enumerar por escrito todas as funções da carta e nomear quem cobre cada uma. Frases como "criatura fraca" ou "corpo pequeno" não são justificativa de corte — são sinal de que a ficha não foi feita. O mesmo critério usado para defender a entrada vale para a carta que sai.

5. **Registro de decisão** — todo corte e toda entrada vão para `decks/<slug>/decisions.md`. Antes de propor carta que já esteve no deck, consulte o registro: a proposta precisa dizer quem cortou, por quê, e **o que mudou desde então**. Se o motivo original continua válido, a carta não volta.

6. **Puxe o texto oracle na hora, sempre.** Nunca julgue carta de memória — nem as do próprio deck. Use **`bin/mtgdb`** (banco local com o bulk data do Scryfall — ver `references/mtgdb.md`): `mtgdb oracle "<nome>" ...` para cartas e `mtgdb deck <slug>` para o deck inteiro. Caia para o MCP do Scryfall só quando a carta for mais nova que o último dump. Se o banco não existir, rode `make db` (~15 s).

7. **Coleção pessoal primeiro — prioridade de análise, não obrigação de uso.** Toda carta que o usuário já possui é avaliada **antes** de qualquer compra: consulte com `mtgdb collection <nomes...>` (fonte: `data/collection.tsv`). Mas possuir a carta não a torna elegível: se a peça da coleção não servir ao deck, **comprar é a decisão correta**. O que a regra exige é que a coleção seja *considerada primeiro* e que a dispensa seja *justificada por escrito* — o especialista que propõe uma compra precisa nomear a carta equivalente da coleção e dizer por que ela não cobre a função. Nunca force uma carta ruim no deck só porque ela já está na caixa, e nunca proponha compra sem ter olhado a caixa.

8. **Nomes de cartas sempre em inglês** (nome oficial do Scryfall). Textos, análises e conversa em **português (Brasil)**.
   No `report.md`, todo nome de carta é **link clicável para a ficha da LigaMagic**: `[**Nome**](https://www.ligamagic.com.br/?view=cards/card&card=<nome+percent-encoded>)`, com espaço virando `+` e o resto em percent-encoding UTF-8 (`Krenko%2C+Mob+Boss`). Única exceção: o bloco de exportação padrão MTG Online, que é texto puro. Formato completo em `references/deck-report-template.md`.

9. **O usuário decide**: cada fase termina com o usuário revisando e selecionando cartas. Nenhuma carta entra em `deck.md` sem aprovação.

10. O deck final tem exatamente **100 cartas** (comandante + 99), todas dentro da identidade de cor e legais no formato.

11. Regra da diversão: evite pacotes que impedem os oponentes de jogar (MLD, stax pesado, lock infinito) a menos que o usuário peça explicitamente.

## Referências

- `references/card-evaluation-checklist.md` — **ficha de funções F1–F7, protocolo de corte e registro de decisão** (regras 4 e 5). Leitura obrigatória antes de recomendar ou cortar carta.
- `references/mtgdb.md` — **banco local de cartas, tags e rulings** (`bin/mtgdb`). É por onde passam oracle, busca, tags, rulings, preços e coleção.
- `references/scryfall-search-guide.md` — sintaxe, tags confirmadas, receitas de busca, controle de volume.
- `references/deck-report-template.md` — template do relatório final, incluindo o **formato dos links de carta para a LigaMagic** (regra 8).
