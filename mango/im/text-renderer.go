package im

import (
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/opengl"
)


type TextRenderer struct {


  vao *opengl.VertexArray
  vbo *opengl.VertexBuffer
  layout *opengl.VertexBufferLayout

  ibo *opengl.IndexBuffer
  shader *opengl.Shader


  modelMatrix glm.Mat4

}


func InitTextRenderer() *TextRenderer {

  renderer := new(TextRenderer)

  renderer.vao = opengl.NewVertexArray()
  renderer.vbo = opengl.NewVertexBuffer(sprite_positions)
  renderer.layout = opengl.NewVertexBufferLayout()
  renderer.layout.Pushf(2)
  renderer.layout.Pushf(2)

  renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

  renderer.ibo = opengl.NewIndexBuffer(quad_indeces)

  renderer.shader = opengl.NewShader("TextVertex.glsl", "TextFragment.glsl")

  renderer.modelMatrix = glm.Ident4()

  // build font atlas


  return renderer

}


func (renderer *TextRenderer) RenderText(x, y, size float32, text string, projectionMatrix, viewMatrix glm.Mat4) {
  
  xOffset := x


  // Loop through the letters in text
  
  for _, letter := range text {

    // if letter is a letter uppercase it
    character := _atlas[strings.ToUpper(string(letter))]
    logging.DebugLog(character)
  

	texture := getTexture("BitmapFont.png", true)

  texturePositions := glm.Vec4{
    float32(character.x) / float32(texture.GetWidth()), // X
    float32(character.y) / float32(texture.GetHeight()), // Y
    float32(character.width) / float32(texture.GetWidth()), // WIDTH
    float32(FONT_SIZE) / float32(texture.GetHeight()), // HEIGHT
  }

  logging.DebugLog(texturePositions)

	texture.Bind(1)
  texture.UpdateSubImage(character.x, character.y, character.width, FONT_SIZE)

  renderer.vbo.SetData([]float32{
    0.0, 0.0, texturePositions[0], texturePositions[1], 
    0.0, 1.0, texturePositions[0], texturePositions[1] + texturePositions[3],
    1.0, 1.0, texturePositions[0] + texturePositions[2], texturePositions[1] + texturePositions[3],
    1.0, 0.0, texturePositions[0] + texturePositions[2], texturePositions[1],
  })




	translation := glm.Translate3D(xOffset, y, 0)
	scale := glm.Scale3D(size, size, 1.0)
	model := translation.Mul4(scale).Mul4(renderer.modelMatrix)

	renderer.shader.Bind()
	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", model)
	renderer.shader.SetUniform1i("uTexture", 1)


	renderer.vao.Bind()
	renderer.ibo.Bind()

	// Draw the sprite
	gl.DrawElements(gl.TRIANGLES, int32(renderer.ibo.GetCount()), gl.UNSIGNED_INT, nil)

    xOffset += float32(character.width)
  }

}


