package main

type UserID int

type Key struct {
	Region string
	ID     int
}

type Report struct {
	Values []int
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Delete(key T) {
	delete(s, key)
}

func (s Set[T]) Len() int {
	return len(s)
}

func Unique[T comparable](items []T) []T {
	seen := make(Set[T])
	var result []T

	for _, value := range items {
		if seen.Contains(value) {
			continue
		}
		seen.Add(value)
		result = append(result, value)
	}
	return result
}
