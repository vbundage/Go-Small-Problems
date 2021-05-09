// https://www.codewars.com/kata/583f158ea20cfcbeb400000a/go

package kata

// func Arithmetic(a int, b int, operator string) int {
//   ops := map[string]func(int, int) int {
//     "add": func(a, b int) int {
//       return a + b
//     },
//     "subtract": func(a, b int) int {
//       return a - b
//     },
//     "multiply": func(a, b int) int {
//       return a * b
//     },
//     "divide": func(a, b int) int {
//       return a / b
//     },
//   }
//   return ops[operator](a, b)
// }

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	}
	return 0
}
