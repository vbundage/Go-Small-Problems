// https://www.codewars.com/kata/5b180e9fedaa564a7000009a/go

package kata

import "strings"

func solve(str string) string {
	const aCharCode rune = 97
	var upcaseCount int
	var lowcaseCount int

	for _, charCode := range str {
		if charCode < aCharCode {
			upcaseCount++
		} else {
			lowcaseCount++
		}
	}

	switch {
	case upcaseCount > lowcaseCount:
		return strings.ToUpper(str)
	case upcaseCount < lowcaseCount:
		fallthrough
	default:
		return strings.ToLower(str)
	}
}
