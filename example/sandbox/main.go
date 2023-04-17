package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/core"
	"github.com/micahke/infinite-universe/mango/logging"
)

func init() {
  runtime.LockOSThread()
}

func main() {

  mango.Init(core.RENDER_MODE_IM)
  mango.CreateWindow(800, 600, "Mango Sandbox", false)

  mango.IM.ConnectScene(&Sandbox{})
  mango.Start()

}



type Sandbox struct {}



func (sb *Sandbox) Init() {

}


func (sb *Sandbox) Update(deltaTime float64) {
  logging.Log(deltaTime)
}


func (sb *Sandbox) Draw() {

  mango.IM.DrawText(fmt.Sprint(math.Floor(core.Timer.FPS()), " FPS"), 250, 390)

}
