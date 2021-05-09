// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/go

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	runeCounts := make(map[rune]int)
	for _, charCode := range strings.ToLower(s1) {
		runeCounts[charCode] += 1
	}

	var countOfDups int
	for _, count := range runeCounts {
		if count > 1 {
			countOfDups += 1
		}
	}
	return countOfDups
}

func main() {
	fmt.Println(duplicate_count("") == 0)
	fmt.Println(duplicate_count("abcde") == 0)
	fmt.Println(duplicate_count("abcdea") == 1)
	fmt.Println(duplicate_count("abcdeaB11") == 3)
	fmt.Println(duplicate_count("indivisibility") == 1)
}