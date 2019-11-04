package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Animal struct
type Animal struct {
	food       string
	locomotion string
	sound      string
}

// Eat returns food
func (animal Animal) Eat() string {
	return animal.food
}

// Move returns locomotion
func (animal Animal) Move() string {
	return animal.locomotion
}

// Speak return sound
func (animal Animal) Speak() string {
	return animal.sound
}

func main() {

	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		sound:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		sound:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		sound:      "hsss",
	}

	scanner := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("> ")

		line, err := scanner.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		array := strings.Fields(line)

		if line != "" && len(array) == 2 {

			name := array[0]
			action := array[1]

			switch name {
			case "cow":
				switch action {
				case "eat":
					fmt.Println(cow.Eat())
				case "move":
					fmt.Println(cow.Move())
				case "speak":
					fmt.Println(cow.Speak())
				default:
					fmt.Println("Please ask again")
				}
			case "bird":
				switch action {
				case "eat":
					fmt.Println(bird.Eat())
				case "move":
					fmt.Println(bird.Move())
				case "speak":
					fmt.Println(bird.Speak())
				default:
					fmt.Println("Please ask again")
				}
			case "snake":
				switch action {
				case "eat":
					fmt.Println(snake.Eat())
				case "move":
					fmt.Println(snake.Move())
				case "speak":
					fmt.Println(snake.Speak())
				default:
					fmt.Println("Please ask again")
				}
			}
		}
	}

}
