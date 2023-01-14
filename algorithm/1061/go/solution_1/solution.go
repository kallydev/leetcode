package solution

import "strings"

var representative [26]int

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	for i := 0; i < 26; i++ {
		representative[i] = i
	}

	for i := 0; i < len(s1); i++ {
		performUnion(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	var result strings.Builder

	for _, char := range baseStr {
		result.WriteRune(rune(find(int(char-'a')) + 'a'))
	}

	return result.String()
}

func find(x int) int {
	if representative[x] == x {
		return x
	}

	representative[x] = find(representative[x])

	return representative[x]
}

func performUnion(x, y int) {
	x, y = find(x), find(y)

	if x == y {
		return
	}

	if x < y {
		representative[y] = x
	} else {
		representative[x] = y
	}
}
