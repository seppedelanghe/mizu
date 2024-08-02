package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seppedelanghe/mizu/examples/bunnymark/assets"
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type Background struct{}

func (b *Background) Draw(_ engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)
}
