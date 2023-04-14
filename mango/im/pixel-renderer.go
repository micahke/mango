package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util/color"
)



type PixelRenderer struct {

  vao *opengl.VertexArray
  vbo *opengl.VertexBuffer
  layout *opengl.VertexBufferLayout

  shader *opengl.Shader

  points []float32

  curIndex int

  modelMatrix glm.Vec4

  // No need for index buffer here
  // Note: If stuff doesn't work, just make sure that IBOs are unbound 
    
}

const (
  p_MAX_POINTS int = 5000
  p_VERTEX_SIZE int = 2  // only x y
)

func InitPixelRenderer() *PixelRenderer {

  renderer := new(PixelRenderer)

  renderer.points = make([]float32, p_MAX_POINTS * p_VERTEX_SIZE)

  renderer.vao = opengl.NewVertexArray()
  renderer.vbo = opengl.NewVertexBuffer(renderer.points)
  renderer.layout = opengl.NewVertexBufferLayout()
  renderer.layout.Pushf(2)
  renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

  renderer.shader = opengl.NewShader("PixelVertex.glsl", "PixelFragment.glsl")

  return renderer

}


// TODO: think about exceeding max points

// TODO: think about exceeding max points
func (renderer *PixelRenderer) DrawPixels(pixels []float32, size float32, projectionMatrix, viewMatrix glm.Mat4) {

  renderer.shader.Bind()

	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", glm.Ident4())
  clr := color.WHITE.Vec4
  renderer.shader.SetUniform4f("uColor", clr[0], clr[1], clr[2], 1.0)
  renderer.shader.SetUniform1f("uPixelSize", size)

  var leftoverVertices int = len(pixels) % (p_MAX_POINTS * 2)
  var iterator int

  for iterator < len(pixels) / (p_MAX_POINTS * 2) {
    logging.DebugLog("BATCH: ", iterator)
    copy(renderer.points, pixels[iterator*p_MAX_POINTS*2:(iterator+1)*p_MAX_POINTS*2])

    renderer.vbo.Bind()
    gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4*len(renderer.points), gl.Ptr(&renderer.points[0]))

    renderer.vao.Bind()
    gl.DrawArrays(gl.POINTS, 0, int32(len(renderer.points) / 2))
    iterator++
  }

  if leftoverVertices > 0 {
    copy(renderer.points, pixels[iterator*p_MAX_POINTS*2:])
    
    renderer.vbo.Bind()
    gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4*leftoverVertices, gl.Ptr(&renderer.points[0]))

    renderer.vao.Bind()
    gl.DrawArrays(gl.POINTS, 0, int32(leftoverVertices / 2))
  }
  
}

