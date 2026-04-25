package sprites

import (
	"bytes"
	"fmt"
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"

	// Register image formats
	_ "image/png"
)

// LoadImageFromBytes decodes image data from a byte slice and returns it as
// an ebiten.Image ready for use with the sprites package.
//
// This function is useful when you have image data already loaded in memory,
// such as from an embedded file, a network request, or any other source that
// provides raw image bytes.
//
// Supported image formats are determined by the registered image decoders.
// By default, PNG is supported. Import other image format packages to add
// support for JPEG, GIF, etc.
//
// Parameters:
//   - data: The raw image data as a byte slice
//
// Returns:
//   - *ebiten.Image: The loaded image wrapped as an ebiten.Image
//   - error: An error if the data cannot be decoded as an image
//
// Example using embedded assets:
//
//	//go:embed assets/player.png
//	var playerPNG []byte
//	img, err := sprites.LoadImageFromBytes(playerPNG)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	sprite := sprites.NewSprite(img)
func LoadImageFromBytes(data []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return ebiten.NewImageFromImage(img), nil
}

// LoadImageFromFS loads an image from any fs.FS implementation and returns
// it as an ebiten.Image ready for use with the sprites package.
//
// This function supports any filesystem that implements the fs.FS interface,
// including:
//   - os.DirFS for loading from the local filesystem
//   - embed.FS for loading embedded assets
//   - Custom fs.FS implementations
//
// Supported image formats are determined by the registered image decoders.
// By default, PNG is supported. Import other image format packages to add
// support for JPEG, GIF, etc.
//
// Parameters:
//   - fsys: The filesystem to load from (must implement fs.FS)
//   - path: The path to the image file within the filesystem
//
// Returns:
//   - *ebiten.Image: The loaded image wrapped as an ebiten.Image
//   - error: An error if the file cannot be read or decoded
//
// Example using embedded assets:
//
//	//go:embed assets/*.png
//	var assets embed.FS
//	img, err := sprites.LoadImageFromFS(assets, "assets/player.png")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	sprite := sprites.NewSprite(img)
//
// Example using local filesystem:
//
//	fsys := os.DirFS("./assets")
//	img, err := sprites.LoadImageFromFS(fsys, "enemy.png")
func LoadImageFromFS(fsys fs.FS, path string) (*ebiten.Image, error) {
	b, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read image %s: %w", path, err)
	}

	img, err := LoadImageFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image %s: %w", path, err)
	}

	return img, nil
}

// LoadMetadataFromFS loads sprite sheet metadata from a JSON file.
//
// The format parameter specifies the JSON format:
//   - "texture-packer-hash": Hash format where frames is an object
//   - "texture-packer-array": Array format where frames is an array
//   - "": Empty string enables auto-detection (default)
//
// Parameters:
//   - fsys:   The filesystem to load from (must implement fs.FS)
//   - path:   The path to the JSON file within the filesystem
//   - format: The format identifier, or "" for auto-detection
//
// Returns:
//   - *MetadataLocator: The loaded metadata locator
//   - error: An error if the file cannot be read, parsed, or format is unsupported
//
// Example:
//
//	locator, err := sprites.LoadMetadataFromFS(assets, "sprites.json", "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	rect := locator.GetRect("player_run_01")
func LoadMetadataFromFS(fsys fs.FS, path, format string) (*MetadataLocator, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file %s: %w", path, err)
	}

	// Auto-detect format if not specified
	if format == "" {
		detectedFormat, detectErr := detectFormat(data)
		if detectErr != nil {
			return nil, fmt.Errorf("failed to detect format for %s: %w", path, detectErr)
		}
		format = detectedFormat
	}

	var sheet *sheetMetadata
	switch format {
	case "texture-packer-hash":
		sheet, err = loadHashFormat(data)
	case "texture-packer-array":
		sheet, err = loadArrayFormat(data)
	default:
		return nil, fmt.Errorf("unsupported metadata format: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to parse metadata file %s: %w", path, err)
	}

	return &MetadataLocator{sheet: sheet}, nil
}
