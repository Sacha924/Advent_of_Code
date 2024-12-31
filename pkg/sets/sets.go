package sets

import (
	"maps"
)

// NOTE :
// The current implementation isn't thread-safe. If the Set will be accessed concurrently,
// consider adding a sync.RWMutex to protect reads and writes.

type Set[T comparable] map[T]struct{}

func SetOf[T comparable](s ...T) Set[T] {
	return SetFrom(s)
}

// Slice to Set
func SetFrom[T comparable](s []T) Set[T] {
	set := make(Set[T])
	for _, v := range s {
		set.Add(v)
	}
	return set
}

// Set to Slice
func (s *Set[T]) Slice() []T {
	l := make([]T, 0, len(*s))
	for i := range *s {
		l = append(l, i)
	}
	return l
}

func (s *Set[T]) Add(v T) {
	if *s == nil {
		*s = make(map[T]struct{})
	}
	(*s)[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(*s, v)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *Set[T]) Clone() Set[T] {
	return maps.Clone(*s)
}

func (s *Set[T]) Seq(yield func(T) bool) {
	for v, _ := range *s {
		if !yield(v) {
			return
		}
	}
}

func (s *Set[T]) Intersect(other Set[T]) []T {
	intersect := make([]T, 0)
	for v, _ := range *s {
		if other.Contains(v) {
			intersect = append(intersect, v)
		}
	}
	return intersect
}
