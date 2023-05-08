package system

import (
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
)


type RenderSystem struct {
  Entities  *[]*ecs.Entity
}


func (renderSystem *RenderSystem) Init() {}

// This is about to be worst render system of all time
func (renderSystem *RenderSystem) Tick() {
  // Loop through the entities and detect which of them are renderable
  for _, entity := range(*renderSystem.Entities) {
    if entity.Renderable {
      // Handle the rendering of the system
    } 
  }
}

// This check whether or a given component should be iterated upon
// by the render system. If so, add the component to a cache that
// gets handled when we do the drawing
// Don't use this function for now
func TrySubmitRenderQueue(component ecs.Component) {
  
  sampleComponent, err := component.(*components.SampleComponent)
  if err {
    logging.DebugLogError("Error submitting SampleComponent to render queue")
  }
  
  logging.DebugLog("Got component:", sampleComponent.GetComponentName())

}


// Placeholder
// I want to replace this using a more efficient ECS system that decouples everything
func (renderSystem *RenderSystem) TrySubmitRenderQueue(component *ecs.Component) {
  // renderSystem.renderableComponents = append(renderSystem.renderableComponents, component)
}
