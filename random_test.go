package main

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueRandom(t *testing.T) {
	cases := []struct {
		name    string
		n, max  int
		wantErr bool
	}{
		{"empty", 0, 0, false},
		{"regular", 3, 10, false},
		{"full", 5, 5, false},
		{"negative n", -1, 5, true},
		{"negative max", 0, -1, true},
		{"too many", 6, 5, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := UniqueRandom(tc.n, tc.max, rand.New(rand.NewPCG(42, 42)))
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Len(t, got, tc.n)
			seen := make(map[int]struct{}, len(got))
			for _, value := range got {
				require.GreaterOrEqual(t, value, 0)
				require.Less(t, value, tc.max)
				require.NotContains(t, seen, value, "duplicate: %d", value)
				seen[value] = struct{}{}
			}
			again, err := UniqueRandom(tc.n, tc.max, rand.New(rand.NewPCG(42, 42)))
			require.NoError(t, err)
			require.Equal(t, got, again, "sequence must be reproducible")
		})
	}
}

func TestUniqueRandomNilSource(t *testing.T) {
	for _, n := range []int{0, 1} {
		_, err := UniqueRandom(n, 10, nil)
		require.Error(t, err, "n=%d: expected error for nil source", n)
	}
}
