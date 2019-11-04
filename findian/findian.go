package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Enter string:")
	fmt.Scan(&s)

	if strings.HasPrefix(strings.ToLower(s), "i") && strings.HasSuffix(strings.ToLower(s), "n") && strings.Contains(strings.ToLower(s), "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
