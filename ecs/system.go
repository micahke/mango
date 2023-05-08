package ecs


// Base implementation of a system
type System interface {

  // Initialize the system
  Init()

  // Updates the system every frame
  Tick()
}


