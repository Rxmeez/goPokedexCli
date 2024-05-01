package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rxmeez/goPokedexCli/internal/pokecache"
)

type Config struct {
	Url      string
	Next     string
	Previous string
}

type PokeLocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (p *PokeLocationArea) printLocationNames(body []byte) error {
	err := json.Unmarshal(body, p)
	if err != nil {
		return err
	}
	for _, res := range p.Results {
		fmt.Println(res.Name)
	}
	return nil
}

func (p *PokeLocationArea) updateUrlState(c *Config) {
	if p.Next != nil {
		c.Next = *p.Next
	} else {
		c.Next = ""
	}
	if p.Previous != nil {
		c.Previous = *p.Previous
	} else {
		c.Previous = ""
	}

}

func selectURL(c *Config, direction string) string {
	switch direction {
	case "next":
		if c.Next != "" {
			return c.Next
		}
	case "previous":
		if c.Previous != "" {
			return c.Previous
		}
	}
	return c.Url
}

var pokeCache = pokecache.NewCache(5 * time.Second)

func GetPokeLocationArea(c *Config, direction string) error {
	url := selectURL(c, direction)
	pokeLocationArea := PokeLocationArea{}

	if v, ok := pokeCache.Get(url); ok {

		err := pokeLocationArea.printLocationNames(v)
		if err != nil {
			return err
		}
		pokeLocationArea.updateUrlState(c)

	}

	res, err := http.Get(selectURL(c, direction))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	err = pokeLocationArea.printLocationNames(body)
	if err != nil {
		return err
	}
	pokeLocationArea.updateUrlState(c)
	return nil

}
