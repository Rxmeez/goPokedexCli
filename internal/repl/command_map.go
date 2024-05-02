package repl

import (
	"github.com/rxmeez/goPokedexCli/internal/api"
)

var pokeConfig = &api.Config{
	Url:      api.BaseUrl + "location-area/",
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
