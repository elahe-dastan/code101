package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	l, _ := reader.ReadString('\n')
	characters := []rune(strings.TrimSuffix(l, "\n"))

	b := build(characters)

	if !valid(b){
		fmt.Println("Invalid")
	}else {
		fmt.Println("Valid")
		fmt.Println(apply(b))
	}


}

func build(characters []rune) []rune {
	ans := make([]rune, 0)

	i := 0
	for i < len(characters) {
		if i+7 <= len(characters){
			if string(characters[i:i+6]) == "<down>" {
				ans = append(ans, '(')
				i += 6
			}else if string(characters[i:i+7]) == "</down>" {
				ans = append(ans, ')')
				i += 7
			}else if string(characters[i:i+4]) == "<up>" {
				i += 4
				ans = append(ans, '{')
			}else if string(characters[i:i+5]) == "</up>" {
				i += 5
				ans = append(ans, '}')
			}else {
				ans = append(ans, characters[i])
				i++
			}
		}else if i+6 <= len(characters) {
			if string(characters[i:i+6]) == "<down>" {
				ans = append(ans, '(')
				i += 6
			}else if string(characters[i:i+4]) == "<up>" {
				i += 4
				ans = append(ans, '{')
			}else if string(characters[i:i+5]) == "</up>" {
				i += 5
				ans = append(ans, '}')
			}else {
				ans = append(ans, characters[i])
				i++
			}
		}else if i+5 <= len(characters) {
			if string(characters[i:i+4]) == "<up>" {
				i += 4
				ans = append(ans, '{')
			}else if string(characters[i:i+5]) == "</up>" {
				i += 5
				ans = append(ans, '}')
			}else {
				ans = append(ans, characters[i])
				i++
			}
		}else if i+4 <= len(characters) {
			if string(characters[i:i+4]) == "<up>" {
				i += 4
				ans = append(ans, '{')
			}else {
				ans = append(ans, characters[i])
				i++
			}
		}else {
			ans = append(ans, characters[i])
			i++
		}

	}

	return ans
}

func valid(characters []rune) bool {
	var f rune
	var i int
	var e rune
	for i = 0; i < len(characters); i++ {
		out := false
		switch characters[i] {
		case '(':
			f = '('
			e = ')'
			out = true
		case '{':
			f = '{'
			e = '}'
			out = true
		case ')':
			return false
		case '}':
			return false
		}

		if out {
			break
		}
	}

	if e == 0 {
		return true
	}

	index := findIndex(characters, f, e, i+1)

	if index == -1 {
		return false
	}

	return valid(characters[i+1:index]) && (valid(characters[index+1:]) || index+1 > len(characters))
}

func findIndex(characters []rune,f rune, e rune, begin int) int {
	count := 1

	for i := begin; i < len(characters); i++ {
		switch characters[i] {
		case f:
			count++
		case e:
			count--
		}

		if count == 0 {
			return i
		}
	}

	return -1
}

// 0 khodesh 1 up 2 down
func apply(characters []rune) string {
	statement := ""
	stack := make([]int, 0)

	for i := 0; i < len(characters); i++ {
		switch characters[i] {
		case '(':
			stack = append(stack, 2)
		case '{':
			stack = append(stack, 1)
		case ')':
			stack = stack[: len(stack)-1]
		case '}':
			stack = stack[: len(stack)-1]
		default:
			if len(stack) == 0 {
				statement += string(characters[i])
			}else if stack[len(stack) - 1] == 1 {
				statement += strings.ToUpper(string(characters[i]))
			}else if stack[len(stack)-1] == 2 {
				statement += strings.ToLower(string(characters[i]))
			}
		}
	}

	return statement
}
