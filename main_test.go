package main

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestGetPokemons(t *testing.T) {
	var mocking bool
	for _, i := range GetAPIclient(mocking).GetPokemons() {
		var x interface{} = i
		p, ok := x.(Pokemon)
		if !ok {
			t.Error("Неверный тип покемона", p)
		}
	}
}

func TestGetPokemonsMock(t *testing.T) {
	mocking := true
	start := time.Now()
	pokemons := GetAPIclient(mocking).GetPokemons()
	elapsed := time.Since(start)

	if elapsed.Seconds() >= 3 {
		t.Errorf("%s too slow: %f sec.", runtime.FuncForPC(reflect.ValueOf(TestGetPokemonsMock).Pointer()).Name(), elapsed.Seconds())
	}

	for _, i := range pokemons {
		var x interface{} = i
		p, ok := x.(Pokemon)
		if !ok {
			t.Error("Неверный тип покемона", p)
		}
	}
}
