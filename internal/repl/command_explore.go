package repl

import (
	"github.com/rxmeez/goPokedexCli/internal/api"
)

func commandExplore(location string) error {
	err := api.ExplorePokemonLocationArea(location)
	return err
}
