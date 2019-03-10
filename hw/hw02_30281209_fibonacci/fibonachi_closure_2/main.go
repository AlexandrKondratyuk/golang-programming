package main

import "fmt"

func main() {

	results := fibClosure(9)
	fmt.Println(results())
	fmt.Println(results())
	fmt.Println(results())
}

//func that RETURN func
func fibClosure(num int) func() []int {
	a, b := 1, 0
	fibResult := []int{}
	return func() []int {
		for i := 0; i < num; i++ {
			fibResult = append(fibResult, (func() int {
				a, b = b, a+b
				return a
			})())
		}
		return fibResult
	}
}

// package main

// import "fmt"

// func main() {
// 	results := fibClosure(9)
// 	fmt.Println(results())
// 	fmt.Println(results())
// 	fmt.Println(results())
// }

// //func that RETURN func
// func fibClosure(num int) func() []int {
// 	a, b := 1, 0
// 	fibResult := []int{}
// 	return func() []int {
// 		for i := 0; i < num; i++ {
// 			a, b = b, a+b
// 			fibResult = append(fibResult, a)
// 		}
// 		return fibResult
// 	}
// }
