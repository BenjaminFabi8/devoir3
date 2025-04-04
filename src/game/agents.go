package game

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Position struct {
	X, Y int
}

func (pos Position) Equals(other Position) bool {
	return pos.X == other.X && pos.Y == other.Y
}

func GetRandomMove(pos Position) Position {
	move := directions[rand.Intn(len(directions))]
	return Position{X: pos.X + move.X, Y: pos.Y + move.Y}
}

type Agent interface {
	Move() bool
	GetPosition() Position
	SetPosition(Position)
	GetObjectivePosition() Position
	GetLogEntries() []LogEntry
	AddLogEntry(LogEntry)
	GetId() int
	IsReached() bool
	SetReached(bool)
}

type BaseAgent struct {
	Id                int
	Position          Position
	ObjectivePosition Position
	LogEntries        []LogEntry
	gameGrid          *Grid
	reached           bool
}

func (a *BaseAgent) GetPosition() Position          { return a.Position }
func (a *BaseAgent) SetPosition(pos Position)       { a.Position = pos }
func (a *BaseAgent) GetObjectivePosition() Position { return a.ObjectivePosition }
func (a *BaseAgent) AddLogEntry(entry LogEntry)     { a.LogEntries = append(a.LogEntries, entry) }
func (a *BaseAgent) GetLogEntries() []LogEntry      { return a.LogEntries }
func (a *BaseAgent) GetId() int                     { return a.Id }
func (a *BaseAgent) IsReached() bool                { return a.reached }
func (a *BaseAgent) SetReached(reached bool)        { a.reached = reached }

func StartAgents(agents []Agent) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, agent := range agents {
		go func(a Agent) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if a.Move() {
						cancel()
						return
					}
					time.Sleep(1 * time.Millisecond) // Simulate agent movement delay
				}
			}
		}(agent)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Objective reached!")
		return
	}
}

type RandomAgent struct {
	BaseAgent
}

type AStarAgent struct {
	BaseAgent
}

type AStarWaitingAgent struct {
	AStarAgent
}

func NewRandomAgent(id int, pos Position, grid *Grid) *RandomAgent {
	return &RandomAgent{
		BaseAgent: BaseAgent{
			Id:         id,
			Position:   pos,
			LogEntries: []LogEntry{},
			gameGrid:   grid,
		},
	}
}

func NewAStartAgent(id int, pos Position, grid *Grid) *AStarAgent {
	bestObjective, _ := grid.GetClosestObjective(pos)

	return &AStarAgent{
		BaseAgent: BaseAgent{
			Id:                id,
			Position:          pos,
			ObjectivePosition: bestObjective,
			LogEntries:        []LogEntry{},
			gameGrid:          grid,
		},
	}
}

func NewAStartWaitingAgent(id int, pos Position, grid *Grid) *AStarWaitingAgent {
	bestObjective, _ := grid.GetClosestObjective(pos)

	return &AStarWaitingAgent{
		AStarAgent: AStarAgent{
			BaseAgent: BaseAgent{
				Id:                id,
				Position:          pos,
				ObjectivePosition: bestObjective,
				LogEntries:        []LogEntry{},
				gameGrid:          grid,
			},
		},
	}
}

func (a *RandomAgent) Move() bool {
	newPos := GetRandomMove(a.Position)

	return a.MoveSelf(newPos)
}

func (a *AStarAgent) Move() bool {
	newPos, _ := a.GenerateAStarPoint(a.Position)

	return a.MoveSelf(newPos)
}

func (a *BaseAgent) MoveSelf(newPos Position) bool {
	if a.gameGrid.Cells[newPos.Y][newPos.X].Load() == Objective {
		a.AddLogEntry(LogEntry{Id: a.Id, Position: newPos, Timestamp: time.Now()})
		a.gameGrid.MoveToObjective(a.Position, newPos)
		a.SetPosition(newPos)
		return true
	}

	if a.gameGrid.MoveAgent(a.Position, newPos) {
		a.SetPosition(newPos)
		a.AddLogEntry(LogEntry{Id: a.Id, Position: newPos, Timestamp: time.Now()})
	}

	return false
}

func (a *AStarWaitingAgent) Move() bool {
	result := a.AStarAgent.Move()

	//Not sure where to put it...
	time.Sleep(10 * time.Millisecond)

	return result
}

func (aStarAgent *AStarAgent) GenerateAStarPoint(start Position) (Position, bool) {
	queue := []Position{start}
	visited := make(map[Position]bool)
	parent := make(map[Position]Position)
	visited[start] = true

	for len(queue) > 0 {
		nextQueue := []Position{}
		for _, current := range queue {
			if current == aStarAgent.ObjectivePosition {
				step := aStarAgent.ObjectivePosition
				for parent[step] != start {
					step = parent[step]
				}
				return step, true
			}

			for _, direction := range directions {
				newPos := Position{X: current.X + direction.X, Y: current.Y + direction.Y}
				if aStarAgent.gameGrid.IsValidMove(newPos) && !visited[newPos] {
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
