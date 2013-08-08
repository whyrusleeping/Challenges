package main

/*
Write a function to check if a string is a palindrome
*/

import (
	"fmt"
)

func IsNumPalin(n int) bool {
	var s string
	for n > 0 {
		s += string([]byte{byte('0' + (n % 10))})
		n /= 10
	}
	return IsPalin(s)
}

func IsPalin(s string) bool {
	for i,j := 0, len(s)-1; i < j; i,j = i+1,j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalin("HellolleH"))
	fmt.Println(IsPalin("Not A Palindrome"))
	fmt.Println(IsNumPalin(10101))
	fmt.Println(IsNumPalin(123456))
}
