package repl

import (
	"github.com/rxmeez/goPokedexCli/internal/api"
)

func commandPokedex() error {
	api.PokeStore.List()
	return nil
}
