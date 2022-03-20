package activation

import "math"

func Relu(input float64) float64 {
	return math.Max(0, input)
}

func LeakyRelu(input float64) float64 {
	return math.Max(0.1*input, input)
}

func ParametricRelu(input float64, scalar float64) float64 {
	return math.Max(input*scalar, input)
}

func Sigmoid(input float64) float64 {
	return 1 / (1 + math.Pow(math.E, -input))
}
