package utils

import (
	"sort"

	"pokemon-rest-api/listing"
)

func CacheSearch(pokemonsCache []listing.Pokemon, id int) int {
	for i, p := range pokemonsCache {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func PokemonsSort(pokemons []listing.Pokemon) {
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})
}
