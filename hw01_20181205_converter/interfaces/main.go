package main

import (
	"fmt"
	. "homework/hw01_20181205_converter/interfaces/casting"
)

func main() {
	// test BOOL
	//b := true
	//b := -2
	//b := "true"
	//fmt.Println(ToBool(b))

	// test STRING
	//str := 2
	//str := 2.2
	//str := uint8(2)
	//str := true
	//fmt.Println(ToString(str))

	//f := "2.3"
	f := true
	//fmt.Println(ToFloat64(f))
	fmt.Println(ToFloat32(f))
}
