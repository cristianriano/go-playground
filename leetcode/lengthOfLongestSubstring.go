package leetcode

func LengthOfLongestSubstring(s string) int {
  if len(s) <= 1 {
    return len(s)
  }

	m, x, y := 1, 0, 0
	set := make(map[byte]int)

	for ; y < len(s); y++ {
		c := s[y]
		pos, ok := set[c]

		if ok {
			m = max(m, y - x)
			for ; x <= pos; x++ {
				delete(set, c)
			}
		}

		set[c] = y
	}

	return max(m, y - x)
}

func max(x, y int) int {
	if x > y { return x }
	return y
}
