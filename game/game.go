package game

import (
	"fmt"

	"example.com/words-game/colors"
	"example.com/words-game/input"
	"example.com/words-game/words"
)

type Player struct {
	GuessedLetters map[rune]int
}

// StartGame initiates the word game with given file path and delimiters
func StartGame(filePath, delimiters string) error {

	colors.PrintBlue("Game is starting...\n")
	//time.Sleep(1 * time.Second)
	wordList, err := words.GetWordList(filePath, delimiters)
	if err != nil {
		return err
	}

	answerWord := words.PickRandomWord(wordList)

	// Print the randomly picked word
	fmt.Printf("Answer: %+v\n", answerWord)

	fmt.Println("Word:", answerWord.CryptedContent)

	var userGuess rune
	for {
		colors.PrintBlue("\nPlease enter a letter guess: ")
		userGuess = getUserGuess()

		isGameFinished := checkUserGuess(userGuess, answerWord)

		fmt.Println("Word:", answerWord.CryptedContent)

		if isGameFinished {
			colors.PrintGreen("You win!")
			break
		}

	}

	return nil
}

func getUserGuess() rune {

	for {
		userGuess, err := input.GetLetterInput()
		if err != nil {
			fmt.Printf(colors.Red+"\nErr: %v\n"+colors.Reset, err)
		} else {
			return userGuess
		}
	}
}

func checkUserGuess(userGuess rune, answerWord *words.Word) bool {
	letterExist := answerWord.CheckLetterExist(string(userGuess))
	var isGameFinished bool
	if letterExist {
		colors.PrintGreen("Correct!\n")
		isGameFinished = answerWord.RevealLetter(userGuess)
	} else {
		colors.PrintRed("Wrong!\n")
	}

	return isGameFinished
}
