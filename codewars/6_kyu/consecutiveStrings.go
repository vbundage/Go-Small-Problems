// https://www.codewars.com/kata/56a5d994ac971f1ac500003e

package kata

import "strings"

func LongestConsec(strarr []string, k int) string {
	var longestStr string
	for idx := 0; idx+k <= len(strarr); idx += 1 {
		consecutiveStr := strings.Join(strarr[idx:idx+k], "")
		if len(consecutiveStr) > len(longestStr) {
			longestStr = consecutiveStr
		}
	}
	return longestStr
}
