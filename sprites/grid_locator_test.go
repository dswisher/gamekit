package sprites

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGridLocator_Defaults(t *testing.T) {
	gl := NewGridLocator(32, 48)

	assert.Equal(t, 32, gl.spriteWidth, "spriteWidth should be 32")
	assert.Equal(t, 48, gl.spriteHeight, "spriteHeight should be 48")
	assert.Equal(t, 0, gl.border, "border should default to 0")
	assert.Equal(t, 0, gl.sheetWidth, "sheetWidth should default to 0")
	assert.Equal(t, 0, gl.sheetHeight, "sheetHeight should default to 0")
}

func TestNewGridLocator_WithBorder(t *testing.T) {
	gl := NewGridLocator(32, 32, WithBorder(2))

	assert.Equal(t, 2, gl.border, "border should be 2")
}

func TestNewGridLocator_WithSheetSizePixels(t *testing.T) {
	gl := NewGridLocator(32, 32, WithSheetSizePixels(128, 256))

	assert.Equal(t, 128, gl.sheetWidth, "sheetWidth should be 128")
	assert.Equal(t, 256, gl.sheetHeight, "sheetHeight should be 256")
}

func TestNewGridLocator_WithSheetSizeGrid(t *testing.T) {
	gl := NewGridLocator(32, 32, WithSheetSizeGrid(4, 8))

	assert.Equal(t, 4, gl.sheetCols, "sheetCols should be 4")
	assert.Equal(t, 8, gl.sheetRows, "sheetRows should be 8")
	// Pixel dimensions should be calculated from grid: 4*32 x 8*32 = 128 x 256
	assert.Equal(t, 128, gl.sheetWidth, "sheetWidth should be calculated from grid as 128")
	assert.Equal(t, 256, gl.sheetHeight, "sheetHeight should be calculated from grid as 256")
}

func TestNewGridLocator_WithSheetSizeGrid_WithBorder(t *testing.T) {
	// 4 columns of 32px sprites with 2px borders: 4*32 + 3*2 = 128 + 6 = 134
	// 3 rows of 32px sprites with 2px borders: 3*32 + 2*2 = 96 + 4 = 100
	gl := NewGridLocator(32, 32, WithBorder(2), WithSheetSizeGrid(4, 3))

	assert.Equal(t, 134, gl.sheetWidth, "sheetWidth should account for borders: 4*32 + 3*2 = 134")
	assert.Equal(t, 100, gl.sheetHeight, "sheetHeight should account for borders: 3*32 + 2*2 = 100")
}

func TestNewGridLocator_PixelSizeOverridesGridCalculation(t *testing.T) {
	// If both pixel size and grid size are set, pixel size takes precedence
	// and grid-based calculation is skipped
	gl := NewGridLocator(32, 32, WithSheetSizePixels(100, 200), WithSheetSizeGrid(4, 4))

	assert.Equal(t, 100, gl.sheetWidth, "sheetWidth should be 100 from pixel setting")
	assert.Equal(t, 200, gl.sheetHeight, "sheetHeight should be 200 from pixel setting")
	// Grid dimensions are still stored
	assert.Equal(t, 4, gl.sheetCols, "sheetCols should still be set")
	assert.Equal(t, 4, gl.sheetRows, "sheetRows should still be set")
}

func TestGridLocator_GetRect_NoBorder(t *testing.T) {
	gl := NewGridLocator(32, 32)

	// First sprite at (0, 0)
	r := gl.GetRect(0, 0)
	assert.Equal(t, image.Rect(0, 0, 32, 32), r, "rect at (0,0) should be (0,0)-(32,32)")

	// Second sprite in row 0, col 1
	r = gl.GetRect(1, 0)
	assert.Equal(t, image.Rect(32, 0, 64, 32), r, "rect at (1,0) should be (32,0)-(64,32)")

	// Sprite in row 1, col 0
	r = gl.GetRect(0, 1)
	assert.Equal(t, image.Rect(0, 32, 32, 64), r, "rect at (0,1) should be (0,32)-(32,64)")

	// Sprite in row 2, col 3
	r = gl.GetRect(3, 2)
	assert.Equal(t, image.Rect(96, 64, 128, 96), r, "rect at (3,2) should be (96,64)-(128,96)")
}

