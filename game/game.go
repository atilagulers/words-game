package game

import (
	"fmt"

	"example.com/words-game/colors"
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

	randomWord := words.PickRandomWord(wordList)

	// Print the word list
	for _, word := range wordList {
		fmt.Printf("%+v\n", word)
	}

	// Print the randomly picked word
	fmt.Printf("Random Word: %+v\n", randomWord)
	return nil
}
