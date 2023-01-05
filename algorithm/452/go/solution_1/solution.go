package solution

import "sort"

var _ sort.Interface = (*Slice)(nil)

type Slice [][]int

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func findMinArrowShots(points [][]int) int {
	sort.Sort(Slice(points))

	var (
		result    = 1
		globalEnd = points[0][1]
	)

	for _, point := range points {
		if globalEnd < point[0] {
			result++

			globalEnd = point[1]
		}
	}

	return result
}
