package system

import (
	"github.com/seppedelanghe/mizu/examples/particles/component"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Age struct {
	*component.Root
	*component.Life
}

func (v *Age) Update(_ engine.World) {
	if v.Root.Enabled {
		return
	}

	// Increase age to death.
	if v.Current >= v.Total {
		return
	}
	v.Current++
}
