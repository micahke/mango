package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util"
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
  p_MAX_POINTS int = 1000
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
func (renderer *PixelRenderer) DrawPixels(pixels []float32, size float32, projectionMatrix, viewMatrix glm.Mat4) {

  // copy(renderer.points, pixels)

  renderer.vbo.Bind()
  gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4*len(pixels), gl.Ptr(&pixels[0]))

  renderer.shader.Bind()

	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", glm.Ident4())
  clr := util.WHITE.Vec4
  renderer.shader.SetUniform4f("uColor", clr[0], clr[1], clr[2], 1.0)
  renderer.shader.SetUniform1f("uPixelSize", size)

  renderer.vao.Bind()

  gl.DrawArrays(gl.POINTS, 0, int32(len(pixels) / 2))


}
