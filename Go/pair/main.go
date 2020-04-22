package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	reader := bufio.NewReader(os.Stdin)
	line,_ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	numbers := strings.Split(line, " ")

	series := make([]int, 0)

	for i := 0; i < len(numbers); i++ {
		s,_ := strconv.Atoi(numbers[i])
		if s < k {
			series = append(series, s)
		}
	}

	sort.Ints(series)

	ans := sum(series, k)

	if ans == nil {
		fmt.Print("No")
	}else {
		fmt.Println("Yes")
		fmt.Printf("%d %d", ans[0], ans[1])
	}

}

func sum(series []int, k int) []int {
	if len(series) == 0 {
		return nil
	}

	s := series[0] + series[len(series) - 1]

	if s == k {
		return []int{series[0], series[len(series) - 1]}
	}else if s < k {
		return sum(series[1:], k)
	}else {
		return sum(series[:len(series)-1], k)
	}
}
