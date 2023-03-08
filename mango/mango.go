package mango

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/infinite-universe/mango/core"
	"github.com/micahke/infinite-universe/mango/im"
	"github.com/micahke/infinite-universe/mango/util/loaders"
)

type Mango struct {
	RenderMode core.RenderMode // The currently set render mode
	Window     *core.Window
}

// The main engine instance
var Engine *Mango

// RENDER MODES
var IM *im.IMMEDIATE_MODE

// TIMER
var Time *core.Timer

// Initialization function for the engine
func Init(renderMode core.RenderMode) {
	// Lock the runtime

	// create a nre instance of Mango
	Engine = new(Mango)

	// set the rendering mode
	Engine.RenderMode = renderMode
	if renderMode == core.RENDER_MODE_IM {
		IM = im.Init()
	}

	// Enable GLFW
	core.GLFWInit()

	// Initialize resource loaders
	loaders.InitPNGLoader()

}

// Creates a new window
func CreateWindow(width, height int, title string) {

	Engine.Window = core.CreateWindow(width, height, title, true)

	if Engine.RenderMode == core.RENDER_MODE_IM {
		IM.InitProjectionMatrix(float32(width), float32(height))
	}

}

// Starts the main game loop
func Start() {

	if Engine.Window == nil {
		panic("No window initialized")
	}

	Time = core.NewTimer()

	for !Engine.Window.Window.ShouldClose() {
		Time.Update()

		// Check the rendermode and do appropriate stuff
		if Engine.RenderMode == core.RENDER_MODE_IM {
			IM.NewFrame(Time.DeltaTime())
		}

		Engine.Window.Window.SwapBuffers()

	}

	glfw.Terminate()

}
