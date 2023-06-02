package main

import (
	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/ecs/prefabs"
	// "github.com/micahke/mango/components"
	// "github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/core/settings"
)

func main() {

	mango.Init(core.RENDER_MODE_DEFAULT)

	settings.Settings.SCENE_EDITOR_STARTUP = true
	// settings.Settings.IMGUI_SANDBOX_ON_STARTUP = true
	// core.Settings.CONSOLE_ON_STARTUP = true
	// core.Settings.SHADER_EDITOR_ON_STARTUP = true

	// scene setup
	scene := mango.CreateScene()
	mango.SetScene(scene)

	player := scene.CreateEntity("player")
	player.Tranform().Position.X = 100

  pr := &components.PrimitiveRenderer{}
  pr.Shape = &shape.Rect{
    Width: 100,
    Height: 100,
  }
  player.AddComponent(pr)
  // player.AddComponent(&ecs.TestComponent{})

	// p2 := scene.CreateEntity("player2")
	// p2.Tranform().Position.X = 400
	// p2.Tranform().Position.Y = 300
	// p2Shape := &components.Shape2DComponent{}
	// p2Shape.SetShape(&shape.Rect{
	//   Width: 30,
	//   Height: 20,
	// })
	// p2.AddComponent(p2Shape)
	// p2.AddComponent(&components.PrimitiveRenderer{})

	circle := prefabs.NewCirclePrefab()
	scene.AddPrefab(circle)

	rect := prefabs.NewSquarePrefab()
	scene.AddPrefab(rect)

	mango.CreateWindow(1300, 800, "Retained Mode Rendering", false)
	mango.Start()

}
