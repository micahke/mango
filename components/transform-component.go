package components

import (
	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/util/math"
)

type TransformComponent struct {
  Name string
	Position math.Vec3f
	Rotation math.Vec3f
	Scale    math.Vec3f
}


func (transform *TransformComponent) Init() {}

// Don't really need much functionality from this every frame
func (tranform *TransformComponent) Update() {}


// Allows the component to designate its own control panel within the editor
func (transform *TransformComponent) RenderControlPanel() {

  
  imgui.Text("Position")

  imgui.BeginChildV("Position", imgui.Vec2{X: 0, Y: 20}, false, 0)
  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("X", &transform.Position.X)

  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("Y", &transform.Position.Y)

  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("Z", &transform.Position.Z)

  imgui.EndChild()

  // Rotation

  imgui.Text("Rotation")

  imgui.BeginChildV("Rotation", imgui.Vec2{X: 0, Y: 20}, false, 0)
  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("X", &transform.Rotation.X)

  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("Y", &transform.Rotation.Y)

  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("Z", &transform.Rotation.Z)

  imgui.EndChild()


  imgui.Text("Scale")

  imgui.BeginChildV("Scale", imgui.Vec2{X: 0, Y: 20}, false, 0)
  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("X", &transform.Scale.X)

  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("Y", &transform.Scale.Y)

  imgui.SameLine()
  imgui.PushItemWidth(75)
  imgui.InputFloat("Z", &transform.Scale.Z)

  imgui.EndChild()


}

// Keeping this here to avoid using reflection
func (transform *TransformComponent) GetComponentName() string {
  return transform.Name
}
