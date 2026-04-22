package sprites

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

// DrawOptions specifies the parameters for drawing a sprite.
// It controls the position, rotation, scaling, and origin for transformations.
//
// For simple use cases, use DrawAt to start building options:
//
//	// Simple positioning
//	sprite.Draw(screen, sprites.DrawAt(100, 200))
//
//	// With rotation and scaling (method chaining)
//	sprite.Draw(screen, sprites.DrawAt(100, 200).WithRotation(math.Pi/4).WithScale(2.0))
//
// The OriginX and OriginY fields control the center of rotation and scaling.
// If set to math.NaN(), the sprite's default origin will be used.
type DrawOptions struct {
	// X is the horizontal screen position where the sprite will be drawn.
	X float64
	// Y is the vertical screen position where the sprite will be drawn.
	Y float64
	// Rotate is the rotation angle in radians.
	Rotate float64
	// ScaleX is the horizontal scale factor (1.0 = original size).
	ScaleX float64
	// ScaleY is the vertical scale factor (1.0 = original size).
	ScaleY float64
	// OriginX is the X coordinate of the rotation/scaling origin in pixels.
	// Set to math.NaN() to use the sprite's default origin.
	OriginX float64
	// OriginY is the Y coordinate of the rotation/scaling origin in pixels.
	// Set to math.NaN() to use the sprite's default origin.
	OriginY float64

	// ColorM is the color matrix for color modulation.
	ColorM colorm.ColorM

	// Blend is the blending mode for drawing.
	Blend ebiten.Blend
}

// DrawAt creates a new DrawOptions positioned at (x, y).
// Use the With* methods to add rotation, scaling, and other transformations.
//
// Example:
//
//	// Simple positioning
//	sprite.Draw(screen, sprites.DrawAt(100, 100))
//
//	// With rotation
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithRotation(math.Pi/4))
//
//	// With uniform scaling
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithScale(2.0))
//
//	// With non-uniform scaling
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithScaleXY(2.0, 0.5))
//
//	// Chaining multiple transformations
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithRotation(math.Pi/4).WithScale(2.0))
func DrawAt(x, y float64) DrawOptions {
	return DrawOptions{
		X:       x,
		Y:       y,
		ScaleX:  1,
		ScaleY:  1,
		OriginX: math.NaN(),
		OriginY: math.NaN(),
	}
}

// WithRotation returns a new DrawOptions with the specified rotation.
// The rotation is in radians.
func (opts DrawOptions) WithRotation(rotate float64) DrawOptions {
	opts.Rotate = rotate
	return opts
}

// WithScale returns a new DrawOptions with uniform scaling.
// The scale factor applies to both X and Y dimensions.
func (opts DrawOptions) WithScale(scale float64) DrawOptions {
	opts.ScaleX = scale
	opts.ScaleY = scale
	return opts
}

// WithScaleXY returns a new DrawOptions with non-uniform scaling.
// scaleX and scaleY can be specified independently.
func (opts DrawOptions) WithScaleXY(scaleX, scaleY float64) DrawOptions {
	opts.ScaleX = scaleX
	opts.ScaleY = scaleY
	return opts
}

// WithOrigin returns a new DrawOptions with the specified origin.
// The origin is in pixel coordinates relative to the top-left of the sprite.
// This overrides the sprite's default origin for this draw call.
func (opts DrawOptions) WithOrigin(originX, originY float64) DrawOptions {
	opts.OriginX = originX
	opts.OriginY = originY
	return opts
}

// WithBlend returns a new DrawOptions with the specified blend mode.
// Common blend modes include:
//   - ebiten.BlendSourceOver (default): Normal alpha blending
//   - ebiten.BlendLighter: Additive blending for glow effects
//   - ebiten.BlendCopy: Replace destination completely
func (opts DrawOptions) WithBlend(blend ebiten.Blend) DrawOptions {
	opts.Blend = blend
	return opts
}

// WithColorM returns a new DrawOptions with the specified color matrix.
// The color matrix transforms the sprite's colors when drawn, enabling effects
// like tinting, brightness adjustment, grayscale, and more.
//
// Use colorm.ColorM's methods to build transformations:
//   - Scale(r, g, b, a): Scale each color channel (1.0 = unchanged)
//   - Translate(r, g, b, a): Add to each color channel
//   - ChangeHSV(hue, saturation, value): HSV adjustments
//
// Example:
//
//	// Tint sprite red (keep red, remove green and blue)
//	cm := colorm.ColorM{}
//	cm.Scale(1.0, 0.0, 0.0, 1.0)
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithColorM(cm))
//
//	// Make sprite 50% darker
//	cm := colorm.ColorM{}
//	cm.Scale(0.5, 0.5, 0.5, 1.0)
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithColorM(cm))
//
//	// Flash white (add white to all channels)
//	cm := colorm.ColorM{}
//	cm.Translate(0.5, 0.5, 0.5, 0.0)
//	sprite.Draw(screen, sprites.DrawAt(100, 100).WithColorM(cm))
func (opts DrawOptions) WithColorM(cm colorm.ColorM) DrawOptions {
	opts.ColorM = cm
	return opts
}
