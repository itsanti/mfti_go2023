package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " ")) one-liner

	res := make([]any, len(os.Args)-1)

	for i, v := range os.Args[1:] {
		res[i] = strings.TrimSpace(v)
	}

	fmt.Println(res...)
}
