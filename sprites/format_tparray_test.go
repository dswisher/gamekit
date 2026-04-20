package sprites

import (
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadArrayFormat_ValidData(t *testing.T) {
	jsonData := `{
		"frames": [
			{
				"filename": "RunRight01.png",
				"frame": {"x": 1, "y": 1, "w": 128, "h": 128},
				"rotated": false,
				"trimmed": false,
				"spriteSourceSize": {"x": 0, "y": 0, "w": 128, "h": 128},
				"sourceSize": {"w": 128, "h": 128}
			},
			{
				"filename": "RunRight02.png",
				"frame": {"x": 131, "y": 1, "w": 64, "h": 64},
				"rotated": true,
				"trimmed": true,
				"spriteSourceSize": {"x": 0, "y": 0, "w": 64, "h": 64},
				"sourceSize": {"w": 64, "h": 64}
			}
		],
		"meta": {
			"app": "http://www.codeandweb.com/texturepacker",
			"version": "1.0",
			"image": "texture-packer.png",
			"format": "RGBA8888",
			"size": {"w": 260, "h": 260},
			"scale": "1"
		}
	}`

	sheet, err := loadArrayFormat([]byte(jsonData))
	require.NoError(t, err)
	require.NotNil(t, sheet)

	// Check metadata
	assert.Equal(t, "texture-packer.png", sheet.Image)
	assert.Equal(t, "RGBA8888", sheet.Format)
	assert.Equal(t, image.Pt(260, 260), sheet.Size)
	assert.Equal(t, "1", sheet.Scale)

	// Check sprites
	require.Len(t, sheet.Sprites, 2)

	// First sprite (lowercase key for case-insensitive lookup)
	sprite1, ok := sheet.Sprites["runright01.png"]
	require.True(t, ok, "first sprite should exist with lowercase key")
	assert.Equal(t, "RunRight01.png", sprite1.Name)
	assert.Equal(t, image.Rect(1, 1, 129, 129), sprite1.Rect)
	assert.False(t, sprite1.Rotated)
	assert.False(t, sprite1.Trimmed)
	assert.Equal(t, image.Pt(128, 128), sprite1.SourceSize)

	// Second sprite
	sprite2, ok := sheet.Sprites["runright02.png"]
	require.True(t, ok, "second sprite should exist with lowercase key")
	assert.Equal(t, "RunRight02.png", sprite2.Name)
	assert.Equal(t, image.Rect(131, 1, 195, 65), sprite2.Rect)
	assert.True(t, sprite2.Rotated)
	assert.True(t, sprite2.Trimmed)
	assert.Equal(t, image.Pt(64, 64), sprite2.SourceSize)
}

func TestLoadArrayFormat_InvalidJSON(t *testing.T) {
	jsonData := `{"frames": [invalid json]}`

	sheet, err := loadArrayFormat([]byte(jsonData))
	assert.Error(t, err)
	assert.Nil(t, sheet)
	assert.Contains(t, err.Error(), "failed to unmarshal")
}

func TestLoadArrayFormat_EmptyFrames(t *testing.T) {
	jsonData := `{
		"frames": [],
		"meta": {
			"image": "empty.png",
			"size": {"w": 100, "h": 100}
		}
	}`

	sheet, err := loadArrayFormat([]byte(jsonData))
	require.NoError(t, err)
	require.NotNil(t, sheet)
	assert.Empty(t, sheet.Sprites)
}

func TestLoadArrayFormat_MissingFilename(t *testing.T) {
	// Frame without a filename should still be processed with empty name
	jsonData := `{
		"frames": [
			{
				"frame": {"x": 0, "y": 0, "w": 32, "h": 32},
				"rotated": false,
				"trimmed": false,
				"spriteSourceSize": {"x": 0, "y": 0, "w": 32, "h": 32},
				"sourceSize": {"w": 32, "h": 32}
			}
		],
		"meta": {
			"image": "test.png"
		}
	}`

	sheet, err := loadArrayFormat([]byte(jsonData))
	require.NoError(t, err)
	require.NotNil(t, sheet)
	require.Len(t, sheet.Sprites, 1)

	// Empty string filename stored with lowercase key (still empty)
	sprite, ok := sheet.Sprites[""]
	require.True(t, ok)
	assert.Equal(t, "", sprite.Name)
}

func TestLoadArrayFormat_RealFile(t *testing.T) {
	// This test uses the actual test data file
	data, err := os.ReadFile("testdata/texture-packer-array.json")
	require.NoError(t, err)

	sheet, err := loadArrayFormat(data)
	require.NoError(t, err)
	require.NotNil(t, sheet)

	// Verify the expected sprites exist
	assert.Equal(t, "texture-packer.png", sheet.Image)
	assert.Len(t, sheet.Sprites, 4)

	// Check one of the sprites
	sprite, ok := sheet.Sprites["runright01.png"]
	require.True(t, ok)
	assert.Equal(t, "RunRight01.png", sprite.Name)
	assert.Equal(t, image.Rect(1, 1, 129, 129), sprite.Rect)
}
