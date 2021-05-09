// https://www.codewars.com/kata/56b7f2f3f18876033f000307/go

package kata

func InAscOrder(numbers []int) bool {
	for idx, value := range numbers {
		if idx == 0 {
			continue
		}

		if numbers[idx-1] > value {
			return false
		}
	}
	return true
}
