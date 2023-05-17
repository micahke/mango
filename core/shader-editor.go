package core

import (
	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util"
)

type ShaderEditor struct {
  activeShader shader_item
}

type shader_item struct {
  source string
  name string
}

// Constructor function
func NewShaderEditor() *ShaderEditor {
	editor := new(ShaderEditor)

	return editor
}

func (editor *ShaderEditor) RenderPanel() {
	imgui.PushID("shader_editor")

	imgui.ColumnsV(2, "", true)

	imgui.SetColumnWidth(0, imgui.WindowWidth()/4)
	imgui.SetColumnWidth(1, (imgui.WindowWidth()/4)*3)

  editor.drawSelectionScreen()
    
	imgui.NextColumn()

	editor.drawEditor()

	imgui.Columns()

	imgui.PopID()

}

func (editor *ShaderEditor) drawSelectionScreen() {
	imgui.Text("Select Shader")

  imgui.BeginChildV("shader_select", util.ImguiGenVec2(-1, -1), true, 0)

  for _, value := range(opengl.ShaderNames)  {
    if imgui.SelectableV(value, value == editor.activeShader.name, 0, util.ImguiGenVec2(0, 0)) {
      if value == editor.activeShader.name {
        editor.activeShader = shader_item{}
      }
      source, ok := opengl.ShaderCache[value]
      if !ok {
        logging.DebugLogError("Could not find source for this shader")
        continue
      }
      editor.activeShader = shader_item{
        name: value,
        source: source,
      }
    }
  }

  imgui.EndChild()
}

func (editor *ShaderEditor) drawEditor() {

	if editor.activeShader.name == "" {
		imgui.Text("No shader selected...")
		imgui.Spacing()
	} else {
    imgui.Text(editor.activeShader.name)
		imgui.Spacing()
		imgui.InputTextMultilineV("##editor", &editor.activeShader.source, util.ImguiGenVec2(-1, -1), 0, editor.imguiMultiTextCallback)
	}

		// imgui.InputTextMultilineV("##editor", &editor.activeShaderText, util.ImguiGenVec2(-1, -1), 0, editor.imguiMultiTextCallback)
}

func (editor *ShaderEditor) imguiMultiTextCallback(cb imgui.InputTextCallbackData) int32 {
	return 0
}
