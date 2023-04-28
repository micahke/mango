package main

import (
	"runtime"

	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
	"github.com/micahke/src/scene/galaxymap"
)

func init() {
	runtime.LockOSThread()
}

var WINDOW_WIDTH int = 1300
var WINDOW_HEIGHT int = 800

func main() {
	mango.Init(core.RENDER_MODE_IM)
	mango.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Infinite Universe", false)

	// mango.IM.ConnectScene(&scene.PlanetMap{WINDOW_WIDTH: WINDOW_WIDTH, WINDOW_HEIGHT: WINDOW_HEIGHT})
	mango.IM.ConnectScene(&galaxymap.GalaxyMap{})
	mango.Start()

}
