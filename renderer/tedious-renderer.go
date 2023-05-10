package renderer

import (
	"reflect"

	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
)

// Basically a bad renderer that renders each entity with a draw call
type TediousRenderer struct {
  projectionMatrix glm.Mat4
  viewMatrix glm.Mat4
}



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
  // Get the Shape2D component
  shape2D, err := entity.GetComponent(reflect.TypeOf(&components.Shape2DComponent{}))
  if err != nil {
    logging.DebugLogError("Entity does not have a Shape2D component")
    return
  }

  correctShape := shape2D.(*components.Shape2DComponent).Determine()
  if correctShape == components.SHAPE_RECT {
    logging.DebugLog("Got rectangle shape")
  }

}


func (renderer *TediousRenderer) Render() {
}
