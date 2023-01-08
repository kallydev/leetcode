package solution

import "math"

func maxPoints(points [][]int) int {
	length := len(points)

	if length == 1 {
		return length
	}

	result := 2

	for i := 0; i < length; i++ {
		dict := make(map[float64]int)

		for j := 0; j < length; j++ {
			if i != j {
				dict[math.Atan2(float64(points[j][1]-points[i][1]), float64(points[j][0]-points[i][0]))] += 1
			}

			maxValue := 0

			for _, value := range dict {
				if value > maxValue {
					maxValue = value
				}
			}

			result = max(result, maxValue+1)
		}
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
