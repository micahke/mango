package im

type IM_SCRIPT interface {
	Init()
	Update(deltaTime float32)
	Draw()
}
