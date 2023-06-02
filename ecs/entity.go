package ecs

import (
	"fmt"
	"reflect"

	"github.com/micahke/mango/components"
)


type Entity struct {
  
  Name string
	Components []interface{}

  componentBuffer []interface{}

  // Represents whether or not this entity gets rendered
  Renderable bool

}



func (entity *Entity) Update() {

  // activate new components
  for _, component := range(entity.componentBuffer) {
    if _, ok := component.(Component); ok {
      cmp := component.(Component)

      entity.checkAndSetEmbeddedValue(cmp)

      // Initializing component
      cmp.Init()

      entity.Components = append(entity.Components, cmp)

    }
  }

  entity.componentBuffer = []interface{}{}


  for _, component := range(entity.Components) {

    // Check to see if this is a valid component
    if _, ok := component.(Component); ok {
      
      cmp := component.(Component)

      // If the component is renderable, mark it as such
      if isRenderableComponent(component) {
        entity.Renderable = true
      }

      cmp.Update()

    }

  }

}

func (entity *Entity) checkAndSetEmbeddedValue(component Component) error {
  value := reflect.ValueOf(component)
  if value.Kind() == reflect.Ptr {
    value = value.Elem()
  }
  
  if value.Kind() != reflect.Struct {
    return fmt.Errorf("Underlying type not a struct")
  }

  for i := 0; i < value.NumField(); i++ {
    field := value.Field(i)
    if field.Type() == reflect.TypeOf(&Entity{}) {
      field.Set(reflect.ValueOf(entity))
    }
  }
  return nil
}


func (entity *Entity) AddComponent(component interface{}) {

  entity.componentBuffer = append(entity.componentBuffer, component)

}

func (entity *Entity) RemoveComponentByType(t reflect.Type) {
  for index, component := range(entity.Components) {
    if reflect.TypeOf(component) == t {
      entity.Components = append(entity.Components[:index], entity.Components[index+1:]...)
    }
  }
}

func (entity *Entity) RemoveComponent(component interface{}) {
  for index, c := range(entity.Components) {
    if c == component {
      entity.Components = append(entity.Components[:index], entity.Components[index+1:]...)
    }
  }
}

// Get the transform component of the entity
func (entity *Entity) Tranform() *components.TransformComponent {

  for _, component := range(entity.Components) {

    // Check to see if the component is a TransformComponent
    if transform, ok := component.(*components.TransformComponent); ok {
      return transform
    }
  }

  for _, component := range(entity.componentBuffer) {

    // Check to see if the component is a TransformComponent
    if transform, ok := component.(*components.TransformComponent); ok {
      return transform
    }
  }

  return nil

}



// Creates a default transform component
func (entity *Entity) addTranformComponent() {

  entity.AddComponent(&components.TransformComponent{
    Name: "Transform",
  })

}


// Checks whether the enitity has a given componetn
func (entity *Entity) HasComponent(t reflect.Type) bool {
  for _, component := range(entity.Components) {
    if reflect.TypeOf(component) == t {
      return true
    }
  }
  return false
}


// Gets a component if it exists
func (entity *Entity) GetComponent(t reflect.Type) (interface{}, error) {
  for _, component := range(entity.Components) {
    if reflect.TypeOf(component) == t {
      return component, nil
    }
  }
  return nil, fmt.Errorf("No component found")
}
