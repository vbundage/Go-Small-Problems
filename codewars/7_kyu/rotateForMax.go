// https://www.codewars.com/kata/56a4872cbb65f3a610000026/go

package kata

import "strconv"

// func MaxNumber(numbers []int64) int64 {
// 	max := numbers[0]
// 	for _, num := range numbers {
// 		if max < num {
// 			max = num
// 		}
// 	}
// 	return max
// }

// func MaxRot(n int64) int64 {
// 	currNum := int(n)
// 	numbers := make([]int64, len(strconv.Itoa(int(n)))-1)
// 	if len(numbers) == 0 {
// 		return n
// 	}

// 	for idx := range numbers {
// 		numbers[idx] = int64(currNum)
// 		if idx == len(numbers)-1 {
// 			break
// 		}

// 		numStr := strconv.Itoa(currNum)
// 		numStr = numStr[:idx] + numStr[idx+1:len(numStr)] + string(numStr[idx])
// 		currNum, _ = strconv.Atoi(numStr)
// 	}

// 	return MaxNumber(numbers)
// }

func MaxRot(n int64) int64 {
	currNum := int(n)
	max := currNum
	for idx := 0; idx < len(strconv.Itoa(int(n)))-1; idx++ {
		numStr := strconv.Itoa(currNum)
		numStr = numStr[:idx] + numStr[idx+1:] + string(numStr[idx])
		currNum, _ = strconv.Atoi(numStr)
		if max < currNum {
			max = currNum
		}
	}

	return int64(max)
}
