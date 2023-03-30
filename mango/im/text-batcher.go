package im

import (

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/opengl"
)

type TextBatcher struct {
	vertices []float32

	vao    *opengl.VertexArray
	vbo    *opengl.VertexBuffer
	layout *opengl.VertexBufferLayout
	ibo    *opengl.IndexBuffer

  indeces []uint32

	shader  *opengl.Shader

	BATCH_SIZE   int
	num_vertices int
}

func InitTextBatcher() *TextBatcher {
	batcher := new(TextBatcher)

	batcher.vao = opengl.NewVertexArray()
	batcher.vbo = opengl.NewVertexBuffer(sprite_positions)
	batcher.layout = opengl.NewVertexBufferLayout()
	batcher.layout.Pushf(2) // Push for vertex position
	batcher.layout.Pushf(2) // Push for fragment position

  batcher.vao.AddBuffer(*batcher.vbo, *batcher.layout)

  batcher.ibo = opengl.NewIndexBuffer(quad_indeces)

	batcher.BATCH_SIZE = 100
	batcher.num_vertices = 0

	batcher.shader = opengl.NewShader("TextVertex.glsl", "TextFragment.glsl")

	return batcher

}

func (batch *TextBatcher) InitBatch() {

	batch.vertices = []float32{}
  batch.indeces = []uint32{}

}

func (batch *TextBatcher) AddCharacter(char string, x, y float32) {

	charInfo := _atlas[char]

	fx := x
	fy := y
	fEndX := float32(charInfo.width) + fx
	fEndY := float32(FONT_SIZE) + fy

	quad := []float32{

		fx, fy, charInfo.texturePositions[0], charInfo.texturePositions[1],
		fx, fEndY, charInfo.texturePositions[0], charInfo.texturePositions[3],
		fEndX, fEndY, charInfo.texturePositions[2], charInfo.texturePositions[3],
		fEndX, fy, charInfo.texturePositions[2], charInfo.texturePositions[1],
	}

  fOff := uint32(batch.num_vertices)

  indeces := []uint32{
    0 + fOff, 1 + fOff, 2 + fOff,
    2 + fOff, 3 + fOff, 0 + fOff,
  }

	batch.vertices = append(batch.vertices, quad...)
  batch.indeces = append(batch.indeces, indeces...)

	batch.num_vertices += 4


}

func (batch *TextBatcher) FlushBatch(projectionMatrix, viewMatrix glm.Mat4) {

  texture := getTexture("BitmapFontDebug.png", false)
	texture.Bind(1)

  vao := opengl.NewVertexArray()
  vbo := opengl.NewVertexBuffer(batch.vertices)
  layout := opengl.NewVertexBufferLayout()
  layout.Pushf(2)
  layout.Pushf(2)
  vao.AddBuffer(*vbo, *layout)

  ibo := opengl.NewIndexBuffer(batch.indeces)

	batch.shader.Bind()
	batch.shader.SetUniformMat4f("projection", projectionMatrix)
	batch.shader.SetUniformMat4f("view", viewMatrix)
	batch.shader.SetUniformMat4f("model", glm.Ident4())
  batch.shader.SetUniform1i("uTexture", 1)

  // logging.DebugLog("Vertices: ", batch.num_vertices)
  logging.DebugLog("Indeces: ", batch.indeces)

  ibo.Bind()
	vao.Bind()


	gl.DrawElements(gl.TRIANGLES, int32(len(batch.indeces)), gl.UNSIGNED_INT, nil)

	batch.num_vertices = 0

}
