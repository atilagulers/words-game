package game

import (
	"fmt"

	"example.com/words-game/colors"
	"example.com/words-game/input"
	"example.com/words-game/words"
)

// StartGame initiates the word game with given file path and delimiters
func StartGame(filePath, delimiters string) error {

	colors.PrintBlue("Game is starting...\n")
	//time.Sleep(1 * time.Second)
	wordList, err := words.GetWordList(filePath, delimiters)
	if err != nil {
		return err
	}

	answerWord := words.PickRandomWord(wordList)

	// Print the word list
	//for _, word := range wordList {
	//	fmt.Printf("%+v\n", word)
	//}

	// Print the randomly picked word
	fmt.Printf("Answer: %+v\n", answerWord)

	fmt.Println(answerWord.CryptedContent)

	var userGuess rune
	for {
		var isGameFinished bool = false
		colors.PrintBlue("\nPlease enter a letter guess: ")
		userGuess, err = input.GetLetterInput()
		if err != nil {
			fmt.Printf("Invalid input.\nErr: %v", err)
			continue
		}

		letterExist := answerWord.CheckLetterExist(string(userGuess))

		if letterExist {
			isGameFinished = answerWord.RevealLetter(userGuess)
		}

		fmt.Println(answerWord.CryptedContent)

		if isGameFinished {

			colors.PrintGreen("You win!")
			break

		}

	}

	return nil
}
