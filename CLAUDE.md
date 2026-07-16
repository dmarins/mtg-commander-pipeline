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
2. **Sinergia sobreposta**: só recomende carta que tenha **2+ pontos de sinergia** com o comandante e/ou com outras cartas já escolhidas. Justifique cada recomendação. Evite cartas que só funcionam isoladas.
3. **Coleção pessoal primeiro**: se o briefing indicar um arquivo de coleção, priorize cartas que o usuário já possui antes de sugerir compras.
4. **Nomes de cartas sempre em inglês** (nome oficial do Scryfall). Textos, análises e conversa em **português (Brasil)**.
5. **O usuário decide**: cada fase termina com o usuário revisando e selecionando cartas. Nenhuma carta entra em `deck.md` sem aprovação.
6. O deck final tem exatamente **100 cartas** (comandante + 99), todas dentro da identidade de cor e legais no formato.
7. Regra da diversão: evite pacotes que impedem os oponentes de jogar (MLD, stax pesado, lock infinito) a menos que o usuário peça explicitamente.

## Referências

- `references/scryfall-search-guide.md` — sintaxe, tags confirmadas, receitas de busca, controle de volume.
- `references/deck-report-template.md` — template do relatório final.
