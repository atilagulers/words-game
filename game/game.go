package game

import (
	"fmt"

	"example.com/words-game/words"
)

// StartGame initiates the word game with given file path and delimiters
func StartGame(filePath, delimiters string) error {

	//rand := rand.Intn(len(*wordsMap))
	wordList, err := words.GetWordList(filePath, delimiters)
	if err != nil {
		return err
	}
	fmt.Println(wordList)
	return nil
}
