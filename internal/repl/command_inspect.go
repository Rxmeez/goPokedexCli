package repl

import (
	"github.com/rxmeez/goPokedexCli/internal/api"
)

func commandInspect(pokemon string) error {
	err := api.InspectPokemon(pokemon)
	return err
}
