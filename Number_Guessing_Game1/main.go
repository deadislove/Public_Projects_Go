package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	// Define the range for the random number
	min := 1
	max := 100

	// Generate a secure random secret number in the range [min, max]
	secretNumber, err := generateSecureRandomNumber(min, max)
	if err != nil {
		log.Fatal(err)
	}

	// Print initial message
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a number between", min, "and", max, ". Try to guess it!")

	var guess int
	attempts := 0

	// Start guessing loop
	for {
		attempts++
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanln(&guess) // Use Scanln to read input
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			// Clear any remaining input in the buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		// Check the guess
		if guess < secretNumber {
			fmt.Println("Too Low!")
		} else if guess > secretNumber {
			fmt.Println("Too High!")
		} else {
			fmt.Printf("Congratulations! You've guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}

// generateSecureRandomNumber generates a secure random number in the range [min, max]
func generateSecureRandomNumber(min, max int) (int, error) {
	// Calculate the range size
	rangeSize := big.NewInt(int64(max - min + 1))

	// Generate a secure random number in the range [0, rangeSize)
	secureRandom, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}

	// Add the minimum value to shift the range to [min, max]
	return int(secureRandom.Int64()) + min, nil
}
