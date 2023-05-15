package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"pokemon-rest-api/listing"
	"pokemon-rest-api/storage"
	"pokemon-rest-api/utils"
)

var pokemonsCache []listing.Pokemon

func uploadPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pokemon listing.Pokemon

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pokemon)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	idx := utils.CacheSearch(pokemonsCache, pokemon.ID)

	if idx < 0 {
		if result := storage.DB.Create(&pokemon); result.Error != nil {
			fmt.Println(result.Error)
			return
		} else {
			pokemonsCache = append(pokemonsCache, pokemon)
			utils.PokemonsSort(pokemonsCache)
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func downloadPokemons(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonsCache)
}

func getPokemonById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	idx := utils.CacheSearch(pokemonsCache, id)

	if idx >= 0 {
		json.NewEncoder(w).Encode(pokemonsCache[idx])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func fillCache() {
	if result := storage.DB.Order("id").Find(&pokemonsCache); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func initServer() *httprouter.Router {
	storage.InitDB()
	fillCache()

	router := httprouter.New()
	router.GET("/pokemons", downloadPokemons)
	router.GET("/pokemons/:id", getPokemonById)
	router.POST("/pokemons", uploadPokemon)
	return router
}

func RunServer(host, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), initServer()))
}
