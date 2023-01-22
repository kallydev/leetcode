package solution

func getCommon(nums1 []int, nums2 []int) int {
	cache := make(map[int]struct{})

	for _, number := range nums1 {
		cache[number] = struct{}{}
	}

	for _, number := range nums2 {
		if _, exists := cache[number]; exists {
			return number
		}
	}

	return -1
}
