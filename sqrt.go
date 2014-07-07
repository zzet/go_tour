package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 1000; i++ {
		z = z - (z*z-x)/2*z
		if goodEnaugh(z, x) {
			return z
		}
	}
	return z
}

func goodEnaugh(guess, x float64) bool {
	return abs(guess*guess-x) < 0.01
}

func abs(x float64) float64 {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
