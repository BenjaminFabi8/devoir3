package main

import (
	"devoir3/src/agents"
	"fmt"
	"math/rand"
	"time"
)

const (
	Empty     = ' '
	Obstacle  = '#'
	Objective = 'O'
	AgentChar = 'A'
)

type Grid struct {
	Cells      [][]rune
	Rows       int
	Cols       int
	Objectives []Position
}

var directions = []Position{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func InitializeGrid(input []string) *Grid {
	rows := len(input)
	cols := len(input[0])
	cells := make([][]rune, rows)
	for i := range rows {
		cells[i] = []rune(input[i])
	}

	return &Grid{Cells: cells, Rows: rows, Cols: cols}
}

func (g *Grid) SetObjectives() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if g.Cells[i][j] == Objective {
				g.Objectives = append(g.Objectives, Position{j, i})
			}
		}
	}
}

func (g *Grid) PrintGrid() {
	for _, row := range g.Cells {
		fmt.Println(string(row))
	}
}

// IsValidMove checks if a move is valid
func (g *Grid) IsValidMove(pos agents.Position) bool {
	if pos.X < 0 || pos.X >= g.Rows || pos.Y < 0 || pos.Y >= g.Cols {
		return false
	}
	return g.Cells[pos.Y][pos.X] == Empty || g.Cells[pos.Y][pos.X] == Objective
}

// MoveAgent moves an agent to a new position if valid
func (g *Grid) MoveAgent(current, next agents.Position) bool {
	if g.IsValidMove(next) {
		g.Cells[current.Y][current.X] = Empty
		g.Cells[next.Y][next.X] = AgentChar
		return true
	}
	return false
}

func GetRandomMove(pos Position, rng *rand.Rand) Position {
	move := directions[rng.Intn(len(directions))]
	return Position{X: pos.X + move.X, Y: pos.Y + move.Y}
}

func GenerateAStarPoint(g *Grid, start Position) (Position, bool) {
	//A faire dans le constructeur
	bestObjective, found := g.GetClosestObjective(start)
	if !found {
		fmt.Println("Oups, pas d'objectif mon grand...")
		return Position{}, false
	}
	//fmt.Printf("Objectif Found: %d, %d \n", bestObjective.X, bestObjective.Y)

	queue := []Position{start}
	visited := make(map[Position]bool)
	parent := make(map[Position]Position)
	visited[start] = true

	for len(queue) > 0 {
		nextQueue := []Position{}
		for _, current := range queue {
			if current == bestObjective {
				step := bestObjective
				for parent[step] != start {
					step = parent[step]
				}
				return step, true
			}

			for _, direction := range directions {
				newPos := Position{current.X + direction.X, current.Y + direction.Y}
				if g.IsValidMove(newPos) && !visited[newPos] {
					nextQueue = append(nextQueue, newPos)
					visited[newPos] = true
					parent[newPos] = current
				}
			}
		}
		queue = nextQueue
	}

	return Position{}, false
}

func (g *Grid) GetClosestObjective(start Position) (Position, bool) {

	if len(g.Objectives) == 1 {
		return g.Objectives[0], true
	}

	closest := Position{}
	minSteps := 10000

	for _, obj := range g.Objectives {
		steps := g.GetDistanceForObjectif(start, obj)
		if steps != -1 && steps < minSteps {
			minSteps = steps
			closest = obj
		}
	}

	if minSteps == -1 {
		return Position{}, false
	}

	return closest, true
}

func (g *Grid) GetDistanceForObjectif(start, objective Position) int {
	queue := []Position{start}
	visited := make(map[Position]bool)
	visited[start] = true
	steps := 0

	for len(queue) > 0 {
		nextQueue := []Position{}
		for _, current := range queue {
			if current == objective {
				return steps
			}

			for _, direction := range directions {
				newPos := Position{current.X + direction.X, current.Y + direction.Y}
				if g.IsValidMove(newPos) && !visited[newPos] {
					nextQueue = append(nextQueue, newPos)
					visited[newPos] = true
				}
			}

		}
		queue = nextQueue
		steps++
	}

	return -1
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

	grid := InitializeGrid(input)
	grid.SetObjectives()
	agentRandomPos := Position{X: 4, Y: 4}
	agentAStarPos := Position{X: 1, Y: 1}

	// Create a new random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for !objectiveReached {
		//Clear console
		//fmt.Print("\033[H\033[2J")

		grid.PrintGrid()

		newPos := GetRandomMove(agentRandomPos, rng)

		if grid.Cells[newPos.Y][newPos.X] == Objective {
			objectiveReached = true
		}

		if grid.MoveAgent(agentRandomPos, newPos) {
			agentRandomPos = newPos
		}

		newPos, err := GenerateAStarPoint(grid, agentAStarPos)

		if !err {
			fmt.Print("GO sa gosse pis sa me force a l'utiliser -_- \n")
			continue
		}

		if grid.Cells[newPos.Y][newPos.X] == Objective {
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
