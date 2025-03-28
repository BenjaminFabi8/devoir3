package main

import (
	"fmt"
)

// Constants for grid symbols
const (
	Empty    = ' '
	Obstacle = '#'
	Objective = 'O'
	Agent    = 'A'
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

func main() {
	// Example grid input
	input := []string{
		"#####",
		"#A  #",
		"# O #",
		"#####",
	}

	grid := InitializeGrid(input)
	fmt.Println("Initial Grid:")
	grid.PrintGrid()

	// Example agent movement
	agentPos := Position{X: 1, Y: 1}
	newPos := Position{X: 2, Y: 1}

	if grid.MoveAgent(agentPos, newPos) {
		fmt.Println("\nGrid after moving agent:")
		grid.PrintGrid()
	} else {
		fmt.Println("\nInvalid move!")
	}
}