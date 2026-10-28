package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestAccumulator(t *testing.T) {
	tests := []struct {
		value    int
		expected int
	}{
		{value: 2, expected: 2},
		{value: -1, expected: 1},
		{value: 5, expected: 6},
	}

	acc := Accumulator()

	for _, tt := range tests {
		t.Run(fmt.Sprintf("value: %d", tt.value), func(t *testing.T) {
			result := acc(tt.value)
			assert.Equal(t, tt.expected, result, "Accumulator(%d) = %d; want %d", tt.value, result, tt.expected)
		})
	}
}

func TestReaders(t *testing.T) {
	tests := []struct {
		values   []int
		expected []int
	}{
		{values: []int{10, 20}, expected: []int{10, 20}},
		{values: nil, expected: nil},
		{values: []int{}, expected: []int{}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("values: %v", tt.values), func(t *testing.T) {
			readers := Readers(tt.values)
			for i, read := range readers {
				assert.Equal(t, tt.expected[i], read())
			}
		})
	}
}