// ecs-demo demonstrates a simple Entity Component System with Ebitengine.
//
// Controls:
//   - SPACE: Spawn a new bouncing entity
//   - C: Clear all entities
//   - ESC: Exit
//
// The demo shows entities with Position, Velocity, Size, Color, and Lifetime
// components bouncing around the screen.
package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/dswisher/gamekit/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	maxEntities  = 500
)

// Component types
type Position struct {
	X, Y float64
}

type Velocity struct {
	VX, VY float64
}

type Size struct {
	Width, Height float64
}

type Renderable struct {
	Color color.Color
}

type Lifetime struct {
	Ticks    int
	MaxTicks int
}

// Game manages the ECS world and rendering
type Game struct {
	registry *ecs.Registry
}

func NewGame() *Game {
	return &Game{
		registry: ecs.NewRegistry(),
	}
}

func (g *Game) Update() error {
	// Exit on ESC
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// Spawn new entity on SPACE press
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// Spawn up to 10 entities per press, capped at maxEntities
		existing := g.countEntities()
		toSpawn := min(10, maxEntities-existing)
		for i := 0; i < toSpawn; i++ {
			g.spawnEntity()
		}
	}

	// Clear all entities on C press
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.clearEntities()
	}

	// Update systems
	g.updateMovement()
	g.updateLifetime()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Render all entities with Position, Size, and Renderable components
	positions := ecs.GetStore[Position](g.registry)
	sizes := ecs.GetStore[Size](g.registry)
	renderables := ecs.GetStore[Renderable](g.registry)

	positions.Each(func(id ecs.EntityID, pos *Position) {
		size, hasSize := sizes.Get(id)
		renderable, hasRenderable := renderables.Get(id)

		if hasSize && hasRenderable {
			// Draw the entity as a filled rectangle
			vector.DrawFilledRect(
				screen,
				float32(pos.X),
				float32(pos.Y),
				float32(size.Width),
				float32(size.Height),
				renderable.Color,
				true,
			)
		}
	})

	// Draw UI
	entityCount := g.countEntities()
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"ECS Demo - Entities: %d/%d\nSPACE: Spawn 10 | C: Clear | ESC: Exit",
		entityCount, maxEntities,
	))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// spawnEntity creates a new bouncing entity with random properties
func (g *Game) spawnEntity() {
	entity := g.registry.CreateEntity()

	// Random position within screen bounds
	x := rand.Float64() * (screenWidth - 40)
	y := rand.Float64() * (screenHeight - 40)

	// Random velocity
	vx := (rand.Float64() - 0.5) * 4
	vy := (rand.Float64() - 0.5) * 4

	// Random size
	size := 10 + rand.Float64()*30

	// Random color
	c := color.RGBA{
		R: uint8(rand.Intn(200) + 55),
		G: uint8(rand.Intn(200) + 55),
		B: uint8(rand.Intn(200) + 55),
		A: 255,
	}

	// Lifetime: 10-20 seconds at 60 FPS
	lifetime := 600 + rand.Intn(600)

	ecs.SetComponent(g.registry, entity, Position{X: x, Y: y})
	ecs.SetComponent(g.registry, entity, Velocity{VX: vx, VY: vy})
	ecs.SetComponent(g.registry, entity, Size{Width: size, Height: size})
	ecs.SetComponent(g.registry, entity, Renderable{Color: c})
	ecs.SetComponent(g.registry, entity, Lifetime{Ticks: lifetime, MaxTicks: lifetime})
}

// updateMovement moves entities and handles screen edge bouncing
func (g *Game) updateMovement() {
	positions := ecs.GetStore[Position](g.registry)
	velocities := ecs.GetStore[Velocity](g.registry)
	sizes := ecs.GetStore[Size](g.registry)

	positions.Each(func(id ecs.EntityID, pos *Position) {
		vel, hasVel := velocities.Get(id)
		size, hasSize := sizes.Get(id)

		if !hasVel || !hasSize {
			return
		}

		// Update position
		pos.X += vel.VX
		pos.Y += vel.VY

		// Bounce off edges
		if pos.X <= 0 {
			pos.X = 0
			vel.VX = -vel.VX
			velocities.Set(id, vel)
		} else if pos.X+size.Width >= screenWidth {
			pos.X = screenWidth - size.Width
			vel.VX = -vel.VX
			velocities.Set(id, vel)
		}

		if pos.Y <= 0 {
			pos.Y = 0
			vel.VY = -vel.VY
			velocities.Set(id, vel)
		} else if pos.Y+size.Height >= screenHeight {
			pos.Y = screenHeight - size.Height
			vel.VY = -vel.VY
			velocities.Set(id, vel)
		}
	})
}

// updateLifetime removes entities whose lifetime has expired
func (g *Game) updateLifetime() {
	lifetimes := ecs.GetStore[Lifetime](g.registry)

	// Collect entities to destroy
	toDestroy := make([]ecs.EntityID, 0)

	lifetimes.Each(func(id ecs.EntityID, life *Lifetime) {
		life.Ticks--
		if life.Ticks <= 0 {
			toDestroy = append(toDestroy, id)
		}
	})

	// Destroy expired entities
	for _, id := range toDestroy {
		g.registry.DestroyEntity(id)
	}
}

// countEntities returns the number of entities with Position component
func (g *Game) countEntities() int {
	count := 0
	ecs.GetStore[Position](g.registry).Each(func(id ecs.EntityID, pos *Position) {
		count++
	})
	return count
}

// clearEntities destroys all entities
func (g *Game) clearEntities() {
	positions := ecs.GetStore[Position](g.registry)

	// Collect all entity IDs
	toDestroy := make([]ecs.EntityID, 0)
	positions.Each(func(id ecs.EntityID, pos *Position) {
		toDestroy = append(toDestroy, id)
	})

	// Destroy them all
	for _, id := range toDestroy {
		g.registry.DestroyEntity(id)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("ECS Demo")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
