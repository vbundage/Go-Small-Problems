// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9/go

package kata

import "strings"

// func Accum(s string) (result string) {
// 	for idx, char := range s {
// 		letter := string(char)
// 		capitalLetter := strings.ToUpper(letter)
// 		lowerLetter := strings.ToLower(letter)
// 		result += capitalLetter + strings.Repeat(lowerLetter, idx)

// 		if idx != len(s)-1 {
// 			result += "-"
// 		}
// 	}
// 	return
// }

func Accum(s string) string {
	parts := make([]string, len(s))

	for idx, char := range s {
		letter := string(char)
		capitalLetter := strings.ToUpper(letter)
		lowerLetter := strings.ToLower(letter)
		parts[idx] = capitalLetter + strings.Repeat(lowerLetter, idx)
	}
	return strings.Join(parts, "-")
}
