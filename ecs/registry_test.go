package ecs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dswisher/gamekit/ecs"
)

// Test component types used across tests.
type Position struct {
	X, Y int
}

type Name struct {
	DisplayName string
}

type PlayerControlled struct{}

// --- Entity tests ---

func TestCreateEntity_ReturnsUniqueIDs(t *testing.T) {
	r := ecs.NewRegistry()

	e1 := r.CreateEntity()
	e2 := r.CreateEntity()
	e3 := r.CreateEntity()

	assert.NotEqual(t, e1, e2, "e1 and e2 should be unique")
	assert.NotEqual(t, e2, e3, "e2 and e3 should be unique")
	assert.NotEqual(t, e1, e3, "e1 and e3 should be unique")
}

func TestCreateEntity_IsAlive(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	assert.True(t, r.IsAlive(e), "expected entity %d to be alive", e)
}

func TestDestroyEntity_IsNotAlive(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	r.DestroyEntity(e)

	assert.False(t, r.IsAlive(e), "expected entity %d to be dead after destroy", e)
}

func TestIsAlive_NonExistentEntity(t *testing.T) {
	r := ecs.NewRegistry()

	assert.False(t, r.IsAlive(999), "expected non-existent entity to not be alive")
}

// --- Component CRUD tests ---

func TestSetAndGetComponent(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 5, Y: 10})

	pos, ok := ecs.GetComponent[Position](r, e)
	assert.True(t, ok, "expected entity to have Position component")
	assert.Equal(t, 5, pos.X, "pos.X should be 5")
	assert.Equal(t, 10, pos.Y, "pos.Y should be 10")
}

func TestGetComponent_Missing(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	_, ok := ecs.GetComponent[Position](r, e)
	assert.False(t, ok, "expected entity without Position to return ok=false")
}

func TestHasComponent(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	assert.False(t, ecs.HasComponent[Position](r, e), "expected HasComponent to return false before setting")

	ecs.SetComponent(r, e, Position{X: 1, Y: 2})

	assert.True(t, ecs.HasComponent[Position](r, e), "expected HasComponent to return true after setting")
}

func TestSetComponent_Overwrites(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 1, Y: 2})
	ecs.SetComponent(r, e, Position{X: 99, Y: 100})

	pos, ok := ecs.GetComponent[Position](r, e)
	assert.True(t, ok, "expected entity to have Position component")
	assert.Equal(t, 99, pos.X, "pos.X should be 99")
	assert.Equal(t, 100, pos.Y, "pos.Y should be 100")
}

func TestRemoveComponent(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 1, Y: 2})
	ecs.RemoveComponent[Position](r, e)

	assert.False(t, ecs.HasComponent[Position](r, e), "expected Position to be removed")
}

func TestRemoveComponent_Noop(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	// Should not panic when removing a component that was never added.
	ecs.RemoveComponent[Position](r, e)
}

// --- Multiple component types ---

func TestMultipleComponentTypes(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 3, Y: 7})
	ecs.SetComponent(r, e, Name{DisplayName: "Hero"})
	ecs.SetComponent(r, e, PlayerControlled{})

	pos, ok := ecs.GetComponent[Position](r, e)
	assert.True(t, ok, "expected entity to have Position component")
	assert.Equal(t, 3, pos.X, "pos.X should be 3")
	assert.Equal(t, 7, pos.Y, "pos.Y should be 7")

	name, ok := ecs.GetComponent[Name](r, e)
	assert.True(t, ok, "expected entity to have Name component")
	assert.Equal(t, "Hero", name.DisplayName, "DisplayName should be Hero")

	assert.True(t, ecs.HasComponent[PlayerControlled](r, e), "expected entity to have PlayerControlled tag")
}

// --- Each iterator ---

func TestEach_IteratesAllEntities(t *testing.T) {
	r := ecs.NewRegistry()

	e1 := r.CreateEntity()
	e2 := r.CreateEntity()
	e3 := r.CreateEntity()

	ecs.SetComponent(r, e1, Position{X: 1, Y: 0})
	ecs.SetComponent(r, e2, Position{X: 2, Y: 0})
	ecs.SetComponent(r, e3, Position{X: 3, Y: 0})

	visited := make(map[ecs.EntityID]int)
	ecs.GetStore[Position](r).Each(func(id ecs.EntityID, pos *Position) {
		visited[id] = pos.X
	})

	assert.Len(t, visited, 3, "expected 3 entities")
	assert.Equal(t, 1, visited[e1], "e1 should have X=1")
	assert.Equal(t, 2, visited[e2], "e2 should have X=2")
	assert.Equal(t, 3, visited[e3], "e3 should have X=3")
}

