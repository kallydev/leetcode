package solution

func firstBadVersion(n int) int {
	left, middle, right := 0, 0, n-1

	for left <= right {
		middle = left + (right-left)/2

		if isBadVersion(middle) {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return false
}
