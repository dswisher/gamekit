package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteSheet represents a collection of sprites organized in a single image.
// A sprite sheet (also known as a texture atlas) is an efficient way to store
// multiple sprites in one image file, reducing texture switching and improving
// rendering performance.
//
// Sprite sheets are typically created from sprite sheet images that contain
// multiple frames or individual sprites arranged in a grid or custom layout.
// Individual sprites can be extracted using the Sprite method with the
// appropriate rectangle coordinates.
//
// Planned features:
//   - Grid-based layout support for automatic sprite extraction
//   - JSON metadata loading for complex sprite sheets
//   - Animation sequence support
//
// Example usage (planned):
//
//	sheet := sprites.NewSheet(playerImg, gridLayout)
//	sprite := sheet.Sprite(gridLayout.Cell(0, 0))
type SpriteSheet struct {
	image *ebiten.Image
}

// TODO: add:
//    sprites.NewSheet(img *ebiten.Image, layout Layout)
//       -> sheet := sprites.NewSheet(playerImg, gridLayout)
//
//    sprites.LoadSheetFromMeta(metaPath string)
//       -> sheet := sprites.LoadSheetFromMeta("assets/player.json")

// Sprite extracts a single sprite from the sprite sheet using the specified
// rectangle coordinates. The rectangle defines the portion of the sprite sheet
// image that corresponds to the desired sprite.
//
// The r parameter should be an image.Rectangle with Min and Max points defining
// the top-left and bottom-right corners of the sprite within the sheet.
//
// Example:
//
//	// Extract a 32x32 sprite at position (64, 0)
//	r := image.Rect(64, 0, 96, 32)
//	sprite := sheet.Sprite(r)
//
// Returns nil if the extraction fails or the rectangle is invalid.
func (s *SpriteSheet) Sprite(r image.Rectangle) *Sprite {
	// TODO: implement sprite extraction from the sprite sheet
	return nil
}
