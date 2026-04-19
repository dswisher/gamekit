package main

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/dswisher/gamekit/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

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

	g.turret.Draw(screen, sprites.DrawOpts(100, 400))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) loadAssets() {
	img := ebiten.NewImageFromImage(readImage("assets/turret-03.png"))
	g.turret = sprites.NewSprite(img)
}

//go:embed assets/*
var assets embed.FS

func readImage(file string) image.Image {
	b, err := assets.ReadFile(file)
	if err != nil {
		log.Fatal("Failed to read file: ", err)
	}
	return bytes2Image(b)
}

func bytes2Image(rawImage []byte) image.Image {
	img, format, error := image.Decode(bytes.NewReader(rawImage))
	if error != nil {
		log.Fatal("Bytes2Image Failed: ", format, error)
	}
	return img
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
