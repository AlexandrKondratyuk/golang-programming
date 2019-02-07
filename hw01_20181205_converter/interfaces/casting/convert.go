package casting

func ToBool(i interface{}) (bool, error) {
	return toBoolHandler(i)
}
