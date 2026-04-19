// Package sprites provides sprite sheet management and animation for Ebitengine.
//
// The sprites package offers tools for loading, managing, and rendering sprite
// sheets and individual sprites. It supports loading images from various sources
// including the filesystem and embedded assets, and provides flexible drawing
// options for positioning and transforming sprites.
//
// Basic usage:
//
//	sprite := sprites.NewSprite(img)
//	opts := sprites.DrawOpts(100, 100)
//	sprite.Draw(screen, opts)
//
// For loading images from embedded filesystems:
//
//	//go:embed assets/*.png
//	var assets embed.FS
//	img, err := sprites.LoadImageFromFS(assets, "assets/player.png")
package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite represents a single drawable image that can be rendered to the screen.
// A Sprite wraps an ebiten.Image and provides a simplified API for drawing
// with positioning and transformation options.
type Sprite struct {
	image *ebiten.Image
}

// NewSprite creates a new Sprite from the specified ebiten.Image.
// The image can be loaded using LoadImageFromFS or created directly.
//
// Example:
//
//	img, _ := sprites.LoadImageFromFS(assets, "player.png")
//	sprite := sprites.NewSprite(img)
func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{image: img}
}

// Draw renders the sprite to the provided screen image using the specified options.
// The opts parameter controls the position and other transformations applied
// to the sprite when drawing.
//
// Example:
//
//	opts := sprites.DrawOpts(100, 100) // position at (100, 100)
//	sprite.Draw(screen, opts)
func (spr *Sprite) Draw(screen *ebiten.Image, opts *DrawOptions) {
	// TODO: properly handle the draw options
	// subImage := spr.subImages[index]
	// screen.DrawImage(subImage, op)

	eopts := &ebiten.DrawImageOptions{}
	eopts.GeoM.Translate(opts.X, opts.Y)
	screen.DrawImage(spr.image, eopts)
}
