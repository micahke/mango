package ecs

type Prefab interface {
	GetPrefabName() string
	GetPrefabComponents() []Component
}
