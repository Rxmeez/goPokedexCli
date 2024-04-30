package repl

import (
	"github.com/rxmeez/goPokedexCli/internal/api"
)

var pokeConfig = &api.Config{
	Url:      "https://pokeapi.co/api/v2/location-area/",
	Next:     "",
	Previous: "",
}

func commandMapf() error {
	err := api.GetPokeLocationArea(pokeConfig, "next")
	return err
}

func commandMapb() error {
	err := api.GetPokeLocationArea(pokeConfig, "previous")
	return err
}
