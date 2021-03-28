package leetcode

import "fmt"

func Permute(nums []int) {
	generatePermutation(nums, len(nums))
}

func generatePermutation(nums []int, size int) {
	if size == 1 {
		fmt.Println(nums)
		return
	}

	generatePermutation(nums, size - 1)
	for i := 0; i < size - 1; i++ {
		if i % 2 == 0 {
			nums[i], nums[size - 1] = nums[size - 1], nums[i]
		} else {
			nums[0], nums[size - 1] = nums[size - 1], nums[0]
		}

		generatePermutation(nums, size - 1)
	}
}
