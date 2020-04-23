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
	characters := strings.Split(l, "")

	fmt.Println(calculate(characters))
}

func calculate(characters []string) string {
	r := ""
	i := 0
	for ;i < len(characters); {
		if characters[i] == "(" {
			end := findEnd(characters, i+1)
			r += calculate(characters[i+1:end])
			i = end + 1
		}else {
			r += characters[i]
			i++
		}
	}

	return strconv.Itoa(findOperation(r))
}

func findEnd(characters []string, begin int) int {
	count := 1

	i := begin
	for ; i < len(characters); i++ {
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

	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case "+":
			f,_ := strconv.Atoi(strings.Join(chars[0:i], ""))
			s,_ := strconv.Atoi(strings.Join(chars[i+1:], ""))
			return Add(f, s)
		case "*":
			f,_ := strconv.Atoi(strings.Join(chars[0:i], ""))
			s,_ := strconv.Atoi(strings.Join(chars[i+1:], ""))
			return Multiply(f, s)
		}
	}

	ans,_ := strconv.Atoi(statement)

	return ans
}

func Add(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}
