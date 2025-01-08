package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversed := 0
	temp := x

	for temp != 0 {
		digit := temp % 10
		reversed = reversed*10 + digit
		temp /= 10
	}

	return reversed == x
}
