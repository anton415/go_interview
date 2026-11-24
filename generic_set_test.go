package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	type sample struct {
		number int
		word   string
		key    Key
	}
	first := sample{7, "go", Key{Region: "north", ID: 7}}
	second := sample{3, "java", Key{Region: "south", ID: 7}}

	tests := []struct {
		name     string
		nilSet   bool
		add      []sample
		remove   []sample
		expected []sample
	}{
		{name: "nil", nilSet: true},
		{name: "delete from nil", nilSet: true, remove: []sample{first}},
		{name: "add duplicate", add: []sample{first, first}, expected: []sample{first}},
		{name: "add distinct", add: []sample{first, second}, expected: []sample{first, second}},
		{name: "delete existing", add: []sample{first, second}, remove: []sample{first}, expected: []sample{second}},
		{name: "delete twice", add: []sample{first, second}, remove: []sample{first, first}, expected: []sample{second}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numbers Set[int]
			var words Set[string]
			var keys Set[Key]
			if !tt.nilSet {
				numbers = make(Set[int])
				words = make(Set[string])
				keys = make(Set[Key])
			}
			for _, value := range tt.add {
				numbers.Add(value.number)
				words.Add(value.word)
				keys.Add(value.key)
			}
			for _, value := range tt.remove {
				numbers.Delete(value.number)
				words.Delete(value.word)
				keys.Delete(value.key)
			}

			assert.Equal(t, len(tt.expected), numbers.Len(), "Set[int].Len()")
			assert.Equal(t, len(tt.expected), words.Len(), "Set[string].Len()")
			assert.Equal(t, len(tt.expected), keys.Len(), "Set[Key].Len()")
			for _, value := range []sample{first, second} {
				want := slices.Contains(tt.expected, value)
				assert.Equal(t, want, numbers.Contains(value.number), "Contains(%d)", value.number)
				assert.Equal(t, want, words.Contains(value.word), "Contains(%q)", value.word)
				assert.Equal(t, want, keys.Contains(value.key), "Contains(%v)", value.key)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	intTests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"duplicates", []int{2, 1, 2}, []int{2, 1}},
		{"order of first occurrence", []int{3, 1, 3, 2, 1}, []int{3, 1, 2}},
		{"already unique", []int{3, 1, 2}, []int{3, 1, 2}},
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.input)
			result := Unique(tt.input)
			assert.True(t, slices.Equal(tt.expected, result), "Unique(%v) = %v, ожидалось %v", original, result, tt.expected)
			assert.Equal(t, original, tt.input, "Unique изменила входной срез")

			clear(result)
			assert.Equal(t, original, tt.input, "изменение результата затронуло входной срез")
		})
	}

	stringTests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"strings", []string{"go", "java", "go"}, []string{"go", "java"}},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.input)
			result := Unique(tt.input)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, original, tt.input, "Unique изменила входной срез")

			clear(result)
			assert.Equal(t, original, tt.input, "изменение результата затронуло входной срез")
		})
	}
}
