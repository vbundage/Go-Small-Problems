// https://www.codewars.com/kata/526571aae218b8ee490006f4/go

package kata

import (
	"strconv"
	"strings"
)

func CountBits(f uint) int {
	binary := strconv.FormatUint(uint64(f), 2)
	return strings.Count(binary, "1")
}
