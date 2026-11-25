package main

import (
	"fmt"
	"testing"
)

var sliceResult []int

func BenchmarkBuild(b *testing.B) {
	for _, n := range []int{1, 1000} {
		for _, reserve := range []bool{false, true} {
			name := fmt.Sprintf("n=%d/reserve=%t", n, reserve)
			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					sliceResult = Build(n, reserve)
				}
			})
		}
	}
}

func TestBuild(t *testing.T) {
	for _, n := range []int{0, 1, 1000} {
		for _, reserve := range []bool{false, true} {
			got := Build(n, reserve)
			if len(got) != n {
				t.Fatalf("Build(%d, %t) вернула %d элементов, ожидали %d", n, reserve, len(got), n)
			}
			for i, value := range got {
				if i != value {
					t.Fatalf("Build(%d, %t)[%d] = %d, ожидали %d", n, reserve, i, value, i)
				}
			}
		}
	}
}
