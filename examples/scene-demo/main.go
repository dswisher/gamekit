package main

import (
	"log"

	"github.com/dswisher/gamekit/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Game is the main game struct that implements ebiten.Game.
type Game struct {
	manager *scenes.SceneManager
}

// NewGame creates a new game with a scene manager.
func NewGame() *Game {
	manager := scenes.NewSceneManager()

	// Start with the menu scene
	manager.Push(NewMenuScene(manager))

	return &Game{manager: manager}
}

// Update processes one frame of game logic.
func (g *Game) Update() error {
	// Calculate delta time
	dt := 1.0 / ebiten.ActualTPS()

	// Update the current scene
	if current := g.manager.Current(); current != nil {
		current.Update(dt)
	}

	// Exit if no scenes left (e.g., user selected Exit from menu)
	if g.manager.Current() == nil {
		return ebiten.Termination
	}

	return nil
}

// Draw renders the game.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the current scene
	if current := g.manager.Current(); current != nil {
		current.Draw(screen)
	}
}

// Layout returns the screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Scene Demo - Press ESC to exit")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
