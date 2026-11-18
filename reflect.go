package main

import (
	"fmt"
	"reflect"
	"strings"
)

func NonZeroFields(v any) (map[string]any, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil, fmt.Errorf("expected a non-nil pointer")
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %T", v)
	}

	result := make(map[string]any)
	seen := make(map[string]struct{})
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		if !field.IsExported() {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		if name == "" {
			name = field.Name
		}
		if _, exists := seen[name]; exists {
			return nil, fmt.Errorf("duplicate field name %q", name)
		}
		seen[name] = struct{}{}

		value := rv.Field(i)
		if !value.IsZero() {
			result[name] = value.Interface()
		}
	}
	return result, nil
}
