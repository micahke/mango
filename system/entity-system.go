package system

import "github.com/micahke/mango/ecs"

type EntitySystem struct {
  Entities *[]*ecs.Entity
}



func (system *EntitySystem) Init() {}


func (system *EntitySystem) Tick() {

  for _, entity := range(*system.Entities) {
    entity.Update()
  }

}
