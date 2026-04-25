package sprites

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Animation represents a sequence of frames that can be played back at a specified FPS.
// The animation loops continuously and advances frames automatically when Update() is called.
type Animation struct {
	Frames             []*ebiten.Image // Animation frames
	FPS                float64         // Animation playback speed (Frames Per Second)
	originX, originY   float64         // Default origin for rotation and scaling
	currentIndex       int             // Current frame index
	tick               float64         // Accumulated time tick for frame advancement
}

// NewAnimation creates a new Animation from the specified image and frame rectangles.
// The frames are extracted as sub-images from the provided image using the rectangles.
// FPS is hardcoded to 12 for now; future versions will support configuration via WithFPS().
//
// Example:
//
//	grid := sprites.NewGridLocator(32, 32)
//	rects := grid.GetRowRects(0, 0, 8) // Get 8 frames from row 0
//	anim := sprites.NewAnimation(img, rects)
func NewAnimation(img *ebiten.Image, rects []image.Rectangle) *Animation {
	frames := make([]*ebiten.Image, len(rects))
	for i, rect := range rects {
		frames[i] = img.SubImage(rect).(*ebiten.Image)
	}

	return &Animation{
		Frames:       frames,
		FPS:          12, // Hardcoded default FPS
		originX:      0,
		originY:      0,
		currentIndex: 0,
		tick:         0,
	}
}

// Update advances the animation to the next frame if sufficient time has passed.
// This method should be called once per frame in your game's Update() loop.
// The animation loops continuously from the last frame back to the first.
//
// Frame timing is based on a 60 TPS (ticks per second) assumption.
// At the default 12 FPS, the animation advances one frame every 5 ticks.
func (anim *Animation) Update() {
	if len(anim.Frames) == 0 || anim.FPS <= 0 {
		return
	}

	// Advance tick based on FPS (assuming 60 TPS)
	anim.tick += anim.FPS / 60.0
	anim.currentIndex = int(math.Floor(anim.tick))

	// Loop back to beginning when we reach the end
	if anim.currentIndex >= len(anim.Frames) {
		anim.tick = 0
		anim.currentIndex = 0
	}
}

// Draw renders the current animation frame to the provided screen image.
// The opts parameter controls position, rotation, scaling, and other transformations.
//
// Example:
//
//	// Simple positioning
//	anim.Draw(screen, sprites.DrawAt(100, 100))
//
//	// With scaling
//	anim.Draw(screen, sprites.DrawAt(100, 100).WithScale(2.0))
// SetOrigin sets the default origin for rotation and scaling transformations.
// The origin is specified in pixel coordinates relative to the top-left corner
// of the animation frame. This value will be used when DrawOptions.OriginX/OriginY are NaN.
//
// Example:
//
//	// Set rotation center to the center of a 32x32 animation frame
//	anim.SetOrigin(16, 16)
func (anim *Animation) SetOrigin(x, y float64) {
	anim.originX = x
	anim.originY = y
}

// Origin returns the current origin values for this animation.
func (anim *Animation) Origin() (x, y float64) {
	return anim.originX, anim.originY
}

// Draw renders the current animation frame to the provided screen image.
// The opts parameter controls position, rotation, scaling, and other transformations.
//
// Example:
//
//	// Simple positioning
//	anim.Draw(screen, sprites.DrawAt(100, 100))
//
//	// With scaling
//	anim.Draw(screen, sprites.DrawAt(100, 100).WithScale(2.0))
func (anim *Animation) Draw(screen *ebiten.Image, opts DrawOptions) {
	if len(anim.Frames) == 0 {
		return
	}
	drawImage(anim, screen, anim.Frames[anim.currentIndex], opts)
}
