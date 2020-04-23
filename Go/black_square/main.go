package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	
	matrix := make([][]int, n)
	
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		l,_ := reader.ReadString('\n')
		l = strings.TrimSuffix(l, "\n")
		c := strings.Split(l, "")
		row := make([]int, n)
		
		for j := 0; j < n; j++ {
			cell,_ := strconv.Atoi(c[j])
			row[j] = cell
		}
		
		matrix[i] = row
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
