package coderun

import (
	"cmp"
	"fmt"
	"slices"
)

func Func155() {
	var n int
	fmt.Scan(&n)

	mas := make([]int, n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		mas[i] = a
	}

	slices.SortFunc(mas, func(a, b int) int {
		return cmp.Compare[int](a, b)
	})

	counter := 0
	flag := 0

	for i := range mas {
		if i == 0 {
			counter++
			continue
		}

		if mas[i] == mas[i-1] {
			if flag == 0 {
				counter--
				flag = -1
				continue
			} else if flag == -1 {
				continue
			}

		}

		flag = 0
		counter++

	}

	fmt.Print(counter)
}
