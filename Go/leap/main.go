package main

import "fmt"

func main() {
	var y int

	fmt.Scanf("%d", &y)

	r := 0
	l := isLeap(y)

	for {
		if isLeap(y) {
			r = (r + 366) % 7
		} else {
			r = (r + 365) % 7
		}

		y++

		if r == 0 && l == isLeap(y) {
			break
		}
	}

	fmt.Println(y)
}

func isLeap(y int) bool {
	return (y%400 == 0) || (y%4 == 0 && y%100 != 0)
}
