package main

import (
	"devoir3/src/game"
	"fmt"
)

func handleCommand(command string, args []string) {
	if command == "" {
		return
	}

	switch command {
	case "start":
		if len(args) < 1 {
			fmt.Println("==== Starting agent program ====")
			StartProgram()
			return
		}

	case "help":
		if len(args) < 1 {
			fmt.Println("==================== LIST OF COMMANDS =======================")
			fmt.Println("start\t\t-\tStarts the program")
			fmt.Println("log <filename>\t-\tGenerates a log of the previous game")
			fmt.Println("help\t\t-\tDisplays this help message")
			fmt.Println("=============================================================")
			return
		}

	case "log":
		if !objectiveReached {
			fmt.Println("Start the program before logging.")
			return
		}
		if len(args) < 1 {
			fmt.Println("filename is required for logging")
			return
		}
		game.CreateLogEntriesFile(logs, args[0])
		fmt.Println("==== LOGGING COMPLETE ====")

	default:
		fmt.Println("Command not recognized. Type \"help\" to see all available commands")
	}
}
