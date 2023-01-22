package solution

import (
	"math"
	"strconv"
)

func alternateDigitSum(n int) int {
	var (
		result int
		length = len(strconv.Itoa(n))
	)

	for index := 0; index < length; index++ {
		step := int(math.Pow10(length - index - 1))
		digit := n / step
		n -= step * digit

		if index%2 == 0 {
			result += digit
		} else {
			result -= digit
		}
	}

	return result
}
