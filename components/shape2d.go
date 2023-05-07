package components

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/logging"
)

type Shape2DComponent struct {
	Name  string
	Shape shape.IShape

	selectedShapeIndex int
}

const (
	SHAPE_RECT    int = 0
	SHAPE_ELLIPSE int = 1
	SHAPE_LINE    int = 2
)

var SHAPES_LIST = [3]string{
	"Rect",
	"Ellipse",
	"Line",
}

func (component *Shape2DComponent) Init() {
	fmt.Println("Initializing shape")
	component.Name = "Shape2D"
}

func (component *Shape2DComponent) Update() {}

func (component *Shape2DComponent) GetComponentName() string {
	return component.Name
}

func (component *Shape2DComponent) RenderControlPanel() {

	imgui.Spacing()

	if imgui.BeginCombo("Select Shape", "---") {

		for index, shape := range SHAPES_LIST {

			if imgui.Selectable(shape) {
				component.selectedShapeIndex = index
				component.processShapeSelection(index)
			}

		}

		imgui.EndCombo()

	}
	imgui.Spacing()

  logging.DebugLog(component.Shape)

	if component.Shape == nil {
		return
	}

	imgui.Text(SHAPES_LIST[component.selectedShapeIndex])

  imgui.Spacing()

  // Check to see if the current shape is a Rect
  if rect, ok := component.Shape.(*shape.Rect); ok {
    component.drawRectPanel(rect)
  }

  imgui.Spacing()

}

func (component *Shape2DComponent) drawRectPanel(rect *shape.Rect) {

  imgui.PushItemWidth(75)
  imgui.InputFloat("Width", &rect.Width)
  
  imgui.PushItemWidth(75)
  imgui.InputFloat("Height", &rect.Height)

}

func (component *Shape2DComponent) processShapeSelection(shapeIndex int) {
	switch shapeIndex {
	case SHAPE_RECT:
		component.SetShape(&shape.Rect{
			Width:  100,
			Height: 100,
		})
	default:
    component.SetShape(nil)
	}
}

func (component *Shape2DComponent) SetShape(shape shape.IShape) {
	component.Shape = shape
}
