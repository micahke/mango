package system

import (
	"github.com/micahke/mango/ecs"
)

type EntitySystem struct {
	Entities *[]*ecs.Entity
}

func (system *EntitySystem) InitializeSystem() {}

func (system *EntitySystem) UpdateSystem() {

	for _, entity := range *system.Entities {
		entity.Update()
	}

}
