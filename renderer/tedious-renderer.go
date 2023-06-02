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
		renderer.drawQuad(transform, quad, color, renderer.quadShader, prComponent.Texture)
	case *shape.Circle:
		circle := prComponent.Shape.(*shape.Circle)
		renderer.drawCircle(transform, circle, color, renderer.circleShader, prComponent.Texture)
	default:
	}
}

// Handles the drawing of a quad
func (renderer *TediousRenderer) drawQuad(transform *components.TransformComponent, quad *shape.Rect, color color.Color, shader *opengl.Shader, texture *opengl.Texture) {

	if !renderer.initialized {
		logging.DebugLogError("Renderer not initialized")
		return
	}

	// renderer.sampleTexture.Bind(0)

	quadVerts := generateQuadVertices(color)

	// Move the quad to the origin, rotate, then move it back to its position.
	translationToOrigin := glm.Translate3D(-quad.Width/2, -quad.Height/2, 0)
	rotX := glm.HomogRotate3D(glm.DegToRad(transform.Rotation.X), glm.Vec3{1, 0, 0})
	rotY := glm.HomogRotate3D(glm.DegToRad(transform.Rotation.Y), glm.Vec3{0, 1, 0})
	rotZ := glm.HomogRotate3D(glm.DegToRad(transform.Rotation.Z), glm.Vec3{0, 0, 1})

	rotation := rotX.Mul4(rotY).Mul4(rotZ)

	translationToPosition := glm.Translate3D(transform.Position.X+quad.Width/2, transform.Position.Y+quad.Height/2, transform.Position.Z)

	scale := glm.Scale3D(quad.Width, quad.Height, 1.0)

	// Create the model matrix
	modelMatrix := glm.Ident4()
	model := translationToPosition.Mul4(rotation).Mul4(translationToOrigin).Mul4(scale).Mul4(modelMatrix)

	shader.Bind()
	shader.SetUniformMat4f("u_model", model)
	shader.SetUniformMat4f("u_view", renderer.viewMatrix)
	shader.SetUniformMat4f("u_projection", renderer.projectionMatrix)
	// shader.SetUniform1i("uTexture", 0)
  renderer.decideBindTexture(shader, texture)

	renderer.quadVBO.Bind()
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(quadVerts)*4, gl.Ptr(&quadVerts[0]))

	// logging.DebugLog(quadVerts)

	renderer.quadVAO.Bind()
	renderer.quadIBO.Bind()

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

}

// Decides whether or not to bind a texture for a given entity based on
// whether it has a valid texture in it's primitive renderer
func (renderer *TediousRenderer) decideBindTexture(shader *opengl.Shader, texture *opengl.Texture) {
	if texture == nil {
		shader.SetUniform1i("isTextured", 1)
		return
	}
  texture.Bind(0)
	shader.SetUniform1i("isTextured", 0)
}

// Circle is just a rect with a bunch of color in a pattern just do that
// Will need a new shader
func (renderer *TediousRenderer) drawCircle(transform *components.TransformComponent, circle *shape.Circle, color color.Color, shader *opengl.Shader, texture *opengl.Texture) {
	// Generate a quad with width and height of 2 * radius
	quad := &shape.Rect{
		Width:  circle.Radius * 2,
		Height: circle.Radius * 2,
	}

	renderer.drawQuad(transform, quad, color, shader, texture)

}
