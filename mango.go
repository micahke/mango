package mango

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/im"
	"github.com/micahke/mango/input"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/system"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/util/loaders"
)

type Mango struct {
	RenderMode core.RenderMode // The currently set render mode
	Window     *core.Window
	LogPanel   *logging.LogPanel

	SceneEditor *core.SceneEditor

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
func Init(renderMode core.RenderMode, args ...any) {

	// create a nre instance of Mango
	Engine = new(Mango)

	// DEBUGGING ENABLED AFTER THIS POINT
	logging.DebugLog("LOGGING ENABLED")

	// set the rendering mode
	Engine.RenderMode = renderMode
	if renderMode == core.RENDER_MODE_IM {
		IM = im.Init()
	}

	// Enable GLFW
	core.GLFWInit()

	// Initialize resource loaders
	loaders.InitPNGLoader()

	// Load shaders
	_, error := opengl.LoadShaders()
	if error != nil {
		logging.DebugLogError("Failed to load shaders: ", error)
	}

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

  // Create logging panel
	Engine.LogPanel = logging.InitLogPanel(width, height)
	util.ImguiRegisterPanel("logPanel", Engine.LogPanel)
  if core.Settings.CONSOLE_ON_STARTUP {
    util.ImguiActivatePanel("logPanel")
  }

  // Scene Editor
  if Engine.RenderMode == core.RENDER_MODE_DEFAULT {
    Engine.SceneEditor = core.NewSceneEditor(Engine.scene)
    util.ImguiRegisterPanel("sceneEditor", Engine.SceneEditor)
    if core.Settings.SCENE_EDITOR_STARTUP {
      util.ImguiActivatePanel("sceneEditor")
    }
  }
}

// Creates a scene and sets up an ECS
func CreateScene() *core.Scene {

	logging.DebugLog("Scene creation requested: ")

	scene := core.NewScene()

	return scene

}

// Handles the creation of the core systems that the ECS uses
// These systems are the ones responsible to handling the entity's logic
func setupCoreSystems() {
  
}

func SetScene(scene *core.Scene) {

	Engine.scene = scene

	scene.ECS().AddSystem(&system.EntitySystem{
		Entities: scene.ECS().GetEntities(),
	})

  scene.ECS().AddSystem(&system.RenderSystem{
		Entities: scene.ECS().GetEntities(),
  })

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

    if util.ImguiWantsTextInput(){
      util.ImguiSetDefaultKeyCallback()
    } else {
      Engine.Window.SetKeyCallback(input.KeyCallback)
    }

		glfw.PollEvents()

		core.Timer.Update()
		util.ImguiNewFrame()

		// TODO: find a better place for this
		processInput()

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


	glfw.Terminate()

	cleanup()

}

// TODO: move this
func GetWindow() *core.Window {
	return Engine.Window
}

// handle any engine cleanup that needs to be done
func cleanup() {

}

func processInput() {
  // On left CTRL, show log
	if input.GetKeyDown(input.KEY_LEFT_CTRL) {
		util.ImguiTogglePanel("logPanel")
	}

  // On left tab, open scene editor
  if input.GetKeyDown(input.KEY_TAB) {
    util.ImguiTogglePanel("sceneEditor")
  }

}
