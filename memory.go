package main

func Build(n int, reserve bool) []int {
	var result []int
	if reserve {
		result = make([]int, 0, n)
	}
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}
