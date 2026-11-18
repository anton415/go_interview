package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNonZeroFields(t *testing.T) {
	type Address struct{ City string }
	type profile struct {
		Name   string `json:"name"`
		Age    int    `json:"age,omitempty"`
		Secret string `json:"-"`
		note   string
	}
	type otherFields struct {
		Label string   `json:",omitempty"`
		Tags  []string `json:"tags,omitempty"`
		Dash  string   `json:"-,omitempty"`
		Address
	}

	p := profile{Name: "Alice", Secret: "token", note: "internal"}
	original := p
	address := Address{City: "Kazan"}
	tests := []struct {
		name  string
		input any
		want  map[string]any
	}{
		{"struct", p, map[string]any{"name": "Alice"}},
		{"pointer", &p, map[string]any{"name": "Alice"}},
		{"all zero fields", profile{}, map[string]any{}},
		{"name with options", profile{Name: "Alice", Age: 28}, map[string]any{"name": "Alice", "age": 28}},
		{"empty tag name", otherFields{Label: "Go"}, map[string]any{"Label": "Go"}},
		{"non-nil empty slice", otherFields{Tags: []string{}}, map[string]any{"tags": []string{}}},
		{"dash with options", otherFields{Dash: "value"}, map[string]any{"-": "value"}},
		{"embedded field without a tag", otherFields{Address: address}, map[string]any{"Address": address}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NonZeroFields(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
	assert.Equal(t, original, p)
}

func TestNonZeroFieldsErrors(t *testing.T) {
	type pair struct {
		Name  string
		Alias string `json:"Name,omitempty"`
	}
	p := &struct{}{}
	number := 42
	tests := []struct {
		name  string
		input any
	}{
		{"nil", nil},
		{"nil pointer", (*struct{})(nil)},
		{"pointer to pointer", &p},
		{"number", number},
		{"pointer to number", &number},
		{"duplicate name", pair{Name: "first", Alias: "second"}},
		{"duplicate with a zero field", pair{Alias: "second"}},
		{"duplicate with two zero fields", pair{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NonZeroFields(tt.input)
			assert.Error(t, err)
		})
	}
}
