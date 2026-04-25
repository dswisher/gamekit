package sprites

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

// Drawable represents something that can be drawn with an origin.
type Drawable interface {
	Origin() (x, y float64)
}

// drawImage draws an image to the screen using the specified options.
// The drawable provides default origin values that can be overridden by opts.
func drawImage(drawable Drawable, screen, img *ebiten.Image, opts DrawOptions) {
	eopts := &colorm.DrawImageOptions{}

	// Get default origin from the drawable
	originX, originY := drawable.Origin()

	// Determine effective origin (use DrawOptions if specified, otherwise use defaults)
	ox, oy := opts.OriginX, opts.OriginY
	if math.IsNaN(ox) {
		ox = originX
	}
	if math.IsNaN(oy) {
		oy = originY
	}

	// Apply transformations in order: translate to origin -> scale/rotate -> translate to position
	// First, translate to the origin point (negative because we want to rotate/scale around that point)
	eopts.GeoM.Translate(-ox, -oy)

	// Apply scaling if different from 1
	if opts.ScaleX != 1 || opts.ScaleY != 1 {
		eopts.GeoM.Scale(opts.ScaleX, opts.ScaleY)
	}

	// Apply rotation if non-zero
	if opts.Rotate != 0 {
		eopts.GeoM.Rotate(opts.Rotate)
	}

	// Translate to the final position, accounting for the origin offset
	eopts.GeoM.Translate(opts.X+ox, opts.Y+oy)

	// Apply blend mode
	eopts.Blend = opts.Blend

	colorm.DrawImage(screen, img, opts.ColorM, eopts)
}
