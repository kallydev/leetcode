package solution

func xorBeauty(nums []int) int {
	result, length := 0, len(nums)

	for i := 0; i < length; i++ {
		result ^= nums[i]
	}

	return result
}
