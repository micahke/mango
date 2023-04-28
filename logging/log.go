package logging

import "time"

type LogItem struct {
	logType   LogType
	source    LogSource
	content   []any
	timestamp time.Time
}

type LogSource int
type LogType int

const (
	ENGINE LogSource = 0 // Represents the engine layer
	APP    LogSource = 1 // Represents the game layer

	LOG_LOG   LogType = 0
	LOG_ERROR LogType = 1
)

var _log []*LogItem
