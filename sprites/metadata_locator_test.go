package sprites

import (
	"image"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestLocator() *MetadataLocator {
	sheet := &sheetMetadata{
		Image:  "test-sheet.png",
		Format: "RGBA8888",
		Size:   image.Pt(256, 256),
		Scale:  "1",
		Sprites: map[string]spriteMetadata{
			"sprite1.png": {
				Name:       "Sprite1.png",
				Rect:       image.Rect(0, 0, 32, 32),
				Rotated:    false,
				Trimmed:    false,
				SourceSize: image.Pt(32, 32),
			},
			"sprite2.png": {
				Name:       "Sprite2.png",
				Rect:       image.Rect(32, 0, 64, 32),
				Rotated:    true,
				Trimmed:    true,
				SourceSize: image.Pt(32, 32),
			},
			"run_001.png": {
				Name:       "Run_001.png",
				Rect:       image.Rect(0, 32, 32, 64),
				Rotated:    false,
				Trimmed:    false,
				SourceSize: image.Pt(32, 32),
			},
		},
	}
	return &MetadataLocator{sheet: sheet}
}

func TestMetadataLocator_GetRect_ExistingSprite(t *testing.T) {
	locator := createTestLocator()

	// Test exact case match
	rect := locator.GetRect("Sprite1.png")
	assert.Equal(t, image.Rect(0, 0, 32, 32), rect)

	// Test lowercase (case-insensitive lookup)
	rect = locator.GetRect("sprite1.png")
	assert.Equal(t, image.Rect(0, 0, 32, 32), rect)

	// Test uppercase
	rect = locator.GetRect("SPRITE1.PNG")
	assert.Equal(t, image.Rect(0, 0, 32, 32), rect)

	// Test mixed case
	rect = locator.GetRect("SpRiTe1.PnG")
	assert.Equal(t, image.Rect(0, 0, 32, 32), rect)
}

func TestMetadataLocator_GetRect_DifferentSprites(t *testing.T) {
	locator := createTestLocator()

	// First sprite
	rect := locator.GetRect("Sprite1.png")
	assert.Equal(t, image.Rect(0, 0, 32, 32), rect)

	// Second sprite (rotated and trimmed)
	rect = locator.GetRect("Sprite2.png")
	assert.Equal(t, image.Rect(32, 0, 64, 32), rect)

	// Third sprite
	rect = locator.GetRect("Run_001.png")
	assert.Equal(t, image.Rect(0, 32, 32, 64), rect)
}

func TestMetadataLocator_GetRect_NonExistentSprite(t *testing.T) {
	locator := createTestLocator()

	assert.Panics(t, func() {
		locator.GetRect("nonexistent.png")
	}, "GetRect should panic for non-existent sprite")
}

func TestMetadataLocator_GetRect_EmptyLocator(t *testing.T) {
	locator := &MetadataLocator{}

	assert.Panics(t, func() {
		locator.GetRect("anything.png")
	}, "GetRect should panic when no sheet is loaded")
}

func TestMetadataLocator_HasSprite_Existing(t *testing.T) {
	locator := createTestLocator()

	// Exact case
	assert.True(t, locator.HasSprite("Sprite1.png"))

	// Case insensitive
	assert.True(t, locator.HasSprite("sprite1.png"))
	assert.True(t, locator.HasSprite("SPRITE1.PNG"))
	assert.True(t, locator.HasSprite("SpRiTe1.PnG"))
}

func TestMetadataLocator_HasSprite_NonExistent(t *testing.T) {
	locator := createTestLocator()

	assert.False(t, locator.HasSprite("nonexistent.png"))
	assert.False(t, locator.HasSprite(""))
}

func TestMetadataLocator_HasSprite_EmptyLocator(t *testing.T) {
	locator := &MetadataLocator{}

	assert.False(t, locator.HasSprite("anything.png"))
}

func TestMetadataLocator_SpriteNames(t *testing.T) {
	locator := createTestLocator()

	names := locator.SpriteNames()
	assert.Len(t, names, 3)

	// Sort for consistent comparison
	sort.Strings(names)
	assert.Equal(t, []string{"Run_001.png", "Sprite1.png", "Sprite2.png"}, names)
}

func TestMetadataLocator_SpriteNames_EmptyLocator(t *testing.T) {
	locator := &MetadataLocator{}

	names := locator.SpriteNames()
	assert.Nil(t, names)
}

func TestMetadataLocator_SpriteNames_EmptySheet(t *testing.T) {
	locator := &MetadataLocator{
		sheet: &sheetMetadata{
			Sprites: map[string]spriteMetadata{},
		},
	}

	names := locator.SpriteNames()
	assert.Empty(t, names)
}

func TestMetadataLocator_ImagePath(t *testing.T) {
	locator := createTestLocator()

	assert.Equal(t, "test-sheet.png", locator.ImagePath())
}

func TestMetadataLocator_ImagePath_EmptyLocator(t *testing.T) {
	locator := &MetadataLocator{}

	assert.Equal(t, "", locator.ImagePath())
}
