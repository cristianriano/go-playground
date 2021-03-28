package leetcode

import "fmt"

func Permute(nums []int) {
	generatePermutation(make([]int, len(nums)), nums, 0)
}

func generatePermutation(head, rest []int, deep int) {
	if len(rest) == 0 {
		fmt.Println(head)
		return
	}

	tmp := make([]int, len(rest) - 1)
	for i := range rest {
		head[deep] = rest[i]
		copy(tmp[:i], rest[:i])
		if i + 1 != len(rest) {
			copy(tmp[i:], rest[i+1:])
		}
		generatePermutation(head, tmp, deep + 1)
	}
}
