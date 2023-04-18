package im

import (

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util/color"
)

type QuadBatcher struct {
	vao         *opengl.VertexArray
	vbo         *opengl.VertexBuffer
	layout      *opengl.VertexBufferLayout
	ibo         *opengl.IndexBuffer
	shader      *opengl.Shader
	modelMatrix glm.Mat4

  nVertices int

	vertices []float32
	indeces  []uint32
}

const (
	q_MAX_QUADS   int = 1000
	q_VERTEX_SIZE int = 6
)

type Quad struct {
	X, Y  float32
  Width, Height float32
	Color color.Color
}

func InitQuadBatcher() *QuadBatcher {

	batcher := new(QuadBatcher)

	batcher.vao = opengl.NewVertexArray()
  batcher.layout = opengl.NewVertexBufferLayout()

	batcher.setupVertexBuffer()
	batcher.generateIndexBuffer()

  batcher.shader = opengl.NewShader("QuadBatcherVertex.glsl", "QuadBatcherFragment.glsl")

	return batcher
}

func (batcher *QuadBatcher) setupVertexBuffer() {

	batcher.vertices = make([]float32, q_MAX_QUADS*q_VERTEX_SIZE*4)

  batcher.vbo = opengl.NewVertexBuffer(batcher.vertices)

	batcher.layout.Pushf(2) // X and Y positions
	batcher.layout.Pushf(4) // r, g, b, a values for color

	batcher.vao.AddBuffer(*batcher.vbo, *batcher.layout)
}

func (batcher *QuadBatcher) generateIndexBuffer() {

	batcher.indeces = make([]uint32, q_MAX_QUADS*6)

	for i := 0; i < q_MAX_QUADS; i++ {

		offset := uint32(i * 4)

		batcher.indeces[i*6+0] = 0 + offset
		batcher.indeces[i*6+1] = 1 + offset
		batcher.indeces[i*6+2] = 2 + offset

		batcher.indeces[i*6+3] = 2 + offset
		batcher.indeces[i*6+4] = 3 + offset
		batcher.indeces[i*6+5] = 0 + offset

	}

	batcher.ibo = opengl.NewIndexBuffer(batcher.indeces)
}


func (batcher *QuadBatcher) AddQuad(quad Quad, projectionMatrix, viewMatrix glm.Mat4) {

  if batcher.nVertices + 4 > q_MAX_QUADS * 4 {
    batcher.FlushBatch(projectionMatrix, viewMatrix)
  }

  // find the offset into the current buffer
  offset := batcher.nVertices * q_VERTEX_SIZE
  clr := quad.Color.Vec4
  
  batcher.vertices[offset] = quad.X
  batcher.vertices[offset+1] = quad.Y
  batcher.vertices[offset+2] = clr[0]
  batcher.vertices[offset+3] = clr[1]
  batcher.vertices[offset+4] = clr[2]
  batcher.vertices[offset+5] = clr[3]

  batcher.vertices[offset+6] = quad.X
  batcher.vertices[offset+7] = quad.Y + quad.Height
  batcher.vertices[offset+8] = clr[0]
  batcher.vertices[offset+9] = clr[1]
  batcher.vertices[offset+10] = clr[2]
  batcher.vertices[offset+11] = clr[3]

  batcher.vertices[offset+12] = quad.X + quad.Width
  batcher.vertices[offset+13] = quad.Y + quad.Height
  batcher.vertices[offset+14] = clr[0]
  batcher.vertices[offset+15] = clr[1]
  batcher.vertices[offset+16] = clr[2]
  batcher.vertices[offset+17] = clr[3]

  batcher.vertices[offset+18] = quad.X + quad.Width
  batcher.vertices[offset+19] = quad.Y
  batcher.vertices[offset+20] = clr[0]
  batcher.vertices[offset+21] = clr[1]
  batcher.vertices[offset+22] = clr[2]
  batcher.vertices[offset+23] = clr[3]


  batcher.nVertices += 4

}

// Handles the actual rendering of the batch
func (batcher *QuadBatcher) FlushBatch(projectionMatrix, viewMatrix glm.Mat4) {

  batcher.vbo.Bind()

  // Move this to the vertex buffer implemenation
  // size refers to memory size: 4 bytes
  gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4 * len(batcher.vertices), gl.Ptr(&batcher.vertices[0]))


  batcher.shader.Bind()
	batcher.shader.SetUniformMat4f("projection", projectionMatrix)
	batcher.shader.SetUniformMat4f("view", viewMatrix)
	batcher.shader.SetUniformMat4f("model", glm.Ident4())

	batcher.ibo.Bind()
	batcher.vao.Bind()

	indeces := (batcher.nVertices / 4) * 6

	gl.DrawElements(gl.TRIANGLES, int32(indeces), gl.UNSIGNED_INT, nil)

  batcher.nVertices = 0

}
