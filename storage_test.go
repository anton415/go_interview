package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCache map[int64]any

func (c testCache) Put(id int64, value any) {
	c[id] = value
}

func (c testCache) Get(id int64) (any, bool) {
	value, ok := c[id]
	return value, ok
}

func TestStorageGet(t *testing.T) {
	warehouse := &Warehouse{ID: 7}
	tests := []struct {
		name     string
		id       int64
		expected *Warehouse
	}{
		{"stored warehouse", warehouse.ID, warehouse},
		{"unknown key", 8, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := Storage{cache: testCache{}}
			storage.Set(warehouse)

			result := storage.Get(tt.id)
			assert.Same(t, tt.expected, result)
		})
	}
}
