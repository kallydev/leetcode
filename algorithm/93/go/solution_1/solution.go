package solution

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	var result []string

	depthFirstSearch(&result, []string{}, s)

	return result
}

func depthFirstSearch(result *[]string, path []string, s string) {
	if len(path) == 4 {
		if len(s) == 0 {
			*result = append(*result, strings.Join(path, "."))
		}

		return
	}

	value := 0

	for i := 0; i < len(s) && i < 3; i++ {
		number := int(s[i] - '0')

		value *= 10
		value += number

		if value <= 255 {
			newPath := make([]string, len(path)+1)
			copy(newPath, path)
			newPath[len(path)] = s[:i+1]

			depthFirstSearch(result, newPath, s[i+1:])
		}

		if s[0] == '0' {
			break
		}
	}
}
