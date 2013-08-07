package main
import "fmt"

var a = []int{1,2,3,4,5,6,7,8}
var b = []int{1,2,3,4,6,7,8}

func main() {
	n := 0
	for _,i := range a {
		n = n ^ i
	}
	for _,i := range b {
		n = n ^ i
	}
	fmt.Println(n)
}
