package main

import "unicode/utf8"

func IsPalindrome(s string) bool {
	if !utf8.ValidString(s) {
		return false
	}

	runes := []rune(s)

	left := 0
	right := len(runes) - 1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}
