package main

import (
	"fmt"
)

//GenDisplaceFn function
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return a*t*t/2 + vo*t + so
	}
	return fn
}

func main() {

	var a float64
	fmt.Print("acceleration= ")
	fmt.Scanf("%f", &a)

	var vo float64
	fmt.Print("initial velocity= ")
	fmt.Scanf("%f", &vo)

	var so float64
	fmt.Print("initial displacement= ")
	fmt.Scanf("%f", &so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Println(fn(3))

	fmt.Println(fn(5))

}
