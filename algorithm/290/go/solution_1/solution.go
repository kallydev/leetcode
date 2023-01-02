package solution

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, "\x20")

	if len(pattern) != len(words) {
		return false
	}

	var (
		charToWord = make(map[rune]string)
		wordToChar = make(map[string]rune)
	)

	for index, char := range pattern {
		word := words[index]

		internalChar, exists := wordToChar[word]

		switch {
		case !exists:
			wordToChar[word] = char
		case exists && internalChar != char:
			return false
		}

		internalWord, exists := charToWord[char]
		if !exists {
			internalWord = word
			charToWord[char] = internalWord
		}

		if internalWord != word {
			return false
		}
	}

	return true
}
