package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var input string
	sli := make([]int, 0)

	for {
		fmt.Println("Enter an int (or X to exit):")
		fmt.Scan(&input)

		if v, err := strconv.Atoi(input); err == nil {

			sli = append(sli, v)
			sort.Ints(sli)

		} else {
			if input == "X" {
				break

			} else {
				fmt.Println("...not an integer, try again...")
				continue
			}
		}
		fmt.Println(sli)
	}
}
