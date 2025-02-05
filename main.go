package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
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

const leaderboardFileName = "leaderboard.json"

type Leaderboard struct {
	// Easy   []RecordType `json:"easy"`
	// Medium []RecordType `json:"medium"`
	// Hard   []RecordType `json:"hard"`
	// RecordType
	TimesPlayed int
	Date        time.Time
	Difficulty  string
}

type RecordType struct {
	TimesPlayed int
	Date        time.Time
	Difficulty  string
}

type Difficulty struct {
	Easy   string `json:"easy"`
	Medium string `json:"medium"`
	Hard   string `json:"hard"`
}

var leaderboardTable []Leaderboard

// var leaderboard []int

var currentUserGuess int
var totalTimesPlayed int
var correctAnswerTimePlayed int
var choiceDifficulty int

func displayHelpfulTexts() {
	fmt.Println("guess [flags]")
	fmt.Println("guess [command]")
	fmt.Println("")

	fmt.Println("Available Commands:")
	fmt.Println("leaderboard           Displays the leaderboard")
	fmt.Println("help 	              Displays this helpful text")
}

func main() {

	args := os.Args
	fmt.Println("args", args)
	if len(args) == 2 {
		fmt.Println("-----leaderboard-----")
		if args[1] == "--leaderboard" {
			showLeaderBoard()
			os.Exit(1)
		}
	}

	createLeaderboardFileIfNotExists()
	welcomeMessage()
	displayHelpfulTexts()
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
	//computer 80
	//user 52

	//user is low
	//pick a number between
	// hint should be 80-52/= 28
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

func createLeaderboardFileIfNotExists() {
	fsys := os.DirFS(".")
	_, err := fs.Stat(fsys, leaderboardFileName)
	if err != nil {
		fmt.Println(leaderboardFileName, " does not exist")
		fmt.Println("creating", leaderboardFileName)
		lfile, err := os.Create(leaderboardFileName)
		if err != nil {
			log.Fatal("error creating leaderboard file")
		}
		lfile.Close()
	}
}
func saveToLeaderboard(timesPlayed int, difficulty int) {
	fmt.Println("saving to leaderboard")
	if difficulty == 1 {

		leaderboardTable = append(leaderboardTable, Leaderboard{TimesPlayed: timesPlayed, Date: time.Now(), Difficulty: "easy"})
	} else if difficulty == 2 {
		leaderboardTable = append(leaderboardTable, Leaderboard{TimesPlayed: timesPlayed, Date: time.Now(), Difficulty: "medium"})

	} else if difficulty == 3 {
		leaderboardTable = append(leaderboardTable, Leaderboard{TimesPlayed: timesPlayed, Date: time.Now(), Difficulty: "hard"})
	} else {
		log.Fatal("Invalid difficulty level")
	}

	fsys := os.DirFS(".")
	fileByte, err := fs.ReadFile(fsys, leaderboardFileName)
	if err != nil {
		createLeaderboardFileIfNotExists()
	}
	err = json.Unmarshal(fileByte, &leaderboardTable)
	fmt.Println("leaderboardTable", leaderboardTable)
	if err != nil {
		log.Fatal("error bytes to struct", err)
	}

	byteToSave, err := json.Marshal(leaderboardTable)
	if err != nil {
		log.Fatal("unable to convert struct to byte")
	}
	err = os.WriteFile(leaderboardFileName, byteToSave, 0644)
	if err != nil {
		log.Fatal("unable to save to leaderboard")
	}
	showLeaderBoard()
}

func showLeaderBoard() {
	fsByte, err := fs.ReadFile(os.DirFS("."), leaderboardFileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fsByte, &leaderboardTable)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range leaderboardTable {
		fmt.Printf("%d: %v\n", key+1, value)
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
