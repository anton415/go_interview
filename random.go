package main

import (
	"fmt"
	"math/rand/v2"
)

func UniqueRandom(n int, max int, rnd *rand.Rand) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("количество не может быть отрицательным")
	}
	if max < 0 {
		return nil, fmt.Errorf("недопустимая граница")
	}
	if n > max {
		return nil, fmt.Errorf("уникальных чисел недостаточно")
	}
	if rnd == nil {
		return nil, fmt.Errorf("генератор не передан")
	}

	seen := make(map[int]struct{})
	result := make([]int, 0, n)

	for len(result) < n {
		value := rnd.IntN(max)

		if _, exists := seen[value]; exists {
			continue
		} else {
			seen[value] = struct{}{}
			result = append(result, value)
		}

	}
	return result, nil
}
