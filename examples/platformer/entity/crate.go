package entity

import "github.com/seppedelanghe/mizu/examples/platformer/component"

// The object that can fall.

type Crate struct {
	component.Pos     // Index for the tile.
	component.Vel     // Vel for the tile.
	component.Size    // Current size in tiles
	component.Solid   // Collision group.
	component.Gravity // Current gravity.
	component.Sprite  // Current image.
}
