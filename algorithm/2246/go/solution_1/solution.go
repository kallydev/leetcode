package solution

func longestPath(parent []int, s string) int {
	var (
		length   = len(parent)
		children = make(map[int][]int, length)
	)

	for i := 1; i < length; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	longestPath := 1

	depthFirstSearch(0, children, s, &longestPath)

	return longestPath
}

func depthFirstSearch(currentNode int, children map[int][]int, s string, longestPath *int) int {
	longestChain, secondLongestChain := 0, 0

	for _, child := range children[currentNode] {
		longestChainStartingFromChild := depthFirstSearch(child, children, s, longestPath)

		if s[currentNode] == s[child] {
			continue
		}

		switch {
		case longestChainStartingFromChild > longestChain:
			secondLongestChain, longestChain = longestChain, longestChainStartingFromChild
		case longestChainStartingFromChild > secondLongestChain:
			secondLongestChain = longestChainStartingFromChild
		}
	}

	*longestPath = max(*longestPath, longestChain+secondLongestChain+1)

	return longestChain + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
