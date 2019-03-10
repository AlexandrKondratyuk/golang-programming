package convert

import (
	"fmt"
	"strconv"
)

func StringToInt(value string) (result int) {
	if temp, err := strconv.Atoi(value); err == nil {
		result = temp
	}
	return result
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func StringToFloat(value string) (result float64) {
	if temp, err := strconv.ParseFloat(value, 64); err == nil {
		result = temp
	}
	return result
}

func StringToBool(value string) (result bool) {
	if temp, err := strconv.ParseBool(value); err == nil {
		result = temp
	}
	return result
}

func StringToUint(value string) (result uint64) {
	if temp, err := strconv.ParseUint(value, 10, 64); err == nil {
		result = temp
	}
	return result
}

func IntToFloat(value int) (result float64) {
	result = float64(value)
	return result
}

func IntToBool(value int) (result bool) {
	if value != 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func IntToUint(value int) (result uint) {
	result = uint(value)
	return result
}

func UintToInt(value uint) (result int) {
	result = int(value)
	return result
}

func FloatToInt(value float32) (result int) {
	result = int(value)
	return result
}

func BoolToInt(value bool) (result int) {
	if value == true {
		result = 1
	} else {
		result = 0
	}
	return result
}

func FloatToBool(value float64) (result bool) {
	if value != 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func UintToBool(value uint) (result bool) {
	if value != 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func BoolToFloat(value bool) (result float64) {
	if value == true {
		result = 1
	} else {
		result = 0
	}
	return result
}

func FloatToString(value float64) (result string) {
	result = fmt.Sprintf("%f", value)
	return result
}

func FloatToUint(value float32) (result uint) {
	result = uint(value)
	return result
}

func UintToFloat(value uint) (result float64) {
	result = float64(value)
	return result
}

func UintToString(value uint) (result string) {
	result = fmt.Sprintf("%d", value)
	return result
}

func BoolToUint(value bool) (result uint) {
	if value == true {
		result = 1
	} else {
		result = 0
	}
	return result
}

func BoolToString(value bool) (result string) {
	result = fmt.Sprintf("%t", value)
	return result
}
