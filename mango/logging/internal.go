package logging

import "time"



func DebugLog(args ...any) {

  item := new(LogItem)

  item.logType = LOG_LOG
  item.source = ENGINE
  item.content = args
  item.timestamp = time.Now()

  _log = append(_log, item)
}


func DebugLogError(args ...any) {

  item := new(LogItem)

  item.logType = LOG_ERROR
  item.source = ENGINE
  item.content = args
  item.timestamp = time.Now()

  _log = append(_log, item)

}


