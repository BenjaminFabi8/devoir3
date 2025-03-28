package main

import (
	"devoir3/src/customAtomic"
	"devoir3/src/game"
	"fmt"
	"math/rand"
	"time"
)

var directions = []game.Position{
	{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1},
}

func GetRandomMove(pos game.Position, rng *rand.Rand) game.Position {
	move := directions[rng.Intn(len(directions))]
	return game.Position{X: pos.X + move.X, Y: pos.Y + move.Y}
}

func main() {

	objectiveReached := false
	input := []string{
		"##########",
		"#A     # #",
		"###      #",
		"#   #    #",
		"#   A    #",
		"#      O #",
		"##########",
	}

	grid := game.NewGameGrid(input)
	grid.SetObjectives()
	agentRandomPos := game.Position{X: 4, Y: 4}
	agentAStarPos := game.Position{X: 1, Y: 1}

	// Create a new random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for !objectiveReached {
		//Clear console
		//fmt.Print("\033[H\033[2J")

		grid.PrintGrid()

		newPos := GetRandomMove(agentRandomPos, rng)

		if grid.Cells[newPos.Y][newPos.X] == customAtomic.NewSwappableRune(game.Objective) {
			objectiveReached = true
		}

		if grid.MoveAgent(agentRandomPos, newPos) {
			agentRandomPos = newPos
		}

		newPos, err := grid.GenerateAStarPoint(agentAStarPos)

		if !err && !objectiveReached {
			fmt.Print("GO sa gosse pis sa me force a l'utiliser -_- \n")
			continue
		}

		if grid.Cells[newPos.Y][newPos.X] == customAtomic.NewSwappableRune(game.Objective) {
			objectiveReached = true
		}

		if grid.MoveAgent(agentAStarPos, newPos) {
			agentAStarPos = newPos
		}

		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Objective reached!")
	grid.PrintGrid()
}
