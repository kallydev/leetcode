package solution

func maximumCount(nums []int) int {
	var positive, negative int

	for _, num := range nums {
		switch {
		case num > 0:
			positive++
		case num < 0:
			negative++
		}
	}

	if positive > negative {
		return positive
	} else {
		return negative
	}
}
