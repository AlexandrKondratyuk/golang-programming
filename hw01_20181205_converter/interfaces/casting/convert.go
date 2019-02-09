package casting

func ToBool(i interface{}) (bool, error) {
	return toBoolHandler(i)
}

func ToString(i interface{}) (string, error) {
	return toStringHandler(i)
}

func ToFloat64(i interface{}) (float64, error) {
	return toFloat64Handler(i)
}

func ToFloat32(i interface{}) (float32, error) {
	return toFloat32Handler(i)
}

func ToIntHandler (i interface{}) (int, error) {
	return toIntHandler(i)
}

func ToInt64Handler (i interface{}) (int64, error) {
	return toInt64Handler(i)
}

func ToInt32Handler (i interface{}) (int32, error) {
	return toInt32Handler(i)
}


func ToInt16Handler (i interface{}) (int16, error) {
	return toInt16Handler(i)
}

func ToInt8Handler (i interface{}) (int8, error) {
	return toInt8Handler(i)
}

func ToUintHandler (i interface{}) (uint, error) {
	return toUintHandler(i)
}

func ToUint64Handler (i interface{}) (uint64, error) {
	return toUint64Handler(i)
}

func ToUint32Handler (i interface{}) (uint32, error) {
	return toUint32Handler(i)
}

func ToUint16Handler (i interface{}) (uint16, error) {
	return toUint16Handler(i)
}

func ToUint8Handler (i interface{}) (uint8, error) {
	return toUint8Handler(i)
}