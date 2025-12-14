package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var number string = "" // random number between 1 and 100

var chances int = 0 // remaining chances

var difficulty string = "" // number put in by player to choose difficulty

var str_difficulty string = "" // difficulty name based on difficulty number, for printing

var attempts int = 0 // number of attempts made

var guess string = "" // player's guess

func main() {
	fmt.Println("Starting Number Guessing Game")
	intro()
}

func intro() {
	// Introduction and difficulty selection //

	number = strconv.Itoa(rand.Intn(100) + 1) // generate random number between 1 and 100

	fmt.Println("\nWelcome to my number guessing game!\nI'm thinking of a number between 1 and 100.\nCan you guess what it is?")
	fmt.Println("\nChoose your difficulty: \n 1. Easy (10 chances) \n 2. Hard (5 chances) \n 3. Almost Impossible (1 chance) \n 4. Impossible (0 chances)")

	fmt.Scan(&difficulty)

	switch difficulty {
	case "1":
		chances = 10
		str_difficulty = "Easy"
	case "2":
		chances = 5
		str_difficulty = "Hard"
	case "3":
		chances = 1
		str_difficulty = "Almost Impossible"
	case "4":
		chances = 0
		str_difficulty = "Impossible"
	default:
		fmt.Println("\nYou didn't pick a valid number, so number 4 it is.")
		chances = 0

	}

	fmt.Println("\nYou chose number ", difficulty, ", which is ", str_difficulty, ". Now let's begin the game!")

	if difficulty == "4" {
		fmt.Println("\noh, would you look at that, you lost. I told you it was impossible. The number was ", number, ".")
	}

	gameLoop()
}

//--- Game loop ---//

func gameLoop() {
	if chances == 0 { // lose state

		fmt.Println("\nYou ran out of guesses, The number was ", number, "!")
		fmt.Println("\nYou made ", attempts, " attempts. Type 1 to play again, or type anything else to quit:")

		var playAgain string

		fmt.Scan(&playAgain)

		if playAgain == "1" {
			// reset game state
			attempts = 0
			intro()
		} else {
			return
		}
	}

	fmt.Println("\nYou have ", chances, " chances left. Type your guess:")
	fmt.Scan(&guess)
	if guess == number { // right guess, win state
		fmt.Println("\nCongratulations! You guessed the number ", number, " correctly!")
		fmt.Println("\nYou made ", attempts, " attempts. Type 1 to play again, or type anything else to quit:")

		var playAgain string

		fmt.Scan(&playAgain)

		if playAgain == "1" {
			// reset game state
			attempts = 0
			intro()
		} else {
			return
		}

	} else { // wrong guess

		attempts += 1
		chances -= 1

		fmt.Println("\nWrong guess. ")
		if guess < number {
			fmt.Println("Too low!") // lower than number
			gameLoop()
		} else {
			fmt.Println("Try a lower number!") // higher than number
			gameLoop()
		}
	}
}
