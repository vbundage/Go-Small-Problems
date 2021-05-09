// https://www.codewars.com/kata/55fd2d567d94ac3bc9000064

package kata

func RowSumOddNumbers(n int) int {
	rows := make([][]int, n)
	currNum := 1

	for idx, row := range rows {
		rowNum := idx + 1
		for rowNum > 0 {
			rowNum--
			row = append(row, currNum)
			currNum += 2
		}
		rows[idx] = row
	}

	var sum int
	for _, num := range rows[n-1] {
		sum += num
	}

	return sum
}
