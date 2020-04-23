package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	var n, m, k int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Scanf("%d", &k)

	if k < m * m {
		fmt.Println("No")
		return
	}

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		row := make([]int, n)
		matrix[i] = row
	}

	reader := bufio.NewReader(os.Stdin)


	for i := 0; i < k; i++ {
		l,_ := reader.ReadString('\n')
		l = strings.TrimSuffix(l, "\n")
		var x, y int
		fmt.Sscanf(l,"%d %d", &x, &y)

		matrix[x-1][y-1] = 1
	}

	fmt.Println(find(matrix, m))
}

func find(matrix [][]int, m int) string {
	for i := 0; i <= len(matrix)-m; i++ {
		for j := 0; j <= len(matrix)-m; j++ {
			if matrix[i][j] == 0 {
				continue
			}

			if check(matrix, i, j, m) {
				return "Yes"
			}
		}
	}

	return "No"
}

func check(matrix [][]int, i int, j int, m int) bool {
	for x := i; x < i + m; x++ {
		for y := j; y < j + m; y++ {
			if matrix[x][y] == 0 {
				return false
			}
		}
	}

	return true
}
