package core

import "github.com/go-gl/glfw/v3.3/glfw"


type Timer struct {

  programTime float32
  deltaTime float32
  timeDilation float32

}

func NewTimer() *Timer {
  timer := new(Timer)
  timer.programTime = float32(glfw.GetTime())
  return timer
}

func (timer *Timer) SetTimeDilation(dilation float32) {
  timer.timeDilation = dilation
}


func (timer *Timer) ProgramTime() float32 {
  return timer.programTime
}


func (timer *Timer) DeltaTime() float32 {
  return timer.deltaTime
}


// Should only be called from the engine
func (timer *Timer) Update() {
  now := float32(glfw.GetTime())
  timer.deltaTime = now - timer.programTime
  timer.programTime = now
}
