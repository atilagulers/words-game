package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var delimiters *regexp.Regexp = regexp.MustCompile(`[,:;.\s]+`)

func main() {
	const filePath string = "words.txt"

	writeFileToMap(filePath)

	//for key, _ := range words {
	//	fmt.Println(result)
	//}

}

func writeFileToMap(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result := delimiters.Split(line, -1)
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	return []string{}
}
