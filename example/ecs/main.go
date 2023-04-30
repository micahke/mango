package main

import (
	"runtime"

	"github.com/micahke/mango"
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/ecs"
	"github.com/micahke/mango/logging"
)

func init() {
  runtime.LockOSThread()
}

func main() {
  
  mango.Init(core.RENDER_MODE_IM)
  mango.CreateWindow(800, 600, "ECS Test", true)

  mango.IM.ConnectScene(&Game{})
  mango.Start()

}


type Game struct {

  ecs *ecs.ECS

  player *ecs.Entity

}


func (game *Game) Init() {

  game.ecs = &ecs.ECS{}
  
  game.player = game.ecs.CreateEntity("player")
  game.player.Tranform().Position.X = 100
  game.player.Tranform().Position.Y = 100

}

func (game *Game) Update() {

  tranform := game.player.Tranform()
  logging.Log(tranform)


}

func (game *Game) Draw() {

}
