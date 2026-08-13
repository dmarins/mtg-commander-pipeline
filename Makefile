BIN := bin/mtgdb
DB  := data/scryfall.db

.PHONY: build db refresh status test fmt vet clean help

## build: compila o binário mtgdb
build:
	@go build -o $(BIN) ./cmd/mtgdb
	@echo "→ $(BIN)"

## db: constrói o banco a partir do bulk data do Scryfall (só se ainda não existir)
db: build
	@test -f $(DB) && echo "$(DB) já existe — use 'make refresh' para atualizar" || ./$(BIN) build

## refresh: rebaixa o bulk data e reconstrói o banco do zero
refresh: build
	@rm -f $(DB) $(DB)-wal $(DB)-shm
	@./$(BIN) build

## status: mostra o estado do banco e a data dos dumps
status: build
	@./$(BIN) status

## test: roda os testes
test:
	@go test ./...

## fmt: formata o código
fmt:
	@go fmt ./...

## vet: análise estática
vet:
	@go vet ./...

## clean: remove binário e banco (os TSV de data/ são preservados)
clean:
	@rm -rf bin $(DB) $(DB)-wal $(DB)-shm
	@echo "removidos bin/ e $(DB) — data/*.tsv preservados"

## help: lista os alvos
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'
