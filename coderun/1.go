package coderun

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// первое решение
func One() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	text = strings.TrimSpace(text)

	masStr := strings.Split(text, " ")
	massive := make([]int, len(masStr))

	for i, numberStr := range masStr {
		fmt.Println(numberStr)
		numberInt, err := strconv.Atoi(numberStr)
		if err != nil {
			return
		}

		massive[i] = numberInt
	}

	sort.Slice(massive, func(i, j int) bool {
		return massive[i] < massive[j]
	})

	fmt.Println(massive[1])
}

// оптимальное решение
func OneOptimize(a, b, c int) {
	sum := a + b + c

	minVal := min(a, b, c)
	maxVal := max(a, b, c)

	fmt.Println(sum - minVal - maxVal)
}
