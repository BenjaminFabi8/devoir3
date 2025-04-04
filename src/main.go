package main

import (
	"bufio"
	"devoir3/src/game"
	"devoir3/src/utils"
	"fmt"
	"os"
	"strings"
)

var (
	objectiveReached bool            = false
	logs             []game.LogEntry = make([]game.LogEntry, 0)
)

func main() {

	consoleReader := bufio.NewReader(os.Stdin)
	handleCommand("help", nil)
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

func StartProgram() {
	objectiveReached = false
	logs = make([]game.LogEntry, 0)

	input := utils.ReadInputGridFromFile("medium.txt")
	grid := game.NewGameGrid(input)
	grid.SetObjectives()
	if len(grid.Objectives) == 0 {
		fmt.Println("Need at least 1 objectif")
		return
	}
	agentsPositions := grid.GetAgentsPositions()
	agents := make([]game.Agent, len(agentsPositions))

	if len(agents) >= 3 {
		for i, agentPos := range agentsPositions {
			if i == 0 {
				agents[i] = game.NewAStartAgent(i, agentPos, grid)
			} else if i == 1 {
				agents[i] = game.NewAStartWaitingAgent(i, agentPos, grid)
			} else {
				agents[i] = game.NewRandomAgent(i, agentPos, grid)
			}
		}
	} else {
		fmt.Println("Need at least 3 agents")
		return
	}

	fmt.Printf("Objectives: %v\n", grid.Objectives)
	game.StartAgents(agents)
	objectiveReached = true

	for _, agent := range agents {
		fmt.Printf("Agent %d Position: (%d, %d)\n", agent.GetId(), agent.GetPosition().X, agent.GetPosition().Y)
	}

	grid.PrintGrid()

	for _, agent := range agents {
		fmt.Printf("Agent %d Log Entries: %d\n", agent.GetId(), len(agent.GetLogEntries()))
		logs = append(logs, agent.GetLogEntries()...)
	}
}
