package leetcode

import "strings"

func countPrefixSuffixPairs(words []string) int {
	result := 0

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
				result++
			}
		}
	}

	return result
}
