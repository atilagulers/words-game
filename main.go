package main

import (
	"log"

	"example.com/words-game/game"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func main() {
	const filePath string = "words.txt"
	delimiters := `[',:;.\s()]+`

	err := game.StartGame(filePath, delimiters)
	if err != nil {
		log.Fatal(err)
	}
}
