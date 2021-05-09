package kata

/*
Input: string of digits, integer size of chunks
Output: string with chunks reversed or rotated
requirements:
  - create chunks out of the input string of size n
    - if n <= 0 then return an empty string
    - if n > size of string return empty string
  - check each chunk
    - if the sum of the cubes of each digit % 2 is equal to 0
      - reverse the chunk
     - else
       - rotate the chunk to the left by one position
         - rotating is moving the first digit to the end of the string

Examples/Test Cases:
"1234", 2
"12", "34"
12 => 1 + 8 = 9 % 2 = 1
34 => 27 + 64 = 91 % 2 = 1
"21", "43"
"2143"

"2352", 0 -> ""
"", 5 -> ""
"23", 3 -> ""

Data Structure:
Slice
String
Integer

Algorithm:
  revrot
  - initialize a slice of chunks
  - initialize a result string
  - iterate until input string is empty
    - append chunks with a slice of string from 0 to n
  - iterate over the chunks
    - sum the digits of the chunk
    - if the sum % 2 == 0
      - reverse chunk
    - else
      - rotate the chunk 1 position to the left
        - move the first digit of the chunk to the end of the chunk
    - concatenate chunk to the result
  - return result

  sumDigits
    - initailize slice of digits
    - initialize sum to 0
    - for each digit in digits
      - add digit to the sum
    - return sum

  rotate
    - return slice of string from 1:len
      - concatenate idx 0 to 1:len slice

  reverseStr
    - initialize slice of digits
    - initialize result string
    - for each digit in slice
      - prepend result with digit
    - return result
*/

import (
	"math"
	"strconv"
	"strings"
)

func createChunks(s string, n int) []string {
	chunks := make([]string, 0)
	for s != "" {
		if n > len(s) {
			break
		}
		chunks = append(chunks, s[:n])
		s = s[n:]
	}

	return chunks
}

func sumDigits(digitStr string) int {
	digits := strings.Split(digitStr, "")
	var sum int
	for _, digit := range digits {
		num, _ := strconv.Atoi(digit)
		num = int(math.Pow(float64(num), 3))
		sum += num
	}
	return sum
}

func reverseDigits(str string) string {
	var result string
	for _, digit := range strings.Split(str, "") {
		result = digit + result
	}
	return result
}

func Revrot(s string, n int) string {
	if len(s) < n {
		return ""
	} else if n <= 0 {
		return ""
	}

	chunks := createChunks(s, n)
	for idx, chunk := range chunks {
		if sumDigits(chunk)%2 == 0 {
			chunks[idx] = reverseStr(chunk)
		} else {
			chunks[idx] = chunk[1:] + string(chunk[0])
		}
	}

	return strings.Join(chunks, "")
}
