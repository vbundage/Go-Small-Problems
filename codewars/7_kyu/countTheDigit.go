// https://www.codewars.com/kata/566fc12495810954b1000030

package kata

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	var digitStr string
	compareStrNum := strconv.Itoa(d)
	for k := 0; k <= n; k++ {
		strNum := strconv.Itoa(k * k)
		if strings.Contains(strNum, compareStrNum) {
			digitStr += strNum
		}
	}
	return strings.Count(digitStr, compareStrNum)
}
