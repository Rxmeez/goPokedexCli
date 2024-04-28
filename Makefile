start: ./bin/goPokedexCli 

build:
	@go build -o ./bin ./cmd/goPokedexCli

.PHONY: build start 