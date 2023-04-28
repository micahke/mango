package logging

import "time"

func Log(args ...any) {

	item := new(LogItem)

	item.logType = LOG_LOG
	item.source = APP
	item.content = args
	item.timestamp = time.Now()

	_log = append(_log, item)
}

func LogError(args ...any) {

	item := new(LogItem)

	item.logType = LOG_ERROR
	item.source = APP
	item.content = args
	item.timestamp = time.Now()

	_log = append(_log, item)

}
