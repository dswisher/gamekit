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
// Example usage:
//
//	sheet := sprites.NewSheet(playerImg)
//	locator := sprites.NewGridLocator(32, 32)
//	sprite := sheet.Sprite(locator.GetRect(0, 0))
type SpriteSheet struct {
	image *ebiten.Image
}

// NewSheet creates a new SpriteSheet from the specified ebiten.Image.
// The image can be loaded using LoadImageFromFS or created directly.
//
// Example:
//
//	img, _ := sprites.LoadImageFromFS(assets, "player.png")
//	sheet := sprites.NewSheet(img)
func NewSheet(img *ebiten.Image) *SpriteSheet {
	return &SpriteSheet{image: img}
}

// Sprite extracts a single sprite from the sprite sheet using the specified
// rectangle coordinates. The rectangle defines the portion of the sprite sheet
// image that corresponds to the desired sprite.
//
// The r parameter should be an image.Rectangle with Min and Max points defining
// the top-left and bottom-right corners of the sprite within the sheet.
//
// The rectangle would typically be specified using a Locator, rather than being
// created directly.
//
// Example:
//
//	// Extract a 32x32 sprite at position (64, 0), two different ways
//	r := image.Rect(64, 0, 96, 32)
//	sprite := sheet.Sprite(r)
//
//	locator := sprites.NewGridLocator(32, 32)
//	sprite = sheet.Sprite(locator.GetRect(1, 0))
func (s *SpriteSheet) Sprite(r image.Rectangle) *Sprite {
	subImg := s.image.SubImage(r).(*ebiten.Image)
	return NewSprite(subImg)
}
