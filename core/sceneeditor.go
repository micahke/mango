package core

import (
	"fmt"
	"reflect"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/util"
)

type SceneEditor struct {
	Scene *Scene

	currentEntityIndex int32
	windowWidth        int
	windowHeight       int

  currentEntity *ecs.Entity
}

var component_list []ecs.NamedComponent = []ecs.NamedComponent{
  &components.SampleComponent{},
  // &components.TransformComponent{},
  // &components.Shape2DComponent{},
  &components.PrimitiveRenderer{},
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
    newName := fmt.Sprint("entity", len(*editor.Scene.ecs.GetEntities()))
		editor.Scene.CreateEntity(newName)
	}

  if len(*editor.Scene.ECS().GetEntities()) == 0 {
    imgui.End()
    return
  }

	entityNames := editor.getEntityNames()
	currentEntity := editor.Scene.ECS().GetEntity(entityNames[editor.currentEntityIndex])
  editor.currentEntity = currentEntity

	if currentEntity == nil {
		return
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

  editor.renderAddComponentList()

	for _, component := range currentEntity.Components {
		// Get the name of the component
		name := ecs.GetComponentName(component)
		// Draw the tree node for the component

		// If it is a transform component, set the element to open (once)
		if name == "Transform" {
			imgui.SetNextItemOpen(true, imgui.ConditionOnce)
		}
		if imgui.TreeNodeV(name, 2) {

      // TODO: fix this so that we're passing references because
      // I think it's fucking with which entity instance we're actually
      // in control of
			editor.renderControlPanel(component)

      if reflect.TypeOf(component) != reflect.TypeOf(&components.TransformComponent{}) {
        editor.drawRemoveButton(component)
      }

			imgui.TreePop()
		}

	}

	imgui.EndChild()

	imgui.End()

}

func (editor *SceneEditor) drawRemoveButton(component interface{}) {
  
  c, ok := component.(ecs.Component)
  if !ok {
    logging.DebugLogError("Failed converting component")
  }

    imgui.Spacing()

  // Make the button font size smaller
  if imgui.ButtonV("Remove Component", util.ImguiGenVec2(-1, 0)) {
    editor.currentEntity.RemoveComponentByType(reflect.TypeOf(c))
    // editor.currentEntity.RemoveComponent(component)
  }

}

func (editor *SceneEditor) renderAddComponentList() {
  if imgui.IsMouseReleased(1) {
    imgui.OpenPopup("component_list")
  }

  if imgui.BeginPopupContextItemV("component_list", 0) {
    imgui.Text("Add Component")
    imgui.Spacing()
    imgui.Separator()
    imgui.Spacing()

    for _, component  := range(component_list) {

      if imgui.Selectable(component.GetComponentName()) {
        editor.addComponentToEntity(component, editor.currentEntity)
      }

    }

    imgui.EndPopup()
  }
}

func (editor *SceneEditor) getEntityNames() []string {

	names := make([]string, len(*editor.Scene.ecs.GetEntities()))

	for i, entity := range *editor.Scene.ecs.GetEntities() {
		names[i] = entity.Name
	}

	return names

}


func (editor *SceneEditor) addComponentToEntity(component interface{}, entity *ecs.Entity) {

  cmpt := reflect.New(reflect.TypeOf(component).Elem()).Interface().(ecs.Component)

  entity.AddComponent(cmpt)

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
