package components

import (
	"sort"
	"strings"

	"github.com/AllenDang/imgui-go"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/util/color"
	"github.com/micahke/mango/util/loaders"
)

type PrimitiveRenderer struct {
	Color color.Color
	Shape shape.IShape
  Texture *opengl.Texture

	selectedIndex int
}

type Shape2D int

const (
	SHAPE_NONE   Shape2D = -1
	SHAPE_RECT   Shape2D = 0
	SHAPE_CIRCLE Shape2D = 1
	SHAPE_LINE   Shape2D = 2
)

type shape_list_item struct {
	name      string
	shapeType Shape2D
}

var shapeList []shape_list_item = []shape_list_item{
	{name: "NONE", shapeType: -1},
	{name: "RECT", shapeType: 0},
	{name: "CIRCLE", shapeType: 1},
	{name: "LINE", shapeType: 2},
}

func (component *PrimitiveRenderer) Init() {
	// component.selectedIndex = -1
	component.Color = color.WHITE
}

func (component *PrimitiveRenderer) Update() {

}

func (component *PrimitiveRenderer) RenderControlPanel() {

	imgui.Spacing()

	if imgui.BeginCombo("Shape Select", shapeList[component.selectedIndex].name) {
		for index, shape := range shapeList {

			if imgui.Selectable(shape.name) {
				component.selectedIndex = index
				component.handleShapeSelection(shape)
			}

		}
		imgui.EndCombo()
	}


	imgui.Spacing()

	component.drawControls()

	imgui.Spacing()

  component.drawTextureSelection()

  imgui.Spacing()

	// Draw the color field which works for any primitive
	component.drawColorField()


}

func (component *PrimitiveRenderer) drawTextureSelection() {

  selectedName := ""
  files := loaders.GetFilesList()
  sort.Strings(files)

  if component.Texture == nil  {
    selectedName = "None"
  } else {
    selectedName = component.Texture.GetPath()
  }
	if imgui.BeginCombo("Select Texture", selectedName) {
		for index, textureName := range(files)  {

			if imgui.Selectable(textureName) {
        logging.DebugLog(index)
        component.buildAndSetTexture(textureName)
			}

		}
		imgui.EndCombo()
	}
}

func (component *PrimitiveRenderer) SetTexture(textureName string) {
  component.buildAndSetTexture(textureName)
}

func (component *PrimitiveRenderer) buildAndSetTexture(textureName string) {
  var imageLoader *loaders.ImageLoader
  textureLowercase := strings.ToLower(textureName)
  if strings.Contains(textureLowercase, "png") {
    imageLoader = loaders.NewImageLoader(loaders.PNG)
  } else if strings.Contains(textureLowercase, "jpg") || strings.Contains(textureLowercase, "jpeg") {
    imageLoader = loaders.NewImageLoader(loaders.JPEG)
  } else {
    return
  }

  imageLoader.LoadImage(textureName)
  nrgbaData, _ := imageLoader.ToNRGBA()
  component.Texture = opengl.NewTextureFromData(textureName, nrgbaData, false)
}


func (component *PrimitiveRenderer) SetShape(shape Shape2D) {
	switch shape {
	case SHAPE_RECT:
		component.setShapeSquare()

	case SHAPE_CIRCLE:
		component.setShapeCircle()
	default:
		logging.DebugLog("shape not implemented")
	}

}

func (component *PrimitiveRenderer) drawColorField() {

	colorTempArr := [4]float32{
		component.Color.X(),
		component.Color.Y(),
		component.Color.Z(),
		component.Color.W(),
	}

	imgui.PushItemWidth(200)
	imgui.ColorEdit4("Color", &colorTempArr)
	// imgui.ColorPicker4V("Color", &colorTempArr, 0)
	// imgui.ColorPicker4V("Color", &colorTempArr, imgui.ColorPickerFlagsPickerHueWheel)
	imgui.Spacing()

	component.Color.Vec4 = glm.Vec4{
		colorTempArr[0],
		colorTempArr[1],
		colorTempArr[2],
		colorTempArr[3],
	}
}

func (component *PrimitiveRenderer) handleShapeSelection(shape shape_list_item) {
	switch shape.shapeType {
	case SHAPE_RECT:
		component.setShapeSquare()
	case SHAPE_CIRCLE:
		component.setShapeCircle()
	default:
		component.Shape = nil
		logging.DebugLogError("No pipeline set for this shape")
	}
}

func (component *PrimitiveRenderer) drawControls() {
	// Find out what shape we have
	switch component.Shape.(type) {
	case *shape.Rect:
		component.drawRectControls()
	case *shape.Circle:
		component.drawCircleControls()
	default:
		// No controls available for shape
	}
}

func (component *PrimitiveRenderer) drawRectControls() {

	imgui.BeginChildV("shape_handler", util.ImguiGenVec2(0, imgui.TextLineHeightWithSpacing()*3.5), true, 0)

	// Get the rect (should work because we're only here if the casting has already worked)
	rect := component.Shape.(*shape.Rect)

	imgui.PushID("rect_controls")

	imgui.ColumnsV(2, "width_controls", false)
	imgui.SetColumnWidth(0, 100)
	imgui.SetColumnWidth(1, imgui.WindowWidth()-100)

	imgui.Text("Width")
	imgui.NextColumn()
	imgui.DragFloatV("##width", &rect.Width, 1.0, 0.0, 0.0, "%.3f", 1.0)

	imgui.Columns()

	imgui.ColumnsV(2, "height_controls", false)
	imgui.SetColumnWidth(0, 100)
	imgui.SetColumnWidth(1, imgui.WindowWidth()-100)

	imgui.Text("Height")
	imgui.NextColumn()
	imgui.DragFloatV("##height", &rect.Height, 1.0, 0.0, 0.0, "%.3f", 1.0)

	imgui.Columns()

	imgui.PopID()

	imgui.EndChild()
}

func (component *PrimitiveRenderer) drawCircleControls() {

	imgui.BeginChildV("shape_handler", util.ImguiGenVec2(0, imgui.TextLineHeightWithSpacing()*3.5), true, 0)

	// Get the circle (should work because we're only here if the casting has already worked)
	circle := component.Shape.(*shape.Circle)

	imgui.PushID("rect_controls")

	imgui.ColumnsV(2, "radius_controls", false)
	imgui.SetColumnWidth(0, 100)
	imgui.SetColumnWidth(1, imgui.WindowWidth()-100)

	imgui.Text("Radius")
	imgui.NextColumn()
	imgui.DragFloatV("##radius", &circle.Radius, 1.0, 0.0, 0.0, "%.3f", 1.0)

	imgui.Columns()

	imgui.PopID()

	imgui.EndChild()
}

// Sets the current shape to be a square
func (component *PrimitiveRenderer) setShapeSquare() {
	rect := &shape.Rect{
		Width:  100,
		Height: 100,
	}
	component.selectedIndex = 1
	// This works because Rect implements the IShape interface
	// Convert rect to ishape
	component.Shape = rect
}

func (component *PrimitiveRenderer) setShapeCircle() {
	circle := &shape.Circle{
		Radius: 50,
	}
	component.selectedIndex = 2
	component.Shape = circle
}

func (component *PrimitiveRenderer) GetComponentName() string {
	return "Primitive Renderer"
}

func (component *PrimitiveRenderer) MarkRender() {}
