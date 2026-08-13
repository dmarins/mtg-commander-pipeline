# Checklist de avaliação de carta

Leitura obrigatória antes de **recomendar** ou **cortar** qualquer carta, em `build` ou `improve`.

> **A falha que este documento existe para impedir.** Um deck de Commander é multifacetado: a mesma carta costuma exercer 3, 4 ou 5 funções ao mesmo tempo. Quando cada agente julga a carta só pela lente da sua especialidade, ele enxerga uma dessas funções e conclui que a carta é fraca. O resultado é corte indevido — e, uma rodada depois, a mesma carta volta por outro motivo. **Vai-e-vem de cartas é sintoma de análise parcial, não de mudança de contexto.**
>
> Caso real (deck Inspirit, v3): *Etherium Sculptor* foi proposta para corte como "criatura 1/2, o pior corpo do deck". A ficha completa mostrava outra coisa — ela **reduz o custo de todo artifact spell** (o deck tem 33 artefatos), **pode ser tapada para Station e crew**, **conta como artefato** para Master of Etherium/metalcraft/affinity, **recebe os dois anthems** (vira 3/4) e **entra no turno 2**. Cinco funções; a análise tinha olhado uma.

---

## 1. Ficha de funções

Antes de qualquer veredito, preencha as sete linhas abaixo para a carta. Só depois julgue.

| # | Eixo | Pergunta |
|---|---|---|
| **F1** | **Texto** | O que **cada linha** do oracle faz? Habilidades estáticas, disparadas e ativadas, todas — inclusive as que parecem irrelevantes. |
| **F2** | **Corpo** | É criatura (ou vira criatura)? Poder/resistência? Pode **ser tapada** para um custo — crew, station, convoke, improvise, habilidades de tap, exert? Pode ser sacrificada como recurso? Bloqueia? |
| **F3** | **Tipo como recurso** | O tipo dela alimenta contagens e custos do deck? (artefato → affinity/metalcraft/improvise/"number of artifacts you control"; encantamento → constelação; criatura → convoke, lordes.) |
| **F4** | **Receptor** | Ela **recebe** algo? Contadores, anthems e lordes, buffs estáticos, gatilhos do comandante, proteção estática (hexproof/indestructible concedidos por outra carta), auras/equipamentos. |
| **F5** | **Facilitador** | Ela **dá** algo a outras cartas? Redução de custo, fixação de cor, haste/vigilância/evasão, proteção, aumento de poder. |
| **F6** | **Curva** | Em que turno ela entra, e o que ela destrava a partir dali? Uma carta barata que acelera o resto vale mais do que o efeito isolado sugere. |
| **F7** | **Atrito** | Anti-sinergias e competição interna: ela disputa **tap**, **mana**, **alvo único de um gatilho** ou **slot de ataque** com outra peça? Alguma estática do próprio deck **anula** parte dela? |

**Regra de ouro do F1:** puxe o texto oracle na hora. **Nunca julgue por memória.** Numa única auditoria do deck Inspirit, cinco afirmações escritas de memória estavam erradas — e três delas invertiam a decisão de corte.

```bash
bin/mtgdb oracle "Etherium Sculptor" "Cloud Key"   # ficha F2–F7 + oracle completo
bin/mtgdb deck <slug>                              # o deck inteiro de uma vez, com agregados
bin/mtgdb rulings "<nome>"                         # quando a dúvida for de regra
```

O `mtgdb` imprime os eixos F2–F7 já preenchidos, mas eles são **sinais extraídos por casamento de padrão, não vereditos** — e a ferramenta diz isso na própria saída. O texto oracle vem junto justamente para ser lido. Ela economiza o trabalho mecânico de levantar os fatos; a leitura e o julgamento continuam seus.

---

## 2. Protocolo de corte

Uma carta só pode ser proposta para corte depois de passar por isto, **por escrito**:

1. **Enumere** todas as funções que a ficha revelou (F1–F6).
2. Para **cada função**, nomeie quem a cobre depois do corte — outra carta do deck, ou a entrada proposta.
3. Se alguma função ficar **descoberta**, declare o custo explicitamente. O corte ainda pode valer a pena; o que não vale é fingir que a função não existia.
4. Só é corte "limpo" quando **todas** as funções estão cobertas ou declaradas como dispensáveis.

Formato mínimo na tabela de swap:

```
Sai: <carta> — funções: [f1, f2, f3] → f1 coberta por <X>, f2 coberta por <Y>, f3 fica descoberta (custo aceito: …)
```

**Não conta como justificativa de corte:** "criatura fraca", "corpo pequeno", "não faz nada pelo tema", "carta genérica" — nenhuma dessas frases sobrevive a uma ficha completa. Se for o único argumento disponível, a ficha não foi feita.

---

## 3. Simetria de critério

O critério usado para **defender uma entrada** tem de ser aplicado às cartas que **saem**, e vice-versa.

Se você argumenta que a carta X entra porque "com os dois anthems ela vira 5/6", então a carta Y que sai também precisa ser avaliada com os anthems em campo. Aplicar o bônus só de um lado da troca produz uma comparação falsa — foi exatamente assim que a v3 do Inspirit quase cortou um redutor de custo que era, na verdade, um estacionador de poder 3.

Antes de fechar qualquer swap, releia os dois lados e confirme que o mesmo conjunto de condições foi assumido para ambos.

---

## 4. Registro de decisão — evitar o vai-e-vem

Todo corte e toda entrada vão para `decks/<slug>/decisions.md`, em ordem cronológica:

```markdown
| Data | Carta | Ação | Fase/agente | Motivo | Funções descobertas pelo corte |
|---|---|---|---|---|---|
| 2026-07-24 | Emissary Escort | corte | v1 · manabase | espaço para 36→38 terrenos | nenhuma avaliada |
```

**Antes de propor uma carta que já esteve no deck, consulte esse registro.** Se ela foi cortada, a proposta precisa responder por escrito:

- quem cortou, em que fase e por qual motivo;
- **o que mudou desde então** — nova carta no deck, correção de leitura de regra, métrica que não existia antes, mudança de objetivo do usuário.

Se o motivo original do corte continua válido, **a carta não volta**. "Reavaliei e agora gostei" não é mudança de contexto.

---

## 5. Verificação antes de entregar

Antes de apresentar qualquer pacote de trocas ao usuário:

- [ ] Todas as cartas envolvidas (entram **e** saem) têm ficha F1–F7 preenchida.
- [ ] Todo corte passou pelo protocolo da seção 2, com as funções mapeadas para quem as cobre.
- [ ] Os dois lados de cada troca foram avaliados sob as **mesmas** condições (seção 3).
- [ ] Nenhuma entrada é reposição de corte anterior sem a justificativa da seção 4.
- [ ] Todos os textos oracle foram puxados do Scryfall nesta sessão, não citados de memória.
- [ ] Nenhum corte atinge a peça que o **próprio usuário** disse querer usar — quando ele reclama que algo é lento, o pedido é para **fazer aquilo funcionar**, não para remover.
