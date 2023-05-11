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
}

// Initialize the renderer
func (renderer *TediousRenderer) Init(windowWidth, windowHeight int) {

	// Initialize the renderer with a projection matrix and a view matrix
	renderer.projectionMatrix = glm.Ortho(0, float32(windowWidth), 0, float32(windowHeight), -1.0, 1.0)
	renderer.viewMatrix = glm.Ident4()

	// Initialize shaders
	renderer.quadShader = opengl.NewShader("RMQuadVertex.glsl", "RMQuadFragment.glsl")

	fmt.Println("Valid shader found")

	// Initialize vertex arrays
	renderer.quadVAO = opengl.NewVertexArray()
	renderer.quadVBO = opengl.NewVertexBuffer(quadVertices)
	renderer.quadIBO = opengl.NewIndexBuffer(quadIndeces)
	quadLayout := opengl.NewVertexBufferLayout()
	quadLayout.Pushf(2)
	quadLayout.Pushf(4)
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

// Handles the drawing of a quad
func (renderer *TediousRenderer) drawQuad(transform *components.TransformComponent, quad *shape.Rect, color color.Color) {

	if !renderer.initialized {
		logging.DebugLogError("Renderer not initialized")
		return
	}

	// quadVerts := generateQuadVertices(color)

	translation := glm.Translate3D(transform.Position.X, transform.Position.Y, 0)
	scale := glm.Scale3D(quad.Width, quad.Height, 1.0)

	modelMatrix := glm.Ident4()
	model := translation.Mul4(scale).Mul4(modelMatrix)

	renderer.quadShader.Bind()
	renderer.quadShader.SetUniformMat4f("u_model", model)
	renderer.quadShader.SetUniformMat4f("u_view", renderer.viewMatrix)
	renderer.quadShader.SetUniformMat4f("u_projection", renderer.projectionMatrix)

	// renderer.quadVBO.Bind()
	// gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(quadVerts)*4, gl.Ptr(&quadVerts[0]))

	logging.DebugLog(model)

	renderer.quadVAO.Bind()
	renderer.quadIBO.Bind()

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

}

func (renderer *TediousRenderer) Render() {
}
