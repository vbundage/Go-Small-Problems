// https://www.codewars.com/kata/517abf86da9663f1d2000003/go

package kata

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`[_-]`)
	s = re.ReplaceAllString(s, " ")
	var newStr string
	for idx, word := range strings.Fields(s) {
		if idx == 0 {
			newStr += word
			continue
		}
		newStr += strings.ToUpper(string(word[0])) + word[1:]
	}
	return newStr
}
