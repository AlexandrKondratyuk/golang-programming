package main

import (
	"fmt"
	"homework/hw01_20181205/simple/convert"
	"reflect"
)

func main() {
	fmt.Println("Start Go converter")

	resStrToInt := convert.StringToInt("0")
	fmt.Printf("Type of value = %T, value = %d\n", resStrToInt, resStrToInt)

	resIntToString := convert.IntToString(1)
	fmt.Printf("Type of value: %T, value: %s\n", resIntToString, resIntToString)

	resStringToFloat := convert.StringToFloat("2")
	fmt.Printf("Type of value: %T, value: %f\n", resStringToFloat, resStringToFloat)

	resStringToBool := convert.StringToBool("true")
	fmt.Printf("Type of value: %T, value: %t\n", resStringToBool, resStringToBool)

	resStringToUint := convert.StringToUint("3")
	fmt.Printf("Type of value: %T, value: %d\n", resStringToUint, resStringToUint)

	resIntToFloat := convert.IntToFloat(4)
	fmt.Printf("Type of value: %T, value: %f\n", resIntToFloat, resIntToFloat)

	resIntToBool := convert.IntToBool(5)
	fmt.Printf("Type of value: %T, value: %t\n", resIntToBool, resIntToBool)

	resIntToUint := convert.IntToUint(-6)
	fmt.Printf("Type of value: %T, value: %d\n", resIntToUint, resIntToUint)

	resUintToInt := convert.UintToInt(7)
	fmt.Printf("Type of value: %T, value: %d\n", resUintToInt, resUintToInt)

	resFloatToInt := convert.FloatToInt(8.789)
	fmt.Printf("Type of value: %T, value: %d\n", resFloatToInt, resFloatToInt)

	resBoolToInt := convert.BoolToInt(true)
	fmt.Printf("Type of value: %T, value: %d\n", resBoolToInt, resBoolToInt)

	resFloatToBool := convert.FloatToBool(10.62764)
	fmt.Printf("Type of value: %T, value: %t\n", resFloatToBool, resFloatToBool)

	resUintToBool := convert.UintToBool(11)
	fmt.Printf("Type of value: %T, value: %t\n", resUintToBool, resUintToBool)

	resBoolToFloat := convert.BoolToFloat(true)
	fmt.Printf("Type of value: %T, value: %f\n", resBoolToFloat, resBoolToFloat)

	resFloatToString := convert.FloatToString(13.54)
	fmt.Printf("Type of value: %T, value: %s\n", resFloatToString, resFloatToString)

	resFloatToUint := convert.FloatToUint(14.54)
	fmt.Printf("Type of value: %T, value: %d\n", resFloatToUint, resFloatToUint)

	resUintToFloat := convert.UintToFloat(15)
	fmt.Printf("Type of value: %T, value: %f\n", resUintToFloat, resUintToFloat)

	resUintToString := convert.UintToString(16)
	fmt.Printf("Type of value: %T, value: %s\n", resUintToString, resUintToString)

	resBoolToUint := convert.BoolToUint(true)
	fmt.Printf("Type of value: %T, value: %d\n", resBoolToUint, resBoolToUint)

	resBoolToString := convert.BoolToString(false)
	fmt.Printf("Type of value: %T, value: %s\n", resBoolToString, resBoolToString)

	fmt.Println(String(3.4))


	fmt.Println("Finished Go converter")
}
