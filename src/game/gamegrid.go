package game

import (
	"devoir3/src/customAtomic"
	"fmt"
	"strings"
)

const (
	Empty     = ' '
	Obstacle  = '#'
	Objective = 'O'
	AgentChar = 'A'
)

var directions = []Position{
	{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1},
}

type Grid struct {
	Cells      [][]customAtomic.SwappableRune
	Rows       int
	Cols       int
	Objectives []Position
}

func NewGameGrid(input []string) *Grid {
	rows := len(input)
	cols := len(input[0])
	cells := make([][]customAtomic.SwappableRune, rows)

	for i := range rows {
		cells[i] = make([]customAtomic.SwappableRune, cols)
		for j := range cols {
			cells[i][j] = customAtomic.NewSwappableRune(rune(input[i][j]))
		}
	}

	return &Grid{Cells: cells, Rows: rows, Cols: cols}
}

func (g *Grid) SetObjectives() {
	for i := range g.Rows {
		for j := range g.Cols {
			if g.Cells[i][j].Load() == Objective {
				g.Objectives = append(g.Objectives, Position{X: j, Y: i})
			}
		}
	}
}

func (g *Grid) GetAgentsPositions() []Position {
	positions := []Position{}
	for i := range g.Rows {
		for j := range g.Cols {
			if g.Cells[i][j].Load() == AgentChar {
				positions = append(positions, Position{X: j, Y: i})
			}
		}
	}
	return positions
}

func (g *Grid) PrintGrid() {
	builder := strings.Builder{}
	for _, row := range g.Cells {
		for _, cell := range row {
			builder.WriteRune(cell.Load())
		}
		builder.WriteRune('\n')
	}
	fmt.Println(builder.String())
}

// IsValidMove checks if a move is valid
func (g *Grid) IsValidMove(pos Position) bool {
	if pos.Y < 0 || pos.Y >= g.Rows || pos.X < 0 || pos.X >= g.Cols {
		return false
	}
	return g.Cells[pos.Y][pos.X].Load() == Empty || g.Cells[pos.Y][pos.X].Load() == Objective
}

// MoveAgent moves an agent to a new position if valid
func (g *Grid) MoveAgent(current, next Position) bool {
	if g.IsValidMove(next) {
		g.Cells[current.Y][current.X].SwapAtom(&g.Cells[next.Y][next.X], Empty)
		return true
	}
	return false
}

func (g *Grid) MoveToObjective(agentPos Position, objectivePos Position) {
	g.Cells[agentPos.Y][agentPos.X].Swap(Empty)
	g.Cells[objectivePos.Y][objectivePos.X].Swap(AgentChar)
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

func (g *Grid) GenerateAStarPoint(start Position) (Position, bool) {
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
				newPos := Position{X: current.X + direction.X, Y: current.Y + direction.Y}
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
				newPos := Position{X: current.X + direction.X, Y: current.Y + direction.Y}
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
