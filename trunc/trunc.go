package main

import "fmt"

func main() {
	var n float32
	fmt.Println("Enter a floating number:")
	fmt.Scan(&n)
	fmt.Printf("Truncated number is: ")
	fmt.Printf("%d\n", int(n))
}
