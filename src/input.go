package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputGridFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

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

func outputLogsToFile(logs string, filepath string) {
	file, err := os.Create(filepath)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(logs)
	if err != nil {
		fmt.Println(err)
	}
}
