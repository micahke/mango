package ecs

import (
	"fmt"

	"github.com/micahke/infinite-universe/mango/logging"
)


type ECS struct {

  entities []*Entity
  systems []interface{}

}

// Run every frame
// Handles the logic for all game systems and entities
func (ecs *ECS) Update() {

  // Update the entity's logic
  // This might have to be (re)moved
  ecs.updateEntities()

  // Update various systems
  ecs.tickSystems()

}

// Creates and entity, adds it to the scene
// You can pass a nil id
func (ecs *ECS) CreateEntity(id string) *Entity {

  entity := new(Entity)
  ecs.addEntityToECS(entity)

  return entity

}

func (ecs *ECS) GetEntity(id string) *Entity {

  for _, entity := range(ecs.entities) {
    if entity.id == id {
      return entity
    }
  }
  return nil

}


func (ecs *ECS) AddSystem(system interface{}) (*System, error) {

  // Check if the system implements the system interface
  if _, ok := system.(System); ok {

    ecs.addSystem(system.(*System))
    return system.(*System), nil
  }

  errorMessage := "The system you added does not implement interface System"
  logging.DebugLog(errorMessage)

  return nil, fmt.Errorf(errorMessage)
    
}


// ++++++++++++++++ PRIVATE ++++++++++++++++++++=


func (ecs *ECS) updateEntities() {
  
  // for _, entity := range(ecs.entities) {
  //
  //   // Update the entity
  //
  // }

}


// Handles the updating of various systems every frame
func (ecs *ECS) tickSystems() {

  for _, system := range(ecs.systems) {

    system.(System).Tick()
      
  }

}

func (ecs *ECS) addEntityToECS(entity *Entity) {

  // Initialize the entity with a transform component
  entity.addTranformComponent()

  ecs.entities = append(ecs.entities, entity)

}


func (ecs *ECS) addSystem(system *System) {

  ecs.systems = append(ecs.systems, system)

}

