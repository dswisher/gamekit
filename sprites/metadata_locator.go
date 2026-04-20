package sprites

import (
	"fmt"
	"image"
	"strings"
)

// MetadataLocator provides name-based sprite lookup from metadata files.
// It loads sprite sheet metadata from JSON files and allows retrieving
// sprite rectangles by name.
type MetadataLocator struct {
	sheet *sheetMetadata
}

// GetRect returns the image rectangle for the sprite with the given name.
// The name lookup is case-insensitive.
//
// Panics if the sprite is not found, consistent with GridLocator behavior.
// Use HasSprite to check existence before calling if needed.
func (ml *MetadataLocator) GetRect(name string) image.Rectangle {
	if ml.sheet == nil {
		panic("sprites: MetadataLocator has no loaded data")
	}

	key := strings.ToLower(name)
	sprite, ok := ml.sheet.Sprites[key]
	if !ok {
		panic(fmt.Sprintf("sprites: sprite %q not found in metadata", name))
	}

	return sprite.Rect
}

// HasSprite checks if a sprite with the given name exists in the metadata.
// The name lookup is case-insensitive.
func (ml *MetadataLocator) HasSprite(name string) bool {
	if ml.sheet == nil {
		return false
	}

	key := strings.ToLower(name)
	_, ok := ml.sheet.Sprites[key]
	return ok
}

// SpriteNames returns a slice of all available sprite names in their original case.
// The order of names is not guaranteed.
func (ml *MetadataLocator) SpriteNames() []string {
	if ml.sheet == nil {
		return nil
	}

	names := make([]string, 0, len(ml.sheet.Sprites))
	for _, sprite := range ml.sheet.Sprites {
		names = append(names, sprite.Name)
	}
	return names
}

// ImagePath returns the path to the sprite sheet image as specified in the metadata.
func (ml *MetadataLocator) ImagePath() string {
	if ml.sheet == nil {
		return ""
	}
	return ml.sheet.Image
}
