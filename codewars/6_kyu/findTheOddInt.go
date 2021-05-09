// https://www.codewars.com/kata/54da5a58ea159efa38000836/go

package kata

func FindOdd(seq []int) int {
	counts := make(map[int]int)
	for _, value := range seq {
		counts[value]++
	}
	for value, count := range counts {
		if count%2 == 1 {
			return value
		}
	}
	return 0
}
