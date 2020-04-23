package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type xy struct {
	x, y int
}

func main() {
	var n, m, k int

	reader := bufio.NewReader(os.Stdin)

	l1, _ := reader.ReadString('\n')
	fmt.Sscanf(l1, "%d %d", &n, &m)

	l2, _ := reader.ReadString('\n')
	fmt.Sscanf(l2, "%d", &k)

	if k < m*m {
		fmt.Println("No")
		return
	}

	matrix := make(map[xy]bool)

	for i := 0; i < k; i++ {
		l, _ := reader.ReadString('\n')
		l = strings.TrimSuffix(l, "\n")
		var x, y int
		fmt.Sscanf(l, "%d %d", &x, &y)

		point := xy{x - 1, y - 1}

		matrix[point] = true
	}

	fmt.Println(find(matrix, m))
}

func find(matrix map[xy]bool, m int) string {
	for point := range matrix {
		if check(matrix, point.x, point.y, m) {
			return "Yes"
		}
	}

	return "No"
}

func check(matrix map[xy]bool, i int, j int, m int) bool {
	for x := i; x < i+m; x++ {
		for y := j; y < j+m; y++ {
			if !matrix[xy{x, y}] {
				return false
			}
		}
	}

	return true
}
