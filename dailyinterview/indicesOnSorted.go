package dailyinterview

func IndicesOnSorted(nums []int, target int) (int, int) {
	i := BinarySearch(nums, target)
	if i == -1 { return -1, -1 }

	x, y := i, i
	x0, x1 := 0, i-1
	y0, y1 := i+1, len(nums) - 1

	for {
		if x0 > x1 { break }
		tmp := (x0 + x1) / 2

		if nums[tmp] == target {
			x = tmp
			x1 = tmp - 1
		} else {
			x0 = tmp + 1
		}
	}

	for {
		if y0 > y1 { break }
		tmp := (y0 + y1) / 2

		if nums[tmp] == target {
			y = tmp
			y0 = tmp + 1
		} else {
			y1 = tmp - 1
		}
	}

	return x, y
}

func BinarySearch(nums []int, target int) int {
	x, y := 0, len(nums) - 1

	for {
		if x > y { break }
		index := (x + y) / 2

		if nums[index] == target {
			return index
		} else if target > nums[index] {
			x = index + 1
		} else {
			y = index - 1
		}
	}

	return -1
}