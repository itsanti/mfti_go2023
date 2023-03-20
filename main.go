package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i, v := range os.Args[1:] {
		os.Args[i] = strings.TrimSpace(v)
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
}
