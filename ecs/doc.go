// Package ecs provides a simple entity component system (ECS) for Go.
//
// An ECS architecture separates data (components) from behavior (systems)
// using entities as lightweight identifiers. This package provides type-safe
// component storage via generics.
//
// Basic usage:
//
//	registry := ecs.NewRegistry()
//	entity := registry.CreateEntity()
//	ecs.SetComponent(registry, entity, Position{X: 10, Y: 20})
//	ecs.SetComponent(registry, entity, Velocity{DX: 1, DY: 1})
//
//	// Query and iterate
//	ecs.GetStore[Position](registry).Each(func(id ecs.EntityID, pos *Position) {
//	    pos.X += 1
//	})
//
// The Registry manages entity lifecycle and provides access to typed
// component stores. ComponentStore offers efficient storage, querying,
// and iteration with optional sorting.
package ecs
