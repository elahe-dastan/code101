package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bulbSwitch(9))
}

func bulbSwitch(n int) int {
	numOfOn := 0

	for i := 1; i <= n; i++ {
		if oddNumOfDivider(i) {
			numOfOn++
		}
	}

	return numOfOn
}

func oddNumOfDivider(n int) bool {
	a := int(math.Sqrt(float64(n)))

	if a * a == n {
		return true
	}

	return false
}
