package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var DIFFICULTY = 1
var mapDifficulties = map[int]string{
	1: "Easy",
	2: "Medium",
	3: "Hard",
}

func main() {
	welcomeMessage()
	computerSelection := computerPlays()
	displayDifficultyLevels()
	var choice int
	fmt.Printf("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Fatal("invalid option selected")
	}
	choice = int(choice)
	var retries int
	if choice == 1 {
		retries = 10
	} else if choice == 2 {
		retries = 5
	} else if choice == 3 {
		retries = 3
	} else {
		log.Fatal("invalid selection")
	}

	fmt.Printf("Great! You have selected the %v difficulty level. \n", mapDifficulties[choice])
	fmt.Println("")
	setDifficulty(choice)
	startGame()

	for i := 0; i < retries; i++ {
		usersGuess := collectGuess()
		isGreater := computerSelection > usersGuess
		if usersGuess == computerSelection {
			stopGame()
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", i)
			break
		}
		var msg string
		if isGreater {
			msg = fmt.Sprintf("Incorrect! The number is greater than %v.\n ", usersGuess)
		} else {
			msg = fmt.Sprintf("Incorrect! The number is less than %v.\n", usersGuess)
		}
		fmt.Println(msg)
		var playAgain string
		fmt.Printf("Want to play again Y/N: ")
		_, err := fmt.Scan(&playAgain)
		if err != nil {
			log.Fatal(err)
		}
		playAgain = string(playAgain)
		if playAgain == "Y" {
			// play game again
		}
	}
}

type Leaderboard struct {
	Easy   []int
	Medium []int
	Hard   []int
}

var leaderboardTable Leaderboard

var leaderboard []int

func timer() {
	start := time.Now()
	endTime := time.Since(start)
	fmt.Println("endTime", endTime)
}
func hint() {
	fmt.Println("let me give you a hint") // the number is either 9 or 2 or 13
}
func saveRetries(count int, difficulty int) {
	leaderboard = append(leaderboard, count)
	if difficulty == 1 {
		leaderboardTable.Easy = append(leaderboardTable.Easy, count)
	} else if difficulty == 2 {
		leaderboardTable.Medium = append(leaderboardTable.Medium, count)
	} else if difficulty == 3 {
		leaderboardTable.Hard = append(leaderboardTable.Hard, count)
	} else {
		log.Fatal("Invalid difficulty level")
	}
}

func showLeaderBoard() {
	for _, value := range leaderboardTable.Easy {
		fmt.Printf("Retries: %d", value)
	}
	for _, value := range leaderboardTable.Medium {
		fmt.Printf("Retries: %d", value)
	}
	for _, value := range leaderboardTable.Hard {
		fmt.Printf("Retries: %d", value)
	}
}

type Difficulty struct {
	Easy   string
	Medium string
	Hard   string
}

func collectGuess() int {
	var guess int
	fmt.Printf("Enter your guess: ")
	_, err := fmt.Scan(&guess)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return int(guess)
}

func computerPlays() int {
	return rand.Intn(100)
}

func setDifficulty(difficulty int) {
	if difficulty != 1 && difficulty != 2 && difficulty != 3 {
		log.Fatalf("Invalid selection: %v", difficulty)
	}
	DIFFICULTY = difficulty
}
func startGame() {
	fmt.Println("Let's start the game!")
}
func stopGame() {
	DIFFICULTY = 1
}
func welcomeMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have a maximum of 10 chances to guess the correct number.")
	fmt.Println("")
}
func displayDifficultyLevels() {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println("")
}
