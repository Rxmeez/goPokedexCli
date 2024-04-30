package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func (p *PokeLocationArea) printLocationNames() {
	for _, res := range p.Results {
		fmt.Println(res.Name)
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

func GetPokeLocationArea(c *Config, direction string) error {
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

	pokeLocationArea := PokeLocationArea{}

	err = json.Unmarshal(body, &pokeLocationArea)
	if err != nil {
		return err
	}

	if pokeLocationArea.Next != nil {
		c.Next = *pokeLocationArea.Next
	} else {
		c.Next = ""
	}
	if pokeLocationArea.Previous != nil {
		c.Previous = *pokeLocationArea.Previous
	} else {
		c.Previous = ""
	}

	pokeLocationArea.printLocationNames()
	return nil

}
