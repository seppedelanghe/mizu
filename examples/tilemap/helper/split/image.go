package split

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seppedelanghe/mizu/examples/tilemap/helper/geo"
)

func Single(source *ebiten.Image, x, y, w, h int) *ebiten.Image {
	return source.SubImage(geo.Rect(x, y, w, h)).(*ebiten.Image)
}
