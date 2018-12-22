package main

import (
	"fmt"
)

func main() {

	results := fibClosure(8)
	fmt.Println(results)
}

func fibClosure(num int) []int {
	a, b := 0, 1
	fibResult := []int{}
	for i := 0; i < num; i++ {
		fibResult = append(fibResult, a)
		a, b = b, a+b
	}
	return fibResult
}
