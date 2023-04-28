package main

import (
	"inifinite-universe/scene/galaxymap"
	"runtime"

	"github.com/micahke/mango"
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
