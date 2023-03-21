package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util"
)


type LineRenderer struct {

  vao *opengl.VertexArray
  vbo *opengl.VertexBuffer
  layout *opengl.VertexBufferLayout

  shader *opengl.Shader
  modelMatrix glm.Mat4

}


var points []float32 = []float32{
  0.0, 0.0,
  1.0, 1.0,
}


func InitLineRenderer() *LineRenderer {

  renderer := new(LineRenderer)

  renderer.vao = opengl.NewVertexArray()
  renderer.vbo = opengl.NewVertexBuffer(points)
  renderer.layout = opengl.NewVertexBufferLayout()
  renderer.layout.Pushf(2)

  renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

  renderer.shader = opengl.NewShader("LineVertex.glsl", "LineFragment.glsl")
  
  renderer.modelMatrix = glm.Ident4()

  return renderer


}


func (renderer *LineRenderer) RenderLine(p1, p2 glm.Vec2, color util.Color, projectionMatrix, viewMatrix glm.Mat4) {

  renderer.vbo = opengl.NewVertexBuffer([]float32{
    p1.X(), p1.Y(),
    p2.X(), p2.Y(),
  })

  renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

	renderer.shader.Bind()
	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", renderer.modelMatrix)

  
	renderer.shader.SetUniform4f("uColor", color.Vec4[0], color.Vec4[1], color.Vec4[2], color.Vec4[3])

	renderer.vao.Bind()

  gl.DrawArrays(gl.LINES, 0, 2)

}
