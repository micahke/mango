package components

import "github.com/micahke/infinite-universe/mango/logging"


type TransformComponent struct {
	X float32
	Y float32
}

func (transform *TransformComponent) Init() {

}

func (tranform *TransformComponent) Update() {

  logging.DebugLog("Hello from tranform component")

}
