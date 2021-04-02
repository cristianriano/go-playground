package main

import (
	"github.com/cristianriano/go-playground/dailyinterview"
	"github.com/cristianriano/go-playground/linkedlist"
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

	// root := leetcode.BinaryThree{Value: 5}
	// root.Insert(4)
	// root.PrintInOrder()
	// fmt.Println(root.Includes(3))

	// fmt.Println(leetcode.Merge([]int{6, 9, 3, 5, 1, 4, 8}))

	// fmt.Println(leetcode.SolveEquation("1+x=2"))

	// fmt.Println(leetcode.InsertSort([]int{6, 9, 3, 5, 1, 4, 8}))

	// leetcode.DuplicateZeros([]int{0,1,2})
	// leetcode.MergeArrays([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	// fmt.Println(leetcode.RemoveElement([]int{3,3,3}, 3))
	// fmt.Println(leetcode.RemoveDuplicates([]int{1,2,3,4,5}))
	// fmt.Println(leetcode.CheckIfExist([]int{-2,0,10,-19,4,6,-8}))

	list1 := linkedlist.Constructor()
	list2 := linkedlist.Constructor()
	list1.AddAtTail(2)
	// list1.AddAtTail(4)
	// list1.AddAtTail(3)

	list2.AddAtTail(9)
	list2.AddAtTail(9)
	list2.AddAtTail(9)
	list2.AddAtTail(9)
	list1.Print()
	list2.Print()

	sum := dailyinterview.SumLinkedList(list1, list2)
	sum.Print()
}