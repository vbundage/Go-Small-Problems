// https://www.codewars.com/kata/56541980fa08ab47a0000040/train/go

package kata

import "fmt"

func PrinterError(s string) string {
	const mCharCode = 109
	var errorCount int
	for _, charCode := range s {
		if charCode > mCharCode {
			errorCount++
		}
	}
	return fmt.Sprintf("%d/%d", errorCount, len(s))
}
