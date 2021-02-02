package solver

import (
	"bufio"
	"os"
)

// LoadWords reads words line by line from a file and returns all words that have the
// required letter, only contain valid letters, and are at least of a certain length
func LoadWords(filePath string, requiredLetter rune, validLetters []rune, minWordSize int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	letters := make(map[rune]struct{}, len(validLetters))
	for _, l := range validLetters {
		letters[l] = struct{}{}
	}

	words := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()

		reqFound := false
		allValid := true
		reqLength := false

		for i, l := range word {
			if i == minWordSize-1 {
				reqLength = true
			}

			if l == requiredLetter {
				reqFound = true
				continue
			}

			if _, ok := letters[l]; !ok {
				allValid = false
				break
			}
		}

		if reqLength && reqFound && allValid {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
