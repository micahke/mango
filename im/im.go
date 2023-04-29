package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/res"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/util/color"
	"github.com/micahke/mango/util/loaders"
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
	textRenderer   *TextRenderer
	textBatcher    *TextBatcher
	pixelRenderer  *PixelRenderer
	quadBatcher    *QuadBatcher
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


	scene, _ := im.scene.(IM_SCRIPT)
	if !im.sceneActivated {
		scene.Init()
		im.sceneActivated = true
	}
	scene.Update()

  // RIP: forgot about the rendering order
  // To do: make this more general
  im.quadBatcher.FlushBatch(im.projectionMatrix, im.viewMatrix)


	scene.Draw()

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

  // Load the bitmap font before we start any text rendering
  fontData, _ := res.LoadEngineResource("BitmapFont.png")
  loaders.LoadImageFromData("BitmapFont.png", fontData)
	im.textRenderer = InitTextRenderer()
	im.textBatcher = InitTextBatcher()

	im.pixelRenderer = InitPixelRenderer()
  im.quadBatcher = InitQuadBatcher()
}

func (im *IMMEDIATE_MODE) SetBackgroundColor(color color.Color) {

	colorVec := color.Vec4

	gl.ClearColor(colorVec[0], colorVec[1], colorVec[2], colorVec[3])

}

// Draw a quad to the screen
func (im *IMMEDIATE_MODE) DrawRect(x, y, width, height float32) {
	im.quadRenderer.RenderQuad(x, y, width, height, color.WHITE, im.projectionMatrix, im.viewMatrix)
}

// Draw a filled rectangle to the screen
func (im *IMMEDIATE_MODE) FillRect(x, y, width, height float32, color color.Color) {
	im.quadRenderer.RenderQuad(x, y, width, height, color, im.projectionMatrix, im.viewMatrix)
}

// Draw a sprite to the screen
func (im *IMMEDIATE_MODE) DrawSprite(x, y, width, height float32, texturePath string) {

	im.spriteRenderer.RenderSprite(x, y, width, height, texturePath, im.projectionMatrix, im.viewMatrix)

}

// Draw a UV sprite to the screen defined by a UV spritemap
func (im *IMMEDIATE_MODE) DrawUVSprite(x, y, width, height float32, texturePath string, uv util.UVSpriteMap) {

	im.spriteRenderer.RenderUVSprite(x, y, width, height, texturePath, uv, im.projectionMatrix, im.viewMatrix)

}

// Draw a circle to the screen
// Centered at bottom left
func (im *IMMEDIATE_MODE) DrawCircle(x, y, width, height float32, color color.Color) {
	im.circleRenderer.RenderCircle(x, y, width, height, color.Vec4, im.projectionMatrix, im.viewMatrix)
}

// Draw a line to the screen
func (im *IMMEDIATE_MODE) DrawGLine(x1, y1, x2, y2 float32, color color.Color, thickness float32) {

	p1 := glm.Vec2{x1, y1}
	p2 := glm.Vec2{x2, y2}

	im.lineRenderer.RenderLine(p1, p2, color, thickness, im.projectionMatrix, im.viewMatrix)

}

func (im *IMMEDIATE_MODE) DrawLine(x0, y0, x1, y1, size int) {
  linePoints := im.lineRenderer.GenerateBresenhamPoints(x0, y0, x1, y1, size)
  
  // TODO: fix unecessary conversions between float32 and int
  im.DrawPixels(linePoints, float32(size))
}

// Draw text within the world space
// This is not batch optimized
func (im *IMMEDIATE_MODE) DrawWorldText(x, y, size float32, text string) {
	im.textRenderer.RenderText(x, y, size, text, im.projectionMatrix, im.viewMatrix)
}

// Draw text to the screen layer
// This is batch renedered and optimized
func (im *IMMEDIATE_MODE) DrawText(text string, x, y float32) {
	im.textBatcher.AddText(text, x, y, im.projectionMatrix, im.viewMatrix)
}

func (im *IMMEDIATE_MODE) DrawPixels(buffer []float32, size float32) {

	im.pixelRenderer.DrawPixels(buffer, 2*size, im.projectionMatrix, im.viewMatrix)

}

func (im *IMMEDIATE_MODE) DrawQuad(quad Quad) {

  im.quadBatcher.AddQuad(quad, im.projectionMatrix, im.viewMatrix)

}
