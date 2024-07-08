package words

import (
	"math/rand"
	"strings"
	"unicode"

	"example.com/words-game/fileops"
	"example.com/words-game/textproc"
)

const encryptChar = '*'

// Word represents a word with its content, length, and usage count.
type Word struct {
	Content        string
	CryptedContent string
	Length         int
	UsageCount     int
	Used           bool
}

// NewWord creates a new Word instance.
func NewWord(content string) *Word {
	return &Word{
		Content:        content,
		CryptedContent: encryptWord(content),
		Length:         len(content),
		UsageCount:     0,
		Used:           false,
	}
}

// Use increments the usage count of the word.
func (w *Word) Use() {
	w.UsageCount++
	w.Used = true
}

func (w *Word) ToLowerContent() {
	w.Content = strings.ToLower(w.Content)
}

// CheckLetter checks if the letter is in the Content
func (w *Word) CheckLetterExist(letter string) bool {
	letter = strings.ToLower(letter)
	return strings.Contains(w.Content, letter)
}

// RevealLetter updates the CryptedContent with the given letter if it exists in Content.
func (w *Word) RevealLetter(letter rune) bool {
	letter = unicode.ToLower(letter)

	cryptedRunes := []rune(w.CryptedContent)
	contentRunes := []rune(w.Content)

	for i, char := range contentRunes {
		if char == letter {
			cryptedRunes[i] = letter
		}
	}

	w.CryptedContent = string(cryptedRunes)

	// Check if there are any encrypted characters left
	for _, char := range cryptedRunes {
		if char == encryptChar {
			return false
		}
	}
	return true
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
		newWord := NewWord(word)
		newWord.ToLowerContent()

		wordList = append(wordList, newWord)

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

func encryptWord(word string) string {
	return strings.Repeat(string(encryptChar), len(word))
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
