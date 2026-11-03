package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		value    int
		expected []int
	}{
		{"remove 2 from [1, 2, 3, 2, 4]", []int{1, 2, 3, 2, 4}, 2, []int{1, 3, 4}},
		{"remove 5 from [1, 2, 3]", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"remove -1 from [-1, -1, -1]", []int{-1, -1, -1}, -1, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveValue(tt.input, tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestKeepPositive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"keep positive from [1, -2, 3, 0, 4]", []int{1, -2, 3, 0, 4}, []int{1, 3, 4}},
		{"keep positive from [-1, -2, -3]", []int{-1, -2, -3}, []int{}},
		{"keep positive from [0, 0, 0]", []int{0, 0, 0}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeepPositive(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestColapseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"collapse 'aaafbbbb'", "aaafbbbb", "a3fb4"},
		{"collapse 'a'", "a", "a"},
		{"collapse 'aa'", "aa", "a2"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ColapseString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
