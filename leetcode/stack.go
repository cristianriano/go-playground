package leetcode

type Stack []int

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 { return true }
	return false
}

func (s *Stack) Size() int {
	return len(*s) - 1
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	x := (*s)[s.Size()]
	*s = (*s)[:s.Size()]

	return x, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return (*s)[s.Size()], true
}