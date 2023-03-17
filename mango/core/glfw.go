package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var glfwEnabled bool = false

func GLFWInit() {

	if err := glfw.Init(); err != nil {
		panic("Error initializing GLFW")
	}

	// enable OpenGL
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfwEnabled = true

}
