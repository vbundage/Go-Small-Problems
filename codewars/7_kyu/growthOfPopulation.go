// https://www.codewars.com/kata/563b662a59afc2b5120000c6/train/go

package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
	var count int
	percent = percent / 100
	for p0 < p {
		count++
		p0 += int(float64(p0)*percent) + aug
	}
	return count
}
