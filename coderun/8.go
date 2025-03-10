package coderun

import (
	"fmt"
	"strings"
)

func Eight() {
	var N, M int
	fmt.Scan(&N, &M)

	graph := make([][]int, N+1)
	for range M {
		var a, b int
		fmt.Scan(&a, &b)

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	viseted := make([]bool, N+1)
	components := make([][]int, N)
	counter := 0
	var dfs func(v int)
	dfs = func(v int) {
		viseted[v] = true
		components[counter] = append(components[counter], v)
		for _, val := range graph[v] {
			if !viseted[val] {
				dfs(val)
			}
		}
	}

	for i := 1; i <= N; i++ {
		if viseted[i] {
			continue
		}
		dfs(i)
		counter++
	}

	filtered := make([][]int, 0, len(components))
	for _, slice := range components {
		if len(slice) > 0 {
			filtered = append(filtered, slice)
		}
	}

	fmt.Println(len(filtered))
	for i, slice := range filtered {
		fmt.Println(len(slice))
		slice1 := strings.Trim(fmt.Sprint(slice), "]")
		if i == len(filtered)-1 {
			fmt.Print(strings.Trim(slice1, "["))
		} else {
			fmt.Println(strings.Trim(slice1, "["))
		}
	}
}
