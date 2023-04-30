package components

import (
	"github.com/micahke/mango/util/math"
)

type TransformComponent struct {
	Position math.Vec3f
	Rotation float64
	Scale    float64
}


func (transform *TransformComponent) Init() {}

// Don't really need much functionality from this every frame
func (tranform *TransformComponent) Update() {}
