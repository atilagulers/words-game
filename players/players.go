package players

import (
	"fmt"
	"sort"

	"example.com/words-game/colors"
	"example.com/words-game/input"
)

// Player represents a player in the game with an alphabet map, guess count, and maximum guesses allowed.
type Player struct {
	Alphabet   map[rune]bool
	GuessCount int
	MaxGuess   int
}

// New creates and returns a new Player instance with initialized alphabet, guess count, and max guesses.
func New() *Player {
	alphabet := createAlphabet()
	return &Player{
		Alphabet:   alphabet,
		GuessCount: 0,
		MaxGuess:   10,
	}
}

// GetPlayerGuess prompts the player to enter a letter guess and returns the guessed letter.
func (p *Player) GetPlayerGuess() (rune, error) {
	for {
		colors.PrintBlue("\nPlease enter a letter guess: ")
		playerGuess, err := input.GetLetterInput()
		if err != nil {
			return 0, err
		} else {
			return playerGuess, nil
		}
	}
}

// PrintAlphabet prints the player's alphabet, showing used letters as "-" and unused letters in yellow.
func (p *Player) PrintAlphabet() {
	colors.PrintWhite("Your alphabet: ")
	// Collect keys and sort them to print in alphabetical order
	keys := make([]rune, 0, len(p.Alphabet))
	for char := range p.Alphabet {
		keys = append(keys, char)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Print each letter in sorted order, marking used letters with "-"
	for _, char := range keys {
		if p.Alphabet[char] {
			fmt.Printf(colors.Yellow + "- " + colors.Reset)
		} else {
			fmt.Printf(colors.Yellow+"%c "+colors.Reset, char)
		}
	}
	fmt.Println()
}

// UseLetter marks a letter as used in the player's alphabet and returns an error if the letter is invalid or already used.
func (p *Player) UseLetter(letter rune) error {
	if !checkLetterInRange(letter) {
		return fmt.Errorf("invalid letter: %c. Letter must be between 'a' and 'z'", letter)
	}

	for char, used := range p.Alphabet {
		if char == letter {
			if used {
				return fmt.Errorf("%c is already used", letter)
			} else {
				p.Alphabet[char] = true
			}
		}
	}

	return nil
}

// createAlphabet initializes and returns a map representing the alphabet with all letters marked as unused.
func createAlphabet() map[rune]bool {
	alphabet := make(map[rune]bool)

	for ch := 'a'; ch <= 'z'; ch++ {
		alphabet[ch] = false
	}

	return alphabet
}

// checkLetterInRange checks if a given letter is between 'a' and 'z'.
func checkLetterInRange(letter rune) bool {
	if letter < 'a' || letter > 'z' {
		return false
	}
	return true
}
