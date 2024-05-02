package repl

import (
	"github.com/rxmeez/goPokedexCli/internal/api"
)

func commandCatch(pokemon string) error {
	err := api.CatchPokemon(pokemon)
	return err
}
