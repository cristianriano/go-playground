package leetcode

import "fmt"

type BinaryThree struct {
	Value int
	left *BinaryThree
	right *BinaryThree
}

func (node *BinaryThree) Insert(x int) {
	if x < node.Value {
		if node.left == nil {
			node.left = &BinaryThree{Value: x}
		} else {
			node.left.Insert(x)
		}
	} else {
		if node.right == nil {
			node.right = &BinaryThree{Value: x}
		} else {
			node.right.Insert(x)
		}
	}
}

func (node *BinaryThree) Includes(x int) bool {
	if node == nil {
		return false
	}

	if node.Value == x {
		return true
	} else if x < node.Value {
		return node.left.Includes(x)
	} else {
		return node.right.Includes(x)
	}
}

func (node *BinaryThree) PrintInOrder() {
	if node == nil {
		return
	}

	node.left.PrintInOrder()
	fmt.Printf(" (%d) ", node.Value)
	node.right.PrintInOrder()
}
