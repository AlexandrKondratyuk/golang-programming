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
