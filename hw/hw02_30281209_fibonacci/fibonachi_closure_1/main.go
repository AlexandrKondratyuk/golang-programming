// package main

// import "fmt"

// func main() {

// 	results := fibClosure(9)
// 	fmt.Println(results)
// }

// // func that USE another func
// func fibClosure(num int) []int {
// 	a, b := 1, 0
// 	fibResult := []int{}
// 	for i := 0; i < num; i++ {
// 		fibResult = append(fibResult, (func() int {
// 			a, b = b, a+b
// 			return a
// 		})())
// 	}
// 	return fibResult
// }

package main

import (
	"fmt"
	"strconv"
)

func main() {

	results := fibClosure(9)
	fmt.Println(results())
	fmt.Println(results())
}

// func that USE another func
func fibClosure(num int) func() string {
	a, b := 1, 0
	fibResult := ""

	return func() string {
		for i := 0; i < num; i++ {
			a, b = b, a+b
			fibResult += strconv.Itoa(a) + " "
		}
		return fibResult
	}
}
