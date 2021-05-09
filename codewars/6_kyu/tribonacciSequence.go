// https://www.codewars.com/kata/556deca17c58da83c00002db/go

package kata

func sum(numbers []float64) (sum float64) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Tribonacci(signature [3]float64, n int) []float64 {
	if n <= 3 {
		return signature[:n]
	}

	sequence := signature[:]
	for idx := 3; idx < n; idx += 1 {
		sequence = append(sequence, sum(sequence[idx-3:idx]))
	}
	return sequence[:n]
}
