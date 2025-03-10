package coderun

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Seven() {
	reader := bufio.NewReader(os.Stdin)

	// читаем N и M
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	N, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	M, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	graph := make([][]int, N+1)

	// читаем ребра
	for i := 0; i < M; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, N+1)
	var component []int

	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		component = append(component, v)

		for _, val := range graph[v] {
			if !visited[val] {
				dfs(val)
			}
		}
	}

	dfs(1)

	sort.Ints(component)

	fmt.Println(len(component))
	lin1 := strings.Trim(fmt.Sprint(component), "]")
	lin2 := strings.Trim(lin1, "[")
	fmt.Print(lin2)
}
