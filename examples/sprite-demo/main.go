package main

import (
	"embed"
	"log"
	"math"

	"github.com/dswisher/gamekit/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

//go:embed assets/*
var assets embed.FS

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	turret        *sprites.Sprite
	runRightGrid  *sprites.Sprite
	runRightMeta  *sprites.Sprite
	yellowCircle  *sprites.Sprite
	rotation      float64
	gridAnimation *sprites.Animation
}

func NewGame() *Game {
	g := &Game{}
	g.loadAssets()

	return g
}

func (g *Game) Update() error {
	// Exit on ESC key
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// Slowly rotate the turret
	g.rotation += 0.02
	if g.rotation > 2*math.Pi {
		g.rotation -= 2 * math.Pi
	}

	// Update the animation
	g.gridAnimation.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	// Draw turret with rotation around its center (34, 34)
	// Uses the sprite's default origin set in loadAssets()
	g.turret.Draw(screen, sprites.DrawAt(100, 50).WithRotation(g.rotation))

	// Simple positioning with DrawAt
	g.runRightGrid.Draw(screen, sprites.DrawAt(300, 20))
	g.runRightMeta.Draw(screen, sprites.DrawAt(450, 20))

	// Draw runRightMeta scaled to half-size
	g.runRightMeta.Draw(screen, sprites.DrawAt(600, 50).WithScale(0.5))

	// Demonstrate BlendLighter with scaled-up circles
	g.drawBlendDemo(screen, 50, 200)

	// Demonstrate ColorM (color matrix) transformations
	g.drawColorMDemo(screen, 50, 300)

	// Draw the animation
	g.gridAnimation.Draw(screen, sprites.DrawAt(30, 400).WithScale(3.0))
}

// drawBlendDemo demonstrates BlendLighter with overlapping circles.
// Shows how overlapping sprites accumulate brightness using additive blending.
// Each yellow circle has alpha 64, so:
//
//	1 circle = 64 alpha (transparent)
//	2 circles = 128 alpha (semi-transparent)
//	3 circles = 192 alpha (mostly solid)
//	4 circles = 255 alpha (solid yellow)
func (g *Game) drawBlendDemo(screen *ebiten.Image, baseX, baseY float64) {
	scale := 2.0
	// With 2x scale, circles are 48x48 pixels.
	// For ~50% overlap, offset by half the width (~24 pixels).
	halfWidth := 24.0

	// Single circle (0 overlaps)
	g.yellowCircle.Draw(screen, sprites.DrawAt(baseX, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))

	// Two circles (1 overlap at 50%)
	g.yellowCircle.Draw(screen, sprites.DrawAt(baseX+80, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))
	g.yellowCircle.Draw(screen, sprites.DrawAt(baseX+80+halfWidth, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))

	// Three circles (2 overlaps at 50% each) - arranged in a triangle
	triX := baseX + 180.0
	g.yellowCircle.Draw(screen, sprites.DrawAt(triX, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))
	g.yellowCircle.Draw(screen, sprites.DrawAt(triX+halfWidth, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))
	g.yellowCircle.Draw(screen, sprites.DrawAt(triX+halfWidth/2, baseY+halfWidth).WithScale(scale).WithBlend(ebiten.BlendLighter))

	// Four circles (3 overlaps at 50% each) - arranged in a 2x2 grid pattern
	quadX := baseX + 280.0
	g.yellowCircle.Draw(screen, sprites.DrawAt(quadX, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))
	g.yellowCircle.Draw(screen, sprites.DrawAt(quadX+halfWidth, baseY).WithScale(scale).WithBlend(ebiten.BlendLighter))
	g.yellowCircle.Draw(screen, sprites.DrawAt(quadX, baseY+halfWidth).WithScale(scale).WithBlend(ebiten.BlendLighter))
	g.yellowCircle.Draw(screen, sprites.DrawAt(quadX+halfWidth, baseY+halfWidth).WithScale(scale).WithBlend(ebiten.BlendLighter))
}

// drawColorMDemo demonstrates ColorM (color matrix) transformations.
// Shows various color effects like tinting, brightness adjustment, and grayscale.
func (g *Game) drawColorMDemo(screen *ebiten.Image, baseX, baseY float64) {
	scale := 0.5
	spacing := 80.0

	// Original (no color transformation)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX, baseY).WithScale(scale))

	// Red tint - scale green and blue to 0, keep red and alpha
	cm := colorm.ColorM{}
	cm.Scale(1.0, 0.0, 0.0, 1.0)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX+spacing, baseY).WithScale(scale).WithColorM(cm))

	// Green tint - scale red and blue to 0, keep green and alpha
	cm = colorm.ColorM{}
	cm.Scale(0.0, 1.0, 0.0, 1.0)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX+spacing*2, baseY).WithScale(scale).WithColorM(cm))

	// Blue tint - scale red and green to 0, keep blue and alpha
	cm = colorm.ColorM{}
	cm.Scale(0.0, 0.0, 1.0, 1.0)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX+spacing*3, baseY).WithScale(scale).WithColorM(cm))

	// 50% darker - scale all RGB channels by 0.5
	cm = colorm.ColorM{}
	cm.Scale(0.5, 0.5, 0.5, 1.0)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX+spacing*4, baseY).WithScale(scale).WithColorM(cm))

	// Grayscale - average all channels (using standard weights: 0.299*R + 0.587*G + 0.114*B)
	// This uses Translate and Scale to create a grayscale matrix
	cm = colorm.ColorM{}
	cm.Scale(0.0, 0.0, 0.0, 1.0)           // Clear original colors
	cm.Translate(0.333, 0.333, 0.333, 0.0) // Add equal amounts (simplified grayscale)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX+spacing*5, baseY).WithScale(scale).WithColorM(cm))

	// Flash effect (brighter) - scale up and add white
	cm = colorm.ColorM{}
	cm.Scale(1.5, 1.5, 1.5, 1.0)
	cm.Translate(0.2, 0.2, 0.2, 0.0)
	g.runRightGrid.Draw(screen, sprites.DrawAt(baseX+spacing*6, baseY).WithScale(scale).WithColorM(cm))
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
	// Set the rotation center to (34, 34) - the center of the turret
	g.turret.SetOrigin(34, 34)

	// Load a yellow circle (with alpha 64 ~ 25%) for demonstrating blend modes
	img, err = sprites.LoadImageFromFS(assets, "assets/yellow-circle.png")
	if err != nil {
		log.Fatal(err)
	}
	g.yellowCircle = sprites.NewSprite(img)

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

	// Use a grid locator to set up an animation using the ebiten "runner" example
	img, err = sprites.LoadImageFromBytes(images.Runner_png)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: should we have a "sprites.NoBorder" or some such??
	grid = sprites.NewGridLocator(32, 32, sprites.WithBorder(0))

	g.gridAnimation = sprites.NewAnimation(img, grid.GetRowRects(0, 1, 8))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
