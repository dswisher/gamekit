# ECS Demo

Interactive demonstration of the ECS (Entity Component System) package.

## What It Shows

This demo illustrates core ECS concepts:

- **Entities**: Lightweight IDs that represent game objects
- **Components**: Data-only structs (Position, Velocity, Size, Color, Lifetime)
- **Systems**: Logic that processes entities with specific components:
  - Movement System: Updates position based on velocity, handles bouncing
  - Lifetime System: Removes entities after they expire
  - Render System: Draws entities as colored rectangles

## Controls

| Key | Action |
|-----|--------|
| `SPACE` | Spawn 10 new bouncing entities (up to 500 max) |
| `C` | Clear all entities |
| `ESC` | Exit |

## Running

```bash
cd examples/ecs-demo
go run .
```

Each spawned entity has:
- Random position, size, and color
- Random velocity that bounces off screen edges
- A limited lifetime (10-20 seconds)

## Code Highlights

### Component Definitions

```go
type Position struct { X, Y float64 }
type Velocity struct { VX, VY float64 }
type Size struct { Width, Height float64 }
type Renderable struct { Color color.Color }
type Lifetime struct { Ticks, MaxTicks int }
```

### System Example: Movement

```go
positions := ecs.GetStore[Position](registry)
velocities := ecs.GetStore[Velocity](registry)

positions.Each(func(id ecs.EntityID, pos *Position) {
    if vel, ok := velocities.Get(id); ok {
        pos.X += vel.VX
        pos.Y += vel.VY
        // Bounce off edges...
    }
})
```

### Entity Creation

```go
entity := registry.CreateEntity()
ecs.SetComponent(registry, entity, Position{X: 100, Y: 200})
ecs.SetComponent(registry, entity, Velocity{VX: 2, VY: -1})
ecs.SetComponent(registry, entity, Size{Width: 20, Height: 20})
```
