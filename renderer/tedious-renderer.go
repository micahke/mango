package renderer

import (
	"reflect"

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util"
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

  fbData *FrameBufferData
  framebufferID uint32
  framebufferTextureID uint32

  windowWidth int
  windowHeight int

  outputViewer *OutputViewer 
}

type FrameBufferData struct {
	Shader *opengl.Shader
	VAO    *opengl.VertexArray
	VBO    *opengl.VertexBuffer
	IBO    *opengl.IndexBuffer
}

// Initialize the renderer
func (renderer *TediousRenderer) Init(windowWidth, windowHeight int) {

	// Initialize the renderer with a projection matrix and a view matrix
	renderer.projectionMatrix = glm.Ortho(0, float32(windowWidth), 0, float32(windowHeight), -1.0, 1.0)
	renderer.viewMatrix = glm.Ident4()

  renderer.windowWidth = windowWidth
  renderer.windowHeight = windowHeight

  renderer.initFrameBuffer()
  renderer.outputViewer = NewOutputViewer(&renderer.framebufferTextureID)

  util.ImguiRegisterPanel("framebuffer", renderer.outputViewer)
  util.ImguiActivatePanel("framebuffer")

	// Initialize shaders
	renderer.quadShader = opengl.NewShader("RMQuadVertex.glsl", "RMQuadFragment.glsl")

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

  prComponent := primitveRenderer.(*components.PrimitiveRenderer)
  switch prComponent.Shape.(type) {
    case *shape.Rect:
      quad := prComponent.Shape.(*shape.Rect)
      renderer.drawQuad(transform, quad, color)
    default:
  }
}

// Handles the drawing of a quad
func (renderer *TediousRenderer) drawQuad(transform *components.TransformComponent, quad *shape.Rect, color color.Color) {


	if !renderer.initialized {
		logging.DebugLogError("Renderer not initialized")
		return
	}

	quadVerts := generateQuadVertices(color)

	translation := glm.Translate3D(transform.Position.X, transform.Position.Y, transform.Position.Z)
	scale := glm.Scale3D(quad.Width, quad.Height, 1.0)

	modelMatrix := glm.Ident4()
	model := translation.Mul4(scale).Mul4(modelMatrix)

	renderer.quadShader.Bind()
	renderer.quadShader.SetUniformMat4f("u_model", model)
	renderer.quadShader.SetUniformMat4f("u_view", renderer.viewMatrix)
	renderer.quadShader.SetUniformMat4f("u_projection", renderer.projectionMatrix)

	renderer.quadVBO.Bind()
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(quadVerts)*4, gl.Ptr(&quadVerts[0]))

	// logging.DebugLog(quadVerts)

	renderer.quadVAO.Bind()
	renderer.quadIBO.Bind()

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

}

func (renderer *TediousRenderer) initFrameBuffer() {
  // Framebuffer
  gl.GenFramebuffers(1, &renderer.framebufferID)
  gl.BindFramebuffer(gl.FRAMEBUFFER, renderer.framebufferID)

  // Framebuffer Texture
  gl.GenTextures(1, &renderer.framebufferTextureID)
  gl.BindTexture(gl.TEXTURE_2D, renderer.framebufferTextureID)

  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(renderer.windowWidth), int32(renderer.windowHeight), 0, gl.RGBA, gl.UNSIGNED_INT, nil)
  gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, renderer.framebufferTextureID, 0)

  status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)
  if status != gl.FRAMEBUFFER_COMPLETE {
    logging.DebugLogError("Error setting up framebuffer")
  }

  gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

  data := &FrameBufferData{}
  data.Shader = opengl.NewShader("FramebufferVertex.glsl", "FramebufferFragment.glsl")
  data.VAO = opengl.NewVertexArray()
  data.VBO = opengl.NewVertexBuffer(framebufferVertices)
  layout := opengl.NewVertexBufferLayout()
  layout.Pushf(2)
  layout.Pushf(2)
  data.VAO.AddBuffer(*data.VBO, *layout)

  data.IBO = opengl.NewIndexBuffer(quadIndeces)
  renderer.fbData = data
}

func (renderer *TediousRenderer) NewFrame() {
  gl.BindFramebuffer(gl.FRAMEBUFFER, renderer.framebufferID)
  // gl.Viewport(0, 0, int32(renderer.windowWidth), int32(renderer.windowHeight))
  gl.Enable(gl.DEPTH_TEST)

  gl.ClearColor(0.5, 0.5, 0.5, 1.0)
  gl.Clear(gl.COLOR_BUFFER_BIT)
}



// Find a better name for this function (this is experimental)
func (renderer *TediousRenderer) FlushFrame(){
  // Bind the default framebuffer
  gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
  gl.Disable(gl.DEPTH_TEST)
  // gl.Viewport(0, 0, int32(renderer.windowWidth), int32(renderer.windowHeight))

  gl.ClearColor(0.5, 0.5, 0.5, 1.0)
  gl.Clear(gl.COLOR_BUFFER_BIT)

  gl.ActiveTexture(gl.TEXTURE0)
  gl.BindTexture(gl.TEXTURE_2D, renderer.framebufferTextureID)
  // Next step, render full screen quad
  renderer.renderFullQuad()
}

func (renderer *TediousRenderer) renderFullQuad() {
  renderer.fbData.Shader.Bind()
  renderer.fbData.Shader.SetUniform1i("uTexture", 0)
  
  renderer.fbData.VAO.Bind()
  renderer.fbData.IBO.Bind()

  gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
}
