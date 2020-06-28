package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	return recursive(n, 1)
}

func recursive(n int, i int) bool {
	sumOfSquares := 0

	for {
		a := n % 10
		sumOfSquares += a * a
		n /= 10

		if n == 0 {
			break
		}
	}

	if sumOfSquares == 1 {
		return true
	}

	if i == 100 {
		return false
	}

	i++
	return recursive(sumOfSquares, i)
}