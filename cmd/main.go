package main

import (
	"encoding/json"
	"log"
	"net/http"
	"pokemon-rest-api/listing"
	"sort"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var pokemons []listing.Pokemon

func UploadPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var pokemon listing.Pokemon
	err := decoder.Decode(&pokemon)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pokemons = append(pokemons, pokemon)

	w.WriteHeader(http.StatusCreated)
}

func DownloadPokemons(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id--

	if id < 0 || id >= len(pokemons) {
		http.Error(w, "Bad id param", 400)
		return
	}

	pokemonsSort(pokemons)
	json.NewEncoder(w).Encode(pokemons[id])
}

func pokemonsSort(pokemons []listing.Pokemon) {
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})
}

func main() {
	router := httprouter.New()
	router.GET("/pokemons", DownloadPokemons)
	router.GET("/pokemons/:id", GetPokemonById)
	router.POST("/pokemons", UploadPokemon)

	log.Fatal(http.ListenAndServe(":8080", router))
}
