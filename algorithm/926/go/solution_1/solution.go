package solution

func minFlipsMonoIncr(s string) int {
	m := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			m++
		}
	}

	result := m

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			m--
			result = min(result, m)
		} else {
			m++
		}
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
