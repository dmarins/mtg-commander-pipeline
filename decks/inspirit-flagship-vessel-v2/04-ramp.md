# Aceleração (Ramp) — v2, escopo estreito

> Esta fase não reavalia o pacote de ramp inteiro (já auditado em `02-theme.md`): trata **apenas** do swap apontado pela auditoria temática — redundância de fixação W/U + ausência de fixação de R nos mana rocks.

Ramp padrão já no deck: 10/10–11 (dentro da meta, sem alteração de contagem) · Explosivo: 1/2–3 (herdado da fase de interação, fora de escopo aqui)

## Diagnóstico

O deck tem **Talisman of Progress** e **Azorius Signet**, ambos com `color_identity: U, W` — nenhum dos dois toca vermelho. Nos terrenos básicos, Mountain (5) é a cor mais escassa (vs. 6 Plains / 9 Island), e a demanda de pips R é real e distribuída (Chainsaw, Rip Apart, Alibou, Kilo, Third Path Iconoclast, Enthusiastic Mechanaut, Jhoira, Saheeli, Warmaker Gunship, Chain Reaction). Manter os dois fixadores W/U não ajuda essa lacuna; um dos dois pode virar um fixador que inclua R sem perder nenhuma fixação líquida de W ou U (o outro permanece).

### Diferença real entre os dois candidatos ao corte

Apesar de ambos fixarem só W/U por {2} de custo de elenco, não são idênticos em eficiência:

- **Talisman of Progress**: `{T}: Add {C}` **ou** `{T}: Add {W} or {U}, this deals 1 damage to you` — ativação de 1 símbolo de mana (o próprio artefato), custo pago em **vida**, mana colorido disponível já no turno em que entra (mesmo sem mana extra disponível).
- **Azorius Signet**: `{1}, {T}: Add {W}{U}` — precisa de **{1} adicional** para ativar; entrega 2 mana coloridos de uma vez, mas exige ter mana sobrando primeiro (pior no turno em que entra, quando geralmente não há mana livre).

Isso favorece levemente o Talisman como o mais eficiente dos dois (custo em vida é normalmente mais barato que custo em mana num deck sem drenos de vida relevantes). **Decisão**: cortar **Azorius Signet**, manter Talisman of Progress.

## Busca (Scryfall)

`legal:commander id<=wur otag:mana-rock mv<=2 c:r usd<0.5 order:edhrec` — candidatos avaliados: Izzet Signet (U/R, ativação `{1},{T}`, US$0,46), Boros Signet (R/W, ativação `{1},{T}`, US$0,42), Talisman of Creativity (U/R, US$1,59), Talisman of Conviction (R/W, US$1,18).

Critério de escolha entre Izzet Signet e Boros Signet: o deck já tem **Island em excesso relativo** (9, a maior base de básicos + Talisman of Progress + Arcane Signet + Midnight Clock já cobrem U fartamente). W (6 Plains) e R (5 Mountain) são as duas cores mais carentes — um fixador **R/W** ataca as duas lacunas ao mesmo tempo, em vez de reforçar ainda mais a cor que já é a mais abundante.

## Swap proposto

| Sai | Entra | CMC | Cores fixadas | US$ | Faixa | Justificativa |
|---|---|---|---|---|---|---|
| Azorius Signet | Boros Signet | 2 | R/W | 0.42 | torneio | Corta a redundância de fixação W/U (Talisman of Progress permanece cobrindo W/U). Boros Signet fixa exatamente as duas cores mais escassas do deck nos básicos (6 Plains, 5 Mountain) sem tirar nem uma gota de fixação líquida de W (ainda coberto por Talisman + Arcane Signet) — e resolve a lacuna real apontada em `02-theme.md`: nenhum rock do deck ajudava especificamente R. Mesmo custo de elenco (2), mesma velocidade (T2), mesmo padrão de ativação `{1},{T}` do card que substitui — zero perda de curva ou de "feel" de jogo. Preço na faixa torneio (US$0,42 ≈ R$2,31), dentro do teto de R$200.|

## Efeito na contagem

- Ramp padrão: 10 → **10** (sem mudança de quantidade, apenas qualidade de fixação).
- Ramp explosivo: sem alteração (1/2–3, fora de escopo desta fase).
- Fixação de R: 0 rocks → **1 rock** (Boros Signet) ajudando especificamente a cor mais escassa do deck.
- Redundância W/U: 2 fixadores idênticos → **1** (Talisman of Progress).

## Reservas (não usadas, registradas para referência)

- **Izzet Signet** (U/R, US$0,46) — segunda opção equivalente em preço/velocidade, mas reforçaria U (já a cor mais bem servida) em vez de W; preterida.
- **Talisman of Conviction** (R/W, US$1,18) / **Talisman of Creativity** (U/R, US$1,59) — mesmo perfil de eficiência do Talisman de Progress (ativação por vida), mas fora da faixa torneio preferencial; candidatas de reserva na faixa `mesão` se o usuário preferir manter o padrão "Talisman" em vez de "Signet" por consistência de nomenclatura/arte, sem diferença funcional relevante.
