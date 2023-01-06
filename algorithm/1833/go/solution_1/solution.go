package solution

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	var number, total int

	for _, cost := range costs {
		if total+cost > coins {
			break
		}

		total += cost
		number++
	}

	return number
}
