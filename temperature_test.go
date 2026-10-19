package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestParseCelsius(t *testing.T) {
	tests := []struct {
		raw      string
		expected int16
		err      error
	}{
		{raw: "25", expected: 25, err: nil},
		{raw: "32768", expected: 0, err: strconv.ErrRange},
		{raw: "cold", expected: 0, err: strconv.ErrSyntax},
		{raw: "-32769", expected: 0, err: strconv.ErrRange},
		{raw: "0", expected: 0, err: nil},
		{raw: "-32768", expected: -32768, err: nil},
		{raw: "32767", expected: 32767, err: nil},
	}

	for _, tt := range tests {
		t.Run(tt.raw, func(t *testing.T) {
			value, err := ParseCelsius(tt.raw)
			if value != tt.expected {
				t.Errorf("ParseCelsius(%q) = %d; want %d", tt.raw, value, tt.expected)
			}
			if !errors.Is(err, tt.err) {
				t.Errorf("ParseCelsius(%q) error = %v; want %v", tt.raw, err, tt.err)
			}
		})
	}
}

func TestToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius  int16
		expected float64
	}{
		{celsius: 25, expected: 77.0},
		{celsius: 0, expected: 32.0},
		{celsius: -40, expected: -40.0},
		{celsius: 100, expected: 212.0},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(int64(tt.celsius), 10), func(t *testing.T) {
			result := ToFahrenheit(tt.celsius)
			if result != tt.expected {
				t.Errorf("ToFahrenheit(%d) = %f; want %f", tt.celsius, result, tt.expected)
			}
		})
	}
}
