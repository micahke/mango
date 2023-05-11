package renderer

import (
	"reflect"

	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/util/color"
)

// Basically a bad renderer that renders each entity with a draw call
type TediousRenderer struct {
  projectionMatrix glm.Mat4
  viewMatrix glm.Mat4
}



// Initialize the renderer
func (renderer *TediousRenderer) Init(windowWidth, windowHeight int) {

  // Initialize the renderer with a projection matrix and a view matrix
  renderer.projectionMatrix = glm.Ortho(0, float32(windowWidth), 0, float32(windowHeight), -1.0, 1.0)
  renderer.viewMatrix = glm.Ident4()

}


// Handles the entity and starts to work on its available data
func (renderer *TediousRenderer) Submit(entity *ecs.Entity, renderableComponent RenderableComponent) {

  switch renderableComponent {
  case PRIMITIVE_RENDERER:
    renderer.handlePrimitiveRenderer(entity)
  }

}

func (renderer *TediousRenderer) handlePrimitiveRenderer(entity *ecs.Entity) {
  // stopwatch := &util.Stopwatch{}
  // stopwatch.Start()
  // Get the Shape2D component
  primitveRenderer, err := entity.GetComponent(reflect.TypeOf(&components.PrimitiveRenderer{}))
  color := primitveRenderer.(*components.PrimitiveRenderer).Color
  transform := entity.Tranform()
  if err != nil {
    logging.DebugLogError("No primitive component found for entity")
  }
  rawShape, err := entity.GetComponent(reflect.TypeOf(&components.Shape2DComponent{}))
  if err != nil {
    logging.DebugLogError("Entity does not have a Shape2D component")
    return
  }

  shape2D := rawShape.(*components.Shape2DComponent)

  correctShape := shape2D.Determine()
  if correctShape == components.SHAPE_RECT {
    quad := shape2D.Shape.(*shape.Rect)
    renderer.drawQuad(transform, quad, color)
  }
  // end := stopwatch.Stop()
  // logging.DebugLog("Filtering render data took", end, "MS")
}

func (renderer *TediousRenderer) drawQuad(tranform *components.TransformComponent, quad *shape.Rect, color color.Color) {
  // Handles the drawing of a quad
  logging.DebugLog("Render data:", tranform, quad, color) 
}


func (renderer *TediousRenderer) Render() {
}
