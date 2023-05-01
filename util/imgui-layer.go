package util

import (
	"github.com/AllenDang/imgui-go"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/glfw_imgui_backend"
)

type ImguiLayer struct {
	context *imgui.Context
	io      imgui.IO
	impl    *glfw_imgui_backend.ImguiGlfw3

	panels []*ImguiPanel
}

type ImguiPanel struct {
	name   string
	panel  interface{}
	active bool
}

type ImguiRenderer interface {
	RenderPanel()
}

var imgui_layer *ImguiLayer

func InitImguiLayer(window *glfw.Window) {

	imgui_layer = new(ImguiLayer)

	imgui_layer.context = imgui.CreateContext(nil)
	imgui_layer.io = imgui.CurrentIO()
  imgui_layer.io.SetIniFilename("")

	imgui_layer.impl = glfw_imgui_backend.ImguiGlfw3Init(window, imgui_layer.io)

}

func GetImguiLayer() *ImguiLayer {
	return imgui_layer
}

// Revert to the default key callback
func ImguiSetDefaultKeyCallback() {
  imgui_layer.impl.SetDefaultKeyCallback()
}



func ImguiNewFrame() {
	imgui_layer.impl.NewFrame()
}

func ImguiRender() {
	// Render all the panels that should be rendered

	for i := 0; i < len(imgui_layer.panels); i++ {
		panel := imgui_layer.panels[i].panel.(ImguiRenderer)
		if imgui_layer.panels[i].active {
			panel.RenderPanel()
		}
	}

	imgui.Render()
	imgui_layer.impl.Render(imgui.RenderedDrawData())
}

func ImguiRegisterPanel(name string, i interface{}) {
	panel := new(ImguiPanel)
	panel.active = false
	panel.name = name
	panel.panel = i

	imgui_layer.panels = append(imgui_layer.panels, panel)

}

func ImguiActivatePanel(name string) {

	for i := 0; i < len(imgui_layer.panels); i++ {
		if imgui_layer.panels[i].name == name {
			imgui_layer.panels[i].active = true
			break
		}
	}

}

func ImguiTogglePanel(name string) {
	for i := 0; i < len(imgui_layer.panels); i++ {
		if imgui_layer.panels[i].name == name {
			imgui_layer.panels[i].active = !imgui_layer.panels[i].active
			break
		}
	}
}

func ImguiDeactivatePanel(name string) {

	for i := 0; i < len(imgui_layer.panels); i++ {
		if imgui_layer.panels[i].name == name {
			imgui_layer.panels[i].active = false
			break
		}
	}

}

func ImguiPanelStatus(name string) *bool {
	for i := 0; i < len(imgui_layer.panels); i++ {
		if imgui_layer.panels[i].name == name {
			return &imgui_layer.panels[i].active
		}
	}
	return nil
}

func ImguiDestroy() {
	imgui_layer.context.Destroy()
	imgui_layer.impl.Shutdown()
}

func ImguiWantsMouse() bool {

	return imgui_layer.io.WantCaptureMouse() == true

}

func ImguiWantsTextInput() bool {
	return imgui_layer.io.WantTextInput()
}

func ImguiWantsKeyInput() bool {
	return imgui_layer.io.WantCaptureKeyboard()
}
