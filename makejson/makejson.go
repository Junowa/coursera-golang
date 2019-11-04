package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var name string
	fmt.Println("Enter a name:")
	scanner.Scan()
	name = scanner.Text()

	var address string
	fmt.Println("Enter an address:")
	scanner.Scan()
	address = scanner.Text()

	var user map[string]string
	user = make(map[string]string)

	user["name"] = name
	user["address"] = address

	b, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(b))
	}

}
