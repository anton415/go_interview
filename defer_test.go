package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeferredOrder(t *testing.T) {
	order := deferredOrder()
	expected := []int{3, 2, 1}
	assert.Equal(t, expected, order)
}
