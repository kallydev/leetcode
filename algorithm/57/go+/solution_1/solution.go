package solution

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)

	for index, interval := range intervals {
		if interval[1] < newInterval[0] {
			result = append(result, interval)

			continue
		}

		if interval[0] > newInterval[1] {
			result = append(result, newInterval)
			result = append(result, intervals[index:]...)

			return result
		}

		newInterval[0], newInterval[1] = min(newInterval[0], interval[0]), max(newInterval[1], interval[1])
	}

	return append(result, newInterval)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
