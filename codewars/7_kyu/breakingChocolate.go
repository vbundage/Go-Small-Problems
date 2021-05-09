// https://www.codewars.com/kata/534ea96ebb17181947000ada

package kata

func BreakChocolate(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	return ((n - 1) * m) + (m - 1)
}
