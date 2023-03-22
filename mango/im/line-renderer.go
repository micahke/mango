package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util"
)

type LineRenderer struct {
	vao    *opengl.VertexArray
	vbo    *opengl.VertexBuffer
	layout *opengl.VertexBufferLayout

	shader      *opengl.Shader
	modelMatrix glm.Mat4
}

var points []float32 = []float32{
	0.0, 0.0,
	0.0, 0.0,
}

func InitLineRenderer() *LineRenderer {

	renderer := new(LineRenderer)

	renderer.vao = opengl.NewVertexArray()
	renderer.vbo = opengl.NewVertexBuffer(points)
	renderer.layout = opengl.NewVertexBufferLayout()
	renderer.layout.Pushf(2)

	renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

	// renderer.shader = opengl.NewShaderG("LineVertex.glsl", "LineFragment.glsl", "LineGeometry.glsl")
	renderer.shader = opengl.NewShader("LineVertex.glsl", "LineFragment.glsl")

	renderer.modelMatrix = glm.Ident4()

	return renderer

}

func (renderer *LineRenderer) RenderLine(p1, p2 glm.Vec2, color util.Color, thickness float32, projectionMatrix, viewMatrix glm.Mat4) {

	renderer.vbo.SetData([]float32{
		p1.X(), p1.Y(),
		p2.X(), p2.Y(),
	})

	renderer.shader.Bind()
	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", renderer.modelMatrix)

	tk := thickness / float32(1300)

	renderer.shader.SetUniform1f("thickness", tk)

	renderer.shader.SetUniform4f("uColor", color.Vec4[0], color.Vec4[1], color.Vec4[2], color.Vec4[3])

	renderer.vao.Bind()

	gl.DrawArrays(gl.LINES, 0, 2)

}
