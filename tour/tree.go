package tour

import (
	"fmt"
	"math/rand"
)

// Tree is a struct which represents a Node of a Binary tree
type Tree struct {
   Left  *Tree
   Value int
   Right *Tree
}

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *Tree, ch chan<- int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go walker(t1, ch1)
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			break
		}
		if v1 != v2 {
			return false
		}
	}

	return true
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func walker(t *Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
