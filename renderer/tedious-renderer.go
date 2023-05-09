package renderer

import glm "github.com/go-gl/mathgl/mgl32"

// Basically a bad renderer that renders each entity with a draw call
type TediousRenderer struct {
  projectionMatrix glm.Mat4
  viewMatrix glm.Mat4
}


func (renderer *TediousRenderer) Init(windowWidth, windowHeight int) {

  // Initialize the renderer with a projection matrix and a view matrix
  renderer.projectionMatrix = glm.Ortho(0, float32(windowWidth), 0, float32(windowHeight), -1.0, 1.0)
  renderer.viewMatrix = glm.Ident4()

}


func (renderer *TediousRenderer) Render() {

}
