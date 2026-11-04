package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencies(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{"{go, java, go}", []string{"go", "java", "go"}, map[string]int{"go": 2, "java": 1}},
		{"nil", nil, map[string]int{}},
		{"empty", []string{}, map[string]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Frequencies(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMostFrequent(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		exp_word string
		exp_count int
		exp_ok bool
	}{
		{"{go 2, java 1}", map[string]int{"go": 2, "java": 1}, "go", 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word, count, ok := MostFrequent(tt.input)
			assert.Equal(t, tt.exp_word, word)
			assert.Equal(t, tt.exp_count, count)
			assert.Equal(t, tt.exp_ok, ok)
		})
	}
}
