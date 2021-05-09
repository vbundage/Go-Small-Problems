// https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go

package kata

// import (
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func HighAndLow(in string) string {
// 	strNums := strings.Split(in, " ")
// 	nums := make([]int, len(strNums))
// 	for idx, numChar := range strNums {
// 		nums[idx], _ = strconv.Atoi(numChar)
// 	}

// 	sort.Ints(nums)
// 	return strconv.Itoa(nums[len(nums)-1]) + " " + strconv.Itoa(nums[0])
// }

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	var max, min int
	for idx, val := range strings.Fields(in) {
		num, _ := strconv.Atoi(val)
		if idx == 0 {
			max = num
			min = num
		}

		if max < num {
			max = num
		} else if min > num {
			min = num
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}
