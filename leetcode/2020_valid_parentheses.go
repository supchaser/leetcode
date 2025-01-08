package leetcode

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[val] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
