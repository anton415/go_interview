package main

import (
	"fmt"
)

type ValidationError struct {
	Field  string
	Reason string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", v.Field, v.Reason)
}

func handle() error {
	return &ValidationError{
		Field:  "email",
		Reason: "empty",
	}
}

func registerUser() error {
	handErr := handle()
	if handErr != nil {
		return fmt.Errorf("register user: %w", handErr)
	}
	return nil
}
