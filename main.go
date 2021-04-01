package main

import (
	"fmt"

	"github.com/cristianriano/go-playground/leetcode"
)

func main() {
	// fmt.Println(leetcode.LengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(leetcode.ReverseInteger(1234))

	// x := []int{3, -9, -1, 3, 4, 5, -2}
	// fmt.Println(leetcode.ThreeSum(x))

	// var st leetcode.Stack
	// st.Push(1)
	// fmt.Println(st.Pop())

	// x := "asdefghhgfedty"
	// fmt.Println(x)
	// fmt.Println(leetcode.LongestPalindrome(x))

	// fmt.Println(leetcode.TwoSum([]int{2, 7, 11, 15}, 26))

	// leetcode.Permute([]int{1, 2, 3, 4})

	root := leetcode.BinaryThree{Value: 5}
	root.Insert(4)
	root.Insert(2)
	root.Insert(8)
	root.Insert(7)
	root.Insert(1)
	root.PrintInOrder()

	fmt.Println(root.Includes(3))
	fmt.Println(root.Includes(1))
}