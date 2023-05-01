package components

import (
	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/util/math"
)

type TransformComponent struct {
  Name string
	Position math.Vec3f
	Rotation float64
	Scale    float64
}


func (transform *TransformComponent) Init() {}

// Don't really need much functionality from this every frame
func (tranform *TransformComponent) Update() {}


// Allows the component to designate its own control panel within the editor
func (transform *TransformComponent) RenderControlPanel() {
  imgui.Text("Hello, world!")
}

// Keeping this here to avoid using reflection
func (transform *TransformComponent) GetComponentName() string {
  return transform.Name
}
