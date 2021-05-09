// https://www.codewars.com/kata/59cfc000aeb2844d16000075/go

package kata

import "strings"

func Capitalize(st string) []string {
	result := make([]string, 2)
	for idx, charCode := range st {
		if idx%2 == 0 {
			result[0] += strings.ToUpper(string(charCode))
			result[1] += string(charCode)
		} else {
			result[0] += string(charCode)
			result[1] += strings.ToUpper(string(charCode))
		}
	}
	return result
}
