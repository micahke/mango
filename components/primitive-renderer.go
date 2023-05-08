package components

import (
	"github.com/AllenDang/imgui-go"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/util/color"
)


type PrimitiveRenderer struct {

  Color color.Color
  
}


func (component *PrimitiveRenderer) Init() {
  component.Color = color.WHITE

}



func (component *PrimitiveRenderer) Update() {

}


func (component *PrimitiveRenderer) RenderControlPanel() {

  imgui.Spacing()

  colorTempArr := [4]float32{
    component.Color.X(),
    component.Color.Y(),
    component.Color.Z(),
    component.Color.W(),
  }

  imgui.PushItemWidth(200)
  imgui.ColorPicker4V("Color", &colorTempArr, imgui.ColorPickerFlagsPickerHueWheel)
  imgui.Spacing()

  component.Color.Vec4 = glm.Vec4{
    colorTempArr[0],
    colorTempArr[1],
    colorTempArr[2],
    colorTempArr[3],
  }


} 


func (component *PrimitiveRenderer) GetComponentName() string {
  return "Primitive Renderer"
}
