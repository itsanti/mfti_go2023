package main

import (
	"math/rand"
	"time"
)

type MockingHTTPclient struct {
	Count int
}

const RND_MAX_VALUE = 100

var POCKEMON_NAMES = [...]string{
	"caterpie", "beedrill", "metapod", "butterfree", "charizard",
	"charmeleon", "bulbasaur", "blastoise", "weedle", "kakuna",
}

func (mc MockingHTTPclient) GetPokemons() []Pokemon {
	pockemons := make([]Pokemon, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < mc.Count; i++ {
		pockemon := Pokemon{
			ID:             rand.Intn(RND_MAX_VALUE),
			Name:           POCKEMON_NAMES[rand.Intn(len(POCKEMON_NAMES))],
			BaseExperience: rand.Intn(RND_MAX_VALUE),
			Height:         rand.Intn(RND_MAX_VALUE),
			Weight:         rand.Intn(RND_MAX_VALUE),
			Order:          rand.Intn(RND_MAX_VALUE),
			IsDefault:      rand.Intn(2) == 1,
		}
		pockemons = append(pockemons, pockemon)
	}
	return pockemons
}
