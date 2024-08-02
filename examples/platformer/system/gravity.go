package system

import (
	"github.com/seppedelanghe/mizu/examples/platformer/component"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Gravity struct {
	*component.Vel
	*component.Gravity
}

func NewGravity() *Gravity {
	return &Gravity{}
}

func (g *Gravity) Update(_ engine.World) {
	// Increase vertical speed.
	g.Vel.M += g.Gravity.Value
}
