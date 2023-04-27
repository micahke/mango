package core

import "github.com/micahke/infinite-universe/mango/ecs"


type Scene struct {

  ecs *ecs.ECS

}


func NewScene() *Scene {

  scene := new(Scene)
  scene.ecs = new(ecs.ECS)

  return scene

}


// Creates a new entity and adds it to the scene
func (scene *Scene) CreateEntity() *ecs.Entity {
  
  return scene.ecs.CreateEntity()

}

