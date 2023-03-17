package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util"
)

type Quad struct {
	x      float32
	y      float32
	width  float32
	height float32

	color glm.Vec4
}

type QuadRenderer struct {
	vao    *opengl.VertexArray
	vbo    *opengl.VertexBuffer
	layout *opengl.VertexBufferLayout

	ibo    *opengl.IndexBuffer
	shader *opengl.Shader

	modelMatrix glm.Mat4
}

var quad_positions []float32 = []float32{
	0.0, 0.0,
	0.0, 1.0,
	1.0, 1.0,
	1.0, 0.0,
}

var quad_indeces []uint32 = []uint32{
	0, 1, 2,
	2, 3, 0,
}

func InitQuadRenderer() *QuadRenderer {
	renderer := new(QuadRenderer)

	renderer.vao = opengl.NewVertexArray()
	renderer.vbo = opengl.NewVertexBuffer(quad_positions)
	renderer.layout = opengl.NewVertexBufferLayout()
	renderer.layout.Pushf(2)
	renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

	renderer.ibo = opengl.NewIndexBuffer(quad_indeces)
	renderer.shader = opengl.NewShader("QuadVertex.glsl", "QuadFragment.glsl")

	renderer.modelMatrix = glm.Ident4()

	return renderer
}

func (renderer *QuadRenderer) RenderQuad(x, y, width, height float32, color util.Color, projectionMatrix, viewMatrix glm.Mat4) {

	translation := glm.Translate3D(x, y, 0)
	scale := glm.Scale3D(width, height, 1.0)
	model := translation.Mul4(scale).Mul4(renderer.modelMatrix)

	renderer.shader.Bind()
	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", model)

	renderer.shader.SetUniform4f("uColor", color.Vec4[0], color.Vec4[1], color.Vec4[2], color.Vec4[3])

	renderer.vao.Bind()
	renderer.ibo.Bind()

	gl.DrawElements(gl.TRIANGLES, int32(renderer.ibo.GetCount()), gl.UNSIGNED_INT, nil)

}
