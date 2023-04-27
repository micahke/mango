package mango

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/infinite-universe/mango/core"
	"github.com/micahke/infinite-universe/mango/im"
	"github.com/micahke/infinite-universe/mango/input"
	"github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/mango/util/loaders"
)

type Mango struct {
	RenderMode core.RenderMode // The currently set render mode
	Window     *core.Window
	LogPanel   *logging.LogPanel
}

// The main engine instance
var Engine *Mango

// RENDER MODES
var IM *im.IMMEDIATE_MODE

// TIMER

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
func CreateWindow(width, height int, title string, vsync bool) {

	Engine.Window = core.CreateWindow(width, height, title, vsync)

	Engine.Window.SetCursorPosCallback(input.CursorCallback)
	Engine.Window.SetMouseButtonCallback(input.MouseButtonCallback)
	Engine.Window.SetKeyCallback(input.KeyCallback)

	// Initialize the keyboard input system
	input.InitKeyboardInput(Engine.Window)

	if Engine.RenderMode == core.RENDER_MODE_IM {
		IM.InitProjectionMatrix(float32(width), float32(height))
	}

	// At this point, OpenGL is ready to be used anywhere in the program

	util.InitImguiLayer(Engine.Window.Window)
	Engine.LogPanel = logging.InitLogPanel(width, height)
	util.ImguiRegisterPanel("logPanel", Engine.LogPanel)

}

func GetWindow() *core.Window {
	return Engine.Window
}

// Starts the main game loop
func Start() {

	if Engine.Window == nil {
		panic("No window initialized")
	}

	core.InitTimer()

	for !Engine.Window.Window.ShouldClose() {
		start := glfw.GetTime()

		Engine.Window.SetMouseButtonCallback(input.MouseButtonCallback)
		Engine.Window.SetKeyCallback(input.KeyCallback)

		glfw.PollEvents()

		core.Timer.Update()
		util.ImguiNewFrame()

    // TODO: find a better place for this
		update()
  

		// Check the rendermode and do appropriate stuff
		if Engine.RenderMode == core.RENDER_MODE_IM {
			IM.NewFrame(core.Timer.DeltaTime())
		}

		util.ImguiRender()

		// Input handler reset
		input.MouseInputCleanup()
		input.ResetKeyInput()

		Engine.Window.Window.SwapBuffers()

		end := glfw.GetTime()
		core.Timer.UpdateFrameData(start, end)
	}

	// util.ImguiDestroy()

	glfw.Terminate()

}

func update() {

	if input.GetKeyDown(input.KEY_LEFT_CTRL) {
		util.ImguiTogglePanel("logPanel")
	}
}
