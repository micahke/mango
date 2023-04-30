package main

import (
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

  logging.Log(player.Tranform().Position)

  mango.CreateWindow(800, 600, "Retained Mode Rendering", true)
  mango.Start()

}
