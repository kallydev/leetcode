package solution

func minDeletionSize(strs []string) int {
	var (
		result int
		length = len(strs[0])
	)

	for index := 0; index < length; index++ {
		var left uint8

		for _, str := range strs {
			right := str[index]

			if left > right {
				result++

				break
			}

			left = right
		}
	}

	return result
}
