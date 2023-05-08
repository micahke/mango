package ecs

import "github.com/micahke/mango/components"


type Component interface {

  // This function initializes the component
  Init()

  // This function updated every frame
  Update()


}

type NamedComponent interface {
  GetComponentName() string
}


type UIEditableComponent interface {

  RenderControlPanel()

}


// Checks whetehr a given component is renerable or not
func isRenderableComponent(component Component) bool {
  // For testing, we're setting the sample component as something that
  // that can be rendered
  _, sampleError := component.(*components.SampleComponent)
  if !sampleError {
    return true
  }


  _, primitiveError := component.(*components.PrimitiveRenderer)
  if !primitiveError {
    return true
  }

  return false

}
