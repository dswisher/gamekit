package sprites

import (
	"fmt"
	"image"
)

// GridLocator allows for translation of col,row coordinates to an image rectangle,
// primarily intended to be used to extract sprites from a sprite sheet.
type GridLocator struct {
	spriteWidth, spriteHeight int
	border                    int
	sheetWidth, sheetHeight   int // in pixels, 0 means no bounds checking
	sheetCols, sheetRows      int // in grid coords, 0 means not set, only used during initialization
}

type GridLocatorOption func(*GridLocator)

// NewGridLocator creates a new GridLocator with the specified sprite size
// and optional border settings.
//
// Example — simple sheet with no border:
//
//	grid := sprites.NewGridLocator(32, 32)
//	sprite := sheet.Sprite(grid.GetRect(0, 0))
//
// Example — sheet with a 1-pixel border between sprites, bounded to 4×2 grid:
//
//	grid := sprites.NewGridLocator(128, 128,
//	    sprites.WithBorder(1),
//	    sprites.WithSheetSizeGrid(4, 2),
//	)
//	sprite := sheet.Sprite(grid.GetRect(1, 0))
func NewGridLocator(spriteWidth, spriteHeight int, opts ...GridLocatorOption) *GridLocator {
	// Initialize with defaults
	gl := &GridLocator{
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		border:       0,
		sheetWidth:   0,
		sheetHeight:  0,
	}

	// Apply all specified options
	for _, opt := range opts {
		opt(gl)
	}

	// If grid dimensions were set, calculate pixel dimensions now
	if gl.sheetCols > 0 && gl.sheetRows > 0 && gl.sheetWidth == 0 && gl.sheetHeight == 0 {
		gl.sheetWidth = gl.sheetCols*gl.spriteWidth + (gl.sheetCols-1)*gl.border
		gl.sheetHeight = gl.sheetRows*gl.spriteHeight + (gl.sheetRows-1)*gl.border
	}

	return gl
}

// WithBorder sets the border size between sprites in pixels.
// The border is the spacing between adjacent sprites in the grid.
func WithBorder(border int) GridLocatorOption {
	return func(gl *GridLocator) {
		gl.border = border
	}
}

// WithSheetSizePixels sets the sprite sheet dimensions in pixels, enabling
// bounds checking in GetRect. Use this when you know the pixel dimensions of
// the sprite sheet image directly (e.g. from the image metadata).
//
// Example:
//
//	grid := sprites.NewGridLocator(32, 32, sprites.WithSheetSizePixels(256, 128))
func WithSheetSizePixels(width, height int) GridLocatorOption {
	return func(gl *GridLocator) {
		gl.sheetWidth = width
		gl.sheetHeight = height
	}
}

// WithSheetSizeGrid sets the sprite sheet dimensions in grid coordinates,
// enabling bounds checking in GetRect. Use this when you know the number of
// columns and rows rather than the pixel dimensions.
//
// Example:
//
//	grid := sprites.NewGridLocator(32, 32, sprites.WithSheetSizeGrid(8, 4))
func WithSheetSizeGrid(cols, rows int) GridLocatorOption {
	return func(gl *GridLocator) {
		gl.sheetCols = cols
		gl.sheetRows = rows
	}
}

// GetRect returns the image rectangle for the sprite at the specified grid coordinates.
//
// Panics if the coordinates are out of bounds when sheet dimensions have been configured.
// This behavior is optional, and it panics rather than returning error semantics to keep
// the calls clean. It is assumed all sprites will be loaded when game/scene starts,
// so if there is an issue, it will show itself early in development.
func (gl *GridLocator) GetRect(col, row int) image.Rectangle {
	// Validate bounds if sheet dimensions are configured
	if gl.sheetWidth > 0 && gl.sheetHeight > 0 {
		maxCols := (gl.sheetWidth + gl.border) / (gl.spriteWidth + gl.border)
		maxRows := (gl.sheetHeight + gl.border) / (gl.spriteHeight + gl.border)

		if col < 0 || col >= maxCols || row < 0 || row >= maxRows {
			panic(fmt.Sprintf("sprites: grid coordinate (%d, %d) out of bounds (max %d, %d)", col, row, maxCols, maxRows))
		}
	}

	x0 := col*gl.spriteWidth + col*gl.border
	y0 := row*gl.spriteHeight + row*gl.border

	return image.Rect(x0, y0, x0+gl.spriteWidth, y0+gl.spriteHeight)
}
