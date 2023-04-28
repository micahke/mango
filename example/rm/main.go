package main

import (
	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
)


func main() {
  
  mango.Init(core.RENDER_MODE_DEFAULT)

  // scene setup
  scene := mango.CreateScene()
  mango.SetScene(scene)

  scene.CreateEntity("test")

  mango.CreateWindow(800, 600, "Retained Mode Rendering", true)
  mango.Start()

}
