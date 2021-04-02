package leetcode

func QuickSort(nums []int) {
	quick(nums, 0, len(nums) - 1)
}

func quick(nums []int, first, last int) {
	i, j, pivot := first, last, (first + last)/2

	for ; i < j; {
		for ; nums[i] < nums[pivot]; i++ {}
		for ; nums[j] > nums[pivot]; j-- {}

		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	if first < j { quick(nums, first, j) }
	if last > i { quick(nums, i, last) }
}

func Bubble(nums []int) {
	swaps := 1

	for swaps != 0 {
		swaps = 0

		for i := 0; i < len(nums) - 1; i++ {
			if nums[i] > nums[i + 1] {
				nums[i], nums[i + 1] = nums[i + 1], nums[i]
				swaps++
			}
		}
	}
}

func BubbleImproved(nums []int) {
	swaps, runs := 1, 0

	for swaps != 0 {
		swaps = 0

		for i := 0; i < len(nums) - 1 - runs; i++ {
			if nums[i] > nums[i + 1] {
				nums[i], nums[i + 1] = nums[i + 1], nums[i]
				swaps++
			}
		}
		runs++
	}
}

func Merge(nums []int) []int {
	return directMerge(nums)
}

func directMerge(nums []int) []int {
	if len(nums) > 1 {
		center := len(nums) / 2
		left, right := make([]int, center), make([]int, len(nums) - center)
		copy(left[:], nums[0:center])
		copy(right[:], nums[center:])

		left = directMerge(left)
		right = directMerge(right)

		nums = intercalate(left, right)
	}
	return nums
}

func intercalate(x, y []int) []int {
	res := make([]int, len(x) + len(y))
	i, j, k := 0, 0, 0

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			res[k] = x[i]
			i++
		} else {
			res[k] = y[j]
			j++
		}
		k++
	}

	// Add remaining elements
	for ; i < len(x); i++ {
		res[k] = x[i]
	}
	for ; j < len(y); j++ {
		res[k] = y[j]
	}

	return res
}

func InsertSort(nums []int) []int {
	res := make([]int, len(nums))

	for i, x := range nums {
		insertSorted(res, x, i)
	}
	return res
}

func insertSorted(list []int, x, index int) {
	j := 0

  for j = index - 1; j >= 0 && list[j] > x; j-- {
		list[j + 1] = list[j]
	}
	list[j + 1] = x
}
