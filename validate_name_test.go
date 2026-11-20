package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateNameSuccess(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"name", "Anton"},
		{"space", " "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)
			assert.NoError(t, err, "error type: %T", err)
		})
	}
}

func TestValidateNameEmpty(t *testing.T) {
	err := ValidateName("")

	var validationErr *ValidationError
	require.ErrorAs(t, err, &validationErr)
	require.NotNil(t, validationErr)
	assert.Equal(t, "name", validationErr.Field)
}
