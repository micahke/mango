package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/util"
)

type IMMEDIATE_MODE struct {
	// These deal with scripts
	scene          interface{}
	sceneActivated bool

	// This deals with camera
	projectionMatrix glm.Mat4
	viewMatrix       glm.Mat4

	// Renderers
	quadRenderer   *QuadRenderer
	spriteRenderer *SpriteRenderer
	circleRenderer *CircleRenderer
	lineRenderer   *LineRenderer
  // textRenderer   *TextRenderer
  textBatcher   *TextBatcher
}

func Init() *IMMEDIATE_MODE {
	im_mode := new(IMMEDIATE_MODE)

	InitTextureCache()
  InitFontAtlas()
	im_mode.viewMatrix = glm.Ident4()

	return im_mode
}

// Starts a new frame in IMMEDIATE MODE
// This function should only be used internally by the engine
func (im *IMMEDIATE_MODE) NewFrame(deltaTime float64) {

	// Poll for events
	// glfw.PollEvents()

	// Clear the screen
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	scene, _ := im.scene.(IM_SCRIPT)
	if !im.sceneActivated {
		scene.Init()
		im.sceneActivated = true
	}
	scene.Update(deltaTime)
	scene.Draw()

  im.textBatcher.InitBatch()
  im.textBatcher.AddCharacter("T", 0, 0)
  im.textBatcher.AddCharacter("9", 100, 100)
  // im.textBatcher.AddCharacter("W", 100, 100)
  im.textBatcher.FlushBatch(im.projectionMatrix, im.viewMatrix)

}

func (im *IMMEDIATE_MODE) ConnectScene(i interface{}) {
	// Pass in whatever interface we get
	im.scene = i
}

func (im *IMMEDIATE_MODE) InitProjectionMatrix(width, height float32) {
	im.projectionMatrix = glm.Ortho(0, width, 0, height, -1.0, 1.0)

	// Also set up renderers here because OpenGL has just been set up
	im.setupRenderers()
}

func (im *IMMEDIATE_MODE) setupRenderers() {
	im.quadRenderer = InitQuadRenderer()
	im.spriteRenderer = InitSpriteRenderer()
	im.circleRenderer = InitCircleRenderer()
	im.lineRenderer = InitLineRenderer()
  // im.textRenderer = InitTextRenderer()
  im.textBatcher = InitTextBatcher()
}

func (im *IMMEDIATE_MODE) SetBackgroundColor(color util.Color) {

	colorVec := color.Vec4

	gl.ClearColor(colorVec[0], colorVec[1], colorVec[2], colorVec[3])

}

// API FOR DRAWING QUADA

func (im *IMMEDIATE_MODE) DrawRect(x, y, width, height float32) {
	im.quadRenderer.RenderQuad(x, y, width, height, util.WHITE, im.projectionMatrix, im.viewMatrix)
}

func (im *IMMEDIATE_MODE) FillRect(x, y, width, height float32, color util.Color) {
	im.quadRenderer.RenderQuad(x, y, width, height, color, im.projectionMatrix, im.viewMatrix)
}

func (im *IMMEDIATE_MODE) DrawSprite(x, y, width, height float32, texturePath string) {

	im.spriteRenderer.RenderSprite(x, y, width, height, texturePath, im.projectionMatrix, im.viewMatrix)

}

func (im *IMMEDIATE_MODE) DrawUVSprite(x, y, width, height float32, texturePath string, uv util.UVSpriteMap) {

	im.spriteRenderer.RenderUVSprite(x, y, width, height, texturePath, uv, im.projectionMatrix, im.viewMatrix)

}

func (im *IMMEDIATE_MODE) DrawCircle(x, y, width, height float32, color util.Color) {
	im.circleRenderer.RenderCircle(x, y, width, height, color.Vec4, im.projectionMatrix, im.viewMatrix)
}

func (im *IMMEDIATE_MODE) DrawLine(x1, y1, x2, y2 float32, color util.Color, thickness float32) {

	p1 := glm.Vec2{x1, y1}
	p2 := glm.Vec2{x2, y2}

	im.lineRenderer.RenderLine(p1, p2, color, thickness, im.projectionMatrix, im.viewMatrix)

}


func (im *IMMEDIATE_MODE) DrawText(x, y, size float32, text string) {
  // im.textRenderer.RenderText(x, y, size, text, im.projectionMatrix, im.viewMatrix)
}
