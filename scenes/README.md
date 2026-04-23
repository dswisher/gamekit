# Scenes

[![Go Reference](https://pkg.go.dev/badge/github.com/dswisher/gamekit/scenes.svg)](https://pkg.go.dev/github.com/dswisher/gamekit/scenes) ![Version](https://img.shields.io/github/v/tag/dswisher/gamekit?filter=scenes/*&label=)

Scene management library for Ebitengine games.

## Installation

```bash
go get github.com/dswisher/gamekit/scenes
```

## Features

- Stack-based scene management
- Clean scene lifecycle (Enter, Exit, Update, Draw)
- Easy scene navigation (Push, Pop, Replace)
- Interface-based design for testability

## Usage

```go
import "github.com/dswisher/gamekit/scenes"

// Create a scene manager
manager := scenes.NewSceneManager()

// Push your initial scene
manager.Push(NewMenuScene(manager))

// In your ebiten.Game Update:
if current := manager.Current(); current != nil {
    current.Update(dt)
}

// In your ebiten.Game Draw:
if current := manager.Current(); current != nil {
    current.Draw(screen)
}
```

## Example Scene

```go
type MenuScene struct {
    controller scenes.SceneController
}

func NewMenuScene(controller scenes.SceneController) *MenuScene {
    return &MenuScene{controller: controller}
}

func (m *MenuScene) Update(dt float64) {
    if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
        m.controller.Push(NewGameScene(m.controller))
    }
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
    // Draw your scene
}

func (m *MenuScene) Enter() {
    // Called when scene becomes active
}

func (m *MenuScene) Exit() {
    // Called when scene is replaced/popped
}
```

## Related Examples

- [`scene-demo`](../examples/scene-demo/) - Two-scene demo showing menu and game navigation

## License

MIT
