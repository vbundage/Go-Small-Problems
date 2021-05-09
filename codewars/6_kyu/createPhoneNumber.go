// https://www.codewars.com/kata/525f50e3b73515a6db000b83

package kata

import (
	"fmt"
	"strconv"
	"strings"
)

// func CreatePhoneNumber(numbers [10]uint) string {
// 	digitSlice := make([]string, len(numbers))
// 	for idx, num := range numbers {
// 		digitSlice[idx] = strconv.Itoa(int(num))
// 	}
// 	return fmt.Sprintf("(%s) %s-%s",
// 		strings.Join(digitSlice[:3], ""),
// 		strings.Join(digitSlice[3:6], ""),
// 		strings.Join(digitSlice[6:], ""))
// }

func CreatePhoneNumber(numbers [10]uint) string {
	var digitStr strings.Builder
	for _, num := range numbers {
		digitStr.WriteString(strconv.Itoa(int(num)))
	}
	return fmt.Sprintf("(%s) %s-%s",
		digitStr.String()[:3],
		digitStr.String()[3:6],
		digitStr.String()[6:])
}
