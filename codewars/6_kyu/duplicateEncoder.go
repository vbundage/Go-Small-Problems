// https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/go

package kata

import "strings"

func DuplicateEncode(word string) string {
	lowerCaseWord := strings.ToLower(word)
	var encoded strings.Builder
	for _, char := range strings.Split(lowerCaseWord, "") {
		if strings.Count(lowerCaseWord, char) > 1 {
			encoded.WriteString(")")
		} else {
			encoded.WriteString("(")
		}
	}
	return encoded.String()
}
