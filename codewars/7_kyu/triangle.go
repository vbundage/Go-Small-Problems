// https://www.codewars.com/kata/56606694ec01347ce800001b
package kata

import "sort"

func IsTriangle(a, b, c int) bool {
	var sides = []int{a, b, c}
	sort.Ints(sides)
	a, b, c = sides[0], sides[1], sides[2]

	if a <= 0 {
		return false
	} else if (a + b) > c {
		return true
	}
	return false
}
