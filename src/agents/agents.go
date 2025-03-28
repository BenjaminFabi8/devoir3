package agents

type Position struct {
	X, Y int
}

type Agent interface {
	Move()
}

type GridAgent struct {
	Id                int
	Appearance        rune
	Position          Position
	ObjectivePosition Position
	LogEntries        []LogEntry
	moveFunc          func()
}

func (a *GridAgent) Move() {
	a.moveFunc()
}

func NewRandomAgent(id int, appearance rune, pos Position, objPos Position) *GridAgent {
	return &GridAgent{
		Id:                id,
		Appearance:        appearance,
		Position:          pos,
		ObjectivePosition: objPos,
	}
}

func NewAStarAgent(id int, name rune, pos Position, objPos Position) *GridAgent {
	return &GridAgent{
		Id:                id,
		Appearance:        name,
		Position:          pos,
		ObjectivePosition: objPos,
	}
}

func NewSlowAStarAgent(id int, name rune, pos Position, objPos Position) *GridAgent {
	return &GridAgent{
		Id:                id,
		Appearance:        name,
		Position:          pos,
		ObjectivePosition: objPos,
	}
}
