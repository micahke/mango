package core

import "github.com/micahke/mango/ecs"

type Scene struct {
	ecs *ecs.ECS
}

func NewScene() *Scene {

	scene := new(Scene)
	scene.ecs = new(ecs.ECS)

	return scene

}

// Get the scene's Entity Component System
func (scene *Scene) ECS() *ecs.ECS {
	return scene.ecs
}

func (scene *Scene) Update() {

	// update the entity component system
	scene.ecs.Update()

}

// Creates a new entity and adds it to the scene
func (scene *Scene) CreateEntity(id string) *ecs.Entity {

	return scene.ecs.CreateEntity(id)

}

func (scene *Scene) AddPrefab(prefab ecs.Prefab) *ecs.Entity {
	entity := scene.ecs.CreateEntity(prefab.GetPrefabName())
	for _, component := range prefab.GetPrefabComponents() {
		entity.AddComponent(component)
	}
  return entity
}
