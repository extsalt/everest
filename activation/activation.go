package activation

func Relu(input float64) float64 {
	if input < 0 {
		return 0
	}

	return input
}
