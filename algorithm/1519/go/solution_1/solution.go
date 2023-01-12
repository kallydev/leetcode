package solution

func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := make(map[int][]int, n)

	for _, edge := range edges {
		left, right := edge[0], edge[1]

		graph[left] = append(graph[left], right)
		graph[right] = append(graph[right], left)
	}

	var (
		counts = make([]int, n)
		chars  = make([]rune, 0, len(labels))
	)

	for _, char := range labels {
		chars = append(chars, char)
	}

	depthFirstSearch(0, -1, graph, chars, counts)

	return counts
}

func depthFirstSearch(node, parent int, graph map[int][]int, labels []rune, counts []int) []int {
	nodeCounts := make([]int, 26)
	nodeCounts[labels[node]-'a'] = 1

	for _, child := range graph[node] {
		if child == parent {
			continue
		}

		childCounts := depthFirstSearch(child, node, graph, labels, counts)

		for i := 0; i < 26; i++ {
			nodeCounts[i] += childCounts[i]
		}
	}

	counts[node] = nodeCounts[labels[node]-'a']

	return nodeCounts
}
