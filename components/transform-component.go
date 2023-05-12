package components

import (
	"github.com/AllenDang/imgui-go"
	// "github.com/micahke/mango/util"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/util/math"
)

type TransformComponent struct {
	Name     string
	Position math.Vec3f
	Rotation math.Vec3f
	Scale    math.Vec3f
}

func (transform *TransformComponent) Init() {}

// Don't really need much functionality from this every frame
func (tranform *TransformComponent) Update() {}


func (transform *TransformComponent) renderDragFloats(id string, first, second, third *float32) {


	imgui.PushID(id)

	imgui.ColumnsV(2, "", false)
	imgui.SetColumnWidth(0, 100)
	imgui.SetColumnWidth(1, imgui.WindowWidth() - 100)

  // var itemWidth float32 = (300 - (3 * 32.5)) / 3
  var itemWidth float32 = ((imgui.WindowWidth() - 100) - (3 * 25)) / 3

	imgui.Text(id)
	imgui.NextColumn()

  imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0, Y:0})

  imgui.PushStyleColor(imgui.StyleColorButton, util.ImguiGenVec4(0.91, 0.30, 0.21, 1.0))
  imgui.PushStyleColor(imgui.StyleColorButtonHovered, util.ImguiGenVec4(1.0, 0.40, 0.31, 1.0))
  imgui.PushStyleColor(imgui.StyleColorButtonActive, util.ImguiGenVec4(1.0, 0.40, 0.31, 1.0))
	if imgui.Button("X") {
    *first = 0.0
	}
  imgui.PopStyleColorV(3)
	imgui.SameLine()
	imgui.PushItemWidth(itemWidth) // set a specific width for the DragFloat

	imgui.DragFloatV("##X", first, 1.0, 0.0, 0.0, "%.3f", 1.0)

	imgui.SameLine()

  imgui.PushStyleColor(imgui.StyleColorButton, util.ImguiGenVec4(0.15, 0.68, 0.38, 1.0))
  imgui.PushStyleColor(imgui.StyleColorButtonHovered, util.ImguiGenVec4(0.25, 0.78, 0.48, 1.0))
  imgui.PushStyleColor(imgui.StyleColorButtonActive, util.ImguiGenVec4(0.25, 0.78, 0.48, 1.0))
	if imgui.Button("Y") {
    *second = 0.0
	}
  imgui.PopStyleColorV(3)
  imgui.PopItemWidth()
	imgui.SameLine()
	imgui.PushItemWidth(itemWidth) // set a specific width for the DragFloat
	imgui.DragFloatV("##Y", second, 1.0, 0.0, 0.0, "%.3f", 1.0)
  imgui.PopItemWidth()

	imgui.SameLine()


  imgui.PushStyleColor(imgui.StyleColorButton, util.ImguiGenVec4(0.16, 0.50, 0.73, 1.0))
  imgui.PushStyleColor(imgui.StyleColorButtonHovered, util.ImguiGenVec4(0.26, 0.83, 0.83, 1.0))
  imgui.PushStyleColor(imgui.StyleColorButtonActive, util.ImguiGenVec4(0.26, 0.83, 0.83, 1.0))
	if imgui.Button("Z") {
    *third = 0.0
	}
  imgui.PopStyleColorV(3)
	imgui.SameLine()
	imgui.PushItemWidth(itemWidth) // set a specific width for the DragFloat
	imgui.DragFloatV("##Z", third, 1.0, 0.0, 0.0, "%.3f", 1.0)
  imgui.PopItemWidth()

	imgui.PopStyleVar()

	imgui.Columns()

	imgui.PopID()
}

// Allows the component to designate its own control panel within the editor
func (transform *TransformComponent) RenderControlPanel() {
  imgui.Spacing()

  transform.renderDragFloats("Translation", &transform.Position.X, &transform.Position.Y, &transform.Position.Z)
  imgui.Spacing()
  transform.renderDragFloats("Rotation", &transform.Rotation.X, &transform.Rotation.Y, &transform.Rotation.Z)
  imgui.Spacing()
  transform.renderDragFloats("Scale", &transform.Scale.X, &transform.Scale.Y, &transform.Scale.Z)

  imgui.Spacing()
  imgui.Spacing()
}




// Keeping this here to avoid using reflection
func (transform *TransformComponent) GetComponentName() string {
	return transform.Name
}
