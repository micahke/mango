package ecs

import "github.com/micahke/infinite-universe/mango/logging"

// Base implementation of a system
type System interface {

  // Initialize the system
  Init()

  // Updates the system every frame
  Tick()
}


type EntitySystem struct {

  

}



func (system *EntitySystem) Init() {

  

}


func (system *EntitySystem) Tick() {

  logging.DebugLog("Ticking from entity system")

}
