// https://www.codewars.com/kata/585d7d5adb20cf33cb000235/go

package kata

func FindUniq(arr []float32) float32 {
	counts := map[float32]int{}
	for _, val := range arr {
		counts[val] += 1
	}

	for key, count := range counts {
		if count == 1 {
			return key
		}
	}
	return -1
}
