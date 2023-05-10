package main

import (
	"fmt"

	"github.com/micahke/mango"
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/components/shape"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/logging"
	// "github.com/micahke/mango/logging"
)


func main() {
  
  mango.Init(core.RENDER_MODE_DEFAULT)

  core.Settings.SCENE_EDITOR_STARTUP = true
  core.Settings.CONSOLE_ON_STARTUP = true

  // scene setup
  scene := mango.CreateScene()
  mango.SetScene(scene)

  player := scene.CreateEntity("player")
  player.Tranform().Position.X = 100

  shapeComponent := &components.Shape2DComponent{}
  rect := shape.Rect{
    Width: 100,
    Height: 100,
  }
  shapeComponent.SetShape(&rect)
  player.AddComponent(shapeComponent)

  player.AddComponent(&components.PrimitiveRenderer{})
  player.AddComponent(&components.SampleComponent{})
  // player.AddComponent(&CustomComponent{})

  addEntities(scene, 10)

  logging.Log(player.Tranform().Position)

  mango.CreateWindow(1300, 800, "Retained Mode Rendering", false)
  mango.Start()

}


func addEntities(scene*core.Scene, num int) {

  for i := 0; i < num; i++ {
    scene.CreateEntity(fmt.Sprint("entity", i))
  }

}



type CustomComponent struct {}


func (cc *CustomComponent) Init() {}

func (cc *CustomComponent) Update() {
}

func (cc *CustomComponent) GetComponentName() string {
  return "Custom Component"
}