func TestEach_CanModifyComponents(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 5, Y: 10})

	ecs.GetStore[Position](r).Each(func(id ecs.EntityID, pos *Position) {
		pos.X += 1
		pos.Y += 1
	})

	pos, ok := ecs.GetComponent[Position](r, e)
	assert.True(t, ok, "expected entity to have Position component")
	assert.Equal(t, 6, pos.X, "pos.X should be 6")
	assert.Equal(t, 11, pos.Y, "pos.Y should be 11")
}

func TestEach_EmptyStore(t *testing.T) {
	r := ecs.NewRegistry()
	count := 0

	ecs.GetStore[Position](r).Each(func(id ecs.EntityID, pos *Position) {
		count++
	})

	assert.Zero(t, count, "expected 0 iterations on empty store")
}

// --- EachSorted iterator ---

func TestEachSorted_IteratesInSortedOrder(t *testing.T) {
	r := ecs.NewRegistry()

	e1 := r.CreateEntity()
	e2 := r.CreateEntity()
	e3 := r.CreateEntity()

	ecs.SetComponent(r, e1, Position{X: 1, Y: 0})
	ecs.SetComponent(r, e2, Position{X: 2, Y: 0})
	ecs.SetComponent(r, e3, Position{X: 3, Y: 0})

	order := make([]ecs.EntityID, 0, 3)
	ecs.GetStore[Position](r).EachSorted(
		func(id ecs.EntityID, pos *Position) int {
			return pos.X
		},
		func(id ecs.EntityID, _ *Position) {
			order = append(order, id)
		},
	)

	assert.Len(t, order, 3, "expected 3 entities")
	assert.Equal(t, e1, order[0], "first should be entity with X=1")
	assert.Equal(t, e2, order[1], "second should be entity with X=2")
	assert.Equal(t, e3, order[2], "third should be entity with X=3")
}

func TestEachSorted_EmptyStore(t *testing.T) {
	r := ecs.NewRegistry()
	count := 0

	ecs.GetStore[Position](r).EachSorted(
		func(id ecs.EntityID, pos *Position) int { return pos.X },
		func(id ecs.EntityID, pos *Position) { count++ },
	)

	assert.Zero(t, count, "expected 0 iterations on empty store")
}

// --- DestroyEntity cleans up components ---

func TestDestroyEntity_RemovesAllComponents(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 1, Y: 2})
	ecs.SetComponent(r, e, Name{DisplayName: "Goblin"})

	r.DestroyEntity(e)

	assert.False(t, ecs.HasComponent[Position](r, e), "expected Position to be removed after DestroyEntity")
	assert.False(t, ecs.HasComponent[Name](r, e), "expected Name to be removed after DestroyEntity")
}

func TestDestroyEntity_DoesNotAffectOtherEntities(t *testing.T) {
	r := ecs.NewRegistry()

	e1 := r.CreateEntity()
	e2 := r.CreateEntity()

	ecs.SetComponent(r, e1, Position{X: 1, Y: 1})
	ecs.SetComponent(r, e2, Position{X: 2, Y: 2})

	r.DestroyEntity(e1)

	assert.False(t, ecs.HasComponent[Position](r, e1), "expected e1 Position to be removed")

	pos, ok := ecs.GetComponent[Position](r, e2)
	assert.True(t, ok, "expected e2 to still have Position")
	assert.Equal(t, 2, pos.X, "e2 pos.X should be 2")
	assert.Equal(t, 2, pos.Y, "e2 pos.Y should be 2")
}

// --- GetStore returns the same store on repeated calls ---

func TestGetStore_ReturnsSameInstance(t *testing.T) {
	r := ecs.NewRegistry()

	s1 := ecs.GetStore[Position](r)
	s2 := ecs.GetStore[Position](r)

	assert.Same(t, s1, s2, "expected GetStore to return the same instance on repeated calls")
}

// --- ModifyComponent can actually modify components

func TestModifyComponent_CanModifyComponents(t *testing.T) {
	r := ecs.NewRegistry()
	e := r.CreateEntity()

	ecs.SetComponent(r, e, Position{X: 5, Y: 10})

	ecs.ModifyComponent(r, e, func(p *Position) {
		p.X += 1
		p.Y += 1
	})

	pos, _ := ecs.GetComponent[Position](r, e)
	assert.Equal(t, 6, pos.X, "pos.X should be 6")
	assert.Equal(t, 11, pos.Y, "pos.Y should be 11")
}
