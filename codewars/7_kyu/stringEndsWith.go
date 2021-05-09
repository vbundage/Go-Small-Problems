// https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d

package kata

func solution(str, ending string) bool {
	endLength := len(str) - len(ending)
	if endLength < 0 {
		return false
	}
	return str[endLength:] == ending
}
