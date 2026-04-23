package main

import (
	"image/color"
	"math"

	"github.com/dswisher/gamekit/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// GameScene displays a simple game-like view with a bouncing ball.
type GameScene struct {
	controller scenes.SceneController
	ballX      float64
	ballY      float64
	ballVX     float64
	ballVY     float64
	ballRadius float64
	time       float64
}

// NewGameScene creates a new game scene.
func NewGameScene(controller scenes.SceneController) *GameScene {
	return &GameScene{
		controller: controller,
		ballX:      screenWidth / 2,
		ballY:      screenHeight / 2,
		ballVX:     150,
		ballVY:     100,
		ballRadius: 20,
	}
}

// Update processes game logic.
func (g *GameScene) Update(dt float64) {
	// Update ball position
	g.ballX += g.ballVX * dt
	g.ballY += g.ballVY * dt

	// Bounce off walls
	if g.ballX-g.ballRadius < 0 {
		g.ballX = g.ballRadius
		g.ballVX = -g.ballVX
	}
	if g.ballX+g.ballRadius > screenWidth {
		g.ballX = screenWidth - g.ballRadius
		g.ballVX = -g.ballVX
	}
	if g.ballY-g.ballRadius < 0 {
		g.ballY = g.ballRadius
		g.ballVY = -g.ballVY
	}
	if g.ballY+g.ballRadius > screenHeight {
		g.ballY = screenHeight - g.ballRadius
		g.ballVY = -g.ballVY
	}

	g.time += dt

	// Return to menu on ESC (only on key press, not hold)
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.controller.Pop()
	}
}

// Draw renders the game scene.
func (g *GameScene) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(color.RGBA{R: 20, G: 40, B: 60, A: 255})

	// Draw bouncing ball
	ballColor := color.RGBA{R: 255, G: 100, B: 100, A: 255}
	vector.DrawFilledCircle(screen, float32(g.ballX), float32(g.ballY), float32(g.ballRadius), ballColor, true)

	// Draw a trail effect
	for i := 1; i <= 5; i++ {
		offsetX := float64(i) * -g.ballVX * 0.01
		offsetY := float64(i) * -g.ballVY * 0.01
		alpha := uint8(255 - i*40)
		trailColor := color.RGBA{R: 255, G: 100, B: 100, A: alpha}
		vector.DrawFilledCircle(screen, float32(g.ballX+offsetX), float32(g.ballY+offsetY), float32(g.ballRadius)-float32(i)*3, trailColor, true)
	}

	// Draw title using debug print
	ebitenutil.DebugPrintAt(screen, "GAME SCENE", screenWidth/2-40, 40)

	// Draw instructions
	ebitenutil.DebugPrintAt(screen, "Press ESC to return to menu", screenWidth/2-120, screenHeight-60)

	// Draw some animated shapes in the background
	for i := 0; i < 3; i++ {
		angle := g.time + float64(i)*2*math.Pi/3
		x := float32(screenWidth/2 + math.Cos(angle)*150)
		y := float32(screenHeight/2 + math.Sin(angle)*100)
		vector.DrawFilledCircle(screen, x, y, 10, color.RGBA{R: 100, G: 150, B: 255, A: 100}, true)
	}
}

// Enter is called when the scene becomes active.
func (g *GameScene) Enter() {
	// Reset ball position when entering
	g.ballX = screenWidth / 2
	g.ballY = screenHeight / 2
}

// Exit is called when the scene is being replaced/popped.
func (g *GameScene) Exit() {}
