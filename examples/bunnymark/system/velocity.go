package system

import (
	"github.com/seppedelanghe/mizu/examples/bunnymark/component"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Velocity struct {
	*component.Position
	*component.Velocity
}

func (v *Velocity) Update(_ engine.World) {
	v.Position.X += v.Velocity.X
	v.Position.Y += v.Velocity.Y
}
