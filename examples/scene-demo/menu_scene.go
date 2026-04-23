package main

import (
	"image/color"

	"github.com/dswisher/gamekit/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// MenuScene displays a simple menu with options to start the game or exit.
type MenuScene struct {
	controller scenes.SceneController
	selected   int
	options    []string
}

// NewMenuScene creates a new menu scene.
func NewMenuScene(controller scenes.SceneController) *MenuScene {
	return &MenuScene{
		controller: controller,
		selected:   0,
		options:    []string{"Start Game", "Exit"},
	}
}

// Update processes menu input.
func (m *MenuScene) Update(dt float64) {
	// Navigate menu with arrow keys (only on key press, not hold)
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.selected = (m.selected + 1) % len(m.options)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.selected--
		if m.selected < 0 {
			m.selected = len(m.options) - 1
		}
	}

	// Select with Enter (only on key press, not hold)
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch m.selected {
		case 0: // Start Game - push so we can pop back to menu
			m.controller.Push(NewGameScene(m.controller))
		case 1: // Exit
			m.controller.Pop()
		}
	}
}

// Draw renders the menu.
func (m *MenuScene) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(color.RGBA{R: 30, G: 30, B: 50, A: 255})

	// Draw title using ebiten's debug print
	ebitenutil.DebugPrintAt(screen, "MAIN MENU", screenWidth/2-40, 80)

	// Draw options
	for i, option := range m.options {
		y := 180 + i*50
		prefix := "  "
		if i == m.selected {
			prefix = "> "
		}
		text := prefix + option
		ebitenutil.DebugPrintAt(screen, text, screenWidth/2-50, y)
	}

	// Draw instructions
	ebitenutil.DebugPrintAt(screen, "Use Arrow Keys + Enter", screenWidth/2-80, screenHeight-60)
}

// Enter is called when the scene becomes active.
func (m *MenuScene) Enter() {
	m.selected = 0
}

// Exit is called when the scene is being replaced/popped.
func (m *MenuScene) Exit() {}
