package solution

import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
	sort.SliceStable(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}
