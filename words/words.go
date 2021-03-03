package words

import "strings"

// Count returns a map with word count
func Count(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range(strings.Fields(s)) {
		m[word]++
	}
	return m
}