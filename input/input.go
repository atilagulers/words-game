package input

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// GetLetterInput reads a single letter from the standard input.
func GetLetterInput() (rune, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	if len(input) != 1 {
		return 0, errors.New("please enter a single letter")
	}

	return rune(input[0]), nil
}
