package prefabs

import (
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/ecs"
)

type CirclePrefab struct {
	Name     string
	Renderer *components.PrimitiveRenderer
}

func NewCirclePrefab() *CirclePrefab {
	prefab := new(CirclePrefab)
	prefab.Name = "circle"
	prefab.Renderer = &components.PrimitiveRenderer{}
	prefab.Renderer.SetShape(components.SHAPE_CIRCLE)
	return prefab
}

func (prefab *CirclePrefab) GetPrefabName() string {
	return prefab.Name
}

func (prefab *CirclePrefab) GetPrefabComponents() []ecs.Component {
	return []ecs.Component{prefab.Renderer}
}
