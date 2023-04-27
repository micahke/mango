package main

import (
	"runtime"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/core"
	"github.com/micahke/infinite-universe/mango/ecs"
	"github.com/micahke/infinite-universe/mango/util/color"
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
  
  game.player = game.ecs.CreateEntity()
  game.player.Tranform().X = 100
  game.player.Tranform().Y = 100

}

func (game *Game) Update(deltaTime float64) {

}

func (game *Game) Draw() {

  mango.IM.FillRect(game.player.Tranform().X, game.player.Tranform().Y, 100, 100, color.ELECTRON_BLUE)

}
