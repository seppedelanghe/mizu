package scene

import (
	"time"

	"github.com/seppedelanghe/mizu/examples/bunnymark/assets"
	"github.com/seppedelanghe/mizu/examples/bunnymark/component"
	"github.com/seppedelanghe/mizu/examples/bunnymark/entity"
	"github.com/seppedelanghe/mizu/examples/bunnymark/helper"
	"github.com/seppedelanghe/mizu/examples/bunnymark/system"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Position{}, component.Velocity{}, component.Gravity{},
		component.Sprite{}, component.Hue{}, component.Settings{},
	)
	w.AddSystems(
		&system.Background{}, &system.Velocity{}, &system.Gravity{},
		&system.Bounce{}, &system.Render{}, &system.Metrics{}, &system.Spawn{},
	)
	w.AddEntities(
		&entity.Settings{
			Settings: component.Settings{
				Ticker:   time.NewTicker(500 * time.Millisecond),
				Gpu:      helper.GpuInfo(),
				Tps:      helper.NewPlot(20, 60),
				Fps:      helper.NewPlot(20, 60),
				Objects:  helper.NewPlot(20, 60000),
				Sprite:   assets.Bunny,
				Colorful: false,
				Amount:   100,
			},
		},
	)
}
