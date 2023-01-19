package solution

func subarraysDivByK(nums []int, k int) int {
	prefixMod, result := 0, 0

	modGroups := make([]int, k)
	modGroups[0] = 1

	for _, number := range nums {
		prefixMod = (prefixMod + number%k + k) % k

		result += modGroups[prefixMod]
		modGroups[prefixMod]++
	}

	return result
}
