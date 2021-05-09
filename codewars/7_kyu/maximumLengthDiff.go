// https://www.codewars.com/kata/5663f5305102699bad000056/go

package kata

import "math"

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	var maxLength int
	for _, strA1 := range a1 {
		for _, strA2 := range a2 {
			difference := int(math.Abs(float64(len(strA1) - len(strA2))))
			if maxLength < difference {
				maxLength = difference
			}
		}
	}
	return maxLength
}
