package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/mango/core"
)

var window *core.Window

func InitKeyboardInput(w *core.Window) {
	window = w
}

func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if action == glfw.Press {
		KEY_PRESSED = true
	}

}

// Returns true is the key is being held down
func GetKey(key Key) bool {

	if window.GetKey(glfw.Key(key)) == glfw.Press {
		return true
	}
	return false

}

// Returns true if the key is in a pressed state
func GetKeyDown(key Key) bool {
	if GetKey(key) && KEY_PRESSED {
		return true
	}
	return false
}

func ResetKeyInput() {
	KEY_PRESSED = false
}
