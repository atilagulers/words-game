package input

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// GetLetterInput reads a single letter from the standard input.
func GetLetterInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if len(input) != 1 {
		return "", errors.New("please enter a single letter")
	}

	return string(input[0]), nil
}
