package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Print(first("Code"))
	//fmt.Print(second(23946456))
	//fmt.Print(palindrome(345543))
	fmt.Print(fourth([][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 7}}))
}

func first(input string) string {
	output := ""

	for i := 0; i <= len(input); i++ {
		output += input[0:i]
	}

	return output
}

func second(input int) int {
	sum := 0

	for {
		sum += input % 10
		input /= 10

		if input == 0 {
			break
		}
	}

	if sum < 10 {
		return sum
	}

	return second(sum)
}

func palindrome(input int) bool {
	end := numberOfDigits(input) - 1

	return check(input, end)
}

func numberOfDigits(input int) float64 {
	n := 0.0

	for {
		n++
		input /= 10
		if input == 0 {
			break
		}
	}

	return n
}

func check(number int, end float64) bool {
	if number < 10 {
		return true
	}

	f := number % 10
	e := number / int(math.Pow(10, end))

	if f != e {
		return false
	}

	number %= int(math.Pow(10, end))
	number /= 10

	return check(number, end-2)
}

func fourth(matrix [][]int) [][]int {
	n := len(matrix)
	output := make([][]int, n)

	for i := 0; i < n; i++ {
		output[i] = make([]int, n)
	}

	for i, row := range matrix {
		for j, r := range output {
			r[i] = row[n - j - 1]
		}
	}

	return output
}