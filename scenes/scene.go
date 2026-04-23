package scenes

import "github.com/hajimehoshi/ebiten/v2"

// Scene defines the lifecycle of a scene in a game.
type Scene interface {
	// Update processes one frame of logic
	Update(dt float64)

	// Draw renders the scene
	Draw(screen *ebiten.Image)

	// Enter is called when scene becomes active
	Enter()

	// Exit is called when scene is being replaced/popped
	Exit()
}

// SceneController provides scene navigation methods for scenes to use.
// Scenes receive this interface (rather than *SceneManager) to allow
// for easier testing with mock implementations.
type SceneController interface {
	// Push adds a new scene on top of the stack
	Push(sc Scene)

	// Pop removes the current scene from the stack
	Pop()

	// Replace pops the current scene and pushes a new one
	Replace(sc Scene)
}
