package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

var walked map[coordinate]bool

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	matrix := make([][]int, n)
	walked = make(map[coordinate]bool, n*m)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		l,_ := reader.ReadString('\n')
		l = strings.TrimSuffix(l, "\n")
		r := strings.Split(l, " ")
		row := make([]int, 0)

		for j := 0; j < m; j++ {
			num,_ := strconv.Atoi(r[j])
			walked[coordinate{
				x: i,
				y: j,
			}] = false
			row = append(row, num)
		}

		matrix[i] = row
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if walked[coordinate {
				x: i,
				y: j,
			}] || matrix[i][j] == 0{
				continue
			}

			v := sum(matrix, i, j)
			if v > max {
				max = v
			}
		}
	}

	fmt.Print(max)
}

func sum(matrix [][]int, i int, j int) int {
	if i == len(matrix) || j == len(matrix[0]) || i < 0 || j < 0{
		return 0
	}

	if matrix[i][j] == 0 || walked[coordinate {
		x: i,
		y: j,
	}]{
		return 0
	}

	walked[coordinate {
		x: i,
		y: j,
	}] = true

	return matrix[i][j] + sum(matrix, i+1, j) + sum(matrix, i, j+1) + sum(matrix, i-1, j) + sum(matrix, i, j-1)
}