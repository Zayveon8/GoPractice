package main

//Import needed packages
import "fmt"
import "math/rand"
import "time"

//Main function
func main() {

	// Start the count of the number of tries
	rand.Seed(time.Now().UnixNano())

	// generate random number between 1 and 100
	randomNumber := rand.Intn(100) + 1

	// Show welcome message
	fmt.Println("Welcome to the Guessing Game!")

	// Prompt the user to guess a number 1-100
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")

	// Initialize count
	guessCount := 0

	// For loop that only terminates when user guess correctly
	for {
		// Tell user to enter guess
		fmt.Print("Enter your guess: ")
		var guess int
		fmt.Scan(&guess)

		// Increment the guess count
		guessCount++

		// Check if the guess is correct
		if guess == randomNumber {
			fmt.Printf("Congratulations! You guessed the number in %d tries.\n", guessCount)
			break
		} else if guess < randomNumber {
			fmt.Println("Too low. Guess again!")
		} else {
			fmt.Println("Too high. Guess again!")
		}
	}
}