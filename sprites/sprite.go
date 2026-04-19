// Package sprites provides sprite sheet management and animation for Ebitengine.
package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image *ebiten.Image
}

// Create a simple sprite from the specified image
func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{image: img}
}

// Draw the sprite with the specified options
func (spr *Sprite) Draw(screen *ebiten.Image, opts *DrawOptions) {
	// TODO: properly handle the draw options
	// subImage := spr.subImages[index]
	// screen.DrawImage(subImage, op)

	eopts := &ebiten.DrawImageOptions{}
	eopts.GeoM.Translate(opts.X, opts.Y)
	screen.DrawImage(spr.image, eopts)
}
