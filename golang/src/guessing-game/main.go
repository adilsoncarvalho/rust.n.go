package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Guess the number!")

	// let secret_number = rand::thread_rng().gen_range(1..=100);
	var secretNumber = rand.Intn(100) + 1

	// println!("The secret number is: {}", secret_number);
	fmt.Println("The secret number is: ", secretNumber)

	for {
		fmt.Println("Please input your guess.")

		var guess string

		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Failed to read line")
			os.Exit(1)
		}

		if guess == "quit" {
			fmt.Println("Quitting...")
			break
		}

		var guessInt int

		guessInt, err = strconv.Atoi(guess)

		if err != nil {
			fmt.Println("Invalid entry, please try again.")
			continue
		}

		if guessInt < secretNumber {
			fmt.Println("Too small!")
		} else if guessInt > secretNumber {
			fmt.Println("Too big!")
		} else {
			fmt.Println("You win!")
			break
		}

		fmt.Println("You guessed: ", guessInt)
	}
}
