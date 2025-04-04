package main

import (
	"devoir3/src/game"
	"fmt"
	"time"
    "path/filepath"
)

func main() {

	input := readInputGridFromFile("medium.txt")
	grid := game.NewGameGrid(input)
	grid.SetObjectives()
	fmt.Printf("Objectives: %v\n", grid.Objectives)
	agentsPositions := grid.GetAgentsPositions()
	agents := make([]game.Agent, len(agentsPositions))

	for i, agentPos := range agentsPositions {
		agents[i] = game.NewRandomAgent(i, agentPos, grid.Objectives[0], grid)
	}

	// agentAStarPos := game.Position{X: 1, Y: 1}

	game.StartAgents(agents)

	for _, agent := range agents {
		fmt.Printf("Agent %d Position: (%d, %d)\n", agent.GetId(), agent.GetPosition().X, agent.GetPosition().Y)
	}

	grid.PrintGrid()
	logs := make([]game.LogEntry, 0)

	for _, agent := range agents {
		fmt.Printf("Agent %d Log Entries: %d\n", agent.GetId(), len(agent.GetLogEntries()))
		logs = append(logs, agent.GetLogEntries()...)
		for _, entry := range agent.GetLogEntries() {
			fmt.Printf("Position: (%d, %d), Timestamp: %s\n", entry.Position.X, entry.Position.Y, entry.Timestamp.Format(time.RFC3339))
		}
	}

	/*===MERGE LOGS===*/
	mergedLogs := game.GetMergedLogEntriesString(logs)
	//fmt.Println("merged logs : \n" + mergedLogs)
	outputFile := filepath.Join("logs", "log_"+time.Now().Format("2006-01-02_15-04-05")+".txt")
	outputLogsToFile(mergedLogs, outputFile)
	fmt.Println("Objective reached!")

	// for !objectiveReached {
	// 	time.Sleep(100 * time.Millisecond)
	// 	grid.PrintGrid()

	// 	for _, agent := range agents {
	// 		objectiveReached = agent.Move()
	// 		if objectiveReached {
	// 			break
	// 		}
	// 	}

	// 	// newPos, err := grid.GenerateAStarPoint(agentAStarPos)

	// 	// if !err && !objectiveReached {
	// 	// 	fmt.Print("GO sa gosse pis sa me force a l'utiliser -_- \n")
	// 	// 	continue
	// 	// }

	// 	// if grid.Cells[newPos.Y][newPos.X].Load() == game.Objective {
	// 	// 	objectiveReached = true
	// 	// }

	// 	// if grid.MoveAgent(agentAStarPos, newPos) {
	// 	// 	agentAStarPos = newPos
	// 	// }
	// }
}
