package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	a := make([]int, 0)
	b := make([]int, 0)

	for i := 0; i < n; i++ {
		var new int
		fmt.Scanf("%d", &new)
		a = append(a, new)
	}

	sort.Ints(a)

	for i := 0; i < m; i++ {
		var new int
		fmt.Scanf("%d", &new)
		b = append(b, new)
	}

	sort.Ints(b)

	intersection := find(a, b)

	fmt.Println(len(intersection))

	for i := 0; i < len(intersection); i++ {
		fmt.Printf("%d ",intersection[i])
	}

	if len(intersection) == 0 {
		fmt.Println()
	}
}

func binarySearch(set []int, l int, r int, x int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2

	if set[mid] == x {
		return mid
	}else if set[mid] < x {
		return binarySearch(set, mid+1, r, x)
	}else {
		return binarySearch(set, l, mid-1, x)
	}
}

func find(a []int, b []int) []int {
	intersection := make([]int, 0)

	var small []int
	var big []int
	if len(a) < len(b) {
		small = a
		big = b
	}else {
		small = b
		big = a
	}

	be := 0
	for i := 0; i < len(small); i++ {
		j := binarySearch(big, be, len(big), small[i])
		if j != -1 {
			intersection = append(intersection, small[i])
			be = j+1
		}
	}

	return intersection
}
