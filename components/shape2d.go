package components

import (
	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/components/shape"
)

type Shape2DComponent struct {
	Name  string
	Shape shape.IShape

	selectedShapeIndex int
}

type Shape2D int

const (
	SHAPE_NONE    Shape2D = -1
	SHAPE_RECT    Shape2D = 0
	SHAPE_ELLIPSE Shape2D = 1
	SHAPE_LINE    Shape2D = 2
)

var SHAPES_LIST = [3]string{
	"Rect",
	"Ellipse",
	"Line",
}

func (component *Shape2DComponent) Init() {
	component.Name = "Shape2D"
}

func (component *Shape2DComponent) Update() {}

func (component *Shape2DComponent) GetComponentName() string {
	return component.Name
}

// Determines what specific shape this is
func (component *Shape2DComponent) Determine() Shape2D {

	_, ok := component.Shape.(*shape.Rect)
	if ok {
		return SHAPE_RECT
	}

	return SHAPE_NONE

}

func (component *Shape2DComponent) RenderControlPanel() {

	imgui.Spacing()

	if imgui.BeginCombo("Select Shape", SHAPES_LIST[component.selectedShapeIndex]) {

		for index, shape := range SHAPES_LIST {

			if imgui.Selectable(shape) {
				component.selectedShapeIndex = index
				component.processShapeSelection(index)
			}

		}

		imgui.EndCombo()

	}
	imgui.Spacing()

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

  imgui.PushID("rect_controls")

  imgui.ColumnsV(2, "width_controls", false) 
  imgui.SetColumnWidth(0, 100)
  imgui.SetColumnWidth(1, imgui.WindowWidth() - 100)

  imgui.Text("Width")
  imgui.NextColumn()
  imgui.DragFloatV("##width", &rect.Width, 1.0, 0.0, 0.0, "%.3f", 1.0)

  imgui.Columns()

  imgui.ColumnsV(2, "height_controls", false) 
  imgui.SetColumnWidth(0, 100)
  imgui.SetColumnWidth(1, imgui.WindowWidth() - 100)

  imgui.Text("Height")
  imgui.NextColumn()
  imgui.DragFloatV("##height", &rect.Height, 1.0, 0.0, 0.0, "%.3f", 1.0)

  imgui.Columns()

  imgui.PopID()

}

func (component *Shape2DComponent) processShapeSelection(shapeIndex int) {
	switch shapeIndex {
	case int(SHAPE_RECT):
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
