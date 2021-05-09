// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9/train/go

package kata

import "strings"

func FindShort(s string) int {
	shortestLen := len(strings.Split(s, " ")[0])
	for _, word := range strings.Fields(s) {
		if len(word) < shortestLen {
			shortestLen = len(word)
		}
	}
	return shortestLen
}
