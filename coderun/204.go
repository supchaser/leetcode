package coderun

import (
	"fmt"
)

func Func204() {
	var n int
	fmt.Scan(&n)

	m := make(map[int][2]int)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		m[i] = [2]int{a, b}
	}

	mRes := make(map[int]float64)
	sum := 0.0
	for key, v := range m {
		mas := v
		a := float64(mas[0])
		b := float64(mas[1])
		mRes[key] = a * b
		sum += a * b
	}

	for i := 1; i <= len(mRes); i++ {
		if sum == 0 {
			fmt.Println(0.0)
		} else {
			fmt.Println(mRes[i] / sum)
		}
	}
}
