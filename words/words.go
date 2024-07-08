package words

import (
	"math/rand"

	"example.com/words-game/fileops"
	"example.com/words-game/textproc"
)

// Word represents a word with its content, length, and usage count.
type Word struct {
	Content    string
	Length     int
	UsageCount int
	Used       bool
}

// NewWord creates a new Word instance.
func NewWord(content string) *Word {
	return &Word{
		Content:    content,
		Length:     len(content),
		UsageCount: 0,
		Used:       false,
	}
}

// Use increments the usage count of the word.
func (w *Word) Use() {
	w.UsageCount++
	w.Used = true
}

// GetWordList processes the file and returns a list of Word instances
func GetWordList(filePath, delimiters string) ([]*Word, error) {
	lines, err := getLines(filePath, delimiters)
	if err != nil {
		return nil, err
	}
	wordCounts := textproc.CreateMapFromSlice(lines)
	wordList := make([]*Word, 0, len(*wordCounts))

	for word := range *wordCounts {
		wordList = append(wordList, NewWord(word))

	}

	return wordList, nil
}

// PickRandomWord picks a random unused word from the list.
func PickRandomWord(wordList []*Word) *Word {
	unusedWords := getUnusedWordList(wordList)
	if len(unusedWords) == 0 {
		return nil // If no unused words are available
	}

	randIndex := rand.Intn(len(unusedWords))
	randomWord := unusedWords[randIndex]
	randomWord.Use()

	return randomWord
}

func getLines(filePath, delimiters string) ([]string, error) {
	lines, err := fileops.ReadLinesFromFile(filePath)
	if err != nil {
		return nil, err
	}
	lines = textproc.SplitByDelimiters(lines, delimiters)
	lines = textproc.FilterSliceByLength(lines, 3)

	return lines, nil
}

func getUnusedWordList(wordList []*Word) []*Word {
	unusedWords := make([]*Word, 0)
	for _, word := range wordList {
		if !word.Used {
			unusedWords = append(unusedWords, word)
		}
	}
	return unusedWords
}
