package game

import (
	"fmt"
	"sort"
	"time"
	"strings"
)

type LogEntry struct {
	Id        int
	Position  Position
	Timestamp time.Time
}

func (log LogEntry) GetStringPosition() string {
	return fmt.Sprintf("%d|%d:%d", log.Id, log.Position.X, log.Position.Y)
}

func GetMergedLogEntriesString(logs []LogEntry) string {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Timestamp.Before(logs[j].Timestamp)
	})

	var builder strings.Builder

	WelcomeMessage := 
	"===========================\nAGENT LOGS\n===========================\nAgent | Position (X:Y)\n---------------------------\n"
    builder.WriteString(WelcomeMessage)
	for i, log := range logs {
        builder.WriteString(log.GetStringPosition())
        if i < len(logs)-1 {
            builder.WriteString("\n")
        }
    }

    return builder.String()
}

