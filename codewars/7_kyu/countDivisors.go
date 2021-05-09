// https://www.codewars.com/kata/542c0f198e077084c0000c2e/go

package kata

func Divisors(n int) int {
	var count int
	for currNum := 1; currNum < n; currNum++ {
		if n%currNum == 0 {
			count++
		}
	}
	return count + 1
}
