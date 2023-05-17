package components

import (
	"github.com/AllenDang/imgui-go"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/util/color"
)


type PrimitiveRenderer struct {

  Color color.Color
  Shape *shape.IShape

  selectedIndex int
  
}


type Shape2D int
const (
	SHAPE_NONE    Shape2D = -1
	SHAPE_RECT    Shape2D = 0
	SHAPE_ELLIPSE Shape2D = 1
	SHAPE_LINE    Shape2D = 2
)

type shape_list_item struct {
  name string
  shapeType Shape2D 
}

var shapeList []shape_list_item =  []shape_list_item{
  {name: "NONE", shapeType: -1},
  {name: "RECT", shapeType:0},
  {name: "ELLIPSE", shapeType:1},
  {name: "LINE", shapeType:2},
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
    for index, shape := range(shapeList) {

      if imgui.Selectable(shape.name) {
        component.selectedIndex = index
      }

    }
  imgui.EndCombo()
  }

  imgui.Spacing()


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

  temp := [2]float32{2, 3}

  imgui.PushID("rect_controls")

  imgui.ColumnsV(2, "width_controls", false) 
  imgui.SetColumnWidth(0, 100)
  imgui.SetColumnWidth(1, imgui.WindowWidth() - 100)

  imgui.Text("Width")
  imgui.NextColumn()
  imgui.DragFloatV("##width", &temp[0], 1.0, 0.0, 0.0, "%.3f", 1.0)

  imgui.Columns()

  imgui.ColumnsV(2, "height_controls", false) 
  imgui.SetColumnWidth(0, 100)
  imgui.SetColumnWidth(1, imgui.WindowWidth() - 100)

  imgui.Text("Height")
  imgui.NextColumn()
  imgui.DragFloatV("##height", &temp[1], 1.0, 0.0, 0.0, "%.3f", 1.0)

  imgui.Columns()

  imgui.PopID()


} 


func (component *PrimitiveRenderer) GetComponentName() string {
  return "Primitive Renderer"
}
