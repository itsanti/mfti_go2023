package main

import (
	"fmt"
	"time"
)

type APIclient interface {
	GetPokemons() []Pokemon
}

func GetAPIclient(mock bool) APIclient {
	var client APIclient

	if mock {
		client = MockingHTTPclient{Count: 3}
	} else {
		client = RetryableHTTPclient{
			Timeout: 2 * time.Second,
			Retries: 3,
			Debug:   false,
		}
	}
	return client
}

func main() {
	mocking := true
	pokemons := GetAPIclient(mocking).GetPokemons()
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
}
