package main

import "strings"

func RemoveValue(items []int, value int) []int {
	write := 0
	for read := 0; read < len(items); read++ {
		if items[read] == value {
			continue
		}
		items[write] = items[read]
		write++
	}
	clear(items[write:])
	return items[:write]
}

func KeepPositive(items []int) []int {
	write := 0
	for read := 0; read < len(items); read++ {
		if items[read] <= 0 {
			continue
		}
		items[write] = items[read]
		write++
	}
	clear(items[write:])
	return items[:write]
}

// дана строка нужно ее сжать по правилу, если идет подряд повторяющиеся символы, то их заменить числом.
// например: "aaafbbbb" -> "a3fb4"
func ColapseString(s string) string {
	if len(s) == 0 {
		return ""
	}
	var result strings.Builder
	result.WriteByte(s[0])
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count > 1 {
				result.WriteString(string(byte('0' + count)))
			}
			result.WriteString(string(s[i]))
			count = 1
		}
	}
	if count > 1 {
		result.WriteString(string(byte('0' + count)))
	}
	return result.String()
}
