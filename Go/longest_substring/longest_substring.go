package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

var max = 0

func lengthOfLongestSubstring(s string) int {
	length := 0

	for i := 0; i < len(s); i++ {
		length++
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				length--
				if length > max {
					max = length
				}

				b := lengthOfLongestSubstring(s[i:])
				if b > max {
					max = b
				}

				length = 0
			}
		}
	}

	return max
}