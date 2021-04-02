package dailyinterview

import "github.com/cristianriano/go-playground/linkedlist"

func SumLinkedList(list1, list2 linkedlist.MyLinkedList) linkedlist.MyLinkedList {
	sum := linkedlist.Constructor()
	carry, i := 0, 0

	for {
		x, y := list1.Get(i), list2.Get(i)
		if x == -1 && y == -1 {
			if carry > 0 { sum.AddAtTail(carry) }
			break
		}

		if x == -1 { x = 0 }
		if y == -1 { y = 0 }

		val := x + y + carry
		carry = val/10
		sum.AddAtTail(val % 10)
		i++
	}

	return sum
}