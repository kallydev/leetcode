package solution

func pivotIndex(nums []int) int {
	var (
		left, right int
		length      = len(nums)
	)

	for index := range nums {
		if index == length-1 {
			break
		}

		right += nums[length-index-1]
	}

	if left == right {
		return 0
	}

	for index, number := range nums {
		if index == 0 {
			continue
		}

		if index == length-1 {
			right = 0
		} else {
			right -= number
		}

		left += nums[index-1]

		if left == right {
			return index
		}
	}

	return -1
}
