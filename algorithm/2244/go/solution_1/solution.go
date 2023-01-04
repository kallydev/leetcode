package solution

func minimumRounds(tasks []int) int {
	taskMap := make(map[int]int)

	for _, task := range tasks {
		taskMap[task]++
	}

	var result int

	for _, value := range taskMap {
		if value == 1 {
			return -1
		}

		if value%3 == 0 {
			result += value / 3
		} else {
			result += value/3 + 1
		}
	}

	return result
}
