package system

import (
	"reflect"

	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/renderer"
)


type RenderSystem struct {
  Entities  *[]*ecs.Entity
  tediousRenderer *renderer.TediousRenderer
}


func (renderSystem *RenderSystem) Init() {
  renderSystem.tediousRenderer = &renderer.TediousRenderer{}
}

// This is about to be worst render system of all time
func (renderSystem *RenderSystem) Tick() {
  // Loop through the entities and detect which of them are renderable
  for _, entity := range(*renderSystem.Entities) {
    if entity.Renderable {
      renderSystem.determineRenderer(entity)
    } 
  }
}

// Detmines which renderer we send things to
func (renderSystem *RenderSystem) determineRenderer(entity *ecs.Entity) {

  if entity.HasComponent(reflect.TypeOf(&components.PrimitiveRenderer{})) {
    logging.DebugLog("Would be sending to the active 2D renderer")
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



