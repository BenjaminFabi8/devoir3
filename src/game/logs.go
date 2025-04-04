package game

import (
	"fmt"
	"sort"
	"time"
	"strings"
    "path/filepath"
	"devoir3/src/utils"
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

func CreateLogEntriesFile(logs []LogEntry) {
	/*===MERGE LOGS===*/
	mergedLogs := GetMergedLogEntriesString(logs)
	//fmt.Println("merged logs : \n" + mergedLogs)
	outputFile := filepath.Join("logs", "log_"+time.Now().Format("2006-01-02_15-04-05")+".txt")
	utils.OutputStringToFile(mergedLogs, outputFile)
}

