package sprites

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

// Sprite represents a single drawable image that can be rendered to the screen.
// A Sprite wraps an ebiten.Image and provides a simplified API for drawing
// with positioning and transformation options.
type Sprite struct {
	image          *ebiten.Image
	originX        float64
	originY        float64
	width, height  float64
}

// NewSprite creates a new Sprite from the specified ebiten.Image.
// The image can be loaded using LoadImageFromFS or created directly.
//
// Example:
//
//	img, _ := sprites.LoadImageFromFS(assets, "player.png")
//	sprite := sprites.NewSprite(img)
func NewSprite(img *ebiten.Image) *Sprite {
	bounds := img.Bounds()
	return &Sprite{
		image:   img,
		originX: 0,
		originY: 0,
		width:   float64(bounds.Dx()),
		height:  float64(bounds.Dy()),
	}
}

// SetOrigin sets the default origin for rotation and scaling transformations.
// The origin is specified in pixel coordinates relative to the top-left corner
// of the sprite. This value will be used when DrawOptions.OriginX/OriginY are NaN.
//
// Example:
//
//	// Set rotation center to the center of a 64x64 sprite
//	sprite.SetOrigin(32, 32)
func (spr *Sprite) SetOrigin(x, y float64) {
	spr.originX = x
	spr.originY = y
}

// Origin returns the current origin values for this sprite.
func (spr *Sprite) Origin() (x, y float64) {
	return spr.originX, spr.originY
}

// Draw renders the sprite to the provided screen image using the specified options.
// The opts parameter controls the position, rotation, scaling, and other transformations
// applied to the sprite when drawing.
//
// Example:
//
//	// Simple positioning
//	sprite.Draw(screen, sprites.DrawAt(100, 100))
//
//	// With rotation and scaling using method chaining
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithRotation(math.Pi/4).WithScale(2.0))
func (spr *Sprite) Draw(screen *ebiten.Image, opts DrawOptions) {
	drawImage(screen, spr.image, opts, spr.originX, spr.originY)
}

// drawImage draws an image to the screen using the specified options.
// originX and originY are default origin values that can be overridden by opts.
func drawImage(screen, img *ebiten.Image, opts DrawOptions, originX, originY float64) {
	eopts := &colorm.DrawImageOptions{}

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
