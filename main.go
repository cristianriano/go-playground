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

	x := "asdefghhgfedty"
	fmt.Println(x)
	fmt.Println(leetcode.LongestPalindrome(x))
}