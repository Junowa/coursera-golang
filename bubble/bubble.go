package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// BubbleSort function
func bubbleSort(sp *[]int) {
	arrayLength := len(*sp)

	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-i-1; j++ {

			if (*sp)[j] > (*sp)[j+1] {
				Swap(sp, j)
			}
		}
	}
}

// Swap function
func Swap(sp *[]int, i int) {
	(*sp)[i], (*sp)[i+1] = (*sp)[i+1], (*sp)[i]
}

func main() {

	// User input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 10 integers:")
	scanner.Split(bufio.ScanWords)

	myList := []int{}
	count := 0
	for scanner.Scan() {

		// string -> int64
		res, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		// int64 -> int32
		myInt := int(res)

		myList = append(myList, myInt)
		count++

		if count == 10 {
			break
		}

	}

	bubbleSort(&myList)

	fmt.Println(myList)
}
