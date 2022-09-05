package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Println("Int Non-Generic Sum=", IntSum(ints))
	fmt.Println("Floats Non-Generic Sum=", FloatSum(floats))
	fmt.Println("Int Generic Sum=", IntFloatSum(ints))
	fmt.Println("Floats Generic Sum=", IntFloatSum(floats))

}
func IntSum(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}
func FloatSum(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func IntFloatSum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
