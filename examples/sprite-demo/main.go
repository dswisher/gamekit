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
	turret *sprites.Sprite
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

	g.turret.Draw(screen, sprites.DrawOpts(100, 300))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) loadAssets() {
	img, err := sprites.LoadImageFromFS(assets, "assets/turret-03.png")
	if err != nil {
		log.Fatal(err)
	}
	g.turret = sprites.NewSprite(img)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
