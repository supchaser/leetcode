package coderun

import (
	"fmt"
	"slices"
	"strings"
)

func Ten() {
	var N, M int
	fmt.Scan(&N, &M)
	graph := make([][]int, N+1)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		graph[a] = append(graph[a], b)
	}

	state := make([]int, N+1)
	var hasCycle bool
	sortMas := make([]int, 0, N)

	var dfs func(v int)
	dfs = func(v int) {
		state[v] = 1

		for _, to := range graph[v] {
			if state[to] == 0 {
				dfs(to)
				if hasCycle {
					return
				}
			} else if state[to] == 1 {
				hasCycle = true
				return
			}
		}
		state[v] = 2
		sortMas = append(sortMas, v)
	}

	for i := 1; i <= N; i++ {
		if state[i] == 0 {
			dfs(i)
			if hasCycle {
				break
			}
		}
	}

	if hasCycle {
		return
	}

	slices.Reverse(sortMas)
	mas1 := strings.Trim(fmt.Sprint(sortMas), "]")
	fmt.Print(strings.Trim(mas1, "["))

}
