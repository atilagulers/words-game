package player

import (
	"errors"
	"fmt"

	"example.com/words-game/colors"
)

type Player struct {
	Alphabet   []rune
	GuessCount int
	MaxGuess   int
}

func New() *Player {
	alphabet := createAlphabet()
	return &Player{
		alphabet,
		0,
		10,
	}
}

func (p *Player) PrintAlphabet() {
	colors.PrintWhite("Your alphabet: ")
	for _, char := range p.Alphabet {
		fmt.Printf(colors.Yellow+"%c "+colors.Reset, char)
	}
}

func (p *Player) UseLetter(letter rune) error {
	if !checkLetterInRange(letter) {
		return fmt.Errorf("invalid letter: %c. Letter must be between 'a' and 'z'", letter)
	}

	for i, char := range p.Alphabet {
		if char == letter {
			p.Alphabet[i] = '-'
		}
	}

	return nil
}

func createAlphabet() []rune {
	alphabet := []rune{}

	for ch := 'a'; ch <= 'z'; ch++ {
		alphabet = append(alphabet, ch)
	}

	return alphabet
}

func checkLetterInRange(letter rune) bool {
	if letter < 'a' || letter > 'z' {
		return false
	}
	return true
}
