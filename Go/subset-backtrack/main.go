package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	subset(n)
}

func subset(s int) {
	subsetRec("", s, 0)
}

func subsetRec(prefix string, s int, index int) {
	if s+1 == index {
		return
	}

	fmt.Println(prefix)

	for i := index + 1; i <= s; i++ {
		prefix += strconv.Itoa(i)

		subsetRec(prefix, s, i)

		prefix = prefix[:len(prefix)-1]
	}
}
