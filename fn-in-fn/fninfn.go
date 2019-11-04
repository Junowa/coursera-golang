package main

import (
	"fmt"
	"math"
)

// MakeDistOrigin function returns function
func MakeDistOrigin(ox, oy float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-ox, 2) + math.Sqrt(math.Pow(y-oy, 2)))
	}
	return fn
}

func main() {
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(0, 2)

	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))
}
