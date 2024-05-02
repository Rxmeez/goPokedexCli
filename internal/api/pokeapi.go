package api

import (
	"time"

	"github.com/rxmeez/goPokedexCli/internal/pokecache"
)

const BaseUrl string = "https://pokeapi.co/api/v2/"

var PokeCache = pokecache.NewCache(5 * time.Second)
