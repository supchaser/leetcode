package main

import (
	"fmt"
)

func main() {
	// str := "pwwkew"

	// for i, val := range str {
	// 	if i != 0 {
	// 		fmt.Printf("i: %v, val: %v, type: %T", str[i-1], val, val)
	// 	}
	// 	fmt.Printf("i: %v, val: %v, type: %T", str[i], val, val)
	// }

	str1 := "({})"

	for _, val := range str1 {
		fmt.Println(string(val))
	}

	// result := leetcode.LengthOfLongestSubstring(str)
	// fmt.Println(result)
}
