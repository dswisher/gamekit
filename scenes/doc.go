// Package scenes provides scene management for Ebiten-based games.
//
// A scene represents a distinct game state (e.g., main menu, gameplay, pause screen).
// The SceneManager maintains a stack of scenes, allowing for modal overlays
// and easy navigation between game states.
//
// Basic usage:
//
//	manager := scenes.NewSceneManager()
//	manager.Push(NewMainMenuScene(manager))
//
//	// In your ebiten.Game Update:
//	if current := manager.Current(); current != nil {
//	    current.Update(dt)
//	}
//
//	// In your ebiten.Game Draw:
//	if current := manager.Current(); current != nil {
//	    current.Draw(screen)
//	}
//
// Scenes should receive a SceneController in their constructor for navigation:
//
//	type MenuScene struct {
//	    controller scenes.SceneController
//	}
//
//	func NewMenuScene(controller scenes.SceneController) *MenuScene {
//	    return &MenuScene{controller: controller}
//	}
//
//	func (m *MenuScene) startGame() {
//	    m.controller.Replace(NewGameScene(m.controller))
//	}
package scenes
