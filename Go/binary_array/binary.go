package main

import "fmt"

func main() {
	binary := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1}

	max := 0

	ones := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == 1 {
			ones++
		}else {
			if ones > max {
				max = ones
				ones = 0
			}
		}
	}

	if ones > max {
		max = ones
	}

	fmt.Println(max)
}
