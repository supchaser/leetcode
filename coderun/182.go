package coderun

import (
	"fmt"
)

func decomposition(nod, nok int) int {
	count := 0
	if nok%nod != 0 {
		return 0
	}

	nokDivideNod := nok / nod
	uniquePairs := make(map[int]bool)
	temp := nokDivideNod
	for i := 2; i*i <= temp; i++ {
		if temp%i == 0 {
			uniquePairs[i] = true
			for temp%i == 0 {
				temp /= i
			}
		}
	}

	if temp > 1 {
		uniquePairs[temp] = true
	}

	count = len(uniquePairs)

	return 1 << count
}

func Func182() {
	var nod, nok int
	fmt.Scan(&nod, &nok)

	fmt.Println(decomposition(nod, nok))
}
