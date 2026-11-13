package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationErrorWrapping(t *testing.T) {
	original := &ValidationError{
		Field:  "email",
		Reason: "empty",
	}

	wrapped := fmt.Errorf("register user: %w", original)
	assert.True(t, errors.Is(wrapped, original))

	var target *ValidationError
	if assert.True(t, errors.As(wrapped, &target)) {
		assert.Equal(t, "email", target.Field)
		assert.Equal(t, "empty", target.Reason)
	}

	other := &ValidationError{
		Field:  "email",
		Reason: "empty",
	}
	assert.False(t, errors.Is(wrapped, other))
}
