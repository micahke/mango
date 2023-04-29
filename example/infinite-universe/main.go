package main

import (
	"github.com/micahke/mango/example/infinite-universe/scene/galaxymap"
	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
)


var WINDOW_WIDTH int = 1300
var WINDOW_HEIGHT int = 800

func main() {
	mango.Init(core.RENDER_MODE_IM)
	mango.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Infinite Universe", false)

	mango.IM.ConnectScene(&galaxymap.GalaxyMap{})
	mango.Start()

}