package ecs

import (
	"github.com/micahke/mango/components"
	"github.com/micahke/mango/components/shape"
)

type SquarePrefab struct {
	renderer *components.PrimitiveRenderer
	name string
}

func NewSquarePrefab() *SquarePrefab{
	prefab := new(SquarePrefab);
	prefab.name = "square"
	prefab.renderer = 

	// 
	rect = prefab.renderer.Shape.(*shape.Rect)
	rect.Width = 10
}