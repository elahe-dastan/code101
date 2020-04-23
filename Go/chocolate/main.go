package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	
	reader := bufio.NewReader(os.Stdin)
	l,_ := reader.ReadString('\n')
	l = strings.TrimSuffix(l, "\n")
	chocolates := strings.Split(l, " ")
	
	c := make([]int, n)
	sum := 0 
	for i := 0; i < n; i++ {
		newNumber,_ := strconv.Atoi(chocolates[i])
		c[i] = newNumber
		sum += newNumber
	}
	
	if sum % n != 0 {
		fmt.Println(-1)
		return
	}

	fmt.Println(adjust(c, sum/n, 0))
	
}

func adjust(chocolates []int, number int, level int) int {
	if len(chocolates) < 2 {
		return level
	}

	f := chocolates[0] - number
	chocolates[1] += f

	s := chocolates[len(chocolates) - 1] - number
	chocolates[len(chocolates) - 2] += s

	return adjust(chocolates[1:len(chocolates)-1], number, level + int(math.Abs(float64(f))) + int(math.Abs(float64(s))))
}
