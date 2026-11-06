package main

func Frequencies(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func MostFrequent(freq map[string]int) (string, int, bool) {
	var word string
	var count int
	var ok bool
	for candidate, n := range freq {
		if !ok {
			word = candidate
			count = n
			ok = true
		}
		if n > count {
			word = candidate
			count = n
		}
		if n == count && candidate < word {
			word = candidate
			count = n
		}
	}
	return word, count, ok
}
