package im

import (
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util/loaders"
)

type TextBatcher struct {
	vao    *opengl.VertexArray
	vbo    *opengl.VertexBuffer
	layout *opengl.VertexBufferLayout

	ibo *opengl.IndexBuffer

	vertices []float32
	indeces  []uint32

	shader  *opengl.Shader
	texture *opengl.Texture

	num_vertices int
}

const (
	t_BATCH_SIZE  int = 500 // Basically "max quads"
	t_VERTEX_SIZE int = 4
)

func InitTextBatcher() *TextBatcher {

	batch := new(TextBatcher)

	batch.num_vertices = 0

	batch.vertices = make([]float32, t_BATCH_SIZE*t_VERTEX_SIZE*4)

	batch.vao = opengl.NewVertexArray()
	batch.vbo = opengl.NewVertexBuffer(batch.vertices)
	batch.layout = opengl.NewVertexBufferLayout()
	batch.layout.Pushf(2)
	batch.layout.Pushf(2)
	batch.vao.AddBuffer(*batch.vbo, *batch.layout)

	batch.generateIndexBuffer()

	batch.shader = opengl.NewShader("TextVertex.glsl", "TextFragment.glsl")


  fontImageData := loaders.LoadPNGFromResources("BitmapFont.png") 
  batch.texture = opengl.NewTextureFromData("BitmapFont.png", fontImageData, false)

	return batch

}

func (batch *TextBatcher) generateIndexBuffer() {

	batch.indeces = make([]uint32, t_BATCH_SIZE*6)

	for i := 0; i < t_BATCH_SIZE; i++ {

		offset := uint32(i * 4)

		batch.indeces[i*6+0] = 0 + offset
		batch.indeces[i*6+1] = 1 + offset
		batch.indeces[i*6+2] = 2 + offset

		batch.indeces[i*6+3] = 2 + offset
		batch.indeces[i*6+4] = 3 + offset
		batch.indeces[i*6+5] = 0 + offset

	}

	// logging.DebugLog(batch.indeces)
	batch.ibo = opengl.NewIndexBuffer(batch.indeces)
}

func (batch *TextBatcher) AddCharacter(char *FontAtlasItem, x, y float32, projectionMatrix, viewMatrix glm.Mat4) {

	if (batch.num_vertices + 4) > t_BATCH_SIZE*t_VERTEX_SIZE {
		// logging.DebugLog("Batch is full, flushing early")
    batch.FlushBatch(projectionMatrix, viewMatrix)
	}

	numVerts := batch.num_vertices * 4

	fx := x
	fy := y
	fEndX := float32(char.width) + fx
	fEndY := float32(FONT_SIZE) + fy

	// Bottom left
	batch.vertices[numVerts+0] = fx
	batch.vertices[numVerts+1] = fy
	batch.vertices[numVerts+2] = char.texturePositions[0]
	batch.vertices[numVerts+3] = char.texturePositions[1]

	// Top left
	batch.vertices[numVerts+4] = fx
	batch.vertices[numVerts+5] = fEndY
	batch.vertices[numVerts+6] = char.texturePositions[0]
	batch.vertices[numVerts+7] = char.texturePositions[3]

	// Top right
	batch.vertices[numVerts+8] = fEndX
	batch.vertices[numVerts+9] = fEndY
	batch.vertices[numVerts+10] = char.texturePositions[2]
	batch.vertices[numVerts+11] = char.texturePositions[3]

	// Bottom right
	batch.vertices[numVerts+12] = fEndX
	batch.vertices[numVerts+13] = fy
	batch.vertices[numVerts+14] = char.texturePositions[2]
	batch.vertices[numVerts+15] = char.texturePositions[1]

	batch.num_vertices += 4

	// logging.DebugLog(batch.vertices)

}

func (batch *TextBatcher) AddText(text string, x, y float32, projectionMatrix, viewMatrix glm.Mat4) {

	var offset int = 0

	for _, char := range text {

		character := _atlas[strings.ToUpper(string(char))]

		if character == nil {
			logging.DebugLogError("Could not find character in atlas:", character)
			return
		}

		batch.AddCharacter(character, x+float32(offset), y, projectionMatrix, viewMatrix)

		offset += 24

	}

}

func (batch *TextBatcher) FlushBatch(projectionMatrix, viewMatrix glm.Mat4) {

	batch.texture.Bind(1)

	batch.vbo.Bind()
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4*len(batch.vertices), gl.Ptr(&batch.vertices[0]))

	batch.shader.Bind()
	batch.shader.SetUniformMat4f("projection", projectionMatrix)
	batch.shader.SetUniformMat4f("view", viewMatrix)
	batch.shader.SetUniformMat4f("model", glm.Ident4())
	batch.shader.SetUniform1i("uTexture", 1)

	batch.ibo.Bind()
	batch.vao.Bind()

	indeces := (batch.num_vertices / 4) * 6

	gl.DrawElements(gl.TRIANGLES, int32(indeces), gl.UNSIGNED_INT, nil)


	batch.num_vertices = 0

}
