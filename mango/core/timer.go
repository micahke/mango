package core

import "github.com/go-gl/glfw/v3.3/glfw"

type MTimer struct {
	programTime  float64
	deltaTime    float64
	timeDilation float64

	frameTime float64
	fps       float64
}

var Timer *MTimer

func InitTimer() {
	Timer = new(MTimer)
	Timer.programTime = glfw.GetTime()
}

func (timer *MTimer) SetTimeDilation(dilation float64) {
	timer.timeDilation = dilation
}

func (timer *MTimer) ProgramTime() float64 {
	return timer.programTime
}

func (timer *MTimer) DeltaTime() float64 {
	return timer.deltaTime
}

// Should only be called from the engine
func (timer *MTimer) Update() {
	now := glfw.GetTime()
	timer.deltaTime = now - timer.programTime
	timer.programTime = now
}

func (timer *MTimer) UpdateFrameData(start, end float64) {
	timer.frameTime = end - start

	timer.fps = 1.0 / timer.frameTime
}

func (timer *MTimer) FrameTime() float64 {
	return timer.frameTime
}

func (timer *MTimer) FPS() float64 {
	return timer.fps
}
