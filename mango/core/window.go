package core

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)


type Window struct {
  *glfw.Window
  width int
  height int
  title string

  vsyncEnabled bool
}


func CreateWindow(width, height int, title string, vsync bool) *Window {
  if !glfwEnabled {
    panic("Cant make window: GLFW not enabled")
  }

  // create new Window instance
  window := new(Window)

  w, err := glfw.CreateWindow(width, height, title, nil, nil)
  if err != nil {
    panic(err)
  }

  w.MakeContextCurrent()
  if vsync {
    glfw.SwapInterval(1)
  } else {
    glfw.SwapInterval(0)
  }

  if err := gl.Init(); err != nil {
    panic("Error initializing OpenGL")
  }

  // Enable blending
  gl.Enable(gl.BLEND)
  gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

  // Add data to window object
  window.Window = w
  window.width = width
  window.height = height
  window.title = title
  window.vsyncEnabled = vsync

  return window

}
