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

func NewRandomAgent(id int, pos Position, objectivePos Position, grid *Grid) *RandomAgent {
	return &RandomAgent{
		BaseAgent: BaseAgent{
			Id:                id,
			Position:          pos,
			ObjectivePosition: objectivePos,
			LogEntries:        []LogEntry{},
			gameGrid:          grid,
		},
	}
}

func (a *RandomAgent) Move() bool {
	if a.reached {
		return false
	}
	newPos := GetRandomMove(a.Position)
	if a.gameGrid.Cells[newPos.Y][newPos.X].Load() == Objective {
		a.AddLogEntry(LogEntry{Id: a.Id, Position: newPos, Timestamp: time.Now()})
		a.gameGrid.MoveToObjective(a.Position, newPos)
		return true
	}

	if a.gameGrid.MoveAgent(a.Position, newPos) {
		a.SetPosition(newPos)
		a.AddLogEntry(LogEntry{Id: a.Id, Position: newPos, Timestamp: time.Now()})
	}

	return false
}
