package solution

func search(nums []int, target int) int {
	left, middle, right := 0, 0, len(nums)-1

	for left <= right {
		middle = left + (right-left)/2
		number := nums[middle]

		switch {
		case number == target:
			return middle
		case number > target:
			right = middle - 1
		case number < target:
			left = middle + 1
		}
	}

	return -1
}
