package main

import (
	"fmt"
    "math/rand"
    "time"
)

// Constants for grid symbols
const (
	Empty    = ' '
	Obstacle = '#'
	Objective = 'O'
	Agent   = 'A'
)

// Grid represents the game grid
type Grid struct {
	Cells [][]rune
	Rows  int
	Cols  int
}

// Position represents a position in the grid
type Position struct {
	X, Y int
}




// InitializeGrid initializes the grid from a given ASCII representation
func InitializeGrid(input []string) *Grid {
	rows := len(input)
	cols := len(input[0])
	cells := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		cells[i] = []rune(input[i])
	}
	return &Grid{Cells: cells, Rows: rows, Cols: cols}
}

// PrintGrid displays the grid
func (g *Grid) PrintGrid() {
	for _, row := range g.Cells {
		fmt.Println(string(row))
	}
}

// IsValidMove checks if a move is valid
func (g *Grid) IsValidMove(pos Position) bool {
	if pos.X < 0 || pos.X >= g.Rows || pos.Y < 0 || pos.Y >= g.Cols {
		return false
	}
	return g.Cells[pos.X][pos.Y] == Empty || g.Cells[pos.X][pos.Y] == Objective
}

// MoveAgent moves an agent to a new position if valid
func (g *Grid) MoveAgent(current, next Position) bool {
	if g.IsValidMove(next) {
		g.Cells[current.X][current.Y] = Empty
		g.Cells[next.X][next.Y] = Agent
		return true
	}
	return false
}

func GetRandomMove(pos Position, rng *rand.Rand) Position {
    directions := []Position{
        {X: -1, Y: 0}, // Up
        {X: 1, Y: 0},  // Down
        {X: 0, Y: -1}, // Left
        {X: 0, Y: 1},  // Right
    }
    move := directions[rng.Intn(len(directions))]
    return Position{X: pos.X + move.X, Y: pos.Y + move.Y}
}

func main() {

	objectiveReached := false
	// Example grid input
	input := []string{
		"##########",
		"#A O   # #",
		"# #      #",
		"#   #    #",
		"#        #",
		"#        #",	
		"##########",
	}

    grid := InitializeGrid(input)
    agentPos := Position{X: 1, Y: 1}

	 // Create a new random number generator
	 rng := rand.New(rand.NewSource(time.Now().UnixNano()))

    for !objectiveReached {
        // Clear the console
        fmt.Print("\033[H\033[2J")

        // Print the grid
        grid.PrintGrid()


		//RANDOM AGENT
        // Generate a random move
        newPos := GetRandomMove(agentPos, rng)
		// Check if the agent will reached the objective
		if grid.Cells[newPos.X][newPos.Y] == Objective {
			objectiveReached = true
		}
        // Move the agent if the move is valid
        if grid.MoveAgent(agentPos, newPos)  {
            agentPos = newPos
        }


        time.Sleep(100 * time.Millisecond)
    }
	fmt.Println("Objective reached!")
}