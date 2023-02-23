package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := make(map[string]int)

	if len(os.Args) < 2 {
		fmt.Println("use: cmd <file_path>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error opening file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines[line]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file: ", err)
	}

	for line, count := range lines {
		if count > 1 {
			fmt.Printf("%s = %d\n", line, count)
		}
	}
}
