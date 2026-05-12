# Sprites

[![Go Reference](https://pkg.go.dev/badge/github.com/dswisher/gamekit/sprites.svg)](https://pkg.go.dev/github.com/dswisher/gamekit/sprites) ![Version](https://img.shields.io/github/v/tag/dswisher/gamekit?filter=sprites/*&label=)

Sprite sheet management and animation library for Ebitengine.

## Installation

```bash
go get github.com/dswisher/gamekit/sprites
```

## Features

- Sprite sheet loading and management
- Sprite animation support
- Ebitengine integration

## Usage

### Basic Sprite Loading and Drawing

```go
package main

import (
    "embed"
    "log"

    "github.com/dswisher/gamekit/sprites"
    "github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

type Game struct {
    player *sprites.Sprite
}

func NewGame() *Game {
    // Load an image from embedded assets
    img, err := sprites.LoadImageFromFS(assets, "assets/player.png")
    if err != nil {
        log.Fatal(err)
    }

    return &Game{
        player: sprites.NewSprite(img),
    }
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw at position (100, 100)
    g.player.Draw(screen, sprites.DrawAt(100, 100))
}
```

### Transformations: Rotation and Scaling

```go
func (g *Game) Draw(screen *ebiten.Image) {
    // Set default origin for rotation (e.g., center of 64x64 sprite)
    g.player.SetOrigin(32, 32)

    // Draw with rotation (angle in radians)
    g.player.Draw(screen, sprites.DrawAt(200, 200).WithRotation(math.Pi/4))

    // Draw with uniform scaling
    g.player.Draw(screen, sprites.DrawAt(300, 200).WithScale(2.0))

    // Draw with non-uniform scaling
    g.player.Draw(screen, sprites.DrawAt(400, 200).WithScaleXY(2.0, 0.5))

    // Chain multiple transformations
    g.player.Draw(screen, sprites.DrawAt(500, 200).
        WithRotation(math.Pi/4).
        WithScale(1.5))

    // Override origin for a single draw call
    g.player.Draw(screen, sprites.DrawAt(600, 200).
        WithOrigin(0, 0).      // Use top-left corner
        WithRotation(math.Pi/8))
}
```

### Sprite Sheets with Grid Locator

For sprite sheets arranged in a uniform grid:

```go
func NewGame() *Game {
    // Load the sprite sheet image
    sheetImg, err := sprites.LoadImageFromFS(assets, "assets/characters.png")
    if err != nil {
        log.Fatal(err)
    }

    sheet := sprites.NewSheet(sheetImg)

    // Create a grid locator for 32x32 sprites with 1-pixel borders
    grid := sprites.NewGridLocator(32, 32, sprites.WithBorder(1))

    // Extract sprites by grid coordinates (col, row)
    playerSprite := sheet.Sprite(grid.GetRect(0, 0))
    enemySprite := sheet.Sprite(grid.GetRect(1, 0))

    return &Game{
        player: playerSprite,
        enemy:  enemySprite,
    }
}
```

For bounded grid checking (panics if out of bounds):

```go
// Limit to 8 columns × 4 rows
grid := sprites.NewGridLocator(32, 32, 
    sprites.WithBorder(1),
    sprites.WithSheetSizeGrid(8, 4),
)
```

### Sprite Sheets with Metadata (TexturePacker)

For sprite sheets exported from TexturePacker with JSON metadata:

```go
func NewGame() *Game {
    // Load the sprite sheet image
    sheetImg, err := sprites.LoadImageFromFS(assets, "assets/sprites.png")
    if err != nil {
        log.Fatal(err)
    }

    sheet := sprites.NewSheet(sheetImg)

    // Load metadata (auto-detects TexturePacker format)
    meta, err := sprites.LoadMetadataFromFS(assets, "assets/sprites.json", "")
    if err != nil {
        log.Fatal(err)
    }

    // Extract sprites by name (case-insensitive)
    playerSprite := sheet.Sprite(meta.GetRect("player_run_01"))
    enemySprite := sheet.Sprite(meta.GetRect("enemy_idle"))

    // Check if a sprite exists
    if meta.HasSprite("boss_attack") {
        bossSprite := sheet.Sprite(meta.GetRect("boss_attack"))
        _ = bossSprite
    }

    return &Game{
        player: playerSprite,
        enemy:  enemySprite,
    }
}
```

### Animation

Create frame-based animations from sprite sheets:

```go
type Game struct {
    runAnimation *sprites.Animation
}

func NewGame() *Game {
    // Load the sprite sheet
    sheetImg, err := sprites.LoadImageFromFS(assets, "assets/player.png")
    if err != nil {
        log.Fatal(err)
    }

    // Create grid locator for 32x32 sprites
    grid := sprites.NewGridLocator(32, 32)

    // Extract 8 frames from row 0, starting at column 0
    rects := grid.GetRowRects(0, 0, 8)

    // Create animation at 12 FPS (default)
    anim := sprites.NewAnimation(sheetImg, rects)
    anim.SetOrigin(16, 16) // Center for rotation

    return &Game{
        runAnimation: anim,
    }
}

func (g *Game) Update() error {
    // Advance animation frames
    g.runAnimation.Update()
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw current animation frame
    g.runAnimation.Draw(screen, sprites.DrawAt(400, 300).WithScale(2.0))
}
```

### Blend Modes and Color Effects

```go
import "github.com/hajimehoshi/ebiten/v2/colorm"

func (g *Game) Draw(screen *ebiten.Image) {
    // Additive blending for glow effects
    g.sprite.Draw(screen, sprites.DrawAt(100, 100).
        WithBlend(ebiten.BlendLighter))

    // Red tint using color matrix
    cm := colorm.ColorM{}
    cm.Scale(1.0, 0.0, 0.0, 1.0) // Keep red, remove green and blue
    g.sprite.Draw(screen, sprites.DrawAt(200, 100).
        WithColorM(cm))

    // 50% darker
    cm = colorm.ColorM{}
    cm.Scale(0.5, 0.5, 0.5, 1.0)
    g.sprite.Draw(screen, sprites.DrawAt(300, 100).
        WithColorM(cm))

    // Flash effect (brighter)
    cm = colorm.ColorM{}
    cm.Scale(1.5, 1.5, 1.5, 1.0)
    cm.Translate(0.2, 0.2, 0.2, 0.0)
    g.sprite.Draw(screen, sprites.DrawAt(400, 100).
        WithColorM(cm))
}
```

### Loading Images from Different Sources

```go
// From embedded assets (recommended for distribution)
//go:embed assets/*
var assets embed.FS
img, _ := sprites.LoadImageFromFS(assets, "assets/player.png")

// From local filesystem
fsys := os.DirFS("./assets")
img, _ := sprites.LoadImageFromFS(fsys, "player.png")

// From bytes (e.g., downloaded, procedurally generated)
var pngData []byte // your image data
img, _ := sprites.LoadImageFromBytes(pngData)
```

## Related Tools

- [`spritetool`](../cmd/spritetool/) - CLI tool for building sprite sheets

## License

MIT

## Inspiration

- [yottahmd/ganim8-lib](https://github.com/yottahmd/ganim8-lib) - Sprite animation library for Ebitengine inspired by anim8
- [fzipp/texturepacker](https://github.com/fzipp/texturepacker) - Read sprite sheets created and exported as JSON by TexturePacker.
- [setanarut/anim](https://github.com/setanarut/anim) - Animation player for Ebitengine v2
- [Kangaroux/go-spritesheet](https://github.com/Kangaroux/go-spritesheet) - Use YAML to make working with sprite sheets easy. Compatible with ebiten
- [solarlune/goaseprite](https://github.com/solarlune/goaseprite) - Aseprite JSON loader for Go

## Editors

An incomplete list of editing tools, none of which I have used, so these are not recommendations.

- [LibreSprite](https://github.com/LibreSprite/LibreSprite) - fork of [AseSprite](https://www.aseprite.org/) circa 2016
- [Pixelorama](https://github.com/Orama-Interactive/Pixelorama) - built on Godot
- [mtPaint](https://github.com/wjaguar/mtPaint) - Mark Tyler's Painting Program - old, but still active
- [GrafX2](https://pulkomandy.tk/projects/GrafX2) - [gitlab](https://gitlab.com/GrafX2/grafX2) - The ultimate 256 color painting program - even older, but still active

