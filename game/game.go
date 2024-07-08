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

	var userGuess string
	for {
		colors.PrintBlue("\nPlease enter a letter guess")
		userGuess, err = input.GetLetterInput()
		if err != nil {
			fmt.Printf("Invalid input. Your guess: %v, Err: %v", userGuess, err)
			continue
		}
		break
	}

	letterExist := answerWord.CheckLetterExist("e")

	if letterExist {
		fmt.Println("LETTER FOUND !!!!!")
		answerWord.RevealLetter('e')
	}

	fmt.Printf("Random Word: %+v\n", answerWord)

	return nil
}
