package main

import (
	"fmt"
	"devoir3/src/game"
)

func handleCommand(command string, args []string) {
	if command == "" {
		return
	}

	switch command {
	case "start":
		if len(args) < 1 {
			fmt.Println("Usage: start <input_file>")
			return
		}
		
	case "log":
		if objectiveReached {
			game.CreateLogEntriesFile(logs)
			fmt.Println("Objective reached!")
			return
		}

}
}
