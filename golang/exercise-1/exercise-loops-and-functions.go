package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		newz := z - (z*z - x) / (2*z)
		if d := math.Abs(newz - z); d < 1e-9 {
			break;
		}
		z = newz
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("math.Sqrt :", math.Sqrt(2))
}
