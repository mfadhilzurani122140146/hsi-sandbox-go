package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)

	// this logic will works if length sting is odd, if even it wont works
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Input: ")
	fmt.Scan(&input)

	result := isPalindrome(input)

	if result {
		fmt.Printf("Output: true (karena \"%s\" dibaca sama dari depan dan belakang)\n", input)
	} else {
		fmt.Printf("Output: false (karena \"%s\" tidak sama jika dibaca terbalik)\n", input)
	}
}
