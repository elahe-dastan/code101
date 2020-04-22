package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	if Valid([]rune(s)) {
		fmt.Print("Valid")
	}else {
		fmt.Print("Invalid")
	}
}

func Valid(characters []rune) bool {
	if len(characters) == 0 {
		return true
	}

	if len(characters) % 2 == 1{
		return false
	}

	i := findEnd(characters)
	if i == -1 {
		return false
	}

	if Valid(characters[1:i]) && ((Valid(characters[i+1:])) || i+1 == len(characters)){
		return true
	}

	return false
}

func findEnd(characters []rune) int {
	count := 1
	f := characters[0]

	var e rune
	if f == '(' {
		e = ')'
	}else if f == '{' {
		e = '}'
	}else if f == '[' {
		e = ']'
	}

	for i := 1; i < len(characters); i++ {
		if f == characters[i] {
			count++
		}else if e == characters[i] {
			count--
		}

		if count == 0 {
			return i
		}
	}

	return -1
}