package words

import (
	"example.com/words-game/fileops"
	"example.com/words-game/textproc"
)

// Word represents a word with its content and length
type Word struct {
	content string
	length  int
}

// NewWord creates a new Word instance
func NewWord(content string) Word {
	length := len(content)
	return Word{
		content,
		length,
	}
}

// GetWordList processes the file and returns a list of Word instances
func GetWordList(filePath, delimiters string) (*[]Word, error) {
	lines, err := getLines(filePath, delimiters)
	if err != nil {
		return nil, err
	}
	wordsMap := textproc.CreateMapFromSlice(lines)
	wordList := make([]Word, len(*wordsMap))

	idx := 0
	for word := range *wordsMap {
		wordList[idx] = NewWord(word)
		idx++
	}

	return &wordList, nil
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
