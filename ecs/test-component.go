package ecs

import "github.com/micahke/mango/logging"

// A sample component that we can use around the engine for various tasks
type TestComponent struct {
  *Entity
}



func (component *TestComponent) Init() {
  logging.Log(component.Entity)
}


func (component *TestComponent) Update() {}


func (component *TestComponent) GetComponentName() string {
  return "Test Component"
}
