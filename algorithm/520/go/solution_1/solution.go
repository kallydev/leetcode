package solution

func detectCapitalUse(word string) bool {
	var isUppercase, isLowercase, isTitle bool

	for index, character := range word {
		switch {
		case index == 0 && character >= 65 && character <= 90:
			isTitle = true
			isUppercase = true
		case index == 0 && character >= 97 && character <= 122:
			isLowercase = true
		case character >= 65 && character <= 90:
			if index == 1 {
				isTitle = false
			}

			if isTitle || isLowercase {
				return false
			}
		case character >= 97 && character <= 122:
			if isTitle {
				isUppercase = false
			}

			if isUppercase {
				return false
			}
		default:
			return false
		}
	}

	return true
}
