package leetcode

import "math"

func ReverseInteger(x int) int {
	if x == 0 {
		return 0;
	}

	var res int

	for x != 0 {
		pop := x % 10
		x /= 10

		if (res > math.MaxInt32/10 || (res == math.MaxInt32/10 && pop > 7)) {
			return 0
		}
		if (res < math.MinInt32/10 || (res == math.MinInt32/10 && pop < -8)){
			return 0
		}

		res = (res * 10) + pop
	}

	return res
}
