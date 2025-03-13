package coderun

import "fmt"

func Twelve() {
	var N, start, end int
	fmt.Scan(&N)

	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Scan(&start, &end)

	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = -1
	}

	dist[start] = 0
	queue := []int{start}

	for len(queue) > 0 {
		curVertex := queue[0]
		queue = queue[1:]

		for u := 1; u <= N; u++ {
			if matrix[curVertex-1][u-1] == 1 && dist[u] == -1 {
				dist[u] = dist[curVertex] + 1
				queue = append(queue, u)
				if u == end {
					break
				}
			}
		}
	}

	fmt.Print(dist[end])
}
