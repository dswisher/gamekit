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
	turret   *sprites.Sprite
	runRight *sprites.Sprite
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

	g.turret.Draw(screen, sprites.DrawOpts(50, 50))
	g.runRight.Draw(screen, sprites.DrawOpts(300, 50))
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

	// Load a sprite sheet and use JSON metadata to create a sprite from it
	// TODO
	img, err = sprites.LoadImageFromFS(assets, "assets/texture-packer.png")
	if err != nil {
		log.Fatal(err)
	}
	g.runRight = sprites.NewSprite(img)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
