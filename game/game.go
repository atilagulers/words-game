package game

import (
	"fmt"

	"example.com/words-game/colors"
	"example.com/words-game/players"
	"example.com/words-game/words"
)

// StartGame initiates the word game with given file path and delimiters
func StartGame(filePath, delimiters string) error {
	colors.PrintBlue("Game is starting...\n")

	wordList, err := words.GetWordList(filePath, delimiters)
	if err != nil {
		return err
	}

	answerWord := words.PickRandomWord(wordList)

	player := players.New()
	var playerGuess rune

	for {
		printWordAndAlphabet(answerWord, player)

		// Get player guess
		playerGuess, err = player.GetPlayerGuess()
		if err != nil {
			fmt.Printf(colors.Red+"\nErr: %v\n"+colors.Reset, err)
			continue
		}

		// Use player letter from alphabet
		err := player.UseLetter(playerGuess)
		if err != nil {
			fmt.Printf(colors.Red+"\nErr: %v\n"+colors.Reset, err)
			continue
		}

		isGameFinished := chooseLetter(playerGuess, answerWord)

		if isGameFinished {
			colors.PrintGreen("You win!")
			break
		}

	}

	return nil
}

func chooseLetter(playerGuess rune, answerWord *words.Word) bool {
	letterExist := answerWord.CheckLetterExist(string(playerGuess))
	var isGameFinished bool
	if letterExist {
		colors.PrintGreen("Correct!\n")
		isGameFinished = answerWord.RevealLetter(playerGuess)
	} else {
		colors.PrintRed("Wrong!\n")
	}

	return isGameFinished
}

func printWordAndAlphabet(answerWord *words.Word, player *players.Player) {
	fmt.Printf(colors.White + "\nWord: " + colors.Reset)
	fmt.Printf(colors.Magenta+"%v\n"+colors.Reset, answerWord.CryptedContent)
	player.PrintAlphabet()

}
