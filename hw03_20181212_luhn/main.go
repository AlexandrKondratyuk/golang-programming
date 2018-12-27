package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	numberForValidate := "1234 5678 9012 3456 7894"
	// numberForValidate := "1234 5678 9012 3456 78dssdsdsds94"
	// numberForValidate := "1234 5678 9012 3456 7890"

	numerWithoutSpaces := ""
	numberModifyed := ""
	valid := false

	// stripe spaces
	firstStripSpaces(&numberForValidate, &numerWithoutSpaces)

	// check is value length more then 1 and is value digit
	valid = secondEnoughBigAndIsNumber(numerWithoutSpaces)

	if valid {
		// modifying here string using doubling number
		thirdModifyByDoubbling(&numerWithoutSpaces, &numberModifyed)

		// check here -> divide sum of element by 10
		valid = fourthPastValidateByDiv10(&numberModifyed)
	}

	if valid {
		fmt.Printf("Value %s is VALID !!! \n", numberForValidate)
	} else {
		fmt.Printf("Value %s is INVALID, try another value !!! \n", numberForValidate)
	}

}

// 1) func delete spaces
func firstStripSpaces(num *string, newNum *string) {
	for _, val := range *num {
		if !unicode.IsSpace(val) {
			*newNum += string(val)
		}
	}
}

// 2) func check => is more then 1
// 3) func also check => is digit
func secondEnoughBigAndIsNumber(str string) bool {
	if len(str) <= 1 {
		fmt.Println("ERROR here - value is less or equal 1")
		return false
	}

	for _, c := range str {
		if !unicode.IsDigit(c) {
			fmt.Println("ERROR here - value is not number")
			return false
		}
	}

	return true
}

func thirdModifyByDoubbling(numerWithoutSpaces *string, numberModifyed *string) {
	modifyedNumber := strings.Builder{}
	newStr := ""

	for i, v := range *numerWithoutSpaces {

		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, v)
		currentSimbol, _ := strconv.Atoi(string(buf))
		if len(*numerWithoutSpaces)%2 == 0 {
			if i%2 != 0 {
				modifyedNumber.WriteString(string(currentSimbol))
			} else {
				if currentSimbol*2 > 9 {
					modifyedNumber.WriteString(string(currentSimbol*2 - 9))
				} else {
					modifyedNumber.WriteString(string(currentSimbol * 2))
				}
			}
		} else {
			if i%2 == 0 {
				modifyedNumber.WriteString(string(currentSimbol))
			} else {
				if currentSimbol*2 > 9 {
					modifyedNumber.WriteString(string(currentSimbol*2 - 9))
				} else {
					modifyedNumber.WriteString(string(currentSimbol * 2))
				}
			}
		}
	}

	for _, v := range modifyedNumber.String() {
		newStr += strconv.FormatInt(int64(v), 10)
	}

	*numberModifyed = newStr
}

func fourthPastValidateByDiv10(numberModifyed *string) bool {
	sum := 0

	for _, val := range *numberModifyed {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, val)
		currentSimbol, _ := strconv.Atoi(string(buf))
		sum += currentSimbol
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
