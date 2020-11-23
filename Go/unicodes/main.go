package main

import (
	"fmt"
)

func main() {
	es := "elahe1373"
	bytesLength := len(es)
	unicodesLength := len([]rune(es))
	fmt.Println(bytesLength)
	fmt.Println(unicodesLength)
	// ----> for english words each character will be in one byte so 9 characters = 9 bytes

	fs := "الهه۱۳۷۳"
	bytesLength = len(fs)
	unicodesLength = len([]rune(fs))
	fmt.Println(bytesLength)
	fmt.Println(unicodesLength)
	// ----> for persian words each character will be in two bytes so 8 characters = 16 bytes

	// if you want to have the length of characters independent from the language use unicodes not bytes
}
