package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}} //8
	wanted := 16

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		if j := search(row, wanted, 0); j != -1 {
			fmt.Print(i,j)
			break
		}
	}
}

func search(row []int, wanted int, c int) int {
	n := len(row)
	if n == 1 && row[0] != wanted {
		return -1
	}
	if row[n/2] == wanted {
		return c+n/2
	}
	if row[n/2] < wanted {
		return search(row[n/2:], wanted, c+n/2)
	}else{
		return search(row[0:n/2], wanted, c)
	}
}
