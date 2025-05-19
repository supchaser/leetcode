package coderun

import (
	"bufio"
	"fmt"
	"os"
)

func Func42back1() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	m := make(map[int]int)

	maxVal := 0
	maxKey := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		m[a]++
		if m[a] > maxVal || (m[a] == maxVal && a > maxKey) {
			maxVal = m[a]
			maxKey = a
		}
	}

	fmt.Fprintln(writer, maxKey)
}
