package main

import (
	"runtime"

	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/im"
	"github.com/micahke/mango/util/color"
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
}


func (sb *Sandbox) Draw() {

  quad := im.Quad{
    X: 100,
    Y: 100,
    Width: 100,
    Height: 100,
    Color: color.WHITE,
  }
  quad2 := im.Quad{
    X: 400,
    Y: 400,
    Width: 100,
    Height: 100,
    Color: color.ELECTRON_BLUE,
  }
  mango.IM.DrawQuad(quad)
  mango.IM.DrawQuad(quad2)

}
