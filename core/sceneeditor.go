package core

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/util"
)

type SceneEditor struct {
	Scene *Scene

  currentEntityIndex int32
}

func NewSceneEditor(scene *Scene) *SceneEditor {

	editor := &SceneEditor{}

	editor.Scene = scene

	return editor

}

// Render the editor panel
func (editor *SceneEditor) RenderPanel() {

  entityNames := editor.getEntityNames()

  imgui.SetNextWindowSizeV(imgui.Vec2{400, 400}, imgui.ConditionOnce)
	imgui.BeginV("Scene Editor", util.ImguiPanelStatus("sceneEditor"), 0)

  imgui.Text("Entity List")
  imgui.PushItemWidth(-1.0)
	imgui.ListBox("List Box", &editor.currentEntityIndex, entityNames)

  imgui.Separator()
  
  imgui.Text(fmt.Sprint("Selected entity:\t", entityNames[editor.currentEntityIndex]))

	imgui.End()

}


func (editor *SceneEditor) getEntityNames() []string {

  names := make([]string, len(*editor.Scene.ecs.GetEntities()))

  for i, entity := range(*editor.Scene.ecs.GetEntities()) {
    names[i] = entity.Name
  }

  return names

}
