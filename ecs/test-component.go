package ecs

import (
	"reflect"

	"github.com/micahke/mango/components"
	"github.com/micahke/mango/logging"
)

// A sample component that we can use around the engine for various tasks
type TestComponent struct {
  *Entity
  pr *components.PrimitiveRenderer
}



func (component *TestComponent) Init() {
  pr, err := component.Entity.GetComponent(reflect.TypeOf(&components.PrimitiveRenderer{}))
  if err != nil {
    logging.DebugLog("No entity found")
  }
  renderer, ok := pr.(*components.PrimitiveRenderer)
  if !ok {
    logging.DebugLog("Couldn't convert")
  }
  component.pr = renderer
}


func (component *TestComponent) Update() {

  component.Entity.Tranform().Rotation.Z += 0.5

}


func (component *TestComponent) GetComponentName() string {
  return "Test Component"
}
