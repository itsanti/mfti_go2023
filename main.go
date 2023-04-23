package main

import (
	"fmt"
	"time"
)

type APIclient interface {
	GetPokemons() []Pokemon
}

func GetAPIclient(mock bool) APIclient {
	var a APIclient

	if mock {
		a = MockingHTTPclient{Count: 3}
	} else {
		const timeout = 2 * time.Second
		const retries = 3
		const debug = false
		a = RetryableHTTPclient{Timeout: timeout, Retries: retries, Debug: debug}
	}
	return a
}

func main() {
	var mocking bool
	mocking = true
	pokemons := GetAPIclient(mocking).GetPokemons()
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
}
