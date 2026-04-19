package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	image *ebiten.Image
}

// TODO: add:
//    sprites.NewSheet(img *ebiten.Image, layout Layout)
//       -> sheet := sprites.NewSheet(playerImg, gridLayout)
//
//    sprites.LoadSheetFromMeta(metaPath string)
//       -> sheet := sprites.LoadSheetFromMeta("assets/player.json")

// Create a sprite from the sprite sheet and give rectangle
func (s *SpriteSheet) Sprite(r image.Rectangle) *Sprite {
	// TODO
	return nil
}
