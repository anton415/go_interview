package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBag(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{value: "iPhone", expected: true},
		{value: "headphones", expected: true},
		{value: "pencil", expected: true},
		{value: "", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			var bag Bag
			assert.Equal(t, 0, bag.Len(), "начальная длина должна быть 0")
			assert.False(t, bag.Contains(tt.value), "вещь не должна быть в сумке %s ", tt.value)
			bag.Add(tt.value)
			result := bag.Contains(tt.value)
			assert.Equal(t, tt.expected, result, "Contains(%s) = %v; want %v", tt.value, result, tt.expected)
			assert.Equal(t, 1, bag.Len(), "после добавления длина должна быть 1")
		})
	}
}