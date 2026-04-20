package main

import (
	"embed"
	"log"

	"github.com/dswisher/gamekit/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	turret       *sprites.Sprite
	runRightGrid *sprites.Sprite
	runRightMeta *sprites.Sprite
}

func NewGame() *Game {
	g := &Game{}
	g.loadAssets()

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	g.turret.Draw(screen, sprites.DrawOpts(100, 50))
	g.runRightGrid.Draw(screen, sprites.DrawOpts(300, 50))
	g.runRightMeta.Draw(screen, sprites.DrawOpts(500, 50))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) loadAssets() {
	// Load a simple standalone sprite
	img, err := sprites.LoadImageFromFS(assets, "assets/turret-03.png")
	if err != nil {
		log.Fatal(err)
	}
	g.turret = sprites.NewSprite(img)

	// Load the sprite sheet image
	img, err = sprites.LoadImageFromFS(assets, "assets/texture-packer.png")
	if err != nil {
		log.Fatal(err)
	}

	sheet := sprites.NewSheet(img)

	// Set up a grid locator and use it to extract a sprite from the sheet
	grid := sprites.NewGridLocator(128, 128, sprites.WithBorder(1))
	g.runRightGrid = sheet.Sprite(grid.GetRect(1, 0))

	// Load the metadata for the sprite sheet and use it to extract a sprite from the sheet
	meta, err := sprites.LoadMetadataFromFS(assets, "assets/texture-packer-hash.json", "")
	if err != nil {
		log.Fatal(err)
	}

	g.runRightMeta = sheet.Sprite(meta.GetRect("RunRight01.png"))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
