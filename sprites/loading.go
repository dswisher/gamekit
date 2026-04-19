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

// LoadImageFromFS loads an image from any fs.FS implementation.
// It returns an ebiten.Image ready for use with sprites.
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
