package system

import (
	"github.com/seppedelanghe/mizu/examples/particles/component"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Acceleration struct {
	*component.Root
	*component.Vel
	*component.Accel
}

func (a *Acceleration) Update(_ engine.World) {
	if a.Root.Enabled {
		return
	}

	// Increase velocity
	a.Vel.L += a.Accel.O
	a.Vel.M += a.Accel.P
}
