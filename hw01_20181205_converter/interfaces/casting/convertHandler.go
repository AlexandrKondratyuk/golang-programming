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
	case int64:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case int32:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case int16:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case int8:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case uint:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case uint64:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case uint32:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case uint16:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case uint8:
		if bol == 0 {
			return false, nil
		}
		return true, nil
	case string:
		return strconv.ParseBool(bol)
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'BOOL' ", bol, bol)
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
		return "", fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'string' ", str, str)
	}
}

func toFloat64Handler(i interface{}) (float64, error) {
	switch flo := i.(type) {
	case float64:
		return flo, nil
	case float32:
		return float64(flo), nil
	case int:
		return float64(flo), nil
	case int64:
		return float64(flo), nil
	case int32:
		return float64(flo), nil
	case int16:
		return float64(flo), nil
	case int8:
		return float64(flo), nil
	case uint:
		return float64(flo), nil
	case uint64:
		return float64(flo), nil
	case uint32:
		return float64(flo), nil
	case uint16:
		return float64(flo), nil
	case uint8:
		return float64(flo), nil
	case bool:
		if flo {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if value, err := strconv.ParseFloat(flo, 64); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'float64' ", flo, flo)
		} else {
			return value, nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'float64' ", flo, flo)
	}
}

func toFloat32Handler(i interface{}) (float32, error) {
	switch flo := i.(type) {
	case float64:
		return float32(flo), nil
	case float32:
		return flo, nil
	case int:
		return float32(flo), nil
	case int64:
		return float32(flo), nil
	case int32:
		return float32(flo), nil
	case int16:
		return float32(flo), nil
	case int8:
		return float32(flo), nil
	case uint:
		return float32(flo), nil
	case uint64:
		return float32(flo), nil
	case uint32:
		return float32(flo), nil
	case uint16:
		return float32(flo), nil
	case uint8:
		return float32(flo), nil
	case bool:
		if flo {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if value, err := strconv.ParseFloat(flo, 32); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'float32' ", flo, flo)
		} else {
			return float32(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'float32' ", flo, flo)
	}
}

func toIntHandler(i interface{}) (int, error) {
	switch integer := i.(type) {
	case bool:
		if integer {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return int(integer), nil
	case float32:
		return int(integer), nil
	case int:
		return integer, nil
	case int64:
		return int(integer), nil
	case int32:
		return int(integer), nil
	case int16:
		return int(integer), nil
	case int8:
		return int(integer), nil
	case uint:
		return int(integer), nil
	case uint64:
		return int(integer), nil
	case uint32:
		return int(integer), nil
	case uint16:
		return int(integer), nil
	case uint8:
		return int(integer), nil
	case string:
		if value, err := strconv.Atoi(integer); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int' ", integer, integer)
		} else {
			return value, nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int' ", integer, integer)
	}
}

func toInt64Handler(i interface{}) (int64, error) {
	switch integer := i.(type) {
	case bool:
		if integer {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return int64(integer), nil
	case float32:
		return int64(integer), nil
	case int:
		return int64(integer), nil
	case int64:
		return integer, nil
	case int32:
		return int64(integer), nil
	case int16:
		return int64(integer), nil
	case int8:
		return int64(integer), nil
	case uint:
		return int64(integer), nil
	case uint64:
		return int64(integer), nil
	case uint32:
		return int64(integer), nil
	case uint16:
		return int64(integer), nil
	case uint8:
		return int64(integer), nil
	case string:
		if value, err := strconv.Atoi(integer); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int' ", integer, integer)
		} else {
			return int64(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int64' ", integer, integer)
	}
}

func toInt32Handler(i interface{}) (int32, error) {
	switch integer := i.(type) {
	case bool:
		if integer {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return int32(integer), nil
	case float32:
		return int32(integer), nil
	case int:
		return int32(integer), nil
	case int64:
		return int32(integer), nil
	case int32:
		return integer, nil
	case int16:
		return int32(integer), nil
	case int8:
		return int32(integer), nil
	case uint:
		return int32(integer), nil
	case uint64:
		return int32(integer), nil
	case uint32:
		return int32(integer), nil
	case uint16:
		return int32(integer), nil
	case uint8:
		return int32(integer), nil
	case string:
		if value, err := strconv.Atoi(integer); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int32' ", integer, integer)
		} else {
			return int32(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int32' ", integer, integer)
	}
}

func toInt16Handler(i interface{}) (int16, error) {
	switch integer := i.(type) {
	case bool:
		if integer {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return int16(integer), nil
	case float32:
		return int16(integer), nil
	case int:
		return int16(integer), nil
	case int64:
		return int16(integer), nil
	case int32:
		return int16(integer), nil
	case int16:
		return integer, nil
	case int8:
		return int16(integer), nil
	case uint:
		return int16(integer), nil
	case uint64:
		return int16(integer), nil
	case uint32:
		return int16(integer), nil
	case uint16:
		return int16(integer), nil
	case uint8:
		return int16(integer), nil
	case string:
		if value, err := strconv.Atoi(integer); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int16' ", integer, integer)
		} else {
			return int16(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int16' ", integer, integer)
	}
}

func toInt8Handler(i interface{}) (int8, error) {
	switch integer := i.(type) {
	case bool:
		if integer {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return int8(integer), nil
	case float32:
		return int8(integer), nil
	case int:
		return int8(integer), nil
	case int64:
		return int8(integer), nil
	case int32:
		return int8(integer), nil
	case int16:
		return int8(integer), nil
	case int8:
		return integer, nil
	case uint:
		return int8(integer), nil
	case uint64:
		return int8(integer), nil
	case uint32:
		return int8(integer), nil
	case uint16:
		return int8(integer), nil
	case uint8:
		return int8(integer), nil
	case string:
		if value, err := strconv.Atoi(integer); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int8' ", integer, integer)
		} else {
			return int8(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'int8' ", integer, integer)
	}
}

func toUintHandler(i interface{}) (uint, error) {
	switch ui := i.(type) {
	case bool:
		if ui {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return uint(ui), nil
	case float32:
		return uint(ui), nil
	case int:
		return uint(ui), nil
	case int64:
		return uint(ui), nil
	case int32:
		return uint(ui), nil
	case int16:
		return uint(ui), nil
	case int8:
		return uint(ui), nil
	case uint:
		return ui, nil
	case uint64:
		return uint(ui), nil
	case uint32:
		return uint(ui), nil
	case uint16:
		return uint(ui), nil
	case uint8:
		return uint(ui), nil
	case string:
		if value, err := strconv.Atoi(ui); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint' ", ui, ui)
		} else {
			return uint(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint' ", ui, ui)
	}
}

func toUint64Handler(i interface{}) (uint64, error) {
	switch ui := i.(type) {
	case bool:
		if ui {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return uint64(ui), nil
	case float32:
		return uint64(ui), nil
	case int:
		return uint64(ui), nil
	case int64:
		return uint64(ui), nil
	case int32:
		return uint64(ui), nil
	case int16:
		return uint64(ui), nil
	case int8:
		return uint64(ui), nil
	case uint:
		return uint64(ui), nil
	case uint64:
		return ui, nil
	case uint32:
		return uint64(ui), nil
	case uint16:
		return uint64(ui), nil
	case uint8:
		return uint64(ui), nil
	case string:
		if value, err := strconv.Atoi(ui); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint64' ", ui, ui)
		} else {
			return uint64(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint64' ", ui, ui)
	}
}

func toUint32Handler(i interface{}) (uint32, error) {
	switch ui := i.(type) {
	case bool:
		if ui {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return uint32(ui), nil
	case float32:
		return uint32(ui), nil
	case int:
		return uint32(ui), nil
	case int64:
		return uint32(ui), nil
	case int32:
		return uint32(ui), nil
	case int16:
		return uint32(ui), nil
	case int8:
		return uint32(ui), nil
	case uint:
		return uint32(ui), nil
	case uint64:
		return uint32(ui), nil
	case uint32:
		return ui, nil
	case uint16:
		return uint32(ui), nil
	case uint8:
		return uint32(ui), nil
	case string:
		if value, err := strconv.Atoi(ui); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint32' ", ui, ui)
		} else {
			return uint32(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint32' ", ui, ui)
	}
}

func toUint16Handler(i interface{}) (uint16, error) {
	switch ui := i.(type) {
	case bool:
		if ui {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return uint16(ui), nil
	case float32:
		return uint16(ui), nil
	case int:
		return uint16(ui), nil
	case int64:
		return uint16(ui), nil
	case int32:
		return uint16(ui), nil
	case int16:
		return uint16(ui), nil
	case int8:
		return uint16(ui), nil
	case uint:
		return uint16(ui), nil
	case uint64:
		return uint16(ui), nil
	case uint32:
		return uint16(ui), nil
	case uint16:
		return ui, nil
	case uint8:
		return uint16(ui), nil
	case string:
		if value, err := strconv.Atoi(ui); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint16' ", ui, ui)
		} else {
			return uint16(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint16' ", ui, ui)
	}
}

func toUint8Handler(i interface{}) (uint8, error) {
	switch ui := i.(type) {
	case bool:
		if ui {
			return 1, nil
		} else {
			return 0, nil
		}
	case float64:
		return uint8(ui), nil
	case float32:
		return uint8(ui), nil
	case int:
		return uint8(ui), nil
	case int64:
		return uint8(ui), nil
	case int32:
		return uint8(ui), nil
	case int16:
		return uint8(ui), nil
	case int8:
		return uint8(ui), nil
	case uint:
		return uint8(ui), nil
	case uint64:
		return uint8(ui), nil
	case uint32:
		return uint8(ui), nil
	case uint16:
		return uint8(ui), nil
	case uint8:
		return ui, nil
	case string:
		if value, err := strconv.Atoi(ui); err != nil {
			return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint8' ", ui, ui)
		} else {
			return uint8(value), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Error: can't convert type: %T of value: %#v to type 'uint8' ", ui, ui)
	}
}