package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var delimiters *regexp.Regexp = regexp.MustCompile(`[,:;.\s]+()`)

func main() {
	const filePath string = "words.txt"

	results := writeFileToMap(filePath)

	for _, val := range results {
		fmt.Println(val)
	}

}

func writeFileToMap(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	results := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, delimiters.Split(line, -1)...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	return results
}
