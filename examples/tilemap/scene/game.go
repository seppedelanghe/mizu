package scene

import (
	"github.com/seppedelanghe/mizu/examples/tilemap/assets"
	"github.com/seppedelanghe/mizu/examples/tilemap/component"
	"github.com/seppedelanghe/mizu/examples/tilemap/entity"
	"github.com/seppedelanghe/mizu/examples/tilemap/helper/random"
	"github.com/seppedelanghe/mizu/examples/tilemap/system"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(component.Pos{}, component.Zoom{}, component.Solid{}, component.Construct{}, component.Sprite{})

	// Creates empty tiles for the tilemap with the prepared size.
	construct := component.NewConstruct(8, 6)
	w.AddSystems(system.NewRendering(construct.Width, construct.Height))
	w.AddSystems(system.NewFocus(), system.NewBuilding(), system.NewConnection(), system.NewDebug())
	for y := 0; y < construct.Height; y++ {
		for x := 0; x < construct.Width; x++ {
			w.AddEntities(&entity.Tile{
				Pos:    component.NewPosI(x, y),
				Solid:  component.NewSolid(random.Chance(0.6)),
				Sprite: component.NewSprite(assets.Images["road_0"]),
			})
		}
	}
	w.AddEntities(&entity.Construct{Construct: construct}, &entity.Camera{})
}
