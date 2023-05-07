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
const PAGE_COUNT = 5

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

func GetPokemonsPages(wg *sync.WaitGroup, chanPokemons chan pokemonResult) {
	defer wg.Done()

	var nextURL = make(chan string, 20)
	nextURL <- API_START_URL

	for i := 0; i < PAGE_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var data pokemonsResult

			url := <-nextURL

			if url == "" {
				return
			}

			jsonData, err := makeRequest(url)
			if err != nil {
				panic(err.Error())
			}

			err = json.Unmarshal(jsonData, &data)
			if err != nil {
				panic(err.Error())
			}

			if data.Next != "" {
				nextURL <- data.Next
			} else {
				close(nextURL)
			}

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
					chanPokemons <- item
				}(item.Url)
			}
		}()
	}
}

func GetPokemons() []Pokemon {
	var pokemons []Pokemon
	var chanPokemons = make(chan pokemonResult)
	var wg sync.WaitGroup

	wg.Add(1)
	go GetPokemonsPages(&wg, chanPokemons)

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
