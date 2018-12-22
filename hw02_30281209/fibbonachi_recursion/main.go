package main

import (
	"fmt"
	"strconv"
)

func main() {
	fibLoop(9)
}

func fib(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

func fibLoop(num int) []string {
	var newStr []string
	for i := 0; i < num; i++ {
		newStr = append(newStr, strconv.Itoa(fib(i)))
	}
	fmt.Println(newStr)
	return newStr
}
