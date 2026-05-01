package ecs

import "slices"

// ComponentStore holds all instances of a single component type,
// keyed by EntityID for fast lookup.
type ComponentStore[T any] struct {
	components map[EntityID]T
}

// NewComponentStore creates an empty ComponentStore for the given
// component type.
func NewComponentStore[T any]() *ComponentStore[T] {
	return &ComponentStore[T]{
		components: make(map[EntityID]T),
	}
}

// Set attaches a component to an entity, replacing any existing
// component of this type on that entity.
func (s *ComponentStore[T]) Set(id EntityID, c T) {
	s.components[id] = c
}

// Get retrieves the component for an entity. The second return value
// reports whether the entity has this component.
func (s *ComponentStore[T]) Get(id EntityID) (T, bool) {
	c, ok := s.components[id]
	return c, ok
}

// Has reports whether an entity has this component.
func (s *ComponentStore[T]) Has(id EntityID) bool {
	_, ok := s.components[id]
	return ok
}

// Remove detaches this component from an entity. It is safe to call
// even if the entity does not have this component.
func (s *ComponentStore[T]) Remove(id EntityID) {
	delete(s.components, id)
}

// Each calls fn for every entity that has this component. The
// component is passed by pointer so the callback can modify it in
// place.
func (s *ComponentStore[T]) Each(fn func(EntityID, *T)) {
	for id := range s.components {
		c := s.components[id]
		fn(id, &c)
		s.components[id] = c
	}
}

// EachSorted calls fn for every entity that has this component, in the
// order determined by the key function. Entities are sorted by ascending
// key value. The component is passed by pointer so the callback can
// modify it in place.
func (s *ComponentStore[T]) EachSorted(key func(EntityID, *T) int, fn func(EntityID, *T)) {
	type entry struct {
		id  EntityID
		val T
		k   int
	}

	entries := make([]entry, 0, len(s.components))
	for id, c := range s.components {
		entries = append(entries, entry{id, c, key(id, &c)})
	}

	slices.SortStableFunc(entries, func(a, b entry) int {
		return a.k - b.k
	})

	for i := range entries {
		fn(entries[i].id, &entries[i].val)
		s.components[entries[i].id] = entries[i].val
	}
}

// First returns the first entity that has the specified component
// The second return value reports whether any entity was found.
func (s *ComponentStore[T]) First() (EntityID, bool) {
	for id := range s.components {
		return id, true
	}
	return 0, false
}

// FirstInList returns the first entity from a list that has this component.
// The component and a found flag are returned.
func (s *ComponentStore[T]) FirstInList(ids []EntityID) (T, bool) {
	for _, id := range ids {
		if c, ok := s.components[id]; ok {
			return c, true
		}
	}
	var zero T
	return zero, false
}

// Any returns true if the store contains 1 or more entities.
func (s *ComponentStore[T]) Any() bool {
	return len(s.components) > 0
}

// Modify applies fn to the entity's component, if it exists.
func (s *ComponentStore[T]) Modify(id EntityID, fn func(*T)) bool {
	c, ok := s.components[id]
	if !ok {
		return false
	}
	fn(&c)
	s.components[id] = c
	return true
}
