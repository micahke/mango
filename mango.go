package mango

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/im"
	"github.com/micahke/mango/input"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/util/loaders"
)

type Mango struct {
	RenderMode core.RenderMode // The currently set render mode
	Window     *core.Window
	LogPanel   *logging.LogPanel

  scene *core.Scene
}

// The main engine instance
var Engine *Mango

// RENDER MODES
var IM *im.IMMEDIATE_MODE

func init() {
  runtime.LockOSThread()
}

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

  // DEBUGGING ENABLED AFTER THIS POINT

  // Load shaders
  shaders, error := opengl.LoadShaders()
  if error != nil {
    logging.DebugLogError("Failed to load shaders: ", error)
  } else {
    logging.DebugLog(shaders["CircleVertex.glsl"])
  }

}


// Creates a scene and sets up an ECS
func CreateScene() *core.Scene {

  scene := core.NewScene()
  scene.ECS().AddSystem(&ecs.EntitySystem{
    Entities: scene.ECS().GetEntities(),
  })

  return scene

}


func SetScene(scene *core.Scene) {
  
  Engine.scene = scene

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
  

    // Clear the screen
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Check the rendermode and do appropriate stuff
		if Engine.RenderMode == core.RENDER_MODE_IM {
			IM.NewFrame(core.Timer.DeltaTime())
		}

    if Engine.RenderMode == core.RENDER_MODE_DEFAULT {
      Engine.scene.ECS().Update()
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


func GetWindow() *core.Window {
	return Engine.Window
}




func update() {

	if input.GetKeyDown(input.KEY_LEFT_CTRL) {
		util.ImguiTogglePanel("logPanel")
	}
}
