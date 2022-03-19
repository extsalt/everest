package main

import (
	"fmt"
	"math"
)

const LEN = 100

var trainData []float64
var trainTarget []float64

func fit(input float64) float64 {
	return input + 10
}


type Neuron struct {
	//input for neuron
	 input [] float64
	 //activation func
	 activation string
}

func (n *Neuron) setInput (input [] float64) {
	n.input = input
}

func (n *Neuron) setActivationFunction (function string) {
	n.activation = function
}

func (n *Neuron) output () float64 {
	sum := float64(0)

	for _, i  := range n.input {
		sum += i
	}

	return sum
}

func initTrainingData() {
	for i := 0; i < LEN; i++ {
		trainData = append(trainData, float64(i))
	}

	for i := 0; i < LEN; i++ {
		trainTarget = append(trainTarget, predict(trainData[i]))
	}
}

func dump() {
	for i := range trainData {
		fmt.Println(trainData[i])
	}

	for j := range trainTarget {
		fmt.Println(trainTarget[j])
	}
}

func predict(input float64) float64 {
	return 2*input + 10
}

func main() {
	initTrainingData()
	dump()
	var target []float64
	totalError := 0.0
	for i := 0; i < LEN; i++ {
		target = append(target, fit(trainData[i]))
		totalError += trainTarget[i] - target[i]
	}

	avgError := totalError / LEN

	fmt.Println("avg err: " + fmt.Sprint(avgError))
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
	return 1 / (1 + math.Pow(math.E, -input))
}
