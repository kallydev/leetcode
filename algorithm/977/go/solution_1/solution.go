package solution

import "sort"

var _ sort.Interface = (*Slice)(nil)

type Slice []int

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Less(i, j int) bool {
	return abs(s[i]) < abs(s[j])
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sortedSquares(nums []int) []int {
	slice := Slice(nums)

	sort.Stable(slice)

	nums = slice

	for index := range nums {
		nums[index] *= nums[index]
	}

	return nums
}
