package coderun

import "fmt"

func Func187() {
	var T int // T - число тестов, n - количество лежаков
	fmt.Scan(&T)

	result := make([]int, T)
	for i := range T {

		var n int
		fmt.Scan(&n)

		sunbedRatings := make([]int, n+1) // чтобы индекс соответствовал позиции лежака

		for i := 1; i <= n; i++ {
			var a int
			fmt.Scan(&a)
			sunbedRatings[i] = a
		}

		temp := sunbedRatings[1] ^ sunbedRatings[2]
		for i := 1; i < len(sunbedRatings)-1; i++ {
			for j := i + 1; j < len(sunbedRatings); j++ {
				if (sunbedRatings[i] ^ sunbedRatings[j]) < temp {
					temp = sunbedRatings[i] ^ sunbedRatings[j]
				}
			}
		}

		result[i] = temp
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
