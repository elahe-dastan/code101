package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
		log.Fatal(err)
	}

	a := make([]int, 0)
	b := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)

	aMembers, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	aMembers = strings.TrimSuffix(aMembers, "\n")

	bMembers, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	bMembers = strings.TrimSuffix(bMembers, "\n")

	aNumbers := strings.Split(aMembers, " ")
	bNumbers := strings.Split(bMembers, " ")

	for i := 0; i < n; i++ {
		ne, _ := strconv.Atoi(aNumbers[i])
		a = append(a, ne)
	}

	sort.Ints(a)

	for i := 0; i < m; i++ {
		ne, _ := strconv.Atoi(bNumbers[i])
		b = append(b, ne)
	}

	sort.Ints(b)

	intersection := find(a, b)

	fmt.Println(len(intersection))

	for i := 0; i < len(intersection); i++ {
		fmt.Printf("%d ", intersection[i])
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
	} else if set[mid] < x {
		return binarySearch(set, mid+1, r, x)
	} else {
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
	} else {
		small = b
		big = a
	}

	be := 0
	for i := 0; i < len(small); i++ {
		j := binarySearch(big, be, len(big)-1, small[i])
		if j != -1 {
			intersection = append(intersection, small[i])
			be = j + 1
		}
	}

	return intersection
}
