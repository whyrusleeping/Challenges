/*
Write a function to reverse all the words in a string.
example input:
	Hello there I like to program
expected output:
	olleH ereht I ekil ot margorp
*/

package main

import (
	"fmt"
)

func swapBetween(a,b int, arr []byte) {
	if a + 1 == b {
		return
	}
	b--
	for a < b {
		 arr[a],  arr[b] =  arr[b],  arr[a]
		a++
		b--
	}
}

func revWords(s string) string {
	buf := []byte(s)
	wordbound := 0
	for i := 0; i < len(buf); i++ {
		if buf[i] == ' ' {
			swapBetween(wordbound, i, buf)
			wordbound = i + 1
		}
	}
	swapBetween(wordbound, len(buf), buf)
	return string(buf)
}

func main() {
	fmt.Println(revWords("Hello My name is jeromy johnson"))
}


