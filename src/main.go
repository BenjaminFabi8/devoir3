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
	Objectives []agents.Position
}

var directions = []agents.Position{
	{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1},
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
	for i := range g.Rows {
		for j := range g.Cols {
			if g.Cells[i][j] == Objective {
				g.Objectives = append(g.Objectives, agents.Position{X: j, Y: i})
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
	if pos.Y < 0 || pos.Y >= g.Rows || pos.X < 0 || pos.X >= g.Cols {
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

func GetRandomMove(pos agents.Position, rng *rand.Rand) agents.Position {
	move := directions[rng.Intn(len(directions))]
	return agents.Position{X: pos.X + move.X, Y: pos.Y + move.Y}
}

func GenerateAStarPoint(g *Grid, start agents.Position) (agents.Position, bool) {
	//A faire dans le constructeur
	bestObjective, found := g.GetClosestObjective(start)
	if !found {
		fmt.Println("Oups, pas d'objectif mon grand...")
		return agents.Position{}, false
	}
	//fmt.Printf("Objectif Found: %d, %d \n", bestObjective.X, bestObjective.Y)

	queue := []agents.Position{start}
	visited := make(map[agents.Position]bool)
	parent := make(map[agents.Position]agents.Position)
	visited[start] = true

	for len(queue) > 0 {
		nextQueue := []agents.Position{}
		for _, current := range queue {
			if current == bestObjective {
				step := bestObjective
				for parent[step] != start {
					step = parent[step]
				}
				return step, true
			}

			for _, direction := range directions {
				newPos := agents.Position{X: current.X + direction.X, Y: current.Y + direction.Y}
				if g.IsValidMove(newPos) && !visited[newPos] {
					nextQueue = append(nextQueue, newPos)
					visited[newPos] = true
					parent[newPos] = current
				}
			}
		}
		queue = nextQueue
	}

	return agents.Position{}, false
}

func (g *Grid) GetClosestObjective(start agents.Position) (agents.Position, bool) {

	if len(g.Objectives) == 1 {
		return g.Objectives[0], true
	}

	closest := agents.Position{}
	minSteps := 10000

	for _, obj := range g.Objectives {
		steps := g.GetDistanceForObjectif(start, obj)
		if steps != -1 && steps < minSteps {
			minSteps = steps
			closest = obj
		}
	}

	if minSteps == -1 {
		return agents.Position{}, false
	}

	return closest, true
}

func (g *Grid) GetDistanceForObjectif(start, objective agents.Position) int {
	queue := []agents.Position{start}
	visited := make(map[agents.Position]bool)
	visited[start] = true
	steps := 0

	for len(queue) > 0 {
		nextQueue := []agents.Position{}
		for _, current := range queue {
			if current == objective {
				return steps
			}

			for _, direction := range directions {
				newPos := agents.Position{X: current.X + direction.X, Y: current.Y + direction.Y}
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
	agentRandomPos := agents.Position{X: 4, Y: 4}
	agentAStarPos := agents.Position{X: 1, Y: 1}

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

		if !err && !objectiveReached {
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
