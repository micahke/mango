package main

import "github.com/micahke/mango/core"


func main() {

  Init(core.RENDER_MODE_DEFAULT)
  
  scene := CreateScene()
  SetScene(scene)

  CreateWindow(800, 600, "Mango Sandbox", true)

  Start()


}
