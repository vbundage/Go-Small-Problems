package kata

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	var sum int
	digitStr := strconv.Itoa(n)
	for _, charCode := range digitStr {
		num, _ := strconv.Atoi(string(charCode))
		sum += int(math.Pow(float64(num), float64(p)))
		p += 1
	}

	k := 1
	for k*n < sum {
		k++
	}

	if k*n != sum {
		return -1
	}
	return k
}
