// Package sprites provides sprite sheet management and animation for Ebitengine.
package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image *ebiten.Image
}

// Create a simple sprite from the specified image
// TODO: I don't like this name, as it isn't clear what is being created
func FromImage(img *ebiten.Image) *Sprite {
	return &Sprite{image: img}
}

// Draw the sprite with the specified options
func (spr *Sprite) Draw(screen *ebiten.Image, opts *DrawOptions) {
	// TODO
	// subImage := spr.subImages[index]
	// screen.DrawImage(subImage, op)

	eopts := &ebiten.DrawImageOptions{}
	eopts.GeoM.Translate(opts.X, opts.Y)
	screen.DrawImage(spr.image, eopts)
}
