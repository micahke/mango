package components

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/components/shape"
)

type Shape2DComponent struct {

  Name string
  Shape shape.IShape

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
  imgui.Text(component.Shape.GetShapeName())
}

func (component *Shape2DComponent) SetShape(shape shape.IShape) {
  component.Shape = shape
}
