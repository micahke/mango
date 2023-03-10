package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/infinite-universe/mango/util"
)

// +++++++++++ EXPORTED +++++++++++++++++++

var MouseX float64 = 0.0
var MouseY float64 = 0.0

var MouseLeftPressed bool = false
var MouseMiddlePressed bool = false
var MouseRightPressed bool = false

var MouseLeftReleased bool = false
var MouseMiddleReleased bool = false
var MouseRightReleased bool = false

type MOUSE_BUTTON int

const (
  MOUSE_BUTTON_LEFT MOUSE_BUTTON = 0
  MOUSE_BUTTON_MIDDLE  MOUSE_BUTTON = 1
  MOUSE_BUTTON_RIGHT MOUSE_BUTTON = 2
)


func CursorCallback(window *glfw.Window, xpos float64, ypos float64) {

  _, height := window.GetSize()
  MouseX = xpos
  MouseY = float64(height) - ypos

}

func MouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

  if util.ImguiWantsMouse() {
    return
  }

  if button == glfw.MouseButtonLeft {
    if action == glfw.Press {
      MouseLeftPressed = true
    }
    if action == glfw.Release {
      MouseLeftReleased = true
    }
  }

  if button == glfw.MouseButtonMiddle {
    if action == glfw.Press {
      MouseMiddlePressed = true
    }
    if action == glfw.Release {
      MouseMiddleReleased = true
    }
  }

  if button == glfw.MouseButtonRight {
    if action == glfw.Press {
      MouseRightPressed = true
    }
    if action == glfw.Release {
      MouseRightReleased = true
    }
  }

}

func MouseButtonPressed(mouseButton MOUSE_BUTTON) bool {
  if mouseButton == MOUSE_BUTTON_LEFT {
    return MouseLeftPressed
  }
  if mouseButton == MOUSE_BUTTON_MIDDLE {
    return MouseMiddlePressed
  }
  if mouseButton == MOUSE_BUTTON_RIGHT {
    return MouseRightPressed
  }
  return false
}

func MouseInputCleanup() {
  MouseLeftPressed = false
  MouseMiddlePressed = false
  MouseRightPressed = false

  MouseLeftReleased = false
  MouseMiddleReleased = false
  MouseRightReleased = false

}
