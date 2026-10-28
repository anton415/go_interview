package main

func Accumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Readers(values []int) []func() int {
	readers := make([]func() int, len(values))
	for i, v := range values {
		readers[i] = func() int {
			return v
		}
	}
	return readers
}
