package main

import (
	"testing"
)

type pokemon map[string]any

func TestGetPokemons(t *testing.T) {
	for _, i := range GetPokemons() {
		var x interface{} = i
		p, ok := x.(Pokemon)
		if !ok {
			t.Error("Неверный тип покемона", p)
		}
	}
}
