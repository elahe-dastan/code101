package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// extracts a. from the string and then searches the rest of the string
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("abac", -1))
	fmt.Println(re.FindAllString("aab", -1))
	fmt.Println(strings.Fields("elahe	12"))
}
