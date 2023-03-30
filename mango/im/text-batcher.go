package im

import (
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/opengl"
)

type TextBatcher struct {
	vertices []float32
	indeces []uint32

	shader *opengl.Shader

	BATCH_SIZE   int
	num_vertices int
}

func InitTextBatcher() *TextBatcher {
	batcher := new(TextBatcher)

	batcher.BATCH_SIZE = 100
	batcher.num_vertices = 0

	batcher.shader = opengl.NewShader("TextVertex.glsl", "TextFragment.glsl")

	return batcher

}

func (batch *TextBatcher) InitBatch() {

	batch.vertices = []float32{}
	batch.indeces = []uint32{}

}

func (batch *TextBatcher) AddCharacter(char *FontAtlasItem, x, y float32) {

	fx := x
	fy := y
	fEndX := float32(char.width) + fx
	fEndY := float32(FONT_SIZE) + fy

	quad := []float32{

		fx, fy, char.texturePositions[0], char.texturePositions[1],
		fx, fEndY, char.texturePositions[0], char.texturePositions[3],
		fEndX, fEndY, char.texturePositions[2], char.texturePositions[3],
		fEndX, fy, char.texturePositions[2], char.texturePositions[1],
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

func (batch *TextBatcher) AddText(text string, x, y float32) {

	var offset int = 0

	for _, char := range text {

		character := _atlas[strings.ToUpper(string(char))]

		if character == nil {
			logging.DebugLogError("Could not find character in atlas:", character)
			return
		}

		batch.AddCharacter(character, x+float32(offset), y)

		offset += 24

	}

}

func (batch *TextBatcher) FlushBatch(projectionMatrix, viewMatrix glm.Mat4) {

  if (batch.num_vertices < 4) {
    return
  }

	texture := getTexture("BitmapFont.png", false)
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
