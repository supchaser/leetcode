package coderun

import (
	"fmt"
)

func Nine() {
	var N, M int
	fmt.Scan(&N, &M)

	graph := make([][]int, N+1)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// 0 - не покрашена, 1 - красный, 2 - черный
	color := make([]int, N+1)
	result := 0
	var dfs func(v int)
	dfs = func(v int) {
		for _, val := range graph[v] {
			if color[val] == 0 {
				color[val] = 3 - color[v]
				dfs(val)
				if result == 1 {
					return
				}
			} else if color[val] == color[v] {
				result = 1
				return
			}
		}
	}

	for i := 1; i <= N; i++ {
		if color[i] == 0 {
			color[i] = 1
			dfs(i)
			if result == 1 {
				break
			}
		}
	}

	if result == 1 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
