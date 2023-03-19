package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const API_BASE_ENDPOINT = "https://pokeapi.co/api/v2"

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Order          int    `json:"order"`
	IsDefault      bool   `json:"is_default"`
}

type pokemonResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string
		Url  string
	} `json:"results"`
}

func makeRequest(path string) ([]byte, error) {
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func GetPokemons() []Pokemon {
	var data pokemonResult
	var pokemon Pokemon
	var pokemons []Pokemon
	var chanPokemons = make(chan Pokemon)
	var wg sync.WaitGroup

	jsonData, err := makeRequest(API_BASE_ENDPOINT + "/pokemon")
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(jsonData, &data)

	for _, item := range data.Results {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			jsonData, err = makeRequest(url)
			json.Unmarshal(jsonData, &pokemon)
			chanPokemons <- pokemon
		}(item.Url)
	}

	go func() {
		wg.Wait()
		close(chanPokemons)
	}()

	for pokemon := range chanPokemons {
		pokemons = append(pokemons, pokemon)
	}

	for _, pok := range pokemons {
		fmt.Println(pok)
	}

	return pokemons
}

func main() {
	GetPokemons()
}
