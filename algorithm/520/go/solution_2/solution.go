package solution

import "strings"

func detectCapitalUse(word string) bool {
	return word == strings.ToLower(word) || word == strings.ToUpper(word) || word == strings.ToUpper(word[:1])+strings.ToLower(word[1:])
}
