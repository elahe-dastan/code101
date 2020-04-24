package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	l,_ := reader.ReadString('\n')
	l = strings.TrimSuffix(l, "\n")

	var n, m, d int
	fmt.Sscanf(l, "%d %d %d", &n, &m, &d)

	l,_ = reader.ReadString('\n')
	l = strings.TrimSuffix(l, "\n")
	h := strings.Split(l, " ")

	boxes := make([][]int, n-m+1)
	tallests := make([]int, n-m+1)

	highest := 0
	//shortestTallest := 100000000
	//var shortestTallestIndex int

	for i := 0; i < n-m+1; i++ {
		row := make([]int, m)
		tallest := 0
		for j := i; j < i+m; j++ {
			height,_ := strconv.Atoi(h[j])
			if height > highest {
				highest = height
			}
			if height > tallest {
				tallest = height
			}
			row[j - i] = height
		}

		tallests[i] = tallest
		//if tallest < shortestTallest {
		//	shortestTallest = tallest
		//	shortestTallestIndex = i
		//}

		boxes[i] = row
	}

	fmt.Printf("%d %d", highest+d, build(boxes, tallests, d, highest))
}

func build(boxes [][]int, tallests []int, d int, highest int) int {
	passedDays := 0
	for {
		for k := 0; k < len(tallests); k++{
			if tallests[k] < highest {
				var diff int
				if highest - tallests[k] < d - passedDays {
					diff = highest - tallests[k]
				}else {
					diff = d - passedDays
				}

				row := boxes[k]
				for i := 0; i < len(row); i++ {
					row[i] += diff
				}
				tallests[k]+=diff

				for ghabli := k - len(boxes[0]) + 1; ghabli < k; ghabli++ {
					if ghabli < 0 {
						continue
					}
					row = boxes[ghabli]
					tallest := tallests[ghabli]
					for i := k - ghabli; i < len(row); i++ {
						row[i] += diff
						if row[i] > tallest {
							tallest = row[i]
						}
					}
					tallests[ghabli] = tallest

				}

				for badi := k + 1; badi < k + len(boxes[0]); badi++ {
					if badi == len(boxes) {
						break
					}
					row = boxes[badi]
					tallest := tallests[badi]
					for i := 0; i < len(row) - (badi - k); i++ {
						row[i] += diff
						if row[i] > tallest {
							tallest = row[i]
						}
					}
					tallests[badi] = tallest

				}

				passedDays += diff
			}

			if passedDays == d {
				break
			}
		}

		if passedDays == d {
			break
		}
		highest++
	}
	tallest := 0
	for i := 0; i < len(tallests); i++ {
		if tallests[i] > tallest {
			tallest = tallests[i]
		}
	}

	return tallest
}

func findShortest(boxes [][]int, tallests []int, shortestTallestIndex int, d int) int {
	if d == 0 {
		tallest := 0
		for i := 0; i < len(tallests); i++ {
			if tallests[i] > tallest {
				tallest = tallests[i]
			}
		}

		return tallest
	}

	row := boxes[shortestTallestIndex]
	for i := 0; i < len(row); i++ {
		row[i]++
	}
	tallests[shortestTallestIndex]++

	if shortestTallestIndex > 0 {
		row = boxes[shortestTallestIndex - 1]
		tallest := tallests[shortestTallestIndex - 1]
		for i := 1; i < len(row); i++ {
			row[i]++
			if row[i] > tallest {
				tallest = row[i]
			}
		}
		tallests[shortestTallestIndex - 1] = tallest

	}

	if shortestTallestIndex < len(tallests) - 1 {
		row = boxes[shortestTallestIndex + 1]
		tallest := tallests[shortestTallestIndex + 1]
		for i := 0; i < len(row) - 1; i++ {
			row[i]++
			if row[i] > tallest {
				tallest = row[i]
			}
		}
		tallests[shortestTallestIndex + 1] = tallest

	}

	var newIndex int
	shortest := 100000000
	for i := 0; i < len(tallests); i++ {
		if tallests[i] < shortest {
			shortest = tallests[i]
			newIndex = i
		}
		//else if tallests[i] == shortest {
		//	if i == 0 || i == len(tallests) - 1 {
		//		newIndex = i
		//	}
		//}
	}

	return findShortest(boxes, tallests, newIndex, d - 1)

}