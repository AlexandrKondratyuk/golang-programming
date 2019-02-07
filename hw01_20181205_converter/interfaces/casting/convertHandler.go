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
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case string:
		return strconv.ParseBool(bol)
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("Error: can't convert type %T of %#v to type BOOL",bol, bol)
	}
}
