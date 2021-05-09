package kata

import (
	"strconv"
	"strings"
)

func DigitalRoot(n int) (sum int) {
	currNum := n
	for currNum >= 10 {
		sum = 0
		for _, digitStr := range strings.Split(strconv.Itoa(currNum), "") {
			num, _ := strconv.Atoi(digitStr)
			sum += num
		}
		currNum = sum
	}
	return currNum
}
