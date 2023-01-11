package solution

func minTime(n int, edges [][]int, hasApple []bool) int {
	adjacency := make(map[int][]int, n)

	for _, edge := range edges {
		left, right := edge[0], edge[1]

		adjacency[left] = append(adjacency[left], right)
		adjacency[right] = append(adjacency[right], left)
	}

	return depthFirstSearch(0, -1, adjacency, hasApple)
}

func depthFirstSearch(node, parent int, adjacency map[int][]int, hasApple []bool) int {
	if _, exists := adjacency[node]; !exists {
		return 0
	}

	var totalTime, childTime int

	for _, child := range adjacency[node] {
		if child == parent {
			continue
		}

		childTime = depthFirstSearch(child, node, adjacency, hasApple)

		if childTime > 0 || hasApple[child] {
			totalTime += childTime + 2
		}
	}

	return totalTime
}
