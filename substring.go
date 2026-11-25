package main

func LongestUniqueSubstring(s string) int {
	best := 0
	left := 0
	lastSeen := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		ch := s[right]
		prev, ok := lastSeen[ch]

		if ok && prev >= left {
			left = prev + 1
		}

		lastSeen[ch] = right

		windowLen := right - left + 1

		if best < windowLen {
			best = windowLen
		}
	}

	return best
}
