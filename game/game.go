package game

import (
	"fmt"

	"example.com/words-game/colors"
	"example.com/words-game/input"
	"example.com/words-game/player"
	"example.com/words-game/words"
)

// StartGame initiates the word game with given file path and delimiters
func StartGame(filePath, delimiters string) error {
	player := player.New()

	colors.PrintBlue("Game is starting...\n")
	//time.Sleep(1 * time.Second)
	wordList, err := words.GetWordList(filePath, delimiters)
	if err != nil {
		return err
	}

	answerWord := words.PickRandomWord(wordList)

	var playerGuess rune
	for {
		fmt.Printf(colors.Cyan+"\nWord: %v\n"+colors.Reset, answerWord.CryptedContent)

		player.PrintAlphabet()

		colors.PrintBlue("\nPlease enter a letter guess: ")

		playerGuess = getPlayerGuess()

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

func getPlayerGuess() rune {

	for {
		playerGuess, err := input.GetLetterInput()
		if err != nil {
			fmt.Printf(colors.Red+"\nErr: %v\n"+colors.Reset, err)
		} else {
			return playerGuess
		}
	}
}

func checkPlayerGuess(playerGuess rune, player *player.Player) error {
	err := player.UseLetter(playerGuess)

	if err != nil {
		return err
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
