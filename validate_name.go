package main

func ValidateName(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Reason: "empty"}
	}
	return nil
}
