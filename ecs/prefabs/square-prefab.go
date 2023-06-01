package ecs

import (
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
)

type SquarePrefab struct {
	Renderer *components.PrimitiveRenderer
	Name     string
}

func NewSquarePrefab() *SquarePrefab {
	prefab := new(SquarePrefab)
	prefab.Name = "square"
	prefab.Renderer = &components.PrimitiveRenderer{}
	prefab.Renderer.SetShape(components.SHAPE_RECT)
	return prefab
}

func (prefab *SquarePrefab) GetPrefabName() string {
	return prefab.Name
}

func (prefab *SquarePrefab) GetPrefabComponents() []ecs.Component {
	return []ecs.Component{prefab.Renderer}
}