func TestGridLocator_GetRect_WithBorder(t *testing.T) {
	gl := NewGridLocator(32, 48, WithBorder(2))

	// First sprite at (0, 0) - no border before it
	r := gl.GetRect(0, 0)
	assert.Equal(t, image.Rect(0, 0, 32, 48), r, "rect at (0,0) with border should be (0,0)-(32,48)")

	// Second sprite in row 0, col 1 - has 2px border before it
	// x0 = 1*32 + 1*2 = 34
	r = gl.GetRect(1, 0)
	assert.Equal(t, image.Rect(34, 0, 66, 48), r, "rect at (1,0) with border should be (34,0)-(66,48)")

	// Sprite in row 1, col 0
	// y0 = 1*48 + 1*2 = 50
	r = gl.GetRect(0, 1)
	assert.Equal(t, image.Rect(0, 50, 32, 98), r, "rect at (0,1) with border should be (0,50)-(32,98)")

	// Sprite in row 2, col 3
	// x0 = 3*32 + 3*2 = 96 + 6 = 102
	// y0 = 2*48 + 2*2 = 96 + 4 = 100
	r = gl.GetRect(3, 2)
	assert.Equal(t, image.Rect(102, 100, 134, 148), r, "rect at (3,2) with border should be (102,100)-(134,148)")
}

func TestGridLocator_GetRect_BoundsChecking_NoSheetSize(t *testing.T) {
	// When no sheet size is configured, bounds checking is skipped
	// and any coordinates should work without panic
	gl := NewGridLocator(32, 32)

	// Large coordinates should work fine
	assert.NotPanics(t, func() {
		r := gl.GetRect(100, 200)
		assert.Equal(t, image.Rect(3200, 6400, 3232, 6432), r)
	}, "should not panic with large coordinates when no sheet size is set")

	// Negative coordinates - current implementation allows these
	// This documents current behavior; may want to change in future
	assert.NotPanics(t, func() {
		r := gl.GetRect(-1, -1)
		assert.Equal(t, image.Rect(-32, -32, 0, 0), r)
	}, "current implementation allows negative coordinates")
}

func TestGridLocator_GetRect_BoundsChecking_WithPixelSheetSize(t *testing.T) {
	// 4x4 grid of 32x32 sprites: 128x128 pixels
	gl := NewGridLocator(32, 32, WithSheetSizePixels(128, 128))

	// Valid coordinates should work
	assert.NotPanics(t, func() {
		r := gl.GetRect(0, 0)
		assert.Equal(t, image.Rect(0, 0, 32, 32), r)
	}, "(0,0) should be valid")

	assert.NotPanics(t, func() {
		r := gl.GetRect(3, 3)
		assert.Equal(t, image.Rect(96, 96, 128, 128), r)
	}, "(3,3) should be valid")

	// Out of bounds should panic
	assert.Panics(t, func() {
		gl.GetRect(4, 0)
	}, "(4,0) should panic - col out of bounds")

	assert.Panics(t, func() {
		gl.GetRect(0, 4)
	}, "(0,4) should panic - row out of bounds")

	assert.Panics(t, func() {
		gl.GetRect(-1, 0)
	}, "(-1,0) should panic - negative col")

	assert.Panics(t, func() {
		gl.GetRect(0, -1)
	}, "(0,-1) should panic - negative row")
}

func TestGridLocator_GetRect_BoundsChecking_WithGridSheetSize(t *testing.T) {
	// 4x3 grid of 32x48 sprites
	gl := NewGridLocator(32, 48, WithSheetSizeGrid(4, 3))

	// Valid coordinates should work
	assert.NotPanics(t, func() {
		r := gl.GetRect(3, 2)
		assert.Equal(t, image.Rect(96, 96, 128, 144), r)
	}, "(3,2) should be valid")

	// Out of bounds should panic
	assert.Panics(t, func() {
		gl.GetRect(4, 2)
	}, "(4,2) should panic - col out of bounds")

	assert.Panics(t, func() {
		gl.GetRect(3, 3)
	}, "(3,3) should panic - row out of bounds")
}

