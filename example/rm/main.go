package main

import (
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/core"
)





func main() {
  
  mango.Init(core.RENDER_MODE_DEFAULT)

  // scene setup
  scene := mango.CreateScene()
  mango.SetScene(scene)

  mango.CreateWindow(800, 600, "Retained Mode Rendering", true)
  mango.Start()

}
