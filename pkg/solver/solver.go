package solver

import (
	"bufio"
	"fmt"
	"os"
)

// LoadWords reads words line by line from a file and returns all words that have the
// required letter, only contain valid letters, and are at least of a certain length
func LoadWords(filePath string, requiredLetter rune, validLetters []rune, minWordSize int) ([]string, error) {
	if minWordSize < 1 {
		return nil, fmt.Errorf("Min word size must be at least 1")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	letters := make(map[rune]bool, len(validLetters))
	for _, l := range validLetters {
		letters[l] = false
	}

	words := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()

		reqFound := false
		allValid := true
		reqLength := false
		for l, _ := range letters {
			letters[l] = false
		}

		for i, l := range word {
			if i == minWordSize-1 {
				reqLength = true
			}

			if l == requiredLetter {
				reqFound = true
			}

			if _, ok := letters[l]; !ok {
				allValid = false
				break
			}

			letters[l] = true
		}

		allFound := true
		for _, f := range letters {
			allFound = allFound && f
		}

		if allFound {
			word = fmt.Sprintf("%s (A)", word)
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
