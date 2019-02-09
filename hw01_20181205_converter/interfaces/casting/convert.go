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

func ToInt (i interface{}) (int, error) {
	return toIntHandler(i)
}

func ToInt64 (i interface{}) (int64, error) {
	return toInt64Handler(i)
}

func ToInt32 (i interface{}) (int32, error) {
	return toInt32Handler(i)
}


func ToInt16 (i interface{}) (int16, error) {
	return toInt16Handler(i)
}

func ToInt8 (i interface{}) (int8, error) {
	return toInt8Handler(i)
}

func ToUint (i interface{}) (uint, error) {
	return toUintHandler(i)
}

func ToUint64 (i interface{}) (uint64, error) {
	return toUint64Handler(i)
}

func ToUint32 (i interface{}) (uint32, error) {
	return toUint32Handler(i)
}

func ToUint16 (i interface{}) (uint16, error) {
	return toUint16Handler(i)
}

func ToUint8 (i interface{}) (uint8, error) {
	return toUint8Handler(i)
}