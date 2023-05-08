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
  windowWidth int
  windowHeight int
}

func NewSceneEditor(scene *Scene, windowWidth, windowHeight int) *SceneEditor {

	editor := &SceneEditor{}

	editor.Scene = scene
  editor.windowWidth = windowWidth
  editor.windowHeight = windowHeight

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
		Y: float32(editor.windowHeight),
	}

  position := imgui.Vec2{
    X: float32(editor.windowWidth) - size.X,
    Y: 0,
  }

  imgui.SetNextWindowPosV(position, imgui.ConditionOnce, imgui.Vec2{X: 0, Y: 0})
	imgui.SetNextWindowSizeV(size, imgui.ConditionOnce)
	imgui.BeginV("Scene Editor", util.ImguiPanelStatus("sceneEditor"), 0)

  if imgui.Button("Add Entity") {
    editor.Scene.CreateEntity("Unnamed Entity")
  }

  var heightItems int = 5
  if len(entityNames) > 10 {
    heightItems = 10
  }

	imgui.ListBoxV("Entity List", &editor.currentEntityIndex, entityNames, heightItems)

	imgui.Spacing()
	imgui.Separator()

	imgui.Spacing()
	imgui.Text(fmt.Sprint("Selected entity:\t", currentEntity.Name))
	imgui.Spacing()

  imgui.InputText("Edit Name", &currentEntity.Name)

	imgui.Spacing()

  imgui.BeginChild("Components")

	for _, component := range currentEntity.Components {
		// Get the name of the component
		name := ecs.GetComponentName(component)
		// Draw the tree node for the component
		if imgui.TreeNodeV(name, 2) {

			editor.renderControlPanel(component)

			imgui.TreePop()
		}

	}

  imgui.EndChild()

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
