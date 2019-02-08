package casting

import (
	"fmt"
	"strconv"
)

func toBoolHandler(i interface{}) (bool, error) {
	switch bol := i.(type) {
	case bool:
		return bol, nil
	case int:
		if bol == 0 { return false, nil }
		return true, nil
	case int64:
		if bol == 0 { return false, nil }
		return true, nil
	case int32:
		if bol == 0 { return false, nil }
		return true, nil
	case int16:
		if bol == 0 { return false, nil }
		return true, nil
	case int8:
		if bol == 0 { return false, nil }
		return true, nil
	case uint:
		if bol == 0 { return false, nil }
		return true, nil
	case uint64:
		if bol == 0 { return false, nil }
		return true, nil
	case uint32:
		if bol == 0 { return false, nil }
		return true, nil
	case uint16:
		if bol == 0 { return false, nil }
		return true, nil
	case uint8:
		if bol == 0 { return false, nil }
		return true, nil
	case string:
		return strconv.ParseBool(bol)
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("Error: can't convert type %T of %#v to type BOOL", bol, bol)
	}
}

func toStringHandler(i interface{}) (string, error) {
	switch str := i.(type) {
	case string:
		return str, nil
	case int:
		return strconv.Itoa(str), nil
	case float64:
		return strconv.FormatFloat(str, 'f', 2, 64), nil
	case float32:
		return strconv.FormatFloat(float64(str), 'f', 2, 32), nil
	case int64:
		return strconv.FormatInt(str, 64), nil
	case int32:
		return strconv.FormatInt(int64(str), 10), nil
	case int16:
		return strconv.FormatInt(int64(str), 10), nil
	case int8:
		return strconv.FormatInt(int64(str), 10), nil
	case uint:
		return strconv.FormatInt(int64(str), 10), nil
	case uint64:
		return strconv.FormatInt(int64(str), 10), nil
	case uint32:
		return strconv.FormatInt(int64(str), 10), nil
	case uint16:
		return strconv.FormatInt(int64(str), 10), nil
	case uint8:
		return strconv.FormatInt(int64(str), 10), nil
	case bool:
		return strconv.FormatBool(str), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("Error: can't convert type %T of %#v to type 'string'", str, str)
	}
}
