package ecs


type ECS struct {

  entities []*Entity

}


// Creates and entity, adds it to the scene
func (ecs *ECS) CreateEntity() *Entity {

  entity := new(Entity)
  ecs.addEntityToECS(entity)

  return entity

}



func (ecs *ECS) addEntityToECS(entity *Entity) {

  // Initialize the entity with a transform component
  entity.addTranformComponent()

  ecs.entities = append(ecs.entities, entity)

}


