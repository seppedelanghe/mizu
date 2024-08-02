package system

import (
	"github.com/seppedelanghe/mizu/examples/bunnymark/component"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Gravity struct {
	*component.Velocity
	*component.Gravity
}

func (g *Gravity) Update(_ engine.World) {
	g.Velocity.Y += g.Gravity.Value
}
