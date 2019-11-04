package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// Define struct
	type Person struct {
		fname [20]byte //byte array to store 20 characters strings
		lname [20]byte
	}

	// Create slice
	mySlice := []Person{}

	// Get filename
	var filename string
	fmt.Println("Enter your filename:")
	// Read until first space
	fmt.Scan(&filename)

	// open file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read File
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		// split line in fields
		fields := strings.Split(string(line), " ")

		// init struct
		p := new(Person)

		// convert string to array
		copy(p.fname[:], fields[0])
		copy(p.lname[:], fields[1])

		// Append Struct to slice
		mySlice = append(mySlice, *p)
	}

	// Iterate through slice to d
	for _, elem := range mySlice {
		fmt.Printf("%v %v\n", string(elem.fname[:20]), string(elem.lname[:20]))
	}
}
