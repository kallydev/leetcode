package solution

func runningSum(nums []int) []int {
	result := make([]int, len(nums))

	for index, number := range nums {
		if index == 0 {
			result[index] = number
		} else {
			result[index] = result[index-1] + number
		}
	}

	return result
}
