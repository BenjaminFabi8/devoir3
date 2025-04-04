package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputGridFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	gameGrid := make([]string, 0)

	for scanner.Scan() {
		gameGrid = append(gameGrid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)

	}

	return gameGrid
}
