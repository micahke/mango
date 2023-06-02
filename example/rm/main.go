package main

import (
	"github.com/micahke/mango"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/ecs/prefabs"

	// "github.com/micahke/mango/components"
	// "github.com/micahke/mango/components/shape"
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


	// circle := prefabs.NewCirclePrefab()
 //  circleEntity := scene.AddPrefab(circle)
 //  circleEntity.Tranform().Position.X = 300



	mango.CreateWindow(1300, 800, "Retained Mode Rendering", false)

	rect := prefabs.NewSquarePrefab()
  rect.Renderer.SetTexture("man.png")
  rect.Renderer.Shape = &shape.Rect{
    Width: 100,
    Height: 150,
  }
  scene.AddPrefab(rect)

	mango.Start()

}
