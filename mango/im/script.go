package im

type IM_SCRIPT interface {
	Init()
	Update(deltaTime float64)
	Draw()
}
