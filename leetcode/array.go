package leetcode

import "fmt"

func DuplicateZeros(arr []int) {
	x := 0

	fmt.Println(arr)
	for ; x < len(arr) - 1; {
		if arr[x] == 0 {
			shiftRight(arr[x:])
			x++
		}
		x++
	}
	fmt.Println(arr)
}

func shiftRight(arr []int) {
	for i := len(arr) - 2; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
}

func MergeArrays(nums1 []int, m int, nums2 []int, n int)  {
	tmp := make([]int, m)
	copy(tmp, nums1[:m])

	i, j, k := 0, 0, 0
	for ; i < len(tmp) && j < len(nums2); {
		if tmp[i] < nums2[j] {
			nums1[k] = tmp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for ; i < len(tmp); i++ {
		nums1[k] = tmp[i]
		k++
	}
	for ; j < len(nums2); j++ {
		nums1[k] = nums2[j]
		k++
	}
}

func RemoveElement(nums []int, val int) int {
	fmt.Println(nums)
	i, j := 0, len(nums) - 1

	for {
		for ; i < len(nums) && nums[i] != val; i++ {}
		for ; j > 0 && nums[j] == val; j-- {}

		if i >= j { break }

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return i
}

func RemoveDuplicates(nums []int) int {
	i, swaps := len(nums) - 2, 0

	for ; i >= 0 ; i-- {
		if nums[i] != nums[i+1] { continue }

		j := i + 2
		for ; j < len(nums); j++ {
			nums[j-1] = nums[j]
		}
		swaps++
	}
	return len(nums) - swaps
}

func CheckIfExist(arr []int) bool {
	fmt.Println(arr)
	dict := make(map[int]int)

	for i, v := range arr {
		_, ok := dict[v*2]
		if ok {
			fmt.Printf("Double %v", v)
			return true
		}

		if v % 2 == 0 {
			_, ok := dict[v/2]
			if ok {
				fmt.Printf("Half %v\n", v)
				return true
			}
		}

		dict[v] = i
	}
	return false
}