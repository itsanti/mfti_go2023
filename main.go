package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const API_START_URL = "https://pokeapi.co/api/v2/pokemon/?limit=100&offset=0"

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

func GetPokemonsPages(wg *sync.WaitGroup, outPokemons chan<- pokemonsResult) {
	defer wg.Done()

	var nextURL = make(chan string, 1)
	nextURL <- API_START_URL

	for url := range nextURL {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			var data pokemonsResult
			if url == "" {
				close(nextURL)
				close(outPokemons)
			} else {
				jsonData, err := makeRequest(url)
				if err != nil {
					panic(err.Error())
				}

				err = json.Unmarshal(jsonData, &data)
				if err != nil {
					panic(err.Error())
				}

				outPokemons <- data
				nextURL <- data.Next
			}
		}(url)
	}

}

func GetPokemonPage(wg *sync.WaitGroup, inPokemons <-chan pokemonsResult, outPokemon chan<- pokemonResult) {
	for data := range inPokemons {
		for _, item := range data.Results {
			wg.Add(1)
			go func(url string) {
				var item pokemonResult
				var pokemon Pokemon
				defer wg.Done()
				jsonData, err := makeRequest(url)
				if err != nil {
					item.Err = err
				}
				err = json.Unmarshal(jsonData, &pokemon)
				if err != nil {
					item.Err = err
				} else {
					item.Pokemon = pokemon
				}
				outPokemon <- item
			}(item.Url)
		}
	}
}

func GetPokemons() []Pokemon {
	var pokemons []Pokemon
	var wg sync.WaitGroup

	var chanPokemonsResult = make(chan pokemonsResult)
	var chanPokemons = make(chan pokemonResult)

	wg.Add(1)
	go GetPokemonsPages(&wg, chanPokemonsResult)
	go GetPokemonPage(&wg, chanPokemonsResult, chanPokemons)

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

func main() {
	start := time.Now()
	pokemons := GetPokemons()
	s := float64(time.Since(start).Microseconds())

	fmt.Println(len(pokemons))
	fmt.Printf("GetPokemons time: %gs.\n", s/1000000)
}
