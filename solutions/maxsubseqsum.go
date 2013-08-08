package main

import (
	"fmt"
)

var test = []int{1,5,6,-9,-21,3,4,2,9,8,-5,-14,-45,4,8,1,9}

func MaxSumOn(a []int) int {
	i, max  := 0, 0
	for _,n := range a {
		i += n
		if i > max { max = i }
		if i < 1 { i = 0 }
	}
	return max
}

func main() {
	fmt.Println(MaxSumOn(test))
}
