package main

import (
	"runtime"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/core"
	"github.com/micahke/infinite-universe/src/scene"
)

func init() {
	runtime.LockOSThread()
}

var WINDOW_WIDTH int = 1300
var WINDOW_HEIGHT int = 800

func main() {

	mango.Init(core.RENDER_MODE_IM)
	mango.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Infinite Universe")
	mango.IM.ConnectScene(&scene.PlanetMap{WINDOW_WIDTH: WINDOW_WIDTH, WINDOW_HEIGHT: WINDOW_HEIGHT})
	mango.Start()

}
