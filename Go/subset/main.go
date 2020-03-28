package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	powSetSize := 1 << n
	fmt.Println(powSetSize)

	for counter := 0; counter < powSetSize; counter++ {
		for j := 0; j < n; j++ {
			k := 1 << j
			if counter&k > 0 {
				fmt.Print(j + 1)
			}
		}
		fmt.Println()
	}
}
