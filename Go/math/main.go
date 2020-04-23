package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	l,_ := reader.ReadString('\n')
	l = strings.TrimSuffix(l, "\n")
	characters := strings.Split(l, "")

	calculate(characters)
}

func calculate(characters []string) string {
	r := ""
	i := 0
	for ;i < len(characters); {
		if characters[i] == "(" {
			end := findEnd(characters)
			r += calculate(characters[i:end])
			i = end
		}else {
			r += characters[i]
			i++
		}
	}

	return strconv.Itoa(findOperation(r))
}

func findEnd(characters []string) int {
	count := 1

	var i int
	for i = 1; i < len(characters); i++ {
		if characters[i] == "("{
			count++
		}else if characters[i] == ")" {
			count--
			if count == 0 {
				break
			}
		}
	}

	return i
}

func findOperation(statement string) int {
	chars := strings.Split(statement, "")
	f,_ := strconv.Atoi(chars[0])
	s,_ := strconv.Atoi(chars[2])

	switch chars[1] {
	case "+":
		return Add(f, s)
	case "*":
		return Multiply(f, s)
	}
}

func Add(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}
