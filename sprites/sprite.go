package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite represents a single drawable image that can be rendered to the screen.
// A Sprite wraps an ebiten.Image and provides a simplified API for drawing
// with positioning and transformation options.
type Sprite struct {
	image            *ebiten.Image
	originX, originY float64
	width, height    float64
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
	drawImage(spr, screen, spr.image, opts)
}
