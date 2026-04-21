// Package sprites provides sprite sheet management and animation for Ebitengine.
//
// The sprites package offers tools for loading, managing, and rendering sprite
// sheets and individual sprites. It supports loading images from various sources
// including the filesystem and embedded assets, and provides flexible drawing
// options for positioning and transforming sprites.
//
// # Loading a standalone sprite
//
//	//go:embed assets/*.png
//	var assets embed.FS
//
//	img, err := sprites.LoadImageFromFS(assets, "assets/player.png")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	sprite := sprites.NewSprite(img)
//
//	// In your Draw method:
//	sprite.Draw(screen, sprites.DrawOpts(100, 100))
//
// # Using a sprite sheet with a grid locator
//
// A [GridLocator] translates (col, row) coordinates into image rectangles,
// suitable for sprite sheets where all frames are the same size.
//
//	img, err := sprites.LoadImageFromFS(assets, "assets/sheet.png")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	sheet := sprites.NewSheet(img)
//
//	// 128×128 sprites with a 1-pixel border between frames
//	grid := sprites.NewGridLocator(128, 128, sprites.WithBorder(1))
//	runSprite := sheet.Sprite(grid.GetRect(1, 0))
//
// # Using a sprite sheet with JSON metadata
//
// A [MetadataLocator] translates sprite names into image rectangles, loaded
// from a TexturePacker JSON metadata file. Both hash and array formats are
// supported; pass an empty string to auto-detect.
//
//	meta, err := sprites.LoadMetadataFromFS(assets, "assets/sheet.json", "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	idleSprite := sheet.Sprite(meta.GetRect("Idle01.png"))
package sprites
