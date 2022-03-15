package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello everest")
}

// ReLU /**
func ReLU(input float64) float64 {
	if input < 0 {
		return 0
	}

	return input
}

func LeackyRelU(input float64) float64 {
	if input < 0 {
		return input / 10
	}

	return input
}

func Sigmoid(input float64) float64 {
	return 1 / (1 + math.Pow(math.E, input))
}
