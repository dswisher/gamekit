# Gamekit - Game Development Toolkit for Go

A collection of modular packages for building games with Go and Ebitengine.

## Packages

| Package | Docs | Latest |
|---------|------|--------|
| [scenes](./scenes/) - Scene management for Ebitengine games | [![Go Reference](https://pkg.go.dev/badge/github.com/dswisher/gamekit/scenes.svg)](https://pkg.go.dev/github.com/dswisher/gamekit/scenes) | ![Version](https://img.shields.io/github/v/tag/dswisher/gamekit?filter=scenes/*&label=)
| [sprites](./sprites/) - Sprite sheet management and animation for Ebitengine | [![Go Reference](https://pkg.go.dev/badge/github.com/dswisher/gamekit/sprites.svg)](https://pkg.go.dev/github.com/dswisher/gamekit/sprites) | ![Version](https://img.shields.io/github/v/tag/dswisher/gamekit?filter=sprites/*&label=)
| [ecs](./ecs/) - Entity Component System for Go | [![Go Reference](https://pkg.go.dev/badge/github.com/dswisher/gamekit/ecs.svg)](https://pkg.go.dev/github.com/dswisher/gamekit/ecs) | ![Version](https://img.shields.io/github/v/tag/dswisher/gamekit?filter=ecs/*&label=)
| [spritetool](./cmd/spritetool/) - CLI tool for building sprite sheets | - | - |

## Examples

See the [examples](./examples/) directory for usage examples.

## Development

This repo uses a Go workspace for local development:

```bash
go work sync
```

## License

MIT

