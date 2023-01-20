package solution

func findSubsequences(nums []int) [][]int {
	var result [][]int

	depthFirstSearch(nums, []int{}, &result)

	return result
}

func depthFirstSearch(numbers []int, path []int, result *[][]int) {
	length := len(path)

	if length > 1 {
		cache := make([]int, len(path))

		copy(cache, path)

		*result = append(*result, cache)
	}

	visited := make(map[int]bool)

	for index, number := range numbers {
		if length > 0 && number < path[length-1] || visited[number] {
			continue
		}

		visited[number] = true

		depthFirstSearch(numbers[index+1:], append(path, number), result)
	}
}
