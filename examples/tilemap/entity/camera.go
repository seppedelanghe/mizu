package entity

import "github.com/seppedelanghe/mizu/examples/tilemap/component"

// An object that stores a point and distance for focusing on other objects.

type Camera struct {
	component.Pos
	component.Zoom
}
