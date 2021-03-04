package tour

import "strings"

// WordCount returns a map with word count
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range(strings.Fields(s)) {
		m[word]++
	}
	return m
}