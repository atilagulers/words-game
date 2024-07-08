package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type Word struct {
	index   int
	content string
	count   int
	length  int
}

func NewWord(index, count, length int, content string) Word {
	return Word{
		index,
		content,
		count,
		length,
	}
}

func main() {
	const filePath string = "words.txt"
	delimiters := `[',:;.\s()]+`

	lines := readLinesFromFile(filePath)
	lines = splitByDelimiters(lines, delimiters)
	lines = filterSliceByLength(lines, 3)

	wordsMap := mapSlice(lines)

	//for key, val := range *wordsMap {
	//	fmt.Println(key, val)
	//}

	rand := rand.Intn(len(*wordsMap))
	fmt.Println(wordsMap)
}

// Reads a file and returns its content as a slice of strings
func readLinesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	return lines
}

// Splits a slice of strings by given delimiters
func splitByDelimiters(text []string, delimiters string) []string {
	delimitersRegex := regexp.MustCompile(delimiters)
	// convert slice to string
	textStr := strings.Join(text, " ")

	splitText := delimitersRegex.Split(textStr, -1)

	return splitText
}

// Filters strings in a slice by a minimum length
func filterSliceByLength(slice []string, minLength int) []string {
	newSlice := []string{}
	for _, val := range slice {
		if len(val) >= minLength {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}

// Creates a map of word counts from a slice of strings
func mapSlice(slice []string) *map[string]int {
	resultMap := make(map[string]int)

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
