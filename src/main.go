package main

import (
	"devoir3/src/game"
	"fmt"
	"devoir3/src/utils"
	"bufio"
	"os"
	"strings"
)

var ( objectiveReached bool = false
	  logs []game.LogEntry = make([]game.LogEntry, 0)
	)


func main() {

	consoleReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := consoleReader.ReadString('\n')

		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		command := strings.TrimSpace(parts[0])
		args := []string{}
		if len(parts) > 1 {
			args = strings.Fields(parts[1])
		}
		handleCommand(command, args)
	}
}

func StartProgram(){
	//initialize variables
	objectiveReached = false
	logs = make([]game.LogEntry, 0)

	input := utils.ReadInputGridFromFile("medium.txt")
	grid := game.NewGameGrid(input)
	grid.SetObjectives()
	fmt.Printf("Objectives: %v\n", grid.Objectives)
	agentsPositions := grid.GetAgentsPositions()
	agents := make([]game.Agent, len(agentsPositions))

	if len(agents) >= 3 {
		for i, agentPos := range agentsPositions {
			if i == 0 {
				//A*
				agents[i] = game.NewAStartAgent(i, agentPos, grid)
			} else if i == 1 {
				//A* wait
				agents[i] = game.NewAStartWaitingAgent(i, agentPos, grid)
			} else {
				agents[i] = game.NewRandomAgent(i, agentPos, grid.Objectives[0], grid)
			}
		}
	}

	game.StartAgents(agents)

	for _, agent := range agents {
		fmt.Printf("Agent %d Position: (%d, %d)\n", agent.GetId(), agent.GetPosition().X, agent.GetPosition().Y)
	}

	grid.PrintGrid()

	for _, agent := range agents {
		fmt.Printf("Agent %d Log Entries: %d\n", agent.GetId(), len(agent.GetLogEntries()))
		logs = append(logs, agent.GetLogEntries()...)
	}

	objectiveReached = true
	fmt.Println("Objective reached!")
}