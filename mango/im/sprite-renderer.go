package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
)

type Sprite struct {
  x float32
  y float32
  width float32
  height float32

  texturePath string
  color glm.Vec4
}


type SpriteRenderer struct {

  vao *opengl.VertexArray
  vbo *opengl.VertexBuffer
  layout *opengl.VertexBufferLayout

  ibo *opengl.IndexBuffer
  shader *opengl.Shader

  modelMatrix glm.Mat4

}

// We can use the quad positions from the quad renderer
// TODO: be less lazy about this next time

var sprite_positions []float32 = []float32{
  0.0, 0.0, 0.0, 0.0,
  0.0, 1.0, 0.0, 1.0,
  1.0, 1.0, 1.0, 1.0,
  1.0, 0.0, 1.0, 0.0,
}


func InitSpriteRenderer() *SpriteRenderer {
  renderer := new(SpriteRenderer)
  
  renderer.vao = opengl.NewVertexArray()
  renderer.vbo = opengl.NewVertexBuffer(sprite_positions)
  renderer.layout = opengl.NewVertexBufferLayout()
  renderer.layout.Pushf(2)
  renderer.layout.Pushf(2)
  renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

  renderer.ibo = opengl.NewIndexBuffer(quad_indeces)
  renderer.shader = opengl.NewShader("SpriteVertex.glsl", "SpriteFragment.glsl")

  renderer.modelMatrix = glm.Ident4()

  return renderer

}


func (renderer *SpriteRenderer) RenderSprite(x, y, width, height float32, texturePath string, projectionMatrix, viewMatrix glm.Mat4) {

  texture := getTexture(texturePath)
  texture.Bind(0)

  translation := glm.Translate3D(x, y, 0)
  scale := glm.Scale3D(width, height, 1.0)
  model := translation.Mul4(scale).Mul4(renderer.modelMatrix)

  renderer.shader.Bind()
  renderer.shader.SetUniformMat4f("projection", projectionMatrix)
  renderer.shader.SetUniformMat4f("view", viewMatrix)
  renderer.shader.SetUniformMat4f("model", model)
  renderer.shader.SetUniform1i("uTexture", 0)
  
  // Should probably bind the texture

  renderer.vao.Bind()
  renderer.ibo.Bind()

  // Draw the sprite
  gl.DrawElements(gl.TRIANGLES, int32(renderer.ibo.GetCount()), gl.UNSIGNED_INT, nil)

}
