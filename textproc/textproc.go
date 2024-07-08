package textproc

import (
	"regexp"
	"strings"
)

// splitByDelimiters splits a slice of strings by given delimiters.
func SplitByDelimiters(text []string, delimiters string) []string {
	delimitersRegex := regexp.MustCompile(delimiters)
	// convert slice to string
	textStr := strings.Join(text, " ")

	splitText := delimitersRegex.Split(textStr, -1)

	return splitText
}

// filterSliceByLength filters strings in a slice by a minimum length.
func FilterSliceByLength(slice []string, minLength int) []string {
	newSlice := []string{}
	for _, val := range slice {
		if len(val) >= minLength {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}

// createMapFromSlice creates a map counting the occurrences of each element in the slice.
func CreateMapFromSlice[T string | int | float64](slice []T) *map[T]int {
	resultMap := make(map[T]int)

	for _, val := range slice {

		// if key exist increase count else create
		_, ok := resultMap[val]
		if ok {
			resultMap[val]++
		} else {
			resultMap[val] = 1
		}
	}

	return &resultMap
}
