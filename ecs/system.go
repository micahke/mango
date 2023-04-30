package ecs


// Base implementation of a system
type System interface {

  // Initialize the system
  Init()

  // Updates the system every frame
  Tick()
}


type EntitySystem struct {
  Entities *[]*Entity
}



func (system *EntitySystem) Init() {}


func (system *EntitySystem) Tick() {

  for _, entity := range(*system.Entities) {

    entity.Update()


  }

}
