package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	oneToMany, manyToOne := make(map[int]int), make(map[int]int)

	for _, people := range trust {
		oneToMany[people[0]] += 1
		manyToOne[people[1]] += 1
	}

	for people, value := range manyToOne {
		if oneToMany[people] == 0 && value == n-1 {
			return people
		}
	}

	return -1
}
