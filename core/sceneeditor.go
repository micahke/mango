package core

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/ecs"
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
	currentEntity := editor.Scene.ECS().GetEntity(entityNames[editor.currentEntityIndex])

	if currentEntity == nil {
		return
	}

	size := imgui.Vec2{
		X: 400,
		Y: 400,
	}

	imgui.SetNextWindowSizeV(size, imgui.ConditionOnce)
	imgui.BeginV("Scene Editor", util.ImguiPanelStatus("sceneEditor"), 0)

  if imgui.Button("Add Entity") {
    editor.Scene.CreateEntity("Unnamed Entity")
  }

	imgui.ListBox("Entity List", &editor.currentEntityIndex, entityNames)

	imgui.Spacing()
	imgui.Separator()

	imgui.Spacing()
	imgui.Text(fmt.Sprint("Selected entity:\t", currentEntity.Name))
	imgui.Spacing()

  imgui.InputText("Edit Name", &currentEntity.Name)

	imgui.Spacing()

	for _, component := range currentEntity.Components {
		// Get the name of the component
		name := ecs.GetComponentName(component)
		// Draw the tree node for the component
		if imgui.TreeNodeV(name, 2) {

			editor.renderControlPanel(component)

			imgui.TreePop()
		}

	}

	imgui.End()

}

func (editor *SceneEditor) getEntityNames() []string {

	names := make([]string, len(*editor.Scene.ecs.GetEntities()))

	for i, entity := range *editor.Scene.ecs.GetEntities() {
		names[i] = entity.Name
	}

	return names

}

func (editor *SceneEditor) renderControlPanel(component interface{}) {

	// Check to see whether the entity implements the UIEditableComponent interface

	// imgui.PushItemWidth(-1.0)
	// imgui.BeginChildV("Control Panel", imgui.Vec2{
	// 	X: 0,
	// 	Y: 0,
	// }, true, 0)

	controlPanel, ok := component.(ecs.UIEditableComponent)

	if ok {
		controlPanel.RenderControlPanel()
	} else {
		imgui.Text("This component does not implement a control panel")
	}

	// imgui.EndChild()

}
