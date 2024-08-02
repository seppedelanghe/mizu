package component

import (
	"github.com/seppedelanghe/mizu/examples/platformer/helper/graphics"
)

type Sprite struct {
	Frameset *graphics.Frameset
}

func NewSprite(frameset *graphics.Frameset) Sprite {
	return Sprite{Frameset: frameset}
}
