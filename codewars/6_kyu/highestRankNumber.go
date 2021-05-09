// https://www.codewars.com/kata/5420fc9bb5b2c7fd57000004

import "fmt"

func HighestRank(nums []int) int {
	numCounts := map[int]int{}
	for _, num := range nums {
		numCounts[num] += 1
	}

	var highestNum int
	for key, count := range numCounts {
		if count > numCounts[highestNum] {
			highestNum = key
		} else if count == numCounts[highestNum] {
			if key > highestNum {
				highestNum = key
			}
		}
	}

	return highestNum
}

func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}) == 12)
	fmt.Println(HighestRank([]int{2}) == 2)
	fmt.Println(HighestRank([]int{2, 1, 1, 5, 3}) == 1)
}