// https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go

package kata

import "strings"

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, val := range vowels {
		count += strings.Count(str, val)
	}
	return count
}
