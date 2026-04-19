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

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image %s: %w", path, err)
	}

	return ebiten.NewImageFromImage(img), nil
}
