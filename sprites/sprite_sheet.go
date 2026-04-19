package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	image *ebiten.Image
}

func (s *SpriteSheet) Sprite(r image.Rectangle) *Sprite {
	// TODO
	return nil
}
