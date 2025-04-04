package main

import (
	"devoir3/src/game"
	"fmt"
	"sync"
	"time"
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

	// Create a new random number generator
	objectiveReached := make(chan bool)
	dones := make([]chan bool, len(agents))

	wg := &sync.WaitGroup{}

	for i, agent := range agents {
		wg.Add(1)
		dones[i] = game.StartAgent(agent, objectiveReached)
	}

	go func() {
		select {
		case <-objectiveReached:
			grid.PrintGrid()
			for i := range agents {
				wg.Done()
				dones[i] <- true
			}
		}
	}()

	wg.Wait()

	for _, agent := range agents {
		fmt.Printf("Agent %d Log Entries: %d\n", agent.GetId(), len(agent.GetLogEntries()))
		for _, entry := range agent.GetLogEntries() {
			fmt.Printf("Position: (%d, %d), Timestamp: %s\n", entry.Position.X, entry.Position.Y, entry.Timestamp.Format(time.RFC3339))
		}
	}

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
