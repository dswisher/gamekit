package sprites

// import "github.com/hajimehoshi/ebiten/v2"

// DrawOptions specifies the parameters for drawing a sprite.
// It controls the position and will eventually support rotation,
// scaling, color modulation, and blending modes.
//
// For simple use cases, the DrawOpts function provides a convenient
// way to create DrawOptions with just X and Y coordinates.
//
// Example using DrawOpts shortcut:
//
//	opts := sprites.DrawOpts(100, 200)
//	sprite.Draw(screen, opts)
//
// Example creating DrawOptions directly:
//
//	opts := &sprites.DrawOptions{X: 100, Y: 200}
//	sprite.Draw(screen, opts)
//
// Planned features:
//   - Rotation (in degrees or radians)
//   - Scaling (ScaleX, ScaleY)
//   - Origin point for transformations
//   - Color modulation using the colorm package
//   - Custom blend modes
type DrawOptions struct {
	// X is the horizontal screen position where the sprite will be drawn.
	X float64
	// Y is the vertical screen position where the sprite will be drawn.
	Y float64

	// Rotate           float64
	// ScaleX, ScaleY   float64
	// OriginX, OriginY float64

	// TODO: deprecated - use colorm package instead
	// ColorM           ebiten.ColorM

	// TODO: deprecated - use Blend instead
	// CompositeMode    ebiten.CompositeMode
}

// DrawOpts creates a new DrawOptions with the specified X and Y coordinates.
// This is a convenience function for the common case where you only need
// to specify the position without rotation, scaling, or other effects.
//
// For more advanced options, create DrawOptions directly:
//
//	opts := &sprites.DrawOptions{X: x, Y: y}
//
// Parameters:
//   - x: The horizontal screen position
//   - y: The vertical screen position
//   - args: Reserved for future use (rotation, scale, origin)
//
// Returns a pointer to a new DrawOptions instance.
//
// Example:
//
//	sprite.Draw(screen, sprites.DrawOpts(100, 100))
func DrawOpts(x, y float64, args ...float64) *DrawOptions {
	// r, sx, sy, ox, oy := 0., 1., 1., 0., 0.
	// switch len(args) {
	// case 5:
	// 	oy = args[4]
	// 	fallthrough
	// case 4:
	// 	ox = args[3]
	// 	fallthrough
	// case 3:
	// 	sy = args[2]
	// 	fallthrough
	// case 2:
	// 	sx = args[1]
	// 	fallthrough
	// case 1:
	// 	r = args[0]
	// }

	return &DrawOptions{
		X: x,
		Y: y,
		// Rotate:        r,
		// ScaleX:        sx,
		// ScaleY:        sy,
		// OriginX:       ox,
		// OriginY:       oy,
		// ColorM:        ebiten.ColorM{},
		// CompositeMode: ebiten.CompositeModeSourceOver,
	}
}
