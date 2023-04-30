package components

import (
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/util/math"
)

type TransformComponent struct {
	Position math.Vec3f
	Rotation float64
	Scale    float64
}

func (transform *TransformComponent) Init() {

}

func (tranform *TransformComponent) Update() {

	logging.DebugLog("Hello from tranform component")

}
