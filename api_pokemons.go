package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

const API_BASE_ENDPOINT = "https://pokeapi.co/api/v2/pokemon"

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Order          int    `json:"order"`
	IsDefault      bool   `json:"is_default"`
}

type pokemonsResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string
		Url  string
	} `json:"results"`
}

type pokemonResult struct {
	Pokemon Pokemon
	Err     error
}

func (rc RetryableHTTPclient) makeRequest(path string) ([]byte, error) {
	res := rc.Get(path)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (rc RetryableHTTPclient) GetPokemons() []Pokemon {
	var data pokemonsResult
	var pokemons []Pokemon
	var chanPokemons = make(chan pokemonResult)
	var wg sync.WaitGroup

	jsonData, err := rc.makeRequest(API_BASE_ENDPOINT)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err.Error())
	}

	for _, item := range data.Results {
		wg.Add(1)
		go func(url string) {
			var item pokemonResult
			var pokemon Pokemon
			defer wg.Done()
			jsonData, err := rc.makeRequest(url)
			if err != nil {
				item.Err = err
			}
			err = json.Unmarshal(jsonData, &pokemon)
			if err != nil {
				item.Err = err
			} else {
				item.Pokemon = pokemon
			}
			chanPokemons <- item
		}(item.Url)
	}

	go func() {
		wg.Wait()
		close(chanPokemons)
	}()

	for item := range chanPokemons {
		if item.Err != nil {
			continue
		}
		pokemons = append(pokemons, item.Pokemon)
	}

	return pokemons
}
