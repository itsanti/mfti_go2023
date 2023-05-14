package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	i "pokemon-rest-api/init"
	"pokemon-rest-api/listing"
	"sort"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var pokemonsCache []listing.Pokemon

func init() {
	i.ConnectDB()
	if result := i.DB.Order("id").Find(&pokemonsCache); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func UploadPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pokemon listing.Pokemon

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pokemon)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	idx := cacheSearch(pokemonsCache, pokemon.ID)

	if idx < 0 {
		if result := i.DB.Create(&pokemon); result.Error != nil {
			fmt.Println(result.Error)
			return
		} else {
			pokemonsCache = append(pokemonsCache, pokemon)
			pokemonsSort(pokemonsCache)
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func DownloadPokemons(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonsCache)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, "Param id must be integer", 400)
		return
	}

	if id <= 0 {
		http.Error(w, "Param id must be greater than 0", 400)
		return
	}

	idx := cacheSearch(pokemonsCache, id)

	if idx >= 0 {
		json.NewEncoder(w).Encode(pokemonsCache[idx])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func cacheSearch(pokemonsCache []listing.Pokemon, id int) int {
	for i, p := range pokemonsCache {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func pokemonsSort(pokemons []listing.Pokemon) {
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})
}

func main() {
	i.DB.AutoMigrate(&listing.Pokemon{})
	fmt.Println("? Migration complete")

	router := httprouter.New()
	router.GET("/pokemons", DownloadPokemons)
	router.GET("/pokemons/:id", GetPokemonById)
	router.POST("/pokemons", UploadPokemon)

	log.Fatal(http.ListenAndServe(":8080", router))
}
