package ecs


// "github.com/micahke/mango/components"
// "github.com/micahke/mango/logging"


type Component interface {

  // This function initializes the component
  Init()

  // This function updated every frame
  Update()


}

type RenderComponent interface {
  MarkRender()
}

type NamedComponent interface {
  GetComponentName() string
}


type UIEditableComponent interface {

  RenderControlPanel()

}


// Checks whetehr a given component is renerable or not
func isRenderableComponent(component interface{}) bool {
  
  // For testing, we're setting the sample component as something that
  // that can be rendered
  // _, sampleError := component.(*components.SampleComponent)
  // if !sampleError {
  //
  //   return true
  // }
  //
  //
  //
  // _, primitiveError := component.(*components.PrimitiveRenderer)
  // if !primitiveError {
  //   return true
  // }

  _, ok := component.(RenderComponent)
  if ok  {
    return true
  }



  return false

}
