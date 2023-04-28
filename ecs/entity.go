package ecs

import (
	"github.com/micahke/mango/components"
)




type Entity struct {
  
  id string
	components []interface{}

}



func (entity *Entity) Update() {


  for _, component := range(entity.components) {

    // Check to see if this is a valid component
    if _, ok := component.(Component); ok {
      
      cmp := component.(Component)
      cmp.Update()

    }

  }

}


func (entity *Entity) AddComponent(component interface{}) {

  entity.components = append(entity.components, component)

}


// Get the transform component of the entity
func (entity *Entity) Tranform() *components.TransformComponent {

  for _, component := range(entity.components) {

    // Check to see if the component is a TransformComponent
    if transform, ok := component.(components.TransformComponent); ok {
      return &transform
    }
  }

  return nil

}



// Creates a default transform component
func (entity *Entity) addTranformComponent() {

  entity.AddComponent(components.TransformComponent{
    X: 0,
    Y: 0,
  })

}


