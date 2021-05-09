// https://www.codewars.com/kata/5aba780a6a176b029800041c/go

package kata

func MaxMultiple(d, b int) int {
	if b%d == 0 {
		return b
	}

	curr := b - 1
	for curr >= d {
		if curr%d == 0 {
			break
		}
		curr--
	}
	return curr
}
