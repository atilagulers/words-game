package main

import (
	"log"

	"example.com/words-game/game"
)

func main() {
	const filePath string = "words.txt"
	delimiters := `[',:;.\s()]+`

	err := game.StartGame(filePath, delimiters)
	if err != nil {
		log.Fatal(err)
	}
}
