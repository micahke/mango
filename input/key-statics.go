package input

import "github.com/go-gl/glfw/v3.3/glfw"

type Key glfw.Key

var KEY_PRESSED bool
var KEY_REPEATED bool

// NUMBER KEYS
var KEY_0 Key = Key(glfw.Key0)
var KEY_1 Key = Key(glfw.Key1)
var KEY_2 Key = Key(glfw.Key2)
var KEY_3 Key = Key(glfw.Key3)
var KEY_4 Key = Key(glfw.Key4)
var KEY_5 Key = Key(glfw.Key5)
var KEY_6 Key = Key(glfw.Key6)
var KEY_7 Key = Key(glfw.Key7)
var KEY_8 Key = Key(glfw.Key8)
var KEY_9 Key = Key(glfw.Key9)

// LETTER KEYS
var KEY_A Key = Key(glfw.KeyA)
var KEY_B Key = Key(glfw.KeyB)
var KEY_C Key = Key(glfw.KeyC)
var KEY_D Key = Key(glfw.KeyD)
var KEY_E Key = Key(glfw.KeyE)
var KEY_F Key = Key(glfw.KeyF)
var KEY_G Key = Key(glfw.KeyG)
var KEY_H Key = Key(glfw.KeyH)
var KEY_I Key = Key(glfw.KeyI)
var KEY_J Key = Key(glfw.KeyJ)
var KEY_K Key = Key(glfw.KeyK)
var KEY_L Key = Key(glfw.KeyL)
var KEY_M Key = Key(glfw.KeyM)
var KEY_N Key = Key(glfw.KeyN)
var KEY_O Key = Key(glfw.KeyO)
var KEY_P Key = Key(glfw.KeyP)
var KEY_Q Key = Key(glfw.KeyQ)
var KEY_R Key = Key(glfw.KeyR)
var KEY_S Key = Key(glfw.KeyS)
var KEY_T Key = Key(glfw.KeyT)
var KEY_U Key = Key(glfw.KeyU)
var KEY_V Key = Key(glfw.KeyV)
var KEY_W Key = Key(glfw.KeyW)
var KEY_X Key = Key(glfw.KeyX)
var KEY_Y Key = Key(glfw.KeyY)
var KEY_Z Key = Key(glfw.KeyZ)

// Directional keys
var KEY_LEFT Key = Key(glfw.KeyLeft)
var KEY_RIGHT Key = Key(glfw.KeyRight)
var KEY_UP Key = Key(glfw.KeyUp)
var KEY_DOWN Key = Key(glfw.KeyDown)

// Function keys
var KEY_F1 Key = Key(glfw.KeyF1)
var KEY_F2 Key = Key(glfw.KeyF2)
var KEY_F3 Key = Key(glfw.KeyF3)
var KEY_F4 Key = Key(glfw.KeyF4)
var KEY_F5 Key = Key(glfw.KeyF5)
var KEY_F6 Key = Key(glfw.KeyF6)
var KEY_F7 Key = Key(glfw.KeyF7)
var KEY_F8 Key = Key(glfw.KeyF8)
var KEY_F9 Key = Key(glfw.KeyF9)

// Other keys
var KEY_ENTER Key = Key(glfw.KeyEnter)
var KEY_DELETE Key = Key(glfw.KeyDelete)
var KEY_LEFT_CTRL Key = Key(glfw.KeyLeftControl)
var KEY_RIGHT_CTRL Key = Key(glfw.KeyRightControl)
var KEY_LEFT_SHIFT Key = Key(glfw.KeyLeftShift)
var KEY_RIGHT_SHIFT Key = Key(glfw.KeyRightShift)
var KEY_LEFT_ALT Key = Key(glfw.KeyLeftAlt)
var KEY_RIGHT_ALT Key = Key(glfw.KeyRightAlt)
var KEY_TAB Key = Key(glfw.KeyTab)
