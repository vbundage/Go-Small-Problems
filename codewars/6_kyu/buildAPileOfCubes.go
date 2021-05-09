// https://www.codewars.com/kata/5592e3bd57b64d00f3000047/go

package kata

import "math"

func FindNb(m int) int {
	count := 1
	total := 0
	for total < m {
		total += int(math.Pow(float64(count), 3))
		count += 1
	}

	if total == m {
		return count - 1
	}
	return -1
}
