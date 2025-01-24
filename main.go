package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var DIFFICULTY = 1
var mapDifficulties = map[int]string{
	1: "Easy",
	2: "Medium",
	3: "Hard",
}

type Leaderboard struct {
	Easy   []int
	Medium []int
	Hard   []int
}

type Difficulty struct {
	Easy   string
	Medium string
	Hard   string
}

var leaderboardTable Leaderboard

var leaderboard []int

var currentUserGuess int
var totalTimesPlayed int
var correctAnswerTimePlayed int
var choiceDifficulty int

func main() {

	args := os.Args
	if len(args) == 2 {
		if args[1] == "--leaderboard" {
			showLeaderBoard()
		}
	}

	welcomeMessage()
	computerSelection := computerPlays()
	displayDifficultyLevels()

	fmt.Printf("Enter Difficulty: ")
	_, err := fmt.Scan(&choiceDifficulty)
	if err != nil {
		log.Fatal("invalid option selected")
	}
	choiceDifficulty = int(choiceDifficulty)
	var retries int
	if choiceDifficulty == 1 {
		retries = 10
	} else if choiceDifficulty == 2 {
		retries = 5
	} else if choiceDifficulty == 3 {
		retries = 3
	} else {
		log.Fatal("invalid selection")
	}

	fmt.Printf("Great! You have selected the %v difficulty level. \n", mapDifficulties[choiceDifficulty])
	fmt.Println("")
	setDifficulty(choiceDifficulty)
	startGame()
	var res string

	for {
		timesPlayed := playGame(retries, computerSelection)
		totalTimesPlayed += timesPlayed
		if correctAnswerTimePlayed < 0 {
			hintValue := hintUser(currentUserGuess, computerSelection)
			fmt.Println(hintValue)
		}
		if totalTimesPlayed > 3 {
			fmt.Printf("You have played %v amount of times\n", totalTimesPlayed)
		}
		res = playAgain()
		if res != "Y" && res != "y" {
			fmt.Println("Quitting ....")
			break
		}

	}

}

func playAgain() string {
	var playAgain string
	fmt.Printf("Want to play again Y/N: ")
	_, err := fmt.Scan(&playAgain)
	if err != nil {
		log.Fatal(err)
	}
	playAgain = string(playAgain)
	return playAgain
}

func playGame(tries int, computerSelection int) int {
	var timesPlayed int

	for i := 0; i < tries; i++ {
		usersGuess, duration := collectGuess()
		currentUserGuess += usersGuess
		fmt.Printf("Your guess took: %v seconds \n", duration)
		fmt.Println("")
		isGreater := computerSelection > usersGuess
		if usersGuess == computerSelection {
			totalTimesPlayed++
			correctAnswerTimePlayed += totalTimesPlayed
			saveToLeaderboard(correctAnswerTimePlayed, choiceDifficulty)
			stopGame()
			fmt.Printf("Congratulations! You guessed the correct number on your %v attempt.\n", i)
			break
		}
		var msg string
		if isGreater {
			msg = fmt.Sprintf("Incorrect! The number is greater than %v.\n ", usersGuess)
		} else {
			msg = fmt.Sprintf("Incorrect! The number is less than %v.\n", usersGuess)
		}
		fmt.Println(msg)
	}
	timesPlayed++
	return timesPlayed
}

func hintUser(userGuess int, computerGuess int) string {
	fmt.Println("let me give you a hint") // the number is either 9 or 2 or 13
	var lowPoint int
	var highPoint int
	if computerGuess > userGuess {
		var v0 = computerGuess - userGuess
		v1 := rand.Intn(v0)
		v2 := rand.Intn(v0)
		if v2 > v1 {
			highPoint = v2
			lowPoint = v1
		} else {
			lowPoint = v2
			highPoint = v1
		}
	} else {
		var v0 = userGuess - computerGuess
		v1 := rand.Intn(v0)
		v2 := rand.Intn(v0)
		if v2 > v1 {
			highPoint = v2
			lowPoint = v1
		} else {
			lowPoint = v2
			highPoint = v1
		}
	}

	return fmt.Sprintf("The value is between: %v and %v", highPoint, lowPoint)
}
func saveToLeaderboard(timesPlayed int, difficulty int) {
	fmt.Println("saving to leaderboard")
	leaderboard = append(leaderboard, timesPlayed)
	if difficulty == 1 {
		leaderboardTable.Easy = append(leaderboardTable.Easy, timesPlayed)
	} else if difficulty == 2 {
		leaderboardTable.Medium = append(leaderboardTable.Medium, timesPlayed)
	} else if difficulty == 3 {
		leaderboardTable.Hard = append(leaderboardTable.Hard, timesPlayed)
	} else {
		log.Fatal("Invalid difficulty level")
	}
	showLeaderBoard()
}

func showLeaderBoard() {
	if len(leaderboardTable.Easy) > 0 {
		for key, value := range leaderboardTable.Easy {
			fmt.Println("Easy")
			fmt.Printf("%d: %d\n", key+1, value)
		}
	}
	if len(leaderboardTable.Medium) > 0 {
		for _, value := range leaderboardTable.Medium {
			fmt.Printf("Retries: %d\n", value)
		}
	}
	if len(leaderboardTable.Hard) > 0 {
		for _, value := range leaderboardTable.Hard {
			fmt.Printf("Retries: %d\n", value)
		}
	}
}

func collectGuess() (int, time.Duration) {
	start := time.Now()

	var guess int
	fmt.Printf("Enter your guess: ")
	_, err := fmt.Scan(&guess)
	endTime := time.Since(start)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return int(guess), endTime
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
