package core

import "github.com/go-gl/glfw/v3.3/glfw"

type Timer struct {
	programTime  float64
	deltaTime    float64
	timeDilation float64

	frameTime float64
	fps       float64
}

func NewTimer() *Timer {
	timer := new(Timer)
	timer.programTime = glfw.GetTime()
	return timer
}

func (timer *Timer) SetTimeDilation(dilation float64) {
	timer.timeDilation = dilation
}

func (timer *Timer) ProgramTime() float64 {
	return timer.programTime
}

func (timer *Timer) DeltaTime() float64 {
	return timer.deltaTime
}

// Should only be called from the engine
func (timer *Timer) Update() {
	now := glfw.GetTime()
	timer.deltaTime = now - timer.programTime
	timer.programTime = now
}

func (timer *Timer) UpdateFrameData(start, end float64) {
	timer.frameTime = end - start

	timer.fps = 1.0 / timer.frameTime
}

func (timer *Timer) FrameTime() float64 {
	return timer.frameTime
}

func (timer *Timer) FPS() float64 {
	return timer.fps
}
