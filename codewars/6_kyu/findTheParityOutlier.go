// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/go

package kata

func FindOutlier(integers []int) int {
	odds := []int{}
	evens := []int{}
	for _, val := range integers {
		if val%2 == 0 {
			evens = append(evens, val)
		} else {
			odds = append(odds, val)
		}

		if len(evens) >= 2 && len(odds) == 1 {
			return odds[0]
		} else if len(odds) >= 2 && len(evens) == 1 {
			return evens[0]
		}
	}
	return -1
}
