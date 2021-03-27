package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) (res [][]int) {
	nums = threeUnique(nums)
	l := len(nums)

	if len(nums) < 2 {
		return res
	}

	sums := make(map[[3]int]bool)
	for i := 0; i < l - 2; i++ {
		for j := i + 1; j < l - 1; j++ {
			for k := j + 1; k < l; k++ {
				if (nums[i] + nums[j] + nums[k] == 0) {
					val := [3]int{nums[i], nums[j], nums[k]}

					sort.Ints(val[:])
					_, ok := sums[val]
					if !ok {
						sums[val] = true
						res = append(res, val[:])
					}
				}
			}
		}
	}

	return res;
}

func threeUnique(x []int) (uni []int) {
	count := make(map[int]int)

	for _, i := range x {
		if count[i] < 3 {
			count[i]+= 1
			uni = append(uni, i)
		}
	}
	return uni
}
