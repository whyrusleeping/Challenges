package main
import "fmt"

func findDup(a []int) int {
	i := a[0]
	for {
		if a[i] == i {
			return i
		} else {
			i, a[i] = a[i], i
		}
	}
}

var arr = []int{6,2,8,3,1,6,4,5,7}

func main() {
	fmt.Println(findDup(arr))
}