func TestGridLocator_GetRect_BoundsChecking_WithBorder(t *testing.T) {
	// 2x2 grid of 32x32 sprites with 2px border
	// Sheet size: 2*32 + 1*2 = 66 pixels wide, 2*32 + 1*2 = 66 pixels tall
	gl := NewGridLocator(32, 32, WithBorder(2), WithSheetSizePixels(66, 66))

	// Valid coordinates
	assert.NotPanics(t, func() {
		r := gl.GetRect(0, 0)
		assert.Equal(t, image.Rect(0, 0, 32, 32), r)
	}, "(0,0) should be valid")

	assert.NotPanics(t, func() {
		r := gl.GetRect(1, 0)
		// x0 = 1*32 + 1*2 = 34
		assert.Equal(t, image.Rect(34, 0, 66, 32), r)
	}, "(1,0) should be valid")

	assert.NotPanics(t, func() {
		r := gl.GetRect(1, 1)
		// x0 = 34, y0 = 1*32 + 1*2 = 34
		assert.Equal(t, image.Rect(34, 34, 66, 66), r)
	}, "(1,1) should be valid")

	// Out of bounds
	assert.Panics(t, func() {
		gl.GetRect(2, 0)
	}, "(2,0) should panic - only 2 columns available")

	assert.Panics(t, func() {
		gl.GetRect(0, 2)
	}, "(0,2) should panic - only 2 rows available")
}

func TestGridLocator_GetRect_NonUniformSpriteSize(t *testing.T) {
	// Sprites that are wider than tall
	gl := NewGridLocator(64, 32)

	r := gl.GetRect(1, 2)
	// x0 = 1*64 = 64, y0 = 2*32 = 64
	assert.Equal(t, image.Rect(64, 64, 128, 96), r, "rect should account for non-uniform sprite size")
}

func TestGridLocator_OptionChaining(t *testing.T) {
	// Apply multiple options in sequence
	gl := NewGridLocator(16, 16,
		WithBorder(1),
		WithSheetSizeGrid(10, 10),
	)

	// With border and 10x10 grid:
	// sheetWidth = 10*16 + 9*1 = 160 + 9 = 169
	// sheetHeight = 10*16 + 9*1 = 160 + 9 = 169
	assert.Equal(t, 1, gl.border)
	assert.Equal(t, 10, gl.sheetCols)
	assert.Equal(t, 10, gl.sheetRows)
	assert.Equal(t, 169, gl.sheetWidth)
	assert.Equal(t, 169, gl.sheetHeight)

	// Test coordinate at edge
	r := gl.GetRect(9, 9)
	// x0 = 9*16 + 9*1 = 144 + 9 = 153
	// y0 = 9*16 + 9*1 = 153
	assert.Equal(t, image.Rect(153, 153, 169, 169), r)
}

func TestGridLocator_GetRect_BoundsEdgeCases(t *testing.T) {
	// Test with exact-fit sheet size
	gl := NewGridLocator(32, 32, WithSheetSizePixels(64, 96))

	// 64x96 sheet with 32x32 sprites = 2 cols, 3 rows exactly
	// Valid: (0,0), (1,0), (0,1), (1,1), (0,2), (1,2)
	testCases := []struct {
		col         int
		row         int
		shouldPanic bool
		desc        string
	}{
		{0, 0, false, "first sprite"},
		{1, 2, false, "last valid sprite"},
		{2, 0, true, "col exactly at width"},
		{0, 3, true, "row exactly at height"},
		{2, 3, true, "both out of bounds"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if tc.shouldPanic {
				assert.Panics(t, func() { gl.GetRect(tc.col, tc.row) })
			} else {
				assert.NotPanics(t, func() { gl.GetRect(tc.col, tc.row) })
			}
		})
	}
}

func TestGridLocator_ZeroSizeSprites(t *testing.T) {
	// Edge case: zero-sized sprites
	gl := NewGridLocator(0, 0, WithSheetSizePixels(100, 100))

	// This would cause division by zero in bounds checking
	// The current implementation would panic when dividing by zero
	// This test documents this behavior
	require.Panics(t, func() {
		gl.GetRect(0, 0)
	}, "zero-sized sprites cause division by zero in bounds check")
}
