package renderer

import (
	"fmt"
	"reflect"

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util/color"
)

// Basically a bad renderer that renders each entity with a draw call
type TediousRenderer struct {
	initialized bool

	projectionMatrix glm.Mat4
	viewMatrix       glm.Mat4

	quadShader *opengl.Shader
	quadVAO    *opengl.VertexArray
	quadVBO    *opengl.VertexBuffer
	quadIBO    *opengl.IndexBuffer

	circleShader *opengl.Shader
}

// Initialize the renderer
func (renderer *TediousRenderer) Init(windowWidth, windowHeight int) {

	// Initialize the renderer with a projection matrix and a view matrix
	renderer.projectionMatrix = glm.Ortho(0, float32(windowWidth), 0, float32(windowHeight), -1.0, 1.0)
	renderer.viewMatrix = glm.Ident4()

	// Initialize shaders
	renderer.quadShader = opengl.NewShader("RMQuadVertex.glsl", "RMQuadFragment.glsl")
	renderer.circleShader = opengl.NewShader("RMCircleVertex.glsl", "RMCircleFragment.glsl")

	fmt.Println("Valid shader found")

	// Initialize vertex arrays
	renderer.quadVAO = opengl.NewVertexArray()
	renderer.quadVBO = opengl.NewVertexBuffer(quadVertices)
	renderer.quadIBO = opengl.NewIndexBuffer(quadIndeces)
	quadLayout := opengl.NewVertexBufferLayout()
	quadLayout.Pushf(2)
	quadLayout.Pushf(4)
	quadLayout.Pushf(2)
	renderer.quadVAO.AddBuffer(*renderer.quadVBO, *quadLayout)

	renderer.initialized = true

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

  prComponent := primitveRenderer.(*components.PrimitiveRenderer)
  switch prComponent.Shape.(type) {
    case *shape.Rect:
      quad := prComponent.Shape.(*shape.Rect)
      renderer.drawQuad(transform, quad, color, renderer.quadShader)
  case *shape.Circle:
      circle := prComponent.Shape.(*shape.Circle)
      renderer.drawCircle(transform, circle, color, renderer.circleShader)
    default:
  }
}

// Handles the drawing of a quad
func (renderer *TediousRenderer) drawQuad(transform *components.TransformComponent, quad *shape.Rect, color color.Color, shader *opengl.Shader) {


	if !renderer.initialized {
		logging.DebugLogError("Renderer not initialized")
		return
	}

	quadVerts := generateQuadVertices(color)

	translation := glm.Translate3D(transform.Position.X, transform.Position.Y, transform.Position.Z)
	scale := glm.Scale3D(quad.Width, quad.Height, 1.0)

	modelMatrix := glm.Ident4()
	model := translation.Mul4(scale).Mul4(modelMatrix)

	shader.Bind()
	shader.SetUniformMat4f("u_model", model)
	shader.SetUniformMat4f("u_view", renderer.viewMatrix)
	shader.SetUniformMat4f("u_projection", renderer.projectionMatrix)

	renderer.quadVBO.Bind()
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(quadVerts)*4, gl.Ptr(&quadVerts[0]))

	// logging.DebugLog(quadVerts)

	renderer.quadVAO.Bind()
	renderer.quadIBO.Bind()

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

}


// Circle is just a rect with a bunch of color in a pattern just do that
// Will need a new shader
func (renderer *TediousRenderer) drawCircle(transform *components.TransformComponent, circle *shape.Circle, color color.Color, shader *opengl.Shader) {
  // Generate a quad with width and height of 2 * radius
  quad := &shape.Rect{
    Width: circle.Radius * 2,
    Height: circle.Radius * 2,
  }

  renderer.drawQuad(transform, quad, color, shader)

}


