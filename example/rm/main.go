package main

import (
	"fmt"

	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/logging"
)


func main() {
  
  mango.Init(core.RENDER_MODE_DEFAULT)

  // scene setup
  scene := mango.CreateScene()
  mango.SetScene(scene)

  player := scene.CreateEntity("test")
  player.Tranform().Position.X = 100

  player.AddComponent(&CustomComponent{})

  // addEntities(scene, 32)

  logging.Log(player.Tranform().Position)

  mango.CreateWindow(800, 600, "Retained Mode Rendering", true)
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
