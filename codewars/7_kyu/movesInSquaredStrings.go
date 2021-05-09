// https://www.codewars.com/kata/56dbe0e313c2f63be4000b25/go

package kata

import (
	"strings"
)

func reverseCharacters(s string) (result string) {
	for _, char := range strings.Split(s, "") {
		result = char + result
	}
	return
}

func reverseSubstrings(s []string) []string {
	result := make([]string, len(s))
	for idx := range s {
		result[idx] = s[len(s)-idx-1]
	}
	return result
}

func VertMirror(s string) string {
	substrings := strings.Split(s, "\n")
	for idx, substring := range substrings {
		substrings[idx] = reverseCharacters(substring)
	}
	return strings.Join(substrings, "\n")
}

func HorMirror(s string) string {
	substrings := strings.Split(s, "\n")
	return strings.Join(reverseSubstrings(substrings), "\n")
}

type fParam func(string) string

func Oper(f fParam, x string) string {
	return f(x)
}
