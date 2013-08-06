package main

import "fmt"

var strs = []string{"a","b","c","d"}

func permute(s []string, b []bool, cur string) {
	if len(cur) == len(s) {
		fmt.Println(cur)
		return
	}
	for i := 0; i < len(b); i++ {
		if !b[i] {
			b[i] = true
			permute(s, b, cur + s[i])
			b[i] = false
		}
	}
}

func main() {
	permute(strs, make([]bool, len(strs)), "")
}
