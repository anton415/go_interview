package main

import (
	"fmt"
	"strconv"
)

const (
	minInt16 = -1 << 15
	maxInt16 = 1<<15 - 1
)

func ParseCelsius(raw string) (int16, error) {
	n, err := strconv.ParseInt(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("ParseCelsius: parse %q as int16: %w", raw, err)
	}
	if n < minInt16 || n > maxInt16 {
		return 0, fmt.Errorf("ParseCelsius: value %d is outside int16 range", n)
	}
	return int16(n), nil
}

func ToFahrenheit(c int16) float64 {
	return float64(c)*9.0/5.0 + 32.0
}
