package coderun

import (
	"fmt"
	"slices"
)

func Func40back1() {
	// n - кол-во предложений, m - кол-во покупателей
	var n, m int
	fmt.Scan(&n, &m)

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}

	coasts := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&coasts[i])
	}

	slices.Sort(prices)

	slices.SortFunc(coasts, func(a, b int) int {
		return b - a
	})

	i := 0
	j := 0
	sum := 0

	for i < len(prices) && j < len(coasts) {
		if coasts[j]-prices[i] <= 0 {
			break
		}
		sum += coasts[j] - prices[i]
		i++
		j++
	}

	fmt.Println(sum)
}
