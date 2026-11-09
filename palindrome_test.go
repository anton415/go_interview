package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"aba", "aba", true},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ab", false},
		{"abcba", "abcba", true},
		{"abccba", "abccba", true},
		{"abcd", "abcd", false},
		{"топот", "топот", true},
		{"Аа", "Аа", false},
		{"топот ", "топот ", false},
		{"aba!", "aba!", false},
		{"e\u0301", "e\u0301", false},
		{"\xff", "\xff", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}