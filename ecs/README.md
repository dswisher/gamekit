# ECS

[![Go Reference](https://pkg.go.dev/badge/github.com/dswisher/gamekit/ecs.svg)](https://pkg.go.dev/github.com/dswisher/gamekit/ecs) ![Version](https://img.shields.io/github/v/tag/dswisher/gamekit?filter=ecs/*&label=)

A simple Entity Component System (ECS) library for Go.

## Installation

```bash
go get github.com/dswisher/gamekit/ecs
```

## Features

- **Type-safe component storage** using Go generics
- **Lightweight entities** - just uint64 IDs
- **Efficient querying** with typed component stores
- **Flexible iteration** with `Each` and `EachSorted`
- **Zero external dependencies** for the core package

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/dswisher/gamekit/ecs"
)

// Define your component types
type Position struct {
    X, Y float64
}

type Velocity struct {
    VX, VY float64
}

func main() {
    // Create a registry to manage entities and components
    registry := ecs.NewRegistry()

    // Create an entity
    player := registry.CreateEntity()

    // Add components to the entity
    ecs.SetComponent(registry, player, Position{X: 10, Y: 20})
    ecs.SetComponent(registry, player, Velocity{VX: 1, VY: 0.5})

    // Query a component
    pos, ok := ecs.GetComponent[Position](registry, player)
    if ok {
        fmt.Printf("Player is at (%f, %f)\n", pos.X, pos.Y)
    }

    // Check if an entity has a component
    if ecs.HasComponent[Velocity](registry, player) {
        fmt.Println("Player has velocity!")
    }

    // Modify a component
    ecs.ModifyComponent(registry, player, func(v *Velocity) {
        v.VX *= 2
    })

    // Remove a component
    ecs.RemoveComponent[Velocity](registry, player)

    // Destroy the entity (removes all components)
    registry.DestroyEntity(player)
}
```

### Iterating Over Components

Use `Each` to process all entities with a specific component:

```go
// Move all entities with Position and Velocity
positions := ecs.GetStore[Position](registry)
velocities := ecs.GetStore[Velocity](registry)

positions.Each(func(id ecs.EntityID, pos *Position) {
    if vel, ok := velocities.Get(id); ok {
        pos.X += vel.VX
        pos.Y += vel.VY
    }
})
```

Use `EachSorted` for ordered iteration:

```go
// Process entities by Y position (e.g., for rendering order)
positions.EachSorted(
    func(id ecs.EntityID, pos *Position) int {
        return int(pos.Y) // Sort by Y coordinate
    },
    func(id ecs.EntityID, pos *Position) {
        // Draw entity at pos.X, pos.Y
    },
)
```

### Tag Components

Use empty structs as tags:

```go
type PlayerControlled struct{}
type Enemy struct{}

// Add tag to entity
ecs.SetComponent(registry, entity, PlayerControlled{})

// Check for tag
if ecs.HasComponent[PlayerControlled](registry, entity) {
    // Process player input
}
```

## Architecture

The ECS follows a traditional architecture:

- **Entity**: A unique identifier (`EntityID` - uint64)
- **Component**: Plain data structs (position, velocity, health, etc.)
- **System**: Logic that operates on entities with specific components (you implement these)
- **Registry**: Central manager for entities and component storage

## Design Philosophy

This ECS prioritizes **simplicity and type safety** over raw performance:

- Uses Go generics for compile-time type safety
- No query caching or complex filtering (yet)
- Simple map-based storage
- Explicit iteration patterns

For high-performance scenarios with thousands of entities, you may need a more specialized ECS. This library is designed for small-to-medium games where code clarity matters.

## Related Examples

- [`ecs-demo`](../examples/ecs-demo/) - Interactive demo showing entity spawning, movement, and rendering

## License

MIT
