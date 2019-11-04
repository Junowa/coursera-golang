package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct
type Cow struct {
	food       string
	locomotion string
	sound      string
}

// Eat prints cow food
func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

// Move prints cow locomotion
func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

// Speak prints cow sound
func (cow Cow) Speak() {
	fmt.Println(cow.sound)
}

// Bird struct
type Bird struct {
	food       string
	locomotion string
	sound      string
}

// Eat prints bird food
func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

// Move prints bird locomotion
func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

// Speak prints bird sound
func (bird Bird) Speak() {
	fmt.Println(bird.sound)
}

// Snake struct
type Snake struct {
	food       string
	locomotion string
	sound      string
}

// Eat prints snake food
func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

// Move prints snake locomotion
func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

// Speak prints snake sound
func (snake Snake) Speak() {
	fmt.Println(snake.sound)
}

func main() {

	// Map to store animals
	// Key is the name of the animal
	// Value is the interface type called Animal, which is a struct containing three fields, all of which are strings.
	animalmap := make(map[string]Animal)

	scanner := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("> ")

		line, err := scanner.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		array := strings.Fields(line)

		if line != "" && len(array) == 3 {

			command := array[0]
			name := array[1]
			info := array[2]

			switch command {
			case "newanimal":
				switch info {
				case "cow":

					cow := Cow{food: "grass", locomotion: "walk", sound: "moo"}
					animalmap[name] = cow

					fmt.Println("Created it!")

				case "bird":

					bird := Bird{food: "worms", locomotion: "fly", sound: "beep"}
					animalmap[name] = bird

					fmt.Println("Created it!")

				case "snake":

					snake := Snake{food: "mice", locomotion: "slither", sound: "hsss"}
					animalmap[name] = snake

					fmt.Println("Created it!")

				default:
					fmt.Println("Please ask again")
				}
			case "query":
				switch info {
				case "eat":
					animalmap[name].Eat()

				case "move":
					animalmap[name].Move()

				case "speak":
					animalmap[name].Speak()

				default:
					fmt.Println("Please ask again")
				}
			}
		}
	}
}
