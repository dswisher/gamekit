package ecs

import "fmt"

// Registry manages entities and provides access to typed component
// stores. It is the central data structure of the ECS.
type Registry struct {
	nextID   EntityID
	alive    map[EntityID]bool
	stores   map[string]any
	cleaners []func(EntityID)
}

// NewRegistry creates an empty Registry ready for use.
func NewRegistry() *Registry {
	return &Registry{
		nextID: 1,
		alive:  make(map[EntityID]bool),
		stores: make(map[string]any),
	}
}

// CreateEntity allocates a new entity and returns its ID.
func (r *Registry) CreateEntity() EntityID {
	id := r.nextID
	r.nextID++
	r.alive[id] = true
	return id
}

// DestroyEntity removes an entity and all of its components from
// every registered component store.
func (r *Registry) DestroyEntity(id EntityID) {
	for _, clean := range r.cleaners {
		clean(id)
	}
	delete(r.alive, id)
}

// IsAlive reports whether an entity exists in the registry.
func (r *Registry) IsAlive(id EntityID) bool {
	return r.alive[id]
}

// GetStore returns the ComponentStore for a given component type,
// creating it if it does not already exist.
//
// This is a package-level generic function rather than a method on
// Registry because Go does not support generic methods.
func GetStore[T any](r *Registry) *ComponentStore[T] {
	key := storeKey[T]()
	if s, ok := r.stores[key]; ok {
		return s.(*ComponentStore[T])
	}
	s := NewComponentStore[T]()
	r.stores[key] = s
	r.cleaners = append(r.cleaners, func(id EntityID) {
		s.Remove(id)
	})
	return s
}

// storeKey returns a unique string key for a component type, used to
// look up the correct store in the registry's map.
func storeKey[T any]() string {
	var zero T
	return fmt.Sprintf("%T", zero)
}

// Convenience functions that wrap GetStore for common operations.

// SetComponent attaches a component to an entity.
func SetComponent[T any](r *Registry, id EntityID, c T) {
	GetStore[T](r).Set(id, c)
}

// GetComponent retrieves a component from an entity. The second
// return value reports whether the entity has this component.
func GetComponent[T any](r *Registry, id EntityID) (T, bool) {
	return GetStore[T](r).Get(id)
}

// HasComponent reports whether an entity has a given component type.
func HasComponent[T any](r *Registry, id EntityID) bool {
	return GetStore[T](r).Has(id)
}

// RemoveComponent detaches a component from an entity.
func RemoveComponent[T any](r *Registry, id EntityID) {
	GetStore[T](r).Remove(id)
}

// ModifyComponent updates a component in place via a callback.
// Returns true if the entity had the component and was modified, false otherwise.
func ModifyComponent[T any](r *Registry, id EntityID, fn func(*T)) bool {
	return GetStore[T](r).Modify(id, fn)
}

// FirstEntity returns the first entity that has the given component type.
// The second return value reports whether any entity was found.
func FirstEntity[T any](r *Registry) (EntityID, bool) {
	return GetStore[T](r).First()
}

// FirstEntityInList returns the first entity from a list that has the given component type.
// The component and a found flag are returned.
func FirstEntityInList[T any](r *Registry, ids []EntityID) (EntityID, T, bool) {
	c, found := GetStore[T](r).FirstInList(ids)
	var zero T
	if !found {
		return 0, zero, false
	}
	// We need to find the entity ID - iterate through the list to find it
	for _, id := range ids {
		if GetStore[T](r).Has(id) {
			return id, c, true
		}
	}
	return 0, zero, false
}
