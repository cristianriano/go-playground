package leetcode

func LongestPalindrome(s string) string {
	positions := make(map[rune][]int)
	res := s[:1]

	for index, letter := range s {
		pos, ok := positions[letter]

		if ok {
			for _, i := range pos {
				if testCandidate(i, index, s) {
					if (index - i) > len(res) - 1 { res = s[i:index + 1] }
					break
				}
			}
		}

		positions[letter] = append(positions[letter], index)
	}

	return res
}

func testCandidate(start, current int, s string) bool {
	for {
		if start >= current { return true }

		if s[start] != s[current] { return false }

		start++
		current--
	}
}


// Old Solution using a Stack
// func longestPalindrome(s string) string {
//   var stack Stack
//   positions := make(map[rune][]int)
//   res := s[:1]

//   for index, letter := range s {
//     stack.Push(letter)
//     pos, ok := positions[letter]

//     if ok {
//       for _, i := range pos {
//         if testCandidate(i, index, stack, s) {
//           if (index - i) > len(res) - 1 { res = s[i:index + 1] }
//           break
//         }
//       }
//     }

//     positions[letter] = append(positions[letter], index)
//   }

//   return res
// }

// func testCandidate(start, current int, stack Stack, s string) bool {
// 	end := (start + (current - start) / 2) + 1

// 	for _, letter := range s[start:end] {
// 		val, _ := stack.Pop()
// 		if val != letter {
// 			return false
// 		}
// 	}
// 	return true
// }