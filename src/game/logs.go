package game

import (
	"fmt"
	"time"
)

type LogEntry struct {
	Id        int
	Position  Position
	TimeStamp time.Time
}

func (log LogEntry) GetStringPosition() string {
	return fmt.Sprintf("%d|%d:%d", log.Id, log.Position.X, log.Position.Y)
}
