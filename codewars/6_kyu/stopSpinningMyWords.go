// https://www.codewars.com/kata/5264d2b162488dc400000001/go

package kata

import (
	"fmt"
	"strings"
)

func reverseStr(str string) string {
	reversedStr := ""
	for _, char := range strings.Split(str, "") {
		fmt.Println([]rune(char))
		reversedStr = char + reversedStr
	}
	return reversedStr
}

func SpinWords(str string) string {
	strSlice := strings.Fields(str)
	for idx, word := range strSlice {
		if len(word) >= 5 {
			strSlice[idx] = reverseStr(word)
		}
	}
	return strings.Join(strSlice, " ")
}
