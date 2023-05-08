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
  _, err := component.(*components.SampleComponent)
  if err {
    return false
  }
  return true
}
