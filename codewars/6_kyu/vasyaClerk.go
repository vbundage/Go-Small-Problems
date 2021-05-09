// https://www.codewars.com/kata/555615a77ebc7c2c8a0000b8/solutions/go/me/best_practice

package kata

func Tickets(peopleInLine []int) string {
	change := make(map[int]int)
	for _, bill := range peopleInLine {
		switch {
		case bill == 25:
			change[bill] += 1
		case bill == 50:
			if change[25] != 0 {
				change[25] -= 1
				change[bill] += 1
			} else {
				return "NO"
			}
		case bill == 100:
			if change[25] >= 1 && change[50] >= 1 {
				change[25] -= 1
				change[50] -= 1
				change[bill] += 1
			} else if change[25] >= 3 {
				change[25] -= 3
				change[bill] += 1
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}
