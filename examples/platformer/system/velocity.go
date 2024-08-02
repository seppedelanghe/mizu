package system

import (
	"github.com/seppedelanghe/mizu/examples/platformer/component"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Velocity struct {
	*component.Pos
	*component.Vel
}

func NewVelocity() *Velocity {
	return &Velocity{}
}

func (v *Velocity) Update(_ engine.World) {
	// Increase position.
	v.Pos.X += v.Vel.L
	v.Pos.Y += v.Vel.M
}
