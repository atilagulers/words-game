package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	const filePath string = "words.txt"
	delimiters := `[,:;.\s()]+`

	results := readLinesFromFile(filePath)

	results = splitByDelimiters(results, delimiters)

	for _, val := range results {
		fmt.Println(val)
	}

}

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

func splitByDelimiters(text []string, delimiters string) []string {
	delimitersRegex := regexp.MustCompile(delimiters)
	textStr := strings.Join(text, " ")
	splitText := delimitersRegex.Split(textStr, -1)

	return splitText
}
