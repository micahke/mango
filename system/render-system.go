package system

import (
	"reflect"

	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/renderer"
)

type RenderSystem struct {
	Entities        *[]*ecs.Entity
	tediousRenderer *renderer.TediousRenderer
}

type mango_renderer int

const (
	RENDERER_NONE    mango_renderer = -1
	TEDIOUS_RENDERER mango_renderer = 0
)

func (renderSystem *RenderSystem) Init() {
	renderSystem.tediousRenderer = &renderer.TediousRenderer{}
	renderSystem.tediousRenderer.Init(1300, 800)
}

// This is about to be worst render system of all time
func (renderSystem *RenderSystem) Tick() {
	// Loop through the entities and detect which of them are renderable
  renderSystem.NewFrame()
	for _, entity := range *renderSystem.Entities {
		if entity.Renderable {
			// Get the appropriate renderer and the right renderable component
			renderer, renderComponent := renderSystem.determineRenderer(entity)
			renderSystem.submitToRenderer(renderer, entity, renderComponent)
		}
	}
  renderSystem.EndFrame()
}

func (renderSystem *RenderSystem) EndFrame() {
  renderSystem.tediousRenderer.FlushFrame()
}

func (renderSystem *RenderSystem) NewFrame() {
  renderSystem.tediousRenderer.NewFrame()
}

  
// Detmines which renderer we send things to
func (renderSystem *RenderSystem) determineRenderer(entity *ecs.Entity) (mango_renderer, renderer.RenderableComponent) {

	if entity.HasComponent(reflect.TypeOf(&components.PrimitiveRenderer{})) {
		return TEDIOUS_RENDERER, renderer.PRIMITIVE_RENDERER
	}

	return RENDERER_NONE, renderer.NO_RENDER

}

// Submit the entity to the appropriate renderer
// TODO: rework this to more efficiently obtain the correct components to send over
func (renderSystem *RenderSystem) submitToRenderer(renderer mango_renderer, entity *ecs.Entity, component renderer.RenderableComponent) {
	if renderer == RENDERER_NONE {
		return
	}

	if renderer == TEDIOUS_RENDERER {
		// Can be pulled out to another thread
		renderSystem.tediousRenderer.Submit(entity, component)
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
