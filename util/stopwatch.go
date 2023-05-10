package util

import "time"



type Stopwatch struct {
  start time.Time
  duration time.Duration
}


func (stopwatch *Stopwatch) Start() *Stopwatch {
  return &Stopwatch{
    start: time.Now(),
  }
}


func (stopwatch *Stopwatch) Stop() time.Duration {
  stopwatch.duration = time.Since(stopwatch.start)
  return stopwatch.duration 
}


func (stopwatch *Stopwatch) Peek() time.Duration {
  return time.Since(stopwatch.start)
}


func (stopwatch *Stopwatch) Reset() {
  stopwatch.start = time.Time{}

  // FIXME: take a look at how we reset stopwatch duration
  stopwatch.duration = 0
}
