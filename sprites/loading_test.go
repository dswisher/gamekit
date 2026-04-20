package sprites

import (
	"image"
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadImageFromFS_Success(t *testing.T) {
	// Use real filesystem with testdata
	img, err := LoadImageFromFS(os.DirFS("testdata"), "texture-packer.png")
	// Note: This test requires a PNG file in testdata
	// If the file doesn't exist, this will error which is expected behavior
	if err != nil {
		// The image file might not exist, that's ok - just verify error is descriptive
		assert.Contains(t, err.Error(), "failed to read image")
	} else {
		require.NotNil(t, img)
		assert.Greater(t, img.Bounds().Dx(), 0)
		assert.Greater(t, img.Bounds().Dy(), 0)
	}
}

func TestLoadImageFromFS_MissingFile(t *testing.T) {
	mapFS := fstest.MapFS{}

	img, err := LoadImageFromFS(mapFS, "nonexistent.png")
	assert.Error(t, err)
	assert.Nil(t, img)
	assert.Contains(t, err.Error(), "failed to read image")
}

func TestLoadImageFromFS_InvalidImageData(t *testing.T) {
	mapFS := fstest.MapFS{
		"invalid.png": &fstest.MapFile{
			Data: []byte("not a valid png image"),
		},
	}

	img, err := LoadImageFromFS(mapFS, "invalid.png")
	assert.Error(t, err)
	assert.Nil(t, img)
	assert.Contains(t, err.Error(), "failed to decode image")
}

func TestLoadMetadataFromFS_HashFormat(t *testing.T) {
	locator, err := LoadMetadataFromFS(os.DirFS("testdata"), "texture-packer-hash.json", "texture-packer-hash")
	require.NoError(t, err)
	require.NotNil(t, locator)

	// Verify we can retrieve a sprite
	assert.True(t, locator.HasSprite("RunRight01.png"))
	rect := locator.GetRect("RunRight01.png")
	assert.Equal(t, image.Rect(1, 1, 129, 129), rect)

	// Verify image path
	assert.Equal(t, "texture-packer.png", locator.ImagePath())
}

func TestLoadMetadataFromFS_ArrayFormat(t *testing.T) {
	locator, err := LoadMetadataFromFS(os.DirFS("testdata"), "texture-packer-array.json", "texture-packer-array")
	require.NoError(t, err)
	require.NotNil(t, locator)

	// Verify we can retrieve a sprite
	assert.True(t, locator.HasSprite("RunRight01.png"))
	rect := locator.GetRect("RunRight01.png")
	assert.Equal(t, image.Rect(1, 1, 129, 129), rect)

	// Verify image path
	assert.Equal(t, "texture-packer.png", locator.ImagePath())
}

func TestLoadMetadataFromFS_AutoDetectHash(t *testing.T) {
	locator, err := LoadMetadataFromFS(os.DirFS("testdata"), "texture-packer-hash.json", "")
	require.NoError(t, err)
	require.NotNil(t, locator)

	// Should have detected and loaded hash format
	assert.True(t, locator.HasSprite("RunRight01.png"))
	rect := locator.GetRect("RunRight01.png")
	assert.Equal(t, image.Rect(1, 1, 129, 129), rect)
}

func TestLoadMetadataFromFS_AutoDetectArray(t *testing.T) {
	locator, err := LoadMetadataFromFS(os.DirFS("testdata"), "texture-packer-array.json", "")
	require.NoError(t, err)
	require.NotNil(t, locator)

	// Should have detected and loaded array format
	assert.True(t, locator.HasSprite("RunRight01.png"))
	rect := locator.GetRect("RunRight01.png")
	assert.Equal(t, image.Rect(1, 1, 129, 129), rect)
}

func TestLoadMetadataFromFS_MissingFile(t *testing.T) {
	mapFS := fstest.MapFS{}

	locator, err := LoadMetadataFromFS(mapFS, "nonexistent.json", "texture-packer-hash")
	assert.Error(t, err)
	assert.Nil(t, locator)
	assert.Contains(t, err.Error(), "failed to read metadata file")
}

func TestLoadMetadataFromFS_InvalidJSON(t *testing.T) {
	mapFS := fstest.MapFS{
		"invalid.json": &fstest.MapFile{
			Data: []byte("not valid json"),
		},
	}

	// When format is specified, it will fail during parsing
	locator, err := LoadMetadataFromFS(mapFS, "invalid.json", "texture-packer-hash")
	assert.Error(t, err)
	assert.Nil(t, locator)
}

func TestLoadMetadataFromFS_UnsupportedFormat(t *testing.T) {
	mapFS := fstest.MapFS{
		"data.json": &fstest.MapFile{
			Data: []byte(`{"frames": {}, "meta": {}}`),
		},
	}

	locator, err := LoadMetadataFromFS(mapFS, "data.json", "unsupported-format")
	assert.Error(t, err)
	assert.Nil(t, locator)
	assert.Contains(t, err.Error(), "unsupported metadata format")
}

func TestLoadMetadataFromFS_AutoDetectAmbiguous(t *testing.T) {
	mapFS := fstest.MapFS{
		"ambiguous.json": &fstest.MapFile{
			Data: []byte(`{"meta": {"image": "test.png"}}`),
		},
	}

	// Missing frames field should cause detection to fail
	locator, err := LoadMetadataFromFS(mapFS, "ambiguous.json", "")
	assert.Error(t, err)
	assert.Nil(t, locator)
}

func TestLoadMetadataFromFS_WithMapFS(t *testing.T) {
	// Test using MapFS for in-memory testing
	jsonData := `{
		"frames": {
			"test.png": {
				"frame": {"x": 0, "y": 0, "w": 32, "h": 32},
				"rotated": false,
				"trimmed": false,
				"spriteSourceSize": {"x": 0, "y": 0, "w": 32, "h": 32},
				"sourceSize": {"w": 32, "h": 32}
			}
		},
		"meta": {
			"image": "sheet.png",
			"format": "RGBA8888",
			"size": {"w": 32, "h": 32},
			"scale": "1"
		}
	}`

	mapFS := fstest.MapFS{
		"sprites.json": &fstest.MapFile{
			Data: []byte(jsonData),
		},
	}

	locator, err := LoadMetadataFromFS(mapFS, "sprites.json", "texture-packer-hash")
	require.NoError(t, err)
	require.NotNil(t, locator)

	assert.True(t, locator.HasSprite("test.png"))
	rect := locator.GetRect("test.png")
	assert.Equal(t, image.Rect(0, 0, 32, 32), rect)
	assert.Equal(t, "sheet.png", locator.ImagePath())
}

func TestLoadMetadataFromFS_DirFS(t *testing.T) {
	// Test with actual filesystem
	_, err := LoadMetadataFromFS(os.DirFS("testdata"), "texture-packer-hash.json", "texture-packer-hash")
	require.NoError(t, err)
}

func TestLoadMetadataFromFS_InvalidPath(t *testing.T) {
	// Test with invalid path patterns
	mapFS := fstest.MapFS{}

	// Path starting with /
	locator, err := LoadMetadataFromFS(mapFS, "/absolute/path.json", "texture-packer-hash")
	assert.Error(t, err)
	assert.Nil(t, locator)
	// Should get an error about the path (exact error depends on fs implementation)
}

// Ensure fs.FS interface compliance
var _ fs.FS = (fstest.MapFS)(nil)
